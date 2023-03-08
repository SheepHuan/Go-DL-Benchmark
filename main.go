package main

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_physical_devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	"os"
)

func main() {

	//ip, port := "127.0.0.1", 10001

	devDesc := protos.PhysicalDeviceDescription{}
	if jsonStr, err := os.ReadFile("config.json"); err == nil {
		err = jsonpb.UnmarshalString(string(jsonStr), &devDesc)
		if err != nil {
			return
		}
	}

	err := g_physical_devices.RegisterDevicesSelf([]*protos.PhysicalDeviceDescription{&devDesc})
	if err != nil {
		return
	}
	//config := &protos.ModelBenchmarkTestArgs{
	//	ModelPath:        "D:\\code\\Go-DL-Benchmark\\res\\resnet18-12.onnx",
	//	Framework:        protos.FrameworkType_onnxruntime,
	//	InputTensorShape: "1,3,512,512",
	//	InputTensorType:  protos.TensorDataType_float32,
	//	WarmupRounds:     1,
	//	RunRounds:        1,
	//}
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
	//go terminal.LaunchRemoteTerminalService(ip, port)
	//time.Sleep(2 * 1e9)
	//res, err := ability.ModelBenchmarkTest(config, device)
	//if err == nil {
	//	fmt.Println(res)
	//}
	//
	//var s string
	//_, _ = fmt.Scanln(&s)

}
