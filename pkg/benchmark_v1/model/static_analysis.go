package model

import (
	"encoding/json"
	"fmt"
	"github.com/sheephuan/go-dl-benchmark/pkg/devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	log "github.com/sirupsen/logrus"
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

func OnnxruntimeStaticAnalyse(config *protos.ModelBenchmarkTestArgs, device devices.HardwareDevice) *protos.ModelStaticAnalysisResult {
	script := fmt.Sprintf("conda activate ModelProfiler\n"+
		"cd tools/ModelProfileTool \n"+
		"python -m profile.onnxruntime.onnx_static "+
		"--model_path=%s "+
		"--shape=%s "+
		"--type=%d \n", config.GetModelPath(), config.GetInputTensorShape(), config.GetInputTensorType())
	result := protos.ModelStaticAnalysisResult{}
	stdout, stderr, err := device.ExecuteCommand(script)

	if err != nil {
		log.Error(stderr)
		log.Error(err)
		return &result
	}
	//err = proto.Unmarshal([]byte(stdout), &result)
	err = json.Unmarshal([]byte(stdout), &result)
	if err != nil {
		log.Error(err)
		return &result
	}
	return &result
}
