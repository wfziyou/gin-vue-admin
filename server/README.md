## server项目结构

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── packfile
├── resource
│   ├── excel
│   ├── page
│   └── template
├── router
├── service
├── source
└── utils
    ├── timer
    └── upload
```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`        | api层                   | api层 |
| `--v1`       | v1版本接口              | v1版本接口                  |
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `docs`       | swagger文档目录         | swagger文档目录 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | router,redis,gorm,validator, timer的初始化 |
| `--internal` | 初始化内部函数 | gorm 的 longger 自定义,在此文件夹的函数只能由 `initialize` 层进行调用 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `packfile`   | 静态文件打包            | 静态文件打包 |
| `resource`   | 静态资源文件夹          | 负责存放静态文件                |
| `--excel` | excel导入导出默认路径 | excel导入导出默认路径 |
| `--page` | 表单生成器 | 表单生成器 打包后的dist |
| `--template` | 模板 | 模板文件夹,存放的是代码生成器的模板 |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `source` | source层 | 存放初始化数据的函数 |
| `utils`      | 工具包                  | 工具函数封装            |
| `--timer` | timer | 定时器接口封装 |
| `--upload`      | oss                  | oss接口封装        |

1、首先需要注册微信开放平台，然后获取开发者认证。审批通过之后再创建一个移动应用同样还是需要审批。通过之后就可以给这个应用添加微信授权登陆以及相应功能了。这里移动应用审批通过之后会给你两个参数，一个叫AppId，一个叫Secret。这两个参数在后面用的到。


https://github.com/go-oauth2/gin-server

文档  
[GORM 指南](https://gorm.io/zh_CN/docs/)  
# 按照go环境

# 下载源码：
```
git clone https://github.com/wfziyou/gin-vue-admin.git
```

# 安装swagger
```
go install github.com/swaggo/swag/cmd/swag@latest
```
swag 安装在 /root/bin/go/中的，所以，需要在环境变量（/etc/profile）中增加上
`export PATH=$PATH:$GOROOT/bin:$GOPATH/bin:/root/bin/go`
让环境变量生效
`source /etc/profile`
安装依赖
```
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles
```
无法翻墙 由于国内没法安装 go.org/x 包下面的东西，推荐使用 goproxy.cn 或者 goproxy.cn/
```
# Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,direct

# 使用如下命令下载swag
go install github.com/swaggo/swag/cmd/swag@latest
```
# 生成swag文档
```
swag init
```
拉依赖
```
go mod tidy
# 选择上下命令其一执行即可
go generate -x // -x 显示并执行命令
```

# 打包
```
windows:
go build -o myserver.exe main.go
linux:
go build -o server .
```
docker 打包
```
修改配置：config.yaml 中mysql.path 为本地地址10.0.4.10

docker build -t server .

docker run -d --name myserver -p 8888:8888 server
```
# 常见问题
## 服务端是否启动
netstat -anp | grep 8888
## go环境是已安装，却提示没有安装
配置go的环境变量生效 `source /etc/profile`
