# docker-web
Dockerfileの練習

# ファイル
`Dockerfile`と`main.go`だけ。

# コマンド
イメージのビルド
```
docker build --no-cache -t dockerized-go-app .
```

コンテナの起動
```
docker run -p 8080:8080 --rm dockerized-go-app
```

# その他
testする時、warningっぽいの出る。実行はできたけど。
環境変数変えなあかんかったっぽい。無視でも良さそう。

```
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH
```

```bash
CGO_ENABLED="0" go test -v ./...
```

参考
[記事](https://whatsnewqiita.com/20190417/Go)

# 確認
[http://localhost:8080/](http://localhost:8080/)

