---
type: source
title: ST-Bridge 歴代仕様と change log
tags: [st-bridge, history, changelog, xsd]
related:
  - ./stbridge-local-corpus.md
  - ./transcriptions/st-bridge-2-0-0-xml-spec-transcription.md
  - ./transcriptions/st-bridge-2-0-1-xml-spec-transcription.md
  - ./transcriptions/st-bridge-2-0-2-xml-spec-transcription.md
  - ./transcriptions/st-bridge-calculation-2-0-2-xml-spec-transcription.md
  - ./transcriptions/st-bridge-2-1-xml-spec-transcription.md
  - ./transcriptions/st-bridge-calculation-2-1-xml-spec-transcription.md
  - ../synthesis/stbridge-application-utility-opportunities.md
created: 2026-04-19
updated: 2026-04-19
authors: []
year: 2026
venue: local files
---

# ST-Bridge 歴代仕様と change log

## 参照元

- `https://www.building-smart.or.jp/meeting/buildall/structural-design/`
- `https://www.building-smart.or.jp/wp-content/uploads/2018/07/b8736bb76b760bc8add6b99b29d6bfba.0%EF%BC%89.pdf`
- `https://www.building-smart.or.jp/wp-content/uploads/2019/04/ST-Bridge_XML%E4%BB%95%E6%A7%98%E8%AA%AC%E6%98%8E%E6%9B%B8%EF%BC%88ver.2.0.1%EF%BC%89.pdf`
- `st-bridge/ST-Bridge_XML仕様説明書（ver.2.0.2）.pdf`
- `st-bridge/ST-Bridge計算データ編_XML仕様説明書ver.2.0.2.pdf`
- `st-bridge/ST-Bridge2.1_XML仕様説明書ver.2.1.0.pdf`
- `st-bridge/ST-Bridge計算データ編_XML仕様説明書ver.2.1.0.pdf`
- `https://www.building-smart.or.jp/wp-content/uploads/2018/08/ST-Bridge_v2.zip`
- `https://www.building-smart.or.jp/wp-content/uploads/2022/03/ST-Bridge_v201_20220316.zip`
- `st-bridge/ST-Bridge_v202.zip`
- `st-bridge/ST-Bridge210.zip`
- `st-bridge/20171030_st.zip`
- `https://www.building-smart.or.jp/old/download/files/20150602_seminar_3.pdf`
- `https://www.building-smart.or.jp/old/download/files/20150603_st1.pdf`
- `https://www.building-smart.or.jp/old/download/files/20160421_st.pdf`
- `https://www.building-smart.or.jp/meeting/buildall/structural-design/`
- `st-bridge/STBCONVERTER_Archicad28.zip`
- `st-bridge/STBCONVERTER_Archicad29.zip`

上のうち full spec PDF として確認できた `2.0`, `2.0.1`, `2.0.2`, `2.1.0`, `2.0.2 計算データ編`, `2.1.0 計算データ編`, `Link guide` は Markdown に転記済みである。  
`1.x` の standalone PDF は今回アクセス可能な公開 corpus では直接取得できなかったが、`2.0` / `2.0.1` の修正履歴と `2.0.2` / `2.1.0` の序文で、`Draft1` から `ver.1.4` までの change chain を公式に再構成できた。

さらに、`2015年度セミナー資料` と `2015年版 / 2016年版 / 2017年版 / 2018年版 / 2020年版 / 2023年版` の情報連携マップを併読すると、

- `1.0` から `1.3` までの contemporaneous 公開履歴
- `1.4` と `2.0` が実務アプリ群でどう併存したか
- `2.x` 系がどの連携図で主流化したか

まで、仕様本文以外の公式公開物でも追跡できる。

## 公開確認できた full spec corpus

| 系列 | 版 | 公開日 | 根拠 | Markdown |
| --- | --- | --- | --- | --- |
| 本編 | `2.0.0` | `2018-07-25` | buildingSMART Japan 構造設計小委員会ページ + PDF 本文 | `st-bridge-2-0-0-xml-spec-transcription.md` |
| 本編 | `2.0.1` | `2019-04-09` | buildingSMART Japan 構造設計小委員会ページ + PDF 本文 | `st-bridge-2-0-1-xml-spec-transcription.md` |
| 本編 | `2.0.2` | `2021-02-28` | ローカル PDF + XSD | `st-bridge-2-0-2-xml-spec-transcription.md` |
| 計算 | `2.0.2` | `2021-03-23` | ローカル PDF | `st-bridge-calculation-2-0-2-xml-spec-transcription.md` |
| 本編 | `2.1.0` | `2023-05-17` | ローカル PDF + XSD | `st-bridge-2-1-xml-spec-transcription.md` |
| 計算 | `2.1.0` | `2023-05-17` | ローカル PDF | `st-bridge-calculation-2-1-xml-spec-transcription.md` |
| 連携ガイド | Link guide | 不明 | ローカル PDF | `st-bridge-link-start-up-guide-transcription.md` |

