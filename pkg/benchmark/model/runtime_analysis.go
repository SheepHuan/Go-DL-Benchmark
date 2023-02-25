package model

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-dl-benchmark/pkg/devices"
)

type OperateRuntimeInfo struct {
	Name    string  `json:"Name"`
	Latency float64 `json:"Latency"`
	Memory  uint64  `json:"Memory"`
	Power   float64 `json:"Power"`
}

type SingleRoundProfileResult struct {
	// Total
	RoundIdenx   uint64  `json:"RoundIndex"`
	TotalLatency float64 `json:"TotalLatency"`
	PeakMemory   uint64  `json:"PeakMemory"`
	AvgPower     float64 `json:"AvgPower"`

	OperateInfoList []OperateRuntimeInfo `json:"OperateRuntimeInfo"`
}

type RuntimeAnalysisResult struct {
	// average
	AvgTotalLatency float64 `json:"AvgTotalLatency"`
	AvgPeakMemory   float64 `json:"AvgPeakMemory"`
	AvgPeakPower    float64 `json:"AvgPeakPower"`

	InitTime                 float64                    `json:"InitTime"`
	InitMemory               uint64                     `json:"InitMemory"`
	MultiRoundsProfileResult []SingleRoundProfileResult `json:"MultiRoundsProfileResult"`
}

// v1 版本不考虑逐层转发，仅考虑本地执行
// TODO 未来加入设备转发功能。
func nnMeterRuntimeAnalyse(config ModelBenchmarkTestConfig, device devices.HardwareDevice) RuntimeAnalysisResult {
	cmd := fmt.Sprintf("nn-meter predict --predictor %s --predictor-version 1.0", config.NNMeterPredictor)
	if config.NNMeterInferenceFramework == TFlite {
		cmd = fmt.Sprintf("%s --tensorflow %s", cmd, config.ModelPath)
	} else if config.NNMeterInferenceFramework == Onnxruntime {
		cmd = fmt.Sprintf("%s --onnx %s", cmd, config.ModelPath)
	} else {
		return RuntimeAnalysisResult{}
	}
	if device.IsAlive() {
		stdout, stderr, err := device.ExecuteCommand("conda activate tf26")
		if err == nil {
			fmt.Println(fmt.Sprintf("out:%s\nerr:%s\n", stdout, stderr))
		} else {
			log.Error(err)
			return RuntimeAnalysisResult{}
		}
		log.Info(cmd)
		stdout, stderr, err = device.ExecuteCommand(cmd)
		fmt.Println(fmt.Sprintf("out:%s\nerr:%s\n", stdout, stderr))
		if err != nil {
			log.Error(err)
			return RuntimeAnalysisResult{}
		}

	} else {
		log.Error(fmt.Sprintf("Device:%s:%d is not alive!", device.Description.Ip, device.Description.Port))
	}

	return RuntimeAnalysisResult{}
}

func onnxruntimeRuntimeAnalyse(config ModelBenchmarkTestConfig, device devices.HardwareDevice) RuntimeAnalysisResult {
	script := fmt.Sprintf("conda activate ModelProfiler\npython tools/onnxruntime/onnx_runtime_for_pc.py "+
		"--model_path=%s "+
		"--input_tensor_shape=%s "+
		"--input_tensor_type=%s "+
		"--device=%s "+
		"--rounds=%d",
		config.ModelPath,
		config.InputShape,
		config.InputTensorType,
		device.Description.ComputableChip,
		config.RunRounds,
	)

	if device.IsAlive() {
		result := RuntimeAnalysisResult{}
		stdout, stderr, err := device.ExecuteCommand(script)
		//fmt.Println(fmt.Sprintf("out:%s\nerr:%s\n", stdout, stderr))
		if err != nil {
			log.Error(stderr)
			log.Error(err)
			return RuntimeAnalysisResult{}
		}
		err = json.Unmarshal([]byte(stdout), &result)
		if err != nil {
			log.Error(err)
			return result
		}
		return result

	} else {
		log.Error(fmt.Sprintf("Device:%s:%d is not alive!", device.Description.Ip, device.Description.Port))
	}

	return RuntimeAnalysisResult{}
}
