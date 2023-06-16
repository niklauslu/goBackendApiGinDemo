### 项目环境配置

[godotenv](https://github.com/joho/godotenv)

```
go get github.com/joho/godotenv
```

在更目录创建一个`.env`文件
```
CONF_NAME=CONF_VALUE
```

```
// 读取配置文件
err := godotenv.Load()
if err != nil {
log.Fatal("Error loading .env file")
}

// 读取配置项
confVal := os.Getenv("CONF_NAME")
```