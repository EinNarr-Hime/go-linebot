# Go-Linebot

単純に送信したメッセージに対して同様の内容をおうむ返しするだけの line-bot です。

## 手順

1. [Line デベロッパーページ](https://developers.line.biz/)アプリケーションを作成 & 登録する。

2. ルートディレクトリに`.env`ファイルを作成して必要事項を追記する。

```sh
# webサーバーがlistenするポート番号
PORT=2525
# アクセストークン
ACCESS_TOKEN=ABCDEFG-123456789
# アクセスシークレット
ACESS_SECRET=HIJKLMNOP-987654321
```

3. 任意のデブロイ先にデプロイして動作確認する。
