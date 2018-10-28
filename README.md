# wildcat-server

## 実行環境

- dep
- golang1.1.4
- gin
- SQLite3
- nginx

## 外部API

- Twitter REST API

## 実行方法

```
$ dep ensure // install dependency
$ go run migrations/main.go // migration
$ go run tweet_img_getter/main.go // get data from twitter
$ go run main.go // activate server
```

## こだわった点
デザインアーキテクチャとしてMVCを採用しました
今回の場合APIサーバーのため利用者側から見たモデルのインターフェースとして
ServiceというものをVの部分においたためMVSかもしれません。

Router -> Controller -> Model -> Controller -> Service

の流れになります。


