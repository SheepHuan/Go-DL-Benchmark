package main

import (
	"encoding/json"
	"fmt"
	"go-dl-benchmark/pkg/benchmark/model"
	"go-dl-benchmark/pkg/devices"
	"go-dl-benchmark/pkg/terminal"
	"time"
)

func main() {
	//

	ip, port := "127.0.0.1", 10001
	config := model.ModelBenchmarkTestConfig{
		ModelPath:                 "D:\\code\\Go-DL-Benchmark\\res\\mobilenet_quant_v1_224.tflite",
		Framework:                 model.NNMeter,
		WarmupRounds:              5,
		RunRounds:                 10,
		RunInterval:               0,
		NNMeterPredictor:          "cortexA76cpu_tflite21",
		NNMeterInferenceFramework: model.TFlite,
		NNMeterPredictorVersion:   1.0,
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

	go terminal.LaunchRemoteTerminalService(ip, port)
	time.Sleep(2 * 1e9)
	bytes, err := json.Marshal(config)
	if err == nil {
		model.ModelBenchmarkTest(bytes, device)
	}
	var s string
	_, _ = fmt.Scanln(&s)

}
