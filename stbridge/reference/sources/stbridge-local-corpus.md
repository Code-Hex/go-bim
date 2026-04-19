---
type: source
title: ST-Bridge ローカル資料棚卸し
tags: [st-bridge, xsd, stb, local-corpus]
related:
  - ./stbridge-historical-specification-and-changelog.md
  - ../synthesis/stbridge-application-utility-opportunities.md
created: 2026-04-19
updated: 2026-04-19
authors: []
year: 2026
venue: local files
---

# ST-Bridge ローカル資料棚卸し

## ローカルにある主要資料

### PDF

- `st-bridge/ST-Bridge2.1_XML仕様説明書ver.2.1.0.pdf`
- `st-bridge/ST-Bridge計算データ編_XML仕様説明書ver.2.1.0.pdf`
- `st-bridge/ST-Bridge_XML仕様説明書（ver.2.0.2）.pdf`
- `st-bridge/ST-Bridge計算データ編_XML仕様説明書ver.2.0.2.pdf`
- `st-bridge/ST_Bridge_Link_start_up_guilde.pdf`

### 公開取得して追加確認した PDF

- `https://www.building-smart.or.jp/wp-content/uploads/2018/07/b8736bb76b760bc8add6b99b29d6bfba.0%EF%BC%89.pdf` (`ver.2.0`)
- `https://www.building-smart.or.jp/wp-content/uploads/2019/04/ST-Bridge_XML%E4%BB%95%E6%A7%98%E8%AA%AC%E6%98%8E%E6%9B%B8%EF%BC%88ver.2.0.1%EF%BC%89.pdf` (`ver.2.0.1`)

これらはすべて Markdown に転記済み:

- [2.0 仕様書 PDF 転記](./transcriptions/st-bridge-2-0-0-xml-spec-transcription.md)
- [2.0.1 仕様書 PDF 転記](./transcriptions/st-bridge-2-0-1-xml-spec-transcription.md)
- [2.1 仕様書 PDF 転記](./transcriptions/st-bridge-2-1-xml-spec-transcription.md)
- [2.1 計算データ編 PDF 転記](./transcriptions/st-bridge-calculation-2-1-xml-spec-transcription.md)
- [2.0.2 仕様書 PDF 転記](./transcriptions/st-bridge-2-0-2-xml-spec-transcription.md)
- [2.0.2 計算データ編 PDF 転記](./transcriptions/st-bridge-calculation-2-0-2-xml-spec-transcription.md)
- [ST-Bridge Link ガイド PDF 転記](./transcriptions/st-bridge-link-start-up-guide-transcription.md)

### XSD / sample

- `st-bridge/ST-Bridge210.zip`
- `st-bridge/ST-Bridge_v202.zip`
- `st-bridge/20171030_st.zip`
- `st-bridge/STBCONVERTER_Archicad28.zip`
- `st-bridge/STBCONVERTER_Archicad29.zip`

### 公開取得して追加確認した XSD

- `https://www.building-smart.or.jp/wp-content/uploads/2018/08/ST-Bridge_v2.zip`
- `https://www.building-smart.or.jp/wp-content/uploads/2022/03/ST-Bridge_v201_20220316.zip`

歴代仕様と change log のまとめは [ST-Bridge 歴代仕様と change log](./stbridge-historical-specification-and-changelog.md) に分離した。

## 仕様本文以外の公式運用資料

- `https://www.building-smart.or.jp/old/download/files/20150602_seminar_3.pdf`
  - `1.0 -> 1.3` の更新履歴と追加テーマ
- `https://www.building-smart.or.jp/old/download/files/20150603_st1.pdf`
  - `2015年版 情報連携マップ`
- `https://www.building-smart.or.jp/old/download/files/20160421_st.pdf`
  - `2016年版 情報連携マップ`
- `https://www.building-smart.or.jp/meeting/buildall/structural-design/`
  - `2017年版`, `2018年版`, `2020年版`, `2023年版` 情報連携マップ

これらは full spec PDF ではないが、各版が実務アプリ群でどう使われていたかを追う operational corpus として有効である。

## 後期版の sample proxy

standalone sample zip は `20171030_st.zip` しか無い。  
そのため `2.0.2` / `2.1.0` の late fields は、仕様書本文に埋め込まれた XML 例とコンバータ配布物で補完した。

`2.1.0` 本編転記で直接確認した埋め込み例:

- `StbExportInformation`
- `StbIsolatingDevice`
- `StbDampingDevices`
- `StbJoints`
- `StbExtensions`
- `StbConversionWarning`

`2.0.2` / `2.1.0` 計算データ編転記で直接確認した埋め込み例:

- `StbCalData`
- `StbAnaRelations`

`STBCONVERTER_Archicad28.zip` / `STBCONVERTER_Archicad29.zip` の `改変履歴.txt` で確認した後期版論点:

- `デッキプレート`
- `S梁の5断面`
- `WALL_SHEAR`
- `床の構造解析モデルは作成しない`
- `ST-Bridge version 1 は変換対象外`