## まず結論

ST-Bridge は単に要素を足してきたのではない。版が上がるたびに次の方向へ広がっている。

1. 物理モデルだけの交換から、計算条件・解析モデル・出力 provenance まで抱える方向
2. 日本実務の細かい構造詳細, 特に杭, 接合, 配筋, 免震, 制振を厚くする方向
3. ベンダ実装差を減らすため、ID 一意性, 子要素構造, 命名規則, 単位, 文字列表現を詰める方向

したがって utility を作るなら、単なる XSD バリデータでは足りず、`版差を吸収する正規化器` が必須になる。

## 年表

### 本編の履歴

- `2010.11.27` Draft1
- `2011.05.20` Draft2
- `2011.09.xx` Draft3
- `2012.03.14` Draft4
- `2012.07.25` Ver.1.0 リリース
- `2019.03.20` Ver.2.0 Revision 1
- `2020.12.16` Ver.2.0 Revision 2
- `2021.02.28` Ver.2.0.2 系 PDF
- `2023.01.31` Ver.2.1 change set
- `2023.03.31` Ver.2.1.0 PDF

### 計算データ編の履歴

- `2015.03.31` Ver.1.0.0 仮リリース
- `2015.04.27` Ver.1.3.0 仮リリース
- `2018.06.27` Ver.1.4.0 仮リリース
- `2021.02.28` Ver.2.0.2 リリース
- `2023.03.31` Ver.2.1.0 リリース

## 1.x 系の contemporaneous public evidence

`2015-06-02` 公開の公式セミナー資料「ST-Bridgeの普及と構造情報連携マップのご紹介」では、更新履歴として次が明示されていた。

- `2012-07` `Ver.1.0`
- `2013-07` `Ver.1.1`
  - `Extension`
  - `3次元的な部材の寄り`
- `2014-06` `Ver.1.2`
  - `基礎柱`
  - `パラペット`
  - `梁配筋の内外端`
- `2015-01` `Ver.1.3`
  - `土間スラブ`
  - `土圧壁`
  - `フーチング形状`

このスライドは `1.0` から `1.3` までの版番号と追加テーマを contemporaneous に押さえる一次資料として使える。

また、構造設計小委員会ページで公開されている情報連携マップから、版の実運用史も追える。

- `2015年版 情報連携マップ`: `1.3` 時点の連携図
- `2016年版 情報連携マップ`: `1.4` 時点の連携図
- `2017年版 情報連携マップ`: `1.4` 運用継続
- `2018年版 情報連携マップ`: `2.0` 公開直前の連携図
- `2020年版 情報連携マップ`: `STB1.4` と `STB2.0` を並記
- `2023年版 情報連携マップ`: `2.x` 系中心の連携図

つまり `1.x` 系は standalone PDF 原本を全版取得できていなくても、公式セミナー資料と公式情報連携マップを使うことで、

1. 版番号
2. 公開時期
3. 当時の追加テーマ
4. 実務アプリへの浸透状況

までは公式公開資料から復元できる。

## 1.x 系は 2.0 / 2.0.1 の修正履歴から再構成できる

`ver.2.0` の巻末 `【修正履歴】` は `Draft1` から `ver.1.4` までを明示している。  
`ver.2.0.1` の巻末ではそこに `2019.3.20 ver.2.0 (Revision 1)` が追加されている。

1.x 系の要点だけ抜くと次の通りである。

- `Draft1` (`2010-11-27`)
  - 初期ドラフト
- `Draft2` (`2011-05-20`)
  - 属性名の小文字化、型見直し、必須条件見直し、鉄骨規格追加
- `Draft3` (`2011-09`)
  - 鉄骨断面追加、`StbBraces` 追加、配筋位置 `pos` 追加
- `Draft4` (`2012-03-14`)
  - `SRC/CFT` 柱梁断面、床組、`ON_SLAB`、RC 杭の `strength_concrete`
- `ver.1.0` (`2012-07-25`)
  - 拡張子定義、材料表記ルール、`isFoundation`
