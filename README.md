### CVMotion（异动监测） 
 - (1) 本地显示摄像头内容，并且标记异动
 - (2) 远程访问摄像头内容
 - (3) 分时段存储相关摄像头内容,仅仅存储异动内容，去除静止图形
 - (4) 保持的视频中图像和时间标识
 - (5) 支持windows,Linux,树莓派32位/64w位

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

### Config CVMotion

config by prod.yaml

```
camera:
  deviceid: 0    # deviceID of Camera
window:
  enable: true   # when choose ture open window on desktop. other false
  title: Motion Window
  width: 1024
  height: 768
web:
  enable: true   # open web server or not 
  host: 0.0.0.0
  post: 9000
video:
  enable: false  # open auto save motion to video file or not 
  fps: 25
  prefix: motion
```

