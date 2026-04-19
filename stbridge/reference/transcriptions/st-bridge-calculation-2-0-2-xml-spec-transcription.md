---
type: source
title: ST-Bridge XML ファイル仕様説明書 計算データ編 ver.2.0.2 PDF転記
tags: [pdf-transcription, st-bridge, calculation-data]
related: []
created: 2026-04-19
updated: 2026-04-19
authors: []
year: 2021
url: "/Users/keikamikawa/go/src/github.com/Code-Hex/bim/st-bridge/ST-Bridge計算データ編_XML仕様説明書ver.2.0.2.pdf"
venue: "local PDF"
---

# ST-Bridge XML ファイル仕様説明書 計算データ編 ver.2.0.2 PDF転記

> 機械抽出による Markdown 転記。代表ページは vision で目視確認済み。表組みは PDF 上の罫線や段組みを完全には保持していない。

- 元 PDF: `/Users/keikamikawa/go/src/github.com/Code-Hex/bim/st-bridge/ST-Bridge計算データ編_XML仕様説明書ver.2.0.2.pdf`
- ページ数: 297

## PDF ブックマーク

- p.11: 1. はじめに
  - p.13: 1.1. 2018年6月改訂にあたり
  - p.15: 1.2. 2021年2月改訂にあたり
- p.16: 2. 概要
  - p.16: 2.1. コンセプト
  - p.16: 2.2. XML形式
  - p.17: 2.3. 表記法
  - p.18: 2.4. 命名規則
  - p.18: 2.5. 属性値の型と表現範囲
  - p.19: 2.6. 単位系と数値範囲
  - p.20: 2.7. グローバル一意識別子（GUID）
  - p.20: 2.8. 集合型（monolist）
  - p.21: 2.9. ファイルの拡張子
  - p.21: 2.10. バージョン番号
  - p.22: 2.11. 仕様書の見方
- p.25: 3. 要素リファレンス　特記事項
  - p.25: 3.1. 全体構成
  - p.26: 3.2. ST-Bridge本編の全体構成
  - p.27: 3.3. 構造設計モデルと解析モデル
  - p.30: 3.4. 座標系について
  - p.32: 3.5. 主要な要素のIDと一意性
- p.34: 4. StbCalData 要素リファレンス
  - p.34: 4.1. 計算条件：StbCalData
  - p.35: 4.2. 計算共通条件：StbCalCommon
    - p.36: 4.2.1. 荷重条件：StbCalLoadCondition
    - p.37: 4.2.2. 地震荷重条件：StbCalSeismicCondition
    - p.38: 4.2.3. 風荷重条件：StbCalWindCondition
    - p.39: 4.2.4. 積雪荷重条件：StbCalSnowCondition
    - p.40: 4.2.5. 積載荷重（複数）：StbCalLiveloads
    - p.41: 4.2.6. 積載荷重：StbCalLiveload
    - p.42: 4.2.7. 床区分（複数）：StbCalFloorDividedAreas
    - p.43: 4.2.8. 床区分：StbCalFloorDividedArea
    - p.44: 4.2.9. 床区分を定義する階：StbCalStoryDivided
    - p.45: 4.2.10. 柱(計算用・分割)（複数）：StbCalColumns
    - p.46: 4.2.11. 柱(計算用・分割)：StbCalColumn
    - p.48: 4.2.12. 大梁(計算用・分割)（複数）：StbCalGirders
    - p.49: 4.2.13. 大梁(計算用・分割)：StbCalGirder
  - p.53: 4.3. 荷重定義：StbCalLoad
    - p.54: 4.3.1. 仕上げ定義：StbCalFinish
    - p.55: 4.3.2. 標準仕上げ（ＲＣ部材）：StbCalFinish_RC
    - p.57: 4.3.3. 標準仕上げ（Ｓ部材）：StbCalFinish_S
    - p.58: 4.3.4. 部材種別ごとの標準仕上げ（Ｓ部材）：StbCalFinish_S_member
    - p.59: 4.3.5. 階別仕上げ（複数）：StbCalFloorFinishes
    - p.60: 4.3.6. 階別仕上げ：StbCalFloorFinish
    - p.61: 4.3.7. 階別標準仕上げ（ＲＣ部材）：StbCalFloorFinish_RC
    - p.63: 4.3.8. 階別標準仕上げ（Ｓ部材）：StbCalFloorFinish_S
    - p.64: 4.3.9. 部材種別ごとの階別標準仕上げ（Ｓ部材）：StbCalFloorFinish_S_member
    - p.65: 4.3.10. 部材仕上げ（ＲＣ部材）：StbCalMemberFinishes_RC
    - p.66: 4.3.11. 柱仕上げ（ＲＣ部材）：StbCalColumnFinish_RC
    - p.67: 4.3.12. 梁仕上げ（ＲＣ部材）：StbCalGirderFinish_RC
    - p.68: 4.3.13. スラブ仕上げ：StbCalSlabFinish_RC
    - p.69: 4.3.14. 壁仕上げ：StbCalWallFinish_RC
    - p.70: 4.3.15. 部材仕上げ（Ｓ部材）：StbCalMemberFinishes_S
    - p.71: 4.3.16. 柱仕上げ（Ｓ部材）：StbCalColumnFinish_S
    - p.73: 4.3.17. 梁仕上げ（Ｓ部材）：StbCalGirderFinish_S
    - p.75: 4.3.18. ブレース仕上げ：StbCalBraceFinish_S
    - p.77: 4.3.19. 荷重ケース（複数）：StbCalLoadCases
    - p.78: 4.3.20. 荷重ケース：StbCalLoadCase
    - p.80: 4.3.21. 特殊荷重：StbCalAdditionalLoads
    - p.81: 4.3.22. 部材特殊荷重：StbCalMemberLoad
    - p.85: 4.3.23. 面特殊荷重：StbCalAreaLoad
    - p.86: 4.3.24. 節点特殊荷重：StbCalPointLoad
    - p.88: 4.3.25. 任意点特殊荷重：StbCalSelectedPointLoad
    - p.89: 4.3.26. 土圧・水圧荷重：StbCalEarthHydrostaticPressureLoad
    - p.90: 4.3.27. 追加・補正重量：StbCalAddedWeights
    - p.91: 4.3.28. 節点追加・補正重量：StbCalNodeAddedWeight
    - p.92: 4.3.29. 任意点追加・補正重量：StbCalSelectedNodeAddedWeight
    - p.93: 4.3.30. 地震荷重：StbCalSeismic
    - p.94: 4.3.31. 地震力計算グループ（複数）：StbCalSeismicConditionGroups
    - p.95: 4.3.32. 地震力計算グループ：StbCalSeismicConditionGroup
    - p.96: 4.3.33. 方向別 地震力計算条件（複数）：StbCalSeismicDirections
    - p.97: 4.3.34. 方向別 地震力計算条件：StbCalSeismicDirection
    - p.99: 4.3.35. 補正地震用重量（複数）：StbCalSeismicWeights
    - p.100: 4.3.36. 補正地震用重量：StbCalSeismicWeight
    - p.101: 4.3.37. 補正重心計算用重量（複数）：StbCalGravityPointWeights
    - p.102: 4.3.38. 補正重心計算用重量：StbCalGravityPointWeight
    - p.103: 4.3.39. 層せん断力係数（複数）：StbCalShearcoefficients
    - p.104: 4.3.40. 層せん断力係数：StbCalShearcoefficient
    - p.105: 4.3.41. 層せん断力（複数）：StbCalShearforces
    - p.106: 4.3.42. 層せん断力：StbCalShearforce
    - p.107: 4.3.43. 地震水平力（複数）：StbCalEarthquakeforces
    - p.108: 4.3.44. 地震水平力：StbCalEarthquakeforce
  - p.109: 4.4. 計算条件定義：StbCalCondition
    - p.110: 4.4.1. 部材材端条件：StbCalMemberConditions
    - p.111: 4.4.2. 柱材端条件：StbCalColumnCondition
    - p.113: 4.4.3. 大梁材端条件：StbCalGirderCondition
    - p.114: 4.4.4. 部材剛域長さ：StbCalMemberRigidzones
    - p.115: 4.4.5. 柱剛域長さ：StbCalColumnRigidzone
    - p.116: 4.4.6. 大梁剛域長さ：StbCalGirderRigidzone
    - p.117: 4.4.7. 部材危険断面位置：StbCalMemberCriticalPositions
    - p.118: 4.4.8. 柱危険断面位置：StbCalColumnCriticalPosition
    - p.119: 4.4.9. 大梁危険断面位置：StbCalGirderCriticalPosition
    - p.120: 4.4.10. 部材断面性能（数値指定）：StbCalSectionProperties
    - p.121: 4.4.11. 柱断面性能（数値指定）：StbCalColumnSectionProperty
    - p.122: 4.4.12. 大梁断面性能（数値指定）：StbCalGirderSectionProperty
    - p.123: 4.4.13. 壁断面性能（数値指定）：StbCalWallSectionProperty
    - p.124: 4.4.14. ブレース断面性能（数値指定）：StbCalBraceSectionProperty
    - p.125: 4.4.15. スラブ断面性能（数値指定）：StbCalSlabSectionProperty
    - p.126: 4.4.16. 部材剛性倍率：StbCalMemberStiffnesses
    - p.127: 4.4.17. 柱剛性倍率：StbCalColumnStiffness
    - p.128: 4.4.18. 大梁剛性倍率：StbCalGirderStiffness
    - p.129: 4.4.19. 壁剛性倍率：StbCalWallStiffness
    - p.130: 4.4.20. ブレース剛性倍率：StbCalBraceStiffness
    - p.131: 4.4.21. 節点拘束条件（複数）：StbCalNodeRestrictions
    - p.132: 4.4.22. 節点拘束条件：StbCalNodeRestriction
    - p.134: 4.4.23. 接合部パネル（複数）：StbCalNodePanels
    - p.135: 4.4.24. 接合部パネル：StbCalNodePanel
    - p.136: 4.4.25. 剛床条件（複数）：StbCalFloorDiaphragms
    - p.137: 4.4.26. 剛床条件：StbCalFloorDiaphragm
  - p.138: 4.5. 荷重配置：StbCalLoadArrangements
    - p.139: 4.5.1. 柱仕上げ（ＲＣ部材）配置：StbCalColumnFinish_RC_Arr
    - p.140: 4.5.2. 柱仕上げ（ＲＣ部材）配置・荷重：StbCalColumnFinish_RC_LoadList
    - p.140: 4.5.3. 柱仕上げ（ＲＣ部材）配置・部材：StbCalColumnFinish_RC_MemList
    - p.141: 4.5.4. 柱仕上げ（ＲＣ部材）配置・部材(計算用・分割)：StbCalColumnFinish_RC_CalMemList
    - p.142: 4.5.5. 柱仕上げ（Ｓ部材）配置：StbCalColumnFinish_S_Arr
    - p.143: 4.5.6. 柱仕上げ（Ｓ部材）配置・荷重：StbCalColumnFinish_S_LoadList
    - p.143: 4.5.7. 柱仕上げ（Ｓ部材）配置・部材：StbCalColumnFinish_S_MemList
    - p.144: 4.5.8. 柱仕上げ（Ｓ部材）配置・部材(計算用・分割)：StbCalColumnFinish_S_CalMemList
    - p.145: 4.5.9. 柱特殊荷重配置：StbCalColumnMemberLoadArr
    - p.146: 4.5.10. 柱特殊荷重配置・荷重：StbCalColumnMemberLoadList
    - p.146: 4.5.11. 柱特殊荷重配置・部材：StbCalColumnMemberLoadMemList
    - p.147: 4.5.12. 柱特殊荷重配置・部材(計算用・分割)：StbCalColumnMemberLoadCalMemList
    - p.148: 4.5.13. 大梁仕上げ（ＲＣ部材）配置：StbCalGirderFinish_RC_Arr
    - p.149: 4.5.14. 大梁仕上げ（ＲＣ部材）配置・荷重：StbCalGirderFinish_RC_LoadList
    - p.149: 4.5.15. 大梁仕上げ（ＲＣ部材）配置・部材：StbCalGirderFinish_RC_MemList
    - p.150: 4.5.16. 大梁仕上げ（ＲＣ部材）配置・部材(計算用・分割)：StbCalGirderFinish_RC_CalMemList
    - p.151: 4.5.17. 大梁仕上げ（Ｓ部材）配置：StbCalGirderFinish_S_Arr
    - p.152: 4.5.18. 大梁仕上げ（Ｓ部材）配置・荷重：StbCalGirderFinish_S_LoadList
    - p.152: 4.5.19. 大梁仕上げ（Ｓ部材）配置・部材：StbCalGirderFinish_S_MemList
    - p.153: 4.5.20. 大梁仕上げ（Ｓ部材）配置・部材(計算用・分割)：StbCalGirderFinish_S_CalMemList
    - p.154: 4.5.21. 大梁特殊荷重配置：StbCalGirderMemberLoadArr
    - p.155: 4.5.22. 大梁特殊荷重配置・荷重：StbCalGirderMemberLoadList
    - p.155: 4.5.23. 大梁特殊荷重配置・部材：StbCalGirderMemberLoadMemList
    - p.156: 4.5.24. 大梁特殊荷重配置・部材(計算用・分割)：StbCalGirderMemberLoadCalMemList
    - p.157: 4.5.25. 小梁仕上げ（ＲＣ部材）配置：StbCalBeamFinish_RC_Arr
    - p.158: 4.5.26. 小梁仕上げ（ＲＣ部材）配置・荷重：StbCalBeamFinish_RC_LoadList
    - p.158: 4.5.27. 小梁仕上げ（ＲＣ部材）配置・部材：StbCalBeamFinish_RC_MemList
    - p.159: 4.5.28. 小梁仕上げ（Ｓ部材）配置：StbCalBeamFinish_S_Arr
    - p.160: 4.5.29. 小梁仕上げ（Ｓ部材）配置・荷重：StbCalBeamFinish_S_LoadList
    - p.160: 4.5.30. 小梁仕上げ（Ｓ部材）配置・部材：StbCalBeamFinish_S_MemList
    - p.161: 4.5.31. 小梁特殊荷重配置：StbCalBeamMemberLoadArr
    - p.162: 4.5.32. 小梁特殊荷重配置・荷重：StbCalBeamMemberLoadList
    - p.162: 4.5.33. 小梁特殊荷重配置・部材：StbCalBeamMemberLoadMemList
    - p.163: 4.5.34. ブレース仕上げ配置：StbCalBraceFinish_S_Arr
    - p.164: 4.5.35. ブレース仕上げ配置・荷重：StbCalBraceFinish_S_LoadList
    - p.164: 4.5.36. ブレース仕上げ配置・部材：StbCalBraceFinish_S_MemList
    - p.165: 4.5.37. 床積載荷重配置：StbCalSlabLiveLoadArr
    - p.166: 4.5.38. 床積載荷重配置・荷重：StbCalSlabLiveLoadList
    - p.166: 4.5.39. 床積載荷重配置・部材：StbCalSlabLiveLoadMemList
    - p.167: 4.5.40. 床スラブ仕上げ配置：StbCalSlabFinish_RC_Arr
    - p.168: 4.5.41. 床スラブ仕上げ配置・荷重：StbCalSlabFinish_RC_LoadList
    - p.168: 4.5.42. 床スラブ仕上げ配置・部材：StbCalSlabFinish_RC_MemList
    - p.169: 4.5.43. 床面特殊荷重配置：StbCalSlabAreaLoadArr
    - p.170: 4.5.44. 床面特殊荷重配置・荷重：StbCalSlabAreaLoadList
    - p.170: 4.5.45. 床面特殊荷重配置・部材：StbCalSlabAreaLoadMemList
    - p.171: 4.5.46. 床土圧・水圧荷重配置：StbCalSlabPressureLoadArr
    - p.172: 4.5.47. 床土圧・水圧荷重配置・荷重：StbCalSlabPressureLoadList
    - p.172: 4.5.48. 床土圧・水圧荷重配置・部材：StbCalSlabPressureLoadMemList
    - p.173: 4.5.49. 壁仕上げ配置：StbCalWallFinish_RC_Arr
    - p.174: 4.5.50. 壁仕上げ配置・荷重：StbCalWallFinish_RC_LoadList
    - p.174: 4.5.51. 壁仕上げ配置・部材：StbCalWallFinish_RC_MemList
    - p.175: 4.5.52. 壁面特殊荷重配置：StbCalWallAreaLoadArr
    - p.176: 4.5.53. 壁面特殊荷重配置・荷重：StbCalWallAreaLoadList
    - p.176: 4.5.54. 壁面特殊荷重配置・部材：StbCalWallAreaLoadMemList
    - p.177: 4.5.55. 壁土圧・水圧荷重配置：StbCalWallPressureLoadArr
    - p.178: 4.5.56. 壁土圧・水圧荷重配置・荷重：StbCalWallPressureLoadList
    - p.178: 4.5.57. 壁土圧・水圧荷重配置・部材：StbCalWallPressureLoadMemList
    - p.179: 4.5.58. 節点追加・補正重量配置：StbCalNodeWeightArr
    - p.180: 4.5.59. 節点追加・補正重量配置・荷重：StbCalNodeWeightLoadList
    - p.180: 4.5.60. 節点追加・補正重量配置・節点：StbCalNodeWeightNodeList
    - p.181: 4.5.61. 節点特殊荷重配置：StbCalNodePointLoadArr
    - p.182: 4.5.62. 節点特殊荷重配置・荷重：StbCalNodePointLoadList
    - p.182: 4.5.63. 節点特殊荷重配置・節点：StbCalNodePointLoadNodeList
  - p.183: 4.6. 計算条件配置：StbCalConditionArrangements
    - p.184: 4.6.1. 柱材端条件配置：StbCalColumnConditionArr
    - p.185: 4.6.2. 柱材端条件配置・条件：StbCalColumnConditionList
    - p.185: 4.6.3. 柱材端条件配置・部材：StbCalColumnConditionMemList
    - p.186: 4.6.4. 柱材端条件配置・部材(計算用・分割)：StbCalColumnConditionCalMemList
    - p.187: 4.6.5. 柱剛域長さ配置：StbCalColumnRigidzoneArr
    - p.188: 4.6.6. 柱剛域長さ配置・条件：StbCalColumnRigidzoneList
    - p.188: 4.6.7. 柱剛域長さ配置・部材：StbCalColumnRigidzoneMemList
    - p.189: 4.6.8. 柱剛域長さ配置・部材(計算用・分割)：StbCalColumnRigidzoneCalMemList
    - p.190: 4.6.9. 柱危険断面位置配置：StbCalColumnCriticalPositionArr
    - p.191: 4.6.10. 柱危険断面位置配置・条件：StbCalColumnCriticalPositionList
    - p.191: 4.6.11. 柱危険断面位置配置・部材：StbCalColumnCriticalPositionMemList
    - p.192: 4.6.12. 柱危険断面位置配置・部材(計算用・分割)：StbCalColumnCriticalPositionCalMemList
    - p.193: 4.6.13. 柱剛性倍率配置：StbCalColumnStiffnessArr
    - p.194: 4.6.14. 柱剛性倍率配置・条件：StbCalColumnStiffnessList
    - p.194: 4.6.15. 柱剛性倍率配置・部材：StbCalColumnStiffnessMemList
    - p.195: 4.6.16. 柱剛性倍率配置・部材(計算用・分割)：StbCalColumnStiffnessCalMemList
    - p.196: 4.6.17. 大梁材端条件配置：StbCalGirderConditionArr
    - p.197: 4.6.18. 大梁材端条件配置・条件：StbCalGirderConditionList
    - p.197: 4.6.19. 大梁材端条件配置・部材：StbCalGirderConditionMemList
    - p.198: 4.6.20. 大梁材端条件配置・部材(計算用・分割)：StbCalGirderConditionCalMemList
    - p.199: 4.6.21. 大梁剛域長さ配置：StbCalGirderRigidzoneArr
    - p.200: 4.6.22. 大梁剛域長さ配置・条件：StbCalGirderRigidzoneList
    - p.200: 4.6.23. 大梁剛域長さ配置・部材：StbCalGirderRigidzoneMemList
    - p.201: 4.6.24. 大梁剛域長さ配置・部材(計算用・分割)：StbCalGirderRigidzoneCalMemList
    - p.202: 4.6.25. 大梁危険断面位置配置：StbCalGirderCriticalPositionArr
    - p.203: 4.6.26. 大梁危険断面位置配置・条件：StbCalGirderCriticalPositionList
    - p.203: 4.6.27. 大梁危険断面位置配置・部材：StbCalGirderCriticalPositionMemList
    - p.204: 4.6.28. 大梁危険断面位置配置・部材(計算用・分割)：StbCalGirderCriticalPositionCalMemList
    - p.205: 4.6.29. 大梁剛性倍率配置：StbCalGirderStiffnessArr
    - p.206: 4.6.30. 大梁剛性倍率配置・条件：StbCalGirderStiffnessList
    - p.206: 4.6.31. 大梁剛性倍率配置・部材：StbCalGirderStiffnessMemList
    - p.207: 4.6.32. 大梁剛性倍率配置・部材(計算用・分割)：StbCalGirderStiffnessCalMemList
    - p.208: 4.6.33. ブレース剛性倍率配置：StbCalBraceStiffnessArr
    - p.209: 4.6.34. ブレース剛性倍率配置・条件：StbCalBraceStiffnessList
    - p.209: 4.6.35. ブレース剛性倍率配置・部材：StbCalBraceStiffnessMemList
    - p.210: 4.6.36. 壁剛性倍率配置：StbCalWallStiffnessArr
    - p.211: 4.6.37. 壁剛性倍率配置・条件：StbCalWallStiffnessList
    - p.211: 4.6.38. 壁剛性倍率配置・部材：StbCalWallStiffnessMemList
    - p.212: 4.6.39. 節点拘束条件配置：StbCalNodeRestrictionArr
    - p.213: 4.6.40. 節点拘束条件配置・条件：StbCalNodeRestrictionList
    - p.213: 4.6.41. 節点拘束条件配置・節点：StbCalNodeRestrictionNodeList
    - p.214: 4.6.42. 節点接合部パネル配置：StbCalNodePanelArr
    - p.215: 4.6.43. 節点接合部パネル配置・条件：StbCalNodePanelList
    - p.215: 4.6.44. 節点接合部パネル配置・節点：StbCalNodePanelNodeList
    - p.216: 4.6.45. 柱断面性能配置：StbCalColumnSecPropertyArr
    - p.217: 4.6.46. 柱断面性能配置・条件：StbCalColumnSecPropertyList
    - p.217: 4.6.47. 柱断面性能配置・ＲＣ断面：StbCalColumnSecProperty_RC_List
    - p.218: 4.6.48. 柱断面性能配置・Ｓ断面：StbCalColumnSecProperty_S_List
    - p.218: 4.6.49. 柱断面性能配置・ＳＲＣ断面：StbCalColumnSecProperty_SRC_List
    - p.219: 4.6.50. 柱断面性能配置・ＣＦＴ断面：StbCalColumnSecProperty_CFT_List
    - p.220: 4.6.51. 大梁断面性能配置：StbCalGirderSecPropertyArr
    - p.221: 4.6.52. 大梁断面性能配置・条件：StbCalGirderSecPropertyList
    - p.221: 4.6.53. 大梁断面性能配置・ＲＣ断面：StbCalGirderSecProperty_RC_List
    - p.222: 4.6.54. 大梁断面性能配置・Ｓ断面：StbCalGirderSecProperty_S_List
    - p.222: 4.6.55. 大梁断面性能配置・ＳＲＣ断面：StbCalGirderSecProperty_SRC_List
    - p.223: 4.6.56. ブレース断面性能配置：StbCalBraceSecPropertyArr
    - p.224: 4.6.57. ブレース断面性能配置・条件：StbCalBraceSecPropertyList
    - p.224: 4.6.58. ブレース断面性能配置・Ｓ断面：StbCalBraceSecProperty_S_List
    - p.225: 4.6.59. スラブ断面性能配置：StbCalSlabSecPropertyArr
    - p.226: 4.6.60. スラブ断面性能配置・条件：StbCalSlabSecPropertyList
    - p.226: 4.6.61. スラブ断面性能配置・ＲＣ断面：StbCalSlabSecProperty_RC_List
    - p.227: 4.6.62. スラブ断面性能配置・デッキプレートスラブ断面：StbCalSlabSecPropertyDeckList
    - p.227: 4.6.63. スラブ断面性能配置・既製スラブ断面：StbCalSlabSecPropertyPrecastList
    - p.228: 4.6.64. 壁断面性能配置：StbCalWallSecPropertyArr
    - p.229: 4.6.65. 壁断面性能配置・条件：StbCalWallSecPropertyList
    - p.229: 4.6.66. 壁断面性能配置・ＲＣ断面：StbCalWallSecProperty_RC_List
- p.230: 5. StbAnaModels要素リファレンス
  - p.230: 5.1. 解析モデル（複数）：StbAnaModels
  - p.231: 5.2. 解析モデル：StbAnaModel
    - p.232: 5.2.1. 解析用節点（複数）：StbAnaNodes
    - p.233: 5.2.2. 解析用節点：StbAnaNode
    - p.234: 5.2.3. 境界条件：StbAnaBoundary
    - p.235: 5.2.4. 解析用階（複数）：StbAnaStories
    - p.236: 5.2.5. 解析用階：StbAnaStory
    - p.237: 5.2.6. 解析用部材IDリスト：StbAnaMemberid_List
    - p.238: 5.2.7. 解析用部材ID：StbAnaMemberid
    - p.239: 5.2.8. 解析用部材：StbAnaMembers
    - p.240: 5.2.9. 梁要素（複数）：StbAnaBeams
    - p.241: 5.2.10. 梁要素：StbAnaBeam
    - p.242: 5.2.11. 梁要素剛域長さ：StbAnaBeamRigidzone
    - p.243: 5.2.12. 梁要素危険断面位置：StbAnaBeamCriticalPosition
    - p.244: 5.2.13. 梁要素材端ばね：StbAnaBeamEndSpring
    - p.245: 5.2.14. トラス要素（複数）：StbAnaTrusses
    - p.246: 5.2.15. トラス要素：StbAnaTruss
    - p.247: 5.2.16. 支点ばね要素（複数）：StbAnaSupports
    - p.248: 5.2.17. 支点ばね要素：StbAnaSupport
    - p.249: 5.2.18. ばね要素（複数）：StbAnaSprings
    - p.250: 5.2.19. ばね要素：StbAnaSpring
    - p.251: 5.2.20. 壁エレメント要素（複数）：StbAnaWalls
    - p.252: 5.2.21. 壁エレメント要素：StbAnaWall
    - p.253: 5.2.22. 平面要素（三角形、複数）： StbAnaPlaneTriangles
    - p.254: 5.2.23. 平面要素（三角形）：StbAnaPlaneTriangle
    - p.255: 5.2.24. 平面要素（四角形、複数）：StbAnaPlaneRectangles
    - p.256: 5.2.25. 平面要素（四角形）：StbAnaPlaneRectangle
    - p.257: 5.2.26. 接合部パネル要素（複数）：StbAnaNodePanels
    - p.258: 5.2.27. 接合部パネル要素：StbAnaNodePanel
    - p.259: 5.2.28. 解析用部材性能：StbAnaProperties
    - p.260: 5.2.29. 梁要素性能（複数）：StbAnaBeamProperties
    - p.261: 5.2.30. 梁要素性能：StbAnaBeamProperty
    - p.262: 5.2.31. トラス要素性能（複数）：StbAnaTrussProperties
    - p.263: 5.2.32. トラス要素性能：StbAnaTrussProperty
    - p.264: 5.2.33. ばね要素性能（複数）：StbAnaSpringProperties
    - p.265: 5.2.34. ばね要素性能：StbAnaSpringProperty
    - p.266: 5.2.35. 壁エレメント要素性能（複数）：StbAnaWallProperties
    - p.267: 5.2.36. 壁エレメント要素性能：StbAnaWallProperty
    - p.268: 5.2.37. 平面要素性能（複数）：StbAnaPlaneProperties
    - p.269: 5.2.38. 平面要素性能：StbAnaPlaneProperty
    - p.270: 5.2.39. 接合部パネル要素性能（複数）：StbAnaNodePanelProperties
    - p.271: 5.2.40. 接合部パネル要素性能：StbAnaNodePanelProperty
    - p.273: 5.2.41. 剛床指定（複数）：StbAnaFloorDiaphragms
    - p.274: 5.2.42. 剛床指定：StbAnaFloorDiaphragm
    - p.275: 5.2.43. 解析用節点IDリスト：StbAnaNodeid_List
    - p.275: 5.2.44. 解析用節点ID：StbAnaNodeid
    - p.276: 5.2.45. 材料性能（複数）：StbAnaMaterials
    - p.277: 5.2.46. 材料性能：StbAnaMaterial
    - p.278: 5.2.47. 断面性能（複数）：StbAnaSections
    - p.279: 5.2.48. 断面性能：StbAnaSection
    - p.280: 5.2.49. 荷重ケース（複数）：StbAnaLoadCases
    - p.280: 5.2.50. 荷重ケース：StbAnaLoadCase
    - p.281: 5.2.51. 節点荷重：StbAnaLoadNode
    - p.282: 5.2.52. 梁荷重：StbAnaLoadBeam
    - p.283: 5.2.53. 梁CMQ荷重：StbAnaLoadBeamCMQ
    - p.284: 5.2.54. 梁要素初期応力：StbAnaLoadBeamInitialStress
    - p.285: 5.2.55. トラス荷重：StbAnaLoadTruss
    - p.285: 5.2.56. トラス要素初期応力：StbAnaLoadTrussInitialStress
    - p.286: 5.2.57. 壁荷重：StbAnaLoadWall
    - p.286: 5.2.58. 壁エレメント要素初期応力：StbAnaLoadWallInitialStress
    - p.287: 5.2.59. ばね荷重：StbAnaLoadSpring
    - p.287: 5.2.60. ばね要素初期応力：StbAnaLoadSpringInitialStress
    - p.288: 5.2.61. 解析ケース（複数）：StbAnaAnalyses
    - p.288: 5.2.62. 静的線形解析ケース：StbAnaAnalysisStaticLinear
  - p.290: 5.3. 解析モデル関連定義：StbAnaRelations
    - p.291: 5.3.1. 解析用節点関連定義：StbAnaNodeRel
    - p.292: 5.3.2. 解析用階関連定義：StbAnaStoryRel
    - p.293: 5.3.3. 解析用部材関連定義：StbAnaMemberRel
    - p.294: 5.3.4. 解析用部材関連定義(計算用・分割)：StbAnaCalMemberRel
    - p.295: 5.3.5. 解析用部材性能関連定義：StbAnaPropertyRel
- p.297: 6. 修正履歴

## Page 1

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
i 
 
 
 
 
 
 
 
 
 
 
 
 
 
ST-Bridge XML ファイル仕様書 
計算データ編（ver.2.0） 
 
 
 
 
 
 
 
2021. 02. 28 
buildingSMART Japan 
構造設計小委員会

## Page 2

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
ii 
 
目次 
1. 
はじめに ..................................................................................................................................................... 1 
1.1. 
2018 年6 月改訂にあたり .................................................................................................................. 3 
1.2. 
2020 年12 月改訂にあたり ................................................................................................................ 5 
2. 
概要 ............................................................................................................................................................ 6 
2.1. 
コンセプト .......................................................................................................................................... 6 
2.2. 
XML 形式 ........................................................................................................................................... 6 
2.3. 
表記法 ................................................................................................................................................. 7 
2.4. 
命名規則 ............................................................................................................................................. 8 
2.5. 
属性値の型と表現範囲 ........................................................................................................................ 8 
2.6. 
単位系と数値範囲 ............................................................................................................................... 9 
2.7. 
グローバル一意識別子（GUID） .................................................................................................... 10 
2.8. 
集合型（monolist） ......................................................................................................................... 10 
2.9. 
ファイルの拡張子 .............................................................................................................................. 11 
2.10. 
バージョン番号 .............................................................................................................................. 11 
2.11. 
仕様書の見方 ................................................................................................................................ 12 
3. 
要素リファレンス 特記事項 ................................................................................................................... 15 
3.1. 
全体構成 ........................................................................................................................................... 15 
3.2. 
ST-Bridge 本編の全体構成 ............................................................................................................... 16 
3.3. 
構造設計モデルと解析モデル ........................................................................................................... 17 
3.4. 
座標系について ................................................................................................................................ 20 
3.5. 
主要な要素のID と一意性 ............................................................................................................... 22 
4. 
StbCalData 要素リファレンス ............................................................................................................... 24 
4.1. 
計算条件：StbCalData .................................................................................................................... 24 
4.2. 
計算共通条件：StbCalCommon ...................................................................................................... 25 
4.2.1. 
荷重条件：StbCalLoadCondition ............................................................................................ 26 
4.2.2. 
地震荷重条件：StbCalSeismicCondition ................................................................................ 27 
4.2.3. 
風荷重条件：StbCalWindCondition........................................................................................ 28 
4.2.4. 
積雪荷重条件：StbCalSnowCondition .................................................................................... 29 
4.2.5. 
積載荷重（複数）：StbCalLiveloads ...................................................................................... 30 
4.2.6. 
積載荷重：StbCalLiveload ...................................................................................................... 31 
4.2.7. 
床区分（複数）：StbCalFloorDividedAreas .......................................................................... 32 
4.2.8. 
床区分：StbCalFloorDividedArea .......................................................................................... 33 
4.2.9. 
床区分を定義する階：StbCalStoryDivided ............................................................................ 34 
4.2.10. 
柱(計算用・分割)（複数）：StbCalColumns ......................................................................... 35

## Page 3

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
iii 
 
4.2.11. 
柱(計算用・分割)：StbCalColumn ......................................................................................... 36 
4.2.12. 
大梁(計算用・分割)（複数）：StbCalGirders ....................................................................... 38 
4.2.13. 
大梁(計算用・分割)：StbCalGirder ....................................................................................... 39 
4.3. 
荷重定義：StbCalLoad .................................................................................................................... 43 
4.3.1. 
仕上げ定義：StbCalFinish ...................................................................................................... 44 
4.3.2. 
標準仕上げ（ＲＣ部材）：StbCalFinish_RC ......................................................................... 45 
4.3.3. 
標準仕上げ（Ｓ部材）：StbCalFinish_S ................................................................................ 47 
4.3.4. 
部材種別ごとの標準仕上げ（Ｓ部材）：StbCalFinish_S_member ....................................... 48 
4.3.5. 
階別仕上げ（複数）：StbCalFloorFinishes ........................................................................... 49 
4.3.6. 
階別仕上げ：StbCalFloorFinish ............................................................................................. 50 
4.3.7. 
階別標準仕上げ（ＲＣ部材）：StbCalFloorFinish_RC ......................................................... 51 
4.3.8. 
階別標準仕上げ（Ｓ部材）：StbCalFloorFinish_S ................................................................ 53 
4.3.9. 
部材種別ごとの階別標準仕上げ（Ｓ部材）：StbCalFloorFinish_S_member ....................... 54 
4.3.10. 
部材仕上げ（ＲＣ部材）：StbCalMemberFinishes_RC ........................................................ 55 
4.3.11. 
柱仕上げ（ＲＣ部材）：StbCalColumnFinish_RC ................................................................ 56 
4.3.12. 
梁仕上げ（ＲＣ部材）：StbCalGirderFinish_RC .................................................................. 57 
4.3.13. 
スラブ仕上げ：StbCalSlabFinish_RC .................................................................................... 58 
4.3.14. 
壁仕上げ：StbCalWallFinish_RC ........................................................................................... 59 
4.3.15. 
部材仕上げ（Ｓ部材）：StbCalMemberFinishes_S ............................................................... 60 
4.3.16. 
柱仕上げ（Ｓ部材）：StbCalColumnFinish_S ...................................................................... 61 
4.3.17. 
梁仕上げ（Ｓ部材）：StbCalGirderFinish_S ........................................................................ 63 
4.3.18. 
ブレース仕上げ：StbCalBraceFinish_S ................................................................................. 65 
4.3.19. 
荷重ケース（複数）：StbCalLoadCases ................................................................................ 67 
4.3.20. 
荷重ケース：StbCalLoadCase ................................................................................................ 68 
4.3.21. 
特殊荷重：StbCalAdditionalLoads ......................................................................................... 70 
4.3.22. 
部材特殊荷重：StbCalMemberLoad ....................................................................................... 71 
4.3.23. 
面特殊荷重：StbCalAreaLoad ................................................................................................. 75 
4.3.24. 
節点特殊荷重：StbCalPointLoad ............................................................................................ 76 
4.3.25. 
任意点特殊荷重：StbCalSelectedPointLoad .......................................................................... 78 
4.3.26. 
土圧・水圧荷重：StbCalEarthHydrostaticPressureLoad ..................................................... 79 
4.3.27. 
追加・補正重量：StbCalAddedWeights ................................................................................. 80 
4.3.28. 
節点追加・補正重量：StbCalNodeAddedWeight ................................................................... 81 
4.3.29. 
任意点追加・補正重量：StbCalSelectedNodeAddedWeight ................................................. 82 
4.3.30. 
地震荷重：StbCalSeismic ........................................................................................................ 83 
4.3.31. 
地震力計算グループ（複数）：StbCalSeismicConditionGroups .......................................... 84 
4.3.32. 
地震力計算グループ：StbCalSeismicConditionGroup ........................................................... 85

## Page 4

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
iv 
 
