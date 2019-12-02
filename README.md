# Slack-Cli

## Description

Golang 勉強用のレポジトリです。Slack にメッセージを送る CLI を Golang で作りました。

## Installation

```bash
go get github.com/jedipunkz/slack-cli
go build
```

## Usage

### メッセージを送る

```bash
slack-cli msg <チャンネル名> 'メッセージをここに書く'
```

### コマンドの実行結果を Slack に送る

```bash
slack-cli exec <チャンネル名> 'コマンド名'
ex) slack-cli exec botchannel 'uname -a'
```
