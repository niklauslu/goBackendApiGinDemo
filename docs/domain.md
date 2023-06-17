### Nginx域名绑定

docker服务配合nginx

nginx配置文件(https)
```nginx
// file: /etc/nginx/conf.d/domain.conf
server {
    listen 443 ssl;
    server_name domain.com;

    ssl_certificate   /var/www/cert/domain.com.cert.pem;
    ssl_certificate_key  /var/www/cert/domain.com.key.pem;
    ssl_session_timeout 5m;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_prefer_server_ciphers on;

    client_max_body_size 5m;

    location / {
        proxy_pass http://127.0.0.1:8081;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

域名证书通过[acme](https://github.com/acmesh-official/acme.sh)获取

```bash
// 安装acme
curl https://get.acme.sh | sh -s email=name@doamin.com
alias acme.sh=~/.acme.sh/acme.sh

// 获取证书，这里关闭80端口
acme.sh --issue -d domain.com --standalone

// 将证书放到指定位置
acme.sh --install-cert -d domain.com \
--key-file       /var/www/cert/domain.com.key.pem  \
--fullchain-file /var/www/cert/domain.com.cert.pem 
```