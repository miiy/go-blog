
## create database

```sql
CREATE DATABASE `up` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
```

sudo docker run --name mysql -v /home/debian/Documents/data/mysql/data:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8
sudo docker run --name redis -v /home/debian/Documents/data/redis/data:/data -p 6379:6379 -d redis:6