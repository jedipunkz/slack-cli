# Slack-Cli

[![Build Status](https://travis-ci.org/jedipunkz/slack-cli.svg?branch=master)](https://travis-ci.org/jedipunkz/slack-cli)

## Description

Golang 勉強用のレポジトリです。Slack にメッセージを送る CLI を Golang で作りました。

## Installation

```bash
go get github.com/jedipunkz/slack-cli
```

## Usage

### メッセージを送る

```bash
slack-cli msg -c <チャンネル名> 'メッセージをここに書く'
```

### コマンドの実行結果を Slack に送る

```bash
slack-cli exec -c <チャンネル名> 'コマンド名'
ex) slack-cli exec -c botchannel 'uname -a'
```

### ユーザ ID に従って情報を表示

```bash
slack-cli user <USER_ID>
```

### ユーザ一覧を取得し表示

```bash
slack-cli get users
```

### チャンネル一覧を取得し表示

```bash
slack-cli get channels
```

### デーモンで起動し Slack の特定の文字列に反応する

```bash
slack-cli listen
```
