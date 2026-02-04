| name       | 说明                    |
|------------|-----------------------|
| controller | 接收 HTTP 请求，调用 service |
| dao        | 数据库模型                 |
| service    | 纯业务逻辑，无context        |
| job        | 定时任务                  |
| client     | rpc相关，与外部互动           |
| route      | 路由                    |


