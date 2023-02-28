package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-dl-benchmark/pkg/devices"
	"go-dl-benchmark/pkg/protos"
)

type BenchmarkAbility struct {
	IsSupportHardwareBenchmarkTest bool `json:"IsSupportHardwareBenchmarkTest"`
	IsSupportModelBenchmarkTest    bool `json:"IsSupportModelBenchmarkTest"`

	SupportedFrameworksForStaticAnalysis  []protos.FrameworkType `json:"SupportedFrameworksForStaticAnalysis"`
	SupportedFrameworksForRuntimeAnalysis []protos.FrameworkType `json:"SupportedFrameworksForRuntimeAnalysis"`
}

func LoadAndPrintBenchmarkTestAbility(content string) BenchmarkAbility {
	return BenchmarkAbility{}
}

func (s *BenchmarkAbility) queryFrameworksSupportStatic(framework protos.FrameworkType) bool {
	found := false
	for _, supportedFramework := range s.SupportedFrameworksForStaticAnalysis {
		if framework == supportedFramework {
			found = true
			break
		}

	}
	return found
}

func (s *BenchmarkAbility) queryFrameworksSupportRuntime(framework protos.FrameworkType) bool {
	found := false
	for _, supportedFramework := range s.SupportedFrameworksForRuntimeAnalysis {
		if framework == supportedFramework {
			found = true
			break
		}
	}
	return found
}

func (s *BenchmarkAbility) ModelBenchmarkTest(config *protos.ModelBenchmarkTestArgs, device devices.HardwareDevice) (*protos.ModelAnalysisResult, error) {
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
			//if config.Framework == protos.FrameworkType_onnxruntime {
			//	runtimeResult := onnxruntimeRuntimeAnalyse(config, device)
			//	modelTestResult.RuntimeResult = runtimeResult
			//}
		} else {
			log.Warn(fmt.Sprintf("Don't support runtime analyse for %s now!", config.Framework))
		}
	}
	err = device.Disconnect()
	//if err != nil {
	//	return ModelBenchmarkTestResult{}, err
	//}
	return &modelTestResult, nil
}
