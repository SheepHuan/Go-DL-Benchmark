package model

const (
	PaddleLite  = "paddleLite"
	TFlite      = "tfLite"
	Onnxruntime = "onnxruntime"
	NNMeter     = "nnMeter"
)

type ModelBenchmarkTestConfig struct {
	ModelPath string `json:"ModelPath"`
	Framework string `json:"Framework"`

	// input setting
	InputShape      string `json:"InputShape"`
	InputTensorType string `json:"InputTensorType"`

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
