##	交叉编译


1.	compile：编译所有目标环境

		CGO_ENABLED=0
		GOOS=linux
		GOARCH=amd64
		$GOROOT/src/make.bash 
	
	_or_ 

		$GOROOT/src/all.bash

2.	build: go build

		CGO_ENABLED=0
		GOOS=linux
		GOARCH=amd64
		go build

3.	cross compile


		//linux
		CGO_ENABLED=1 GOOS=linux GOARCH=amd64 ./make.bash --no-clean
		CGO_ENABLED=1 GOOS=linux GOARCH=386 ./make.bash --no-clean
		 
		//嵌入式linux
		CGO_ENABLED=1 GOOS=linux GOARCH=arm ./make.bash --no-clean
		 
		//嵌入式freebsd
		CGO_ENABLED=1 GOOS=freebse GOARCH=amd64 ./make.bash --no-clean
		CGO_ENABLED=1 GOOS=freebse GOARCH=386 ./make.bash --no-clean
		//mac os
		CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 ./make.bash --no-clean
		CGO_ENABLED=1 GOOS=darwin GOARCH=386 ./make.bash --no-clean