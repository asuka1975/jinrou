# jinrou

## 人狼について

夜と昼がある

夜の各プレイヤーが行動

昼の最初か最後で勝利条件判定，話し合った後投票

## プログラム全体のアーキテクチャ
 - 人狼モデル
 - 人狼APIサーバ
 - Webアプリケーション
 
 ![architecture](./READMEimg/architecture.png)


## ゲームの流れ

### 人狼モデル
人狼ゲームの内部の動作を規定する。ログイン画面や役職確認などUIやゲームの本質以外の部分は規定しない。
 - Session 
   - Morning: 特に何もしない
   - Noon: 特に何もしない
   - Evening: 投票
   - Night: 夜の行動（人狼は誰を殺すか、占いは誰を占うかなど）

### 人狼APIサーバ
人狼モデルを内部に持ちWebアプリとの協調を行う。
WebアプリからのリクエストからJSON形式のメッセージを作成しWebアプリへ返す。
ここでログインや役職決定などの人狼モデルで規定しない部分を規定する。
Morning Sessionにおける死んだ人間の確認やNoonセッションにおける会話、時間の管理などもここで行う。

### Webアプリ
UIに専念。人狼APIサーバから受け取ったJSONデータを元にUIの更新を行う。

### その他情報
- ログイン画面
- 参加時の役職確認

- 夜の活動で入力

  ユーザの属性ごとに違うHTML，夜のユーザの行動が終わったら同じHTMLを配信

- 投票機能

  今いるユーザの分だけボタンを用意してそれを押す．



## 必要なクラス(人狼モデル)
- Jinrou 
  > 人狼モデル。プレイヤーや「人狼における」セッションの管理を行う。
- Session 

  - 朝Session : 人狼モデルでは何もしない
  - 昼Session : 人狼モデルでは何もしない
  - 夕方Session : 投票
  - 夜Session : 各役が各々行動

- Player 名前, 役割, プレイヤーの生死

- Role 役名，
  - 各役のClass
    - GetName　：　役名
    - GetAction ：　役固有の行動

- IActiveCommand 夜の行動（占う等）
  - NoneCommand : 何もしない
  - KillCommand : 他のプレイヤーを殺す
  - ProtectCommand : 他のプレイヤーを守る
  - PredictCommand : 他のプレイヤーを占う

- PassiveCommand 夜の行動に対するリアクション（Knightに守られる等）
  - Cancel : ActiveCommandをキャンセルするか
  - Execute : コマンドの実行
  
> 生死の状態はPlayerが管理．


