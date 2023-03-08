package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	"github.com/sheephuan/go-dl-benchmark/pkg/utils"
)

func main() {

	//ip, port := "127.0.0.1", 10001
	//go terminal.LaunchRemoteTerminalService(ip, port)
	//devDesc := protos.PhysicalDeviceDescription{}
	//if jsonStr, err := os.ReadFile("config.json"); err == nil {
	//	err = proto.Unmarshal(jsonStr, &devDesc)
	//	//fmt.Println(jsonStr)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//}
	//
	//err := g_physical_devices.RegisterDevicesSelf([]*protos.PhysicalDeviceDescription{&devDesc})
	//if err != nil {
	//	return
	//}

	devDesc := protos.PhysicalDeviceDescription{
		SerialNumber:  "M8NRKD050090333",
		DeviceName:    "ROG 幻162021",
		OSType:        protos.DeviceOSType_windows,
		ArchType:      protos.ArchitectureType_x86_64,
		DeviceAddress: &protos.PhysicalDeviceDescription_PcAddr{PcAddr: &protos.PCDeviceAddress{DeviceIp: "127.0.0.1", DevicePort: 10001}},
	}

	b, err := proto.Marshal(&devDesc)
	if err != nil {
		return
	}
	fmt.Println(utils.Pb2Base64(b))

	//config := &protos.ModelBenchmarkTestArgs{
	//	ModelPath:        "files/2023-03-08_20-07-09/resnet18-12.onnx",
	//	Framework:        protos.FrameworkType_onnxruntime,
	//	InputTensorShape: "1,3,512,512",
	//	InputTensorType:  protos.TensorDataType_float32,
	//	WarmupRounds:     1,
	//	RunRounds:        1,
	//}
	//m := jsonpb.Marshaler{}
	//str, err := m.MarshalToString(config)
	//fmt.Println(str)
	//
	//device := &g_physical_devices.PhysicalDeviceClient{
	//	DeviceDescription: devDesc,
	//}
	//
	//ability := model.ModelBenchmarkTestAbility{
	//	IsSupportModelBenchmarkTest:           true,
	//	SupportedFrameworksForRuntimeAnalysis: []protos.FrameworkType{protos.FrameworkType_onnxruntime},
	//	SupportedFrameworksForStaticAnalysis:  []protos.FrameworkType{protos.FrameworkType_onnxruntime},
	//}
	//

	//time.Sleep(2 * 1e9)
	//res, err := ability.ModelBenchmarkTest(config, device)
	//if err == nil {
	//	fmt.Println(res)
	//}
	//
	//var s string
	//_, _ = fmt.Scanln(&s)

}
