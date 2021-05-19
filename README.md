# [Together](https://github.com/ixugo/together)

Grpc 学习项目，前端使用 Flutter 技术，后端使用 Go Grpc。

名称来源「战斗天使阿丽塔」，原电影中位于 `01:52:45` 的画面

![image-20210317214638290](http://img.golang.space/shot-1618549116202.png)

## 目录结构

```bash
.
├── app										# 客户端
│   ├── cmd								# main 函数
│   ├── docs
│   ├── internal
│   │   ├── dao						# 数据库
│   │   ├── middleware		# 中间件
│   │   ├── routers				# 路由
│   │   └── service				# 业务
│   ├── pkg								# 外部包
│   └── scripts						# 自动化脚本
├── blog_server						# grpc 博客模块
├── im_server							# grpc IM通讯模块
├── configs
├── docs
├── global
├── model
└── proto
```

## 后端技术栈

- [Grpc](https://github.com/grpc/grpc-go)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm)
- [Colly](https://github.com/gocolly/colly)
- [ElasticSearch](https://github.com/elastic/go-elasticsearch)
- Redis
- PostgreSQL

## 接口规范

...
