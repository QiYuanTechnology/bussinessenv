# 兰州市营商环境评价的信息化平台

采用了go-micro微服务框架

## 约定
NAMESPACE = com.lzqysoft.bussinessenv


## 目录说明

<pre><code>
bussinessenv--项目根目录
|
|——api--公开的api接口服务，每个目录是一个服务直接处理HTTP请求的服务，通过RPC完全控制http请求/响应。
|
|——bin--文件格式化，项目编译脚本
|
|——common--项目公共资源
|
|——srv --底层微服务
</code></pre>