# 議事録生成用CLIツールmami

## 使い方

1. `mami config`
    * 議事録を生成したいpathを入力する
    * example: `~/Documents/minutes`
1. `mami set-editor`
    * markdownファイルを編集する際に使用するエディタのアプリケーションpathを入力する
    * example: `/usr/local/bin/vim`
    * もしエディタのアプリケーションpathが分からない場合は `which` コマンドを使用して調べて下さい
1. `mami make hoge`
    * 1で指定したpath直下に `YYYYMM` というディレクトリが生成される
    * `1で指定したpath/YYYYMM/` 直下に `hoge.md` ファイルが生成され、2で指定したエディタで開かれる

基本的にはこれで議事録を作成したいと思ったタイミングで `mami make` と入力すれば使用可

## テンプレートファイルを元に新規でファイルを作成したい場合
既にテンプレートファイルが用意されていて、そのファイルを元に新規ファイルを作成しないといけない場合。<br>

1. `mami set template.md template1`
    * 第一引数:元にしたいテンプレートファイル
    * 第二引数:テンプレートファイルの別名
1. `mami make foo template1`
    * fooファイルはtemplate.mdファイルを元に作成される

議事録を生成のテンプレートが決まっている場合は `default` でテンプレートファイルを登録する。

`mami set template.md default`

`mami make bar`

bar.mdはtemplate.mdを元に生成される
