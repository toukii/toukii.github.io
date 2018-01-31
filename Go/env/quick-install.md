# Quick install Golang for linux

## install
	
	
```bash
curl http://www.golangtc.com/static/go/1.6.2/go1.6.2.linux-386.tar.gz > go1.6.tar.gz
tar -zxvf go1.6.tar.gz
export HOME=/home/shaalx
export GOROOT=$HOME/go6
mv go $GOROOT/
```


## env-setting
	

```bash
export HOME=/home/shaalx
echo "export HOME=/home/shaalx" >> $HOME/.bashrc
export GOROOT=$HOME/go
echo "export GOROOT=$HOME/go" >> $HOME/.bashrc
mkdir -p $HOME/GOPATH
export GOPATH=$HOME/GOPATH
echo "export GOPATH=$HOME/GOPATH" >> $HOME/.bashrc
mkdir -p $GOPATH/bin $GOPAHT/src $GOPATH/pkg
export GOBIN=$GOPATH/bin
echo 'export GOBIN=$GOPATH/bin' >> $HOME/.bashrc
echo 'export PATH=$PATH:$GOBIN:$GOROOT/bin' >> $HOME/.bashrc
source $HOME/.bashrc
```


## sublimeText
	
	curl https://download.sublimetext.com/sublime-text_build-3114_amd64.deb