# documents-api

## セットアップ
### docker-compose
コンテナ立ち上げ
`docker-compose up`

コンテナ削除コマンド
`docker-compose down`

### MySQL
接続
`mysql -uuser -ppassword`
データベース指定（接続成功後）
`USE documents_api;`

### Postman
### Go
version 1.14

起動コマンド
`go run ./cmd/main.go`

## チーム開発の進め方

### タスクの流れ
1. issueにタスクを追加。名前は`001 DB立ち上げ`のような形で３桁の数字を連番で記入。
2. ローカルPCでブランチを作成。ブランチ名は`feature/001-init-db`のようにしてissueの数字の後に作業内容を記入。
3. ローカルでタスクが完了したら`git push origin HEAD`でpush。
4. プルリクエストを作成し、slackにプルリクエストのURLを貼る。
5. レビュー完了後マージ。
6. `develop`ブランチに移動し、`git pull origin develop`