4.3.33. 
方向別 地震力計算条件（複数）：StbCalSeismicDirections ................................................ 86 
4.3.34. 
方向別 地震力計算条件：StbCalSeismicDirection ................................................................. 87 
4.3.35. 
補正地震用重量（複数）：StbCalSeismicWeights ................................................................. 89 
4.3.36. 
補正地震用重量：StbCalSeismicWeight ................................................................................. 90 
4.3.37. 
補正重心計算用重量（複数）：StbCalGravityPointWeights ................................................ 91 
4.3.38. 
補正重心計算用重量：StbCalGravityPointWeight ................................................................ 92 
4.3.39. 
層せん断力係数（複数）：StbCalShearcoefficients ............................................................... 93 
4.3.40. 
層せん断力係数：StbCalShearcoefficient ............................................................................... 94 
4.3.41. 
層せん断力（複数）：StbCalShearforces ............................................................................... 95 
4.3.42. 
層せん断力：StbCalShearforce ............................................................................................... 96 
4.3.43. 
地震水平力（複数）：StbCalEarthquakeforces ..................................................................... 97 
4.3.44. 
地震水平力：StbCalEarthquakeforce ..................................................................................... 98 
4.4. 
計算条件定義：StbCalCondition..................................................................................................... 99 
4.4.1. 
部材材端条件：StbCalMemberConditions ........................................................................... 100 
4.4.2. 
柱材端条件：StbCalColumnCondition ................................................................................. 101 
4.4.3. 
大梁材端条件：StbCalGirderCondition ................................................................................ 103 
4.4.4. 
部材剛域長さ：StbCalMemberRigidzones ........................................................................... 104 
4.4.5. 
柱剛域長さ：StbCalColumnRigidzone ................................................................................. 105 
4.4.6. 
大梁剛域長さ：StbCalGirderRigidzone ................................................................................ 106 
4.4.7. 
部材危険断面位置：StbCalMemberCriticalPositions .......................................................... 107 
4.4.8. 
柱危険断面位置：StbCalColumnCriticalPosition ................................................................ 108 
4.4.9. 
大梁危険断面位置：StbCalGirderCriticalPosition ............................................................... 109 
4.4.10. 
部材断面性能（数値指定）：StbCalSectionProperties ........................................................ 110 
4.4.11. 
柱断面性能（数値指定）：StbCalColumnSectionProperty ................................................. 111 
4.4.12. 
大梁断面性能（数値指定）：StbCalGirderSectionProperty................................................ 112 
4.4.13. 
壁断面性能（数値指定）：StbCalWallSectionProperty ...................................................... 113 
4.4.14. 
ブレース断面性能（数値指定）：StbCalBraceSectionProperty .......................................... 114 
4.4.15. 
スラブ断面性能（数値指定）：StbCalSlabSectionProperty................................................ 115 
4.4.16. 
部材剛性倍率：StbCalMemberStiffnesses ........................................................................... 116 
4.4.17. 
柱剛性倍率：StbCalColumnStiffness ................................................................................... 117 
4.4.18. 
大梁剛性倍率：StbCalGirderStiffness .................................................................................. 118 
4.4.19. 
壁剛性倍率：StbCalWallStiffness ........................................................................................ 119 
4.4.20. 
ブレース剛性倍率：StbCalBraceStiffness ............................................................................ 120 
4.4.21. 
節点拘束条件（複数）：StbCalNodeRestrictions ................................................................ 121 
4.4.22. 
節点拘束条件：StbCalNodeRestriction ................................................................................ 122 
4.4.23. 
接合部パネル（複数）：StbCalNodePanels ......................................................................... 124

## Page 5

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
v 
 
4.4.24. 
接合部パネル：StbCalNodePanel ......................................................................................... 125 
4.4.25. 
剛床条件（複数）：StbCalFloorDiaphragms ....................................................................... 126 
4.4.26. 
剛床条件：StbCalFloorDiaphragm ....................................................................................... 127 
4.5. 
荷重配置：StbCalLoadArrangements .......................................................................................... 128 
4.5.1. 
柱仕上げ（ＲＣ部材）配置：StbCalColumnFinish_RC_Arr ............................................... 129 
4.5.2. 
柱仕上げ（ＲＣ部材）配置・荷重：StbCalColumnFinish_RC_LoadList ........................... 130 
4.5.3. 
柱仕上げ（ＲＣ部材）配置・部材：StbCalColumnFinish_RC_MemList ........................... 130 
4.5.4. 
柱仕上げ（ＲＣ部材）配置・部材(計算用・分割)：StbCalColumnFinish_RC_CalMemList
 
131 
4.5.5. 
柱仕上げ（Ｓ部材）配置：StbCalColumnFinish_S_Arr ..................................................... 132 
4.5.6. 
柱仕上げ（Ｓ部材）配置・荷重：StbCalColumnFinish_S_LoadList ................................. 133 
4.5.7. 
柱仕上げ（Ｓ部材）配置・部材：StbCalColumnFinish_S_MemList ................................. 133 
4.5.8. 
柱仕上げ（Ｓ部材）配置・部材(計算用・分割)：StbCalColumnFinish_S_CalMemList ... 134 
4.5.9. 
柱特殊荷重配置：StbCalColumnMemberLoadArr .............................................................. 135 
4.5.10. 
柱特殊荷重配置・荷重：StbCalColumnMemberLoadList ................................................... 136 
4.5.11. 
柱特殊荷重配置・部材：StbCalColumnMemberLoadMemList .......................................... 136 
4.5.12. 
柱特殊荷重配置・部材(計算用・分割)：StbCalColumnMemberLoadCalMemList ............ 137 
4.5.13. 
大梁仕上げ（ＲＣ部材）配置：StbCalGirderFinish_RC_Arr ............................................. 138 
4.5.14. 
大梁仕上げ（ＲＣ部材）配置・荷重：StbCalGirderFinish_RC_LoadList ......................... 139 
4.5.15. 
大梁仕上げ（ＲＣ部材）配置・部材：StbCalGirderFinish_RC_MemList ......................... 139 
4.5.16. 
大梁仕上げ（ＲＣ部材）配置・部材(計算用・分割)：StbCalGirderFinish_RC_CalMemList
 
140 
4.5.17. 
大梁仕上げ（Ｓ部材）配置：StbCalGirderFinish_S_Arr .................................................... 141 
4.5.18. 
大梁仕上げ（Ｓ部材）配置・荷重：StbCalGirderFinish_S_LoadList ................................ 142 
4.5.19. 
大梁仕上げ（Ｓ部材）配置・部材：StbCalGirderFinish_S_MemList ................................ 142 
4.5.20. 
大梁仕上げ（Ｓ部材）配置・部材(計算用・分割)：StbCalGirderFinish_S_CalMemList .. 143 
4.5.21. 
大梁特殊荷重配置：StbCalGirderMemberLoadArr ............................................................. 144 
4.5.22. 
大梁特殊荷重配置・荷重：StbCalGirderMemberLoadList ................................................. 145 
4.5.23. 
大梁特殊荷重配置・部材：StbCalGirderMemberLoadMemList ......................................... 145 
4.5.24. 
大梁特殊荷重配置・部材(計算用・分割)：StbCalGirderMemberLoadCalMemList ........... 146 
4.5.25. 
小梁仕上げ（ＲＣ部材）配置：StbCalBeamFinish_RC_Arr ............................................... 147 
4.5.26. 
小梁仕上げ（ＲＣ部材）配置・荷重：StbCalBeamFinish_RC_LoadList ........................... 148 
4.5.27. 
小梁仕上げ（ＲＣ部材）配置・部材：StbCalBeamFinish_RC_MemList ........................... 148 
4.5.28. 
小梁仕上げ（Ｓ部材）配置：StbCalBeamFinish_S_Arr ..................................................... 149 
4.5.29. 
小梁仕上げ（Ｓ部材）配置・荷重：StbCalBeamFinish_S_LoadList ................................. 150 
4.5.30. 
小梁仕上げ（Ｓ部材）配置・部材：StbCalBeamFinish_S_MemList ................................. 150

## Page 6

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
vi 
 
4.5.31. 
小梁特殊荷重配置：StbCalBeamMemberLoadArr .............................................................. 151 
4.5.32. 
小梁特殊荷重配置・荷重：StbCalBeamMemberLoadList ................................................... 152 
4.5.33. 
小梁特殊荷重配置・部材：StbCalBeamMemberLoadMemList .......................................... 152 
4.5.34. 
ブレース仕上げ配置：StbCalBraceFinish_S_Arr ................................................................ 153 
4.5.35. 
ブレース仕上げ配置・荷重：StbCalBraceFinish_S_LoadList ............................................ 154 
4.5.36. 
ブレース仕上げ配置・部材：StbCalBraceFinish_S_MemList ............................................ 154 
4.5.37. 
床積載荷重配置：StbCalSlabLiveLoadArr ........................................................................... 155 
4.5.38. 
床積載荷重配置・荷重：StbCalSlabLiveLoadList ............................................................... 156 
4.5.39. 
床積載荷重配置・部材：StbCalSlabLiveLoadMemList ....................................................... 156 
4.5.40. 
床スラブ仕上げ配置：StbCalSlabFinish_RC_Arr ............................................................... 157 
4.5.41. 
床スラブ仕上げ配置・荷重：StbCalSlabFinish_RC_LoadList............................................ 158 
4.5.42. 
床スラブ仕上げ配置・部材：StbCalSlabFinish_RC_MemList ........................................... 158 
4.5.43. 
床面特殊荷重配置：StbCalSlabAreaLoadArr ...................................................................... 159 
4.5.44. 
床面特殊荷重配置・荷重：StbCalSlabAreaLoadList ........................................................... 160 
4.5.45. 
床面特殊荷重配置・部材：StbCalSlabAreaLoadMemList .................................................. 160 
4.5.46. 
床土圧・水圧荷重配置：StbCalSlabPressureLoadArr ........................................................ 161 
4.5.47. 
床土圧・水圧荷重配置・荷重：StbCalSlabPressureLoadList ............................................. 162 
4.5.48. 
床土圧・水圧荷重配置・部材：StbCalSlabPressureLoadMemList .................................... 162 
4.5.49. 
壁仕上げ配置：StbCalWallFinish_RC_Arr .......................................................................... 163 
4.5.50. 
壁仕上げ配置・荷重：StbCalWallFinish_RC_LoadList ...................................................... 164 
4.5.51. 
壁仕上げ配置・部材：StbCalWallFinish_RC_MemList ...................................................... 164 
4.5.52. 
壁面特殊荷重配置：StbCalWallAreaLoadArr ...................................................................... 165 
4.5.53. 
壁面特殊荷重配置・荷重：StbCalWallAreaLoadList .......................................................... 166 
4.5.54. 
壁面特殊荷重配置・部材：StbCalWallAreaLoadMemList .................................................. 166 
4.5.55. 
壁土圧・水圧荷重配置：StbCalWallPressureLoadArr ........................................................ 167 
4.5.56. 
壁土圧・水圧荷重配置・荷重：StbCalWallPressureLoadList ............................................ 168 
4.5.57. 
壁土圧・水圧荷重配置・部材：StbCalWallPressureLoadMemList .................................... 168 
4.5.58. 
節点追加・補正重量配置：StbCalNodeWeightArr ............................................................... 169 
4.5.59. 
節点追加・補正重量配置・荷重：StbCalNodeWeightLoadList ........................................... 170 
4.5.60. 
節点追加・補正重量配置・節点：StbCalNodeWeightNodeList .......................................... 170 
4.5.61. 
節点特殊荷重配置：StbCalNodePointLoadArr .................................................................... 171 
4.5.62. 
節点特殊荷重配置・荷重：StbCalNodePointLoadList ......................................................... 172 
4.5.63. 
節点特殊荷重配置・節点：StbCalNodePointLoadNodeList ................................................ 172 
4.6. 
計算条件配置：StbCalConditionArrangements .......................................................................... 173 
4.6.1. 
柱材端条件配置：StbCalColumnConditionArr .................................................................... 174 
4.6.2. 
柱材端条件配置・条件：StbCalColumnConditionList......................................................... 175

## Page 7

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
vii 
 
4.6.3. 
柱材端条件配置・部材：StbCalColumnConditionMemList ................................................ 175 
4.6.4. 
柱材端条件配置・部材(計算用・分割)：StbCalColumnConditionCalMemList .................. 176 
4.6.5. 
柱剛域長さ配置：StbCalColumnRigidzoneArr .................................................................... 177 
4.6.6. 
柱剛域長さ配置・条件：StbCalColumnRigidzoneList ........................................................ 178 
4.6.7. 
柱剛域長さ配置・部材：StbCalColumnRigidzoneMemList ................................................ 178 
4.6.8. 
柱剛域長さ配置・部材(計算用・分割)：StbCalColumnRigidzoneCalMemList .................. 179 
4.6.9. 
柱危険断面位置配置：StbCalColumnCriticalPositionArr ................................................... 180 
4.6.10. 
柱危険断面位置配置・条件：StbCalColumnCriticalPositionList ....................................... 181 
4.6.11. 
柱危険断面位置配置・部材：StbCalColumnCriticalPositionMemList ............................... 181 
4.6.12. 
柱危険断面位置配置・部材(計算用・分割)：StbCalColumnCriticalPositionCalMemList
 
182 
4.6.13. 
柱剛性倍率配置：StbCalColumnStiffnessArr ...................................................................... 183 
4.6.14. 
柱剛性倍率配置・条件：StbCalColumnStiffnessList .......................................................... 184 
4.6.15. 
柱剛性倍率配置・部材：StbCalColumnStiffnessMemList .................................................. 184 
4.6.16. 
柱剛性倍率配置・部材(計算用・分割)：StbCalColumnStiffnessCalMemList .................... 185 
4.6.17. 
大梁材端条件配置：StbCalGirderConditionArr ................................................................... 186 
4.6.18. 
大梁材端条件配置・条件：StbCalGirderConditionList ....................................................... 187 
4.6.19. 
大梁材端条件配置・部材：StbCalGirderConditionMemList ............................................... 187 
4.6.20. 
大梁材端条件配置・部材(計算用・分割)：StbCalGirderConditionCalMemList ................. 188 
4.6.21. 
大梁剛域長さ配置：StbCalGirderRigidzoneArr ................................................................... 189 
4.6.22. 
大梁剛域長さ配置・条件：StbCalGirderRigidzoneList ....................................................... 190 
4.6.23. 
大梁剛域長さ配置・部材：StbCalGirderRigidzoneMemList............................................... 190 
4.6.24. 
大梁剛域長さ配置・部材(計算用・分割)：StbCalGirderRigidzoneCalMemList ................. 191 
4.6.25. 
大梁危険断面位置配置：StbCalGirderCriticalPositionArr .................................................. 192 
4.6.26. 
大梁危険断面位置配置・条件：StbCalGirderCriticalPositionList ...................................... 193 
4.6.27. 
大梁危険断面位置配置・部材：StbCalGirderCriticalPositionMemList ............................. 193 
4.6.28. 
大梁危険断面位置配置・部材(計算用・分割)：StbCalGirderCriticalPositionCalMemList
 
194 
4.6.29. 
大梁剛性倍率配置：StbCalGirderStiffnessArr ..................................................................... 195 
4.6.30. 
大梁剛性倍率配置・条件：StbCalGirderStiffnessList ......................................................... 196 
4.6.31. 
大梁剛性倍率配置・部材：StbCalGirderStiffnessMemList ................................................ 196 
4.6.32. 
大梁剛性倍率配置・部材(計算用・分割)：StbCalGirderStiffnessCalMemList................... 197 
4.6.33. 
ブレース剛性倍率配置：StbCalBraceStiffnessArr ............................................................... 198 
4.6.34. 
ブレース剛性倍率配置・条件：StbCalBraceStiffnessList ................................................... 199 
4.6.35. 
ブレース剛性倍率配置・部材：StbCalBraceStiffnessMemList ........................................... 199 
4.6.36. 
壁剛性倍率配置：StbCalWallStiffnessArr ........................................................................... 200

## Page 8

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
viii 
 
4.6.37. 
壁剛性倍率配置・条件：StbCalWallStiffnessList ................................................................ 201 
4.6.38. 
壁剛性倍率配置・部材：StbCalWallStiffnessMemList ....................................................... 201 
4.6.39. 
節点拘束条件配置：StbCalNodeRestrictionArr ................................................................... 202 
4.6.40. 
節点拘束条件配置・条件：StbCalNodeRestrictionList ....................................................... 203 
4.6.41. 
節点拘束条件配置・節点：StbCalNodeRestrictionNodeList ............................................... 203 
4.6.42. 
節点接合部パネル配置：StbCalNodePanelArr ..................................................................... 204 
4.6.43. 
節点接合部パネル配置・条件：StbCalNodePanelList ......................................................... 205 
4.6.44. 
節点接合部パネル配置・節点：StbCalNodePanelNodeList ................................................ 205 
4.6.45. 
柱断面性能配置：StbCalColumnSecPropertyArr ................................................................ 206 
4.6.46. 
柱断面性能配置・条件：StbCalColumnSecPropertyList ..................................................... 207 
4.6.47. 
柱断面性能配置・ＲＣ断面：StbCalColumnSecProperty_RC_List .................................... 207 
4.6.48. 
柱断面性能配置・Ｓ断面：StbCalColumnSecProperty_S_List ........................................... 208 
4.6.49. 
柱断面性能配置・ＳＲＣ断面：StbCalColumnSecProperty_SRC_List ............................... 208 
4.6.50. 
柱断面性能配置・ＣＦＴ断面：StbCalColumnSecProperty_CFT_List ............................... 209 
4.6.51. 
大梁断面性能配置：StbCalGirderSecPropertyArr ............................................................... 210 
4.6.52. 
大梁断面性能配置・条件：StbCalGirderSecPropertyList ................................................... 211 
4.6.53. 
大梁断面性能配置・ＲＣ断面：StbCalGirderSecProperty_RC_List ................................... 211 
4.6.54. 
大梁断面性能配置・Ｓ断面：StbCalGirderSecProperty_S_List .......................................... 212 
4.6.55. 
大梁断面性能配置・ＳＲＣ断面：StbCalGirderSecProperty_SRC_List ............................. 212 
4.6.56. 
ブレース断面性能配置：StbCalBraceSecPropertyArr ......................................................... 213 
4.6.57. 
ブレース断面性能配置・条件：StbCalBraceSecPropertyList ............................................. 214 
4.6.58. 
ブレース断面性能配置・Ｓ断面：StbCalBraceSecProperty_S_List .................................... 214 
4.6.59. 
スラブ断面性能配置：StbCalSlabSecPropertyArr ............................................................... 215 
4.6.60. 
スラブ断面性能配置・条件：StbCalSlabSecPropertyList ................................................... 216 
4.6.61. 
スラブ断面性能配置・ＲＣ断面：StbCalSlabSecProperty_RC_List ................................... 216 
4.6.62. 
スラブ断面性能配置・デッキプレートスラブ断面：StbCalSlabSecPropertyDeckList ...... 217 
4.6.63. 
スラブ断面性能配置・既製スラブ断面：StbCalSlabSecPropertyPrecastList .................... 217 
4.6.64. 
壁断面性能配置：StbCalWallSecPropertyArr ...................................................................... 218 
4.6.65. 
壁断面性能配置・条件：StbCalWallSecPropertyList .......................................................... 219 
4.6.66. 
壁断面性能配置・ＲＣ断面：StbCalWallSecProperty_RC_List .......................................... 219 
5. 
StbAnaModels 要素リファレンス ......................................................................................................... 220 
5.1. 
解析モデル（複数）：StbAnaModels ............................................................................................. 220 
5.2. 
解析モデル：StbAnaModel ........................................................................................................... 221 
5.2.1. 
解析用節点（複数）：StbAnaNodes ..................................................................................... 222 
5.2.2. 
解析用節点：StbAnaNode ..................................................................................................... 223 
5.2.3. 
境界条件：StbAnaBoundary ................................................................................................. 224

## Page 9

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
ix 
 
5.2.4. 
解析用階（複数）：StbAnaStories ....................................................................................... 225 
5.2.5. 
解析用階：StbAnaStory ........................................................................................................ 226 
5.2.6. 
解析用部材ID リスト：StbAnaMemberid_List ................................................................... 227 
5.2.7. 
解析用部材ID：StbAnaMemberid ........................................................................................ 228 
5.2.8. 
解析用部材：StbAnaMembers .............................................................................................. 229 
5.2.9. 
梁要素（複数）：StbAnaBeams ........................................................................................... 230 
5.2.10. 
梁要素：StbAnaBeam ........................................................................................................... 231 
5.2.11. 
梁要素剛域長さ：StbAnaBeamRigidzone ............................................................................ 232 
5.2.12. 
梁要素危険断面位置：StbAnaBeamCriticalPosition ........................................................... 233 
5.2.13. 
梁要素材端ばね：StbAnaBeamEndSpring ........................................................................... 234 
5.2.14. 
トラス要素（複数）：StbAnaTrusses .................................................................................. 235 
5.2.15. 
トラス要素：StbAnaTruss .................................................................................................... 236 
5.2.16. 
支点ばね要素（複数）：StbAnaSupports ............................................................................. 237 
5.2.17. 
支点ばね要素：StbAnaSupport ............................................................................................. 238 
5.2.18. 
ばね要素（複数）：StbAnaSprings ...................................................................................... 239 
5.2.19. 
ばね要素：StbAnaSpring ...................................................................................................... 240 
5.2.20. 
壁エレメント要素（複数）：StbAnaWalls ........................................................................... 241 
5.2.21. 
壁エレメント要素：StbAnaWall ........................................................................................... 242 
5.2.22. 
平面要素（三角形、複数）： StbAnaPlaneTriangles ......................................................... 243 
5.2.23. 
平面要素（三角形）：StbAnaPlaneTriangle ....................................................................... 244 
5.2.24. 
平面要素（四角形、複数）：StbAnaPlaneRectangles ......................................................... 245 
5.2.25. 
平面要素（四角形）：StbAnaPlaneRectangle ..................................................................... 246 
5.2.26. 
接合部パネル要素（複数）：StbAnaNodePanels ................................................................. 247 
5.2.27. 
接合部パネル要素：StbAnaNodePanel ................................................................................. 248 
5.2.28. 
解析用部材性能：StbAnaProperties ..................................................................................... 249 
5.2.29. 
梁要素性能（複数）：StbAnaBeamProperties .................................................................... 250 
5.2.30. 
梁要素性能：StbAnaBeamProperty...................................................................................... 251 
5.2.31. 
トラス要素性能（複数）：StbAnaTrussProperties ............................................................. 252 
5.2.32. 
トラス要素性能：StbAnaTrussProperty .............................................................................. 253 
5.2.33. 
ばね要素性能（複数）：StbAnaSpringProperties ............................................................... 254 
5.2.34. 
ばね要素性能：StbAnaSpringProperty ................................................................................ 255 
5.2.35. 
壁エレメント要素性能（複数）：StbAnaWallProperties .................................................... 256 
5.2.36. 
壁エレメント要素性能：StbAnaWallProperty ..................................................................... 257 
5.2.37. 
平面要素性能（複数）：StbAnaPlaneProperties ................................................................. 258 
5.2.38. 
平面要素性能：StbAnaPlaneProperty .................................................................................. 259 
5.2.39. 
接合部パネル要素性能（複数）：StbAnaNodePanelProperties .......................................... 260

## Page 10

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
x 
 
5.2.40. 
接合部パネル要素性能：StbAnaNodePanelProperty ........................................................... 261 
5.2.41. 
剛床指定（複数）：StbAnaFloorDiaphragms...................................................................... 263 
5.2.42. 
剛床指定：StbAnaFloorDiaphragm ...................................................................................... 264 
5.2.43. 
解析用節点ID リスト：StbAnaNodeid_List ......................................................................... 265 
5.2.44. 
解析用節点ID：StbAnaNodeid ............................................................................................. 265 
5.2.45. 
材料性能（複数）：StbAnaMaterials ................................................................................... 266 
5.2.46. 
材料性能：StbAnaMaterial ................................................................................................... 267 
5.2.47. 
断面性能（複数）：StbAnaSections ..................................................................................... 268 
5.2.48. 
断面性能：StbAnaSection ..................................................................................................... 269 
5.2.49. 
荷重ケース（複数）：StbAnaLoadCases ............................................................................. 270 
5.2.50. 
荷重ケース：StbAnaLoadCase ............................................................................................. 270 
5.2.51. 
節点荷重：StbAnaLoadNode................................................................................................. 271 
5.2.52. 
梁荷重：StbAnaLoadBeam ................................................................................................... 272 
5.2.53. 
梁CMQ 荷重：StbAnaLoadBeamCMQ ................................................................................ 273 
5.2.54. 
梁要素初期応力：StbAnaLoadBeamInitialStress ................................................................ 274 
5.2.55. 
トラス荷重：StbAnaLoadTruss ............................................................................................ 275 
5.2.56. 
トラス要素初期応力：StbAnaLoadTrussInitialStress ........................................................ 275 
5.2.57. 
壁荷重：StbAnaLoadWall ..................................................................................................... 276 
5.2.58. 
壁エレメント要素初期応力：StbAnaLoadWallInitialStress ............................................... 276 
5.2.59. 
ばね荷重：StbAnaLoadSpring .............................................................................................. 277 
5.2.60. 
ばね要素初期応力：StbAnaLoadSpringInitialStress .......................................................... 277 
5.2.61. 
解析ケース（複数）：StbAnaAnalyses ................................................................................ 278 
5.2.62. 
静的線形解析ケース：StbAnaAnalysisStaticLinear ............................................................ 278 
5.3. 
解析モデル関連定義：StbAnaRelations ....................................................................................... 280 
5.3.1. 
解析用節点関連定義：StbAnaNodeRel ................................................................................. 281 
5.3.2. 
解析用階関連定義：StbAnaStoryRel .................................................................................... 282 
5.3.3. 
解析用部材関連定義：StbAnaMemberRel ............................................................................ 283 
5.3.4. 
解析用部材関連定義(計算用・分割)：StbAnaCalMemberRel .............................................. 284 
5.3.5. 
解析用部材性能関連定義：StbAnaPropertyRel .................................................................... 285 
6. 
修正履歴 ................................................................................................................................................. 287

## Page 11

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
1 
 
1. はじめに 
 
「ST-Bridge」は、国内の建築構造分野のソフトウェア間のデータ交換、情報共有に利用することを目指し、
一般社団法人IAI 日本 構造分科会が開発している標準データ交換形式である。 
本書は、そのST-Bridge を構造計算や構造解析で使用される計算のために必要なデータに関する拡張として
開発されたものを「計算データ編」としてまとめた仕様書である。 
 
ソフトウェアベンダー各社の協力により、国内の主要な一貫構造計算ソフトウェアや構造解析において利用
できるものを目指し、幾度も議論を重ね、まとめることができた。 
ただ、本書のカバーしている適用範囲はまだ十分なものではないだろう。これを出発点に適用範囲を広げ、
さらに有用なものにしていくことを提案したい。 
 
ご協力いただいた構造分科会各社、特にSTB 計算WG のメンバーの方々に心から感謝をしたい。 
 
2015 年3 月31 日 
一般社団法人IAI 日本 構造分科会 
STB 計算WG リーダー 千葉 貴史 
（リーダー任期：2013 年4 月1 日～2015 年3 月31 日） 
 
STB 計算WG 委員（敬称略・会社五十音順） 
アークデータ研究所 
木股信男 
吉沢俊正 
安藤ハザマ 
梅村美孝 
NTT ファシリティーズ総合研究所 
荒川裕信 
奥村幸司 
平賀章 
平尾卓也 
建築ピボット 
千葉貴史 
構造計画研究所 
安藤靖人 
宇佐美祐人 
鈴木壮 
瀬戸隆之 
藤田真可 
構造システム 
松本憲英 
構造ソフト 
高 洪 
コミュニケーションシステム 
鈴木剛 
堀岡昇 
清水建設 
大津聡 
ソフトウェアセンター 
阿部潔 
竹中工務店 
鹿島孝 
マイダスアイティジャパン 
藤田啓 
ユニオンシステム 
奥平裕信 
酒井竜志

## Page 12

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
2 
 
執筆者リスト（敬称略・五十音順） 
第1 章～第3 章 
千葉 貴史（建築ピボット） 
 
第4 章 
奥平 裕信（ユニオンシステム） 
 
木股 信男（アークデータ研究所） 
 
千葉 貴史（建築ピボット） 
 
平尾 卓也（NTT ファシリティーズ総合研究所） 
 
松本 憲英（構造システム） 
 
第5 章 
鈴木 壮（構造計画研究所） 
 
第6 章 
高  洪（構造ソフト） 
 
堀岡 昇（コミュニケーションシステム） 
 
（2015 年3 月末時点の所属）

## Page 13

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
3 
 
1.1. 
2018 年6 月改訂にあたり 
 
 2015 年4 月27 日のST-Bridge XML ファイル仕様書 計算データ編初版 (Ver.1.3) 発行後より、IAI 日本
構造分科会STB 計算WG にて改訂作業に着手した。 
 改訂の目的は、初版仕様によるデータ交換を実現するにあたり、不足する事項や、不統一・不都合と思わ
れる点を改善することである。 
 主な改訂作業は、以下のとおりである。 
 
・属性名、単語名の不統一修正 
 
・ST-Bridge 本体との整合性の確認および修正 
 
・XML 形式、表記法、命名規則への準拠確認、修正 
 
・データ型、単位の確認、修正 
 
・サンプルコード、例、図の充実 
 適用範囲の拡大は対象外としている。 
 また、本仕様の利用者は主に日本国内の構造計算プログラム開発者と想定した。 
 
一般社団法人IAI 日本 構造分科会 
STB 計算WG リーダー 宇佐美 祐人 
（リーダー任期：2015 年4 月1 日～2020 年12 月31 日） 
 
 
STB 計算WG メンバー・オブザーバー（2015 年4 月～2016 年3 月、敬称略・会社五十音順） 
アークデータ研究所 
木股信男 
安藤ハザマ 
梅村美孝 
NTT ファシリティーズ総合研究所 
荒川延夫 
辻井一晃 
建築ピボット 
千葉貴史 
構造計画研究所 
宇佐美祐人 
鈴木壮 
構造システム 
松本憲英 
構造ソフト 
高 洪 
コミュニケーションシステム 
堀岡昇 
清水建設 
大津聡 
竹中工務店 
鹿島孝 
マイダスアイティジャパン 
加藤元樹  金炅奐 
ユニオンシステム 
奥平裕信

## Page 14

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
4 
 
改訂担当（敬称略・五十音順） 
第1 章～第3 章 
宇佐美 祐人（構造計画研究所） 
 
第4 章 
奥平 裕信（ユニオンシステム） 
 
木股 信男（アークデータ研究所） 
 
鈴木 壮（構造計画研究所） 
 
堀岡 昇（コミュニケーションシステム） 
 
第5 章 
松本 憲英（構造システム） 
 
高  洪（構造ソフト） 
 
第6 章 
辻井 一晃（NTT ファシリティーズ総合研究所） 
 
第7 章～第8 章 
宇佐美 祐人（構造計画研究所） 
 
（2016 年3 月末時点の所属）

## Page 15

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
5 
 
1.2. 
2021 年2 月改訂にあたり 
 
 2018 年6 月に、「ST-Bridge XML ファイル仕様書」（以下「本編」）のVer.1.4 発行に合わせた改訂を行っ
たが、その後の本編Ver.2.0 以降の改訂に合わせて引き続き改訂作業を行い、本編Ver.2.0.2 の発行に合わせ
て発行することとした。 
 改訂にあたっては、本編Ver.2.0 の改訂主旨に沿って、id の一意性明確化、子要素の持ち方統一を図って
要素を再構成し、また、要素にguid（オプション）の追加を行った。 
荷重や計算条件について本編の要素を補足するStbCalData については、国内の主要な一貫構造計算プロ
グラムの荷重および計算条件を改めて洗い出し、特に複数のプログラムで保持している条件を極力反映でき
るよう、要素の見直しを行った。 
骨組構造解析に即したStbAnaModels については、構造解析モデルとしての独立性を保ちながら、本編の
StbModel 各要素が持つ条件を参照し易くするため、要素同士を関連付けるStbAnaRelations の構成を見直
した。 
 
一般社団法人buildingSMART Japan 構造設計小委員会 
STB 計算WG リーダー 宇佐美 祐人 
（リーダー任期：2015 年4 月1 日～2020 年12 月31 日） 
STB 計算WG リーダー 武居 秀樹 
（リーダー任期：2021 年1 月1 日～） 
 
 
STB 計算WG 改訂作業メンバー（2020 年12 月時点、敬称略・会社五十音順） 
 
アークデータ研究所 
木股信男 
構造計画研究所 
宇佐美祐人 
鈴木壮 
構造システム 
松本憲英 
中嶋明日翔 
構造ソフト 
高 洪 
コミュニケーションシステム 
堀岡昇 
清水建設 
大津聡 
竹中工務店 
鹿島孝 
日本設計 
武居秀樹

## Page 16

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
6 
 
2. 概要 
 
2.1. 
コンセプト 
ST-Bridge は、IFC のみでは表現が難しい日本国内の建築構造設計情報について、建築構造分野のソフト
ウェア間における橋渡しの実現を目指す標準フォーマットである。おもに、日本国内の一貫構造計算プログ
ラムと汎用の応力解析プログラム、構造図作成プログラムとの連携に重点を置きながら、３次元オブジェク
トCAD や積算プログラムなどと、構造躯体に関する情報を連携することも想定している。 
本仕様書は、「ST-Bridge XML ファイル仕様書」（以下、「ST-Bridge 本編」と表記）の躯体情報定義に付
随する、荷重、設計条件および解析モデルに関する事項を記したものである。 
「計算データ編」では、ST-Bridge 本編で定義可能な適用範囲を元に、現在建設されている建物のうち、
ある程度整形な形式を念頭に置き、あまりに複雑な形状や構造形式については適用範囲外としている。 
そのため、解析においても当面は弾性範囲の計算を主な対象とし、弾塑性解析などについては検討項目とし
て今後の開発にゆだねることとした。 
 
2.2. 
XML 形式 
ST-Bridge はXML 形式を採用している。 
XML (Extensible Markup Language) は、データ交換に使用可能なマークアップ言語を新たに作成するため
の基礎として使用できる、簡単で柔軟なテキスト形式の言語である。 XML は W3C (World Wide Web 
Consortium) のワーキンググループから発行された一連の勧告に基づいており、ST-Bridge もそれに倣う。 
XML 形式を表す用語として、「要素」（Element）、「属性」（Attribute）、「内容」（Content）がある。 
「要素」「属性」「内容」の、実際のタグとの対応関係は、以下となる。 
<Element Attribute=“属性値”>Content</Element> 
「内容」がない場合は、以下としてよい。 
<Element Attribute=“属性値”/> 
 
内容部分に別の要素を表記し、要素を階層構造とすることができる。 
要素名は、大文字と小文字が区別され、文字かアンダースコア（_）で始まる必要がある。また、要素名に
は、文字、数字、ハイフン、アンダースコア、およびピリオドを含めることができる。 
属性値は、  ”  ”  または  ‘ ’  で囲まれた文字列とする。従って、属性値の文字列に「”」または「’」を含
んではならない。

## Page 17

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
7 
 
XML 形式に関する基本的なルールと、ST-Bridge における扱いを以下に列記する。 
 
・XML バージョン番号 
<?xml version="1.0"?> 
 これは必須。 将来の XML バージョンでは番号が変わることがあるが、現在のバージョンは 1.0。 
 
・encoding 宣言 
<?xml version="1.0" encoding="UTF-8"?> 
 この属性は省略可能であるが、ST-Bridge では省略しないこととする。 使用する場合は、encoding 宣言
は XML 宣言中でバージョン情報の直後になければならず、既存の文字エンコードを示す値を含んでいる必
要がある。当面、"UTF-8"、 "Shift_JIS" を適用対象とする。 
 
・XML コメント 
<!--     --> 
 ドキュメントの構造や注釈など、XML パーサーに対する内容でないものは、コメント内に含める。 
 コメントは <!-- で始まり --> で終わる。 
 
・空白（スペース） 
 W3C (World Wide Web Consortium) XML 仕様では、・属性値の中を除き、すべての空白を維持する。 
 したがって、開発者はXML パーサーが空白をどのように処理するかを意識する必要がある。 
 
2.3. 
表記法 
要素の表記法は、キャメルケース（アッパーキャメルケース）を採用する。 
 
例 
<StbCalData> 
   <StbCalElement> 
   </StbCalElement> 
</StbCalData> 
 
属性の表記法は、原則スネーク記法を採用し（「_（アンダースコア）」でつなぐ）、小文字とする。 
 
例 
<StbCalData> 
   <StbCalElement attribite_data1="a" attribute_data2="2.0"> 
   </StbCalElement> 
</StbCalData>

## Page 18

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
8 
 
2.4. 
命名規則 
・ST-Bridge 計算データ編に関する「要素」名は原則として、 
 構造計算条件要素は「StbCal」 
 構造解析モデル要素は「StbAna」 
 で始める。 
・属性名は原則小文字とする 
・「内容」が選択型の文字データの場合は、全て大文字で記述する。 
 
接頭語、接尾語の命名規則 
・断面タイプは大文字とする（BOX, H 等） 
・構造形式は大文字とする（RC, SRC, S 等） 
・位置を表す場合は小文字とする（left, right, top, bottom） 
・方向を表す場合は大文字とする（X, Y, Z, A） 
・荷重を示す記号は大文字とする（C,M,Q,T,N,P1～Pn） 
 
2.5. 
属性値の型と表現範囲 
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
コンピュータ言語の「ブーリアン型」で、’true’または’false’とする 
 
属性値は、特記がない限りnull 値、スペースのみの表記および型に合わない値による表記は認めない。

## Page 19

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
9 
 
2.6. 
単位系と数値範囲 
長さ、力等に関する単位は特記がない限り以下とする。 
・単位系 
座標値 
mm 
○長さ、厚さ 
mm 
○断面積 
mm2 
○断面2 次モーメント mm4 
○比重 
N/mm3 
線荷重 
N/mm 
面荷重 
N/mm2 
力 
N 
モーメント 
Nmm 
角度 
度（degree） 
比率（無次元化数） 小数 
回転ばね定数 
Nmm/rad 
並進ばね定数 
N/mm 
○ヤング係数、せん断弾性係数 N/mm2 
○固有周期 
sec 
○速度 
mm/sec 
 
 ○印の値は、0 より大きい値とする。 
角度は、0～360 度の範囲の値とする。

## Page 20

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
10 
 
2.7. 
グローバル一意識別子（GUID） 
 GUID (Globally Unique IDentifier) は、UUID (Universally Unique IDentifier) としても知られる、128
ビットの符号なし整数で、空間および時間において一意である識別子である。UUID の仕様は RFC (Request 
for Comments) 4122 に規定されている。表記は、ifc における表記法にならい、32 桁の16 進数値を文字列
表現した値とする（16 進数値の 'a' から 'f' は、小文字とする）。 
例 
<Element  guid="78fd87737db64372bf0e7ede42393577"/> 
 
2.8. 
集合型（monolist） 
順番を区別する必要のある値集合については、順番をスペース区切りで続けて記述する。読み取り側アプ
リケーションが要素をスペースごとに切り分け、順に配列に格納することになる。 
特記がない限り、ST-Bridge においては、この表現は「内容」においてのみ利用することとし、内容のnull
値およびスペースのみの表記は認めない。 
例 
<monolist>100  101  102  103</monolist> 
 
（補足） 
元来、XML 形式は冗長さを許容するものであるが、同一の性質を表す集合値についてはスペースで区切る
ことでコンパクト化を図るケースもよく見られる。ここでは、1 対多の対応関係を示すような場合での利用
を想定している。 
地図情報で利用されるKML（Keyhole Markup Language）形式の<coordinates>要素なども同じような
考え方を採用しているので、参考になる。 
 
 
<coordinates> 
 
-122.365662,37.826988,0 
 
-122.365202,37.826302,0 
 
-122.364581,37.82655,0 
 
-122.365038,37.827237,0 
 
-122.365662,37.826988,0 
 
</coordinates> 
（参考） 
 
https://developers.google.com/kml/documentation/kmlreference?hl=ja#coordinates

## Page 21

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
11 
 
2.9. 
ファイルの拡張子 
ST-Bridge データ、及びST-Bridge 計算データのXML ファイルの拡張子は、「.stb」とする。 
 
