FROM golang:1.22.3-alpine3.19  AS builder
#
WORKDIR /app/
## 设置golang代理
ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    CGO_ENABLED=0  \
    GOARCH="amd64" \
    GOOS=linux

COPY . .
RUN go mod tidy && go build --ldflags "-extldflags -static" -o main .

FROM alpine:3.19
WORKDIR /app/

# 修改时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# Install tzdata and set timezone to Asia/Shanghai
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

COPY --from=builder /app/main /app/main
COPY --from=builder /app/ui /app/ui
COPY --from=builder /app/config.yaml /app/config.yaml

EXPOSE 3000
CMD ["./main","serve"]