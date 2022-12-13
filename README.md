# chatroom-service

## 簡述

使用 GoLang 打造簡單聊天室功能。

## 安裝 docker 環境及執行程式

* clone GitHub repository
```
$ git https://github.com/yuyuancha/chatroom-service.git
```

* 複製 .env.example 為 .env

```
$ cp .env.example .env
```

* 複製 env.yml.example 為 env.yml

```
$ cp env.yml.example env.yml
```

* 透過 docker-compose 開啟 docker 容器
```
$ docker-compose up -d
```

* 執行 main.go
```
$ docker-compose exec app go run main.go
```

* 關閉 docker 容器
```
docker-compose down
```
