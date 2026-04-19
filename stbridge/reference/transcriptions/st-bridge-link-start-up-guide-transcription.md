---
type: source
title: ST-Bridge Link Start Up Guide PDF転記
tags: [pdf-transcription, st-bridge, revit, workflow]
related: []
created: 2026-04-19
updated: 2026-04-19
authors: []
year: 2021
url: "/Users/keikamikawa/go/src/github.com/Code-Hex/bim/st-bridge/ST_Bridge_Link_start_up_guilde.pdf"
venue: "local PDF"
---

# ST-Bridge Link Start Up Guide PDF転記

> 機械抽出による Markdown 転記。代表ページは vision で目視確認済み。表組みは PDF 上の罫線や段組みを完全には保持していない。

- 元 PDF: `/Users/keikamikawa/go/src/github.com/Code-Hex/bim/st-bridge/ST_Bridge_Link_start_up_guilde.pdf`
- ページ数: 15

## Page 1

ST-Bridge Link 
スタートアップガイド 
 
 
 
 
 
 
 
 
 
 
 
 
2021 年12 月6 日 
Ver.1.0.0

## Page 2

ST-Bridge Link Start-Up Guide 
 
 
目次 
ST-Bridge Link スタートアップガイドについて .................................................................... 1 
1. ST-Bridge Link の概要 ....................................................................................................... 2 
1.1 Revit モデルと解析モデル ............................................................................................ 2 
1.2 ST-Bridge ファイルとST-Bridge Link .......................................................................... 2 
1.3 Revit での構造設計ワークフロー ................................................................................. 3 
2. ST-Bridge Link で変換 ....................................................................................................... 4 
2.1 ST-Bridge Link の入手 .................................................................................................. 4 
2.2 ST-Bridge Link のインストール .................................................................................... 5 
2.3 テンプレートとマッピングテーブルの確認 ................................................................. 6 
2.4 ST-Bridge Link の実行 .................................................................................................. 7 
2.5 差分インポートについて .............................................................................................12 
3. 変換できないとき .............................................................................................................13

## Page 3

ST-Bridge Link Start-Up Guide                                            1 
 
 
ST-Bridge Link スタートアップガイドについて 
本スタートアップガイド資料は，建築の構造設計分野においてAutodesk Revit（以下Revit）
を使い始める方に向けて，全体のワークフローおよびST-Bridge ファイルを用いた構造計算
ソフトとの連携方法を紹介するものです。Revit そのものの操作法等について解説された，
Revit の構造設計向けトレーニングテキストと合わせてご利用ください。 
 
本ガイドでは，上記トレーニングテキストで解説されていますRevit のUI 等の説明につい
ては記載していませんので，予めご了承ください。

## Page 4

2                                           ST-Bridge Link Start-Up Guide 
 
 
1.  ST-Bridge Link の概要 
1.1  Revit モデルと解析モデル 
Revit は，建物データベースとしてのBIM モデルを作成できるツールです。さらに，作成
したBIM モデルから，様々な2 次元図面を出力可能です。出力された図面とBIM モデルは
連動しています。 
 
一方で，構造設計者は，構造解析用の建物データベースとして，構造解析モデルを既に使
用しているケースが多いです。構造解析モデルのもつ建物情報と，構造BIM モデル，構造図
に必要な建物情報は共通点が非常に多く，個別で作成すると，二重の管理の手間や整合性チ
ェックの必要が生じます。構造設計におけるBIM 活用では，これらのBIM モデルと解析モ
デル間の情報連携を実現させることが重要です。 
 
1.2  ST-Bridge ファイルとST-Bridge Link 
オートデスクでは，構造解析ソフトとの連携用の日本仕様Revit アドインとして，ST-
Bridge Link をサブスクリプションユーザー向けに無料でリリースしています。ST-Bridge 
Link は，一般社団法人buildingSMART Japan による，日本の一貫構造計算ソフトの連携に重
きを置いた中間ファイル形式である，ST-Bridge ファイルをRevit にインポートおよびエク
スポートできるアドインです。 
 
 
ST-Bridge ファイルは，BIM モデル向け中間ファイルとして広く世界で用いられている
IFC ファイルと異なり，日本の建築構造分野での情報交換を目的に定義されたファイルのた
め，部材符号や配筋情報といった日本での構造分野特有の情報にも対応しています。ST-
Bridgeファイルの詳細な仕様は，buildingSMART JapanのHPで公開されている「ST-Bridge 
XML ファイル仕様書（ver.2.0）」をご参照ください。 
 
