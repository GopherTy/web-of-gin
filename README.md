# Template of Gin

此项目基于 Go 语言中 Gin web框架搭建的模板。

架构：

整个项目结构模块化。内置 `config`,`logger`,`db` 模块，其中 `logger` 为 `zap` 包的 `logger`；`db` 使用 `xorm` ,数据库为

`MySQL` ; 配置文件支持 `yaml`, `json`,`jsonnet`。默认配置文件格式为 `json`

