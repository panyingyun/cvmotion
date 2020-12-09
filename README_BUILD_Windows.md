### 安装 cmake/gcc
```
前提条件：cmake和TDM-GCC-64

并且环境变量PATH中包含

C:\Program Files\CMake\bin 和 C:\TDM-GCC-64\bin
```

### 安装OpenCV 

```
win_build_opencv.cmd 修改GoCV自动安装脚本
中的make的"set PATH"的那一行内容
通常本地已安装了cmake和TDM-GCC-64即可 无需再设置PATH

编译提示缺少文件“fatal error: boostdesc_bgm.i: No such file or directory”，请
cd opencv/opencv_contrib-4.5.0/modules/xfeatures2d/src/
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_lbgm.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_binboost_256.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_binboost_128.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_binboost_064.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_bgm_hd.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_bgm_bi.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/34e4206aef44d50e6bbcd0ab06354b52e7466d26/boostdesc_bgm.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_120.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_64.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_48.i
wget https://raw.githubusercontent.com/opencv/opencv_3rdparty/fccf7cd6a4b12079f73bbfb21745f9babcd4eb1d/vgg_generated_80.i

cd C:\opencv\build
make -j4 
make install 
即可安装完成
编译后的DLL请查看
C:\opencv\build\install\x64\mingw\bin
```

### 安装golang 
```
reference to golang.org,下载对应的安装包进行安装
```

### GoCV
```
将编译好的opencv添加到PATH环境变量
C:\opencv\build\install\x64\mingw\bin即可
```


### 编译
```
https://github.com/panyingyun/cvmotion.git
go build
```