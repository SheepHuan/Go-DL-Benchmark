package model

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sheephuan/go-dl-benchmark/pkg/devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	log "github.com/sirupsen/logrus"
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

func onnxruntimeRuntimeAnalyse(config *protos.ModelBenchmarkTestArgs, device devices.HardwareDevice) *protos.ModelRuntimeAnalysisResult {
	script := fmt.Sprintf("conda activate ModelProfiler\n"+
		"cd  tools/ModelProfileTool/\n"+
		"python -m profile.onnxruntime.onnx_runtime_for_pc "+
		"--model_path=%s "+
		"--input_tensor_shape=%s "+
		"--input_tensor_type=%d "+
		"--device=%s "+
		"--rounds=%d",
		config.GetModelPath(),
		config.GetInputTensorShape(),
		config.GetInputTensorType(),
		device.Description.ComputableChip,
		config.GetRunRounds(),
	)

	if device.IsAlive() {
		result := protos.ModelRuntimeAnalysisResult{}
		stdout, stderr, err := device.ExecuteCommand(script)
		if err != nil {
			log.Error(stderr)
			log.Error(err)
			return &result
		}
		dec_str, err := base64.StdEncoding.DecodeString(stdout)
		if err != nil {
			log.Error(err)
			return &result
		}
		err = proto.Unmarshal(dec_str, &result)
		if err != nil {
			log.Error(err)
			return &result
		}
		return &result

	} else {
		log.Error(fmt.Sprintf("Device:%s:%d is not alive!", device.Description.Ip, device.Description.Port))
	}

	return &protos.ModelRuntimeAnalysisResult{}
}
