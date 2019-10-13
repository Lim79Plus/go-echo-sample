# このアプリについて

- goを使ったWebアプリです。
‐ FWにはechoを使用しています。

## 起動方法

docker-compose.ymlのあるディレクトリで以下を実行する。
docker imageがビルドされて、バックグラウンドで実行されます。

```:docker-cmd
docker-compose up -d
```

docker for desktopの場合

```:cmd
curl http://192.168.99.100:1323/hello
```

上記以外の場合

```:cmd
curl http://loalhost:1323/hello
```
