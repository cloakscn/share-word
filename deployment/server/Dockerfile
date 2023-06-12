# 启动编译环境
FROM golang:1.20-alpine AS builder

# 拷贝源代码到编译环境
COPY . /go/src/share-word

# 配置编译环境&&编译
WORKDIR /go/src/share-word
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go install ../share-word/...

# 多阶段编译
FROM alpine:3.16
COPY --from=builder /go/bin/share-word /bin/share-word

# 设置服务入口
EXPOSE 8080
ENTRYPOINT [ "/bin/share-word" ]