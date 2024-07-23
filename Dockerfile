# Golang镜像版本(构建阶段)
FROM golang:1.22.5-alpine AS build
# 工作目录
WORKDIR /app
# 依赖下载
COPY go.mod go.sum ./
RUN go mod download
# 复制项目代码到工作目录
COPY . .
# Golang 主应用打包
RUN go build -o main ./cmd/main.go
# Golang api 文档打包
RUN go build -o swagger ./cmd/swagger.go

# 使用官方的更小的镜像（运行阶段）
FROM alpine:latest
# 设置工作目录
WORKDIR /root/
# 从构建阶段复制编译后的二进制文件
COPY --from=build /app/main .
COPY --from=build /app/swagger .

# 启动Go 程序
CMD ["sh", "-c", "./main & ./swagger"]