このST-Bridge ファイルを中間ファイルとすることで，一貫構造計算ソフト等の多くの構
造解析ソフトとRevit のデータを連携することができます。 
 
 
解析モデル
3D
構造図
2D
BIMモデル
3D
ST-Bridge
中間ファイル
ST-Bridge Link

## Page 5

ST-Bridge Link Start-Up Guide                                            3 
 
 
1.3  Revit での構造設計ワークフロー 
Revit の構造設計向けトレーニングテキストでは，Revit で０から構造BIM モデルを作成し，
その構造BIM モデルから，各構造図を出力しています。このテキストの手順は，Revit での
構造設計向けの操作を習得するのに非常に有用ですので，一度トレーニングとして実施する
ことをお勧めします。 
 
しかし，現在の実務上のワークフローでは，まず一貫構造計算ソフトで解析モデルを作成
し，ST-Bridge Link でRevit の構造BIM モデルに変換した後に，Revit で構造図を作成するこ
とが多いです。これは，ST-Bridge ファイルの出力には，ほとんどの一貫構造計算ソフトが
対応していますが，ST-Bridge ファイルの読込みに対応した一貫構造計算ソフトが非常に少
ないためです。最初に０からRevit で構造BIM モデルを作成しても，それを解析モデルとし
て一貫構造計算ソフトに読み込めず，別途計算ソフト側で解析モデルを作成しなければなり
ません。 
現状では，下図のように解析モデルを変換してBIM モデルを作成し，BIM モデルから構造
図を作成するフローを推奨しています。 
 
 
 
ただし，主要部材は解析モデルからの変換でBIM モデルを作成しますが，解析モデルに通
常含まれない二次部材に関してはRevit 側でモデルを作成する必要があります。また，BIM
モデル作成後は，構造図をRevit で出力する必要があります。これらの手順は上述のRevit の
構造設計向けトレーニングテキストに詳細がありますので，そちらをご参照ください。 
 
また，構造システム社様の「BUS-6 + Revit Op.」では，Revit のBIM モデルとBUS-6 の解
析モデルの双方向連携が既に実現していますので，Revit からモデルを作成しても，BUS-6
からモデルを作成しても相互に連携が可能になっています。詳細については，弊社と構造シ
ステム社様の共同ウェビナーの動画をご覧ください。 
 
 
軸組図
2D
解析モデル
3D
ST-Bridge
中間ファイル
ST-Bridge Link
リスト
2D
伏図
2D
BIMモデル
3D
株式会社構造ソフト
ユニオンシステム株式会社
株式会社NTTファシリティーズ総合研究所
株式会社構造システム
×

## Page 6

4                                           ST-Bridge Link Start-Up Guide 
 
 
2.  ST-Bridge Link で変換 
2.1  ST-Bridge Link の入手 
ST-Bridge Link は，Autodesk App Store にて公開されていますが，その入手方法はRevit
のバージョンによって異なります。また，ST-Bridge Link の使用には，Mapping Table も必
要になりますが，こちらもバージョンにより入手方法が異なります。 
 
Revit2022・Revit2021 
構造設計向け日本仕様アドオン「Revit Extension for Structure Japan 20XX」内に，ST-
Bridge Link およびMapping Table が含まれています。使用しているバージョンに応じたアプ
リをDL してください。これらのアプリ内には，断面リスト作成機能等の日本仕様の図面化
補助ツールも含まれています。 
 
Revit Extension for Structure Japan 2022 
 
Revit Extension for Structure Japan 2021 
 
Revit2020・Revit2019 
「ST-Bridge Link 20XX」および「Mapping Table 20XX」のそれぞれが必要となります。
使用しているバージョンに応じたアプリをDL してください。断面リスト作成機能等の日本
仕様の図面化補助ツールは，別アプリとなっています。Revit2020 用は「Rex J for Structure 
2020」，Revit2019 用は意匠向けツールと一体化された「Japan Standard Extension for 
2019」をご利用ください。 
 
ST-Bridge Link 2020 
Mapping Table 2020 
Rex J for Structure 2020 
 
ST-Bridge Link 2019 
Mapping Table 2019 
Japan Standard Extension for 2019 
 
