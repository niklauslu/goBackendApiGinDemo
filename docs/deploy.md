### 服务器部署

复制配置文件，并进行修改
```bash
cp .env.example .env
```
#### 方式一:直接使用docker部署

安装好docker以后，

```bash
// 从docker hub拉去镜像到本地
docker pull $(org)/$(name):$(tag) $(hub)/$(org)/$(name):$(tag)
```

使用`docker run`命令运行

```bash
docker run -itd --rm --name $(org)-$(name) \
	-p 9090:8080 \
	-v $(PWD)/logs:/www/logs \
	-v $(PWD)/.env:/www/.env \
	-v $(PWD)/uploads:/www/uplaods \
	$(org)/$(name):$(tag) 
```

-p:绑定端口  
-v:映射文件夹  
-e:设置环境变量(项目使用`.env`文件，这里就需要使用-e了)  

```bash
// 停止服务
docker container stop $(org)-$(name)
```

#### 方式二：使用`docker compose`

```yaml
// file:docker-compose.yml
version: '3'
services:
  backend:
    image: niklaslu/backend-api-gin-demo:1.0.0
    ports:
      - "8081:8080"
    volumes:
      - ./logs:/www/logs
      - ./.env:/www/.env 
```

启动服务
```bash
docker compose up -d
```
停止服务
```bash
docker compose stop
```