つまり、`2.0.2` / `2.1.0` の新領域は「sample zip が無いから確認不能」ではなく、埋め込み例と配布物 change log で追える。

## 2.1 XSD の骨格

`ST_BRIDGE`:

- `StbCommon`
- `StbModel`
- `StbExtensions`
- `StbExportInformation`
- `StbCalData`
- `StbAnaModels`

`StbModel`:

- `StbNodes`
- `StbAxes`
- `StbStories`
- `StbMembers`
- `StbSections`
- `StbJoints`
- `StbConnections`
- `StbWeld`

`StbMembers`:

- `StbColumns`
- `StbPosts`
- `StbGirders`
- `StbBeams`
- `StbBraces`
- `StbSlabs`
- `StbWalls`
- `StbIsolatingDevices`
- `StbDampingDevices`
- `StbFrameDampingDevices`
- `StbFootings`
- `StbStripFootings`
- `StbPiles`
- `StbFoundationColumns`
- `StbParapets`
- `StbOpenArrangements`
- `StbPenetrationArrangements`
- `StbJointArrangements`
- `StbPanelZoneArrangements`
- `StbConnectionArrangements`

`StbCalData`:

- `StbCalCommon`
- `StbCalLoad`
- `StbCalCondition`
- `StbCalLoadArrangements`
- `StbCalConditionArrangements`

`StbAnaModel`:

- `StbAnaNodes`
- `StbAnaStories`
- `StbAnaMembers`
- `StbAnaProperties`
- `StbAnaFloorDiaphragms`
- `StbAnaMaterials`
- `StbAnaSections`
- `StbAnaLoadCases`
- `StbAnaAnalyses`
- `StbAnaRelations`

結論として、ST-Bridge は「構造部材 XML」ではなく、「構造の物理モデル + 計算条件 + 解析モデル + provenance」の交換形式である。

## XSD 2.0.0 -> 2.0.1 -> 2.0.2 -> 2.1.0 差分

ローカル比較:

- `2.0.0 = 300`
- `2.0.1 = 301`
- `2.0.2 = 595`
- `2.1.0 = 829`

`2.0.0 -> 2.0.1` は大改定ではなく、主に命名整理と訂正である。

- add `9`
- remove `8`
- 代表例
  - `StbDrawingAxis -> StbDrawingAxes`
  - `StbSecBarArrangement_Beam_RC -> StbSecBarArrangementBeam_RC`
  - `StbSecFoundation_RC_Equi_Triangle -> StbSecFoundation_RC_EquiTriangle`
  - `StbSecFigureBeam_S -> StbSecSteelFigureBeam_S`

- 要素名数
  - `2.0.2 = 595`
  - `2.1.0 = 829`
  - `+357 / -123`
- simple type 数
  - `2.0.2 = 6`
  - `2.1.0 = 8`

大きく増えた領域:

- `StbExportInformation`
- `StbConnections`, `StbWeld`
- 免震 / 制振
- 杭一般工法 / 認定工法
- 詳細配筋
- arrangement 系要素
- 解析 relation 系

`ST-Bridge_v202.xsd` の先頭には余計な 3 バイトがあり、そのままでは XML パースできなかった。normalizer 実装ではこの種の前処理も要る。

## STB サンプル

`20171030_st.zip` のサンプルから確認した代表要素数:

### RC

- `StbNodeid`: 1188
- `StbNode`: 333
- `StbGirder`: 104
- `StbSlab`: 70
- `StbColumn`: 61
- `StbWall`: 58
- `StbBeam`: 48
- `StbPile`: 40
- `StbOpen`: 20
- `StbFooting`: 15

### S

- `StbNodeid`: 812
- `StbNode`: 321
- `StbBeam`: 91
- `StbSlab`: 90
- `StbGirder`: 74
- `StbColumn`: 34
- `StbPost`: 17
- `StbSecColumn_S`: 24
- `StbSecBeam_S`: 22
- `StbBrace`: 4

このサンプルは物理モデル中心で、2.1.0 で厚くなった接合, 溶接, 免震 / 制振, 詳細 `StbCalData`, `StbExportInformation` は十分確認できない。

## ST-Bridge Link ガイドから分かること

[ST-Bridge Link Start Up Guide PDF 転記](./transcriptions/st-bridge-link-start-up-guide-transcription.md) では、次が確認できた。

- `解析モデル 3D -> ST-Bridge -> Revit ST-Bridge Link -> BIM モデル 3D -> 構造図 2D`
- Revit 取込みには family / level / material mapping が必要
- 差分 import には GUID が必要
- ST-Bridge は日本の構造分野向けに部材記号や配筋を意識した形式

つまり ST-Bridge は、構造計算と BIM の間を埋める日本ローカルな中間ファイルとして設計されている。

## この source から言えること

1. ST-Bridge は XML として扱いやすい
2. ただし版差が大きく、単純な固定スキーマ前提は危険
3. 守備範囲は狭いが、その狭い範囲ではかなり深い
4. utility の主戦場は、BIM 全体保存ではなく、構造モデルの抽出, 検証, 差分, quality gate になる
