# Tesla_calculator
用于计算特斯拉线圈各项参数的小计算工具(https://github.com/visualfc/goqt 作为ui基础), 需要使用go x86版本

Need to use [go-bindata](https://github.com/go-bindata/go-bindata)

//go-bindata -pkg main -o images_bindata.go images

具备的功能：

    1. 计算初级线圈电感，谐振频率
    2. 计算次级线圈的电感，谐振频率，寄生电容
    3. 计算次级线圈与初级线圈的耦合系数
    4. 计算电弧长度
    
准备添加的功能：

    1. 添加计算次级线圈品质系数（Q）的计算功能， 公式为
                 Q = (2*Pi*f*L)/R


Support platform(Qt 5.11.1)
   Linux
   Windows
   Raspberry pi 3b(not test other pi)
 

## License
#### [MIT](https://sndnvaps.mit-license.org/2017)
