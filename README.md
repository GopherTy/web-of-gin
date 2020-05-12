# Go web

### 此项目为 `Go` 语言搭建的 `web` 模板项目。

**技术栈**：

- 框架：`Gin`

- 数据库：`MySQL`
- `ORM`：`XORM`

**项目运行**：项目启动前通过配置文件（`.json` 文件）初始化数据库对象，然后启动 `Gin Web  ` 服务。通过配置文件可以决定启动的是服务类型是否为 `HTTPS` 。