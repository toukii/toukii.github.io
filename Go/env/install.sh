#!/bin/sh
#cd ~
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
#curl -C - -O http://7xku3c.com1.z0.glb.clouddn.com/go1.4.min.tar.gz 
tar -zxf go1.4.min.tar.gz && cd go1.4/src && ./make.bash || ./make.bat
cd ../../
export CGO_ENABLED=1
export GOROOT_BOOTSTRAPT=go1.4
#curl -C - -O http://7xku3c.com1.z0.glb.clouddn.com/go1.8.min.tar.gz 
curl -C - -O http://upload.daoapp.io/loadfile/go1.8.min.tar.gz
tar -zxf go1.8.min.tar.gz && cd go/src && ./make.bash || make.bat
../bin/go version


