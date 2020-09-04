# 跨平台网站克隆器

支持网站克隆，网页克隆，网站复制，网页复制到本地。

只需输入网页路径即可复制


所需要的go版本 >= go1.14.4 


**1.调试运行** 

git clone https://github.com/panglove/webcloner.git

cd webcloner


go run main.go

**2.跨平台编译**


首先得安装docker容器

go get github.com/lucor/fyne-cross/v2/cmd/fyne-cross


windows:

64位编译:

fyne-cross --targets=windows/amd64 webcloner

32位编译:

fyne-cross --targets=windows/386 webcloner


linux:
fyne-cross --targets=linux/amd64 webcloner


macos:

fyne-cross --targets=darwin/amd64 webcloner


**3.打包**

go get fyne.io/fyne/cmd/fyne



fyne package -os darwin -icon logo.png

fyne package -os linux -icon logo.png
fyne package -os windows -icon logo.png




**4.预览**


![image](https://gitee.com/seelove792/GoEthWallet/raw/master/image/5.png)
![image](https://gitee.com/seelove792/GoEthWallet/raw/master/image/6.png)
