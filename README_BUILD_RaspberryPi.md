### gocv安装
```
git clone https://github.com/hybridgroup/gocv.git
cd gocv
#安装依赖
sudo make deps 
#下载源码
sudo make download
#编译树莓派
sudo make build_raspi


遇到错误
“/tmp/opencv/opencv_contrib-4.5.0/modules/xfeatures2d/src/boostdesc.cpp:654:20: fatal error: boostdesc_bgm.i: No such file or directory
           #include "boostdesc_bgm.i"”
		   
cd /tmp/opencv/opencv_contrib-4.5.0/modules/xfeatures2d/src/
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_lbgm.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_binboost_256.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_binboost_128.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_binboost_064.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_bgm_hd.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_bgm_bi.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_bgm.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_120.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_64.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_48.i
sudo wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_80.i

make sudo_install

参考：
https://blog.csdn.net/z1314520cz/article/details/103501274
```

### 测试gocv
```
package main

import (
        "gocv.io/x/gocv"
)

func main() {
        webcam, _ := gocv.VideoCaptureDevice(0)
        window := gocv.NewWindow("Hello")
        img := gocv.NewMat()

        for {
                webcam.Read(&img)
                window.IMShow(img)
                window.WaitKey(1)
        }
}


树莓派32位linux系统会出现下面的错误，需要修改相应的代码才可以编译过
../../gopro/pkg/mod/gocv.io/x/gocv@v0.25.0/core.go:2009:16: type [1073741824]*_Ctype_char larger than address space
../../gopro/pkg/mod/gocv.io/x/gocv@v0.25.0/core.go:2009:16: type [1073741824]*_Ctype_char too large
修改
tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(strs.strs))[:length:length] 
为
tmpslice := (*[(1 << 29)-1]*C.char)(unsafe.Pointer(strs.strs))[:length:length] 

再次go build 
运行即可看到自己的视频了

参考：
https://github.com/docker/docker-credential-helpers/pull/61/files/cdde65956310ad369b40016dc7fff867b587877e
```


### 安装golang
```
下载go1.15.6.linux-armv6l.tar.gz，解压到/home/pi/go 
建立工程目录 mkdir -p /home/pi/gopro 

vim /etc/profile
在最后添加
export GOROOT=/home/pi/go 
export GOPATH=/home/pi/gopro 
export GOPROXY=https://goproxy.cn,direct 
export GO111MODULE=on 
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
生效
source /etc/profile
测试
pi@camera:~$ go version
go version go1.15.6 linux/arm
```

### 编译cvmotion
```
https://github.com/panyingyun/cvmotion.git
go build
```