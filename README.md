[English](README_en.md)
# HPE OneView Event Logger
HPE OneView Event Loggerは単純にOneViewのイベントをAPI経由で取得してログファイルに書き出します。


## クイックスタート
[バイナリ](https://github.com/fideltak/oneview-event-logger/releases)または[コンテナイメージ](https://hub.docker.com/repository/docker/fideltak/oneview-event-logger)を実行することで使用できます。各種パラメータは環境変数で設定してください。

```
# tar xvfz oneview-event-logger-<VERSION>-<OS>-amd64.tar.gz 
# OV_ADDR=192.168.2.6 OV_USER=golang OV_PASSWORD=golangtest ./oneview-event-logger
```

数分するとOneViewイベントがログファイルに書き出されます。

```
# cat /tmp/oneview_evnet.log
2021/04/02 15:09:55 OneView:192.168.2.6 Created:2021-04-01T10:43:00.238Z Severity:OK Category:drive-enclosure Desc:"Drive removed from drive bay 17."
2021/04/02 15:09:55 OneView:192.168.2.6 Created:2021-04-01T10:42:58.042Z Severity:Critical Category:drive-enclosures Desc:"Unable to communicate with the drive in drive bay 11. "
2021/04/02 15:09:55 OneView:192.168.2.6 Created:2021-04-01T10:42:50.352Z Severity:OK Category:drive-enclosure Desc:"Drive inserted into drive bay 37."
2021/04/02 15:09:55 OneView:192.168.2.6 Created:2021-04-01T10:42:49.441Z Severity:OK Category:drive-enclosure Desc:"Drive inserted into drive bay 39."
```

## パラメータ
OSの環境変数に以下のパラメータを設定できます。  

| Key | Default | Description |
| :---: | :---: | :---: |
| OV\_INTERVAL | 60 | イベントをスキャンする間隔 |
| OV\_ADDR |  | OneViewのIPアドレスまたはホスト名 |
| OV\_USER |  | OneViewのユーザー名|
| OV\_PASSWORD |  | OneViewユーザーのパスワード|
| OV\_VERSION | 1200 | OneViewのAPIバージョン|
| OV\_LOG\_PATH | /tmp/oneview_evnet.log | ログファイルのパス|
| OV\_LOG\_MAX\_SIZE\_MB | 50 | 何MBになったらログローテンションさせるか|
| OV\_LOG\_MAX\_BACKUPS | 5 | ログローテンション後何世代残すか|
| OV\_LOG\_MAX\_AGE | 365 | ログを何日間保存するか| 
| OV\_LOG\_COMPRESS | true | 古い世代のログは圧縮するか |

## Wiz Zabbix
Zabbixと共につかうことができます。本ツールを作った実際の理由はZabbix用につくりました。  

Zabbixでは外部スクリプトやJavascriptでOneView APIからイベントを収集できますが、Zabbixは複数のイベントを１つのエントリーに挿入しました。監視システムとしてあまり見栄えが良くありませんでした。(ZabbixでJSONリストを分けてイベントとして保存する方法があるかをいまだに探しています。)  

そのため、Zabbix agentとこのツールを使ってOneViewのイベントを監視する方法を思いつきました。OneViewからのイベントはZabbix上でそれぞれのインシデントとして登録されます。[こちら](https://github.com/fideltak/zabbix_oneview_sample)を参照すれば、私がこのツールを作った理由がわかると思います。

以下の例のように、k8s上で本ツールとZabbix-zgentを統合することをお勧めします。k8s用の[サンプルマニフェストは](deploy/k8s/wiz_zabbix_agent)にあります。
  
また、[コンテナイメージはこちら](https://hub.docker.com/repository/docker/fideltak/oneview-event-logger)にあります。

```
┌────────────────────────────────────────────────────────────────┐
│ Pod                                                            │
│ ┌───────────────────────────┐   ┌────────────────────────────┐ │
│ │<Container>                │   │<Container>                 │ │
│ │zabbix-agent               │   │oneview-event-logger        │ │
│ │                           │   │                            │ │
│ │                           │   │                            │ │
│ └───────────────┬───────────┘   └────┬───────────────────────┘ │
│                 │                    │                         │
│           ┌─────▼────────────────────▼──────────────┐          │
│           │<Persistent Volume:RWX>                  │          │
│           │ /var/log/oneview/events.log             │          │
│           │                                         │          │
│           └─────────────────────────────────────────┘          │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

Zabbix上では以下のようにOneViewのイベントを見ることができます。
以下の例では`log[/var/log/oneview/events.log, "Critical"]`と設定して*Critical*なOneViewのイベントのみを監視しています。
![oneview-critical-events](docs/zabbix/oneview-critical-events.png)
