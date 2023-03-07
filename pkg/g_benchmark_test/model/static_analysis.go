package model

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_physical_devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	log "github.com/sirupsen/logrus"
)

// https://github.com/ThanatosShinji/onnx-tool/issues/16
func OnnxruntimeStaticAnalyse(config *protos.ModelBenchmarkTestArgs, device *g_physical_devices.PhysicalDeviceClient) *protos.ModelStaticAnalysisResult {
	script := fmt.Sprintf("conda activate ModelProfiler\n"+
		"cd D:\\code\\Go-DL-Benchmark\\tools\\ModelProfileTool \n"+
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
	if stderr != "" {
		log.Error(stderr)
		return &result
	}
	//err = json.Unmarshal([]byte(stdout), &result)
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
}
