package model

import (
	"fmt"
	"github.com/sheephuan/go-dl-benchmark/pkg/g_physical_devices"
	"github.com/sheephuan/go-dl-benchmark/pkg/protos"
	log "github.com/sirupsen/logrus"
)

var tools_path = "D:\\code\\Go-DL-Benchmark\\tools\\ModelProfileTool"

type ModelBenchmarkTestAbility struct {
	//IsSupportHardwareBenchmarkTest bool `json:"IsSupportHardwareBenchmarkTest"`
	IsSupportModelBenchmarkTest bool `json:"IsSupportModelBenchmarkTest"`

	SupportedFrameworksForStaticAnalysis  []protos.FrameworkType `json:"SupportedFrameworksForStaticAnalysis"`
	SupportedFrameworksForRuntimeAnalysis []protos.FrameworkType `json:"SupportedFrameworksForRuntimeAnalysis"`
}

func (s *ModelBenchmarkTestAbility) queryFrameworksSupportStatic(framework protos.FrameworkType) bool {
	found := false
	for _, supportedFramework := range s.SupportedFrameworksForStaticAnalysis {
		if framework == supportedFramework {
			found = true
			break
		}

	}
	return found
}

func (s *ModelBenchmarkTestAbility) queryFrameworksSupportRuntime(framework protos.FrameworkType) bool {
	found := false
	for _, supportedFramework := range s.SupportedFrameworksForRuntimeAnalysis {
		if framework == supportedFramework {
			found = true
			break
		}
	}
	return found
}

func (s *ModelBenchmarkTestAbility) ModelBenchmarkTest(config *protos.ModelBenchmarkTestArgs, device *g_physical_devices.PhysicalDeviceClient) (*protos.ModelAnalysisResult, error) {
	// 1. 保证设备已连接
	err := device.Connect()
	if err != nil {
		log.Error("Can't connect device!.")
		return &protos.ModelAnalysisResult{}, err
	}
	modelTestResult := protos.ModelAnalysisResult{}
	if s.IsSupportModelBenchmarkTest {
		if s.queryFrameworksSupportStatic(config.GetFramework()) {
			if config.Framework == protos.FrameworkType_onnxruntime {
				staticResult := OnnxruntimeStaticAnalyse(config, device)
				modelTestResult.StaticResult = staticResult
			}
		} else {
			log.Warn(fmt.Sprintf("Don't support static analyse for %s now!", config.Framework))
		}
		if s.queryFrameworksSupportRuntime(config.Framework) {
			if config.Framework == protos.FrameworkType_onnxruntime {
				runtimeResult := onnxruntimeRuntimeAnalyse(config, device)
				modelTestResult.RuntimeResult = runtimeResult
			}
		} else {
			log.Warn(fmt.Sprintf("Don't support runtime analyse for %s now!", config.Framework))
		}
	}
	err = device.Disconnect()
	//if err != nil {
	//	return ModelBenchmarkTestResult{}, err
	//}
	return &modelTestResult, err
}
