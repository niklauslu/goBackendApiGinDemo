FROM golang:latest as builder
WORKDIR /app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o bin/app .

FROM alpine
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++ curl
ENV TZ Asia/Shanghai
WORKDIR /www/
COPY --from=builder /app/bin/app .
# EXPOSE 8080
ENTRYPOINT ./app