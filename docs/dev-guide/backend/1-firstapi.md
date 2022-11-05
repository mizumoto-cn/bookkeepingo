# Define the First API - Account

## Proto Definition

At `api/account/v1/account.proto`.

Then `make api` to generate the code.

## DB Definition

run `make set-env` to set the container environment.

```sql
create table user
(
    id       int(64) not null auto_increment,
    username varchar(64)  not null,
    password varchar(128) not null,
    phone    varchar(18),
    nickname varchar(20),
    PRIMARY KEY (id)
) engine=innodb default charset=utf8mb4;
```