2.10. バージョン番号 
バージョン番号は、「.」で区切り、３つの数字に意味を持たせる。 
 
Ver. X. Y. Z 
X : 大項目（情報アーキテクチャなど大きな項目が変更になった場合） 
Y : 中項目（要素や属性に変更があった場合） 
Z : 細項目（字句修正、説明追加など、リビジョンとして改訂があった場合） 
 
ST-Bridge 計算データ編においても、仕様書のバージョン表記のうち前の2 ケタ（Ver. X.Y）については、
本体のST-Bridge 本編と同じバージョン表記を用いる。 それにより、対応する設計モデルと同期しつつバー
ジョン管理を行うこととする。 
つまり、本体のST-Bridge 仕様が変更になった場合には、計算データ編に変更がない場合でも、計算デー
タ編使用を確認の上、バージョン番号を本体に合わせて変更する。 
 
また、改訂の段階に応じて、Z: 細項目をリビジョンとして規定する。ただし、通称として採用するのは、
バージョン表記のうち前の2 ケタ（Ver. X.Y）とする。ファイル名など、識別に細項目が必要な場合は、X.Y.Z
のように細項目を表記してもよい。

## Page 22

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
12 
 
2.11. 仕様書の見方 
 本仕様書は、各XML 要素ごとに以下のような書式で記述されている。 
 
 
X.X.X 
 計算条件：StbCalXXX 
 
・概要 
説明  ：計算条件 
親要素 ：StbCalData 
   
   
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
欄内に書けない記述は (1)…等、番号を
付けて後の「・補足」に記述

## Page 23

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
13 
 
<StbCalData> 
      <StbCalElem1 attr1="any_string" attr2="1"> 
        <StbCalElemChild1 （略）> 
          （略） 
        </StbCalElemChild1> 
      </StbCalElem1> 
</StbCalData> 
 
 
 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCalElemChild1 
0 
1 
条件１ 
 
StbCalElemChild2 
0 
制限無し 
条件２ 
 
 
 
 
 
※ 最小回数：1 以上である子要素は必須であり、省略できない 
 
・補足 
・・・・・・・・ 
 
・例 
 
 
 
 
 
 
・条件付きで必須となる要素がある場合は「補足」「・補足」に記述する。 
 
・個数、回数に制限がある場合、数値の範囲に制限がある場合は「補足」「・補足」に記述する。 
 
 
必要に応じて記述 
XML 要素が子要素を持つ場合はこの表を、持たない場
合には「無し」を記述 
必要に応じて記述

## Page 24

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
14 
 
 
例えば、XML 定義が 
          <StbCalElementName attribute1="aa"> content </StbCalElementName> 
 
の場合、仕様書の表記との関係は、 
 
名前 
仕様書の表記 
StbCalElementName 
「要素」 
attribute1 
「属性」 
content 
「内容」 
 
 
となり、実際のXML タグとの対応は <要素 属性> 内容 </要素>となる。 
 
「内容」がない場合は、<要素 属性 />となる。

## Page 25

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
15 
 
3. 要素リファレンス 特記事項 
 
3.1. 
全体構成 
 ST-Bridge データは、ルート要素を <ST_BRIDGE> とした階層構造であり、直下に以下の子要素を有す
る構成となっている。 
要素 
要素名 
説明 
共通情報 
StbCommon 
材料など、建物の共通情報 
位置・断面情報 
StbModel 
節点および構造躯体情報 
拡張情報 
StbExtensions 
ST-Bridge の要素に定義のない属性や子要素をアプリケ
ーションが独自に拡張する際に利用 
 
 「計算データ編」で扱うデータは、<ST_BRIDGE> 直下に以下の子要素を追加する形で構成する。 
要素 
要素名 
説明 
計算条件情報 
StbCalData 
構造計算に必要な荷重や設計条件などを、StbModel を
補足する形で定義 
解析モデル情報 
StbAnaModels 
骨組構造解析に即した、StbModel とは別の節点・部材
情報、およびStbModel との関連付けなどを定義 
 
 
 
<ST_BRIDGE> 
<StbCommon> 
・・・ 
</StbCommon> 
<StbModel> 
・・・ 
</StbModel> 
<StbExtensions> 
・・・ 
</StbExtensions> 
<StbCalData> 
・・・ 
</StbCalData> 
<StbAnaModels> 
・・・ 
</StbAnaModels> 
</ST_BRIDGE> 
共通情報 
拡張情報 
位置・断面情報 
計算条件情報 
解析モデル情報 
「計算データ編」で
扱うデータ

## Page 26

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
16 
 
3.2. 
ST-Bridge 本編の全体構成 
本編で扱う要素の階層(has-a 関係) を抜粋したものを以下に示す（ST-Bridge 本編から転記）。 
 
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
StbSecWall_RC 
StbSecFoundation_RC 
StbSecPile***** 
StbSecOpen_RC 
StbSecSteel 
StbSlabs 
StbParapets 
StbSecSlab*** 
StbSecParapet_RC 
*：RC or S or SRC (or CFT when Column) 
**：Parallel or Arc or Radial 
***：_RC or Deck or Precast 
****：same  as  parent  Element 
*****：_RC or _S or Product 
StbOpens 
StbSecFigure**** 
StbSecBarArrangement**** 
StbSecSteelFigure**** 
StbJoints 
If necessary

## Page 27

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
17 
 
3.3. 
構造設計モデルと解析モデル 
構造計算に必要なモデル化の対象には、主に構造躯体本体による構造設計モデルと、工学判断により理想
化された骨組による解析モデルがある。構造設計モデルのうち、部材の位置、断面情報および使用材料につ
いてはST-Bridge 本編のStbModel によって定義されている。 
構造計算においては、StbModel による定義のほかに、荷重や計算条件が必要となる。本仕様書「計算デー
タ編」では、このような構造計算に必要な条件を定義する「計算条件StbCalData」と、骨組による解析モ
デルを定義する「解析モデルStbAnaModels」を扱う。 
計算条件は、親要素StbCalData 以下の子要素にStbCal* と接頭語を付けて定義する。解析モデルは、親
要素StbAnaModels 以下に複数の子要素StbAnaModel を配し、StbAnaModel 以下の子要素にStbAna* と
接頭語を付けて定義する。 
ST-Bridge 本編と、計算条件および解析モデルで想定するモデル化の対象および、主要な子要素名を以下
に示す。 
 
 
構造設計モデル 
（ST-Bridge 本編） 
計算条件 
（StbCalData） 
解析モデル 
（StbAnaModels） 
接頭語 
Stb* 
StbCal* 
StbAna* 
モデル 
 
 
 
節点 
StbNodes 
――― 
StbAnaNodes 
階 
StbStories 
――― 
StbAnaStories 
部材情報 
StbMembers 
――― 
StbAnaMembers 
部材性能 
――― 
――― 
StbAnaProperties 
(断面性能、材料性能) 
断面形状情報 
StbSections 
――― 
――― 
荷重定義 
――― 
StbCalLoad 
StbAnaLoadCases 
(CM0Q,初期応力) 
計算条件定義 
――― 
StbCalCondition 
――― 
 
 
ratio 
A, I, J

## Page 28

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
18 
 
StbModel とStbCalData との関連 
計算条件StbCalData は、主に一貫構造計算プログラムのデータ交換を想定して定義する。 
StbCalData の各子要素は、ST-Bridge 本編のモデルに対して荷重や条件を補足する形で与えるものであり、 
・構造設計モデル全体に対する要素（地震力計算条件など） 
・ST-Bridge 本編の階と関連を持つ要素（床区分や地震力など） 
・ST-Bridge 本編の節点と関連を持つ要素（節点特殊荷重など） 
・ST-Bridge 本編の部材および断面と関連を持つ要素（材端条件、部材断面性能など） 
に、大別される。 
階と関連を持つ要素は、属性にStbStory の階ID を有し、この属性を介して階情報と関連付ける。節点と
関連を持つ要素、部材および断面と関連を持つ要素には、荷重配置StbCalLoadArrangements、計算条件配
置StbCalConditionArrangements があり、それぞれ属性にStbModel の節点、部材および断面ID を有し、
この属性を介してStbModel の各要素と関連付ける。 
 
StbModel とStbAnaModels との関連 
解析モデルStbAnaModels は、主に骨組構造解析プログラムのデータ交換を想定して定義する。 
StbAnaModels は、子要素StbAnaModel を複数個有する構成とし、一つの構造設計モデルに複数の構造解
析モデルを保持することを想定している。それぞれの構造解析モデルの中に、解析用の節点、階、部材、要
素、荷重、条件などを束ねて管理する。 
ここで、節点や部材、断面情報は、StbModel とは別に定義することが出来、設計データであるStbModel
が存在しない場合でも、構造解析モデルとして独立して定義することが可能である。解析モデルと設計モデ
ルに関連がある場合には、StbAnaModel とStbModel を関連付ける要素として、StbAnaRelations 要素を記
述する。

## Page 29

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
19 
 
例えば、ST-Bridge 本編の大梁StbGirder の大梁ID について、StbModel とStbCalData および
StbAnaModels の参照関係は、以下となる。 
 
 
 
 
解析モデル
StbAnaModels
ST-Bridge本編
StbGirder
大梁GUID
大梁ID
ST-Bridge計算編
計算条件
StbCalData
解析モデル
StbAnaModel
大梁
部材特殊荷重
StbCalMemberLoad
特殊荷重ID
大梁特殊荷重配置
配置関係を
定義
StbCalGirderMemberLoadArr
本文4.3節
特殊荷重ID
本文4.5節
StbAnaModels
本文5.2節
梁要素
StbAnaBeam
梁要素ID
参照関係を
定義
本文5.3節
解析用部材関連定義
StbAnaMemberRel
梁要素ID
大梁ID
大梁ID

## Page 30

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
20 
 
3.4. 
座標系について 
座標系の定義は、計算条件要素StbCalData と解析モデル要素StbAnaModels で異なる。StbCalData の
場合は、ST-Bridge 本編と同様であり、全体座標系と各構造部材の座標系（部材座標系）は、下記とする。 
全体座標系と部材座標系を区別する場合、全体座標系は X
－
, Y
－
, Z
－
 のように上線付で表記する。 
・全体座標系・部材座標系ともに直交座標系とする。 
・「節点」は、全体座標系で表記する。 
・構造部材の部材座標系は、各構造部材種別ごとに、節点を結ぶ配置基準線（面）に対して適用する。 
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
柱、接合部パネル、杭 
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

## Page 31

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
21 
 
StbAnaModels の場合は、下記とする。 
 
・全体座標系は原則的に以下の属性名により定義する 
 
Ｘ方向 
（X 軸回り） 
Y 方向 
（Y 軸回り） 
Z 方向 
（Z 軸回り） 
並進 
X 
Y 
Z 
回転 
TX 
TY 
TZ 
 
・引張および圧縮軸力の符号は、引張を＋、圧縮を－として定義する。 
・StbAnaModel 内でのid 参照はStbAnaModel 内で定義されたid とし、ST-Bridge の情報を参照する場
合には「StbAnaRelations」を用いる。

## Page 32

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
22 
 
3.5. 
主要な要素のID と一意性 
 本編の要素と同様に、それぞれ個々の要素を一意に特定する場合は、ID（属性名id）を定義する。 
id はinteger 型の値（１以上の整数値）とし、必須の属性とする。 
ID を定義する要素には、同時に属性 としてGUID（属性名guid）を定義する。guid は必須の属性ではな
いが、持たせる場合は、ST-Bridge データ全体で一意となる識別子の文字列表現とし、ST-Bridge を介した
データ連携においてGUID が変化しないことが望ましい。ST-Bridge においては、一意の識別子はGUID で
統一したいが、現状、GUID を扱わないプログラムが多いため、当面の措置として整数値のid を必須として
いる。 
 
StbCalData の場合 
下表の要素のID は一意とし、ID が重複してはならない。 
 
StbCalData 要素名 
StbCalLiveload 
StbCalSeismicConditionGroup 
StbCalFloorDividedArea 
 
 
StbCalColumnCondition 
StbCalColumn 
StbCalGirderCondition 
StbCalGirder 
StbCalColumnRigidzone 
 
StbCalGirderRigidzone 
StbCalColumnFinish_RC 
StbCalColumnCriticalPosition 
StbCalGirderFinish_RC 
StbCalGirderCriticalPosition 
StbCalSlabFinish_RC 
StbCalColumnSectionProperty 
StbCalWallFinish_RC 
StbCalGirderSectionProperty 
StbCalColumnFinish_S 
StbCalWallSectionProperty 
StbCalGirderFinish_S 
StbCalBraceSectionProperty 
StbCalBraceFinish_S 
StbCalSlabSectionProperty 
 
 
StbCalLoadCase 
StbCalColumnStiffness 
StbCalMemberLoad 
StbCalGirderStiffness 
StbCalAreaLoad 
StbCalWallStiffness 
StbCalPointLoad 
StbCalBraceStiffness 
StbCalSelectedPointLoad 
StbCalNodeRestriction 
StbCalEarthHydrostaticPressureLoad 
StbCalNodePanel 
StbCalNodeAddedWeight 
 
StbCalSelectedNodeAddedWeight 
StbCalFloorDiaphragm

## Page 33

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
23 
 
StbAnaModels の場合 
下表の要素のID は各々のStbAnaModel 要素内部で一意とし、ID が重複してはならない。 
 
StbAnaModel 要素名 
StbAnaNode 
StbAnaBeamProperty 
StbAnaStory 
StbAnaTrussProperty 
StbAnaBeam 
StbAnaSpringProperty 
StbAnaTruss 
StbAnaWallProperty 
StbAnaSupport 
StbAnaPlaneProperty 
StbAnaSpring 
StbAnaNodePanelProperty 
StbAnaWall 
StbAnaFloorDiaphragm 
StbAnaPlaneTriangle 
StbAnaMaterial 
StbAnaPlaneRectangle 
StbAnaSection 
StbAnaNodePanel 
StbAnaLoadCase 
 
StbAnaAnalysisStaticLinear 
 
StbAnaModels 要素の子要素StbAnaModel については、解析モデルID はStbAnaModels 要素内部で一
意とする。

## Page 34

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
24 
 
4. StbCalData 要素リファレンス 
 
4.1. 
計算条件：StbCalData 
・概要 
説明  
：計算条件 
親要素 
：ST_BRIDGE 
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
StbCalCommon 
0 
1 
計算共通条件 
※(1) 
StbCalLoad 
0 
1 
荷重定義 
※(2) 
StbCalCondition 
0 
1 
計算条件定義 
※(3) 
StbCalLoadArrangements 
0 
1 
荷重配置 
※(4) 
StbCalConditionArrangements 
0 
1 
計算条件配置 
※(5) 
・補足 
StbCalData では、主に一貫構造計算プログラムにおける計算条件に関するデータを保持、交換
するための値を定義する。 
(1) 対象建築物全般にわたる共通条件を指定する。 
(2) 節点や部材に配置する荷重の値を指定する。 
(3) 節点、断面および部材に配置する計算条件の値を指定する。 
(4) (2)で定義した値を配置する、ST-Bridge 本編の節点ID、部材ID を指定する。 
(5) (3)で定義した値を配置する、ST-Bridge 本編の節点ID、断面ID および部材ID を指定する。

## Page 35

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
25 
 
4.2. 
計算共通条件：StbCalCommon 
・概要 
説明  
：計算共通条件 
親要素 
：StbCalData 
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
StbCalLoadCondition 
1 
1 
荷重条件 
※(1) 
StbCalFloorDividedAreas 
0 
1 
床区分（複数） 
※(2) 
StbCalColumns 
0 
1 
柱(計算用・分割)（複数） ※(3) 
StbCalGirders 
0 
1 
大梁(計算用・分割)（複数） ※(4) 
・補足 
対象建築物全般にわたる共通条件を定義する（詳細は、各子要素の項を参照）。 
(1) 荷重条件のうち、個々の節点や部材に配置する荷重以外の、対象建築物全般にわたる共通条件の
値を指定する。 
(2) ST-Bridge 本編の階要素における所属節点に対し、計算上グループ分け（床区分）が必要な場合に
指定する。 
(3) ST-Bridge 本編の柱要素に対し、計算上分割が必要な場合に指定する。 
(4) ST-Bridge 本編の大梁要素に対し、計算上分割が必要な場合に指定する。

## Page 36

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
26 
 
4.2.1. 荷重条件：StbCalLoadCondition 
・概要 
説明  
：荷重条件集合 
親要素 
：StbCalCommon 
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
StbCalSeismicCondition 
0 
1 
地震荷重条件 
 
StbCalWindCondition 
0 
1 
風荷重条件 
 
StbCalSnowCondition 
0 
1 
積雪荷重条件 
 
StbCalLiveloads 
0 
1 
積載荷重（複数） 
 
・補足 
無し

## Page 37

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
27 
 
4.2.2. 地震荷重条件：StbCalSeismicCondition 
・概要 
説明  
：地震荷重の計算条件 
親要素 
：StbCalLoadCondition 
・属性 
属性名 
型 
必須 
説明 
補足 
zone 
double 
○ 
地域係数 
※(1) 
importance 
double 
 
重要度係数 
※(2) 
soil 
integer 
○ 
地盤種別 
1(第1 種地盤)、2(第2 種地盤)、3(第3 種地
盤) から選択 
※(1) 
Tc 
double 
 
地盤周期直接入力値（秒） 
※(3) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 日本の建築基準法施行令による「地域係数：Ｚ」、および同告示による「地盤種別」の値を指
定する。 
(2) 省略された場合は、1.0 とする。 
(3) Tc は地盤種別による値と異なる場合に指定するが、プログラムによっては実際に計算に用いら
れる採用値は異なることがある。 
・例 
<StbCalData> 
  <StbCalCommon> 
    <StbCalLoadCondition> 
      <StbCalSeismicCondition zone="1.0" importance="1.0" soil="2"/> 
     （略） 
    </StbCalLoadCondition> 
    （略） 
  </StbCalCommon> 
</StbCalData>

## Page 38

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
28 
 
4.2.3. 風荷重条件：StbCalWindCondition 
・概要 
説明  
：風荷重の計算条件 
親要素 
：StbCalLoadCondition 
・属性 
属性名 
型 
必須 
説明 
補足 
roughness 
integer 
○ 
地表面粗度区分 
1(Ⅰ) 2(Ⅱ) 3(Ⅲ) 4(Ⅳ) から選択 
※(1) 
wind_speed 
double 
○ 
基準風速(mm/sec) 
※(1) 
height 
double 
○ 
風荷重計算用建物高さ(mm)  
※(1) 
reduction_coefficient 
double 
 
低減率 
※(2) 
velocity_pressure 
double 
 
速度圧(N/mm2) 
※(3) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 速度圧を、日本の建築基準法施行令および同告示による、地表面粗度区分、基準風速および風
荷重計算用建物高さから計算する場合、これらの値は必須とする。ただし、風荷重計算用建物
高さは、建築物の形状から計算するプログラムでは、省略されることがある。 
(2) 省略された場合は、1.0 とする。 
(3) 速度圧が、地表面粗度区分、基準風速、風荷重計算用建物高さから計算する計算値と異なる場
合に指定する。指定する場合は、0 より大きな値とする。ただし、プログラムによっては実際に
計算に用いられる採用値は異なることがある。 
・例 
<StbCalData> 
  <StbCalCommon> 
    <StbCalLoadCondition> 
      <StbCalWindCondition roughness="2" wind_speed="38000" height="15600"/> 
     （略） 
    </StbCalLoadCondition> 
    （略） 
  </StbCalCommon> 
</StbCalData>

## Page 39

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
29 
 
4.2.4. 積雪荷重条件：StbCalSnowCondition 
・概要 
説明  
：積雪荷重の計算条件 
親要素 
：StbCalCondition 
・属性 
属性名 
型 
必須 
説明 
補足 
unit_weight 
double 
○ 
積雪単位重量(N/mm2/mm) 
 
snow_depth 
double 
○ 
垂直積雪量(mm) 
※(1) 
region 
integer 
 
区域番号 
※(2) 
altitude 
double 
 
標高(mm) 
※(2) 
sea_coverage 
double 
 
海率 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 垂直積雪量を(2)の計算とするプログラムの場合、計算に用いられる垂直積雪量の採用値はこの
値と異なることがある。 
(2) 垂直積雪量を、日本の建築基準法施行令および告示により区域番号、標高および海率から計算
する場合、これらの値は必須とする。区域番号は、告示「平12 建告第1455 号 別表」に示す
番号とする。 
・例 
<StbCalData> 
  <StbCalCommon> 
    <StbCalLoadCondition> 
      <StbCalSnowCondition unit_weight="2.0e-6" snow_depth="400"/> 
     （略） 
    </StbCalLoadCondition> 
    （略） 
  </StbCalCommon> 
</StbCalData>

## Page 40

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
30 
 
4.2.5. 積載荷重（複数）：StbCalLiveloads 
・概要 
説明  
：積載荷重（複数） 
親要素 
：StbCalCondition 
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
StbCalLiveload 
1 
制限無し 
積載荷重 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCommon> 
    <StbCalLoadCondition> 
      <StbCalLiveloads> 
  <StbCalLiveload id="1" name="居室" type="1"/> 
  <StbCalLiveload id="2" name="屋根" type="0" liveload_slab="0.001" 
        liveload_beam="0.001" liveload_frame="0.0013" liveload_seismic="0.0006"/> 
      </StbCalLiveloads> 
     （略） 
    </StbCalLoadCondition> 
    （略） 
  </StbCalCommon> 
</StbCalData>

## Page 41

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
31 
 
4.2.6. 
積載荷重：StbCalLiveload 
・概要 
説明  
：単位面積当たりの積載荷重値 
親要素 
：StbCalLiveloads 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
積載荷重ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
用途名 
 
type 
integer 
○ 
1～7：基準法施行令 0：その他 
※(1) 
liveload_slab 
double 
 
スラブ用積載荷重(N/mm2) 
 
liveload_beam 
double 
 
小梁用積載荷重(N/mm2) 
 
liveload_frame 
double 
 
ラーメン用積載荷重(N/mm2) 
 
liveload_seismic 
double 
 
地震用積載荷重(N/mm2) 
 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 1～7：日本の建築基準法施行令による標準的な積載荷重値を、下表のtype 番号で指定し、スラ
ブ用積載荷重から地震用積載荷重までの属性は指定しない。 
0：積載荷重値を、スラブ用積載荷重から地震用積載荷重までの属性で指定する。４つの属性を
全て指定し、省略してはならない。 
 
type 
用途 
積載荷重(N/mm2) 
スラブ用 
小梁用 
ラーメン用 
地震用 
１ 
居室、寝室、病室 
0.0018 
0.0018 
0.0013 
0.0006 
2 
事務室 
0.0029 
0.0029 
0.0018 
0.0008 
3 
教室 
0.0023 
0.0023 
0.0021 
0.0011 
4 
百貨店、店舗の売場 
0.0029 
0.0029 
0.0024 
0.0013 
5 
客席、集会室（固定席） 
0.0029 
0.0029 
0.0026 
0.0016 
6 
客席、集会室（その他） 
0.0035 
0.0035 
0.0032 
0.0021 
7 
車庫、自動車通路 
0.0054 
0.0054 
0.0039 
0.0020

## Page 42

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
32 
 
4.2.7. 床区分（複数）：StbCalFloorDividedAreas 
・概要 
説明  
：床区分（複数） 
親要素 
：StbCalCommon 
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
StbCalFloorDividedArea 
1 
制限無し 
床区分 
 
・補足 
構造計算プログラムが、地震による水平力に対する検討を行うとき、基本的に剛床仮定（同一層
の節点が同一の水平変位となる）が成り立つとして応力解析される場合が多い。 
一方、プログラムによっては、拡張機能として剛床仮定の条件を変更することができ、大きな吹
抜けがある場合、ツインタワーのような形状の場合などのモデル化方法として利用されている。仕
様策定にあたり、主要な各社のプログラムを調査したところ、 
・応力解析時に、特定節点の剛床を解除できるもの 
・応力解析時に、同一層に複数の剛床を考慮できるもの 
・加えて、層にまたがる剛床をグループ化して、それぞれのグループについて地震力の計算 
（および、条件の指定）ができるものがあった。 
 
本仕様では、可能な限りこれらに対応するため、剛床を構成する節点の集まりを階ごとに指定し
たものを「床区分」としてID を設けて定義し、剛床は、該当する「床区分」を指定することで行う
こととする。 
また、地震による水平力を定義する条件を複数指定できるよう、「地震力計算グループ」にID を
設けて定義し、「地震力計算グループ」定義の際に、対象とする「床区分」を指定することとする。 
このような定義方法により、基本的な計算条件の場合も簡易に定義しながら、各プログラムの拡
張機能にも対応可能とした。

## Page 43

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
33 
 
4.2.8. 
床区分：StbCalFloorDividedArea 
・概要 
説明  
：床区分（節点の集まりを階ごとに指定したもの） 
親要素 
：StbCalFloorDividedAreas 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
床区分ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
床区分名称 
 
isInclude 
boolean 
 
床区分条件 
子要素の内容に示す節点が床区分に 
true：含まれる  false：含まれない 
※(1) 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCalStoryDivided 
0 
制限無し 
床区分を定義する階 
 
・補足 
子要素にて、階ごとに節点の集まりを指定する。 
(1) 省略された場合は、true とする。 
・例 
<StbCalData> 
  <StbCalCommon> 
    <StbCalFloorDividedAreas> 
      <StbCalFloorDividedArea id="1" name="zone1" isInclude="true"> 
  <StbCalStoryDivided id_story="2" name_floor="zone1-2F">10 20 30</StbCalStoryDivided> 
  <StbCalStoryDivided id_story="3" name_floor="zone1-3F">11 21 31</StbCalStoryDivided> 
      </StbCalFloorDividedArea> 
      <StbCalFloorDividedArea id="2" name="kaijyo1" isInclude="false"> 
   <StbCalStoryDivided id_story="3" name_floor="kaijyo-3F">16</StbCalStoryDivided> 
      </StbCalFloorDividedArea> 
      <StbCalFloorDividedArea id="3" name="hi-gosyo" isInclude="true"/> 
 
      <StbCalFloorDividedArea id="4" name="gosyo" isInclude="false"/> 
    </StbCalFloorDividedAreas> 
  </StbCalCommon> 
</StbCalData> 
 
id="1"の例：含まれる節点を指定 → 部分剛床の指定を想定 
id="2"の例：含まれない節点を指定 → 剛床解除節点の指定を想定 
id="3"の例：含まれる節点がひとつもない指定 → 全層非剛床の指定を想定 
id="4"の例：含まれない節点がひとつもない指定 → 全層剛床の指定を想定

## Page 44

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
34 
 
4.2.9. 
床区分を定義する階：StbCalStoryDivided 
・概要 
説明  
：床区分を定義する階の指定と構成する節点の指定 
親要素 
：StbCalFloorDividedArea 
・属性 
属性名 
型 
必須 
説明 
補足 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
name_floor 
string 
 
床区分用階名称 
 
・内容 
内容 
型 
必須 
説明 
補足 
StbNode.id 
[monolist] 
integer 
○ 
節点ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
無し 
・補足 
(1) 該当する節点がない階の場合は、この子要素自体を省略してよい 
・例 
<StbCalData> 
  <StbCalCommon> 
    （略） 
    <StbCalFloorDividedAreas> 
      <StbCalFloorDividedArea id="1" name="zone1" isInclude="true"> 
  <StbCalStoryDivided id_story="2" name_floor="zone1-2F">10 20 30</StbCalStoryDivided> 
  <StbCalStoryDivided id_story="3" name_floor="zone1-3F">11 21 31</StbCalStoryDivided> 
      </StbCalFloorDividedArea> 
    </StbCalFloorDividedAreas> 
  </StbCalCommon> 
</StbCalData>

## Page 45

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
35 
 
4.2.10. 柱(計算用・分割)（複数）：StbCalColumns 
・概要 
説明  
：計算上の指定で分割した柱（複数） 
親要素 
：StbCalCommon 
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
StbCalColumn 
1 
制限無し 
柱(計算用・分割) 
 
・補足 
構造計算プログラムにおいては、断面計算上または作図上は１台の柱であるが、中間に大梁やブ
レースが取付く場合に、応力解析上は節点を設けて２台に分割する場合がある。StbCalColumn は、
２台以上の分割された柱に、計算編における条件を指定する必要がある場合に指定する。 
（具体例は、「大梁(計算用・分割)（複数）：StbCalGirders」の項を参照。）

## Page 46

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
36 
 
4.2.11. 柱(計算用・分割)：StbCalColumn 
・概要 
説明  
：計算上の指定で分割した柱 
親要素 
：StbCalColumns 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
柱(計算用・分割)ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
名称 
 
id_parent 
integer 
○ 
親部材 柱ID(StbColumn.id) 
 
id_node_bottom 
integer 
○ 
始端節点ID(StbNode) 
※(1) 
id_node_top 
integer 
○ 
終端節点ID(StbNode) 
※(1) 
rotate 
double 
 
回転角 
※(3),※(6) 
id_section 
integer 
○ 
断面ID 
※(5) 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値をとる。 
RC、S、SRC、CFT、UNDEFINED 
※(5) 
strength_concrete 
string 
 
コンクリート強度 
※(6) 
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
・内容 
無し 
・子要素 
無し 
・補足 
(1) 軸線が曲率を有する柱は対象外。

## Page 47

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
37 
 
(2) 柱の基準点は図心位置とし、オフセットは、節点から基準点への相対距離を記述する。省略され
た場合は、基準点が節点となる。 
(3) 回転はX 方向を0 度とし、始端から終端への進行方向時計回りを正とし、断面を回転した後にオ
フセットを考慮する。省略した時は0 度（回転なし）とみなす。 
(4) 円形の場合は「Ｘ始」のみ記述する。省略された場合、当該属性がないものとする。 
(5) 構造種別と、断面ID が参照するST-Bridge 本編の要素名との対応は、それぞれRC－
StbSecColumn_RC 、S －StbSecColumn_S 、SRC －StbSecColumn_SRC 、CFT －
StbSecColumn_CFT、UNDEFINED－StbSecUndefined とする。 
(6) 省略された場合は、親部材の同一属性による。 
 
・例 
<StbCalData> 
  <StbCalCommon> 
    （略） 
    <StbCalColumns> 
      <StbCalColumn id="101" id_parent="1" name="1CA1" 
id_node_bottom="15" id_node_top="24" id_section="3" kind_structure="RC"/> 
     （略） 
    </StbCalColumns> 
  </StbCalCommon> 
</StbCalData>

## Page 48

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
38 
 
4.2.12. 大梁(計算用・分割)（複数）：StbCalGirders 
・概要 
説明  
：計算上の指定で分割した大梁（複数） 
親要素 
：StbCalCommon 
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
StbCalGirder 
1 
制限無し 
大梁(計算用・分割) 
 
・補足 
構造計算プログラムにおいては、断面計算上または作図上は１台の大梁であるが、中間に大梁や
ブレースが取付く場合に、応力解析上は節点を設けて２台に分割する場合がある。ST-Bridge 本編
では、そのような節点がある場合の措置として、StbGirder に子要素を追加して経由する「中間節
点」のID を記述する仕様としている。 
 
 
 
計算編においては、荷重定義および計算条件定義の指定について、１台の大梁に対する場合と、
２台以上の大梁それぞれに別個に指定する場合が起こり得る。StbCalGirder は、２台以上の大梁に
指定する必要がある場合に、分割された大梁のID を独自に定義し、親となる１台との関係性を明示
するために指定する。 
 
 
 
始点
終点
中間節点
StbGirder
StbCalGirder
StbCalGirder
：StbNode
直交方向から中間節点を
介して大梁が取り付く
StbGirder
StbGirder
StbGirder

## Page 49

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
39 
 
4.2.13. 大梁(計算用・分割)：StbCalGirder 
・概要 
説明  
：計算上の指定で分割した大梁 
親要素 
：StbCalGirders 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
大梁(計算用・分割)ID 
 
guid 
string 
 
GUID 
 
name 
string 
○ 
名称 
 
id_parent 
integer 
○ 
親部材 大梁ID(StbGirder.id) 
 
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
※(3), ※(11) 
id_section 
integer 
○ 
断面ID 
※(10) 
section_io_start 
string 
 
始端側断面の外端・内端指定 
以下のいずれかの値をとる。 
OUT、IN、CENTER 
※(4) 
section_io_end 
string 
 
終端側断面の外端・内端指定 
以下のいずれかの値をとる。 
OUT、IN、CENTER 
※(4) 
kind_structure 
string 
○ 
構造種別 
以下のいずれかの値をとる。 
RC、S、SRC、UNDEFINED 
※(10) 
isFoundation 
boolean 
○ 
基礎か否か 
 
strength_concrete 
string 
 
コンクリート強度 
※(11) 
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
thickness_add_top 
double 
 
ふかし厚さ（上） 
※(5) 
thickness_add_bottom 
double 
 
ふかし厚さ（下） 
※(5) 
thickness_add_right 
double 
 
ふかし厚さ（右） 
※(5) 
thickness_add_left 
double 
 
ふかし厚さ（左） 
※(5)

## Page 50

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
40 
 
haunch_start 
double 
 
ハンチ位置（始端） 
※(6) 
haunch_end 
double 
 
ハンチ位置（終端） 
※(6) 
kind_haunch_start 
string 
 
ハンチ種類（始端） 
以下のいずれかの値をとる。 
SLOPE、DROP 
※(7) 
kind_haunch_end 
string 
 
ハンチ種類（終端） 
以下のいずれかの値をとる。 
SLOPE、DROP 
※(7) 
type_haunch_H 
string 
 
水平ハンチ形状 
以下のいずれかの値をとる。 
BOTH、RIGHT、LEFT 
※(8) 
type_haunch_V 
string 
 
鉛直ハンチ形状 
以下のいずれかの値をとる。 
BOTH、TOP、BOTTOM 
※(9) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 梁の部材座標系は、始端から終端をX、鉛直上向きがZ、の右手座標系とする。曲がり梁（軸線が
曲率を有する梁）は対象外とする。 
(2) 梁の基準点は、梁天端の幅中心位置とし、オフセットは、節点から基準点への相対距離を記述す
る。省略された場合は、基準点が節点となる。 
(3) 回転の中心は梁の天端で、回転した後でオフセット値を考慮する。省略した時は0 度（回転なし）
とみなす。 
(4) 断面の外端・内端指定は、断面の始端・終端を外端または内端（片持ち梁の場合は元端または先
端）と読み替える場合に指定する。この指定がある場合、梁断面（StbSecBeam_RC・
StbSecBeam_S・StbSecBeam_SRC）のisOutin は必ずtrue でなければならない。逆に梁断面の
isOutin がtrue でこの指定が省略された場合はアプリケーションが外端・内端を判断する。 
(5) 省略された場合、当該属性がないものとする。ふかし厚さはハンチ部分も同様の厚さとする。 
(6) 省略された場合、ハンチがない（部材の全部位で一様な形状）とする。ハンチ位置は梁始終端の
基準点からの距離とする。 
(7) 省略された場合、SLOPE（テーパーハンチ）とする。 
(8) 省略された場合、BOTH（水平両ハンチ）とする。 
(9) 省略された場合、BOTTOM（鉛直下ハンチ）とする。

## Page 51

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
41 
 
(10) 構造種別と、断面ID が参照する要素名との対応は、それぞれRC－StbSecBeam_RC、S－
StbSecBeam_S、SRC－StbSecBeam_SRC、UNDEFINED－StbSecUndefined とする。 
(11) 省略された場合は、親部材の同一属性による。 
・例 
以下の接続関係の場合について示す。 
 
 
 
 
<StbGirders> 
  <StbGirder id="101" name="2G1" id_node_start="1" id_node_end="7" id_section="1" ・・>  
<StbGirderViaNode> 
      <StbNodeIdOrder>4</StbNodeIdOrder> 
</StbGirderViaNode> 
  </StbGirder> 
  <StbGirder id="102" name="2G3" id_node_start="3" id_node_end="9" id_section="3" ・・> 
<StbGirderViaNode> 
      <StbNodeIdOrder>6</StbNodeIdOrder> 
</StbGirderViaNode> 
  </StbGirder> 
</StbGirders> 
 
<StbCalData> 
  <StbCalCommon> 
    <StbCalGirders> 
      <StbCalGirder id="1" id_parent="101" name="2G1_1" 
id_node_start="1" id_node_end="4" id_section="35" kind_structure="S" 
isFoundation="false"/> 
StbGirder id = 101
StbCalGirder
id = 1
StbNode id = 1
StbNode id = 4
StbCalGirder
id = 2
StbNode id = 7
StbGirder
id = 102
StbCalGirder
id = 3
StbNode id = 3
StbNode id = 6
StbCalGirder
id = 4
StbNode id = 9

## Page 52

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
42 
 
      <StbCalGirder id="2" id_parent="101" name="2G1_2" 
id_node_start="4" id_node_end="7" id_section="35" kind_structure="S" 
isFoundation="false"/> 
      <StbCalGirder id="3" id_parent="102" name="2G3_1" 
id_node_start="3" id_node_end="5" id_section="36" kind_structure="S" 
isFoundation="false"/> 
      <StbCalGirder id="4" id_parent="102" name="2G3_2" 
id_node_start="6" id_node_end="9" id_section="36" kind_structure="S" 
isFoundation="false"/> 
    </StbCalGirders> 
  </StbCalCommon> 
</StbCalData>

## Page 53

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
43 
 
4.3. 
荷重定義：StbCalLoad 
・概要 
説明  
：荷重定義 
親要素 
：StbCalData 
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
StbCalFinish 
0 
1 
仕上げ定義 
※(1) 
StbCalLoadCases 
0 
1 
荷重ケース（複数） 
※(2) 
StbCalAdditionalLoads 
0 
1 
特殊荷重 
※(3) 
StbCalAddedWeights 
0 
1 
追加・補正重量 
※(4) 
StbCalSeismic 
0 
1 
地震荷重 
※(5) 
・補足 
節点や部材に配置する荷重の値を指定する（詳細は、各子要素の項を参照）。 
(1) 部材に配置する仕上げ重量を指定する。 
(2) 荷重種類を分類する単位となる「荷重ケース」を定義する。 
(3) 部材および節点に直接作用する荷重を、タイプおよび数値で指定する。 
(4) 節点および任意点に直接付加する重量を、数値で指定する。 
(5) 地震荷重に関する計算条件および荷重値を指定する。

## Page 54

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
44 
 
4.3.1. 仕上げ定義：StbCalFinish 
・概要 
説明  
：部材の仕上げ定義 
親要素 
：StbCalLoad 
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
StbCalFinish_RC 
0 
1 
標準仕上げ（ＲＣ部材） 
 
StbCalFinish_S 
0 
1 
標準仕上げ（Ｓ部材） 
 
StbCalFloorFinishes 
0 
1 
階別仕上げ（複数） 
 
