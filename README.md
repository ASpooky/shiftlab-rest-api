## 概要
シフトを管理するアプリのRESTAPIをGo＆ECHOを用いて開発しました.
アプリにログインし、自分の職場の時給を設定、シフトの時間を登録することで年間の給与をバイトごとに可視化するアプリになっています。

## アーキテクチャ
バックエンドは、クリーンアーキテクチャを用いて開発しました。クリーンアーキテクチャを採用した理由としては、機能追加がしやすい点や保守性が高いコードが書けるためです。この強みを生かして今後も自分なりに考えたアイデアでプロダクトをより良くしていこうと考えています。
![OIP](https://github.com/ASpooky/shiftlab-rest-api/assets/100838115/5b7d26f9-fe56-4e22-bbce-105659b6390d)

## ER図
モデルはユーザ・ワークスペース・シフトの3つです。

![https://github.com/ASpooky/shiftlab-rest-api/er](https://github.com/ASpooky/shiftlab-rest-api/blob/main/er.png)
