# go-memo-server
一个乱七八糟的备忘录后端作业

## 部署
1. 在`configs/conf.yaml`配置PostgreSQL、Redis连接信息，以及JWT密钥和监听信息
2. 使用`docs/sql`中的SQL文件建表
3. 启动服务器

## API文档

用`swag`来自动生成`Swagger`文档，虽然效果也许有点一言难尽
- [Swagger API](docs/API.md)

## 项目结构
三层架构没啥经验，也许搞得乱七八糟的...
```text
💾 go-memo-server/
├── 📄 .gitignore
├── 📄 .hz
├── 📄 go.mod
├── 📄 go.sum
├── 📄 main.go               # 服务器入口点
├── 📄 README.md
│
├── 📁 docs/                 # API文档和建表SQL
│   ├── 📁 sql/              # 建表SQL
│   ├── 📄 API.md            # ApiFox生成的API文档
│   ├── 📄 swagger.json      # swag生成的API文档
│   ├── 📄 swagger.yaml
│   └── 📄 swagger.go
│
├── 📁 configs/                 # 配置文件目录
│   ├── 📄 conf.yaml
│   └── 📄 conf.yaml.example
│
├── 📁 internal/
│   ├── 📁 dal/                 # 数据访问层
│   │   ├── 📁 cache/
│   │   │   ├── 📄 cache.go     # 自己封装的RedisClient
│   │   │   └── 📄 redis.go
│   │   ├── 📁 db/
│   │   │   └── 📄 postgre.go
│   │   └── 📄 dal.go
│   │
│   ├── 📁 handler/             # 请求处理层
│   │   ├── 📁 auth/            # 认证
│   │   ├── 📁 memo/            # 备忘录
│   │   └── 📄 base.go
│   │
│   ├── 📁 model/               # 数据模型
│   ├── 📁 mw/                  # 中间件
│   │   └── 📄 jwt.go
│   │
│   ├── 📁 repository/          # 仓储层
│   ├── 📁 service/             # 业务逻辑层
│   └── 📄 router.go            # 路由
│
└── 📁 pkg/                     # 可公开使用的包 (好怪)
    └── 📁 config/              # 配置文件处理
```