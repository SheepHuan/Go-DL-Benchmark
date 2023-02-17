package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"go-dl-benchmark/pkg/devices"
)

const (
	PaddleLite  = "PaddleLite"
	TFlite      = "TFLite"
	Onnxruntime = "onnxruntime"
	NNMeter     = "nnMeter"
)

type ModelBenchmarkTestConfig struct {
	ModelPath string `json:"ModelPath"`
	Framework string `json:"Framework"`

	// runtime setting
	WarmupRounds int     `json:"WarmupRounds"`
	RunRounds    int     `json:"RunRounds"`
	RunInterval  float64 `json:"RunInterval"`
	// limitaion
	PeakMemory        float64 `json:"PeakMemory"`
	PeakPower         float64 `json:"PeakPower"`
	CpuPeakFrequency  uint64  `json:"CpuPeakFrequency"`
	CpuMaxBigCores    int     `json:"CpuMaxBigCores"`
	CpuMaxLittleCores int     `json:"CpuMaxLittleCores"`
	GpuPeakFrequency  uint64  `json:"GpuPeakFrequency"`
	PeakBandwidth     float64 `json:"PeakBandwidth"`

	//nn-Meter predictor
	NNMeterPredictor          string  `json:"NNMeterPredictor"`
	NNMeterPredictorVersion   float32 `json:"NNMeterPredictorVersion"`
	NNMeterInferenceFramework string  `json:"NNMeterInferenceFramework"`
	//
}

type ModelBenchmarkTestResult struct {
	TestConfig    ModelBenchmarkTestConfig `json:"ModelBenchmarkTestConfig"`
	StaticResult  StaticAnalysisResult     `json:"StaticAnalysisResult"`
	RuntimeResult RuntimeAnalysisResult    `json:"RuntimeAnalysisResult"`
}

func ModelBenchmarkTest(configBytes []byte, device devices.HardwareDevice) (error, ModelBenchmarkTestResult) {
	_config := ModelBenchmarkTestConfig{}
	err := json.Unmarshal(configBytes, &_config)

	if err != nil {
		log.Error(err.Error())
		return nil, ModelBenchmarkTestResult{}
	}
	device.Connect()
	//TODO 识别模型地址是URL还是一个本地路径,若是URL则先下载模型
	//TODO 首先展开静态分析
	//展开动态分析
	//log.Info(string(configBytes))
	if _config.Framework == NNMeter {
		nnMeterRuntimeAnalyse(_config, device)
	}
	device.Disconnect()
	return nil, ModelBenchmarkTestResult{}
}