StbCalMemberFinishes_RC 
0 
1 
部材仕上げ（ＲＣ部材） 
 
StbCalMemberFinishes_S 
0 
1 
部材仕上げ（Ｓ部材） 
 
・補足 
無し

## Page 55

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
45 
 
4.3.2. 標準仕上げ（ＲＣ部材）：StbCalFinish_RC 
・概要 
説明  
：ＲＣ部材の標準仕上げ 
親要素 
：StbCalFinish 
・属性 
属性名 
型 
必須 
説明 
補足 
weight_girder 
double 
 
大梁仕上げ重量(N/mm2) 
※(2) 
type_girder 
integer 
 
大梁仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
※(1) 
weight_column 
double 
 
柱仕上げ重量(N/mm2) 
※(2) 
type_column 
integer 
 
柱仕上げタイプ 
1：4 面仕上げ 2：2 面仕上げ 
3：仕上げ無し 
※(1) 
weight_beam 
double 
 
小梁仕上げ重量(N/mm2) 
※(2) 
type_beam 
integer 
 
小梁仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
※(1) 
weight_canti 
double 
 
片持梁仕上げ重量(N/mm2) 
※(2) 
type_canti 
integer 
 
片持梁仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
※(1) 
weight_wall 
double 
 
壁仕上げ重量(N/mm2) 
※(2) 
type_wall 
integer 
 
仕上げタイプ 
1：両面仕上げ 2：片面仕上げ 
3：仕上げ無し 
※(1) 
weight_slab 
double 
 
スラブ仕上げ重量(N/mm2) 
※(3) 
・内容 
無し 
・子要素 
無し 
・補足 
ＲＣ部材の単位面積あたり仕上げ重量で、個別指定または階別指定のない、該当部材種別の部材
に適用される標準値がある場合、その値とタイプを指定する。 
(1) それぞれの部材種別について、標準値がある場合は必須とする。

## Page 56

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
46 
 
(2) それぞれの部材種別について、仕上タイプが「仕上げ無し」以外のとき、仕上げ重量は必須とす
る。 
(3) 仕上げの総和の値とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalFinish_RC weight_girder="0.0006" type_girder="1" 
weight_column="0.0006" type_column="1" 
weight_beam="0.0006" type_beam="1" 
weight_canti="0.0006" type_canti="1" 
weight_wall="0.0008" type_wall="1" weight_slab="0.0008"/> 
     （略） 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 57

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
47 
 
4.3.3. 標準仕上げ（Ｓ部材）：StbCalFinish_S 
・概要 
説明  
：Ｓ部材の標準仕上げ 
親要素 
：StbCalFinish 
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
StbCalFinish_S_member 
1 
5 
部材種別ごとの標準仕上げ（Ｓ部材） 
 
・補足 
Ｓ部材の単位面積あたり仕上げ重量で、個別または階別指定のない該当部材に適用される標準値
がある場合、その値を指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalFinish_S> 
        <StbCalFinish_S_member member_type="COLUMN" finishing_weight="0.0006" 
                covering_type="B" covering_size="30" covering_unit_weight="1.35e-7"/> 
<StbCalFinish_S_member member_type="BEAM" finishing_weight="0.0006" 
 covering_type="A" covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalFinish_S> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 58

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
48 
 
4.3.4. 部材種別ごとの標準仕上げ（Ｓ部材）：StbCalFinish_S_member 
・概要 
説明  
：部材種別ごとのＳ部材の標準仕上げ 
親要素 
：StbCalFinish_S 
・属性 
属性名 
型 
必須 
説明 
補足 
member_type 
string 
○ 
部材種別 以下のいずれかの値をとる 
COLUMN(柱)  GIRDER(大梁) 
BEAM(小梁)   CANTI(片持梁) 
BRACE(ブレース) 
 
finishing_weight 
double 
 
仕上げ重量(N/mm2) 
※(1) 
covering_type 
string 
○ 
被覆形式 A(形式A) B(形式B) 
C(形式C) D(形式D) 
※(2) 
covering_size 
double 
 
被覆寸法(mm) 
※(2) 
covering_unit_weight 
double 
 
被覆材単位重量(N/mm3) 
※(2) 
load_ratio 
double 
 
鉄骨重量の割増率 
※(3) 
・内容 
無し 
・子要素 
無し 
・補足 
仕上げ面と、耐火被覆の組合せで指定する。 
(1) 省略された場合は、0.0 とする。 
(2) 被覆形式は、以下とする。A,B,C の場合、被覆寸法と被覆材単位重量は必須とする。 
 
形式A 
形式B 
形式C 
形式D 
 
d 
 
 
 
 
 
 
 
 
 
 
                                       ：仕上げ面      
                    ：被覆材         d：被覆寸法 
鋼管の場合は、B またはD とする。 
(3) 省略された場合は、1.0 とする。

## Page 59

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
49 
 
4.3.5. 階別仕上げ（複数）：StbCalFloorFinishes 
・概要 
説明  
：階別の標準仕上げ 
親要素 
：StbCalFinish 
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
StbCalFloorFinish 
1 
制限無し 
階別仕上げ 
※(1) 
・補足 
指定した階について、 
大梁および小梁の場合  ・・・始端節点 
柱およびブレースの場合 ・・・始端節点 
床および壁の場合    ・・・第１節点 
が、それぞれ属する場合に、該当する標準仕上げが適用される。 
 
(1) ST-Bridge 本編にて、StbStories がない場合は、StbCalFloorFinishes も存在しないものとし、 
StbStories がある場合は、子要素の最大回数は、StbStory の要素数を超えないものとする。

## Page 60

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
50 
 
4.3.6. 階別仕上げ：StbCalFloorFinish 
・概要 
説明  
：階別の標準仕上げ 
親要素 
：StbCalFloorFinishes 
・属性 
属性名 
型 
必須 
説明 
補足 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCalFloorFinish_RC 
0 
1 
階別標準仕上げ（ＲＣ部材） 
 
StbCalFloorFinish_S 
0 
1 
階別標準仕上げ（Ｓ部材） 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalFloorFinishes> 
        <StbCalFloorFinish id_story="1"> 
          <StbCalFloorFinish_S> 
            <StbCalFloorFinish_S_member member_type="COLUMN" finishing_weight="0.0006" 
                    covering_type="B" covering_size="30" covering_unit_weight="1.35e-7"/> 
<StbCalFloorFinish_S_member member_type="BEAM" finishing_weight="0.0006" 
 covering_type="A" covering_size="30" covering_unit_weight="1.35e-7"/> 
        </StbCalFloorFinish_S> 
        </StbCalFloorFinish> 
       （略） 
      </StbCalFloorFinishes> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 61

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
51 
 
4.3.7. 階別標準仕上げ（ＲＣ部材）：StbCalFloorFinish_RC 
・概要 
説明  
：ＲＣ部材の階別標準仕上げ 
親要素 
：StbCalFloorFinish 
・属性 
属性名 
型 
必須 
説明 
補足 
weight_girder 
double 
 
大梁仕上げ重量(N/mm2) 
※(2) 
type_girder 
integer 
 
大梁仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
※(1) 
weight_column 
double 
 
柱仕上げ重量(N/mm2) 
※(2) 
type_column 
integer 
 
柱仕上げタイプ 
1：4 面仕上げ 2：2 面仕上げ 
3：仕上げ無し 
※(1) 
weight_beam 
double 
 
小梁仕上げ重量(N/mm2) 
※(2) 
type_beam 
integer 
 
小梁仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
※(1) 
weight_canti 
double 
 
片持梁仕上げ重量(N/mm2) 
※(1) 
type_canti 
integer 
 
片持梁仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
※(1) 
weight_wall 
double 
 
壁仕上げ重量(N/mm2) 
※(2) 
type_wall 
integer 
 
仕上げタイプ 
1：両面仕上げ 2：片面仕上げ 
3：仕上げ無し 
※(1) 
weight_slab 
double 
 
スラブ仕上げ重量(N/mm2) 
※(3) 
・内容 
無し 
・子要素 
無し 
・補足 
ＲＣ部材の単位面積あたり仕上げ重量で、親要素で指定した階に所属し、個別指定のない該当部
材に適用される標準値がある場合、その値とタイプを指定する。 
(1) それぞれの部材種別について、標準値がある場合は必須とする。

## Page 62

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
52 
 
(2) それぞれの部材種別について、仕上タイプが「仕上げ無し」以外のとき、仕上げ重量は必須とす
る。 
(3) 仕上げの総和の値とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalFloorFinishes> 
        <StbCalFloorFinish id_story="1"> 
          <StbCalFloorFinish_RC weight_girder="0.0006" type_girder="1" 
weight_column="0.0006" type_column="1" 
weight_beam="0.0006" type_beam="1" 
weight_canti="0.0006" type_canti="1" 
weight_wall="0.0008" type_wall="1" weight_slab="0.0008"/> 
        </StbCalFloorFinish> 
       （略） 
      </StbCalFloorFinishes> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 63

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
53 
 
4.3.8. 階別標準仕上げ（Ｓ部材）：StbCalFloorFinish_S 
・概要 
説明  
：Ｓ部材の階別標準仕上げ 
親要素 
：StbCalFloorFinish 
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
StbCalFloorFinish_S_member 
1 
5 
部材種別ごとの階別標準仕上げ（Ｓ部材） 
 
・補足 
Ｓ部材の単位面積あたり仕上げ重量で、親要素で指定した階に所属し、個別指定のない該当部材
に適用される標準値がある場合、その値を指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalFloorFinishes> 
        <StbCalFloorFinish id_story="1"> 
          <StbCalFloorFinish_S> 
            <StbCalFloorFinish_S_member member_type="COLUMN" finishing_weight="0.0006" 
                    covering_type="B" covering_size="30" covering_unit_weight="1.35e-7"/> 
<StbCalFloorFinish_S_member member_type="BEAM" finishing_weight="0.0006" 
 covering_type="A" covering_size="30" covering_unit_weight="1.35e-7"/> 
        </StbCalFloorFinish_S> 
        </StbCalFloorFinish> 
       （略） 
      </StbCalFloorFinishes> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 64

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
54 
 
4.3.9. 部材種別ごとの階別標準仕上げ（Ｓ部材）：StbCalFloorFinish_S_member 
・概要 
説明  
：部材種別ごとのＳ部材の階別標準仕上げ 
親要素 
：StbCalFloorFinish_S 
・属性 
属性名 
型 
必須 
説明 
補足 
member_type 
string 
○ 
部材種別 以下のいずれかの値をとる 
COLUMN(柱)  GIRDER(大梁) 
BEAM(小梁)   CANTI(片持梁) 
BRACE(ブレース) 
 
finishing_weight 
double 
 
仕上げ重量(N/mm2) 
※(1) 
covering_type 
string 
○ 
被覆形式 A(形式A) B(形式B) 
C(形式C) D(形式D) 
※(2) 
covering_size 
double 
 
被覆寸法(mm) 
※(2) 
covering_unit_weight 
double 
 
被覆材単位重量(N/mm3) 
※(2) 
load_ratio 
double 
 
鉄骨重量の割増率 
※(3) 
・内容 
無し 
・子要素 
無し 
・補足 
仕上げ面と、耐火被覆の組合せで指定する。 
(1) 省略された場合は、0.0 とする。 
(2) 被覆形式は、以下とする。A,B,C の場合、被覆寸法と被覆材単位重量は必須とする。 
 
形式A 
形式B 
形式C 
形式D 
 
d 
 
 
 
 
 
 
 
 
 
 
                                       ：仕上げ面      
                    ：被覆材         d：被覆寸法 
鋼管の場合は、B またはD とする。 
(3) 省略された場合は、1.0 とする。

## Page 65

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
55 
 
4.3.10. 部材仕上げ（ＲＣ部材）：StbCalMemberFinishes_RC 
・概要 
説明  
：部材仕上げ（ＲＣ部材、個別指定用） 
親要素 
：StbCalFinish 
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
StbCalColumnFinish_RC 
0 
制限無し 
柱仕上げ（ＲＣ部材） 
 
StbCalGirderFinish_RC 
0 
制限無し 
梁仕上げ（ＲＣ部材） 
 
StbCalSlabFinish_RC 
0 
制限無し 
スラブ仕上げ（ＲＣ部材） 
 
StbCalWallFinish_RC 
0 
制限無し 
壁仕上げ（ＲＣ部材） 
 
・補足 
部材に参照される仕上げ重量の値を指定する。 
子要素の並びは、上表に示す順番としなければならない。

## Page 66

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
56 
 
4.3.11. 柱仕上げ（ＲＣ部材）：StbCalColumnFinish_RC 
・概要 
説明  
：柱部材ごとの仕上げ（ＲＣ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_RC 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
type 
integer 
○ 
仕上げタイプ 
1：4 面仕上げ 2：2 面仕上げ 
3：仕上げ無し 
 
weight 
double 
 
仕上げ重量(N/mm2) 
 
・内容 
無し 
・子要素 
無し 
・補足 
柱部材の仕上げ重量の値を指定する。「4 面」「2 面」仕上げのとき、仕上げ重量は必須とする。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID を柱部材に適用する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalMemberFinishes_RC> 
        <StbCalColumnFinish_RC id="1" type="1" weight="0.0005"/> 
        <StbCalColumnFinish_RC id="2" type="2" weight="0.00095"/> 
        （略） 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 67

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
57 
 
4.3.12. 梁仕上げ（ＲＣ部材）：StbCalGirderFinish_RC 
・概要 
説明  
：梁部材ごとの仕上げ（ＲＣ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_RC 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
type 
integer 
○ 
仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
 
weight 
double 
 
仕上げ重量(N/mm2) 
 
・内容 
無し 
・子要素 
無し 
・補足 
梁部材の仕上げ重量の値を指定する。「4 面」「2 面」仕上げのとき、仕上げ重量は必須とする。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID を大梁、小梁部材に適用する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalMemberFinishes_RC> 
        <StbCalGirderFinish_RC id="3" type="1" weight="0.0005"/> 
        <StbCalGirderFinish_RC id="4" type="2" weight="0.00095"/> 
        （略） 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 68

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
58 
 
4.3.13. スラブ仕上げ：StbCalSlabFinish_RC 
・概要 
説明  
：スラブ部材ごとの仕上げ（ＲＣ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_RC 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
weight 
double 
 
仕上げ重量(N/mm2) 
 
・内容 
無し 
・子要素 
無し 
・補足 
スラブ部材の仕上げ重量の値を指定する。仕上げ重量は、仕上げの総和の値とする。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID をスラブ部材に適用する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalMemberFinishes_RC> 
        <StbCalSlabFinish_RC id="5" weight="0.001"/> 
        <StbCalSlabFinish_RC id="6" weight="0.002"/> 
        （略） 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 69

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
59 
 
4.3.14. 壁仕上げ：StbCalWallFinish_RC 
・概要 
説明  
：壁部材ごとの仕上げ（ＲＣ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_RC 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
type 
integer 
○ 
仕上げタイプ 
1：両面仕上げ 2：片面仕上げ 
3：仕上げ無し 
 
weight 
double 
 
仕上げ重量(N/mm2) 
 
・内容 
無し 
・子要素 
無し 
・補足 
壁部材の仕上げ重量の値を指定する。「両面」「2 面」仕上げのとき、仕上げ重量は必須とする。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID を壁部材に適用する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
     （略） 
      <StbCalMemberFinishes_RC> 
        <StbCalWallFinish_RC id="7" type="1" weight="0.001"/> 
        <StbCalWallFinish_RC id="8" type="1" weight="0.002"/> 
        （略） 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 70

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
60 
 
4.3.15. 部材仕上げ（Ｓ部材）：StbCalMemberFinishes_S 
・概要 
説明  
：部材ごとの仕上げ（Ｓ部材、個別指定用） 
親要素 
：StbCalFinish 
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
StbCalColumnFinish_S 
0 
制限無し 
柱仕上げ（Ｓ部材） 
 
StbCalGirderFinish_S 
0 
制限無し 
梁仕上げ（Ｓ部材） 
 
StbCalBraceFinish_S 
0 
制限無し 
ブレース仕上げ（Ｓ部材） 
 
・補足 
部材に参照される仕上げ重量の値を指定する。 
子要素の並びは、上表に示す順番としなければならない。

## Page 71

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
61 
 
4.3.16. 柱仕上げ（Ｓ部材）：StbCalColumnFinish_S 
・概要 
説明  
：柱部材ごとの仕上げ（Ｓ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_S 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
type 
integer 
○ 
仕上げタイプ 
1：4 面仕上げ 2：2 面仕上げ 
3：仕上げ無し 
 
weight 
double 
 
仕上げ重量(N/mm2) 
※(1) 
covering_unit_weight 
double 
 
被覆材単位重量(N/mm3) 
※(2) 
covering_size 
double 
 
被覆寸法(mm) 
※(2) 
covering_type 
string 
 
被覆形式 A(形式A) B(形式B) 
C(形式C) D(形式D) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
柱部材の仕上げ重量の値を、仕上げ面と耐火被覆の組合せで指定する。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID を柱部材に適用する。 
(1) 省略された場合は、0.0 とする。 
(2) 被覆形式は、「部材種別ごとの標準仕上げ（Ｓ部材）：StbCalFinish_S_member」を参照。被覆
形式がA,B,C の場合、被覆寸法と被覆材単位重量は必須とする。

## Page 72

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
62 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
        <StbCalColumnFinish_S id="1" type="1" weight="0.0006" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
        <StbCalColumnFinish_S id="2" type="2" weight="0.0008" covering_type="A" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
</StbCalData>

## Page 73

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
63 
 
4.3.17. 梁仕上げ（Ｓ部材）：StbCalGirderFinish_S 
・概要 
説明  
：梁部材ごとの仕上げ（Ｓ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_S 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
type 
integer 
○ 
仕上げタイプ 
1：両側仕上げ 2：片側仕上げ 
3：仕上げ無し 
 
weight 
double 
 
仕上げ重量(N/mm2) 
※(1) 
covering_unit_weight 
double 
 
被覆材単位重量(N/mm3) 
※(2) 
covering_size 
double 
 
被覆寸法(mm) 
※(2) 
covering_type 
string 
 
被覆形式 A(形式A) B(形式B) 
C(形式C) D(形式D) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
梁部材の仕上げ重量の値を、仕上げ面と耐火被覆の組合せで指定する。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID を大梁、小梁部材に適用する。 
(1) 省略された場合は、0.0 とする。 
(2) 被覆形式は、「部材種別ごとの標準仕上げ（Ｓ部材）：StbCalFinish_S_member」を参照。被覆
形式がA,B,C の場合、被覆寸法と被覆材単位重量は必須とする。

## Page 74

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
64 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
        <StbCalGirderFinish_S id="3" type="1" weight="0.0005" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
        <StbCalGirderFinish_S id="4" type="2" weight="0.0005" covering_type="A" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
</StbCalData>

## Page 75

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
65 
 
4.3.18. ブレース仕上げ：StbCalBraceFinish_S 
・概要 
説明  
：ブレース部材ごとの仕上げ（Ｓ部材、個別指定用） 
親要素 
：StbCalMemberFinishes_S 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
仕上げID 
 
guid 
string 
 
GUID 
 
type 
integer 
○ 
仕上げタイプ 
1：4 面仕上げ 2：2 面仕上げ 
3：仕上げ無し 
 
weight 
double 
 
仕上げ重量(N/mm2) 
※(1) 
covering_unit_weight 
double 
 
被覆材単位重量(N/mm3) 
※(2) 
covering_size 
double 
 
被覆寸法(mm) 
※(2) 
covering_type 
string 
 
被覆形式 A(形式A) B(形式B) 
C(形式C) D(形式D) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
ブレース部材の仕上げ重量の値を、仕上げ面と耐火被覆の組合せで指定する。 
「荷重配置：StbCalLoadArrangements」にて、ここで定義したID をブレース部材に適用する。 
(1) 省略された場合は、0.0 とする。 
(2) 被覆形式は、「部材種別ごとの標準仕上げ（Ｓ部材）：StbCalFinish_S_member」を参照。被覆
形式がA,B,C の場合、被覆寸法と被覆材単位重量は必須とする。

## Page 76

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
66 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
        <StbCalBraceFinish_S id="5" type="1" weight="0.0005" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
        <StbCalBraceFinish_S id="6" type="2" weight="0.0005" covering_type="A" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
</StbCalData>

## Page 77

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
67 
 
4.3.19. 荷重ケース（複数）：StbCalLoadCases 
・概要 
説明  
：荷重ケースの定義（複数） 
親要素 
：StbCalLoad 
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
StbCalLoadCase 
1 
制限無し 
荷重ケース 
 
・補足 
StbCalLoad で定義される荷重のうち、特殊荷重・追加荷重や付加重量は、主に、構造躯体以外
の重量を、固定荷重や積載荷重に追加したい場合、地震時の水平力に付加したい場合などに用いら
れる。 
解析時には、検討する荷重の種類に応じて特殊荷重・追加荷重や付加重量を与える必要がある。
そのため、計算条件要素（StbCalData）内では、各荷重に「荷重ケース」として属性を付与して分
類に用い、解析時にはその属性を参照して荷重を設定することとする。 
 
StbCalLoadCase では、各荷重に付与する「荷重ケース」を定義する。

## Page 78

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
68 
 
4.3.20. 荷重ケース：StbCalLoadCase 
・概要 
説明  
：荷重ケースの定義 
親要素 
：StbCalLoadCases 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
荷重ケースID 
※(1) 
guid 
string 
 
GUID 
 
category 
string 
○ 
荷重分類 STANDARD または 
ANALYSIS 
※(2) 
kind 
string 
○ 
荷重種類 
※(3) 
name 
string 
 
荷重ケース名 
※(4) 
direction 
integer 
 
水平系荷重のとき荷重方向 
※(5) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 各荷重データから参照される荷重ケースID を指定する。  
(2) 主として、日本の建築基準法施行令で用いられる荷重の区分による場合STANDARD、応力解
析における外力の定義による場合ANALYSIS とする。 
(3) Ver2.0 以降では、以下の名称とする。 
・荷重分類 STANDARD の場合 
名称 
荷重種類 
DL 
令第８４条に規定する固定荷重 
LLf 
令第８５条に規定するラーメン用積載荷重 
LLe 
令第８５条に規定する地震用積載荷重 
TL 
常時荷重 TL (=DL+LLf) 
S 
令第８６条に規定する積雪荷重 
 
・荷重分類 ANALYSIS の場合 
名称 
荷重種類 
TL 
常時荷重時解析に追加する荷重 
S 
積雪荷重時解析に追加する荷重 
K 
地震時解析に追加する荷重

## Page 79

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
69 
 
W 
風荷重解析に追加する荷重 
OTHER 
その他の荷重時解析に追加する荷重 
 
(4) 各プログラムで、独自に名称を付けたい場合に記述する。 
(5) 荷重分類 ANALYSIS の場合で水平系外力の場合必須とし、以下の数値とする。 
 
名称 
荷重種類 
-180 
X
―
 加力正負方向 
0 
X
―
 加力正方向 
180 
X
―
 加力負方向 
-270 
Y
―
 加力正負方向 
90 
Y
―
 加力正方向 
270 
Y
―
 加力負方向 
その他の正数 
全体座標系 X
―
 軸とのなす角度方向（<360） 
その他の正数 
全体座標系 X
―
 軸とのなす角度の正負方向 
（<-180、> 360） 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalLoadCases> 
      <StbCalLoadCase id="1" catergory="STANDARD" kind="DL" name="case-G"/> 
      <StbCalLoadCase id="2" catergory="STANDARD" kind="LLf" name="case-P"/> 
      <StbCalLoadCase id="3" catergory="STANDARD" kind="S" name="case-S"/> 
      <StbCalLoadCase id="4" catergory="ANALYSIS" kind="K" name="case-Xp" direction ="0"/> 
      <StbCalLoadCase id="5" catergory="ANALYSIS" kind="K" name="case-Yp" direction ="90"/> 
      <StbCalLoadCase id="6" catergory="ANALYSIS" kind="K" name="case-Xm" direction ="180"/> 
      <StbCalLoadCase id="7" catergory="ANALYSIS" kind="K" name="case-Ym" direction ="270"/> 
  </StbCalLoadCases> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 80

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
70 
 
4.3.21. 特殊荷重：StbCalAdditionalLoads 
・概要 
説明  
：特殊荷重の定義 
親要素 
：StbCalLoad 
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
StbCalMemberLoad 
0 
制限無し 
部材特殊荷重 
 
StbCalAreaLoad 
0 
制限無し 
面特殊荷重 
 
StbCalPointLoad 
0 
制限無し 
節点特殊荷重 
 
StbCalSelectedPointLoad 
0 
制限無し 
任意点特殊荷重 
 
StbCalEarthHydrostaticPressureLoad 
0 
制限無し 
土圧・水圧荷重 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 81

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
71 
 
4.3.22. 部材特殊荷重：StbCalMemberLoad 
・概要 
説明  
：はり、柱など線材に直接（追加）考慮する特殊荷重 
親要素 
：StbCalAdditionalLoads 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
特殊荷重ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名 
※(1) 
type 
integer 
○ 
荷重タイプ 
※(2) 
P1 
double 
○ 
荷重パラメータ P1 
※(2) 
P2 
double 
 
荷重パラメータ P2 
※(2) 
P3 
double 
 
荷重パラメータ P3 
※(2) 
P4 
double 
 
荷重パラメータ P4 
※(2) 
P5 
double 
 
荷重パラメータ P5 
※(2) 
P6 
double 
 
荷重パラメータ P6 
※(2) 
coordinate_load 
string 
 
荷重作用方向 
LOCAL：部材座標系 
GLOBAL：全体座標系(長さ実長) 
PROJECTION：全体座標系(長さ投影) 
※(3) 
direction_load 
string 
 
荷重作用軸 
X：部材座標系X 軸方向 
Y：部材座標系Y 軸方向 
※(4) 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。

## Page 82

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
72 
 
(2) 荷重タイプの定義は以下とする。 
 
Type 
荷重図 
P1 
P2 
P3 
P4 
P5 
P6 
1 
 
P1 
(N) 
L1 
(mm) 
P2 
(N) 
L2 
(mm) 
P3 
(N) 
L3 
(mm) 
2 
 
M1 
(Nmm) 
L1 
(mm) 
M2 
(Nmm) 
L2 
(mm) 
M3 
(Nmm) 
L3 
(mm) 
3 
 
P 
(N) 
個数 
 
 
 
 
4 
 
w 
(N/mm) 
 
 
 
 
 
5 
 
w 
(N/mm) 
L 
(mm) 
 
 
 
 
6 
 
w1 
(N/mm) 
w2 
(N/mm) 
L1 
(mm) 
L2 
(mm) 
 
 
7 
 
w1 
(N/mm) 
w2 
(N/mm) 
L1 
(mm) 
L2 
(mm) 
 
 
8 
 
w1 
(N/mm) 
w2 
(N/mm) 
w3 
(N/mm) 
L1 
(mm) 
L2 
(mm) 
L3 
(mm) 
9 
 
w1 
(N/mm) 
w2 
(N/mm) 
w3 
(N/mm) 
L1 
(mm) 
L2 
(mm) 
L3 
(mm) 
10 
 
Ci 
(Nmm) 
Cj 
(Nmm) 
Qoi 
(N) 
Qoj 
(N) 
Mo 
(Nmm) 
 
11 
 
q1 
(N/mm2) 
q2 
(N/mm2) 
スパン長 
(mm)

## Page 83

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
73 
 
12 
 
q 
(N/mm2) 
L 
(mm) 
スパン長 
(mm) 
 
 
 
13 
 
q1 
(N/mm2) 
q2 
(N/mm2) 
q3 
(N/mm2) 
L1 
(mm) 
L2 
(mm) 
スパン長 
(mm) 
14 
 
q 
(N/mm2) 
個数 
スパン長 
(mm) 
 
 
 
 
① 
表の空欄以外のパラメータは省略しないものとする。ただし、Type 1、Type 2 において、
荷重１か所の場合はP3～P6 を、２か所の場合はP5～P6 を省略する。 
② 
Type 10 は、一貫構造計算プログラム等でデータとして荷重値を指定された場合を想定し
ており、各荷重形から計算されたCMQo 値を指定するものではない。 
③ 
荷重図において、線材の左端および右端は、梁の場合は、StbGirder,StbBeam の「始端」
「終端」を、柱の場合は、StbColumn,StbPost の「柱脚」「柱頭」を示す。 
④ 
荷重の符号（＋，－）は、梁の場合は、部材座標系Z 軸において図の下向き（負方向）を
正とし、柱の場合は属性：荷重作用軸による。 
⑤ 
部材座標系の定義は、ST-Bridge 本編StbGirder およびStbColumn の「補足」による。 
 
(3) 全体座標系に対し、部材が傾いている場合の扱いを指定する。省略時は LOCAL とする。 
LOCAL：             GLOBAL：      PROJECTION： 
 
 
(4) 柱の場合、荷重が作用する軸の向きを指定する。それぞれ、座標軸に対して図の下向き（負方
向）を正とする。省略された場合は X とする。

## Page 84

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
74 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalAdditionalLoads> 
     <StbCalMemberLoad id="1" id_loadcase="1" loadcase="DL" type="1" 
  P1="10000" P2="1000" description="Point load A"/> 
     <StbCalMemberLoad id="2" id_loadcase="1" loadcase="DL" type="1" 
  P1="10000" P2="2000" description="Point load B"/> 
     <StbCalMemberLoad id="3" id_loadcase="1" loadcase="DL" type="4" 
  P1="5.0" description="Equipment load C"/> 
     <StbCalAreaLoad id="4" id_loadcase="1" loadcase="DL" type="51" 
  P1="0.0005" description="additional load A"/> 
     <StbCalAreaLoad id="5" id_loadcase="1" loadcase="DL" type="51" 
  P1="0.0008" description="additional load B"/> 
      （略） 
</StbCalAdditionalLoads> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 85

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
75 
 
4.3.23. 面特殊荷重：StbCalAreaLoad 
・概要 
説明  
：壁、スラブなど面材に直接（追加）考慮する特殊荷重 
親要素 
：StbCalAdditionalLoads 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
特殊荷重ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名 
※(1) 
type 
integer 
○ 
荷重タイプ 
※(2) 
P1 
double 
○ 
荷重パラメータ P1 
※(2) 
coordinate_load 
string 
 
荷重作用方向 
LOCAL：部材座標系 
GLOBAL：全体座標系(長さ実長) 
PROJECTION：全体座標系(長さ投影) 
※(3) 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。 
 
(2) 荷重タイプの定義は以下とする。荷重の符号（＋，－）は、部材座標系Z 軸において図の下向
き（負方向）を正とする。部材座標系の定義は、ST-Bridge 本編StbWall およびStbSlab の「補
足」による。 
 
Type 
荷重図 
P1 
51 
 
 
等分布 q1 (N/mm2) で指定 
52 
 
総重量 W (N)  で指定 
 
(3) 全体座標系に対し、部材が傾いている場合の扱いを指定する。省略時は "LOCAL" とする。

## Page 86

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
76 
 
4.3.24. 節点特殊荷重：StbCalPointLoad 
・概要 
説明  
：節点に直接（追加）考慮する特殊荷重 
親要素 
：StbCalAdditionalLoads 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
特殊荷重ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名 
※(1) 
P1 
double 
 
荷重パラメータ P1 
※(2) 
P2 
double 
 
荷重パラメータ P2 
※(2) 
P3 
double 
 
荷重パラメータ P3 
※(2) 
P4 
double 
 
荷重パラメータ P4 
※(2) 
P5 
double 
 
荷重パラメータ P5 
※(2) 
P6 
double 
 
荷重パラメータ P6 
※(2) 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
節点位置に作用する荷重について、応力解析用として追加指定する場合に用いる。なお、重量計
算用に追加する重量の値については、StbCalNodeAddedWeight（節点追加・補正重量）にて指定
する。 
 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。 
(2) 各パラメータの定義は以下とする。荷重の符号は、全体座標系において、荷重図に示す矢印の
向きを正とする。省略された場合は0.0 とする。ただし、P1～P6 のうち、少なくとも１個の荷
重パラメータ属性を記述するものとする。 
荷重図 
P1 
P2 
P3 
P4 
P5 
P6 
 
Px 
(N) 
Py 
(N) 
Pz 
(N) 
Mx 
(Nmm) 
My 
(Nmm) 
Mz 
(Nmm)

## Page 87

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
77 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalPointLoad id="501" id_loadcase="1" loadcase="DL" 
  P3="-1000" description="additional load C"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
</StbCalData>

## Page 88

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
78 
 
4.3.25. 任意点特殊荷重：StbCalSelectedPointLoad 
・概要 
説明  
：任意位置に直接（追加）考慮する特殊荷重 
親要素 
：StbCalAdditionalLoads 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
特殊荷重ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名 
※(1) 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
 
地震力計算グループID 
※(2) 
X 
double 
○ 
X
―
 座標(mm) 
 
Y 
double 
○ 
Y
―
 座標(mm) 
 
P1 
double 
 
荷重パラメータ P1 
※(3) 
P2 
double 
 
荷重パラメータ P2 
※(3) 
P3 
double 
 
荷重パラメータ P3 
※(3) 
P4 
double 
 
荷重パラメータ P4 
※(3) 
P5 
double 
 
荷重パラメータ P5 
※(3) 
P6 
double 
 
荷重パラメータ P6 
※(3) 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
節点以外の任意位置に作用する荷重について、応力解析用として追加指定する場合に用いる。 
 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。 
(2) 荷重ケースが水平系荷重で、階に対して地震力計算グループ条件の指定が必要な場合、指定す
る。 
(3) 各パラメータの定義は、「節点特殊荷重：StbCalPointLoad」にならう。

## Page 89

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
79 
 
4.3.26. 土圧・水圧荷重：StbCalEarthHydrostaticPressureLoad 
・概要 
説明  
：土圧・水圧荷重 
親要素 
：StbCalAdditionalLoads 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
特殊荷重ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名 
※(1) 
P1 
double 
 
荷重パラメータ P1 
※(2) 
P2 
double 
 
荷重パラメータ P2 
※(2) 
P3 
double 
 
荷重パラメータ P3 
※(2) 
P4 
double 
 
荷重パラメータ P4 
※(2) 
coordinate_load 
string 
 
作用面方向 plus または minus 
※(3) 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
壁位置に水平方向に作用する、土圧や水圧に相当する荷重について、応力解析用として追加指定
する場合に用いる。柱に対し、壁を介して亀甲分割で作用する荷重の場合は、「部材特殊荷重：
StbCalMemberLoad」で指定する。 
 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。 
(2) 荷重パラメータの定義は以下とする。 
荷重図 
P1 
P2 
P3 
P4 
 
下端からの 
距離 
L1(mm) 
上端からの 
距離 
L2(mm) 
上部分布 
荷重値 
W1(N/mm2) 
下部分布 
荷重値 
W2(N/mm2)

## Page 90

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
80 
 
(3) 壁に対して配置する場合は、部材座標系Z 方向正側加力の場合plus、負側加力の場合minus と
する。部材座標系の定義は、ST-Bridge 本編StbWall の「補足」による。 
4.3.27. 追加・補正重量：StbCalAddedWeights 
・概要 
説明  
：追加・補正重量の定義 
親要素 
：StbCalLoad 
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
StbCalNodeAddedWeight 
0 
制限無し 
節点追加・補正重量 
 
StbCalSelectedNodeAddedWeight 
0 
制限無し 
任意点追加・補正重量 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 91

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
81 
 
4.3.28. 節点追加・補正重量：StbCalNodeAddedWeight 
・概要 
説明  
：重量計算用として節点位置に追加する重量 
親要素 
：StbCalAddedWeights 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
追加・補正重量ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名（補足欄参照） 
※(1) 
weight 
double 
○ 
追加重量(N) 
 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
節点位置に作用する荷重について、重量計算用に追加する重量の値を指定する場合に用いる。 
なお、応力解析用として外力を直接指定する場合は、StbCalPointLoad（節点特殊荷重）にて指
定する。 
 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalAddedWeights> 
     <StbCalNodeAddedWeight id="6" id_loadcase="1" loadcase="DL" weight ="1000"/> 
     <StbCalNodeAddedWeight id="7" id_loadcase="1" loadcase="DL" weight ="1200"/> 
</StbCalAddedWeights> 
    （略） 
  </StbCalLoad> 
</StbCalData>

## Page 92

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
82 
 
4.3.29. 任意点追加・補正重量：StbCalSelectedNodeAddedWeight 
・概要 
説明  
：重量計算用として任意位置に追加する重量 
親要素 
：StbCalAddedWeights 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
追加・補正重量ID 
 
guid 
string 
 
GUID 
 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
loadcase 
string 
 
荷重ケース名（補足欄参照） 
※(1) 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
 
地震力計算グループID 
※(2) 
X 
double 
○ 
X
―
 座標(mm) 
 
Y 
double 
○ 
Y
―
 座標(mm) 
 
weight 
double 
○ 
追加重量(N) 
 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
節点以外の任意位置に作用する荷重について、重量計算用に追加する重量の値を指定する場合に
用いる。 
 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ID とする。「荷重ケース名」はプログラ
ム独自で名前を付けたい場合に用いる。 
(2) 荷重ケースが水平系荷重で、階に対して地震力計算グループ条件の指定が必要な場合、指定す
る。

## Page 93

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
83 
 
4.3.30. 地震荷重：StbCalSeismic 
・概要 
説明  
：地震荷重に関する計算条件および荷重値 
親要素 
：StbCalLoad 
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
StbCalSeismicConditionGroups 
1 
1 
地震力計算グループ（複数） 
 
StbCalSeismicDirections 
0 
1 
方向別 地震力計算条件（複数）  
StbCalSeismicWeights 
0 
1 
補正地震用重量（複数） 
 
StbCalGravityPointWeights 
0 
1 
補正重心計算用重量（複数） 
 
StbCalShearcoefficients 
0 
制限無し 
層せん断力係数（複数） 
 
StbCalShearforces 
0 
制限無し 
層せん断力（複数） 
 
StbCalEarthquakeforces 
0 
制限無し 
地震水平力（複数） 
 
・補足 
地震荷重に関する計算条件のうち、各層（階）および各節点に配置する荷重等の値を指定する（詳
細は、各子要素の項を参照）。

## Page 94

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
84 
 
4.3.31. 地震力計算グループ（複数）：StbCalSeismicConditionGroups 
・概要 
説明  
：地震力計算グループの定義（複数） 
親要素 
：StbCalSeismic 
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
StbCalSeismicConditionGroup 
1 
制限無し 
地震力計算グループ 
 
・補足 
構造計算プログラムによっては、ツインタワーのような形状の場合のモデル化方法として、同一
層に複数の剛床を考慮でき、層にまたがる剛床をグループ化して、それぞれのグループについて地
震力計算ができるものがある。このように、地震による水平力を定義する条件を複数指定する場合、
「地震力計算グループ」にID を設けて定義する。 
条件が１個の場合も、「地震力計算グループ」要素を１個指定する。

