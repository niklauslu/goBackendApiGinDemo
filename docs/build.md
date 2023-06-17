### 部署之镜像打包

在项目根目录新建一个`Dockerfile`，外加一个`.dockerignore`(打包不需要的文件路径添加到这里)

由于直接用golang镜像打包出来的会image文件会很大，所以这里做的是两步式。  
+ 先用`golang`的镜像打包二进制文件
+ 然后再把二进制文件添加到`alpine`（非常轻量级Linux发行版）环境运行


具体Dockerfile示例如下
```dockerfile
// file: Dockerfile
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
```

docker命令:
```bash
// 打包命令
docker build -t $(org)/$(name):$(tag) . --no-cache

// 打包完后，打tag上传到hub
docker tag $(org)/$(name):$(tag) $(hub)/$(org)/$(name):$(tag)
docker push $(org)/$(name):$(tag)
```