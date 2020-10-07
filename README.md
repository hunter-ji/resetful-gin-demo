# Resetful Gin Demo



## 简介

使用`gin`开发的resetful风格的todo项目，包括用户管理和todo的增删改查。使用query params的方式携带token来验证身份，中间件拦截请求，验证token，将相关用户信息写到上下文中。数据库连接使用`gin`的`GIN_MODE`环境变量来切换生产和开发环境。



## 路由

| 路由         | 说明                          |
| ------------ | ----------------------------- |
| /ping        | ping/pong检测mysql和redis连接 |
| /user/login  | 登录                          |
| /user/info   | 用户信息                      |
| /user/logout | 注销                          |
| /todo        | todo的增删改查                |
| /todo/single | 单个todo的详情                |

