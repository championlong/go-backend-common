# go-backend-common

通用组件

## 介绍

| 模块       | 说明                                                                        | 引用                                                                  |
|----------|---------------------------------------------------------------------------|---------------------------------------------------------------------|
| crypto   | 支持加解密，当前仅支持rsa公钥私钥加密解密                                                    |                                                                     |
| handler  | 初始化prometheus和pprof相关接口，支持服务集成监控                                          | https://github.com/prometheus/client_golang                         |
| hbase    | 使用[thrift2](https://thrift.apache.org/tutorial)批量操作hbase，支持连接池，减少频繁创建施放链接 | https://hbase.apache.org/                                           |
| postgres | 在pgv10基础上进行封装，支持数据库主从和分片配置读写                                              | https://github.com/go-pg/pg/v10                                     |
| slog     | 在zap日志打印基础进行封装，支持默认初始化调用                                                  | https://github.com/uber-go/zap                                      |
| viper    | 集成 fsnotify 和 vipe 实现多种格式的配置文件                                            | https://github.com/fsnotify/fsnotify https://github.com/spf13/viper |
| sql      | 实现动态拼接where条件，生成查询sql                                                     |                                                                     |
| app      | 实现统一服务初始化模版                                                               |                                                                     |

## 目录结构

```
├── app             (应用启动)
├── crypto          (加解密算法)
│   └── rsa
├── handler         (debug相关监控端口)
├── hbase           (hbase客户端)
├── postgres        (PostgreSQL初始化)
├── runner          (服务初始化接口)
├── slog            (日志服务)
├── sql             (sql动态拼接)
└── viper           (读取config)
```

