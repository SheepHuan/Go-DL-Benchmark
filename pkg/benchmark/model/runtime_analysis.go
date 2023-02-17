package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-dl-benchmark/pkg/devices"
)

type OperateProfileResult struct {
	Latency float64 `json:"Latency"`
	Memory  float64 `json:"Memoey"`
	Power   float64 `json:"Power"`
}

type RuntimeAnalysisResult struct {
	// total
	Latency float64 `json:"Latency"`
	Memory  float64 `json:"Memoey"`
	Power   float64 `json:"Power"`

	OperateInfoList []OperateProfileResult `json:"OperateInfo"`
}

// v1 版本不考虑逐层转发，仅考虑本地执行
// TODO 未来加入设备转发功能。
func nnMeterRuntimeAnalyse(config ModelBenchmarkTestConfig, device devices.HardwareDevice) (error, RuntimeAnalysisResult) {
	cmd := fmt.Sprintf("nn-meter predict --predictor %s --predictor-version 1.0", config.NNMeterPredictor)
	if config.NNMeterInferenceFramework == TFlite {
		cmd = fmt.Sprintf("%s --tensorflow %s", cmd, config.ModelPath)
	} else if config.NNMeterInferenceFramework == Onnxruntime {
		cmd = fmt.Sprintf("%s --onnx %s", cmd, config.ModelPath)
	} else {
		return nil, RuntimeAnalysisResult{}
	}
	if device.IsAlive() {
		stdout, stderr, err := device.ExecuteCommand("conda activate tf26")
		if err == nil {
			fmt.Println(fmt.Sprintf("out:%s\nerr:%s\n", stdout, stderr))
		} else {
			log.Error(err)
			return err, RuntimeAnalysisResult{}
		}
		log.Info(cmd)
		stdout, stderr, err = device.ExecuteCommand(cmd)
		fmt.Println(fmt.Sprintf("out:%s\nerr:%s\n", stdout, stderr))
		if err != nil {
			log.Error(err)
			return err, RuntimeAnalysisResult{}
		}

	} else {
		log.Error(fmt.Sprintf("Device:%s:%d is not alive!", device.Description.Ip, device.Description.Port))
	}

	return nil, RuntimeAnalysisResult{}
}
