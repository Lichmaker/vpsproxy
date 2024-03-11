# vpsproxy

vpsproxy 可以提供简单的HTTPS和gRPC反向代理，并且gRPC可设置带secret的校验。


# 构建程序

```
go build -o vpsproxy main.go
```

# 配置
```
cp settings.yaml.example settings.yaml
```

```
# settings.yaml


app:
    host: 0.0.0.0
    # 对外提供gRPC管理服务的端口
    port: 10099
    debug: true
    env: dev
    name: vpsproxy
    timezone: Asia/Shanghai
log:
    compress: true
    encoder: console
    filename: logs/vpsproxy.log
    level: debug
    max_age: 30
    max_backup: 5
    max_size: 64
proxy:
    https:
        # 证书和私钥的路径，为空时则仅进行HTTP代理
        cert: ""
        key: ""
        # HTTPS反向代理的端口
        port: 12000
        # 代理目标地址
        target: http://127.0.0.1:12315
    grpc:
        # gRPC的反向代理端口
        port: 12010
        # 要求从metadata中带有的密钥，不正确时拒绝访问
        secret: qwert
        # gRPC代理目标地址
        target: 127.0.0.1:10010

```

# 安装并启动
```
sudo ./vpsproxy install

sudo systemctl enable vpsproxy
sudo systemctl start vpsproxy
sudo systemctl status vpsproxy
sudo systemctl stop vpsproxy
```