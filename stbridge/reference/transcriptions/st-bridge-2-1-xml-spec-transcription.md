---
type: source
title: ST-Bridge XML ファイル仕様説明書 ver.2.1.0 PDF転記
tags: [pdf-transcription, st-bridge, specification]
related: []
created: 2026-04-19
updated: 2026-04-19
authors: []
year: 2023
url: "/Users/keikamikawa/go/src/github.com/Code-Hex/bim/st-bridge/ST-Bridge2.1_XML仕様説明書ver.2.1.0.pdf"
venue: "local PDF"
---

# ST-Bridge XML ファイル仕様説明書 ver.2.1.0 PDF転記

> 機械抽出による Markdown 転記。代表ページは vision で目視確認済み。表組みは PDF 上の罫線や段組みを完全には保持していない。

- 元 PDF: `/Users/keikamikawa/go/src/github.com/Code-Hex/bim/st-bridge/ST-Bridge2.1_XML仕様説明書ver.2.1.0.pdf`
- ページ数: 666

## PDF ブックマーク

- p.1: 00. 表紙・目次
- p.10: 01. 概要
  - p.10: 1. 概要
    - p.10: 1.1. コンセプト
    - p.11: 1.2. XML形式
    - p.12: 1.3. 表記法
    - p.13: 1.4. 命名規則
    - p.13: 1.5. 属性値の型と表現範囲
    - p.13: 1.6. 単位系と数値範囲
    - p.14: 1.7. グローバル一意識別子（GUID）
    - p.14: 1.8. 集合型（monolist）
    - p.15: 1.9. ファイルの拡張子
    - p.15: 1.10. バージョン番号
    - p.15: 1.11. 名前空間（Namespace）
    - p.15: 1.12. XML Schema の利用
    - p.16: 1.13. 製品型番一覧表
    - p.16: 1.14. 工法一覧表
- p.17: 02. 要素リファレンス 特記事項
  - p.17: 2. 要素リファレンス　特記事項
    - p.17: 2.1. 全体構成
    - p.19: 2.2. 主要な要素のIDと一意性
    - p.21: 2.3. 部材の配置
    - p.22: 2.4. 材料の表記
    - p.23: 2.5. 鉄筋コンクリート部材における鉄筋の重心位置とかぶり厚さについて
    - p.24: 2.6. 座標系
    - p.26: 2.7. 仕様書の見方
    - p.28: 2.8. ST-Bridgeで同じ内容を示す場合の優先順位と推奨出力について
- p.29: 03. 要素リファレンス 共通情報
  - p.29: 3. 要素リファレンス　共通情報
    - p.29: 3.1. ST-Bridge：ST_BRIDGE
    - p.30: 3.2. 共通情報：StbCommon
      - p.33: 3.2.1.  径別鉄筋強度情報リスト：StbReinforcementStrengthList
        - p.33: 3.2.1.1.   径別鉄筋強度情報：StbReinforcementStrength
      - p.34: 3.2.2. 杭の径別鉄筋強度情報リスト：StbReinforcementStrengthListPile
      - p.35: 3.2.3. 属性・条件適用リスト： StbApplyConditionsList
        - p.37: 3.2.3.1. StbApplyConditionList_RC：RC要領
          - p.38: 3.2.3.1.1. StbApply_RC_General：RC要領　一般事項
          - p.39: 3.2.3.1.2. StbApply_RC_Pile：RC要領　杭
          - p.40: 3.2.3.1.3. StbApply_RC_Foundation：RC要領　基礎
          - p.41: 3.2.3.1.4. StbApply_RC_FoundationGirder：RC要領　基礎大梁
          - p.43: 3.2.3.1.5. StbApply_RC_FoundationBeam：RC要領　基礎小梁
          - p.45: 3.2.3.1.6. StbApply_RC_Column：RC要領　柱
          - p.47: 3.2.3.1.7. StbApply_RC_Girder：RC要領　大梁
          - p.49: 3.2.3.1.8. StbApply_RC_Beam：RC要領　小梁
          - p.51: 3.2.3.1.9. StbApply_RC_Wall：RC要領　壁
          - p.52: 3.2.3.1.10. StbApply_RC_Slab：RC要領　スラブ
        - p.53: 3.2.3.2. StbApplyConditionList_S：S要領
          - p.54: 3.2.3.2.1. StbApply_S_General：S要領　一般事項
          - p.55: 3.2.3.2.2. StbApply_S_Column：S要領　柱
          - p.56: 3.2.3.2.3. StbApply_S_Girder：S要領　大梁
          - p.57: 3.2.3.2.4. StbApply_S_Beam：S要領　小梁
      - p.58: 3.2.4. システム標準板厚リスト：StbStandardPlateThicknessList
      - p.59: 3.2.5. コネクションスペック情報：StbConnectionSpecs
        - p.60: 3.2.5.1. Ｓ梁仕口・Ｈ形コネクション仕様情報：StbConnectionSpecBeamH
          - p.61: 3.2.5.1.1. ガセットプレート仕様：StbConnectionSpecGussetPlate
            - p.65: 3.2.5.1.1.1. スプライスプレート仕様：StbConnectionSpecSplice
          - p.66: 3.2.5.1.2. リブプレート仕様：StbConnectionSpecRibPlate
        - p.67: 3.2.5.2. Ｓ柱仕口・Ｈ形コネクション仕様情報：StbConnectionSpecColumnH
          - p.68: 3.2.5.2.1. スチフナープレート仕様：StbConnectionSpecStiffner
        - p.69: 3.2.5.3. Ｓ柱仕口・角形鋼管コネクション仕様情報：StbConnectionSpecColumnBox
          - p.70: 3.2.5.3.1. ダイアフラム仕様：StbConnectionSpecDiaphragm
          - p.72: 3.2.5.3.2. パネルゾーン仕様：StbConnectionSpecPanel
            - p.73: 3.2.5.3.2.1. パネルゾーン材種：StbConnectionSpecPanelStrength
        - p.74: 3.2.5.4. Ｓ柱仕口・円形鋼管コネクション仕様情報：StbConnectionSpecColumnPipe
      - p.75: 3.2.6. 溶接情報：StbWeldCommon
      - p.76: 3.2.7. 追加情報：StbAdditionalInformation
- p.77: 04. 要素リファレンス 位置情報
  - p.77: 4. 要素リファレンス　位置情報
    - p.77: 4.1. 位置・断面情報：StbModel
    - p.78: 4.2. 節点（複数）：StbNodes
      - p.79: 4.2.1. 節点：StbNode
      - p.82: 4.2.2. 節点IDリスト：StbNodeIdList
      - p.83: 4.2.3. 節点ID：StbNodeId
      - p.84: 4.2.4. 順序のある節点ID：StbNodeIdOrder
    - p.85: 4.3. 軸（複数）：StbAxes
      - p.86: 4.3.1. 平行軸(複数)：StbParallelAxes
        - p.87: 4.3.1.1. 平行軸：StbParallelAxis
      - p.89: 4.3.2. 円弧軸(複数)：StbArcAxes
        - p.90: 4.3.2.1. 円弧軸：StbArcAxis
      - p.91: 4.3.3. 放射軸(複数)：StbRadialAxes
        - p.92: 4.3.3.1. 放射軸：StbRadialAxis
      - p.93: 4.3.4. 作図用軸(複数)：StbDrawingAxes
        - p.94: 4.3.4.1. 作図用直線軸：StbDrawingLineAxis
        - p.95: 4.3.4.2. 作図用円弧軸：StbDrawingArcAxis
    - p.96: 4.4. 階（複数）：StbStories
      - p.97: 4.4.1. 階：StbStory
- p.99: 05. 要素リファレンス 部材情報
  - p.99: 5. 要素リファレンス　部材情報
    - p.99: 5.1. 部材情報：StbMembers
    - p.101: 5.2. 柱（複数）：StbColumns
      - p.102: 5.2.1. 柱：StbColumn
        - p.106: 5.2.1.1. 柱中間節点：StbColumnViaNode
        - p.107: 5.2.1.2. 中間節点オフセットリスト：StbMemberOffsetList
    - p.108: 5.3. 間柱（複数）：StbPosts
      - p.109: 5.3.1. 間柱：StbPost
        - p.110: 5.3.1.1. 間柱中間節点：StbPostViaNode
    - p.111: 5.4. 大梁（複数）：StbGirders
      - p.112: 5.4.1. 大梁：StbGirder
        - p.115: 5.4.1.1. 大梁中間節点：StbGirderViaNode
        - p.116: 5.4.1.2.  RC躯体位置：StbGirderConcreteSwitch
        - p.117: 5.4.1.3. 配筋切り替え位置：StbGirderBarSwitch
        - p.118: 5.4.1.4. S躯体位置：StbGirderSteelSwitch
    - p.119: 5.5. 小梁（複数）：StbBeams
      - p.119: 5.5.1. 小梁：StbBeam
        - p.120: 5.5.1.1. 小梁中間節点：StbBeamViaNode
        - p.120: 5.5.1.2. RC躯体位置：StbBeamConcreteSwitch
        - p.121: 5.5.1.3. 配筋切り替え位置：StbBeamBarSwitch
        - p.121: 5.5.1.4. S躯体切り替え位置：StbBeamSteelSwitch
    - p.122: 5.6. ブレース（複数）：StbBraces
      - p.123: 5.6.1. ブレース：StbBrace
    - p.126: 5.7. スラブ（複数）：StbSlabs
      - p.127: 5.7.1. スラブ：StbSlab
        - p.130: 5.7.1.1. スラブオフセットリスト：StbSlabOffsetList
          - p.130: 5.7.1.1.1. スラブオフセット：StbSlabOffset
    - p.131: 5.8. 壁（複数）：StbWalls
      - p.132: 5.8.1. 壁：StbWall
        - p.136: 5.8.1.1. 壁オフセットリスト：StbWallOffsetList
          - p.136: 5.8.1.1.1. 壁オフセット：StbWallOffset
    - p.137: 5.9. 免震装置（複数）：StbIsolatingDevices
      - p.138: 5.9.1. 免震装置：StbIsolatingDevice
        - p.142: 5.9.1.1. 免震装置配置位置（複数台数配置時）：StbIsolatingDevicePosition
    - p.143: 5.10. 制振装置（複数）： StbDampingDevices
      - p.145: 5.10.1. 制振装置：StbDampingDevice
      - p.150: 5.10.2. 制振装置配置位置（複数台数配置時）：StbDampingDevicePosition
    - p.151: 5.11. 制振装置（フレーム）（複数）： StbFrameDampingDevices
      - p.152: 5.11.1. 制振装置（フレーム）：StbFrameDampingDevice
        - p.156: 5.11.1.1. 制振装置（フレーム）オフセットリスト：StbFrameDampingDeviceOffsetList
          - p.156: 5.11.1.1.1. 制振装置（フレーム）オフセット：StbFrameDampingDeviceOffset
        - p.157: 5.11.1.2. 制振装置（フレーム）構成：StbFrameDampingDeviceConfigulation
          - p.159: 5.11.1.2.1. 制振装置（フレーム）構成・制振装置部材：StbFrameDampingDeviceMember
          - p.160: 5.11.1.2.2. 制振装置（フレーム）構成・接合部部材：StbFrameDampingDeviceConnection
    - p.161: 5.12. フーチング（複数）：StbFootings
      - p.162: 5.12.1. フーチング：StbFooting
    - p.164: 5.13. 布基礎（複数）：StbStripFootings
      - p.165: 5.13.1. 布基礎：StbStripFooting
    - p.167: 5.14. 杭基礎（複数）：StbPiles
      - p.168: 5.14.1. 杭基礎：StbPile
    - p.171: 5.15. 基礎柱（複数）：StbFoundationColumns
      - p.172: 5.15.1. 基礎柱：StbFoundationColumn
    - p.175: 5.16. パラペット（複数）：StbParapets
      - p.176: 5.16.1. パラペット：StbParapet
    - p.178: 5.17. 開口配置情報（複数）：StbOpenArrangements
      - p.179: 5.17.1. 開口配置情報：StbOpenArrangement
        - p.180: 5.17.1.1.
    - p.181: 5.18. 梁貫通孔配置情報（複数）：StbPenetrationArrangements
      - p.182: 5.18.1. 梁貫通孔配置情報：StbPenetrationArrangement
    - p.184: 5.19. 継手配置情報（複数）：StbJointArrangements
      - p.185: 5.19.1.  継手配置情報：StbJointArrangement
    - p.186: 5.20. コンクリート柱梁接合部配置情報（複数）：StbPanelZoneArrangements
      - p.187: 5.20.1.  コンクリート柱梁接合部：StbPanelZoneArrangement
    - p.188: 5.21. コネクション配置情報（複数）：StbConnectionArrangements
      - p.189: 5.21.1. コネクション配置情報：StbConnectionArrangement
        - p.190: 5.21.1.1. 主材となる柱：StbConnectedColumn
        - p.191: 5.21.1.2. 主材となる間柱：StbConnectedPost
        - p.192: 5.21.1.3. 主材となる大梁：StbConnectedGirder
        - p.193: 5.21.1.4. 主材となる小梁：StbConnectedBeam
        - p.194: 5.21.1.5. 主材となるブレース：StbConnectedBrace
        - p.195: 5.21.1.6. ガセットプレート配置情報：StbConnectionGussetPlate
          - p.196: 5.21.1.6.1. 接続する間柱：StbConnectingPost
          - p.197: 5.21.1.6.2. 接続する大梁：StbConnectingGirder
          - p.198: 5.21.1.6.3. 接続する小梁：StbConnectingBeam
          - p.199: 5.21.1.6.4. 接続するブレース：StbConnectingBrace
          - p.200: 5.21.1.6.5. リブプレート配置情報：StbConnectionRibPlate
        - p.201: 5.21.1.7. ダイアフラム配置情報：StbConnectionDiaphragm
        - p.202: 5.21.1.8. スチフナー配置情報：StbConnectionStiffner
- p.203: 06.01. 要素リファレンス 断面情報
  - p.203: 6. 要素リファレンス　断面情報
    - p.203: 6.1. 断面情報：StbSections
- p.205: 06.02. RC柱断面
  - p.205: 6.2. ＲＣ柱断面：StbSecColumn_RC
    - p.207: 6.2.1. ＲＣ柱断面形状：StbSecFigureColumn_RC
      - p.208: 6.2.1.1. コンクリート柱断面形状・矩形：StbSecColumnRect
      - p.209: 6.2.1.2. コンクリート柱断面形状・円形：StbSecColumnCircle
    - p.210: 6.2.2. ＲＣ柱断面配筋：StbSecBarArrangementColumn_RC
      - p.212: 6.2.2.1. コンクリート柱断面配筋・矩形・同一：StbSecBarColumnRectSame
      - p.214: 6.2.2.2. コンクリート柱断面配筋・矩形・同一・簡易：StbSecBarColumnRectSameSimple
        - p.217: 6.2.2.2.1.1. コンクリート柱断面配筋・追加鉄筋：StbSecBarColumnAdditional
        - p.218: 6.2.2.2.2. コンクリート柱断面配筋・矩形・同一・詳細：StbSecBarColumnRectSameComplex
          - p.219: 6.2.2.2.2.1. コンクリート柱断面配筋・矩形・詳細・主筋：StbSecBarColumnRectComplexMain
            - p.221: 6.2.2.2.2.1.1. コンクリート柱断面配筋・矩形・詳細・主筋列：StbSecBarColumnRectComplexMainLine
            - p.222: 6.2.2.2.2.1.2. コンクリート柱断面配筋・矩形・詳細・主筋位置：StbSecBarColumnRectComplexMainLoc
          - p.223: 6.2.2.2.2.2. コンクリート柱断面配筋・矩形・詳細・帯筋：StbSecBarColumnRectComplexHoop
            - p.224: 6.2.2.2.2.2.1.  コンクリート柱断面配筋・矩形・詳細・X方向帯筋位置：StbSecBarColumnRectComplexHoopLocX
            - p.225: 6.2.2.2.2.2.2. コンクリート柱断面配筋・矩形・詳細・Y方向帯筋位置：StbSecBarColumnRectComplexHoopLocY
          - p.226: 6.2.2.2.2.3. コンクリート柱断面配筋・矩形・詳細・軸筋：StbSecBarColumnRectComplexAxial
            - p.227: 6.2.2.2.2.3.1. コンクリート柱断面配筋・矩形・詳細・軸筋位置：StbSecBarColumnRectComplexAxialLoc
      - p.228: 6.2.2.3.  コンクリート矩形柱　Ｘ形配筋：StbSecBarColumnXReinforced
      - p.229: 6.2.2.4. コンクリート柱断面配筋・矩形・柱頭脚別：StbSecBarColumnRectNotSame
        - p.231: 6.2.2.4.1. コンクリート柱断面配筋・矩形・柱頭脚別・簡易：StbSecBarColumnRectNotSameSimple
        - p.234: 6.2.2.4.2. コンクリート柱断面配筋・矩形・詳細：StbSecBarColumnRectNotSameComplex
      - p.235: 6.2.2.5. コンクリート柱断面配筋・円形・同一：StbSecBarColumnCircleSame
        - p.236: 6.2.2.5.1. コンクリート柱断面配筋・円形・同一・簡易：StbSecBarColumnCircleSameSimple
        - p.238: 6.2.2.5.2. コンクリート柱断面配筋・円形・同一・詳細：StbSecBarColumnCircleSameComplex
          - p.239: 6.2.2.5.2.1. コンクリート柱断面配筋・円形・詳細・主筋：StbSecBarColumnCircleComplexMain
            - p.240: 6.2.2.5.2.1.1. コンクリート柱断面配筋・円形・詳細・主筋位置：StbSecBarColumnCircleComplexMainLoc
          - p.241: 6.2.2.5.2.2. コンクリート柱断面配筋・円形・詳細・帯筋：StbSecBarColumnCircleComplexHoop
            - p.242: 6.2.2.5.2.2.1. コンクリート柱断面配筋・円形・詳細・帯筋中子位置：StbSecBarColumnCircleComplexHoopLoc
          - p.243: 6.2.2.5.2.3. コンクリート柱断面配筋・円形・詳細・軸筋：StbSecBarColumnCircleComplexAxial
            - p.244: 6.2.2.5.2.3.1. コンクリート柱断面配筋・円形・詳細・軸筋位置：StbSecBarColumnCircleComplexAxialLoc
      - p.245: 6.2.2.6.  StbSecBarColumnCircleNotSame：コンクリート柱断面配筋・円形・柱頭脚別
        - p.246: 6.2.2.6.1. コンクリート柱円形断面配筋・簡易：StbSecBarColumnCircleNotSameSimple
        - p.248: 6.2.2.6.2. コンクリート柱円形断面配筋・詳細（柱頭・柱脚が別配筋の場合）：StbSecBarCircleCircleNotSameComplex
- p.249: 06.03. S柱断面
  - p.249: 6.3. Ｓ柱断面：StbSecColumn_S
    - p.251: 6.3.1. Ｓ柱断面鉄骨形状：StbSecSteelFigureColumn_S
      - p.253: 6.3.1.1. Ｓ柱断面鉄骨形状・同一：StbSecSteelColumn_S_Same
      - p.254: 6.3.1.2. Ｓ柱断面鉄骨形状・柱頭脚別：StbSecSteelColumn_S_NotSame
      - p.255: 6.3.1.3. Ｓ柱断面鉄骨形状・３種類：StbSecSteelColumn_S_ThreeTypes
    - p.256: 6.3.2. 鉄骨断面柱脚製品：StbSecBaseProduct
    - p.258: 6.3.3. 鉄骨断面柱脚在来工法：StbSecBaseConventional
      - p.259: 6.3.3.1. 鉄骨断面柱脚在来工法・ベースプレート：StbSecBaseConventionalPlate
      - p.261: 6.3.3.2. 鉄骨断面柱脚在来工法・アンカーボルト：StbSecBaseConventionalAnchorBolts
        - p.263: 6.3.3.2.1. 鉄骨断面柱脚在来工法・アンカーボルト詳細：StbSecBaseConventionalAnchorBolt
      - p.264: 6.3.3.3. 鉄骨断面柱脚在来工法・リブプレート：StbSecBaseConventionalRibPlates
        - p.265: 6.3.3.3.1. 鉄骨断面柱脚在来工法・リブプレート：StbSecBaseConventionalRibPlate
- p.266: 06.04. SRC柱断面
  - p.266: 6.4. ＳＲＣ柱断面：StbSecColumn_SRC
    - p.268: 6.4.1. ＳＲＣ柱断面形状：StbSecFigureColumn_SRC
    - p.269: 6.4.2. ＳＲＣ柱断面配筋：StbSecBarArrangementColumn_SRC
    - p.270: 6.4.3. ＳＲＣ柱断面鉄骨形状：StbSecSteelFigureColumn_SRC
      - p.272: 6.4.3.1. ＳＲＣ柱断面鉄骨形状・同一：StbSecSteelColumn_SRC_Same
        - p.274: 6.4.3.1.1. ＳＲＣ柱断面鉄骨形状・Ｈ形：StbSecSteelColumn_SRC_ShapeH
        - p.276: 6.4.3.1.2. ＳＲＣ柱断面鉄骨形状・□形：StbSecSteelColumn_SRC_ShapeBox
        - p.277: 6.4.3.1.3. ＳＲＣ柱断面鉄骨形状・○形：StbSecSteelColumn_SRC_ShapePipe
        - p.278: 6.4.3.1.4. ＳＲＣ柱断面鉄骨形状・＋形（H+H）：StbSecSteelColumn_SRC_ShapeCross1
        - p.280: 6.4.3.1.5. ＳＲＣ柱断面鉄骨形状・＋形（H+T+T）：StbSecSteelColumn_SRC_ShapeCross2
        - p.282: 6.4.3.1.6. ＳＲＣ柱断面鉄骨形状・Ｔ形：StbSecSteelColumn_SRC_ShapeT
      - p.284: 6.4.3.2. ＳＲＣ柱断面鉄骨形状・柱頭脚別：StbSecSteelColumn_SRC_NotSame
      - p.286: 6.4.3.3. ＳＲＣ柱断面鉄骨形状・３種類：StbSecSteelColumn_SRC_ThreeTypes
- p.288: 06.05. CFT柱断面
  - p.288: 6.5. ＣＦＴ柱断面：StbSecColumn_CFT
    - p.290: 6.5.1. ＣＦＴ柱断面鉄骨形状：StbSecSteelFigureColumn_CFT
      - p.292: 6.5.1.1. ＣＦＴ柱断面鉄骨形状・同一：StbSecSteelColumn_CFT_Same
      - p.293: 6.5.1.2. ＣＦＴ柱断面鉄骨形状・柱頭脚別：StbSecSteelColumn_CFT_NotSame
      - p.294: 6.5.1.3. ＣＦＴ柱断面鉄骨形状・３種類：StbSecSteelColumn_CFT_ThreeTypes
- p.295: 06.06. RC梁断面
  - p.295: 6.6. ＲＣ梁断面：StbSecBeam_RC
    - p.297: 6.6.1. ＲＣ梁断面形状：StbSecFigureBeam_RC
      - p.298: 6.6.1.1. コンクリート梁断面形状・ストレート：StbSecBeamStraight
      - p.299: 6.6.1.2. コンクリート梁断面形状・テーパー：StbSecBeamTaper
    - p.301: 6.6.2. ＲＣ梁断面配筋：StbSecBarArrangementBeam_RC
      - p.303: 6.6.2.1. コンクリート梁断面配筋・簡易：StbSecBarBeamSimple
        - p.305: 6.6.2.1.1. コンクリート梁断面配筋・簡易主筋：StbSecBarBeamSimpleMain
        - p.306: 6.6.2.1.2. コンクリート梁断面配筋・追加鉄筋：StbSecBarBeamAdditional
      - p.307: 6.6.2.2. コンクリート梁断面配筋・３種類・詳細：StbSecBarBeamComplex
        - p.308: 6.6.2.2.1.  コンクリート梁断面配筋・詳細・主筋：StbSecBarBeamComplexMain
          - p.310: 6.6.2.2.1.1. コンクリート梁断面配筋・詳細・主筋・各行：StbSecBarBeamComplexMainLine
          - p.311: 6.6.2.2.1.2. コンクリート梁断面配筋・詳細・主筋位置：StbSecBarBeamComplexMainLoc
        - p.312: 6.6.2.2.2. コンクリート梁断面配筋・詳細・あばら筋：StbSecBarBeamComplexStirrup
          - p.313: 6.6.2.2.2.1. コンクリート梁断面配筋・詳細・あばら筋位置：StbSecBarBeamComplexStirrupLoc
        - p.314: 6.6.2.2.3. コンクリート梁断面配筋・詳細・腹筋：StbSecBarBeamComplexWeb
          - p.315: 6.6.2.2.3.1. コンクリート梁断面配筋・詳細・腹筋位置：StbSecBarBeamComplexWebLoc
      - p.316: 6.6.2.3. コンクリート梁　Ｘ形配筋：StbSecBarBeamXReinforced
- p.317: 06.07. S梁断面
  - p.317: 6.7. Ｓ梁断面：StbSecBeam_S
    - p.319: 6.7.1. Ｓ梁断面鉄骨形状：StbSecSteelFigureBeam_S
      - p.320: 6.7.1.1. Ｓ梁断面任意鉄骨形状：StbSecSteelBeam_S_Shape
        - p.321: 6.7.1.1.1. ストレート要素：StbSecSteelBeamStraight
        - p.322: 6.7.1.1.2. テーパー要素：StbSecSteelBeamTaper
      - p.323: 6.7.1.2. 拡幅要素：StbSecSteelBeamWidening
- p.324: 06.08. SRC梁断面
  - p.324: 6.8. ＳＲＣ梁断面：StbSecBeam_SRC
    - p.326: 6.8.1. ＳＲＣ梁断面形状：StbSecFigureBeam_SRC
    - p.327: 6.8.2. ＳＲＣ梁断面配筋：StbSecBarArrangementBeam_SRC
    - p.328: 6.8.3. ＳＲＣ梁断面鉄骨形状：StbSecSteelFigureBeam_SRC
      - p.329: 6.8.3.1. SRC梁断面任意鉄骨形状：StbSecSteelBeam_SRC_Shape
- p.330: 06.09. Sブレース断面
  - p.330: 6.9. Ｓブレース断面：StbSecBrace_S
    - p.332: 6.9.1. Ｓブレース断面鉄骨形状：StbSecSteelFigureBrace_S
      - p.333: 6.9.1.1. Ｓブレース断面鉄骨形状・同一：StbSecSteelBrace_S_Same
      - p.334: 6.9.1.2. Ｓブレース断面鉄骨形状・頭脚部別：StbSecSteelBrace_S_NotSame
      - p.335: 6.9.1.3. Ｓブレース断面鉄骨形状・３種類：StbSecSteelBrace_S_ThreeTypes
- p.336: 06.10. RCスラブ断面
  - p.336: 6.10. ＲＣスラブ断面：StbSecSlab_RC
    - p.338: 6.10.1.  ＲＣスラブ在来断面：StbSecSlab_RC_Conventional
      - p.339: 6.10.1.1. ＲＣ在来スラブ断面形状：StbSecFigureSlab_RC_Conventional
        - p.340: 6.10.1.1.1. ＲＣ在来スラブ断面形状・ストレート：StbSecSlab_RC_ConventionalStraight
        - p.341: 6.10.1.1.2. ＲＣ在来スラブ断面形状・テーパー：StbSecSlab_RC_ConventionalTaper
        - p.342: 6.10.1.1.3. ＲＣ在来スラブ断面形状・ハンチ：StbSecSlab_RC_ConventionalHaunch
      - p.344: 6.10.1.2. ＲＣ在来スラブ断面配筋：StbSecBarArrangementSlab_RC_Conventional
        - p.346: 6.10.1.2.1. ＲＣ在来スラブ断面配筋・標準：StbSecBarSlab_RC_ConventionalStandard
        - p.349: 6.10.1.2.2. ＲＣ在来スラブ断面配筋・２方向：StbSecBarSlab_RC_Conventional2Way
        - p.351: 6.10.1.2.3. ＲＣ在来スラブ断面配筋・１方向１：StbSecBarSlab_RC_Conventional1Way1
        - p.353: 6.10.1.2.4. ＲＣ在来スラブ断面配筋・１方向２：StbSecBarSlab_RC_Conventional1Way2
        - p.355: 6.10.1.2.5. スラブ開口配筋：StbSecBarSlab_RC_Open
      - p.356: 6.10.1.3. 型枠要素：StbSecFormworkSlab
        - p.357: 6.10.1.3.1.  在来型枠：StbSecFormworkSlabWood
        - p.358: 6.10.1.3.2. フラットデッキ型枠：StbSecFormworkSlabFlatDeck
          - p.359: 6.10.1.3.2.1. フラットデッキ型枠仕様指定：StbSecFormworkSlabFlatDeckSpec
          - p.360: 6.10.1.3.2.2.  フラットデッキ型番指定：StbSecFormworkSlabFlatDeckProduct
    - p.361: 6.10.2. ＲＣスラブトラス断面：StbSecSlab_RC_Truss
      - p.362: 6.10.2.1. トラス筋付きデッキスラブ製品：StbSecSlab_RC_TrussProduct
        - p.363: 6.10.2.1.1. トラス筋付きデッキスラブ断面配筋：StbSecBarArrangementSlab_RC_Truss
          - p.364: 6.10.2.1.1.1. トラス筋付きデッキスラブ断面配筋・１方向：StbSecBarSlab_RC_Truss1Way
- p.366: 06.11. デッキ合成スラブ断面
  - p.366: 6.11. デッキ合成スラブ断面：StbSecSlabDeck
    - p.368: 6.11.1. 合成デッキ製品：StbSecSlabDeckProduct
      - p.369: 6.11.1.1. デッキ合成スラブ断面配筋：StbSecBarArrangementSlabDeck
        - p.370: 6.11.1.1.1. デッキ合成スラブ断面配筋・１方向：StbSecBarSlabDeck1Way
      - p.371: 6.11.1.2. 認定番号：StbCertificationNumber
- p.372: 06.12. 既製スラブ断面
  - p.372: 6.12. 既製スラブ断面：StbSecSlabPrecast
    - p.374: 6.12.1. 既製スラブトップ部分断面形状：StbSecFigureSlabPrecast
      - p.374: 6.12.1.1. 既製スラブトップ部分断面形状・ストレート：StbSecSlabPrecastStraight
    - p.375: 6.12.2. 既製スラブトップ部分断面配筋：StbSecBarArrangementSlabPrecast
      - p.376: 6.12.2.1. 既製スラブ断面配筋・標準：StbSecBarSlabPrecastStandard
      - p.376: 6.12.2.2. 既製スラブ断面配筋・２方向：StbSecBarSlabPrecast2Way
      - p.376: 6.12.2.3. 既製スラブ断面配筋・１方向：StbSecBarSlabPrecast1Way
    - p.377: 6.12.3. 既製スラブ製品：StbSecProductSlabPrecast
- p.378: 06.13. 荷重用スラブ断面
  - p.378: 6.13. 荷重用スラブ断面：StbSecSlabLoad
- p.379: 06.14. RC壁断面
  - p.379: 6.14. ＲＣ壁断面：StbSecWall_RC
    - p.380: 6.14.1. ＲＣ壁断面形状：StbSecFigureWall_RC
      - p.381: 6.14.1.1.  ＲＣ壁断面形状・ストレート：StbSecWall_RC_Straight
      - p.382: 6.14.1.2.  ＲＣ壁断面形状・テーパー：StbSecWall_RC_Taper
    - p.384: 6.14.2. ＲＣ壁断面配筋：StbSecBarArrangementWall_RC
      - p.386: 6.14.2.1. ＲＣ壁断面配筋・シングル：StbSecBarWall_RC_Single
      - p.387: 6.14.2.2. ＲＣ壁断面配筋・千鳥：StbSecBarWall_RC_Zigzag
      - p.388: 6.14.2.3. ＲＣ壁断面配筋・ダブル：StbSecBarWall_RC_DoubleNet
      - p.389: 6.14.2.4. ＲＣ壁断面配筋（内外異なる）：StbSecBarWall_RC_InsideAndOutside
        - p.393: 6.14.2.4.1. ＲＣ壁断面配筋（内外異なる）全断面一様な配筋：StbSecBarWall_RC_All
        - p.394: 6.14.2.4.2. ＲＣ壁断面配筋（内外異なる）縦筋上端・横筋始端：StbSecBarWall_RC_TopStart
        - p.395: 6.14.2.4.3. ＲＣ壁断面配筋（内外異なる）縦筋中央・横筋中央：StbSecBarWall_RC_Middle
        - p.396: 6.14.2.4.4. ＲＣ壁断面配筋（内外異なる）縦筋下端・横筋終端：StbSecBarWall_RC_BottomEnd
      - p.397: 6.14.2.5. 端部補強筋（コ型補強筋）：StbSecBarWall_RC_Edge
      - p.398: 6.14.2.6. 壁開口配筋：StbSecBarWall_RC_Open
- p.400: 06.15. 荷重用壁断面
  - p.400: 6.15. 荷重用壁断面：StbSecWallLoad
- p.401: 06.16. 免震装置断面
  - p.401: 6.16. 免震装置断面：StbSecIsolatingDevice
    - p.403: 6.16.1. 免震装置断面・製品：StbSecProductIsolatingDevice
      - p.404: 6.16.1.1. 天然ゴム系積層ゴム支承：StbSecIsolatingDeviceNRB
        - p.405: 6.16.1.1.1. 免震装置・規定仕様の変更：StbSecIsolatingDeviceSpecificationChange
      - p.407: 6.16.1.2. 高減衰ゴム系積層ゴム支承：StbSecIsolatingDeviceHDR
      - p.408: 6.16.1.3. 鉛プラグ入り積層ゴム支承：StbSecIsolatingDeviceLRB
      - p.409: 6.16.1.4. 錫プラグ入り積層ゴム支承：StbSecIsolatingDeviceTRB
      - p.410: 6.16.1.5. ダンパー一体型積層ゴム支承：StbSecIsolatingDeviceSDRB
      - p.411: 6.16.1.6. 弾性すべり支承：StbSecIsolatingDeviceESB
      - p.412: 6.16.1.7. 剛すべり支承：StbSecIsolatingDeviceRSB
      - p.413: 6.16.1.8. 曲面すべり支承：StbSecIsolatingDeviceCSB
      - p.414: 6.16.1.9. レール式すべり支承：StbSecIsolatingDeviceCLSB
      - p.415: 6.16.1.10. レール式転がり支承：StbSecIsolatingDeviceCLB
    - p.416: 6.16.2. 免震装置断面・仕様指定：StbSecSpecificationIsolatingDevice
    - p.417: 6.16.3. 免震装置断面・接合部：StbSecConnectionIsolatingDevice
      - p.418: 6.16.3.1. 積層ゴム支承接合部：StbSecConnectionIsolatingDeviceRB
      - p.420: 6.16.3.2. すべり板設置型支承接合部：StbSecConnectionIsolatingDeviceSP
      - p.422: 6.16.3.3. 免震層上下接合型支承接合部：StbSecConnectionIsolatingDeviceLB
- p.424: 06.17. 制振装置断面
  - p.424: 6.17. 制振装置断面：StbSecDampingDevice
    - p.426: 6.17.1. 制振装置断面・製品：StbSecProductDampingDevice
      - p.427: 6.17.1.1. 流体系ダンパー：StbSecDampingDeviceOil
        - p.428: 6.17.1.1.1. 制振装置・規定仕様の変更：StbSecDampingDeviceSpecificationChange
      - p.430: 6.17.1.2. 粘性体ダンパー：StbSecDampingDeviceViscous
      - p.431: 6.17.1.3. 粘弾性体ダンパー：StbSecDampingDeviceViscoelastic
      - p.432: 6.17.1.4. 履歴系ダンパー：StbSecDampingDeviceHistory
      - p.433: 6.17.1.5. 摩擦ダンパー：StbSecDampingDeviceFriction
      - p.434: 6.17.1.6. 質量系ダンパー：StbSecDampingDeviceMass
    - p.435: 6.17.2. 制振装置断面・鉄骨断面形状：StbSecSteelFigureDampingDevice
      - p.436: 6.17.2.1. 履歴系ダンパー・鉄骨断面形状：StbSecSteelFigureDampingDeviceHistory
    - p.437: 6.17.3. 制振装置断面・仕様指定：StbSecSpecificationDampingDevice
    - p.438: 6.17.4. 制振装置断面・接合部：StbSecConnectionDampingDevice
      - p.439: 6.17.4.1. 水平方向接合型接合部：StbSecConnectionDampingDeviceHorizontal
      - p.441: 6.17.4.2. 層上下接合型接合部：StbSecConnectionDampingDeviceVertical
- p.443: 06.18. RC基礎断面
  - p.443: 6.18. ＲＣ基礎断面：StbSecFoundation_RC
    - p.445: 6.18.1. ＲＣ基礎断面形状：StbSecFigureFoundation_RC
      - p.446: 6.18.1.1. ＲＣ基礎断面形状・矩形：StbSecFoundation_RC_Rect
      - p.447: 6.18.1.2. ＲＣ基礎断面形状・矩形テーパー：StbSecFoundation_RC_TaperedRect
      - p.449: 6.18.1.3. ＲＣ基礎断面形状・直角三角形：StbSecFoundation_RC_Triangle
      - p.451: 6.18.1.4. ＲＣ基礎断面形状・正三角形：StbSecFoundation_RC_EquiTriangle
      - p.452: 6.18.1.5. ＲＣ基礎断面形状・八角形：StbSecFoundation_RC_Octagon
      - p.453: 6.18.1.6. ＲＣ連続基礎断面形状：StbSecFoundation_RC_Continuous
    - p.454: 6.18.2. ＲＣ基礎断面配筋：StbSecBarArrangementFoundation_RC
      - p.456: 6.18.2.1. ＲＣ基礎断面配筋・矩形：StbSecBarFoundation_RC_Rect
      - p.459: 6.18.2.2. ＲＣ基礎断面配筋・三角：StbSecBarFoundation_RC_Triangle
      - p.461: 6.18.2.3. ＲＣ基礎断面配筋・三方：StbSecBarFoundation_RC_ThreeWay
      - p.463: 6.18.2.4. ＲＣ基礎断面配筋・連続：StbSecBarFoundation_RC_Continuous
- p.466: 06.19. RC杭断面
  - p.466: 6.19. ＲＣ杭断面：StbSecPile_RC
    - p.467: 6.19.1. ＲＣ杭断面　一般工法：StbSecPile_RC_Conventional
      - p.468: 6.19.1.1. ＲＣ杭断面形状　一般工法：StbSecFigurePile_RC_Conventional
        - p.470: 6.19.1.1.1. ＲＣ杭断面形状・ストレート：StbSecPile_RC_ConventionalStraight
        - p.471: 6.19.1.1.2. ＲＣ杭断面形状・脚部拡大：StbSecPile_RC_ConventionalExtendedFoot
        - p.472: 6.19.1.1.3. ＲＣ杭断面形状・頂部拡大：StbSecPile_RC_ConventionalExtendedTop
        - p.473: 6.19.1.1.4. ＲＣ杭断面形状・頂部脚部拡大：StbSecPile_RC_ConventionalExtendedTopFoot
      - p.474: 6.19.1.2. ＲＣ杭断面配筋：StbSecBarArrangementPile_RC_Conventional
        - p.476: 6.19.1.2.1. ＲＣ杭断面配筋・全断面：StbSecBarPile_RC_Same
        - p.478: 6.19.1.2.2. ＲＣ杭断面配筋・杭頭脚別：StbSecBarPile_RC_TopBottom
        - p.480: 6.19.1.2.3. ＲＣ杭断面配筋・杭頭軸部杭脚：StbSecBarPile_RC_TopCenterBottom
      - p.482: 6.19.1.3.  StbSecPile_RC_Connection：場所打ちコンクリート杭　杭頭接合
        - p.483: 6.19.1.3.1. StbSecPile_RC_ConnectionCertified：場所打ちコンクリート杭　杭頭接合　認定工法
          - p.484: 6.19.1.3.1.1. 認定工法特有の属性と値：StbCertifiedMethodKeyValue
    - p.485: 6.19.2. RC杭断面　認定工法：StbSecPile_RC_Certified
      - p.486: 6.19.2.1. ＲＣ杭断面形状　認定工法：StbSecFigurePile_RC_Certified
      - p.487: 6.19.2.2. ＲＣ杭断面配筋：StbSecBarArrangementPile_RC_Certified
- p.488: 06.20. 鋼管杭断面
  - p.488: 6.20. 鋼管杭断面：StbSecPile_S
    - p.489: 6.20.1. 鋼管杭　一般工法：StbSecPile_S_Conventional
      - p.490: 6.20.1.1. 鋼管杭断面形状：StbSecFigurePile_S
        - p.491: 6.20.1.1.1. 鋼管杭断面形状・ストレート：StbSecPile_S_Straight
        - p.492: 6.20.1.1.2. 鋼管杭断面形状・回転貫入杭（先端拡翼杭）：StbSecPile_S_Rotational
        - p.493: 6.20.1.1.3. 鋼管杭断面形状・テーパー管杭：StbSecPile_S_Taper
        - p.494: 6.20.1.1.4. 鋼管杭断面形状・製品指定：StbSecPile_S_Product
      - p.495: 6.20.1.2. 鋼管杭継手：StbSecPile_S_Joint
        - p.496: 6.20.1.2.1. 鋼管杭溶接継手：StbSecPile_S_JointWeld
        - p.497: 6.20.1.2.2. 鋼管杭機械式継手：StbSecPile_S_JointMecanical
      - p.498: 6.20.1.3. 鋼管杭杭頭接合：StbSecPile_S_Connection
    - p.499: 6.20.2. 鋼管杭　認定工法：StbSecPile_S_Certified
- p.500: 06.21. 既製コンクリート杭断面
  - p.500: 6.21. 既製コンクリート杭断面：StbSecPilePrecast
    - p.501: 6.21.1. 既製コンクリート　一般工法：StbSecPilePrecastConventional
      - p.502: 6.21.1.1. 既成コンクリート杭断面形状：StbSecFigurePilePrecast
        - p.503: 6.21.1.1.1. 既製コンクリート杭断面形状・PHC杭：StbSecPilePrecast_PHC
        - p.504: 6.21.1.1.2. 既製コンクリート杭断面形状・ST杭：StbSecPilePrecast_ST
        - p.505: 6.21.1.1.3. 既製コンクリート杭断面形状・SC杭：StbSecPilePrecast_SC
        - p.506: 6.21.1.1.4. 既製コンクリート杭断面形状・PRC杭：StbSecPilePrecast_PRC
        - p.507: 6.21.1.1.5. 既製コンクリート杭断面形状・CPRC杭：StbSecPilePrecast_CPRC
        - p.508: 6.21.1.1.6. 既製コンクリート杭断面形状・節付PHC杭：StbSecPilePrecastNodular_PHC
        - p.509: 6.21.1.1.7. 既製コンクリート杭断面形状・節付PRC杭：StbSecPilePrecastNodular_PRC
        - p.510: 6.21.1.1.8. 既製コンクリート杭断面形状・節付CPRC杭：StbSecPilePrecastNodular_CPRC
        - p.511: 6.21.1.1.9. 既成コンクリート杭断面形状・製品指定：StbSecPileConventionalPrecastProduct
      - p.512: 6.21.1.2. 継手：StbSecPilePrecastJoint
        - p.513: 6.21.1.2.1. 溶接継手：StbSecPilePrecastJointWeld
        - p.514: 6.21.1.2.2. 鋼管杭機械式継手：StbSecPilePrecastJointMecanical
      - p.515: 6.21.1.3. 既製杭　杭頭接合：StbSecPilePrecastConnection
        - p.516: 6.21.1.3.1. 既製杭　杭頭接合　在来工法：StbSecPilePrecastConnectionConventional
        - p.517: 6.21.1.3.2. 既製杭　杭頭接合　認定工法：StbSecPilePrecastConnectionCertified
    - p.518: 6.21.2. 既製コンクリート杭　認定工法：StbSecPilePrecastCertified
- p.519: 06.22. RCパラペット断面
  - p.519: 6.22. ＲＣパラペット断面：StbSecParapet_RC
    - p.521: 6.22.1. ＲＣパラペット断面形状：StbSecFigureParapet_RC
      - p.522: 6.22.1.1. ＲＣパラペット断面形状・L型：StbSecParapet_RC_TypeL
      - p.523: 6.22.1.2. ＲＣパラペット断面形状・T型：StbSecParapet_RC_TypeT
      - p.524: 6.22.1.3. ＲＣパラペット断面形状・I型：StbSecParapet_RC_TypeI
    - p.525: 6.22.2. ＲＣパラペット断面配筋：StbSecBarArrangementParapet_RC
      - p.527: 6.22.2.1. ＲＣパラペット断面配筋・シングル：StbSecBarParapet_RC_Single
      - p.528: 6.22.2.2. ＲＣパラペット断面配筋・千鳥：StbSecBarParapet_RC_Zigzag
      - p.529: 6.22.2.3. ＲＣパラペット断面配筋・ダブル：StbSecBarParapet_RC_DoubleNet
      - p.530: 6.22.2.4. パラペット先端補強筋（アゴ筋）：StbSecBarParapet_RC_Tip
      - p.531: 6.22.2.5. 端部補強筋：StbSecBarParapet_RC_Edge
- p.532: 06.23. RC開口断面
  - p.532: 6.23. ＲＣ開口断面：StbSecOpen_RC
    - p.534: 6.23.1.  ＲＣ開口断面配筋：StbSecBarArrangementOpen_RC
      - p.535: 6.23.1.1. ＲＣスラブ開口配筋：StbSecBarOpen_RC_Slab
      - p.536: 6.23.1.2. ＲＣ壁開口配筋：StbSecBarOpen_RC_Wall
- p.537: 06.24. S梁貫通断面
  - p.537: 6.24. S梁貫通孔補強仕様：StbSecPenetration_S
    - p.539: 6.24.1. ハイリング拡張情報：StbSecPenetration_S_HiringExtension
    - p.540: 6.24.2. EGリング拡張情報：StbSecPenetration_S_EGringExtension
    - p.541: 6.24.3. フリードーナツ拡張情報：StbSecPenetration_S_FRringExtension
    - p.542: 6.24.4. OSリング拡張情報：StbSecPenetration_S_OSringExtension
- p.543: 06.25. 柱梁接合部断面
  - p.543: 6.25. 柱梁接合部断面：StbSecPanelZone
    - p.544: 6.25.1. 柱梁接合部断面形状：StbSecFigurePanelZone
      - p.545: 6.25.1.1. 柱梁接合部断面形状・矩形：StbSecPanelZoneRect
      - p.546: 6.25.1.2. 柱梁接合部断面形状・円形：StbSecPanelZoneCircle
- p.547: 06.26. 鉄骨断面
  - p.547: 6.26. 鉄骨断面：StbSecSteel
    - p.549: 6.26.1.  H形鋼：StbSecRoll-H
    - p.550: 6.26.2. 組立H形鋼：StbSecBuild-H
    - p.551: 6.26.3. 組立H形鋼（上下異なる）：StbSecBuild-HAssymmetric
    - p.552: 6.26.4. 角形鋼管：StbSecRoll-BOX
    - p.553: 6.26.5. 組立角形鋼管：StbSecBuild-BOX
    - p.554: 6.26.6. 円形鋼管：StbSecPipe
    - p.555: 6.26.7. T形鋼：StbSecRoll-T
    - p.556: 6.26.8. 組立T形鋼：StbSecBuild-T
    - p.557: 6.26.9. 溝形鋼：StbSecRoll-C
    - p.558: 6.26.10.  溝形鋼(2丁)：StbSecRoll-2C
    - p.559: 6.26.11. 山形鋼：StbSecRoll-L
    - p.560: 6.26.12. 山形鋼(2丁)：StbSecRoll-2L
    - p.561: 6.26.13. リップ溝形鋼：StbSecLipC
    - p.562: 6.26.14. リップ溝形鋼(2丁)：StbSecLip2C
    - p.563: 6.26.15. フラットバー：StbSecFlatBar
    - p.564: 6.26.16. 丸鋼：StbSecRoundBar
    - p.565: 6.26.17.  鉄骨製品：StbSecSteelProduct
    - p.566: 6.26.18.  未定義鉄骨断面：StbSecSteelUndefined
- p.567: 06.27. 構造種別に依存しない断面
  - p.567: 6.27. 構造種別に依存しない断面：StbSecUndefined
- p.568: 07. 要素リファレンス 継手情報
  - p.568: 7. 要素リファレンス　継手情報
    - p.568: 7.1. 継手情報：StbJoints
      - p.569: 7.1.1. Ｓ梁継手・Ｈ形：StbJointBeamShapeH
        - p.570: 7.1.1.1. Ｈ形継手詳細：StbJointShapeH
        - p.571: 7.1.1.2. Ｈ形継手詳細・フランジ：StbJointShapeHFlange
          - p.573: 7.1.1.2.1.  Ｈ形継手詳細・フランジボルト詳細：StbJointShapeHFlangeBolt
        - p.574: 7.1.1.3. Ｈ形継手詳細・ウェブ：StbJointShapeHWeb
          - p.576: 7.1.1.3.1. Ｈ形継手詳細・ウェブボルト詳細：StbJointShapeHWebBolt
      - p.578: 7.1.2. Ｓ柱継手・Ｈ形：StbJointColumnShapeH
      - p.579: 7.1.3. Ｓ柱継手・Ｔ形：StbJointColumnShapeT
        - p.581: 7.1.3.1. Ｔ形継手詳細：StbJointShapeT
        - p.582: 7.1.3.2. Ｔ形継手詳細・Ｈ部分フランジ：StbJointShapeTFlangeH
          - p.584: 7.1.3.2.1. Ｔ形継手詳細・フランジボルト詳細：StbJointShapeTFlangeBolt
        - p.585: 7.1.3.3. Ｔ形継手詳細・Ｈ部分ウェブ(長)：StbJointShapeTWebHLong
        - p.587: 7.1.3.4. Ｔ形継手詳細・ウェブボルト詳細：StbJointShapeTWebBolt
        - p.588: 7.1.3.5. Ｔ形継手詳細・Ｈ部分ウェブ(短)：StbJointShapeTWebHShort
        - p.590: 7.1.3.6. Ｔ形継手詳細・Ｔ部分フランジ：StbJointShapeTFlangeT
        - p.592: 7.1.3.7. Ｔ形継手詳細・Ｔ部分ウェブ：StbJointShapeTWebT
      - p.594: 7.1.4. Ｓ柱継手・＋形：StbJointColumnShapeCross
        - p.596: 7.1.4.1. ＋形継手詳細：StbJointShapeCross
        - p.597: 7.1.4.2. ＋形継手詳細・T部分フランジ（長）：StbJointShapeCrossTFlangeLong
          - p.599: 7.1.4.2.1. ＋形継手詳細・フランジボルト詳細：StbJointShapeCrossFlangeBolt
        - p.600: 7.1.4.3. ＋形継手詳細・T部分フランジ（短）：StbJointShapeCrossTFlangeShort
        - p.602: 7.1.4.4. ＋形継手詳細・T部分ウェブ(長)：StbJointShapeCrossTWebLong
          - p.604: 7.1.4.4.1. ＋形継手詳細・ウェブボルト詳細：StbJointShapeCrossWebBolt
        - p.605: 7.1.4.5. ＋形継手詳細・T部分ウェブ(短)：StbJointShapeCrossTWebShort
        - p.607: 7.1.4.6. ＋形継手詳細・H部分フランジ：StbJointShapeCrossHFlange
        - p.609: 7.1.4.7. ＋形継手詳細・H部分ウェブ(長)：StbJointShapeCrossHWebLong
        - p.611: 7.1.4.8. ＋形継手詳細・H部分ウェブ(短)：StbJointShapeCrossHWebShort
- p.613: 08. 要素リファレンス コネクション情報
  - p.613: 8. 要素リファレンス　コネクション情報
    - p.613: 8.1. コネクション情報：StbConnections
    - p.614: 8.2. ガセットプレート詳細（複数）：StbGussetPlates
      - p.615: 8.2.1. ガセットプレート詳細：StbGussetPlate
        - p.617: 8.2.1.1. ボルト詳細：StbGussetPlateBoltArray
        - p.618: 8.2.1.2. スプライスプレート詳細：StbGussetPlateSplice
    - p.619: 8.3. リブプレート詳細（複数）：StbRibPlates
      - p.620: 8.3.1. リブプレート詳細：StbRibPlate
    - p.621: 8.4. ダイアフラム詳細（複数）：StbDiaphragms
      - p.622: 8.4.1. ダイアフラム詳細：StbDiaphragm
    - p.623: 8.5. スチフナープレート詳細（複数）：StbStiffners
      - p.624: 8.5.1. スチフナープレート詳細：StbStiffner
- p.625: 09. 要素リファレンス 溶接情報
  - p.625: 9. 要素リファレンス　溶接情報
    - p.625: 9.1. 溶接：StbWeld
    - p.626: 9.2. 完全溶込み溶接：StbWeldFullPenetration
      - p.638: 9.2.1.  溶接仕様詳細：StbWeldSpec
    - p.639: 9.3. 部分溶け込み溶接：StbWeldPartialPenetration
    - p.641: 9.4. 隅肉溶接：StbWeldFillet
    - p.643: 9.5. フレア溶接：StbWeldFlare
- p.644: 10. 要素リファレンス 拡張情報
  - p.644: 10. 要素リファレンス　拡張情報
    - p.644: 10.1. 拡張情報（複数）：StbExtensions
    - p.646: 10.2. 拡張情報：StbExtension
      - p.647: 10.2.1. 対象オブジェクト：StbExtObject
        - p.648: 10.2.1.1. 拡張属性：StbExtProperty
      - p.649: 10.2.2. 拡張子要素：StbExtElement
        - p.650: 10.2.2.1. 拡張属性定義：StbExtPropertyDef
- p.651: 11. 要素リファレンス 出力情報
  - p.651: 11. 要素リファレンス　出力情報
    - p.651: 11.1. 出力情報：StbExportInformation
    - p.652: 11.2. 変換方針：StbExportPolicy
    - p.653: 11.3. 変換ログ：StbExportLog
      - p.654: 11.3.1. エラーログ：StbExportError
      - p.655: 11.3.2. 通知ログ：StbExportWarning
      - p.656: 11.3.3. 通知ログ：StbExportNotice
- p.657: 修正履歴
  - p.657: 【修正履歴】

## Page 1

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
i 
 
 
 
 
 
 
 
 
 
 
 
 
 
ST-Bridge XML ファイル仕様書（ver.2.1） 
 
 
 
 
 
 
 
2023. 03. 31 
 buildingSMART Japan 
構造設計小委員会

## Page 2

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
ii 
 
はじめに 
 
「ST-Bridge」は、日本国内の建築構造分野のソフトウェア間のデータ交換、情報共有に利用することを目指
し、一般社団法人buildingSMART Japan 構造設計小委員会（旧 一般社団法人IAI 日本 構造分科会）
が開発している標準データ交換形式である。 
 
ST-Bridge は2012 年7 月にver.1.0 をリリースした後、いろいろな場面で活用が広がる中で、適用範囲を拡
張してきたが、今回のver.2.0 はこれまでの拡張だけでなく、XML としての厳密性やデータ交換精度の向上
を目的とした改定である。 
 
ご協力いただいた構造設計小委員会各社、特にSTB 普及WG のメンバーの方々に心から感謝をしたい。 
 
2018 年7 月1 日

## Page 3

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
  
 
目次 
1. 概要 .............................................................................................................................................................. 1-1 
1.1. コンセプト .......................................................................................................................................... 1-1 
1.2. XML 形式 ............................................................................................................................................. 1-2 
1.3. 表記法 .................................................................................................................................................. 1-3 
1.4. 命名規則 .............................................................................................................................................. 1-4 
1.5. 属性値の型と表現範囲 ........................................................................................................................ 1-4 
1.6. 単位系と数値範囲................................................................................................................................ 1-4 
1.7. グローバル一意識別子（GUID） ....................................................................................................... 1-5 
1.8. 集合型（monolist） ............................................................................................................................ 1-5 
1.9. ファイルの拡張子................................................................................................................................ 1-6 
1.10. バージョン番号 ................................................................................................................................. 1-6 
1.11. 名前空間（Namespace） .................................................................................................................. 1-6 
1.12. XML Schema の利用 ........................................................................................................................ 1-6 
1.13. 製品型番一覧表 ................................................................................................................................. 1-7 
1.14. 工法一覧表 ........................................................................................................................................ 1-7 
2. 要素リファレンス 特記事項 ...................................................................................................................... 2-1 
2.1. 全体構成 .............................................................................................................................................. 2-1 
2.2. 主要な要素のID と一意性 .................................................................................................................. 2-3 
2.3. 部材の配置 .......................................................................................................................................... 2-5 
2.4. 材料の表記 .......................................................................................................................................... 2-6 
2.5. 鉄筋コンクリート部材における鉄筋の重心位置とかぶり厚さについて ............................................ 2-7 
2.6. 座標系 .................................................................................................................................................. 2-8 
2.7. 仕様書の見方 ..................................................................................................................................... 2-10 
2.8. ST-Bridge で同じ内容を示す場合の優先順位と推奨出力について .................................................. 2-12 
3. 要素リファレンス 共通情報 ...................................................................................................................... 3-1 
3.1. ST-Bridge：ST_BRIDGE ................................................................................................................... 3-1 
3.2. 共通情報：StbCommon ...................................................................................................................... 3-2 
3.2.1. 径別鉄筋強度情報リスト：StbReinforcementStrengthList ....................................................... 3-5 
3.2.2. 杭の径別鉄筋強度情報リスト：StbReinforcementStrengthListPile ......................................... 3-6 
3.2.3. 属性・条件適用リスト： StbApplyConditionsList .................................................................... 3-7 
3.2.4. システム標準板厚リスト：StbStandardPlateThicknessList ................................................... 3-30 
3.2.5. コネクションスペック情報：StbConnectionSpecs ................................................................... 3-31 
3.2.6. 溶接情報：StbWeldCommon ..................................................................................................... 3-47

## Page 4

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
i 
 
3.2.7. 追加情報：StbAdditionalInformation ...................................................................................... 3-48 
4. 要素リファレンス 位置情報 ...................................................................................................................... 4-1 
4.1. 位置・断面情報：StbModel ................................................................................................................ 4-1 
4.2. 節点（複数）：StbNodes ................................................................................................................... 4-2 
4.2.1. 節点：StbNode ............................................................................................................................. 4-3 
4.2.2. 節点ID リスト：StbNodeIdList ................................................................................................. 4-6 
4.2.3. 節点ID：StbNodeId .................................................................................................................... 4-7 
4.2.4. 順序のある節点ID：StbNodeIdOrder ........................................................................................ 4-8 
4.3. 軸（複数）：StbAxes ......................................................................................................................... 4-9 
4.3.1. 平行軸(複数)：StbParallelAxes ................................................................................................. 4-10 
4.3.2. 円弧軸(複数)：StbArcAxes ........................................................................................................ 4-13 
4.3.3. 放射軸(複数)：StbRadialAxes ................................................................................................... 4-15 
4.3.4. 作図用軸(複数)：StbDrawingAxes ............................................................................................ 4-17 
4.4. 階（複数）：StbStories ................................................................................................................... 4-20 
4.4.1. 階：StbStory .............................................................................................................................. 4-21 
5. 要素リファレンス 部材情報 ...................................................................................................................... 5-1 
5.1. 部材情報：StbMembers ..................................................................................................................... 5-1 
5.2. 柱（複数）：StbColumns .................................................................................................................. 5-3 
5.2.1. 柱：StbColumn ............................................................................................................................ 5-4 
5.3. 間柱（複数）：StbPosts .................................................................................................................. 5-10 
5.3.1. 間柱：StbPost ............................................................................................................................ 5-11 
5.4. 大梁（複数）：StbGirders ............................................................................................................... 5-13 
5.4.1. 大梁：StbGirder ........................................................................................................................ 5-14 
5.5. 小梁（複数）：StbBeams ................................................................................................................ 5-21 
5.5.1. 小梁：StbBeam .......................................................................................................................... 5-21 
5.6. ブレース（複数）：StbBraces ......................................................................................................... 5-24 
5.6.1. ブレース：StbBrace .................................................................................................................. 5-25 
5.7. スラブ（複数）：StbSlabs ............................................................................................................... 5-27 
5.7.1. スラブ：StbSlab ........................................................................................................................ 5-28 
5.8. 壁（複数）：StbWalls...................................................................................................................... 5-32 
5.8.1. 壁：StbWall ............................................................................................................................... 5-33 
5.9. 免震装置（複数）：StbIsolatingDevices ......................................................................................... 5-38 
5.9.1. 免震装置：StbIsolatingDevice .................................................................................................. 5-39 
5.10. 制振装置（複数）： StbDampingDevices .................................................................................... 5-44 
5.10.1. 制振装置：StbDampingDevice ............................................................................................... 5-46 
5.10.2. 制振装置配置位置（複数台数配置時）：StbDampingDevicePosition ................................... 5-51

## Page 5

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
ii 
 
5.11. 制振装置（フレーム）（複数）： StbFrameDampingDevices ................................................... 5-52 
5.11.1. 制振装置（フレーム）：StbFrameDampingDevice ............................................................... 5-53 
5.12. フーチング（複数）：StbFootings ................................................................................................ 5-62 
5.12.1. フーチング：StbFooting .......................................................................................................... 5-63 
5.13. 布基礎（複数）：StbStripFootings ............................................................................................... 5-65 
5.13.1. 布基礎：StbStripFooting ........................................................................................................ 5-66 
5.14. 杭基礎（複数）：StbPiles .............................................................................................................. 5-68 
5.14.1. 杭基礎：StbPile ....................................................................................................................... 5-69 
5.15. 基礎柱（複数）：StbFoundationColumns .................................................................................... 5-72 
5.15.1. 基礎柱：StbFoundationColumn ............................................................................................. 5-73 
5.16. パラペット（複数）：StbParapets ............................................................................................... 5-76 
5.16.1. パラペット：StbParapet ......................................................................................................... 5-77 
5.17. 開口配置情報（複数）：StbOpenArrangements .......................................................................... 5-79 
5.17.1. 開口配置情報：StbOpenArrangement ................................................................................... 5-80 
5.18. 梁貫通孔配置情報（複数）：StbPenetrationArrangements ........................................................ 5-82 
5.18.1. 梁貫通孔配置情報：StbPenetrationArrangement ................................................................. 5-83 
5.19. 継手配置情報（複数）：StbJointArrangements .......................................................................... 5-85 
5.19.1. 継手配置情報：StbJointArrangement .................................................................................... 5-86 
5.20. コンクリート柱梁接合部配置情報（複数）：StbPanelZoneArrangements................................. 5-87 
5.20.1. コンクリート柱梁接合部：StbPanelZoneArrangement ......................................................... 5-88 
5.21. コネクション配置情報（複数）：StbConnectionArrangements ................................................. 5-89 
5.21.1. コネクション配置情報：StbConnectionArrangement ........................................................... 5-90 
6. 要素リファレンス 断面情報 ................................................................................................................... 6.1-1 
6.1. 断面情報：StbSections .................................................................................................................... 6.1-1 
6.2. ＲＣ柱断面：StbSecColumn_RC .................................................................................................... 6.2-1 
6.2.1. ＲＣ柱断面形状：StbSecFigureColumn_RC ........................................................................... 6.2-3 
6.2.2. ＲＣ柱断面配筋：StbSecBarArrangementColumn_RC ......................................................... 6.2-6 
6.3. Ｓ柱断面：StbSecColumn_S ........................................................................................................... 6.3-1 
6.3.1. Ｓ柱断面鉄骨形状：StbSecSteelFigureColumn_S .................................................................. 6.3-3 
6.3.2. 鉄骨断面柱脚製品：StbSecBaseProduct .................................................................................. 6.3-8 
6.3.3. 鉄骨断面柱脚在来工法：StbSecBaseConventional ............................................................... 6.3-10 
6.4. ＳＲＣ柱断面：StbSecColumn_SRC .............................................................................................. 6.4-1 
6.4.1. ＳＲＣ柱断面形状：StbSecFigureColumn_SRC ..................................................................... 6.4-3 
6.4.2. ＳＲＣ柱断面配筋：StbSecBarArrangementColumn_SRC .................................................... 6.4-4 
6.4.3. ＳＲＣ柱断面鉄骨形状：StbSecSteelFigureColumn_SRC ..................................................... 6.4-5 
6.5. ＣＦＴ柱断面：StbSecColumn_CFT .............................................................................................. 6.5-1

## Page 6

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
iii 
 
6.5.1. ＣＦＴ柱断面鉄骨形状：StbSecSteelFigureColumn_CFT ...................................................... 6.5-3 
6.6. ＲＣ梁断面：StbSecBeam_RC ........................................................................................................ 6.6-1 
6.6.1. ＲＣ梁断面形状：StbSecFigureBeam_RC ............................................................................... 6.6-3 
6.6.2. ＲＣ梁断面配筋：StbSecBarArrangementBeam_RC ............................................................. 6.6-7 
6.7. Ｓ梁断面：StbSecBeam_S .............................................................................................................. 6.7-1 
6.7.1. Ｓ梁断面鉄骨形状：StbSecSteelFigureBeam_S ..................................................................... 6.7-3 
6.8. ＳＲＣ梁断面：StbSecBeam_SRC .................................................................................................. 6.8-1 
6.8.1. ＳＲＣ梁断面形状：StbSecFigureBeam_SRC ......................................................................... 6.8-3 
6.8.2. ＳＲＣ梁断面配筋：StbSecBarArrangementBeam_SRC ....................................................... 6.8-4 
6.8.3. ＳＲＣ梁断面鉄骨形状：StbSecSteelFigureBeam_SRC ......................................................... 6.8-5 
6.9. Ｓブレース断面：StbSecBrace_S ................................................................................................... 6.9-1 
6.9.1. Ｓブレース断面鉄骨形状：StbSecSteelFigureBrace_S ........................................................... 6.9-3 
6.10. ＲＣスラブ断面：StbSecSlab_RC ............................................................................................... 6.10-1 
6.10.1. ＲＣスラブ在来断面：StbSecSlab_RC_Conventional ......................................................... 6.10-3 
6.10.2. ＲＣスラブトラス断面：StbSecSlab_RC_Truss ................................................................ 6.10-26 
6.11. デッキ合成スラブ断面：StbSecSlabDeck .................................................................................. 6.11-1 
6.11.1. 合成デッキ製品：StbSecSlabDeckProduct ......................................................................... 6.11-3 
6.12. 既製スラブ断面：StbSecSlabPrecast ......................................................................................... 6.12-1 
6.12.1. 既製スラブトップ部分断面形状：StbSecFigureSlabPrecast .............................................. 6.12-3 
6.12.2. 既製スラブトップ部分断面配筋：StbSecBarArrangementSlabPrecast ............................ 6.12-4 
6.12.3. 既製スラブ製品：StbSecProductSlabPrecast ..................................................................... 6.12-6 
6.13. 荷重用スラブ断面：StbSecSlabLoad .......................................................................................... 6.13-1 
6.14. ＲＣ壁断面：StbSecWall_RC ..................................................................................................... 6.14-1 
6.14.1. ＲＣ壁断面形状：StbSecFigureWall_RC ............................................................................ 6.14-2 
6.14.2. ＲＣ壁断面配筋：StbSecBarArrangementWall_RC ........................................................... 6.14-6 
6.15. 荷重用壁断面：StbSecWallLoad ................................................................................................. 6.15-1 
 免震装置断面：StbSecIsolatingDevice ..................................................................................... 6.16-1 
6.16.1. 免震装置断面・製品：StbSecProductIsolatingDevice ......................................................... 6.16-3 
6.16.2. 免震装置断面・仕様指定：StbSecSpecificationIsolatingDevice ....................................... 6.16-16 
6.16.3. 免震装置断面・接合部：StbSecConnectionIsolatingDevice ............................................. 6.16-17 
 制振装置断面：StbSecDampingDevice .................................................................................... 6.17-1 
6.17.1. 制振装置断面・製品：StbSecProductDampingDevice ........................................................ 6.17-3 
6.17.2. 制振装置断面・鉄骨断面形状：StbSecSteelFigureDampingDevice ................................. 6.17-12 
6.17.3. 制振装置断面・仕様指定：StbSecSpecificationDampingDevice ....................................... 6.17-14 
6.17.4. 制振装置断面・接合部：StbSecConnectionDampingDevice ............................................. 6.17-15 
6.18. ＲＣ基礎断面：StbSecFoundation_RC ...................................................................................... 6.18-1

## Page 7

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
iv 
 
6.18.1. ＲＣ基礎断面形状：StbSecFigureFoundation_RC ............................................................. 6.18-3 
6.18.2. ＲＣ基礎断面配筋：StbSecBarArrangementFoundation_RC .......................................... 6.18-12 
6.19. ＲＣ杭断面：StbSecPile_RC ....................................................................................................... 6.19-1 
6.19.1. ＲＣ杭断面 一般工法：StbSecPile_RC_Conventional ...................................................... 6.19-2 
6.19.2. RC 杭断面 認定工法：StbSecPile_RC_Certified ............................................................. 6.19-20 
6.20. 鋼管杭断面：StbSecPile_S .......................................................................................................... 6.20-1 
6.20.1. 鋼管杭 一般工法：StbSecPile_S_Conventional ................................................................ 6.20-2 
6.20.2. 鋼管杭 認定工法：StbSecPile_S_Certified ...................................................................... 6.20-12 
6.21. 既製コンクリート杭断面：StbSecPilePrecast............................................................................ 6.21-1 
6.21.1. 既製コンクリート 一般工法：StbSecPilePrecastConventional ....................................... 6.21-2 
6.21.2. 既製コンクリート杭 認定工法：StbSecPilePrecastCertified ......................................... 6.21-19 
6.22. ＲＣパラペット断面：StbSecParapet_RC ................................................................................. 6.22-1 
6.22.1. ＲＣパラペット断面形状：StbSecFigureParapet_RC ........................................................ 6.22-3 
6.22.2. ＲＣパラペット断面配筋：StbSecBarArrangementParapet_RC ....................................... 6.22-7 
6.23. ＲＣ開口断面：StbSecOpen_RC ................................................................................................. 6.23-1 
6.23.1. ＲＣ開口断面配筋：StbSecBarArrangementOpen_RC ...................................................... 6.23-3 
6.24. S 梁貫通孔補強仕様：StbSecPenetration_S ............................................................................... 6.24-1 
6.24.1. ハイリング拡張情報：StbSecPenetration_S_HiringExtension .......................................... 6.24-3 
6.24.2. EG リング拡張情報：StbSecPenetration_S_EGringExtension .......................................... 6.24-4 
6.24.3. フリードーナツ拡張情報：StbSecPenetration_S_FRringExtension ................................. 6.24-5 
6.24.4. OS リング拡張情報：StbSecPenetration_S_OSringExtension ........................................... 6.24-6 
6.25. 柱梁接合部断面：StbSecPanelZone ........................................................................................... 6.25-1 
6.25.1. 柱梁接合部断面形状：StbSecFigurePanelZone .................................................................. 6.25-2 
6.26. 鉄骨断面：StbSecSteel ............................................................................................................... 6.26-1 
6.26.1. H 形鋼：StbSecRoll-H........................................................................................................... 6.26-3 
6.26.2. 組立H 形鋼：StbSecBuild-H ............................................................................................... 6.26-4 
6.26.3. 組立H 形鋼（上下異なる）：StbSecBuild-HAssymmetric ................................................ 6.26-5 
6.26.4. 角形鋼管：StbSecRoll-BOX ................................................................................................. 6.26-6 
6.26.5. 組立角形鋼管：StbSecBuild-BOX ........................................................................................ 6.26-7 
6.26.6. 円形鋼管：StbSecPipe .......................................................................................................... 6.26-8 
6.26.7. T 形鋼：StbSecRoll-T ............................................................................................................ 6.26-9 
6.26.8. 組立T 形鋼：StbSecBuild-T .............................................................................................. 6.26-10 
6.26.9. 溝形鋼：StbSecRoll-C ........................................................................................................ 6.26-11 
6.26.10. 溝形鋼(2 丁)：StbSecRoll-2C ........................................................................................... 6.26-12 
6.26.11. 山形鋼：StbSecRoll-L ....................................................................................................... 6.26-13 
6.26.12. 山形鋼(2 丁)：StbSecRoll-2L ............................................................................................ 6.26-14

## Page 8

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
v 
 
6.26.13. リップ溝形鋼：StbSecLipC .............................................................................................. 6.26-15 
6.26.14. リップ溝形鋼(2 丁)：StbSecLip2C ................................................................................... 6.26-16 
6.26.15. フラットバー：StbSecFlatBar ......................................................................................... 6.26-17 
6.26.16. 丸鋼：StbSecRoundBar ................................................................................................... 6.26-18 
6.26.17. 鉄骨製品：StbSecSteelProduct ........................................................................................ 6.26-19 
6.26.18. 未定義鉄骨断面：StbSecSteelUndefined ........................................................................ 6.26-20 
6.27. 構造種別に依存しない断面：StbSecUndefined ......................................................................... 6.27-1 
7. 要素リファレンス 継手情報 ...................................................................................................................... 7-1 
7.1. 継手情報：StbJoints .......................................................................................................................... 7-1 
7.1.1. Ｓ梁継手・Ｈ形：StbJointBeamShapeH.................................................................................... 7-2 
7.1.2. Ｓ柱継手・Ｈ形：StbJointColumnShapeH .............................................................................. 7-11 
7.1.3. Ｓ柱継手・Ｔ形：StbJointColumnShapeT ............................................................................... 7-12 
7.1.4. Ｓ柱継手・＋形：StbJointColumnShapeCross ........................................................................ 7-27 
8. 要素リファレンス コネクション情報 ........................................................................................................ 8-1 
8.1. コネクション情報：StbConnections .................................................................................................. 8-1 
8.2. ガセットプレート詳細（複数）：StbGussetPlates .......................................................................... 8-2 
8.2.1. ガセットプレート詳細：StbGussetPlate .................................................................................... 8-3 
8.3. リブプレート詳細（複数）：StbRibPlates ....................................................................................... 8-7 
8.3.1. リブプレート詳細：StbRibPlate ................................................................................................. 8-8 
8.4. ダイアフラム詳細（複数）：StbDiaphragms ................................................................................... 8-9 
8.4.1. ダイアフラム詳細：StbDiaphragm .......................................................................................... 8-10 
8.5. スチフナープレート詳細（複数）：StbStiffners ............................................................................ 8-11 
8.5.1. スチフナープレート詳細：StbStiffner ...................................................................................... 8-12 
9. 要素リファレンス 溶接情報 ...................................................................................................................... 9-1 
9.1. 溶接：StbWeld ................................................................................................................................... 9-1 
9.2. 完全溶込み溶接：StbWeldFullPenetration ...................................................................................... 9-2 
9.2.1. 溶接仕様詳細：StbWeldSpec .................................................................................................... 9-14 
9.3. 部分溶け込み溶接：StbWeldPartialPenetration ............................................................................ 9-15 
9.4. 隅肉溶接：StbWeldFillet ................................................................................................................. 9-17 
9.5. フレア溶接：StbWeldFlare ............................................................................................................. 9-19 
10. 要素リファレンス 拡張情報 .................................................................................................................. 10-1 
10.1. 拡張情報（複数）：StbExtensions ................................................................................................ 10-1 
10.2. 拡張情報：StbExtension ................................................................................................................ 10-3 
10.2.1. 対象オブジェクト：StbExtObject ........................................................................................... 10-4 
10.2.2. 拡張子要素：StbExtElement .................................................................................................. 10-6 
11. 要素リファレンス 出力情報 .................................................................................................................. 11-1

## Page 9

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
vi 
 
11.1. 出力情報：StbExportInformation ................................................................................................. 11-1 
11.2. 変換方針：StbExportPolicy ........................................................................................................... 11-2 
11.3. 変換ログ：StbExportLog ............................................................................................................... 11-3 
11.3.1. エラーログ：StbExportError .................................................................................................. 11-4 
11.3.2. 通知ログ：StbExportWarning ................................................................................................ 11-5 
11.3.3. 通知ログ：StbExportNotice .................................................................................................... 11-6 
【修正履歴】 ........................................................................................................................................................1

## Page 10

1-1 
 
1. 概要 
1.1. コンセプト 
ST-Bridge は、IFC のみでは表現が難しい日本国内の建築構造設計情報について、建築構造分野のソフト
ウェア間における橋渡しの実現を目指す標準フォーマットである。2.0.2 版以前は、日本国内の一貫構造計算
プログラムと汎用の応力解析プログラム、構造図作成プログラムとの連携に重点を置きながら、３次元オブ
ジェクトCAD や積算プログラムなどと、構造躯体に関する情報を連携することも想定していた。2.1 版から
は、設計から施工へのデータ連携や鉄骨製作段階でのデータ共有など、より建築施工に近い領域でのデータ
連携を担うことができるように仕様を拡張することとし、一般社団法人 日本建築構造技術者協会やBIM ラ
イブラリ技術研究組合の要望を取りまとめ、BIM ソフトや鉄骨CAD ソフトなどの連携も視野に入れた拡張
を行った。 
建築の構造設計者にとって連携に必要でありながら、IFC では表現が難しい情報の表現として、ST-Bridge
では特に以下の表現方法を採用している。 
・構造検討に必要な構造部材の接続関係明示 
構造部材（柱、梁・・）の端点位置（座標値）を「節点」として要素定義し、節点要素を介して 
構造部材の接続関係を明示的に表現 
・鉄筋コンクリート部材配筋の属性表現 
   鉄筋コンクリート部材について、建築の構造設計図における断面表のような、種類・径・本数 
またはピッチによる配筋の表現 
 
本仕様書は、建築構造設計情報のうち、建築の構造設計図に表現される範囲の躯体情報について主に表記
する。荷重、設計条件および解析モデルに関する事項の詳細は、別途「計算編」による。

## Page 11

1-2 
 
1.2. XML 形式 
ST-Bridge はXML 形式を採用している。 
XML (Extensible Markup Language) は、データ交換に使用可能なマークアップ言語を新たに作成するため
の基礎として使用できる、簡単で柔軟なテキスト形式の言語である。 XML は W3C (World Wide Web 
Consortium) のワーキンググループから発行された一連の勧告に基づいており、ST-Bridge もそれにならう。 
XML 形式を表す用語として、「要素」(Element)、「属性」(Attribute)、「内容」(Content) がある。 
「要素」「属性」「内容」の、実際のタグとの対応関係は、以下となる。 
<Element Attribute=“属性値”>Content</Element> 
「内容」がない場合は、以下としてよい。 
<Element Attribute=“属性値” /> 
 
内容部分に別の要素を表記し、要素を階層構造とすることができる。 
要素名は、大文字と小文字が区別され、文字かアンダースコア（_）で始まる必要がある。また、要素名に
は、文字、数字、ハイフン、アンダースコア、およびピリオドを含めることができる。 
属性値は、  ”  ”  または  ‘ ’  で囲まれた文字列とする。従って、属性値の文字列に「”」または「’」を
含んではならない。 
 
XML 形式に関する基本的なルールと、ST-Bridge における扱いを以下に列記する。 
 
・XML バージョン番号 
<?xml version="1.0"?> 
 必須である。 将来の XML バージョンでは番号が変わることがあるが、現在のバージョンは 1.0。 
 
・encoding 宣言 
<?xml version="1.0" encoding="UTF-8"?> 
 この属性は省略可能であるが、ST-Bridge では省略しないこととする。 使用する場合は、encoding 宣言
は XML 宣言中でバージョン情報の直後になければならず、既存の文字エンコードを示す値を含んでいる必
要がある。当面、”UTF-8”、 ”Shift_JIS” を適用対象とする。 
 
・XML コメント 
<!--     --> 
 ドキュメントの構造や注釈など、XML パーサーに対する内容でないものは、コメント内に含める。 
 コメントは <!-- で始まり --> で終わる。 
 
・空白（スペース） 
 W3C (World Wide Web Consortium) XML 仕様では、属性値の中を除き、すべての空白を維持する。 
 したがって、開発者はXML パーサーが空白をどのように処理するかを意識する必要がある。

## Page 12

1-3 
 
1.3. 表記法 
 要素の表記法は、原則キャメルケース（アッパーキャメルケース）を採用する。 
 
例 
<StbMembers> 
   <StbGirders> 
   </StbGirders> 
</StbMembers> 
 
 属性の表記法は、原則スネーク記法を採用し（「_（アンダースコア）」でつなぐ）、小文字とする。 
 
例 
<StbColumns> 
   <StbColumn id_node_bottom="1" id_node_top="2" id_section="1"> 
   </StbColumn> 
</StbColumns>

## Page 13

1-4 
 
1.4. 命名規則 
・ST-Bridge に関する「要素」名は原則として「Stb」で始める。「要素」名および「属性」名の表記は、原
則として前節による。ただし、視認性向上および強調表現上の観点より、以下の場合は除く。 
・構造種別は大文字とし、要素名に用いる場合は、前後をアンダースコア( _ )でつなぐ（RC, SRC, S 等） 
・方向および座標値を表す場合は大文字とする（X, Y, Z, H, V） 
・本数（N_）、鉄筋径（D）、円の直径（D）は大文字とする 
・boolean の属性のうち、否かどうかを示す値（isXxxx）はローワーキャメルケースとする 
・基礎柱において、基礎柱（FD）根巻柱（WR）の区別に大文字を用いる 
・既製コンクリート杭断面において、杭分類名は大文字とし、要素名に用いる場合は、前をアンダースコア 
( _ )でつなぐ（PHC, PRC, SC 等） 
・鉄骨断面において、鋼材断面タイプは大文字とする（BOX, H 等） 
・鉄骨断面において、鋼材の外形寸法に大文字を用いる（A, B 等） 
・SRC 断面および継手情報において、形鋼の偏心（offset）部位の区別に大文字を用いる（T,HY,HX） 
 
1.5. 属性値の型と表現範囲 
 属性値は、XML 形式では文字列であるが、ST-Bridge では属性ごとに型を定め、特記がない限り表現範囲
を以下として、文字列を解釈する。 
 
型 
意味および表現範囲 
string 
ST-Bridge においても文字列であり、様式はencoding 宣言による 
integer 
コンピュータ言語の「符号なし４バイト整数型」による 
double 
コンピュータ言語の「８バイト実数型」で、表記は固定小数点型式とする 
boolean 
コンピュータ言語の「ブーリアン型」で、”true”または”false”とする 
 
属性値は、特記がない限りnull 値、スペースのみの表記および型に合わない値による表記は認めない。 
 
1.6. 単位系と数値範囲 
 位置を表す座標値、長さおよび角度に関する単位は、特記がない限り以下とする。 
座標値 
mm 
長さ 
mm 
角度 
度（degree） 
 
特記がない限り、長さ（寸法）は、0 より大きい値とし、角度は、0～360 度未満の範囲の値として反時計回
りを正とする。例えば、Z 軸回りの角度の場合は、Z 軸を下から見てX 軸位置を0 度とした、時計回りの数
値で表す。

## Page 14

1-5 
 
1.7. グローバル一意識別子（GUID） 
 GUID (Globally Unique IDentifier) は、UUID (Universally Unique IDentifier) としても知られる、128
ビットの符号なし整数で、空間および時間において一意である識別子である。UUID の仕様は RFC (Request 
for Comments) 4122 に規定されている。表記は、ifc における表記法にならい、32 桁の16 進数値を文字列
表現した値とする（16 進数値の 'a' から 'f' は、小文字とする）。 
例 
<Element  guid="78fd87737db64372bf0e7ede42393577"/> 
 
1.8. 集合型（monolist） 
 順番を区別する必要のある値集合については、順番をスペース区切りで続けて表記する。読み取り側アプ
リケーションは、要素をスペースごとに切り分け、順に配列に格納することになる。 
特記がない限り、ST-Bridge においては、この表現は「内容」においてのみ利用することとし、内容のnull
値およびスペースのみの表記は認めない。 
例 
<monolist>100  101  102  103</monolist> 
 
（補足） 
 元来、XML 形式は冗長さを許容するものであるが、同一の性質を表す集合値についてはスペースで区切る
ことでコンパクト化を図るケースもよく見られる。ここでは、1 対多の対応関係を示すような場合での利用
を想定している。 
地図情報で利用されるKML（Keyhole Markup Language）形式の<coordinates>要素なども同じような考
え方を採用しているので、参考になる。 
 
<coordinates> 
-122.365662,37.826988,0 
-122.365202,37.826302,0 
-122.364581,37.82655,0 
-122.365038,37.827237,0 
-122.365662,37.826988,0 
</coordinates> 
 
（参考） 
https://developers.google.com/kml/documentation/kmlreference?hl=ja#coordinates

## Page 15

1-6 
 
1.9. ファイルの拡張子 
 ST-Bridge データのXML ファイルの拡張子は、「.stb」とする。 
 
1.10. バージョン番号 
 バージョン番号は、「.」で区切り、３つの数字に意味を持たせる。 
 
ver. X. Y. Z 
X : 大項目（情報アーキテクチャなど大きな項目が変更になった場合） 
Y : 中項目（要素や属性に変更があった場合） 
Z : 細項目（字句修正、説明追加など、リビジョンとして改訂があった場合） 
 
 改訂の段階に応じて、Z: 細項目をリビジョンとして規定する。ただし、通称として採用するのは、バージ
ョン表記のうち前の2 ケタ（ver. X.Y）とする。ファイル名など、識別に細項目が必要な場合は、X.Y.Z のよ
うに細項目を表記してもよい。 
 
1.11. 名前空間（Namespace） 
 XML 形式では、要素や属性の重複を避けるために、属する集合に名前を付けて有効な範囲を定める、名前
空間（Namespace）の概念が導入されている。 
名前空間は、特にXML Schema を使用する際に、１つのスキーマが１つの名前空間を持つことで名前の
重複を回避する際に用いられており、通常、親要素の属性として記述する。 
 ST-Bridge では、親要素<ST_BRIDGE>にて、XML Schema に関する属性を宣言し、ST-Bridge 以下の要
素や属性は、デフォルトの名前空間に属するように宣言する。すなわち、以下のように記述する。 
 
<ST_BRIDGE xmlns:xs="http://www.w3.org/2001/XMLSchema" 
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" version="2.0.2" 
xmlns="https://www.building-smart.or.jp/dl"> 
 
 名前空間に関する属性の記述は、省略不可とする。 
 
1.12. XML Schema の利用 
 ST-Bridge データの妥当性を保証するために、ST-Bridge 仕様を記述したXML Schema を使用する。 
仕様書と対応したXML Schema はbuildingSMART Japan のWeb サイトからダウンロードできる。

## Page 16

1-7 
 
1.13. 製品型番一覧表 
 Ver2.0.2 以前では既製品に対して、メーカー名と製品型番をキーにして製品を特定していたが、製品型番
に規定がなかったため、アプリケーション間の連携がとりにくい仕様となっていた。Ver2.1 ではアプリケー
ション間のデータ連携を正しく行うために製品型番一覧表を整備する。 
 メーカーヒアリングから、同じ製品型番でも製品仕様が異なる場合があることがわかったため、製品型番
一覧表では、「製品型番」と「リリース時期」の2 つのキーで製品を一意に指定する仕様とした。 
 製品型番一覧表はアプリケーション間での一意な製品受け渡しまでを受け持つものと考えている。製品の
詳細データはカタログを参考に入出力アプリケーションが確認して利用してもらいたい。そのため、製品型
番一覧表にはデータ追加のみを行い、製品の廃番等の情報更新は行わない。 
 Ver2.1.0 は試用期間と考え、製品型番一覧表はbSJ 加盟者と製品公開メーカーに対してのみ公開するが、
今後の改定では仕様書とあわせて製品型番一覧表も公開していく予定である。 
 
1.14. 工法一覧表 
 杭などの認定工法では、各工法で指定しなければいけない固有の属性を持つことがあるが、任意の属性に
対してST-Bridge の属性として準備しておくことは難しいため、工法名と固有属性をセットで工法一覧表と
して整備する。 
 製品型番一覧表と同様、Ver2.1.0 は試用期間と考え、工法一覧表はbSJ 加盟者と工法公開メーカーに対し
てのみ公開するが、今後の改定では仕様書とあわせて工法一覧表も公開していく予定である。

## Page 17

2-1 
 
2. 要素リファレンス 特記事項 
 
2.1. 全体構成 
 ST-Bridge データは、ルート要素を <ST_BRIDGE> とした階層構造であり、直下に以下の子要素を有す
る構成となっている。 
要素 
要素名 
説明 
備考 
共通情報 
StbCommon 
材料など、建物の共通情報 
 
位置・断面情報 
StbModel 
節点および構造躯体情報 
 
拡張情報 
StbExtensions 
ST-Bridge の要素に定義のない属性
や子要素をアプリケーションが独
自に拡張する際に利用 
 
出力情報 
StbExportInformation 
出力方針と出力ログ 
 
計算データ要素 
StbCalData 
構造計算に必要な荷重や設計条件
などを、StbModel を補足する形で
定義 
計算編参照 
解析モデル要素 
StbAnaModels 
骨組構造解析に即した、StbModel
とは別の節点・部材情報、および
StbModel との関連付けなどを定義 
計算編参照 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 <StbCommon> および <StbModel> は、必須であり、省略することはできない。 
本仕様書で扱う主要な要素の階層(has-a 関係) を抜粋したものを以下に示す。 
<ST_BRIDGE> 
<StbCommon> 
</StbCommon> 
<StbModel> 
</StbModel> 
<StbExtensions> 
</StbExtensions> 
<StbExportInformation> 
</StbExportInformation> 
<StbCalData> 
</StbCalData> 
<StbAnaModels> 
</StbAnaModels> 
</ST_BRIDGE> 
共通情報 
拡張情報 
位置・断面情報 
計算データ要素 
解析モデル要素 
「計算データ編」で
扱うデータ 
出力情報

## Page 18

2-2 
 
 
Stb**Axis 
ST_BRIDGE 
StbCommon 
StbNodes 
StbNode 
StbAxes 
StbModel 
StbExtensions 
StbStories 
StbMembers 
StbSections 
Stb**Axes 
StbNodeidList 
StbNodeid 
StbStory 
StbColumns 
StbPosts 
StbGirders 
StbWalls 
StbFootings 
StbStripFootings 
StbBeams 
StbPiles 
StbBraces 
StbFoundationColumns 
StbSecColumn_* 
StbSecBeam_* 
StbSecBrace_S 
StbSecWall****** 
StbSecFoundation_RC 
StbSecPile***** 
StbSecOpen_RC 
StbSlabs 
StbParapets 
StbSecSlab*** 
StbSecParapet_RC 
*：RC or S or SRC (or CFT when Column) 
**：Parallel or Arc or Radial 
***：_RC or Deck or Precast or Load 
****：same  as  parent  Element 
*****：_RC or _S or Precast 
******：_RC or Load 
StbFrameDampingD
evices 
StbSecFigure**** 
StbSecBarArrangement**** 
StbSecSteelFigure**** 
StbJoints 
If necessary 
StbExportInformation 
StbSecSteel 
StbSecUndefined 
StbSecIsolatingDevice 
StbSecDampingDevice 
StbWeld 
StbConnectionArrangements 
StbDampingDevices 
StbIsolatingDevices 
StbOpenArrangements 
StbConnections 
StbPanelZoneArrangements 
StbSecPenetration_S 
StbPenetrationArrangements 
StbJointArrangements 
StbSecPanelZone

## Page 19

2-3 
 
2.2. 主要な要素のID と一意性 
 
 ST-Bridge の主要な要素においては、それぞれ個々の要素を一意に特定するために、ID（属性名id）を定
義する。id はinteger 型の値（１以上の整数値）とし、必須の属性とする。以下の要素のID は同じ要素でID
が重複してはならない。 
 
要素名 
 
要素名 
StbNode 
 
StbSecColumn_RC 
StbParallelAxis 
 
StbSecColumn_S 
StbArcAxis 
 
StbSecColumn_SRC 
StbRadialAxis 
 
StbSecColumn_CFT 
StbStory 
 
StbSecBeam_RC 
StbColumn 
 
StbSecBeam_S 
StbPost 
 
StbSecBeam_SRC 
StbGirder 
 
StbSecBrace_S 
StbBeam 
 
StbSecSlab_RC 
StbBrace 
 
StbSecSlabDeck 
StbSlab 
 
StbSecSlabPrecast 
StbWall 
 
StbSecSlabLoad 
StbIsolatingDevice 
 
StbSecWall_RC 
StbDampingDevice 
 
StbSecWallLoad 
StbFrameDampingDevice 
 
StbSecIsolatingDevice 
StbFooting 
 
StbSecDampingDevice 
StbStripFooting  
 
StbSecFoundation_RC 
StbPile  
 
StbSecPile_RC 
StbFoundationColumn  
 
StbSecPile_S 
StbParapet  
 
StbSecPilePrecast 
StbOpenArrangements  
 
StbSecParapet_RC 
StbPenetrationArrangements 
 
StbSecOpen_RC 
StbJointArrangements 
 
StbSecPenetration_S 
StbConnectionArrangements 
 
StbSecPanelZone 
 
 
StbJointBeamShapeH 
 
 
StbJointColumnShapeH 
StbJointColumnShapeT 
StbJointColumnShapeCross

## Page 20

2-4 
 
 
 
StbConnections 
 
 
StbWeld 
 
 
StbSecUndefined 
 
 
StbSecSteel  ※1 
                         ※1  StbSecSteel は、id を持たず形状名（文字列）で一意とする。 
 
また、特に構造部材については、ifc その他、GUID で部材を特定するデータと連携することもあり得るた
め、ID を定義する要素には、同時に属性 としてGUID（属性名guid）を定義する。guid は必須の属性で
はないが、持たせる場合は、ST-Bridge データ全体で一意となる識別子の文字列表現とし、ST-Bridge を介
したデータ連携においてGUID が変化しないことが望ましい。ST-Bridge においては、一意の識別子は
GUID で統一したいが、現状、GUID を扱わないプログラムが多いため、当面の措置として整数値のid を
必須としている。

## Page 21

2-5 
 
2.3. 部材の配置 
節点について 
構造部材（柱、梁・・）の配置は「節点」要素で定義する。 
節点は、構造部材の接続関係を明示するために用いる仮想の点であり、属性として全体座標系の３次元座
標値とID（GUID）を持つ。 
構造部材は、部材種別ごとに、線材においては配置基準線、面材においては配置基準面をそれぞれ定める。
節点は、配置基準線（面）を結ぶ点と定義する。このとき、構造部材は、属性として、結ぶ節点のID を持つ。 
例えば、線材である梁の場合、持つ節点は始端、終端の２点となり、配置基準線は下図のように定義して
いる（座標系については、次々節参照）。 
 
節点と配置基準線                  節点と配置基準線による配置例 
 
節点を介した配置は、一貫構造計算プログラムなどにおいて重量の流れ方を評価する際の基準となり、応
力解析プログラムなどにおいて解析モデル化を行う場合に、骨組解析の節点を定める際の基準となる。 
 
部材のオフセットと基準点について 
配置基準線（面）で配置した構造部材の位置は、実際の位置とは異なる場合もある。その場合は、実際の
端点位置と、節点座標値との差を、構造部材が持つ属性「オフセット」で指定する。実際の端点を、部材の
基準点と呼ぶ。オフセットの数値の定義は、全体座標系による。 
なお、線材の材軸回転については、別の属性「回転角」で指定する。 
 
仕様書に記述される属性のうち、長さや相対距離については、節点位置からの場合と、基準点位置からの
場合があるので、注意が必要である。 
 
（A 終端）
（A 始端）
（B 始端）
（B 終端）
Ｚ
Y
X
梁Ａ
梁B
Ｚ
Y
X
梁B 座標系
梁A 座標系
配置基準線
節点
節点
ID（GUID）
X 座標値
Y 座標値
Z 座標値
ID（GUID）
X 座標値
Y 座標値
Z 座標値
ID（GUID）
・・・・
始端ID
終端ID
ID（GUID）
・・・・
始端ID
終端ID

## Page 22

2-6 
 
 
例えば、下図のように鉄筋コンクリート造の梁同志が接続する場合は、取付く梁Ｂの実際の端点位置は、
梁Ａの躯体側面と交差する位置（●印）と考えることができる。その場合は、梁Ｂの「始端基準点」は●印
の位置となり、「始端側オフセット」に、節点座標位置と梁Ｂ端点位置との材軸方向の距離を指定する。 
一方、梁Ｂにハンチがある場合は、仕様書ではハンチ位置（○印）は前記「●印」からの長さとしている
ので、「ハンチ位置（始端）」には、始端基準点位置から、オフセットを反映した材軸方向の距離を指定す
る。 
 
 
３次元オブジェクトCAD などと、構造躯体の実際の位置情報を連携する場合は、オフセットにて特定し
た実際の位置を用いることが可能である。 
但し、オフセットされた実際の端点であっても、例えば、上の梁では部材の形状が六面体の組合せである
ことを基本と考えており、斜交や部分的なフカシなど、実際に施工される躯体の状態は明示されていないの
で注意が必要である。これは、定義の範囲を構造検討において最低限必要な情報に留めているためであり、
それより詳しい情報の連携についてはifc などによる必要がある。 
 
2.4. 材料の表記 
コンクリート、鉄筋および鉄骨材料の表記方法は、日本の建築基準法による指定建築材料においては、下
記とする。 
コンクリート強度は、「FC○○」と表記し、○○は設計基準強度（N/mm2）とする。ただし、軽量コンク
リートの場合は「LC○○」と表記する。 
鉄筋強度および鉄骨強度は、JIS 規格適合品や大臣認定品における規格の呼称にならう。 
鉄筋の径は、異形鉄筋の場合は呼び名「D○○」とし、丸鋼の場合は「R○○」（○○は呼び径 [mm] ）と
する。ただし、スラブや壁などで径の異なる鉄筋を交互に配筋する場合は、2 種類の呼び名を合成して「D○
○D△△」と表記する。 
高強度せん断補強筋などの大臣認定品については、それぞれの製品の呼び名（「S○○」等）を用いてよい。 
指定建築材料の表記においては、○○の数値に、規格にふさわしくない値を用いてはならない。 
 
 
 
（A 終端）
（A 始端）
（B 始端）
（B 終端）
梁Ａ
梁B
Ｚ
Y
X
梁B 座標系
ID（GUID）
・・・・
始端ID
終端ID
・・・・
始端側オフセット
・・・・
ハンチ位置（始端）
ハンチ位置（始端）
始端側オフセット
始端基準点

## Page 23

2-7 
 
2.5. 鉄筋コンクリート部材における鉄筋の重心位置とかぶり厚さについて 
 
ST-Bridge の断面データで定義する、鉄筋コンクリート部材の鉄筋位置を表す配筋情報には「鉄筋のかぶ
り厚さ」と「鉄筋の重心位置」がある（下図）。多段筋の場合、「重心位置」は１段目の重心位置を示す。 
 
 
これらは、おもに一貫構造計算プログラムにおいて断面の耐力を計算するために用いる、コンクリート外
面からの鉄筋重心までの距離を表すために定める値である。「鉄筋のかぶり厚さ」を定めた場合は、かぶり
厚さと鉄筋の寸法からプログラムが重心位置を計算する場合が多い。配筋が多段にわたる場合の「段筋のあ
き」と「段筋重心間距離」の関係も同様である。 
ST-Bridge の断面データでは、どちらか一方を記述することを想定しており、両者の整合については特に
制限は設けていない。したがって、プログラムが両方を併記する場合は、整合性に関して扱いを明確にする
必要がある。 
また、構造図作成プログラムとの連携においては、ここに示す「鉄筋のかぶり厚さ」は、設計図書に表示
する「鉄筋のかぶり厚さ」と必ずしも整合しない場合があるので注意が必要である。 
 
 
 
かぶり厚さ
段筋重心間距離
重心位置
段筋のあき

## Page 24

2-8 
 
2.6. 座標系 
 全体座標系と各構造部材の座標系（部材座標系）は、下記とする。 
全体座標系と部材座標系を区別する場合、全体座標系は X
－
, Y
－
, Z
－
 のように上線付で表記する。 
・全体座標系・部材座標系ともに直交座標系とする。 
・「節点」は、全体座標系で表記する。 
・構造部材の部材座標系は、構造部材種別ごとに、節点を結ぶ配置基準線（面）に対して適用する。 
 【軸の定義】 
【部材の基準軸】  構造部材の始点節点から次点節点方向に向かう部材座標系の軸 
【部材の断面軸】  主に構造部材断面の一方の軸を決めるために必要な部材座標系の軸 
【対応断面軸】   主に構造部材断面の一方の軸を決めるために必要な全体座標系の軸 
  定義方法 
１． 構造部材の始点節点を原点とし、【部材の基準軸】を下表のように定める。 
２． 【部材の断面軸】は、【部材の基準軸】と、原点から全体座標の【対応断面軸】方向に伸ばした線（梁
の場合、下図の赤線）の面内にあるものとして定める。 
３． 右手系で【部材の基準軸】でも【部材の断面軸】でもない軸を残りの一軸と定める。 
 
     定義方法（梁の場合） 
 
 
 
 
 
 
 
 
構造部材種別ごとの部材の基準軸、部材の断面軸および全体座標系の対応断面軸 
構造部材 
部材の基準軸 
全体座標系と部材座標系の対応 
対応断面軸 
部材の断面軸 
梁、ブレース、布基礎 
X 
Z
－
 
Z 
柱、杭、免震装置 
Z 
Y
－
 
Y 
スラブ 
X 
Z
－
 
Z 
壁 
X 
Z
－
 
Y 
フーチング 
Z (鉛直上向きとする) 
Y
－
 
Y 
     制振装置は、当該節の「・補足」による。 
 
Z
－ 
X
－
 
Y
－
 
Z 
Y 
X 
部材の基準軸 
対応断面軸

## Page 25

2-9 
 
 ST-Bridge で定義される建物の位置に対し、他のモデルと位置関係を調整するために、共通のグローバル
座標を定め、その原点（プロジェクトにおける測量点など）との相対位置を使用する場合がある。 
その場合は、グローバル座標を定めるグローバル座標系（下図、右手直交座標系 X
－
G, Y
－
G, Z
－
G ）に対する、
ST-Bridge 全体座標系の原点との相対位置（ΔX, ΔY, ΔZ）およびZ
－
G まわりの角度（θ、反時計まわりを
正）を指定する。 
 
 
これらの指定は、共通情報：StbCommon にて行う。

## Page 26

2-10 
 
2.7. 仕様書の見方 
 本仕様書は、XML 要素ごとに以下のような書式で記述されている。 
 
 
X.X.X 要素データ：StbElem1 
 
・概要 
説明 ：要素のデータ 
親要素：StbElemParent1 
   
   
・属性 
属性名 
型 
必須 
説明 
補足 
attr1 
string 
○ 
属性１ 
※(1) 
attr2 
integer 
 
属性２ 
 
   
   
   
必須に「○」がある場合は、属性はこの要素に必須であり、属性名・属性値ともに省略 
できない。 
必須に「○」がない場合は、属性はこの要素に必須ではない（出現条件は必ず「補足」 
に記述する。） 
   
・内容 
内容 
型 
必須 
説明 
補足 
cont1 
[monolist] 
integer 
○ 
内容１ 
 
 
 
 
※ 内容がある場合、記述は必須であり、省略できない 
 
 
この要素データの解説を記述 
このXML 要素の親要素を記述 
XML 要素に属性（attribute）がある場合はこの表を、 
ない場合には「無し」を記述 
XML 要素に内容（content）がある場合はこの表を、 
ない場合には「無し」を記述 
欄内に書けない記述は (1)…等、番号
を付けて後の「・補足」に記述

## Page 27

2-11 
 
 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbElemChild1 
0 
1 
条件１ 
 
StbElemChild2 
0 
制限なし 
条件２ 
 
 
 
 
 
※ 最小回数：1 以上である子要素は必須であり、省略できない 
 
・補足 
・・・・・・・・ 
 
・例 
 
 
 
 
 
・条件付きで必須となる要素がある場合は「補足」「・補足」に記述する。 
・個数、回数に制限がある場合、数値の範囲に制限がある場合は「補足」「・補足」に記述する。 
  ・要素数が0 の場合、出力アプリケーション上で未定義か出力できないことを表す。 
  ・必須属性は要素を成立させるために必要な属性である。必須属性が出力できないアプリケーションは
仮の値を出力し、仮の値を出力したことをStbExportInformation で明示する必要がある。 
・1.6 節で制限がある数値（角度や長さなど）や制限範囲が明確な値をもつ属性については、XML Schema
で表に示す属性値よりも制限をかけている場合がある。 
 
必要に応じて記述 
XML 要素が子要素を持つ場合はこの表を、持たない
場合には「無し」を記述 
必要に応じて記述 
<StbElemParent1> 
      <StbElem1 attr1="any_string" attr2="1"> 
        <StbElemChild1 （略）> 
          （略） 
        </StbElemChild1> 
      </StbElem1> 
</StbElemParent1>

## Page 28

2-12 
 
2.8. ST-Bridge で同じ内容を示す場合の優先順位と推奨出力について 
 一部の属性については、建物全体・階ごと・断面ごと・部材ごとのように同一の属性に対して複数個所で
設定が可能である。優先度は(部材ごと) > (断面ごと) > (階ごと) >(建物全体)とし、値が異なる場合には優先
度の高い側の値を採用するものとする。 
 ある属性について建物全体で一回設定しても、すべての部材について同じ属性を設定してもST-Bridge が
表現する建物としては同じ意味を有するが、インポートするアプリケーションによっては扱いが異なる場合
があるため、過度に細かな単位に分解せず出力アプリケーションのデータ構造に沿った形での出力を推奨す
る。

## Page 29

3-1 
 
3. 要素リファレンス 共通情報 
 
3.1. ST-Bridge：ST_BRIDGE 
・概要 
説明 ：ST-Bridge 全体 
親要素： － 
 
・属性 
属性名 
型 
必須 
説明 
補足 
version 
string 
○ 
ST-Bridge のバージョン 
2.1.0  
XML Schema に関する属性については、第１節「名前空間」の項による。 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCommon 
1 
1 
共通情報 
 
StbModel 
1 
1 
位置・断面情報 
 
StbExtensions 
0 
1 
拡張情報 
 
StbExportInformation 
0 
1 
出力情報 
 
StbCalData 
0 
1 
計算データ要素 
計算編参照 
StbAnaModels 
0 
1 
解析モデル要素 
計算編参照

## Page 30

3-2 
 
3.2. 共通情報：StbCommon 
・概要 
説明 ：共通情報 
親要素：ST_BRIDGE 
 
・属性 
属性名 
型 
必須 
説明 
補足 
guid 
string 
 
グローバルID 
IFC のGUID を流用 
project_name 
string 
○ 
プロジェクト名 
 
app_name 
string 
○ 
アプリケーション名 
※(1) 
app_version 
string 
○ 
アプリケーションの
バージョン 
※(1) 
convert_app_name 
string 
 
変換プログラム名 
※(2) 
convert_app_version 
string 
 
変換プログラムのバ
ージョン 
※(2) 
strength_concrete 
string 
 
建物全体のコンクリ
ート強度 
例：FC24 
※(3) 
global_offset_X 
double 
 
グローバル座標系と
のずれ（ΔX） 
※(4) 
global_offset_Y 
double 
 
グローバル座標系と
のずれ（ΔY） 
※(4) 
global_offset_Z 
double 
 
グローバル座標系と
のずれ（ΔZ） 
※(4) 
global_rotation 
double 
 
グローバル座標系と
の回転角度（θ） 
※(4) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbReinforcementStrengthList 
0 
1 
径別鉄筋強度情報リスト 
 
StbReinforcementStrengthListPile 
0 
1 
径別鉄筋強度情報リスト
(杭) 
 
StbApplyConditionList 
0 
1 
要領リスト

## Page 31

3-3 
 
StbStandardPlateThicknessList 
0 
1 
システム標準板厚リスト 
 
StbConnectionSpecs 
0 
1 
コネクションスペック情
報 
 
StbWeldCommon 
0 
1 
溶接における共通情報 
 
StbAdditionalInformation 
0 
1 
追加情報 
 
 
 
・補足 
(1) 「アプリケーション」は、このファイルを作成したアプリケーションプログラムとする。 
(2) 「変換プログラム」は、アプリケーションプログラムの総合的な名称と、変換プログラムの名称が
異なることを強調したい場合に記述する。 
(3) コンクリート強度は、建物全体の定義を省略した場合、StbStory で階別に定義、またはコンクリ
ートを用いる部材の断面別（RC 柱-StbSecColumn_RC、SRC 柱-StbSecColumn_SRC、CFT 柱-
StbSecColumn_CFT 、RC 梁-StbSecBeam_RC 、SRC 梁-StbSecBeam_SRC 、RC スラブ-
StbSecSlab_RC、デッキ合成スラブ-StbSecSlabDeck、既製スラブ-StbSecSlabPrecast、RC 壁-
StbSecWall_RC、RC 基礎-StbSecFoundation_RC、RC 杭-StbSecPile_RC、RC パラペット-
StbSecParapet_RC）に定義する必要がある。コンクリートを用いる部材自体に定義がある場合は、
部材の定義を優先する。優先順位は、次ページ表の通りとなる。 
(4) グローバル座標系とST-Bridge 全体座標系との位置関係を指定する。グローバル座標系の扱いお
よび各指定値の内容は、第２節「座標系」の説明による。 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
<StbCommon guid="9fcb952bb06242e58b0f96aecfcbd770" project_name="Prj名" 
 app_name="アプリ名" app_version=”アプリのバージョン” 
strength_concrete="FC24"> 
<StbReinforcementStrengthList> 
<StbReinforcementStrength D="D10" strength="SD295A"/> 
・・・ 
<StbReinforcementStrength D="D19" strength="SD345"/> 
・・・ 
</StbReinforcementStrengthList> 
</StbCommon>

## Page 32

3-4 
 
 
コンクリート強度の持ち方
定着方法は以下の4種類とする。
1) 部材が持つconcrete_strength
2) 部材 → 断面(id) → 断面が持つconcrete_strength
3) 部材 → 断面(id) → 節点が所属する階(id) → 階が持つconcrete_strength
4) 全体(StbCommon) が持つconcrete_strength
優先順位は、それぞれの部材について下表となる。
〇
 : concrete_strength to be adopted
〇
〇
〇
〇
〇
〇
〇
〇
×
×
×
×
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
〇
〇
〇
〇
〇
〇
〇
×
×
×
×
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
〇
〇
〇
〇
〇
〇
〇
×
×
×
×
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
〇
〇
〇
〇
〇
〇
〇
×
×
×
×
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
〇
×
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
×
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
〇
〇
×
×
〇
〇
×
×
〇
〇
〇
×
○ : concrete_strength is exist
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
〇
〇
×
×
〇
〇
×
×
〇
〇
〇
×
〇
〇
×
×
〇
×
〇
×
〇
〇
〇
〇
×
×
×
×
〇
〇
〇
〇
×
×
×
×
〇
〇
×
×
〇
〇
×
×
〇
×
〇
×
〇
×
〇
×
〇
〇
〇
〇
×
×
×
×
〇
〇
×
×
〇
〇
×
×
〇
×
〇
×
〇
×
〇
×
3) id_node_top → StbStory
柱 : StbColumn, 間柱 : StbPost
1) concrete_strength
2) id_section → StbSecColumn_RC(SRC,CFT)
3) id_node_top → StbStory
4) StbCommon
○ : concrete_strength is exist
大梁 : StbGirder, 小梁 : StbBeam
○ : concrete_strength is exist
1) concrete_strength
2) id_section → StbSecBeam_RC(SRC)
3) last node of StbNodeIdOrder.id → StbStory
4) StbCommon
スラブ : StbSlab
○ : concrete_strength is exist
1) concrete_strength
2) id_section → StbSecSlab_RC(Deck,Precast
3) first node of StbNodeIdOrder.id → StbStory
4) StbCommon
壁 : StbWall
○ : concrete_strength is exist
1) concrete_strength
2) id_section → StbSecWall_RC
○ : concrete_strength is exist
布基礎 : StbStripFooting
2) id_section → StbSecFoundation_RC
3) id_node_start → StbStory
4) StbCommon
フーチング : StbFooting
2) id_section → StbSecFoundation_RC
3) id_node → StbStory
4) StbCommon
4) StbCommon
杭基礎 : StbPile
2) id_section → StbSecPile_RC
4) StbCommon
○ : concrete_strength is exist
基礎柱 : StbFoundationColumn
○ : concrete_strength is exist
2) id_section_WR → StbSecColumn_RC
3) id_node → StbStory
4) StbCommon
2) id_section_FD → StbSecColumn_RC
パラペット : StbParapet
○ : concrete_strength is exist
2) id_section → StbSecParapet_RC
3) id_node_start → StbStory
4) StbCommon

## Page 33

3-5 
 
3.2.1.  径別鉄筋強度情報リスト：StbReinforcementStrengthList 
・概要 
説明 ：径別鉄筋強度 
親要素：StbCommon 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbReinforcementStrength 
1 
制限なし 
径別鉄筋強度情報 
 
 
・補足 
  杭に用いる鉄筋強度は上部構造と異なることがあるため、杭とそれ以外の鉄筋を分けて定義する。 
 
3.2.1.1.   径別鉄筋強度情報：StbReinforcementStrength 
・概要 
説明 ：径別鉄筋強度情報 
親要素：StbReinforcementStrengthList, StbReinforcementStrengthListPile 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
○ 
鉄筋径 
例：D25 
strength 
string 
○ 
SD(鉄筋強度) 
例：SD345 
 
・内容 
無し 
 
・子要素 
無し

## Page 34

3-6 
 
3.2.2. 杭の径別鉄筋強度情報リスト：StbReinforcementStrengthListPile 
・概要 
説明 ：杭の径別鉄筋強度 
親要素：StbCommon 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbReinforcementStrength 
1 
制限なし 
径別鉄筋強度情報 
 
 
・補足 
  杭に用いる鉄筋強度は上部構造と異なることがあるため、杭とそれ以外の鉄筋を分けて定義する。

## Page 35

3-7 
 
3.2.3. 属性・条件適用リスト： StbApplyConditionsList 
・概要 
説明 ：属性・条件適用リスト 
親要素：StbCommon 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbApplyConditionList_RC 
0 
1 
RC 要領 
 
StbApplyConditionList_S 
0 
1 
鉄骨要領 
 
 
・補足 
  要領として指定する項目をコンクリート・鉄骨ごとに定義する。 
  指定内容が細かなほど優先順位が高いものとし、階定義・断面定義・部材定義に記載がある場合はより
指定内容が細か詳細なものを優先する。 
  現段階でどこまで詳細な定義が必要かわからない項目については、string 値として利用者に委ねること
にした。次のバージョン以降に実情を反映して詳細定義を行うこととする。 
 
・例 
 以下の例は、配筋のデータについてそれぞれ、RC 柱およびRC 梁は適用対象で初期値を規定し、RC 杭は適用対象で
初期値を規定せず、RC 床他は適用対象としない場合を示す。

## Page 36

3-8 
 
 
 
  <StbCommon （略）> 
<StbReinforcementStrengthList> 
（略） 
</StbReinforcementStrengthList> 
<StbApplyConditionsList> 
  <StbApplyConditionList_RC> 
    <StbApply_RC_General comment='コメント'/> 
    <StbApply_RC_Pile comment='コメント'/> 
    <StbApply_RC_Foundation comment='コメント'/> 
    <StbApply_RC_FoundationGirder isOnSite='true' 
                depth_cover_left='50’ depth_cover_right='50’ depth_cover_top='40’ 
                depth_cover_bottom='50’ interval='70’ 
                D_bar_spacing='D10’ pitch_bar_spacing='1000'/> 
    <StbApply_RC_FoundationBeam isOnSite='true' 
                depth_cover_left='40’ depth_cover_right='40’ depth_cover_top='40’ 
                depth_cover_bottom='50’ interval='70’ 
                D_bar_spacing='D10’ pitch_bar_spacing='1000'/> 
    <StbApply_RC_Column isOnSite='true' 
                depth_cover_start_X='40’ depth_cover_end_X='40’ 
                depth_cover_start_Y='40’ depth_cover_end_Y='40’ interval='70’ 
                D_bar_spacing='D10’ pitch_bar_spacing='1000'/> 
    <StbApply_RC_Girder isOnSite='true' 
                depth_cover_left='40’ depth_cover_right='40’ depth_cover_top='40’ 
                depth_cover_bottom='40’ interval='70’/> 
    <StbApply_RC_Beam isOnSite='true' 
                depth_cover_left='40’ depth_cover_right='40’ depth_cover_top='40’ 
                depth_cover_bottom='40’ interval='70’/> 
    <StbApply_RC_Wall depth_cover='40’ D_bar_spacing='D10' pitch_bar_spacing='1000' 
                horizontal_bar_switch_start_ratio='1/4' 
                horizontal_bar_switch_end_ratio='1/4' 
                vertical_bar_switch_start_ratio='1/4' 
                vertical_bar_switch_end_ratio='1/4'/> 
  </StbApplyConditionList_RC> 
</StbApplyConditionsList> 
</StbCommon>

## Page 37

3-9 
 
3.2.3.1. StbApplyConditionList_RC：RC 要領 
・概要 
説明 ：RC 要領 
親要素：StbApplyConditionsList 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbApply_RC_General 
0 
1 
一般事項 
 
StbApply_RC_Pile 
0 
1 
杭 
 
StbApply_RC_Foundation 
0 
1 
基礎 
 
StbApply_RC_FoundationGirder 
0 
1 
基礎大梁 
 
StbApply_RC_FoundationBeam 
0 
1 
基礎小梁 
 
StbApply_RC_Column 
0 
1 
柱 
 
StbApply_RC_Girder 
0 
1 
大梁 
 
StbApply_RC_Beam 
0 
1 
小梁 
 
StbApply_RC_Wall 
0 
1 
壁 
 
StbApply_RC_Slab 
0 
1 
スラブ 
 
 
 
・補足 
(1) 上記のうち１種類以上の子要素を持つものとし、全子要素の最小回数が0 であってはならない。

## Page 38

3-10 
 
3.2.3.1.1. StbApply_RC_General：RC 要領 一般事項 
・概要 
説明 ：RC 要領 一般事項 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
comment 
string 
〇 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 鉄筋の加工・組立、定着長等をあたえることを想定している。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 39

3-11 
 
3.2.3.1.2. StbApply_RC_Pile：RC 要領 杭 
・概要 
説明 ：RC 要領 杭 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
comment 
string 
〇 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 既製杭・場所打ちコンクリート杭の要領を想定している。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 40

3-12 
 
3.2.3.1.3. StbApply_RC_Foundation：RC 要領 基礎 
・概要 
説明 ：RC 要領 基礎 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
comment 
string 
〇 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 独立基礎・布基礎の要領を想定している。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 41

3-13 
 
3.2.3.1.4. StbApply_RC_FoundationGirder：RC 要領 基礎大梁 
・概要 
説明 ：RC 要領 基礎大梁 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isOnSite 
boolean 
 
現場打ちかどうか 
 
depth_cover_left 
double 
 
かぶり厚さ（左） 
 
depth_cover_right 
double 
 
かぶり厚さ（右） 
 
depth_cover_top 
double 
 
かぶり厚さ（上） 
 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
 
interval 
double 
 
２段筋のあき 
 
center_top 
double 
 
主筋重心位置（上） 
 
center_bottom 
double 
 
主筋重心位置（下） 
 
center_side 
double 
 
主筋重心位置（側） 
 
center_interval 
double 
 
２段筋重心間距離 
 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
allocation_rule_stirrup 
string 
 
帯筋：割り付けルール 
 
D_additional 
string 
 
補助筋径 
 
strength_additional 
string 
 
補助筋強度 
 
number_rule_additional 
string 
 
補助筋本数ルール 
※(2) 
anchorage_rule 
string 
 
定着ルール 
※(2) 
cut_off_rule 
string 
 
カットオフルール 
※(2) 
figure_switch_start_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_始端側 
※(3) 
figure_switch_end_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_終端側 
※(3) 
bar_switch_start_ratio 
double 
 
3 断面配筋の切替位置
(比)_始端側 
※(4) 
bar_switch_end_ratio 
double 
 
3 断面配筋の切替位置
(比)_終端側 
※(4) 
comment 
string 
 
コメント

## Page 42

3-14 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbGirder のkind_structure がRC またはSRC でかつisFoundation がtrue の部材とその断面
に適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。 
(3) figure_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定す
る。figure_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定
する。 
(4) bar_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定する。
bar_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定する。

## Page 43

3-15 
 
3.2.3.1.5. StbApply_RC_FoundationBeam：RC 要領 基礎小梁 
・概要 
説明 ：RC 要領 基礎小梁 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isOnSite 
boolean 
 
現場打ちかどうか 
 
depth_cover_left 
double 
 
かぶり厚さ（左） 
 
depth_cover_right 
double 
 
かぶり厚さ（右） 
 
depth_cover_top 
double 
 
かぶり厚さ（上） 
 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
 
interval 
double 
 
２段筋のあき 
 
center_top 
double 
 
主筋重心位置（上） 
 
center_bottom 
double 
 
主筋重心位置（下） 
 
center_side 
double 
 
主筋重心位置（側） 
 
center_interval 
double 
 
２段筋重心間距離 
 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
allocation_rule_stirrup 
string 
 
帯筋：割り付けルール 
 
D_additional 
string 
 
補助筋径 
 
strength_additional 
string 
 
補助筋強度 
 
number_rule_additional 
string 
 
補助筋本数ルール 
※(2) 
anchorage_rule 
string 
 
定着ルール 
※(2) 
cut_off_rule 
string 
 
カットオフルール 
※(2) 
figure_switch_start_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_始端側 
※(3) 
figure_switch_end_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_終端側 
※(3) 
bar_switch_start_ratio 
double 
 
3 断面配筋の切替位置
(比)_始端側 
※(4) 
bar_switch_end_ratio 
double 
 
3 断面配筋の切替位置
(比)_終端側 
※(4) 
comment 
string 
 
コメント

## Page 44

3-16 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbBeam のkind_structure がRC またはSRC でかつisFoundation がtrue の部材とその断面に
適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。 
(3) figure_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定す
る。figure_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定
する。 
(4) bar_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定する。
bar_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定する。

## Page 45

3-17 
 
3.2.3.1.6. StbApply_RC_Column：RC 要領 柱 
・概要 
説明 ：RC 要領 柱 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isOnSite 
boolean 
 
現場打ちかどうか 
 
depth_cover_start_X 
double 
 
かぶり厚さ（X 始） 
 
depth_cover_end_X 
double 
 
かぶり厚さ（X 終） 
 
depth_cover_start_Y 
double 
 
かぶり厚さ（Y 始） 
 
depth_cover_end_Y 
double 
 
かぶり厚さ（Y 終） 
 
interval 
double 
 
２段筋のあき 
 
center_start_X 
double 
 
主筋重心位置（X 始） 
 
center_end_X 
double 
 
主筋重心位置（X 終） 
 
center_start_Y 
double 
 
主筋重心位置（Y 始） 
 
center_end_Y 
double 
 
主筋重心位置（Y 終） 
 
center_interval 
double 
 
２段筋重心間距離 
 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
allocation_rule_stirrup 
string 
 
帯筋：割り付けルール 
 
D_additional 
string 
 
補助筋径 
 
strength_additional 
string 
 
補助筋強度 
 
number_rule_additional 
string 
 
補助筋本数ルール 
※(2) 
anchorage_rule 
string 
 
定着ルール 
※(2) 
bar_switch_ratio 
double 
 
配筋の切替位置 
※(3) 
comment 
string 
 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 46

3-18 
 
・補足 
(1) StbColumn のkind_structure がRC またはSRC の部材とその断面に適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。 
(3) 鉄筋切替位置を始端オフセット位置から切り替え位置までの部材内法に対する比を指定する。

## Page 47

3-19 
 
3.2.3.1.7. StbApply_RC_Girder：RC 要領 大梁 
・概要 
説明 ：RC 要領 大梁 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isOnSite 
boolean 
 
現場打ちかどうか 
 
depth_cover_left 
double 
 
かぶり厚さ（左） 
 
depth_cover_right 
double 
 
かぶり厚さ（右） 
 
depth_cover_top 
double 
 
かぶり厚さ（上） 
 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
 
interval 
double 
 
２段筋のあき 
 
center_top 
double 
 
主筋重心位置（上） 
 
center_bottom 
double 
 
主筋重心位置（下） 
 
center_side 
double 
 
主筋重心位置（側） 
 
center_interval 
double 
 
２段筋重心間距離 
 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
allocation_rule_stirrup 
string 
 
帯筋：割り付けルール 
 
D_additional 
string 
 
補助筋径 
 
strength_additional 
string 
 
補助筋強度 
 
number_rule_additional 
string 
 
補助筋本数ルール 
※(2) 
anchorage_rule 
string 
 
定着ルール 
※(2) 
cut_off_rule 
string 
 
カットオフルール 
※(2) 
figure_switch_start_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_始端側 
※(3) 
figure_switch_end_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_終端側 
※(3) 
bar_switch_start_ratio 
double 
 
3 断面配筋の切替位置
(比)_始端側 
※(4) 
bar_switch_end_ratio 
double 
 
3 断面配筋の切替位置
(比)_終端側 
※(4) 
comment 
string 
 
コメント

## Page 48

3-20 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbGirder のkind_structure がRC またはSRC でかつisFoundation がfalse の部材とその断面
に適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。 
(3) figure_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定す
る。figure_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定
する。 
(4) bar_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定する。
bar_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定する。

## Page 49

3-21 
 
3.2.3.1.8. StbApply_RC_Beam：RC 要領 小梁 
・概要 
説明 ：RC 要領 小梁 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isOnSite 
boolean 
 
現場打ちかどうか 
 
depth_cover_left 
double 
 
かぶり厚さ（左） 
 
depth_cover_right 
double 
 
かぶり厚さ（右） 
 
depth_cover_top 
double 
 
かぶり厚さ（上） 
 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
 
interval 
double 
 
２段筋のあき 
 
center_top 
double 
 
主筋重心位置（上） 
 
center_bottom 
double 
 
主筋重心位置（下） 
 
center_side 
double 
 
主筋重心位置（側） 
 
center_interval 
double 
 
２段筋重心間距離 
 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
allocation_rule_stirrup 
string 
 
帯筋：割り付けルール 
 
D_additional 
string 
 
補助筋径 
 
strength_additional 
string 
 
補助筋強度 
 
number_rule_additional 
string 
 
補助筋本数ルール 
※(2) 
anchorage_rule 
string 
 
定着ルール 
※(2) 
cut_off_rule 
string 
 
カットオフルール 
※(2) 
figure_switch_start_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_始端側 
※(3) 
figure_switch_end_ratio 
double 
 
3 断面の場合の断面切替
位置(比)_終端側 
※(3) 
bar_switch_start_ratio 
double 
 
3 断面配筋の切替位置
(比)_始端側 
※(4) 
bar_switch_end_ratio 
double 
 
3 断面配筋の切替位置
(比)_終端側 
※(4) 
comment 
string 
 
コメント

## Page 50

3-22 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbBeam のkind_structure がRC またはSRC でかつisFoundation がfalse の部材とその断面に
適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。 
(3) figure_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定す
る。figure_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定
する。 
(4) bar_switch_start は始端オフセット位置から切り替え位置までの部材内法に対する比を指定する。
bar_switch_end は終端オフセット位置から切り替え位置までの部材内法に対する比を指定する。

## Page 51

3-23 
 
3.2.3.1.9. StbApply_RC_Wall：RC 要領 壁 
・概要 
説明 ：RC 要領 壁 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover 
double 
 
かぶり厚さ 
※(2) 
D_bar_spacing 
double 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
horizontal_bar_switch_start_ratio double 
 
横筋始端側の鉄筋切替
位置(比) 
※(3) 
horizontal_bar_switch_end_ratio 
double 
 
横筋終端側の鉄筋切替
位置(比) 
※(3) 
vertical_bar_switch_bottom_ratio 
double 
 
縦筋下端側の鉄筋切替
位置(比) 
※(3) 
vertical_bar_switch_top_ratio 
double 
 
縦筋上端側の鉄筋切替
位置(比) 
※(3) 
comment 
string 
 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbWall の部材と断面に適用する。 
(2) depth_covern_inside, depth_cover_outside に適用する。 
(3) 配筋が断面内で異なる場合の壁内法に対する比を指定する。

## Page 52

3-24 
 
3.2.3.1.10. StbApply_RC_Slab：RC 要領 スラブ 
・概要 
説明 ：RC 要領 スラブ 
親要素：StbApplyConditionList_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover 
double 
 
かぶり厚さ 
※(2) 
beam_reinforcement_rule 
string 
 
梁上の補強筋(合成床
板)ルール 
※(3) 
column_reinforcement_rule 
string 
 
柱周りの補強筋(合成
床板)ルール 
※(3) 
opening_reinforcement_rule string 
 
開口補強ルール 
※(3) 
corner_reinforcement_rule 
string 
 
出隅入隅補強ルール 
※(3) 
step_reinforcement_rule 
string 
 
段差補強ルール 
※(3) 
comment 
string 
 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbSlab の部材と断面に適用する。 
(2) depth_cover, depth_cover_top に適用する。 
(3) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 53

3-25 
 
3.2.3.2. StbApplyConditionList_S：S 要領 
・概要 
説明 ：S 要領 
親要素：StbApplyConditionsList 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbApply_S_General 
0 
1 
一般事項 
 
StbApply_S_Column 
0 
1 
柱 
 
StbApply_S_Girder 
0 
1 
大梁 
 
StbApply_S_Beam 
0 
1 
小梁 
 
 
 
・補足 
(1) 上記のうち１種類以上の子要素を持つものとし、全子要素の最小回数が0 であってはならない。

## Page 54

3-26 
 
3.2.3.2.1. StbApply_S_General：S 要領 一般事項 
・概要 
説明 ：S 要領 一般事項 
親要素：StbApplyConditionList_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
planting_rule 
string 
 
メッキルール 
※(2) 
covering_rule 
string 
 
耐火被覆ルール 
※(2) 
rustproof_rule 
string 
 
防錆処理ルール 
※(2) 
zone_rule 
string 
 
工区ルール 
※(2) 
comment 
string 
 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 鉄骨要領の一般事項を想定している。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 55

3-27 
 
3.2.3.2.2. StbApply_S_Column：S 要領 柱 
・概要 
説明 ：S 要領 柱 
親要素：StbApplyConditionList_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
comment 
string 
〇 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbColumn のkind_structure がS またはSRC の部材とその断面に適用する。

## Page 56

3-28 
 
3.2.3.2.3. StbApply_S_Girder：S 要領 大梁 
・概要 
説明 ：S 要領 大梁 
親要素：StbApplyConditionList_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
stud_rule 
string 
 
頭付きスタッドルール 
※(2) 
comment 
string 
 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbGirder のkind_structure がS またはSRC の部材とその断面に適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 57

3-29 
 
3.2.3.2.4. StbApply_S_Beam：S 要領 小梁 
・概要 
説明 ：S 要領 小梁 
親要素：StbApplyConditionList_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
stud_rule 
string 
 
頭付きスタッドルール 
※(2) 
comment 
string 
 
コメント 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbGirder のkind_structure がS またはSRC の部材とその断面に適用する。 
(2) 次のバージョン以降に実情を反映して詳細定義を行うこととする。

## Page 58

3-30 
 
3.2.4. システム標準板厚リスト：StbStandardPlateThicknessList 
・概要 
説明 ：板厚情報リスト 
親要素：StbCommon 
 
・属性 
無し 
 
・内容 
システム標準として使用する切板の板厚をブランクで区切って列挙する。 
内容 
型 
必須 
説明 
補足 
thickness 
[monolist] 
double 
○ 
使用する板厚 
※(1) 
 
・子要素 
無し 
 
・補足 
(1) 「1.8 集合型（monolist）」より、以下のように記載する（板厚は記入例）。 
<StbPlateThicknessList>1.2 1.4 1.6 1.8 2.0 2.3 2.5 2.8 3.2</StbPlateThicknessList> 
(2) 建物で使用する切板の板厚ではなく、ガセットプレートのサイズアップ等、システム標準として
計算等に使用する切板の板厚のサイズテーブルを記述する。

## Page 59

3-31 
 
3.2.5. コネクションスペック情報：StbConnectionSpecs 
・概要 
説明 ：コネクションスペック情報 
親要素：StbCommon 
 
・属性 
属性名 
型 
必須 
説明 
補足 
size_up_min_thickness 
double 
〇 
サイズアップの最小板厚差 
例: 3mm 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecBeamH 
0 
1 
S 梁仕口・H 形 
 
StbConnectionSpecColumnH 
0 
1 
S 柱仕口・H 形 
 
StbConnectionSpecColumnBox 
0 
1 
S 柱仕口・角形鋼管 
 
StbConnectionSpecColumnPipe 
0 
1 
S 柱仕口・円形鋼管 
 
 
・補足 
(1) 子要素の並びは、上表に示す順番としなければならない。

## Page 60

3-32 
 
3.2.5.1. Ｓ梁仕口・Ｈ形コネクション仕様情報：StbConnectionSpecBeamH 
概要 
説明 ：Ｓ梁の仕口・Ｈ形コネクションの仕様情報 
親要素：StbConnectionSpecs 
 
・属性 
属性名 
型 
必須 
説明 
補足 
gusset_size_up 
integer 
 
ガセットプレートのサイズアップ 
※(1) 
rib_size_up 
integer 
 
リブプレートのサイズアップ 
※(2) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecGussetPla
te 
0 
制限なし 
ガセットプレート仕様 
 
StbConnectionSpecRibPlate 
0 
制限なし 
リブプレート仕様 
 
 
・補足 
(1) 接続するH 形鋼ウェブ厚さに対するサイズアップについて記述する。 
(2) 接続するガセットプレート厚さに対するサイズアップについて記述する。

## Page 61

3-33 
 
3.2.5.1.1. ガセットプレート仕様：StbConnectionSpecGussetPlate 
概要 
説明 ：ガセットプレート仕様 
親要素：StbConnectionSpecBeamH、StbConnectionSpecColumnH 
 
・属性 
属性名 
型 
必須 
説明 
補足 
connection_type 
string 
○ 
仕口タイプ 
gap、web、splice 
※(1) 
flange_type 
string 
○ 
フランジ刃落としタイプ 
none、one_side、both_sides 
※(2) 
strength_plate 
string 
○ 
ガセットプレートの材種 
例：SS400 
strength_bolt 
string 
○ 
ボルト材種 
例：S10T 
name_bolt 
string 
○ 
ボルト径（呼び名） 
例：M20 
clearance 
double 
 
母材と接続部材の間隔 
※(3) 
flange_cut 
double 
 
フランジ刃落としの長さ 
フランジ刃落としタイ
プ
(flange_type)
が
none 以外の時、必須 
※(4) 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
m≧2 の時、必須 
※(4) 
pitch 
double 
 
部材長手方向のボルトピッチ (pL) 
n≧2 の時、必須 
※(4) 
e1 
double 
○ 
縁端距離1 (e1) 
※(4) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(4) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
※(4) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
※(4) 
h1 
double 
 
ガセットプレートのあき1 (h1) 
e3、e4 の数値を優先す
る ※(4) 
h2 
double 
 
ガセットプレートのあき2 (h2) 
e3、e4 の数値を優先す
る ※(4) 
ey 
double 
 
嵩上げ点の幅方向の距離 
※(6)

## Page 62

3-34 
 
ez 
double 
 
嵩上げ点の高さ方向の距離 
※(6) 
R 
double 
 
入隅半径 (R) 
※(6) 
id_weld 
integer 
 
溶接ID 
 
cutback 
integer 
 
溶接控え 
※(7) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecSplice 
0 
1 
スプライスプレート詳細 
 
 
・補足 
(1) 仕口タイプは、下図の通りとする。 
 
 
あき優先(gap) 
ウェブサイズ優先(web) 
2 面せん断(splice) 
 
(2) 刃落としのタイプは下図の通りとする。 
 
 
刃落としなし(none) 
片刃落とし(one_side) 
両刃落とし(both_sides) 
 
(3) 省略値は、10 ㎜とする。

## Page 63

3-35 
 
(4) 各属性は、下図の通りとする。 
 
 
(5) ガセットプレート、リブプレートの角処理のパラメータは定義しない。 
(6) 嵩上げの場合の各属性は下図の通りとする。 
 
右図のように嵩上げ点（左図の左赤丸）が存在しない場合、ey, ez を省略（ey=0, ez=0）してよい。 
 
 
h1 
e1 e2 
ｸﾘｱﾗﾝｽ 
pC 
pC 
pC 
e3 
e4 
h2 
m4 
m3 
m2 
m1 
m 
刃落としの長さ

## Page 64

3-36 
 
(7) 溶接控えの有無は下図の通りとし、その数値（mm）を記述する。 
 
控え無             控え有

## Page 65

3-37 
 
3.2.5.1.1.1. スプライスプレート仕様：StbConnectionSpecSplice 
概要 
説明 ：スプレイスプレート仕様 
親要素：StbConnectionSpecGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength_plate 
string 
○ 
添え板の材種 
例：SN490B 
plate_thickness 
double 
○ 
添え板の厚さ 
 
strength_filler 
string 
 
フィラープレートの材種 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) JASS6 の規定ではフィラープレートの材種は400N/mm2 級で良いことになっており、特別に指定す
る場合のみ記述する。

## Page 66

3-38 
 
3.2.5.1.2. リブプレート仕様：StbConnectionSpecRibPlate 
概要 
説明 ：リブプレート仕様 
親要素：StbConnectionSpecBeamH、StbConnectionSpecColumnH 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength_plate 
string 
○ 
リブプレートの材種 
例：SN490B 
id_weld 
integer 
 
溶接ID 
 
cutback 
integer 
 
溶接控え 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は3.2.22. ガセットプレート仕様補足を参照

## Page 67

3-39 
 
3.2.5.2. Ｓ柱仕口・Ｈ形コネクション仕様情報：StbConnectionSpecColumnH 
概要 
説明 ：Ｓ柱の仕口・Ｈ形コネクションの仕様情報 
親要素：StbConnectionSpecs 
 
・属性 
属性名 
型 
必須 
説明 
補足 
gusset_size_up 
integer 
 
ガセットプレートのサイズアップ 
※(1) 
stiffner_size_up_connected 
integer 
 
柱フランジ厚さに対する水平スチ
フナーのサイズアップ 
※(2) 
stiffner_size_up_connecting 
integer 
 
梁フランジ厚さに対する水平スチ
フナーのサイズアップ 
※(3) 
rib_size_up 
integer 
 
リブプレートのサイズアップ 
※(4) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecGussetPla
te 
0 
制限なし 
ガセットプレート仕様 
 
StbConnectionSpecStiffner 
0 
制限なし 
スチフナー仕様 
※(5) 
StbConnectionSpecRibPlate 
0 
1 
リブプレート仕様 
 
 
・補足 
(1) 接続する梁H 形鋼ウェブ厚さに対するサイズアップについて記述する。 
(2) 接続する柱H 形鋼フランジ厚さに対するサイズアップについて記述する。 
(3) 接続する梁H 形鋼フランジ厚さに対するサイズアップについて記述する。 
(4) 接続するガセットプレート厚さに対するサイズアップについて記述する。 
(5) 水平スチフナーの記述は、上から順に行う。

## Page 68

3-40 
 
3.2.5.2.1. スチフナープレート仕様：StbConnectionSpecStiffner 
概要 
説明 ：スチフナープレートの仕様 
親要素：StbConnectionSpecColumnH 
 
・属性 
 
属性名 
型 
必須 
説明 
補足 
strength_plate 
string 
〇 
スチフナーの材種 
例：SN490B 
id_weld 
integer 
 
溶接ID 
 
cutback 
integer 
 
溶接控え 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は3.2.22. ガセットプレート仕様補足を参照

## Page 69

3-41 
 
3.2.5.3. Ｓ柱仕口・角形鋼管コネクション仕様情報：StbConnectionSpecColumnBox 
概要 
説明 ：Ｓ柱の仕口・角形鋼管コネクションの仕様情報 
親要素：StbConnectionSpecs 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecDiaphrag
m 
0 
2 
ダイアフラム仕様 
※(1) 
StbConnectionSpecPanel 
0 
1 
パネル仕様 
 
 
・補足 
(1) 上下・中間で最大2 回繰り返す。 
(2) 中間の指定が無い場合、上下と中間の仕様は同じとする。

## Page 70

3-42 
 
3.2.5.3.1. ダイアフラム仕様：StbConnectionSpecDiaphragm 
概要 
説明 ：ダイアフラム仕様 
親要素：StbConnectionSpecColumnBox、StbConnectionSpecColumnPipe 
 
・属性 
属性名 
型 
必須 
説明 
補足 
location 
string 
○ 
ダイアフラムの位置 
以下のいずれかの値をとる。 
Outer（上下） 
Inner（中間） 
 
type 
string 
○ 
ダイアフラム形式 
以下のいずれかの値をとる。 
Through（通しダイア）、 
Internal（内ダイア）、 
External（外ダイア） 
 
isProduct 
boolean 
〇 
既製品か否か 
※(1) 
strength_border_thickness 
double 
 
材種の板厚のしきい値 
例：40mm 
strength_min 
string 
 
しきい値未満の板厚の場合のダ
イアフラムの材種 
SN490C 
strength_max 
string 
 
しきい値以上の板厚の場合のダ
イアフラムの材種 
TMCP325C 
e_border_dia_thickness 
double 
 
ダイアフラム厚に対するダイア
フラム出寸法のしきい値 
 
dia_size_up_panel 
integer 
 
パネル材の最大板厚に対するサ
イズアップ 
 
dia_thickness_up_panel 
double 
 
パネル材の最大板厚に加算する
板厚 
 
dia_size_up_flange 
integer 
 
梁フランジの最大板厚に対する
サイズアップ 
 
dia_thickness_up_flange 
double 
 
梁フランジの最大板厚に加算す
る板厚 
 
e_min_dia_thickness 
double 
 
ダイアフラム厚に対するダイア
フラム出寸法の最小値 
※(2) 
e_max_dia_thickness 
double 
 
ダイアフラム厚に対するダイア
フラム出寸法の最大値 
※(2)

## Page 71

3-43 
 
e_border_panel_thickness 
double 
 
パネル材の最大板厚に対するダ
イアフラム出寸法のしきい値 
 
e_min_panel_thickness 
double 
 
パネル材の最大板厚に対するダ
イアフラム出寸法の最小値 
※(2) 
e_max_panel_thickness 
double 
 
パネル材の最大板厚に対するダ
イアフラム出寸法の最大値 
※(2) 
max_flange_gap 
double 
 
梁フランジの目違いの最大値 
 
min_diaphragm_distance 
double 
 
ダイアフラム間の距離の最小値 
 
id_weld 
integer 
 
溶接ID 
※(3) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 既製品を使用するかどうかを記述するが、既製品のメーカー・品番等の記述は行わない。 
(2) 柱スキンプレートからダイアフラムの出寸法を記述する。 
(3) 通しダイア形式以外の場合、必須とする。

## Page 72

3-44 
 
3.2.5.3.2. パネルゾーン仕様：StbConnectionSpecPanel 
概要 
説明 ：パネル仕様 
親要素：StbConnectionSpecColumnBox、StbConnectionSpecColumnPipe 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isTaper 
boolean 
〇 
絞りの有無 
 
border_gap 
double 
 
絞りのしきい値 
※(1) 
panel_type 
string 
〇 
パネルのタイプ 
以下のいずれかの値をとる。 
Same（柱と同一） 
Product（既製品） 
Plate（プレート） 
※(2)、(3) 
dia_size_up_panel 
integer 
 
柱の最大板厚に対するサイズア
ップ 
 
dia_thickness_up_panel 
double 
 
柱の最大板厚に加算する板厚 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecPanelStren
gth 
0 
2 
パネル材種 
※(4) 
 
・補足 
(1) 上下の柱スキンプレート同士の離れがしきい値以上の場合絞る。 
(2) パネルの構成材を指定する。下の柱と同じサイズ（Same）、既製品を使う場合（Product）、プレー
トで構成する場合（Plate）の別を記述する。 
(3) 既製品を使う場合でも、既製品のメーカー・品番等の記述は行わない。 
(4) 既製品を使う場合材種は指定しない。

## Page 73

3-45 
 
3.2.5.3.2.1. パネルゾーン材種：StbConnectionSpecPanelStrength 
概要 
説明 ：パネル材種仕様 
親要素：StbConnectionSpecPanel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength 
string 
〇 
パネルの材種 
以下のいずれかの値をとる。 
Same（同一材種） 
Class_C（同一強度でC 種） 
 
strength_internal 
string 
 
内ダイアが溶接される場合のパ
ネルの材種 
以下のいずれかの値をとる。 
Same（同一材種） 
Class_C（同一強度でC 種） 
省略値はSame とする 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 74

3-46 
 
3.2.5.4. Ｓ柱仕口・円形鋼管コネクション仕様情報：StbConnectionSpecColumnPipe 
概要 
説明 ：Ｓ柱の仕口・円形鋼管コネクションの仕様情報 
親要素：StbConnectionSpecs 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionSpecDiaphragm 
0 
2 
ダイアフラム仕様 
※(1) 
StbConnectionSpecPanel 
0 
1 
パネル仕様 
 
 
・補足 
(1) 上下・中間で最大2 回繰り返す。 
(2) 中間の指定が無い場合、上下と中間の仕様は同じとする。

## Page 75

3-47 
 
3.2.6. 溶接情報：StbWeldCommon 
・概要 
説明 ：溶接における共通情報 
親要素：StbCommon 
 
・属性 
属性名 
型 
必須 
説明 
補足 
kind_end_tab 
string 
 
エンドタブの種類 
以下のいずれかの値をとる。 
Steel（鋼製エンドタブ）、Flux（固形エンド
タブ） 
 
strength_backup 
string 
 
裏当て金の材質 
 
shape_backup 
string 
 
裏当て金の材種 
以下の値をとる。 
PL（プレート）、FB（フラットバー） 
 
size_backup 
string 
 
裏当て金のサイズ 
※(1) 
 
・内容、子要素 
無し 
 
・補足 
(1) 裏当て金のサイズは「板厚x 幅」の記述とする。

## Page 76

3-48 
 
3.2.7. 追加情報：StbAdditionalInformation 
・概要 
説明 ：追加情報 
親要素：StbCommon 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pile_allowable_eccentricity_X double 
○ 
杭許容偏心量(X 方向) 
 
pile_allowable_eccentricity_Y double 
○ 
杭許容偏心量(Y 方向) 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 77

4-1 
 
4. 要素リファレンス 位置情報 
4.1. 位置・断面情報：StbModel  
・概要 
説明 ：位置・断面情報（節点・部材・階・軸） 
親要素：ST_BRIDGE 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodes 
0 
1 
節点 
※(1) 
StbAxes 
0 
1 
軸 
 
StbStories 
0 
1 
階 
 
StbMembers 
0 
1 
部材 
部材情報 
StbSections 
0 
1 
断面 
断面情報 
StbJoints 
0 
1 
継手 
継手情報 
StbConnections 
0 
1 
コネクション 
コネクション情報 
StbWeld 
0 
1 
溶接 
溶接情報 
 
・補足 
(1) 軸や階のみ受け渡す場合は、節点の未定義を許容する。

## Page 78

4-2 
 
4.2. 節点（複数）：StbNodes 
・概要 
説明 ：節点（複数） 
親要素：StbModel 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNode 
1 
制限なし 
節点

## Page 79

4-3 
 
4.2.1. 節点：StbNode 
・概要 
説明 ：節点 
親要素：StbNodes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
X 
double 
○ 
全体座標系 X
－
 
※(1) 
Y 
double 
○ 
全体座標系 Y
－
 
※(1) 
Z 
double 
○ 
全体座標系 Z
－
 
※(1) 
kind 
string 
○ 
以下のいずれかの値をとる 
ON_GIRDER：大梁上 
ON_BEAM：小梁上 
ON_COLUMN：柱上 
ON_POST：間柱上 
ON_GRID：グリッド上 
ON_CANTI：片持ち大梁先端 
ON_SLAB：スラブ上 
OTHER：その他 
※(2)～※(7) 
id_member 
integer 
 
リンクする部材のID 
※(2)～※(7) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 全体座標原点からの相対座標値とする。 
(2) 大梁上、小梁上の節点（ON_GIRDER／ON_BEAM）は、梁の中間に他の梁などが取付く場合の、
取付く側の端点を指定する。このとき、「リンクする部材のID」には取付かれる側の梁のID を
記述する。スラブ上（ON_SLAB）、柱上（ON_COLUMN／ON_POST）の節点も同様とする。 
(3) 柱上、間柱上の節点（ON_COLUMN／ON_POST）は、階段の踊り場受け梁など階の途中に水平
材が取りつく場合に用いる。

## Page 80

4-4 
 
(4) グリッド上の節点（ON_GRID）は、複数の軸および階との交点を示す概念であり、一貫構造計算
プログラムなどにおいて構造耐力上主要な柱や大梁などが接続されることを想定している。ただ
し、その座標は必ずしも交点と一致する必要はない。例えば、下図で節点X4-Y3 のY 座標はY3
軸の代表距離と一致しなくてもよい。「リンクする部材のID」は、指定しない。 
(5) 下図のX2 軸とY2 軸の交点のように、梁の中間に他の梁が取付く節点がグリッド上にある場合
は、ON_GRID とする。 
(6) 片持ち先端の節点（ON_CANTI）は、片持ち大梁の先端節点に用いる。このとき、「リンクする
部材のID」には、大梁自身のID を記述する。 
(7) その他の節点（OTHER）は、腰壁の上側節点・垂れ壁の下側節点やY 型ブレースの中間節点、片
持ち小梁・片持ちスラブの先端節点などに用いる。 
 
      
 
 
 
原点 
X
－
 
Y
－
 
Z
－

## Page 81

4-5 
 
 
    片持ち梁と片持ちスラブの節点が重なる場合は、下図の配置とする。 
 
 
・例 
 
 
 
 
 
 
 
 
 
ON_CANTI
OTHER
：ON_GIRDER
節点
：ON_GRID
：ON_CANTI
：OTHER
ON_CANTI
ON_GIRDER
ON_CANTI
ON_GIRDER
OTHER
<StbNodes> 
<StbNode id="15" X="0.000000" Y="0.000000" Z="0.000000" kind="ON_GRID"/> 
・・・ 
<StbNode id="276" X="9000.000000" Y="0.000000" Z="0.000000" kind="ON_GIRDER" 
id_member="59"/> 
・・・ 
</StbNodes>

## Page 82

4-6 
 
4.2.2. 節点ID リスト：StbNodeIdList 
・概要 
説明 ：節点ID リスト 
親要素： StbParallelAxis、StbArcAxis、StbRadialAxis、StbStory 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeId 
1 
制限なし 
節点ID 
 
 
・補足 
子要素StbNodeId の属性 id は、StbNodeIdList 内で一意とする。 
・例 
 
 
 
 
 
 
 
 
 
 
 
 
 
<StbAxes> 
<StbParallelAxes group_name ="X_Axes" X="0.00" Y="0.00" angle="270.00"> 
<StbParallelAxis id ="3" name="1" distance="0.00"> 
<StbNodeIdList> 
<StbNodeId id="34"/> 
・・・ 
</StbNodeIdList>  
</StbParallelAxis> 
  ・・・ 
</StbParallelAxes> 
</StbAxes>

## Page 83

4-7 
 
4.2.3. 節点ID：StbNodeId 
・概要 
説明 ：節点ID 
親要素：StbNodeIdList 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
StbNode のID 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 84

4-8 
 
4.2.4. 順序のある節点ID：StbNodeIdOrder 
・概要 
説明 ：順序のある節点ID  
親要素：StbColumnViaNode, StbPostViaNode, StbGirderViaNode, StbBeamViaNode,  
StbSlab, StbWall, StbFrameDampingDevice 
 
・属性 
無し 
 
・内容 
節点ID をブランクで区切って列挙する。 
内容 
型 
必須 
説明 
補足 
id 
[monolist] 
integer 
○ 
StbNode のID 
※(1) 
 
・子要素 
無し 
 
・補足 
(1) 「1.8 集合型（monolist）」より、以下のように記載する（ID は記入例）。 
<StbNodeIdOrder>101 102 103 104</StbNodeIdOrder>

## Page 85

4-9 
 
4.3. 軸（複数）：StbAxes 
・概要 
説明 ：軸情報 
親要素：StbModel 
 
・属性 
無し 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbParallelAxes 
0 
制限なし 
平行軸(複数) 
※(2) 
StbArcAxes 
0 
制限なし 
円弧軸(複数) 
※(2) 
StbRadialAxes 
0 
制限なし 
放射軸(複数) 
※(2) 
StbDrawingAxes 
0 
1 
作図用軸(複数) 
※(1) 
 
・補足 
(1) 作図用軸（StbDrawingAxis）は、建物形状の基本情報である平行軸（StbParallelAxes）、円弧軸
（StbArcAxes）、放射軸(StbRadialAxes)とは別に、作図用の軸を定義する場合に使用する。 
(2) 上記のうち１種類以上の子要素を持つものとし、全子要素の最小回数が0 であってはならない。

## Page 86

4-10 
 
4.3.1. 平行軸(複数)：StbParallelAxes 
・概要 
説明 ：平行軸(複数) 
親要素：StbAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
group_name 
string 
○ 
平行軸グループの名称 
 
X 
double 
○ 
基準座標 X
－
 
 
Y 
double 
○ 
基準座標 Y
－
 
 
angle 
double 
○ 
角度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbParallelAxis 
1 
制限なし 
平行軸 
※(1) 
 
・補足 
(1) 軸を面として定義している。（階の情報をもたない。） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
Ｘ１
Ｘ２
Ｘ３
Ｘ４
Ｙ１
Ｙ２
Ｙ３
Ｚ１
Ｚ２
Ｚ３
Ｚ４
Ｚ５
Ｚ６
Ｚ７
ＧＬ
Ｘ１
Ｘ２
Ｘ３
Ｘ４
Ｙ１
Ｙ２
Ｙ３
Ｚ１
Ｚ２
Ｚ３
Ｚ４
Ｚ５
Ｚ６
Ｚ７
ＧＬ
ＧＬ
X
－
 
Y
－
 
Z
－
 
例：X 軸グループ(X1～X4)の
基準座標点を原点とし
た場合

## Page 87

4-11 
 
4.3.1.1. 平行軸：StbParallelAxis 
・概要 
説明 ：平行軸 
親要素：StbParallelAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
名称 
 
distance 
double 
○ 
基準座標点からの距離 
※(2) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdList 
0 
1 
節点ID リスト 
※(1) 
 
・補足 
(1) 属する節点（StbNodeIdList）はグリッド上の節点（ON_GRID）のみ列挙する。 
(2) 平行軸は、親要素StbParallelAxes の属性 angle で示される角度に平行な（複数の）面で定義さ
れる。基準座標点は、親要素StbParallelAxes の属性 X,Y で示される、X
―
Y
―
 平面上の点である。
各平行軸の位置は、この「基準座標点からの距離」で規定され、基準座標点を中心に全体座標系を 
Z
―
 軸について属性 angle で示される角度に回転した後の座標系において、Y 軸への距離を指定す
る。Y 軸負方向の位置に面がある場合は、負の値を指定する。 
 
 
 
 
Y 
X 
回転後の座標系 
 
Y1 
＋distance 
angle=0 
 
Y2 
Y2 
  
Y3 
Y 
X 
回転後の座標系 
 
Y1 
＋distance 
angle=180 
 
 Y2 
  
Y3 
Y 
X 
回転後の座標系 
 
Y1 
＋distance 
angle=270 
 
 Y2 
  
Y3 
 
Y1 
＋distance 
angle=90 
 
 Y2 
  
Y3 
X 
Y 
回転後の座標系 
原点 
X
－
 
Y
－
 
Z
－
 
全体座標系

## Page 88

4-12 
 
・例 右図の場合 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
<StbParallelAxes group_name ="X" X="0.00" Y="0.00" angle="270.00"> 
<StbParallelAxis id ="3" name="X1" distance="0.00"> 
<StbNodeIdList> 
・・・ 
</StbNodeIdList>  
</StbParallelAxis> 
<StbParallelAxis id ="4" name="X2" distance="6000.00"> 
<StbNodeIdList> 
・・・ 
</StbNodeIdList>  
</StbParallelAxis> 
・・・ 
</StbParallelAxes> 
<StbParallelAxes group_name ="Y" X="0.00" Y="0.00" angle="0.00"> 
<StbParallelAxis id ="7" name="Y1" distance="0.00"> 
<StbNodeIdList> 
・・・ 
</StbNodeIdList>  
</StbParallelAxis> 
<StbParallelAxis id ="8" name="Y2" distance="9000.00"> 
<StbNodeIdList> 
・・・ 
</StbNodeIdList>  
</StbParallelAxis> 
  ・・・ 
</StbParallelAxes>

## Page 89

4-13 
 
4.3.2. 円弧軸(複数)：StbArcAxes 
・概要 
説明 ：円弧軸(複数) 
親要素：StbAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
group_name 
string 
○ 
円弧軸グループの名称 
 
X 
double 
○ 
中心座標 X
－
 
 
Y 
double 
○ 
中心座標 Y
－
 
 
start_angle 
double 
 
開始角度 
※(2) 
end_angle 
double 
 
終了角度 
※(2) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbArcAxis 
1 
制限なし 
円弧軸 
 
 
・補足 
(1) 軸を面として定義している。（階の情報をもたない。） 
(2) 円弧形状の場合は必須とし、円状の場合は開始角度・終了角度を指定しない。 
 
 
 
X
－
 
Y
－
 
Z
－
 
開始角度 
終了角度

## Page 90

4-14 
 
4.3.2.1. 円弧軸：StbArcAxis 
・概要 
説明 ：円弧軸 
親要素：StbArcAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
名称 
 
radius 
double 
○ 
中心座標からの半径距離 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdList 
0 
1 
節点ID リスト 
 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
<StbArcAxes group_name ="Arc_Axes" X="0.00" Y="0.00" start_angle="0.00" 
end_angle="270.00"> 
<StbArcAxis id ="3" name="A" radius="3000.00"> 
<StbNodeIdList> 
<StbNodeId id="34"/> 
・・・ 
</StbNodeIdList>  
</StbArcAxis> 
・・・ 
</StbArcAxes>

## Page 91

4-15 
 
4.3.3. 放射軸(複数)：StbRadialAxes 
・概要 
説明 ：放射軸(複数) 
親要素：StbAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
group_name 
string 
○ 
放射軸グループの名称 
 
X 
double 
○ 
中心座標 X
－
 
 
Y 
double 
○ 
中心座標 Y
－
 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbRadialAxis 
1 
制限なし 
放射軸 
 
 
・補足 
(1) 軸を面として定義している。（階の情報をもたない。） 
 
 
 
X
－
 
Y
－
 
Z
－

## Page 92

4-16 
 
4.3.3.1. 放射軸：StbRadialAxis 
・概要 
説明 ：放射軸 
親要素：StbRadialAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
名称 
 
angle 
double 
○ 
中心座標からの角度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdList 
0 
1 
節点ID リスト 
 
 
・例 
 
 
 
 
 
 
 
 
 
 
<StbRadialAxes group_name ="Rad_Axes" X="0.00" Y="0.00"> 
<StbRadialAxis id ="3" name="R1" angle="0.00"> 
<StbNodeIdList> 
<StbNodeId id="34"/> 
・・・ 
</StbNodeIdList>  
</StbRadialAxis> 
・・・ 
</StbRadialAxes>

## Page 93

4-17 
 
4.3.4. 作図用軸(複数)：StbDrawingAxes 
・概要 
説明 ：作図用軸(複数) 
親要素：StbAxes 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbDrawingLineAxis 
0 
制限なし 
作図用直線軸 
 
StbDrawingArcAxis 
0 
制限なし 
作図用円弧軸 
 
 
・補足 
上記のうち１種類以上の子要素を持つものとし、全子要素の最小回数が0 であってはならない。

## Page 94

4-18 
 
4.3.4.1. 作図用直線軸：StbDrawingLineAxis 
・概要 
説明 ：作図用直線軸 
親要素：StbDrawingAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
group_name 
string 
 
軸グループ名称 
作図用軸にまとまりを作り
たいときに使用する 
name 
string 
○ 
名称 
 
start_X 
double 
○ 
始点座標 X
－
 
作図用 
start_Y 
double 
○ 
始点座標 Y
－
 
作図用 
end_X 
double 
○ 
終点座標 X
－
 
作図用 
end_Y 
double 
○ 
終点座標 Y
－
 
作図用 
 
・内容 
無し 
 
・子要素 
無し 
 
・例 
 
 
 
 
 
 
 
 
 
<StbDrawingAxes> 
<StbDrawingLineAxis name="1" start_X="0.00" start_Y="0.00" end_X="0.00" 
end_Y="16000.00"/> 
・・・ 
<StbDrawingLineAxis name="A" start_X="0.00" start_Y="0.00" end_X="16000.00" 
end_Y="0.00"/> 
</StbDrawingAxes>

## Page 95

4-19 
 
4.3.4.2. 作図用円弧軸：StbDrawingArcAxis 
・概要 
説明 ：作図用円弧軸 
親要素：StbDrawingAxes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
group_name 
string 
 
軸グループ名称 
作図用軸にまとまりを作り
たいときに使用する 
name 
string 
○ 
名称 
 
X 
double 
○ 
中心座標 X
－
 
作図用 
Y 
double 
○ 
中心座標 Y
－
 
作図用 
radius 
double 
○ 
半径 
作図用 
start_angle 
double 
 
開始角度 
作図用※(1) 
end_angle 
double 
 
終了角度 
作図用※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 円弧形状の場合は必須とし、円状の場合は開始角度・終了角度を指定しない。 
 
・例 
 
 
 
 
 
 
 
 
<StbDrawingAxes> 
<StbDrawingArcAxis name="A" X="0.00" Y="0.00" radius="3000.00" start_angle ="0.00" 
end_angle ="270.00"/> 
・・・ 
</StbDrawingAxes>

## Page 96

4-20 
 
4.4. 階（複数）：StbStories 
・概要 
説明 ：階情報（複数） 
親要素：StbModel 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbStory 
1 
制限なし 
各階の情報

## Page 97

4-21 
 
4.4.1. 階：StbStory 
・概要 
説明 ：階情報 
親要素：StbStories 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
階名称 
※(1) 
level_name 
string 
 
フロアレベル名称 
※(1) 
height 
double 
○ 
代表高さ 
※(2) 
kind 
string 
○ 
階属性 
以下のいずれかの値を取る 
GENERAL（一般階） 
BASEMENT（地下階） 
ROOF（屋上階） 
PENTHOUSE（塔屋階） 
ISOLATION（免震階） 
DEPENDENCE（従属階） 
 
id_dependence 
integer 
 
従属する階のID 
※(3) 
strength_concrete string 
 
コンクリート強度 
※(4), ※(5) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdList 
0 
1 
節点ID リスト 
※(6) 
 
・補足 
(1) name はStbSection のfloor と対応する名称を想定している (例、3 階であれば3)。level_name
はフロアレベル名称(例：3FL)を記載する。 
(2) 各階の代表高さはＧＬから意匠上のＦＬまでの高さとする。代表高さが同じStbStory を二重に定
義してはならない。 
(3) 属性kind が DEPENDENCE の場合にのみ必須。

## Page 98

4-22 
 
(4) 階毎でコンクリート強度を設定したい場合に記述（例：FC24） 
(5) コンクリート強度を参照する部材の範囲は「見上げ」とする。例えば、StbStory がZ1～Z3 の場
合、以下とする。 
  Z1 のコンクリート強度は、Z1 の梁・スラブ・フーチング 
  Z2 のコンクリート強度は、Z2 の梁・スラブとZ1～Z2 間の柱・壁 
  Z3 のコンクリート強度は、Z3 の梁・スラブとZ2～Z3 間の柱・壁、およびZ3 のパラペット 
 
コンクリート強度を参照する階は、 
 
柱、間柱の場合   ・・・終端節点 
大梁、小梁の場合  ・・・始端節点 
床の場合      ・・・第１節点 
壁の場合      ・・・最終節点 
フーチングの場合  ・・・節点 
布基礎の場合    ・・・始端節点 
基礎柱の場合    ・・・節点 
パラペットの場合  ・・・始端節点 
 
が、それぞれ属する階とする。 
 
(6) 属する節点（StbNodeIdList）はグリッド上の節点以外も列挙する。「4.2.2 節点ID リスト：
StbNodeIdList」と同様。 
 
・例  
 
 
 
 
 
 
 
 
 
<StbStories> 
<StbStory id="12" name="1" height="0.00" kind="GENERAL"> 
<StbNodeIdList> 
<StbNodeId id="289"/> 
・・・ 
</StbNodeIdList>  
</StbStory> 
 ・・・ 
</StbStories>

## Page 99

5-1 
 
5. 要素リファレンス 部材情報 
5.1. 部材情報：StbMembers 
・概要 
説明 ：柱・梁・スラブ・壁などの部材情報 
親要素：StbModel 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbColumns 
0 
1 
柱（複数） 
 
StbPosts 
0 
1 
間柱（複数） 
 
StbGirders 
0 
1 
大梁（複数） 
 
StbBeams 
0 
1 
小梁（複数） 
 
StbBraces 
0 
1 
ブレース（複数） 
 
StbSlabs 
0 
1 
スラブ（複数） 
 
StbWalls 
0 
1 
壁（複数） 
 
StbIsolatingDevices 
0 
1 
免震装置（複数） 
 
StbDampingDevices 
0 
1 
制振装置（複数） 
 
StbFrameDampingDevices 
0 
1 
制振装置（フレーム） （複
数） 
 
StbFootings 
0 
1 
フーチング（複数） 
 
StbStripFootings 
0 
1 
布基礎（複数） 
 
StbPiles 
0 
1 
杭基礎（複数） 
 
StbFoundationColumns 
0 
1 
基礎柱（複数） 
 
StbParapets 
0 
1 
パラペット（複数） 
 
StbOpenArrangements 
0 
1 
開口（複数） 
 
StbPenetrationArrangements 
0 
1 
貫通孔（複数） 
 
StbJointArrangements 
0 
1 
継手(複数) 
 
StbPanelZoneArrangements 
0 
1 
柱梁接合部(複数) 
 
StbConnectionArrangements 
0 
1 
コネクション(複数)

## Page 100

5-2 
 
・補足 
上記のうち１種類以上の子要素を持つものとし、全子要素の回数が0 であってはならない。

## Page 101

5-3 
 
5.2. 柱（複数）：StbColumns 
・概要 
説明 ：柱情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbColumn 
1 
制限なし 
柱情報

## Page 102

5-4 
 
5.2.1. 柱：StbColumn 
・概要 
説明 ：柱情報 
親要素：StbColumns 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_bottom 
integer 
○ 
始端節点ID 
※(1) 
id_node_top 
integer 
○ 
終端節点ID 
※(1) 
rotate 
double 
 
回転角 
※(3) 
id_section 
integer 
○ 
断面ID 
※(6) 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値を
とる。 
RC、S、SRC、CFT、
UNDEFINED 
※(6) 
strength_concrete 
string 
 
コンクリート強度 
※(7) 
offset_bottom_X 
double 
 
始端側オフセット（X
－
） 
※(2) 
offset_bottom_Y 
double 
 
始端側オフセット（Y
－
） 
※(2) 
offset_bottom_Z 
double 
 
始端側オフセット（Z
－
） 
※(2) 
offset_top_X 
double 
 
終端側オフセット（X
－
） 
※(2) 
offset_top_Y 
double 
 
終端側オフセット（Y
－
） 
※(2) 
offset_top_Z 
double 
 
終端側オフセット（Z
－
） 
※(2) 
thickness_add_start_X 
double 
 
ふかし厚さ（X 始） 
円形の場合はＸ始のみ 
※(4) 
thickness_add_end_X 
double 
 
ふかし厚さ（X 終） 
※(4) 
thickness_add_start_Y 
double 
 
ふかし厚さ（Y 始） 
※(4) 
thickness_add_end_Y 
double 
 
ふかし厚さ（Y 終） 
※(4) 
seam 
string 
 
シーム方向 
以下のいずれかの値を
とる 
X始（X_P）、X終（X_N） 
ボックス柱のシームの
方向 
※(11)

## Page 103

5-5 
 
Y始（Y_P）、X終（Y_N） 
X、Y 
bar_switch_height 
double 
 
配筋切り替え高さ 
※(8) 
steel_cutback_bottom 
double 
 
鉄骨カットバック（始
端側） 
※(9) 
steel_cutback_top 
double 
 
鉄骨カットバック（終
端側） 
※(9) 
steel_switch_height_bottom 
double 
 
鉄骨断面切り替え高さ
(始端側) 
※(10) 
steel_switch_height_top 
double 
 
鉄骨断面切り替え高さ
(終端側) 
※(10) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbColumnViaNode 
0 
1 
柱中間節点 
※(5) 
 
 
・補足 
(1) 軸線が曲率を有する柱は対象外。 
(2) 柱の基準点は図心位置とする。オフセットが省略された場合、始終端の基準点は節点とする。 
S 柱で、柱脚でベースプレートを有する場合は、モルタル下端の位置を指定するものとする。 
(3) 回転はX 方向を0 度とし、始端から終端への進行方向時計回りを正とし、断面を回転した後にオ
フセットを考慮する。省略した時は0 度（回転なし）とみなす。 
(4) 省略された場合、当該属性がないものとする。 
(5) 該当節点がある場合は、節点ID を、始端に近い順に記述する。当要素は柱の折れ曲がりを表現す
るための要素であり、折れ曲がりがなく他の部材の端点となるだけの場合は記載する必要はない。 
(6) 構造種別と、断面ID が参照する要素名との対応は、それぞれRC－StbSecColumn_RC、S－
StbSecColumn_S、SRC－StbSecColumn_SRC、CFT－StbSecColumn_CFT、UNDEFINED－
StbSecUndefined とする。 
(7) 省略された場合の扱いについては、StbCommon の補足説明を参照のこと。 
(8) StbApply_RC_Column で指定したbar_switch_ratio から変更する場合、始端基準点からの距離
を指定する。

## Page 104

5-6 
 
(9) SRC の鉄骨で最上階柱の柱頭や最下階の柱脚など、鉄骨位置が基準点と異なる場合に基準点から
反対側の基準点方向への距離を指定する。柱脚でベースプレートを有する場合は、モルタル下端
までの距離を指定するものとする。 
(10) 鉄骨断面が柱頭・柱脚が別形状の場合はsteel_switch_height_bottom に記入し、柱頭・中央・柱
脚が別形状の場合はsteel_switch_height_bottom、steel_switch_height_top 両方に記入する。い
ずれも基準点からの距離とする。 
(11) シームの方向については下図による 
 
 
 
 
 
 
 
 
 
 
 
・ 
 
 
 
 
 
 
X_N 
（X 終） 
X_P 
（X 始） 
Y_P 
（Y 始） 
Y_N 
（Y 終） 
X 
Y 
X
Y 
ボックス柱のシームの方向 
（終端側から始端側を見た図） 
Y 
X 
終端基準点 
終端節点 
始端基準点 
始端節点 
始端オフセット(X
－
,Y
－
,Z
－
) 
終端オフセット(X
－
,Y
－
,Z
－
) 
X 
Z 
Y 
ふかし厚さ(Y 終) 
ふかし厚さ(X 終) 
ふかし厚さ(Y 始) 
ふかし厚さ(X 始) 
オフセット(X
－
,Y
－
,Z
－
) 
節点 
図心 
回転＋ 
鉄骨カットバ
ック(終端側) 
鉄骨カットバ
ック(始端側)

## Page 105

5-7 
 
例 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
<StbColumns> 
<StbColumn id="56" name="2CA3" id_node_bottom="30" id_node_top="39" rotate="0.00" 
id_section="1" kind_structure="RC" offset_bottom_X="0.00" offset_bottom_Y="0.00" 
offset_bottom_Z="0.00" offset_top_X="0.00" offset_top_Y="0.00" offset_top_Z="0.00"/> 
・・・ 
<StbColumn id="131" name="1CA1" id_node_bottom="15" id_node_top="24" rotate="0.00" 
id_section="3" kind_structure="SRC" offset_bottom_X="0.00" offset_bottom_Y="0.00" 
offset_bottom_Z="-650.00" offset_top_X="0.00" offset_top_Y="0.00" offset_top_Z="0.00"/> 
・・・ 
<StbColumn id="141" name="2CA2" id_node_bottom="27" id_node_top="36" rotate="0.00" 
id_section="12" kind_structure="S" offset_bottom_X="0.00" offset_bottom_Y="0.00" 
offset_bottom_Z="0.00" offset_top_X="0.00" offset_top_Y="0.00" offset_top_Z="0.00"/> 
・・・ 
<StbColumn id="157" name="2CA1" id_node_bottom="24" id_node_top="33" rotate="0.00" 
id_section="16" kind_structure="CFT" offset_bottom_X="0.00" offset_bottom_Y="0.00" 
offset_bottom_Z="0.00" offset_top_X="0.00" offset_top_Y="0.00" offset_top_Z="-150.00"/> 
・・・ 
</StbColumns>

## Page 106

5-8 
 
5.2.1.1. 柱中間節点：StbColumnViaNode 
・概要 
説明 ：柱の中間節点 
親要素：StbColumn 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
中間節点ID 
※(1) 
StbMemberOffsetList 
0 
制限なし 
中間節点オフセットリスト 
※(2) 
 
・補足 
中間節点とは、柱の始端節点から終端節点の間に有る、始終端以外の節点である。柱部材が折れ曲が
りのある直線の組合せの場合に、その形状を示すことができる。また、中間節点が、他の大梁およびブ
レースの端点となっている場合は、構造計算上、端点を有する部材と接続関係があることを示す。 
 
(1) 該当節点ID を、始端に近い順に記述する。始端自身および終端自身は記述しない。 
(2) 中間節点が、StbColumn の始終点と同様のオフセットを有する場合に記述する。省略された場合、
各オフセットの値は0 とする。

## Page 107

5-9 
 
5.2.1.2. 中間節点オフセットリスト：StbMemberOffsetList 
・概要 
説明 ：中間節点のオフセット 
親要素：StbColumnViaNode, StbGirderViaNode 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_node 
integer 
○ 
<StbNodeIdOrder>の中間節点ID 
 
offset_X 
double 
○ 
X
－
方向のオフセット 
 
offset_Y 
double 
○ 
Y
－
方向のオフセット 
 
offset_Z 
double 
○ 
Z
－
方向のオフセット 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 108

5-10 
 
5.3. 間柱（複数）：StbPosts 
・概要 
説明 ：間柱情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbPost 
1 
制限なし 
間柱情報

## Page 109

5-11 
 
5.3.1. 間柱：StbPost 
・概要 
説明 ：間柱情報 
親要素：StbPosts 
 
・属性 
「5.2.1 柱：StbColumn」と同様 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbPostViaNode 
0 
1 
柱中間節点 
※(9) 
 
・補足 
無し

## Page 110

5-12 
 
5.3.1.1. 間柱中間節点：StbPostViaNode 
・概要 
説明 ：間柱の中間節点 
親要素：StbPost 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
中間節点ID 
 
StbMemberOffsetList 
0 
制限なし 
中間節点オフセットリスト 
 
 
・補足 
「5.2.1 柱中間節点：StbColumnViaNode」と同様

## Page 111

5-13 
 
5.4. 大梁（複数）：StbGirders 
・概要 
説明 ：大梁情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbGirder 
1 
制限なし 
大梁情報

## Page 112

5-14 
 
5.4.1. 大梁：StbGirder 
・概要 
説明 ：大梁情報 
親要素：StbGirders 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_start 
integer 
○ 
始端節点ID 
※(1), ※(2) 
id_node_end 
integer 
○ 
終端節点ID 
※(1), ※(2) 
rotate 
double 
 
回転角 
※(3), ※(4) 
id_section 
integer 
○ 
断面ID 
※(8) 
section_io_start 
string 
 
断面の外端・内端指定 
以下のいずれかの値をとる。 
OUT、IN 
※(5) 
section_io_end 
string 
 
断面の外端・内端指定 
以下のいずれかの値をとる。 
OUT、IN 
※(5) 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値をとる。 
RC、S、SRC、UNDEFINED 
※(8) 
isFoundation 
boolean 
○ 
基礎か否か 
 
strength_concrete 
string 
 
コンクリート強度 
※(9) 
offset_start_X 
double 
 
始端側オフセット（X
－
） 
※(3) 
offset_start_Y 
double 
 
始端側オフセット（Y
－
） 
※(3) 
offset_start_Z 
double 
 
始端側オフセット（Z
－
） 
※(3) 
offset_end_X 
double 
 
終端側オフセット（X
－
） 
※(3) 
offset_end_Y 
double 
 
終端側オフセット（Y
－
） 
※(3) 
offset_end_Z 
double 
 
終端側オフセット（Z
－
） 
※(3) 
thickness_add_top 
double 
 
ふかし厚さ（上） 
※(6) 
thickness_add_bottom 
double 
 
ふかし厚さ（下） 
※(6) 
thickness_add_right 
double 
 
ふかし厚さ（右） 
※(6) 
thickness_add_left 
double 
 
ふかし厚さ（左） 
※(6)

## Page 113

5-15 
 
isOnSite 
boolean 
 
現場打ちかどうか 
 
composite_beam 
string 
 
合成梁 
以下のいずれかの値をとる。 
NONE：合成梁として期待しない 
INCOMPLETE：不完全合成梁 
COMPLETE：完全合成梁 
※(10) 
steel_cutback_start 
double 
 
鉄骨カットバック（始端） 
※(11) 
steel_cutback_end 
double 
 
鉄骨カットバック（終端） 
※(11) 
cover_plate 
string 
 
カバープレート 
※(12) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbGirderViaNode 
0 
1 
大梁中間節点 
※(7) 
StbGirderConcreteSwitch 
0 
制限なし 
RC 躯体切替位置 
 
StbGirderBarSwitch 
0 
制限なし 
配筋切替位置 
 
StbGirderSteelSwitch 
0 
制限なし 
S 躯体切替位置 
 
 
・補足 
(1) 曲がり梁（軸線が曲率を有する梁）は対象外。 
(2) 梁の部材座標系は、始端から終端をX、鉛直上向きがZ、の右手座標系とする。 
(3) 梁の基準点は、梁天端の幅中心位置とする。オフセットが省略された場合、始終端の基準点は節点
となる。 
(4) 回転の中心は梁の天端で、回転した後でオフセット・レベルを考慮する。省略した時は0 度（回転
なし）とみなす。 
(5) 断面の外端・内端指定は、断面の始端・終端を外端または内端（片持ち梁の場合は元端または先端）
と読み替える場合に指定する。この指定がある場合、梁断面（StbSecBeam_RC・StbSecBeam_S・
StbSecBeam_SRC）のisOutin は必ずtrue でなければならない。逆に梁断面のisOutin がtrue
でこの指定が省略された場合はアプリケーションが外端・内端を判断する。 
(6) ふかし厚さはハンチ部分も同様とする。

## Page 114

5-16 
 
 
 
 
 
 
 
 
 
 
 
(7) 該当節点がある場合は、節点ID を、始端に近い順に記述する。当要素は梁の折れ曲がりを表現す
るための要素であり、折れ曲がりがなく他の部材の端点となるだけの場合は記載する必要はない。 
(8) 構造種別と、断面ID が参照する要素名との対応は、それぞれRC－StbSecBeam_RC、S－
StbSecBeam_S、SRC－StbSecBeam_SRC、UNDEFINED－StbSecUndefined とする。 
(9) 省略された場合の扱いについては、StbCommon の補足説明を参照のこと。 
(10) 構造耐力上合成梁としてスタッド・床スラブを考慮しているかどうかをあらわす。 
(11) SRC の鉄骨で片持ち梁先端など、鉄骨位置が基準点と異なる場合に基準点から反対側の基準点方
向への距離を指定する。SRC 柱に剛接合するSRC 大梁を表現する場合は負の値となる。 
(12) 次のバージョン以降に実情を反映して詳細定義を行うこととする。 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
<StbGirders> 
<StbGirder id="59" name="1BA1" id_node_start="15" id_node_end="16" rotate="0.00" 
id_section="19" kind_structure="RC" isFoundation="false" offset_start_X="0.00" 
offset_start_Y="350.00" offset_start_Z="0.00" offset_end_X="0.00" offset_end_Y="-350.00" 
offset_end_Z="0.00"/> 
・・・ 
<StbGirder id="145" name="RGB1" id_node_start="34" id_node_end="37" rotate="0.00" 
id_section="35" kind_structure="S" isFoundation="false" offset_start_X="250.00" 
offset_start_Y="0.00" offset_start_Z="-150.00" offset_end_X="-200.00" offset_end_Y="0.00" 
offset_end_Z="-150.00"/> 
・・・ 
<StbGirder id="164" name="2BA1" id_node_start="24" id_node_end="25" rotate="0.00" 
id_section="43" kind_structure="SRC" isFoundation="false" offset_start_X="0.00" 
offset_start_Y="350.00" offset_start_Z="0.00" offset_end_X="0.00" offset_end_Y="-350.00" 
offset_end_Z="0.00"/> 
・・・ 
</StbGirders> 
終端基準点 
終端基準点 
始端基準点 
節点 
(始) 
節点 
(終) 
X 
Z 
Z 
Y 
（始端側から終端側を見た図） 
節点 
ふかし 
厚さ(上) 
オフセット(X
－
,Y
－
,Z
－
) 
ふかし 
厚さ(下) 
ふかし 
厚さ(左) 
ふかし 
厚さ(右)

## Page 115

5-17 
 
 
5.4.1.1. 大梁中間節点：StbGirderViaNode 
・概要 
説明 ：大梁の中間節点 
親要素：StbGirder 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
中間節点ID 
※(1) 
StbMemberOffsetList 
0 
制限なし 
中間節点オフセットリスト 
※(2) 
 
・補足 
中間節点とは、大梁の始端節点から終端節点の間に有る、始終端以外の節点である。大梁部材が折れ
曲がりのある直線の組合せの場合に、その形状を示すことができる。また、中間節点が、他の大梁、柱
およびブレースの端点となっている場合は、構造計算上、端点を有する部材と接続関係があることを示
す。 
 
(1) 該当節点ID を、始端に近い順に記述する。始端自身および終端自身は記述しない。 
(2) 中間節点が、StbColumn の始終点と同様のオフセットを有する場合に記述する。省略された場合、
各オフセットの値は0 とする。

## Page 116

5-18 
 
5.4.1.2. RC 躯体位置：StbGirderConcreteSwitch 
・概要 
説明 ：RC 躯体切替位置 
親要素：StbGirder 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
〇 
要素の順番 
※(1) 
distance 
double 
〇 
始端基準点からの距離 
※(2) 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbSecFigureBeam_RC、StbSecFigureBeam_SRC のorder と一致させる。番号は1 から始まる
連番とする。 
(2) order で示した部材の終了位置を梁始端の基準点からの距離で指定する。中間節点があり直線でな
い場合は、中間節点を経由した距離とする。

## Page 117

5-19 
 
5.4.1.3. 配筋切り替え位置：StbGirderBarSwitch 
・概要 
説明 ：配筋切り替え位置 
親要素：StbGirder 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
〇 
要素の順番 
※(1) 
distance 
double 
〇 
基準点からの距離 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbSecBarArrangementBeam_RC、StbSecBarArrangementBeam_SRCのorderと一致させる。 
(2) order で示した部材の終了位置を梁始端の基準点からの距離で指定する。中間節点があり直線でな
い場合は、中間節点を経由した距離とする。

## Page 118

5-20 
 
5.4.1.4. S 躯体位置：StbGirderSteelSwitch 
・概要 
説明 ：S 躯体位置 
親要素：StbGirder 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
〇 
要素の順番 
※(1) 
distance 
double 
〇 
基準点からの距離 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) StbSecSteelFigureBeam_S_Shape、StbSecSteelFigureBeam_SRC_Shape のorder と一致させ
る。 
(2) order で示した部材の終了位置を梁始端の基準点からの距離で指定する。中間節点があり直線でな
い場合は、中間節点を経由した距離とする。

## Page 119

5-21 
 
5.5. 小梁（複数）：StbBeams 
・概要 
説明 ：小梁情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbBeam 
1 
制限なし 
小梁情報 
 
 
 
5.5.1. 小梁：StbBeam 
・概要 
説明 ：小梁情報 
親要素：StbBeams 
 
・属性 
「5.4.1 大梁：StbGirder」と同様 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbBeamViaNode 
0 
1 
小梁中間節点 
 
StbBeamConcreteSwitch 
0 
制限なし 
RC 躯体切替位置 
 
StbBeamBarSwitch 
0 
制限なし 
配筋切替位置 
 
StbBeamSteelSwitch 
0 
制限なし 
S 躯体切替位置

## Page 120

5-22 
 
5.5.1.1. 小梁中間節点：StbBeamViaNode 
・概要 
説明 ：小梁中間節点 
親要素：StbBeam 
 
・属性 
無し 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
中間節点ID 
※(1) 
StbMemberOffsetList 
0 
制限なし 
中間節点オフセットリスト 
※(2) 
 
・補足 
「5.4.2. 大梁中間節点：StbGirderViaNode」と同様 
 
 
5.5.1.2. RC 躯体位置：StbBeamConcreteSwitch 
・概要 
説明 ：RC 躯体位置 
親要素：StbBeam 
 
・属性 
「RC 躯体位置：StbGirderConcreteSwitch」と同様 
 
・内容 
無し 
 
・子要素 
無し

## Page 121

5-23 
 
5.5.1.3. 配筋切り替え位置：StbBeamBarSwitch 
・概要 
説明 ：配筋切り替え位置 
親要素：StbBeam 
 
・属性 
「5.4.1 配筋切り替え位置：StbBeamBarSwitch」と同様 
 
・内容 
無し 
 
・子要素 
無し 
 
5.5.1.4. S 躯体切り替え位置：StbBeamSteelSwitch 
・概要 
説明 ：配筋切り替え位置 
親要素：StbBeam 
 
・属性 
「S 躯体位置：StbGirderSteelSwitch」と同様

## Page 122

5-24 
 
5.6. ブレース（複数）：StbBraces 
・概要 
説明 ：ブレース情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbBrace 
1 
制限なし 
ブレース情報

## Page 123

5-25 
 
5.6.1. ブレース：StbBrace 
・概要 
説明 ：ブレース情報 
親要素：StbBraces 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_start 
integer 
○ 
始端節点ID 
※(1) 
id_node_end 
integer 
○ 
終端節点ID 
※(1) 
rotate 
double 
 
回転角 
※(2), ※(4) 
id_section 
integer 
○ 
断面ID 
 
kind_structure 
string 
○ 
構造種別 
RC、S、SRC のいずれかの値 
当面はS のみ 
aim_offset_start_X 
double 
 
始端側狙い点オフセット（X
－
） 
※(2), ※(3) 
aim_offset_start_Y 
double 
 
始端側狙い点オフセット（Y
－
） 
※(2), ※(3) 
aim_offset_start_Z 
double 
 
始端側狙い点オフセット（Z
－
） 
※(2), ※(3) 
aim_offset_end_X 
double 
 
終端側狙い点オフセット（X
－
） 
※(2), ※(3) 
aim_offset_end_Y 
double 
 
終端側狙い点オフセット（Y
－
） 
※(2), ※(3) 
aim_offset_end_Z 
double 
 
終端側狙い点オフセット（Z） 
※(2), ※(3) 
cutback_start 
double 
 
始端側カットバック 
※(6) 
cutback_end 
double 
 
終端側カットバック 
※(6) 
feature_brace 
string 
 
ブレース特性 
引張り：TENSION、 
引張り圧縮： 
TENSIONANDCOMPRESSION 
のいずれかの値 
※(5) 
 
・内容 
無し 
 
・子要素 
無し

## Page 124

5-26 
 
・補足 
(1) 梁と同様に、部材座標系Z を全体座標系の上向きZ
－
とする。 
(2) ブレースの狙い点は、断面の外形を包絡する長方形の中心とする（下図参照）。 
 
 
 
 
 
 
 
 
(3) ブレースの狙い点オフセットは、節点からブレース狙い点までの距離で、全体座標系で定義す
る。オフセットが省略された場合、始終端の狙い点は節点となる（下図参照）。 
 
Y 
Z 
狙い点 
 (始端側から終端側を見た図) 
節点 
 狙い点オフセット(X
－
,Y
－
,Z
－
) 
 (始端側から終端側を見た図) 
狙い点 
節点 
 狙い点オフセット(X
－
,Y
－
,Z
－
)

## Page 125

5-27 
 
(4) 省略した時は0 度（回転なし）とみなす。回転はX 方向を0 度とし、始端から終端への進行方向
時計回りを正とし、回転した後でオフセットを考慮する。 
(5) 省略された場合、TENSION（引張り）とする。 
(6) 省略された場合、0 とする。 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
<StbBraces> 
<StbBrace id="180" name="1XA1B1" id_node_start="15" id_node_end="25" rotate="0.00" 
id_section="55" kind_structure="S" offset_start_X="0.00" offset_start_Y="0.00" 
offset_start_Z="-325.00" offset_end_X="0.00" offset_end_Y="0.00" offset_end_Z="-325.00" 
feature_brace="TENSION"/> 
・・・ 
</StbBraces>

## Page 126

5-28 
 
5.7. スラブ（複数）：StbSlabs 
・概要 
説明 ：スラブ情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSlab 
1 
制限なし 
スラブ情報

## Page 127

5-29 
 
5.7.1. スラブ：StbSlab 
・概要 
説明 ：スラブ情報 
親要素：StbSlabs 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_section 
integer 
○ 
断面ID 
 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値をとる。 
RC（RC スラブ） 
DECK（デッキ合成スラブ） 
PRECAST（既製スラブ） 
LOAD(荷重のみ) 
※(10) 
kind_slab 
string 
○ 
スラブ種類 
以下のいずれかの値をとる。 
NORMAL、CANTI 
 
strength_concrete 
string 
 
コンクリート強度 
※(9) 
thickness_add_top 
double 
 
ふかし厚さ（上） 
※(1) 
thickness_add_bottom 
double 
 
ふかし厚さ（下） 
※(1) 
direction_load 
string 
 
荷重伝達方向 
以下のいずれかの値をとる。 
1WAY、2WAY 
※(2) 
angle_load 
double 
 
荷重伝達方向「1WAY」の場合
の角度 
※(3) 
angle_main_bar_direction 
double 
 
主筋方向角度 
※(6), ※(4) 
angle_deck 
double 
 
デッキ角度 
※(7) 
isFoundation 
boolean 
○ 
基礎か否か 
基礎とは・・・ 
地反力や土圧･水圧を受
けている基礎スラブ 
 
・内容 
無し

## Page 128

5-30 
 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
節点ID 順序 
「4.2.4 順序のある節点
ID：StbNodeIdOrder」
と同様 
※(4), ※(5) 
StbSlabOffsetList 
0 
1 
オフセットリスト 
※(8) 
 
・補足 
(1) 省略された場合、当該属性がないものとする。 
(2) スラブ種類がNORMAL かつ荷重伝達方向が省略された場合、2WAY とする。スラブ種類が
CANTI かつ荷重伝達方向が省略された場合、1WAY とする。 
(3) 荷重伝達方向が1WAY の場合に記述するが、省略された場合は短編方向を荷重伝達方向とする。
ただし、正方形スラブの場合は必須とする。スラブ種類がCANTI かつ荷重伝達方向が省略された
場合（1WAY）、記述は必須とする。角度の記述は、部材座標X 軸に対するZ 軸回りの角度を記
述する。 
(4) スラブの部材座標系は、１点目から２点目がX、鉛直上向きがZ の右手座標系とする。 
(5) 節点の並びは下から見て時計まわりとする。 
(6) 省略された場合、短辺方向を主筋方向とする。正方形スラブで省略された場合は未指定とする。角
度の記述は、部材座標X 軸に対するZ 軸回りの角度を記述する。 
(7) フラットデッキスラブ、トラス筋付きデッキスラブ、デッキ合成スラブ、既製スラブに適応する。
省略された場合、短辺方向を敷き込み方向とする。正方形スラブで省略された場合は未指定とす
る。角度の記述は、部材座標X 軸に対するZ 軸回りの角度を記述する。 
(8) オフセットリストが省略された場合、スラブ上端の周辺基準点は節点とする。 
(9) 省略された場合の扱いについては、StbCommon の補足説明を参照のこと。 
(10) RC（RC スラブ）を選択した場合、RC スラブ断面StbSecSlab_RC の子要素で在来スラブおよび
トラス筋付きデッキスラブを指定することができる。また、在来スラブを指定した場合は、さらに
その子要素で合板型枠またはフラットデッキを指定することができる。詳細はStbSecSlab_RC を
参照のこと。

## Page 129

5-31 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
（1 点目） 
（2 点目） 
（3 点目） 
（4 点目） 
右手座標系 
・通常スラブ 
スラブ種類 
・キャンチスラブ 
<StbSlabs> 
<StbSlab id="107" name="1SA1.1" id_section="71" kind_structure="RC" 
kind_slab="NORMAL" direction_load="2WAY" isFoundation="false"> 
<StbNodeIdOrder>15 18 277 276</StbNodeIdOrder> 
<StbSlabOffsetList> 
<StbSlabOffset id_node="15" offset_X="175.00" offset_Y="175.00" offset_Z="0.00"/> 
<StbSlabOffset id_node="18" offset_X="-175.00" offset_Y="175.00" offset_Z="0.00"/> 
<StbSlabOffset id_node="277" offset_X="-175.00" offset_Y="-150.00" offset_Z="0.00"/> 
<StbSlabOffset id_node="276" offset_X="175.00" offset_Y="-150.00" offset_Z="0.00"/> 
</StbSlabOffsetList> 
</StbSlab> 
・・・ 
</StbSlabs> 
オフセット(X
－
,Y
－
,Z
－
) 
節点 
節点 
節点 
節点 
Z 
Y 
X 
ふかし厚さ（上） 
ふかし厚さ（下）

## Page 130

5-32 
 
5.7.1.1. スラブオフセットリスト：StbSlabOffsetList 
・概要 
説明 ：スラブのオフセットリスト 
親要素：StbSlab 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSlabOffset 
1 
制限なし 
スラブのオフセット 
 
 
 
5.7.1.1.1. スラブオフセット：StbSlabOffset 
・概要 
説明 ：スラブのオフセット 
親要素：StbSlabOffsetList 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_node 
integer 
○ 
節点ID 
 
offset_X 
double 
○ 
X
－方向のオフセット 
 
offset_Y 
double 
○ 
Y
－方向のオフセット 
 
offset_Z 
double 
○ 
Z
－方向のオフセット 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 131

5-33 
 
5.8. 壁（複数）：StbWalls 
・概要 
説明 ：壁情報（複数） 
親要素：StbMembers 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWall 
1 
制限なし 
壁情報

## Page 132

5-34 
 
5.8.1. 壁：StbWall 
・概要 
説明 ：壁情報 
親要素：StbWalls 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_section 
integer 
○ 
断面ID 
 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値をとる。 
RC(RC 壁) 
LOAD(荷重のみ) 
 
kind_layout 
string 
○ 
壁種別 
以下のいずれかの値をとる。 
ON_GIRDER（大梁上）、 
ON_BEAM（小梁上）、 
ON_SLAB（スラブ上） 
 
strength_concrete 
string 
 
コンクリート強度 
※(9) 
thickness_add_right 
double 
 
ふかし厚さ（右） 
※(1), ※(3), ※(5) 
thickness_add_left 
double 
 
ふかし厚さ（左） 
※(1), ※(3), ※(5) 
kind_wall 
string 
 
耐力区分 
以下のいずれかの値をとる。 
WALL_NORMAL（一般壁）、
WALL_SHEAR（耐力壁） 
※(2) 
slit_upper 
double 
 
構造スリット（上） 
※(1) 梁底からの
距離,※(10) 
slit_bottom 
double 
 
構造スリット（下） 
※(1) 梁天からの
距離,※(10) 
slit_right 
double 
 
構造スリット（右） 
※(1) 柱面からの
距離,※(10) 
slit_left 
double 
 
構造スリット（左） 
※(1) 柱面からの
距離,※(10)

## Page 133

5-35 
 
type_outside 
string 
 
外側のタイプ 
以下のいずれかの値をとる。 
TYPE_PLUS（正側） 
TYPE_MINUS（負側） 
※(6), ※(3) 
type_press 
string 
 
土圧を受ける状況 
以下のいずれかの値をとる。 
NONE（土圧壁ではない） 
ONESIDE（片側土圧） 
BOTH（両側土圧） 
※(7) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
節点ID リスト 
「4.2.4 順序のある節点ID：
StbNodeIdOrder」と同様 
※(4), ※(3) 
StbWallOffsetList 
0 
1 
オフセットリスト 
※(8) 
 
・補足 
(1) 省略された場合、当該属性がないものとする。 
(2) 省略された場合、WALL_NORMAL（一般壁）とする。 
(3) 壁の部材座標系は、１点目から２点目がX、鉛直上向きがY の右手座標系とする。 
(4) StbNodeIdOrder に納める節点の順番は梁上なら梁の始端に近い方、スラブ上なら原点に近い方
から順番に格納する。 
(5) ふかしの厚さの右左の判定は部材座標系Z 方向の正側を右、負側を左とする。 
 
 
 
 
 
 
 
 
 
 
Y 
X 
Z 
節点 
節点 
節点 
節点 
右手座標系 
オフセット 
（X
－
,Y
－
,Z
－
） 
（1 点目） 
（2 点目） 
（3 点目） 
（4 点目） 
FL 
FL 
ふかし厚さ 
（左） 
ふかし厚さ 
（右）

## Page 134

5-36 
 
 
(6)  type_Press がONESIDE（片側土圧）の場合、RC 壁断面形状（StbSecFigureWall_RC）がテー
パー（StbSecWall_RC_Taper）の場合、RC 壁断面配筋（StbSecBarArrangementWall_RC）の
かぶり厚さが内外異なる場合、また、配筋がRC 壁断面配筋（内外異なる）
（StbSecBarWall_RC_InsideAndOutside）の場合、記述は必須とする。それ以外の場合は、記述
は無効とする。外側タイプに関しては部材座標系を元に正側・負側を設定する。 
 
 
(7) ONESIDE（片側土圧）の場合、type_outside で指定された「外側」が土圧を受ける面とする。省
略された場合、土圧壁でないものとする。 
(8) オフセットリストが省略された場合、壁厚中央の周辺基準点は節点とする。 
(9) 省略された場合の扱いについては、StbCommon の補足説明を参照のこと。 
(10) 構造スリットの距離はスリットの梁面（柱面）までの近い方とする。（上であれば垂壁長さ、下で
あれば腰壁長さ、左右であれば袖壁長さ） 
 
 
 
▼梁底 
▲梁天 
▲ 
柱面 
▲ 
柱面 
～ 
～ 
距離

## Page 135

5-37 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
<StbWalls> 
<StbWall id="160" name="1WC1G" id_section="95" kind_structure="RC" 
kind_layout="ON_GIRDER" kind_wall="WALL_NORMAL"> 
<StbNodeIdOrder>17 20 29 26</StbNodeIdOrder> 
<StbWallOffsetList> 
<StbWallOffset id_node="17" offset_X="350.00" offset_Y="0.00" offset_Z="0.00"/> 
<StbWallOffset id_node="20"offset_X="-350.00" offset_Y="0.00" offset_Z="0.00"/> 
<StbWallOffset id_node="29" offset_X="-350.00" offset_Y="0.00" offset_Z="-650.00"/> 
<StbWallOffset id_node="26" offset_X="350.00" offset_Y="0.00" offset_Z="-650.00"/> 
</StbWallOffsetList> 
</StbWall> 
・・・ 
</StbWalls>

## Page 136

5-38 
 
5.8.1.1. 壁オフセットリスト：StbWallOffsetList 
・概要 
説明 ：壁のオフセットリスト 
親要素：StbWall 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWallOffset 
1 
制限なし 
壁のオフセット 
 
 
 
5.8.1.1.1. 壁オフセット：StbWallOffset 
・概要 
説明 ：壁のオフセット 
親要素：StbWallOffsetList 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_node 
integer 
○ 
節点ID 
 
offset_X 
double 
○ 
X
－方向のオフセット 
 
offset_Y 
double 
○ 
Y
－方向のオフセット 
 
offset_Z 
double 
○ 
Z
－方向のオフセット 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 137

5-39 
 
5.9. 免震装置（複数）：StbIsolatingDevices 
・概要 
説明 ：免震装置部材情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbIsolatingDevice 
1 
制限なし 
免震装置 
 
 
・補足 
免震装置StbIsolatingDevice は、いわゆる支承材として扱われる免震装置部材を対象とする。なお、
免震層に配置される減衰材は制振装置の節で扱い、復元材は対象外とする。 
当面、国内の主要な免震装置部材を扱う応力解析プログラム、構造計算プログラムが適用範囲とする
支承材を対象とし、それ以外の免震装置については、今後の各プログラムの対応状況や装置の普及状況
に応じて対象とする。

## Page 138

5-40 
 
5.9.1. 免震装置：StbIsolatingDevice 
・概要 
説明 ：免震装置部材情報 
親要素：StbIsolatingDevices 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_start 
integer 
○ 
始端節点ID 
 
id_node_end 
integer 
○ 
終端節点ID 
 
rotate 
double 
 
回転角 
※(1) 
id_section 
integer 
○ 
免震装置断面ID 
 
offset_start_X 
double 
 
始端側オフセット（X
－
） 
※(2) 
offset_start_Y 
double 
 
始端側オフセット（Y
－
） 
※(2) 
offset_start_Z 
double 
 
始端側オフセット（Z
－
） 
※(2) 
offset_end_X 
double 
 
終端側オフセット（X
－
） 
※(2) 
offset_end_Y 
double 
 
終端側オフセット（Y
－
） 
※(2) 
offset_end_Z 
double 
 
終端側オフセット（Z
－
） 
※(2) 
number 
integer 
 
装置台数 
※(6) 
isMemberArrangement_start 
boolean 
 
始端側接合部・部材配置あるか否か 
※(3) 
kind_structure_start 
string 
 
始端側接合部・構造種別 
以下のいずれかとする。 
RC,S,SRC,CFT,UNDEFINED 
※(4) 
id_section_start 
integer 
 
始端側接合部・断面ID 
※(4) 
connection_length_start 
double 
 
始端側接合部・接合長さ 
※(5) 
isMemberArrangement_end 
boolean 
 
終端側接合部・部材配置あるか否か 
※(3) 
kind_structure_end 
string 
 
終端側接合部・構造種別 
以下のいずれかとする。 
RC,S,SRC,CFT,UNDEFINED 
※(4) 
id_section_end 
integer 
 
終端側接合部・断面ID 
※(4) 
connection_length_end 
double 
 
終端側接合部・接合長さ 
※(5)

## Page 139

5-41 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbIsolatingDevicePosition 
0 
制限なし 
免震装置配置位置（複数台数配置時） 
※(6) 
 
・補足 
免震装置部材の形状例を下図に示す。「始端」は装置本体の下側、「終端」は装置本体の上側に定義す
るものとし、始終端における基準点は、装置の平面形状に対して図心の位置とする。 
免震装置部材は、通常、下図のように、装置本体の上下に装置を固定する取付け躯体（以下、接合部）
を配置することが多く、免震装置部材のモデル化においては、図のように接合部を含めた一体のモデル
（以下、接合部一体型）として定義する場合がある。 
 
 
一方、下図のように、装置本体のみをこの要素で定義して、接合部は柱部材など別の部材要素とし、
節点を介して定義することも可能である。 
 
本要素ではどちらも適用範囲とする。 
接合部については、取付け躯体部分を柱断面とみなせるモデルを対象とする。 
 
 
基準点＝図心
節点
節点
終端
X
Y
Ｚ
Y
X
Z
接合部
接合部
装置本体
始端
節点
基準点
基準点
終端
始端
StbIsolatingDevice
接合部（StbColumn）
装置本体（StbIsolatingDevice）
節点
基準点
基準点
節点
節点
節点
接合部（StbColumn）
終端
始端

## Page 140

5-42 
 
(1) 回転はX 方向を0 度とし、始端から終端への進行方向時計回りを正とし、断面を回転した後にオ
フセットを考慮する。省略した場合は0 度（回転なし）とみなす。 
 
(2) オフセットは、節点位置から基準点までの長さを全体座標系で表した値とする。省略された場合
は、0 とする。始終端のそれぞれXYZ 全ての値が0 の場合、基準点と節点は一致する。 
 
 
(3) 始端または終端に、接合部一体型として部材を配置する場合はtrue、接合部一体型として部材を
配置しない場合はfalse とする。省略した場合はfalse とする。 
 
 
(4) 接合部一体型として部材を配置する場合、接合部の構造種別としてRC、S、SRC、CFT、およ
びUNDEFINED のいずれかの値を指定し、対応する柱断面ID を指定する。 
構造種別と、対応する柱断面ID が参照する要素名は、それぞれRC－StbSecColumn_RC、S－
StbSecColumn_S、SRC－StbSecColumn_SRC、CFT－StbSecColumn_CFT、UNDEFINED－
StbSecUndefined とする。接合部一体型として部材を配置する場合, 構造種別および断面ID は
必須とし、接合部一体型として部材を配置しない場合は指定してはならない。 
 
 
回転＋
節点
基準点
オフセット（X）
オフセット（Y）
節点
節点
オフセット（X , Y）
オフセット（Z）
終端
始端
基準点
基準点
節点
節点
節点
基準点
基準点
節点
true：
false：

## Page 141

5-43 
 
(5) 接合部一体型として部材を配置する場合、装置本体と接合部材との接合位置から基準点までの長
さ（以下、接合長さ）を指定する。接合長さは、参照する免震装置断面から定義される装置本体の
高さと階の高さ等から決まるものと想定されるため、始終端双方を接合部一体型として部材配置
する場合は、どちらかを必須とする。接合部一体型として部材を配置しない場合は指定してはな
らない。 
 
 
(6) 装置が複数台ある場合、台数を指定する。省略された場合は、1 とする。装置が複数台あって、そ
の各々の位置を指定する必要がある場合は、子要素StbIsolatingDevicePosition にてオフセット
値を指定する。子要素を指定する場合、その回数は台数と一致しなければならない。 
 
 
・例 
 
 
 
 
 
 
 
 
始端節点
終端節点
接合長さ（終端）
接合長さ（始端）
装置本体の高さ
基準点
基準点
接合位置
接合位置
  <StbIsolatingDevices> 
<StbIsolatingDevice id="67" name="LRB1" id_node_start="40" id_node_end="49" 
id_section="1"/> 
・・・ 
  </StbIsolatingDevices>

## Page 142

5-44 
 
5.9.1.1. 免震装置配置位置（複数台数配置時）：StbIsolatingDevicePosition 
・概要 
説明 ：装置台数が複数台ある場合の配置位置 
親要素：StbIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
offset_start_X 
double 
 
始端側オフセット（X） 
 
offset_start_Y 
double 
 
始端側オフセット（Y） 
 
offset_end_X 
double 
 
終端側オフセット（X） 
 
offset_end_Y 
double 
 
終端側オフセット（Y） 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
免震装置が複数台あって、その各々の位置を指定する必要がある場合は、この要素にて指定する。下
図に示す、始端、終端それぞれの接合位置からのオフセット値を、装置台数分、部材座標系にて指定す
る。従って、子要素回数は台数と一致しなければならない。 
 
 
 
 
接合長さ（終端）
接合長さ（始端）
装置本体の高さ
接合位置
接合部
接合部
装置本体
StbIsolatingDevice
始端基準点
終端基準点
終端側オフセット（X , Y）
接合位置
始端側オフセット（X , Y）

## Page 143

5-45 
 
5.10. 制振装置（複数）： StbDampingDevices 
・概要 
説明 ：２節点指定型の制振装置部材情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbDampingDevice 
1 
制限なし 
制振装置 
 
 
・補足 
制振装置StbDampingDevice は、主架構に組み込まれる制振ダンパーを対象とする。取付け部材を介
して組み込まれる場合、取付け部材も対象とすることができる。また、免震層に配置される減衰材も対
象とする。 
当面、これらの制振装置部材を扱う、国内の主要な応力解析プログラム、構造計算プログラムが適用
範囲とする制振装置部材を対象とする。それ以外の制振装置ならびに、戸建て住宅向けに開発されてい
る装置等については、今後の各プログラムの対応状況や装置の普及状況に応じて対象とする。

## Page 144

5-46 
 
対象となる制振装置部材は、通常、主架構に様々な形状の部材として組み込まれており、多様な分類
方法が考えられるが、組み込み形状を、国内の主要な応力解析プログラム、構造計算プログラムの指定
方法を参考にして、下表のように分類する。 
分類タイプ 
適用される形状の例 
赤：制振装置 黒：取付け部材 
想定される解析モデル 
２節点
指定型 
始端、終
端の２
節点を
指定 
ブレースタイプ 
 
  
材軸方向考慮 
柱（間柱）タイプ 
 
 
せん断方向考慮 
梁タイプ 
 
 
せん断 or 材軸方向考慮 
４節点
指定型 
壁面を
構成す
る４節
点を指
定し、パ
ターン
配置 
壁タイプ 
 
 
組合せブレースタイプ 
 
 
間柱タイプ 
 
 
シアリンクタイプ 
 
 
免震層水平配置タイプ 
 
 
 
   本仕様書では、以上の分類を踏まえて、以下の２つの要素で定義する。 
・「２節点指定型」StbDampingDevice（本節） 
・「４節点指定型」StbFrameDampingDevice（次節） 
 
制振装置部材を扱うプログラムにおいては、「４節点指定型」を採用していない場合もある。その場
合は、制振装置を「２節点指定型」の要素として、取付け部材をStbBrace などのStbMembers におけ
る他の子要素として、それぞれ節点を介して定義することとする。また、上記分類に当てはまらない、
制振装置部材を方杖のように配置、柱部材の中間に配置、等の場合についても、同様に節点を介して定
義することとする。

## Page 145

5-47 
 
5.10.1. 制振装置：StbDampingDevice 
・概要 
説明 ：２節点指定型の制振装置部材情報 
親要素：StbDampingDevices 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_start 
integer 
○ 
始端節点ID 
 
id_node_end 
integer 
○ 
終端節点ID 
 
rotate 
double 
 
回転角 
 
id_section 
integer 
○ 
制振装置断面ID 
 
type_shape 
string 
○ 
部材形状の分類 
以下のいずれかとする。 
BRACE,POST,GIRDER 
※(1) 
offset_start_X 
double 
 
始端側オフセット（X
－
） 
 
offset_start_Y 
double 
 
始端側オフセット（Y
－
） 
 
offset_start_Z 
double 
 
始端側オフセット（Z
－
） 
 
offset_end_X 
double 
 
終端側オフセット（X
－
） 
 
offset_end_Y 
double 
 
終端側オフセット（Y
－
） 
 
offset_end_Z 
double 
 
終端側オフセット（Z
－
） 
 
number 
integer 
 
装置台数 
※(5) 
isMemberArrangement_start 
boolean 
 
始端側接合部・部材配置あるか否か 
※(2) 
kind_structure_start 
string 
 
始端側接合部・構造種別 
以下のいずれかとする。 
RC,S,SRC,CFT,UNDEFINED 
※(3) 
id_section_start 
integer 
 
始端側接合部・断面ID 
※(3) 
connection_length_start 
double 
 
始端側接合部・接合長さ 
※(4) 
isMemberArrangement_end 
boolean 
 
終端側接合部・部材配置あるか否か 
※(2) 
kind_structure_end 
string 
 
終端側接合部・構造種別 
以下のいずれかとする。 
RC,S,SRC,CFT,UNDEFINED 
※(3) 
id_section_end 
integer 
 
終端側接合部・断面ID 
※(3)

## Page 146

5-48 
 
connection_length_end 
double 
 
終端側接合部・接合長さ 
※(4) 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbDampingDevicePosition 
0 
制限なし 
制振装置配置位置（複数台数配置時） 
※(5) 
 
・補足 
前節に示した分類のうち、「２節点指定型」の要素を定義する。 
 
(1) 配置形状に応じた分類を指定する。 
 
部材形状の分類 
適用される形状の例 
赤：制振装置 黒：取付け部材 
想定される 
解析モデル 
分類に対応する 
StbMembers 子要素 
BRACE 
ブレース
タイプ 
 
  
材軸方向考慮 
StbBrace 
POST 
柱（間柱）
タイプ 
 
 
せん断方向考慮 
StbPost 
GIRDER 
梁タイプ 
 
 
せん断 or 材軸方向考慮 
StbGirder 
 
    「２節点指定型」の分類を、「想定される解析モデル」に応じて指定する。「分類に対応する
StbMembers 子要素」とは、それぞれ、所属属性の構成が類似するStbMembers の子要素を指し、
「適用される形状の例」に示す形状を想定している。 
始終端における基準点の定義および部材座標系の定義は、それぞれ「分類に対応するStbMembers
子要素」の定義方法にならうものとする。「始端」「終端」「回転角」および「オフセット」の定義
についても、同様とする。

## Page 147

5-49 
 
(2) 始端または終端に、接合部一体型として部材を配置する場合はtrue、接合部一体型として部材を
配置しない場合はfalse とする。省略した場合はfalse とする。 
 
免震装置部材と同様に、制振装置部材においても、装置本体の取付け部に装置を固定する取付
け躯体（以下、接合部）を配置し、モデル化においては、接合部を含めた一体のモデル（以下、接
合部一体型）として定義する場合がある。柱（間柱）タイプの場合、下図のようになる。 
 
 
 
一方、装置本体のみをこの要素で定義して、接合部は別の部材要素とし、節点を介して定義す
ることも可能である。 
 
 
本要素ではどちらも適用範囲とする。 
接合部については、取付け躯体部分を柱断面とみなせるモデルを対象とする。ただし、梁タイ
プの場合は、大梁断面とみせるモデルを対象とする。 
 
 
 
Y
X
Z
接合部
接合部
装置本体
節点
節点
基準点
基準点
終端
始端
終端
始端
StbDampingDevice
接合部（StbColumn）
装置本体（StbDampingDevice）
節点
基準点
接合部（StbColumn）
終端
始端
基準点
節点
節点
節点

## Page 148

5-50 
 
(3) 接合部一体型として部材を配置する場合、部材形状の分類に応じて、下表に示す構造種別と対応
する断面ID を指定する。接合部一体型として部材を配置する場合、構造種別および断面ID は必
須とし、接合部一体型として部材を配置しない場合は指定してはならない。 
部材形状の分類 
接合部・構造種別 
断面ID が参照する要素 
POST 
BRACE 
RC 
StbSecColumn_RC 
S 
StbSecColumn_S 
SRC 
StbSecColumn_SRC 
CFT 
StbSecColumn_CFT 
UNDEFINED 
StbSecUndefined 
GIRDER 
RC 
StbSecBeam_RC 
S 
StbSecBeam_S 
SRC 
StbSecBeam_SRC 
UNDEFINED 
StbSecUndefined 
 
(4) 接合部一体型として部材を配置する場合、装置本体と接合部材との接合位置から基準点までの長
さ（接合長さ）を指定する。柱（間柱）タイプの場合、下図のようになる。接合部一体型として部
材を配置する場合は必須とし、接合部一体型として部材を配置しない場合は指定してはならない。 
 
 
(5) 装置が複数台ある場合、台数を指定する。省略された場合は、1 とする。装置が複数台あって、そ
の各々の位置を指定する必要がある場合は、子要素StbDampingDevicePosition にてオフセット
値を指定する。子要素を指定する場合、その回数は台数と一致しなければならない。 
 
 
基準点
始端節点
終端節点
基準点＝図心
接合位置
接合位置
終端接合長さ
始端接合長さ
装置本体高さ

## Page 149

5-51 
 
・例 
 
 
 
 
 
 
 
 
 
<StbDampingDevices> 
  <StbDampingDevice id="68" name="DMP1" 
        id_node_start="101" id_node_end="102" id_section="1" type_shape="BRACE"/> 
  <StbDampingDevice id="69" name="DMP2" 
        id_node_start="103" id_node_end="104" id_section="2" type_shape="POST" 
        isConnection_start="true" kind_structure_start="RC" id_section_start="10" 
connection_length_start="1000.0" 
isConnection_end="true" kind_structure_end="RC" id_section_end="10" 
connection_length_end="1000.0"/> 
</StbDampingDevices>

## Page 150

5-52 
 
5.10.2. 制振装置配置位置（複数台数配置時）：StbDampingDevicePosition 
・概要 
説明 ：装置台数が複数台ある場合の配置位置 
親要素：StbDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
offset_start_X 
double 
 
始端側オフセット（X） 
 
offset_start_Y 
double 
 
始端側オフセット（Y） 
 
offset_end_X 
double 
 
終端側オフセット（X） 
 
offset_end_Y 
double 
 
終端側オフセット（Y） 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
制振装置が複数台あって、その各々の位置を指定する必要がある場合は、この要素にて指定する。始
端、終端それぞれの接合位置からのオフセット値を、装置台数分、部材座標系にて指定する。従って、
子要素回数は台数と一致しなければならない。

## Page 151

5-53 
 
5.11. 制振装置（フレーム）（複数）： StbFrameDampingDevices 
・概要 
説明 ：４節点指定型の制振装置部材情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbFrameDampingDevice 
1 
制限なし 
制振部材（フレーム） 
 
 
・補足 
前節に示した分類のうち、「４節点指定型」の要素を定義する。

## Page 152

5-54 
 
5.11.1. 制振装置（フレーム）：StbFrameDampingDevice 
・概要 
説明 ：４節点指定型の制振装置部材情報 
親要素：StbDampingDevices 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_section 
integer 
○ 
制振装置断面ID 
 
type_shape 
string 
○ 
部材形状の分類 
以下のいずれかとする。 
WALL,BRACES,POSTS,SHEARLINK, 
HORIZONTALLY_SI_LAYER 
※(1) 
minor_type_shape 
string 
 
同、小分類 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbNodeIdOrder 
1 
1 
節点ID リスト 
(順序のある節点ID) 
※(2) 
StbFrameDampingDeviceOffsetList 
0 
1 
オフセットリスト 
※(3) 
StbFrameDampingDeviceConfiguration 
0 
1 
制振装置（フレーム）構成 
※(4)

## Page 153

5-55 
 
・補足 
(1) 配置形状に応じた分類を指定する。 
 
部材形状の分類 
適用される形状の例 
赤：制振装置 
黒：取付け部材 
想定される 
解析モデル 
WALL 
壁タイプ 
 
 
BRACES 
組合せブレースタイプ 
 
 
POSTS 
間柱タイプ 
 
 
SHEARLINK 
シアリンクタイプ 
 
 
HORIZONTALLY 
_SI_LAYER 
免震層水平配置タイプ 
 
 
 
以下の分類においては、小分類にて形状特定してもよい。 
部材形状の分類 
小分類：想定される解析モデル 
BRACES 
追加節点：1  
取付け部材：0  
NORMAL_V 
 
INVERTED_V 
 
 
BRACES 
追加節点：2  
取付け部材：0  
SEPARATED_V 
 
INVERTED_SEPARATED_V 
 
 
SHEARLINK 
追加節点：1  
取付け部材：2  
LOWER_BOTH 
 
LOWER_2ND_SIDE 
LOWER_1ST_SIDE 
 
UPPER_BOTH 
 
UPPER_3RD_SIDE 
 
UPPER_4TH_SIDE

## Page 154

5-56 
 
部材形状の分類 
小分類：想定される解析モデル 
SHEARLINK 
追加節点：2  
取付け部材：2  
LOWER_VERTICAL 
 
 
 
UPPER_VERTICAL 
 
 
 
 
(2) ４節点を構成する節点ID を指定する。ここに掲げたStbNodeIdOrder を用いる指定方法は、壁：
StbWall の指定方法と同一であり、基準点および部材座標系の定義方法も、StbWall にならい、オ
フセットの扱いも同様とする。StbNodeIdOrder に納める節点の順番は梁上なら梁の始端に近い
方、スラブ上なら原点に近い方から順番に格納する。 
 
(3) ４節点について、オフセットがある場合に指定する。省略された場合は、基準点と節点は一致する
ものとする。 
(4) 構成する制振装置、および制振装置と架構を接合する取付け部材（以下、接合部部材）の配置を、
節点ID およびオフセットにて指定する。 
接合部部材の配置によっては、節点ID リストで指定する４節点以外の節点（以下、追加節点）が
必要な場合がある。追加節点は、この要素以外の他要素との関連がない場合は、StbNode にて属
性 kind をOTHER として、別途定義する。追加節点が所属する階を、１・２点目または３・４点
目と同じ階とする必要がある場合は、階のID が同一なStbStory の子要素StbNodeIdList に指定
する。 
組合せブレースタイプにおいては、「２節点指定型」に示すような「接合部一体型」の指定は、本
要素の対象外とする。 
 
 
（２点目）
（４点目）
（３点目）
（１点目）
X
Y
Z
基準点
基準点
基準点
基準点
StbFrameDamperを構成する面

## Page 155

5-57 
 
・例 
 
 
 
 
 
 
 
 
 
 
<StbFrameDampingDevices> 
  <StbFrameDampingDevice id="78" name="DMP3" id_section="2" type_shape="WALL"> 
   <StbNodeIdOrder>1001 274 139 1002</StbNodeIdOrder> 
  </StbFrameDampingDevice> 
  <StbFrameDampingDevice id="79" name="DMP4" id_section="3" type_shape="POSTS"> 
   <StbNodeIdOrder>1003 354 287 1004</StbNodeIdOrder> 
   <StbFrameDampingDeviceConfiguration> 
   <StbFrameDampingDeviceMember id_node_start="1005" id_node_end="1006"/> 
   </StbFrameDampingDeviceConfiguration> 
  </StbFrameDampingDevice> 
</StbFrameDampingDevices>

## Page 156

5-58 
 
5.11.1.1. 制振装置（フレーム）オフセットリスト：StbFrameDampingDeviceOffsetList 
・概要 
説明 ：４節点指定型の制振装置端点節点オフセットリスト 
親要素：StbFrameDampingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbFrameDampingDeviceOffset 
1 
4 
制振装置（フレーム）のオフセット 
 
 
 
5.11.1.1.1. 制振装置（フレーム）オフセット：StbFrameDampingDeviceOffset 
・概要 
説明 ：４節点指定型の制振装置端点節点オフセット 
親要素：StbFrameDampingDeviceOffsetList 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_node 
integer 
○ 
節点ID 
 
offset_X 
double 
○ 
X
－
方向のオフセット 
 
offset_Y 
double 
○ 
Y
－
方向のオフセット 
 
offset_Z 
double 
○ 
Z
－
方向のオフセット 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 157

5-59 
 
5.11.1.2. 制振装置（フレーム）構成：StbFrameDampingDeviceConfigulation 
・概要 
説明 ：４節点指定型の制振装置部材構成 
親要素：StbFrameDampingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbFrameDampingDeviceMember 
1 
制限なし 
制振装置（フレーム）構成・制振
装置部材 
※(1) 
StbFrameDampingDeviceConnection 
0 
制限なし 
制振装置（フレーム）構成・接合
部部材 
※(2) 
 
・補足 
(1) 構成する要素のうち、制振装置部材の配置を指定する。 
(2) 構成する要素のうち、制振装置部材以外の配置を指定する。取付け部材のほか、装置機構の一部と
なるような部材に関しても、接合部部材としてこの要素で指定する。

## Page 158

5-60 
 
・例 
   間柱タイプ、シアリンクタイプにおける配置例 
 
間柱タイプ（id  1）            シアリンクタイプ（id  2） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
１
２ 
４
３ 
５ 
６ 
８ 
７ 
1001 
1003 
1004 
1001 
1002 
<StbFrameDampingDevices> 
  <StbFrameDampingDevice id="1" name="DMP1" id_section="101" kind="POSTS"> 
    <StbNodeIdOrder>1 2 3 4</StbNodeIdOrder> 
    <StbFrameDampingDeviceConfiguration> 
      <StbFrameDampingDeviceMember id_node_start="1001" id_node_end="1002"/> 
      <StbFrameDampingDeviceConnection 
       kind="POST" id_member="13" id_node_start="1003" id_node_end="1001"/> 
      <StbFrameDampingDeviceConnection 
       kind="POST" id_member="13" id_node_start="1002" id_node_end="1004"/> 
    </StbDampingDeviceConfiguration> 
  </StbFrameDampingDevice> 
  <StbFrameDampingDevice id="2" name="DMP2" id_section="102" kind="SHEARLINK"> 
    <StbNodeIdOrder>5 6 7 8</StbNodeIdOrder> 
    <StbFrameDampingDeviceConfiguration> 
      <StbFrameDampingDeviceMember id_node_start="1001" id_node_end="6"/> 
      <StbFrameDampingDeviceConnection 
       kind="BRACE" id_member="14" id_node_start="1001" id_node_end="8"/> 
      <StbFrameDampingDeviceConnection 
       kind="BRACE" id_member="14" id_node_start="1001" id_node_end="7"/> 
    </StbDampingDeviceConfiguration> 
  </StbFrameDampingDevice> 
</StbFrameDampingDevices>

## Page 159

5-61 
 
5.11.1.2.1. 制振装置（フレーム）構成・制振装置部材：StbFrameDampingDeviceMember 
・概要 
説明 ：４節点指定型の制振装置部材構成における制振装置部材の配置 
親要素：StbFrameDampingDeviceConfigulation 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_node_start 
integer 
○ 
構成部材始端節点ID 
 
id_node_end 
integer 
○ 
構成部材終端節点ID 
 
offset_start_X 
double 
 
構成部材始端側オフセット（X
－
） 
 
offset_start_Y 
double 
 
構成部材始端側オフセット（Y
－
） 
 
offset_start_Z 
double 
 
構成部材始端側オフセット（Z
－
） 
 
offset_end_X 
double 
 
構成部材終端側オフセット（X
－
） 
 
offset_end_Y 
double 
 
構成部材終端側オフセット（Y
－
） 
 
offset_end_Z 
double 
 
構成部材終端側オフセット（Z
－
） 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
構成する制振装置部材の配置を、節点ID およびオフセットにて指定する。

## Page 160

5-62 
 
5.11.1.2.2. 制振装置（フレーム）構成・接合部部材：StbFrameDampingDeviceConnection 
・概要 
説明 ：４節点指定型の制振装置部材構成における接合部部材の配置 
親要素：StbFrameDampingDeviceConfigulation 
 
・属性 
属性名 
型 
必須 
説明 
補足 
kind_structure 
string 
○ 
接合部・部材種類 
以下のいずれかとする 
BRACE,POST,GIRDER 
 
id_member 
integer 
○ 
接合部・部材ID 
 
id_node_start 
integer 
○ 
接合部・始端節点ID 
 
id_node_end 
integer 
○ 
接合部・終端節点ID 
 
offset_start_X 
double 
 
接合部・始端側オフセット（X
－
） 
 
offset_start_Y 
double 
 
接合部・始端側オフセット（Y
－
） 
 
offset_start_Z 
double 
 
接合部・始端側オフセット（Z
－
） 
 
offset_end_X 
double 
 
接合部・終端側オフセット（X
－
） 
 
offset_end_Y 
double 
 
接合部・終端側オフセット（Y
－
） 
 
offset_end_Z 
double 
 
接合部・終端側オフセット（Z
－
） 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
構成する接合部部材の配置を、節点ID およびオフセットにて指定する。

## Page 161

5-63 
 
5.12. フーチング（複数）：StbFootings 
・概要 
説明 ：フーチング情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbFooting 
1 
制限なし 
フーチング情報

## Page 162

5-64 
 
5.12.1. フーチング：StbFooting 
・概要 
説明 ：フーチング情報 
親要素：StbFootings 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node 
integer 
○ 
節点ID 
 
rotate 
double 
 
回転角 
※(3), ※(1), ※(2) 
id_section 
integer 
○ 
断面ID 
 
offset_X 
double 
 
オフセット（X
－
） 
※(2), ※(1) 
offset_Y 
double 
 
オフセット（Y
－
） 
※(2), ※(1) 
level_bottom 
double 
 
レベル（下） 
※(2), 説明図参照 
thickness_add_start_X 
double 
 
ふかし厚さ（X 始） 
※(4) 
thickness_add_end_X 
double 
 
ふかし厚さ（X 終） 
※(4) 
thickness_add_start_Y 
double 
 
ふかし厚さ（Y 始） 
※(4) 
thickness_add_end_Y 
double 
 
ふかし厚さ（Y 終） 
※(4) 
thickness_add_top 
double 
 
ふかし厚さ（上） 
※(4) 
thickness_add_bottom 
double 
 
ふかし厚さ（下） 
※(4) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) フーチングの基準点は図心位置とする。 
(2) フーチングのオフセットは節点からフーチング基準点までの距離（全体座標系X
－
,Y
－
）とする。オフ
セットとレベルが省略された場合、フーチング下端の基準点は節点とする。 
(3) 回転角はX 方向を0 度とし、Z 方向時計回りを正とし、断面を回転した後にオフセットを考慮す
る。省略した時は0 度（回転なし）とみなす。なお、配置するRC 基礎断面が矩形・矩形テーパ
ー・八角形の場合は0 度以上180 度未満とする。

## Page 163

5-65 
 
(4) 省略された場合、当該属性がないものとする。 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
・例 
 
 
 
 
 
 
 
 
<StbFootings> 
<StbFooting id="32" name="1FA1" id_node="15" rotate="0.00" id_section="107" 
offset_X="0.00" offset_Y="0.00" level_bottom="-700.00"/> 
・・・ 
</StbFootings> 
 
X 
Y 
Z 
ふかし厚さ（下） 
ふかし厚さ（上） 
節点 
レベル（下） 
ふかし厚さ(Y 終) 
ふかし厚さ(X 終) 
ふかし厚さ(Y 始) 
ふかし厚さ(X 始) 
オフセット(Y
－
) 
節点 
図心 
回転角＋ 
Y 
X 
オフセット(X
－
)

## Page 164

5-66 
 
5.13. 布基礎（複数）：StbStripFootings 
・概要 
説明 ：布基礎情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbStripFooting 
1 
制限なし 
布基礎情報

## Page 165

5-67 
 
5.13.1. 布基礎：StbStripFooting 
・概要 
説明 ：布基礎情報 
親要素：StbStripFootings 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_start 
integer 
○ 
始端節点ID 
※(2) 
id_node_end 
integer 
○ 
終端節点ID 
※(2) 
id_section 
integer 
○ 
断面ID 
 
kind_structure 
string 
○ 
構造種別 
RC のみ 
level 
double 
 
レベル 
※(1), ※(2) 
offset 
double 
 
オフセット 
※(1)～※(4) 
length_ex_start 
double 
 
始点側余長 
※(5) 
length_ex_end 
double 
 
終点側余長 
※(5) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 布基礎は底板部分を表し、基礎梁とは別にオフセットとレベルを指定する。オフセットとレベル
が省略された場合、基礎底板の中心線は梁の始終端の基準点を結ぶ線とする。 
(2) 布基礎の部材座標系は梁と同様に進行方向をX とし、鉛直上向きをZ とする。 
(3) 布基礎の基準点は、基礎天端の幅中心位置とする。 
(4) 布基礎のオフセットは梁の始終端の基準点を結ぶ線に対して平行移動（始端から終端に向かって
左側を正）とする。 
(5) 余長は、梁始終端の基準点から外側に出る寸法を正として記述する。省略された場合、布基礎の始
終端位置は、梁始終端の基準点位置とする。

## Page 166

5-68 
 
 
 
 
 
・例 
 
 
 
 
 
 
 
 
 
 
始端節点 
レベル 
終端節点 
 
Y 
X 
Z 
 
 
 
中心  
オフセット  
 
レベル   
Y 
Z  
<StbStripFootings> 
<StbStripFooting id="32" name="1FA1G" id_node_start="15" id_node_end="18" 
id_section="107" kind_structure="RC" level="-700.00" offset="0.00" length_ex_start="-
750.00" length_ex_end="0.00"/> 
・・・ 
</StbStripFootings> 
梁基準点 
梁終端基準点 
梁始端基準点

## Page 167

5-69 
 
5.14. 杭基礎（複数）：StbPiles 
・概要 
説明 ：杭基礎情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbPile 
1 
制限なし 
杭基礎情報

## Page 168

5-70 
 
5.14.1. 杭基礎：StbPile 
・概要 
説明 ：杭基礎情報 
親要素：StbPiles 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node 
integer 
○ 
節点ID 
 
id_section 
integer 
○ 
断面ID 
 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値を
とる。 
RC、S、PC 
RC：場所打ち杭 
S：鋼管杭 
PC：既製コンクリート杭 
offset_X 
double 
 
オフセット（X
－
） 
※(1), ※(2) 
offset_Y 
double 
 
オフセット（Y
－
） 
※(1), ※(2) 
level_top 
double 
 
レベル（杭天） 
※(1), ※(2) 
length_all 
double 
 
杭全長 
※(3) 
length_head 
double 
 
杭頭（拡頭）長さ 
※(4) 
length_foot 
double 
 
杭脚長さ 
※(5) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 杭の基準点は図心とする。 
(2) 杭のオフセットは節点から図心までの距離（全体座標系X
－
,Y
－
）とする。オフセットとレベルが省略
された場合、杭頭の基準点は節点とする。 
(3) 構造種別がRC（場所打ち杭）の場合、記述は必須とする。鋼管杭と既製コンクリート杭は継ぎ杭
本数と継ぎ杭長さで定義する。

## Page 169

5-71 
 
(4) 杭断面形状（StbSecFigurePile_RC）が頂部拡大（StbSecPile_RC_ExtendedTop）、杭断面配筋
（StbSecBarArrangementPile_RC）が杭頭脚別（StbSecBarPile_RC_TopBottom）または杭頭軸
部杭脚（StbSecBarPile_RC_TopCenterBottom）の場合、記述は必須とする。 
(5) 杭
断
面
配
筋
（
StbSecBarArrangementPile_RC
）
が
杭
頭
軸
部
杭
脚
（StbSecBarPile_RC_TopCenterBottom）の場合、記述は必須とする。 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
X 
Y 
Z 
基準点 
（杭心） 
杭頭長さ 
杭全長 
レベル 
拡頭長さ 
節点 
節点 
基準点 
（杭心） 
杭脚長さ 
杭脚長さ

## Page 170

5-72 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
・例 
 
 
 
 
 
 
 
 
<StbPiles> 
<StbPile id="32" name="1PA1" id_node="15" id_section="107" kind_structure="RC" 
offset_X="0.00" offset_Y="0.00" level_top="-1100.00" length_all="10000.00"/> 
・・・ 
</StbPiles> 
X 
Y 
基準点 
（図心） 
節点 
オフセット（X
－
） 
オフセット（Y
－
）

## Page 171

5-73 
 
5.15. 基礎柱（複数）：StbFoundationColumns 
・概要 
説明 ：基礎柱情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbFoundationColumn 
1 
制限なし 
基礎柱情報

## Page 172

5-74 
 
5.15.1. 基礎柱：StbFoundationColumn 
・概要 
説明 ：基礎柱情報（根巻柱を含む） 
親要素：StbFoundationColumns 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
※(1) 
guid 
string 
 
GUID 
※(1) 
name 
string 
 
名称 
※(1) 
id_node 
integer 
○ 
節点ID 
Ｓ柱の柱脚節点ID 
rotate 
double 
 
回転角 
※(4) 
offset_Z 
double 
 
基礎柱・根巻柱の基
準点のオフセット（Z
－
） 
※(3) 
kind_structure 
string 
○ 
構造種別 
以下の値をとる。 
RC 
 
id_section_FD 
integer 
 
基礎柱断面ＩＤ 
※(5), ※(6) , ※(7) 
length_FD 
double 
 
基礎柱高さ 
※(8), ※(9) 
offset_FD_X 
double 
 
基礎柱オフセット（X
－
） 
※(2), ※(3) 
offset_FD_Y 
double 
 
基礎柱オフセット（Y
－
） 
※(2), ※(3) 
thickness_add_FD_start_X 
double 
 
基礎柱ふかし厚さ（X
始） 
※(9) 
thickness_add_FD_end_X 
double 
 
基礎柱ふかし厚さ（X
終） 
※(9) 
thickness_add_FD_start_Y 
double 
 
基礎柱ふかし厚さ（Y
始） 
※(9) 
thickness_add_FD_end_Y 
double 
 
基礎柱ふかし厚さ（Y
終） 
※(9) 
id_section_WR 
integer 
 
根巻柱断面ＩＤ 
※(5), ※(6) 
length_WR 
double 
 
根巻柱高さ 
※(8), ※(9)

## Page 173

5-75 
 
offset_WR_X 
double 
 
根巻柱オフセット（X
－
） 
※(2), ※(3) 
offset_WR_Y 
double 
 
根巻柱オフセット（Y
－
） 
※(2), ※(3) 
thickness_add_WR_start_X 
double 
 
根巻柱ふかし厚さ（X
始） 
※(9) 
thickness_add_WR_end_X 
double 
 
根巻柱ふかし厚さ（X
終） 
※(9) 
thickness_add_WR_start_Y 
double 
 
根巻柱ふかし厚さ（Y
始） 
※(9) 
thickness_add_WR_end_Y 
double 
 
根巻柱ふかし厚さ（Y
終） 
※(9) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) Ｓ柱下の基礎柱形状と根巻柱脚の根巻形状を定義する。 
(2) 基礎柱・根巻柱の基準点は図心位置とする。 
(3) オフセットはＳ柱の始端節点から基礎柱・根巻柱の基準点までの距離とする。オフセットが省略
された場合、基礎柱・根巻柱の基準点は柱脚節点とする。 
(4) 回転はX 方向を0 度とし、-Z
－
方向に見て反時計回りを正とする。断面を回転した後にオフセット
を考慮する。省略した時は0 度（回転なし）とみなす。 
(5) 断面ID の参照先はStbSecColumn_RC とする。id_section_FD とid_section_WR のうち、どちら
か一方は必ず記述する。 
(6) 配筋は、StbSecBarColumn_RC_RectSame （矩形柱 
柱頭柱脚同一配筋）、
StbSecBarColumn_RC_CircleSame（円形柱 柱頭柱脚同一配筋）のいずれかとする。 
(7) 基礎柱下に基礎が存在する場合、配筋はフーチング底面まであるものとする。 
(8) 基礎柱断面ＩＤ(id_section_FD)がある場合は基礎柱高さを、根巻柱断面ＩＤ(id_section_WR)があ
る場合は基礎柱高さを、必ず記述する。基礎柱高さ／根巻柱高さは、基礎柱・根巻柱の基準点から
下側／上側への寸法を正として記述する。 
(9) 省略された場合、当該属性がないものとする。

## Page 174

5-76 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
・例 
 
 
 
 
 
 
 
 
（-Z 方向に見た図） 
<StbFoundationColumns> 
<StbFoundationColumn id="32" name="1CA1_1FA1" id_node="15" rotate="0.00" 
offset_Z="-0.00" kind_structure="RC" id_section_FD="107" length_FD="1200.00" 
offset_FD_X="0.00" offset_FD_Y="0.00" id_section_WR="1107" length_WR="-500.00" 
offset_WR_X="0.00" offset_WR_Y="0.00"/> 
・・・ 
</StbFoundationColumns> 
Y 
X 
Z 
基礎柱高さ 
根巻き柱高さ 
節点（S 柱の始端節点) 
基礎柱・根巻柱の基準点 
節点 
基礎柱高さ 
根巻柱高さ 
オフセット 
基準点 
ふかし厚さ(Y 終) 
ふかし厚さ(X 終) 
ふかし厚さ(Y 始) 
ふかし厚さ(X 始) 
オフセット(Y
－
) 
節点 
図心 
回転角＋ 
Y 
X 
オフセット(X
－
)

## Page 175

5-77 
 
5.16. パラペット（複数）：StbParapets 
・概要 
説明 ：パラペット情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbParapet 
1 
制限なし 
パラペット情報

## Page 176

5-78 
 
5.16.1. パラペット：StbParapet 
・概要 
説明 ：パラペット情報 
親要素：StbParapets 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node_start 
integer 
○ 
始端節点ID 
 
id_node_end 
integer 
○ 
終端節点ID 
 
id_section 
integer 
○ 
断面ID 
 
kind_structure 
string 
○ 
構造種別 
RC のみ 
kind_layout 
string 
○ 
壁種別 
以下のいずれかの値をとる。 
ON_GIRDER（大梁上）、 
ON_BEAM（小梁上）、 
ON_SLAB（スラブ上） 
 
direction 
string 
 
アゴの方向を示す(R/L) 
R（右側） 
L（左側） 
※(3) 
offset 
double 
 
オフセット 
※(1), ※(2) 
level 
double 
 
レベル 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) パラペットのoffset の符号は、梁（壁）と同様とし、正側を左とする。 
(2) レベルは、パラペット下端の位置を示す。オフセットとレベルが省略された場合、パラペット下端
中央の基準点は節点とする。

## Page 177

5-79 
 
(3) パラペット断面形状（StbSecFigureParapet_RC）がL 型（StbSecParapet_RC_TypeL）、また
はT 型（StbSecParapet_RC_TypeT）の場合、アゴの方向の記述は必須とする。 
 
 
 
・例 
 
 
 
 
 
 
 
 
X 
Y
Z 
始端基準点 
終端基準点 
× 
節点 
Y 
Z
offset（正値）
level（正値） 
direction：R（逆向きがL） 
×
×
基準点 
<StbParapets> 
<StbParapet id="184" name="RWB2.b1" id_node_start="330" id_node_end="331" 
id_section="103" kind_structure="RC" kind_layout="ON_GIRDER" direction="L" 
level="0"/> 
・・・ 
</StbParapets>

## Page 178

5-80 
 
5.17. 開口配置情報（複数）：StbOpenArrangements 
・概要 
説明 ：開口配置情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbOpenArrangement 
1 
制限なし 
開口情報

## Page 179

5-81 
 
5.17.1. 開口配置情報：StbOpenArrangement 
・概要 
説明 ：開口配置情報 
親要素：StbOpenArrangements 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_section 
integer 
○ 
断面ID 
 
kind_member 
integer 
○ 
部材の種別 
以下のいずれか 
WALL（壁） 
SLAB（スラブ） 
 
id_member 
integer 
○ 
部材ID 
※(3) 
position_X 
double 
○ 
開口位置（X） 
１点目から開口始点
までのX 方向距離 
※(1), ※(2) 
position_Y 
double 
○ 
開口位置（Y） 
１点目から開口始点
までのY 方向距離 
※(1) , ※(2) 
rotate 
double 
○ 
回転角度（度） 
開口がX 軸となす 
角度 
※(1) , ※(2) 
 
・内容 
無し 
 
・子要素 
    無し 
 
・補足 
(1) 開口位置・開口寸法はスラブ・壁の第１基準点からの距離で、１点目から2 点目に向かう方向を
X 方向、それに直交する方向（基準点の並びが反時計まわりになる側から見て反時計回りに90 度）
をY 方向とする。 
(2) スラブ・壁の周辺節点にオフセットが定義されている場合、オフセット後の基準点位置で定義す
る。

## Page 180

5-82 
 
(3) このID の参照先は、kind が壁「WALL」の場合<StbWall>を、床「SLAB」の場合<StbSlab>と
する。 
 
 
 
 
・例 
 
 
 
 
 
 
 
<StbOpenArrangements> 
<StbOpenArrangement id="188" name="1WC1G_Opn01" id_section="96" 
position_X="2650.00" position_Y="800.00" length_X="2000.00" length_Y="1000.00" 
rotate="0.00"/> 
・・・ 
</StbOpenArrangements>

## Page 181

5-83 
 
5.18. 梁貫通孔配置情報（複数）：StbPenetrationArrangements 
・概要 
説明 ：梁貫通孔配置情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbPenetrationArrangement 
0 
制限なし 
梁貫通配置情報

## Page 182

5-84 
 
5.18.1. 梁貫通孔配置情報：StbPenetrationArrangement 
・概要 
説明 ：梁貫通配置情報 
親要素：StbPenetrationArrangements 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_section 
integer 
〇 
梁貫通孔仕様ID 
※(1) 
kind_member 
string 
 
部材の種別 
以下のいずれか 
GIRDER（大梁） 
BEAM（小梁） 
※(2) 
id_member 
integer 
〇 
部材ID 
※(3) 
floor_load 
double 
 
床荷重（kN/m2） 
梁貫通補強検討用 
B 
double 
 
床の負担幅 
 
DL 
double 
 
長期分布荷重（N/mm） 
 
N 
double 
 
設計用軸力（N） 
 
Lh 
double 
〇 
梁始端から貫通孔芯まで
の距離 
※(4) 
e 
double 
〇 
梁天端から貫通孔芯まで
の距離 
※(5) 
isEccentricity 
boolean 
〇 
偏芯判定 
OK（true）、NG（false） 
 
isPosition 
boolean 
〇 
位置判定 
OK（true）、NG（false） 
 
isPitch 
boolean 
〇 
ピッチ判定 
OK（true）、NG（false） 
 
isAxial 
boolean 
〇 
軸力判定 
OK（true）、NG（false） 
 
isOther 
boolean 
〇 
その他判定 
OK（true）、NG（false） 
 
comment 
string 
 
メーカーコメント

## Page 183

5-85 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) このID の参照先は、鉄骨梁貫通孔情報<StbSecPenetration_S>とする。 
(2) 省略された場合は、GIRDER とする 
(3) このID の参照先は、kind_beam が大梁「GIRDER」の場合<StbGirder>を、小梁「BEAM」の場
合<StbBeam>とする。 
(4) 傾斜梁の場合、貫通孔芯から材軸との垂線が梁天端と交わる交点と、始端側基点までの長さとす
る。 
(5) 傾斜梁の場合、貫通孔芯から材軸との垂線が梁天端と交わる交点までの長さとする。 
 
 
 
 
 
梁始端基準点 
梁終端基準点 
Lh 
e 
梁貫通孔 
× 
×

## Page 184

5-86 
 
5.19. 継手配置情報（複数）：StbJointArrangements 
・概要 
説明 ：継手配置情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointArrangement 
1 
制限なし 
継手配置情報

## Page 185

5-87 
 
5.19.1. 継手配置情報：StbJointArrangement 
・概要 
説明 ：継手配置情報  
親要素：StbJointArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_section 
integer 
○ 
継手ID 
※(1) 
kind_member 
integer 
○ 
部材の種別 
以下のいずれか 
COLUMN(柱) 
POST(間柱) 
GIRDER（大梁） 
BEAM（小梁） 
BRACE(ブレース) 
※(1)※(2) 
id_member 
integer 
〇 
部材ID 
※(2) 
starting_point 
string 
〇 
起点の位置 
以下のいずれか 
START(始端) 
END(終端) 
 
distance 
double 
〇 
距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) このID の参照先は、kind_member がCOLUMN・POST の場合StbJointColumn*とし、
kind_member がGIRDER・BEAM・BRACE の場合、StbJointBeam*とする。 
(2) このID の参照先は、kind_member がCOLUMN の場合StbColumn、POST の場合StbPost、
GIRDER の場合StbGirder、BEAM の場合StbBeam、Brace の場合StbBrace とする。

## Page 186

5-88 
 
5.20. コンクリート柱梁接合部配置情報（複数）：StbPanelZoneArrangements 
・概要 
説明 ：コンクリート柱梁接合部情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbPanelZoneArrangement 
1 
制限なし 
コンクリート柱梁
接合部配置情報

## Page 187

5-89 
 
5.20.1. コンクリート柱梁接合部：StbPanelZoneArrangement 
・概要 
説明 ：コンクリート柱梁接合部 
親要素：StbPanelZoneArrangements 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
 
名称 
 
id_node 
integer 
○ 
節点ID 
 
offset_X 
double 
 
オフセット(Xഥ) 
※(1) 
offset_Y 
double 
 
オフセット(Yഥ) 
※(1) 
offset_Z 
double 
 
オフセット(Zത) 
※(1) 
height 
double 
 
高さ 
 
rotate 
double 
 
回転角 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) パネルゾーンの図心・天端位置を示すものとする。 
(2) 回転はX 方向を0 度とし、Z 軸の負側から正側を見た場合の時計回りを正とし、断面を回転した後
にオフセットを考慮する。省略した時は0 度（回転なし）とみなす。 
 
 
 
× 
オフセット 
基準点 
節点 
height

## Page 188

5-90 
 
5.21. コネクション配置情報（複数）：StbConnectionArrangements 
概要 
説明 ：コネクション配置情報（複数） 
親要素：StbMembers 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectionArrangement 
1 
制限なし 
コネクション情報

## Page 189

5-91 
 
5.21.1. コネクション配置情報：StbConnectionArrangement 
概要 
説明 ：コネクション配置情報 
親要素：StbConnectionArrangements 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_node 
integer 
○ 
節点ID 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectedColumn 
0 
1 
主材となる柱 
※(2) 
StbConnectedPost 
0 
1 
主材となる間柱 
※(2) 
StbConnectedGirder 
0 
1 
主材となる大梁 
※(2) 
StbConnectedBeam 
0 
1 
主材となる小梁 
※(2) 
StbConnectedBrace 
0 
1 
主材となるブレース 
※(2) 
StbConnectionGussetPlate 
0 
制限なし 
 
 
StbConnectionDiaphragm 
0 
制限なし 
 
 
StbConnectionStiffner 
0 
制限なし 
 
 
 
・補足 
(1) 親となる部材の節点ID とする。 
(2) いずれかを必須とする

## Page 190

5-92 
 
5.21.1.1. 主材となる柱：StbConnectedColumn 
概要 
説明 ：コネクション配置情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 191

5-93 
 
5.21.1.2. 主材となる間柱：StbConnectedPost 
概要 
説明 ：コネクション情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 192

5-94 
 
5.21.1.3. 主材となる大梁：StbConnectedGirder 
概要 
説明 ：コネクション情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 193

5-95 
 
5.21.1.4. 主材となる小梁：StbConnectedBeam 
概要 
説明 ：コネクション配置 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 194

5-96 
 
5.21.1.5. 主材となるブレース：StbConnectedBrace 
概要 
説明 ：コネクション情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 195

5-97 
 
5.21.1.6. ガセットプレート配置情報：StbConnectionGussetPlate 
概要 
説明 ：ガセットプレート配置情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_gusset_plate 
integer 
○ 
ガセットプレート詳細ID 
 
side 
string 
 
払い込み方向 
front、back 
※(1) 
 
・内容 
無し 
 
子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectingPost 
0 
1 
接続する間柱 
※(2) 
StbConnectingGirder 
0 
1 
接続する大梁 
※(2) 
StbConnectingBeam 
0 
1 
接続する小梁 
※(2) 
StbConnectingBrace 
0 
1 
接続するブレース 
※(2) 
StbConnectionRibPlate 
0 
1 
控えのリブプレート 
 
 
・補足 
(1) ガセットプレートにシングルシアで接合する場合などの払い込み方向は、取りつく材の部材座標系
Y 方向プラス方向を表、Y 方向マイナス方向を裏とする。 
 
(2) いずれかを必須とする 
 
表 
表 
終点節点 
始点節点

## Page 196

5-98 
 
5.21.1.6.1. 接続する間柱：StbConnectingPost 
概要 
説明 ：接続する間柱情報 
親要素：StbConnectionGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
pos 
string 
 
接続する端部 
TOP、BOTTOM 
 
 
・内容 
無し 
 
子要素 
無し

## Page 197

5-99 
 
5.21.1.6.2. 接続する大梁：StbConnectingGirder 
概要 
説明 ：接続する大梁情報 
親要素：StbConnectionGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
pos 
string 
 
接続する端部 
START、END 
 
 
・内容 
無し 
 
子要素 
無し

## Page 198

5-100 
 
5.21.1.6.3. 接続する小梁：StbConnectingBeam 
概要 
説明 ：接続する小梁情報 
親要素：StbConnectionGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
pos 
string 
 
接続する端部 
START、END 
 
 
・内容 
無し 
 
子要素 
無し

## Page 199

5-101 
 
5.21.1.6.4. 接続するブレース：StbConnectingBrace 
概要 
説明 ：接続するブレース情報 
親要素：StbConnectionGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
pos 
string 
 
接続する端部 
START、END 
 
 
・内容 
無し 
 
子要素 
無し

## Page 200

5-102 
 
5.21.1.6.5. リブプレート配置情報：StbConnectionRibPlate 
概要 
説明 ：リブプレート配置情報 
親要素：StbConnectionGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_rib_plate 
integer 
○ 
リブプレート詳細ID 
 
mounting_angle 
string 
○ 
リブプレートの取り付け角度 
以下のいずれかの値をとる。 
Parallel（取り付き部材に平行）、 
Right-angled（母材と直行） 
※(1) 
offset 
double 
 
ガセットプレート芯に対するオ
フセット 
省略値は0 とする 
 
・内容 
無し 
 
子要素 
無し 
 
・補足 
(1) リブプレートの取り付け角度は、取り付く子材と平行または、母材の材軸と直行方向とする。

## Page 201

5-103 
 
5.21.1.7. ダイアフラム配置情報：StbConnectionDiaphragm 
概要 
説明 ：ダイアフラム配置情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_diaphragm 
integer 
○ 
ID 
 
offset 
double 
 
節点からのオフセット 
省略値は0 とする 
 
・内容 
無し 
 
子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectingGirder 
1 
制限なし 
接続する大梁 
 
 
・補足

## Page 202

5-104 
 
5.21.1.8. スチフナー配置情報：StbConnectionStiffner 
概要 
説明 ：ダイアフラム配置情報 
親要素：StbConnectionArrangement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_stiffner 
integer 
○ 
ID 
 
offset 
double 
 
節点からのオフセット 
省略値は0 とする 
 
・内容 
無し 
 
子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbConnectingGirder 
1 
1 
接続する大梁 
 
 
・補足 
無し

## Page 203

6.1-1 
 
6. 要素リファレンス 断面情報  
6.1. 断面情報：StbSections 
・概要 
説明 ：断面情報 
親要素：StbModel 
 
・属性、内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecColumn_RC 
0 
制限なし 
ＲＣ柱断面 
 
StbSecColumn_S 
0 
制限なし 
Ｓ柱断面 
 
StbSecColumn_SRC 
0 
制限なし 
ＳＲＣ柱断面 
 
StbSecColumn_CFT 
0 
制限なし 
ＣＦＴ柱断面 
 
StbSecBeam_RC 
0 
制限なし 
ＲＣ梁断面 
 
StbSecBeam_S 
0 
制限なし 
Ｓ梁断面 
 
StbSecBeam_SRC 
0 
制限なし 
ＳＲＣ梁断面 
 
StbSecBrace_S 
0 
制限なし 
Ｓブレース断面 
 
StbSecSlab_RC 
0 
制限なし 
ＲＣスラブ断面 
 
StbSecSlabDeck 
0 
制限なし 
デッキプレートスラブ断面 
 
StbSecSlabPrecast 
0 
制限なし 
既製スラブ断面 
 
StbSecSlabLoad 
0 
制限なし 
荷重用スラブ断面 
 
StbSecWall_RC 
0 
制限なし 
ＲＣ壁断面 
 
StbSecWallLoad 
0 
制限なし 
荷重用壁断面 
 
StbSecIsolatingDevice 
0 
制限なし 
免震装置断面 
 
StbSecDampingDevice 
0 
制限なし 
制振装置断面 
 
StbSecFoundation_RC 
0 
制限なし 
ＲＣ基礎断面 
 
StbSecPile_RC 
0 
制限なし 
ＲＣ杭断面 
 
StbSecPile_S 
0 
制限なし 
鋼管杭断面 
 
StbSecPilePrecast 
0 
制限なし 
既製コンクリート杭断面 
 
StbSecParapet_RC 
0 
制限なし 
ＲＣパラペット断面 
 
StbSecOpen_RC 
0 
制限なし 
ＲＣ開口断面 
 
StbSecPenetration_S 
0 
制限なし 
鉄骨梁貫通孔補強仕様 
 
StbSecPanelZone 
0 
制限なし 
柱梁接合部断面

## Page 204

6.1-2 
 
StbSecSteel 
0 
1 
鉄骨断面 
 
StbSecUndefined 
0 
制限なし 
構造種別に依存しない断面 
 
 
・補足 
上記のいずれか１種類の子要素を持つものとし、全子要素の最小回数が0 であってはならない。 
子要素の並びは、上表に示す順番としなければならない。

## Page 205

6.2-1 
 
6.2. ＲＣ柱断面：StbSecColumn_RC 
・概要 
説明 ：ＲＣ柱断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
〇 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※(2) 
kind_column 
string 
 
柱の種別 
以下のいずれか 
COLUMN（柱） 
POST（間柱） 
※(3) 
strength_concrete 
string 
 
コンクリート強度 
※(4) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureColumn_RC 
0 
1 
ＲＣ柱断面形状 
※(5) 
StbSecBarArrangementColumn_RC 
0 
1 
ＲＣ柱断面配筋 
※(6) 
 
・補足 
(1)  「断面名称」は、部材リスト（構造図の柱断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「C1」で「1C1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、COLUMN とする。 
(4)  省略された場合は、参照する<StbColumn>の「終端節点ID」id_node_top が所属する<StbStory>
のコンクリート強度を、この要素のコンクリート強度とする。参照した<StbStory>のコンクリー
ト強度が省略されていた場合は、共通情報の属性「建物全体のコンクリート強度」
strength_concrete をこの要素のコンクリート強度とする。

## Page 206

6.2-2 
 
(5)  子要素<StbSecFigureColumn_RC>の回数が0 となる場合は、構造計算プログラムが計算対象
としない形状で、断面性能を直接指定する場合を想定している。この場合は、他の子要素も指定し
てはならない。 
(6) 子要素<StbSecBarArrangementColumn_RC>の回数が0 となる場合は、鉄筋を扱わないプログ
ラムが一時的に作成する場合を想定しており、無筋であることを示すものではない。 
 
・例 
 
 
 
  <StbSecColumn_RC id="13" name="C1" floor="1" strength_concrete="FC24"> 
    <StbSecFigureColumn_RC> 
      <StbSecColumnRect width_X="650" width_Y="650"/> 
    </StbSecFigureColumn_RC> 
    <StbSecBarArrangementColumn_RC> 
        （略） 
    </StbSecBarArrangementColumn_RC> 
  </StbSecColumn_RC> 
 
  <StbSecColumn_RC id="15" name="C2" floor="1" kind_column="COLUMN"> 
    <StbSecFigureColumn_RC> 
      <StbSecColumnCircle D="450"/> 
    </StbSecFigureColumn_RC> 
    <StbSecBarArrangementColumn_RC> 
        （略） 
    </StbSecBarArrangementColumn_RC> 
  </StbSecColumn_RC>

## Page 207

6.2-3 
 
6.2.1. ＲＣ柱断面形状：StbSecFigureColumn_RC 
・概要 
説明 ：ＲＣ柱断面の形状 
親要素：StbSecColumn_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecColumnRect 
1 
1 
コンクリート柱断面形状・矩形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecColumnCircle 
1 
1 
コンクリート柱断面形状・円形 
 
 
・補足

## Page 208

6.2-4 
 
6.2.1.1. コンクリート柱断面形状・矩形：StbSecColumnRect 
・概要 
説明 ：コンクリート柱矩形断面の形状 
親要素：StbSecFigureColumn_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_X 
double 
○ 
Ｘ幅 
 
width_Y 
double 
○ 
Ｙ幅 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「Ｘ幅」および「Ｙ幅」の定義は下図による。 
 
 
・例 
 
 
 
X幅
Y幅
X
Y
    <StbSecColumn_RC id="13" name="C1" （略）> 
      <StbSecFigureColumn_RC> 
        <StbSecColumnRect width_X="750" width_Y="650"/> 
      </StbSecFigureColumn_RC> 
      <StbSecBarArrangementColumn_RC （略）> 
        （略） 
      </StbSecBarArrangementColumn_RC> 
    </StbSecColumn_RC>

## Page 209

6.2-5 
 
6.2.1.2. コンクリート柱断面形状・円形：StbSecColumnCircle 
・概要 
説明 ：コンクリート柱円形断面の形状 
親要素：StbSecFigureColumn_RC、StbSecFigureColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
double 
〇 
直径 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「直径」の定義は下図による。 
 
・例 
 
・ 
 
 
 
直径
X
Y
    <StbSecColumn_RC id="15" name="C2" （略）> 
      <StbSecFigureColumn_RC> 
        <StbSecColumnCircle D="450"/> 
      </StbSecFigureColumn_RC> 
      <StbSecBarArrangementColumn_RC （略）> 
        （略） 
      </StbSecBarArrangementColumn_RC> 
    </StbSecColumn_RC>

## Page 210

6.2-6 
 
6.2.2. ＲＣ柱断面配筋：StbSecBarArrangementColumn_RC 
・概要 
説明 ：ＲＣ柱断面の配筋 
親要素：StbSecColumn_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
  矩形の場合 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectSame 
1 
1 
コンクリート柱断面配筋・矩
形・同一 
 
StbSecBarColumnXReinforced 
0 
1 
コンクリート矩形柱 Ｘ形配
筋 
※(1) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectNotSame 
2 
2 
コンクリート柱断面配筋・矩
形・柱頭脚別 
 
StbSecBarColumnXReinforced 
0 
1 
コンクリート矩形柱 Ｘ形配
筋 
※(1) 
  円形の場合 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleSame 
1 
1 
コンクリート柱断面配筋・円
形・同一 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleNotSame 
2 
2 
コンクリート柱断面配筋・円
形・柱頭脚別 
 
 
・補足

## Page 211

6.2-7 
 
(1) 子要素<StbSecBarColumnXReinforced> は、X 形配筋を使用する場合に用いる。X 形配筋部を
この要素で追加し、平行配筋部は子要素<StbSecBarColumnRectSame>などで定義する。 
・例 
 
 
 
  <StbSecColumn_RC id="13" name="C1" （略）> 
    <StbSecFigureColumn_RC> 
      <StbSecColumnRect width_X="750" width_Y="650"/> 
    </StbSecFigureColumn_RC> 
    <StbSecBarArrangementColumn_RC> 
      <StbSecBarColumnRectSame> 
       <StbSecBarColumnRectSameSimple （略）/> 
      </StbSecBarColumnRectSame> 
    </StbSecBarArrangementColumn_RC> 
  </StbSecColumn_RC> 
 
  <StbSecColumn_RC id="15" name="C2" （略）> 
    <StbSecFigureColumn_RC> 
      <StbSecColumnCircle D="600"/> 
    </StbSecFigureColumn_RC> 
    <StbSecBarArrangementColumn_RC> 
      <StbSecBarColumnCircleSame> 
       <StbSecBarColumnCircleSameSimple （略）/> 
      </StbSecBarColumnCircleSame> 
    </StbSecBarArrangementColumn_RC> 
  </StbSecColumn_RC>

## Page 212

6.2-8 
 
6.2.2.1. コンクリート柱断面配筋・矩形・同一：StbSecBarColumnRectSame 
・概要 
説明 ：コンクリート柱矩形断面の配筋（柱頭・柱脚が同一配筋の場合） 
親要素：StbSecBarArrangementColumn_RC、StbSecBarArrangementColumn_SRC 
 
・属性 
 
属性名 
型 
必須 
説明 
補足 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
 
・内容、子要素 
無し 
 
・内容 
無し 
 
・子要素 
     
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectSameSimple 
1 
1 
コンクリート柱断面配筋・矩
形・同一・簡易 
※(1) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamRectSameComplex 
1 
1 
コンクリート柱断面配筋・矩
形・同一・詳細 
※(2) 
 
・補足 
(1) 対象は一段配筋までとする。 
(2) 方向毎に鉄筋径と配筋位置まで指定する。複数段指定することができる。

## Page 213

6.2-9 
 
・例 
 
 
 <StbSecColumn_RC id="13" name="C1" （略）> 
    <StbSecFigureColumn_RC> 
      <StbSecColumnRect width_X="750" width_Y="650"/> 
    </StbSecFigureColumn_RC> 
    <StbSecBarArrangementColumn_RC> 
      <StbSecBarColumnRectSame> 
       <StbSecBarColumnRectSameSimple D_main=”D25” N_X=”5” N_Y=”5” D_hoop=”D13” 
N_hoop_X=”2” N_hoop_Y=”2” pitch_hoop=”100”> 
       </StbSecBarColumnRectSameSimple> 
      </StbSecBarColumnRectSame> 
    </StbSecBarArrangementColumn_RC> 
  </StbSecColumn_RC>

## Page 214

6.2-10 
 
6.2.2.2. コンクリート柱断面配筋・矩形・同一・簡易：StbSecBarColumnRectSameSimple 
・概要 
説明 ：コンクリート柱矩形断面の配筋・簡易（柱頭・柱脚が同一配筋の場合） 
親要素：StbSecBarColumnRectSame 
 
・属性 
属性名 
型 
必須 
説明 
補足 
main_direction 
string 
 
主方向 
X、Y のいずれか 
※(1)※(2) 
depth_cover_start_X 
double 
 
かぶり厚さ（X 始） 
※(3) 
depth_cover_end_X 
double 
 
かぶり厚さ（X 終） 
※(3) 
depth_cover_start_Y 
double 
 
かぶり厚さ（Y 始） 
※(3) 
depth_cover_end_Y 
double 
 
かぶり厚さ（Y 終） 
※(3) 
interval 
double 
 
２段筋のあき 
※(3) 
center_start_X 
double 
 
主筋重心位置（X 始） 
※(3) 
center_start_Y 
double 
 
主筋重心位置（X 終） 
※(3) 
center_end_X 
double 
 
主筋重心位置（Y 始） 
※(3) 
center_end_Y 
double 
 
主筋重心位置（Y 終） 
※(3) 
center_interval 
double 
 
２段筋重心間距離 
※(3) 
D_main 
string 
〇 
主鉄筋径 
 
D_sub 
string 
 
副主鉄筋径 
※(4) 
strength_main 
string 
 
主鉄筋強度 
 
strength_sub 
string 
 
副主鉄筋強度 
※(4) 
N_X 
integer 
〇 
X 方向片側一段本数 
 
N_start_X 
integer  
 
X 始端側に寄せて配置する本数 
※(5) 
N_end_X 
integer  
 
X 終端側に寄せて配置する本数 
※(5) 
N_Y 
integer  
〇 
Y 方向片側一段本数 
 
N_start_Y 
integer 
 
Y 始端側に寄せて配置する本数 
※(5) 
N_end_Y 
integer 
 
Y 終端側に寄せて配置する本数 
※(5) 
hoop_type 
string 
 
帯筋種別 
NORMAL：普通配筋 
WELD：溶接閉鎖 
SPIRAL：スパイラル筋 
 
D_hoop 
string 
〇 
帯筋径 
 
strength_hoop 
string 
 
帯筋強度

## Page 215

6.2-11 
 
N_hoop_X 
integer 
〇 
X 方向帯筋本数 
 
N_hoop_Y 
integer 
〇 
Y 方向帯筋本数 
 
pitch_hoop 
double 
〇 
帯筋ピッチ 
 
D_axial 
string 
 
軸筋径 
 
strength_axial 
string 
 
軸筋強度 
 
N_axial 
integer 
 
軸筋本数 
 
 
・内容 
無し 
 
・子要素 
     
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋 
 
 
    
・補足 
(1) 記入無しの場合、X 方向とする。 
(2) コーナー筋は主方向の鉄筋とする。 
(3) 「かぶり厚さ（…）」および「２段筋のあき」の定義は下図による。 
 
 
「かぶり厚さ（…）」「主筋重心位置（…）」および「２段筋のあき」が省略された場合の扱いは、
<StbApplyConditionsList >の補足説明による。 
(4) 記入がない場合は主鉄筋径、主鉄筋強度とする 
(5) 寄筋で考慮する本数を記入する。 
かぶり厚さ（X始）
かぶり厚さ（X終）
かぶり厚さ（Y終）
かぶり厚さ（Y始）
X 
Y

## Page 216

6.2-12 
 
・例 
 
 
 <StbSecColumn_RC id="13" name="C1" （略）> 
    <StbSecFigureColumn_RC> 
      <StbSecColumnRect width_X="850" width_Y="550"/> 
    </StbSecFigureColumn_RC> 
    <StbSecBarArrangementColumn_RC> 
      <StbSecBarColumnRectSame> 
       <StbSecBarColumnRectSameSimple main_direction=”X” D_main=”D25” D_sub=”D22”  
N_X=”5” N_startX=”2” N_startY=”3” N_Y=”7” D_hoop=”D13” 
N_hoop_X=”3” N_hoop_Y=”4” pitch_hoop=”100”> 
       </StbSecBarColumnRectSameSimple> 
      </StbSecBarColumnRectSame> 
    </StbSecBarArrangementColumn_RC>

## Page 217

6.2-13 
 
6.2.2.2.1.1. コンクリート柱断面配筋・追加鉄筋：StbSecBarColumnAdditional 
・概要 
説明 ：コンクリート柱断面配筋・追加鉄筋 
親要素：StbSecBarColumnRectSameSimple, StbSecBarColumnRectSameComplex, 
StbSecBarColumnRectNotSameSimple, StbSecBarColumnRectNotSameComplex, 
StbSecBarColumnCircleSameSimple, StbSecBarColumnCircleSameComplex, 
StbSecBarColumnCircleNotSameSimple, StbSecBarColumnCircleNotSameComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
径 
 
strength 
string 
 
強度 
 
X 
double 
〇 
X 位置 
※(1) 
Y 
double 
〇 
Y 位置 
※(1) 
isStructual 
boolean 
 
構造計算に考慮するかどうか 
 
 
・内容 
無し 
 
・子要素 
    無し 
 
    
・補足 
(1) 図心を原点とし、鉄筋芯位置を定義する。

## Page 218

6.2-14 
 
6.2.2.2.2. コンクリート柱断面配筋・矩形・同一・詳細：StbSecBarColumnRectSameComplex 
・概要 
説明 ：コンクリート柱矩形断面配筋・詳細（柱頭・柱脚が同一配筋の場合） 
親要素：StbSecBarColumnRectSame 
 
・属性 
属性名 
型 
必須 
説明 
補足 
N_mainBar 
integer 
〇 
主筋総本数 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectComplexMain 
4 
4 
コンクリート柱断面配
筋・矩形・詳細・主筋 
※(1) 
StbSecBarColumnRectComplexHoop 
1 
1 
コンクリート柱断面配
筋・矩形・詳細・帯筋 
 
StbSecBarColumnRectComplexAxial 
0 
1 
コンクリート柱断面配
筋・矩形・詳細・軸筋 
 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配
筋・追加鉄筋 
 
    
 
・補足 
(1) X,Y 方向ともに計算に使用している鉄筋を定義する。同じ位置に鉄筋が配筋される場合は一つと
して検討し、主筋総本数と比較する。ST-Bridge 内でデータの整合性は担保できないので、アプ
リケーション側でチェックする。

## Page 219

6.2-15 
 
6.2.2.2.2.1. コンクリート柱断面配筋・矩形・詳細・主筋：StbSecBarColumnRectComplexMain 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・主筋 
親要素：StbSecBarColumnRectSameComplex, StbSecBarColumnRectNotSameComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
〇 
配筋位置 以下のいずれか 
STARTX：X 始 
ENDX：X 終 
STARTY：Y 始 
ENDY：Y 終 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectComplexMainLine 
1 
制限なし 
コンクリート柱断面配筋・矩
形・詳細・主筋列 
※(1) 
 
・補足 
(1) 鉄筋径が混在する場合は別要素とする。

## Page 220

6.2-16 
 
・例 
 
 
    <StbSecBarArrangementColumn_RC> 
      <StbSecBarColumnRectSame> 
       <StbSecBarColumnRectSameComplex N_mainBar=”12”> 
        <StbSecBarColumnRectComplexMain pos=”STARTX”> 
         <StbSecBarColumnRectComplexMainLine step=”1” D=”D25” distance=”70”> 
          <StbSecBarColumnRectComplexMainLoc distance=”75”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”170”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”330”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”425”/> 
         </StbSecBarColumnRectComplexMainLine >          
        </StbSecBarColumnRectComplexMain> 
        <StbSecBarColumnRectComplexMain pos=”ENDX”> 
         <StbSecBarColumnRectComplexMainLine step=”1” D=”D25” distance=”70”> 
          <StbSecBarColumnRectComplexMainLoc distance=”75”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”170”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”330”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”425”/> 
         </StbSecBarColumnRectComplexMainLine >          
        </StbSecBarColumnRectComplexMain> 
        <StbSecBarColumnRectComplexMain pos=”STARTY”> 
         <StbSecBarColumnRectComplexMainLine step=”1” D=”D25” distance=”70”> 
          <StbSecBarColumnRectComplexMainLoc distance=”75”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”170”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”330”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”425”/> 
         </StbSecBarColumnRectComplexMainLine >          
        </StbSecBarColumnRectComplexMain> 
        <StbSecBarColumnRectComplexMain pos=”ENDY”> 
         <StbSecBarColumnRectComplexMainLine step=”1” D=”D25” distance=”70”> 
          <StbSecBarColumnRectComplexMainLoc distance=”75”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”170”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”330”/> 
          <StbSecBarColumnRectComplexMainLoc distance=”425”/> 
         </StbSecBarColumnRectComplexMainLine >          
        </StbSecBarColumnRectComplexMain> 
        <StbSecBarColumnRectComplexHoop D=”D13” pitch=”100”>   
         <StbSecBarColumnRectComplexHoopLocX distance=”55/>   
         <StbSecBarColumnRectComplexHoopLocX distance=”445”/>   
         <StbSecBarColumnRectComplexHoopLocY distance=”55”/>   
         <StbSecBarColumnRectComplexHoopLocY distance=”445”/>          
        </StbSecBarColumnRectComplexHoop >          
  </StbSecBarColumnRectSameComplex> 
      </StbSecBarColumnRectSame> 
    </StbSecBarArrangementColumn_RC>

## Page 221

6.2-17 
 
6.2.2.2.2.1.1. コンクリート柱断面配筋・矩形・詳細・主筋列：StbSecBarColumnRectComplexMainLine 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・主筋列 
親要素：StbSecBarColumnRectComplexMain 
 
・属性 
属性名 
型 
必須 
説明 
補足 
step 
integer 
〇 
段数 
 
D 
string 
〇 
鉄筋径 
 
strength 
string 
 
鉄筋強度 
 
distance 
double 
〇 
コンクリート端から鉄筋列までの距離 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectComplexMainLoc 
1 
制限なし 
コンクリート柱断面配筋・矩
形・詳細・主筋位置 
※(1) 
 
・補足 
(1) 本数は要素数とする

## Page 222

6.2-18 
 
6.2.2.2.2.1.2. コンクリート柱断面配筋・矩形・詳細・主筋位置：StbSecBarColumnRectComplexMainLoc 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・主筋位置 
親要素：StbSecBarColumnRectComplexMainLine 
 
・属性 
属性名 
型 
必須 
説明 
補足 
distance 
double 
〇 
コンクリート端から鉄筋芯までの距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 223

6.2-19 
 
6.2.2.2.2.2. コンクリート柱断面配筋・矩形・詳細・帯筋：StbSecBarColumnRectComplexHoop 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・帯筋 
親要素：StbSecBarColumnRectSameComplex, StbSecBarColumnRectNotSameComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
鉄筋径 
 
strength 
string 
 
鉄筋強度 
 
pitch 
double 
   〇 
ピッチ 
 
hoop_type 
string 
 
帯筋種別 
NORMAL：普通配筋 
WELD：溶接閉鎖 
SPIRAL：スパイラル筋 
 
 
・内容 
無し 
 
・子要素 
 
 
 
 
 
 
・補足 
(1) 帯筋径はすべて同じとする。 
 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectComplexHoopLocX 
1 
制限なし 
コンクリート柱断面配筋・
矩形・詳細・X 方向帯筋位置 
 
StbSecBarColumnRectComplexHoopLocY 
1 
制限なし 
コンクリート柱断面配筋・
矩形・詳細・Y 方向帯筋位置

## Page 224

6.2-20 
 
6.2.2.2.2.2.1. コンクリート柱断面配筋・矩形・詳細・X 方向帯筋位置：
StbSecBarColumnRectComplexHoopLocX 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・X 方向帯筋位置 
親要素：StbSecBarColumnRectComplexHoop 
 
・属性 
属性名 
型 
必須 
説明 
補足 
distance 
double 
〇 
コンクリート端から鉄筋芯までの距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 225

6.2-21 
 
6.2.2.2.2.2.2. コンクリート柱断面配筋・矩形・詳細・Y 方向帯筋位置：
StbSecBarColumnRectComplexHoopLocY 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・Y 方向帯筋位置 
親要素：StbSecBarColumnRectComplexHoop 
 
・属性 
属性名 
型 
必須 
説明 
補足 
distance 
double 
〇 
コンクリート端から鉄筋芯までの距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 226

6.2-22 
 
6.2.2.2.2.3. コンクリート柱断面配筋・矩形・詳細・軸筋：StbSecBarColumnRectComplexAxial 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・軸筋 
親要素：StbSecBarColumnRectSameComplex, StbSecBarColumnRectNotSameComplex 
 
属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
径 
 
strength 
string 
 
強度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectComplexAxialLoc 
1 
制限なし 
コンクリート柱断面配筋・矩
形・詳細・軸筋位置 
 
 
    
・補足 
(1) 軸筋径はすべて同じとする。 
 
・例 
 
 
 <StbSecBarColumnRectComplexAxial D=”D25”> 
 <StbSecBarColumnRectComplexAxialLoc X=”-100” Y=”-100”/> 
<StbSecBarColumnRectComplexAxialLoc X=”-100” Y=”100”/> 
<StbSecBarColumnRectComplexAxialLoc X=”100” Y=”-100”/> 
<StbSecBarColumnRectComplexAxialLoc X=”100” Y=”100”/> 
</StbSecBarColumnRectComplexAxial>

## Page 227

6.2-23 
 
6.2.2.2.2.3.1. コンクリート柱断面配筋・矩形・詳細・軸筋位置：StbSecBarColumnRectComplexAxialLoc 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細・軸筋位置 
親要素：StbSecBarColumnRectComplexAxial 
 
属性 
属性名 
型 
必須 
説明 
補足 
X 
double 
〇 
X 位置 
※(1) 
Y 
double 
〇 
Y 位置 
※(1) 
 
・内容 
無し 
 
・子要素 
    無し 
    
・補足 
(1) 図心を原点とし、鉄筋芯位置を定義する。 
 
 
 
Y 
（50,100） 
X 
（－100,－50）

## Page 228

6.2-24 
 
6.2.2.3. コンクリート矩形柱 Ｘ形配筋：StbSecBarColumnXReinforced 
・概要 
説明 ：コンクリート矩形柱 X 形配筋 
親要素：StbSecBarArrangementColumn_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_main 
string 
 
主筋径 
 
N_main_X 
integer 
 
主筋：X 方向本数 
 
N_main_Y 
integer 
 
主筋：Y 方向本数 
 
N_main_total 
integer 
 
主筋：X 形配筋の総本数 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し 
 
・例（断面、形状、帯筋等は省略） 
 
 
 
 
X 
Y 
X 形配筋 
平行配筋 
<StbSecBarArrangementColumn_RC > 
<StbSecBarColumnRectSame (略)> 
</StbSecBarColumnRectSame> 
< StbSecBarColumnXReinforced  D_main=”D25” N_main_X="2"  N_main_total ="4"/ > 
< /StbSecBarArrangementColumn_RC >

## Page 229

6.2-25 
 
6.2.2.4. コンクリート柱断面配筋・矩形・柱頭脚別：StbSecBarColumnRectNotSame 
・概要 
説明 ：コンクリート柱矩形断面の配筋（柱頭・柱脚が別配筋の場合） 
親要素：StbSecBarArrangementColumn_RC, StbSecBarArrangementColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
 
・内容 
無し 
 
・子要素 
     
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectNotSameSimple 
2 
2 
コンクリート柱断面配筋・矩
形・柱頭脚別・簡易 
※(1) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamRectNotSameComplex 
2 
2 
コンクリート柱断面配筋・矩
形・柱頭脚別・詳細 
※(2) 
 
・補足 
(1) 柱脚、柱頭配筋について、属性をそれぞれpos= “BASE” およびpos= “TOP” としたこの子要素
を各１回記述する。 
(2) 対象は一段配筋までとする。 
(3) 方向毎に鉄筋径と配筋位置まで指定する。複数段指定することができる。

## Page 230

6.2-26

## Page 231

6.2-27 
 
6.2.2.4.1. コンクリート柱断面配筋・矩形・柱頭脚別・簡易：StbSecBarColumnRectNotSameSimple 
・概要 
説明 ：コンクリート柱断面の配筋・簡易（柱頭・柱脚が別配筋の場合） 
親要素：StbSecBarColumnRectNotSame 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
〇 
柱頭か柱脚か 以下のいずれか 
TOP：柱頭 
BOTTOM：柱脚 
 
main_direction 
string 
 
主方向 以下のいずれか 
X 
Y 
※(1)※(2) 
depth_cover_start_X 
double 
 
かぶり厚さ（X 始） 
※(3) 
depth_cover_end_X 
double 
 
かぶり厚さ（X 終） 
※(3) 
depth_cover_start_Y 
double 
 
かぶり厚さ（Y 始） 
※(3) 
depth_cover_end_Y 
double 
 
かぶり厚さ（Y 終） 
※(3) 
interval 
double 
 
２段筋のあき 
※(3) 
center_start_X 
double 
 
主筋重心位置（X 始） 
※(3) 
center_start_Y 
double 
 
主筋重心位置（X 終） 
※(3) 
center_end_X 
double 
 
主筋重心位置（Y 始） 
※(3) 
center_end_Y 
double 
 
主筋重心位置（Y 終） 
※(3) 
center_interval 
double 
 
２段筋重心間距離 
※(3) 
D_main 
string 
〇 
主方向鉄筋径 
 
D_sub 
string 
 
副方向鉄筋径 
※(4) 
strength_main 
string 
 
主方向鉄筋強度 
 
strength_sub 
string 
 
副方向鉄筋強度 
※(4) 
N_X 
integer 
〇 
X 方向片側一段本数 
 
N_start_X 
integer  
 
X 始端側に寄せて配置する本数 
※(5) 
N_end_X 
integer  
 
X 終端側に寄せて配置する本数 
※(5) 
N_Y 
integer  
〇 
Y 方向片側一段本数 
 
N_start_Y 
integer 
 
Y 始端側に寄せて配置する本数 
※(5) 
N_end_Y 
integer 
 
Y 終端側に寄せて配置する本数 
※(5) 
hoop_type 
string 
 
帯筋種別 
NORMAL：普通配筋

## Page 232

6.2-28 
 
WELD：溶接閉鎖 
SPIRAL：スパイラル筋 
D_hoop 
string 
〇 
帯筋径 
 
strength_hoop 
string 
 
帯筋強度 
 
N_hoop_X 
integer 
〇 
X 方向帯筋本数 
 
N_hoop_Y 
integer 
〇 
Y 方向帯筋本数 
 
pitch_hoop 
double 
〇 
帯筋ピッチ 
 
D_axial 
string 
 
軸筋径 
 
strength_axial 
string 
 
軸筋強度 
 
N_axial 
integer 
 
軸筋本数 
 
 
・内容 
無し 
 
・子要素 
     
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋 
 
 
    
・補足 
(1) 記入無しの場合、X 方向とする 
(2) コーナー筋は主方向の鉄筋とする 
(3) 「かぶり厚さ（…）」および「２段筋のあき」の定義は下図による。 
 
 
X
Y
かぶり厚さ（X始）
かぶり厚さ（X終）
かぶり厚さ（Y終）
かぶり厚さ（Y始）
2段筋のあき
2段筋のあき
かぶり厚さ

## Page 233

6.2-29 
 
「かぶり厚さ（…）」「主筋重心位置（…）」および「２段筋のあき」が省略された場合の扱いは、
<StbApplyConditionsList >の補足説明による。 
(4) 記入がない場合は主方向鉄筋径、主方向鉄筋強度とする。 
(5) 寄筋で考慮する本数を記入する。

## Page 234

6.2-30 
 
6.2.2.4.2. コンクリート柱断面配筋・矩形・詳細：StbSecBarColumnRectNotSameComplex 
・概要 
説明 ：コンクリート柱断面配筋・矩形・詳細（柱頭・柱脚が別配筋の場合） 
親要素：StbSecBarColumnRectNotSame 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
〇 
配筋位置 以下のいずれか 
BOTTOM：柱脚 
TOP：柱頭 
 
N_mainBar 
integer 
〇 
主筋総本数 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectComplexMain 
4 
4 
コンクリート柱断面配筋・矩
形・詳細・主筋 
※(1) 
StbSecBarColumnRectComplexHoop 
1 
1 
コンクリート柱断面配筋・矩
形・詳細・帯筋 
 
StbSecBarColumnRectComplexAxial 
0 
1 
コンクリート柱断面配筋・矩
形・詳細・軸筋 
 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋 
 
    
 
・補足 
(1) X,Y 方向ともに計算に使用している鉄筋を定義する。同じ位置に鉄筋が配筋される場合は一つと
して検討し、主筋総本数と比較する。ST-Bridge 内でデータの整合性は担保できないので、アプ
リケーション側でチェックする。

## Page 235

6.2-31 
 
6.2.2.5. コンクリート柱断面配筋・円形・同一：StbSecBarColumnCircleSame 
・概要 
説明 ：コンクリート柱断面配筋・円形・同一 
親要素：StbSecBarArrangementColumn_RC, StbSecBarArrangementColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
 
・内容 
無し 
 
・子要素 
     
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleSameSimple 
1 
1 
コンクリート柱断面配筋・円
形・同一・簡易 
※(1) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamCircleSameComplex 
1 
1 
コンクリート柱断面配筋・円
形・同一・詳細 
※(2) 
 
・補足 
(1) 対象は一段配筋までとする。 
(2) 方向毎に鉄筋径と配筋位置まで指定する。複数段指定することができる。 
 
・例

## Page 236

6.2-32 
 
6.2.2.5.1. コンクリート柱断面配筋・円形・同一・簡易：StbSecBarColumnCircleSameSimple 
・概要 
説明 ：コンクリート柱円形断面の配筋・簡易（柱頭・柱脚が同一配筋の場合） 
親要素：StbSecBarArrangementColumn_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover 
double 
 
かぶり厚さ 
 
center 
double 
 
主筋重心位置 
 
D_main 
string 
〇 
主筋：径 
 
D_hoop 
string 
〇 
帯筋：径 
 
D_axial 
string 
 
軸筋：径 
 
strength_main 
string 
 
主筋：鉄筋強度 
 
strength_hoop 
string 
 
帯筋：鉄筋強度 
 
strength_axial 
string 
 
軸筋：鉄筋強度 
 
N_main 
integer 
〇 
主筋：本数 
 
N_hoop_X 
integer 
〇 
帯筋：X 方向本数 
 
N_hoop_Y 
integer 
〇 
帯筋：Y 方向本数 
 
N_axial 
integer 
 
軸筋：本数 
 
pitch_hoop 
double 
〇 
帯筋：ピッチ 
 
hoop_type 
string 
 
帯筋種別 
NORMAL：普通配筋 
WELD：溶接閉鎖 
SPIRAL：スパイラル筋 
 
bar_start_angle 
double 
 
鉄筋配置する際の開始角度 
※(3) 
 
・内容 
無し 
 
・子要素 
  
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋

## Page 237

6.2-33 
 
 
・補足 
主筋、軸筋、帯筋および巾止筋の定義は下図による。下図の場合、帯筋本数は2 本と数える。
 
 
(1)  軸筋は、位置に関する情報を有しない。軸筋がある場合は、軸筋に関する属性をすべて記述す
る。 
(2)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
(3) （図による。） 
 
・例 
 
 
 
X
Y
軸筋
巾止筋
帯筋
主筋
X 
Y 
bar_start_angle 
  <StbSecBarColumnCircleSameSimple D_main=”D25” D_hoop=”D13” N_main=”12” 
N_hoopX=”4” N_HoopY=”2 pitch_hoop=”100” bar_start_angle=”15” />

## Page 238

6.2-34 
 
6.2.2.5.2. コンクリート柱断面配筋・円形・同一・詳細：StbSecBarColumnCircleSameComplex 
・概要 
説明 ：コンクリート柱円形断面の配筋・詳細（柱頭・柱脚が同一配筋の場合） 
親要素：StbSecBarColumnCircleSame 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleComplexMain 
1 
制限なし 
コンクリート柱断面配筋・円
形・詳細・主筋 
 
StbSecBarColumnCircleComplexHoop 
1 
1 
コンクリート柱断面配筋・円
形・詳細・帯筋 
 
StbSecBarColumnCircleComplexAxial 
0 
制限なし 
コンクリート柱断面配筋・円
形・詳細・軸筋 
 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋 
 
 
・補足 
    無し

## Page 239

6.2-35 
 
6.2.2.5.2.1. コンクリート柱断面配筋・円形・詳細・主筋：StbSecBarColumnCircleComplexMain 
・概要 
説明 ：コンクリート柱断面配筋・円形・詳細・主筋 
親要素：StbSecBarColumnCircleSameComplex、StbSecBarCircleCircleNotSameComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
鉄筋径 
 
strength 
string 
 
鉄筋強度 
 
step 
integer 
〇 
段数 
 
distance 
double 
〇 
コンクリート面から鉄筋列までの距離 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleComplexMainLoc 
1 
制限なし 
コンクリート柱断面配筋・円
形・詳細・主筋位置 
※(1) 
 
・補足 
(1) 本数は要素数とする 
 
・例 
 
 
<StbSecBarColumnCircleComplexMain D=”D25” step = “1” distance=”75”> 
<StbSecBarColumnCircleComplexMainLoc angle=”15”> 
<StbSecBarColumnCircleComplexMainLoc angle=”45”> 
<StbSecBarColumnCircleComplexMainLoc angle=”75”> 
<StbSecBarColumnCircleComplexMainLoc angle=”115”> 
<StbSecBarColumnCircleComplexMainLoc angle=”135”> 
<StbSecBarColumnCircleComplexMainLoc angle=”155”> 
<StbSecBarColumnCircleComplexMainLoc angle=”205”> 
<StbSecBarColumnCircleComplexMainLoc angle=”225”> 
<StbSecBarColumnCircleComplexMainLoc angle=”245”> 
<StbSecBarColumnCircleComplexMainLoc angle=”285”> 
<StbSecBarColumnCircleComplexMainLoc angle=”315”> 
<StbSecBarColumnCircleComplexMainLoc angle=”345”> 
</StbSecBarColumnCircleComplexMain>

## Page 240

6.2-36 
 
6.2.2.5.2.1.1. コンクリート柱断面配筋・円形・詳細・主筋位置：StbSecBarColumnCircleComplexMainLoc 
・概要 
説明 ：コンクリート柱断面配筋・円形・詳細・主筋位置 
親要素：StbSecBarColumnCircleComplexMain 
 
・属性 
属性名 
型 
必須 
説明 
補足 
angle 
double 
〇 
鉄筋配置角度 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 図による。 
 
 
angle=75

## Page 241

6.2-37 
 
6.2.2.5.2.2. コンクリート柱断面配筋・円形・詳細・帯筋：StbSecBarColumnCircleComplexHoop 
・概要 
説明 ：コンクリート柱断面配筋・円形・詳細・帯筋 
親要素：StbSecBarColumnCircleSameComplex、StbSecBarCircleCircleNotSameComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
鉄筋径 
 
strength 
string 
 
鉄筋強度 
 
pitch 
double 
〇 
ピッチ 
 
hoop_type 
string 
 
帯筋種別 
NORMAL：普通配筋 
WELD：溶接閉鎖 
SPIRAL：スパイラル筋 
 
distance 
double 
〇 
コンクリート面から芯までの距離 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleComplexHoopL
oc 
0 
制限なし 
コンクリート柱断面配筋・円
形・詳細・帯筋中子位置 
※(1) 
 
・補足 
(1) 外側の帯筋の他に中子筋を配置する場合に指定する。 
・例 
 
 
<StbSecBarColumnCircleComplexHoop D=”D13” pitch=”100” distance=”75”> 
  <StbSecBarColumnCircleComplexHoopLoc angle =”90” distance=”-100”/> 
</StbSecBarColumnCircleComplexHoop>

## Page 242

6.2-38 
 
6.2.2.5.2.2.1. コンクリート柱断面配筋・円形・詳細・帯筋中子位置：
StbSecBarColumnCircleComplexHoopLoc 
・概要 
説明 ：コンクリート柱断面配筋・円形・詳細・帯筋中子位置 
親要素：StbSecBarColumnCircleComplexHoop  
 
・属性 
属性名 
型 
必須 
説明 
補足 
angle 
double 
〇 
鉄筋配置角度 
※(1) 
distance 
double 
〇 
コンクリート面から芯までの距離 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 図心を中心にZ 軸について属性angle で指定した角度に回転した後の座標系において、Y 軸正方
向への距離を指定する。Y 軸負方向の位置にある場合は負の値を指定する。 
 
 
 
 
X 
Y 
angle=90°、distance=－100 
100

## Page 243

6.2-39 
 
6.2.2.5.2.3. コンクリート柱断面配筋・円形・詳細・軸筋：StbSecBarColumnCircleComplexAxial 
・概要 
説明 ：コンクリート柱断面配筋・円形・詳細・軸筋 
親要素：StbSecBarColumnCircleSameComplex、StbSecBarCircleCircleNotSameComplex 
 
属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
径 
 
strength 
string 
 
強度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleComplexAxialLoc 
1 
制限なし 
コンクリート柱断面配筋・円
形・詳細・軸筋位置 
 
 
    
・補足 
 
・例

## Page 244

6.2-40 
 
6.2.2.5.2.3.1. コンクリート柱断面配筋・円形・詳細・軸筋位置：StbSecBarColumnCircleComplexAxialLoc 
・概要 
説明 ：コンクリート柱断面配筋・円形・詳細・軸筋位置 
親要素：StbSecBarColumnCircleComplexAxial 
 
属性 
属性名 
型 
必須 
説明 
補足 
X 
double 
〇 
X 位置 
 
Y 
double 
〇 
Y 位置 
 
 
・内容 
無し 
 
・子要素 
    無し 
    
・補足 
(1) 図心を原点とし、鉄筋芯位置を定義する。 
 
・例

## Page 245

6.2-41 
 
6.2.2.6. StbSecBarColumnCircleNotSame：コンクリート柱断面配筋・円形・柱頭脚別 
・概要 
説明 ：コンクリート柱円形断面の配筋（柱頭・柱脚が別配筋の場合） 
親要素：StbSecBarArrangementColumn_RC, StbSecBarArrangementColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
 
・内容 
無し 
 
・子要素 
     
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleNotSameSimple 
2 
2 
コンクリート柱断面配筋・円
形・柱頭脚別・簡易 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleNotSameComplex 
2 
2 
コンクリート柱断面配筋・円
形・柱頭脚別・詳細 
 
 
・補足 
(1) 柱脚、柱頭配筋について、属性をそれぞれpos= “BOTTOM” およびpos= “TOP” としたこの子要
素を各１回記述する。 
(2) 対象は一段配筋までとする。 
(3) 方向毎に鉄筋径と配筋位置まで指定する。複数段指定することができる。

## Page 246

6.2-42 
 
6.2.2.6.1. コンクリート柱円形断面配筋・簡易：StbSecBarColumnCircleNotSameSimple 
・概要 
説明 ：コンクリート柱円形断面の配筋・簡易（柱頭・柱脚が別配筋の場合） 
親要素：StbSecBarColumnCircleNotSame 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
〇 
配筋位置 以下のいずれか 
BOTTOM：柱脚 
TOP：柱頭 
 
depth_cover 
double 
 
かぶり厚さ 
 
center 
double 
 
主筋重心位置 
 
D_main 
string 
〇 
主筋：径 
 
D_hoop 
string 
〇 
帯筋：径 
 
D_axial 
string 
 
軸筋：径 
 
strength_main 
string 
 
主筋：鉄筋強度 
 
strength_hoop 
string 
 
帯筋：鉄筋強度 
 
strength_axial 
string 
 
軸筋：鉄筋強度 
 
N_main 
integer 
〇 
主筋：本数 
 
N_hoop_X 
integer 
〇 
帯筋：X 方向本数 
 
N_hoop_Y 
integer 
〇 
帯筋：Y 方向本数 
 
N_axial 
integer 
 
軸筋：本数 
 
pitch_hoop 
double 
〇 
帯筋：ピッチ 
 
hoop_type 
string 
 
帯筋種別 
NORMAL：普通配筋 
WELD：溶接閉鎖 
SPIRAL：スパイラル筋 
 
bar_start_angle 
double 
 
鉄筋配置する際の開始角度 
※(3) 
 
・内容 
無し 
 
・子要素

## Page 247

6.2-43 
 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋 
 
 
・補足 
主筋、軸筋、帯筋および巾止筋の定義は下図による。下図の場合、帯筋本数は2 本と数える。
 
 
(1)  軸筋は、位置に関する情報を有しない。軸筋がある場合は、軸筋に関する属性をすべて記述す
る。 
(2)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
(3) 図による。 
 
 
 
 
 
X
Y
軸筋
巾止筋
帯筋
主筋
X 
Y 
bar_start_angle

## Page 248

6.2-44 
 
6.2.2.6.2. コンクリート柱円形断面配筋・詳細（柱頭・柱脚が別配筋の場合）：StbSecBarCircleCircleNotSameComplex 
・概要 
説明 ：コンクリート柱円形断面の配筋・詳細（柱頭・柱脚が別配筋の場合） 
親要素：StbSecBarColumnCircleNotSame 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
BOTTOM：柱脚 
TOP：柱頭 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleComplexMain 
1 
制限なし 
コンクリート柱断面配筋・円
形・詳細・主筋 
 
StbSecBarColumnCircleComplexBar 
1 
1 
コンクリート柱断面配筋・円
形・詳細・帯筋 
 
StbSecBarColumnCircleComplexAxial 
0 
制限なし 
コンクリート柱断面配筋・円
形・詳細・軸筋 
 
StbSecBarColumnAdditional 
0 
制限なし 
コンクリート柱断面配筋・追加
鉄筋 
 
 
・補足 
無し

## Page 249

6.3-1 
 
6.3. Ｓ柱断面：StbSecColumn_S 
・概要 
説明 ：Ｓ柱断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※
(2) 
kind_column 
string 
 
柱の種別 
以下のいずれか 
COLUMN（柱） 
POST（間柱） 
※(3) 
isReferenceDirection 
boolean 
 
鉄骨向き 
※(4) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelFigureColumn_S 
1 
1 
Ｓ柱断面鉄骨形状 
 
 
   上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseProduct 
0 
1 
鉄骨断面柱脚製品 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseConventional 
0 
1 
鉄骨断面柱脚在来工法 
 
 
・補足 
鉄骨断面の形状は、鉄骨断面要素 <StbSecSteel> で用意した断面形状を、この要素の子要素で参照す
ることで定義する。

## Page 250

6.3-2 
 
(1)  「断面名称」は、部材リスト（構造図の柱断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「C1」で「1C1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、COLUMN とする。 
(4)  true の場合は、<StbSecSteel> で用意した断面形状の「鉄骨断面の基準方向」を、部材の断面
軸（Ｙ軸）とする。false の場合は、90 度傾けた断面形状を部材の断面軸とする。Ｈ型鋼において、
同一断面を部材リストに90 度傾けて表記する場合などを想定しており、true の場合は下図のよ
うになる。省略された場合は、true とする。なお、この属性はver.1.X の「direction」属性に対
応するが、ver.1.X の定義とはtrue とfalse が逆となっているので、注意を要する。 
 
・例 
 
 
 
  <StbSecColumn_S id="1" name="C1" floor="1" kind_column="COLUMN"> 
    <StbSecSteelFigureColumn_S> 
      <StbSecSteelColumn_S_NotSame pos="BOTTOM" 
 shape="□-400x400x25x25" strength_main="BCP325"/> 
      <StbSecSteelColumn_S_NotSame pos="TOP" 
shape="□-400x400x22x22" strength_main="BCP325"/> 
    </StbSecSteelFigureColumn_S> 
  </StbSecColumn_S> 
 
  <StbSecColumn_S id="19" name="P1" floor="1" kind_column="POST" 
 isReferenceDirection ="false"> 
    <StbSecSteelFigureColumn_S> 
      <StbSecSteelColumn_S_Same shape="H-100x100x6x8x8" strength_main="SN400"/> 
    </StbSecSteelFigureColumn_S> 
  </StbSecColumn_S>

## Page 251

6.3-3 
 
6.3.1. Ｓ柱断面鉄骨形状：StbSecSteelFigureColumn_S 
・概要 
説明 ：Ｓ柱断面の形状 
親要素：StbSecColumn_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
base_type 
string 
 
柱脚形式 以下のいずれか 
NONE（鉄骨柱脚なし） 
EXPOSE（露出） 
EMBEDDED（埋込） 
WRAP（根巻） 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_S_Same 
1 
1 
Ｓ柱断面鉄骨形状・同一 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_S_NotSame 
2 
2 
Ｓ柱断面鉄骨形状・柱頭脚別 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_S_ThreeTypes 
3 
3 
Ｓ柱断面鉄骨形状・３種類 
※(2)

## Page 252

6.3-4 
 
・補足 
(1)  省略された場合は、NONE とする。 
(2)  ハンチ状の変断面状態は想定しない。下図のように、断面数に応じて、記述する子要素を選択す
る。 
 
 
 
 
 
柱頭
柱脚
中央
柱頭
柱脚
子要素：
同一
柱頭脚別
３種類

## Page 253

6.3-5 
 
6.3.1.1. Ｓ柱断面鉄骨形状・同一：StbSecSteelColumn_S_Same 
・概要 
説明 ：Ｓ柱鉄骨断面の形状（全断面同一の場合） 
親要素：StbSecSteelFigureColumn_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(2)  省略された場合は、「鉄骨強度（主）」と同一とする。 
 
・例 
 
 
 
  <StbSecColumn_S id="19" name="P1" （略）> 
    <StbSecSteelFigureColumn_S> 
      <StbSecSteelColumn_S_Same shape="H-100x100x6x8x8" strength_main="SN400"/> 
    </StbSecSteelFigureColumn_S> 
  </StbSecColumn_S>

## Page 254

6.3-6 
 
6.3.1.2. Ｓ柱断面鉄骨形状・柱頭脚別：StbSecSteelColumn_S_NotSame 
・概要 
説明 ：Ｓ柱鉄骨断面の形状（柱頭・柱脚が別形状の場合） 
親要素：StbSecSteelFigureColumn_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 
以下のいずれか 
BOTTOM（柱脚） 
TOP（柱頭） 
 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
属性をそれぞれpos= “BOTTOM” およびpos= “TOP” とした子要素を各１回記述する。 
定義と補足内容 (1)～(2) は、「Ｓ柱断面鉄骨形状・同一」による。 
 
・例 
 
 
 
  <StbSecColumn_S id="1" name="C1" （略）> 
    <StbSecSteelFigureColumn_S> 
      <StbSecSteelColumn_S_NotSame pos="BOTTOM" 
 shape="□-400x400x25x25" strength_main="BCP325"/> 
      <StbSecSteelColumn_S_NotSame pos="TOP" 
shape="□-400x400x22x22" strength_main="BCP325"/> 
    </StbSecSteelFigureColumn_S> 
  </StbSecColumn_S>

## Page 255

6.3-7 
 
6.3.1.3. Ｓ柱断面鉄骨形状・３種類：StbSecSteelColumn_S_ThreeTypes 
・概要 
説明 ：Ｓ柱鉄骨断面の形状（柱頭・中央・柱脚が別形状の場合） 
親要素：StbSecSteelFigureColumn_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 
以下のいずれか 
BOTTOM（柱脚） 
CENTER（中央） 
TOP（柱頭） 
 
shape 
string 
○ 
鉄骨形状 
 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
属性をpos= “BOTTOM” , pos= “CENTER” およびpos= “TOP” とした子要素を各１回記述する。 
定義と補足内容 (1)～(2) は、「Ｓ柱断面鉄骨形状・同一」による。 
 
・例 
 
 
  <StbSecColumn_S id="2" name="C2" （略）> 
    <StbSecSteelFigureColumn_S> 
      <StbSecSteelColumn_S_ThreeTypes pos="BOTTOM" 
shape="BH-200x200x9x12" strength_main="SN400"/> 
      <StbSecSteelColumn_S_ThreeTypes pos="CENTER" 
shape="H-200x200x8x12" strength_main="SN400"/> 
      <StbSecSteelColumn_S_ThreeTypes pos="TOP" 
shape="BH-200x200x9x12" strength_main="SN400"/>  
    </StbSecSteelFigureColumn_S> 
  </StbSecColumn_S>

## Page 256

6.3-8 
 
6.3.2. 鉄骨断面柱脚製品：StbSecBaseProduct 
・概要 
説明 ：鉄骨の柱脚部分（既製品） 
親要素：StbSecColumn_S、StbSecColumn_SRC、StbSecColumn_CFT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
〇 
リリース時期 
※(1) 
direction_type 
integer 
 
偏心タイプの場合、ベースプレートの向
きで、以下のいずれか（度） 
0、90、180、270 
※(2) 
height_mortar 
double 
○ 
モルタル高さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
記述できるのは、鉄骨形状の属性「柱脚形式」が、EXPOSE（露出）のときのみとする。 
 
(1)  製品型番一覧表の柱脚カテゴリから「製品型番」と「リリース時期」の2 つのキーで製品を一意
に指定する 
(2)  偏心タイプ（柱鉄骨の図心とベースプレートの図心が一致しない形式）の場合、ベースプレート
の偏心する向きを指定する。向きに応じた配置は、1 方向偏心タイプ、2 方向偏心タイプについて
それぞれ次ページとする（角度は、- Z 方向に見て反時計回りで定義）。 
 省略された場合は、0 度とする。

## Page 257

6.3-9 
 
 
       0 度                             90 度 
 
 
 
270
度                                                
 
 
 
180 度 
 
 
 
 
 
 
１方向偏心の場合 
 
      0 度                             90 度 
 
 
 
270
度                                                
                                  
 
 
180 度 
 
 
 
 
 
 
２方向偏心の場合

## Page 258

6.3-10 
 
 
6.3.3. 鉄骨断面柱脚在来工法：StbSecBaseConventional 
・概要 
説明 ：鉄骨の柱脚部分（在来工法） 
親要素：StbSecColumn_S、StbSecColumn_SRC、StbSecColumn_CFT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
height_mortar 
double 
○ 
モルタル高さ 
 
cut_wide 
double 
 
切込み幅 
 
cut_height 
double 
 
切込み高さ 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseConventionalPlate 
1 
1 
ベースプレート 
 
StbSecBaseConventionalAnchorBolts 
1 
1 
アンカーボルト 
 
StbSecBaseConventionalRibPlates 
0 
制限なし 
リブプレート 
 
 
・補足 
切込み寸法 
 
 
cut_height 
cut_wide

## Page 259

6.3-11 
 
6.3.3.1. 鉄骨断面柱脚在来工法・ベースプレート：StbSecBaseConventionalPlate 
・概要 
説明 ：鉄骨柱脚（在来工法）のベースプレート部分 
親要素：StbSecBaseConventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
B_X 
double 
○ 
ベースプレートの寸法(Bx) 
 
B_Y 
double 
○ 
ベースプレートの寸法(By) 
 
C1_X 
double 
 
面取りX 幅(1) 
※(1) 
C1_Y 
double 
 
面取りY 幅(1) 
※(1) 
C2_X 
double 
 
面取りX 幅(2) 
※(1) 
C2_Y 
double 
 
面取りY 幅(2) 
※(1) 
C3_X 
double 
 
面取りX 幅(3) 
※(1) 
C3_Y 
double 
 
面取りY 幅(3) 
※(1) 
C4_X 
double 
 
面取りX 幅(4) 
※(1) 
C4_Y 
double 
 
面取りY 幅(4) 
※(1) 
t 
double 
○ 
ベースプレートの板厚 
 
strength 
string 
○ 
ベースプレートの鉄骨強度 
 
D_bolthole 
double 
○ 
アンカーボルトの孔径 
※(2) 
offset_X 
double 
 
オフセット(X) 
※(1) 
offset_Y 
double 
 
オフセット(Y) 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  省略された場合は、0mm とする。 
(2)  ボルト径に応じた適切な数値とする。

## Page 260

6.3-12 
 
 
ベースプレートの寸法と面取り幅 
 
 
 
 
 
 
 
 
 
 
 
 
 
オフセット・・節点位置よりベースプレート中央（Bx とBy のそれぞれ1/2 の位置＝アンカーボ
ルトの基準位置）までの距離 
 
 
 
 
ｙ
Ｂｘ
Ｂ
Ｙ
Ｘ
ベースプレート中央 
節点 
offset_y 
offset_x 
Bx 
By 
X 
Y

## Page 261

6.3-13 
 
6.3.3.2. 鉄骨断面柱脚在来工法・アンカーボルト：StbSecBaseConventionalAnchorBolts 
・概要 
説明 ：鉄骨柱脚（在来工法）のアンカーボルト部分 
親要素：StbSecBaseConventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
kind_bolt 
string 
○ 
アンカーボルト種別で以下のいずれか 
STD（建方用アンカーボルト） 
ABR（JIS B 1220 ABR アンカーボルト
[転造ねじ]） 
ABM（JIS B 1221 ABM アンカーボルト
[切削ねじ]） 
 
name_bolt 
string 
○ 
アンカーボルト径（ねじの呼びd） 
例：M30 
L 
double 
○ 
アンカーボルト定着長 
※(1) 
strength_bolt 
string 
○ 
アンカーボルト強度 
 
type_bolt 
string 
〇 
アンカーボルト形状で以下のいずれか 
I：I 型 
J：J 型 
L：L 型 
LHOOK：L 型フック 
HOLEIN：ホールインアンカー 
※(2) 
R1 
double 
 
アンカーボルト曲げ半径 
アンカーボルト形
状がJ 型、L 型、L
型フックの場合必
須 
R2 
double 
 
アンカーボルト曲げ半径2 
アンカーボルト形
状がL 型フックの
場合必須 
Lt 
double 
 
アンカーボルト全長 
重量長 
S1 
double 
 
アンカーボルト余長（上） 
 
S2 
double 
 
アンカーボルト余長（下） 
 
L1 
double 
 
埋め込み長さ1 
 
L2 
double 
 
埋め込み長さ2 
 
type_flame 
string 
 
アンカーフレーム形状 
例：FB-9x90 
(SS400)

## Page 262

6.3-14 
 
 
・内容 
無し 
 
・子要素 
 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseConventionalAnchorBolt 
1 
制限なし 
アンカーボルト詳細 
 
 
・補足 
(1)  アンカーボルト定着長・・躯体コンクリート天端から定着金物までの長さとする。（下図Ｌ） 
 
(2)  アンカーボルト形状・・左よりI 型、J 型、L 型、L 型フック、ホールインアンカー 
 
 
 
 
 
 
 
 
 
 
 
S2 
L 
S1 
L1 
L 
S1 
R1 
L1 
L 
S1 
R1 
L1
L
S1 
R1 
L2 
R2 
L 
S1

## Page 263

6.3-15 
 
6.3.3.2.1. 鉄骨断面柱脚在来工法・アンカーボルト詳細：StbSecBaseConventionalAnchorBolt 
・概要 
説明 ：鉄骨柱脚（在来工法）のアンカーボルト詳細 
親要素：StbSecBaseConventionalAnchorBolts 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
〇 
ボルトの番号 
 
offset_X 
double 
 
オフセット(X) 
※(1) 
offset_Y 
double 
 
オフセット(Y) 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) オフセット・・ベースプレート中央からボルト心までのオフセット距離。 
 
 
 
offset_X 
Offset_Y 
Ｙ 
Ｘ

## Page 264

6.3-16 
 
6.3.3.3. 鉄骨断面柱脚在来工法・リブプレート：StbSecBaseConventionalRibPlates 
・概要 
説明 ：鉄骨柱脚（在来工法）のリブプレート部分 
親要素：StbSecBaseConventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
A1 
double 
○ 
リブプレートの長さ 
 
A2 
double 
 
リブプレートの長さ 
※(1) 
B1 
double 
○ 
リブプレートの高さ 
 
B2 
double 
 
リブプレートの高さ 
※(1) 
t 
double 
○ 
リブプレートの厚さ 
 
strength 
string 
○ 
リブプレート鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseConventionalRibPlate 
1 
制限なし 
リブプレート 
 
 
・補足 
(1)  省略された場合は、0mm とする。  リブプレート寸法（スカラップは対象外）

## Page 265

6.3-17 
 
6.3.3.3.1. 鉄骨断面柱脚在来工法・リブプレート：StbSecBaseConventionalRibPlate 
・概要 
説明 ：鉄骨柱脚（在来工法）のリブプレート詳細 
親要素：StbSecBaseConventionalRibPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
〇 
リブプレートの番号 
 
offset_X 
double 
 
オフセット(X) 
※(1) 
offset_Y 
double 
 
オフセット(Y) 
※(1) 
angle 
double 
 
角度 
※(2) 
 
・子要素 
無し 
 
・補足 
(1) オフセット・・ベースプレート中央からリブプレート心までのオフセット距離。 
 
 
 
 
 
 
 
 
 
 
 
(2) 角度・・リブプレート挿入基点における回転角 
上記例では180 度 
offset_X 
offset_Y 
X 
Y 
ベースプレート中央 
リブプレート挿入起点

## Page 266

6.4-1 
 
6.4. ＳＲＣ柱断面：StbSecColumn_SRC 
・概要 
説明 ：ＳＲＣ柱断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※(2) 
kind_column 
string 
 
柱の種別 以下のいずれか 
COLUMN（柱） 
POST（間柱） 
※(3) 
strength_concrete 
string 
 
コンクリート強度 
※(4) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureColumn_SRC 
0 
1 
ＳＲＣ柱断面形状 
※(6) 
StbSecBarArrangementColumn_SRC 
0 
1 
ＳＲＣ柱断面配筋 
※(5) 
StbSecSteelFigureColumn_SRC 
0 
1 
ＳＲＣ柱断面鉄骨形状 
※(5) 
 
   上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseProduct 
0 
1 
鉄骨断面柱脚製品 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseConventional 
0 
1 
鉄骨断面柱脚在来工法

## Page 267

6.4-2 
 
・補足 
(1)  「断面名称」は、部材リスト（構造図の柱断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「C1」で「1C1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、COLUMN とする。 
(4)  省略された場合は、参照する<StbColumn>の「終端節点ID」id_node_top が所属する<StbStory>
のコンクリート強度をこの要素のコンクリート強度とする。「参照した<StbStory>のコンクリー
ト強度が省略されていた場合は、共通情報の属性「建物全体のコンクリート強度」
strength_concrete をこの要素のコンクリート強度とする。 
(5)  子要素<StbSecBarArrangementColumn_SRC>および<StbSecSteelFigureColumn_SRC>の
回数が0 となる場合は、鉄筋や内部鉄骨を扱わないプログラムが一時的に作成する場合を想定し
ており、無筋や鉄骨無配置であることを示すものではない。 
(6)  子要素<StbSecFigureColumn_SRC>の回数が0 となる場合は、構造計算プログラムが計算対象
としない形状で、断面性能を直接指定する場合を想定している。この場合は、他の子要素も指定し
てはならない。 
 
・例 
 
 
  <StbSecColumn_SRC id="23" name="C11" floor="1" strength_concrete="FC24"> 
    <StbSecFigureColumn_SRC> 
      <StbSecColumnRect width_X="900" width_Y="900"/> 
    </StbSecFigureColumn_SRC> 
    <StbSecBarArrangementColumn_SRC> 
      <StbSecBarColumnRectSame  （略）/> 
    </StbSecBarArrangementColumn_SRC> 
    <StbSecSteelFigureColumn_SRC> 
      <StbSecSteelColumn_SRC_Same> 
        <StbSecSteelColumn_SRC_ShapeCross1 
shape_X="BH-600x200x12x25" shape_Y="BH-600x200x12x25" 
strength_main_X="SN490" strength_main_Y="SN490" 
offset_XX="0" offset_XY="0" offset_YX="0" offset_YY="0"/> 
</StbSecSteelColumn_SRC_Same> 
    </StbSecSteelFigureColumn_SRC> 
  </StbSecColumn_SRC>

## Page 268

6.4-3 
 
6.4.1. ＳＲＣ柱断面形状：StbSecFigureColumn_SRC 
・概要 
説明 ：ＳＲＣ柱断面の形状 
親要素：StbSecColumn_SRC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecColumnRect 
1 
1 
コンクリート柱断面形状・矩形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecColumnCircle 
1 
1 
コンクリート柱断面形状・円形 
 
 
・補足 
無し

## Page 269

6.4-4 
 
6.4.2. ＳＲＣ柱断面配筋：StbSecBarArrangementColumn_SRC 
・概要 
説明 ：ＳＲＣ柱断面の配筋 
親要素：StbSecColumn_SRC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
  矩形の場合 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectSame 
1 
1 
コンクリート柱断面配筋矩形・
同一 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnRectNotSame 
2 
2 
コンクリート柱断面配筋矩形 
・柱頭脚別 
 
 
  円形の場合 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleSame 
1 
1 
コンクリート柱断面配筋円形・
同一 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarColumnCircleNotSame 
2 
2 
コンクリート柱断面配筋円形 
・柱頭脚別 
 
 
・補足 
無し

## Page 270

6.4-5 
 
6.4.3. ＳＲＣ柱断面鉄骨形状：StbSecSteelFigureColumn_SRC 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状 
親要素：StbSecColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
base_type 
string 
 
柱脚形式 以下のいずれか 
NONE（鉄骨柱脚なし） 
UNEMBEDDED（非埋込） 
UNEMBEDDED2（非埋込） 
EMBEDDED（埋込） 
※(1) 、※(2)、※(4) 
length_embedded 
double 
 
柱脚埋め込み長さ 
※(3) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_Same 
1 
1 
ＳＲＣ柱断面鉄骨形状・同一 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_NotSame 
1 
2 
ＳＲＣ柱断面鉄骨形状 
・柱頭脚別 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ThreeTypes 
2 
3 
ＳＲＣ柱断面鉄骨形状・３種類 
※(2)

## Page 271

6.4-6 
 
・補足 
(1)  省略された場合は、NONE とする。 
(2)  ハンチ状の変断面状態は想定しない。下図のように、断面数に応じて、記述する子要素を選択
する。柱脚形式が非埋込で、柱脚をRC 柱断面としたい場合は、「柱頭脚別」の場合最小回数
１、および「3 種類」の場合最小回数2 として、pos = ‘BOTTOM’（柱脚）となる子要素を省略
してもよい。 
 
 
 
(3)  EMBEDDED とした場合の柱脚埋込長さを指定する。柱脚埋め込み長さは、始端節点からの長さと
し下向きを正とする。 
(4)  非埋込UNEMBEDDED2 は、外形上はUNEMBEDDED と同じであり、構造計算上、柱脚をRC 柱断面
としたい場合を示す。 
 
 
 
子要素：
同一
柱頭脚別
柱頭
柱脚
3種類
中央
柱脚
柱頭

## Page 272

6.4-7 
 
6.4.3.1. ＳＲＣ柱断面鉄骨形状・同一：StbSecSteelColumn_SRC_Same 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状（全断面同一の場合） 
親要素：StbSecSteelFigureColumn_SRC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeH 
1 
1 
ＳＲＣ柱断面鉄骨形状・Ｈ形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeBox 
1 
1 
ＳＲＣ柱断面鉄骨形状・□形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapePipe 
1 
1 
ＳＲＣ柱断面鉄骨形状・○形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeCross1 
1 
1 
ＳＲＣ柱断面鉄骨形状・＋形 
（H+H） 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeCross2 
1 
1 
ＳＲＣ柱断面鉄骨形状・＋形 
（H+T+T） 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeT 
1 
1 
ＳＲＣ柱断面鉄骨形状・Ｔ形

## Page 273

6.4-8 
 
・補足 
鉄骨形状は下記による。

## Page 274

6.4-9 
 
6.4.3.1.1. ＳＲＣ柱断面鉄骨形状・Ｈ形：StbSecSteelColumn_SRC_ShapeH 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状・Ｈ形 
親要素：
StbSecSteelColumn_SRC_Same
、
StbSecSteelColumn_SRC_NotSame
、
StbSecSteelColumn_SRC_ThreeTypes 
 
・属性 
属性名 
型 
必須 
説明 
補足 
direction_type 
string 
○ 
鉄骨の向き 以下のいずれか 
H（同方向）、I（直交） 
※(1) 
shape 
string 
○ 
鉄骨形状 
※(2) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(3) 
offset_X 
double 
 
鉄骨の偏心（X 方向） 
※(1) 
offset_Y 
double 
 
鉄骨の偏心（Y 方向） 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  ＲＣ断面に対する鉄骨の向きは、鉄骨断面要素 <StbSecSteel> で用意した断面形状の「鉄骨断
面の基準方向」と部材座標系Ｘ軸との関係で定義し、ＲＣ断面に対する鉄骨の位置関係は、相互の
図心間の部材座標系における距離（鉄骨の偏心）で表す。距離0 の場合は「鉄骨の偏心」を省略し
てもよい。

## Page 275

6.4-10 
 
(2)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(3)  省略された場合は、「鉄骨強度（主）」と同一とする。 
 
・例 
 
 
  <StbSecColumn_SRC id="23" name="C11"  （略）> 
    <StbSecFigureColumn_SRC> 
      <StbSecColumnRect width_X="900" width_Y="900"/> 
    </StbSecFigureColumn_SRC> 
    <StbSecBarArrangementColumn_SRC  （略）> 
      <StbSecBarColumnRectSame  （略）/> 
    </StbSecBarArrangementColumn_SRC> 
    <StbSecSteelFigureColumn_SRC> 
      <StbSecSteelColumn_SRC_Same> 
        <StbSecSteelColumn_SRC_ShapeH direction_type ="H" shape="BH-600x200x12x25" 
strength_main="SN490" offset_X="100" offset_Y="100"/> 
</StbSecSteelColumn_SRC_Same> 
    </StbSecSteelFigureColumn_SRC> 
  </StbSecColumn_SRC>

## Page 276

6.4-11 
 
6.4.3.1.2. ＳＲＣ柱断面鉄骨形状・□形：StbSecSteelColumn_SRC_ShapeBox 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状・□形 
親要素：
StbSecSteelColumn_SRC_Same
、
StbSecSteelColumn_SRC_NotSame
、
StbSecSteelColumn_SRC_ThreeTypes 
 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
角形鋼管形状 
※(1) 
encase_type 
string 
○ 
鋼管コンクリートのタイプ 
以下のいずれか 
ENCASED（被覆形） 
ENCASEDANDINFILLED 
（充填被覆形） 
 
strength 
string 
○ 
鉄骨強度 
 
offset_X 
double 
 
鉄骨の偏心（X 方向） 
※(2) 
offset_Y 
double 
 
鉄骨の偏心（Y 方向） 
※(2) 
strength_inner_concrete 
string 
 
充填コンクリート強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(2)  ＲＣ断面に対する鉄骨の位置関係は、
相互の図心間の部材座標系における距
離（鉄骨の偏心）で表す。距離0 の場合
は「鉄骨の偏心」を省略してもよい。

## Page 277

6.4-12 
 
6.4.3.1.3. ＳＲＣ柱断面鉄骨形状・○形：StbSecSteelColumn_SRC_ShapePipe 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状・○形 
親要素：
StbSecSteelColumn_SRC_Same
、
StbSecSteelColumn_SRC_NotSame
、
StbSecSteelColumn_SRC_ThreeTypes 
 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
鋼管形状 
※(1) 
encase_type 
string 
○ 
鋼管コンクリートのタイプ 
以下のいずれか 
ENCASED（被覆形） 
ENCASEDANDINFILLED 
（充填被覆形） 
 
strength 
string 
○ 
鉄骨強度 
 
offset_X 
double 
 
鉄骨の偏心（X 方向） 
※(2) 
offset_Y 
double 
 
鉄骨の偏心（Y 方向） 
※(2) 
strength_inner_concrete 
string 
 
充填コンクリート強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(2)  ＲＣ断面に対する鉄骨の位置関係は、
相互の図心間の部材座標系における距
離（鉄骨の偏心）で表す。距離0 の場合
は「鉄骨の偏心」を省略してもよい。

## Page 278

6.4-13 
 
6.4.3.1.4. ＳＲＣ柱断面鉄骨形状・＋形（H+H）：StbSecSteelColumn_SRC_ShapeCross1 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状・＋形（H+H） 
親要素：
StbSecSteelColumn_SRC_Same
、
StbSecSteelColumn_SRC_NotSame
、
StbSecSteelColumn_SRC_ThreeTypes 
 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape_X 
string 
○ 
X 方向鉄骨形状 
※(1) ※(2) 
shape_Y 
string 
○ 
Y 方向鉄骨形状 
※(1) ※(2) 
strength_main_X 
string 
○ 
X 方向鉄骨強度（主） 
 
strength_web_X 
string 
 
X 方向鉄骨強度（ウェブ） 
※(3) 
strength_main_Y 
string 
 
Y 方向鉄骨強度（主） 
※(3) 
strength_web_Y 
string 
 
Y 方向鉄骨強度（ウェブ） 
※(3) 
offset_XX 
double 
 
X 方向鉄骨の偏心（X 方向） 
※(1) 
offset_XY 
double 
 
X 方向鉄骨の偏心（Y 方向） 
※(1) 
offset_YX 
double 
 
Y 方向鉄骨の偏心（X 方向） 
※(1) 
offset_YY 
double 
 
Y 方向鉄骨の偏心（Y 方向） 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  下図に示す、＋形を構成する2 個のＨ形鋼をそれぞれ「X 方向鉄骨」「Y 方向鉄骨」とし、ＲＣ
断面に対する各々の鉄骨の位置関係は、相互の図心間の部材座標系における距離（鉄骨の偏心）で
表す。距離0 の場合は「鉄骨の偏心」を省略してもよい。

## Page 279

6.4-14 
 
 
 
(2)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(3)  省略された場合は、「X 方向鉄骨強度（主）」と同一とする。 
 
・例 
 
 
  <StbSecColumn_SRC id="24" name="C12"  （略）> 
    <StbSecFigureColumn_SRC> 
      <StbSecColumn_SRC_Rect width_X="900" width_Y="900"/> 
    </StbSecFigureColumn_SRC> 
    <StbSecBarArrangementColumn_SRC  （略）> 
      <StbSecBarColumn_SRC_RectSame  （略）/> 
    </StbSecBarArrangementColumn_SRC> 
    <StbSecSteelFigureColumn_SRC> 
      <StbSecSteelColumn_SRC_Same> 
        <StbSecSteelColumn_SRC_ShapeCross1 shape_X="BH-600x200x12x25" 
                 shape_Y="BH-600x200x12x25" strength_main_X="SN490"/> 
</StbSecSteelColumn_SRC_Same> 
    </StbSecSteelFigureColumn_SRC> 
  </StbSecColumn_SRC>

## Page 280

6.4-15 
 
6.4.3.1.5. ＳＲＣ柱断面鉄骨形状・＋形（H+T+T）：StbSecSteelColumn_SRC_ShapeCross2 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状・＋形（H+T+T） 
親要素：
StbSecSteelColumn_SRC_Same
、
StbSecSteelColumn_SRC_NotSame
、
StbSecSteelColumn_SRC_ThreeTypes 
 
 
・属性 
属性名 
型 
必須 
説明 
補足 
direction_type 
string 
〇 
H 形鋼鉄骨の向き 以下のいずれか 
H（同方向）、I（直交） 
※(1) 
shape_H 
string 
○ 
H 形鋼鉄骨形状 
※(1) ※(2) 
shape_T1 
string 
○ 
T 形鋼鉄骨形状1 
※(1) ※(2) 
shape_T2 
string 
〇 
T 形鋼鉄骨形状2 
※(1) ※(2) 
strength_main_H 
string 
○ 
H 形鋼鉄骨強度（主） 
 
strength_web_H 
string 
 
H 形鋼鉄骨強度（ウェブ） 
※(3) 
strength_main_T 
string 
 
T 形鋼鉄骨強度（主） 
※(3) 
strength_web_T 
String 
 
T 形鋼鉄骨強度（ウェブ） 
※(3) 
offset_HX 
double 
 
H 形鋼鉄骨の偏心（X 方向） 
※(1) 
offset_HY 
double 
 
H 形鋼鉄骨の偏心（Y 方向） 
※(1) 
offset_T 
double 
 
T 形鋼鉄骨の偏心 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  下図に示す、＋形を構成するＨ形鋼の向きは、鉄骨断面要素 <StbSecSteel> で用意した断面形
状の「鉄骨断面の基準方向」と部材座標系Ｘ軸との関係で定義する。ＲＣ断面に対する各々の鉄骨
の位置関係は、相互の図心間の部材座標系における距離（鉄骨の偏心）で表す。距離0 の場合は
「鉄骨の偏心」を省略してもよい。

## Page 281

6.4-16 
 
 
 
(2)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(3)  省略された場合は、「H 形鋼鉄骨強度（主）」と同一とする。 
 
・例 
 
 
  <StbSecColumn_SRC id="24" name="C12"  （略）> 
    <StbSecFigureColumn_SRC> 
      <StbSecColumnRect width_X="900" width_Y="900"/> 
    </StbSecFigureColumn_SRC> 
    <StbSecBarArrangementColumn_SRC  （略）> 
      <StbSecBarColumnRectSame  （略）/> 
    </StbSecBarArrangementColumn_SRC> 
    <StbSecSteelFigureColumn_SRC> 
      <StbSecSteelColumn_SRC_Same> 
        <StbSecSteelColumn_SRC_ShapeCross2 shape_H="BH-600x200x12x25" 
                 shape_T1="T-300x200x12x25" shape_T2="T-300x200x12x25" 
                 strength_main_H="SN490"/> 
</StbSecSteelColumn_SRC_Same> 
    </StbSecSteelFigureColumn_SRC> 
/S bS
C l
S C

## Page 282

6.4-17 
 
6.4.3.1.6. ＳＲＣ柱断面鉄骨形状・Ｔ形：StbSecSteelColumn_SRC_ShapeT 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状・Ｔ形 
親要素：
StbSecSteelColumn_SRC_Same
、
StbSecSteelColumn_SRC_NotSame
、
StbSecSteelColumn_SRC_ThreeTypes 
 
 
・属性 
属性名 
型 
必須 
説明 
補足 
direction_type 
string 
○ 
鉄骨の向き 以下のいずれか 
T1：┬、T2：┤、T3：┴、T4：├ 
※(1) 
shape_H 
string 
○ 
H 形鋼鉄骨形状 
※(2) 
shape_T 
string 
○ 
T 形鋼鉄骨形状 
※(2) 
strength_main_H 
string 
○ 
H 形鋼鉄骨強度（主） 
 
strength_web_H 
string 
 
H 形鋼鉄骨強度（ウェブ） 
※(3) 
strength_main_T 
string 
○ 
T 形鋼鉄骨強度（主） 
※(3) 
strength_web_T 
string 
 
T 形鋼鉄骨強度（ウェブ） 
※(3) 
offset_HX 
double 
 
H 形鋼鉄骨の偏心（X 方向） 
※(4) 
offset_HY 
double 
 
H 形鋼鉄骨の偏心（Y 方向） 
※(4) 
offset_T 
double 
 
T 形鋼鉄骨の偏心（Y 方向） 
※(4) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  ＲＣ断面に対する鉄骨の向きは、それぞれ下図による。

## Page 283

6.4-18 
 
(2)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(3)  省略された場合は、「H 形鋼鉄骨強度（主）」と同一とする。 
(4)  ＲＣ断面に対する各々の鉄骨の位置関係は、Ｔ形を構成するＨ形鋼については、相互の図心間
の部材座標系における距離（鉄骨の偏心）で表す。Ｔ形を構成するＴ形鋼については、下図に示す
距離で表す。鉄骨断面の計算上の図心位置がＲＣ断面と一致する場合は「鉄骨の偏心」を省略して
もよい。

## Page 284

6.4-19 
 
6.4.3.2. ＳＲＣ柱断面鉄骨形状・柱頭脚別：StbSecSteelColumn_SRC_NotSame 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状（柱頭・柱脚が別形状の場合） 
親要素：StbSecSteelFigureColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 以下のいずれか 
BOTTOM（柱脚） 
TOP（柱頭） 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeH 
１ 
１ 
ＳＲＣ柱断面鉄骨形状 
・Ｈ形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeBox 
１ 
１ 
ＳＲＣ柱断面鉄骨形状 
・□形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapePipe 
１ 
１ 
ＳＲＣ柱断面鉄骨形状 
・○形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeCross1 
１ 
１ 
ＳＲＣ柱断面鉄骨形状 
・＋形（H+H） 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeCross2 
１ 
１ 
ＳＲＣ柱断面鉄骨形状 
・＋形（H+T+T） 
 
   または

## Page 285

6.4-20 
 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeT 
１ 
１ 
ＳＲＣ柱断面鉄骨形状 
・Ｔ形 
 
 
・補足 
鉄骨形状は「ＳＲＣ柱断面鉄骨形状・同一」による。 
属性をそれぞれpos= “BOTTOM” およびpos= “TOP” とした子要素を各１回記述する。 
 
・例 
 
 
  <StbSecColumn_SRC id="25" name="C13" floor="1" strength_concrete="FC24"> 
    <StbSecFigureColumn_SRC> 
      <StbSecColumnRect width_X="900" width_Y="900"/> 
    </StbSecFigureColumn_SRC> 
    <StbSecBarArrangementColumn_SRC  （略）> 
      <StbSecBarColumnRectSame  （略）/> 
    </StbSecBarArrangementColumn_SRC> 
    <StbSecSteelFigureColumn_SRC> 
      <StbSecSteelColumn_SRC_NotSame pos="BOTTOM"> 
        <StbSecSteelColumn_SRC_ShapeCross1 shape_X="BH-600x200x12x25" 
                 shape_Y="BH-600x200x12x25" strength_main_X="SN490"/> 
</StbSecSteelColumn_SRC_NotSame> 
<StbSecSteelColumn_SRC_NotSame pos="TOP"> 
        <StbSecSteelColumn_SRC_ShapeCross1 shape_X="BH-600x200x12x22" 
                 shape_Y="BH-600x200x12x22" strength_main_X="SN490"/> 
</StbSecSteelColumn_SRC_NotSame> 
    </StbSecSteelFigureColumn_SRC> 
  </StbSecColumn_SRC>

## Page 286

6.4-21 
 
6.4.3.3. ＳＲＣ柱断面鉄骨形状・３種類：StbSecSteelColumn_SRC_ThreeTypes 
・概要 
説明 ：ＳＲＣ柱鉄骨断面の形状（柱頭・中央・柱脚が別形状の場合） 
親要素：StbSecSteelFigureColumn_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 以下のいずれか 
BOTTOM（柱脚） CENTER（中央） 
TOP（柱頭） 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeH 
1 
1 
ＳＲＣ柱断面鉄骨形状 
・Ｈ形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeBox 
1 
1 
ＳＲＣ柱断面鉄骨形状 
・□形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapePipe 
1 
1 
ＳＲＣ柱断面鉄骨形状 
・○形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeCross1 
1 
1 
ＳＲＣ柱断面鉄骨形状 
・＋形（H+H） 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeCross2 
1 
1 
ＳＲＣ柱断面鉄骨形状 
・＋形（H+T+T） 
 
または

## Page 287

6.4-22 
 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_SRC_ShapeT 
1 
1 
ＳＲＣ柱断面鉄骨形状 
・Ｔ形 
 
 
・補足 
鉄骨形状は「ＳＲＣ柱断面鉄骨形状・同一」による。 
属性をpos= “BOTTOM” , pos= “CENTER” およびpos= “TOP” とした子要素を各１回記述する。 
 
・例 
  <StbSecColumn_SRC id="26" name="C14" floor="1" strength_concrete="FC24"> 
    <StbSecFigureColumn_SRC> 
      <StbSecColumnRect width_X="900" width_Y="900"/> 
    </StbSecFigureColumn_SRC> 
    <StbSecBarArrangementColumn_SRC  （略）> 
      <StbSecBarColumnRectSame  （略）/> 
    </StbSecBarArrangementColumn_SRC> 
    <StbSecSteelFigureColumn_SRC> 
      <StbSecSteelColumn_SRC_ThreeTypes pos="BOTTOM"> 
        <StbSecSteelColumn_SRC_ShapeCross1 shape_X="BH-600x200x12x25" 
                 shape_Y="BH-600x200x12x25" strength_main_X="SN490"/> 
</StbSecSteelColumn_SRC_ThreeTypes> 
<StbSecSteelColumn_SRC_ThreeTypes pos="CENTER"> 
        <StbSecSteelColumn_SRC_ShapeCross1 shape_X="BH-600x200x12x22" 
                 shape_Y="BH-600x200x12x22" strength_main_X="SN490"/> 
</StbSecSteelColumn_SRC_ThreeTypes> 
<StbSecSteelColumn_SRC_ThreeTypes pos="TOP"> 
        <StbSecSteelColumn_SRC_ShapeCross1 shape_X="BH-600x200x12x25" 
                 shape_Y="BH-600x200x12x25" strength_main_X="SN490"/> 
</StbSecSteelColumn_SRC_ThreeTypes> 
    </StbSecSteelFigureColumn_SRC> 
  </StbSecColumn_SRC>

## Page 288

6.5-1 
 
6.5. ＣＦＴ柱断面：StbSecColumn_CFT 
・概要 
説明 ：ＣＦＴ柱断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※
(2) 
kind_column 
string 
 
柱の種別 
以下のいずれか 
COLUMN（柱） 
POST（間柱） 
※(3) 
strength_concrete 
string 
 
コンクリート強度 
※(4) 
isReferenceDirection 
boolean 
 
鉄骨向き 
※(5) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelFigureColumn_CFT 
1 
1 
ＣＦＴ柱断面鉄骨形状 
 
 
   上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseProduct 
0 
1 
鉄骨断面柱脚製品 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBaseConventional 
0 
1 
鉄骨断面柱脚在来工法

## Page 289

6.5-2 
 
・補足 
CFT の鉄骨部分は、鉄骨断面要素 <StbSecSteel> で用意した断面形状を、この要素の子要素で参照
することで定義する。 
 
(1)  「断面名称」は、部材リスト（構造図の柱断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「C1」で「1C1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、COLUMN とする。 
(4)  省略された場合は、参照する<StbColumn>の「終端節点ID」id_node_top が所属する<StbStory>
のコンクリート強度を、この要素のコンクリート強度とする。参照した<StbStory>のコンクリー
ト強度が省略されていた場合は、共通情報の属性「建物全体のコンクリート強度」
strength_concrete をこの要素のコンクリート強度とする。 
(5)  true の場合は、<StbSecSteel> で用意した断面形状の「鉄骨断面の基準方向」を、部材の断面
軸（Ｙ軸）とする。false の場合は、90 度傾けた断面形状を部材の断面軸とする。Ｈ型鋼において、
同一断面を部材リストに90 度傾けて表記する場合などを想定している。省略された場合は、true 
とする。 
 
・例 
 
 
 
 
  <StbSecColumn_CFT id="121" name="C1" floor="1" kind_column="COLUMN" 
strength_concrete="FC24"> 
    <StbSecSteelFigureColumn_CFT> 
      <StbSecSteelColumn_CFT_Same shape="□-700x700x25x25" 
strength_main="BCP325"/> 
    </StbSecSteelFigureColumn_CFT> 
  </StbSecColumn_CFT>

## Page 290

6.5-3 
 
6.5.1. ＣＦＴ柱断面鉄骨形状：StbSecSteelFigureColumn_CFT 
・概要 
説明 ：ＣＦＴ柱断面の形状 
親要素：StbSecColumn_CFT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
base_type 
string 
 
柱脚形式 以下のいずれか 
NONE（鉄骨柱脚なし） 
EXPOSE（露出） 
EMBEDDED（埋込） 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_CFT_Same 
1 
1 
ＣＦＴ柱断面鉄骨形状・同一 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_CFT_NotSame 
2 
2 
ＣＦＴ柱断面鉄骨形状・柱頭脚別 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelColumn_CFT_ThreeTypes 
3 
3 
ＣＦＴ柱断面鉄骨形状・３種類 
※(2)

## Page 291

6.5-4 
 
・補足 
(1)  省略された場合は、NONE とする。 
(2)  断面形状は、鉄骨の継手位置で切り替わるものとし、ハンチ状の変断面状態は想定しない。下図
のように、継手位置の箇所数に応じて、記述する子要素を選択する。 
 
 
 
 
 
 
柱頭
柱脚
中央
柱頭
柱脚
子要素：
同一
柱頭脚別
３種類

## Page 292

6.5-5 
 
6.5.1.1. ＣＦＴ柱断面鉄骨形状・同一：StbSecSteelColumn_CFT_Same 
・概要 
説明 ：ＣＦＴ柱鉄骨断面の形状（全断面同一の場合） 
親要素：StbSecSteelFigureColumn_CFT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength 
string 
○ 
鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
 
・例 
 
 
 
  <StbSecColumn_CFT id="121" name="C1" floor="1" kind_column="COLUMN" 
strength_concrete="FC24"> 
    <StbSecSteelFigureColumn_CFT> 
      <StbSecSteelColumn_CFT_Same shape="□-700x700x25x25" 
strength ="BCP325"/> 
    </StbSecSteelFigureColumn_CFT> 
  </StbSecColumn_CFT>

## Page 293

6.5-6 
 
6.5.1.2. ＣＦＴ柱断面鉄骨形状・柱頭脚別：StbSecSteelColumn_CFT_NotSame 
・概要 
説明 ：ＣＦＴ柱鉄骨断面の形状（柱頭・柱脚が別形状の場合） 
親要素：StbSecSteelFigureColumn_CFT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 
以下のいずれか 
BOTTOM（柱脚） 
TOP（柱頭） 
 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength 
string 
○ 
鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
属性をそれぞれpos= “BOTTOM” およびpos= “TOP” とした子要素を各１回記述する。 
定義と補足内容 (1) は、「ＣＦＴ柱断面鉄骨形状・同一」による。

## Page 294

6.5-7 
 
6.5.1.3. ＣＦＴ柱断面鉄骨形状・３種類：StbSecSteelColumn_CFT_ThreeTypes 
・概要 
説明 ：ＣＦＴ柱鉄骨断面の形状（柱頭・中央・柱脚が別形状の場合） 
親要素：StbSecSteelFigureColumn_CFT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 
以下のいずれか 
BOTTOM（柱脚） 
CENTER（中央） 
TOP（柱頭） 
 
shape 
string 
○ 
鉄骨形状 
 
strength 
string 
○ 
鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
属性をpos= “BOTTOM” , pos= “CENTER” およびpos= “TOP” とした子要素を各１回記述する。 
定義と補足内容 (1) は、「ＣＦＴ柱断面鉄骨形状・同一」による。

## Page 295

6.6-1 
 
6.6. ＲＣ梁断面：StbSecBeam_RC 
・概要 
説明 ：ＲＣ梁断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※(2) 
kind_beam 
string 
 
梁の種別 
以下のいずれか 
GIRDER（大梁） 
BEAM（小梁） 
※(3) 
isFoundation 
boolean 
 
基礎梁か否か 
部材リスト用※(4) 
isCanti 
boolean 
 
片持ち梁か否か 
部材リスト用※(4) 
isOutin 
boolean 
 
外端・内端指定 
※(4) ※(6) 
strength_concrete 
string 
 
コンクリート強度 
※(5) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureBeam_RC 
0 
制限なし 
ＲＣ梁断面形状 
※(8) 
StbSecBarArrangementBeam_RC 
0 
制限なし 
ＲＣ梁断面配筋 
※(7) 
 
・補足 
(1)  「断面名称」は、部材リスト（構造図の梁断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「G1」で「1G1」となる）。 
(2) 「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、GIRDER とする。 
(4)  省略された場合は、false とする。

## Page 296

6.6-2 
 
(5)  省略された場合は、参照する<StbGirder>の「始端節点ID」id_node_start が所属階する
<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した<StbStory>の
コンクリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリート強度」
strength_concrete をこの要素のコンクリート強度とする。 
(6)  「外端・内端指定」がtrue の場合、子要素の記述中、「始端」を「外端」（片持ち梁の場合は
元端）、「終端」を「内端」（片持ち梁の場合は先端）と読み替える。 
(7)  子要素<StbSecBarArrangementBeam_RC>の回数が0 となる場合は、鉄筋を扱わないプログ
ラムが一時的に作成する場合を想定しており、無筋であることを示すものではない。 
(8)  子要素<StbSecFigureGirder_RC>の回数が0 となる場合は、構造計算プログラムが計算対象と
しない形状で、断面性能を直接指定する場合を想定している。この場合は、他の子要素も指定して
はならない。 
・例 
   
<StbSecBeam_RC id="26" name="G1" floor="2" strength_concrete="FC24"> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1"> 
      <StbSecBarBeamSimple （略）/> 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC> 
 
  <StbSecBeam_RC id="27" name="G2" floor="2" strength_concrete="FC24"> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecFigureBeam_RC order="2"> 
      <StbSecBeamStraight width="400" depth="500"/> 
    </StbSecFigureBeam_RC> 
    <StbSecFigureBeam_RC order="3"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1"> 
      <StbSecBarBeamSimple （略）/> 
    </StbSecBarArrangementBeam_RC> 
    <StbSecBarArrangementBeam_RC order="2"> 
      <StbSecBarBeamSimple （略）/> 
    </StbSecBarArrangementBeam_RC> 
    <StbSecBarArrangementBeam_RC order="3"> 
      <StbSecBarBeamSimple （略）/> 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 297

6.6-3 
 
6.6.1. ＲＣ梁断面形状：StbSecFigureBeam_RC 
・概要 
説明 ：ＲＣ梁断面の形状 
親要素：StbSecBeam_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
○ 
始端からの順番 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBeamStraight 
1 
1 
コンクリート梁断面形状・スト
レート 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBeamTaper 
1 
1 
コンクリート梁断面形状・テー
パー 
※(1) 
    
・補足 
(1) StbSecBeamStraight で表現できないものにのみ使用する。

## Page 298

6.6-4 
 
6.6.1.1. コンクリート梁断面形状・ストレート：StbSecBeamStraight 
・概要 
説明 ：コンクリート梁断面の形状（全部位同一寸法の場合） 
親要素：StbSecFigureBeam_RC、StbSecFigureBeam_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width 
double 
○ 
幅 
※(1) 
depth 
double 
○ 
成 
※(1) 
horizontal_offset 
double 
 
水平方向のオフセット 
※(2) 
vertical_offset 
double 
 
鉛直方向のオフセット 
※(3) 
 
・内容、子要素 
無し 
 
・補足 
(1) 「幅」および「成」の定義は下図による。 
(2) 始端基準点から終端基準点に向かう方向をX 軸とし、Y 方向に移動した距離を表す。 
(3) 始端基準点と終端基準点を結ぶ線から梁上端がZ 方向に移動した距離を表す 
 
 
・例 
 
 
 
Y
Z
成
幅
   <StbSecBeam_RC id="26" name="G1" （略）> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC （略）> 
        （略） 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 299

6.6-5 
 
6.6.1.2. コンクリート梁断面形状・テーパー：StbSecBeamTaper 
・概要 
説明 ：コンクリート梁断面の形状（始端、終端異寸法の場合） 
親要素：StbSecFigureBeam_RC、StbSecFigureBeam_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
start_width 
double 
○ 
始端側の幅 
※(1) 
start_depth 
double 
○ 
始端側の成 
※(1) 
end_width 
double 
〇 
終端側の幅 
※(1) 
end_depth 
double 
〇 
終端側の成 
※(1) 
start_horizontal_offset 
double 
 
始端側の水平方向のオフセット 
※(2) 
start_vertucal_offset 
double 
 
始端側の鉛直方向のオフセット 
※(3) 
end_horizontal_offset 
double 
 
終端側の水平方向のオフセット 
※(2) 
end_vertical_offset 
double 
 
終端側の鉛直方向のオフセット 
※(3) 
 
・内容、子要素 
無し 
 
・補足 
(1) 始端と終端を結んだ変断面形状とする。「幅」および「成」の定義は下図による。 
(2) 始端基準点から終端基準点に向かう方向をX 軸とし、Y 方向に移動した距離を表す。 
(3) 始端基準点と終端基準点を結ぶ線から梁上端がZ 方向に移動した距離を表す 
 
 
始端
終端
Y
Z
幅：始端
幅：終端
成：終端
成：始端

## Page 300

6.6-6 
 
・例 
 
 
 
  <StbSecBeam_RC id="28" name="G3" （略）> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamTaper start_width="400" start_depth="650" end_width="400"  
end_depth="500"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC （略）> 
        （略） 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 301

6.6-7 
 
6.6.2. ＲＣ梁断面配筋：StbSecBarArrangementBeam_RC 
・概要 
説明 ：ＲＣ梁断面の配筋 
親要素：StbSecBeam_RC 
 
・属性 
属性名 
型 
必須 
説明 
補
足 
order 
integer 
○ 
配筋位置 
始端からの番号 
 
D_bar_spacing 
double 
 
巾止筋：径 
 
strength_bar_spacing 
double 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
D_catch_bar 
string 
 
段取り筋 
 
strength_catch_bar 
string 
 
段取り筋強度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamSimple 
1 
1 
コンクリート梁断面配筋・簡易 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamComplex 
1 
1 
コンクリート梁断面配筋・詳細 
※(2) 
 
   上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamXReinforced 
0 
1 
コンクリート梁 Ｘ形配筋 
※(1) 
 
・補足 
(1)  子要素<StbSecBarBeamXReinforced> は、X 形配筋を使用する場合に用いる。X 形配筋部をこ
の要素で追加し、平行配筋部は子要素<StbSecBarBeamSimple> または
<StbSecBarBeamComplex>で定義する。 
(2) 鉄筋位置をすべて規定する場合に使用する。

## Page 302

6.6-8 
 
 
・例 
 
 
 
  <StbSecBeam_RC id="26" name="G1" （略）> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1" D_bar_spacing="D10"  
pitch_bar_spacing ="300" > 
      <StbSecBarBeamSimple depth_cover_left="40" depth_cover_right="40"  
depth_cover_top="40" depth_cover_bottom="40" interval="70"  
D_stirrup="D13" N_stirrup="4" pitch_stirrup="150" D_web="D10" N_web="4"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="1" D="D22" N="7"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="2" D="D22" N="6"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="3" D="D22" N="4"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="1" D="D22" N="7"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="2" D="D22" N="6"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="3" D="D22" N="4"/> 
      <StbSecBarBeamSimple/> 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 303

6.6-9 
 
6.6.2.1. コンクリート梁断面配筋・簡易：StbSecBarBeamSimple 
・概要 
説明 ：コンクリート梁断面配筋・簡易 
親要素：StbSecBarArrangementBeam_RC、StbSecBarArrangementBeam_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_left 
double 
 
かぶり厚さ（左） 
※(1) 
depth_cover_right 
double 
 
かぶり厚さ（右） 
※(1) 
depth_cover_top 
double 
 
かぶり厚さ（上） 
※(1)  
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
※(1) 
interval 
double 
 
２段筋のあき 
※(1) 
center_top 
double 
 
主筋重心位置（上） 
※(2) 
center_bottom 
double 
 
主筋重心位置（下） 
※(2) 
center_side 
double 
 
主筋重心位置（側） 
※(2) 
center_interval 
double 
 
２段筋重心間距離 
※(2) 
D_stirrup 
string 
〇 
あばら筋：径 
 
N_stirrup 
integer 
〇 
あばら筋：本数 
 
strength_stirrup 
string 
 
あばら筋：鉄筋強度 
 
pitch_stirrup 
double 
〇 
あばら筋：ピッチ 
 
D_web 
string 
 
腹筋：径 
 
N_web 
integer 
 
腹筋：本数 
 
strength_web 
string 
 
腹筋：鉄筋強度 
 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamSimpleMain 
2 
制限なし 
コンクリート梁簡易主筋 
※(3) 
StbSecBarBeamAdditional 
0 
制限なし 
コンクリート梁断面配筋・追加
鉄筋 
 
 
・補足

## Page 304

6.6-10 
 
(1) 「かぶり厚さ（…）」および「２段筋のあき」の定義は下図による。 
 
 
「かぶり厚さ（…）」「主筋重心位置（…）」および「２段筋のあき」が省略された場合の扱いは、
<StbApplyConditionsList>の補足説明による。 
(2)  「かぶり厚さ」と「重心位置」の扱いについては、第２章を参照のこと。 
(3) 同じ段に複数種類の径がある場合は別要素とする。 
 
 
かぶり厚さ
（下）
かぶり厚さ
（上）
Y
Z
かぶり厚さ
（左）
かぶり厚さ
（右）
2段筋のあき
  <StbSecBeam_RC id="26" name="G1" （略）> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1" D_bar_spacing="D10"  
pitch_bar_spacing ="300" > 
      <StbSecBarBeamSimple depth_cover_left="40" depth_cover_right="40"  
depth_cover_top="40" depth_cover_bottom="40" interval="70"  
D_stirrup="D13" N_stirrup="4" pitch_stirrup="150" D_web="D10" N_web="4"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="1" D="D22" N="7"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="2" D="D22" N="6"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="3" D="D22" N="4"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="1" D="D22" N="7"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="2" D="D22" N="6"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="3" D="D22" N="4"/> 
      <StbSecBarBeamSimple/> 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 305

6.6-11 
 
6.6.2.1.1. コンクリート梁断面配筋・簡易主筋：StbSecBarBeamSimpleMain 
・概要 
説明 ： コンクリート梁断面配筋・簡易主筋 
親要素： StbSecBarBeamSimple 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
〇 
配置位置 
以下のいずれか 
BOTTOM：下端 
TOP：上端 
 
step 
integer 
〇 
段数 
 
D 
string 
〇 
主筋径 
 
strength 
string 
 
主筋強度 
 
N 
integer 
〇 
主筋本数 
 
N_left 
integer  
 
左側に寄せて配置する本数 
※(1) 
N_right 
integer 
 
右側に寄せて配置する本数 
※(1) 
interval 
double 
 
寄せ筋間隔 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 寄筋とする本数を記入する。

## Page 306

6.6-12 
 
6.6.2.1.2. コンクリート梁断面配筋・追加鉄筋：StbSecBarBeamAdditional 
・概要 
説明 ：コンクリート梁断面配筋・追加鉄筋 
親要素：StbSecBarBeamSimple、StbSecBarBeamComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
径 
 
strength 
string 
 
強度 
 
Y 
double 
〇 
Y 座標 
※(1) 
Z 
double 
〇 
Z 座標 
※(1) 
isStructual 
boolean 
 
構造耐力上必要な鉄筋かどうか 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 図心を原点とし、鉄筋芯位置を指定する。 
  <StbSecBeam_RC id="26" name="G1" （略）> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="650"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1" D_bar_spacing="D10"  
pitch_bar_spacing ="300" > 
      <StbSecBarBeamSimple depth_cover_left="40" depth_cover_right="40"  
depth_cover_top="40" depth_cover_bottom="40" interval="70"  
D_stirrup="D13" N_stirrup="4" pitch_stirrup="150" D_web="D10" N_web="4"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="1" D="D22" N="7"/> 
        <StbSecBarBeamSimpleMain pos="TOP" step="2" D="D22" N="6"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="1" D="D22" N="7"/> 
        <StbSecBarBeamSimpleMain pos="BOTTOM" step="2" D="D22" N="6"/> 
        <StbSecBarBeamAdditional D="D19" Y="138" Z="60"/> 
        <StbSecBarBeamAdditional D="D19" Y="-138" Z="-60"/> 
      <StbSecBarBeamSimple/> 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 307

6.6-13 
 
 
 
6.6.2.2. コンクリート梁断面配筋・３種類・詳細：StbSecBarBeamComplex 
・概要 
説明 ：コンクリート梁断面配筋・３種類・詳細 
親要素：StbSecBarArrangementBeam_RC、StbSecBarArrangementBeam_SRC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamComplexMain 
1 
1 
コンクリート梁断面配筋・詳
細・主筋 
 
StbSecBarBeamComplexStirrup 
1 
1 
コンクリート梁断面配筋・詳
細・あばら筋 
 
StbSecBarBeamComplexWeb 
0 
1 
コンクリート梁断面配筋・詳
細・腹筋 
※(1) 
StbSecBarBeamAdditional 
0 
制限なし 
コンクリート梁断面配筋・
追加鉄筋 
 
    
・補足 
(1) 腹筋がない場合は回数を0 とする。

## Page 308

6.6-14 
 
6.6.2.2.1. 
  <StbSecBeam_RC id="26" name="G1" （略）> 
    <StbSecFigureBeam_RC order="1"> 
      <StbSecBeamStraight width="400" depth="800"/> 
    </StbSecFigureBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1" D_bar_spacing="D10"  
pitch_bar_spacing ="300" > 
      <StbSecBarBeamComplex> 
        <StbSecBarBeamComplexMain pos="TOP"> 
          <StbSecBarBeamComplexMainLine step="1" depth="64" D="D22"> 
            <StbSecBarBeamComplexMainLoc distance="64"> 
            <StbSecBarBeamComplexMainLoc distance="200"> 
            <StbSecBarBeamComplexMainLoc distance="336"> 
          </StbSecBarBeamComplexMainLine> 
        </StbSecBarBeamComplexMain> 
        <StbSecBarBeamComplexMain pos="BOTTOM"> 
          <StbSecBarBeamComplexMainLine step="1" depth="64" D="D22"> 
            <StbSecBarBeamComplexMainLoc distance="64"> 
            <StbSecBarBeamComplexMainLoc distance="154"> 
            <StbSecBarBeamComplexMainLoc distance="246"> 
            <StbSecBarBeamComplexMainLoc distance="336"> 
          </StbSecBarBeamComplexMainLine> 
        </StbSecBarBeamComplexMain> 
        <StbSecBarBeamComplexStirrup D="D13" pitch="150"> 
          <StbSecBarBeamComplexStirrupLoc distance="47"> 
          <StbSecBarBeamComplexStirrupLoc distance="200"> 
          <StbSecBarBeamComplexStirrupLoc distance="353"> 
        </StbSecBarBeamComplexStirrup> 
        <StbSecBarBeamComplexWeb D="D10"> 
          <StbSecBarBeamComplexWebLoc distance="400"> 
        </StbSecBarBeamComplexWeb> 
      </StbSecBarBeamComplex> 
    </StbSecBarArrangementBeam_RC> 
  </StbSecBeam_RC>

## Page 309

6.6-15 
 
コンクリート梁断面配筋・詳細・主筋：StbSecBarBeamComplexMain 
・概要 
説明 ： コンクリート梁断面配筋・詳細・主筋 
親要素：StbSecBarBeamComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
〇 
配置位置 
以下のいずれか 
BOTTOM：下端 
TOP：上端 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamComplexMainLine 
1 
制限なし 
コンクリート梁断面配筋・詳
細・主筋列 
※(1) 
 
・補足 
(1) 1 段に異なる鉄筋を配置する場合は別要素とする

## Page 310

6.6-16 
 
6.6.2.2.1.1. コンクリート梁断面配筋・詳細・主筋・各行：StbSecBarBeamComplexMainLine 
・概要 
説明 ：コンクリート梁断面配筋・詳細・主筋・各行 
親要素：StbSecBarBeamComplexMain 
 
・属性 
属性名 
型 
必須 
説明 
補足 
step 
integer 
〇 
段数 
 
depth 
double 
〇 
コンクリート端（上端の場合は上面、下
端の場合は下面）から鉄筋芯までの距離 
 
D 
string 
〇 
主筋径 
 
strength 
string 
 
主筋強度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamComplexMainLoc 
1 
制限なし 
コンクリート梁断面配筋・詳
細・主筋位置 
※(1) 
 
・補足 
(1) 本数は要素数とする

## Page 311

6.6-17 
 
6.6.2.2.1.2. コンクリート梁断面配筋・詳細・主筋位置：StbSecBarBeamComplexMainLoc 
・概要 
説明 ： コンクリート梁断面配筋・詳細・主筋位置 
親要素：StbSecBarBeamComplexMainLine 
 
・属性 
属性名 
型 
必須 
説明 
補足 
distance 
double 
〇 
コンクリート左端から鉄筋芯までの距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 312

6.6-18 
 
6.6.2.2.2. コンクリート梁断面配筋・詳細・あばら筋：StbSecBarBeamComplexStirrup 
・概要 
説明 ： コンクリート梁断面配筋・詳細・あばら筋 
親要素：StbSecBarBeamComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
鉄筋径 
 
strength 
string 
 
鉄筋強度 
 
pitch 
double 
〇 
ピッチ 
 
 
 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecBarBeamComplexStirrupLoc 
1 
制限なし 
コンクリート梁断面配筋・詳細・あばら
筋位置 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 上下の位置はかぶり厚から求めるものとする。

## Page 313

6.6-19 
 
6.6.2.2.2.1. コンクリート梁断面配筋・詳細・あばら筋位置：StbSecBarBeamComplexStirrupLoc 
・概要 
説明 ： コンクリート梁断面配筋・詳細・あばら筋位置 
親要素：StbSecBarBeamComplexStirrup 
 
・属性 
属性名 
型 
必須 
説明 
補足 
distance 
double 
〇 
コンクリート左端から鉄筋芯までの距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 314

6.6-20 
 
6.6.2.2.3. コンクリート梁断面配筋・詳細・腹筋：StbSecBarBeamComplexWeb 
・概要 
説明 ： コンクリート梁断面配筋・詳細・腹筋 
親要素：StbSecBarBeamComplex 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
string 
〇 
鉄筋筋径 
 
strength 
string 
 
鉄筋強度 
 
 
・内容 
無し 
 
・子要素 
属性名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamComplexWebLoc 
1 
制限なし 
コンクリート梁断面配筋・詳細・腹筋位
置 
 
 
・補足 
(1) 左右の位置はあばら筋から求めるものとする。

## Page 315

6.6-21 
 
6.6.2.2.3.1. コンクリート梁断面配筋・詳細・腹筋位置：StbSecBarBeamComplexWebLoc 
・概要 
説明 ：コンクリート梁断面配筋・詳細・腹筋位置 
親要素：StbSecBarBeamComplexWeb 
 
・属性 
属性名 
型 
必須 
説明 
補足 
distance 
double 
〇 
コンクリート上端から鉄筋芯までの距離 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 316

6.6-22 
 
6.6.2.3. コンクリート梁 Ｘ形配筋：StbSecBarBeamXReinforced 
・概要 
説明 ：コンクリート梁 X 形配筋 
親要素：StbSecBarArrangementBeam_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_main 
string 
○ 
主筋径 
 
N_main_top 
integer 
○ 
主筋：上端1 段目 
 
N_main_bottom 
integer 
○ 
主筋：下端1 段目 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
・例（断面、形状、あばら筋等は省略） 
 
 
Y 
Z 
X 形配筋 
平行配筋 
  <StbSecBeam_RC> 
    <StbSecBarArrangementBeam_RC order="1" （略）> 
      <StbSecBarBeamSimple （略）> 
      <StbSecBarBeamXReinforced D_main=”D25” N_main_top="2" N_main_bottom="2"/> 
    </StbSecBarArrangementBeam_RC > 
  </StbSecBeam_RC>

## Page 317

6.7-1 
 
6.7. Ｓ梁断面：StbSecBeam_S 
・概要 
説明 ：Ｓ梁断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※(2) 
kind_beam 
string 
 
梁の種別 
以下のいずれか 
GIRDER（大梁） 
BEAM（小梁） 
※(3) 
isCanti 
boolean 
 
片持ち梁か否か 
部材リスト用※(4) 
isOutin 
boolean 
 
外端・内端指定 
※(4) ※(5) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelFigureBeam_S 
1 
1 
Ｓ梁断面鉄骨形状 
 
 
・補足 
鉄骨断面の形状は、鉄骨断面要素 <StbSecSteel> で用意した断面形状を、この要素の子要素で参照す
ることで定義する。このとき <StbSecSteel> で用意した断面形状の「鉄骨断面の基準方向」を、部材の
断面軸（Ｚ軸）とする。 
 
(1)  「断面名称」は、部材リスト（構造図の梁断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「G1」で「1G1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、GIRDER とする。

## Page 318

6.7-2 
 
(4)  省略された場合は、false とする。 
(5)  「外端・内端指定」がtrue の場合、子要素の記述中、「始端」を「外端」（片持ち梁の場合は
基端）、「終端」を「内端」（片持ち梁の場合は先端）と読み替える。 
 
・例 
 
  <StbSecBeam_S id="5" name="SG1" floor="2" kind_beam="GIRDER"> 
    <StbSecSteelFigureBeam_S> 
<StbSecSteelBeam_S_Shape order=”1”> 
      <StbSecSteelBeamStraight shape="H-700x300x13x24x28" strength_main="SN490"/> 
</StbSecSteelBeam_S_Shape> 
    </StbSecSteelFigureBeam_S> 
  </StbSecBeam_S> 
 
  <StbSecBeam_S id="6" name="SB1" kind_beam="BEAM"> 
    <StbSecSteelFigureBeam_S> 
<StbSecSteelBeam_S_Shape order=”1”>  
   <StbSecSteelBeamStraight shape="H-400x200x8x13" strength_main="SS400"/> 
</StbSecSteelBeam_S_Shape> 
    </StbSecSteelFigureBeam_S> 
  </StbSecBeam_S> 
 
  <StbSecBeam_S id="7" name="SG2" floor="2"> 
    <StbSecSteelFigureBeam_S> 
<StbSecSteelBeam_S_Shape order=”1”> 
      <StbSecSteelBeamStraight shape="BH-700x300x14x25" strength_main="SN490"/> 
</StbSecSteelBeam_S_Shape> 
<StbSecSteelBeam_S_Shape order=”2”>   
 <StbSecSteelBeamStraight shape="H-700x300x13x24x28" strength_main="SN490"/> 
</StbSecSteelBeam_S_Shape> 
<StbSecSteelBeam_S_Shape order=”3”> 
        <StbSecSteelBeamStraight shape="BH-700x300x14x25" strength_main="SN490"/> 
</StbSecSteelBeam_S_Shape> 
    </StbSecSteelFigureBeam_S> 
  </StbSecBeam_S> 
 
  <StbSecBeam_S id="8" name="SCG1" isCanti="true"> 
    <StbSecSteelFigureBeam_S> 
<StbSecSteelBeam_S_Shape order=”1”> 
      <StbSecSteelBeamTaper start_shape="BH-600x200x12x25" end_shape="BH-
450x200x12x25" strength_main="SN490"/> 
</StbSecSteelBeam_S_Shape> 
    </StbSecSteelFigureBeam_S> 
  </StbSecBeam_S>

## Page 319

6.7-3 
 
6.7.1. Ｓ梁断面鉄骨形状：StbSecSteelFigureBeam_S 
・概要 
説明 ：Ｓ梁断面の形状 
親要素：StbSecBeam_S 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBeam_S_Shape 
1 
制限なし 
S 梁断面任意鉄骨形状 
 
StbSecSteelBeamWidening 
0 
2 
拡幅要素 
※(1) 
 
・補足 
(1) 拡幅する梁については、形状で表現するのではなく拡幅要素を定義する。

## Page 320

6.7-4 
 
6.7.1.1. Ｓ梁断面任意鉄骨形状：StbSecSteelBeam_S_Shape 
・概要 
説明 ：Ｓ梁断面任意鉄骨形状 
親要素：StbSecSteelFigureBeam_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
○ 
順番 
※(1) 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBeamStraight 
1 
1 
ストレート 
 
または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBeamTaper 
1 
1 
テーパー 
 
 
・補足 
(1) StbSecSteelBeamStraight で表現できないものにのみ使用する。

## Page 321

6.7-5 
 
6.7.1.1.1. ストレート要素：StbSecSteelBeamStraight 
・概要 
説明 ：ストレート要素（始端・終端同一形状の要素） 
親要素：StbSecSteelBeamShape 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
horizontal_offset 
double 
 
左右方向のオフセット 
※(3) 
vertical_offset 
double 
 
上下方向のオフセット 
※(4) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(2) 省略された場合は、「鉄骨強度（主）」と同一とする。 
(3) 始端基準点から終端基準点に向かう方向をX 軸とし、Y 方向に移動した距離を表す。 
(4) 始端基準点と終端基準点を結ぶ線から梁上端がZ 方向に移動した距離を表す

## Page 322

6.7-6 
 
6.7.1.1.2. テーパー要素：StbSecSteelBeamTaper 
・概要 
説明 ：テーパー要素（始端・終端で異なる形状の要素） 
親要素：StbSecSteelBeamShape 
 
・属性 
属性名 
型 
必須 
説明 
補足 
start_shape 
string 
○ 
始端側形状 
※(1) 
end_shape 
string 
○ 
終端側形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
start_horizontal_offset 
double 
 
始端側の左右方向のオフセット 
※(3) 
start_vertical_offset 
double 
 
始端側の上下方向のオフセット 
※(4) 
end_horizontal_offset 
double 
 
終端側の左右方向のオフセット 
※(3) 
end_vertical_offset 
double 
 
終端側の上下方向のオフセット 
※(4) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(2) 省略された場合は、「鉄骨強度（主）」と同一とする。 
(3) 始端基準点から終端基準点に向かう方向をX 軸とし、Y 方向に移動した距離を表す。 
(4) 始端基準点と終端基準点を結ぶ線から梁上端がZ 方向に移動した距離を表す。

## Page 323

6.7-7 
 
6.7.1.2. 拡幅要素：StbSecSteelBeamWidening 
・概要 
説明 ：拡幅要素 
親要素：StbSecSteelFigureBeam_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
形状位置 
以下のいずれか 
START（始端） 
END（終端） 
 
l1 
double 
〇 
水平部長さ 
 
l2 
double 
〇 
サイドPL 長さ 
 
l3 
double 
〇 
サイドPL 水平長さ 
 
l4 
double 
〇 
サイドPL 先端立ち上がり長さ 
 
isWeld 
boolean 
 
溶接する／一枚板 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 324

6.8-1 
 
6.8. ＳＲＣ梁断面：StbSecBeam_SRC 
・概要 
説明 ：ＳＲＣ梁断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※(2) 
kind_beam 
string 
 
梁の種別 
以下のいずれか 
GIRDER（大梁） 
BEAM（小梁） 
※(3) 
isFoundation 
boolean 
 
基礎梁か否か 
部材リスト用※(4) 
isCanti 
boolean 
 
片持ち梁か否か 
部材リスト用※(4) 
isOutin 
boolean 
 
外端・内端指定 
※(4) ※(6) 
strength_concrete 
string 
 
コンクリート強度 
※(5) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureBeam_SRC 
0 
制限なし 
ＳＲＣ梁断面形状 
※(8) 
StbSecBarArrangementBeam_SRC 
0 
制限なし 
ＳＲＣ梁断面配筋 
※(7) 
StbSecSteelFigureBeam_SRC 
0 
1 
ＳＲＣ梁断面鉄骨形状 
※(7) 
 
・補足 
(1)  「断面名称」は、部材リスト（構造図の梁断面表）における、所属階を付けない名称を想定して
いる（所属階「1」と断面名称「G1」で「1G1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3)  省略された場合は、GIRDER とする。

## Page 325

6.8-2 
 
(4)  省略された場合は、false とする。 
(5)  省略された場合は、参照する<StbGirder>の「始端節点ID」id_node_startが所属する<StbStory>
のコンクリート強度を、この要素のコンクリート強度とする。参照した<StbStory>のコンクリー
ト強度が省略されていた場合は、共通情報の属性「建物全体のコンクリート強度」
strength_concrete をこの要素のコンクリート強度とする。 
(6)  「外端・内端指定」がtrue の場合、子要素の記述中、「始端」を「外端」（片持ち梁の場合は
元端）、「終端」を「内端」（片持ち梁の場合は先端）と読み替える。 
(7)  子要素<StbSecBarArrangementBeam_SRC>および<StbSecSteelFigureBeam_SRC>の回数
が0 となる場合は、鉄筋や内部鉄骨を扱わないプログラムが一時的に作成する場合を想定してお
り、無筋や鉄骨無配置であることを示すものではない。 
(8)  子要素<StbSecFigureGirder_SRC>の回数が0 となる場合は、構造計算プログラムが計算対象
としない形状で、断面性能を直接指定する場合を想定している。この場合は、他の子要素も指定し
てはならない。 
・例 
 
 
 
  <StbSecBeam_SRC id="36" name="G11" floor="2" strength_concrete="FC24"> 
    <StbSecFigureBeam_SRC order=”1”> 
      <StbSecBeamStraight width="350" depth="750"/> 
    </StbSecFigureBeam_SRC> 
    <StbSecBarArrangementBeam_SRC order=”1”> 
      <StbSecBarBeamSimple (略)/> 
    </StbSecBarArrangementBeam_SRC> 
    <StbSecSteelFigureBeam_SRC> 
    <StbSecSteelBeam_SRC_Shape order=”1”> 
      <StbSecSteelBeamStraight shape="BH-450x200x12x25" strength_main="SN490"/> 
    </StbSecSteelBeam_SRC_Shape> 
    </StbSecSteelFigureBeam_SRC> 
  </StbSecBeam_SRC>

## Page 326

6.8-3 
 
6.8.1. ＳＲＣ梁断面形状：StbSecFigureBeam_SRC 
・概要 
説明 ：ＳＲＣ梁断面の形状 
親要素：StbSecBeam_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
○ 
始端からの順番 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBeamStraight 
0 
1 
ストレート形状 
※(1) 
StbSecBeamTaper 
0 
1 
テーパー形状 
※ (1),
※(2) 
    
・補足 
(1) 要素がない場合は、その位置のコンクリートがないものとする。 
(2) StbSecBeamStraight で表現できないものにのみ使用する。

## Page 327

6.8-4 
 
6.8.2. ＳＲＣ梁断面配筋：StbSecBarArrangementBeam_SRC 
・概要 
説明 ：ＳＲＣ梁断面の配筋 
親要素：StbSecBeam_SRC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
○ 
始端からの順番 
 
D_bar_spacing 
string 
 
巾止筋：径 
 
strength_bar_spacing 
string 
 
巾止筋：鉄筋強度 
 
pitch_bar_spacing 
double 
 
巾止筋：ピッチ 
 
D_catch_bar 
string 
 
段取り筋 
 
strength_catch_bar 
string 
 
段取り筋強度 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarBeamSimple 
0 
1 
コンクリート梁断面配筋・簡
易 
※(1) 
StbSecBarBeamComplex 
0 
1 
コンクリート梁断面配筋・詳
細 
※(1) 
 
・補足 
(1) 要素がない場合はその位置の配筋がないものとする。

## Page 328

6.8-5 
 
6.8.3. ＳＲＣ梁断面鉄骨形状：StbSecSteelFigureBeam_SRC 
・概要 
説明 ：ＳＲＣ梁断面の形状 
親要素：StbSecBeam_SRC 
 
・属性 
無し 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBeam_SRC_Shape 
1 
制限なし 
 
 
StbSecSteelBeamWidening 
0 
2 
拡幅要素 
※(1) 
 
・補足 
(1) 拡幅する梁については、形状で表現するのではなく拡幅要素を定義する。

## Page 329

6.8-6 
 
6.8.3.1. SRC 梁断面任意鉄骨形状：StbSecSteelBeam_SRC_Shape 
・概要 
説明 ：Ｓ梁断面の形状（全部位同一寸法の場合） 
親要素：StbSecSteelFigureBeam_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
order 
integer 
○ 
左端からの順番 
※(1) 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBeamStraight 
0 
1 
ストレート 
※(1) 
StbSecSteelBeamTaper 
0 
1 
テーパー 
※(1), 
※(2) 
 
・補足 
(1) 要素がない場合は、鉄骨がないものとする。 
(2) StbSecSteelBeamStraight で表現できないものにのみ使用する。

## Page 330

6.9-1 
 
6.9. Ｓブレース断面：StbSecBrace_S 
・概要 
説明 ：Ｓブレース断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
※(1) 
floor 
string 
 
所属階 
部材リスト用 ※(2) 
kind_brace 
string 
 
ブレースの種別 
以下のいずれか 
VERTICAL（鉛直ブレース） 
HORIZONTAL（水平ブレース） 
※(3) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelFigureBrace_S 
1 
1 
Ｓブレース断面鉄骨形状 
 
 
・補足 
鉄骨断面の形状は、鉄骨断面要素 <StbSecSteel> で用意した断面形状を、この要素の子要素で参照す
ることで定義する。このとき <StbSecSteel> で用意した断面形状の「鉄骨断面の基準方向」を、部材の
断面軸（Ｚ軸）とする。 
 
(1)  「断面名称」は、部材リスト（構造図のブレース断面表）における、所属階を付けない名称を想
定している（所属階「1」と断面名称「V1」で「1V1」となる）。 
(2)  「所属階」は、部材の配置情報を検索しなくても部材リストが作成できるようにするための属性
という位置付けであり、省略された場合、所属する階が特定されない部材リスト名が作成される
こととなる。 
(3) 省略された場合は、VERTICAL とする。

## Page 331

6.9-2 
 
・例 
 
 
  <StbSecBrace_S id="56" name="V1" floor="R" kind_brace="VERTICAL"> 
    <StbSecSteelFigureBrace_S> 
      <StbSecSteelBrace_S_Same shape="H-100x100x6x8x8" strength_main="SN400"/> 
    </StbSecSteelFigureBrace_S> 
  </StbSecBrace_S>

## Page 332

6.9-3 
 
6.9.1. Ｓブレース断面鉄骨形状：StbSecSteelFigureBrace_S 
・概要 
説明 ：Ｓブレース断面の形状 
親要素：StbSecBrace_S 
 
・属性 
    無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBrace_S_Same 
1 
1 
Ｓブレース断面鉄骨形状 
・同一 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBrace_S_NotSame 
2 
2 
Ｓブレース断面鉄骨形状 
・頭脚部別 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelBrace_S_ThreeTypes 
3 
3 
Ｓブレース断面鉄骨形状 
・３種類 
 
 
・補足 
(1) 断面形状はハンチ状の変断面状態は想定しない。

## Page 333

6.9-4 
 
6.9.1.1. Ｓブレース断面鉄骨形状・同一：StbSecSteelBrace_S_Same 
・概要 
説明 ：Ｓブレース鉄骨断面の形状（全断面同一の場合） 
親要素：StbSecSteelFigureBrace_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  鉄骨断面要素 <StbSecSteel> の子要素の属性name「形状名」と一致する文字列を記述する。
ID による参照とはなっていないので、文字列を完全一致させるように注意を要する。 
(2)  省略された場合は、「鉄骨強度（主）」と同一とする。 
 
・例 
 
 
 
  <StbSecBrace_S id="56" name="V1" floor="R" kind_brace="VERTICAL"> 
    <StbSecSteelFigureBrace_S> 
      <StbSecSteelBrace_S_Same shape="H-100x100x6x8x8" strength_main="SN400"/> 
    </StbSecSteelFigureBrace_S> 
  </StbSecBrace_S>

## Page 334

6.9-5 
 
6.9.1.2. Ｓブレース断面鉄骨形状・頭脚部別：StbSecSteelBrace_S_NotSame 
・概要 
説明 ：Ｓブレース鉄骨断面の形状（頭部・脚部が別形状の場合） 
親要素：StbSecSteelFigureBrace_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 
以下のいずれか 
BOTTOM（脚部） 
TOP（頭部） 
 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
属性をそれぞれpos= “BOTTOM” およびpos= “TOP” とした子要素を各１回記述する。 
定義と補足内容 (1)～(2) は、「Ｓブレース断面鉄骨形状・同一」による。 
 
・例 
 
 
  <StbSecBrace_S id="56" name="V1" floor="R" kind_brace="VERTICAL"> 
    <StbSecSteelFigureBrace_S> 
      <StbSecSteelBrace_S_NotSame pos=”BOTTOM” 
shape="H-100x100x6x8x8" strength_main="SN400"/> 
      <StbSecSteelBrace_S_NotSame pos=”TOP” 
shape="BH-100x100x6x12" strength_main="SN400"/> 
    </StbSecSteelFigureBrace_S> 
  </StbSecBrace_S>

## Page 335

6.9-6 
 
6.9.1.3. Ｓブレース断面鉄骨形状・３種類：StbSecSteelBrace_S_ThreeTypes 
・概要 
説明 ：Ｓブレース鉄骨断面の形状（頭部・中央・脚部が別形状の場合） 
親要素：StbSecSteelFigureBrace_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配置位置 
以下のいずれか 
BOTTOM（脚部） 
CENTER（中央） 
TOP（頭部） 
 
shape 
string 
○ 
鉄骨形状 
※(1) 
strength_main 
string 
○ 
鉄骨強度（主） 
 
strength_web 
string 
 
鉄骨強度（ウェブ） 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
属性をpos= “BOTTOM” , pos= “CENTER” およびpos= “TOP” とした子要素を各１回記述する。 
定義と補足内容 (1)～(2) は、「Ｓブレース断面鉄骨形状・同一」による。 
 
・例 
  <StbSecBrace_S id="56" name="V1" floor="R" kind_brace="VERTICAL"> 
    <StbSecSteelFigureBrace_S> 
      <StbSecSteelBrace_S_ThreeTypes pos=”BOTTOM” 
shape="BH-100x100x6x12" strength_main="SN400"/> 
      <StbSecSteelBrace_S_ThreeTypes pos=”CENTER” 
shape="H-100x100x6x8x8" strength_main="SN400"/> 
      <StbSecSteelBrace_S_ThreeTypes pos=”TOP” 
shape="BH-100x100x6x12" strength_main="SN400"/> 
    </StbSecSteelFigureBrace_S> 
  </StbSecBrace_S>

## Page 336

6.10-1 
 
6.10. ＲＣスラブ断面：StbSecSlab_RC 
・概要 
説明 ：ＲＣスラブ断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
isFoundation 
boolean 
 
基礎スラブか否か 
部材リスト用※(1) 
isEarthen 
boolean 
 
土間か否か 
部材リスト用※(1) 
isCanti 
boolean 
 
片持ちスラブか否か 
部材リスト用※(1) 
strength_concrete 
string 
 
コンクリート強度 
※(2) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSlab_RC_Conventional 
0 
1 
在来スラブ 
※(3) 
StbSecSlab_RC_Truss 
0 
1 
トラス筋付きデッキスラブ 
※(3) 
 
・補足 
ここで記述する、基礎、土間、片持ちおよびハンチの情報は、部材リスト（構造図の断面表）の作成
上、分類のために用いることを想定しており、実際に形状が片持ちかどうかなどの指定は位置情報
<StbSlab>の属性による。 
 
(1)  省略された場合は、false とする。 
(2)  省略された場合は、参照する<StbSlab>の「節点ID 順序」StbNodeIdOrder の第１節点が所属
する<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した
<StbStory>のコンクリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリ
ート強度」strength_concrete をこの要素のコンクリート強度とする。 
(3)  設計段階でトラス筋付きデッキが在来型枠スラブと併記されることがあるからどちらも選択で
きるようにしている。二つとも指定することができる。

## Page 337

6.10-2 
 
・例 
  <StbSecSlab_RC id="120" name="S1" strength_concrete="FC21"> 
    <StbSecSlab_RC_Conventional> 
      <StbSecFigureSlab_RC_Conventional> 
        <StbSecSlab_RC_ConventionalStraight depth=”150”/> 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="MAIN_TOP" D="D13" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="MAIN_BOTTOM" D="D10" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="TRANSVERSE_TOP" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="TRANSVERSE_BOTTOM" D="D10" pitch="200"/> 
      </StbSecBarArrangementSlab_RC_Conventional> 
      <StbSecFormworkSlab> 
        <StbSecFormWorkSlabWood thickness=”12” type=”A”/> 
      </StbSecFormworkSlab> 
    </StbSecSlab_RC_Conventional> 
  </StbSecSlab_RC> 
 
  <StbSecSlab_RC id="117" name="CS1" isFoundation="false" isCanti="true"> 
    <StbSecSlab_RC_Conventional> 
      <StbSecFigureSlab_RC_Conventional> 
        <StbSecSlab_RC_ConventionalTaper base_depth="300" tip_depth="150"/> 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="MAIN_TOP" D="D13" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="MAIN_BOTTOM" D="D10" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="TRANSVERSE_TOP" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="TRANSVERSE_BOTTOM" D="D10" pitch="200"/> 
      </StbSecBarArrangementSlab_RC_Conventional> 
      <StbSecFormworkSlab> 
        <StbSecFormWorkSlabWood thickness=”12” type=”A”/> 
      </StbSecFormworkSlab> 
    </StbSecSlab_RC_Conventional> 
  </StbSecSlab_RC> 
 
  <StbSecSlab_RC id="130" name="FS1" isFoundation="true" isCanti="false" strength_concrete="FC24"> 
    <StbSecSlab_RC_Conventional> 
      <StbSecFigureSlab_RC_Conventional> 
        <StbSecSlab_RC_ConventionalStraight depth="300"/> 
      </StbSecFigureSlab_RC_Conventional> 
      （略） 
    </StbSecSlab_RC_Conventional> 
  </StbSecSlab_RC>

## Page 338

6.10-3 
 
6.10.1. ＲＣスラブ在来断面：StbSecSlab_RC_Conventional 
・概要 
説明 ：ＲＣスラブ在来断面 
親要素：StbSecSlab_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecFigureSlab_RC_Conventional 
1 
1 
RC 在来スラブ断面形状 
 
StbSecBarArrangementSlab_RC_Conventional 
0 
1 
RC 在来スラブ断面配筋 
 
StbSecFormworkSlab 
0 
1 
RC スラブ型枠 
 
 
・補足 
    無し

## Page 339

6.10-4 
 
6.10.1.1. ＲＣ在来スラブ断面形状：StbSecFigureSlab_RC_Conventional 
・概要 
説明 ：ＲＣ在来スラブ断面の形状 
親要素：StbSecSlab_RC_Conventional 
 
・属性 
無し 
 
・内容 
無し 
 
    
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecSlab_RC_ConventionalStraight 
1 
1 
ＲＣ在来スラブ断面形状・ストレート 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecSlab_RC_ConventionalTaper 
1 
1 
ＲＣ在来スラブ断面形状・テーパー 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecSlab_RC_ConventionalHaunch 
1 
1 
ＲＣ在来スラブ断面形状・ハンチ 
 
 
・補足 
子要素<StbSecSlab_RC_Taper> は、親要素<StbSecSlab_RC>の属性「片持ちスラブか否か」がtrue
のときのみ適用する。

## Page 340

6.10-5 
 
6.10.1.1.1. ＲＣ在来スラブ断面形状・ストレート：StbSecSlab_RC_ConventionalStraight 
・概要 
説明 ：ＲＣ在来スラブ断面の形状（全部位同一厚さの場合） 
親要素：StbSecFigureSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth 
double 
○ 
厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「厚さ」の定義は下図による。 
 
・例 
 
 
厚さ
  <StbSecSlab_RC id="120" name="S1" （略）> 
    < StbSecSlab_RC_Conventional> 
<StbSecFigureSlab_RC_Conventional > 
        <StbSecSlab_RC_ConventionalStraight depth="150"/> 
      </StbSecFigureSlab_RC_Conventional> 
 
      <StbSecBarArrangementSlab_RC_Conventional> 
        （略） 
      </StbSecBarArrangementSlab_RC_Conventional > 
<StbSecFormworkSlab> 
        <StbSecFormWorkSlabWood thickness=”12” type=”A”/> 
      </StbSecFormworkSlab> 
    </StbSecSlab_RC_Conventional> 
  </StbSecSlab_RC>

## Page 341

6.10-6 
 
6.10.1.1.2. ＲＣ在来スラブ断面形状・テーパー：StbSecSlab_RC_ConventionalTaper 
・概要 
説明 ：ＲＣ在来スラブ断面の形状（片持ちスラブで、根元、先端が異なる厚さの場合） 
親要素：StbSecFigureSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
base_depth 
double 
○ 
根元厚さ 
 
tip_depth 
double 
○ 
先端厚さ 
 
tip_offset 
 
 
オフセット平面からの下がり寸法 
※(1) 
 
・内容、子要素 
無し 
 
・補足 
根元および先端の「厚さ」の定義は下図による。 
(1) 記入がない場合は0 とする。 
(2) 片持ちスラブの先端は、StbNode のkind で判別する。 
先端厚さ
オフセット平面からの下がり寸法
根元厚さ
 
・例 
 
 
  <StbSecSlab_RC id="117" name="CS1" isFoundation="false" isCanti="true">    
<StbSecSlab_RC_Conventional> 
      <StbSecFigureSlab_RC_Conventional> 
        <StbSecSlab_RC_ConventionalTaper base_depth="300" tip_depth="150"/> 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional> 
              (略） 
      </StbSecBarArrangementSlab_RC_Conventional> 
       （略） 
    </StbSecSlab_RC_Conventional> 
  </StbSecSlab_RC>

## Page 342

6.10-7 
 
6.10.1.1.3. ＲＣ在来スラブ断面形状・ハンチ：StbSecSlab_RC_ConventionalHaunch 
・概要 
説明 ：ＲＣ在来スラブ断面の形状（元端にハンチがあり、根元が中央と異なる厚さの場合） 
親要素：StbSecFigureSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
base_depth 
double 
〇 
根元厚さ 
 
tip_depth 
double 
〇 
中央厚さ 
 
haunch_length 
double 
〇 
ハンチ長さ 
※(1) 
tip_offset 
double 
 
オフセット平面からの下がり寸法 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
根元および中央の「厚さ」、および「ハンチ長さ」の定義は下図による。ハンチは、全ての支持辺につ
いて同一形状であるものとする。 
ハンチの形状に応じて、属性をそれぞれpos= “BASE”、pos= “CENTER” およびpos= “HAUNCH” 
としたこの子要素を各１回記述する。 
(1) ハンチ長さ起点はオフセット位置からとする。 
(2) 記入がない場合は0 とする 
中央厚さ
オフセット平面からの下がり寸法
根元厚さ
ハンチ長さ

## Page 343

6.10-8 
 
・例 
 
 
 
  <StbSecSlab_RC id=”117” name=”S1” （略）> 
   <StbSecSlab_RC_Conventional> 
      <StbSecFigureSlab_RC_Conventional> 
        <StbSecSlab_RC_ConventionalHaunch base_depth=”300” tip_depth=”150” haunch_length=”1000”/> 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional> 
              (略） 
      </StbSecBarArrangementSlab_RC_Conventional> 
       （略） 
    </StbSecSlab_RC_Conventional> 
  </StbSecSlab_RC>

## Page 344

6.10-9 
 
6.10.1.2. ＲＣ在来スラブ断面配筋：StbSecBarArrangementSlab_RC_Conventional 
・概要 
説明 ：ＲＣ在来スラブ断面の配筋 
親要素：StbSecSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_top 
double 
 
かぶり厚さ（上） 
※(1) 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlab_RC_ConventionalStandard 
12 
12 
ＲＣ在来スラブ断面配筋・標準 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlab_RC_Conventional2Way 
4 
4 
ＲＣ在来スラブ断面配筋・２方
向 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlab_RC_Conventional1Way1 
4 
4 
ＲＣ在来スラブ断面配筋・１方
向１ 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlab_RC_Conventional1Way2 
6 
6 
ＲＣ在来スラブ断面配筋・１方
向２ 
※(2) 
 
上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlab_RC_Open 
1 
6 
スラブ開口配筋 
※(3) 
 
・補足

## Page 345

6.10-10 
 
(1)  「かぶり厚さ（…）」が省略された場合の扱いは、<StbApplyConditionsList>の補足説明によ
る。 
(2)  「標準」「２方向」は、スラブの平面形状に応じて、短辺・長辺方向に配筋する場合、「１方向」
は、主筋方向を指定して配筋する場合に用いる。 
(3) 子要素<StbSecBarSlab_RC_Open>によるスラブ開口配筋は、スラブ断面に共通の補強筋で、該当
する補強筋が存在する場合に、追加で記述する。ただし、開口ごとに配筋が異なる場合は、この要
素ではなく、<StbSecOpen_RC>のスラブ開口配筋を用いる。

## Page 346

6.10-11 
 
6.10.1.2.1. ＲＣ在来スラブ断面配筋・標準：StbSecBarSlab_RC_ConventionalStandard 
・概要 
説明 ：ＲＣ在来スラブ断面の配筋（標準） 
親要素：StbSecBarArrangementSlab_RC_Conventional 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 
以下のいずれか 
SHORT_TOP_COLUMN 
（①短辺上端柱列帯） 
SHORT_TOP_MID_END 
（②短辺上端柱間帯端部） 
SHORT_TOP_MID_CENTER 
（③短辺上端柱間帯中央） 
SHORT_BOTTOM_COLUMN 
（④短辺下端柱列帯） 
SHORT_BOTTOM_MID_END 
（⑤短辺下端柱間帯端部） 
SHORT_BOTTOM_MID_CENTER 
（⑥短辺下端柱間帯中央） 
LONG_TOP_COLUMN 
（⑦長辺上端柱列帯） 
LONG_TOP_MID_END 
（⑧長辺上端柱間帯端部） 
LONG_TOP_MID_CENTER 
（⑨長辺上端柱間帯中央） 
LONG_BOTTOM_COLUMN 
（⑩長辺下端柱列帯） 
LONG_BOTTOM_MID_END 
（⑪長辺下端柱間帯端部） 
LONG_BOTTOM_MID_CENTER 
（⑫長辺下端柱間帯中央） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ

## Page 347

6.10-12 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。部分的な回数の省略
はできない。 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。

## Page 348

6.10-13 
 
・例 
 
 
 
  <StbSecSlab_RC id="121" name="S2" （略）> 
    <StbSecSlab_RC_Conventional> 
      <StbSecFigureSlab_RC_Conventional>        （略） 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional> 
 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="SHORT_TOP_COLUMN" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="SHORT_TOP_MID_END" D="D10D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="SHORT_TOP_MID_CENTER" D="D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="SHORT_BOTTOM_COLUMN" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="SHORT_BOTTOM_MID_END" D="D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="SHORT_BOTTOM_MID_CENTER" D="D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="LONG_TOP_COLUMN" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="LONG_TOP_MID_END" D="D10D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="LONG_TOP_MID_CENTER" D="D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="LONG_BOTTOM_COLUMN" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="LONG_BOTTOM_MID_END" D="D13" pitch="200"/> 
        <StbSecBarSlab_RC_ConventionalStandard  
pos="LONG_BOTTOM_MID_CENTER" D="D13" pitch="200"/> 
      </StbSecBarArrangementSlab_RC_Conventional>    </StbSecSlab_RC_Conventional> 
 
  </StbSecSlab_RC>

## Page 349

6.10-14 
 
6.10.1.2.2. ＲＣ在来スラブ断面配筋・２方向：StbSecBarSlab_RC_Conventional2Way 
・概要 
説明 ：ＲＣ在来スラブ断面の配筋（２方向） 
親要素：StbSecBarArrangementSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 
以下のいずれか 
SHORT_TOP_END 
（①短辺上端端部） 
SHORT_BOTTOM_END 
（②短辺下端端部） 
SHORT_TOP_CENTER 
（③短辺上端中央） 
SHORT_BOTTOM_CENTER 
（④短辺下端中央） 
LONG_TOP_END 
（⑤長辺上端端部） 
LONG _BOTTOM_END 
（⑥長辺下端端部） 
LONG_TOP_CENTER 
（⑦長辺上端中央） 
LONG_BOTTOM_CENTER 
（⑧長辺下端中央） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 350

6.10-15 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。部分的な回数の省略
はできない。 
 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
 
・例

## Page 351

6.10-16 
 
6.10.1.2.3. ＲＣ在来スラブ断面配筋・１方向１：StbSecBarSlab_RC_Conventional1Way1 
・概要 
説明 ：ＲＣ在来スラブ断面の配筋（１方向、主筋方向の配筋が全断面同一の場合） 
親要素：StbSecBarArrangementSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 
以下のいずれか 
MAIN_TOP 
（①主筋方向上端） 
MAIN_BOTTOM 
（②主筋方向下端） 
TRANSVERSE_TOP 
（③配力筋方向上端） 
TRANSVERSE_BOTTOM 
（④配力筋方向下端） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 352

6.10-17 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。部分的な回数の省略
はできない。 
 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
  <StbSecSlab_RC id="117" name="CS1" （略）> 
    <StbSecSlab_RC_Conventional> 
<StbSecFigureSlab_RC_Conventional> 
        （略） 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional> 
 
        <StbSecBarSlab_RC_Conventional1Way1 pos="MAIN_TOP" D="D13" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="MAIN_BOTTOM" D="D10" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="TRANSVERSE_TOP" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_Conventional1Way1 pos="TRANSVERSE_BOTTOM" D="D10" pitch="20"/> 
      </StbSecBarArrangementSlab_RC_Conventional> 
    </StbSecSlab_RC_Conventional> 
 
  </StbSecSlab_RC>

## Page 353

6.10-18 
 
6.10.1.2.4. ＲＣ在来スラブ断面配筋・１方向２：StbSecBarSlab_RC_Conventional1Way2 
・概要 
説明 ：ＲＣ在来スラブ断面の配筋（１方向、主筋方向の根元と先端が別配筋の場合） 
親要素：StbSecBarArrangementSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 
以下のいずれか 
MAIN_BASE_TOP 
（①主筋方向根元上端） 
MAIN_BASE_BOTTOM 
（②主筋方向根元下端） 
MAIN_TIP_TOP 
（③主筋方向先端上端） 
MAIN_TIP_BOTTOM 
（④主筋方向先端下端） 
TRANSVERSE_TOP 
（⑤配力筋方向上端） 
TRANSVERSE_BOTTOM 
（⑥配力筋方向下端） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 354

6.10-19 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。部分的な回数の省略
はできない。 
 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
  <StbSecSlab_RC id="122" name="CS2" （略）> 
    <StbSecSlab_RC_Conventional> 
<StbSecFigureSlab_RC_Conventional> 
        （略） 
      </StbSecFigureSlab_RC_Conventional> 
      <StbSecBarArrangementSlab_RC_Conventional 
        <StbSecBarSlab_RC_Conventional1Way2 pos="MAIN_BASE_TOP" D="D13" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way2 pos="MAIN_BASE_BOTTOM" D="D10" pitch="100"/> 
        <StbSecBarSlab_RC_Conventional1Way2 pos="MAIN_TIP_TOP" D="D13" pitch="200"/> 
        <StbSecBarSlab_RC_Conventional1Way2 pos="MAIN_TIP_BOTTOM" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_Conventional1Way2 pos="TRANSVERSE_TOP" D="D10" pitch="200"/> 
        <StbSecBarSlab_RC_Conventional1Way2 pos="TRANSVERSE_BOTTOM" D="D10" pitch="200"/> 
      </StbSecBarArrangementSlab_RC_Conventional> 
    </StbSecSlab_RC_Conventional> 
 
  </StbSecSlab_RC> 
①主筋方向根元上端 
②主筋方向根元下端 
③主筋方向先端上端 
④主筋方向先端下端 
⑤配力筋方向上端 
⑥配力筋方向下端

## Page 355

6.10-20 
 
6.10.1.2.5. スラブ開口配筋：StbSecBarSlab_RC_Open 
・概要 
説明 ：ＲＣスラブ断面の配筋（スラブ断面符号ごとの開口部配筋） 
親要素：StbSecBarArrangementSlab_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
X_TOP（X 方向上端） 
X_BOTTOM（X 方向下端） 
Y_TOP（Y 方向上端） 
Y_BOTTOM（Y 方向下端） 
DIAGONAL_TOP（斜め方向上端） 
DIAGONAL_BOTTOM（斜め方向下端） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
※(2) 
length 
double 
 
長さ 
 
 
・内容、子要素 
無し 
 
・補足 
必要に応じて、補強筋のある位置ごとにこの子要素を各１回記述する。補強筋がその位置にない場合
はこの要素を省略してよい。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
(2)  「本数」は、１辺および１隅あたりの合計本数とする。 
 
 
（1点目）
（2点目）
（3点目）
（4点目）
開口
X
Y
Y方向上端（下端）開口補強筋
X方向上端（下端）開口補強筋
斜め方向上端（下端）開口補強筋

## Page 356

6.10-21 
 
6.10.1.3. 型枠要素：StbSecFormworkSlab 
・概要 
説明 ：型枠要素 
親要素：StbSecSlab_RC_Conventional 
 
・属性 
無し 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFormworkSlabWood 
1 
1 
合板型枠 
※(2) 
または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFormworkSlabFlatDeck 
1 
1 
フラットデッキ 
※(2) 
 
・補足

## Page 357

6.10-22 
 
6.10.1.3.1. 在来型枠：StbSecFormworkSlabWood 
・概要 
説明 ：合板型枠 
親要素：StbSecFormworkSlab 
 
・概要 
説明 ：  
親要素： 
 
・属性 
属性名 
型 
必須 
説明 
補足 
thickness 
double 
 
木の厚さ 
 
type 
string 
 
せき板 
以下のいずれかの値をとる 
A：A 種 
B：B 種 
C：C 種 
 
 
・内容、子要素 
無し 
 
・補足

## Page 358

6.10-23 
 
6.10.1.3.2. フラットデッキ型枠：StbSecFormworkSlabFlatDeck 
・概要 
説明 ：フラットデッキ型枠 
親要素：StbSecFormworkSlab 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFormworkSlabFlatDeckSpec 
1 
1 
仕様指定 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFormworkSlabFlatDeckProduct 
1 
1 
型番指定 
 
 
・補足

## Page 359

6.10-24 
 
6.10.1.3.2.1. フラットデッキ型枠仕様指定：StbSecFormworkSlabFlatDeckSpec 
・概要 
説明 ：フラットデッキ型枠仕様指定 
親要素：StbSecFormworkSlabFlatDeck 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth 
double 
〇 
デッキ厚 
 
surface_finishing 
double 
 
表面処理(Z12 など) 
 
 
・内容、子要素 
無し 
 
・補足 
 
デッキ厚

## Page 360

6.10-25 
 
6.10.1.3.2.2. フラットデッキ型番指定：StbSecFormworkSlabFlatDeckProduct 
・概要 
説明 ：フラットデッキ型番指定 
親要素：StbSecFormworkSlabFlatDeck 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
〇 
製品型番 
※(1) 
release_time 
string 
〇 
リリース時期 
※(1) 
surface_finishing 
double 
 
表面処理(Z12 など) 
 
 
・内容 
    無し 
 
・子要素 
無し 
 
・補足 
(1) 製品型番一覧表のフラットデッキカテゴリから「製品型番」と「リリース時期」の2 つのキー
で製品を一意に指定する。

## Page 361

6.10-26 
 
6.10.2. ＲＣスラブトラス断面：StbSecSlab_RC_Truss 
・概要 
説明 ：ＲＣスラブ断面 
親要素：StbSections 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSlab_RC_TrussProduct 
1 
制限なし 
トラス筋付きデッキプレート
製品 
※(1) 
 
・補足 
(1) 回数が2 以上の場合は、同等品として複数の製品を指定したものとする。

## Page 362

6.10-27 
 
6.10.2.1. トラス筋付きデッキスラブ製品：StbSecSlab_RC_TrussProduct 
・概要 
説明 ：トラス筋付きデッキスラブ製品 
親要素：StbSecSlab_RC_Truss 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
〇 
製品型番 
※(1) 
release_time 
string  
〇 
リリース時期 
※(1) 
thickness 
double  
〇 
コンクリート厚 
 
surface_finishing 
string 
 
表面処理(Z12 など) 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecBarArrangementSlab_RC_Truss 
0 
1 
トラス筋付きデッキスラブ断面の配
筋 
 
 
・補足 
(1) 製品型番一覧表のトラス筋付きデッキカテゴリから「製品型番」と「リリース時期」の2 つの
キーで製品を一意に指定する。 
 
かぶり厚さ(上)
かぶり厚さ(下)
コンクリート厚

## Page 363

6.10-28 
 
6.10.2.1.1. トラス筋付きデッキスラブ断面配筋：StbSecBarArrangementSlab_RC_Truss 
・概要 
説明 ：RC スラブ(トラス筋付きデッキスラブ)断面の配筋 
親要素：StbSecSlab_RC_TrussProduct 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_top 
double 
 
かぶり厚さ（上） 
※(1) 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlab_RC_Truss1Way 
2 
4 
トラス筋付きデッキスラブ断
面配筋・１方向 
 
 
・補足 
(1) 「かぶり厚さ（…）」が省略された場合の扱いは、<StbApplyConditionsList>の補足説明による。

## Page 364

6.10-29 
 
6.10.2.1.1.1. トラス筋付きデッキスラブ断面配筋・１方向：StbSecBarSlab_RC_Truss1Way 
・概要 
説明 ：デッキプレートスラブ断面の配筋（１方向） 
親要素：StbSecBarArrangementSlab_RC_Truss 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 
以下のいずれか 
MAIN_TOP 
（①主筋方向上端） 
MAIN_BOTTOM 
（②主筋方向下端） 
TRANSVERSE_TOP 
（③配力筋方向上端） 
TRANSVERSE_BOTTOM 
（④配力筋方向下端） 
※(2) 
strength 
string 
 
鉄筋強度 
※(3) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) デッキプレートの主方向は、配筋の主筋方向と一致するものとする。 
(2) 「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。ただし、配筋
が１段の場合は、②および④を省略する。 
(3) 「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にあ
る場合は、省略してもよい。

## Page 365

6.10-30

## Page 366

6.11-1 
 
6.11. デッキ合成スラブ断面：StbSecSlabDeck 
・概要 
説明 ：デッキ合成スラブ断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
strength_concrete 
string 
 
コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSlabDeckProduct 
1 
制限なし 
合成デッキ製品 
※(2) 
 
・補足 
(1)  省略された場合は、参照する<StbSlab>の「節点ID 順序」StbNodeIdOrder の第１節点が所属
する<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した
<StbStory>のコンクリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリ
ート強度」strength_concrete をこの要素のコンクリート強度とする。 
(2)  回数が2 以上の場合は、同等品として複数の製品を指定したものとする。

## Page 367

6.11-2 
 
・例 
 
 
 
  <StbSecSlabDeck id="117" name="DS1" strength_concrete="FC21"> 
    <StbSecSlabDeckProduct product_code="XXXX" release_time="2023" top_concrete="80"> 
      <StbSecBarArrangementSlabDeck> 
        <StbSecBarSlabDeck1Way pos="MAIN_TOP" D="D13" pitch="200"/> 
        <StbSecBarSlabDeck1Way pos="TRANSVERSE_TOP" D="D13" pitch="200"/> 
      </StbSecBarArrangementSlabDeck> 
    </StbSecSlabDeckProduct> 
  </StbSecSlabDeck>

## Page 368

6.11-3 
 
6.11.1. 合成デッキ製品：StbSecSlabDeckProduct 
・概要 
説明 ：合成デッキ製品 
親要素：StbSecSlabDeck 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
〇 
製品型番 
※(1) 
release_time 
string  
〇 
リリース時期 
※(1) 
top_concrete 
double  
〇 
トップコンクリート 
 
surface_finishing 
string 
 
表面処理(Z12 など) 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarArrangementSlabDeck 
0 
1 
デッキ合成スラブ断面の配筋 
 
StbCertificationNumber 
0 
制限なし 
認定番号 
※(2) 
 
・補足 
(1) 製品型番一覧表の合成デッキカテゴリから「製品型番」と「リリース時期」の2 つのキーで製
品を一意に指定する。 
(2) 該当する耐火認定番号が複数ある場合は列挙することができる。

## Page 369

6.11-4 
 
  
6.11.1.1. デッキ合成スラブ断面配筋：StbSecBarArrangementSlabDeck 
・概要 
説明 ：デッキ合成スラブ断面の配筋 
親要素：StbSecSlabDeckProduct 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_top 
double 
 
かぶり厚さ（上） 
※(1) 
depth_cover_bottom 
double 
 
かぶり厚さ（下） 
D_refactory_bar 
string 
 
耐火補強筋径 
※(2) 
strength_refactory_bar 
string 
 
耐火補強筋強度 
※(2)、※(3) 
 
・内容 
無し 
 
・子要素 
   要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlabDeck1Way 
1 
2 
デッキ合成スラブ断面配筋・１
方向 
 
 
・補足 
(1) 「かぶり厚さ（…）」が省略された場合の扱いは、<StbApplyConditionsList>の補足説明による。 
(2) 耐火補強筋が必要な場合に指定する。記入された場合は各溝に1 本耐火補強筋を配筋するものと
する。 
(3) 「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある場
合は、省略してもよい

## Page 370

6.11-5 
 
6.11.1.1.1. デッキ合成スラブ断面配筋・１方向：StbSecBarSlabDeck1Way 
・概要 
説明 ：デッキ合成スラブ断面の配筋（１方向） 
親要素：StbSecBarArrangementSlabDeck 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 
以下のいずれか 
MAIN_TOP 
（①主筋方向上端） 
TRANSVERSE_TOP 
（②配力筋方向上端） 
MESH 
（③溶接金網） 
※(2),※(3) 
strength 
string 
 
鉄筋強度 
※(4) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) デッキ合成スラブ用デッキプレートの主方向は、配筋の主筋方向と一致するものとする。 
(2) ひび割れ補強筋を鉄筋で配筋する場合は①、②を使用し、溶接金網を使用する場合は③を使用
する。溶接金網の場合、ピッチは網目と読み替える。 
(3) 「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にあ
る場合は、省略してもよい。

## Page 371

6.11-6 
 
6.11.1.2. 認定番号：StbCertificationNumber 
・概要 
説明 ：認定番号 
親要素：StbSecSlabDeckProduct、StbSecPile_RC_Certified、 
StbSecPile_S_Certified、 StbPilePrecastCertified 
・属性 
属性名 
型 
必須 
説明 
補足 
certification_number 
string 
〇 
認定番号 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 372

6.12-1 
 
6.12. 既製スラブ断面：StbSecSlabPrecast 
・概要 
説明 ：既製スラブ断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
precast_type 
string 
○ 
工法種別  
以下のいずれか 
FULL（フルPC工法） 
HALF（ハーフPC 工法） 
FORM（型枠利用） 
 
strength_top_concrete 
string 
 
トップ部分コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecFigureSlabPrecast 
0 
1 
既製スラブトップ部分断面形
状 
※
(1) 
StbSecBarArrangementSlabPrecast 
0 
1 
既製スラブトップ部分断面配
筋 
※
(2) 
StbSecProductSlabPrecast 
1 
1 
既製スラブ製品 
 
 
・補足 
(1)  「工法種別」がFULL（フルPC 工法）の場合は、「トップ部分コンクリート強度」および子要
素 <StbSecFigureSlabPrecast> <StbSecBarArrangementSlabPrecast> を記述してはならない。
HALF（ハーフPC 工法）、FORM（型枠利用）の場合は記述するが、「トップ部分コンクリート
強度」は省略した場合、参照する<StbSlab>の「節点ID 順序」StbNodeIdOrder の第１節点が所
属する<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した

## Page 373

6.12-2 
 
<StbStory>のコンクリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリ
ート強度」strength_concrete をこの要素のコンクリート強度とする。 
(2)  記述が必要な場合で、子要素<StbSecBarArrangementSlabPrecast>の回数が0 となる場合は、
鉄筋を扱わないプログラムが一時的に作成する場合を想定しており、無筋であることを示すもの
ではない。 
 
・例 
 
 
 
  <StbSecSlabPrecast id="119" name="S1" precast_type="HALF" strength_concrete="FC21"> 
    <StbSecFigureSlabPrecast> 
      <StbSecSlabPrecastStraight depth_top_concrete="100"/> 
    </StbSecFigureSlabPrecast> 
    <StbSecBarArrangementSlabPrecast> 
      <StbSecBarSlabPrecast2Way pos="SHORT_TOP" D="D13" pitch="200"/> 
      <StbSecBarSlabPrecast2Way pos="LONG_TOP" D="D13" pitch="200"/> 
    </StbSecBarArrangementSlabPrecast> 
    <StbSecProductSlabPrecast product_code="XXXX" release_time="XXXX"/> 
  </StbSecSlabPrecast>

## Page 374

6.12-3 
 
6.12.1. 既製スラブトップ部分断面形状：StbSecFigureSlabPrecast 
・概要 
説明 ：既製スラブトップ部分断面の形状 
親要素：StbSecSlabPrecast 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSlabPrecastStraight 
1 
1 
既製スラブトップ部分断面形状・ストレート 
 
 
・補足 
 
6.12.1.1. 既製スラブトップ部分断面形状・ストレート：StbSecSlabPrecastStraight 
・概要 
説明 ：既製スラブトップ部分の断面の形状 
親要素：StbSecFigureSlabPrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_top_concrete 
double 
○ 
トップコンクリート厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
 
既成スラブ 
トップコンクリートの厚さ 
既成スラブトップ部分

## Page 375

6.12-4 
 
  
6.12.2. 既製スラブトップ部分断面配筋：StbSecBarArrangementSlabPrecast 
・概要 
説明 ：既製スラブトップ部分断面の配筋 
親要素：StbSecSlabPrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_top 
double 
 
かぶり厚さ（上） 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlabPrecastStandard 
12 
12 
既製スラブトップ部分断面配筋・標準 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlabPrecast2Way 
4 
4 
既製スラブトップ部分断面配筋・２方向 
※(2) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarSlabPrecast1Way 
2 
5 
既製スラブトップ部分断面配筋・１方向 
※(2) 
 
・補足 
(1)  「かぶり厚さ（上）」が省略された場合の扱いは、<StbApplyConditionsList>の補足説明によ
る。 
(2)  「標準」「２方向」は、既製スラブの方向に関係なく、コンクリート部分の配筋においてスラブ
の平面形状に応じて、短辺・長辺方向に配筋する場合に用い、「１方向」は、既製スラブの主方向
を指定する場合に用いる。

## Page 376

6.12-5 
 
6.12.2.1. 既製スラブ断面配筋・標準：StbSecBarSlabPrecastStandard 
・概要 
説明 ：既製スラブトップ部分断面の配筋（標準） 
親要素：StbSecBarArrangementSlabPrecast 
 
属性、内容、補足事項とも「ＲＣスラブ断面配筋・標準：StbSecBarSlab_RC_ConventionalStandard」
と同じである。 
 
6.12.2.2. 既製スラブ断面配筋・２方向：StbSecBarSlabPrecast2Way 
・概要 
説明 ：既製スラブトップ部分断面の配筋（２方向） 
親要素：StbSecBarArrangementSlabPrecast 
 
属性、内容、補足事項とも「ＲＣスラブ断面配筋・２方向：StbSecBarSlab_RC_Conventional2Way」
と同じである。 
 
6.12.2.3. 既製スラブ断面配筋・１方向：StbSecBarSlabPrecast1Way 
・概要 
説明 ：既製スラブトップ部分断面の配筋（１方向） 
親要素：StbSecBarArrangementSlabPrecast 
 
属性、内容、補足事項とも「デッキスラブ断面配筋・１方向：StbSecBarSlabDeck1Way」と同じであ
る。ただし、「デッキプレート」「デッキプレートスラブ」とある箇所は「既製スラブ」と読み替える。

## Page 377

6.12-6 
 
6.12.3. 既製スラブ製品：StbSecProductSlabPrecast 
・概要 
説明 ：既製スラブ断面の既製（プレキャストコンクリート製品）部分 
親要素：StbSecSlabPrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
〇 
製品型番 
※(1) 
release_time 
string  
〇 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  製品型番一覧表の既成スラブカテゴリから「製品型番」と「リリース時期」の2 つのキーで製品
を一意に指定する。

## Page 378

6.13-1 
 
6.13. 荷重用スラブ断面：StbSecSlabLoad 
・概要 
説明 ：荷重用スラブ断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
一貫構造計算プログラムなどにおける床・屋根の指定で、折板屋根などコンクリート部分がない場合
は、仕上げ相当の荷重値のみで考慮する場合がある。ＲＣスラブ断面StbSecSlab_RC の子要素では、コ
ンクリート部分の厚さdepth の値を0 より大きい値で必須と規定しているため、このような場合は、本
要素にて指定する。仕上げ相当の荷重値の指定は「計算データ編：StbCalSlabFinish_RC」による。 
 
・例 
 
  <StbSecSlabLoad id="130" name="Load1"/>

## Page 379

6.14-1 
 
6.14. ＲＣ壁断面：StbSecWall_RC 
・概要 
説明 ：ＲＣ壁断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
strength_concrete 
string 
 
コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureWall_RC 
1 
1 
ＲＣ壁断面形状 
 
StbSecBarArrangementWall_RC 
0 
1 
ＲＣ壁断面配筋 
※(2) 
 
・補足 
(1) 省略された場合は、参照する<StbWall>の「節点ID 順序」StbNodeIdOrder の最終節点が所属す
る<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した<StbStory>のコン
クリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリート強度」
strength_concrete をこの要素のコンクリート強度とする。 
(2) 子要素<StbSecBarArrangementWall_RC>の回数が0 となる場合は、鉄筋を扱わないプログラ
ムが一時的に作成する場合を想定しており、無筋であることを示すものではない。 
 
・例 
 
  <StbSecWall_RC id="93" name="EW16" strength_concrete="FC36"> 
    <StbSecFigureWall_RC> 
      <StbSecWall_RC_Straight t="160"/> 
    </StbSecFigureWall_RC> 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 380

6.14-2 
 
6.14.1. ＲＣ壁断面形状：StbSecFigureWall_RC 
・概要 
説明 ：ＲＣ壁断面の形状 
親要素：StbSecWall_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecWall_RC_Straight 
1 
1 
ＲＣ壁断面形状・ストレート 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecWall_RC_Taper 
1 
1 
ＲＣ壁断面形状・テーパー 
 
 
・補足

## Page 381

6.14-3 
 
6.14.1.1.  ＲＣ壁断面形状・ストレート：StbSecWall_RC_Straight 
・概要 
説明 ：ＲＣ壁断面の形状（全部位同一厚さの場合） 
親要素：StbSecFigureWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
t 
double 
○ 
厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「厚さ」の定義は下図による。 
 
・例 
 
 
 
Ｚ
Y
厚さ
  <StbSecWall_RC id="93" name="EW16" （略）> 
    <StbSecFigureWall_RC> 
      <StbSecWall_RC_Straight t="160"/> 
    </StbSecFigureWall_RC> 
    <StbSecBarArrangementWall_RC （略）> 
        （略） 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 382

6.14-4 
 
6.14.1.2.  ＲＣ壁断面形状・テーパー：StbSecWall_RC_Taper 
・概要 
説明 ：ＲＣ壁断面の形状（ＲＣ壁の下端と上端で厚さが異なる場合） 
親要素：StbSecFigureWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
t_bottom 
double 
○ 
RC 壁下端の厚さTb 
※(1) 
t_top 
double 
〇 
RC 壁上端の厚さTt 
※(1) 
depth_Hb 
double 
 
形状の切り替え位置Hb 
※(1), ※(2) 
depth_Ht 
double 
 
形状の切り替え位置Ht 
※(1), ※(2) 
type_straight 
string 
 
ストレートな面 
以下のいずれかの値をとる。 
OUTSIDE（外側） 
INSIDE（内側） 
※(3), ※(4) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 各寸法の定義は右図による。 
部材座標系Z 方向の基準点位置はRC 壁下端の板厚中央とする。 
形状の切り替え位置を示すdepth_Hb、 depth_Ht は基準点からの
距離とする。 
(2) 省略した時は0 とする。 
(3) type_straight の外側・内側と部材座標系の関係は
<StbWall>のtype_outside で指定する。 
例として、<StbWall>のtype_outside が
TYPE_MINUS の場合の、type_straight による違
いを右図に示す。 
(4) type_straight を省略した場合はOUTSIDE とする。 
t_top 
t_bottom 
depth_Ht 
depth_Hb 
基準点 
基準点 
Ｚ
Y
Ｚ
Y
Ｚ
Y
OUTSIDE           INSIDE 
type_straight 
外 
外

## Page 383

6.14-5 
 
・例 
 
 
 
  <StbSecWall_RC id="93" name="EW16" （略）> 
    <StbSecFigureWall_RC> 
      <StbSecWall_RC_Taper t_bottom="300" t_top="150" depth_Hb="1000"  
depth_Ht="500" type_straight="OUTSIDE"/> 
    </StbSecFigureWall_RC> 
    <StbSecBarArrangementWall_RC （略）> 
        （略） 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 384

6.14-6 
 
6.14.2. ＲＣ壁断面配筋：StbSecBarArrangementWall_RC 
・概要 
説明 ：ＲＣ壁断面の配筋 
親要素：StbSecWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_outside 
double 
 
かぶり厚さ（内外異なるときは外） 
※(1) 
depth_cover_inside 
double 
 
内外異なるとき、内のかぶり厚さ 
outer_bar_direction 
string 
 
外側に配置する鉄筋 
以下のいずれかの値をとる。 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
※(2) 
D_bar_spacing 
string 
 
幅止め筋：径 
※(3) 
strength_bar_spacing 
string 
 
幅止め筋：鉄筋強度 
※(3) 
pitch_bar_spacing 
double 
 
幅止め筋：ピッチ 
※(3), ※(4) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_Single 
2 
2 
ＲＣ壁断面配筋・シングル 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_Zigzag 
2 
2 
ＲＣ壁断面配筋・千鳥 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_DoubleNet 
2 
2 
ＲＣ壁断面配筋・ダブル 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_InsideAndOutside 
4 
制限なし 
ＲＣ壁断面配筋(内外異なる) 
※(5)

## Page 385

6.14-7 
 
上記子要素に、以下を追加してもよい。 
  要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_Edge 
1 
4 
壁端部補強筋 
※(6) 
StbSecBarWall_RC_Open 
1 
3 
壁開口配筋 
※(7) 
 
・補足 
(1)  かぶり厚さの定義は、下図による。配筋がシングルのときは省略してよい。それ以外で省略され
た場合の扱いは、<StbApplyConditionsList>の補足説明による。 
 
(2) 省略された場合は、縦筋を外側に配置する鉄筋とする。 
(3) 省略された場合の扱いについては、StbCommon の補足説明を参照のこと。 
(4) 幅止め筋のピッチは縦方向横方向共通とする。 
(5)  子要素<StbSecBarWall_RC_InsideAndOutside>は、土圧、水圧などを受ける壁において、壁の
外部および内部で配筋の異なる状態を想定している（内外部の定義は、子要素の節を参照）。 
(6)  子要素<StbSecBarWall_RC_Edge>による壁端部補強筋は、該当する補強筋が存在する場合に、追
加で記述する。 
(7)  子要素<StbSecBarWall_RC_Open>による壁開口配筋は、壁断面に共通の補強筋で、該当する補強
筋が存在する場合に、追加で記述する。ただし、開口ごとに配筋が異なる場合は、この要素ではな
く、<StbSecOpen_RC>の壁開口配筋を用いる。 
・例 
 
 
かぶり厚さ
かぶり厚さ
Ｚ
Y
ダブル配筋
ダブル配筋（内外異なる）
かぶり厚さ（外）
かぶり厚さ（内）
かぶり厚さ
千鳥配筋
かぶり厚さ
  <StbSecWall_RC id="93" name="EW16"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC depth_cover_outside="30" depth_cover_inside="30" 
outer_bar_direction="HORIZONTAL" D_bar_spacing="D10"  
strength_bar_spacing="SD295A" pitch_bar_spacing="1000"> 
      <StbSecBarWall_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 386

6.14-8 
 
6.14.2.1. ＲＣ壁断面配筋・シングル：StbSecBarWall_RC_Single 
・概要 
説明 ：ＲＣ壁断面の配筋（シングル配筋） 
親要素：StbSecBarArrangementWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
縦筋、横筋について、属性をそれぞれpos= “VERTICAL” およびpos= “HORIZONTAL” としたこの
子要素を各１回記述する。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
 
横筋（ピッチ）
Ｚ
Y
  <StbSecWall_RC id="94" name="EW15"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_Single pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_Single pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 387

6.14-9 
 
6.14.2.2. ＲＣ壁断面配筋・千鳥：StbSecBarWall_RC_Zigzag 
・概要 
説明 ：ＲＣ壁断面の配筋（千鳥配筋） 
親要素：StbSecBarArrangementWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
縦筋、横筋について、属性をそれぞれpos= “VERTICAL” およびpos= “HORIZONTAL” としたこの
子要素を各１回記述する。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
 
横筋（ピッチ）
Ｚ
Y
  <StbSecWall_RC id="95" name="EW18"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_Zigzag pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_Zigzag pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 388

6.14-10 
 
6.14.2.3. ＲＣ壁断面配筋・ダブル：StbSecBarWall_RC_DoubleNet 
・概要 
説明 ：ＲＣ壁断面の配筋（ダブル配筋） 
親要素：StbSecBarArrangementWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
縦筋、横筋について、属性をそれぞれpos= “VERTICAL” およびpos= “HORIZONTAL” としたこの
子要素を各１回記述する。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
横筋（ピッチ）
Ｚ
Y
  <StbSecWall_RC id="96" name="EW20"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_DoubleNet pos="VERTICAL" D="D13" pitch="200"/> 
      <StbSecBarWall_RC_DoubleNet pos="HORIZONTAL" D="D13" pitch="200"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 389

6.14-11 
 
6.14.2.4. ＲＣ壁断面配筋（内外異なる）：StbSecBarWall_RC_InsideAndOutside 
・概要 
説明 ：ＲＣ壁断面の配筋（内外の配筋が異なる場合） 
親要素：StbSecBarArrangementWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL_OUTSIDE（縦筋外側） 
VERTICAL_INSIDE（縦筋内側） 
HORIZONTAL_OUTSIDE（横筋外側） 
HORIZONTAL_INSIDE（横筋内側） 
※(1) 
※(3) 
pos2 
integer 
〇 
鉄筋の段位置 
※(2) 
length1 
double 
 
縦筋上端・横筋始端側の鉄筋切替位置 
※(4) 
length1_ex 
double 
 
縦筋上端・横筋始端側の余長 
※(4) 
length2 
double 
 
縦筋下端・横筋終端側の鉄筋切替位置 
※(4) 
length2_ex 
double 
 
縦筋下端・横筋終端側の余長 
※(4) 
 
・内容、子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_All 
1 
1 
全断面一様な配筋 
※(5) 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarWall_RC_TopStart 
0 
1 
縦筋上端・横筋始端 
※(3) 
※(5) 
※(6) 
StbSecBarWall_RC_Middle 
0 
1 
縦筋中央・横筋中央 
StbSecBarWall_RC_BottomEnd 
0 
1 
縦筋下端・横筋終端 
 
 
・補足 
(1) 外側・内側と部材座標系の関係は<StbWall>の
type_outside で指定する。 
(2) 鉄筋の段位置は右図を参照のこと。  
 
 
 
外 
内 
横筋外側 
横筋内側 
Ｚ
Y
縦筋外側 
縦筋内側 
鉄筋の段位置
縦筋：1 2    2 1 
横筋： 1 2   2 1

## Page 390

6.14-12 
 
(3) 縦筋上端・中央・下端および横筋始端・中央・終端は下図による。 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
外 
内 
外側下端縦筋 
外側中央縦筋 
外側上端縦筋 
内側下端縦筋 
内側中央縦筋 
内側上端縦筋 
梁
梁
壁
横筋 
始端 
横筋 
中央 
横筋 
終端 
外側始端横筋 
外側終端横筋 
外側中央横筋 
内側始端横筋 
内側中央横筋 
内側終端横筋 
柱 
柱 
壁 
外 
内

## Page 391

6.14-13 
 
(4) 鉄筋切替位置のlength1、length2 及び余長length1_ex、length2_ex は下図を参照する。図A、
B は縦筋を例として図示しており、図A が余長length1_ex、length2_ex を用いない（値を記述し
ない）場合の配筋パターン例であり、図B が余長length1_ex、length2_ex を用いる場合の配筋パ
ーン例である。鉄筋切替位置までの距離は基準点からの距離とする。配筋の切り替えが無い場合
は、記述しないものとする。ただし、配筋の切り替えがあり、かつlength1、length2 が省略され
た場合の扱いについては、StbCommon の補足説明を参照のこと。 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
(5) StbSecBarWall_RC_InsideAndOutside では、属性で壁の縦筋・横筋、壁の外側・内側、段位置お
よび鉄筋の切り替え位置を指定し、それぞれの鉄筋に対する鉄筋径・強度・ピッチは子要素
（StbSecBarWall_RC_All 、StbSecBarWall_RC_Top_Start 、StbSecBarWall_RC_Middle 、
StbSecBarWall_RC_Bottom_End ）で指定する。鉄筋の切り替えが無い場合は、
StbSecBarWall_RC_All を用い、鉄筋の切り替えがある場合は、該当位置に応じて
StbSecBarWall_RC_Top_Start、StbSecBarWall_RC_Middle、StbSecBarWall_RC_Bottom_End
で指定する。 
(6) 鉄筋の切り替えがある場合の子要素
StbSecBarWall_RC_Top_Start 、
StbSecBarWall_RC_Middle、StbSecBarWall_RC_Bottom_End は、いずれか1 種類以上の子要
素を持つものとし、全子要素の最小回数が0 であってはならない。 
柱 
柱 
梁 
梁 
壁 
基準点 
基準点 
縦筋 
上端 
縦筋 
中央 
縦筋 
下端 
length1 
length2 
（鉄筋切替位置） 
（鉄筋切替位置） 
図A 余長length1_ex、length2_ex を用いない配筋パターン例 
柱 
柱 
梁 
梁 
壁 
基準点 
基準点 
縦筋 
上端 
縦筋 
中央 
縦筋 
下端 
length1 
length2 
（鉄筋切替位置） 
図B 余長length1_ex、length2_ex を用いる配筋パターン例 
length1_ex（余長） 
(鉄筋切替位置) 
length2_ex(余長)

## Page 392

6.14-14 
 
・例 
 
 
  <StbSecWall_RC id="97" name="EW50"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="2" 
length2="1000"> 
<StbSecBarWall_RC_Top_Start D="D25" pitch="400"/> 
<StbSecBarWall_RC_Middle D="D25" pitch="400"/> 
<StbSecBarWall_RC_Bottom_End D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
<StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_INSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
<StbSecBarWall_RC_InsideAndOutside pos="HORIZONTAL_OUTSIDE" pos2="1" 
length1="1000" length1_ex="330" length2="1000" length2_ex="330"> 
<StbSecBarWall_RC_Top_Start D="D22" pitch="100"/> 
<StbSecBarWall_RC_Middle D="D22" pitch="150"/> 
<StbSecBarWall_RC_Bottom_End D="D22" pitch="100"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
<StbSecBarWall_RC_InsideAndOutside pos="HORIZONTAL_INSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
</StbSecBarArrangementWall_RC> 
</StbSecWall_RC>

## Page 393

6.14-15 
 
6.14.2.4.1. ＲＣ壁断面配筋（内外異なる）全断面一様な配筋：StbSecBarWall_RC_All 
・概要 
説明 ：ＲＣ壁断面の配筋（内外の配筋が異なる場合）の全断面一様な配筋 
親要素：StbSecBarWall_RC_InsideAndOutside 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
 
  <StbSecWall_RC id="97" name="EW30"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D22D25" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
       （略） 
</StbSecBarArrangementWall_RC> 
</StbSecWall_RC>

## Page 394

6.14-16 
 
6.14.2.4.2. ＲＣ壁断面配筋（内外異なる）縦筋上端・横筋始端：StbSecBarWall_RC_TopStart 
・概要 
説明 ：ＲＣ壁断面の配筋（内外の配筋が異なる場合）の縦筋上端もしくは横筋始端 
親要素：StbSecBarWall_RC_InsideAndOutside 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
(1) 「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある場
合は、省略してもよい。 
 
・例 
 
 
 
  <StbSecWall_RC id="97" name="EW30"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="2" 
length2="1000"> 
<StbSecBarWall_RC_Top_Start D="D25" pitch="400"/> 
<StbSecBarWall_RC_Middle D="D25" pitch="400"/> 
<StbSecBarWall_RC_Bottom_End D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
       （略） 
</StbSecBarArrangementWall_RC> 
</StbSecWall_RC>

## Page 395

6.14-17 
 
6.14.2.4.3. ＲＣ壁断面配筋（内外異なる）縦筋中央・横筋中央：StbSecBarWall_RC_Middle 
・概要 
説明 ：ＲＣ壁断面の配筋（内外の配筋が異なる場合）の縦筋中央もしくは横筋中央 
親要素：StbSecBarWall_RC_InsideAndOutside 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
(1) 「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある場
合は、省略してもよい。 
 
・例 
 
 
 
  <StbSecWall_RC id="97" name="EW30"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="2" 
length2="1000"> 
<StbSecBarWall_RC_Top_Start D="D25" pitch="400"/> 
<StbSecBarWall_RC_Middle D="D25" pitch="400"/> 
<StbSecBarWall_RC_Bottom_End D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
       （略） 
</StbSecBarArrangementWall_RC> 
</StbSecWall_RC>

## Page 396

6.14-18 
 
6.14.2.4.4. ＲＣ壁断面配筋（内外異なる）縦筋下端・横筋終端：StbSecBarWall_RC_BottomEnd 
・概要 
説明 ：ＲＣ壁断面の配筋（内外の配筋が異なる場合）の縦筋下端もしくは横筋終端 
親要素：StbSecBarWall_RC_InsideAndOutside 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
(1) 「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある場
合は、省略してもよい。 
・例 
 
 
 
  <StbSecWall_RC id="97" name="EW30"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="1"> 
<StbSecBarWall_RC_All D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
      <StbSecBarWall_RC_InsideAndOutside pos="VERTICAL_OUTSIDE" pos2="2" 
length2="1000"> 
<StbSecBarWall_RC_Top_Start D="D25" pitch="400"/> 
<StbSecBarWall_RC_Middle D="D25" pitch="400"/> 
<StbSecBarWall_RC_Bottom_End D="D25D22" pitch="200"/> 
      </ StbSecBarWall_RC_InsideAndOutside> 
       （略） 
</StbSecBarArrangementWall_RC> 
</StbSecWall_RC>

## Page 397

6.14-19 
 
6.14.2.5. 端部補強筋（コ型補強筋）：StbSecBarWall_RC_Edge 
・概要 
説明 ：ＲＣ壁断面の配筋（壁端部のコ型補強筋） 
親要素：StbSecBarArrangementWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL_START（袖壁始端） 
VERTICAL_END（袖壁終端） 
HORIZONTAL_BOTTOM（たれ壁下端） 
HORIZONTAL_TOP（腰壁上端） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
必要に応じて、補強筋のある位置ごとにこの子要素を各１回記述する。補強筋がその位置にない場合
はこの要素を省略してよい。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
・例 
 
 
  <StbSecWall_RC id="95" name="EW18"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_Edge pos="HORIZONTAL_BOTTOM" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_Edge pos="HORIZONTAL_TOP" D="D10" pitch="200"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 398

6.14-20 
 
6.14.2.6. 壁開口配筋：StbSecBarWall_RC_Open 
・概要 
説明 ：ＲＣ壁断面の配筋（壁断面符号ごとの開口部配筋） 
親要素：StbSecBarArrangementWall_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
DIAGONAL（斜め筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
※(2) 
length 
double 
 
長さ 
 
 
・内容、子要素 
無し 
 
・補足 
必要に応じて、補強筋のある位置ごとにこの子要素を各１回記述する。補強筋がその位置にない場合
はこの要素を省略してよい。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
(2)  「本数」は、１辺および１隅あたりの合計本数とする。 
 
 
 
（1点目）
（2点目）
（3点目）
（4点目）
開口
X
Y
縦筋
横筋
斜め筋

## Page 399

6.14-21 
 
・例 
 
  <StbSecWall_RC id="95" name="EW18"  （略）> 
        （略） 
    <StbSecBarArrangementWall_RC> 
      <StbSecBarWall_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
      <StbSecBarWall_RC_Open pos="VERTICAL" D="D16" N="4"/> 
      <StbSecBarWall_RC_Open pos="HORIZONTAL" D="D16" N="4"/> 
      <StbSecBarWall_RC_Open pos="DIAGONAL" D="D13" N="2"/> 
    </StbSecBarArrangementWall_RC> 
  </StbSecWall_RC>

## Page 400

6.15-1 
 
6.15. 荷重用壁断面：StbSecWallLoad 
・概要 
説明 ：荷重用壁断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
一貫構造計算プログラムなどにおける壁の指定で、ALC 版などコンクリート以外の場合は、仕上げ相
当のm2 当たりの荷重値で考慮する場合がある。ＲＣ壁断面StbSecWall_RC の子要素では、コンクリー
ト部分の厚さt の値を0 より大きい値で必須と規定しているため、このような場合は、本要素にて指定
する。仕上げ相当の荷重値の指定は「計算データ編：StbCalWallFinish_RC」による。 
 
・例 
  <StbSecWallLoad id="103" name="Load1"/>

## Page 401

6.16-1 
 
 免震装置断面：StbSecIsolatingDevice 
・概要 
説明  
：免震装置の製品および接合部 
親要素 
：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecProductIsolatingDevice 
1 
1 
免震装置断面・製品 
※(1) 
      または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSpecificationIsolatingDevice 
1 
1 
免震装置断面・仕様指定 
※(2) 
 
   上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecConnectionIsolatingDevice 
0 
1 
免震装置断面・接合部 
※(3) 
 
・補足 
(1) 配置される装置の名称を、免震装置の種類に応じてそれぞれ指定する。装置の名称は、「製品型番
リスト」で定義される製品型番による（「製品型番一覧表」は、１章を参照）。 
(2) 「製品型番一覧表」によらない免震装置の場合は、この子要素を指定する。製品型番が一覧表に無
く、性能値を直接指定する場合を想定している。その場合、性能値の指定は「計算データ編」の「免
震装置性能（数値指定）」による。 
(3) 免震装置と、装置を固定する取付け躯体または別の部材要素と接合する接合方法（ベースプレー
トを介した接合、ボルトによる接合）とその仕様を指定する。回数が0 となる場合は、構造解析プ
ログラム等で、免震装置の接合方法を対象としない場合を想定している。

## Page 402

6.16-2 
 
・例 
 
 
 
  <StbSecIsolatingDevice id="1" name="NRB1"> 
     <StbSecProductIsolatingDevice> 
       <StbSecIsolatingDeviceNRB product_code="XXXXXX" release_time="2022"/> 
     </StbSecProductIsolatingDevice> 
   </StbSecIsolatingDevice> 
 
   <StbSecIsolatingDevice id="101" name="NRB1series"> 
     <StbSecProductIsolatingDevice> 
       <StbSecIsolatingDeviceNRB product_code="XXXXXX" release_time="2010"/> 
       <StbSecIsolatingDeviceNRB product_code="YYYYYY" release_time="2012"/> 
       <StbSecIsolatingDeviceNRB product_code="ZZZZZZ" release_time="2002"/> 
     </StbSecProductIsolatingDevice> 
   </StbSecIsolatingDevice> 
 
   <StbSecIsolatingDevice id="201" name="NRB1change"> 
     <StbSecProductIsolatingDevice> 
       <StbSecIsolatingDeviceNRB product_code="XXXXXX" release_time="2022"> 
<StbSecIsolatingDeviceSpecificationChange 
 specification_changeable="T_FRANGE" value="24"/> 
       </StbSecIsolatingDeviceNRB> 
     </StbSecProductIsolatingDevice> 
     <StbSecConnectionIsolatingDevice> 
       <StbSecConnectionIsolatingDeviceRB 
   method_start="BASEPLATE" shape_plate_start="SQUARE" size_plate_start="1200" 
t_plate_start="32" strength_plate_start="SN400B" 
method_end="BASEPLATE" shape_plate_end="SQUARE" size_plate_end="1200" 
t_plate_end="32" strength_plate_end="SN400B"/> 
     </StbSecConnectionIsolatingDevice> 
   </StbSecIsolatingDevice> 
 
   <StbSecIsolatingDevice id="301" name="NRB1direct"> 
     <StbSecSpecificationIsolatingDevice/> 
   </StbSecIsolatingDevice>

## Page 403

6.16-3 
 
6.16.1. 免震装置断面・製品：StbSecProductIsolatingDevice 
・概要 
説明  
：免震装置の製品定義 
親要素 
：StbSecIsolatingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
  いずれかひとつを選択する。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceNRB 
1 
制限なし 
天然ゴム系積層ゴム支承 
 
StbSecIsolatingDeviceHDR 
1 
制限なし 
高減衰ゴム系積層ゴム支承 
 
StbSecIsolatingDeviceLRB 
1 
制限なし 
鉛プラグ入り積層ゴム支承 
 
StbSecIsolatingDeviceTRB 
1 
制限なし 
錫プラグ入り積層ゴム支承 
 
StbSecIsolatingDeviceSDRB 
1 
制限なし 
ダンパー一体型積層ゴム支承 
 
StbSecIsolatingDeviceESB 
1 
制限なし 
弾性すべり支承 
 
StbSecIsolatingDeviceRSB 
1 
制限なし 
剛すべり支承 
 
StbSecIsolatingDeviceCSB 
1 
制限なし 
曲面すべり支承 
 
StbSecIsolatingDeviceCLSB 
1 
制限なし 
レール式すべり支承 
 
StbSecIsolatingDeviceCLB 
1 
制限なし 
レール式転がり支承 
 
 
・補足 
同等性能を有する免震装置を複数個示す必要がある場合は、ひとつの免震装置断面で子要素を複
数回数定義する。ただし、複数回数定義する子要素は、いずれか１種類で同一とする。

## Page 404

6.16-4 
 
6.16.1.1. 天然ゴム系積層ゴム支承：StbSecIsolatingDeviceNRB 
・概要 
説明  
：天然ゴム系積層ゴム支承(Natural Rubber Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
天然ゴム系積層ゴム支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」に示される製品型番およびリリース時期にて指定する（「製品型番一覧表」に
ついては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 405

6.16-5 
 
6.16.1.1.1. 免震装置・規定仕様の変更：StbSecIsolatingDeviceSpecificationChange 
・概要 
説明  
：製品型番の形状変更を指定 
親要素 
：StbSecIsolatingDeviceNRB、StbSecIsolatingDeviceHDR、他 
 
・属性 
属性名 
型 
必須 
説明 
補足 
specification_changeable 
string 
○ 
変更する仕様の名称 
 
value 
string 
○ 
変更する値 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「製品型番一覧表」にて特定される製品に対して、一部仕様と異なる形状がある場合、変更する
形状に関する仕様の名称と値の組合せを指定する。各製品を製作するメーカーに応じて、変更可能
とされる形状は異なるが、仕様策定にあたり実施したヒアリング結果をもとに、当面、以下の形状
に関する仕様の名称を対象とする。 
 
変更する仕様の名称 
型 
説明 
補足 
SIZE_FRANGE 
double 
取付けフランジ外形寸法
(mm) 
 
形状円形の時：外径 
形状正方形の時：外形 
形状長方形の時：幅 
LENGTH_FRANGE 
double 
取付けフランジ外形長さ
(mm) 
形状長方形の時 
T_FRANGE 
double 
取付けフランジ厚さ(mm) 
 
PCD_BOLT 
double 
取付けボルトPCD(mm) 
（Pitch Circle Diameter ボ
ルト穴の中心線を結んだ円
の直径） 
形状円形の時：PCD 
形状矩形の時：ボルト穴の中
心間の標準的な間隔 
N_BOLT 
integer 
取付けボルト本数 
 
HOLE_BOLT 
integer 
取付けボルト穴数

## Page 406

6.16-6 
 
NAME_BOLT 
string 
取付けボルト材料名称 
 
HEIGHT 
double 
製品総高さ(mm) 
 
 
すべり板設置型支承の場合は、取付けフランジおよび取付けボルトの値は、支承側（受け材側）
の値とし、すべり板（ソールプレート）板側については以下とする。 
 
変更する仕様の名称 
型 
説明 
補足 
SIZE_SLIDEPLATE 
double 
すべり板寸法(mm) 
 
形状円形の時：外径 
形状正方形の時：外形 
T_ SLIDEPLATE 
double 
すべり板総厚さ(mm) 
 
PCD_BOLT_ 
SLIDEPLATE 
double 
すべり板取付けボルト 
PCD(mm) 
 
形状円形の時：PCD 
形状矩形の時：ボルト穴の中
心間の標準的な間隔 
N_BOLT_ SLIDEPLATE 
integer 
すべり板取付けボルト本数 
 
NAME_BOLT_ 
SLIDEPLATE 
string 
すべり板取付けボルト材料
名称

## Page 407

6.16-7 
 
6.16.1.2. 高減衰ゴム系積層ゴム支承：StbSecIsolatingDeviceHDR 
・概要 
説明  
：高減衰ゴム系積層ゴム支承(High Damping Rubber bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
高減衰ゴム系積層ゴム支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 408

6.16-8 
 
6.16.1.3. 鉛プラグ入り積層ゴム支承：StbSecIsolatingDeviceLRB 
・概要 
説明  
：鉛プラグ入り積層ゴム支承(Lead Rubber Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
鉛プラグ入り積層ゴム支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 409

6.16-9 
 
6.16.1.4. 錫プラグ入り積層ゴム支承：StbSecIsolatingDeviceTRB 
・概要 
説明  
：錫プラグ入り積層ゴム支承(Tin Rubber Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
錫プラグ入り積層ゴム支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 410

6.16-10 
 
6.16.1.5. ダンパー一体型積層ゴム支承：StbSecIsolatingDeviceSDRB 
・概要 
説明 ：鋼材ダンパー一体型積層ゴム支承(Steel Damper Rubber Bearing) の製品定義 
親要素：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
鋼材ダンパー一体型積層ゴム支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 411

6.16-11 
 
6.16.1.6. 弾性すべり支承：StbSecIsolatingDeviceESB 
・概要 
説明  
：弾性すべり支承(Elastic Sliding Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
pos 
string 
○ 
すべり板設置 以下のいずれか 
TOP（支承上側） 
BOTTOM（支承下側） 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
弾性すべり支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 412

6.16-12 
 
6.16.1.7. 剛すべり支承：StbSecIsolatingDeviceRSB 
・概要 
説明  
：剛すべり支承(Rigid Sliding Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
pos 
string 
○ 
すべり板設置 以下のいずれか 
TOP（支承上側） 
BOTTOM（支承下側） 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
剛すべり支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 413

6.16-13 
 
6.16.1.8. 曲面すべり支承：StbSecIsolatingDeviceCSB 
・概要 
説明  
：曲面すべり支承(Curved Sliding Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
曲面すべり支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 414

6.16-14 
 
6.16.1.9. レール式すべり支承：StbSecIsolatingDeviceCLSB 
・概要 
説明  
：レール式すべり支承(Cross Linear Sliding Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
レール式すべり支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 415

6.16-15 
 
6.16.1.10. レール式転がり支承：StbSecIsolatingDeviceCLB 
・概要 
説明  
：レール式転がり支承(Cross Linear Bearing) の製品定義 
親要素 
：StbSecProductIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
isUsedUpsideDown 
boolean 
 
true：キ型(下２列、上1 列)を上下反転し
て設置 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecIsolatingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
レール式転がり支承について、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 416

6.16-16 
 
6.16.2. 免震装置断面・仕様指定：StbSecSpecificationIsolatingDevice 
・概要 
説明  
：免震装置の仕様指定による定義 
親要素 
：StbSecIsolatingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「製品型番一覧表」によらない免震装置の場合は、この子要素を指定する。 
製品型番が一覧表に無く、性能値を直接指定する場合を想定している。その場合、性能値の指定は
「計算データ編」の「免震装置性能（数値指定）」による。

## Page 417

6.16-17 
 
6.16.3. 免震装置断面・接合部：StbSecConnectionIsolatingDevice 
・概要 
説明  
：免震装置の接合部定義 
親要素 
：StbSecIsolatingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
  いずれかひとつを選択する。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecConnectionIsolatingDeviceRB 
1 
1 
積層ゴム支承接合部 
 
StbSecConnectionIsolatingDeviceSP 
1 
1 
すべり板設置型支承接合部 
 
StbSecConnectionIsolatingDeviceLB 
1 
1 
免震層上下接合型支承接合部 
 
 
・補足 
免震装置と、装置を固定する取付け躯体または別の部材要素と接合する接合方法とその仕様を指
定する。当面、以下の接合方法を対象とする（図の青色部分）。 
 
     
   
   
 
        積層ゴム支承       すべり板設置型支承      免震層上下接合型支承 
            （アンカー 
ボルトが 
付く場合）

## Page 418

6.16-18 
 
6.16.3.1. 積層ゴム支承接合部：StbSecConnectionIsolatingDeviceRB 
・概要 
説明  
：積層ゴム支承の接合部定義 
親要素 
：StbSecConnectionIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
method_start 
string 
○ 
始端（下部）接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合） 
 
shape_plate_start 
string 
 
始端（下部）ベースプレート・形状 
以下のいずれか 
CIRCLE（円形）、SQUARE（正方形） 
 
size_plate_start 
double 
 
始端（下部）ベースプレート・外径 
※(3) 
t_plate_start 
double 
 
始端（下部）ベースプレート・厚さ 
 
strength_plate_start 
string 
 
始端（下部）ベースプレート・鉄骨強度 
 
number_studbolt_start 
integer 
 
始端（下部）スタッドボルト・全本数 
※(1) 
name_studbolt_start 
string 
 
始端（下部）スタッドボルト・鉄骨強度 
※(2) 
strength_studbolt_start 
string 
 
始端（下部）スタッドボルト・材料名称 
※(2) 
number_anchorbolt_start 
integer 
 
始端（下部）アンカーボルト類・全本数 
※(1) 
pitch_anchorbolt_start 
double 
 
始端（下部）アンカーボルト類・穴PCD 
※(2)、※(4) 
name_anchorbolt_start 
string 
 
始端（下部）アンカーボルト類・材料名称 
※(2) 
strength_anchorbolt_start 
string 
 
始端（下部）アンカーボルト類・鉄骨強度 
※(2) 
method_end 
string 
○ 
終端（上部）接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合） 
 
shape_plate_end 
string 
 
終端（上部）ベースプレート・形状 
以下のいずれか 
CIRCLE（円形）、SQUARE（正方形） 
 
size_plate_end 
double 
 
終端（上部）ベースプレート・外径 
※(3) 
t_plate_end 
double 
 
終端（上部）ベースプレート・厚さ 
 
strength_plate_end 
string 
 
終端（上部）ベースプレート・鉄骨強度 
 
number_studbolt_end 
integer 
 
終端（上部）スタッドボルト・全本数 
※(1)

## Page 419

6.16-19 
 
name_studbolt_end 
string 
 
終端（上部）スタッドボルト・鉄骨強度 
※(2) 
strength_studbolt_end 
string 
 
終端（上部）スタッドボルト・材料名称 
※(2) 
number_anchorbolt_end 
integer 
 
終端（上部）アンカーボルト類・全本数 
※(1) 
pitch_anchorbolt_end 
double 
 
終端（上部）アンカーボルト類・穴PCD 
※(2)、※(4) 
name_anchorbolt_end 
string 
 
終端（上部）アンカーボルト類・材料名称 
※(2) 
strength_anchorbolt_end 
string 
 
終端（上部）アンカーボルト類・鉄骨強度 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
 
「始端」は装置本体の下側、「終端」は装置本体の上側に定義するものとする。 
 
               終端 
 
 
               始端 
 
(1) 省略された場合は、配置なし（0 本）とする。 
(2) 対応する「全本数」が、0 本より大きい場合は、必須とする。 
(3) 対応する「形状」が、CIRCLE のときは直径、SQUARE のときは外形長さとする。 
(4) 対応する「形状」が、CIRCLE のときはPCD（= Pitch Circle Diameter ボルト穴の中心線を結ん
だ円の直径）、SQUARE のときはボルト穴の中心間の標準的な間隔とする。 
 
 
 
PCD
pitch

## Page 420

6.16-20 
 
6.16.3.2. すべり板設置型支承接合部：StbSecConnectionIsolatingDeviceSP 
・概要 
説明  
：すべり板を設置する型式の接合部定義 
親要素 
：StbSecConnectionIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
method_bearingside 
string 
○ 
支承側接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合） 
 
shape_plate_bearingside 
string 
 
支承側ベースプレート・形状 
以下のいずれか 
CIRCLE（円形）、SQUARE（正方形） 
 
size_plate_bearingside 
double 
 
同・外径 
※(3) 
t_plate_bearingside 
double 
 
同・厚さ 
 
strength_plate_bearingside 
string 
 
同・鉄骨強度 
 
number_studbolt_bearingside 
integer 
 
支承側スタッドボルト・全本数 
※(1) 
name_studbolt_bearingside 
string 
 
同・鉄骨強度 
※(2) 
strength_studbolt_bearingside 
string 
 
同・材料名称 
※(2) 
number_anchorbolt_bearingside 
integer 
 
支承側アンカーボルト類・全本数 
※(1) 
pitch_anchorbolt_bearingside 
double 
 
同・穴PCD 
※(2)、 
※(4) 
name_anchorbolt_bearingside 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_bearingside 
string 
 
同・鉄骨強度 
※(2) 
method_plateside 
string 
○ 
すべり板側接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合） 
 
shape_plate_plateside 
string 
 
すべり板側ベースプレート・形状 
以下のいずれか 
CIRCLE（円形）、SQUARE（正方形） 
 
size_plate_plateside 
double 
 
同・外径 
※(3) 
t_plate_plateside 
double 
 
同・厚さ 
 
strength_plate_plateside 
string 
 
同・鉄骨強度

## Page 421

6.16-21 
 
number_studbolt_plateside 
integer 
 
すべり板側スタッドボルト・全本数 
※(1) 
name_studbolt_plateside 
string 
 
同・鉄骨強度 
※(2) 
strength_studbolt_plateside 
string 
 
同・材料名称 
※(2) 
number_anchorbolt_plateside 
integer 
 
すべり板側アンカーボルト類・全本数 
※(1) 
pitch_anchorbolt_plateside 
double 
 
同・穴PCD 
※(2)、 
※(4) 
name_anchorbolt_plateside 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_plateside 
string 
 
同・鉄骨強度 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 省略された場合は、配置なし（0 本）とする。 
(2) 対応する「全本数」が、0 本より大きい場合は、必須とする。 
(3) 対応する「形状」が、CIRCLE のときは直径、SQUARE のときは外形長さとする。 
(4) 対応する「形状」が、CIRCLE のときはPCD（= Pitch Circle Diameter ボルト穴の中心線を結ん
だ円の直径）、SQUARE のときはボルト穴の中心間の標準的な間隔とする。 
 
 
 
PCD
pitch

## Page 422

6.16-22 
 
6.16.3.3. 免震層上下接合型支承接合部：StbSecConnectionIsolatingDeviceLB 
・概要 
説明  
：鉛直方向を固定する型式の接合部定義 
親要素 
：StbSecConnectionIsolatingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
method_start 
string 
○ 
始端（下部）接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合）、 
WELDING（溶接接合） 
 
t_plate_start 
double 
 
始端（下部）ベースプレート・厚さ 
 
strength_plate_start 
string 
 
同・鉄骨強度 
 
number_studbolt_start 
integer 
 
始端（下部）スタッドボルト・全本数 
※(1) 
name_studbolt_start 
string 
 
同・鉄骨強度 
※(2) 
strength_studbolt_start 
string 
 
同・材料名称 
※(2) 
number_anchorbolt_start 
integer 
 
始端（下部）アンカーボルト類・全本数 
※(1) 
name_anchorbolt_start 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_start 
string 
 
同・鉄骨強度 
※(2) 
method_end 
string 
○ 
終端（上部）接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合）、 
WELDING（溶接接合） 
 
t_plate_end 
double 
 
終端（上部）ベースプレート・厚さ 
 
strength_plate_end 
string 
 
同・鉄骨強度 
 
number_studbolt_end 
integer 
 
終端（上部）スタッドボルト・全本数 
※(1) 
name_studbolt_end 
string 
 
同・鉄骨強度 
※(2) 
strength_studbolt_end 
string 
 
同・材料名称 
※(2) 
number_anchorbolt_end 
integer 
 
終端（上部）アンカーボルト類・全本数 
※(1) 
name_anchorbolt_end 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_end 
string 
 
同・鉄骨強度 
※(2) 
 
・内容 
無し

## Page 423

6.16-23 
 
・子要素 
無し 
・補足 
 
「始端」は装置本体の下側、「終端」は装置本体の上側に定義するものとする。 
 
               終端 
 
 
               始端 
 
 
(1) 省略された場合は、配置なし（0 本）とする。 
(2) 対応する「全本数」が、0 本より大きい場合は、必須とする。

## Page 424

6.17-1 
 
 制振装置断面：StbSecDampingDevice 
・概要 
説明  
：制振装置の製品および接合部 
親要素 
：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecProductDampingDevice 
1 
1 
制振装置断面・製品 
※(1) 
      または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelFigureDampingDevice 
1 
1 
制振装置断面・鉄骨断面形状 
※(2) 
      または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSpecificationDampingDevice 
1 
1 
制振装置断面・仕様指定 
※(3) 
 
   上記子要素に、以下を追加してもよい。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecConnectionDampingDevice 
0 
1 
制振装置断面・接合部 
※(4) 
 
・補足 
(1) 配置される装置の名称を、制振装置の種類に応じてそれぞれ指定する。装置の名称は、「製品型番
一覧表」で定義される製品型番による（「製品型番一覧表」は、X.X 節を参照）。 
(2) StbSecSteel による鋼材形状で定義可能な場合は、この子要素を指定する。 
(3) 「製品型番一覧表」によらず、StbSecSteel にもよらない制振装置の場合は、この子要素を指定す
る。製品型番が一覧表に無く、性能値を直接指定する場合を想定している。その場合、性能値の指
定は「計算データ編」の「制振装置性能（数値指定）」による。

## Page 425

6.17-2 
 
(4) 制振装置と、装置を固定する取付け躯体または別の部材要素と接合する接合方法（ベースプレー
トを介した接合、ボルトによる接合、溶接接合）とその仕様を指定する。回数が0 となる場合は、
構造解析プログラム等で、制振装置の接合方法を対象としない場合を想定している。 
 
・例 
 
 
 
  <StbSecDampingDevice id="1" name="OD1"> 
     <StbSecProductDampingDevice> 
       <StbSecDampingDeviceOil product_code="XXXXXX" release_time="2002"/> 
     </StbSecProductDampingDevice> 
   </StbSecDampingDevice> 
 
  <StbSecDampingDevice id="1" name="OD1series"> 
     <StbSecProductDampingDevice> 
       <StbSecDampingDeviceOil product_code="XXXXXX" release_time="2002"/> 
       <StbSecDampingDeviceOil product_code="YYYYYY" release_time="2022"/> 
     </StbSecProductDampingDevice> 
   </StbSecDampingDevice> 
 
   <StbSecDampingDevice id="201" name=" OD1change"> 
     <StbSecProductDampingDevice> 
       <StbSecDampingDeviceOil product_code=" XXXXXX" release_time="2002"> 
<StbSecDeviceSpecificationChange 
 specification_changeable="T_FRANGE" value="24"/> 
       </StbSecDampingDeviceOil> 
     </StbSecProductDampingDevice> 
     <StbSecConnectionDampingDevice> 
       <StbSecConnectionDampingDeviceHorizontal 
   method_start="THROUGH_BOLT" number_anchorbolt_start="4" 
   name_anchorbolt_start="M39" strength_anchorbolt_start ="SD345" 
method_end="THROUGH_BOLT"" number_anchorbolt_end="4" 
name_anchorbolt_end="M39" strength_anchorbolt_end="SD345"/> 
     </StbSecConnectionDampingDevice> 
   </StbSecDampingDevice> 
 
   <StbSecDampingDevice id="301" name=" OD1direct"> 
     <StbSecSpecificationDampingDevice/> 
   </StbSecDampingDevice>

## Page 426

6.17-3 
 
6.17.1. 制振装置断面・製品：StbSecProductDampingDevice 
・概要 
説明  
：制振装置の製品定義 
親要素 
：StbSecDampingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
  いずれかひとつを選択する。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceOil 
1 
制限なし 
流体系ダンパー 
 
StbSecDampingDeviceViscous 
1 
制限なし 
粘性体ダンパー 
 
StbSecDampingDeviceViscoelastic 
1 
制限なし 
粘弾性体ダンパー 
 
StbSecDampingDeviceHistory 
1 
制限なし 
履歴系ダンパー 
 
StbSecDampingDeviceFriction 
1 
制限なし 
摩擦ダンパー 
 
StbSecDampingDeviceMass 
1 
制限なし 
質量系ダンパー 
 
 
・補足 
同等性能を有する制振装置を複数個示す必要がある場合は、ひとつの制振装置断面で子要素を複
数回数定義する。ただし、複数回数定義する子要素は、いずれか１種類で同一とする。

## Page 427

6.17-4 
 
6.17.1.1. 流体系ダンパー：StbSecDampingDeviceOil 
・概要 
説明  
：流体系ダンパーの製品定義 
親要素 
：StbSecProductDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
流体系ダンパーについて、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 428

6.17-5 
 
6.17.1.1.1. 制振装置・規定仕様の変更：StbSecDampingDeviceSpecificationChange 
・概要 
説明  
：製品型番の形状変更を指定 
親要素 
：StbSecDampingDeviceOil、StbSecDampingDeviceViscous、他 
 
・属性 
属性名 
型 
必須 
説明 
補足 
specification_changeable 
string 
○ 
変更する仕様の名称 
 
value 
string 
○ 
変更する値 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「製品型番一覧表」にて特定される製品に対して、一部仕様と異なる形状がある場合、変更する
形状に関する仕様の名称と値の組合せを指定する。各製品を製作するメーカーに応じて、変更可能
とされる形状は異なるが、仕様策定にあたり実施したヒアリング結果をもとに、当面、以下の形状
に関する仕様の名称を対象とする。 
 
変更する仕様の名称 
型 
説明 
補足 
WIDTH_MOUNTPLATE 
double 
取付け板外形幅 (mm) 
 
LENGTH_MOUNTPLATE 
double 
取付け板外形長さ(mm) 
 
T_MOUNTPLATE 
double 
取付け板厚さ(mm) 
 
NAME_MOUNTPLATE 
string 
取付け板材料名称 
 
WIDTH_SPLICEPLATE 
double 
継ぎ板外形幅 (mm) 
 
LENGTH_SPLICEPLATE 
double 
継ぎ板外形長さ(mm) 
 
T_SPLICEPLATE 
double 
継ぎ板厚さ(mm) 
 
NAME_SPLICEPLATE 
string 
継ぎ板材料名称 
 
PITCH_BOLT 
double 
取付けボルトの標準的な間隔
(mm) 
 
N_BOLT 
integer 
取付けボルト本数 
 
HOLE_BOLT 
integer 
取付けボルト穴数

## Page 429

6.17-6 
 
NAME_BOLT 
string 
取付けボルト材料名称 
 
LENGTH 
double 
製品総長さ（高さ）(mm)

## Page 430

6.17-7 
 
6.17.1.2. 粘性体ダンパー：StbSecDampingDeviceViscous 
・概要 
説明  
：粘性体ダンパーの製品定義 
親要素 
：StbSecProductDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
粘性体ダンパーについて、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 431

6.17-8 
 
6.17.1.3. 粘弾性体ダンパー：StbSecDampingDeviceViscoelastic 
・概要 
説明  
：粘弾性体ダンパーの製品定義 
親要素 
：StbSecProductDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
粘弾性体ダンパーについて、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 432

6.17-9 
 
6.17.1.4. 履歴系ダンパー：StbSecDampingDeviceHistory 
・概要 
説明  
：履歴系ダンパーの製品定義 
親要素 
：StbSecProductDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
履歴系ダンパーについて、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 433

6.17-10 
 
6.17.1.5. 摩擦ダンパー：StbSecDampingDeviceFriction 
・概要 
説明  
：摩擦ダンパーの製品定義 
親要素 
：StbSecProductDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
摩擦ダンパーについて、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 434

6.17-11 
 
6.17.1.6. 質量系ダンパー：StbSecDampingDeviceMass 
・概要 
説明  
：質量系ダンパーの製品定義 
親要素 
：StbSecProductDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecDampingDeviceSpecificationChange 
0 
制限なし 
規定仕様の変更 
※(2) 
 
・補足 
質量系ダンパーについて、配置される装置の名称を指定する。 
 
(1) 「製品型番一覧表」による製品型番およびリリース時期にて指定する（「製品型番一覧表」につい
ては、第１章を参照）。 
(2) 「製品型番一覧表」にて特定される製品に対して、一部異なる形状がある場合、変更する形状に関
する仕様の名称と値の組合せを指定する。要素の最大回数は、変更する属性の数となる。

## Page 435

6.17-12 
 
6.17.2. 制振装置断面・鉄骨断面形状：StbSecSteelFigureDampingDevice 
・概要 
説明  
：制振装置の鉄骨断面形状による定義 
親要素 
：StbSecDampingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecSteelFigureDampingDeviceHistory 
1 
制限なし 
履歴系ダンパー・鉄骨断面形状 
 
 
・補足 
StbSecSteel による鋼材形状で定義可能な場合は、この子要素を指定する。

## Page 436

6.17-13 
 
6.17.2.1. 履歴系ダンパー・鉄骨断面形状：StbSecSteelFigureDampingDeviceHistory 
・概要 
説明  
：履歴系ダンパーの鉄骨断面形状による定義 
親要素 
：StbSecSteelFigureDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
shape 
string 
○ 
芯材鉄骨形状 
※(1) 
shape_sub 
string 
 
芯材鉄骨形状（副） 
※(1) 
strength 
string 
○ 
芯材鉄骨強度 
 
kind_section_stiffener 
string 
 
補剛材・構造種別 
※(2) 
id_section_stiffener 
integer 
 
補剛材・断面ID 
※(2) 
length 
double 
 
装置部分の長さ、または高さ(mm) 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 芯材の鉄骨形状に応じ、鉄骨断面要素 <StbSecSteel> で定義した断面形状を指定する。（副）は、
平鋼を２個組み合わせる場合など、２個目の形状指定がある場合に指定する。 
(2) ブレース形状の制振装置において、補剛材の断面形状を柱断面とみなして、対応する柱断面の構
造種別および断面ID を指定する。構造種別と断面ID が参照する要素の組合せは以下とする。 
 
構造種別 
断面ID が参照する要素 
RC 
StbSecColumn_RC 
S 
StbSecColumn_S 
SRC 
StbSecColumn_SRC 
CFT 
StbSecColumn_CFT

## Page 437

6.17-14 
 
6.17.3. 制振装置断面・仕様指定：StbSecSpecificationDampingDevice 
・概要 
説明  
：制振装置の仕様指定による定義 
親要素 
：StbSecDampingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「製品型番一覧表」によらず、StbSecSteel にもよらない制振装置の場合は、この子要素を指定す
る。製品型番が一覧表に無く、性能値を直接指定する場合を想定している。その場合、性能値の指
定は「計算データ編」の「制振装置性能（数値指定）」による。

## Page 438

6.17-15 
 
6.17.4. 制振装置断面・接合部：StbSecConnectionDampingDevice 
・概要 
説明  
：制振装置の接合部定義 
親要素 
：StbSecDampingDevice 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
  いずれかひとつを選択する。 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecConnectionDampingDeviceHorizontal 
1 
1 
水平方向接合型接合部 
 
StbSecConnectionDampingDeviceVertical 
1 
1 
層上下接合型接合部 
 
 
・補足 
制振装置と、装置を固定する取付け躯体または別の部材要素と接合する接合方法（ベースプレー
トを介した接合、ボルトによる接合、溶接接合）とその仕様を指定する。当面、以下の接合方法を
対象とする（図の青色部分）。 
        
 
  水平方向接合型（免震層において、水平方向に配置する部材を取付け躯体に接合） 
 
 
 
層上下接合型（壁・間柱・ブレース型の部材を上下の取付け躯体に接合）

## Page 439

6.17-16 
 
6.17.4.1. 水平方向接合型接合部：StbSecConnectionDampingDeviceHorizontal 
・概要 
説明  
：水平方向に配置する部材の接合部定義 
親要素 
：StbSecConnectionDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
method_start 
string 
○ 
始端接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合）、 
THROUGH_BOLT（通しボルト接合）、 
WELDING（溶接接合） 
 
t_plate_start 
double 
 
始端ベースプレート・厚さ 
 
strength_plate_start 
string 
 
同・鉄骨強度 
 
number_anchorbolt_start 
integer 
 
始端ボルト類・全本数 
※(1) 
name_anchorbolt_start 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_start 
string 
 
同・鉄骨強度 
※(2) 
method_end 
string 
○ 
終端接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合）、 
THROUGH_BOLT（通しボルト接合）、 
WELDING（溶接接合） 
 
t_plate_end 
double 
 
終端ベースプレート・厚さ 
 
strength_plate_end 
string 
 
同・鉄骨強度 
 
number_anchorbolt_end 
integer 
 
終端ボルト類・全本数 
※(1) 
name_anchorbolt_end 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_end 
string 
 
同・鉄骨強度 
※(2) 
 
・内容 
無し 
 
・子要素 
無し

## Page 440

6.17-17 
 
・補足 
(1) 省略された場合は、配置なし（0 本）とする。 
(2) 対応する「全本数」が、0 本より大きい場合は、必須とする。

## Page 441

6.17-18 
 
6.17.4.2. 層上下接合型接合部：StbSecConnectionDampingDeviceVertical 
・概要 
説明  
：壁・間柱型部材の上下接合部定義 
親要素 
：StbSecConnectionDampingDevice 
 
・属性 
属性名 
型 
必須 
説明 
補足 
method_start 
string 
○ 
始端（下部）接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合）、 
WELDING（溶接接合） 
 
t_plate_start 
double 
 
始端（下部）ベースプレート・厚さ 
 
strength_plate_start 
string 
 
同・鉄骨強度 
 
number_anchorbolt_start 
integer 
 
始端（下部）ボルト類・全本数 
※(1) 
name_anchorbolt_start 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_start 
string 
 
同・鉄骨強度 
※(2) 
method_end 
string 
○ 
終端（上部）接合方法 
以下のいずれか 
BASEPLATE（ベースプレート接合）、 
BOLT（ボルト接合）、 
WELDING（溶接接合） 
 
t_plate_end 
double 
 
終端（上部）ベースプレート・厚さ 
 
strength_plate_end 
string 
 
同・鉄骨強度 
 
number_anchorbolt_end 
integer 
 
終端（上部）ボルト類・全本数 
※(1) 
name_anchorbolt_end 
string 
 
同・材料名称 
※(2) 
strength_anchorbolt_end 
string 
 
同・鉄骨強度 
※(2) 
 
・内容 
無し 
 
・子要素 
無し

## Page 442

6.17-19 
 
・補足 
 
「始端」は装置本体の下側、「終端」は装置本体の上側に定義するものとする。 
 
               終端 
 
 
 
 
               始端 
 
(1) 省略された場合は、配置なし（0 本）とする。 
(2) 対応する「全本数」が、0 本より大きい場合は、必須とする。

## Page 443

6.18-1 
 
6.18. ＲＣ基礎断面：StbSecFoundation_RC 
・概要 
説明 ：ＲＣ基礎断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
strength_concrete 
string 
 
コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureFoundation_RC 
1 
1 
ＲＣ基礎断面形状 
 
StbSecBarArrangementFoundation_RC 
0 
1 
ＲＣ基礎断面配筋 
※(2) 
 
・補足 
(1)  省略された場合は、参照する<StbFooting>の節点ID、<StbStripFooting>の始端節点ID が所
属する<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した
<StbStory>のコンクリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリ
ート強度」strength_concrete をこの要素のコンクリート強度とする。 
(2)  子要素<StbSecBarArrangementSlab_RC>の回数が0 となる場合は、鉄筋を扱わないプログラ
ムが一時的に作成する場合を想定しており、無筋であることを示すものではない。

## Page 444

6.18-2 
 
  <StbSecFoundation_RC id="1" name="F1" strength_concrete="FC24"> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_Rect width_X="1000" width_Y="1000" depth="900"/> 
    </StbSecFigureFoundation_RC> 
    <StbSecBarArrangementFoundation_RC> 
      <StbSecBarFoundation_RC_Rect pos="X_BOTTOM" D="D16" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="Y_BOTTOM" D="D16" N="7"/> 
    </StbSecBarArrangementFoundation_RC> 
  </StbSecFoundation_RC> 
・例

## Page 445

6.18-3 
 
6.18.1. ＲＣ基礎断面形状：StbSecFigureFoundation_RC 
・概要 
説明 ：ＲＣ基礎断面の形状 
親要素：StbSecFoundation_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFoundation_RC_Rect 
1 
1 
ＲＣ基礎断面形状・矩形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFoundation_RC_TaperedRect 
1 
1 
ＲＣ基礎断面形状・矩形テーパー 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFoundation_RC_Triangle 
1 
1 
ＲＣ基礎断面形状・直角三角形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFoundation_RC_EquiTriangle 
1 
1 
ＲＣ基礎断面形状・正三角形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFoundation_RC_Octagon 
1 
1 
ＲＣ基礎断面形状・八角形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFoundation_RC_Continuous 
1 
1 
ＲＣ基礎断面形状・連続 
 
  
・補足 
 
・例

## Page 446

6.18-4 
 
<StbSecFoundation_RC id="1" name="F1" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_Rect width_X="1000" width_Y="1000" depth="900"/> 
    </StbSecFigureFoundation_RC> 
        （略） 
  </StbSecFoundation_RC> 
6.18.1.1. ＲＣ基礎断面形状・矩形：StbSecFoundation_RC_Rect 
・概要 
説明 ：ＲＣ基礎断面の形状・矩形 
親要素：StbSecFigureFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_X 
double 
○ 
Ｘ幅 
 
width_Y 
double 
○ 
Ｙ幅 
 
depth 
double 
○ 
厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
Ｘ幅、Ｙ幅および厚さの定義は下図による。 
下図右の配置を0 度（<StbFooting>の属性rotate= "0"）とする。 
 
 
・例 
 
 
厚さ 
X 
Z 
 
X 
Y 
 
X 幅 
Y 幅

## Page 447

6.18-5 
 
6.18.1.2. ＲＣ基礎断面形状・矩形テーパー：StbSecFoundation_RC_TaperedRect 
・概要 
説明 ：ＲＣ基礎断面・矩形テーパーの形状 
親要素：StbSecFigureFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_X 
double 
○ 
Ｘ幅 
 
width_Y 
double 
○ 
Ｙ幅 
 
depth_base 
double 
○ 
根元厚さ 
 
depth_tip 
double 
○ 
先端厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
Ｘ幅、Ｙ幅および厚さの定義は下図による。 
下図右の配置を0 度（<StbFooting>の属性rotate= "0"）とする。 
 
 
・例 
X 
Z 
 
根元厚さ 
先端厚さ 
X 
Y 
 
X 幅 
Y 幅

## Page 448

6.18-6 
 
<StbSecFoundation_RC id="2" name="F2" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_TaperedRect width_X="1000" width_Y="1000" 
depth_base="900" 
depth_tip="400"/> 
    </StbSecFigureFoundation_RC> 
        （略）

## Page 449

6.18-7 
 
6.18.1.3. ＲＣ基礎断面形状・直角三角形：StbSecFoundation_RC_Triangle 
・概要 
説明 ：ＲＣ基礎断面・直角三角形の形状 
親要素：StbSecFigureFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_X 
double 
○ 
Ｘ幅 
 
width_Y 
double 
○ 
Ｙ幅 
 
width_chamfer_X 
double 
 
面取りＸ幅 
※(1) 
width_chamfer_Y 
double 
 
面取りＹ幅 
※(1) 
depth 
double 
○ 
厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
Ｘ幅、Ｙ幅および厚さの定義は下図による。 
下図右の三角形配置を0 度（<StbFooting>の属性 rotate= "0"） とする。 
(1)  面取り幅が省略された場合は、0 とする。 
 
 
 
 
 
X 
Z 
 
厚さ 
X 
Y 
面取りX 幅 
面取りY 幅 
X 幅 
Y 幅

## Page 450

6.18-8 
 
<StbSecFoundation_RC id="3" name="F3" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_Triangle width_X="1000" width_Y="1000" 
width_chamfer_X="300" width_chamfer_Y="300" depth="900"/> 
    </StbSecFigureFoundation_RC> 
        （略） 
・例

## Page 451

6.18-9 
 
<StbSecFoundation_RC id="4" name="F4" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_EquiTriangle 
 width_base ="800" width_chamfer ="400" depth="900"/> 
    </StbSecFigureFoundation_RC> 
        （略） 
  </StbSecFoundation_RC> 
6.18.1.4. ＲＣ基礎断面形状・正三角形：StbSecFoundation_RC_EquiTriangle 
・概要 
説明 ：ＲＣ基礎断面・正三角形の形状 
親要素：StbSecFigureFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_base 
double 
○ 
底辺幅 
 
width_chamfer 
double 
○ 
面取り幅 
 
depth 
double 
○ 
厚さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
 
・補足 
定義は下図による。底辺幅、面取り幅は３辺とも同じとする 
下図右の三角形配置を0 度（<StbFooting>の属性rotate= "0"） とする。 
 
 
 
 
 
 
 
・例 
 
 
底辺幅 
面取り幅 
X 
Z 
X 
Y 
厚さ

## Page 452

6.18-10 
 
6.18.1.5. ＲＣ基礎断面形状・八角形：StbSecFoundation_RC_Octagon 
・概要 
説明 ：ＲＣ基礎断面・八角形の形状 
親要素：StbSecFigureFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_X 
double 
○ 
X 幅 
 
width_Y 
double 
○ 
Y 幅 
 
width_chamfer1_X 
double 
○ 
面取りX 幅(1) 
※(1) 
width_chamfer1_Y 
double 
○ 
面取りY 幅(1) 
※(1) 
width_chamfer2_X 
double 
○ 
面取りX 幅(2) 
※(1) 
width_chamfer2_Y 
double 
○ 
面取りY 幅(2) 
※(1) 
width_chamfer3_X 
double 
○ 
面取りX 幅(3) 
※(1) 
width_chamfer3_Y 
double 
○ 
面取りY 幅(3) 
※(1) 
width_chamfer4_X 
double 
○ 
面取りX 幅(4) 
※(1) 
width_chamfer4_Y 
double 
○ 
面取りY 幅(4) 
※(1) 
depth 
double 
○ 
厚さ 
 
 
・内容、子要素 
無し 
 
・補足 
（1）定義は下図による。下図右の配置を0 度（<StbFooting>の属性rotate= "0"） とする。 
    （2）五角形、六角形、台形などの基礎形状を、四隅の面取り幅を変えて表してもよい。ただし、
面取り幅が0 となる場合も、属性の省略はできない。 
 
 
 
 
 
 
 
 
 
 
  
 
X 
Z 
X 
Y 
Y 幅 
X 幅 
面取りX 幅(1) 
面取りY 幅(1) 
面取りX 幅(2) 
面取りY 幅(2) 
面取りX 幅(4) 
面取りX 幅(3) 
面取りY 幅(4) 
面取りY 幅(3) 
厚さ

## Page 453

6.18-11 
 
6.18.1.6. ＲＣ連続基礎断面形状：StbSecFoundation_RC_Continuous 
・概要 
説明 ：ＲＣ基礎断面・連続の形状 
親要素：StbSecFigureFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width 
double 
○ 
幅 
 
depth_base 
double 
○ 
根元厚さ 
 
depth_tip 
double 
○ 
先端厚さ 
 
type 
string 
○ 
タイプ 以下のいずれか 
RIGHT_L 
LEFT_L 
REVERSE_T 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  「タイプ」は基礎梁の始端から終端に向かって、脚部がどちら側に取付くかを示す。 
 
      RIGHT-L           LEFT-L        REVERSE-T 
 
 
 
 
 
 
 
 
 
 
 width 
 width 
 width 
depth_base 
depth_tip 
基礎梁 
基礎梁 
基礎梁

## Page 454

6.18-12 
 
6.18.2. ＲＣ基礎断面配筋：StbSecBarArrangementFoundation_RC 
・概要 
説明 ：ＲＣ基礎断面の配筋 
親要素：StbSecFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_top 
double 
 
かぶり厚さ(上) 
※(1) 
depth_cover_bottom 
double 
 
かぶり厚さ(下) 
depth_cover_side 
double 
 
かぶり厚さ(側面) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarFoundation_RC_Rect 
2 
5 
ＲＣ基礎断面配筋 
・矩形、八角形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarFoundation_RC_Triangle 
2 
5 
ＲＣ基礎断面配筋 
・直角三角形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarFoundation_RC_ThreeWay 
2 
5 
ＲＣ基礎断面配筋 
・三方 正三角形、六角形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarFoundation_RC_Continuous 
2 
5 
ＲＣ基礎断面配筋・連続

## Page 455

6.18-13 
 
  <StbSecFoundation_RC id="1" name="F1" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_Rect width_X="1000" width_Y="1000" depth="900"/> 
    </StbSecFigureFoundation_RC> 
    <StbSecBarArrangementFoundation_RC 
 depth_cover_top="50" depth_cover_bottom="60" depth_cover_side="50"> 
      <StbSecBarFoundation_RC_Rect pos="X_TOP" D="D13" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="Y_TOP" D="D13" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="X_BOTTOM" D="D16" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="Y_BOTTOM" D="D16" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="HORIZONTAL" D="D10" N="3"/> 
    </StbSecBarArrangementFoundation_RC> 
  </StbSecFoundation_RC> 
・補足 
(1)  かぶり厚さの定義は、下図による。上端配筋および横配筋がないときは省略してよい。それ以
外で省略された場合の扱いは、<StbApplyConditionsList>の補足説明による。ただし、該当属性
がない場合は、省略してはならない。 
 
 
 
・例 
 
 
 
X
Y
かぶり厚さ（側面）
かぶり厚さ（側面）
X
Z
かぶり厚さ（上）
かぶり厚さ（下）
かぶり厚さ（側面）
かぶり厚さ（側面）

## Page 456

6.18-14 
 
6.18.2.1. ＲＣ基礎断面配筋・矩形：StbSecBarFoundation_RC_Rect 
・概要 
説明 ：ＲＣ基礎断面の配筋（矩形、八角形基礎用） 
親要素：StbSecBarArrangementFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
X_TOP（X 方向上端） 
X_BOTTOM（X 方向下端） 
Y_TOP（Y 方向上端） 
Y_BOTTOM（Y 方向下端） 
HORIZONTAL（横） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
 
isVertical 
boolean 
 
外周の縦筋の扱い 
上端筋の場合は立ち下げるか否か 
下端筋の場合は立ち上げるか否か 
※(2) 
length_vertical 
double 
 
上端筋の場合は立ち下げ長さ 
下端筋の場合は立ち上げ長さ 
※(3) 
 
・内容、子要素 
無し 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。但し、いわゆる
「かご配筋」でない、下端筋のみの場合は、属性をそれぞれpos= “X_BOTTOM” およびpos= 
“Y_BOTTOM” としたこの子要素を各１回記述する。いずれの場合も、部分的な回数の省略はできない。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にあ
る場合は、省略してもよい。 
 
X
Z
Y方向上端筋
X方向上端筋
横筋
Y方向下端筋
X方向下端筋

## Page 457

6.18-15 
 
(2)  isVertical は外周の縦筋に関する設定で、pos が上端筋の場合は立ち下げる（true）か否
（false）か、pos が下端筋の場合は立ち上げる（true）か否（false）かを指定する。省略された
場合はfalse とする。なお、pos が横筋の場合は省略する。下表参照のこと。 
(3)  length_vertical はisVertical がtrue の場合に、立ち下げ長さもしくは立ち上げ長さを指定す
る。isVertical がtrue でlength_vertical が省略された場合は、上端筋の場合は下端筋まで、下
端筋の場合は上端筋までとする。isVertical がfalse もしくは省略された場合は、length_vertical
も省略する。下表参照のこと。 
 
配筋の概要 
配筋が下端筋のみ 
下端筋の立ち上がり無し 
配筋が下端筋のみ 
下端筋の立ち上がり有り 
pos 
X_BOTTOM 
Y_BOTTOM 
X_BOTTOM 
Y_BOTTOM 
isVertical 
省略（false） 
true 
length_vertical 
省略 
長さ（double）を入力 
 
 
 
 
 
 
配筋の概要 
かご配筋 
上端筋を立ち下げ 
かご配筋 
下端筋を立ち上げ 
pos 
X_TOP 
Y_TOP 
X_BOTTOM 
Y_BOTTOM 
X_TOP 
Y_TOP 
X_BOTTOM 
Y_BOTTOM 
isVertical 
true 
省略（false） 
省略（false） 
true 
length_vertical 
省略 
省略 
省略 
省略 
 
 
 
 
 
 
 
 
 
X 
Z 
XX 幅 
Z 
X 
Z 
length 
_vertical 
X 
Z

## Page 458

6.18-16 
 
  <StbSecFoundation_RC id="1" name="F1" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_Rect width_X="1000" width_Y="1000" depth="900"/> 
    </StbSecFigureFoundation_RC> 
    <StbSecBarArrangementFoundation_RC 
 depth_cover_top="50" depth_cover_bottom="60" depth_cover_side="50"> 
      <StbSecBarFoundation_RC_Rect pos="X_TOP" D="D13" N="7" isVertical="true"/> 
      <StbSecBarFoundation_RC_Rect pos="Y_TOP" D="D13" N="7" isVertical="true"/> 
      <StbSecBarFoundation_RC_Rect pos="X_BOTTOM" D="D16" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="Y_BOTTOM" D="D16" N="7"/> 
      <StbSecBarFoundation_RC_Rect pos="HORIZONTAL" D="D10" N="3"/>    
</StbSecBarArrangementFoundation_RC> 
  </StbSecFoundation_RC> 
配筋の概要 
かご配筋 
上端筋を立ち下げ 
下端筋を立ち上げ 
pos 
X_TOP 
Y_TOP 
X_BOTTOM 
Y_BOTTOM 
isVertical 
true 
true 
length_vertical 
長さ（double）を入力 
長さ（double）を入力 
 
 
 
 
 
・例 
 
 
 
X 
Z 
下端筋の 
length_vertical 
上端筋の 
length_vertical

## Page 459

6.18-17 
 
6.18.2.2. ＲＣ基礎断面配筋・三角：StbSecBarFoundation_RC_Triangle 
・概要 
説明 ：ＲＣ基礎断面の配筋（直角三角形基礎用） 
親要素：StbSecBarArrangementFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
MAIN_TOP（主筋方向上端） 
MAIN_BOTTOM（主筋方向下端） 
TRANSVERSE_TOP（配力筋方向上端） 
TRANSVERSE_BOTTOM （配力筋方向下
端） 
HORIZONTAL（横） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
 
isVertical 
boolean 
 
外周の縦筋の扱い 
上端筋の場合は立ち下げるか否か 
下端筋の場合は立ち上げるか否か 
※(2) 
length_vertical 
double 
 
上端筋の場合は立ち下げ長さ 
下端筋の場合は立ち上げ長さ 
※(3) 
 
・内容、子要素 
無し 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。但し、いわゆる
「かご配筋」でない、下端筋のみの場合は、属性をそれぞれpos= “MAIN_BOTTOM” およびpos=  
“TRANSVERSE_BOTTOM” としたこの子要素を２回記述する。いずれの場合も、部分的な回数の省
略はできない。 
 
 
 
 
 
 
 
 
 
X
Z
主筋上端
主筋下端
横筋
配力筋上端
配力筋下端
X 
Y 
配力筋 
主筋

## Page 460

6.18-18 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にあ
る場合は、省略してもよい。 
(2)  isVertical は外周の縦筋に関する設定で、pos が上端筋の場合は立ち下げる（true）か否
（false）か、pos が下端筋の場合は立ち上げる（true）か否（false）かを指定する。省略された
場合はfalse とする。なお、pos が横筋の場合は省略する。6.18.2.1 の補足(3)表を参照のこと。 
(3)  length_vertical はisVertical がtrue の場合に、立ち下げ長さもしくは立ち上げ長さを指定す
る。isVertical がtrue でlength_vertical が省略された場合は、上端筋の場合は下端筋まで、下
端筋の場合は上端筋までとする。isVertical がfalse もしくは省略された場合は、length_vertical
も省略する。6.18.2.1 の補足(3)表を参照のこと。

## Page 461

6.18-19 
 
6.18.2.3. ＲＣ基礎断面配筋・三方：StbSecBarFoundation_RC_ThreeWay 
・概要 
説明 ：ＲＣ基礎断面の配筋（三方 正三角形、六角形基礎用） 
親要素：StbSecBarArrangementFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
MAIN_TOP（主筋方向上端） 
MAIN_BOTTOM（主筋方向下端） 
OUTSIDE_TOP（外周上端） 
OUTSIDE_BOTTOM（外周下端） 
HORIZONTAL（横） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
 
isVertical 
boolean 
 
外周の縦筋の扱い 
上端筋の場合は立ち下げるか否か 
下端筋の場合は立ち上げるか否か 
※(2) 
length_vertical 
double 
 
上端筋の場合は立ち下げ長さ 
下端筋の場合は立ち上げ長さ 
※(3) 
 
・内容、子要素 
無し 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。但し、いわゆる
「かご配筋」でない、下端筋のみの場合は、属性をそれぞれpos= “MAIN_BOTTOM” およびpos=  
“OUTSIDE_BOTTOM” としたこの子要素を各１回記述する。いずれの場合も、部分的な回数の省略は
できない。 
 
 
 
X
Y
X
Z
主筋方向上端
主筋方向下端
横筋
外周上端
外周下端
主筋方向上端
（上端）
（下端）
主筋方向上端
主筋方向上端
外周上端
主筋方向下端
主筋方向下端
主筋方向下端
外周下端

## Page 462

6.18-20 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にあ
る場合は、省略してもよい。  
(2) isVertical は外周の縦筋に関する設定で、pos が上端筋の場合は立ち下げる（true）か否（false）
か、pos が下端筋の場合は立ち上げる（true）か否（false）かを指定する。省略された場合は
false とする。なお、pos が横筋の場合は省略する。6.18.2.1 の補足(3)表を参照のこと。 
(3)  length_vertical はisVertical がtrue の場合に、立ち下げ長さもしくは立ち上げ長さを指定す
る。isVertical がtrue でlength_vertical が省略された場合は、上端筋の場合は下端筋まで、下
端筋の場合は上端筋までとする。isVertical がfalse もしくは省略された場合は、length_vertical
も省略する。6.18.2.1 の補足(3)表を参照のこと。

## Page 463

6.18-21 
 
6.18.2.4. ＲＣ基礎断面配筋・連続：StbSecBarFoundation_RC_Continuous 
・概要 
説明 ：ＲＣ基礎断面の配筋（連続基礎用） 
親要素：StbSecBarArrangementFoundation_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
MAIN_BASE_TOP（主筋方向元端上端筋） 
MAIN_BASE_BOTTOM（主筋方向元端下端
筋） 
MAIN_TIP_TOP（主筋方向先端上端筋） 
MAIN_TIP_BOTTOM （主筋方向先端下端
筋） 
TRANSVERSE_TOP（配力筋方向上端筋） 
TRANSVERSE_BOTTOM（配力筋方向下端
筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(5) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
※(6) 
pitch 
double 
○ 
ピッチ 
※(6) 
length 
double 
 
主筋の鉄筋切替位置 
※(1),※(2) 
main_type 
string 
 
主筋方向元端先端で配筋を切り替える場合の
配筋パターンを示す 
以下のいずれかの値をとる 
SYMMETRICAL 
ASYMMETRICAL 
※(3),※(4) 
isVertical 
boolean 
 
外周の縦筋の扱い 
上端筋の場合は立ち下げるか否か 
下端筋の場合は立ち上げるか否か 
※(7) 
length_vertical 
double 
 
上端筋の場合は立ち下げ長さ 
下端筋の場合は立ち上げ長さ 
※(8) 
 
・内容、子要素 
無し

## Page 464

6.18-22 
 
・補足 
「配筋位置」pos が示す配筋の位置は、下図とし、子要素回数は各１回とする。但し、下端筋のみ
の場合は、属性をそれぞれpos= “MAIN_BOTTOM” およびpos= “TRANSVERSE_BOTTOM” とした
この子要素を２回記述する。いずれの場合も、部分的な回数の省略はできない。 
(1)  「主筋の切替位置」を示すlength は、当該連続基礎の始端基準点と終端基準点を結んだ軸線
から元端主筋の端部までの距離とする。前頁の図を参照のこと。 
(2)  主筋の切り替えが無い場合およびpos が配力筋や横筋の場合は、length は記述しないものとす
る。ただしpos が主筋で配筋の切り替えがあり、かつlength が省略された場合の扱いについて
は、StbCommon の補足説明を参照のこと。 
(3) ＲＣ連続基礎断面形状（StbSecFoundation_RC_Continuous）のtype が「REVERSE_T」で主
筋に切り替えがある場合、以下の配筋パターンが考えられる。いずれの配筋パターンかは、pos
が主筋の場合のmain_type にて明示する。 
(4) ＲＣ連続基礎断面形状（StbSecFoundation_RC_Continuous ）のtype が「RIGHT_L 」
「LEFT_L」の場合はmain_type は記述しないものとする。またＲＣ連続基礎断面形状
（StbSecFoundation_RC_Continuous）のtype が「REVERSE_T」で、主筋の切り替えが無い
場合およびpos が配力筋や横筋の場合はmain_type は記述しないものとする。 
(5)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にあ
る場合は、省略してもよい。 
(6)  主筋はピッチ、配力筋・横筋は本数で指定する。 
X 
主筋上端 
主筋下端 
配力筋上端 
配力筋下端 
横筋 
Y 
Z 
始端基準点と終端基準点を結んだ軸線 
下端 
上端 
配力筋上端 
横筋 
主筋先端上端 
主筋元端上端 
主筋元端下端 
主筋先端下端 
配力筋下端 
横筋 
Y 
主筋の 
鉄筋切替位置 
length 
length 
鉄筋切替位置 
主筋の 
X 
Y 
main_type= SYMMETRICAL 
main_type=ASYMMETRICAL

## Page 465

6.18-23 
 
  <StbSecFoundation_RC id="6" name="F6" （略）> 
    <StbSecFigureFoundation_RC> 
      <StbSecFoundation_RC_Continuous width="1000" depth_base="900"  
depth_tip="400"/> 
    </StbSecFigureFoundation_RC> 
    <StbSecBarArrangementFoundation_RC 
 depth_cover_top="50" depth_cover_bottom="60" depth_cover_side="50"> 
      <StbSecBarFoundation_RC_Continuous pos="MAIN_BASE_TOP" D="D13" pitch="100"  
length="300" main_type="ASYMMETRICAL" isVertical="true"/> 
      <StbSecBarFoundation_RC_Continuous pos="MAIN_TIP_TOP" D="D13" pitch="200"  
length="300" main_type="ASYMMETRICAL" isVertical="true"/> 
      <StbSecBarFoundation_RC_Continuous pos="MAIN_BASE_BOTTOM" D="D13" 
     pitch="200"/> 
      <StbSecBarFoundation_RC_Continuous pos="MAIN_TIP_BOTTOM" D="D13" 
         pitch="200"/> 
      <StbSecBarFoundation_RC_Continuous pos="TRANSVERSE_TOP" D="D13"  
pitch="200"/> 
<StbSecBarFoundation_RC_Continuous pos="TRANSVERSE_BOTTOM" D="D13"  
pitch="200"/> 
      < StbSecBarFoundation_RC_Continuous  pos="HORIZONTAL" D="D10" N="3"/> 
    </StbSecBarArrangementFoundation_RC> 
  </StbSecFoundation_RC> 
(7) isVertical は外周の縦筋に関する設定で、pos が上端筋の場合は立ち下げる（true）か否（false）
か、pos が下端筋の場合は立ち上げる（true）か否（false）かを指定する。省略された場合は
false とする。なお、pos が横筋の場合は省略する。6.18.2.1 の補足(3)表を参照のこと。 
(8)  length_vertical はisVertical がtrue の場合に、立ち下げ長さもしくは立ち上げ長さを指定す
る。isVertical がtrue でlength_vertical が省略された場合は、上端筋の場合は下端筋まで、下
端筋の場合は上端筋までとする。isVertical がfalse もしくは省略された場合は、length_vertical
も省略する。6.18.2.1 の補足(3)表を参照のこと。 
 
・例

## Page 466

6.19-1 
 
6.19. ＲＣ杭断面：StbSecPile_RC 
・概要 
説明 ：ＲＣ杭断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_RC_Conventional 
1 
1 
一般工法 
※(1) 
または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_RC_Certified 
1 
制限なし 
認定工法 
※(2) 
 
 
・補足 
(1) アースドリル工法、リバース工法、オールケーシング工法、BH 工法、深礎工法の一般工法のみを
指定し、特定の工法を指定しない場合に使用する。 
(2) 特定のメーカーしか施工できない認定工法（旧大臣認定工法を含む）を指定する。同等品として複
数の工法を選択する場合は、最大回数を2 以上とする。 
 
・例

## Page 467

6.19-2 
 
6.19.1. ＲＣ杭断面 一般工法：StbSecPile_RC_Conventional 
・概要 
説明 ：RC 杭断面 一般工法 
親要素：StbSecPile_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
construction_method 
string 
 
工法 以下のいずれか 
EARTHDRILL(アースドリル工法) 
REVERSE(リバース工法) 
ALLCASING(オールケーシング工法) 
BH(BH 工法) 
SHINSO(深礎工法) 
 
strength_concrete 
string 
 
コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecFigurePile_RC_Conventional 
1 
1 
ＲＣ杭断面形状：一般工法 
 
StbSecBarArrangementPile_RC_Conventional 
0 
1 
ＲＣ杭断面配筋：一般工法 
 
StbSecPile_RC_Connection 
0 
1 
RC 杭 杭頭接合 
 
 
 
・補足 
(1)  省略された場合は、共通情報の属性「建物全体のコンクリート強度」strength_concrete をこの
要素のコンクリート強度とする。

## Page 468

6.19-3 
 
6.19.1.1. ＲＣ杭断面形状 一般工法：StbSecFigurePile_RC_Conventional 
・概要 
説明 ：ＲＣ杭断面形状（一般工法） 
親要素：StbSecPile_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
length_pipe 
double 
 
鋼管部長さ 
※(1) 
t_pipe 
double 
 
鋼管の厚さ 
※(1) 
strength_pipe 
string 
 
鋼管の鉄骨強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_RC_ConventionalStraight 
1 
1 
ＲＣ杭断面形状・ストレート 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_RC_ConventionalExtendedFoot 
1 
1 
ＲＣ杭断面形状・脚部拡大 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_RC_ConventionalExtendedTop 
1 
1 
ＲＣ杭断面形状・頂部拡大 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecPile_RC_ConventionalExtendedTopFoot 
1 
1 
ＲＣ杭断面形状・両端拡大 
 
 
・補足 
(1)  鋼管巻き場所打ち杭の場合のみ、記述する。記述する場合は、全属性を指定し、一部を省略する
ことはできない。

## Page 469

6.19-4 
 
 
「鋼管部長さ」は、杭頭から鋼管下端部までの長さとする（下図）。

## Page 470

6.19-5 
 
 
 
6.19.1.1.1. ＲＣ杭断面形状・ストレート：StbSecPile_RC_ConventionalStraight 
・概要 
説明 ：ＲＣ杭断面の形状（頂部と脚部の径が同一） 
親要素：StbSecFigurePile_RC_Conventional  
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
double 
○ 
杭径 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「杭径」の定義は下図による。 
 
 
 
 
X
Y
Z
杭径

## Page 471

6.19-6 
 
6.19.1.1.2. ＲＣ杭断面形状・脚部拡大：StbSecPile_RC_ConventionalExtendedFoot 
・概要 
説明 ：ＲＣ杭断面の形状（脚部の径が大きい） 
親要素：StbSecFigurePile_RC_Conventional  
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_axial 
double 
○ 
軸径 
 
D_extended_foot 
double 
○ 
拡底径 
 
length_extended_foot 
double 
○ 
拡底部の立ち上がり長さ 
 
angle_extended_foot_taper 
double 
○ 
拡底部の傾斜角度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「軸径」「拡底径」「拡底部の立ち上がり長さ」および「拡底部の傾斜角度」の定義は下図による。 
           
 
 
 
X
Y
Z
拡底径
軸径

## Page 472

6.19-7 
 
6.19.1.1.3. ＲＣ杭断面形状・頂部拡大：StbSecPile_RC_ConventionalExtendedTop 
・概要 
説明 ：ＲＣ杭断面の形状（頂部の径が大きい） 
親要素：StbSecFigurePile_RC_Conventional  
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_extended_top 
double 
○ 
拡頭径 
 
D_axial 
double 
○ 
軸径 
 
angle_extended_top_taper 
double 
○ 
拡頭部のテーパー角度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「拡頭径」「軸径」および「拡頭部のテーパー角度」の定義は下図による。 
 
          
 
 
 
X
Y
Z
拡頭径
軸径

## Page 473

6.19-8 
 
6.19.1.1.4. ＲＣ杭断面形状・頂部脚部拡大：StbSecPile_RC_ConventionalExtendedTopFoot 
・概要 
説明 ：ＲＣ杭断面の形状（両端が軸径よりも大きい） 
親要素：StbSecFigurePile_RC_Conventional  
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_extended_top 
double 
○ 
拡頭径 
 
D_axial 
double 
○ 
軸径 
 
D_extended_foot 
double 
○ 
拡底径 
 
angle_extended_top_taper 
double 
○ 
拡頭部のテーパー角度 
 
length_extended_foot 
double 
○ 
拡底部の立ち上がり長さ 
 
angle_extended_foot_taper 
double 
○ 
拡底部の傾斜角度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「拡頭径」「軸径」および「拡底径」の定義は下図による。 
「拡頭部のテーパー角度」「拡底部の立ち上がり長さ」および「拡底部の傾斜角度」の定義は頂部
拡大、脚部拡大の図による。 
 
X
Y
Z
拡頭径
軸径
拡底径

## Page 474

6.19-9 
 
6.19.1.2. ＲＣ杭断面配筋：StbSecBarArrangementPile_RC_Conventional 
・概要 
説明 ：ＲＣ杭断面の配筋 
親要素：StbSecPile_RC_Conventional  
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover 
double 
 
かぶり厚さ 
 
depth_cover_top 
double 
 
拡頭部かぶり厚さ 
拡頭杭の場合 
isSpiral 
boolean 
 
帯筋がスパイラルか否か 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarPile_RC_Same 
1 
1 
ＲＣ杭断面配筋・全断面 
 
  または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarPile_RC_TopBottom 
2 
2 
ＲＣ杭断面配筋・杭頭脚別 
 
  または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarPile_RC_TopCenterBottom 
3 
3 
ＲＣ杭断面配筋・杭頭軸部杭脚

## Page 475

6.19-10 
 
 
・補足 
「かぶり厚さ」の定義は下図による。省略された場合の扱いは、<StbApplyConditionsList>の補
足説明による。 
 
 
(1)  帯筋がスパイラルのとき true とし、省略された場合は、false とする。 
 
 
X
Y
かぶり厚さ

## Page 476

6.19-11 
 
6.19.1.2.1. ＲＣ杭断面配筋・全断面：StbSecBarPile_RC_Same 
・概要 
説明 ：ＲＣ杭断面の配筋（全断面同一の場合） 
親要素：StbSecBarArrangementPile_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_main 
string 
○ 
主筋：径 
 
D_2nd_main 
string 
 
束ね筋：径 
 
D_core 
string 
 
芯筋：径 
※(1) 
D_band 
string 
○ 
帯筋：径 
 
D_core_band 
string 
 
芯帯筋：径 
 
strength_main 
string 
 
主筋：鉄筋強度 
※(2) 
strength_2nd_main 
string 
 
束ね筋：鉄筋強度 
strength_core 
string 
 
芯筋：鉄筋強度 
strength_band 
string 
 
帯筋：鉄筋強度 
strength_core_band 
string 
 
芯帯筋：径 
 
N_main 
integer 
○ 
主筋：本数 
 
N_2nd_main 
integer 
 
束ね筋：本数 
 
N_core 
integer 
 
芯筋：本数 
 
pitch_band 
double 
○ 
帯筋：ピッチ 
 
pitch_core_band 
double 
〇 
芯帯筋：ピッチ 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 477

6.19-12 
 
 
・補足 
主筋、芯筋および帯筋の定義は下図による。 
 
 
 
(1)  芯筋は、位置に関する情報を有しない。芯筋がある場合は、芯筋に関する属性をすべて記述す
る。 
(2)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
 
 
主筋 
束ね筋 
帯筋 
芯帯筋 
芯筋

## Page 478

6.19-13 
 
6.19.1.2.2. ＲＣ杭断面配筋・杭頭脚別：StbSecBarPile_RC_TopBottom 
・概要 
説明 ：ＲＣ杭断面の配筋（杭頭・杭脚が別配筋の場合） 
親要素：StbSecBarArrangementPile_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置以下のいずれか 
TOP（杭頭） 
BOTTOM（杭脚） 
 
D_main 
string 
○ 
主筋：径 
 
D_2nd_main 
string 
 
束ね筋：径 
 
D_core 
string 
 
芯筋：径 
※(1) 
D_band 
string 
○ 
帯筋：径 
 
D_core_band 
string 
 
芯帯筋：径 
 
strength_main 
string 
 
主筋：鉄筋強度 
※(2) 
strength_2nd_main 
string 
 
束ね筋：鉄筋強度 
 
strength_core 
string 
 
芯筋：鉄筋強度 
 
strength_band 
string 
 
帯筋：鉄筋強度 
 
strength_core_band 
string 
 
芯帯筋：径 
 
N_main 
integer 
○ 
主筋：本数 
 
N_2nd_main 
integer 
 
束ね筋：本数 
 
N_core 
integer 
 
芯筋：本数 
 
pitch_band 
double 
○ 
帯筋：ピッチ 
 
pitch_core_band 
double 
〇 
芯帯筋：ピッチ 
 
length_bar 
double 
 
配筋長さ 
※(3) 
length_lap_bar 
double 
 
重ね継手長さ 
※(3) 
 
・内容 
無し 
 
・子要素 
無し

## Page 479

6.19-14 
 
 
・補足 
杭頭、杭脚配筋について、属性をそれぞれpos= “TOP” およびpos= “BOTTOM” としたこの子要素を
各１回記述する。定義と補足内容 (1)～(2) は、「ＲＣ杭断面配筋・全断面」による。 
 
 
 
(3)  配筋長さは以下とする。 
杭頭pos= “TOP” の場合・・杭頭位置から杭脚配筋に切替わる位置までの長さ 
（重ね継手長さを含まない） 
杭脚pos= “BOTTOM” の場合・・上記杭頭の配筋長さ位置から杭先端の鉄筋位置までの長さ 
 
重ね継手長さは、杭頭pos= “TOP” の場合のみ指定し、杭頭に近い方側の配筋で定義する。 
 
 
X
Y
杭頭
杭脚
主筋
帯筋
芯筋
主筋
帯筋
芯筋

## Page 480

6.19-15 
 
6.19.1.2.3. ＲＣ杭断面配筋・杭頭軸部杭脚：StbSecBarPile_RC_TopCenterBottom 
・概要 
説明 ：ＲＣ杭断面の配筋（杭頭・軸部・杭脚が別配筋の場合） 
親要素：StbSecBarArrangementPile_RC_Conventional 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置以下のいずれか 
TOP（杭頭） 
CENTER（軸部） 
BOTTOM（杭脚） 
 
D_main 
string 
○ 
主筋：径 
 
D_2nd_main 
string 
 
束ね筋：径 
 
D_core 
string 
 
芯筋：径 
※(1) 
D_band 
string 
○ 
帯筋：径 
 
D_core_band 
string 
 
芯帯筋：径 
 
※(2) 
 
strength_main 
string 
 
主筋：鉄筋強度 
strength_2nd_main 
string 
 
束ね筋：鉄筋強度 
strength_core 
string 
 
芯筋：鉄筋強度 
strength_band 
string 
 
帯筋：鉄筋強度 
 
strength_core_band 
string 
 
芯帯筋：径 
 
N_main 
integer 
○ 
主筋：本数 
 
N_2nd_main 
integer 
 
束ね筋：本数 
 
N_core 
integer 
 
芯筋：本数 
 
pitch_band 
double 
○ 
帯筋：ピッチ 
 
pitch_core_band 
double 
〇 
芯帯筋：ピッチ 
 
length_bar 
double 
 
配筋長さ 
※(3) 
length_lap_bar 
double 
 
重ね継手長さ 
※(3) 
 
・内容 
無し 
 
・子要素 
無し

## Page 481

6.19-16 
 
 
・補足 
杭頭、軸部および杭脚配筋について、属性をそれぞれpos= “TOP” , “CENTER” , “BOTTOM” とした
この子要素を各１回記述する。定義と補足内容 (1)～(2) は、「ＲＣ杭断面配筋・全断面」による。 
 
 
 
 
(3)  配筋長さは以下とする。 
杭頭pos= “TOP” の場合・・杭頭位置から軸部配筋に切替わる位置までの長さ 
             （軸部との重ね継手長さを含まない） 
軸部pos= “CENTER” の場合・・上記、杭頭の配筋長さ位置から杭脚配筋に切替わる 
位置までの長さ（杭脚との重ね継手長さを含まない） 
杭脚pos= “BOTTOM” の場合・・上記、軸部の配筋長さ位置から杭先端の鉄筋位置 
までの長さ 
  
重ね継手長さは、杭頭pos= “TOP” の場合および軸部pos= “CENTER” の場合に指定し、杭頭
に近い方側の配筋で定義する。 
X
Y
杭脚主筋
杭脚芯筋
杭頭主筋
杭頭芯筋
杭頭
杭脚
軸部主筋
軸部芯筋
軸部
帯筋
帯筋
帯筋
杭頭主筋(2段目）
杭頭束ね筋

## Page 482

6.19-17 
 
6.19.1.3. StbSecPile_RC_Connection：場所打ちコンクリート杭 杭頭接合 
・概要 
説明 ：場所打ちコンクリート杭 杭頭接合 
親要素：StbSecPile_RC_Conventional、StbSecPile_RC_Certified 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecPile_RC_ConnectionCertified 
1 
1 
杭頭接合 認定工法 
 
 
 
・補足

## Page 483

6.19-18 
 
6.19.1.3.1. StbSecPile_RC_ConnectionCertified：場所打ちコンクリート杭 杭頭接合 認定工法 
・概要 
説明 ：場所打ちコンクリート杭 杭頭接合 認定工法 
親要素：StbSecPile_RC_Connection 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
〇 
杭頭接合 工法名 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCertifiedMethodKeyValue 
0 
制限なし 
認定工法特有の属性と値 
※(2) 
 
 
・補足 
(1) 製品一覧表に記載した工法名を記載する。 
(2) 工法特有の属性値が製品一覧表に記載があり、その値を指定する場合は列挙する。

## Page 484

6.19-19 
 
6.19.1.3.1.1. 認定工法特有の属性と値：StbCertifiedMethodKeyValue 
・概要 
説明 ：認定工法特有の属性と値 
親要素：StbSecPile_RC_ConnectionCertified、StbSecPile_RC_Certified、Stb 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
属性名 
型 
必須 
説明 
補足 
key 
string 
〇 
属性名 
※(1) 
value 
string 
〇 
値 
※(1) 
 
 
・補足 
(1) 工法一覧表に記載のある工法名と紐づいた属性とその値を列挙する。

## Page 485

6.19-20 
 
6.19.2. RC 杭断面 認定工法：StbSecPile_RC_Certified  
・概要 
説明 ：認定工法 RC 杭断面 
親要素：StbSecPile_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
〇 
認定工法名 
 
strength_concrete 
string 
 
コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigurePile_RC_Certified 
1 
1 
ＲＣ杭断面形状：認定工法 
 
StbSecBarArrangementPile_RC_Certified 
0 
1 
ＲＣ杭断面配筋：認定工法 
 
StbCertifiedMethodKeyValue 
0 
制限なし 
属性の追加 
※(2) 
StbCertificationNumber 
0 
制限なし 
認定番号 
※(3) 
StbSecPile_RC_Connection 
0 
1 
RC 杭 杭頭接合 
 
 
・補足 
(1) 省略された場合は、共通情報の属性「建物全体のコンクリート強度」strength_concrete をこの要
素のコンクリート強度とする。 
(2) 工法特有の属性値が製品一覧表に記載があり、その値を指定する場合は列挙する。 
(3) 該当する認定番号が複数ある場合は列挙することができる。

## Page 486

6.19-21 
 
6.19.2.1. ＲＣ杭断面形状 認定工法：StbSecFigurePile_RC_Certified 
・概要 
説明 ：ＲＣ杭断面形状（認定工法） 
親要素：StbSecPile_RC_Certified 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
double 
○ 
杭径 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 487

6.19-22 
 
6.19.2.2. ＲＣ杭断面配筋：StbSecBarArrangementPile_RC_Certified 
・概要 
説明 ：ＲＣ杭断面の配筋 
親要素：StbSecPile_RC_Certified  
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover 
double 
 
かぶり厚さ 
 
isSpiral 
boolean 
 
帯筋がスパイラルか否か 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarPile_RC_Same 
1 
1 
ＲＣ杭断面配筋・全断面 
 
  または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarPile_RC_TopBottom 
2 
2 
ＲＣ杭断面配筋・杭頭脚別 
 
  または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarPile_RC_TopCenterBottom 
3 
3 
ＲＣ杭断面配筋・杭頭軸部杭脚

## Page 488

6.20-1 
 
6.20. 鋼管杭断面：StbSecPile_S 
・概要 
説明 ：鋼管杭断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_S_Conventional 
1 
1 
鋼管杭：一般工法 
 
または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_S_Certified 
1 
制限なし 
鋼管杭：認定工法 
※(1) 
 
 
・補足 
(1)  特定のメーカーしか施工できない認定工法（旧大臣認定工法を含む）を指定する。同等品として
複数の工法を選択する場合は、最大回数を2 以上とする。 
 
・例 
 
 
  <StbSecPile_S id="235" name="P2"> 
    <StbSecPile_S_Conventional> 
      <StbSecFigurePile_S> 
        <StbSecPile_S_Straight id_order="1" 
              length_pile="13000" D="600" t="12" strength="SKK400"/> 
        <StbSecPile_S_Straight id_order="2" 
              length_pile="10000" D="600" t="12" strength="SKK400"/> 
      </StbSecFigurePile_S> 
    </StbSecPile_S_Conventional> 
  </StbSecPile_S>

## Page 489

6.20-2 
 
6.20.1. 鋼管杭 一般工法：StbSecPile_S_Conventional 
・概要 
説明 ：鋼管杭断面の形状 
親要素：StbSecPile_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
construction_method 
string 
 
工法 以下のいずれか 
DRIVING(打ち込み杭) 
BURIED(埋込杭) 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigurePile_S 
1 
1 
鋼管杭断面形状 
 
StbSecPile_S_Joint 
0 
制限なし 
継手 
※(1) 
StbSecPile_S_Connection 
0 
1 
杭頭接合 
 
 
・補足 
(1)  子要素の最大回数は継杭の本数-1 とする。

## Page 490

6.20-3 
 
6.20.1.1. 鋼管杭断面形状：StbSecFigurePile_S 
・概要 
説明 ：鋼管杭断面の形状 
親要素：StbSecPile_S_Conventional、StbSecPile_S_Certified 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_S_Straight 
0 
制限なし 
鋼管杭断面形状・ストレート 
※(1) 
StbSecPile_S_Rotational 
0 
制限なし 
鋼管杭断面形状・回転貫入杭
（先端拡翼杭） 
※(1) 
StbSecPile_S_Taper 
0 
制限なし 
鋼管杭断面形状・テーパー管杭 
※(1) 
StbSecPile_S_Product 
0 
制限なし 
鋼管杭断面形状・製品指定 
※(1) 
 
・補足 
(1)  上記のいずれか１種類以上の子要素を持つものとし、子要素の回数は、継杭の本数分とする。 
(2)  全子要素の最小回数が0 であってはならない。 
(3)  子要素の並びは、上表に示す順番としなければならない。

## Page 491

6.20-4 
 
6.20.1.1.1. 鋼管杭断面形状・ストレート：StbSecPile_S_Straight 
・概要 
説明 ：鋼管杭断面の形状（ストレート杭） 
親要素：StbSecFigurePile_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
D 
double 
○ 
軸部径 
 
t 
double 
○ 
鋼管の厚さ 
 
strength 
string 
○ 
鋼管の鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。

## Page 492

6.20-5 
 
6.20.1.1.2. 鋼管杭断面形状・回転貫入杭（先端拡翼杭）：StbSecPile_S_Rotational 
・概要 
説明 ：鋼管杭断面の形状（回転貫入杭） 
親要素：StbSecFigurePile_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
D1 
double 
○ 
軸部径 
 
D2 
double 
○ 
先端拡翼径 
 
t 
double 
○ 
鋼管の厚さ 
 
strength 
string 
○ 
鋼管の鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。

## Page 493

6.20-6 
 
6.20.1.1.3. 鋼管杭断面形状・テーパー管杭：StbSecPile_S_Taper 
・概要 
説明 ：鋼管杭断面の形状（テーパー管杭） 
親要素：StbSecFigurePile_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
D1 
double 
○ 
上部径 
 
D2 
double 
○ 
下部径 
 
t 
double 
○ 
鋼管の厚さ 
 
strength 
string 
○ 
鋼管の鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  「製品型番」「メーカー名」は、形状寸法および材料強度にて仕様が特定できる場合は、省略し
てもよい。

## Page 494

6.20-7 
 
6.20.1.1.4. 鋼管杭断面形状・製品指定：StbSecPile_S_Product 
・概要 
説明 ：鋼管杭断面の形状製品指定 
親要素：StbSecFigurePile_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
product_code 
string 
○ 
製品型番 
※(2) 
release_time 
string 
〇 
リリース時期 
※(2) 
length_pile 
double 
○ 
杭の長さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  製品型番一覧表の鋼管杭カテゴリから「製品型番」と「リリース時期」の2 つのキーで製品を一
意に指定する。

## Page 495

6.20-8 
 
6.20.1.2. 鋼管杭継手：StbSecPile_S_Joint 
・概要 
説明 ：継手 
親要素：StbSecPile_S_Conventional、StbSecPile_S_Certified 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継手位置 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_S_JointWeld 
1 
1 
溶接継手 
 
  または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPile_S_JointMecanical 
1 
1 
機械式継手 
 
 
 
・補足 
(1)  「継手位置」は、上部の継杭を指し示すものとし、第1 杭と第2 杭の間に位置する継手は ”1” 、
のように記述する。

## Page 496

6.20-9 
 
6.20.1.2.1. 鋼管杭溶接継手：StbSecPile_S_JointWeld 
・概要 
説明 ：鋼管杭溶接継手 
親要素：StbSecPile_S_Joint 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
 
溶接工法名 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  溶接工法を指定する場合に記入する。

## Page 497

6.20-10 
 
6.20.1.2.2. 鋼管杭機械式継手：StbSecPile_S_JointMecanical 
・概要 
説明 ：鋼管杭機械式継手 
親要素：StbSecPile_S_Joint 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
〇 
製品型番 
※(1) 
release_time 
string 
〇 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  製品型番一覧表の機械式継手(鋼管杭)カテゴリから「製品型番」と「リリース時期」の2 つのキ
ーで製品を一意に指定する。

## Page 498

6.20-11 
 
6.20.1.3. 鋼管杭杭頭接合：StbSecPile_S_Connection 
 
・概要 
説明 ：鋼管杭杭頭接合 
親要素：StbSecPile_S_Conventional、StbSecPile_S_Certified 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_pile_head 
string 
 
杭頭補強筋の径 
 
N_pile_head 
integer 
 
杭頭補強筋の本数 
 
stength_pile_head 
string 
 
杭頭補強筋の強度 
 
D_inner 
string 
 
中詰め鉄筋径 
 
N_inner 
integer 
 
中詰め鉄筋本数 
 
strength_inner 
string 
 
中詰め鉄筋強度 
 
strength_inner_concrete 
string 
 
中詰めコンクリート 
 
concrete_depth 
double 
 
中詰めコンクリート深さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 499

6.20-12 
 
6.20.2. 鋼管杭 認定工法：StbSecPile_S_Certified 
・概要 
説明 ：鋼管杭断面の形状 
親要素：StbSecPile_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
〇 
認定工法名 
※(1) 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigurePile_S 
1 
1 
鋼管杭断面の形状 
 
StbCertifiedMethodKeyValue 
0 
制限なし 
属性の追加 
※(2) 
StbCertificationNumber 
0 
制限なし 
認定番号 
※(3) 
StbSecPile_S_Joint 
0 
制限なし 
継手 
※(4) 
StbSecPile_S_Connection 
0 
1 
杭頭補強 
 
 
・補足 
(1)  工法一覧表の鋼管杭カテゴリに記載した工法名を記載する。 
(2)  工法特有の属性値が工法一覧表に記載があり、その値を指定する場合は列挙する。 
(3)  該当する認定番号が複数ある場合は列挙することができる。 
(4)  子要素の最大回数は継杭の本数-1 とする。

## Page 500

6.21-1 
 
  <StbSecPilePrecast id="236" name="P3"> 
    <StbSecFigurePilePrecast> 
      <StbSecPilePrecast_PHC id_order="1" product_code="xxxxxx" 
              length_pile="12000" kind="C" D="600" t="90"/> 
      <StbSecPilePrecast_PHC id_order="2" product_code="xxxxxx" 
              length_pile="10000" kind="A" D="600" t="90"/> 
   </StbSecFigurePilePrecast> 
  </StbSecPilePrecast> 
6.21. 既製コンクリート杭断面：StbSecPilePrecast 
・概要 
説明 ：既製コンクリート杭断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPilePrecastConventional 
1 
1 
既製コンクリート杭:一般工法 
 
または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPilePrecastCertified 
1 
制限なし 
既製コンクリート杭：認定工法 
 
 
・補足 
無し 
 
・例

## Page 501

6.21-2 
 
6.21.1. 既製コンクリート 一般工法：StbSecPilePrecastConventional 
・概要 
説明 ：既製コンクリート杭断面の形状 
親要素：StbSecPilePrecast 
・属性 
属性名 
型 
必須 
説明 
補足 
construction_method 
string 
 
工法 以下のいずれか 
DRIVING(打ち込み杭) 
BURIED(埋込杭) 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigurePilePrecast 
1 
1 
既成コンクリート杭断面形状 
 
StbSecPilePrecastJoint 
0 
制限なし 
継手 
※(1) 
StbSecPilePrecastConnection 
0 
1 
杭頭接合 
 
 
 
・補足 
(1) 子要素の最大回数は継杭の本数-1 とする。

## Page 502

6.21-3 
 
6.21.1.1. 既成コンクリート杭断面形状：StbSecFigurePilePrecast 
・概要 
説明 ：既製コンクリート杭断面の形状 
親要素：StbSecPilePrecast 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPilePrecast_PHC 
0 
制限なし 
既製コンクリート杭断面形状・PHC 杭 
 
StbSecPilePrecast_ST 
0 
制限なし 
既製コンクリート杭断面形状・ST 杭 
 
StbSecPilePrecast_SC 
0 
制限なし 
既製コンクリート杭断面形状・SC 杭 
 
StbSecPilePrecast_PRC 
0 
制限なし 
既製コンクリート杭断面形状・PRC 杭 
 
StbSecPilePrecast_CPRC 
0 
制限なし 
既製コンクリート杭断面形状 
・CPRC 杭 
 
StbSecPilePrecastNodular_PHC 
0 
制限なし 
既製コンクリート杭断面形状 
・節付PHC 杭 
 
StbSecPilePrecastNodular_PRC 
0 
制限なし 
既製コンクリート杭断面形状 
・節付PRC 杭 
 
StbSecPilePrecastNodular_CPRC 
0 
制限なし 
既製コンクリート杭断面形状 
・節付CPRC 杭 
 
StbSecPilePrecastProduct 
0 
制限なし 
既製コンクリート杭断面形状 製品指
定 
 
 
・補足 
上記のいずれか１種類以上の子要素を持つものとし、子要素の回数は、継杭の本数分とする。 
全子要素の最小回数が0 であってはならない。 
子要素の並びは、上表に示す順番としなければならない。

## Page 503

6.21-4 
 
6.21.1.1.1. 既製コンクリート杭断面形状・PHC 杭：StbSecPilePrecast_PHC 
・概要 
説明 ：既製コンクリート杭断面の形状（PHC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
○ 
種類 
※(2) 
D 
double 
○ 
外径 
 
t 
double 
○ 
厚さ 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
 
・内容 
無し 
 
・子要素 
無し 
length_pile
D
t
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件に応じた種別（A 種であれば ”A” ）を記述する。

## Page 504

6.21-5 
 
6.21.1.1.2. 既製コンクリート杭断面形状・ST 杭：StbSecPilePrecast_ST 
・概要 
説明 ：既製コンクリート杭断面の形状（ST 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
○ 
種類 
※(2) 
D1 
double 
○ 
外径（本体部） 
 
D2 
double 
○ 
外径（拡径部） 
 
t1 
double 
○ 
厚さ（本体部） 
 
t2 
double 
○ 
厚さ（拡径部） 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
 
・内容 
無し 
 
・子要素 
無し 
length_pile
D2
D1
t1
t2
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件に応じた種別（A 種であれば ”A” ）を記述する。

## Page 505

6.21-6 
 
6.21.1.1.3. 既製コンクリート杭断面形状・SC 杭：StbSecPilePrecast_SC 
・概要 
説明 ：既製コンクリート杭断面の形状（SC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
 
種類 
※(2) 
D 
double 
○ 
外径 
 
tc 
double 
○ 
肉厚(含鋼管) 
 
ts 
double 
○ 
鋼管の板厚 
 
strength_concrete 
string 
 
コンクリート強度 
 
strength_pipe 
string 
 
鋼管の鉄骨強度 
 
 
・内容 
無し 
 
・子要素 
無し 
D
tc
ts
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件などに応じた種別がある場合は記述する。

## Page 506

6.21-7 
 
6.21.1.1.4. 既製コンクリート杭断面形状・PRC 杭：StbSecPilePrecast_PRC 
・概要 
説明 ：既製コンクリート杭断面の形状（PRC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
 
種類 
※(2) 
D 
double 
○ 
外径 
 
tc 
double 
○ 
厚さ 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
D_bar 
string 
 
異形棒鋼径 
 
N_bar 
integer 
 
異形棒鋼本数 
 
strength_bar 
string 
 
異形棒鋼強度 
※(3) 
 
・内容 
無し 
 
・子要素 
無し 
length_pile
D
tc
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件などに応じた種別がある場合は記述する。 
(3)  対応する径が、共通情報の要素<StbReinforcementStrength> にある場合、省略してよい。

## Page 507

6.21-8 
 
6.21.1.1.5. 既製コンクリート杭断面形状・CPRC 杭：StbSecPilePrecast_CPRC 
・概要 
説明 ：既製コンクリート杭断面の形状（CPRC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
 
種類 
※(2) 
D 
double 
○ 
外径 
 
tc 
double 
○ 
厚さ 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
D_bar 
string 
 
異形棒鋼径 
 
N_bar 
integer 
 
異形棒鋼本数 
 
strength_bar 
string 
 
異形棒鋼強度 
※(3) 
 
・内容 
無し 
 
・子要素 
無し 
length_pile
D
tc
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件などに応じた種別がある場合は記述する。 
(3)  対応する径が、共通情報の要素<StbReinforcementStrength> にある場合、省略してよい。

## Page 508

6.21-9 
 
6.21.1.1.6. 既製コンクリート杭断面形状・節付PHC 杭：StbSecPilePrecastNodular_PHC 
・概要 
説明 ：既製コンクリート杭断面の形状（節付PHC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
 
種類 
※(2) 
D1 
double 
○ 
外径（軸部） 
 
D2 
double 
○ 
外形（節部） 
 
t 
double 
○ 
厚さ 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
 
・内容 
無し 
 
・子要素 
無し 
length_pile
D2
D1
tc
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件などに応じた種別がある場合は記述する。

## Page 509

6.21-10 
 
6.21.1.1.7. 既製コンクリート杭断面形状・節付PRC 杭：StbSecPilePrecastNodular_PRC 
・概要 
説明 ：既製コンクリート杭断面の形状（節付PRC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
 
種類 
※(2) 
D1 
double 
○ 
外径（軸部） 
 
D2 
double 
○ 
外形（節部） 
 
tc 
double 
○ 
厚さ 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
D_bar 
string 
 
異形棒鋼径 
 
N_bar 
integer 
 
異形棒鋼本数 
 
strength_bar 
string 
 
異形棒鋼強度 
※(3) 
 
・内容、子要素 
無し 
length_pile
D2
D1
tc
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件などに応じた種別がある場合は記述する。 
(3)  対応する径が、共通情報の要素<StbReinforcementStrength> にある場合、省略してよい。

## Page 510

6.21-11 
 
6.21.1.1.8. 既製コンクリート杭断面形状・節付CPRC 杭：StbSecPilePrecastNodular_CPRC 
・概要 
説明 ：既製コンクリート杭断面の形状（節付CPRC 杭） 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
継杭の位置 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
kind 
string 
 
種類 
※(2) 
D1 
double 
○ 
外径（軸部） 
 
D2 
double 
○ 
外形（節部） 
 
tc 
double 
○ 
厚さ 
 
strength_concrete 
string 
 
コンクリート強度 
 
D_PC 
double 
 
PC 鋼棒径 
 
N_PC 
integer 
 
PC 鋼棒本数 
 
strength_PC 
string 
 
PC 鋼棒強度 
 
D_bar 
string 
 
異形棒鋼径 
 
N_bar 
integer 
 
異形棒鋼本数 
 
strength_bar 
string 
 
異形棒鋼強度 
※(3) 
 
・内容、子要素 
無し 
length_pile
D2
D1
tc
 
・補足 
(1)  「継杭の位置」は、杭頭側から杭先端に向かって順に、第1 杭、第2 杭、第3 杭・・・とし、
第1 杭は ”1” 、第2 杭は ”2” ・・・のように記述する。 
(2)  強度条件などに応じた種別がある場合は記述する。 
(3)  対応する径が、共通情報の要素<StbReinforcementStrength> にある場合、省略してよい。

## Page 511

6.21-12 
 
6.21.1.1.9. 既成コンクリート杭断面形状・製品指定：StbSecPileConventionalPrecastProduct 
・概要 
説明 ：既製コンクリート杭断面の製品指定 
親要素：StbSecFigurePilePrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
integer 
○ 
製品型番 
※(1) 
release_time 
string 
○ 
リリース時期 
※(1) 
length_pile 
double 
○ 
杭の長さ 
 
top_length 
double 
 
上部と下部で異なる製品の上部長さ 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 製品型番一覧表の既製杭カテゴリから「製品型番」と「リリース時期」の2 つのキーで製品を一
意に指定する。 
(2) 一つの部材で断面が切り替わる場合に入力する。

## Page 512

6.21-13 
 
6.21.1.2. 継手：StbSecPilePrecastJoint 
・概要 
説明 ：既製コンクリート杭断面の製品指定 
親要素：StbSecPilePrecastConventional、StbSecPilePrecastCertified 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPilePrecastJointWeld 
1 
1 
溶接継手 
 
  または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPilePrecastJointMecanical 
1 
1 
機械式継手 
 
 
 
・補足 
(1)  「継手位置」は、上部の継杭を指し示すものとし、第1 杭と第2 杭の間に位置する継手は ”1” 、
のように記述する。

## Page 513

6.21-14 
 
6.21.1.2.1. 溶接継手：StbSecPilePrecastJointWeld 
・概要 
説明 ：既製コンクリート杭溶接継手 
親要素：StbSecPilePrecastJoint 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
 
溶接工法名 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 溶接工法を指定する場合に記入する。

## Page 514

6.21-15 
 
6.21.1.2.2. 鋼管杭機械式継手：StbSecPilePrecastJointMecanical 
・概要 
説明 ：鋼管杭機械式継手 
親要素：StbSecPileJointPrecast 
 
・属性 
属性名 
型 
必須 
説明 
補足 
product_code 
string 
〇 
製品型番 
※(1) 
release_time 
string 
〇 
リリース時期 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 製品型番一覧表の機械式継手(既製杭)カテゴリから「製品型番」と「リリース時期」の2 つのキー
で製品を一意に指定する。

## Page 515

6.21-16 
 
6.21.1.3. 既製杭 杭頭接合：StbSecPilePrecastConnection 
・概要 
説明 ：既製杭 杭頭接合 
親要素：StbSecPilePrecastConventional、StbSecPilePrecastCertified 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecPilePrecastConnectionConventional 
1 
1 
杭頭補強 在来工法 
 
または 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecPilePrecastConnectionCertified 
1 
1 
杭頭補強 認定工法 
 
 
 
・補足

## Page 516

6.21-17 
 
6.21.1.3.1. 既製杭 杭頭接合 在来工法：StbSecPilePrecastConnectionConventional 
・概要 
説明 ：既製杭 杭頭接合 
親要素：StbSecPilePrecastConnection 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D_pile_head 
string 
 
杭頭補強筋の径 
 
N_pile_head 
integer 
 
杭頭補強筋の本数 
 
stength_pile_head 
string 
 
杭頭補強筋の強度 
 
D_inner 
string 
 
中詰めの鉄筋径 
 
N_inner 
integer 
 
中詰めの鉄筋本数 
 
strength_inner 
string 
 
中詰めの鉄筋強度 
 
strength_inner_concrete 
string 
 
中詰めコンクリート 
 
concrete_depth 
double 
 
中詰めコンクリート深さ 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 517

6.21-18 
 
6.21.1.3.2. 既製杭 杭頭接合 認定工法：StbSecPilePrecastConnectionCertified 
・概要 
説明 ：既製杭 杭頭接合 認定工法 
親要素：StbSecPilePrecastConnection 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
〇 
杭頭接合 工法名 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCertifiedMethodKeyValue 
0 
制限なし 
認定工法特有の属性と値 
※(2) 
 
 
・補足 
(1) 工法一覧表の杭頭接合(既製杭)カテゴリに記載した工法名を記載する。 
(2) 工法特有の属性値が製品一覧表に記載があり、その値を指定する場合は列挙する。

## Page 518

6.21-19 
 
6.21.2. 既製コンクリート杭 認定工法：StbSecPilePrecastCertified 
・概要 
説明 ：既製コンクリート杭：認定工法 
親要素：StbSecPilePrecast 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
〇 
認定工法 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigurePilePrecast 
1 
1 
 
 
StbCertifiedMethodKeyValue 
0 
制限なし 
認定工法特有の属性
と値 
※(2) 
StbCertificationNumber 
0 
制限なし 
認定番号 
※(3) 
StbSecPilePrecastJoint 
0 
制限なし 
継手 
※(4) 
StbSecPilePrecastConnection 
0 
1 
杭頭接合 
 
 
 
・補足 
(1) 工法一覧表の既製杭カテゴリに記載した工法名を記載する。 
(2) 工法特有の属性値が製品一覧表に記載があり、その値を指定する場合は列挙する。 
(3) 該当する認定番号が複数ある場合は列挙することができる。 
(4) 子要素の最大回数は継杭の本数-1 とする。

## Page 519

6.22-1 
 
6.22. ＲＣパラペット断面：StbSecParapet_RC 
・概要 
説明 ：ＲＣパラペット断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
strength_concrete 
string 
 
コンクリート強度 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigureParapet_RC 
1 
1 
ＲＣパラペット断面形状 
 
StbSecBarArrangementParapet_RC 
0 
1 
ＲＣパラペット断面配筋 
※(2) 
 
・補足 
(1)  省略された場合は、参照する<StbParapet>の「始端節点ID」id_node_start が「所属階」と一
致する<StbStory>のコンクリート強度を、この要素のコンクリート強度とする。参照した
<StbStory>のコンクリート強度が省略されていた場合は、共通情報の属性「建物全体のコンクリ
ート強度」strength_concrete をこの要素のコンクリート強度とする。 
(2)  子要素<StbSecBarArrangementParapet_RC>の回数が0 となる場合は、鉄筋を扱わないプロ
グラムが一時的に作成する場合を想定しており、無筋であることを示すものではない。

## Page 520

6.22-2 
 
・例 
 
 
  <StbSecParapet_RC id="393" name="P1" strength_concrete="FC21"> 
    <StbSecFigureParapet_RC> 
      <StbSecParapet_RC_TypeL t_T="200" depth_H="850" 
 t_T1="200" depth_H1="250" depth_H2="0"/> 
    </StbSecFigureParapet_RC> 
    <StbSecBarArrangementParapet_RC> 
      <StbSecBarParapet_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarParapet_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementParapet_RC> 
  </StbSecParapet_RC>

## Page 521

6.22-3 
 
6.22.1. ＲＣパラペット断面形状：StbSecFigureParapet_RC 
・概要 
説明 ：ＲＣパラペット断面の形状 
親要素：StbSecParapet_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecParapet_RC_TypeL 
1 
1 
ＲＣパラペット断面形状・L 型 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecParapet_RC_TypeT 
1 
1 
ＲＣパラペット断面形状・T 型 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecParapet_RC_TypeI 
1 
1 
ＲＣパラペット断面形状・I 型 
 
 
・補足 
 
・例 
 
 
  <StbSecParapet_RC id="393" name="P1" strength_concrete="FC21"> 
    <StbSecFigureParapet_RC> 
      <StbSecParapet_RC_TypeL t_T="200" depth_H="850" 
 t_T1="200" depth_H1="250" depth_H2="0"/> 
    </StbSecFigureParapet_RC> 
    <StbSecBarArrangementParapet_RC> 
      <StbSecBarParapet_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarParapet_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementParapet_RC> 
  </StbSecParapet_RC>

## Page 522

6.22-4 
 
6.22.1.1. ＲＣパラペット断面形状・L 型：StbSecParapet_RC_TypeL 
・概要 
説明 ：ＲＣパラペット断面の形状（L 型） 
親要素：StbSecFigureParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
t_T 
double 
○ 
厚さT 
 
depth_H 
double 
○ 
高さH 
 
t_T1 
double 
○ 
寸法T1 
 
depth_H1 
double 
○ 
寸法H1 
 
depth_H2 
double 
○ 
寸法H2 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各寸法の定義は右図による。 
 
 
 
 
・例 
 
 
 
  <StbSecParapet_RC id="393" name="P1" （略）> 
    <StbSecFigureParapet_RC> 
      <StbSecParapet_RC_TypeL t_T="200" depth_H="850" 
 t_T1="200" depth_H1="250" depth_H2="0"/> 
    </StbSecFigureParapet_RC> 
    <StbSecBarArrangementParapet_RC （略）> 
        （略） 
    </StbSecBarArrangementParapet_RC> 
  </StbSecParapet_RC>

## Page 523

6.22-5 
 
6.22.1.2. ＲＣパラペット断面形状・T 型：StbSecParapet_RC_TypeT 
・概要 
説明 ：ＲＣパラペット断面の形状（T 型） 
親要素：StbSecFigureParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
t_T 
double 
○ 
厚さT 
 
depth_H 
double 
○ 
高さH 
 
t_T1 
double 
○ 
寸法T1 
 
depth_H1 
double 
○ 
寸法H1 
 
depth_H2 
double 
○ 
寸法H2 
 
depth_H3 
double 
○ 
寸法H3 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各寸法の定義は右図による。

## Page 524

6.22-6 
 
6.22.1.3. ＲＣパラペット断面形状・I 型：StbSecParapet_RC_TypeI 
・概要 
説明 ：ＲＣパラペット断面の形状（I 型） 
親要素：StbSecFigureParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
t_T 
double 
○ 
厚さT 
 
depth_H 
double 
○ 
高さH 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各寸法の定義は右図による。

## Page 525

6.22-7 
 
6.22.2. ＲＣパラペット断面配筋：StbSecBarArrangementParapet_RC 
・概要 
説明 ：ＲＣパラペット断面の配筋 
親要素：StbSecParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
depth_cover_outside 
double 
 
かぶり厚さ（外）  ※アゴがない側 
※(1) 
depth_cover_inside 
double 
 
かぶり厚さ（内）  ※アゴがある側 
isTipline 
boolean 
 
垂下の鉄筋の有無 
※(2) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarParapet_RC_Single 
2 
2 
ＲＣパラペット断面配筋・シングル 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarParapet_RC_Zigzag 
2 
2 
ＲＣパラペット断面配筋・千鳥 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarParapet_RC_DoubleNet 
2 
2 
ＲＣパラペット断面配筋・ダブル 
 
 
上記子要素に、以下を追加してもよい。 
  要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarParapet_RC_Tip 
1 
2 
パラペット先端補強筋（アゴ筋） 
※(3) 
StbSecBarParapet_RC_Edge 
1 
4 
パラペット端部補強筋 
※(3)

## Page 526

6.22-8 
 
・補足 
(1)  かぶり厚さの定義は、下図による。配筋がシングルのときは省略してよい。それ以外で省略され
た場合の扱いは、<StbApplyConditionsList>の補足説明による。ただし、該当属性がない場合は、
省略してはならない。 
 
 
(2)  true の場合、「あり」とし、省略された場合は、true とする。 
(3)  パラペット端部補強筋は、該当する補強筋が存在する場合に、追加で記述する。 
 
・例 
 
 
 
かぶり厚さ
かぶり厚さ
Ｚ
Y
ダブル配筋
ダブル配筋（内外異なる）
かぶり厚さ（外）
かぶり厚さ（内）
かぶり厚さ
千鳥配筋
かぶり厚さ
  <StbSecParapet_RC id="393" name="P1"  （略）> 
        （略） 
    <StbSecBarArrangementParapet_RC 
 depth_cover_outside="30" depth_cover_inside="30"> 
      <StbSecBarParapet_RC_DoubleNet pos="VERTICAL" D="D10" pitch="200"/> 
      <StbSecBarParapet_RC_DoubleNet pos="HORIZONTAL" D="D10" pitch="200"/> 
    </StbSecBarArrangementParapet_RC> 
  </StbSecParapet_RC>

## Page 527

6.22-9 
 
6.22.2.1. ＲＣパラペット断面配筋・シングル：StbSecBarParapet_RC_Single 
・概要 
説明 ：ＲＣパラペット断面の配筋（シングル配筋） 
親要素：StbSecBarArrangementParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
縦筋、横筋について、属性をそれぞれpos= “VERTICAL” およびpos= “HORIZONTAL” としたこの
子要素を各１回記述する。 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
 
 
 
横筋（ピッチ）
Ｚ
Y

## Page 528

6.22-10 
 
6.22.2.2. ＲＣパラペット断面配筋・千鳥：StbSecBarParapet_RC_Zigzag 
・概要 
説明 ：ＲＣパラペット断面の配筋（千鳥配筋） 
親要素：StbSecBarArrangementParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
縦筋、横筋について、属性をそれぞれpos= “VERTICAL” およびpos= “HORIZONTAL” としたこの
子要素を各１回記述する。 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
 
 
 
横筋（ピッチ）
Ｚ
Y

## Page 529

6.22-11 
 
6.22.2.3. ＲＣパラペット断面配筋・ダブル：StbSecBarParapet_RC_DoubleNet 
・概要 
説明 ：ＲＣパラペット断面の配筋（ダブル配筋） 
親要素：StbSecBarArrangementParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
○ 
ピッチ 
 
 
・内容、子要素 
無し 
 
・補足 
縦筋、横筋について、属性をそれぞれpos= “VERTICAL” およびpos= “HORIZONTAL” としたこの
子要素を各１回記述する。 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
 
 
 
横筋（ピッチ）
Ｚ
Y

## Page 530

6.22-12 
 
6.22.2.4. パラペット先端補強筋（アゴ筋）：StbSecBarParapet_RC_Tip 
・概要 
説明 ：ＲＣパラペット断面の配筋（パラペット先端の補強筋） 
親要素：StbSecBarArrangementParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
SHORT_SIDE（短辺方向） 
LONG_SIDE（長辺方向） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
pitch 
double 
 
ピッチ 
 
N 
integer 
 
本数 
 
 
・内容、子要素 
無し 
 
・補足 
下図「アゴ筋」において、短辺方向はピッチを、長辺方向は本数を、必要に応じて、位置ごとにこの子
要素として各１回記述する。補強筋がその位置にない場合はこの要素を省略してよい。 
 
 
 
 
 
 
 
 
 
 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
 
 
アゴ筋 短辺方向 
アゴ筋 長辺方向

## Page 531

6.22-13 
 
6.22.2.5. 端部補強筋：StbSecBarParapet_RC_Edge 
・概要 
説明 ：ＲＣパラペット断面の配筋（パラペット端部の補強筋） 
親要素：StbSecBarArrangementParapet_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL_START（パラペット始端） 
VERTICAL_END（パラペット終端） 
HORIZONTAL_TOP（パラペット上端） 
HORIZONTAL_BOTTOM（パラペット下端） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
 
 
・内容、子要素 
無し 
 
・補足 
必要に応じて、補強筋のある位置ごとにこの子要素を各１回記述する。補強筋がその位置にない場合
はこの要素を省略してよい。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。

## Page 532

6.23-1 
 
6.23. ＲＣ開口断面：StbSecOpen_RC 
・概要 
説明 ：ＲＣ開口（補強筋のリスト） 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
※(1) 
guid 
string 
 
GUID 
 
name 
string 
 
断面名称 
 
length_X 
double 
〇 
開口寸法(X) 
 
length_Y 
double 
〇 
開口寸法(Y) 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarArrangementOpen_RC 
0 
1 
ＲＣ開口断面配筋 
※(2) 
 
 
・補足 
(1) このID の参照元は、開口<StbOpenArrangement>とする。ここで示す開口配筋は、参照された
開口を有する壁<StbWall>およびスラブ<StbSlab>に適用される。 
(2) 適用される壁およびスラブは、参照するRC 壁断面ならびにRC スラブ断面も開口配筋を有して
いる場合があるが、子要素の回数が1 の場合にはこちらの配筋情報が優先されるものとする。 
 
 
・例 
 
 
  <StbSecOpen_RC id="84" name="SD-11" length_X=”500”, length_Y=”500”> 
  <StbSecBarArrangementOpen_RC> 
<StbSecBarOpen_RC_Wall pos=”VERTICAL” D=”D16” N=”2”/> 
<StbSecBarOpen_RC_Wall pos=”HORIZONTAL” D=”D16” N=”2”/> 
<StbSecBarOpen_RC_Wall pos=”DIAGONAL” D=”D13” N=”1”/> 
</StbSecBarArrangementOpen_RC> 
  </StbSecOpen_RC>

## Page 533

6.23-2

## Page 534

6.23-3 
 
6.23.1. ＲＣ開口断面配筋：StbSecBarArrangementOpen_RC 
・概要 
説明 ：ＲＣ開口の補強筋 
親要素：StbSecOpen_RC 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarOpen_RC_Slab 
0 
6 
ＲＣスラブ開口配筋 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecBarOpen_RC_Wall 
0 
3 
ＲＣ壁開口配筋 
 
 
・補足 
(1) スラブ、壁ともに補強筋のある位置ごとにこの子要素を記述する。最小回数が0 の場合は、開
口が存在しながら開口補強筋がない状態を示す。

## Page 535

6.23-4 
 
6.23.1.1. ＲＣスラブ開口配筋：StbSecBarOpen_RC_Slab 
・概要 
説明 ：スラブ開口配筋 
親要素：StbSecBarArrangementOpen_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
X_TOP（X 方向上端） 
X_BOTTOM（X 方向下端） 
Y_TOP（Y 方向上端） 
Y_BOTTOM（Y 方向下端） 
DIAGONAL_TOP（斜め方向上端） 
DIAGONAL_BOTTOM（斜め方向下端） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
※(2) 
length 
double 
 
長さ 
 
 
・内容、子要素 
無し 
 
・補足 
必要に応じて、補強筋のある位置ごとにこの子要素を各１回記述する。補強筋がその位置にない場合
はこの要素を省略してよい。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
(2)  「本数」は、１辺および１隅あたりの合計本数とする。 
 
 
（1点目）
（2点目）
（3点目）
（4点目）
開口
X
Y
Y方向上端（下端）開口補強筋
X方向上端（下端）開口補強筋
斜め方向上端（下端）開口補強筋

## Page 536

6.23-5 
 
6.23.1.2. ＲＣ壁開口配筋：StbSecBarOpen_RC_Wall 
・概要 
説明 ：壁開口配筋（壁開口ごとの配筋） 
親要素：StbSecBarArrangementOpen_RC 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pos 
string 
○ 
配筋位置 以下のいずれか 
VERTICAL（縦筋） 
HORIZONTAL（横筋） 
DIAGONAL（斜め筋） 
 
strength 
string 
 
鉄筋強度 
※(1) 
D 
string 
○ 
径 
 
N 
integer 
○ 
本数 
※(2) 
length 
double 
 
長さ 
 
 
・内容、子要素 
無し 
 
・補足 
必要に応じて、補強筋のある位置ごとにこの子要素を各１回記述する。補強筋がその位置にない場合
はこの要素を省略してよい。 
 
(1)  「鉄筋強度」は、それぞれ対応する径が、共通情報の要素<StbReinforcementStrength> にある
場合は、省略してもよい。 
(2)  「本数」は、１辺および１隅あたりの合計本数とする。 
（1点目）
（2点目）
（3点目）
（4点目）
開口
X
Y
縦筋
横筋
斜め筋

## Page 537

6.24-1 
 
6.24. S 梁貫通孔補強仕様：StbSecPenetration_S 
・概要 
説明 ：鉄骨梁貫通孔補強仕様 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
phi 
double 
〇 
貫通孔径 
 
type 
string 
〇 
補強工法 
以下のいずれか 
EG（EG リング） 
FR（フリードーナツ） 
HI（ハイリング） 
OS（OS リング） 
NO（無補強） 
CO（在来補強） 
 
model 
string 
 
リング型式名 
 
phi_ring 
double 
 
リング外径 
 
protrude 
double 
 
リングのWeb 面からの突出寸法 
※(1) 
clear 
double 
 
リング縁からの必要クリアランス 
※(2)、※(3) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補
足 
StbSecPenetration_S_HiringExtension 
0 
1 
ハイリング拡張情報 
 
StbSecPenetration_S_EGringExtension 
0 
1 
EG リング拡張情報 
 
StbSecPenetration_S_FRringExtension 
0 
1 
フリードーナツ拡張情報 
 
StbSecPenetration_S_OSringExtension 
0 
1 
OS リング拡張情報

## Page 538

6.24-2 
 
・補足 
(1) 表裏で突出寸法が異なる場合は、大きい方の値とする。 
(2) ガセットPL、縦リブ、水平リブ等とのクリアランス。 
(3) ハイリングの場合はスティック縁とリング縁からのクリアランスが同じであるため、スティック
のみの場合もリング縁からのクリアの値を入力する（Ｇ寸法）。

## Page 539

6.24-3 
 
6.24.1. ハイリング拡張情報：StbSecPenetration_S_HiringExtension 
・概要 
説明 ：ハイリングの拡張情報 
親要素：StbSecPenetration_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
length_stick 
double 
 
スティック長さ 
 
wide_stick 
double 
 
スティックの幅 
 
protruce_stick 
double 
 
スティックのウェブ面からの突出寸法 
 
Lh_stick 
double 
 
スティック取り付け寸法 
梁天端からスティッ
ク天端までの距離 
clear_stick 
double 
 
スティック端からフランジスプライス端
までの必要クリアランス 
 
phi_bottom 
double 
〇 
下孔径 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 540

6.24-4 
 
6.24.2. EG リング拡張情報：StbSecPenetration_S_EGringExtension 
・概要 
説明 ：EG リングの拡張情報 
親要素：StbSecPenetration_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
phi_bottom 
double 
〇 
下孔径 
 
installation 
integer 
〇 
設置段 
上からの段数 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 541

6.24-5 
 
6.24.3. フリードーナツ拡張情報：StbSecPenetration_S_FRringExtension 
・概要 
説明 ：フリードーナツの拡張情報 
親要素：StbSecPenetration_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
phi_bottom 
double 
〇 
下孔径 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 542

6.24-6 
 
6.24.4. OS リング拡張情報：StbSecPenetration_S_OSringExtension 
・概要 
説明 ：OS リングの拡張情報 
親要素：StbSecPenetration_S 
 
・属性 
属性名 
型 
必須 
説明 
補足 
clear_splice 
double 
〇 
リング縁からウェブスプライス端までの
必要クリアランス 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) ＯＳリングはウェブスプライスとそれ以外のクリアランスの値が異なるため、ウェブスプライス
端までの必要クリアランスを入力する。

## Page 543

6.25-1 
 
6.25. 柱梁接合部断面：StbSecPanelZone 
・概要 
説明 ：柱梁接合部断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
断面名称 
 
N_hoop_X 
integer 
○ 
帯筋：X 方向本数 
 
D_hoop_X 
string 
○ 
帯筋：X 方向径 
 
strength_X 
string 
 
帯筋：X 方向強度 
 
N_hoop_Y 
integer 
○ 
帯筋：Y 方向本数 
 
D_hoop_Y 
string 
○ 
帯筋：Y 方向径 
 
strength_Y 
string 
 
帯筋：Y 方向強度 
 
pitch_hoop 
double 
○ 
帯筋：ピッチ 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecFigurePanelZone 
1 
1 
柱梁接合部断面形状 
 
 
・補足 
無し

## Page 544

6.25-2 
 
6.25.1. 柱梁接合部断面形状：StbSecFigurePanelZone 
・概要 
説明 ：柱梁接合部断面形状 
親要素：StbSecPanelZone 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPanelZoneRect 
1 
1 
柱梁接合部断面形状・矩形 
 
   または 
要素名 
最小回数 
最大回数 
説明 
補足 
StbSecPanelZoneCircle 
1 
1 
柱梁接合部断面形状・円形 
 
 
・補足 
無し

## Page 545

6.25-3 
 
6.25.1.1. 柱梁接合部断面形状・矩形：StbSecPanelZoneRect 
・概要 
説明 ：柱梁接合部断面形状・矩形 
親要素：StbSecFigurePanelZone 
 
・属性 
属性名 
型 
必須 
説明 
補足 
width_X 
double 
○ 
Ｘ幅 
 
width_Y 
double 
○ 
Ｙ幅 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 546

6.25-4 
 
6.25.1.2. 柱梁接合部断面形状・円形：StbSecPanelZoneCircle 
・概要 
説明 ：柱梁接合部断面形状・円形 
親要素：StbSecFigurePanelZone 
 
・属性 
属性名 
型 
必須 
説明 
補足 
D 
double 
○ 
直径 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 547

6.26-1 
 
6.26. 鉄骨断面：StbSecSteel 
・概要 
説明 ：鉄骨断面 
親要素：StbSections 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小
回数 
最大回数 
説明 
補足 
StbSecRoll-H 
0 
制限なし 
H 形鋼 
 
StbSecBuild-H 
0 
制限なし 
組立H 形鋼 
 
StbSecBuild-HAssymmetric 
0 
制限なし 
組立H 形鋼(上下異なる) 
 
StbSecRoll-BOX 
0 
制限なし 
角形鋼管 
 
StbSecBuild-BOX 
0 
制限なし 
組立角形鋼管 
 
StbSecPipe 
0 
制限なし 
円形鋼管 
 
StbSecRoll-T 
0 
制限なし 
T 形鋼 
 
StbSecBuild-T 
0 
制限なし 
組立T 形鋼 
 
StbSecRoll-C 
0 
制限なし 
溝形鋼 
 
StbSecRoll-2C 
0 
制限なし 
溝形鋼(2 丁) 
 
StbSecRoll-L 
0 
制限なし 
山形鋼 
 
StbSecRoll-2L 
0 
制限なし 
山形鋼(2 丁) 
 
StbSecLipC 
0 
制限なし 
リップ溝形鋼 
 
StbSecLip2C 
0 
制限なし 
リップ溝形鋼(2 丁) 
 
StbSecFlatBar 
0 
制限なし 
フラットバー 
 
StbSecRoundBar 
0 
制限なし 
丸鋼 
 
StbSecSteelProduct 
0 
制限なし 
鉄骨製品 
 
StbSecSteelUndefined 
0 
制限なし 
未定義鉄骨断面 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。 
鉄骨断面の定義は、ここに掲げた各子要素内で記述する。柱断面、梁断面などの要素（断面要素）で
は、属性「鉄骨形状」で、各子要素の属性「形状名」の文字列を参照する。

## Page 548

6.26-2 
 
断面要素の属性「鉄骨形状」の文字列において、規格品における呼び方などを用いて、形状を特定で
きるとみなされる場合は、<StbSecSteel> の各子要素は省略してもよい。 
参照時における断面の向き（方向）は、各子要素について、原則として属性「成」の方向を「鉄骨断面
の基準方向」とし、断面要素で参照する際の基準とする。 
 
・例 
  <StbSecSteel> 
    <StbSecRoll-H name=”H-700x300x13x24x18” type=”H” A=”700” B=”300”  
t1=”13” t2=”24” r=”18” /> 
    <StbSecRoll-BOX name=”□-400x400x22x22” type=”BCP” A=”400” B=”400”  
t=”22” r=”77” /> 
  </StbSecSteel>

## Page 549

6.26-3 
 
6.26.1. H 形鋼：StbSecRoll-H 
・概要 
説明 ：H 形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
type 
string 
○ 
形状タイプ 
以下のいずれか 
H（一般H 形鋼） 
SH（外法一定H 形鋼） 
 
A 
double 
○ 
成 
 
B 
double 
○ 
フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2 
double 
○ 
フランジ厚 
 
r 
double 
○ 
フィレット半径 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
 
各属性の定義は下図による。 
  
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
 
 
鉄骨断面の
基準方向

## Page 550

6.26-4 
 
6.26.2. 組立H 形鋼：StbSecBuild-H 
・概要 
説明 ：組立H 形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
A 
double 
○ 
成 
 
B 
double 
○ 
フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2 
double 
○ 
フランジ厚 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性の定義はH 形鋼の図を参照。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 551

6.26-5 
 
6.26.3. 組立H 形鋼（上下異なる）：StbSecBuild-HAssymmetric 
・概要 
説明 ：組立H 形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
A 
double 
○ 
成 
 
B_T 
double 
○ 
上部フランジ幅 
 
B_B 
double 
○ 
下部フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2_T 
double 
○ 
上部フランジ厚 
 
t2_B 
double 
○ 
下部フランジ厚 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
 
各属性の定義は下図による。 
 
 
(1) 親要素 <StbSecSteel> 内で、一意な名称とする。 
(2) 上下のフランジが等しい場合には、StbSecBuild-H を用いる。 
A 
B_B 
B_T 
t2_T 
t2_B 
t1 
鉄骨断面の
基準方向

## Page 552

6.26-6 
 
6.26.4. 角形鋼管：StbSecRoll-BOX 
・概要 
説明 ：角形鋼管 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
type 
string 
○ 
形状タイプ 
以下のいずれか 
BCP、BCR、STKR、ELSE 
 
A 
double 
○ 
成 
 
B 
double 
○ 
幅 
 
t 
double 
○ 
板厚 
 
r 
double 
○ 
コーナー半径(R) 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
 
各属性の定義は下図による。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 553

6.26-7 
 
6.26.5. 組立角形鋼管：StbSecBuild-BOX 
・概要 
説明 ：組立角形鋼管 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
A 
double 
○ 
成 
 
B 
double 
○ 
幅 
 
t1 
double 
○ 
成方向の板厚 
 
t2 
double 
○ 
幅方向の板厚 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
    「成」方向を「鉄骨断面の基準方向」とする。 
 
各属性の定義は下図による。 
 
 
 
 
 
 
 
 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
 
 
A 
B 
t1 
t2

## Page 554

6.26-8 
 
6.26.6. 円形鋼管：StbSecPipe 
・概要 
説明 ：円形鋼管 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
D 
double 
○ 
直径 
 
t 
double 
○ 
板厚 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性の定義は下図による。 
 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 555

6.26-9 
 
6.26.7. T 形鋼：StbSecRoll-T 
・概要 
説明 ：T 形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
type 
string 
○ 
形状タイプ 
以下のいずれか 
T（一般T 形鋼） 
ST（外法一定T 形鋼） 
 
A 
double 
○ 
成 
 
B 
double 
○ 
フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2 
double 
○ 
フランジ厚 
 
r 
double 
○ 
フィレット半径 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
    「成」方向を「鉄骨断面の基準方向」とする。 
 
各属性の定義は下図による。 
 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
t1 
A 
B 
t2 
r 
鉄骨断面の
基準方向

## Page 556

6.26-10 
 
6.26.8. 組立T 形鋼：StbSecBuild-T 
・概要 
説明 ：T 形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
A 
double 
○ 
成 
 
B 
double 
○ 
フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2 
double 
○ 
フランジ厚 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性の定義はT 形鋼の図を参照。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 557

6.26-11 
 
6.26.9. 溝形鋼：StbSecRoll-C 
・概要 
説明 ：溝形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
A 
double 
○ 
成 
 
B 
double 
○ 
フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2 
double 
○ 
フランジ厚 
 
r1 
double 
○ 
フィレット半径 
 
r2 
double 
○ 
フランジ先端半径 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
各属性の定義は右図による。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。  
(2)  右図上のように進行方向に向かって右側が開いた形状とする。

## Page 558

6.26-12 
 
6.26.10. 溝形鋼(2 丁)：StbSecRoll-2C 
・概要 
説明 ：溝形鋼(2 丁) 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
type 
string 
○ 
形状タイプ 
以下のいずれか 
背中合わせ：BACKTOBACK 
表合わせ：FACETOFACE 
※(2) 
A 
double 
○ 
成 
 
B 
double 
○ 
フランジ幅 
 
t1 
double 
○ 
ウェブ厚 
 
t2 
double 
○ 
フランジ厚 
 
r1 
double 
○ 
フィレット半径 
 
r2 
double 
○ 
フランジ先端半径 
 
gap 
double 
 
ギャップ 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
各属性の定義は右図による。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
(2)  背中合わせ、表合わせは右図下とする。 
 
 
 
表合わせ 
背中合わせ

## Page 559

6.26-13 
 
6.26.11. 山形鋼：StbSecRoll-L 
・概要 
説明 ：山形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
A 
double 
○ 
成 
不等辺山形鋼：A＞B 
B 
double 
○ 
幅 
 
t1 
double 
○ 
成方向の板厚 
不等厚山形鋼： 
t1＜t2 
t2 
double 
○ 
幅方向の板厚 
 
r1 
double 
○ 
フィレット半径 
 
r2 
double 
○ 
先端半径 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
各属性の定義は右図による。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
(2)  右図上のように進行方向に向かって右側が開いた 
 形状とする。

## Page 560

6.26-14 
 
6.26.12. 山形鋼(2 丁)：StbSecRoll-2L 
・概要 
説明 ：山形鋼(2 丁) 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
type 
string 
○ 
形状タイプ 
以下のいずれかの値をとる。 
背中合わせ：BACKTOBACK 
表合わせ：FACETOFACE 
※(2) 
A 
double 
○ 
成 
不等辺山形鋼：A＞B 
B 
double 
○ 
幅 
 
t1 
double 
○ 
成方向の板厚 
不等厚山形鋼： 
t1＜t2 
t2 
double 
○ 
幅方向の板厚 
 
r1 
double 
○ 
フィレット半径 
 
r2 
double 
○ 
先端半径 
 
gap 
double 
 
ギャップ 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。 
各属性の定義は右図による。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
(2)  背中合わせ、表合わせは右図下とする。 
 
 
 
 
背中合わせ 
表合わせ

## Page 561

6.26-15 
 
6.26.13. リップ溝形鋼：StbSecLipC 
・概要 
説明 ：リップ溝形鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
H 
double 
○ 
成 
 
A 
double 
○ 
幅 
 
C 
double 
○ 
リップ長 
 
t 
double 
○ 
板厚 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。各属性の定義は右図による。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。 
(2)  右図上のように進行方向に向かって右側が開いた形状とする。

## Page 562

6.26-16 
 
6.26.14. リップ溝形鋼(2 丁)：StbSecLip2C 
 
・概要 
説明 ：リップ溝形鋼(2 丁) 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
type 
string 
○ 
形状タイプ 
以下のいずれか 
背中合わせ：BACKTOBACK 
表合わせ：FACETOFACE 
※(2) 
H 
double 
○ 
成 
 
A 
double 
○ 
幅 
 
C 
double 
○ 
リップ長 
 
t 
double 
○ 
板厚 
 
gap 
double 
 
ギャップ 
 
 
・内容、子要素 
無し 
 
・補足 
「成」方向を「鉄骨断面の基準方向」とする。各属性の定義は右図による。 
 
(1)   親要素 <StbSecSteel> 内で、一意な名称とする。 
(2)  背中合わせ、表合わせは右図下とする。 
 
 
 
 
 
表合わせ 
背中合わせ

## Page 563

6.26-17 
 
6.26.15. フラットバー：StbSecFlatBar 
・概要 
説明 ：フラットバー 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
B 
double 
○ 
幅 
 
t 
double 
○ 
板厚 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
「幅」方向を「鉄骨断面の基準方向」とする。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 564

6.26-18 
 
6.26.16. 丸鋼：StbSecRoundBar 
・概要 
説明 ：丸鋼 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
R 
double 
○ 
直径 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 565

6.26-19 
 
6.26.17. 鉄骨製品：StbSecSteelProduct 
・概要 
説明 ：鉄骨製品 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
product_company 
string 
 
メーカー名 
 
product_name 
string 
 
製品名または種類 
 
product_code 
string 
○ 
製品型番 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 566

6.26-20 
 
6.26.18. 未定義鉄骨断面：StbSecSteelUndefined 
・概要 
説明 ：未定義鉄骨断面 
親要素：StbSecSteel 
 
・属性 
属性名 
型 
必須 
説明 
補足 
name 
string 
○ 
形状名 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
この要素は、構造計算プログラムが計算対象としない形状で、断面性能を直接指定する場合を想定
している。 
 
(1)  親要素 <StbSecSteel> 内で、一意な名称とする。

## Page 567

6.27-1 
 
6.27. 構造種別に依存しない断面：StbSecUndefined 
・概要 
説明 ：構造種別に依存しない断面 
親要素：StbSections 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
断面名称 
 
 
・内容 
無し 
 
・属性 
無し 
 
・内容 
 
無し 
 
・補足 
構造計算プログラムにおいて、形状を入力する際の補助材として、壁や床を構成する多角形の外周位
置に、形状も断面性能も設定しない、いわゆる「ダミー部材」を指定する場合がある。この要素ではこ
のような、構造種別に依存しない断面を記述する。 
参照する部材要素は、StbGirder、StbColumn を対象とし、参照する部材要素の構造種別
（kind_structure）は “UNDEFINED” としなくてはならない。

## Page 568

7-1 
 
7. 要素リファレンス 継手情報  
7.1. 継手情報：StbJoints 
・概要 
説明 ：継手情報 
親要素：StbModel 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointBeamShapeH 
0 
制限なし 
Ｓ梁継手・Ｈ形 
 
StbJointColumnShapeH 
0 
制限なし 
Ｓ柱継手・Ｈ形 
 
StbJointColumnShapeT 
0 
制限なし 
Ｓ柱継手・Ｔ形 
 
StbJointColumnShapeCross 
0 
制限なし 
Ｓ柱継手・＋形 
 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 569

7-2 
 
7.1.1. Ｓ梁継手・Ｈ形：StbJointBeamShapeH 
・概要 
説明 ：Ｓ梁の継手・Ｈ形 
親要素：StbJoints 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
ID 
 
guid 
string 
 
GUID 
 
joint_name 
string 
 
継手呼称 
他のリストへ 
joint_mark 
string 
 
継手符号 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeH 
1 
1 
Ｈ形継手詳細 
 
StbJointShapeHFlange 
0 
1 
Ｈ形継手詳細・フランジ 
※(1) ※(2) 
StbJointShapeHWeb 
1 
1 
Ｈ形継手詳細・ウェブ 
※(2) 
 
・補足 
(1) 「StbJointShapeHFlange」を省略した場合、フランジ溶接とする。 
 
(2) 各子要素とスプライスプレート（添え板）の対応は下図の通りとする。 
 
  StbJointShapeHFlange 
  StbJointShapeHWeb

## Page 570

7-3 
 
7.1.1.1. Ｈ形継手詳細：StbJointShapeH 
・概要 
説明 ：Ｈ形鋼継手の詳細 
親要素：StbJointBeamShapeH、StbJointColumnShapeH 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength_plate_flange 
string 
○ 
添え板の材種（フランジ） 
例：SN490A 
strength_plate_web 
string 
 
添え板の材種（ウェブ） 
省略値はフランジに
同じ 例：SN490A 
strength_bolt 
string 
○ 
ボルト材種 
例：F10T 
name_bolt 
string 
○ 
ボルト径（呼び名） 
例：M22 
clearance 
double 
 
部材の母材間隔 
※(1) 
strength_filler 
string 
 
フィラープレートの材種（共
通） 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 省略値は、10 ㎜とする。 
(2) JASS6 の規定ではフィラープレートの材種は400N 級で良いことになっており、特別に指定する
場合のみ記述する。

## Page 571

7-4 
 
7.1.1.2. Ｈ形継手詳細・フランジ：StbJointShapeHFlange 
・概要 
説明 ：Ｈ形鋼フランジ添え板の寸法とボルト穴位置 
親要素：StbJointBeamShapeH、StbJointColumnShapeH 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isZigzag 
boolean 
 
千鳥配置か否か 
※(1) 
g1 
double 
○ 
ゲージ寸法1 (g1) 
※(2) 
g2 
double 
 
ゲージ寸法2 (g2) 
mf≧4 のとき必須 
※(2) 
pitch 
double 
○ 
長手方向のボルトピッチ (P) ※(2) 
e1 
double 
○ 
縁端距離1 (e1) 
※(2) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(2) ※(3) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe3 
内添え板 
※(2) 
outside_thickness 
double 
○ 
外添え板 厚さ  
 
outside_width 
double 
○ 
外添え板 幅(B) 
※(2) 
outside_length 
double 
 
外添え板 長さ(L) 
※(2) ※(3) 
inside_thickness 
double 
 
内添え板 厚さ 
 
inside_width 
double 
 
内添え板 幅 
※(2) 
inside_length 
double 
 
内添え板 長さ 
※(2) ※(3) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeHFlangeB
olt 
1 
制限なし 
Ｈ形継手フランジボルト詳細

## Page 572

7-5 
 
 
・補足 
(1) 省略値は、否とする。また、千鳥配置の場合は外スタートとする。 
(2) 各属性は、下図の通りとする。 
(3) 「縁端距離2 (e2)」の指定がある場合、「外添え板 長さ(L)」および「内添え板 長さ」は自動
決定するため、値を省略する。「外添え板 長さ(L)」および「内添え板 長さ」を指定する場
合、「縁端距離2 (e2)」の値を省略する。 
 
（フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
e2 
e1 
P 
P 
ｸﾘｱﾗﾝｽ 
L 
B 
g2 
g2 
g1 
e3 
e4 
e4 
e3

## Page 573

7-6 
 
7.1.1.2.1. Ｈ形継手詳細・フランジボルト詳細：StbJointShapeHFlangeBolt 
・概要 
説明 ：Ｈ形鋼フランジのボルト詳細 
親要素：StbJointShapeHFlange 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
継手中心からのボル
トの列数 
mf 
integer 
○ 
ボルトの数 
列当たりのボルトの
数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は、下図の通りとする。 
 
（フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
id_order 
m3 
3 
2 
1 
m2 
m1 
mf

## Page 574

7-7 
 
7.1.1.3. Ｈ形継手詳細・ウェブ：StbJointShapeHWeb 
・概要 
説明 ：Ｈ形鋼ウェブ添え板の寸法とボルト穴位置 
親要素：StbJointBeamShapeH、StbJointColumnShapeH 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
e1 
double 
○ 
縁端距離1 (e1) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe3 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeHWebBolt 
1 
制限なし 
Ｈ形継手ウェブボルト詳細

## Page 575

7-8 
 
・補足 
各属性は、下図の通りとする。 
 
（ウェブの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
mw 
L 
B 
e1 pL e2 
ｸﾘｱﾗﾝｽ 
pC 
pC 
pC 
e3 
e4

## Page 576

7-9 
 
7.1.1.3.1. Ｈ形継手詳細・ウェブボルト詳細：StbJointShapeHWebBolt 
・概要 
説明 ：Ｈ形鋼ウェブのボルト詳細 
親要素：StbJointShapeHWeb 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
継手中心からのボル
トの列数 
mw 
integer 
○ 
ボルトの数 
列当たりのボルトの
数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は、下図の通りとする。 
 
（ウェブの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
id_order 
2 
1 
m2 
m1 
mw

## Page 577

7-10 
 
・例 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
<StbJoints> 
<StbJointBeamShapeH id="1" joint_name="xxx" joint_mark="yyy"> 
<StbJointShapeH strength_plate_flange="SM490" strength_bolt="F10T" name_bolt="M22"/> 
<StbJointShapeHFlange 
isZigzag="true" g1="150" g2="40" pitch="45" e1="40" e3=”35” outside_thickness="19" 
outside_width="300" outside_length="710" inside_thickness="19" inside_width="110" 
inside_length="710"> 
<StbJointShapeHFlangeBolt id_order=”1” mf=”2”> 
<StbJointShapeHFlangeBolt id_order=”2” mf=”2”> 
･･･ 
<StbJointShapeHFlangeBolt id_order=”7” mf=”2”> 
</StbJointShapeHFlange> 
<StbJointShapeHWeb 
pitch_depth="90" pitch="60" e1="40" plate_thickness="16" plate_width="620" plate_length="290"> 
<StbJointShapeHWebBolt id_order=”1” mw=”7”> 
<StbJointShapeHWebBolt id_order=”2” mw=”7”> 
</StbJointShapeHWeb> 
</StbJointBeamShapeH> 
</StbJoints>

## Page 578

7-11 
 
7.1.2. Ｓ柱継手・Ｈ形：StbJointColumnShapeH 
・概要 
説明 ：Ｓ柱、ＳＲＣ柱の継手・Ｈ形 
親要素：StbJoints 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
ID 
 
guid 
string 
 
GUID 
 
joint_name 
string 
 
継手呼称 
他のリストへ 
joint_mark 
string 
 
継手符号 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeH 
1 
1 
Ｈ形継手詳細 
※(1) 
StbJointShapeHFlange 
0 
1 
Ｈ形継手詳細・フランジ 
※(1) ※(2) ※(3) 
StbJointShapeHWeb 
1 
1 
Ｈ形継手詳細・ウェブ 
※(1) ※(3) 
 
・補足 
(1) StbJointBeamShapeH と同じ。 
 
(2) 「StbJointShapeHFlange」を省略した場合、フランジ溶接とする。 
 
(3) 各子要素とスプライスプレートの対応は下図の通りとする。 
 
  StbJointShapeHFlange 
  StbJointShapeHWeb

## Page 579

7-12 
 
7.1.3. Ｓ柱継手・Ｔ形：StbJointColumnShapeT 
・概要 
説明 ：ＳＲＣ柱鉄骨部分の継手・Ｔ形 
親要素：StbJoints 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
ID 
 
guid 
string 
 
GUID 
 
joint_name 
string 
 
継手呼称 
他のリストへ 
joint_mark 
string 
 
継手符号 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeT 
1 
1 
Ｔ形継手詳細 
 
StbJointShapeTFlangeH 
0 
1 
Ｔ形継手詳細・Ｈ部分フランジ 
H 形鋼 
※(1) ※(2) 
StbJointShapeTWebHLong 
1 
1 
Ｔ形継手詳細・Ｈ部分ウェブ
(長) 
H 形鋼 長い方 
※(2) 
StbJointShapeTWebHShort 
1 
1 
Ｔ形継手詳細・Ｈ部分ウェブ
(短) 
H 形鋼 短い方 
※(2) 
StbJointShapeTFlangeT 
0 
1 
Ｔ形継手詳細・Ｔ部分フランジ 
T 形鋼 
※(1) ※(2) 
StbJointShapeTWebT 
1 
1 
Ｔ形継手詳細・Ｔ部分ウェブ 
T 形鋼 
※(2) 
 
・補足 
(1) 「StbJointShapeTFlangeH」、「StbJointShapeTFlangeT」を省略した場合、フランジ溶接とす
る。

## Page 580

7-13 
 
(2) 各子要素とスプライスプレートの対応は下図の通りとする。 
 
  StbJointShapeTFlangeH 
  StbJointShapeTWebHLong 
  StbJointShapeTWebHShort 
  StbJointShapeTFlangeT 
  StbJointShapeTWebT

## Page 581

7-14 
 
7.1.3.1. Ｔ形継手詳細：StbJointShapeT 
・概要 
説明 ：Ｔ形継手の詳細 
親要素：StbJointColumnShapeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength_plate_flange 
string 
○ 
添え板の材種（フランジ） 例：SN490A 
strength_plate_web 
string 
 
添え板の材種（ウェブ） 
省略値はフランジ
に同じ 
例：
SN490A 
strength_bolt 
string 
○ 
ボルト材種 
例：F10T 
name_bolt 
string 
○ 
ボルト径（呼名） 
例：M22 
offset_T 
double 
 
Ｔ形鋼の偏心（Ｈ形鋼の成
の中心からの距離） 
※(1) 
正の値とする 
clearance 
double 
 
部材の母材間隔 
※(2) 
strerngth_filler 
String 
 
フィラープレートの材種
（共通） 
※(3) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 省略値は、0 とする。 
(2) 省略値は、10 ㎜とする。 
(3) JASS6 の規定ではフィラープレートの材種は400N 級で良いことになっており、特別に指定する
場合のみ記述する。

## Page 582

7-15 
 
7.1.3.2. Ｔ形継手詳細・Ｈ部分フランジ：StbJointShapeTFlangeH 
・概要 
説明 ：Ｔ形継手のＨ形鋼部分フランジ添え板の寸法とボルト穴位置 
親要素：StbJointColumnShapeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isZigzag 
boolean 
 
千鳥配置か否か 
※(1) 
g1 
double 
○ 
ゲージ寸法1 (g1) 
※(2) 
g2 
double 
 
ゲージ寸法2 (g2) 
mf≧4 のとき必須 
※(2) 
pitch 
double 
○ 
長手方向のボルトピッチ(P) 
※(2) 
e1 
double 
○ 
縁端距離1 (e1) 
※(2) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(2) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(2) 
outside_thickness 
double 
○ 
外添え板 厚さ 
 
outside_width 
double 
○ 
外添え板 幅(B) 
※(2) 
outside_length 
double 
 
外添え板 長さ(L) 
※(2) ※(3) 
inside_thickness 
double 
 
内添え板 厚さ 
 
inside_width 
double 
 
内添え板 幅 
※(2) 
inside_length 
double 
 
内添え板 長さ 
※(2) ※(3) 
 
・内容 
無し

## Page 583

7-16 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeTFlangeB
olt 
1 
制限なし 
Ｔ形継手のフランジボルト詳細 
 
 
 
・補足 
(1) 省略値は、否とする。また、千鳥配置の場合は外スタートとする。 
(2) 各属性は、下図の通りとする。 
(3) 「縁端距離2 (e2)」の指定がある場合は、「外添え板 長さ(L)」および「内添え板 長さ」は自
動決定するため、値を省略する。「外添え板 長さ(L)」および「内添え板 長さ」を指定する
場合、「縁端距離2 (e2)」の値を省略する。 
 
    H 形鋼  
   （フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
e2 
e1 
P 
P 
ｸﾘｱﾗﾝｽ 
L 
B 
g2 
g2 
g1 
e3 
e4 
e4 
e3

## Page 584

7-17 
 
7.1.3.2.1. Ｔ形継手詳細・フランジボルト詳細：StbJointShapeTFlangeBolt 
概要 
説明 ：Ｔ形継手のフランジのボルト詳細 
親要素：StbJointShapeTFlangeH、StbJointShapeTFlangeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
継手中心からのボル
トの列数 
mf 
integer 
○ 
ボルトの数 
列当たりのボルトの
数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は、下図の通りとする。 
 
（フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
id_order 
m3 
3 
2 
1 
m2 
m1 
mf

## Page 585

7-18 
 
7.1.3.3. Ｔ形継手詳細・Ｈ部分ウェブ(長)：StbJointShapeTWebHLong 
・概要 
説明 ：Ｔ形継手のＨ形鋼部分ウェブ添え板（長い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
※(1) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
※(1) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(1) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
 
・内容 
無し

## Page 586

7-19 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeTWebBolt 
1 
制限なし 
Ｔ形継手のウェブボルト詳細 
 
・補足 
(1) 各属性は、下図の通りとする（HL 側）。 
 
    H 形鋼  
   （ウェブの継手） 
 
 
 
e3 
Offset＿T 
<HL 側> 
<HS 側> 
pC 
e2 
pL 
ｸﾘｱﾗﾝｽ 
e1 
L 
e4 
B 
B 
e5 
H 形鋼の中心

## Page 587

7-20 
 
7.1.3.4. Ｔ形継手詳細・ウェブボルト詳細：StbJointShapeTWebBolt 
概要 
説明 ：Ｔ形継手のウェブのボルト配列 
親要素：StbJointShapeTWebHLong、StbJointShapeTWebHShort、StbJointShapeTWebT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
継手中心からのボル
トの列数 
mw 
integer 
○ 
ボルトの数 
列当たりのボルトの
数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は、下図の通りとする。 
 
（ウェブの継手） 
 
 
 
m1 
m2 
mw 
id_order 
2 
1 
<HL 側> 
<HS 側> 
id_order 
m1 
m2 
2 
1 
mw

## Page 588

7-21 
 
7.1.3.5. Ｔ形継手詳細・Ｈ部分ウェブ(短)：StbJointShapeTWebHShort 
・概要 
説明 ：Ｔ形継手のＨ形鋼部分ウェブ添え板（短い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
※(1) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
※(1) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(1) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
※(1) 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
※(1) 
 
・内容 
無し

## Page 589

7-22 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeTWebBolt 
1 
制限なし 
Ｔ形継手のウェブボルト詳細 
 
 
 
・補足 
(1) 「7.1.3.3. Ｔ形継手詳細・Ｈ部分ウェブ(長)」の補足（HS 側）参照。

## Page 590

7-23 
 
7.1.3.6. Ｔ形継手詳細・Ｔ部分フランジ：StbJointShapeTFlangeT 
・概要 
説明 ：Ｔ形継手のＴ形鋼部分フランジ添え板の寸法とボルト穴位置 
親要素：StbJointColumnShapeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isZigzag 
boolean 
 
千鳥配置か否か 
※(1) 
g1 
double 
○ 
ゲージ寸法1 (g1) 
※(2) 
g2 
double 
 
ゲージ寸法2 (g2) 
mf≧4 のとき必須 
※(2) 
pitch 
double 
○ 
長手方向のボルトピッチ(P) 
※(2) 
e1 
double 
○ 
縁端距離1 (e1) 
※(2) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(2) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
※(2) 
outside_thickness 
double 
○ 
外添え板 厚さ 
 
outside_width 
double 
○ 
外添え板 幅(B) 
※(2) 
outside_length 
double 
 
外添え板 長さ(L) 
※(2) ※(3) 
inside_thickness 
double 
 
内添え板 厚さ 
 
inside_width 
double 
 
内添え板 幅 
 
inside_length 
double 
 
内添え板 長さ 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeTFlangeB
olt 
1 
制限なし 
Ｔ形継手のフランジボルト詳細

## Page 591

7-24 
 
・補足 
(1) 省略値は、否とする。また、千鳥配置の場合は外スタートとする。 
(2) 各属性は、下図の通りとする。 
(3) 「縁端距離2 (e2)」の指定がある場合は、「外添え板 長さ(L)」および「内添え板 長さ」は自
動決定するため、値を省略する。「外添え板 長さ(L)」および「内添え板 長さ」を指定する場
合、「縁端距離2 (e2)」の値を省略する。 
 
    T 形鋼  
   （フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
e2 
e1 
P 
P 
ｸﾘｱﾗﾝｽ 
L 
B 
g2 
g2 
g1 
e3 
e4 
e4 
e3

## Page 592

7-25 
 
7.1.3.7. Ｔ形継手詳細・Ｔ部分ウェブ：StbJointShapeTWebT 
・概要 
説明 ：Ｔ形継手のＴＨ形鋼部分ウェブ添え板の寸法とボルト穴位置 
親要素：StbJointColumnShapeT 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
※(1) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
※(1) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(1) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
※(1) 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
※(1) 
 
・内容 
無し

## Page 593

7-26 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeTWebBolt 
1 
制限なし 
Ｔ形継手のウェブボルト詳細 
 
 
・補足 
(1) 各属性は、下図の通りとする。 
 
    T 形鋼  
   （ウェブの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
B 
L 
e2 
pL 
ｸﾘｱﾗﾝｽ 
e1 
e3 
pC 
e4 
e5

## Page 594

7-27 
 
7.1.4. Ｓ柱継手・＋形：StbJointColumnShapeCross 
・概要 
説明 ：ＳＲＣ柱鉄骨部分の継手・＋形（クロスＨ） 
親要素：StbJoints 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
ID 
 
guid 
string 
 
GUID 
 
joint_name 
string 
 
継手呼称 
 
joint_mark 
string 
 
継手符号 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCross 
1 
1 
＋形継手詳細 
StbJointShapeCrossTFlangeLong 
0 
1 
＋形継手詳細・T 部分フランジ（長） 
長い方 
※(1) ※(2) 
StbJointShapeCrossTFlangeShort 
0 
1 
＋形継手詳細・T 部分フランジ（端） 
短い方 
※(1) ※(2) 
StbJointShapeCrossTWebLong 
1 
1 
＋形継手詳細・T 部分ウェブ(長) 
長い方 
※(2) 
StbJointShapeCrossTWebShort 
1 
1 
＋形継手詳細・T 部分ウェブ(短) 
短い方 
※(2) 
StbJointShapeCrossHFlange 
0 
1 
＋形継手詳細・H 部分フランジ 
※(1) ※(2) 
StbJointShapeCrossHWebLong 
1 
1 
＋形継手詳細・H 部分ウェブ(長) 
長い方 
※(2) 
StbJointShapeCrossHWebShort 
1 
1 
＋形継手詳細・H 部分ウェブ(短) 
短い方 
※(2) 
 
・補足 
(1) 「StbJointCrossTFlangeLong」、「StbJointCrossTFlangeShort」、「StbJointCrossHFlange」
を省略した場合、フランジ溶接とする。

## Page 595

7-28 
 
(2) 各子要素とスプライスプレートの対応は下図の通りとする。 
 
StbJointShapeCrossTFlangeLong 
StbJointShapeCrossTFlangeShort 
StbJointShapeCrossTWebLong 
StbJointShapeCrossTWebShort 
StbJointShapeCrossHFlange 
StbJointShapeCrossHWebLong 
StbJointShapeCrossHWebShort

## Page 596

7-29 
 
7.1.4.1. ＋形継手詳細：StbJointShapeCross 
・概要 
説明 ：＋形（クロスＨ）継手の詳細 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength_plate_flange 
string 
○ 
添え板の材種（フランジ） 例：SN490A 
strength_plate_web 
string 
○ 
添え板の材種（ウェブ） 
省略値はフランジ
に同じ 
例：
SN490A 
strength_bolt 
string 
○ 
ボルト材種 
例：F10T 
name_bolt 
string 
○ 
ボルト径（呼名） 
例：M22 
offset_H 
double 
 
Ｈ形鋼の偏心（T 形鋼部
分の成の中心からの距
離） 
※(1) 
正の値とする 
offset_T 
double 
 
T 形鋼部分の偏心（Ｈ形
鋼の成の中心からの距
離） 
※(1) 
正の値とする 
clearance 
double 
 
部材の母材間隔 
※(2) 
strength_filler 
string 
 
フィラープレートの材種
（共通） 
※(3) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 省略値は、0 とする。 
(2) 省略値は、10 ㎜とする。 
(3) JASS6 の規定ではフィラープレートの材種は400N 級で良いことになっており、特別に指定する
場合のみ記述する。

## Page 597

7-30 
 
7.1.4.2. ＋形継手詳細・T 部分フランジ（長）：StbJointShapeCrossTFlangeLong 
・概要 
説明 ：＋形（クロスＨ）継手のT 形鋼部分 フランジ添え板（長い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isZigzag 
boolean 
 
千鳥配置か否か 
※(1) 
g1 
double 
○ 
ゲージ寸法1 (g1) 
※(2) 
g2 
double 
 
ゲージ寸法2 (g2) 
mf≧4 のとき必須 
※(2) 
pitch 
double 
○ 
長手方向のボルトピッチ(P) 
※(2) 
e1 
double 
○ 
縁端距離1 (e1) 
※(2) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(2) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(2) 
outside_thickness 
double 
○ 
外添え板 厚さ 
 
outside_width 
double 
○ 
外添え板 幅(B) 
※(2) 
outside_length 
double 
 
外添え板 長さ(L) 
※(2) ※(3) 
inside_thickness 
double 
 
内添え板 厚さ 
 
inside_width 
double 
 
内添え板 幅 
※(2) 
inside_length 
double 
 
内添え板 長さ 
※(2) ※(3) 
 
・内容 
無し

## Page 598

7-31 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossFlan
geBolt 
1 
制限なし 
＋形（クロスＨ）継手のフランジ
ボルト詳細 
 
 
・補足 
(1) 省略値は、否とする。また、千鳥配置の場合は外スタートとする。 
(2) 各属性は、下図の通りとする。 
(3) 「縁端距離2 (e2)」の指定がある場合は、「外添え板 長さ(L)」および「内添え板 長さ」は自
動決定するため、値を省略する。「外添え板 長さ(L)」および「内添え板 長さ」を指定する
場合、「縁端距離2 (e2)」の値を省略する。 
 
    T 部分（H 部分） 
   （フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
e2 
e1 
P 
P 
ｸﾘｱﾗﾝｽ 
L 
B 
g2 
g2 
g1 
e3 
e4 
e4 
e3

## Page 599

7-32 
 
7.1.4.2.1. ＋形継手詳細・フランジボルト詳細：StbJointShapeCrossFlangeBolt 
概要 
説明 ：＋形（クロスＨ）継手のフランジのボルト詳細 
親要素： StbJointShapeCrossTFlangeLong、StbJointShapeCrossTFlangeShort、 
 
StbJointShapeCrossHFlange 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
継手中心からのボル
トの列数 
mf 
integer 
○ 
ボルトの数 
列当たりのボルトの
数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は、下図の通りとする。 
 
（フランジの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
id_order 
m3 
3 
2 
1 
m2 
m1 
mf

## Page 600

7-33 
 
7.1.4.3. ＋形継手詳細・T 部分フランジ（短）：StbJointShapeCrossTFlangeShort 
・概要 
説明 ：＋形（クロスＨ）継手のT 形鋼部分 フランジ添え板（短い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isZigzag 
boolean 
 
千鳥配置か否か 
※(1) 
g1 
double 
○ 
ゲージ寸法1 (g1) 
※(2) 
g2 
double 
 
ゲージ寸法2 (g2) 
mf≧4 のとき必須 
※(2) 
pitch 
double 
○ 
長手方向のボルトピッチ(P) 
※(2) 
e1 
double 
○ 
縁端距離1 (e1) 
※(2) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(2) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(2) 
outside_thickness 
double 
○ 
外添え板 厚さ 
 
outside_width 
double 
○ 
外添え板 幅(B) 
※(2) 
outside_length 
double 
 
外添え板 長さ(L) 
※(2) ※(3) 
inside_thickness 
double 
 
内添え板 厚さ 
 
inside_width 
double 
 
内添え板 幅 
※(2) 
inside_length 
double 
 
内添え板 長さ 
※(2) ※(3) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossFlan
geBolt 
1 
制限なし 
＋形（クロスＨ）継手のフランジ
ボルト詳細

## Page 601

7-34 
 
・補足 
(1) 省略値は、否とする。また、千鳥配置の場合は外スタートとする。 
(2) 「7.1.4.2. ＋形継手詳細・T 部分フランジ(長)」の補足参照 
(3) 「縁端距離2 (e2)」の指定がある場合は、「外添え板 長さ(L)」および「内添え板 長さ」は自
動決定するため、値を省略する。「外添え板 長さ(L)」および「内添え板 長さ」を指定する
場合、「縁端距離2 (e2)」の値を省略する。

## Page 602

7-35 
 
7.1.4.4. ＋形継手詳細・T 部分ウェブ(長)：StbJointShapeCrossTWebLong 
・概要 
説明 ：＋形（クロスＨ）継手のT 形鋼部分 ウェブ添え板（長い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
e2 の初期値はe1 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(2) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(2) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
※(1) 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
※(1) 
 
・内容 
無し

## Page 603

7-36 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossWeb
Bolt 
1 
制限なし 
＋形（クロスＨ）継手のウェブボ
ルト詳細 
 
 
・補足 
 
(1) 各属性は、下図の通りとする（TL 側）。 
 
 
T 部分 
 
（ウェブの継手） 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
e3 
Offset＿H 
<TL 側> 
<TS 側> 
pC 
e2 
pL 
ｸﾘｱﾗﾝｽ 
e1 
L 
e4 
B 
B 
e5 
T 形鋼部分の中心

## Page 604

7-37 
 
7.1.4.4.1. ＋形継手詳細・ウェブボルト詳細：StbJointShapeCrossWebBolt 
概要 
説明 ：＋形（クロスＨ）継手のウェブのボルト配列 
親要素： StbJointShapeCrossTWebLong、StbJointShapeCrossTWebShort、 
 
StbJointShapeCrossHWebLong、StbJointShapeCrossHWebShort 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
継手中心からのボル
トの列数 
mw 
integer 
○ 
ボルトの数 
列当たりのボルトの
数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は、下図の通りとする。 
 
（ウェブの継手） 
 
 
 
<TL 側> 
<TS 側> 
id_order 
m1 
m2 
2 
1 
mw 
<HL 側> 
<HS 側> 
id_order 
m1 
m2 
2 
1 
mw

## Page 605

7-38 
 
7.1.4.5. ＋形継手詳細・T 部分ウェブ(短)：StbJointShapeCrossTWebShort 
・概要 
説明 ：＋形（クロスＨ）継手のT 形鋼部分 ウェブ添え板（短い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
e2 の初期値はe1 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(2) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(2) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
※(1) 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
※(1) 
 
・内容 
無し

## Page 606

7-39 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossWeb
Bolt 
1 
制限なし 
＋形（クロスＨ）継手のウェブボ
ルト詳細 
 
 
・補足 
(1) 「」の補足（TS 側）参照。

## Page 607

7-40 
 
7.1.4.6. ＋形継手詳細・H 部分フランジ：StbJointShapeCrossHFlange 
・概要 
説明 ：＋形（クロスＨ）継手のＨ形鋼部分 ウェブ添え板（長い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
isZigzag 
boolean 
 
千鳥配置か否か 
※(1) 
g1 
double 
○ 
ゲージ寸法1 (g1) 
※(2) 
g2 
double 
 
ゲージ寸法2 (g2) 
mf≧4 のとき必須 
※(2) 
pitch 
double 
○ 
長手方向のボルトピッチ(P) 
※(2) 
e1 
double 
○ 
縁端距離1 (e1) 
※(2) 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
※(2) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(2) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(2) 
outside_thickness 
double 
○ 
外添え板 厚さ 
 
outside_width 
double 
○ 
外添え板 幅(B) 
※(2) 
outside_length 
double 
 
外添え板 長さ(L) 
※(2) ※(3) 
inside_thickness 
double 
 
内添え板 厚さ 
 
inside_width 
double 
 
内添え板 幅 
※(2) 
inside_length 
double 
 
内添え板 長さ 
※(2) ※(3) 
 
・内容 
無し

## Page 608

7-41 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossFlan
geBolt 
1 
制限なし 
＋形（クロスＨ）継手のフランジ
ボルト詳細 
 
 
・補足 
(1) 省略値は、否とする。また、千鳥配置の場合は外スタートとする。 
(2) 「1.4.2. ＋形継手詳細・T 部分フランジ（長）」の補足参照。 
(3)  「縁端距離2 (e2)」の指定がある場合は、「外添え板 長さ(L)」および「内添え板 長さ」は自
動決定するため、値を省略する。「外添え板 長さ(L)」および「内添え板 長さ」を指定する
場合、「縁端距離2 (e2)」の値を省略する。

## Page 609

7-42 
 
7.1.4.7. ＋形継手詳細・H 部分ウェブ(長)：StbJointShapeCrossHWebLong 
・概要 
説明 ：＋形（クロスＨ）継手のＨ形鋼部分 ウェブ添え板（長い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
e2 の初期値はe1 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(1) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(1) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(1) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
※(1) 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
※(1) 
 
・内容 
無し

## Page 610

7-43 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossWeb
Bolt 
1 
制限なし 
＋形（クロスＨ）継手のウェブボ
ルト詳細 
 
 
・補足 
 
(1) 各属性は、下図の通りとする（HL 側）。 
 
 
H 部分 
 
（ウェブの継手） 
   
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
e3 
Offset＿T 
<HL 側> 
<HS 側> 
pC 
e2 
pL 
ｸﾘｱﾗﾝｽ 
e1 
L 
e4 
B 
B 
e5 
H 形鋼の中心

## Page 611

7-44 
 
7.1.4.8. ＋形継手詳細・H 部分ウェブ(短)：StbJointShapeCrossHWebShort 
・概要 
説明 ：＋形（クロスＨ）継手のＨ形鋼部分 ウェブ添え板（短い方）の寸法とボルト穴位置 
親要素：StbJointColumnShapeCross 
 
・属性 
属性名 
型 
必須 
説明 
補足 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
mw≧2 の時、必須 
※(1) 
pitch 
double 
 
部材長手方向のボルトピッチ(pL) 
nw≧2 の時、必須 
※(1) 
e1 
double 
○ 
縁端距離1 (e1) 
e2 の初期値はe1 
※(1) 
e2 
double 
 
縁端距離2 (e2) 
※(1) 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
内添え板 
※(1) 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
内添え板 
※(1) 
e5 
double 
 
縁端距離5 (e5) 
省略値は振り分け 
※(1) 
plate_thickness 
double 
○ 
添え板 厚さ 
 
plate_width 
double 
○ 
添え板 幅(B) 
※(1) 
plate_length 
double 
 
添え板 長さ(L) 
縁端距離2 (e2)の
指定がある場合、
この値は自動決定
するため、省略す
る。この値を指定
する場合、e2 の値
を省略する 
※(1) 
 
・内容 
無し

## Page 612

7-45 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbJointShapeCrossWeb
Bolt 
1 
制限なし 
＋形（クロスＨ）継手のウェブボ
ルト詳細 
 
 
・補足 
(1) 「」の補足（HYS 側）参照。

## Page 613

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-1 
 
8. 要素リファレンス コネクション情報 
8.1. コネクション情報：StbConnections 
・概要 
説明 ：鉄骨のコネクション 
親要素：StbModel 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbGussetPlates 
0 
1 
ガセットプレート詳細 
 
StbRibPlates 
0 
1 
リブプレート詳細 
 
StbDiaphragms 
0 
1 
ダイアフラム詳細 
 
StbStiffners 
0 
1 
スチフナープレート詳細 
 
 
・補足 
(1) 子要素の並びは、上表に示す順番としなければならない。

## Page 614

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-2 
 
8.2. ガセットプレート詳細（複数）：StbGussetPlates 
・概要 
説明 ：ガセットプレート情報（複数） 
親要素：StbConnections 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbGussetPlate 
1 
制限なし 
ガセットプレート詳細

## Page 615

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-3 
 
8.2.1. ガセットプレート詳細：StbGussetPlate 
概要 
説明 ：ガセットプレート詳細 
親要素： StbGussetPlates 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
符号 
 
connection_type 
string 
○ 
仕口タイプ 
gap、web、splice 
 
flange_type 
string 
○ 
フランジ刃落としタイプ 
none、one_side、both_sides 
 
strength_plate 
string 
○ 
ガセットプレートの材種 
例：SS400 
thickness 
double 
○ 
ガセットプレートの厚さ 
 
strength_bolt 
string 
○ 
ボルト材種 
例：S10T 
name_bolt 
string 
○ 
ボルト径（呼び名） 
例：M20 
clearance 
double 
 
母材と接続部材の間隔 
 
flange_cut 
double 
 
フランジ刃落としの長さ 
フランジ刃落としタイ
プ
(flange_type)
が
none 以外の時、必須 
 
pitch_depth 
double 
 
部材成方向のボルトピッチ (pC) 
m≧2 の時、必須 
 
pitch 
double 
 
部材長手方向のボルトピッチ (pL) 
n≧2 の時、必須 
 
e1 
double 
○ 
縁端距離1 (e1) 
 
e2 
double 
 
縁端距離2 (e2) 
e2 の初期値はe1 
 
e3 
double 
 
縁端距離3 (e3) 
e3 の初期値はe1 
 
e4 
double 
 
縁端距離4 (e4) 
e4 の初期値はe1 
 
h1 
double 
 
ガセットプレートのあき1 (h1) 
e3、e4 の数値を優先す

## Page 616

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-4 
 
る 
h2 
double 
 
ガセットプレートのあき2 (h2) 
e3、e4 の数値を優先す
る 
ey 
double 
 
嵩上げ点の幅方向の距離 
 
ez 
double 
 
嵩上げ点の高さ方向の距離 
 
R 
double 
 
入隅半径 (R) 
 
id_weld 
integer 
 
溶接ID 
 
cutback 
integer 
 
溶接控え 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbGussetPlateBoltArr
ay 
0 
制限なし 
ボルト詳細 
 
StbGussetPlateSplice 
0 
1 
スプライスプレート詳細 
 
 
・補足 
各属性は3.2.5.1.1. ガセットプレート仕様補足を参照

## Page 617

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-5 
 
8.2.1.1. ボルト詳細：StbGussetPlateBoltArray 
概要 
説明 ：ボルト詳細 
親要素：StbGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id_order 
integer 
○ 
ボルトの列数 
中心からのボルトの列数 
mw 
integer 
○ 
ボルトの数 
列当たりのボルトの数 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
各属性は7.1.1.3.1. Ｈ形継手詳細・ウェブボルト詳細を参照

## Page 618

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-6 
 
8.2.1.2. スプライスプレート詳細：StbGussetPlateSplice 
概要 
説明 ：スプレイスプレート詳細 
親要素：StbGussetPlate 
 
・属性 
属性名 
型 
必須 
説明 
補足 
strength_plate 
string 
○ 
添え板の材種 
例：SN490B 
plate_thickness 
double 
○ 
添え板の厚さ 
 
strength_filler 
string 
 
フィラープレートの材種 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) JASS6 の規定ではフィラープレートの材種は400N 級で良いことになっており、特別に指定する場
合のみ記述する。

## Page 619

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-7 
 
8.3. リブプレート詳細（複数）：StbRibPlates 
・概要 
説明 ：リブプレート情報（複数） 
親要素：StbConnections 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbRibPlate 
1 
制限なし 
リブプレート詳細

## Page 620

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-8 
 
8.3.1. リブプレート詳細：StbRibPlate 
概要 
説明 ：リブプレート詳細 
親要素：StbRibPlates 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
符号 
 
strength_plate 
string 
○ 
リブプレートの材種 
例：SN490B 
thickness 
double 
○ 
リブプレートの厚さ 
 
id_weld 
integer 
 
溶接ID 
 
cutback 
integer 
 
溶接控え 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 621

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-9 
 
8.4. ダイアフラム詳細（複数）：StbDiaphragms 
・概要 
説明 ：ダイアフラム情報（複数） 
親要素：StbConnections 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbDiaphragm 
1 
制限なし 
ダイアフラム詳細

## Page 622

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-10 
 
8.4.1. ダイアフラム詳細：StbDiaphragm 
概要 
説明 ：ダイアフラム詳細 
親要素：StbDiaphragms 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
符号 
 
strength_plate 
string 
○ 
ダイアフラムの材種 
例：SN490B 
thickness 
double 
○ 
ダイアフラムの厚さ 
 
B_X 
double 
〇 
ダイアフラムの寸法(Bx) 
※(1) 
B_Y 
double 
 
ダイアフラムの寸法(By) 
※(1) 
type 
string 
○ 
ダイアフラム形式 
以下のいずれかの値をとる。 
Through（通しダイア）、 
Internal（内ダイア）、 
External（外ダイア） 
 
id_weld 
integer 
 
溶接ID 
※(2) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) ダイアフラムが円形でBx のみを記述した場合円形形状とする。それ以外の場合はBy を必須とす
る。 
(2) 通しダイア形式以外の場合、必須とする。

## Page 623

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-11 
 
8.5. スチフナープレート詳細（複数）：StbStiffners 
・概要 
説明 ：スチフナープレート情報（複数） 
親要素：StbConnections 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbStiffner 
1 
制限なし 
スチフナープレート詳細

## Page 624

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8-12 
 
8.5.1. スチフナープレート詳細：StbStiffner 
概要 
説明 ：スチフナープレートの詳細 
親要素：StbStiffners 
 
・属性 
 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
符号 
 
strength_plate 
string 
〇 
スチフナーの材種 
例：SN490B 
thickness 
double 
○ 
スチフナーの厚さ 
 
B_X 
double 
〇 
スチフナーの寸法(Bx) 
 
B_Y 
double 
〇 
スチフナーの寸法(By) 
 
id_weld 
integer 
 
溶接ID 
 
cutback 
integer 
 
溶接控え 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
無し

## Page 625

9-1 
 
9. 要素リファレンス 溶接情報 
9.1. 溶接：StbWeld 
概要 
説明 ：溶接の仕様 
親要素：StbModel 
 
属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWeldFullPenetration 
0 
制限なし 
完全溶込み溶接 
 
StbWeldPartialPenetration 
0 
制限なし 
部分溶込み溶接 
 
StbWeldFillet 
0 
制限なし 
隅肉溶接 
 
StbWeldFlare 
0 
制限なし 
フレア溶接 
 
・補足 
(1) 子要素の並びは、上表に示す順番としなければならない。 
(2) 子要素のID はStbWeld で一意としなければならない。

## Page 626

9-2 
 
9.2. 完全溶込み溶接：StbWeldFullPenetration 
・概要 
説明 ：完全溶込み溶接 
親要素：StbWeld 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
mark 
string 
〇 
識別マーク 
※(1) 
name 
string 
〇 
略称 
※(1) 
type_weld 
string 
〇 
継手形式 
以下のいずれかの値をとる。 
B（突合せ溶接）、 
T（T 形溶接）、 
C（角溶接） 
E（エレクトロスラグ溶接） 
 
type_backup 
string 
〇 
バックアップ材の種類 
以下のいずれかの値をとる。 
U（裏当金）、 
H（裏はつり） 
 
shape_bevel 
string 
〇 
開先加工の形状 
以下のいずれかの値をとる。 
L（レ形開先）、 
V（V 形開先）、 
K（K 形開先）、 
X（X 形開先） 
I（I 形開先） 
 
side_bevel1 
string 
 
開先の方向（表裏） 
以下のいずれかの値をとる。 
Front（表面） 
Reverse（裏面） 
Both（両面） 
レ形、Ｖ形の場合必須 
side_bevel2 
string 
 
開先の方向（左右） 
以下のいずれかの値をとる。 
Right（右側） 
省略値は右側とする 
※(2)

## Page 627

9-3 
 
Left（左側） 
min_t1 
double 
〇 
板厚の下限 
 
max_t1 
double 
〇 
板厚の上限 
 
G 
double 
 
ルート間隔 
 
R 
double 
 
ルート面 
 
alfa1 
integer 
 
開先角度1（表面） 
 
alfa2 
integer 
 
開先角度2（裏面） 
 
method 
string 
〇 
溶接方法 
以下のいずれかの値をとる。 
G（GMAW-ガスシールドアーク
溶接、SMAW-被覆アーク溶接）、 
S（SAW-サブマージアーク溶
接）、 
E（ESW-エレクトロスラグ溶
接） 
 
location 
string 
〇 
溶接場所 
以下のいずれかの値をとる。 
Factory（工場溶接）、Site（現
場溶接） 
 
robots 
string 
 
ロボット溶接の対応 
 
kind_end_tab 
string 
 
エンドタブの種類 
以下のいずれかの値をとる。 
Steel（鋼製エンドタブ）、Flux
（固形エンドタブ） 
 
strength_backup 
string 
 
裏当て金の材質 
例：SN490B 
shape_backup1 
string 
 
裏当て金の材種1 
以下の値をとる。 
PL（プレート）、FB（フラット
バー） 
 
size_backup1 
string 
 
裏当て金のサイズ1 
※(3) 
shape_backup2 
string 
 
裏当て金の材種2 
以下の値をとる。 
PL（プレート）、FB（フラット
バー） 
 
size_backup2 
string 
 
裏当て金のサイズ2 
※(3) 
taper_processing 
boolean 
 
テーパー加工

## Page 628

9-4 
 
facing_processing 
boolean 
 
フェーシング加工 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWeldSpec 
1 
制限なし 
溶接仕様詳細

## Page 629

9-5 
 
・補足 
(1) 略称および略図は下図による。 
(2) 継手形式がB（突合せ溶接）かつ同厚の場合にどちらの板に開先があるかを指定する。 
(3) 裏当て金のサイズは「板厚x 幅」の記述とする。

## Page 630

9-6 
 
完全溶込み溶接（FP）

## Page 631

9-7

## Page 632

9-8

## Page 633

9-9 
 
部分溶込み溶接（P）

## Page 634

9-10

## Page 635

9-11 
 
隅肉溶接（F）

## Page 636

9-12

## Page 637

9-13 
 
フレア溶接（FL）

## Page 638

9-14 
 
9.2.1. 溶接仕様詳細：StbWeldSpec 
・概要 
説明 ：溶接仕様詳細 
親要素：StbWeldFullPenetration、StbWeldPartialPenetration、StbWeldFillet、StbWeldFlare 
 
・属性 
属性名 
型 
必須 
説明 
補足 
t1 
double 
〇 
板厚 
存在する板厚毎に記載 
d 
double 
 
直径 
溶接する丸鋼の直径 
D1 
double 
 
開先深さ1（表面） 
 
D2 
double 
 
開先深さ2（裏面） 
 
S1 
double 
 
隅肉サイズ（表面） 
 
S2 
double 
 
隅肉サイズ（裏面） 
 
h1 
double 
 
余盛高さ1（表面） 
 
h2 
double 
 
余盛高さ2（裏面） 
 
area 
double 
 
断面積 
※(1) 
length 
double 
 
溶接実長 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 6mm 換算長を算出するため、換算率ではなく溶接断面積を記述する。

## Page 639

9-15 
 
9.3. 部分溶け込み溶接：StbWeldPartialPenetration 
・概要 
説明 ：部分溶込み溶接 
親要素：StbWeld 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
mark 
string 
〇 
識別マーク 
※(1) 
name 
string 
〇 
略称 
※(1) 
type_weld 
string 
〇 
継手形式 
以下のいずれかの値をとる。 
B（突合せ溶接）、 
T（T 形溶接）、 
C（角溶接） 
 
type_backup 
string 
 
バックアップ材の種類 
角溶接で必要な場合以下の値を
とる。 
U（裏当金） 
 
shape_bevel 
string 
〇 
開先加工の形状 
以下のいずれかの値をとる。 
L（レ形開先）、 
V（V 形開先）、 
K（K 形開先） 
 
side_bevel1 
string 
 
開先の方向（表裏） 
以下のいずれかの値をとる。 
Front（表面） 
Reverse（裏面） 
Both（両面） 
レ形、Ｖ形の場合必須 
side_bevel2 
string 
 
開先の方向（左右） 
以下のいずれかの値をとる。 
Right（右側） 
Left（左側） 
省略値は右側とする 
※(2) 
min_t1 
double 
〇 
板厚の下限 
 
max_t1 
double 
〇 
板厚の上限

## Page 640

9-16 
 
R 
double 
〇 
ルート面 
 
alfa1 
integer 
〇 
開先角度1 
 
alfa2 
integer 
 
開先角度2 
K 形の場合必須 
continuous 
boolean 
 
まわし溶接の有無 
 
method 
string 
〇 
溶接方法 
以下のいずれかの値をとる。 
G（GMAW-ガスシールドアーク
溶接、SMAW-被覆アーク溶接）、 
S（SAW-サブマージアーク溶
接） 
 
location 
string 
〇 
溶接場所 
以下のいずれかの値をとる。 
Factory（工場溶接）、Site（現
場溶接） 
 
robots 
string 
 
ロボット溶接の対応 
 
strength_backup 
string 
 
裏当て金の材質 
例：SN490B 
shape_backup 
string 
 
裏当て金の材種 
以下の値をとる。 
PL（プレート）、FB（フラット
バー） 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWeldSpec 
1 
制限なし 
溶接仕様詳細 
 
 
・補足 
(1) 略称および略図は9.2. 完全溶込み溶接による。 
(2) 継手形式がB（突合せ溶接）かつ同厚の場合にどちらの板に開先があるかを指定する。

## Page 641

9-17 
 
9.4. 隅肉溶接：StbWeldFillet 
・概要 
説明 ：隅肉溶接 
親要素：StbWeld 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
mark 
string 
〇 
識別マーク 
※(1) 
name 
string 
〇 
略称 
※(1) 
type_weld 
string 
〇 
継手形式 
以下のいずれかの値をとる。 
B（突合せ溶接）、 
T（T 形溶接）、 
L（重ね溶接） 
 
shape_bevel 
string 
 
開先形状 
以下のいずれかの値をとる。 
L（レ形開先）、 
K（K 形開先） 
 
side_bevel1 
string 
 
開先の方向 
以下のいずれかの値をとる。 
Front（表面） 
Reverse（裏面） 
レ形の場合必須 
side_bevel2 
string 
 
開先の方向（左右） 
以下のいずれかの値をとる。 
Right（右側） 
Left（左側） 
省略値は右側とする 
※(2) 
min_t1 
double 
〇 
板厚の下限 
 
max_t1 
double 
〇 
板厚の上限 
 
G 
double 
 
ルート間隔 
 
R 
double 
 
ルート面 
K 形の場合必須 
alfa1 
integer 
 
開先角度1（表面） 
レ形、K 形の場合必須 
alfa2 
integer 
 
開先角度2（裏面） 
K 形の場合必須 
continuous 
boolean 
 
まわし溶接の有無 
T 形は必須

## Page 642

9-18 
 
location 
string 
〇 
溶接場所 
以下のいずれかの値をとる。 
Factory（工場溶接）、Site（現
場溶接） 
 
robots 
string 
 
ロボット溶接の対応 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWeldSpec 
1 
制限なし 
溶接仕様詳細 
 
 
・補足 
(1) 略称および略図は9.2. 完全溶込み溶接による。 
(2) 継手形式がB（突合せ溶接）かつ同厚の場合にどちらの板に開先があるかを指定する。

## Page 643

9-19 
 
9.5. フレア溶接：StbWeldFlare 
・概要 
説明 ：フレア溶接 
親要素：StbWeld 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ID 
 
guid 
string 
 
GUID 
 
mark 
string 
〇 
識別マーク 
※(1) 
name 
string 
〇 
略称 
※(1) 
type_weld 
string 
〇 
継手形式 
以下のいずれかの値をとる。 
RR（丸鋼―丸鋼）、 
RP（丸鋼―板）、 
CC（角形鋼管―角形鋼管） 
CP（角形鋼管―板） 
 
min_t 
double 
〇 
板厚の下限 
 
max_t 
double 
〇 
板厚の上限 
 
continuous 
boolean 
 
まわし溶接の有無 
 
location 
string 
〇 
溶接場所 
以下のいずれかの値をとる。 
Factory（工場溶接）、Site（現
場溶接） 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbWeldSpec 
1 
制限なし 
溶接仕様詳細 
 
 
・補足 
(1) 略称および略図は9.2. 完全溶込み溶接による。

## Page 644

10-1 
 
10. 要素リファレンス 拡張情報 
 
10.1. 拡張情報（複数）：StbExtensions 
・概要 
説明 ：拡張情報 
親要素：ST_BRIDGE 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbExtension 
1 
制限なし 
拡張情報 
※(1) 
 
・補足 
(1) StbExtension は、ST-Bridge に定義されていない属性をアプリケーションが独自に拡張するフレ
ームワークを提供するオプションである。ST-Bridge の要素に、属性および子要素を拡張する方
法には、下記の２種類がある。 
 
１． 属性の定義と属性値の指定を同時に行う方法（子要素StbExtObject を使用） 
 拡張される元の要素の記述は変更せず、この要素内で追加属性を定義し、属性値を指定する。 
原則として、固有のID を持つST-Bridge の各要素に対して属性を追加することができる。 
 
２． 子要素および属性の定義のみ行う方法（子要素StbExtElement を使用） 
   この要素では追加の定義を行い、拡張される元の要素に追加子要素、追加属性と属性値を記
述する。元の既存の要素（子要素）に、属性を追加することも可能である。 
   ST-Bridge の全ての要素を追加の対象とする。

## Page 645

10-2 
 
・例 
(1) 柱の各面に仕上げ重量について、StbExtObject を用いて、属性の定義と属性値の記述を行う。 
<StbExtensions> 
<StbExtension identifier ="XXX-001" description="Finishing load" > 
<StbExtObject object_name="StbColumn" id_object="1001"> 
<StbExtProperty key="x_start" type="double" value="200.0"/> 
<StbExtProperty key="x_end" type="double" value="200.0"/> 
<StbExtProperty key="y_start" type="double" value="150.0"/> 
<StbExtProperty key="y_end" type="double" value="200.0"/> 
・・・ 
</StbExtObject> 
</StbExtension> 
</StbExtensions> 
 
(2) 柱の各面に仕上げ重量について、StbExtElement を用いて、追加子要素と属性を定義し、拡張さ
れる要素StbColumn に、子要素、属性および属性値の記述を行う。 
<StbExtensions> 
<StbExtension identifier ="XXX-001" description="Finishing load" > 
<StbExtElement object_name="StbColumn" element_name="StbFinishingLoad" /> 
  <StbExtProprtyDef key="x_start" type="double"/> 
  <StbExtProprtyDef key="x_end" type="double"/> 
  <StbExtProprtyDef key="y_start" type="double"/> 
  <StbExtProprtyDef key="y_end" type="double"/> 
</StbExtElement> 
</StbExtension> 
</StbExtensions> 
 
<StbModel> 
  <StbColumns> 
    <StbColumn id=1001 ……..> 
      <StbFinishingLoad x_start="200.0" x_end="200.0" y_start="150.0" y_end="200.0"/> 
    </StbColumn> 
  </StbColumns> 
</StbModel>

## Page 646

10-3 
 
10.2. 拡張情報：StbExtension 
・概要 
説明 ：拡張情報  
親要素：StbExtensions 
 
・属性 
属性名 
型 
必須 
説明 
補足 
identifier 
string 
○ 
拡張情報の識別子 
 
description 
string 
 
拡張情報の説明 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbExtObject 
0 
制限なし 
対象オブジェクト 
 
StbExtElement 
0 
制限なし 
拡張子要素 
 
 
・補足 
  上記のうち1 種類以上の子要素を持つものとし、全子要素の最小回数が0 であってはならない。

## Page 647

10-4 
 
10.2.1. 対象オブジェクト：StbExtObject 
・概要 
説明 ：拡張対象となる要素（属性の定義と属性値の指定を同時に行う場合） 
親要素：StbExtension 
 
・属性 
属性名 
型 
必須 
説明 
補足 
object_name 
string 
○ 
ST-Bridge の要素名 
※(1) 
id_object 
integer 
○ 
要素のID 
 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbExtProperty 
1 
制限なし 
拡張属性 
 
 
・補足 
(1) 要素名は、固有のID を持つ要素のみ指定できる。ただし、共通情報（StbCommon）の属性を拡
張する場合は要素名としてStbCommon を指定し、id_object="0"とする。

## Page 648

10-5 
 
10.2.1.1. 拡張属性：StbExtProperty 
・概要 
説明 ：拡張属性および属性値 
親要素：StbExtObject 
 
・属性 
属性名 
型 
必須 
説明 
補足 
key 
string 
○ 
変数名 
 
type 
string 
○ 
変数型で以下のいずれかの値をとる 
string：文字型 
integer：整数型 
double：実数型 
boolean：論理型 
 
value 
string 
○ 
値 
 
 
・内容 
無し 
 
・子要素 
無し

## Page 649

10-6 
 
10.2.2. 拡張子要素：StbExtElement 
・概要 
説明 ：拡張対象となる子要素（定義のみ行う場合） 
親要素：StbExtension 
 
・属性 
属性名 
型 
必須 
説明 
補足 
object_name 
string 
○ 
ST-Bridge の要素名 
 
element_name 
string 
 
拡張する子要素の名前 
※(1) 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbExtPropertyDef 
0 
制限なし 
拡張属性定義 
※(1) 
 
・補足 
(1) ・子要素を新たに追加する場合 
element_name に、拡張する子要素の名前を記述する。この名前は、拡張する子要素同志で重
複しないほか、既存の子要素名および他のプログラムが使う拡張子要素と重複を避けるよう、注
意する必要がある。 
なお、子要素を追加し、その子要素に属性を設定しない場合はStbExtPropertyDef は記述しな
い。 
      
・既存の要素に、属性を新たに追加する場合 
 element_name を省略する。このとき、StbExtPropertyDef（拡張属性定義）は、要素内におい
て、拡張する属性同志で重複しないほか、既存の属性名および他のプログラムが使う拡張属性と
重複を避けるよう、注意する必要がある。 
 
既存の子要素に、属性を新たに追加する場合は、要素名object_name に該当する子要素を記述
する。

## Page 650

10-7 
 
10.2.2.1. 拡張属性定義：StbExtPropertyDef 
・概要 
説明 ：拡張属性の定義 
親要素：StbExtElement 
 
・属性 
属性名 
型 
必須 
説明 
補足 
key 
string 
○ 
変数名 
 
type 
string 
○ 
変数型で以下のいずれかの値をとる 
string：文字型 
integer：整数型 
double：実数型 
boolean：論理型 
 
default 
type で指
定した型 
 
省略値 
※(1) 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 拡張属性に「省略された場合」の扱いがあるとき、記述する。

## Page 651

11-1 
 
11. 要素リファレンス 出力情報 
  
11.1. 出力情報：StbExportInformation 
・概要 
説明 ：出力情報 
親要素：ST_BRIDGE 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbExportPolicy 
0 
制限なし 
出力方針 
※(1) 
StbExportLog 
0 
制限なし 
出力ログ 
※(2) 
 
・補足 
(1) 出力アプリケーション側がデフォルトで設定している項目や未対応の項目を記入する 
(2) 個々の要素に対して、出力時の注釈等を残す場合に用いる。 
(3) ST-Bridge ファイルを利用するユーザーに対しての補足情報であり、StbExportInformation を用
いて変換することは想定していない。

## Page 652

11-2 
 
11.2. 変換方針：StbExportPolicy 
・概要 
説明 ：変換方針 
親要素：StbExportInformation 
 
・属性 
属性名 
型 
必須 
説明 
補足 
comment 
string 
○ 
変換方針説明 
※(1) 
・内容 
無し 
 
・子要素 
無し 
 
補足 
(1) ST-Bridge ファイルの出力時に必要とした項目のみを列挙し、使用しなかった変換方針について
は出力しない。

## Page 653

11-3 
 
11.3. 変換ログ：StbExportLog 
・概要 
説明 ：変換ログ 
親要素：StbExportInformation 
 
・属性 
無し 
 
・内容 
無し 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbExportError 
0 
制限なし 
エラーログ 
※(1), (3) 
StbExportWarning 
0 
制限なし 
ワーニングログ 
※(2), (3) 
StbExportNotice 
0 
制限なし 
通知ログ 
※(4) 
 
・補足 
(1) ST-Bridge にある要素で代替できないため、要素情報そのものが欠落してしまう場合に使用する。
ただし、ST-Bridge で対応していない要素や属性を除く。 
 
(2) 情報は欠落するがST-Bridge にある要素で代替可能なものに使用する。 
 
(3) ST-Bridge で未定義な断面を用いた部材を使用している場合、その部材がないものとして
StbExportError として記述するか、仮の断面を与えてStbExportWarning として記述するかは
出力アプリケーションによるものとする。 
 
(4) デバッグ用に使用する。ファイルサイズが大きくなることが想定されるため、変換時に
StbExportNotice の有無を選択できることが望ましい。

## Page 654

11-4 
 
11.3.1. エラーログ：StbExportError 
・概要 
説明 ：エラーログ 
親要素：StbExportLog 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
エラーID 
 
object_name 
string 
 
出力アプリケーションでの名前 
 
id_object 
integer 
 
出力アプリケーションでのID 
 
comment 
string 
○ 
説明 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
 
例 
 <StbConversionError id = “1” object_name=”P1” id_object=”1” comment = “杭断面P1 の出力は未実装で
す。”/>

## Page 655

11-5 
 
 
11.3.2. 通知ログ：StbExportWarning 
・概要 
説明 ：警告ログ 
親要素：StbExportLog 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
ワーニングID 
 
object_name 
string 
○ 
ST-Bridge の要素名 
※(1) 
id_object 
integer 
○ 
object_name に該当するのID 
※(1) 
comment 
string 
○ 
説明 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足 
(1) 要素名は、固有のID を持つ要素のみ指定できる。 
 
例 
 <StbConversionWarning id = “1” object_name=”StbSecColumn_S” id_object=”1” comment = “部材名1C1
は鉄骨の材料情報を取得できませんでした。SS400 として設定します。”/> 
<StbConversionWarning id = “2” object_name=”StbSecColumn_RC” id_object=”2” comment = “部材名1C5
は主筋Ｘ方向１段目の本数設定は、断面積(○○mm2)による配筋量指定が行われているためST-BRIDGE で
は扱えません。本数=0 で出力します”/>

## Page 656

11-6 
 
11.3.3. 通知ログ：StbExportNotice 
・概要 
説明 ：通知ログ 
親要素：StbExportLog 
 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
〇 
通知ID 
 
object_name 
string 
 
ST-Bridge の要素名 
 
id_object 
integer 
 
object_name に該当する要素のID 
 
comment 
string 
○ 
説明 
 
 
・内容 
無し 
 
・子要素 
無し 
 
・補足

## Page 657

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
1 
 
【修正履歴】 
2023.01.31  ver.2.1 
 
（新規追加） 
 
・製品一覧表を追加した。 
 
・工法一覧表を追加した。 
 
・鉄骨のコネクション情報の定義を追加 
(StbStandardPlateThicknessList、StbConnectionsSpec、StbConnectionArrangements、 
StbConnections) 
 
・溶接情報の定義を追加（StbWeldCommon、StbWeld） 
 
・免震装置の定義を追加(StbIsolatingDevices、StbSecIsolatingDevice) 
 
・制振装置の定義を追加 
(StbDampingDevices、StbFrameDampingDevices、StbSecDampingDevice) 
 
・床および壁に荷重のみ要素の定義を追加(StbSecSlabLoad、StbSecWallLoad) 
 
・梁貫通孔の定義を追加(StbPenetrationArrangements、StbSecPenetration_S) 
 
・コンクリート柱梁接合部(StbPanelZoneArrangement、StbSecPanelZone)の定義を追加 
 
・出力情報(StbExportInformation)の定義を追加 
 
・StbSecSlab_RC に型枠要素を追加(StbSecFormworkSlab) 
 
 
（変更） 
 
・径別鉄筋強度情報リストを杭とそれ以外に分割 
 
・StbApplyConditionList を要領情報要素として修正 
 
・StbCommon に追加情報(StbAdditionalInformation)を付加 
 
・StbStory にlevel_name を追加 
 
・StbArcAxes で円状に対応できるように修正 
 
・StbColum とStbPost の属性修正 
 
・継手情報を要素(StbJointArrangements)に変更 
 
・梁を任意回数断面に修正し、それぞれの切替位置要素を追加 
(StbGirder、StbBeam、StbSecBeam_RC、StbSecBeam_S、StbSecBeam_SRC) 
 
・StbGirder とStbBeam の属性修正 
 
・スラブのハンチ形状指定は断面で定義するように修正 
 
・開口寸法は断面で定義するように修正 
 
・柱及び梁の配筋定義で鉄筋位置を指定して詳細に入力することもできるように修正 
 
・柱脚ベースプレート仕様修正 
 
・SRC 十字柱でT 形鋼分割されている形状を追加 
 
・鉄骨梁で拡幅形状をあらわせるように修正 
 
・StbSecSlab_RC で在来スラブとトラス筋付きデッキを併記できるように修正。

## Page 658

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
2 
 
 
・RC 壁配筋で複数段定義することができるように修正 
 
・基礎断面配筋を詳細に定義できるように修正 
 
・杭断面定義で工法を追加し、一般工法と認定工法を選択する形式に修正 
 
・RC 杭断面配筋で属性名を変更 
 
・StbSecSteel に組立H 形鋼(上下異なる)、組立T 形鉄骨、溝形鋼(2 丁)、山形鋼(2 丁)、リ
ップ型溝形鋼(2 丁)を追加 
 
・鉄骨H 形鋼のスプライスジョイント仕様の見直し 
 
 
2020.12.16 
ver.2.0(Revision 2) 
 
（補足） 
 
1.11. 名前空間（Namespace）を追記 
 
1.12. XML Schema の利用 を追記 
 
2.1. 全体構成にStbCal とStbAnaModels を追記 
 
2.2. 主要な要素のID と一意性  GUID の位置付けについて確認事項を追記、 
 
StbJoint*****に関する記述を要素名に変更、StbSecOpen_RC・StbSecUndefined を追記 
 
2.6. 座標系  グローバル座標の概念追加に伴い、説明および全体座標系との関係を追記 
 
3.1. ST-BRIDGE の子要素にStbCal とStbAnaModels を追記 
 3.2. 共通情報（StbCommon） 
strength_concrete について、「部材自体に定義がある場合は、部材定義を優先する」 
記述を追記、また、優先順位に関する表を挿入 
 
4.2.1. 節点（StbNode）  ON_GIRDER、ON_BEAM、ON_GRID、ON_CANTI の違い 
が明確となるように図を追加・差し替え、補足説明を追記 
4.2.2. 節点ID リスト（StbNodeIdList） StbNodeId 一意性の補足説明を追記 
4.3.2. 平行軸：StbParallelAxis  例文のXML 記述に対応した図を追記 
       5.14.2. 開口ID リスト（StbOpenIdList） StbOpenId 一意性の補足説明を追記 
6.3. Ｓ柱断面（StbSecColumn_S） 鉄骨向きの補足説明を追記 
 
6.4.9. ＳＲＣ柱断面鉄骨形状（StbSecSteelFigureColumn_SRC） 
 柱脚埋め込み長さの符号の記述（始端から下向き正）を追記 
6.11.7. デッキプレート製品（StbSecProductSlabDeck） デッキ厚さの注釈を追記 
6.10. ＲＣスラブ断面（StbSecSlab_RC）、 
6.11. デッキプレートスラブ断面（StbSecSlabDeck）、 
6.12. 既製スラブ断面（StbSecSlabPrecast）、6.13. ＲＣ壁断面（StbSecWall_RC）、 
6.14. ＲＣ基礎断面（StbSecFoundation_RC）、 
6.19. ＲＣパラペット断面（StbSecParapet_RC） 
 strength_concrete について、省略された場合の記述を追記

## Page 659

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
3 
 
6.20.1. H 形鋼（StbSecRoll-H）、6.20.3. 角形鋼管（StbSecRoll-BOX）、 
6.20.4. 組立角形鋼管（StbSecBuild-BOX）、6.20.6.T 形鋼（StbSecRoll-T） 、 
6.20.7. 溝形鋼（StbSecRoll-C）、6.20.8.  山形鋼（StbSecRoll-L）、 
6.20.9. リップ溝形鋼（StbSecLipC） 
「鉄骨断面の基準方向」の説明を追記 
 （変更） 
 4.2.1. 節点（StbNode） 
 片持ち大梁先端はON_CANTI、片持ち床先端はOTHER とする 
 片持ち小梁先端はOTHER とする 
 5.12.1. 基礎柱（StbFoundationColumn） 
 id_section_FD とid_section_WR の「必須」を取りやめ、「補足」にどちらか一方を 
記述する旨を追記 
 6.2. ＲＣ柱断面（StbSecColumn_RC）、6.4. ＳＲＣ柱断面（StbSecColumn_SRC） 
 6.6. ＲＣ梁断面（StbSecBeam_RC）、6.8. ＳＲＣ梁断面（StbSecBeam_SRC） 
 StbSecFigureXXXX の最小回数1→0 0 となる場合について追記 
 6.2. ＲＣ柱断面（StbSecColumn_RC）、6.4. ＳＲＣ柱断面（StbSecColumn_SRC）、 
 6.5. ＣＦＴ柱断面（StbSecColumn_CFT） 
 strength_concrete の「所属階」と一致する、をid_node_top が所属する、に変更 
 6.6. ＲＣ梁断面（StbSecBeam_RC）、6.8. ＳＲＣ梁断面（StbSecBeam_SRC） 
 strength_concrete の「所属階」と一致する、をid_node_start が所属する、に変更 
 （追加） 
 ST-Bridge（ST_BRIDGE） XML Schema に関する属性追加 
（名前空間（Namespace）の節に説明文追加） 
 3.2. 共通情報（StbCommon） 
アプリケーションのバージョンapp_version を追加（必須） 
変換プログラム名とバージョンconvert_app_name,convert_app_version を追加 
 グローバル座標との関係を表すglobal_offset_X,global_offset_Y,global_offset_Z, 
global_rotation を追加 
 4.2.4. 順序のある節点ID（StbNodeIdOrder） 参照する親要素名を追加 
4.4.1. 階（StbStory） strength_concrete について、指定した階と部材との関係を追加 
 5.2.1. 柱（StbColumn）属性strength_concrete 追加、kind_structure にUNDEFINED 
追加、子要素追加 
 5.2.2. 柱中間節点（StbColumnViaNode）一本部材扱いに伴い要素追加 
 5.2.3. 中間節点オフセットリスト（StbMemberOffsetList）一本部材扱いに伴い要素追加 
 5.4.1. 大梁（StbGirder）属性strength_concrete 追加、kind_structure にUNDEFINED 
追加、子要素追加

## Page 660

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
4 
 
 5.4.2. 大梁中間節点（StbGirderViaNode）一本部材扱いに伴い要素追加 
5.7.1. スラブ（StbSlab） 属性strength_concrete 追加 
5.8.1. 壁（StbWall） 属性strength_concrete 追加 
 6.1. 断面情報（StbSections）子要素StbSecUndefined 追加 
 6.4.9. ＳＲＣ柱断面鉄骨形状（StbSecSteelFigureColumn_SRC） 
非埋込タイプにUNEMBEDDED2（RC 扱い）追加 
 6.20. 鉄骨断面（StbSecSteel）子要素StbSecSteelProduct,StbSecSteelUndefined 追加 
 6.20.12. 鉄骨製品（StbSecSteelProduct）追加 
 6.20.13. 未定義鉄骨断面（StbSecSteelUndefined）ダミー部材追加に伴い要素追加 
 6.21. 構造種別に依存しない断面（StbSecUndefined）ダミー部材追加に伴い要素追加 
 7.2. Ｓ梁継手・Ｈ形（StbJointBeamShapeH）      GUID を追加 
 7.3. Ｓ柱継手・Ｈ形（StbJointColumnShapeH）    GUID を追加 
 7.4. Ｓ柱継手・Ｔ形（StbJointColumnShapeT）    GUID を追加 
 7.5. Ｓ柱継手・＋形（StbJointColumnShapeCross）GUID を追加 
（訂正） 
誤字・脱字を訂正 
2019.3.20 
ver.2.0(Revision 1) 
 
（補足） 
 
節点（StbNode） 図の差し替え、補足説明の見直し 
 
開口（StbOpen） 補足説明を追記 
SRC 柱断面鉄骨形状（StbSecSteelFigureColumn_SRC） 補足説明を追記 
 
RC 開口断面（StbSecOpen） 省略時の扱い等、補足説明を追記 
 
（訂正） 
スラブ（StbSlab） 属性「id_section」の型 string→integer 
 
開口（StbOpen） 属性「id_section」の必須条件 必須→（なし） 
SRC 柱断面鉄骨形状（StbSecSteelFigureColumn_SRC） 
  StbSecSteelColumn_SRC_NotSame 子要素回数 2→1 
  StbSecSteelColumn_SRC_ThreeTypes 子要素回数 3→2 
SRC 柱断面鉄骨形状・柱頭脚別（StbSecSteelColumn_SRC_NotSame） 
  子要素回数 2→1（全子要素） 
SRC 柱断面鉄骨形状・３種類（StbSecSteelColumn_SRC_ThreeTypes） 
  子要素回数 3→1（全子要素） 
梁Ｘ形配筋（StbSecBeamXReinforced） 要素名StbSecBeamXReinforced→StbSecBar- 
BeamXReinforced（目次,要素リファレンス 要素一覧も修正） 
RC スラブ断面配筋（StbSecBarArrangementSlab_RC） 
  子要素「StbSecBarSlab_RC_Open」の記載なし→記載

## Page 661

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
5 
 
スラブ開口配筋（StbSecBarSlab_RC_Open） 記載なし→6.10.10.に記載  
RC 基礎断面形状（StbSecFoundation_RC_Rect） 属性「width_X」の型 string→double 
RC スラブ断面配筋・１方向１、２（StbSecBarSlab_RC_1Way1、2） 
デッキスラブ断面配筋・１方向（StbSecBarSlabDeck1Way） 
RC 基礎断面配筋・三角、連続（StbSecBarFoundation_RC_Triangle、Continuous） 
  属性「pos」の値 TRANSVERS_TOP → TRANSVERSE_TOP 
           TRANSVERS_BOTTOM → TRANSVERSE_BOTTOM 
RC 開口断面（StbSecOpen_RC） 
  子要素「StbSecBarArrangementOpen_RC」の最小回数 1→0 
RC 開口断面配筋（StbSecBarArrangementOpen_RC） 
  子要素「StbSecBarOpen_RC_Slab」「StbSecBarOpen_RC_Wall」の最小回数1→0 
※その他、誤字・脱字を訂正 
 
2018.7.1 
ver.2.0 リリース 
 
修正が多岐にわたるため、主な修正点のみ記載する。 
 
（補足） 
 
省略時の扱いなどを補足に追記するとともに、例を追記 
 
（変更） 
 
id の一意性の範囲を「それぞれの子要素内で一意」に変更 
 
要素名・属性名の命名規則を明確にし、規則に合わせて要素名・属性名を変更 
 
鉄筋情報を子要素に集約し、コンクリート・鉄筋・鉄骨の子要素の持ち方を統一 
 
節点に順序がある場合（床・壁の周辺節点）はモノリストに変更 
 
部材の平行移動オフセットを削除し、3 次元オフセットに統一 
 
XY 軸（StbX_Axis・StbY_Axis）を平行軸に変更 
 
（追加） 
 
節点・部材・断面にguid（オプション）を追加 
 
軸（StbAxes）に、円弧軸、放射軸を追加 
 
基礎柱（StbFoudationColumn）に根巻柱用の属性を追加 
 
杭（StbPile）に既製杭、鋼管杭を追加 
 
（削除） 
 
床組（StbSlabFrames）を削除 
 
2016.6.1 
ver.1.4 リリース 
 
（補足） 
 
軸（StbX_Axis、StbY_Axis）に「属する節点はグリッド上の設定のみ列挙」を追記 
 
節点（StbNode）のkind に関する補足を追記

## Page 662

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
6 
 
 
溝形鋼（StbSecRoll-C）、リップ溝形鋼（StbSecRoll-LipC）の向きを追記 
 
（変更） 
 
大梁（StbGirder）の図を変更し、ジョイント位置を節点からの距離に変更 
RC 基礎断面の連続（StbSecContinuous）の「RIGHT-L・LEFT-L・REVERSE-T」を
「RIGHT_L・LEFT_L・REVERSE_T」に変更 
 
山形鋼（StbSecRoll-L）の図を変更し、断面の上下を変更 
 
（追加） 
節点（StbNode）のkind にON_COLUMN を追加 
柱（StbColumn）・大梁（StbGirder）およびS 柱断面、SRC 柱断面、S 梁断面、SRC 梁断
面に継手ID を追加し、継手情報（StbJoints）以降を追加 
RC 梁断面（StbSecBeam_RC）にカットオフ筋長さを追加 
既製スラブ断面（StbSecSlab_Precast）、スラブ製品（StbSecSlab_Product）を追加 
子要素を拡張する仕様（StbExtElement、StbExtPropertyDef）を追加 
（訂正） 
正三角形基礎（StbSecEqui_Triangle）のスペルミスを訂正 
2015.1.21 
ver.1.3 リリース 
 
（補足） 
 
各要素詳細説明のはじめに「部材位置に関して」を追記 
 
円形柱の帯筋本数の数え方を追記 
 
鉄骨梁の始端のみにハンチがある場合の定義方法を追記 
 
（追加） 
 
スラブ（StbSlab）にハンチ形状（type_haunch）を追加 
 
スラブ断面（StbSecSlab_RC）に土間属性（isEarthen）を追加 
壁（StbWall）に土圧壁属性（isPress）を追加 
壁のダブル配筋（内外異なる）（StbSecInside_And_Outside）に位置２（pos2）を追加 
基礎の形状として正三角形（StbSecEqui_Triangle）と八角形（StbSecOctagon）を追加 
基礎の配筋として三方配筋（StbSecTreeWay）を追加 
 
（訂正） 
柱のX型配筋（StbSecRect_Column_XReinforced）が柱の配筋（StbSecBar_Arrangement）
の子要素に漏れていたのを訂正 
円形柱の配筋（StbSecCircle_Column_Same, StbSecCircle_Column_Not_Same）の帯筋
本数（count_band）を必須からオプションに変更 
パラペット断面（StbSecParapet_RC）の配筋タイプ（StbSecBar_Arrangement）の最大
回数を1 回に訂正 
 
2014.6.3 
ver.1.2 リリース

## Page 663

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
7 
 
（補足） 
StbStories の代表高さ、コンクリート強度に関する補足を追記 
柱に「中折れ柱は対象外」を追記 
部材の座標系、オフセットの座標系に関する補足を追記 
StbSecRect_Column_Same、StbSecRect_Column_Not_Same、 
StbSecCircle_Column_Same、StbSecCircle_Column_Not_Same のcount_axial に「RC
柱のみ」を追記 
   （追加） 
大梁（StbGirder）と梁断面（StbSecBeam_RC、StbSecBeam_S、StbSecBeam_SRC）
に外端・内端指定を追加 
基礎柱（StbFoundationColumns、StbFoundationColumn）を追加 
パラペット（StbParapets、StbParapet）と断面（StbSecParapet_RC、 
StbSecBar_Arrangement、StbSecSingle、StbSecZigzag、StbSecDouble_Net、 
StbSecParapetTip、StbSecParapetEdge）を追加 
壁断面に共通の開口配筋（StbSecOpen_Wall）を追加 
   （訂正） 
StbSlabOffset のoffset_x,offset_y→offset_X,offset_Y に訂正 
StbSecColum_SRC_Shape〇→StbSecColumn_SRC_Shape〇に訂正 
StbSecBeam_XReinforeced→StbSecBeam_XReinforced に訂正 
StbSecSlab_Deck のdepth_concrete の型をinteger→double に訂正 
StbSec1Way_Slab の最大回数を4→5 に訂正 
StbSecDeck_Product のdeck_depth の型をinteger→double に訂正 
StbSecOpen_RC のname を必須→オプションに訂正 
 
2013.7.25 
ver.1.1 リリース 
（補足） 
 
「本章で扱う用語に関して」を補足し、記載例を修正 
 
StbNode のON_GRID に関する補足を追記 
 
X 軸、Y 軸の説明をStbAxes の補足に追記 
（追加） 
 
ST-Bridge に定義されてない属性をアプリケーションが独自に拡張するフレームワーク
（StbExtensions、StbExtension、StbExtObject、StbExtProperty を追加 
 
作図用軸StbDrawingAxis を追加 
 
StbStory の階属性（kind）に免震階（ISOLATION）を追加 
 
StbColumn に始端・終端のオフセット（offset_bottom_X～offset_top_Z）を追加 
 
StbGirder に始端・終端のオフセット（offset_start_X～offset_end_Z）を追加

## Page 664

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
8 
 
 
StbBrace に始端・終端のオフセット（offset_start_X～offset_end_Z）を追加 
 
StbSlab の子要素にスラブオフセットリスト（StbSlabOffset_List）を、その子要素にス
ラブオフセット（StbSlabOffset）追加 
 
StbWall にオフセット（offset）を追加 
 
StbSecColumn_RC
に主筋重心位置（center_reinforcement_start_X ～center_ 
reinforcement _end_Y）を追加 
 
柱のStbSecBar_Arrangement の子要素にX 型配筋（StbSecRect_Column_XRainforced）
を追加 
 
StbSecBeam_RC,StbSecBeam_S,StbSecBeam_SRC に片持ち梁（isCanti）を追加 
 
StbSecBeam_RC,StbSecBeam_SRC に副主筋径（D_reinforcement_2nd_main）、副主
筋強度（strength_reinforcement_2nd_main）を追加 
 
StbSecBeam_RC,StbSecBeam_SRC に主筋重心位置（center_ reinforcement_top, 
center_ reinforcement_bottom）を追加 
 
梁のStbSecBar_Arrangement の子要素にX 型配筋（StbSecBeam_XRainforced）を追
加 
 
StbSecBeam_Same_Section,StbSecBeam_Start_Center_End_Section, 
StbSecBeam_Start_End_Section
に副主筋本数（count_2nd_main_top_1st ～
count_2nd_main_bottom_3rd)を追加 
 
StbSecSlab_RC に片持ちスラブ（isCanti）を追加 
 
壁のStbSecBar_Arrangement の子要素に壁端部補強筋（StbSecWallEdge）を追加 
 
StbSecPile_RC に拡頭部かぶり厚さ（depth_cover_top）を追加 
 
StbSlabFramePattern、StbSlabFrameBeam にスラブ、小梁のレベルを追加 
（訂正） 
 
最小回数の訂正（StbNodeid_List） 
 
属性名の訂正（StbSecRect_Column_Not_Same のcount_total_main → 
count_main_total） 
 
型の訂正（StbSecStandard_Slab 、StbSec2Way_Slab 、StbSec1Way_Slab_1 、
StbSec1Way_Slab_2 のpos、StbSecContinuous のtype：double→string） 
 
StbSlabFramePattern、StbSlabFrameBeam が参照するスラブ、小梁のID を部材ID
（StbSlab、StbBeam のID）から断面ID（StbSecSlab_xx、StbSecBeam_xx）に訂正 
 
2012.7.25 
ver.1.0 リリース 
（追加・変更） 
 
ファイル拡張子の定義を追加 
 
コンクリート、鉄筋、鉄骨の表記ルールを追加

## Page 665

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
9 
 
 
StbSecBeam_RC、StbSecBeam_SRC、StbSecSlab_RC にisFoundation を追加（部材リ
ストで基礎梁・基礎スラブリストに表示する断面かを識別にするため） 
（訂正） 
 
StbSteel をStbSecSteel に訂正 
 
要素名の訂正（StbSecRoll-R→StbSecRoll-Bar） 
 
属性名の訂正（StbSlab のname_section→id_section、StbWall のtype_oudside→
type_outside） 
 
補足の訂正（鉄骨形状の参照StbSteel→StbSecSteel） 
 
2012.3.14 
Draft4 リリース 
（追加・変更） 
 
SRC 柱断面、CFT 柱断面、SRC 梁断面、床組の定義を追加 
 
節点（StbNode）の属性に「ON_SLAB」を追加 
 
杭（StbPile）の構造種別の補足に「当面はRC のみ」を追加 
 
矩形柱の柱頭柱脚別配筋（StbSecRect_Column_Not_Same）の属性名を柱頭柱脚同一配
筋（StbSecRect_Column_Same）に合わせる 
 
円形柱の柱頭柱脚別配筋（StbSecCircle_Column_Not_Same）の属性名を柱頭柱脚同一配
筋（StbSecCircle_Column_Same）に合わせる 
 
S 柱断面（StbSecColumn_S）の柱脚形式の埋込をCFT 柱（StbSecColumn_CFT）に合
わせる 
 
RC 梁断面のハンチ（StbSecHaunch）の補足に「変断面の部分、等断面の部分」を追記
し、省略する場合のルールを追記 
 
RC 杭断面（StbSecPile_RC）の属性として、コンクリート強度（strength_concrete）を
追加 
（訂正） 
 
要素名の訂正（StbSecPile_Top_Middle_Bottom →StbSecPile_Top_Center_Bottom、
StbSecRoll_RipC→StbSecRoll_LipC） 
 
属性名の訂正（StbStory のconcrete_strength 
 
型の訂正（StbGirder、StbStrip_Footing、StbPile のid_section） 
 
選択肢の訂正（StbGirder のkind_haunch_start・kind_haunch_end のSLOPE） 
 
最大回数（StbSecOpen_Slab、StbSecOpen_Wall） 
 
2011.9.xx 
Draft3 リリース 
 
属性の型の統一（例：帯筋ピッチinteger→double） 
 
配筋情報の属性に位置（pos）を追加し、必要回数繰り返す方法に変更 
 
鉄骨断面を追加

## Page 666

ST-Bridge XML ファイル仕様書（ver.2.1）  
2023/03/31 
10 
 
 
StbColumn の属性に「ジョイント種別」を追加 
 
StbGirder の属性に「ジョイント種別」を追加 
 
StbBraces、StbBrace を追加 
 
StbSecColumn_S、StbSecBeam_S、StbSecBrace、StbSecSlab_Deck、StbSecSteel とそ
の子要素を追加 
 
ブレース・スラブについて部材座標系と全体座標系の対応を追加 
2011.5.20 
Draft2 リリース 
 
属性名を小文字に統一（例：StbProjectName→project_name、例外：ID） 
 
属性の型の見直し（例：かぶり厚さinteger→double） 
 
属性の必須項目を見直し（例：巾止筋） 
 
属性の補足に例を記載（例：鉄筋強度 SD345） 
 
属性の選択肢を大文字に統一（例：階属性 GENERAL） 
 
StbCommon の属性に「鉄骨の規格」を追加 
 
StbColumn の属性名で、_top と_bottom を入れ替え 
 
StbColumn のオフセット、ふかし厚さの属性名を修正 
 
StbSecColumn_RC の属性に「副主筋強度」を追加 
 
StbSecRect_Column_Same 等の属性に「副主筋」を追加 
 
StbSecRect_Column_Same 等の「帯筋ピッチ」をXY 別から１つに変更 
 
2010.11.27 
Draft1 リリース