- `ver.1.1` (`2013-07-25`)
  - `StbExtensions`, 作図用軸, 免震階, 3D オフセット, X 型配筋, 片持ち属性
- `ver.1.2` (`2014-06-03`)
  - 基礎柱, パラペット, 壁開口配筋, 大梁 / 梁の外端・内端
- `ver.1.3` (`2015-01-21`)
  - スラブハンチ, `isEarthen`, `isPress`, 基礎形状 / 配筋拡張
- `ver.1.4` (`2016-06-01`)
  - `ON_COLUMN`, 継手情報, 既製スラブ, `StbExtElement`, `StbExtPropertyDef`

つまり、`2.0` は突然生えた版ではなく、`1.0` から `1.4` までに増えてきた部材 / 配筋 / 拡張表現を、XML として整理し直した大改定である。

## 2.0.0 で起きたこと

`ver.2.0` の序文と修正履歴から、主旨はかなり明確である。

- `ver.1.0` から拡張してきた仕様を、XML として厳密に整理する
- データ交換精度を上げる
- 命名規則と子要素構造を整理する

巻末の `2018.7.1 ver.2.0 リリース` では主な変更として次が挙がる。

- id 一意性を「それぞれの子要素内で一意」に変更
- 要素名 / 属性名の命名規則を明確化
- 鉄筋情報を子要素へ集約
- 節点順序をモノリスト化
- 平行移動オフセットを廃止し 3D オフセットへ統一
- `StbX_Axis`, `StbY_Axis` を平行軸へ再整理
- `guid` 追加
- 円弧軸, 放射軸追加
- 既製杭, 鋼管杭追加
- `StbSlabFrames` 削除

要するに `2.0` は、`1.x` の積み上げを「交換フォーマットとして壊れにくい形」に組み替えた版である。

## 2.0.0 -> 2.0.1 で起きたこと

`ver.2.0.1` は新領域追加というより revision release で、修正履歴では次が明記されている。

- `StbNode`, `StbOpen`, `StbSecSteelFigureColumn_SRC`, `StbSecOpen` の補足追記
- `StbSlab.id_section` 型修正 `string -> integer`
- `StbOpen.id_section` の必須条件緩和
- `StbSecSteelColumn_SRC_NotSame`, `ThreeTypes` の子要素回数訂正
- `StbSecBeamXReinforced -> StbSecBarBeamXReinforced` への名称整理
- `StbSecBarSlab_RC_Open` の記載漏れ補完
- `StbSecFoundation_RC_Rect.width_X` 型修正 `string -> double`
- `TRANSVERS_* -> TRANSVERSE_*` の誤記修正
- 開口配筋系の最小回数緩和

XSD 比較でも、`2.0.0 -> 2.0.1` は大規模追加ではなかった。

- 要素名数: `300 -> 301`
- add `9`
- remove `8`

代表的な rename / 整理:

- `StbDrawingAxis -> StbDrawingAxes`
- `StbSecBarArrangement_Beam_RC -> StbSecBarArrangementBeam_RC`
- `StbSecBarArrangement_Pile_RC -> StbSecBarArrangementPile_RC`
- `StbSecFoundation_RC_Equi_Triangle -> StbSecFoundation_RC_EquiTriangle`
- `StbSecFoundation_RC_Tapered_Rect -> StbSecFoundation_RC_TaperedRect`
- `StbSecFigureBeam_S -> StbSecSteelFigureBeam_S`

つまり `2.0.1` は `2.0` の全面改定を受けた first stabilization だと見てよい。

## 2.0.2 で起きたこと

本編 `ver.2.0.2` と計算データ編 `ver.2.0.2` で大きいのは、`StbCalData` と `StbAnaModels` を本体の責務へ明示的に取り込んだ点である。

### 本編側の主要変更

- `ST_BRIDGE` 配下に `StbCalData`, `StbAnaModels` を追記
- ID 一意性と GUID の位置付けを明確化
- グローバル座標との関係 `global_offset_X/Y/Z`, `global_rotation` を追加
- `app_version`, `convert_app_name`, `convert_app_version` を追加
- `strength_concrete` を柱, 大梁, スラブ, 壁などに追加
- `StbSecUndefined`, `StbSecSteelProduct`, `StbSecSteelUndefined` を追加
- 継手要素へ GUID を追加

### 計算データ編側の主要変更

- 本編 `Ver.2.0` 以降に合わせて全面改訂
- `id` の一意性を明確化
- 子要素の持ち方を統一
- 要素への `guid` 追加
- 国内主要構造計算プログラムの荷重 / 計算条件を洗い出して見直し
- `StbAnaRelations` を見直し、`StbModel` と `StbAnaModels` の参照性を改善

