### GOlang的学习笔记

Q1: 学习顺序：

	从noob_go文件夹开始按照顺序阅读编译运行源码

Q2: 如何安装SDK:

	<https://golang.google.cn/dl/>或<https://golang.org/dl/>
	Windows选择go1.x.x.windows-amd64.msi
	下载并运行

Q3: 如何运行代码:

	cd noob_go
	go run .\noob_go\xx.go

Q4: package的用法：

	package main里面只能有一个main函数
	可以有很多go文件是package main

	或者 package 包名
	
	必须放在GOPATH得src里面
	不能有main函数