## Page 95

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
85 
 
4.3.32. 地震力計算グループ：StbCalSeismicConditionGroup 
・概要 
説明  
：地震力計算グループの定義 
親要素 
：StbCalSeismicConditionGroups 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
地震力計算グループID 
 
guid 
string 
 
GUID 
 
name 
string 
 
地震力計算グループ名称 
※(1) 
id_floor_divided_area 
integer 
 
床区分ID(StbCalFloorDividedArea.id) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
地震による水平力を定義する条件を指定する場合に用いる。 
 
(1) プログラム独自で名前を付けたい場合に用いる。 
(2) 「床区分：StbCalFloorDividedArea」で定義される床区分ID とする。「地震力計算グループ」
要素（この要素）が、１個の場合は省略してよい。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
<StbCalSeismicConditionGroups> 
<StbCalSeismicConditionGroup id="1" name="zone-1" id_floor_divided_area="1"/> 
<StbCalSeismicConditionGroups> 
（略） 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 96

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
86 
 
4.3.33. 方向別 地震力計算条件（複数）：StbCalSeismicDirections 
・概要 
説明  
：方向ごとの地震力計算条件（複数） 
親要素 
：StbCalSeismic 
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
StbCalSeismicDirection 
1 
制限無し 
方向別 地震力計算条件 
 
・補足 
無し

## Page 97

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
87 
 
4.3.34. 方向別 地震力計算条件：StbCalSeismicDirection 
・概要 
説明  
：方向ごとの地震力計算条件 
親要素 
：StbCalSeismicDirections 
・属性 
属性名 
型 
必須 
説明 
補足 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
baseshear_coefficient1 
double 
○ 
標準せん断力係数(1 次設計) 
 
baseshear_coefficient2 
double 
○ 
標準せん断力係数(2 次設計) 
 
lateral_coefficient 
double 
 
塔屋水平震度 
※(2) 
height 
double 
 
固有周期計算用建物高さ 
※(3) 
natural_period 
double 
 
固有周期(sec) 
※(4) 
id_seismic_condition 
integer 
○ 
地震力計算グループID 
(StbCalSeismicConditionGroup.id) 
※(5) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ケースID を指定する。指定する荷重ケー
ス要素は、属性direction（水平系荷重のとき荷重方向）が、存在するものとする。 
(2) 省略された場合は、1.0 とする。 
(3) 建築物の形状から計算するプログラムでは、省略されることがある。 
(4) 固有周期が、建物高さから計算する計算値と異なる場合に指定する。プログラムによっては実
際に計算に用いられる採用値は異なることがある。 
(5) 「地震力計算グループ：StbCalSeismicConditionGroup」で定義される地震力計算グループID
とする。地震力計算グループ条件が１個の場合も指定する。

## Page 98

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
88 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
（略） 
<StbCalSeismicDirections> 
        <StbCalSeismicDirection id_loadcase="4" 
  baseshear_coefficient1="0.2" baseshear_coefficient2="1.0" id_seismic_condition ="1"/> 
        <StbCalSeismicDirection id_loadcase="5" 
  baseshear_coefficient1="0.2" baseshear_coefficient2="1.0" id_seismic_condition ="1"/> 
        <StbCalSeismicDirection id_loadcase="6" 
  baseshear_coefficient1="0.2" baseshear_coefficient2="1.0" id_seismic_condition ="1"/> 
        <StbCalSeismicDirection id_loadcase="7" 
  baseshear_coefficient1="0.2" baseshear_coefficient2="1.0" id_seismic_condition ="1"/> 
</StbCalSeismicDirections> 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 99

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
89 
 
4.3.35. 補正地震用重量（複数）：StbCalSeismicWeights 
・概要 
説明  
：地震用重量に直接加算する重量（複数） 
親要素 
：StbCalSeismic 
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
StbCalSeismicWeight 
1 
制限無し 
補正地震用重量 
 
・補足 
無し

## Page 100

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
90 
 
4.3.36. 補正地震用重量：StbCalSeismicWeight 
・概要 
説明  
：地震用重量に直接加算する重量 
親要素 
：StbCalSeismicWeights 
・属性 
属性名 
型 
必須 
説明 
補足 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
○ 
地震力計算グループID 
(StbCalSeismicConditionGroup.id) 
※(1) 
weight 
double 
○ 
加算する地震用重量(N) 
 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 「地震力計算グループ：StbCalSeismicConditionGroup」で定義される地震力計算グループID
とする。地震力計算グループ条件が１個の場合も指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
（略） 
<StbCalSeismicWeights> 
        <StbCalSeismicWeight id_story="2" id_seismic_condition ="1" weight="120000.0"/> 
        <StbCalSeismicWeight id_story="3" id_seismic_condition ="1" weight="223000.0"/> 
</StbCalSeismicWeights> 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 101

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
91 
 
4.3.37. 補正重心計算用重量（複数）：StbCalGravityPointWeights 
・概要 
説明  
：重心計算に使用する補正用重量（複数） 
親要素 
：StbCalSeismic 
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
StbCalGravityPointWeight 
1 
制限無し 
補正重心計算用重量 
 
・補足 
無し

## Page 102

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
92 
 
4.3.38. 補正重心計算用重量：StbCalGravityPointWeight 
・概要 
説明  
：重心計算に使用する補正用重量 
親要素 
：StbCalGravityPointWeights 
・属性 
属性名 
型 
必須 
説明 
補足 
weight 
double 
○ 
重量(N) 
 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
○ 
地震力計算グループID 
(StbCalSeismicConditionGroup.id) 
※(1) 
X 
double 
○ 
X
―
 座標(mm) 
 
Y 
double 
○ 
Y
―
 座標(mm) 
 
description 
string 
 
説明 
 
・内容 
無し 
・子要素 
無し 
・補足 
重心計算に使用する補正用重量で、層重量には加算しない値を指定する。 
 
(1) 「地震力計算グループ：StbCalSeismicConditionGroup」で定義される地震力計算グループID
とする。地震力計算グループ条件が１個の場合も指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
（略） 
<StbCalGravityPointWeights> 
        <StbCalGravityPointWeight id_story="2" id_seismic_condition ="1" 
weight="1000.0" X="15000.0" Y="8500.0"/> 
</StbCalGravityPointWeights> 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 103

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
93 
 
4.3.39. 層せん断力係数（複数）：StbCalShearcoefficients 
・概要 
説明  
：方向ごとの層せん断力係数または震度（地下階、塔屋）（複数） 
親要素 
：StbCalSeismic 
・属性 
属性名 
型 
必須 
説明 
補足 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCalShearcoefficient 
1 
制限無し 
層せん断力係数 
 
・補足 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ケースID とする。属性direction（水平系
荷重のとき荷重方向）が、存在する要素を指定する。

## Page 104

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
94 
 
4.3.40. 層せん断力係数：StbCalShearcoefficient 
・概要 
説明  
：方向ごとの層せん断力係数または震度（地下階、塔屋） 
親要素 
：StbCalShearcoefficients 
・属性 
属性名 
型 
必須 
説明 
補足 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
○ 
地震力計算グループID 
(StbCalSeismicConditionGroup.id) 
※(1) 
coefficient 
double 
○ 
層せん断力係数（震度） 
 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 「地震力計算グループ：StbCalSeismicConditionGroup」で定義される地震力計算グループID
とする。地震力計算グループ条件が１個の場合も指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
（略） 
<StbCalShearcoefficients id_loadcase="4"> 
<StbCalShearcoefficient id_story="1" id_seismic_condition="1" coefficient="0.200"/> 
<StbCalShearcoefficient id_story="2" id_seismic_condition="1" coefficient="0.215"/> 
<StbCalShearcoefficient id_story="3" id_seismic_condition="1" coefficient="0.255"/> 
</StbCalShearcoefficients> 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 105

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
95 
 
4.3.41. 層せん断力（複数）：StbCalShearforces 
・概要 
説明  
：方向ごとの層せん断力（複数） 
親要素 
：StbCalSeismic 
・属性 
属性名 
型 
必須 
説明 
補足 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCalShearforce 
1 
制限無し 
層せん断力 
 
・補足 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ケースID とする。指定する荷重ケース要素
は、属性direction（水平系荷重のとき荷重方向）が、存在するものとする。

## Page 106

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
96 
 
4.3.42. 層せん断力：StbCalShearforce 
・概要 
説明  
：方向ごとの層せん断力 
親要素 
：StbCalShearforces 
・属性 
属性名 
型 
必須 
説明 
補足 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
○ 
地震力計算グループID 
(StbCalSeismicConditionGroup.id) 
※(1) 
force 
double 
○ 
層せん断力 
 
・内容 
無し 
・子要素 
無し 
・補足 
地震による水平力の数値を、層せん断力にて直接指定する。 
 
(1) 「地震力計算グループ：StbCalSeismicConditionGroup」で定義される地震力計算グループID
とする。地震力計算グループ条件が１個の場合も指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
（略） 
<StbCalShearforces id_loadcase="4"> 
<StbCalShearforce id_story="1" id_seismic_condition="1" force="1787900"/> 
<StbCalShearforce id_story="2" id_seismic_condition="1" force="1416200"/> 
<StbCalShearforce id_story="3" id_seismic_condition="1" force="1275500"/> 
</StbCalShearforces> 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 107

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
97 
 
4.3.43. 地震水平力（複数）：StbCalEarthquakeforces 
・概要 
説明  
：方向ごとの地震水平力（複数） 
親要素 
：StbCalSeismic 
・属性 
属性名 
型 
必須 
説明 
補足 
id_loadcase 
integer 
○ 
荷重ケースID(StbCalLoadCase.id) 
※(1) 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbCalEarthquakeforce 
1 
制限無し 
地震水平力 
 
・補足 
(1) 「荷重ケース：StbCalLoadCase」で定義される荷重ケースID とする。指定する荷重ケース要素
は、属性direction（水平系荷重のとき荷重方向）が、存在するものとする。

## Page 108

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
98 
 
4.3.44. 地震水平力：StbCalEarthquakeforce 
・概要 
説明  
：方向ごとの地震水平力 
親要素 
：StbCalEarthquakeforces 
・属性 
属性名 
型 
必須 
説明 
補足 
id_story 
integer 
○ 
階ID(StbStory.id) 
 
id_seismic_condition 
integer 
○ 
地震力計算グループID 
(StbCalSeismicConditionGroup.id) 
※(1) 
force 
double 
○ 
水平力 
 
・内容 
無し 
・子要素 
無し 
・補足 
地震による水平力の数値を、直接指定する。 
 
(1) 「地震力計算グループ：StbCalSeismicConditionGroup」で定義される地震力計算グループID
とする。地震力計算グループ条件が１個の場合も指定する。 
・例 
<StbCalData> 
  <StbCalLoad> 
    （略） 
    <StbCalSeismic> 
（略） 
<StbCalEarthquakeforces id_loadcase="9"> 
<StbCalEarthquakeforce id_stoty="3" id_seismic_condition="1" force="597400"/> 
<StbCalEarthquakeforce id_story="4" id_seismic_condition="1" force="438900"/> 
</StbCalEarthquakeforces> 
   </StbCalSeismic> 
  </StbCalLoad> 
</StbCalData>

## Page 109

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
99 
 
4.4. 
計算条件定義：StbCalCondition 
・概要 
説明  
：計算条件定義 
親要素 
：StbCalData 
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
StbCalMemberConditions 
0 
1 
部材材端条件 
※(1) 
StbCalMemberRigidzones 
0 
1 
部材剛域長さ 
※(2) 
StbCalMemberCriticalPositions 
0 
1 
部材危険断面位置 
※(3) 
StbCalSectionProperties 
0 
1 
部材断面性能（数値指定） 
※(4) 
StbCalMemberStiffnesses 
0 
1 
部材剛性倍率 
※(5) 
StbCalNodeRestrictions 
0 
1 
節点拘束条件（複数） 
※(6) 
StbCalNodePanels 
0 
1 
接合部パネル（複数） 
※(6) 
StbCalFloorDiaphragms 
0 
1 
剛床（複数） 
※(7) 
・補足 
構造計算に必要な計算条件を指定する（詳細は、各子要素の項を参照）。 
各部材種別の部材座標系の定義は、ST-Bridge 本編の大梁、柱、ブレース、スラブおよび壁の項
による。 
(1) 部材の材端条件（剛接合、ピン接合または材端ばね）を指定する。 
(2) 部材の剛域長さを指定する。 
(3) 部材の危険断面位置を指定する。 
(4) 部材の断面形状と異なる断面性能値がある場合に指定する。 
(5) 部材の剛性倍率を指定する。 
(6) 節点に関する計算条件（拘束条件、支点ばね、パネル断面）を指定する。 
(7) 層ごとの剛床条件を、「床区分」にて指定する。

## Page 110

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
100 
 
4.4.1. 部材材端条件：StbCalMemberConditions 
・概要 
説明  
：部材の材端条件（剛接合、ピン接合または材端ばね） 
親要素 
：StbCalCondition 
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
StbCalColumnCondition 
0 
制限無し 
柱材端条件 
 
StbCalGirderCondition 
0 
制限無し 
大梁材端条件 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 111

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
101 
 
4.4.2. 柱材端条件：StbCalColumnCondition 
・概要 
説明  
：柱の材端条件（剛接合、ピン接合または材端ばね） 
親要素 
：StbCalMemberConditions 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
材端条件ID 
 
guid 
string 
 
GUID 
 
bottom_X 
string 
 
柱脚X 軸回り材端条件 以下から選択 
FIX  PIN  SPRING 
※(1) 
bottom_Y 
string 
 
柱脚Y 軸回り材端条件 以下から選択 
FIX  PIN  SPRING 
※(1) 
top_X 
string 
 
柱頭X 軸回り材端条件 以下から選択 
FIX  PIN  SPRING 
※(1) 
top_Y 
string 
 
柱頭Y 軸回り材端条件 以下から選択 
FIX  PIN  SPRING 
※(1) 
bottom_spring_X 
double 
 
柱脚X 軸回りばね定数(Nmm/rad) 
※(2) 
bottom_spring_Y 
double 
 
柱脚Y 軸回りばね定数(Nmm/rad) 
※(2) 
top_spring_X 
double 
 
柱頭X 軸回りばね定数(Nmm/rad) 
※(2) 
top_spring_Y 
double 
 
柱頭Y 軸回りばね定数(Nmm/rad) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、FIX とする。 
(2) 材端条件がSPRING のとき、該当する方向のばね定数を指定する。

## Page 112

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
102 
 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberConditions> 
<StbCalColumnCondition id="1" 
 bottom_X="FIX" bottom_Y="PIN" top_X="FIX" top_Y="PIN"/> 
<StbCalColumnCondition id="2" 
 bottom_X="SPRING" bottom_Y="SPRING" top_Y="PIN" 
               bottom_spring_X="8754242.0" bottom_spring_Y="8754242.0"/> 
</StbCalMemberConditions> 
 </StbCalCondition> 
</StbCalData>

## Page 113

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
103 
 
4.4.3. 大梁材端条件：StbCalGirderCondition 
・概要 
説明  
：大梁の材端条件（剛接合、ピン接合または材端ばね） 
親要素 
：StbCalMemberConditions 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
材端条件ID 
 
guid 
string 
 
GUID 
 
start 
string 
 
始端Y 軸回り材端条件 以下から選択 
FIX  PIN  SPRING 
※(1) 
end 
string 
 
終端Y 軸回り材端条件 以下から選択 
FIX  PIN  SPRING 
※(1) 
start_spring 
double 
 
始端Y 軸回りばね定数(Nmm/rad) 
※(2) 
end_spring 
double 
 
終端Y 軸回りばね定数(Nmm/rad) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、FIX とする。 
(2) 材端条件がSPRING のとき、該当するばね定数を指定する。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberConditions> 
<StbCalColumnCondition id="1" 
 bottom_X="FIX" bottom_Y="PIN" top_X="FIX" top_Y="PIN"/> 
<StbCalColumnCondition id="2" 
 bottom_Y="SPRING" bottom_Y="SPRING" top_Y="PIN" 
               bottom_spring_X="8754242.0" bottom_spring_Y="8754242.0"/> 
<StbCalGirderCondition id="3" start="FIX" end="FIX"/> 
<StbCalGirderCondition id="4" start="PIN" end="PIN"/> 
</StbCalMemberConditions> 
 </StbCalCondition> 
</StbCalData>

## Page 114

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
104 
 
4.4.4. 部材剛域長さ：StbCalMemberRigidzones 
・概要 
説明  
：部材の剛域長さ 
親要素 
：StbCalCondition 
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
StbCalColumnRigidzone 
0 
制限無し 
柱剛域長さ 
 
StbCalGirderRigidzone 
0 
制限無し 
大梁剛域長さ 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 115

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
105 
 
4.4.5. 柱剛域長さ：StbCalColumnRigidzone 
・概要 
説明  
：柱の剛域長さ 
親要素 
：StbCalMemberRigidzones 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛域長さID 
 
guid 
string 
 
GUID 
 
bottom_X 
double 
 
柱脚X 軸剛域長さ(mm) 
※(1) 
bottom_Y 
double 
 
柱脚Y 軸剛域長さ(mm) 
※(1) 
top_X 
double 
 
柱頭X 軸剛域長さ(mm) 
※(1) 
top_Y 
double 
 
柱頭Y 軸剛域長さ(mm) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、0.0 mm とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberRigidzones> 
<StbCalColumnRigidzone 
 id="1" bottom_X="100" bottom_Y="100" top_X="50" top_Y="50"/> 
<StbCalColumnRigidzone 
id="2" bottom_X="85" bottom_Y="100" top_X="85" top_Y="85"/> 
</StbCalMemberRigidzones> 
 </StbCalCondition> 
</StbCalData>

## Page 116

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
106 
 
4.4.6. 大梁剛域長さ：StbCalGirderRigidzone 
・概要 
説明  
：大梁の剛域長さ 
親要素 
：StbCalMemberRigidzones 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛域長さID 
 
guid 
string 
 
GUID 
 
start 
double 
 
始端剛域長さ(mm) 
※(1) 
end 
double 
 
終端剛域長さ(mm) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
剛域長さはZ 方向の値とする。 
 
(1) 省略時は、0.0 mm とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberRigidzones> 
<StbCalColumnRigidzone 
 id="1" bottom_X="100" bottom_Y="100" top_X="50" top_Y="50"/> 
<StbCalColumnRigidzone 
id="2" bottom_X="85" bottom_Y="100" top_X="85" top_Y="85"/> 
<StbCalGirderRigidzone id="3" start="50" end="50"/> 
<StbCalGirderRigidzone id="4" start="75" end="125"/> 
</StbCalMemberRigidzones> 
 </StbCalCondition> 
</StbCalData>

## Page 117

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
107 
 
4.4.7. 部材危険断面位置：StbCalMemberCriticalPositions 
・概要 
説明  
：部材の危険断面位置 
親要素 
：StbCalCondition 
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
StbCalColumnCriticalPosition 
0 
制限無し 
柱危険断面位置 
 
StbCalGirderCriticalPosition 
0 
制限無し 
大梁危険断面位置 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 118

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
108 
 
4.4.8. 柱危険断面位置：StbCalColumnCriticalPosition 
・概要 
説明  
：柱の危険断面位置 
親要素 
：StbCalMemberCriticalPositions 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
危険断面位置ID 
 
guid 
string 
 
GUID 
 
bottom_X 
double 
 
柱脚X 軸危険断面位置(mm) 
※(1) 
bottom_Y 
double 
 
柱脚Y 軸危険断面位置(mm) 
※(1) 
top_X 
double 
 
柱頭X 軸危険断面位置(mm) 
※(1) 
top_Y 
double 
 
柱頭Y 軸危険断面位置(mm) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、0.0 mm とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberCriticalPositions> 
<StbCalColumnCriticalPosition 
id="1" bottom_X="100" bottom_Y="100" top_X="50" top_Y="50"/> 
<StbCalColumnCriticalPosition 
id="2" bottom_X="85" bottom_Y="100" top_X="85" top_Y="85"/> 
</StbCalMemberCriticalPositions> 
 </StbCalCondition> 
</StbCalData>

## Page 119

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
109 
 
4.4.9. 大梁危険断面位置：StbCalGirderCriticalPosition 
・概要 
説明  
：大梁の危険断面位置 
親要素 
：StbCalMemberCriticalPositions 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
危険断面位置ID 
 
guid 
string 
 
GUID 
 
start 
double 
 
始端危険断面位置(mm) 
※(1) 
end 
double 
 
終端危険断面位置(mm) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
危険断面位置はZ 方向の値とする。 
 
(1) 省略時は、0.0 mm とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberCriticalPositions> 
<StbCalColumnCriticalPosition 
id="1" bottom_X="100" bottom_Y="100" top_X="50" top_Y="50"/> 
<StbCalColumnCriticalPosition 
id="2" bottom_X="85" bottom_Y="100" top_X="85" top_Y="85"/> 
<StbCalGirderCriticalPosition id="3" start="50" end="50"/> 
<StbCalGirderCriticalPosition id="4" start="75" end="125"/> 
</StbCalMemberCriticalPositions> 
 </StbCalCondition> 
</StbCalData>

## Page 120

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
110 
 
4.4.10. 部材断面性能（数値指定）：StbCalSectionProperties 
・概要 
説明  
：部材の断面性能（数値指定） 
親要素 
：StbCalCondition 
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
StbCalColumnSectionProperty 
0 
制限無し 
柱断面性能（数値指定） 
 
StbCalGirderSectionProperty 
0 
制限無し 
大梁断面性能（数値指定） 
 
StbCalWallSectionProperty 
0 
制限無し 
壁断面性能（数値指定） 
 
StbCalBraceSectionProperty 
0 
制限無し 
ブレース断面性能（数値指定）  
StbCalSlabSectionProperty 
0 
制限無し 
スラブ断面性能（数値指定）  
・補足 
    形状によらず、断面性能を直接指定する場合に指定する。 
 
子要素の並びは、上表に示す順番としなければならない。

## Page 121

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
111 
 
4.4.11. 柱断面性能（数値指定）：StbCalColumnSectionProperty 
・概要 
説明  
：柱の断面性能（数値指定） 
親要素 
：StbCalSectionProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
断面性能ID 
 
guid 
string 
 
GUID 
 
A 
double 
 
断面積(mm2) 
※(1) 
As_X 
double 
 
X 方向せん断断面積(mm2) 
※(1) 
As_Y 
double 
 
Y 方向せん断断面積(mm2) 
※(1) 
I_X 
double 
 
X 軸回り断面2 次モーメント(mm4) 
※(1) 
I_Y 
double 
 
Y 軸回り断面2 次モーメント(mm4) 
※(1) 
J 
double 
 
断面2 次モーメント(ねじり)(mm4) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、0.0 とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalColumnSectionProperty id="1" A="11840" As_X ="9000" As_Y ="2700" 
I_X ="202000000" I_Y ="67500000" J="375000000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
</StbCalData>

## Page 122

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
112 
 
4.4.12. 大梁断面性能（数値指定）：StbCalGirderSectionProperty 
・概要 
説明  
：大梁の断面性能（数値指定） 
親要素 
：StbCalSectionProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
断面性能ID 
 
guid 
string 
 
GUID 
 
A 
double 
 
断面積(mm2) 
※(1) 
As_Y 
double 
 
Y 方向せん断断面積(mm2) 
※(1) 
As_Z 
double 
 
Z 方向せん断断面積(mm2) 
※(1) 
I_Y 
double 
 
Y 軸回り断面2 次モーメント(mm4) 
※(1) 
I_Z 
double 
 
Z 軸回り断面2 次モーメント(mm4) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、0.0 とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalColumnSectionProperty id="1" A="11840" As_X ="9000" As_Y ="2700" 
I_X ="202000000" I_Y ="67500000" J="375000000"/> 
<StbCalGirderSectionProperty id="3" A="6290" As_Y="3850" As_Z="2230" 
I_Y="135000000" I_Z="9840000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
</StbCalData>

## Page 123

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
113 
 
4.4.13. 壁断面性能（数値指定）：StbCalWallSectionProperty 
・概要 
説明  
：壁の断面性能（数値指定） 
親要素 
：StbCalSectionProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
断面性能ID 
 
guid 
string 
 
GUID 
 
A 
double 
 
軸断面積(mm2) 
※(1) 
As 
double 
 
せん断断面積(mm2) 
※(1) 
I 
double 
 
断面2 次モーメント(mm4) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
壁エレメントモデルの置換柱に対する断面性能値を想定する。 
 
(1) 省略時は、0.0 とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalColumnSectionProperty id="1" A="11840" As_X ="9000" As_Y ="2700" 
I_X ="202000000" I_Y ="67500000" J="375000000"/> 
<StbCalGirderSectionProperty id="3" A="6290" As_Y="3850" As_Z="2230" 
I_Y="135000000" I_Z="9840000"/> 
<StbCalWallSectionProperty id="5" A="780000" As="650000" Iy="175700000000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
</StbCalData>

## Page 124

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
114 
 
4.4.14. ブレース断面性能（数値指定）：StbCalBraceSectionProperty 
・概要 
説明  
：ブレースの断面性能（数値指定） 
親要素 
：StbCalSectionProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
断面性能ID 
 
guid 
string 
 
GUID 
 
A 
double 
 
軸断面積(mm2) 
※(1) 
i 
double 
 
最小断面2 次半径(mm) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、0.0 とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalColumnSectionProperty id="1" A="11840" As_X ="9000" As_Y ="2700" 
I_X ="202000000" I_Y ="67500000" J="375000000"/> 
<StbCalGirderSectionProperty id="3" A="6290" As_Y="3850" As_Z="2230" 
I_Y="135000000" I_Z="9840000"/> 
<StbCalBraceSectionProperty id="6" A="2160" i="25"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
</StbCalData>

## Page 125

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
115 
 
4.4.15. スラブ断面性能（数値指定）：StbCalSlabSectionProperty 
・概要 
説明  
：スラブの断面性能（数値指定） 
親要素 
：StbCalSectionProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
断面性能ID 
 
guid 
string 
 
GUID 
 
t 
double 
 
厚さ(mm) 
※(1) 
E 
double 
 
ヤング係数(N/mm2) 
※(1) 
G 
double 
 
せん断弾性係数(N/mm2) 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、0.0 とする。

## Page 126

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
116 
 
4.4.16. 部材剛性倍率：StbCalMemberStiffnesses 
・概要 
説明  
：部材の剛性倍率 
親要素 
：StbCalCondition 
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
StbCalColumnStiffness 
0 
制限無し 
柱剛性倍率 
 
StbCalGirderStiffness 
0 
制限無し 
大梁剛性倍率 
 
StbCalWallStiffness 
0 
制限無し 
壁剛性倍率 
 
StbCalBraceStiffness 
0 
制限無し 
ブレース剛性倍率 
 
・補足 
    形状より計算した断面性能に対する倍率を、直接指定する場合に指定する。 
 
子要素の並びは、上表に示す順番としなければならない。

## Page 127

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
117 
 
4.4.17. 柱剛性倍率：StbCalColumnStiffness 
・概要 
説明  
：柱の剛性倍率 
親要素 
：StbCalMemberStiffnesses 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛性倍率ID 
 
guid 
string 
 
GUID 
 
ratio_axial_X 
double 
 
軸剛性倍率 
※(1) 
ratio_axial_Y 
double 
 
軸剛性倍率 
※(1) , ※(2) 
ratio_shear_X 
double 
 
X 方向せん断剛性倍率 
※(1) 
ratio_shear_Y 
double 
 
Y 方向せん断剛性倍率 
※(1) 
ratio_bending_X 
double 
 
X 軸回り曲げ剛性倍率 
※(1) 
ratio_bending_Y 
double 
 
Y 軸回り曲げ剛性倍率 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、1.0 とする。 
(2) 平面解析、疑似立体解析を採用する構造計算プログラムで指定する場合がある。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalColumnStiffness id="1" ratio_bending_X="1.2" ratio_bending_Y="1.2"/> 
<StbCalColumnStiffness id="2" ratio_bending_X="1.3" ratio_bending_Y="1.3"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
</StbCalData>

## Page 128

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
118 
 
4.4.18. 大梁剛性倍率：StbCalGirderStiffness 
・概要 
説明  
：大梁の剛性倍率 
親要素 
：StbCalMemberStiffnesses 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛性倍率ID 
 
guid 
string 
 
GUID 
 
ratio_shear_Y 
double 
 
Y 方向せん断剛性倍率 
※(1) 
ratio_shear_Z 
double 
 
Z 方向せん断剛性倍率 
※(1) 
ratio_bending_Y 
double 
 
Y 軸回り曲げ剛性倍率 
※(1) 
ratio_bending_Z 
double 
 
Z 軸回り曲げ剛性倍率 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、1.0 とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalColumnStiffness id="1" ratio_bending_X="1.2" ratio_bending_Y="1.2"/> 
<StbCalColumnStiffness id="2" ratio_bending_X="1.3" ratio_bending_Y="1.3"/> 
<StbCalGirderStiffness id="3" ratio_bending_Y="1.5"/> 
<StbCalGirderStiffness id="4" ratio_bending_Y="2.0"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
</StbCalData>

## Page 129

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
119 
 
4.4.19. 壁剛性倍率：StbCalWallStiffness 
・概要 
説明  
：壁の剛性倍率 
親要素 
：StbCalMemberStiffnesses 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛性倍率ID 
 
guid 
string 
 
GUID 
 
ratio_axial 
double 
 
軸剛性倍率 
※(1) 
ratio_shear 
double 
 
せん断剛性倍率 
※(1) 
ratio_bending 
double 
 
曲げ剛性倍率 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
壁エレメントモデルの置換柱に対する断面性能値の倍率を、直接指定する場合に指定する。 
 
(1) 省略時は、1.0 とする。

## Page 130

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
120 
 
4.4.20. ブレース剛性倍率：StbCalBraceStiffness 
・概要 
説明  
：ブレースの剛性倍率 
親要素 
：StbCalMemberStiffnesses 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛性倍率ID 
 
guid 
string 
 
GUID 
 
ratio_axial 
double 
 
軸剛性倍率 
※(1) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、1.0 とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalColumnStiffness id="1" ratio_bending_X="1.2" ratio_bending_Y="1.2"/> 
<StbCalColumnStiffness id="2" ratio_bending_X="1.3" ratio_bending_Y="1.3"/> 
<StbCalGirderStiffness id="3" ratio_bending_Y="1.5"/> 
<StbCalGirderStiffness id="4" ratio_bending_Y="2.0"/> 
<StbCalBraceStiffness id="5" ratio_axial="2.0"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
</StbCalData>

## Page 131

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
121 
 
4.4.21. 節点拘束条件（複数）：StbCalNodeRestrictions 
・概要 
説明  
：節点の拘束条件（支点剛、支点ピンまたは支点ばね）（複数） 
親要素 
：StbCalCondition 
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
StbCalNodeRestriction 
1 
制限無し 
節点拘束条件 
 
・補足 
無し

## Page 132

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
122 
 
4.4.22. 節点拘束条件：StbCalNodeRestriction 
・概要 
説明  
：節点の拘束条件（支点剛、支点ピンまたは支点ばね） 
親要素 
：StbCalNodeRestrictions 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
節点拘束条件ID 
 
guid 
string 
 
GUID 
 
e_X 
string 
 
X
―
方向拘束条件 以下から選択 
FIX  FREE  SPRING 
※(1) 
e_Y 
string 
 
Y
―
方向拘束条件 以下から選択 
FIX  FREE  SPRING 
※(1) 
e_Z 
string 
 
Z
―
方向拘束条件 以下から選択 
FIX  FREE  SPRING 
※(1) 
r_X 
string 
 
X
―
軸回り拘束条件 以下から選択 
FIX  FREE  SPRING 
※(1) 
r_Y 
string 
 
Y
―
軸回り拘束条件 以下から選択 
FIX  FREE  SPRING 
※(1) 
r_Z 
string 
 
Z
―
軸回り拘束条件 以下から選択 
FIX  FREE  SPRING 
※(1) 
e_spring_X 
double 
 
X
―
方向支点ばね定数(N/mm) 
※(2) 
e_spring_Y 
double 
 
Y
―
方向支点ばね定数(N/mm) 
※(2) 
e_spring_Z 
double 
 
Z
―
方向支点ばね定数(N/mm) 
※(2) 
r_spring_X 
double 
 
X
―
軸回り支点ばね定数(Nmm/rad) 
※(2) 
r_spring_Y 
double 
 
Y
―
軸回り支点ばね定数(Nmm/rad) 
※(2) 
r_spring_Z 
double 
 
Z
―
軸回り支点ばね定数(Nmm/rad) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 省略時は、FIX とする。 
(2) 材端条件がSPRING のとき、該当する方向のばね定数を指定する。

## Page 133

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
123 
 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalNodeRestrictions> 
<StbCalNodeRestriction id="1"/> 
<StbCalNodeRestriction id="2" e_X="FREE" e_Y="FREE" e_Z="FREE"/> 
<StbCalNodeRestriction id="3" e_X="SPRING" e_Y="SPRING" e_Z="FREE" 
                                 e_spring_X="1230000" e_spring_Y="1230000"/> 
</StbCalNodeRestrictions> 
 </StbCalCondition> 
</StbCalData>

## Page 134

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
124 
 
4.4.23. 接合部パネル（複数）：StbCalNodePanels 
・概要 
説明  
：接合部パネルの形状（複数） 
親要素 
：StbCalCondition 
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
StbCalNodePanel 
1 
制限無し 
接合部パネル 
 
・補足 
節点位置に配置される、柱梁接合部パネルの形状を、直接指定する場合に指定する。

## Page 135

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
125 
 
4.4.24. 接合部パネル：StbCalNodePanel 
・概要 
説明  
：接合部パネルの形状 
親要素 
：StbCalNodePanels 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
接合部パネルID 
 
guid 
string 
 
GUID 
 
B_X 
double 
 
X 方向幅(mm) 
※(1) 
D_X 
double 
 
X 方向せい(mm) 
※(1) 
t_X 
double 
 
X 方向厚さ(mm) 
※(1) 
B_Y 
double 
 
Y 方向幅(mm) 
※(1) 
D_Y 
double 
 
Y 方向せい(mm) 
※(1) 
t_Y 
double 
 
Y 方向厚さ(mm) 
※(1) 
G 
double 
○ 
せん断弾性係数(N/mm2) 
 
・内容 
無し 
・子要素 
無し 
・補足 
(1) 指定する場合は、X 方向、Y 方向それぞれについて、幅、せいおよび厚さの３つの値を必須と
する。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalNodePanels> 
<StbCalNodePanel  id="2" B_X="600" D_X="800" t_X="18" 
 B_Y="600" D_Y="800" t_Y="18" G="79000"/> 
</StbCalNodePanels> 
 </StbCalCondition> 
</StbCalData>

## Page 136

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
126 
 
4.4.25. 剛床条件（複数）：StbCalFloorDiaphragms 
・概要 
説明  
：剛床条件（複数） 
親要素 
：StbCalCondition 
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
StbCalFloorDiaphragm 
1 
制限無し 
剛床条件 
 
・補足 
無し

## Page 137

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
127 
 
4.4.26. 剛床条件：StbCalFloorDiaphragm 
・概要 
説明  
：剛床条件 
親要素 
：StbCalFloorDiaphragms 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛床条件ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
剛床名称 
※(1) 
id_floor_divided_area 
integer 
○ 
床区分ID(StbCalFloorDividedArea.id) 
※(2) 
・内容 
無し 
・子要素 
無し 
・補足 
剛床の指定は、剛床を構成する節点の集まりで指定する。節点の集まりは、「床区分」としてID
を設けて定義し、剛床条件では、該当する「床区分」を指定する（詳細は「床区分」の項を参照）。 
 
(1) 剛床に、プログラム独自で名前を付けたい場合に用いる。 
(2) 「床区分：StbCalFloorDividedArea」で定義される、床区分ID を指定する。 
・例 
<StbCalData> 
  <StbCalCommon> 
    <StbCalFloorDividedAreas> 
      <StbCalFloorDividedArea id="1" name="zone1" isInclude="true"> 
  <StbCalStoryDivided id_story="2" name_floor="zone1-2F">10 20</StbCalStoryDivided> 
     <StbCalStoryDivided id_story="3" name_floor="zone1-3F">11 21</StbCalStoryDivided> 
      </StbCalFloorDividedArea> 
    </StbCalFloorDividedAreas> 
  </StbCalCommon> 
   （略） 
  <StbCalCondition> 
<StbCalFloorDiaphragms> 
<StbCalFloorDiaphragm id="1" name="zone1" id_floor_divided_area ="1"/> 
</StbCalFloorDiaphragms> 
 </StbCalCondition> 
</StbCalData>

## Page 138

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
128 
 
4.5. 
荷重配置：StbCalLoadArrangements 
・概要 
説明  
：荷重を配置する部材ID および節点ID を指定 
親要素 
：StbCalData 
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
StbCalColumnFinish_RC_Arr 
0 
制限無し 
柱仕上げ（ＲＣ部材）配置 
 
StbCalColumnFinish_S_Arr 
0 
制限無し 
柱仕上げ（Ｓ部材）配置 
 
StbCalColumnMemberLoadArr 
0 
制限無し 
柱特殊荷重配置 
 
StbCalGirderFinish_RC_Arr 
0 
制限無し 
大梁仕上げ（ＲＣ部材）配置 
 
StbCalGirderFinish_S_Arr 
0 
制限無し 
大梁仕上げ（Ｓ部材）配置 
 
StbCalGirderMemberLoadArr 
0 
制限無し 
大梁特殊荷重配置 
 
StbCalBeamFinish_RC_Arr 
0 
制限無し 
小梁仕上げ（ＲＣ部材）配置 
 
StbCalBeamFinish_S_Arr 
0 
制限無し 
小梁仕上げ（Ｓ部材）配置 
 
StbCalBeamMemberLoadArr 
0 
制限無し 
小梁特殊荷重配置 
 
StbCalBraceFinish_S_Arr 
0 
制限無し 
ブレース仕上げ配置 
 
StbCalSlabLiveLoadArr 
0 
制限無し 
床積載荷重配置 
 
StbCalSlabFinish_RC_Arr 
0 
制限無し 
床スラブ仕上げ配置 
 
StbCalSlabAreaLoadArr 
0 
制限無し 
床面特殊荷重配置 
 
StbCalSlabPressureLoadArr 
0 
制限無し 
床土圧・水圧荷重配置 
 
StbCalWallFinish_RC_Arr 
0 
制限無し 
壁仕上げ配置 
 
StbCalWallAreaLoadArr 
0 
制限無し 
壁面特殊荷重配置 
 
StbCalWallPressureLoadArr 
0 
制限無し 
壁土圧・水圧荷重配置 
 
StbCalNodeWeightArr 
0 
制限無し 
節点追加・補正重量配置 
 
