# Before Development

## Install kratos cli and init project

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
kratos new bookkeepingo
```

And you will see

```raw
ð Creating service bookkeepingo, layout repo is https://github.com/go-kratos/kratos-layout.git, please wait a moment.

Already up to date.

...
...
... ...

ðº Project creation succeeded bookkeepingo
ð» Use the following command to start the project ð:

$ cd bookkeepingo
$ go generate ./...
$ go build -o ./bin/ ./...
$ ./bin/bookkeepingo -conf ./configs

                        ð¤ Thanks for using Kratos
        ð Tutorial: https://go-kratos.dev/docs/getting-started/start
```

## Upgrade kratos and kratos tool

```bash
# 1 ååçº§å°ææ°ç
sudo kratos upgrade

# 2 å®è£ protoc-gen-go-http
go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2

# 3 å®è£ protoc-gen-go-errors
go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2

# 4 å®è£
go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go get -u github.com/envoyproxy/protoc-gen-validate
```

ç¶åä½ å°±è¦è·å»é¼makefileä½æäºäºã
