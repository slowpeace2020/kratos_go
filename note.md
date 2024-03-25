
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

```
config变量配置过程：
首先在config.yaml中定义变量，配置环境变量
在conf.proto中增加全局的变量比如用户token
make config生成配套的文件

```

进度：3/6 05 3:00
进度：3/13 05 34:47
进度：3/14 06 19:03