Revit2018 
「ST-Bridge Link 2018」内にMapping Table が同梱されています。断面リスト作成機能等
の日本仕様の図面化補助ツールは，それぞれ個別のアプリとなっています。 
 
ST-Bridge Link 2018 
Japan Standard Extension RST 2018 
RC 断面リスト作成2018 
S 断面リスト作成2018 
 
Revit2017 以前 
各アドオンは公開終了しています。新しいバージョンのRevit をご利用ください。

## Page 7

ST-Bridge Link Start-Up Guide                                            5 
 
 
2.2  ST-Bridge Link のインストール 
2.1 節を参照し，ブラウザで必要なアプリのページを開いてください。ここでは，
Revit2021 を例に解説します。各アプリのページにおいて，右側の「ダウンロード」ボタン
より，アプリのインストーラをダウンロードします。 
 
 
 
前節でご案内しました，構造設計向けの各アプリは，サブスクリプション特典となってお
りますので，ダウンロードにはサブスクリプションライセンスが必要です。サブスクリプシ
ョンライセンスのある，Autodesk Account でApp Store にログインしてダウンロードしてく
ださい。 
 
ダウンロードされたインストーラ（MSI ファイル）を実行することで，各アプリをインス
トールできます。インストールは画面のダイアログにしたがってください。なお，インスト
ーラ実行の際に，バージョンにかかわらずRevit が起動しているとインストールができませ
ん。事前にRevit を終了してください。 
 
インストール後にRevit を起動し，「JP 構造」タブ（Revit2019 以前は「REXJ」タブ）が
作成され，各ツールのボタンが作成されたことを確認してください。

## Page 8

6                                           ST-Bridge Link Start-Up Guide 
 
 
2.3  テンプレートとマッピングテーブルの確認 
ST-Bridge Link では，ST-Bridge ファイル内の構造材の情報を読込み，Revit 上で各ファミ
リのパラメータに情報を自動的に書き込み，部材を各位置に自動的にモデリングします。こ
のとき，RC 柱やS 梁等の様々な構造部材に対し，どの部材をどのファミリに読み込むのか，
鉄筋径や本数，鉄骨断面等の様々な情報を，どのパラメータに読み込むのかを指示したもの
がマッピングテーブルです。ST-Bridge Link の実行の際には，マッピングテーブルで指示さ
れた各ファミリが，開いているRevit プロジェクトにロードされている必要があります。 
 
初めて使用する際には，マッピングテーブルの初期状態で指示されたファミリは，ST-
Bridge Link に同梱されたテンプレートにロード済ですので，特にマッピングテーブルの指示
は必要ありません。ご自身で作成したファミリを使用したいといった際には，マッピングテ
ーブルを書き換えることで，自由に使用するファミリを指示することが可能です。 
 
各バージョンのST-Bridge Link をインストールすると，以下の場所に同梱のテンプレート
が保存されます。C ドライブ直下の「ProgramData」フォルダは，Windows の初期設定では
隠しフォルダとなっておりますので，ご注意ください。 
 
Revit2022 
"C:\ProgramData\Autodesk\ApplicationPlugins\REXJ_RST2022.bundle\Contents\ST-Bridge 
Link\SampleData\Structure-Template2022.rte" 
 
Revit2021 
"C:\ProgramData\Autodesk\ApplicationPlugins\REXJ_RST2021.bundle\Contents\ST-Bridge 
Link\SampleData\Structure-Template2021.rte" 
 
Revit2020 
"C:\ProgramData\Autodesk\ApplicationPlugins\ST-Bridge Link 2020.bundle\Contents 
\SampleData\Structure-Template2020.rte" 
 
Revit2019 
"C:\ProgramData\Autodesk\ApplicationPlugins\ST-Bridge Link 2019.bundle\Contents 
\SampleData\Structure-Template2019.rte" 
 
Revit2018 
"C:\ProgramData\Autodesk\ApplicationPlugins\ST-Bridge Link 2018.bundle\Contents 
\SampleData\Structure-Template2018.rte"

## Page 9

ST-Bridge Link Start-Up Guide                                            7 
 
 
2.4  ST-Bridge Link の実行 
前節のST-Bridge Link 同梱のテンプレートを使用し，マッピングテーブルを初期設定から
変更していない場合を例として解説します。 
 
 ST-Bridge Link 同梱のテンプレートで新規プロジェクトを作成します。 
