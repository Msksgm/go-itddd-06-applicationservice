# GO-ITDDD-05-REPOSITORY

zenn の記事「[Go でアプリケーションサービスを実装（「入門ドメイン駆動設計」Chapter6）](https://zenn.dev/msksgm/articles/20220422-go-itddd-06-applicationservice)」のサンプルコードです。

# 実行環境

- Go
  - 1.18
- docker compose

# 実行方法

## コンテナを起動・マイグレーション

コンテナの起動

```bash:コンテナの起動
> make up
docker compose up -d
# 完了までまつ
```

```bash:マイグレーション
> make run-migration
docker compose exec app bash db-migration.sh
1/u user (10.031ms)
```

## 実行

Register の 1 回目の実行

```bash:Register（1 回目）
> make run usecase=register
docker compose exec app go run main.go -usecase=register
2022/04/21 23:17:11 successfully connected to database
2022/04/21 23:17:11 register
2022/04/21 23:17:11 user name of test-user is successfully saved
```

Register の 2 回目の実行

```bash:Register（2 回目）
> make run usecase=register
docker compose exec app go run main.go -usecase=register
2022/04/21 23:19:07 successfully connected to database
2022/04/21 23:19:07 register
2022/04/21 23:19:07 userapplicationservice.Register err: user name of test-user is already exists.
```

Get の 1 回目の実行

```bash:Get（1 回目）
> make run usecase=get
docker compose exec app go run main.go -usecase=get
2022/04/21 23:20:24 successfully connected to database
2022/04/21 23:20:24 get
2022/04/21 23:20:24 &{test-id test-user}
```

Update の実行

```bash:Update
> make run usecase=update
docker compose exec app go run main.go -usecase=update
2022/04/21 23:21:26 successfully connected to database
2022/04/21 23:21:26 update
2022/04/21 23:21:26 successfully updated
```

Get の 2 回目の実行

```bash:Get（2回目）
> make run usecase=get
docker compose exec app go run main.go -usecase=get
2022/04/21 23:22:24 successfully connected to database
2022/04/21 23:22:24 get
2022/04/21 23:22:24 &{test-id test-updated-user}
```

Delete の実行

```bash:Delete
> make run usecase=delete
docker compose exec app go run main.go -usecase=delete
2022/04/21 23:23:52 successfully connected to database
2022/04/21 23:23:52 delete
2022/04/21 23:23:52 successfully deleted
```

# テスト

```bash
> make test
docker compose exec app go test ./...
?       github.com/msksgm/go-itddd-06-applicationservice        [no test files]
ok      github.com/msksgm/go-itddd-06-applicationservice/domain/model/user      0.003s
```
