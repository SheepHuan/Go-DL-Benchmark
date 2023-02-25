package model

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-dl-benchmark/pkg/devices"
)

// https://github.com/ThanatosShinji/onnx-tool/issues/16
type OperateStaticInfo struct {
	Name          string  `json:"Name"`
	MACs          uint64  `json:"MACs"`
	MacsPercent   float64 `json:"MACsPercent"` // 单位 百分号
	Memory        uint64  `json:"Memory"`      // 单位
	MemPercent    float64 `json:"MemPercent"`
	Params        uint64  `json:"Params"`
	ParamsPercent float64 `json:"ParamsPercent"`
	InShape       string  `json:"InShape"`
	OutShape      string  `json:"OutShape"`
}

type StaticAnalysisResult struct {
	Framework string `json:"Framework"`
	//From Tensor
	Memory uint64 `json:"Memory"` //计算量,每秒发生百万次计算
	Params uint64 `json:"Params"` //参数量,网络存在MB的参数大小
	MACs   uint64 `json:"MACs"`
	//From Operate
	OnnxOperateInfoList []OperateStaticInfo `json:"OperateStaticInfo"`
}

func OnnxruntimeStaticAnalyse(config ModelBenchmarkTestConfig, device devices.HardwareDevice) StaticAnalysisResult {
	script := fmt.Sprintf("conda activate ModelProfiler\n"+
		"python tools/onnxruntime/onnx_static.py "+
		"--model_path=%s "+
		"--shape=%s "+
		"--type=%s \n", config.ModelPath, config.InputShape, config.InputTensorType)
	result := StaticAnalysisResult{}
	stdout, stderr, err := device.ExecuteCommand(script)

	if err != nil {
		log.Error(stderr)
		log.Error(err)
		return result
	}
	err = json.Unmarshal([]byte(stdout), &result)
	if err != nil {
		return result
	}
	return result
}

func paddleliteStaticAnalyse() {

}

func tfliteStaticAnalyse() {

}
