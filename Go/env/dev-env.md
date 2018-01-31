Dev env setting
===================

## 1 [安装go，配置环境变量](index.html)


## 2 下载gocode实现编码提示

```bash
go get github.com/nsf/gocode
go install github.com/nsf/gocode
```

## 3 下载Sublime Text 3


### sublimeText

```
curl https://download.sublimetext.com/sublime-text_build-3114_amd64.deb
```

>Ctrl+\` 调出命令行，输入如下代码，重启sublime text。

------------------------


```
import urllib.request,os; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); open(os.path.join(ipp, pf), 'wb').write(urllib.request.urlopen( 'http://sublime.wbond.net/' + pf.replace(' ','%20')).read())
```



## 4 这一步也很重要
1. sublimetext 安装插件 GoSublime
2. sublimetext 安装插件 GoDef
3. Preference --> Package Setting --> GoSublime --> Setting Default，修改env：


```json
"env": {
	"path":"G:\\Go\\bin",
	"GOROOT":"G:\\Go",
	"GOPATH":"$GOPATH"
},
```


##	Github Configure

* (1) SSH KEY GEN

>ssh-keygen -t rsa -C "account@gmail.com"



//填写email地址，然后一直“回车”ok
打开本地..\.ssh\id_rsa.pub文件。此文件里面内容为刚才生成人密钥。

* (2) SSH KEY ADD

登陆github系统。点击右上角的 Account Settings--->SSH Public keys ---> add another public keys.

把你本地生成的密钥复制到里面（key文本框中）， 点击 add key 就ok了

* (3) Test

>ssh -T git@github.com



如果提示：Hi defnngj You've successfully authenticated, but GitHub does not provide shell access. 说明你连接成功了

## git 多账号

http://memoryboxes.github.io/blog/2014/12/07/duo-ge-gitzhang-hao-zhi-jian-de-qie-huan/

```
git config user.user toukii
git config user.email toukii@yahoo.com
```


