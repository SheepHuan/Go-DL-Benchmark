package main

import (
	"fmt"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_physical_devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
)

func main() {

	// 注册设备
	devDesc := protos.PhysicalDeviceDescription{
		SerialNumber:    "M8NRKD050090333",
		DeviceName:      "ROG 幻162021",
		OSType:          protos.DeviceOSType_windows,
		ArchType:        protos.ArchitectureType_x86_64,
		ComputableChips: []protos.ComputableChipType{protos.ComputableChipType_cpu},
		DeviceAddress:   &protos.PhysicalDeviceDescription_PcAddr{PcAddr: &protos.PCDeviceAddress{DeviceIp: "127.0.0.1", DevicePort: 10001}},
	}

	g_physical_devices.RegisterDevicesSelf([]*protos.PhysicalDeviceDescription{&devDesc})

	//ip, port := "127.0.0.1", 10001
	//go terminal.LaunchRemoteTerminalService(ip, port)
	//config := &protos.ModelBenchmarkTestArgs{
	//	ModelPath:        "uploads/file/d45044fc495d9e031315dc19e20300a7_20230311194541.onnx",
	//	Framework:        protos.FrameworkType_onnxruntime,
	//	InputTensorShape: "1,3,512,512",
	//	InputTensorType:  protos.TensorDataType_float32,
	//	WarmupRounds:     1,
	//	RunRounds:        1,
	//}
	//d, _ := proto.Marshal(config)
	//s := utils.Pb2Base64(d)
	//fmt.Println(s)
	//
	////ability := model.ModelBenchmarkTestAbility{
	////	IsSupportModelBenchmarkTest:           true,
	////	SupportedFrameworksForRuntimeAnalysis: []protos.FrameworkType{protos.FrameworkType_onnxruntime},
	////	SupportedFrameworksForStaticAnalysis:  []protos.FrameworkType{protos.FrameworkType_onnxruntime},
	////}
	////
	//
	////time.Sleep(2 * 1e9)
	////res, err := ability.ModelBenchmarkTest(config, device)
	////if err == nil {
	////	fmt.Println(res)
	////}
	////
	var ss string
	_, _ = fmt.Scanln(&ss)

}
