# data-platform-api-post-creates-rmq-kube
data-platform-api-post-creates-rmq-kube は、周辺システム　を データ連携基盤 と統合することを目的に、API で投稿取引データを登録するマイクロサービスです。

* https://xxx.xxx.io/api/API_POST_SRV/creates/

## 動作環境
data-platform-api-post-creates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## 本レポジトリ が 対応する API サービス
data-platform-api-post-creates-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス URL: https://xxx.xxx.io/api/API_POST_SRV/creates/

## 本レポジトリ に 含まれる API名
data-platform-api-post-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_Header（投稿 - ヘッダ）
* A_Friend（投稿 - フレンド）

## API への 値入力条件 の 初期値
data-platform-api-post-creates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール
Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMPostCreates",
	"accepter": ["Header"],
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMPostCreates",
	"accepter": ["All"],
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 投稿 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"Post" ～ "IsMarkedForDeletion" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
```
