# senryu-user
## 概要
川柳を共有することができるSNSアプリのユーザーデータを管理しているバックエンド

## 技術
|名前|備考|
|--|--|
|Go kit(Golang)|バックエンド|
|MongoDB|データベース|
|Docker|開発環境|

## Run
### Docker Compose
ルートディレクトリで`docker-compose up`  

## Use
```
// ログイン
curl -XPOST -d'{"username": "test", "password": "test"}' localhost:8080/login

// ユーザー登録
curl -XPOST -d'{"username": "test2", "email": "test2@test.com",  "password": "test2"}' localhost:8080/register
```

## 開発手法（VSCode Remote Container）
`docker-compose.yml`の`command: fresh -c .fresh.conf`をコメントアウトして、
VSCode Remote Containerで実行する。  
コンテナ内で`fresh -c .fresh.conf`を実行するとホットリロードで開発ができる。
