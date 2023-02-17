# 使用nn-Meter进行模型分析

[nn-Meter](https://github.com/microsoft/nn-Meter)
## 构建运行环境
```bash
conda activate -n nnmeter python=3.8
conda activate nnmeter

pip install nn-meter
```


## 执行
nn-meter 2.0中存在四个推理模型,在benchmark中我们会提供四个选项进行选择


| Predictor(device_inferenceframework)	| Processor Category	 |Version|
|--------------------------------|---------------------|-----|
| cortexA76cpu_tflite21      |  CPU	               |1.0|
| adreno640gpu_tflite21      |  GPU	               | 1.0                 |
| adreno630gpu_tflite21      |  GPU	               | 1.0                 |
| myriadvpu_openvino2019r2	 |  VPU	               | 1.0                 |



```bash

# for Tensorflow (*.pb) file
nn-meter predict --predictor <hardware> [--predictor-version <version>] --tensorflow <pb-file_or_folder> 
# Example Usage
nn-meter predict --predictor cortexA76cpu_tflite21 --predictor-version 1.0 --tensorflow mobilenetv3small_0.pb 

# for ONNX (*.onnx) file
nn-meter predict --predictor <hardware> [--predictor-version <version>] --onnx <onnx-file_or_folder>

```