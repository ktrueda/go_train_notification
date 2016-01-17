# 関東の電車の運行情報をSlackに通知するツール

定期的に、電車の運行情報をSlackに垂れ流します。


## 使い方(設定ファイル)

config.json

```json
{
  "Trains": ["東急池上線", "山手線"],
  "Slack" : {
    "Token": "xoxp-xxxxxxxxx-xxxxxxxxx-xxxxxxxx-xxxxxxxx",
    "Channel": "xxxxxxxx"
  }
}

```

## cronの設定

```sh
crontab -e
0 8-23/2 * * * path/to/file
```
