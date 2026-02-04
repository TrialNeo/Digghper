<img src="\ico.png#pic_center" width="250"/>

AI生图真好用

向土木老哥们致敬， dig + gopher = digghper 

## 项目规范
1.每个目录 需要有独立的readme.md changelog.md contributor.md
## 基本目录说明
| file name  | description                 |
|------------|-----------------------------|
| cmd        | 命令，main.go，启动               |
| configs    | 配置文件，无代码                    |   
| internal   | 只允许包内模块互相调用，禁止外部调用，但是可以调用外部 |
| pkg        | 可复用的工具                      |
| global     | 全局变量/常量                     |
| initialize | 初始化                         |

## 更换框架问题
如果要更换web框架要修改的文件，变量，函数

| path                | var      | func          |
|---------------------|----------|---------------|
| global/const.go     | FbConfig | -             |
| global/global       | WebApp   | -             |
| initialize/fiber.go | -        | RunWebService |
| internal/route      | -        | -             |

## 一些可供替换的组件

| path                  | description |
|-----------------------|-------------|
| pkg/crypto/encrypt.go | 登录密码加密函数    |
|                       |             |
