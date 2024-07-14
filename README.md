![GitHub language count](https://img.shields.io/github/languages/count/stormsha/go_backend_template)
![GitHub top language](https://img.shields.io/github/languages/top/stormsha/go_backend_template)
![Repo Size](https://img.shields.io/github/repo-size/stormsha/go_backend_template)
[![License](https://img.shields.io/github/license/stormsha/go_backend_template)](https://github.com/stormsha/go_backend_template/blob/master/LICENSE)

### Golang 实现的基础服务

每次起一个新项目时都需要从老项目中拷贝公共代码，虽然不难搞，但是从创建项目，搭建出基本开发框架，还是需要一点时间的，于是索性，把基本上每个Web
API 服务需要的功能模块整理为一个基础项目模板，用来帮助我快速开始一个项目，而不是重头开始写，浪费大量时间和精力。

> 目前只是搭建一个基础项目框架，未提供快速上线的工作流

### 技术栈

- [x] Web 框架：[Echo](https://echo.labstack.com/)
- [x] 数据库：MySql
- [x] 数据库操作：[GORM](https://gorm.io/)
- [x] 身份认证机制：[golang-jwt](https://golang-jwt.github.io/jwt/)
- [x] 日志器： [logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
- [x] 接口文档：[swaggo](https://github.com/swaggo/swag)

## 如何使用?

确保已经安装 [golang](https://go.dev/) >= 1.14.x

1. 获取代码到本地

```bash
git clone https://github.com/stormsha/go_backend_template.git
```

2. 下载项目依赖

```bash
cd 项目
go mod tidy
```

3. 运行项目

```bash
go run main.go
```

4. 访问API文档

[http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html)

![效果图](https://github.com/stormsha/go_backend_template/blob/master/swagger.png)