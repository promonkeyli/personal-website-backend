# Golang镜像版本(构建阶段)
FROM golang:1.22.5-alpine AS build
# 工作目录
WORKDIR /app
# 拷贝依赖文件
COPY ../go.mod ../go.sum ./
# 下载Golang依赖
RUN go mod download
# 复制项目代码到工作目录
COPY ../ .
# Golang 打包构建
RUN go build -o main ../cmd/my_website_backend/main.go

# 使用官方的更小的镜像（运行阶段）
FROM alpine:latest
# 设置工作目录
WORKDIR /root/
# 从构建阶段复制编译后的二进制文件
COPY --from=build /app/main .
# 暴露运行端口
EXPOSE 8080
# 启动Go 程序
CMD ["./main"]