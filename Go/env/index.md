# Golang 环境搭建

## [快速搭建 fro linux](quick-install.html)

## [源码安装](source-install.html)

### install 

```
#!/bin/sh
dir=$PWD
curl --version
if [ $? -eq 127 ];then
  echo "please install curl first"
  exit 0
fi
gcc --version
if [ $? -eq 127 ];then
  echo "please install gcc first"
  exit 0
fi
curl -C - -O http://7xku3c.com1.z0.glb.clouddn.com/go1.4.min.tar.gz 
tar -zxvf go1.4.min.tar.gz && cd go1.4/src && ./make.bash || ./make.bat
export CGO_ENABLED=1
export GOROOT_BOOTSTRAP=$dir/go1.4
cd $dir
curl -C - -O http://7xku3c.com1.z0.glb.clouddn.com/go1.8.min.tar.gz 
tar -zxvf go1.8.min.tar.gz && cd go/src && ./make.bash || ./make.bat
echo "**********************go version********************"
$dir/go/bin/go version
```


### env 

```
export HOME=/home/toukii 
echo "export HOME=/home/toukii" >> $HOME/.bashrc 
export GOROOT=$HOME/go 
echo "export GOROOT=$HOME/go" >> $HOME/.bashrc 
mkdir -p $HOME/GOPATH 
export GOPATH=$HOME/GOPATH 
echo "export GOPATH=$HOME/GOPATH" >> $HOME/.bashrc 
mkdir -p $GOPATH/bin $GOPATH/src $GOPATH/pkg 
export GOBIN=$GOPATH/bin 
echo 'export GOBIN=$GOPATH/bin' >> $HOME/.bashrc 
echo 'export PATH=$PATH:$GOBIN:$GOROOT/bin' >> $HOME/.bashrc source $HOME/.bashrc 
```


1.	需要用到GCC等工具，Windows用户可安装__[tdm-gcc](http://sourceforge.net/projects/tdm-gcc/)__,将gcc的bin目录添加至$PATH下。

## [交叉编译](cross.html)


## [开发环境配置](dev-env.html)