### 実装上の意味

`2.0.2` は「設計モデル XML」から「設計 + 計算 + 解析の交換ハブ XML」へ踏み込んだ版である。

## 後期版の sample analysis は埋め込み例と配布物 change log で補強できる

`20171030_st.zip` の standalone sample は古く、`2.0.2` / `2.1.0` で増えた要素群を十分にはカバーしない。  
そのため後期版については、仕様書本文の埋め込み例と実運用配布物の change log を sample proxy として併読した。

`2.1.0` 本編転記では、少なくとも次の埋め込み XML 例を確認できる。

- `StbExportInformation`
- `StbIsolatingDevice`
- `StbDampingDevices`
- `StbJoints`
- `StbExtensions`
- `StbConversionWarning`

`2.0.2` / `2.1.0` 計算データ編転記では、少なくとも次の埋め込み XML 例を確認できる。

- `StbCalData`
- `StbAnaRelations`

さらに `STBCONVERTER_Archicad28.zip` / `STBCONVERTER_Archicad29.zip` の `改変履歴.txt` では、

- `デッキプレート`
- `S梁の5断面`
- `WALL_SHEAR`
- `床の構造解析モデルは作成しない`
- `ST-Bridge version 1 は変換対象外`

といった、後期版で実際に問題になった実装論点が残っていた。

したがって、`2.0.2` / `2.1.0` の late fields は「standalone sample zip が無いから完全に未知」ではなく、

1. 仕様本文の埋め込み XML 例
2. XSD 差分
3. 実運用コンバータの change log
4. 情報連携マップの運用史

の 4 層で確認可能である。

## 2.1.0 で起きたこと

`2.1.0` は `2.0.2` の延長ではあるが、追加領域がかなり大きい。

### 本編側の主要追加

- 製品一覧表
- 工法一覧表
- 鉄骨コネクション情報
- 溶接情報
- 免震装置
- 制振装置
- 床 / 壁の荷重のみ要素
- 梁貫通孔
- コンクリート柱梁接合部 (`StbPanelZone`)
- 出力情報 (`StbExportInformation`)
- RC スラブ断面への型枠要素

### 本編側の主要変更

- 径別鉄筋強度情報リストを杭とそれ以外に分割
- `StbApplyConditionList` を要領情報要素として再整理
- `StbCommon` に `StbAdditionalInformation` を追加
- `StbStory` に `level_name` を追加
- 梁を任意回数断面へ拡張
- 開口寸法を断面側で定義
- 柱梁配筋を鉄筋位置指定で詳細入力可能に
- 基礎断面配筋を詳細化
- 杭断面を一般工法 / 認定工法へ再編
- 鉄骨断面に組立 H/T, 2C, 2L, 2LipC などを追加

### 計算データ編側の主要追加 / 変更

- `StbCalFinish` に `StbCalMemberFinishValues` を追加
- 仕上げ荷重の荷重値指定系列を追加
- 装置性能, 装置性能倍率, 免震 / 制振関連の計算条件と配置を追加
- `StbCalLiveload` に `code` を追加
- 地震 / 風 / 雪 / 積載荷重の型や補足を見直し
- `StbAnaModel` に `name` を追加
- `StbAnaMemberid_List -> StbAnaMemberidList`
- `StbAnaNodeid_List -> StbAnaNodeidList`
- 材料性能 / 断面性能 / 解析要素説明へ単位を追加

### 実装上の意味

`2.1.0` は、日本の構造実務で実際にぶつかる「製品型番」「認定工法」「杭」「接合」「制振 / 免震」「荷重のみ要素」をかなり具体的に標準へ寄せた版である。

## XSD 差分

ローカルで `ST-Bridge_v2.xml`, `ST-Bridge_v201_20220316.xsd`, `ST-Bridge_v202.xsd`, `ST-Bridge210.xsd` を比較した結果:

- 要素名数
  - `2.0.0 = 300`
  - `2.0.1 = 301`
  - `2.0.2 = 595`
  - `2.1.0 = 829`

- `2.0.0 -> 2.0.1`
  - `+9 / -8`
- `2.0.1 -> 2.0.2`
  - `+294 / -0`
- `2.0.2 -> 2.1.0`
  - `+357 / -123`
- simple type 数
  - `2.0.2 = 6`
  - `2.1.0 = 8`
  - `monolist` が整理され、`monolist_id`, `monolist_length`, `ratio` が追加

