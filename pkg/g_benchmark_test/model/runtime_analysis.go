package model

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_physical_devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	log "github.com/sirupsen/logrus"
)

func onnxruntimeRuntimeAnalyse(config *protos.ModelBenchmarkTestArgs, device *g_physical_devices.PhysicalDeviceClient) *protos.ModelRuntimeAnalysisResult {
	script := fmt.Sprintf("conda activate ModelProfiler\n"+
		"cd  %s\n"+
		"python -m profile.onnxruntime.onnx_runtime_for_pc "+
		"--model_path=%s "+
		"--input_tensor_shape=%s "+
		"--input_tensor_type=%d "+
		"--device=%d "+
		"--rounds=%d",
		tools_path,
		config.GetModelPath(),
		config.GetInputTensorShape(),
		config.GetInputTensorType(),
		0,
		config.GetRunRounds(),
	)
	//fmt.Println(script)
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
		//log.Error(fmt.Sprintf("Device:%s:%d is not alive!", device.DeviceDescription.GetPcAddr(), device.Description.Port))
	}

	return &protos.ModelRuntimeAnalysisResult{}
}
