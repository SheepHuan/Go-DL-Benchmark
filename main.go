package main

import (
	"fmt"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_benchmark_test/model"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_physical_devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	"github.com/sheephuan/go-dl-benchmark/pkg/terminal"
	"time"
)

func main() {
	ip, port := "127.0.0.1", 10001
	config := &protos.ModelBenchmarkTestArgs{
		ModelPath:        "D:\\code\\Go-DL-Benchmark\\res\\resnet18-12.onnx",
		Framework:        protos.FrameworkType_onnxruntime,
		InputTensorShape: "1,3,512,512",
		InputTensorType:  protos.TensorDataType_float32,
		WarmupRounds:     1,
		RunRounds:        1,
	}

	//device := g_physical_devices.HardwareDevice{
	//	Description: g_physical_devices.HardwareDescription{
	//		DeviceName:     "ROG",
	//		OSType:         g_physical_devices.Windows,
	//		Architecture:   g_physical_devices.X86_64,
	//		ComputableChip: g_physical_devices.Cpu,
	//		Ip:             ip,
	//		Port:           port,
	//	},
	//}

	device := &g_physical_devices.PhysicalDeviceClient{
		DeviceDescription: protos.PhysicalDeviceDescription{
			DeviceName:      "ROG",
			OSType:          protos.DeviceOSType_windows,
			ArchType:        protos.ArchitectureType_x86_64,
			ComputableChips: []protos.ComputableChipType{protos.ComputableChipType_cpu},
			DeviceAddress: &protos.PhysicalDeviceDescription_PcAddr{PcAddr: &protos.PCDeviceAddress{
				DeviceIp:   ip,
				DevicePort: int32(port),
			}},
		},
	}

	ability := model.ModelBenchmarkTestAbility{
		IsSupportModelBenchmarkTest:           true,
		SupportedFrameworksForRuntimeAnalysis: []protos.FrameworkType{protos.FrameworkType_onnxruntime},
		SupportedFrameworksForStaticAnalysis:  []protos.FrameworkType{protos.FrameworkType_onnxruntime},
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