### 2.1.0 で特に増えた領域

#### 1. 出力 / 追加情報

- `StbAdditionalInformation`
- `StbExportInformation`
- `StbExportLog`
- `StbExportError`
- `StbExportWarning`
- `StbExportNotice`
- `StbExportPolicy`

#### 2. 接合 / 溶接 / コネクション

- `StbConnectionArrangements`
- `StbConnections`
- `StbConnectionSpecs`
- `StbConnectionSpecBeamH`
- `StbConnectionSpecColumnH`
- `StbConnectionSpecColumnBox`
- `StbConnectionSpecColumnPipe`
- `StbConnectionSpecDiaphragm`
- `StbConnectionSpecGussetPlate`
- `StbConnectionSpecPanel`
- `StbConnectionSpecRibPlate`
- `StbConnectionSpecSplice`
- `StbConnectionSpecStiffner`
- `StbWeldCommon`
- `StbWeld`, `StbWeldSpec`
- `StbWeldFillet`, `StbWeldFlare`, `StbWeldFullPenetration`, `StbWeldPartialPenetration`
- `StbDiaphragms`, `StbStiffners`, `StbGussetPlates`, `StbRibPlates`

#### 3. 免震 / 制振

- `StbIsolatingDevices`
- `StbDampingDevices`
- `StbFrameDampingDevices`
- `StbSecIsolatingDevice*`
- `StbSecDampingDevice*`
- `StbCalIsolatingDeviceProperty*`
- `StbCalDampingDeviceProperty*`

#### 4. 杭

- `StbSecPilePrecast*`
- `StbSecPile_RC_Certified`
- `StbSecPile_RC_Conventional*`
- `StbSecPile_S_Certified`
- `StbSecPile_S_Conventional`
- `StbSecPile_S_Product`
- `StbCertificationNumber`
- `StbCertifiedMethodKeyValue`

#### 5. 詳細配筋 / 詳細断面

- `StbSecBarBeamComplex*`
- `StbSecBarColumnRectComplex*`
- `StbSecBarColumnCircleComplex*`
- `StbSecBarSlab_RC_Conventional*`
- `StbSecBarWall_RC_*`
- `StbSecFormworkSlab*`

#### 6. 解析関係

- `StbAnaBeamRel`
- `StbAnaWallRel`
- `StbAnaTrussRel`
- `StbAnaPlanePropertyRel`
- `StbAnaFloorDiaphragmRel`
- `StbAnaSupportRel`

### 2.1.0 で整理 / 置換された代表例

- `StbAnaMemberid_List -> StbAnaMemberidList`
- `StbAnaNodeid_List -> StbAnaNodeidList`
- `StbOpen`, `StbOpenId`, `StbOpenIdList`, `StbOpens` から `StbOpenArrangement` 系へ軸足を移した構造
- RC / SRC / Slab / Pile の旧配筋要素が、より詳細な `Conventional`, `Certified`, `Complex`, `Simple` 系へ置換
- 旧 `StbSecBaseConventional_S/SRC/CFT` 群が、構造種別に依らない `StbSecBaseConventional`, `StbSecBaseProduct` 系へ整理

## サンプル STB とのズレ

`20171030_st.zip` のサンプルは、RC / S の物理モデル中心であり、次の領域は実ファイルで確認できるが、2.1.0 の厚みまでは反映していない。

- 節点
- 階
- 軸
- 柱
- 梁
- スラブ
- 壁
- 杭
- 開口
- 断面

逆に、2.1.0 で厚くなった以下はサンプル側で十分確認できない。

- `StbExportInformation`
- `StbConnections`
- `StbWeld`
- `StbIsolatingDevices`
- `StbDampingDevices`
- 詳細な `StbCalData`
- 詳細な `StbAnaRelations`

だから utility のテストは、旧サンプルだけに依存すると普通に穴が開く。

## utility 設計へ落とすと

歴代仕様を見ると、最低でも次の 4 本は必要になる。

1. `version-aware parser`
   - `2.0.2` と `2.1.0` の版差を吸収
2. `normalizer / migrator`
   - renamed elements や arrangement 系へ正規化
3. `domain validator`
   - 接合, 杭, 免震, 制振, 仕上げ荷重, 解析 relations を個別検証
4. `sample gap tester`
   - 旧サンプルに無い 2.1 領域を synthetic fixture で補う

ST-Bridge を使う utility は、最初から `単一 XSD に固定しない` こと。ここをサボると後で確実に詰む。
