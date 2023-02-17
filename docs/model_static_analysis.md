# 模型静态分析

PaddleLite(不能用onnx tool)
TFLite(opt可以)

分类
1. Tensor 从数据出发(计算量、参数量)
   [深度学习中模型计算量(FLOPs)和参数量(Params)的理解以及四种计算方法总结](https://blog.csdn.net/qq_40507857/article/details/118764782)
   - 
2. Operate 从计算OP类型出发
   - Name
   - Macs
   - CPercent
   - Memory
   - MPercent
   - Params
   - PPercent
   - InShape
   - OutShape
3. Graph   分析瓶颈