同梱のテンプレートは階層が深い場所にありますので，お好みの場所にコピーして使用
することをお勧めします。 
 
 
 
 
 「JP 構造」タブの「ST-Bridge Link」パネルをクリックします。 
ダイアログに従い，使用するST-Bridge ファイルを選択します。 
 
 
 
ST-Bridge Link 同梱のテンプレートと同フォルダ内には，サンプルのST-Bridge ファイル
が，RC 造およびS 造それぞれ１つずつ同梱されています。拡張子が「stb」のファイルが
ST-Bridge ファイルです。お手元にST-Bridge ファイルがない場合，変換テストには，これ
らのファイルもご利用頂けます。

## Page 10

8                                           ST-Bridge Link Start-Up Guide 
 
 
 変換確認ダイアログ 
 
 
 
ST-Bridge ファイルを読み込むと，使用するファミリの確認のダイアログが表示されます。
たとえば，S 柱角型鋼管の部材は「S_C_Box_1J」というファミリに変換されることがわか
ります。プロジェクト内に，マッピングテーブルで指示されたファミリがロードされていな
い場合は，ここでアラートが出ます。 
ただし，読み込んだST-Bridge ファイルに存在しない部材（たとえば，上図ではSRC 部
材やCFT 部材は存在していません）は，プロジェクト内にロードされていなくても問題あ
りません。 
必要な部材が変換されることを確認し，「次へ」をクリックします。

## Page 11

ST-Bridge Link Start-Up Guide                                            9 
 
 
 レベルマッピングダイアログ 
 
 
 
続いて，レベルマッピングのダイアログが表示されます。ここでは，ST-Bridge ファイル
内のレベルと，Revit プロジェクト内のレベルをマッピングすることが可能ですが，すべて
のレベルを新規作成とすることも可能です。変換する平面位置を指定したい場合は，座標や
通り芯で指定することも可能です。 
変換する位置を確認し，「次へ」をクリックします。

## Page 12

10                                           ST-Bridge Link Start-Up Guide 
 
 
 マテリアルマッピングダイアログ 
 
 
 
続いて，マテリアルマッピングのダイアログが表示されます。ここでは，ST-Bridge ファ
イル内のどの材料が，Revit プロジェクト内のどのマテリアルとして変換されるか，確認変
更できます。問題がなければ，「次へ」をクリックします。 
 
 
 変換の開始 
 
 
 
「はい」をクリックすることで，変換が開始されます。

## Page 13

ST-Bridge Link Start-Up Guide                                            11 
 
 
 変換結果を確認 
S 造のサンプルST-Bridge ファイルを変換した場合，下記のモデルが作成されます。 
（下図は，ビューテンプレートで「99_部材色分け」を選択した場合の見え方です） 
 
 
また，ST-Bridge Link は，Revit の通常機能と同様，パネル上でF1 キーを押すことで，
ヘルプを表示することができます。こちらのヘルプも合わせてご利用ください。 
 
 
 
 
モデル変換後の各構造図の作成については，Revit の構造設計向けトレーニングテキスト
をご参照ください。

## Page 14

12                                           ST-Bridge Link Start-Up Guide 
 
 
2.5  差分インポートについて 
ST-Bridge Link 2020 以降には，新機能として差分変換を実装しています。本機能を利用す
るためには，GUID に対応したST-Bridge ファイルを使用する必要があります。ユニオンシ
ステム社様のSuper Build/ SS7 では，ST-Bridge ファイル出力のGUID 対応について，2021
年末頃の実装を予定しています。 
 
本機能の詳細については，今後アップデートいたします。 
現時点での最新情報として，2021 年11 月18 日に実施しました，「構造設計者待望の
Revit での差分変換」のウェビナー動画をご覧ください。

## Page 15

ST-Bridge Link Start-Up Guide                                            13 
 
 
3.  変換できないとき 
変換できない等のトラブルは，弊社カスタマーサポートまでご連絡ください。カスタマー
サポートのURL は，各アプリのダウンロードページにも記載がございます。 
 
 
 
ST-Bridge Link にバグがある際は，プログラム修正等の対応をいたします。なお，ST-
Bridge ファイルの出力元にバグ等があり，ST-Bridge ファイルそのものに不具合がある場合
は，弊社では対応できませんので，ST-Bridge ファイルの出力元にご連絡ください。
