<p align="center">
<a href="https://www.wangsu.com/"><img src="https://static-wcs.wangsu.com/portalnav/navLogo_1720611253498_152_网宿logo-23_512.png" height="400" ></a>
</p>
<h1 align="center">WangSU  SDK for Go</h1>

# 目录
1. [简介](#简介)
2. [安装](#获取安装)


# 简介

欢迎使用网宿科技开发者工具套件（SDK），此 SDK 是网宿OpenApi平台的配套开发工具。

# 获取安装

## 依赖环境

1. Go 1.17 版本及以上
2. 部分产品需要在网宿控制台开通后，才能正常调用此产品的接口。
3. 在网宿控制台 [accessKey管理](https://console.wangsu.com/account/accessKey) 页面获取密钥 accessKey ID 和 accessKey Secret，请务必妥善保管，或者使用更安全的临时安全凭证。


## 通过go get安装（推荐）

推荐使用阿里云镜像加速下载：

1. Linux 或 MacOS:

    ```bash
    export GOPROXY=https://mirrors.aliyun.com/goproxy
    ```

2. Windows:

    ```cmd
    set GOPROXY=https://mirrors.aliyun.com/goproxy
    ```

## 按需安装（推荐）

注意：此安装方式仅支持使用 **Go Modules** 模式进行依赖管理，即环境变量 `GO111MODULE=auto`或者`GO111MODULE=on`, 并且在您的项目中执行了 `go mod init xxx`.

如果您使用 GOPATH, 请参考下节： 全部安装

```cmd
    go mod tidy
```