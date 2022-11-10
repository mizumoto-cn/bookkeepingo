# Define the First API - Account

## Preparation

### Proto Definition

At `api/account/v1/account.proto`.

Then `make api` to generate the code.

### DB Definition

run `make set-env` to set the container environment.

```sql
create DATABASE account;
use account;
create table account
(
    id       int(64) not null auto_increment,
    accountmail varchar(64)  not null,
    password varchar(128) not null,
    phone    varchar(18),
    nickname varchar(20),
    PRIMARY KEY (id)
) engine=innodb default charset=utf8mb4;
```

### JWT Definition in `internal/conf/conf.proto`

The configuration of JWT is not written into the source code, but is written into the configuration file.

```proto
message Auth {
    string jwt_secret = 1;
    google.protobuf.Duration jwt_expire_duration = 2;
}
```

And add `auth` to the `Bootstrap` message.

```proto
message Bootstrap {
  Server server = 1;
  Data data = 2;
  Auth auth = 3;
}
```

Then run `make config`:

```bash
make config   
protoc --proto_path=./internal \
       --proto_path=./third_party \
               --go_out=paths=source_relative:./internal \
       internal/conf/conf.proto
```

Then, add the configuration to the `config.yaml` file.

```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
data:
  database:
    driver: mysql
    source: root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
# Authentication
auth:
  jwt_secret: "secret"
  expire_duration: 3600s
```

## Service Implementation

### `biz` Level Business Implementation

目前推荐使用贫血模型，只包含不依赖持久层的业务逻辑。依赖于持久层的业务逻辑将会被放到服务层中。因而贫血模型中的领域对象是不依赖于持久层的。

### `data` Level Data Access Implementation

被操作的实体是定义在biz层的，但是按照kratos的设计，存储到数据库的实体应该是data层自己定义的

#### Switch to `ent`

Install

```bash
go install entgo.io/ent/cmd/ent@latest
```

Generate

```bash
ent init Account
```

### API Interface Implementation - http & grpc

`internal/service/v1/account.go`
