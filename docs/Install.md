### 安装go

#### windows

1. 下载安装包地址 https://go.dev/dl/
2. 配置环境变量，（安装后是否自动配置?待确认）
3. 验证是否安装完成，命令工具执行 go 

#### 配置代理
提高Go modules下载速度
1. 查看镜像
go env GOPROXY
2. 切换淘宝镜像
go env -w GOPROXY=https://goproxy.cn,direct
3. 切换国外镜像
go env -w GOPROXY=https://proxy.golang.org,direct


### 版本管理工具

gvm、goenv、asdf、g


#### 版本管理工具g

1. 下载 (release)[https://github.com/voidint/g/releases]
2. 创建目录~/.g/bin,压缩包解压g.exe放到该目录下
3. 配置环境变量 
```
官方文档，powershell里配置如下（我试了报错，$HOME不是家目录(~)吗？）
$env:GOROOT="$HOME\.g\go"
$env:Path=-join("$HOME\.g\bin;", "$env:GOROOT\bin;", "$env:Path")
最终通过手动添加环境变量
GOROOT %USERPROFILE%\.g\go  // go当前版本目录，g下载多个go版本，当前版本复制到该目录下
path里添加
%GOROOT%\bin // 全局使用go 命令
%USERPROFILE%\.g\bin // 全局使用g命令
```
4. 版本管理命令
```
g ls-remote // 查看全部可下载的版本
g ls-remote stable // 查看远程最近的稳定版本
g install 1.22.6 // 安装某个版本
g ls // 查看本地安装的版本列表
go use 1.22.6 // 使用某个版本
g uninstall 1.22.6 // 卸载某个版本
g clean // 移除下载的go版本压缩包
g self update // 更新g
g self uninstall // 卸载g
```