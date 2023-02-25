package main

import (
	"fmt"
	"go-dl-benchmark/pkg/benchmark/model"
	"go-dl-benchmark/pkg/devices"
	"go-dl-benchmark/pkg/terminal"
	"time"
)

func main() {
	ip, port := "127.0.0.1", 10001
	config := model.ModelBenchmarkTestConfig{
		ModelPath:       "D:\\code\\Go-DL-Benchmark\\res\\resnet18-12.onnx",
		Framework:       model.Onnxruntime,
		InputShape:      "1,3,512,512",
		InputTensorType: "float32",
		WarmupRounds:    5,
		RunRounds:       10,
		RunInterval:     0,
		//NNMeterPredictor:          "cortexA76cpu_tflite21",
		//NNMeterInferenceFramework: model.TFlite,
		//NNMeterPredictorVersion:   1.0,
	}

	device := devices.HardwareDevice{
		Description: devices.HardwareDescription{
			DeviceName:     "ROG",
			OSType:         devices.Windows,
			Architecture:   devices.X86_64,
			ComputableChip: devices.Cpu,
			Ip:             ip,
			Port:           port,
		},
	}

	ability := model.BenchmarkAbility{
		IsSupportModelBenchmarkTest:           true,
		IsSupportHardwareBenchmarkTest:        false,
		SupportedFrameworksForRuntimeAnalysis: []string{model.Onnxruntime},
		SupportedFrameworksForStaticAnalysis:  []string{model.Onnxruntime},
	}

	go terminal.LaunchRemoteTerminalService(ip, port)
	time.Sleep(2 * 1e9)
	res, err := ability.ModelBenchmarkTest(config, device)
	if err == nil {
		fmt.Println(res)
	}

	var s string
	_, _ = fmt.Scanln(&s)

}
