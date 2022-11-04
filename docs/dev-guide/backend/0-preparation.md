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