# bookkeepingo

A cloud bookkeeping ware. **--Development in progress--**

## Architecture

```bullshit
├── api  // API&Error Proto files & Generated codes
│   ├── foo
│   │   ├── job
│   │   └── service
│   └── bar
│       └── interface
├── app  // kratos microservices projects
│   ├── foo
│   │   ├── job
│   │   └── service
│   └── bar
│       └── interface
├── configs  // static kratos configs
├── internal  // internal packages which are not exposed to other projects
│   ├── biz // business logic 业务逻辑的组装层
│   ├── conf // config for applications in protobuf format, after compiled, it will be used to help business logic to read config
│   ├── data // data access of database/redis-cache or other data sources including lower-stream data interfaces 业务数据访问，封装cache/db，实现了biz层的repo接口，偏重业务。作用是把DO重新拿出来。
│   ├── server // server defines the application-layer protocol
│   └── service // service is for the implementation of the apis 实现api定义的服务层，实现DTO->DO(biz领域实体)并协同各类biz交互，但不处理复杂逻辑。
├── pkg  // common used packages
├── deploy  // dockerfile and deployment scripts
├── third_party  // proto files for third party services
├── web  // web frontend files
└── docs // documents
```

## Interface API

## Package API

## Deployment

## Web Frontend

Developed with [Vue.js](https://vuejs.org/) in [typescript](https://www.typescriptlang.org/).

## License

Proudly governed by the [Mizumoto General Public License](./Mizumoto%20General%20Public%20License%20v1.2.md), a license based on Mozilla Public License 2.0 with extra restrictions.
