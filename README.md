# bookkeepingo

A cloud bookkeeping ware. **--Development in progress--**

![work time](https://wakatime.com/badge/user/29d180ce-15fa-4656-8a52-64d706899e5d/project/0dcb4b23-3cd0-4a54-a383-7f58db6716f8.svg?style=default)

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
|   |       └── internal    // internal packages which are not exposed to other projects
|   |           ├── biz     // business logic 业务逻辑的组装层
|   |           ├── conf    // config for applications in protobuf format, after compiled, it will be used to help business logic to read config
|   |           ├── data    // data access of database/redis-cache or other data sources including lower-stream data interfaces
|   |           ├── server  // server logic, including http/grpc server, middleware, etc.
|   |           └── service // service is for the implementation of the apis 实现api定义的服务层，实现DTO->DO(biz领域实体)并协同各类biz交互，但不处理复杂逻辑。
│   └── bar
│       └── interface
├── configs  // static kratos configs
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

Proudly governed by the [Mizumoto General Public License](./LICENSE.md), a license based on Mozilla Public License 2.0 with extra restrictions.
