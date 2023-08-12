# 概要
引数、または標準入力で与えられた文字列内の`:<emoji alias>:`を`UTF-8`の絵文字に変換する。

# 開発環境
```
$ go version
go version go1.18.1 linux/amd64
```

# 使い方
## 引数で与える場合
```bash
$ emoji-converter :smile:
😄
```
## 標準入力で与える場合
```bash
$ echo ":smile:" | emoji-converter
😄

```