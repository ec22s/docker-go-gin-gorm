# docker-go-gin-gorm

https://qiita.com/fujifuji1414/items/b95d3f0d5f79d77360cb と

https://qiita.com/sedori/items/840e39a0cbf9d5bff006 を利用・参照させてもらった、

Dockerで完結する認証付きWeb APIの開発テンプレートです

- 構成物 : MySQL + Go + Gin + GORM + JWT + Air (ホットリロード用)

- 独自に足したもの : Makefileでの便利コマンド

- 各ソフトウェア・ライブラリのバージョンは基本 `latest`

連絡等は[プロフィール](https://github.com/ec22s)記載のe-mailへお願いします

<br>

## 動作確認環境 (2025年12月)

- macOS 15.6 (24G84)

- GNU bash, version 5.3.3(1)-release (x86_64-apple-darwin23.6.0)

- Docker version 29.0.0, build 3d4129b9ea

- Docker Compse version v2.40.3-desktop.1

<br>

## 最小限の使い方

- カレントはリポジトリのルートでいいです

- `make up` を実行します

  - コンテナのビルドと起動が行われます

  - MySQLのテーブル作成はGinの初回起動時にGORMがします

- APIとDBの準備を少し待ちます

  - `make test-root` が `404 page not found` を返し、`make tables` の結果に `users` が出れば準備完了です

- ユーザ名とパスワードをMakefileの冒頭に記入します

  ```
  username=yourname
  password=yoursecret
  ```

- `make test-register` を実行します (ユーザ登録)

  - 正常なら以下のように成功レスポンスが返ります (パスワードは空白)

    ```
    HTTP/1.1 200 OK
    ...
    {"data": ... }
    ```

- `make test-login` を実行します (ログインとJWT取得)

  - 正常なら以下のようにJWTが返ります

    ```
    HTTP/1.1 200 OK
    ...
    {"token": ... }
    ```

- JWTをMakefileの `token=` に記入します

- `make test-api` を実行します (JWTでAPIコール)

  - 正常ならユーザ登録時と同じレスポンスが返ります

    ```
    HTTP/1.1 200 OK
    ...
    {"data": ... }
    ```

<br>

## 便利コマンド

- `make clean-restart` 全て破棄してやり直し

- `make logs-go` APIサーバのログ表示

  - 正常な起動直後の様子が下記です

  - 中ほど `building...` で少し時間がかかります

    ```
    go-1  |
    go-1  |   __    _   ___
    go-1  |  / /\  | | | |_)
    go-1  | /_/--\ |_| |_| \_ v1.63.4, built with Go go1.25.5
    go-1  |
    go-1  | watching .
    go-1  | watching controllers
    go-1  | watching libraries
    go-1  | watching middlewares
    go-1  | watching models
    go-1  | watching utils
    go-1  | watching utils/token
    go-1  | building...
    go-1  | running...
    go-1  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
    go-1  |
    go-1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    go-1  |  - using env:	export GIN_MODE=release
    go-1  |  - using code:	gin.SetMode(gin.ReleaseMode)
    go-1  |
    go-1  | [GIN-debug] POST   /api/register             --> go_gin_gorm/controllers.Register (3 handlers)
    go-1  | [GIN-debug] POST   /api/login                --> go_gin_gorm/controllers.Login (3 handlers)
    go-1  | [GIN-debug] GET    /api/admin/user           --> go_gin_gorm/controllers.CurrentUser (4 handlers)
    go-1  | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
    go-1  | Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.
    go-1  | [GIN-debug] Listening and serving HTTP on :80
    ```

- `make test-root` APIサーバの稼動チェック

  - `404 page not found` が返れば正常

- `make tables` MySQLのユーザテーブルの存在確認

  - `users` が表示されれば正常

- `make table-rows` MySQLのユーザテーブルのレコード確認

- `make err-register-1` `make err-register-2`  ユーザ登録失敗のテスト

- `make err-login-1` `make err-login-2` ログイン失敗テスト

- `make err-api-1` `make err-api-2` APIコール時の認証失敗テスト

  - Makefileにある `dummy_token` を使い、`401 Unauthorized` が返ります

- その他若干のコマンドがMakefileにあります

<br>

## 設定

- 同じ項目で複数ある記述箇所はまとめたいですが、とりあえず

- JWT

  - `/app/.env` に2つ設定あり

    - トークンの有効期限 `TOKEN_HOUR_LIFESPAN`

    - 署名に使用する秘密鍵 `API_SECRET`


- Gin

  - コンテナ内の待ち受けポート番号は下記2箇所に記載

    - 初期値 `80`

    - `/app/main.go`

    - `/compose.yml`

  - ホスト側ポート番号は下記2箇所に記載

    - 初期値 `8008`

    - `/compose.yml`

    - `/Makefile` → `HOST_PORT`

- MySQL

  - ユーザ名とパスワードは下記3箇所に記載

    - 初期値はともに `dev`

    - `/compose.yml` → `MYSQL_USER` `MYSQL_PASSWORD`

    - `/app/.env` → `DB_USER` `DB_PASS`

    - `/mysql_secret.txt` コマンドにパスワードを含めない用

  - ホスト名 (=Docker Composeのサービス名) は下記3箇所に記載

    - 初期値 `mysql`

    - `/compose.yml`

    - `/app.env` → `DB_HOST`

    - `/Makefile` → `DB_HOST`

  - ポート番号はデフォルト (変更する場合は下記2ファイルに要反映)

    - `/compose.yml`

    - `/app.env` → `DB_PORT`

  - データベース名は下記3箇所に記載

    - 初期値 `first`

    - `/compose.yml` → `MYSQL_DATABASE`

    - `/app.env` → `DB_NAME`

    - `/Makefile` → `DB_NAME`

<br>

## TODO

- コンテナ起動時、Airのbuildingが終わるまで待たせたい

- 同じ設定項目の記述箇所をまとめたい

<br>

以上