StbCalNodePointLoadArr 
0 
制限無し 
節点特殊荷重配置 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 139

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
129 
 
4.5.1. 柱仕上げ（ＲＣ部材）配置：StbCalColumnFinish_RC_Arr 
・概要 
説明  
：柱仕上げ（ＲＣ部材）を配置する柱部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalColumnFinish_RC_LoadList 
1 
1 
柱仕上げ（ＲＣ部材）配
置・荷重 
 
StbCalColumnFinish_RC_MemList 
0 
1 
柱仕上げ（ＲＣ部材）配
置・部材 
※(1) 
StbCalColumnFinish_RC_CalMemList 
0 
1 
柱仕上げ（ＲＣ部材）配
置・部材(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_RC> 
        <StbCalColumnFinish_RC id="101" type="2" weight="0.00095"/> 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalColumnFinish_RC_Arr> 
    <StbCalColumnFinish_RC_LoadList>101</StbCalColumnFinish_RC_LoadList> 
      <StbCalColumnFinish_RC_MemList>10 20 30</StbCalColumnFinish_RC_MemList> 
    </StbCalColumnFinish_RC_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 140

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
130 
 
4.5.2. 柱仕上げ（ＲＣ部材）配置・荷重：StbCalColumnFinish_RC_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalColumnFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnFinish_RC.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱仕上げ（ＲＣ部材）」の仕上げID を指定する。 
 
 
4.5.3. 柱仕上げ（ＲＣ部材）配置・部材：StbCalColumnFinish_RC_MemList 
・概要 
説明  
：柱仕上げ（ＲＣ部材）を配置する柱部材ID を指定 
親要素 
：StbCalColumnFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 141

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
131 
 
4.5.4. 柱仕上げ（ＲＣ部材）配置・部材(計算用・分割)：StbCalColumnFinish_RC_CalMemList 
・概要 
説明  
：柱仕上げ（ＲＣ部材）を配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 142

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
132 
 
4.5.5. 柱仕上げ（Ｓ部材）配置：StbCalColumnFinish_S_Arr 
・概要 
説明  
：柱仕上げ（Ｓ部材）を配置する柱部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalColumnFinish_S_LoadList 
1 
1 
柱仕上げ（Ｓ部材）配置・
荷重 
 
StbCalColumnFinish_S_MemList 
0 
1 
柱仕上げ（Ｓ部材）配置・
部材 
※(1) 
StbCalColumnFinish_S_CalMemList 
0 
1 
柱仕上げ（Ｓ部材）配置・
部材(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
         <StbCalColumnFinish_S id="102" type="1" weight="0.0006" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalColumnFinish_S_Arr> 
    <StbCalColumnFinish_S_LoadList>102</StbCalColumnFinish_S_LoadList> 
      <StbCalColumnFinish_S_MemList>11 21 31</StbCalColumnFinish_S_MemList> 
    </StbCalColumnFinish_S_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 143

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
133 
 
4.5.6. 柱仕上げ（Ｓ部材）配置・荷重：StbCalColumnFinish_S_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalColumnFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnFinish_S.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱仕上げ（Ｓ部材）」の仕上げID を指定する。 
 
 
4.5.7. 柱仕上げ（Ｓ部材）配置・部材：StbCalColumnFinish_S_MemList 
・概要 
説明  
：柱仕上げ（Ｓ部材）を配置する柱部材ID を指定 
親要素 
：StbCalColumnFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 144

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
134 
 
4.5.8. 柱仕上げ（Ｓ部材）配置・部材(計算用・分割)：StbCalColumnFinish_S_CalMemList 
・概要 
説明  
：柱仕上げ（Ｓ部材）を配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 145

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
135 
 
4.5.9. 柱特殊荷重配置：StbCalColumnMemberLoadArr 
・概要 
説明  
：部材特殊荷重を配置する柱部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalColumnMemberLoadList 
1 
1 
柱特殊荷重配置・荷重 
 
StbCalColumnMemberLoadMemList 
0 
1 
柱特殊荷重配置・部材 
※(1) 
StbCalColumnMemberLoadCalMemList 
0 
1 
柱特殊荷重配置・部材(計
算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalMemberLoad id="103" id_loadcase="1" loadcase="DL" type="1" 
  P1="10000" P2="1000" direction_load="X" description="Point load A"/> 
     <StbCalMemberLoad id="104" id_loadcase="1" loadcase="DL" type="1" 
  P1="20000" P2="2000" direction_load="X" description="Point load B"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalColumnMemberLoadArr> 
    <StbCalColumnMemberLoadList>103 104</StbCalColumnMemberLoadList> 
      <StbCalColumnMemberLoadMemList>12 22 32</StbCalColumnMemberLoadMemList> 
    </StbCalColumnMemberLoadArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 146

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
136 
 
4.5.10. 柱特殊荷重配置・荷重：StbCalColumnMemberLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalColumnMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalMemberLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「部材特殊荷重」の特殊荷重ID を指定する。 
 
 
4.5.11. 柱特殊荷重配置・部材：StbCalColumnMemberLoadMemList 
・概要 
説明  
：部材特殊荷重を配置する柱部材ID を指定 
親要素 
：StbCalColumnMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 147

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
137 
 
4.5.12. 柱特殊荷重配置・部材(計算用・分割)：StbCalColumnMemberLoadCalMemList 
・概要 
説明  
：部材特殊荷重を配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 148

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
138 
 
4.5.13. 大梁仕上げ（ＲＣ部材）配置：StbCalGirderFinish_RC_Arr 
・概要 
説明  
：大梁仕上げ（ＲＣ部材）を配置する大梁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalGirderFinish_RC_LoadList 
1 
1 
大梁仕上げ（ＲＣ部材）配
置・荷重 
 
StbCalGirderFinish_RC_MemList 
0 
1 
大梁仕上げ（ＲＣ部材）配
置・部材 
※(1) 
StbCalGirderFinish_RC_CalMemList 
0 
1 
大梁仕上げ（ＲＣ部材）配
置・部材(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_RC> 
        <StbCalGirderFinish_RC id="201" type="1" weight="0.0005"/> 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalGirderFinish_RC_Arr> 
    <StbCalGirderFinish_RC_LoadList>201</StbCalGirderFinish_RC_LoadList> 
      <StbCalGirderFinish_RC_MemList>40 50 60</StbCalGirderFinish_RC_MemList> 
    </StbCalGirderFinish_RC_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 149

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
139 
 
4.5.14. 大梁仕上げ（ＲＣ部材）配置・荷重：StbCalGirderFinish_RC_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalGirderFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderFinish_RC.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「梁仕上げ（ＲＣ部材）」の仕上げID を指定する。 
 
 
4.5.15. 大梁仕上げ（ＲＣ部材）配置・部材：StbCalGirderFinish_RC_MemList 
・概要 
説明  
：大梁仕上げ（ＲＣ部材）を配置する大梁部材ID を指定 
親要素 
：StbCalGirderFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 150

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
140 
 
4.5.16. 大梁仕上げ（ＲＣ部材）配置・部材(計算用・分割)：StbCalGirderFinish_RC_CalMemList 
・概要 
説明  
：大梁仕上げ（ＲＣ部材）を配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 151

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
141 
 
4.5.17. 大梁仕上げ（Ｓ部材）配置：StbCalGirderFinish_S_Arr 
・概要 
説明  
：大梁仕上げ（Ｓ部材）を配置する大梁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalGirderFinish_S_LoadList 
1 
1 
大梁仕上げ（Ｓ部材）配
置・荷重 
 
StbCalGirderFinish_S_MemList 
0 
1 
大梁仕上げ（Ｓ部材）配
置・部材 
※(1) 
StbCalGirderFinish_S_CalMemList 
0 
1 
大梁仕上げ（Ｓ部材）配
置・部材(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
         <StbCalGirderFinish_S id="202" type="1" weight="0.0006" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalGirderFinish_S_Arr> 
    <StbCalGirderFinish_S_LoadList>202</StbCalGirderFinish_S_LoadList> 
      <StbCalGirderFinish_S_MemList>41 51 61</StbCalGirderFinish_S_MemList> 
    </StbCalGirderFinish_S_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 152

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
142 
 
4.5.18. 大梁仕上げ（Ｓ部材）配置・荷重：StbCalGirderFinish_S_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalGirderFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderFinish_S.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「梁仕上げ（Ｓ部材）」の仕上げID を指定する。 
 
 
4.5.19. 大梁仕上げ（Ｓ部材）配置・部材：StbCalGirderFinish_S_MemList 
・概要 
説明  
：大梁仕上げ（Ｓ部材）を配置する大梁部材ID を指定 
親要素 
：StbCalGirderFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 153

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
143 
 
4.5.20. 大梁仕上げ（Ｓ部材）配置・部材(計算用・分割)：StbCalGirderFinish_S_CalMemList 
・概要 
説明  
：大梁仕上げ（Ｓ部材）を配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 154

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
144 
 
4.5.21. 大梁特殊荷重配置：StbCalGirderMemberLoadArr 
・概要 
説明  
：部材特殊荷重を配置する大梁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalGirderMemberLoadList 
1 
1 
大梁特殊荷重配置・荷重 
 
StbCalGirderMemberLoadMemList 
0 
1 
大梁特殊荷重配置・部材 
※(1) 
StbCalGirderMemberLoadCalMemList 
0 
1 
大梁特殊荷重配置・部材
(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalMemberLoad id="203" id_loadcase="1" loadcase="DL" type="4" 
  P1="5.0" description="Equipment load C"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalGirderMemberLoadArr> 
    <StbCalGirderMemberLoadList>203</StbCalGirderMemberLoadList> 
      <StbCalGirderMemberLoadMemList>42 52 62</StbCalGirderMemberLoadMemList> 
    </StbCalGirderMemberLoadArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 155

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
145 
 
4.5.22. 大梁特殊荷重配置・荷重：StbCalGirderMemberLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalGirderMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalMemberLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「部材特殊荷重」の特殊荷重ID を指定する。 
 
 
4.5.23. 大梁特殊荷重配置・部材：StbCalGirderMemberLoadMemList 
・概要 
説明  
：部材特殊荷重を配置する大梁部材ID を指定 
親要素 
：StbCalGirderMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 156

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
146 
 
4.5.24. 大梁特殊荷重配置・部材(計算用・分割)：StbCalGirderMemberLoadCalMemList 
・概要 
説明  
：部材特殊荷重を配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 157

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
147 
 
4.5.25. 小梁仕上げ（ＲＣ部材）配置：StbCalBeamFinish_RC_Arr 
・概要 
説明  
：小梁仕上げ（ＲＣ部材）を配置する小梁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalBeamFinish_RC_LoadList 
1 
1 
小梁仕上げ（ＲＣ部材）配
置・荷重 
 
StbCalBeamFinish_RC_MemList 
1 
1 
小梁仕上げ（ＲＣ部材）配
置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_RC> 
        <StbCalGirderFinish_RC id="301" type="1" weight="0.0005"/> 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalBeamFinish_RC_Arr> 
    <StbCalBeamFinish_RC_LoadList>301</StbCalBeamFinish_RC_LoadList> 
      <StbCalBeamFinish_RC_MemList>43 53 63</StbCalBeamFinish_RC_MemList> 
    </StbCalBeamFinish_RC_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 158

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
148 
 
4.5.26. 小梁仕上げ（ＲＣ部材）配置・荷重：StbCalBeamFinish_RC_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalBeamFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderFinish_RC.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「梁仕上げ（ＲＣ部材）」の仕上げID を指定する。 
 
 
4.5.27. 小梁仕上げ（ＲＣ部材）配置・部材：StbCalBeamFinish_RC_MemList 
・概要 
説明  
：小梁仕上げ（ＲＣ部材）を配置する小梁部材ID を指定 
親要素 
：StbCalBeamFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbBeam.id 
[monolist] 
integer 
○ 
小梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbBeam の小梁ID を指定する。

## Page 159

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
149 
 
4.5.28. 小梁仕上げ（Ｓ部材）配置：StbCalBeamFinish_S_Arr 
・概要 
説明  
：小梁仕上げ（Ｓ部材）を配置する小梁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalBeamFinish_S_LoadList 
1 
1 
小梁仕上げ（Ｓ部材）配
置・荷重 
 
StbCalBeamFinish_S_MemList 
1 
1 
小梁仕上げ（Ｓ部材）配
置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
         <StbCalGirderFinish_S id="302" type="1" weight="0.0006" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalBeamFinish_S_Arr> 
    <StbCalBeamFinish_S_LoadList>302</StbCalBeamFinish_S_LoadList> 
      <StbCalBeamFinish_S_MemList>44 54 64</StbCalBeamFinish_S_MemList> 
    </StbCalBeamFinish_S_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 160

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
150 
 
4.5.29. 小梁仕上げ（Ｓ部材）配置・荷重：StbCalBeamFinish_S_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalBeamFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderFinish_S.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「梁仕上げ（Ｓ部材）」の仕上げID を指定する。 
 
 
4.5.30. 小梁仕上げ（Ｓ部材）配置・部材：StbCalBeamFinish_S_MemList 
・概要 
説明  
：小梁仕上げ（Ｓ部材）を配置する小梁部材ID を指定 
親要素 
：StbCalBeamFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbBeam.id 
[monolist] 
integer 
○ 
小梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbBeam の小梁ID を指定する。

## Page 161

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
151 
 
4.5.31. 小梁特殊荷重配置：StbCalBeamMemberLoadArr 
・概要 
説明  
：部材特殊荷重を配置する小梁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalBeamMemberLoadList 
1 
1 
小梁特殊荷重配置・荷重 
 
StbCalBeamMemberLoadMemList 
1 
1 
小梁特殊荷重配置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalMemberLoad id="303" id_loadcase="1" loadcase="DL" type="4" 
  P1="5.0" description="Equipment load C"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalBeamMemberLoadArr> 
    <StbCalBeamMemberLoadList>303</StbCalBeamMemberLoadList> 
      <StbCalBeamMemberLoadMemList>45 55 65</StbCalBeamMemberLoadMemList> 
    </StbCalBeamMemberLoadArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 162

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
152 
 
4.5.32. 小梁特殊荷重配置・荷重：StbCalBeamMemberLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalBeamMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalMemberLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「部材特殊荷重」の特殊荷重ID を指定する。 
 
 
4.5.33. 小梁特殊荷重配置・部材：StbCalBeamMemberLoadMemList 
・概要 
説明  
：部材特殊荷重を配置する小梁部材ID を指定 
親要素 
：StbCalBeamMemberLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbBeam.id 
[monolist] 
integer 
○ 
小梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbBeam の小梁ID を指定する。

## Page 163

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
153 
 
4.5.34. ブレース仕上げ配置：StbCalBraceFinish_S_Arr 
・概要 
説明  
：ブレース仕上げを配置するブレース部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalBraceFinish_S_LoadList 
1 
1 
ブレース仕上げ配置・荷重  
StbCalBraceFinish_S_MemList 
1 
1 
ブレース仕上げ配置・部材  
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_S> 
        <StbCalBraceFinish_S id="401" type="1" weight="0.0005" covering_type="B" 
 covering_size="30" covering_unit_weight="1.35e-7"/> 
      </StbCalMemberFinishes_S> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalBraceFinish_S_Arr> 
    <StbCalBraceFinish_S_LoadList>401</StbCalBraceFinish_S_LoadList> 
      <StbCalBraceFinish_S_MemList>71 72 73</StbCalBraceFinish_S_MemList> 
    </StbCalBraceFinish_S_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 164

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
154 
 
4.5.35. ブレース仕上げ配置・荷重：StbCalBraceFinish_S_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalBraceFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalBraceFinish_S.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「ブレース仕上げ」の仕上げID を指定する。 
 
 
4.5.36. ブレース仕上げ配置・部材：StbCalBraceFinish_S_MemList 
・概要 
説明  
：ブレース仕上げを配置するブレース部材ID を指定 
親要素 
：StbCalBraceFinish_S_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbBrace.id 
[monolist] 
integer 
○ 
ブレースID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbBrace のブレースID を指定する。

## Page 165

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
155 
 
4.5.37. 床積載荷重配置：StbCalSlabLiveLoadArr 
・概要 
説明  
：床積載荷重を配置する床部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalSlabLiveLoadList 
1 
1 
床積載荷重配置・荷重 
 
StbCalSlabLiveLoadMemList 
1 
1 
床積載荷重配置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoadCondition> 
    <StbCalLiveloads> 
  <StbCalLiveload id="1" name="Room Load A" type="1"/> 
     </StbCalLiveloads> 
  </StbCalLoadCondition> 
 
  <StbCalLoadArrangements> 
    <StbCalSlabLiveLoadArr> 
    <StbCalSlabLiveLoadList>1</StbCalSlabLiveLoadList> 
      <StbCalSlabLiveLoadMemList>41 51 61</StbCalSlabLiveLoadMemList> 
    </StbCalSlabLiveLoadArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 166

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
156 
 
4.5.38. 床積載荷重配置・荷重：StbCalSlabLiveLoadList 
・概要 
説明  
：配置される積載荷重ID を指定 
親要素 
：StbCalSlabLiveLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalLiveLoad.id 
[monolist] 
integer 
○ 
積載荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「積載荷重」の積載荷重ID を指定する。 
 
 
4.5.39. 床積載荷重配置・部材：StbCalSlabLiveLoadMemList 
・概要 
説明  
：床積載荷重を配置する床部材ID を指定 
親要素 
：StbCalSlabLiveLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSlab.id 
[monolist] 
integer 
○ 
床ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSlab の床ID を指定する。

## Page 167

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
157 
 
4.5.40. 床スラブ仕上げ配置：StbCalSlabFinish_RC_Arr 
・概要 
説明  
：床スラブ仕上げを配置する床部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalSlabFinish_RC_LoadList 
1 
1 
床スラブ仕上げ配置・荷重  
StbCalSlabFinish_RC_MemList 
1 
1 
床スラブ仕上げ配置・部材  
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_RC> 
        <StbCalSlabFinish_RC id="502" type="1" weight="0.0005"/> 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalSlabFinish_RC_Arr> 
    <StbCalSlabFinish_RC_LoadList>502</StbCalSlabFinish_RC_LoadList> 
      <StbCalSlabFinish_RC_MemList>40 50 60</StbCalSlabFinish_RC_MemList> 
    </StbCalSlabFinish_RC_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 168

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
158 
 
4.5.41. 床スラブ仕上げ配置・荷重：StbCalSlabFinish_RC_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalSlabFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalSlabFinish_RC.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「床スラブ仕上げ」の仕上げID を指定する。 
 
 
4.5.42. 床スラブ仕上げ配置・部材：StbCalSlabFinish_RC_MemList 
・概要 
説明  
：床スラブ仕上げを配置する床部材ID を指定 
親要素 
：StbCalSlabFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSlab.id 
[monolist] 
integer 
○ 
床ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSlab の床ID を指定する。

## Page 169

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
159 
 
 
4.5.43. 床面特殊荷重配置：StbCalSlabAreaLoadArr 
・概要 
説明  
：部材特殊荷重を配置する床部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalSlabAreaLoadList 
1 
1 
床面特殊荷重配置・荷重 
 
StbCalSlabAreaLoadMemList 
1 
1 
床面特殊荷重配置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalAreaLoad id="503" id_loadcase="1" loadcase="DL" type="51" 
  P1="0.0008" description="additional load B"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalSlabAreaLoadArr> 
    <StbCalSlabAreaLoadList>503</StbCalSlabAreaLoadList> 
      <StbCalSlabAreaLoadMemList>42 52 62</StbCalSlabAreaLoadMemList> 
    </StbCalSlabMemberArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 170

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
160 
 
4.5.44. 床面特殊荷重配置・荷重：StbCalSlabAreaLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalSlabAreaLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalAreaLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「面特殊荷重」の特殊荷重ID を指定する。 
 
 
4.5.45. 床面特殊荷重配置・部材：StbCalSlabAreaLoadMemList 
・概要 
説明  
：面特殊荷重を配置する床部材ID を指定 
親要素 
：StbCalSlabAreaLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSlab.id 
[monolist] 
integer 
○ 
床ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSlab の床ID を指定する。

## Page 171

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
161 
 
4.5.46. 床土圧・水圧荷重配置：StbCalSlabPressureLoadArr 
・概要 
説明  
：土圧・水圧荷重を配置する床部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalSlabPressureLoadList 
1 
1 
床土圧・水圧荷重配置・荷重 
 
StbCalSlabPressureLoadMemList 
1 
1 
床土圧・水圧荷重配置・部材 
 
・補足 
無し

## Page 172

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
162 
 
4.5.47. 床土圧・水圧荷重配置・荷重：StbCalSlabPressureLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalSlabPressureLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalEarthHydrostaticPressureLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区
切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「土圧・水圧荷重」の特殊荷重ID を指定する。 
 
 
4.5.48. 床土圧・水圧荷重配置・部材：StbCalSlabPressureLoadMemList 
・概要 
説明  
：土圧・水圧荷重を配置する床部材ID を指定 
親要素 
：StbCalSlabPressureLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSlab.id 
[monolist] 
integer 
○ 
床ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSlab の床ID を指定する。

## Page 173

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
163 
 
4.5.49. 壁仕上げ配置：StbCalWallFinish_RC_Arr 
・概要 
説明  
：壁仕上げを配置する壁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalWallFinish_RC_LoadList 
1 
1 
壁仕上げ配置・荷重 
 
StbCalWallFinish_RC_MemList 
1 
1 
壁仕上げ配置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalFinish> 
      <StbCalMemberFinishes_RC> 
        <StbCalWallFinish_RC id="601" type="1" weight="0.0005"/> 
      </StbCalMemberFinishes_RC> 
    </StbCalFinish> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalWallFinish_RC_Arr> 
    <StbCalWallFinish_RC_LoadList>601</StbCalWallFinish_RC_LoadList> 
      <StbCalWallFinish_RC_MemList>44 54 64</StbCalWallFinish_RC_MemList> 
    </StbCalWallFinish_RC_Arr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 174

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
164 
 
4.5.50. 壁仕上げ配置・荷重：StbCalWallFinish_RC_LoadList 
・概要 
説明  
：配置される仕上げID を指定 
親要素 
：StbCalWallFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalWallFinish_RC.id 
[monolist] 
integer 
○ 
仕上げID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「壁仕上げ」の仕上げID を指定する。 
 
 
4.5.51. 壁仕上げ配置・部材：StbCalWallFinish_RC_MemList 
・概要 
説明  
：壁仕上げを配置する壁部材ID を指定 
親要素 
：StbCalWallFinish_RC_Arr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbWall.id 
[monolist] 
integer 
○ 
壁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbWall の壁ID を指定する。

## Page 175

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
165 
 
 
4.5.52. 壁面特殊荷重配置：StbCalWallAreaLoadArr 
・概要 
説明  
：部材特殊荷重を配置する壁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalWallAreaLoadList 
1 
1 
壁面特殊荷重配置・荷重 
 
StbCalWallAreaLoadMemList 
1 
1 
壁面特殊荷重配置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalAreaLoad id="602" id_loadcase="1" loadcase="DL" type="51" 
  P1="0.0008" description="additional load B"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalWallAreaLoadArr> 
    <StbCalWallAreaLoadList>602</StbCalWallAreaLoadList> 
      <StbCalWallAreaLoadMemList>45 55 65</StbCalWallAreaLoadMemList> 
    </StbCalWallMemberArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 176

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
166 
 
4.5.53. 壁面特殊荷重配置・荷重：StbCalWallAreaLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalWallAreaLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalAreaLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「面特殊荷重」の特殊荷重ID を指定する。 
 
 
4.5.54. 壁面特殊荷重配置・部材：StbCalWallAreaLoadMemList 
・概要 
説明  
：面特殊荷重を配置する壁部材ID を指定 
親要素 
：StbCalWallAreaLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbWall.id 
[monolist] 
integer 
○ 
壁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbWall の壁ID を指定する。

## Page 177

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
167 
 
4.5.55. 壁土圧・水圧荷重配置：StbCalWallPressureLoadArr 
・概要 
説明  
：土圧・水圧荷重を配置する壁部材ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalWallPressureLoadList 
1 
1 
壁土圧・水圧荷重配置・荷重 
 
StbCalWallPressureLoadMemList 
1 
1 
壁土圧・水圧荷重配置・部材 
 
・補足 
無し

## Page 178

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
168 
 
4.5.56. 壁土圧・水圧荷重配置・荷重：StbCalWallPressureLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalWallPressureLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalEarthHydrostaticPressureLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区
切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「土圧・水圧荷重」の特殊荷重ID を指定する。 
 
 
4.5.57. 壁土圧・水圧荷重配置・部材：StbCalWallPressureLoadMemList 
・概要 
説明  
：土圧・水圧荷重を配置する壁部材ID を指定 
親要素 
：StbCalWallPressureLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbWall.id 
[monolist] 
integer 
○ 
壁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbWall の壁ID を指定する。

## Page 179

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
169 
 
4.5.58. 節点追加・補正重量配置：StbCalNodeWeightArr 
・概要 
説明  
：節点追加・補正重量を配置する節点ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalNodeWeightLoadList 
1 
1 
節点追加・補正重量配置・
荷重 
 
StbCalNodeWeightNodeList 
1 
1 
節点追加・補正重量配置・
節点 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAddedWeights> 
     <StbCalNodeAddedWeight id="701" id_loadcase="1" loadcase="DL" weight ="1000"/> 
     <StbCalNodeAddedWeight id="702" id_loadcase="1" loadcase="DL" weight ="1200"/> 
</StbCalAddedWeights> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalNodeWeightArr> 
    <StbCalNodeWeightLoadList>701 702</StbCalNodeWeightLoadList> 
      <StbCalNodeWeightNodeList>81 82 83</StbCalNodeWeightNodeList> 
    </StbCalNodeWeightArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 180

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
170 
 
4.5.59. 節点追加・補正重量配置・荷重：StbCalNodeWeightLoadList 
・概要 
説明  
：配置される追加・補正重量ID を指定 
親要素 
：StbCalNodeWeightArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalNodeAddedWeight.id 
[monolist] 
integer 
○ 
追加・補正重量ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「節点追加・補正重量」の追加・補正重量ID を指定する。 
 
 
4.5.60. 節点追加・補正重量配置・節点：StbCalNodeWeightNodeList 
・概要 
説明  
：節点追加・補正重量を配置する節点ID を指定 
親要素 
：StbCalNodeWeightArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbNode.id 
[monolist] 
integer 
○ 
節点ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbNode の節点ID を指定する。

## Page 181

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
171 
 
4.5.61. 節点特殊荷重配置：StbCalNodePointLoadArr 
・概要 
説明  
：節点特殊荷重を配置する節点ID を指定 
親要素 
：StbCalLoadArrangements 
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
StbCalNodePointLoadList 
1 
1 
節点特殊荷重配置・荷重 
 
StbCalNodePointLoadNodeList 
1 
1 
節点特殊荷重配置・節点 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalLoad> 
    <StbCalAdditionalLoads> 
     <StbCalPointLoad id="703" id_loadcase="1" loadcase="DL" 
  P3="-1000" description="additional load C"/> 
</StbCalAdditionalLoads> 
  </StbCalLoad> 
 
  <StbCalLoadArrangements> 
    <StbCalNodePointLoadArr> 
    <StbCalNodePointLoadList>703</StbCalNodePointLoadList> 
      <StbCalNodePointLoadNodeList>84 85 86</StbCalNodePointLoadNodeList> 
    </StbCalWallMemberArr> 
  </StbCalLoadArrangements> 
</StbCalData>

## Page 182

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
172 
 
4.5.62. 節点特殊荷重配置・荷重：StbCalNodePointLoadList 
・概要 
説明  
：配置される特殊荷重ID を指定 
親要素 
：StbCalNodePointLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalPointLoad.id 
[monolist] 
integer 
○ 
特殊荷重ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「特殊荷重」の特殊荷重ID を指定する。 
 
 
4.5.63. 節点特殊荷重配置・節点：StbCalNodePointLoadNodeList 
・概要 
説明  
：特殊荷重を配置する節点ID を指定 
親要素 
：StbCalNodePointLoadArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbNode.id 
[monolist] 
integer 
○ 
節点ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbNode の節点ID を指定する。

## Page 183

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
173 
 
4.6. 
計算条件配置：StbCalConditionArrangements 
・概要 
説明  
：計算条件を配置する部材ID、節点ID および断面ID を指定 
親要素 
：StbCalData 
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
StbCalColumnConditionArr 
0 
制限無し 
柱材端条件配置 
 
StbCalColumnRigidzoneArr 
0 
制限無し 
柱剛域長さ配置 
 
StbCalColumnCriticalPositionArr 
0 
制限無し 
柱危険断面位置配置 
 
StbCalColumnStiffnessArr 
0 
制限無し 
柱剛性倍率配置 
 
StbCalGirderConditionArr 
0 
制限無し 
大梁材端条件配置 
 
StbCalGirderRigidzoneArr 
0 
制限無し 
大梁剛域長さ配置 
 
StbCalGirderCriticalPositionArr 
0 
制限無し 
大梁危険断面位置配置 
 
StbCalGirderStiffnessArr 
0 
制限無し 
大梁剛性倍率配置 
 
StbCalBraceStiffnessArr 
0 
制限無し 
ブレース剛性倍率配置 
 
StbCalWallStiffnessArr 
0 
制限無し 
壁剛性倍率配置 
 
StbCalNodeRestrictionArr 
0 
制限無し 
節点拘束条件配置 
 
StbCalNodePanelArr 
0 
制限無し 
節点接合部パネル配置 
 
StbCalColumnSecPropertyArr 
0 
制限無し 
柱断面性能配置 
 
StbCalGirderSecPropertyArr 
0 
制限無し 
大梁断面性能配置 
 
StbCalBraceSecPropertyArr 
0 
制限無し 
ブレース断面性能配置 
 
StbCalSlabSecPropertyArr 
0 
制限無し 
スラブ断面性能配置 
 
StbCalWallSecPropertyArr 
0 
制限無し 
壁断面性能配置 
 
・補足 
子要素の並びは、上表に示す順番としなければならない。

## Page 184

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
174 
 
4.6.1. 柱材端条件配置：StbCalColumnConditionArr 
・概要 
説明  
：柱材端条件を配置する柱部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalColumnConditionList 
1 
1 
柱材端条件配置・条件 
 
StbCalColumnConditionMemList 
0 
1 
柱材端条件配置・部材 
※(1) 
StbCalColumnConditionCalMemList 
0 
1 
柱材端条件配置・部材(計
算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberConditions> 
<StbCalColumnCondition id="101" 
 bottom_X="FIX" bottom_Y="PIN" top_X="FIX" top_Y="PIN"/> 
</StbCalMemberConditions> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalColumnConditionArr> 
    <StbCalColumnConditionList>101</StbCalColumnConditionList> 
      <StbCalColumnConditionMemList>12 22 32</StbCalColumnConditionMemList> 
    </StbCalColumnConditionArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 185

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
175 
 
4.6.2. 柱材端条件配置・条件：StbCalColumnConditionList 
・概要 
説明  
：配置される材端条件ID を指定 
親要素 
：StbCalColumnConditionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnCondition.id 
[monolist] 
integer 
○ 
材端条件ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱材端条件」の材端条件ID を指定する。 
 
 
4.6.3. 柱材端条件配置・部材：StbCalColumnConditionMemList 
・概要 
説明  
：柱材端条件を配置する柱部材ID を指定 
親要素 
：StbCalColumnConditionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 186

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
176 
 
4.6.4. 柱材端条件配置・部材(計算用・分割)：StbCalColumnConditionCalMemList 
・概要 
説明  
：柱材端条件を配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnConditionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 187

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
177 
 
4.6.5. 柱剛域長さ配置：StbCalColumnRigidzoneArr 
・概要 
説明  
：柱剛域長さを配置する柱部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalColumnRigidzoneList 
1 
1 
柱剛域長さ配置・条件 
 
StbCalColumnRigidzoneMemList 
0 
1 
柱剛域長さ配置・部材 
※(1) 
StbCalColumnRigidzoneCalMemList 
0 
1 
柱剛域長さ配置・部材(計
算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberRigidzones> 
<StbCalColumnRigidzone 
id="102" bottom_X="85" bottom_Y="100" top_X="85" top_Y="85"/> 
</StbCalMemberRigidzones> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalColumnRigidzoneArr> 
    <StbCalColumnRigidzoneList>102</StbCalColumnRigidzoneList> 
      <StbCalColumnRigidzoneMemList>12 22 32</StbCalColumnRigidzoneMemList> 
    </StbCalColumnRigidzoneArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 188

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
178 
 
4.6.6. 柱剛域長さ配置・条件：StbCalColumnRigidzoneList 
・概要 
説明  
：配置される剛域長さID を指定 
親要素 
：StbCalColumnRigidzoneArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnRigidzone.id 
[monolist] 
integer 
○ 
剛域長さID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱剛域長さ」の剛域長さID を指定する。 
 
 
4.6.7. 柱剛域長さ配置・部材：StbCalColumnRigidzoneMemList 
・概要 
説明  
：柱剛域長さを配置する柱部材ID を指定 
親要素 
：StbCalColumnRigidzoneArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 189

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
179 
 
4.6.8. 柱剛域長さ配置・部材(計算用・分割)：StbCalColumnRigidzoneCalMemList 
・概要 
説明  
：柱剛域長さを配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnRigidzoneArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 190

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
180 
 
4.6.9. 柱危険断面位置配置：StbCalColumnCriticalPositionArr 
・概要 
説明  
：柱危険断面位置を配置する柱部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalColumnCriticalPositionList 
1 
1 
柱危険断面位置配置・条件  
StbCalColumnCriticalPositionMemList 
0 
1 
柱危険断面位置配置・部材 ※(1) 
StbCalColumnCriticalPositionCalMemList 
0 
1 
柱危険断面位置配置・部材
(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberCriticalPositions> 
<StbCalColumnCriticalPosition 
id="103" bottom_X="85" bottom_Y="100" top_X="85" top_Y="85"/> 
</StbCalMemberCriticalPositions> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalColumnCriticalPositionArr> 
    <StbCalColumnCriticalPositionList>103</StbCalColumnCriticalPositionList> 
      <StbCalColumnCriticalPositionMemList>13 23 33</StbCalColumnCriticalPositionMemList> 
    </StbCalColumnCriticalPositionArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 191

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
181 
 
4.6.10. 柱危険断面位置配置・条件：StbCalColumnCriticalPositionList 
・概要 
説明  
：配置される危険断面位置ID を指定 
親要素 
：StbCalColumnCriticalPositionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnCriticalPosition.id 
[monolist] 
integer 
○ 
危険断面位置ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱危険断面位置」の危険断面位置ID を指定する。 
 
 
4.6.11. 柱危険断面位置配置・部材：StbCalColumnCriticalPositionMemList 
・概要 
説明  
：柱危険断面位置を配置する柱部材ID を指定 
親要素 
：StbCalColumnCriticalPositionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 192

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
182 
 
4.6.12. 柱危険断面位置配置・部材(計算用・分割)：StbCalColumnCriticalPositionCalMemList 
・概要 
説明  
：柱危険断面位置を配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnCriticalPositionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 193

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
183 
 
4.6.13. 柱剛性倍率配置：StbCalColumnStiffnessArr 
・概要 
説明  
：柱剛性倍率を配置する柱部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalColumnStiffnessList 
1 
1 
柱剛性倍率配置・条件 
 
StbCalColumnStiffnessMemList 
0 
1 
柱剛性倍率配置・部材 
※(1) 
StbCalColumnStiffnessCalMemList 
0 
1 
柱剛性倍率配置・部材(計
算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalColumnStiffness id="104" ratio_bending_X="1.3" ratio_bending_Y="1.3"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalColumnStiffnessArr> 
    <StbCalColumnStiffnessList>104</StbCalColumnStiffnessList> 
      <StbCalColumnStiffnessMemList>14 24 34</StbCalColumnStiffnessMemList> 
    </StbCalColumnStiffnessArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 194

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
184 
 
4.6.14. 柱剛性倍率配置・条件：StbCalColumnStiffnessList 
・概要 
説明  
：配置される剛性倍率ID を指定 
親要素 
：StbCalColumnStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnStiffness.id 
[monolist] 
integer 
○ 
剛性倍率ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱剛性倍率」の剛性倍率ID を指定する。 
 
 
4.6.15. 柱剛性倍率配置・部材：StbCalColumnStiffnessMemList 
・概要 
説明  
：柱剛性倍率を配置する柱部材ID を指定 
親要素 
：StbCalColumnStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
[monolist] 
integer 
○ 
柱ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbColumn の柱ID を指定する。

## Page 195

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
185 
 
4.6.16. 柱剛性倍率配置・部材(計算用・分割)：StbCalColumnStiffnessCalMemList 
・概要 
説明  
：柱剛性倍率を配置する柱部材ID(計算用・分割)を指定 
親要素 
：StbCalColumnStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
[monolist] 
integer 
○ 
柱(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalColumn の柱(計算用・分割)ID を指定する。

## Page 196

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
186 
 
4.6.17. 大梁材端条件配置：StbCalGirderConditionArr 
・概要 
説明  
：大梁材端条件を配置する大梁部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalGirderConditionList 
1 
1 
大梁材端条件配置・条件 
 
StbCalGirderConditionMemList 
0 
1 
大梁材端条件配置・部材 
※(1) 
StbCalGirderConditionCalMemList 
0 
1 
大梁材端条件配置・部材
(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberConditions> 
<StbCalGirderCondition id="201" start="PIN" end="PIN"/> 
</StbCalMemberConditions> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalGirderConditionArr> 
    <StbCalGirderConditionList>201</StbCalGirderConditionList> 
      <StbCalGirderConditionMemList>41 51 61</StbCalGirderConditionMemList> 
    </StbCalGirderConditionArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 197

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
187 
 
4.6.18. 大梁材端条件配置・条件：StbCalGirderConditionList 
・概要 
説明  
：配置される材端条件ID を指定 
親要素 
：StbCalGirderConditionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderCondition.id 
[monolist] 
integer 
○ 
材端条件ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「大梁材端条件」の材端条件ID を指定する。 
 
 
4.6.19. 大梁材端条件配置・部材：StbCalGirderConditionMemList 
・概要 
説明  
：大梁材端条件を配置する大梁部材ID を指定 
親要素 
：StbCalGirderConditionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 198

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
188 
 
4.6.20. 大梁材端条件配置・部材(計算用・分割)：StbCalGirderConditionCalMemList 
・概要 
説明  
：大梁材端条件を配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderConditionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 199

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
189 
 
4.6.21. 大梁剛域長さ配置：StbCalGirderRigidzoneArr 
・概要 
説明  
：大梁剛域長さを配置する大梁部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalGirderRigidzoneList 
1 
1 
大梁剛域長さ配置・条件 
 
StbCalGirderRigidzoneMemList 
0 
1 
大梁剛域長さ配置・部材 
※(1) 
StbCalGirderRigidzoneCalMemList 
0 
1 
大梁剛域長さ配置・部材
(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberRigidzones> 
<StbCalGirderRigidzone id="202" start="50" end="50"/> 
</StbCalMemberRigidzones> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalGirderRigidzoneArr> 
    <StbCalGirderRigidzoneList>202</StbCalGirderRigidzoneList> 
      <StbCalGirderRigidzoneMemList>42 52 62</StbCalGirderRigidzoneMemList> 
    </StbCalGirderRigidzoneArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 200

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
190 
 
4.6.22. 大梁剛域長さ配置・条件：StbCalGirderRigidzoneList 
・概要 
説明  
：配置される剛域長さID を指定 
親要素 
：StbCalGirderRigidzoneArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderRigidzone.id 
[monolist] 
integer 
○ 
剛域長さID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「大梁剛域長さ」の剛域長さID を指定する。 
 
 
4.6.23. 大梁剛域長さ配置・部材：StbCalGirderRigidzoneMemList 
・概要 
説明  
：大梁剛域長さを配置する大梁部材ID を指定 
親要素 
：StbCalGirderRigidzoneArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 201

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
191 
 
4.6.24. 大梁剛域長さ配置・部材(計算用・分割)：StbCalGirderRigidzoneCalMemList 
・概要 
説明  
：大梁剛域長さを配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderRigidzoneArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 202

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
192 
 
4.6.25. 大梁危険断面位置配置：StbCalGirderCriticalPositionArr 
・概要 
説明  
：大梁危険断面位置を配置する大梁部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalGirderCriticalPositionList 
1 
1 
大梁危険断面位置配置・条
件 
 
StbCalGirderCriticalPositionMemList 
0 
1 
大梁危険断面位置配置・部
材 
※(1) 
StbCalGirderCriticalPositionCalMemList 
0 
1 
大梁危険断面位置配置・部
材(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberCriticalPositions> 
<StbCalGirderCriticalPosition id="203" start="50" end="50"/> 
</StbCalMemberCriticalPositions> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalGirderCriticalPositionArr> 
    <StbCalGirderCriticalPositionList>203</StbCalGirderCriticalPositionList> 
      <StbCalGirderCriticalPositionMemList>43 53 63</StbCalGirderCriticalPositionMemList> 
    </StbCalGirderCriticalPositionArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 203

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
193 
 
4.6.26. 大梁危険断面位置配置・条件：StbCalGirderCriticalPositionList 
・概要 
説明  
：配置される危険断面位置ID を指定 
親要素 
：StbCalGirderCriticalPositionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderCriticalPosition.id 
[monolist] 
integer 
○ 
危険断面位置ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「大梁危険断面位置」の危険断面位置ID を指定する。 
 
 
4.6.27. 大梁危険断面位置配置・部材：StbCalGirderCriticalPositionMemList 
・概要 
説明  
：大梁危険断面位置を配置する大梁部材ID を指定 
親要素 
：StbCalGirderCriticalPositionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 204

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
194 
 
4.6.28. 大梁危険断面位置配置・部材(計算用・分割)：StbCalGirderCriticalPositionCalMemList 
・概要 
説明  
：大梁危険断面位置を配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderCriticalPositionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 205

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
195 
 
4.6.29. 大梁剛性倍率配置：StbCalGirderStiffnessArr 
・概要 
説明  
：大梁剛性倍率を配置する大梁部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalGirderStiffnessList 
1 
1 
大梁剛性倍率配置・条件 
 
StbCalGirderStiffnessMemList 
0 
1 
大梁剛性倍率配置・部材 
※(1) 
StbCalGirderStiffnessCalMemList 
0 
1 
大梁剛性倍率配置・部材
(計算用・分割) 
※(1) 
・補足 
(1) 部材または部材（計算用・分割）のうち、どちらかは必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalGirderStiffness id="204" ratio_bending_Y="1.5"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalGirderStiffnessArr> 
    <StbCalGirderStiffnessList>204</StbCalGirderStiffnessList> 
      <StbCalGirderStiffnessMemList>44 54 64</StbCalGirderStiffnessMemList> 
    </StbCalGirderStiffnessArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 206

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
196 
 
4.6.30. 大梁剛性倍率配置・条件：StbCalGirderStiffnessList 
・概要 
説明  
：配置される剛性倍率ID を指定 
親要素 
：StbCalGirderStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderStiffness.id 
[monolist] 
integer 
○ 
剛性倍率ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「大梁剛性倍率」の剛性倍率ID を指定する。 
 
 
4.6.31. 大梁剛性倍率配置・部材：StbCalGirderStiffnessMemList 
・概要 
説明  
：大梁剛性倍率を配置する大梁部材ID を指定 
親要素 
：StbCalGirderStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbGirder.id 
[monolist] 
integer 
○ 
大梁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbGirder の大梁ID を指定する。

## Page 207

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
197 
 
4.6.32. 大梁剛性倍率配置・部材(計算用・分割)：StbCalGirderStiffnessCalMemList 
・概要 
説明  
：大梁剛性倍率を配置する大梁部材ID(計算用・分割)を指定 
親要素 
：StbCalGirderStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirder.id 
[monolist] 
integer 
○ 
大梁(計算用・分割)ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) StbCalCommon で指定した、StbCalGirder の大梁(計算用・分割)ID を指定する。

## Page 208

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
198 
 
4.6.33. ブレース剛性倍率配置：StbCalBraceStiffnessArr 
・概要 
説明  
：ブレース剛性倍率を配置するブレース部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalBraceStiffnessList 
1 
1 
ブレース剛性倍率配置・条
件 
 
StbCalBraceStiffnessMemList 
1 
1 
ブレース剛性倍率配置・部
材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalBraceStiffness id="301" ratio_axial="2.0"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalBraceStiffnessArr> 
    <StbCalBraceStiffnessList>301</StbCalBraceStiffnessList> 
      <StbCalBraceStiffnessMemList>70 80 90</StbCalBraceStiffnessMemList> 
    </StbCalBraceStiffnessArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 209

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
199 
 
4.6.34. ブレース剛性倍率配置・条件：StbCalBraceStiffnessList 
・概要 
説明  
：配置される剛性倍率ID を指定 
親要素 
：StbCalBraceStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalBraceStiffness.id 
[monolist] 
integer 
○ 
剛性倍率ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「ブレース剛性倍率」の剛性倍率ID を指定する。 
 
 
4.6.35. ブレース剛性倍率配置・部材：StbCalBraceStiffnessMemList 
・概要 
説明  
：ブレース剛性倍率を配置するブレース部材ID を指定 
親要素 
：StbCalBraceStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbBrace.id 
[monolist] 
integer 
○ 
ブレースID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbBrace のブレースID を指定する。

## Page 210

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
200 
 
4.6.36. 壁剛性倍率配置：StbCalWallStiffnessArr 
・概要 
説明  
：壁剛性倍率を配置する壁部材ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalWallStiffnessList 
1 
1 
壁剛性倍率配置・条件 
 
StbCalWallStiffnessMemList 
1 
1 
壁剛性倍率配置・部材 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalMemberStiffnesses> 
<StbCalWallStiffness id="401" ratio_shear="1.1"/> 
</StbCalMemberStiffnesses> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalWallStiffnessArr> 
    <StbCalWallStiffnessList>401</StbCalWallStiffnessList> 
      <StbCalWallStiffnessMemList>72 82 92</StbCalWallStiffnessMemList> 
    </StbCalWallStiffnessArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 211

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
201 
 
4.6.37. 壁剛性倍率配置・条件：StbCalWallStiffnessList 
・概要 
説明  
：配置される剛性倍率ID を指定 
親要素 
：StbCalWallStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalWallStiffness.id 
[monolist] 
integer 
○ 
剛性倍率ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「壁剛性倍率」の剛性倍率ID を指定する。 
 
 
4.6.38. 壁剛性倍率配置・部材：StbCalWallStiffnessMemList 
・概要 
説明  
：壁剛性倍率を配置する壁部材ID を指定 
親要素 
：StbCalWallStiffnessArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbWall.id 
[monolist] 
integer 
○ 
壁ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbWall の壁ID を指定する。

## Page 212

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
202 
 
4.6.39. 節点拘束条件配置：StbCalNodeRestrictionArr 
・概要 
説明  
：節点拘束条件を配置する節点ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalNodeRestrictionList 
1 
1 
節点拘束条件配置・条件 
 
StbCalNodeRestrictionNodeList 
1 
1 
節点拘束条件配置・節点 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCondition> 
    <StbCalAddedRestrictions> 
<StbCalNodeRestriction id="801" e_X="FREE" e_Y="FREE" e_Z="FREE"/> 
</StbCalAddedRestrictions> 
  </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalNodeRestrictionArr> 
    <StbCalNodeRestrictionList>801</StbCalNodeRestrictionList> 
      <StbCalNodeRestrictionNodeList>111 112 113</StbCalNodeRestrictionNodeList> 
    </StbCalNodeRestrictionArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 213

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
203 
 
4.6.40. 節点拘束条件配置・条件：StbCalNodeRestrictionList 
・概要 
説明  
：配置される拘束条件ID を指定 
親要素 
：StbCalNodeRestrictionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalNodeRestriction.id 
[monolist] 
integer 
○ 
拘束条件ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「節点拘束条件」の拘束条件ID を指定する。 
 
 
4.6.41. 節点拘束条件配置・節点：StbCalNodeRestrictionNodeList 
・概要 
説明  
：節点拘束条件を配置する節点ID を指定 
親要素 
：StbCalNodeRestrictionArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbNode.id 
[monolist] 
integer 
○ 
節点ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbNode の節点ID を指定する。

## Page 214

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
204 
 
4.6.42. 節点接合部パネル配置：StbCalNodePanelArr 
・概要 
説明  
：接合部パネルを配置する節点ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalNodePanelList 
1 
1 
節点接合部パネル配置・条件 
 
StbCalNodePanelNodeList 
1 
1 
節点接合部パネル配置・節点 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalNodePanels> 
<StbCalNodePanel id="802" B_X="600" D_X="800" t_X="18" 
 B_Y="600" D_Y="800" t_Y="18" G="79000"/> 
</StbCalNodePanels> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalNodePanelArr> 
    <StbCalNodePanelList>802</StbCalNodePanelList> 
      <StbCalNodePanelNodeList>111 112 113</StbCalNodePanelNodeList> 
    </StbCalNodePanelArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 215

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
205 
 
4.6.43. 節点接合部パネル配置・条件：StbCalNodePanelList 
・概要 
説明  
：配置される接合部パネルID を指定 
親要素 
：StbCalNodePanelArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalNodePanel.id 
[monolist] 
integer 
○ 
接合部パネルID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「接合部パネル」の接合部パネルID を指定する。 
 
 
4.6.44. 節点接合部パネル配置・節点：StbCalNodePanelNodeList 
・概要 
説明  
：接合部パネルを配置する節点ID を指定 
親要素 
：StbCalNodePanelArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbNode.id 
[monolist] 
integer 
○ 
節点ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbNode の節点ID を指定する。

## Page 216

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
206 
 
4.6.45. 柱断面性能配置：StbCalColumnSecPropertyArr 
・概要 
説明  
：柱断面性能（数値指定）を配置する柱断面ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalColumnSecPropertyList 
1 
1 
柱断面性能配置・条件 
 
StbCalColumnSecProperty_RC_List 
0 
1 
柱断面性能配置・ＲＣ断面 
※(1) 
StbCalColumnSecProperty_S_List 
0 
1 
柱断面性能配置・Ｓ断面 
※(1) 
StbCalColumnSecProperty_SRC_List 
0 
1 
柱断面性能配置・ＳＲＣ断面 
※(1) 
StbCalColumnSecProperty_CFT_List 
0 
1 
柱断面性能配置・ＣＦＴ断面 
※(1) 
・補足 
(1) 断面種別の子要素のうち、最低１個は必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalColumnSectionProperty id="1001" A="11840" As_X ="9000" As_Y ="2700" 
I_X ="202000000" I_Y ="67500000" J="375000000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalColumnSecPropertyArr> 
    <StbCalColumnSecPropertyList>1001</StbCalColumnSecPropertyList> 
      <StbCalColumnSecProperty_RC_List>101 102 103</StbCalColumnSecProperty_RC_List> 
    </StbCalColumnSecPropertyArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 217

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
207 
 
4.6.46. 柱断面性能配置・条件：StbCalColumnSecPropertyList 
・概要 
説明  
：配置される断面性能ID を指定 
親要素 
：StbCalColumnSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumnSectionProperty.id 
[monolist] 
integer 
○ 
断面性能ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「柱断面性能（数値指定）」の断面性能ID を指定する。 
 
 
4.6.47. 柱断面性能配置・ＲＣ断面：StbCalColumnSecProperty_RC_List 
・概要 
説明  
：柱断面性能（数値指定）を配置するＲＣ柱断面ID を指定 
親要素 
：StbCalColumnSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecColumn_RC.id 
[monolist] 
integer 
○ 
RC 柱断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecColumn_RC の断面ID を指定する。

## Page 218

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
208 
 
4.6.48. 柱断面性能配置・Ｓ断面：StbCalColumnSecProperty_S_List 
・概要 
説明  
：柱断面性能（数値指定）を配置するＳ柱断面ID を指定 
親要素 
：StbCalColumnSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecColumn_S.id 
[monolist] 
integer 
○ 
S 柱断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecColumn_S の断面ID を指定する。 
 
 
4.6.49. 柱断面性能配置・ＳＲＣ断面：StbCalColumnSecProperty_SRC_List 
・概要 
説明  
：柱断面性能（数値指定）を配置するＳＲＣ柱断面ID を指定 
親要素 
：StbCalColumnSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecColumn_SRC.id 
[monolist] 
integer 
○ 
SRC 柱断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecColumn_SRC の断面ID を指定する。

## Page 219

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
209 
 
4.6.50. 柱断面性能配置・ＣＦＴ断面：StbCalColumnSecProperty_CFT_List 
・概要 
説明  
：柱断面性能（数値指定）を配置するＣＦＴ柱断面ID を指定 
親要素 
：StbCalColumnSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecColumn_CFT.id 
[monolist] 
integer 
○ 
CFT 柱断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecColumn_CFT の断面ID を指定する。

## Page 220

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
210 
 
4.6.51. 大梁断面性能配置：StbCalGirderSecPropertyArr 
・概要 
説明  
：大梁断面性能（数値指定）を配置する大梁断面ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalGirderSecPropertyList 
1 
1 
大梁断面性能配置・条件 
 
StbCalGirderSecProperty_RC_List 
0 
1 
大梁断面性能配置・ＲＣ断面 
※(1) 
StbCalGirderSecProperty_S_List 
0 
1 
大梁断面性能配置・Ｓ断面 
※(1) 
StbCalGirderSecProperty_SRC_List 
0 
1 
大梁断面性能配置・ＳＲＣ断面 
※(1) 
・補足 
(1) 断面種別の子要素のうち、最低１個は必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalGirderSectionProperty id="1002" A="6290" As_Y="3850" As_Z="2230" 
I_Y="135000000" I_Z="9840000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalGirderSecPropertyArr> 
    <StbCalGirderSecPropertyList>1002</StbCalGirderSecPropertyList> 
      <StbCalGirderSecProperty_RC_List>201 202 203</StbCalGirderSecProperty_RC_List> 
    </StbCalGirderSecPropertyArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 221

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
211 
 
4.6.52. 大梁断面性能配置・条件：StbCalGirderSecPropertyList 
・概要 
説明  
：配置される断面性能ID を指定 
親要素 
：StbCalGirderSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalGirderSectionProperty.id 
[monolist] 
integer 
○ 
断面性能ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「大梁断面性能（数値指定）」の断面性能ID を指定する。 
 
 
4.6.53. 大梁断面性能配置・ＲＣ断面：StbCalGirderSecProperty_RC_List 
・概要 
説明  
：大梁断面性能（数値指定）を配置するＲＣ大梁断面ID を指定 
親要素 
：StbCalGirderSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecGirder_RC.id 
[monolist] 
integer 
○ 
RC 大梁断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecGirder_RC の断面ID を指定する。

## Page 222

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
212 
 
4.6.54. 大梁断面性能配置・Ｓ断面：StbCalGirderSecProperty_S_List 
・概要 
説明  
：大梁断面性能（数値指定）を配置するＳ大梁断面ID を指定 
親要素 
：StbCalGirderSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecGirder_S.id 
[monolist] 
integer 
○ 
S 大梁断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecGirder_S の断面ID を指定する。 
 
 
4.6.55. 大梁断面性能配置・ＳＲＣ断面：StbCalGirderSecProperty_SRC_List 
・概要 
説明  
：大梁断面性能（数値指定）を配置するＳＲＣ大梁断面ID を指定 
親要素 
：StbCalGirderSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecGirder_SRC.id 
[monolist] 
integer 
○ 
SRC 大梁断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecGirder_SRC の断面ID を指定する。

## Page 223

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
213 
 
4.6.56. ブレース断面性能配置：StbCalBraceSecPropertyArr 
・概要 
説明  
：ブレース断面性能（数値指定）を配置するブレース断面ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalBraceSecPropertyList 
1 
1 
ブレース断面性能配置・条件 
 
StbCalBraceSecProperty_S_List 
1 
1 
ブレース断面性能配置・Ｓ断面 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalBraceSectionProperty id="1003" A="2160" i="25"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalBraceSecPropertyArr> 
    <StbCalBraceSecPropertyList>1003</StbCalBraceSecPropertyList> 
      <StbCalBraceSecProperty_S_List>301 302 303</StbCalBraceSecProperty_S_List> 
    </StbCalBraceSecPropertyArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 224

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
214 
 
4.6.57. ブレース断面性能配置・条件：StbCalBraceSecPropertyList 
・概要 
説明  
：配置される断面性能ID を指定 
親要素 
：StbCalBraceSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalBraceSectionProperty.id 
[monolist] 
integer 
○ 
断面性能ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「ブレース断面性能（数値指定）」の断面性能ID を指定する。 
 
 
4.6.58. ブレース断面性能配置・Ｓ断面：StbCalBraceSecProperty_S_List 
・概要 
説明  
：ブレース断面性能（数値指定）を配置するＳブレース断面ID を指定 
親要素 
：StbCalBraceSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecBrace_S.id 
[monolist] 
integer 
○ 
S ブレース断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecBrace_S の断面ID を指定する。

## Page 225

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
215 
 
4.6.59. スラブ断面性能配置：StbCalSlabSecPropertyArr 
・概要 
説明  
：スラブ断面性能（数値指定）を配置するスラブ断面ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalSlabSecPropertyList 
1 
1 
スラブ断面性能配置・条件 
 
StbCalSlabSecProperty_RC_List 
0 
1 
スラブ断面性能配置・ＲＣ断面 
※(1) 
StbCalSlabSecPropertyDeckList 
0 
1 
スラブ断面性能配置・デッキプレ
ートスラブ断面 
※(1) 
StbCalSlabSecPropertyPrecastList 
0 
1 
スラブ断面性能配置・既製スラブ
断面 
※(1) 
・補足 
(1) 断面種別の子要素のうち、最低１個は必須とする。 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalSlabSectionProperty id="1004" t="150.0" E="205000" G="79000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalSlabSecPropertyArr> 
    <StbCalSlabSecPropertyList>1004</StbCalSlabSecPropertyList> 
      <StbCalSlabSecProperty_RC_List>401 402 403</StbCalSlabSecProperty_RC_List> 
    </StbCalSlabSecPropertyArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 226

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
216 
 
4.6.60. スラブ断面性能配置・条件：StbCalSlabSecPropertyList 
・概要 
説明  
：配置される断面性能ID を指定 
親要素 
：StbCalSlabSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalSlabSectionProperty.id 
[monolist] 
integer 
○ 
断面性能ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「スラブ断面性能（数値指定）」の断面性能ID を指定する。 
 
 
4.6.61. スラブ断面性能配置・ＲＣ断面：StbCalSlabSecProperty_RC_List 
・概要 
説明  
：スラブ断面性能（数値指定）を配置するＲＣスラブ断面ID を指定 
親要素 
：StbCalSlabSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecSlab_RC.id 
[monolist] 
integer 
○ 
RC スラブ断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecSlab_RC の断面ID を指定する。

## Page 227

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
217 
 
4.6.62. スラブ断面性能配置・デッキプレートスラブ断面：StbCalSlabSecPropertyDeckList 
・概要 
説明  
：スラブ断面性能（数値指定）を配置するデッキプレートスラブ断面ID を指定 
親要素 
：StbCalSlabSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecSlabDeck.id 
[monolist] 
integer 
○ 
デッキプレートスラブ断面ID の列記（スペース区
切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecSlabDeck の断面ID を指定する。 
 
 
4.6.63. スラブ断面性能配置・既製スラブ断面：StbCalSlabSecPropertyPrecastList 
・概要 
説明  
：スラブ断面性能（数値指定）を配置する既製スラブ断面ID を指定 
親要素 
：StbCalSlabSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecSlabPrecast.id 
[monolist] 
integer 
○ 
既製スラブ断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecSlabPrecast の断面ID を指定する。

## Page 228

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
218 
 
4.6.64. 壁断面性能配置：StbCalWallSecPropertyArr 
・概要 
説明  
：壁断面性能（数値指定）を配置する壁断面ID を指定 
親要素 
：StbCalConditionArrangements 
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
StbCalWallSecPropertyList 
1 
1 
壁断面性能配置・条件 
 
StbCalWallSecProperty_RC_List 
1 
1 
壁断面性能配置・ＲＣ断面 
 
・補足 
無し 
・例 
<StbCalData> 
  <StbCalCondition> 
<StbCalSectionProperties> 
<StbCalWallSectionProperty id="1005" A="780000" As="650000" Iy="175700000000"/> 
</StbCalSectionProperties> 
 </StbCalCondition> 
 
  <StbCalConditionArrangements> 
    <StbCalWallSecPropertyArr> 
    <StbCalWallSecPropertyList>1005</StbCalWallSecPropertyList> 
      <StbCalWallSecProperty_RC_List>501 502 503</StbCalWallSecProperty_RC_List> 
    </StbCalWallSecPropertyArr> 
  </StbCalConditionArrangements> 
</StbCalData>

## Page 229

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
219 
 
4.6.65. 壁断面性能配置・条件：StbCalWallSecPropertyList 
・概要 
説明  
：配置される断面性能ID を指定 
親要素 
：StbCalWallSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalWallSectionProperty.id 
[monolist] 
integer 
○ 
断面性能ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 「壁断面性能（数値指定）」の断面性能ID を指定する。 
 
 
4.6.66. 壁断面性能配置・ＲＣ断面：StbCalWallSecProperty_RC_List 
・概要 
説明  
：壁断面性能（数値指定）を配置するＲＣ壁断面ID を指定 
親要素 
：StbCalWallSecPropertyArr 
・属性 
無し 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecWall_RC.id 
[monolist] 
integer 
○ 
RC 壁断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) ST-Bridge 本編における、StbSecWall_RC の断面ID を指定する。

## Page 230

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
220 
 
5. StbAnaModels 要素リファレンス 
 
5.1. 
解析モデル（複数）：StbAnaModels 
・概要 
説明  
：解析モデル（複数） 
親要素 
：ST_BRIDGE 
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
StbAnaModel 
0 
制限無し 
解析モデル 
 
・補足 
無し 
・例 
<StbAnaModels> 
  <StbAnaModel id="1"> 
    （略） 
  </StbAnaModel> 
  <StbAnaModel id="2"> 
    （略） 
  </StbAnaModel> 
</StbAnaModels>

## Page 231

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
221 
 
5.2. 
解析モデル：StbAnaModel 
・概要 
説明  
：解析モデル 
親要素 
：StbAnaModels 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析モデルID 
 
guid 
string 
 
GUID 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaNodes 
0 
1 
解析用節点（複数） 
 
StbAnaStories 
0 
1 
解析用階（複数） 
 
StbAnaMembers 
0 
1 
解析用部材 
 
StbAnaProperties 
0 
1 
解析用部材性能 
 
StbAnaFloorDiaphragms 
0 
1 
剛床指定（複数） 
 
StbAnaMaterials 
0 
1 
材料性能（複数） 
 
StbAnaSections 
0 
1 
断面性能（複数） 
 
StbAnaLoadCases 
0 
1 
荷重ケース（複数） 
 
StbAnaAnalyses 
0 
1 
解析ケース（複数） 
 
StbAnaRelations 
0 
1 
解析モデル配置 
 
・補足 
無し 
・例 
<StbAnaModel id="1"> 
<StbAnaNodes>（略）</StbAnaNodes> 
<StbAnaStories>（略）</StbAnaStories> 
<StbAnaMembers>（略）</StbAnaMembers> 
<StbAnaProperties>（略）</StbAnaProperties> 
<StbAnaMaterials>（略）</StbAnaMaterials> 
<StbAnaSections>（略）</StbAnaSections> 
<StbAnaLoadCases>（略）</StbAnaLoadCases> 
<StbAnaAnalyses>（略）</StbAnaAnalyses> 
<StbAnaRelations>（略）</StbAnaRelations> 
</StbAnaModel>

## Page 232

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
222 
 
5.2.1. 解析用節点（複数）：StbAnaNodes 
・概要 
説明  
：解析用節点（複数） 
親要素 
：StbAnaModel 
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
StbAnaNode 
0 
制限無し 
解析用節点 
 
・補足 
無し 
・例 
<StbAnaNodes> 
<StbAnaNode Id="1" x="0" y="4000" z="3500"> 
   <StbAnaBoundary x="FIX" y="FIX" z="FIX"/> 
  </StbAnaNode> 
</StbAnaNodes>

## Page 233

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
223 
 
5.2.2. 解析用節点：StbAnaNode 
・概要 
説明  
：解析用節点 
親要素 
：StbAnaNodes 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用節点ID 
 
guid 
string 
 
GUID 
 
x 
double 
○ 
全体座標系X 値 
 
y 
double 
○ 
全体座標系Y 値 
 
z 
double 
○ 
全体座標系Z 値 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaBoundary 
0 
1 
境界条件 
 
・補足 
無し 
・例 
<StbAnaNodes> 
<StbAnaNode Id="1" x="0" y="4000" z="3500"> 
   <StbAnaBoundary x="FIX" y="FIX" z="FIX"/> 
  </StbAnaNode> 
</StbAnaNodes>

## Page 234

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
224 
 
5.2.3. 境界条件：StbAnaBoundary 
・概要 
説明  
：境界条件 
親要素 
：StbAnaNode 
・属性 
属性名 
型 
必須 
説明 
補足 
x 
string 
 
X 方向拘束条件 
以下のいずれかの値をとる。 
FIX、FREE（以下同じ） 
 
y 
string 
 
Y 方向拘束条件 
 
z 
string 
 
Z 方向拘束条件 
 
tx 
string 
 
X 軸回り拘束条件 
 
ty 
string 
 
Y 軸回り拘束条件 
 
tz 
string 
 
Z 軸回り拘束条件 
 
・内容 
無し 
・子要素 
無し 
・補足 
各拘束条件属性について、省略した場合はFREE とする。全ての拘束条件がFREE の場合は、この
要素を省略してよい。 
・例 
<StbAnaBoundary x="FIX" y="FIX" z="FIX" tx="FREE" ty="FREE" tz="FREE"/>

## Page 235

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
225 
 
5.2.4. 解析用階（複数）：StbAnaStories 
・概要 
説明  
：解析用階（複数） 
親要素 
：StbAnaModel 
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
StbAnaStory 
0 
制限無し 
解析用階 
 
・補足 
無し 
・例 
<StbAnaStories> 
<StbAnaStory id="1" name="1F" > 
    （略） 
</StbAnaStory> 
</StbAnaStories>

## Page 236

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
226 
 
5.2.5. 解析用階：StbAnaStory 
・概要 
説明  
：解析用階 
親要素 
：StbAnaStories 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用階ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
階名称 
 
id_node_lower 
integer 
 
層間変形角計算用解析節点ID(下側) 
 
id_node_upper 
integer 
 
層間変形角計算用解析節点ID(上側) 
 
height 
double 
 
層間変形角計算用高さ 
 
sum_weight 
double 
 
層せん断力係数計算用上層階重量の
合計 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaMemberid_List 
1 
1 
解析用部材ID リスト 
 
 
・補足 
StbAnaMemberid_List には階に属する部材の集計を指定する。大梁が上階、下階もしくは両方に
含まれるかどうかなどはアプリケーション依存とする。 
・例 
<StbAnaStory id="1" name="1F" id_node_lower="1" id_node_upper="9" sum_weight="45000"> 
  <StbAnaMemberid_List> 
    （略） 
  </StbAnaMemberid_List> 
</StbAnaStory>

## Page 237

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
227 
 
5.2.6. 解析用部材ID リスト：StbAnaMemberid_List 
・概要 
説明  
：解析用部材ＩＤリスト 
親要素 
：StbAnaStory 
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
StbAnaMemberid 
0 
制限無し 
解析用部材ID 
 
・補足 
無し 
・例 
<StbAnaMemberid_List> 
<StbAnaMemberid id="10" kind="StbAnaBeam"/> 
  <StbAnaMemberid id="11" kind="StbAnaBeam"/> 
  <StbAnaMemberid id="12" kind="StbAnaBeam"/> 
</StbAnaMemberid_List>

## Page 238

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
228 
 
5.2.7. 解析用部材ID：StbAnaMemberid 
・概要 
説明  
：解析用部材ＩＤ 
親要素 
：StbAnaMemberid_List 
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
 
kind 
string 
○ 
解析用部材種類 
以下のいずれかの値を取るものとする 
StbAnaBeam,StbAnaTruss, 
StbAnaSupport,StbAnaSpring, 
StbAnaWall,StbAnaPlaneTriangle, 
StbAnaPlaneRectangle,StbAnaNodePanel 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaMemberid_List> 
<StbAnaMemberid id="10" kind="StbAnaBeam"/> 
  <StbAnaMemberid id="11" kind="StbAnaBeam"/> 
  <StbAnaMemberid id="12" kind="StbAnaBeam"/> 
</StbAnaMemberid_List>

## Page 239

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
229 
 
5.2.8. 解析用部材：StbAnaMembers 
・概要 
説明  
：解析用部材 
親要素 
：StbAnaModel 
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
StbAnaBeams 
0 
1 
梁要素（複数） 
 
StbAnaTrusses 
0 
1 
トラス要素（複数） 
 
StbAnaSupports 
0 
1 
支点ばね要素（複数） 
 
StbAnaSprings 
0 
1 
ばね要素（複数） 
 
StbAnaWalls 
0 
1 
壁エレメント要素（複数） 
 
StbAnaPlaneTriangles 
0 
1 
平面要素（三角形、複数） 
 
StbAnaPlaneRectangles 
0 
1 
平面要素（四角形、複数） 
 
StbAnaNodePanels 
0 
1 
接合部パネル要素（複数） 
 
・補足 
無し 
・例 
<StbAnaMembers> 
<StbAnaBeams>（略）</StbAnaBeams> 
<StbAnaTrusses>（略）</StbAnaTrusses> 
<StbAnaSupports>（略）</StbAnaSupports> 
<StbAnaSprings>（略）</StbAnaSprings> 
<StbAnaWalls>（略）</StbAnaWalls> 
<StbAnaPlaneRectangles>（略）</StbAnaPlaneRectangles> 
</StbAnaMembers>

## Page 240

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
230 
 
5.2.9. 梁要素（複数）：StbAnaBeams 
・概要 
説明  
：梁要素（複数） 
親要素 
：StbAnaMembers 
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
StbAnaBeam 
0 
制限無し 
梁要素 
 
・補足 
無し 
・例 
<StbAnaMembers> 
<StbAnaBeams> 
<StbAnaBeam id="3" name="GI001" id_property="4" 
                          id_node_start="12" id_node_end="13" coord_angle="0.0"/> 
<StbAnaBeam id="33" name="CO001" id_property="14" 
                          id_node_start="12" id_node_end="22" coord_angle="0.0"/> 
</StbAnaBeams> 
</StbAnaMembers>

## Page 241

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
231 
 
5.2.10. 梁要素：StbAnaBeam 
・概要 
説明  
：梁要素 
親要素 
：StbAnaBeams 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
梁要素名称 
 
id_property 
integer 
○ 
梁要素性能StbAnaBeamProperty の
ID 
 
id_node_start 
integer 
○ 
始端解析用節点ID 
 
id_node_end 
integer 
○ 
終端解析用節点ID 
 
coord_angle 
double 
 
回転角度 
材軸回りの回転角度 
※(1) 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaBeamRigidzone 
0 
1 
梁要素剛域長さ 
 
StbAnaBeamCriticalPosition 
0 
1 
梁要素危険断面位置 
 
StbAnaBeamEndSpring 
0 
1 
梁要素材端ばね 
 
・補足 
(1) 省略された場合は、0.0 度とする。 
・例 
<StbAnaMembers> 
<StbAnaBeams> 
<StbAnaBeam id="3" name="GI001" id_property="4" 
                          id_node_start="12" id_node_end="13" coord_angle="0.0"/> 
<StbAnaBeam id="33" name="CO001" id_property="14" 
                          id_node_start="12" id_node_end="22" coord_angle="0.0"/> 
</StbAnaBeams> 
</StbAnaMembers>

## Page 242

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
232 
 
5.2.11. 梁要素剛域長さ：StbAnaBeamRigidzone 
・概要 
説明  
：梁要素の剛域長 
親要素 
：StbAnaBeam 
・属性 
属性名 
型 
必須 
説明 
補足 
start_x 
double 
 
始端軸方向 
 
start_y 
double 
 
始端Y 方向(Z 軸回り) 
 
start_z 
double 
 
始端Z 方向(Y 軸回り) 
 
end_x 
double 
 
終端軸方向 
 
end_y 
double 
 
終端Y 方向(Z 軸回り) 
 
end_z 
double 
 
終端Z 方向(Y 軸回り) 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
 
・例 
<StbAnaBeam id="3" name="GI001" id_property="4" id_node_start="12" id_node_end="13"> 
<StbAnaBeamRigidzone start_y="500.0" end_y="500.0"/> 
</StbAnaBeam>

## Page 243

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
233 
 
5.2.12. 梁要素危険断面位置：StbAnaBeamCriticalPosition 
・概要 
説明  
：梁要素の危険断面位置 
親要素 
：StbAnaBeam 
・属性 
属性名 
型 
必須 
説明 
補足 
start_x 
double 
 
始端軸方向 
 
start_y 
double 
 
始端Y 方向(Z 軸回り) 
 
start_z 
double 
 
始端Z 方向(Y 軸回り) 
 
end_x 
double 
 
終端軸方向 
 
end_y 
double 
 
終端Y 方向(Z 軸回り) 
 
end_z 
double 
 
終端Z 方向(Y 軸回り) 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
 
・例 
<StbAnaBeam id="3" name="GI001" id_property="4" id_node_start="12" id_node_end="13"> 
<StbAnaBeamCriticalPosition start_y="500.0" end_y="500.0"/> 
</StbAnaBeam>

## Page 244

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
234 
 
5.2.13. 梁要素材端ばね：StbAnaBeamEndSpring 
・概要 
説明  
：材端ばね 
親要素 
：StbAnaBeam 
・属性 
属性名 
型 
必須 
説明 
補足 
id_property_start_y 
integer 
 
始端Y 軸回りばね要素性能ID 
 
id_property_start_z 
integer 
 
始端Z 軸回りばね要素性能ID 
 
id_property_end_y 
integer 
 
終端Y 軸回りばね要素性能ID 
 
id_property_end_z 
integer 
 
終端Z 軸回りばね要素性能ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
材端ピンの場合は、StbAnaBeamEndSpring でばね値0 のばね性能を設定することで表現する。 
・例 
<StbAnaBeam id="3" name="GI001" id_property="4" id_node_start="12" id_node_end="13"> 
<StbAnaBeamEndSpring id_property_start_y="10" id_property_end_y="10"/> 
</StbAnaBeam>

## Page 245

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
235 
 
5.2.14. トラス要素（複数）：StbAnaTrusses 
・概要 
説明  
：トラス要素（複数） 
親要素 
：StbAnaMembers 
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
StbAnaTruss 
0 
制限無し 
トラス要素 
 
・補足 
無し 
・例 
<StbAnaMembers> 
<StbAnaTrusses> 
<StbAnaTruss id="4" name="BR1" id_property="2" id_node_start="8" id_node_end="13"/> 
</StbAnaTrusses> 
</StbAnaMembers>

## Page 246

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
236 
 
5.2.15. トラス要素：StbAnaTruss 
・概要 
説明  
：トラス要素 
親要素 
：StbAnaTruss 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
トラス要素名称 
 
id_property 
integer 
○ 
トラス要素性能
StbAnaTrussProperty のID 
 
id_node_start 
integer 
○ 
始端解析用節点ID 
 
id_node_end 
integer 
○ 
終端解析用節点ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaMembers> 
<StbAnaTrusses> 
<StbAnaTruss id="4" name="BR1" id_property="2" id_node_start="8" id_node_end="13"/> 
</StbAnaTrusses> 
</StbAnaMembers>

## Page 247

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
237 
 
5.2.16. 支点ばね要素（複数）：StbAnaSupports 
・概要 
説明  
：支点ばね要素（複数） 
親要素 
：StbAnaMembers 
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
StbAnaSupport 
0 
制限無し 
支点ばね要素 
 
・補足 
無し 
・例 
<StbAnaSupports> 
<StbAnaSupport id="6" name="SUP1" id_property="6" id_node="3" direction="UZ"/> 
</StbAnaSupports>

## Page 248

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
238 
 
5.2.17. 支点ばね要素：StbAnaSupport 
・概要 
説明  
：支点ばね要素 
親要素 
：StbAnaSupports 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
支点ばね要素名称 
 
id_property 
integer 
○ 
ばね要素性能 
StbAnaSpringProperty のID 
 
id_node 
integer 
○ 
解析用節点ID 
 
direction 
string 
○ 
ばね方向 
以下のいずれかの値を取るものとする 
UX,UY,UZ,TX,TY,TZ 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaSupports> 
<StbAnaSupport id="6" name="SUP1" id_property="6" id_node="3" direction="UZ"/> 
</StbAnaSupports>

## Page 249

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
239 
 
5.2.18. ばね要素（複数）：StbAnaSprings 
・概要 
説明  
：ばね要素（複数） 
親要素 
：StbAnaMembers 
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
StbAnaSpring 
0 
制限無し 
ばね要素 
 
・補足 
無し 
・例 
<StbAnaSprings> 
<StbAnaSpring id="7" name="SPR1" id_property="7" 
                                   id_node_start="3" id_node_end="4" direction="TY"/> 
</StbAnaSprings>

## Page 250

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
240 
 
5.2.19. ばね要素：StbAnaSpring 
・概要 
説明  
：ばね要素 
親要素 
：StbAnaSprings 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
ばね要素名称 
 
id_property 
integer 
○ 
ばね要素性能 
StbAnaSpringProperty のID 
 
id_node_start 
integer 
○ 
始端解析用節点ID 
 
id_node_end 
integer 
○ 
終端解析用節点ID 
 
direction 
string 
○ 
ばね方向 
以下のいずれかの値を取るものと
する。 
UX,UY,UZ,TX,TY,TZ，AXIAL 
AXIAL：軸方向ば
ね 
・内容 
無し 
・子要素 
無し 
・補足 
id_node_start とid_node_end が同一座標に存在する場合、direction は必須とする。 
・例 
<StbAnaSprings> 
<StbAnaSpring id="7" name="SPR1" id_property="7" 
                                   id_node_start="3" id_node_end="4" direction="TY"/> 
</StbAnaSprings>

## Page 251

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
241 
 
5.2.20. 壁エレメント要素（複数）：StbAnaWalls 
・概要 
説明  
：壁エレメント要素（複数） 
親要素 
：StbAnaMembers 
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
StbAnaWall 
0 
制限無し 
壁エレメント要素 
 
・補足 
無し 
・例 
<StbAnaWalls> 
<StbAnaWall id="8" name="W1" id_property="8" 
 id_node_start_bottom="3" id_node_end_bottom="4" 
 id_node_start_top="14" id_node_end_top="15"/> 
</StbAnaWalls>

## Page 252

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
242 
 
5.2.21. 壁エレメント要素：StbAnaWall 
・概要 
説明  
：壁エレメント要素 
親要素 
：StbAnaWalls 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
壁エレメント要素名称 
 
id_property 
integer 
○ 
壁エレメント要素性能 
StbAnaWallProperty のID 
 
id_node_start_bottom 
integer 
○ 
壁脚始端側節点ID 
 
id_node_end_bottom 
integer 
○ 
壁脚終端側節点ID 
 
id_node_start_top 
integer 
○ 
壁頭始端側節点ID 
 
id_node_end_top 
integer 
○ 
壁頭終端側節点ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaWalls> 
<StbAnaWall id="8" name="W1" id_property="8" 
 id_node_start_bottom="3" id_node_end_bottom="4" 
 id_node_start_top="14" id_node_end_top="15"/> 
</StbAnaWalls>

## Page 253

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
243 
 
5.2.22. 平面要素（三角形、複数）： StbAnaPlaneTriangles 
・概要 
説明  
：平面要素（三角形、複数） 
親要素 
：StbAnaMembers 
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
StbAnaPlaneTriangle 
0 
制限無し 
平面要素(三角形) 
 
・補足 
無し 
・例 
<StbAnaPlaneTriangles> 
<StbAnaPlaneTriangle id="9" name="W3" id_property="9" 
 id_node1="3" id_node2="4" id_node3="14"/> 
</StbAnaPlaneTriangles>

## Page 254

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
244 
 
5.2.23. 平面要素（三角形）：StbAnaPlaneTriangle 
・概要 
説明  
：平面要素（三角形） 
親要素 
：StbAnaPlaneTriangles 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
平面要素名称 
 
id_property 
integer 
○ 
平面要素性能
StbAnaPlaneProperty のID 
 
id_node1 
integer 
○ 
構成解析用節点ID 
 
id_node2 
integer 
○ 
構成解析用節点ID 
 
id_node3 
integer 
○ 
構成解析用節点ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
節点１→節点２を要素内ｘ方向として、節点３をｘ－ｙ平面内として定義する。 
・例 
<StbAnaPlaneTriangles> 
<StbAnaPlaneTriangle id="9" name="W3" id_property="9" 
 id_node1="3" id_node2="4" id_node3="14"/> 
</StbAnaPlaneTriangles>

## Page 255

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
245 
 
5.2.24. 平面要素（四角形、複数）：StbAnaPlaneRectangles 
・概要 
説明  
：平面要素（四角形、複数） 
親要素 
：StbAnaMembers 
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
StbAnaPlaneRectangle 
0 
制限無し 
平面要素(四角形) 
 
・補足 
無し 
・例 
<StbAnaPlaneRectangles> 
<StbAnaPlaneRectangle id="10" name="W4" id_property="10" 
 id_node1="5" id_node2="6" id_node3="16"  id_node4="15"/> 
</StbAnaPlaneRectangles>

## Page 256

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
246 
 
5.2.25. 平面要素（四角形）：StbAnaPlaneRectangle 
・概要 
説明  
：平面要素（四角形） 
親要素 
：StbAnaPlaneRectangles 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
平面要素名称 
 
id_property 
integer 
○ 
平面要素性能 
StbAnaPlaneProperty のID 
 
id_node1 
integer 
○ 
構成解析用節点ID 
 
id_node2 
integer 
○ 
構成解析用節点ID 
 
id_node3 
integer 
○ 
構成解析用節点ID 
 
id_node4 
integer 
○ 
構成解析用節点ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
id_node1,2,3,4 は四角形を構成できる順序で定義する。節点１→節点２を要素内ｘ方向として、節
点３をｘ－ｙ平面内として定義する。 
・例 
<StbAnaPlaneRectangles> 
<StbAnaPlaneRectangle id="10" name="W4" id_property="10" 
 id_node1="5" id_node2="6" id_node3="16"  id_node4="15"/> 
</StbAnaPlaneRectangles>

## Page 257

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
247 
 
5.2.26. 接合部パネル要素（複数）：StbAnaNodePanels 
・概要 
説明  
：接合部パネル要素（複数） 
親要素 
：StbAnaMembers 
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
StbAnaNodePanel 
0 
制限無し 
接合部パネル要素 
 
・補足 
無し 
・例 
<StbAnaNodePanels> 
<StbAnaNodePanel id="11" name="J1" id_property="12" id_node="3"/> 
</StbAnaNodePanels>

## Page 258

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
248 
 
5.2.27. 接合部パネル要素：StbAnaNodePanel 
・概要 
説明  
：接合部パネル要素 
親要素 
：StbAnaNodePanels 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
angle 
double 
 
XY 平面傾斜角 
 
id_property 
integer 
○ 
接合部パネル要素性能 
StbAnaNodePanelProperty のID 
 
id_node 
integer 
○ 
構成解析用節点ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
接合部パネルは全体座標系により配置される。angleにより、XY平面上での回転角を設定する（angle
が省略された場合は0.0 とする）。 
・例 
<StbAnaNodePanels> 
<StbAnaNodePanel id="11" name="J1" id_property="12" id_node="3"/> 
</StbAnaNodePanels>

## Page 259

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
249 
 
5.2.28. 解析用部材性能：StbAnaProperties 
・概要 
説明  
：解析用部材性能 
親要素 
：StbAnaModel 
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
StbAnaBeamProperties 
0 
1 
梁要素性能（複数） 
 
StbAnaTrussProperties 
0 
1 
トラス要素性能（複数） 
 
StbAnaSpringProperties 
0 
1 
ばね要素性能（複数） 
 
StbAnaWallProperties 
0 
1 
壁エレメント要素性能（複数） 
 
StbAnaPlaneProperties 
0 
1 
平面要素性能（複数） 
 
StbAnaNodePanelProperties 
0 
1 
接合部パネル要素性能（複数） 
 
・補足 
無し 
・例 
<StbAnaProperties> 
<StbAnaBeamProperties> 
（略） 
</StbAnaBeamProperties> 
<StbAnaTrussProperties> 
（略） 
</StbAnaTrussProperties> 
<StbAnaSpringProperties> 
（略） 
</StbAnaSpringProperties> 
<StbAnaWallProperties> 
（略） 
</StbAnaWallProperties> 
<StbAnaPlaneProperties> 
（略） 
</StbAnaPlaneProperties> 
<StbAnaNodePanelProperties> 
（略） 
</StbAnaNodePanelProperties> 
</StbAnaProperties>

## Page 260

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
250 
 
5.2.29. 梁要素性能（複数）：StbAnaBeamProperties 
・概要 
説明  
：梁要素性能（複数） 
親要素 
：StbAnaProperties 
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
StbAnaBeamProperty 
0 
制限無し 
梁要素性能 
 
・補足 
無し 
・例 
<StbAnaBeamProperties> 
<StbAnaBeamProperty id="1" name="G0001" id_material="10" id_section="20"/> 
</StbAnaBeamProperties>

## Page 261

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
251 
 
5.2.30. 梁要素性能：StbAnaBeamProperty 
・概要 
説明  
：梁要素性能 
親要素 
：StbAnaBeamProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
梁要素性能名称 
 
id_material 
integer 
○ 
材料性能ID 
 
id_section 
integer 
○ 
断面性能ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaBeamProperties> 
<StbAnaBeamProperty id="1" name="G0001" id_material="10" id_section="20"/> 
</StbAnaBeamProperties>

## Page 262

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
252 
 
5.2.31. トラス要素性能（複数）：StbAnaTrussProperties 
・概要 
説明  
：トラス要素性能（複数） 
親要素 
：StbAnaProperties 
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
StbAnaTrussProperty 
0 
制限無し 
トラス要素性能 
 
・補足 
無し 
・例 
<StbAnaTrussProperties> 
<StbAnaTrussProperty id="2" name="BR0001" id_material="11" id_section="21"/> 
</StbAnaTrussProperties>

## Page 263

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
253 
 
5.2.32. トラス要素性能：StbAnaTrussProperty 
・概要 
説明  
：トラス要素性能 
親要素 
：StbAnaTrussProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
トラス要素性能名称 
 
id_material 
integer 
○ 
材料性能ID 
 
id_section 
integer 
○ 
断面性能ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaTrussProperties> 
<StbAnaTrussProperty id="2" name="BR0001" id_material="11" id_section="21"/> 
</StbAnaTrussProperties>

## Page 264

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
254 
 
5.2.33. ばね要素性能（複数）：StbAnaSpringProperties 
・概要 
説明  
：ばね要素性能（複数） 
親要素 
：StbAnaProperties 
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
StbAnaSpringProperty 
0 
制限無し 
ばね要素性能 
 
・補足 
無し 
・例 
<StbAnaSpringProperties> 
<StbAnaSpringProperty id="3" name="S0001" spring="10000.0"/> 
</StbAnaSpringProperties>

## Page 265

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
255 
 
5.2.34. ばね要素性能：StbAnaSpringProperty 
・概要 
説明  
：ばね要素性能 
親要素 
：StbAnaSpringProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
ばね要素性能名称 
 
spring 
double 
○ 
ばね値 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaSpringProperties> 
<StbAnaSpringProperty id="3" name="S0001" spring="10000.0"/> 
</StbAnaSpringProperties>

## Page 266

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
256 
 
5.2.35. 壁エレメント要素性能（複数）：StbAnaWallProperties 
・概要 
説明  
：壁エレメント要素性能（複数） 
親要素 
：StbAnaProperties 
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
StbAnaWallProperty 
0 
制限無し 
壁エレメント要素性能 
 
・補足 
無し 
・例 
<StbAnaWallProperties> 
<StbAnaWallProperty id="5" name="W0001" id_material="15" id_section="25"/> 
</StbAnaWallProperties>

## Page 267

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
257 
 
5.2.36. 壁エレメント要素性能：StbAnaWallProperty 
・概要 
説明  
：壁エレメント要素性能 
親要素 
：StbAnaWallProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
壁エレメント要素性能名称 
 
id_material 
integer 
○ 
材料性能ID 
 
id_section 
integer 
○ 
断面性能ID 
 
・内容 
無し 
・子要素 
無し 
・補足 
壁の断面性能については、面内方向のせん断、曲げ剛性および軸剛性を考慮する。StbAnaSection
との対応を示す。 
 
 
軸断面積 
せん断断面積 
断面二次 
モーメント 
StbAnaSection 
AX 
AY 
IZ 
 
・例 
<StbAnaWallProperties> 
<StbAnaWallProperty id="5" name="W0001" id_material="15" id_section="25"/> 
</StbAnaWallProperties>

## Page 268

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
258 
 
5.2.37. 平面要素性能（複数）：StbAnaPlaneProperties 
・概要 
説明  
：平面要素性能（複数） 
親要素 
：StbAnaProperties 
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
StbAnaPlaneProperty 
0 
制限無し 
平面要素性能 
 
・補足 
無し 
・例 
<StbAnaPlaneProperties> 
<StbAnaPlaneProperty id="6" name="PL0001" id_material="16" 
element_type="PLANE_STRESS" thickness="200.0"/> 
</StbAnaPlaneProperties>

## Page 269

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
259 
 
5.2.38. 平面要素性能：StbAnaPlaneProperty 
・概要 
説明  
：平面要素性能 
親要素 
：StbAnaPlaneProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
名称 
 
id_material 
integer 
○ 
材料性能ID 
 
element_type 
string 
○ 
平面要素タイプ 
平面応力要素:PLANE_STRESS 
平面歪要素：PLANE_STRAIN 
 
thickness 
double 
○ 
板厚 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaPlaneProperties> 
<StbAnaPlaneProperty id="6" name="PL0001" id_material="16" 
element_type="PLANE_STRESS" thickness="200.0"/> 
</StbAnaPlaneProperties>

## Page 270

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
260 
 
5.2.39. 接合部パネル要素性能（複数）：StbAnaNodePanelProperties 
・概要 
説明  
：接合部パネル要素性能（複数） 
親要素 
：StbAnaProperties 
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
StbAnaNodePanelProperty 
0 
制限無し 
接合部パネル要素性能 
 
・補足 
無し 
・例 
<StbAnaNodePanelProperties> 
<StbAnaNodePanelProperty id="8" name="JP0001" id_material="18" 
B_x="1000.0" B_y="1000.0" B_z="1000.0" t_x="40.0" t_y ="40.0" t_z="40.0" 
D_x ="1000.0" D_y ="1000.0" D_z="1000.0"/> 
</StbAnaNodePanelProperties>

## Page 271

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
261 
 
5.2.40. 接合部パネル要素性能：StbAnaNodePanelProperty 
・概要 
説明  
：接合部パネル要素性能 
親要素 
：StbAnaNodePanelProperties 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析用部材性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
接合部パネル要素性能名称 
 
id_material 
integer 
○ 
材料性能ID 
 
B_x 
double 
 
X 方向パネル幅 
 
B_y 
double 
 
Y 方向パネル幅 
 
B_z 
double 
 
Z 方向パネル幅 
 
t_x 
double 
 
X 方向パネル厚さ 
 
t_y 
double 
 
Y 方向パネル厚さ 
 
t_z 
double 
 
Z 方向パネル厚さ 
 
D_x 
double 
 
X 方向パネルせい 
 
D_y 
double 
 
Y 方向パネルせい 
 
D_z 
double 
 
Z 方向パネルせい 
 
・内容 
無し 
・子要素 
無し 
・補足 
Ｘ方向に加力した場合に変形するパネルはＸ方向パネルとなり、サイズは（幅：B_x、せい：D_x、厚さ：
t_x）で定義される。同様に、Ｙ方向パネルのサイズは（幅：B_y、高さ：D_y、厚さ：t_y）となる。Z方
向パネルについては、全体座標系X方向がB_z,全体座標系Y方向がD_zとする。 
 
 
 
 
 
 
 
 
 
B 
D 
t

## Page 272

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
262 
 
・例 
<<StbAnaNodePanelProperties> 
<StbAnaNodePanelProperty id="8" name="JP0001" id_material="18" 
B_x="1000.0" B_y="1000.0" B_z="1000.0" t_x="40.0" t_y ="40.0" t_z="40.0" 
D_x ="1000.0" D_y ="1000.0" D_z="1000.0"/> 
</StbAnaNodePanelProperties>

## Page 273

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
263 
 
5.2.41. 剛床指定（複数）：StbAnaFloorDiaphragms 
・概要 
説明  
：剛床指定（複数） 
親要素 
：StbAnaModel 
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
StbAnaFloorDiaphragm 
0 
制限無し 
剛床指定 
 
・補足 
無し 
・例 
<StbAnaFloorDiaphragms> 
<StbAnaFloorDiaphragm id="1" name="2FL"> 
<StbAnaNodeid_List> 
  <StbAnaNodeid id="10"/> 
  <StbAnaNodeid id="11"/> 
  <StbAnaNodeid id="12"/> 
</StbAnaNodeid_List> 
</StbAnaFloorDiaphragm> 
</StbAnaFloorDiaphragms>

## Page 274

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
264 
 
5.2.42. 剛床指定：StbAnaFloorDiaphragm 
・概要 
説明  
：剛床指定 
親要素 
：StbAnaFloorDiaphragms 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
剛床指定ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
剛床名称 
 
id_representative_node 
integer 
 
剛床代表節点のID 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaNodeid_List 
1 
1 
解析用節点ID リスト 
 
・補足 
id_representative_node・属性において剛床の慣性重心となる節点を指定する。アプリケーション
によっては入力されている剛床従属節点の質量から慣性重心を求める仕様のものも存在するため、
入力は任意としている。 
・例 
<StbAnaFloorDiaphragm id="1" name="2FL"> 
<StbAnaNodeid_List> 
  <StbAnaNodeid id="10"/> 
  <StbAnaNodeid id="11"/> 
  <StbAnaNodeid id="12"/> 
</StbAnaNodeid_List> 
</StbAnaFloorDiaphragm>

## Page 275

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
265 
 
5.2.43. 解析用節点ID リスト：StbAnaNodeid_List 
・概要 
説明  
：解析用節点ＩＤリスト 
親要素 
：StbAnaFloorDiaphragm 
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
StbAnaNodeid 
0 
制限無し 
解析用節点ID 
 
・補足 
無し 
 
 
5.2.44. 解析用節点ID：StbAnaNodeid 
・概要 
説明  
：解析用節点ＩＤ 
親要素 
：StbAnaNodeid_List 
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
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaFloorDiaphragm id="1" name="2FL"> 
<StbAnaNodeid_List> 
  <StbAnaNodeid id="10"/> 
  <StbAnaNodeid id="11"/> 
  <StbAnaNodeid id="12"/> 
</StbAnaNodeid_List> 
</StbAnaFloorDiaphragm>

## Page 276

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
266 
 
5.2.45. 材料性能（複数）：StbAnaMaterials 
・概要 
説明  
：材料性能（複数） 
親要素 
：StbAnaModel 
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
StbAnaMaterial 
0 
制限無し 
材料性能 
 
・補足 
無し 
・例 
<StbAnaMaterials> 
 <StbAnaMaterial id="1" name="S" E="205000" G="79000" poisson="0.3" 
 thermal="0" density="0.077"/> 
<StbAnaMaterial id="2" name="RC" E="21000" G="8750" poisson="0.2" 
 thermal="0" density="0.023"/> 
</StbAnaMaterials>

## Page 277

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
267 
 
5.2.46. 材料性能：StbAnaMaterial 
・概要 
説明  
：材料性能 
親要素 
：StbAnaMaterials 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
材料性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
材料性能名称 
 
E 
double 
 
ヤング係数 
 
G 
double 
 
せん断弾性係数 
 
poisson 
double 
 
ポアソン比 
 
thermal 
double 
 
線熱膨張係数 
 
density 
double 
 
比重 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaMaterials> 
 <StbAnaMaterial id="1" name="S" E="205000" G="79000" poisson="0.3" 
 thermal="0" density="0.077"/> 
<StbAnaMaterial id="2" name="RC" E="21000" G="8750" poisson="0.2" 
 thermal="0" density="0.023"/> 
</StbAnaMaterials>

## Page 278

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
268 
 
5.2.47. 断面性能（複数）：StbAnaSections 
・概要 
説明  
：断面性能（複数） 
親要素 
：StbAnaModel 
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
StbAnaSection 
0 
制限無し 
断面性能 
 
・補足 
無し 
・例 
<StbAnaSections> 
 <StbAnaSection id="1" name="SEC1" Ax="600000" Ay="500000" Az="500000" 
 Ix="3.5E+10" Iy="7.2E+10" Iz="1.3E+10"/> 
</StbAnaSections>

## Page 279

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
269 
 
5.2.48. 断面性能：StbAnaSection 
・概要 
説明  
：断面性能 
親要素 
：StbAnaSections 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
断面性能ID 
 
guid 
string 
 
GUID 
 
name 
string 
 
断面性能名称 
 
Ax 
double 
 
軸断面積 
 
Ay 
double 
 
せん断断面積 
省略時はせん断剛性∞
として取り扱う 
Az 
double 
 
せん断断面積 
省略時はせん断剛性∞
として取り扱う 
Ix 
double 
 
断面2 次モーメント(ねじり) 
 
Iy 
double 
 
断面2 次モーメント 
 
Iz 
double 
 
断面2 次モーメント 
 
・内容 
無し 
・子要素 
無し 
・補足 
 せん断断面積は、省略値を∞としていることに注意が必要である。 
・例 
<StbAnaSections> 
 <StbAnaSection id="1" name="SEC1" Ax="600000" Ay="500000" Az="500000" 
 Ix="3.5E+10" Iy="7.2E+10" Iz="1.3E+10"/> 
</StbAnaSections>

## Page 280

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
270 
 
5.2.49. 荷重ケース（複数）：StbAnaLoadCases 
・概要 
説明  
：荷重ケース（複数） 
親要素 
：StbAnaModel 
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
StbAnaLoadCase 
0 
制限無し 
荷重ケース 
 
・補足 
無し 
 
5.2.50. 荷重ケース：StbAnaLoadCase 
・概要 
説明  
：荷重ケース 
親要素 
：StbAnaLoadCases 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
荷重ケースID 
 
guid 
string 
 
GUID 
 
name 
string 
 
荷重ケース名称 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaLoadNode 
0 
制限無し 
節点荷重 
 
StbAnaLoadBeam 
0 
制限無し 
梁荷重 
 
StbAnaLoadTruss 
0 
制限無し 
トラス荷重 
 
StbAnaLoadWall 
0 
制限無し 
壁荷重 
 
StbAnaLoadSpring 
0 
制限無し 
ばね荷重 
 
・補足 
無し

## Page 281

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
271 
 
5.2.51. 節点荷重：StbAnaLoadNode 
・概要 
説明  
：節点荷重 
親要素 
：StbAnaLoadCase 
・属性 
属性名 
型 
必須 
説明 
補足 
id_node 
integer 
○ 
対象節点ID 
 
ux 
double 
 
X 方向荷重 
 
uy 
double 
 
Y 方向荷重 
 
uz 
double 
 
Z 方向荷重 
 
tx 
double 
 
X 軸回りモーメント 
 
ty 
double 
 
Y 軸回りモーメント 
 
tz 
double 
 
Z 軸回りモーメント 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaLoadCases> 
  <StbAnaLoadCase id="1" name="G"> 
    <StbAnaLoadNode id_node="1" uz="100"/> 
    <StbAnaLoadNode id_node="2" ty="200"/> 
  </StbAnaLoadCase> 
</StbAnaLoadCases>

## Page 282

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
272 
 
5.2.52. 梁荷重：StbAnaLoadBeam 
・概要 
説明  
：梁荷重 
親要素 
：StbAnaLoadCase 
・属性 
属性名 
型 
必須 
説明 
補足 
id_member 
integer 
○ 
対象解析部材ID 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaLoadBeamCMQ 
0 
制限無し 
梁CM0Q 荷重 
 
StbAnaLoadBeamInitialStress 
0 
制限無し 
梁要素初期応力 
 
・補足 
無し 
・例 
<StbAnaLoadCases> 
  <StbAnaLoadCase id="1" name="G"> 
    <StbAnaLoadBeam id_member="1"> 
      <StbAnaLoadBeamCMQ start_Qz="100.0" start_Cy="-50.0" 
 end_Qz="100.0" end_Cy="50.0" center_M0y="50.0"/> 
    </StbAnaLoadBeam> 
  </StbAnaLoadCase> 
</StbAnaLoadCases>

## Page 283

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
273 
 
5.2.53. 梁CMQ 荷重：StbAnaLoadBeamCMQ 
・概要 
説明  
：梁CMQ 荷重 
親要素 
：StbAnaLoadBeam 
・属性 
属性名 
型 
必須 
説明 
補足 
start_N 
double 
 
始端側X 方向軸力 
 
start_Qy 
double 
 
始端側Y 向せん断力 
 
start_Qz 
double 
 
始端側Z 方向せん断力 
 
start_T 
double 
 
始端側X 軸回りモーメント 
 
start_Cy 
double 
 
始端側Y 軸回りモーメント 
 
start_Cz 
double 
 
始端側Z 軸回りモーメント 
 
end_N 
double 
 
終端側X 方向軸力 
 
end_Qy 
double 
 
終端側Y 方向せん断力 
 
end_Qz 
double 
 
終端側Z 方向せん断力 
 
end_T 
double 
 
終端側X 軸回りモーメント 
 
end_Cy 
double 
 
終端側Y 軸回りモーメント 
 
end_Cz 
double 
 
終端側Z 軸回りモーメント 
 
center_T 
double 
 
中央部X 軸回りモーメント 
 
center_M0y 
double 
 
中央部Y 軸回りモーメント 
 
center_M0z 
double 
 
中央部Z 軸回りモーメント 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaLoadCase id="1" name="G"> 
  <StbAnaLoadBeam id_member="1"> 
    <StbAnaLoadBeamCMQ start_Qz="100.0" start_Cy="-50.0" 
end_Qz="100.0" end_Cy="50.0" center_M0y="50.0"/> 
  </StbAnaLoadBeam> 
</StbAnaLoadCase>

## Page 284

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
274 
 
5.2.54. 梁要素初期応力：StbAnaLoadBeamInitialStress 
・概要 
説明  
：梁要素初期応力 
親要素 
：StbAnaLoadBeam 
・属性 
属性名 
型 
必須 
説明 
補足 
N 
double 
 
軸力（引張：正、圧縮：負） 
 
start_Qy 
double 
 
始端側Y 方向せん断力 
 
start_Qz 
double 
 
始端側Z 方向せん断力 
 
start_My 
double 
 
始端側Y 軸回りモーメント 
 
start_Mz 
double 
 
始端側Z 軸回りモーメント 
 
end_Qy 
double 
 
終端側Y 方向せん断力 
 
end_Qz 
double 
 
終端側Z 方向せん断力 
 
end_My 
double 
 
終端側Y 軸回りモーメント 
 
end_Mz 
double 
 
終端側Z 軸回りモーメント 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaLoadBeam id_member="1"> 
  <StbAnaLoadBeamInitialStress start_Qz="300" start_My="-200" end_Qz="300" end_My="200"/> 
</StbAnaLoadBeam>

## Page 285

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
275 
 
5.2.55. トラス荷重：StbAnaLoadTruss 
・概要 
説明  
：トラス荷重 
親要素 
：StbAnaLoadCase 
・属性 
属性名 
型 
必須 
説明 
補足 
id_member 
integer 
○ 
対象解析部材ID 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaLoadTrussInitialStress 
0 
制限無し 
トラス要素初期応力 
 
・補足 
無し 
 
5.2.56. トラス要素初期応力：StbAnaLoadTrussInitialStress 
・概要 
説明  
：トラス要素初期応力 
親要素 
：StbAnaLoadTruss 
・属性 
属性名 
型 
必須 
説明 
補足 
N 
double 
 
軸力（引張：正、圧縮：負） 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaLoadTruss id_member="1"> 
  <StbAnaLoadTrussInitialStress N="100.0"/> 
</StbAnaLoadTruss>

## Page 286

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
276 
 
5.2.57. 壁荷重：StbAnaLoadWall 
・概要 
説明  
：壁荷重 
親要素 
：StbAnaLoadCase 
・属性 
属性名 
型 
必須 
説明 
補足 
id_member 
integer 
○ 
対象解析部材ID 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaLoadWallInitialStress 
0 
制限無し 
壁エレメント要素初期応力 
 
・補足 
無し 
 
5.2.58. 壁エレメント要素初期応力：StbAnaLoadWallInitialStress 
・概要 
説明  
：壁エレメント要素初期応力 
親要素 
：StbAnaLoadWall 
・属性 
属性名 
型 
必須 
説明 
補足 
N 
double 
 
軸力 
 
Q 
double 
 
せん断力 
 
bottom_M 
double 
 
壁脚モーメント 
 
top_M 
double 
 
壁頭モーメント 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaLoadWall id_member="1"> 
  <StbAnaLoadWallInitialStress N="100.0" Q="500.0" bottom_M="200.0" top_M="300.0" /> 
</StbAnaLoadWall>

## Page 287

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
277 
 
5.2.59. ばね荷重：StbAnaLoadSpring 
・概要 
説明  
：ばね荷重 
親要素 
：StbAnaLoadCase 
・属性 
属性名 
型 
必須 
説明 
補足 
id_member 
integer 
○ 
対象解析部材ID 
 
・内容 
無し 
・子要素 
要素名 
最小回数 
最大回数 
説明 
補足 
StbAnaLoadSpringInitialStress 
0 
制限無し 
ばね要素初期応力 
 
・補足 
無し 
 
 
5.2.60. ばね要素初期応力：StbAnaLoadSpringInitialStress 
・概要 
説明  
：ばね要素初期応力 
親要素 
：StbAnaLoadSpring 
・属性 
属性名 
型 
必須 
説明 
補足 
force 
double 
○ 
ばね応力 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaLoadSpring id_member="1"> 
  <StbAnaLoadSpringInitialStress force="100.0"/> 
</StbAnaLoadSpring>

## Page 288

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
278 
 
5.2.61. 解析ケース（複数）：StbAnaAnalyses 
・概要 
説明  
：解析ケース（複数） 
親要素 
：StbAnaModel 
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
StbAnaAnalysisStaticLinear 
0 
制限無し 
静的線形解析ケース 
 
・補足 
無し 
 
5.2.62. 静的線形解析ケース：StbAnaAnalysisStaticLinear 
・概要 
説明  
：静的線形解析ケース 
親要素 
：StbAnaAnalyses 
・属性 
属性名 
型 
必須 
説明 
補足 
id 
integer 
○ 
解析ケースID 
 
guid 
string 
 
GUID 
 
name 
string 
 
解析ケース名称 
 
id_initial_stress_load_case 
integer 
 
対象初期応力荷重ケースID 
 
id_load_case 
integer 
○ 
対象荷重ケースID 
 
・内容 
無し 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaAnalyses> 
  <StbAnaAnalysisStaticLinear id="1" id_load_case="1"/> 
  <StbAnaAnalysisStaticLinear id="2" id_load_case="2"/> 
</StbAnaAnalyses>

## Page 289

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
279

## Page 290

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
280 
 
5.3. 
解析モデル関連定義：StbAnaRelations 
・概要 
説明  
：解析モデル要素とST-Bridge 本編の要素を関連付ける定義 
親要素 
：StbAnaModel 
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
StbAnaNodeRel 
0 
制限無し 
解析用節点関連定義 
 
StbAnaStoryRel 
0 
制限無し 
解析用階関連定義 
 
StbAnaMemberRel 
0 
制限無し 
解析用部材関連定義 
 
StbAnaCalMemberRel 
0 
制限無し 
解析用部材関連定義 
(計算用・分割) 
 
StbAnaPropertyRel 
0 
制限無し 
解析用部材性能関連定義 
 
・補足 
解析モデルにおける各要素（解析用節点、解析用階、解析用部材、解析用部材性能）とStbModel
の要素を関連付ける必要がある場合に記述する。 
・例 
<StbAnaModel> 
（略） 
  <StbAnaRelations> 
    <StbAnaNodeRel> 
      （略） 
    </StbAnaNodeRel> 
<StbAnaMemberRel> 
      （略） 
    </StbAnaMemberRel> 
<StbAnaPropertyRel> 
      （略） 
    </StbAnaPropertyRel> 
</StbAnaRelations> 
</StbAnaModel>

## Page 291

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
281 
 
5.3.1. 解析用節点関連定義：StbAnaNodeRel 
・概要 
説明  
：解析用節点とStbModel の節点を関連付ける定義 
親要素 
：StbAnaRelations 
・属性 
属性名 
型 
必須 
説明 
補足 
id_ana_node 
integer 
○ 
解析用節点ID 
 
・内容 
内容 
型 
必須 
説明 
補足 
StbNode.id 
[monolist] 
integer 
○ 
節点ID の列記（スペース区切り） 
 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaModel> 
（略） 
  <StbAnaRelations> 
    <StbAnaNodeRel id_ana_node="101">1</StbAnaNodeRel> 
    <StbAnaNodeRel id_ana_node="102">2</StbAnaNodeRel> 
    <StbAnaNodeRel id_ana_node="103">3</StbAnaNodeRel> 
    <StbAnaNodeRel id_ana_node="104">4</StbAnaNodeRel> 
</StbAnaRelations> 
</StbAnaModel>

## Page 292

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
282 
 
5.3.2. 解析用階関連定義：StbAnaStoryRel 
・概要 
説明  
：解析用階とStbModel の階情報を関連付ける定義 
親要素 
：StbAnaRelations 
・属性 
属性名 
型 
必須 
説明 
補足 
id_ana_story 
integer 
○ 
解析用階ID 
 
・内容 
内容 
型 
必須 
説明 
補足 
StbStory.id 
[monolist] 
integer 
○ 
階ID の列記（スペース区切り） 
 
・子要素 
無し 
・補足 
無し 
・例 
<StbAnaModel> 
（略） 
  <StbAnaRelations> 
    <StbAnaStoryRel id_ana_story="1">1</StbAnaStoryRel> 
    <StbAnaStoryRel id_ana_story="2">2</StbAnaStoryRel> 
    <StbAnaStoryRel id_ana_story="3">3</StbAnaStoryRel> 
    <StbAnaStoryRel id_ana_story="4">4</StbAnaStoryRel> 
</StbAnaRelations> 
</StbAnaModel>

## Page 293

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
283 
 
5.3.3. 解析用部材関連定義：StbAnaMemberRel 
・概要 
説明  
：解析用部材とStbModel の部材情報を関連付ける定義 
親要素 
：StbAnaRelations 
・属性 
属性名 
型 
必須 
説明 
補足 
id_ana_member 
integer 
○ 
解析用部材ID 
 
member_kind 
string 
 
部材要素種類 
以下のいずれかの値を取るものとする 
StbColumn,StbPost,StbGirder, 
StbBeam,StbBrace,StbSlab,StbWall 
 
・内容 
内容 
型 
必須 
説明 
補足 
StbColumn.id 
StbPost.id 
StbGirder.id 
StbBeam.id 
StbBrace.id 
StbSlab.id 
StbWall.id 
[monolist] 
integer 
○ 
部材ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 要素種類member_kind に対応したST-Bridge 本編の部材ID を列記する。 
・例 
<StbAnaModel> 
（略） 
  <StbAnaRelations> 
<StbAnaMemberRel id_ana_member="1" member_kind="StbColumn">1</StbAnaMemberRel> 
<StbAnaMemberRel id_ana_member="2" member_kind="StbColumn">2</StbAnaMemberRel> 
<StbAnaMemberRel id_ana_member="3" member_kind="StbColumn">3</StbAnaMemberRel> 
<StbAnaMemberRel id_ana_member="4" member_kind="StbColumn">4</StbAnaMemberRel> 
</StbAnaRelations> 
</StbAnaModel>

## Page 294

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
284 
 
5.3.4. 解析用部材関連定義(計算用・分割)：StbAnaCalMemberRel 
・概要 
説明  
：解析用部材とStbCalGirder およびStbCalColumn を関連付ける定義 
親要素 
：StbAnaRelations 
・属性 
属性名 
型 
必須 
説明 
補足 
id_ana_member 
integer 
○ 
解析用部材ID 
 
member_kind 
string 
 
部材要素種類 
以下のいずれかの値を取るものとする 
StbCalColumn, StbCalGirder 
 
・内容 
内容 
型 
必須 
説明 
補足 
StbCalColumn.id 
StbCalGirder.id 
[monolist] 
integer 
○ 
部材ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 要素種類member_kind に対応したStbCalGirder およびStbCalColumn の部材ID を列記する。 
・例 
<StbAnaModel> 
（略） 
  <StbAnaRelations> 
<StbAnaCalMemberRel 
     id_ana_member="201" member_kind="StbCalColumn">100</StbAnaCalMemberRel> 
</StbAnaRelations> 
</StbAnaModel>

## Page 295

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
285 
 
5.3.5. 解析用部材性能関連定義：StbAnaPropertyRel 
・概要 
説明  
：解析用部材性能とStbModel の断面情報を関連付ける定義 
親要素 
：StbAnaRelations 
・属性 
属性名 
型 
必須 
説明 
補足 
id_ana_property 
integer 
○ 
解析用部材性能ID 
 
section_kind 
string 
 
断面要素種類 
以下のいずれかの値を取るものとする 
StbSecColumn_RC,StbSecColumn_S, 
StbSecColumn_SRC,StbSecColumn_CFT, 
StbSecBeam_RC,StbSecBeam_S, 
StbSecBeam_SRC,StbSecBrace_S, 
StbSecSlab_RC,StbSecSlabDeck, 
StbSecSlabPrecast,StbSecWall_RC 
 
・内容 
内容 
型 
必須 
説明 
補足 
StbSecColumn_RC.id 
StbSecColumn_S.id 
StbSecColumn_SRC.id 
StbSecColumn_CFT.id 
StbSecBeam_RC.id 
StbSecBeam_S.id 
StbSecBeam_SRC.id 
StbSecBrace_S.id 
StbSecSlab_RC.id 
StbSecSlabDeck.id 
StbSecSlabPrecast.id 
StbSecWall_RC.id 
[monolist] 
integer 
○ 
断面ID の列記（スペース区切り） 
※(1) 
・子要素 
無し 
・補足 
(1) 要素種類member_kind に対応したST-Bridge 本編の断面要素のID を列記する。

## Page 296

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
286 
 
 
・例 
<StbAnaModel> 
（略） 
  <StbAnaRelations> 
<StbAnaPropertyRel 
 id_ana_property="401" member_kind ="StbSecColumn_RC">1001</StbAnaPropertyRel> 
</StbAnaRelations> 
</StbAnaModel>

## Page 297

ST-Bridge XML ファイル仕様書 計算データ編（ver.2.0）Revision 2 
2021/02/28 
287 
 
6. 修正履歴 
 
日付 
修正種別 
バージョン 
内容 
備考 
2015.3.31 
追加 
1.0.0 
計算データ編 Ver.1.0 リリース 
一旦、Ver.1.0 として仮リリース 
2015.4.27 
 
変更 
1.3.0 
・StbCalData, StbAnaModel の役割分担につい
て記載（3.2. 設計モデルと計算モデル）。 
・StbCalPointLoad 修正（4.3.14 特殊節点荷重）。 
・計算データ編Ver.1.3 リリース。 
 
 
 
本体仕様書に合わせてVer.1.3として
仮リリース、変更履歴をリセットした。 
2018.6.27 
変更 
1.4.0 
計算データ編 Ver.1.4 リリース 
一旦、Ver.1.4 として仮リリース 
2021.2.28 
変更 
2.0.2 
計算データ編 Ver.2.0.2 リリース 
（本体仕様の見直しに合わせて全面改訂） 
本体仕様書に合わせてVer.2.0.2 とし
てリリース、変更履歴をリセットした。
