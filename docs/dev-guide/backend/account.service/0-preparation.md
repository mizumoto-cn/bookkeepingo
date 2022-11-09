# Before Development

## Install kratos cli and init project

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos new bookkeepingo
```

And you will see

```raw
🚀 Creating service bookkeepingo, layout repo is https://github.com/go-kratos/kratos-layout.git, please wait a moment.

Already up to date.

...
...
... ...

🍺 Project creation succeeded bookkeepingo
💻 Use the following command to start the project 👇:

$ cd bookkeepingo
$ go generate ./...
$ go build -o ./bin/ ./...
$ ./bin/bookkeepingo -conf ./configs

                        🤝 Thanks for using Kratos
        📚 Tutorial: https://go-kratos.dev/docs/getting-started/start
```

## Upgrade kratos and kratos tool

```bash
# 1 先升级到最新版
sudo kratos upgrade

# 2 安装 protoc-gen-go-http
go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2

# 3 安装 protoc-gen-go-errors
go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

# 4 安装
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go get -u github.com/envoyproxy/protoc-gen-validate
```

然后你就要跟傻逼makefile作斗争了。
