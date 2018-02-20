# 議事録生成用CLIツールma-mi

## 使い方

1. `ma-mi config`
    * 議事録を生成したいpathを入力する
    * example: `~/Documents/minutes`
1. `ma-mi set-editor`
    * markdownファイルを編集する際に使用するエディタのアプリケーションpathを入力する
    * example: `/usr/local/bin/vim`
    * もしエディタのアプリケーションpathが分からない場合は `which` コマンドを使用して調べて下さい
1. `ma-mi make hoge`
    * 1で指定したpath直下に `YYYYMM` というディレクトリが生成される
    * `1で指定したpath/YYYYMM/` 直下に `hoge.md` ファイルが生成され、2で指定したエディタで開かれる

基本的にはこれで議事録を作成したいと思ったタイミングで `ma-mi make` と入力すれば使用可

## テンプレートファイルを元に新規でファイルを作成したい場合
既にテンプレートファイルが用意されていて、そのファイルを元に新規ファイルを作成しないといけない場合。<br>

1. `ma-mi set template.md template1`
    * 第一引数:元にしたいテンプレートファイル
    * 第二引数:テンプレートファイルの別名
1. `ma-mi make foo template1`
    * fooファイルはtemplate.mdファイルを元に作成される

議事録を生成のテンプレートが決まっている場合は `default` でテンプレートファイルを登録する。

`ma-mi set template.md default`

`ma-mi make bar`

bar.mdはtemplate.mdを元に生成される
