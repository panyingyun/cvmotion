### CVMotion（异动监测） 
 - (1) 本地显示摄像头内容，并且标记异动
 - (2) 远程访问摄像头内容
 - (3) 分时段存储相关摄像头内容,仅仅存储异动内容，去除静止图形
 - (4) 保持的视频中图像和时间标识

### Run CVMotion 
```
example for windows:
cvmotion.exe 

example for linux:
./cvmotion 

example for 树莓派 32位:
./cvmotion 

change config by prod.yaml
```

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

编译提示缺少文件“fatal error: boostdesc_bgm.i: No such file or directory”，请参考
https://www.michaelapp.com/posts/2018/2018-09-15-%E6%A0%91%E8%8E%93%E6%B4%BE4B-%E8%BD%AF%E4%BB%B6%E5%AE%89%E8%A3%85/

cd C:\opencv\build
make -j4 
make install 
即可安装完成
```

### GoCV
```
将编译好的opencv添加到PATH环境变量
C:\opencv\build\install\x64\mingw\bin即可
```

### 支持视频写入
```
https://github.com/opencv/opencv/releases
下载opencv_videoio_ffmpeg450_64
opencv-4.5.0-vc14_vc15.exe
并且解压
将 opencv\build\bin 下的 opencv_videoio_ffmpeg450_64.dll
拷贝到C:\opencv\build\install\x64\mingw\bin
```