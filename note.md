
```
# 进入docker mysql镜像，mysql-realword-db-1为镜像名
docker exec -it mysql-realword-db-1 sh
```

```
# 启动数据库
cd deploy/mysql
docker-compose up -d
```

```
# 测试
到test所在目录
go test .

```
进度：3/6 05 3:00
进度：3/13 05 34:47
进度：3/14 06 19:03
