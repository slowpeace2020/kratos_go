

# 进入docker mysql镜像，mysql-realword-db-1为镜像名
docker exec -it mysql-realword-db-1 sh
```

#启动数据库
cd deploy/mysql
docker-compose up -d