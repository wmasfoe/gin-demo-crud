# 学习记录

执行 `go env`，查看 `go111module="on"`，没🪜查看是否 配置 `GOPROXY` 推荐两个源：https://goproxy.cn/  https://goproxy.io/

> 这是大坑！可看博客：[go-get-fail](https://blog.justdev.cn/blog/golang/go-get-fail)

项目初始化

```sh
go mod init $project_name
touch main.go
```

下载第三方依赖：

```sh
go get github.com/gin-gonic/gin
```

同一个 package 下不需要 import，但要注意同名，推荐使用 struct 包裹函数名，见 controllers 文件夹。

对于不同 package 引用函数，需要：

```go
import (
  "$project_name/$package_name"
)
```

> $project_name 和上面 `go mod int` 的名字一致，$package_name 和导入的包名一致
