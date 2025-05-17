# Upgrade Library
## 简介
upgrade 是一个用于在运行时升级可执行文件的 Go 语言库。该库针对不同的操作系统（Windows、Linux、macOS）提供了不同的文件替换策略，以确保在运行中的可执行文件能够被安全地替换为新版本。
### 特点
- 跨平台支持：支持 Windows、Linux 和 macOS 操作系统。
- 运行时升级：允许在程序运行时对可执行文件进行升级。
- 简单易用：提供了简单的 API 接口，方便集成到现有项目中。
## 安装
使用 go get 命令安装该库：
```shell
go get github.com/Li-giegie/upgrade
```
## 使用方法
### 示例代码
以下是一个简单的示例，展示了如何使用 upgrade 库进行可执行文件的升级：
```go
package main

import (
    "flag"
    "github.com/Li-giegie/upgrade"
    "log"
    "os"
    "time"
)

var isUpgrade = flag.Bool("upgrade", false, "upgrade to the latest version")
var dstFile = flag.String("dst", "", "dst filename")

func main() {
    flag.Parse()
    if !*isUpgrade {
        // 1. go build -o a.exe
        Start("Process-A")

        // 2. go build -o a.exe
        //Start("Process-B")
        return
    }
    err := upgrade.Upgrade(*dstFile, os.Args[0])
    if err != nil {
        log.Fatal(err)
    }
    log.Println("upgrade to the latest version")
}

func Start(name string) {
    for i := 0; i < 100; i++ {
        log.Printf("%s running\n", name)
        time.Sleep(1 * time.Second)
    }
}
```
首先构建一个a可执行程序文件 go build -o a.exe

把源码中a进程的代码注释掉，b进程的代码解开注释
```go
if !*isUpgrade {
    // 1. go build -o a.exe
    //Start("Process-A")

    // 2. go build -o a.exe
    Start("Process-B")
    return
}
```
在构建一个b可执行程序文件，作为a可执行程序的新版本 go build -o b.exe

执行a.exe 会间隔1秒钟输出Process-A的文本

执行 a.exe -upgrade -dst b.exe 后a可执行程序文件就被替换为b可执行程序文件了

在windows系统中，原本a.exe被命名为 纳秒时间戳+.+随机数+.+源文件名（111111.123123.a.exe）可以使用UpgradeWithOutFilename函数入参传递outFileName指针以获取生成的随机文件名

再次执行a.exe就能看到Process-B的文本输出了
### 代码解释
1. 命令行参数：通过 flag 包解析命令行参数，-upgrade 用于指定是否进行升级操作，-dst 用于指定目标文件的路径。
2. 升级操作：调用 upgrade.Upgrade 函数进行升级操作，该函数接受目标文件路径和源文件路径作为参数。
### 操作系统支持
## Windows
在 Windows 系统中，运行中的可执行文件不能被删除，但可以重命名。因此，upgrade 库会先将当前进程的可执行文件重命名为一个随机的文件名，然后将升级下载的新文件重命名为当前进程的文件名。
### Linux 和 macOS
在 Linux 和 macOS 系统中，直接对可执行程序文件进行覆盖命名。
