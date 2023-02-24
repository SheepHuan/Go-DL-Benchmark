import onnx_tool
from onnx_tool import create_ndarray_f32
import argparse
import time
import re
import json
parser = argparse.ArgumentParser(prog = "onnx_static.py",
        description = "This is a python script for onnxruntime static analysis.",
        epilog = "This is an additional line. Last Line")
parser.add_argument("--model_path",  type = str, default = "NONE", help = "file path for *.onnx")
parser.add_argument("--shape",  type = str, default = "NONE", help = "input shape [batch,c,w,h], such as 1,3,512,512")
parser.add_argument("--type",  type = str, default = "NONE", help = "input data type, int8,float32,float64")

# TODO We Need chech these args is or not valid.
args = parser.parse_args()
model_path = args.model_path
input_shape = [int(item) for item in args.shape.split(",")]
type = args.type
inputs = {
    'inputs':create_ndarray_f32(input_shape)
}
save_path = f"res/{str(int(time.time() * 1000))}.txt"
# python .\onnx_static.py --model_path="D:\code\Go-DL-Benchmark\res\resnet18-12.onnx" --shape=1,3,512,512 --type=float32
onnx_tool.model_profile(model_path,inputs,savenode = save_path) 

lines = open(save_path,"r").readlines()[2:]
lines = [re.sub(r"\s+", " ", line).split(' ') for line in lines]
lines = [item for item in lines if item!='']
result = {}
result["OperateInfo"] = []

def str_remove_comma_to_int(string):
    string = string.replace(",","")
    return int(string)

def str_remove_percent_sign_to_float(string):
    string = string.replace("%","")
    return float(string)

for line in lines[:-1]:
    item = {
        "Name": line[0],
        "MACs": str_remove_comma_to_int(line[1]),
        "MACsPercent": str_remove_percent_sign_to_float(line[2]),
        "Memory": str_remove_comma_to_int(line[3]),
        "MemPercent": str_remove_percent_sign_to_float(line[4]),
        "Params": str_remove_comma_to_int(line[5]),
        "ParamsPercent": str_remove_percent_sign_to_float(line[6]),
        "InShape": line[7],
        "OutShape": line[8]
    }
    result["OperateInfo"].append(item)

result["MACs"] = str_remove_comma_to_int(lines[-1][1])
result["Memory"] = str_remove_comma_to_int(lines[-1][3])
result["Params"] = str_remove_comma_to_int(lines[-1][5])
open("result.json","w").write(json.dumps(result,indent=4,ensure_ascii=False))
