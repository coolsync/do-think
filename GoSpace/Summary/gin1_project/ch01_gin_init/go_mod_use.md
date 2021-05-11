## go-mod包管理

go版本>=1.11

## 一、什么是module？

go中包管理工具

## 二、使用module和不使用的区别

使用环境变量中的GO111MODULE控制是否使用mod

1.开启mod：go env -w GO111MODULE=on，**会将包下载到gopath下的pkg下的mod文件夹中**

2.关闭mod：go env -w GO111MODULE=off，**会将包下载到gopath下的src下**

3.go env GO111MODULE=auto,只有当当前目录在GOPATH/src目录之外而且当前目录包含go.mod文件或者其子目录包含go.mod文件才会启用。

**项目可以不用建在src下了，任何非中文路径下都可以，建议有个统一的代码路径**

## 三、go.mod文件的语法介绍

go help go.mod 查看帮助

示例：

```
module my/thing

go 1.13.4

require (
        new/thing v2.3.4
        old/thing v1.2.3
)
```

1.module：指明根目录

2.go 后面跟版本号是指定go的版本

2.require是个动作指令，对依赖包起作用，比如require(依赖)，还有exclude(排除)，replace(替代)，相同动作的可以放在一个动词+括号组成的结构中，如下：

```
require (
    new/thing v2.3.4
    old/thing v1.2.3
)

require new/thing v2.3.4
require old/thing v1.2.3




// 排除
exclude old/thing         v1.2.3

// 替换，使用箭头后的替换前面的
replace bad/thing         v1.4.5     => good/thing v1.4.5
```

注意：

- exclude和replace仅适用于主的go.mod文件中，其他的依赖中会被忽略、
- 可以使用replace替换无法获取的库，
  - 比如golang.org/x/crypto replace为github.com.com/golang/crypto

3.注释：使用//，没有/* xxx */这种块注释

## 四、go mod 命令

go mod help 查看帮助

```
download    下载模块到本地缓存，go env中的GOCACHE路径，可以通过go clean -cache清空缓存
            多个项目可以共享缓存的包
edit        在工具或脚本中编辑go.mod文件
graph       打印模块需求图
init        在当前目录下初始化新的模块
    go mod init 【项目名】    默认使用当前路径的项目名称
tidy        添加缺失的模块以及移除无用的模块，生成go.sum文件
vendor      会自动下载项目中依赖的包到项目根目录下的vendor文件夹下，并写入go.mod文件，同时生成
            modules.txt文件
            go mod vender -v     

verify      检查当前模块的依赖是否全部下载下来，是否下载下来被修改过
why         解释为什么需要包或模块


注意：-v参数可以查看执行的详细信息
```

已经完成的项目可以这样操作来使用mod

- 项目路径下执行go mod init
- 然后再执行go mod vendor（或者直接运行项目）

项目中可以是这样的执行顺序：

- init初始化 --> tidy 增删模块--> verify 校验模块-->vendor

***注意：项目中引入该项目下的任何路径都要是绝对路径，也就是以改项目名开头的路径***

使用mod的步骤：

1.开启mod:go111module=on

2.进入项目，执行go mod init (在项目根目录生成go.mod文件)

3.启动项目（go.mod添加依赖的包）