# 環境構築
````
$ go env  
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate
````

### Windows版 Scoop
PowerShellプロンプト  
````
PS: Set-ExecutionPolicy RemoteSigned -Scope CurrentUser -Force  
PS: iwr -useb get.scoop.sh | iex  
````

### migrateのインストール
scoop install migrate

### With Go toolchain
Versioned  
````
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate  
$ cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate  
$ git checkout $TAG  # e.g. v4.1.0  
- ### Go 1.15 and below  
$ go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate  
- ### Go 1.16+  
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$TAG  
````

### Gorm関連（DB接続）
v1
````
$ go get github.com/jinzhu/gorm
$ go get github.com/jinzhu/gorm/dialects/mysql
````

### マイグレーションを実行してみる
プロジェクトのルートにいることを想定
````
migrate -database 'mysql://root:パスワード@tcp(127.0.0.1:3306)/go_sample' -path ./migrations/users_table up
````

### test用のライブラリのインストール
https://github.com/stretchr/testify
````
$ go get github.com/stretchr/testify
````
### go-redis install
https://onemuri.space/note/vo6tcv8fq/  
````
$ go get github.com/gomodule/redigo/redis
````