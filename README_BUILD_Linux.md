### gocv安装
```
git clone https://github.com/hybridgroup/gocv.git
cd gocv
#安装依赖
sudo make deps 
#下载源码
sudo make download
#编译Linux 
sudo make build
make sudo_install
```

### 安装golang
```
下载go1.15.6.linux-amd64.tar.gz，解压到/home/pi/go 
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