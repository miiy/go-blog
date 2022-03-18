
## create database

```sql
CREATE DATABASE `up` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
```

sudo docker run --name mysql -v /home/debian/Documents/data/mysql/data:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8

sudo docker run --name redis -v /home/debian/Documents/data/redis/data:/data -p 6379:6379 -d redis:6


## 安装 protoc

https://github.com/protocolbuffers/protobuf

## 安装插件

```bash
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/envoyproxy/protoc-gen-validate
```

protoc 生成器插件会安装到 $GOPATH/bin 目录

## 克隆到本地

https://github.com/googleapis/googleapis.git