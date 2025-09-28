# blog_project
博客系统
概述：是一个基于 Go 语言的博客系统，使用 Gin 框架构建 RESTful API，采用 MySQL 数据库存储数据，使用 GORM 作为 ORM 工具。项目实现了用户注册、文章发布、评论等核心功能。

## 功能
用户功能：
- 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
- 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
文章功能：
- 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
- 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
- 实现文章的更新功能，只有文章的作者才能更新自己的文章。
- 实现文章的删除功能，只有文章的作者才能删除自己的文章。
- 评论功能：
- 实现评论的创建功能，已认证的用户可以对文章发表评论。
- 实现评论的读取功能，支持获取某篇文章的所有评论列表。

## 执行流程
## 数据库连接需要调整 位置在 dbs/dbTable.go 文件 修改数据库连接信息
- 使用 go mod tidy 初始化项目依赖
- 使用 go run main.go 启动项目
- postman文件  博客系统.postman_collection.json  在项目根目录下有postman测试接口文件
- 根路径： localhost:8080
生成相关表结构：
- 请求方式：GET
- URL: /init/createTable 

1. 用户注册：
- 请求方式：POST
- URL: /user/register
- 参数：username, password, email

2. 用户登录：
- 请求方式：POST
- URL: /user/login
- 参数：username, password
- 返回：Authorization  作为herder 放入请求头里，后续请求都需要带上这个token

3. 创建文章：
- 请求方式：POST
- 头部需要包含 Authorization
- URL: /post/createPost
- 参数：title, content

4. 获取文章列表：
- 请求方式：POST
- 头部需要包含 Authorization
- URL: /post/selectPost
- 参数：title, content 模糊查询 也可不传（查全部）

5. 更新文章：
- 请求方式：POST
- 头部需要包含 Authorization
- URL: /post/updatePost
- 参数：title, content，id

6. 删除文章：
- 请求方式：POST
- 头部需要包含 Authorization
- URL: /post/deletePost
- 参数：id

7. 创建评论：
- 请求方式：POST
- 头部需要包含 Authorization
- URL: /comment/createComment
- 参数：postTitle, comment

8. 获取评论列表：
- 请求方式：POST
- 头部需要包含 Authorization
- URL: /comment/selectComment
- 参数：postTitle(可不传查全部)

