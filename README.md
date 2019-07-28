## 说明

这个项目是一个图书馆的图书管理程序。
基于go语言的[beego](https://github.com/beego)框架。
项目还参考了[beego admin](https://github.com/beego/admin)的代码，在此感谢！

## 后台部分部署配置

建议基于linux来部署。当然，基于go语言的跨平台特性，也可以在windows上部署。

首先要下载go语言安装包[下载链接](https://golang.google.cn/)或[download link](https://golang.org)。

配置好两个环境变量。GOPATH和GOROOT。
GOROOT是安装目录，GOPATH是项目目录，下载的工具包、项目代码要放到这里来。

golang的环境配置好之后。执行下面的命令，下载相关需要的包。

```
go get -u -v github.com/beego/bee
go get -u -v github.com/astaxie/beego
go get -u -v github.com/go-sql-driver/mysql
```

