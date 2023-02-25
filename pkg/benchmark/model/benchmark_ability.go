package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-dl-benchmark/pkg/devices"
	"strings"
)

type BenchmarkAbility struct {
	IsSupportHardwareBenchmarkTest bool `json:"IsSupportHardwareBenchmarkTest"`
	IsSupportModelBenchmarkTest    bool `json:"IsSupportModelBenchmarkTest"`

	SupportedFrameworksForStaticAnalysis  []string `json:"SupportedFrameworksForStaticAnalysis"`
	SupportedFrameworksForRuntimeAnalysis []string `json:"SupportedFrameworksForRuntimeAnalysis"`
}

func LoadAndPrintBenchmarkTestAbility(content string) BenchmarkAbility {
	return BenchmarkAbility{}
}

func (s *BenchmarkAbility) queryFrameworksSupportStatic(framwork string) bool {
	found := false
	for _, supportedFramework := range s.SupportedFrameworksForStaticAnalysis {
		if strings.Contains(supportedFramework, framwork) {
			found = true
			break
		}

	}
	return found
}

func (s *BenchmarkAbility) queryFrameworksSupportRuntime(framwork string) bool {
	found := false
	for _, supportedFramework := range s.SupportedFrameworksForRuntimeAnalysis {
		if strings.Contains(supportedFramework, framwork) {
			found = true
			break
		}

	}
	return found
}

func (s *BenchmarkAbility) ModelBenchmarkTest(config ModelBenchmarkTestConfig, device devices.HardwareDevice) (ModelBenchmarkTestResult, error) {
	// 1. 保证设备已连接
	err := device.Connect()
	if err != nil {
		log.Error("Can't connect device!.")
		return ModelBenchmarkTestResult{}, err
	}
	modelTestResult := ModelBenchmarkTestResult{}
	modelTestResult.TestConfig = config
	//print(s)
	if s.IsSupportModelBenchmarkTest {
		if s.queryFrameworksSupportStatic(config.Framework) {

			if config.Framework == Onnxruntime {
				staticResult := OnnxruntimeStaticAnalyse(config, device)
				modelTestResult.StaticResult = staticResult
			}
		} else {
			log.Warn(fmt.Sprintf("Don't support static analyse for %s now!", config.Framework))
		}
		if s.queryFrameworksSupportRuntime(config.Framework) {
			if config.Framework == Onnxruntime {
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
	return modelTestResult, nil
}
