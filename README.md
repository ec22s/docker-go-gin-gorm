# docker-go-gin-gorm

https://qiita.com/fujifuji1414/items/b95d3f0d5f79d77360cb と

https://qiita.com/sedori/items/840e39a0cbf9d5bff006 を元にした、

Dockerで完結する認証付きWeb APIの開発テンプレートです

- 構成物 : MySQL + Go + Gin + GORM + JWT + Air (ホットリロード)

- 独自に足したもの : Makefileでの便利コマンド

- 各ソフトウェア・ライブラリのバージョンは基本 `latest`

- TODO: 認証用ミドルウェア作成、トークン利用のテスト追加

<br>

連絡等は[プロフィール](https://github.com/ec22s)記載のe-mailへお願いします

<br>

## 動作確認環境

- macOS 15.6 (24G84)

- GNU bash, version 5.3.3(1)-release (x86_64-apple-darwin23.6.0)

- Docker version 29.0.0, build 3d4129b9ea

- Docker Compse version v2.40.3-desktop.1

<br>

## 使い方

- 基本、リポジトリトップから各種Makeコマンド実行で完結します

- `make up` コンテナのビルドと起動

  - MySQLのテーブル作成はGinの初回起動時にGORMがします

- `make test-root` Webサーバの稼動チェック

  - `404 page not found` が返れば正常

  - 注 : コンテナ起動直後だとWebサーバが準備中でエラーになります

    - TODO: Webサーバ起動が完了したら自動実行させたい

- `make tables` MySQLのユーザテーブルの存在確認

  - `users` が表示されれば正常

  - 注 : コンテナ起動直後だとGORMのマイグレーション中でエラーになります

    - TODO: マイグレーションが完了したら自動実行させたい

- `make test-register` ユーザ登録用

  - ユーザ名とパスワードはMakefileの先頭あたりで定義しており、自由に変えられます

    ```
    username=
    password=
    ```

  - 成功するとMySQLの `users` テーブルにレコードが追加されます

  - レコード確認は `make table-rows`

  - `make err-register-1` `make err-register-2`  はエラーケースのテスト

- `make test-login` ログイン & JWT確認用

  - 成功するとJWTが返ります

    ```
    {"token":"***************************"}
    ```

  - `make err-login-1` `make err-login-2`  はエラーケースのテスト

- その他若干のコマンドがMakefileにあります

<br>

## 設定

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

    - `/Makefile` → 冒頭あたりの変数名 `fqdn_host`

- MySQL

  - ユーザ名とパスワードは下記3箇所に記載

    - `/compose.yml` → `MYSQL_USER` `MYSQL_PASSWORD`

    - `/app/.env` → `DB_USER` `DB_PASS`

    - `/mysql_secret.txt` コマンドにパスワードを含めない用

    - TODO: できれば一箇所にまとめたい

  - ホスト名 (=Docker Composeのサービス名) は下記3箇所に記載

    - 初期値 `mysql`

    - `/compose.yml`

    - `/app.env` → `DB_HOST`

    - `/Makefile` → 冒頭あたりの変数名 `db_command`

  - ポート番号はデフォルト (変更する場合は下記2ファイルに要反映)

    - `/compose.yml`

    - `/app.env` → `DB_PORT`

  - データベース名は下記3箇所に記載

    - 初期値 `first`

    - `/compose.yml` → `MYSQL_DATABASE`

    - `/app.env` → `DB_NAME`

    - `/Makefile` → 冒頭あたりの変数名 `db_command`

以上
