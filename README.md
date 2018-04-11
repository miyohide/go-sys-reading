# Goならわかるシステムプログラミング 読書メモ
## 2018年3月19日（月）
### P53 - P63
- 久しぶりにメモ。読書会自身はやっていたけど、このメモは完全に放置していた。
- 第三章の課題を解いてみる。Q3.5がわからない。どうやって実装したらいいんだろう。
  - `io.LimitReader`というのがあるっぽいので試してみる。
  - うまくできた気がする。

## 2017年12月18日（月）
### P22 - P34

- ここから先は、io.Writerを使っているサンプル集のような気がする。
  - ただ、最近はGo言語においてインタフェースっていうのはテストのしやすさの観点から非常に重要ということがわかってきたので、ちゃんと把握しておくのも大事な気がする。
- P27の脚注について。`bytes.Buffer`はRubyでいう`StringIO`みたいなものと把握しておけば違いがはっきりした。
- godocで生成するものと、dashで見えるものの違いについて
  - godocは自分のGOPATH以下にあるものを全てを対象にする。
  - dashで見えるのは、標準ライブラリのものだけなのかなと思う。

## 2017年12月4日（月）
### P9 - P22

- `syscall.Write`が呼ばれる場所をデバッカで追う。
  - 本と場所がちょっと違う。
  - どうやらリファクタリングが入ってソースが違うようになっているみたい。
  - https://github.com/golang/go/commit/3792db518327c685da17ca6c6faa4e1d2da4c33c#diff-43f69608b275ac1ddce65e6f77532212
- ファイルディスクリプタについて
  - 「なるほどUNIXプロセス」を今一度読み直そう。
  - https://tatsu-zine.com/books/naruhounix
- Go言語のインタフェースの命名規則が「〜er」「〜or」になっているのがよく分からない。

## 2017年11月20日（月）
### P1 - P9

- Visual Studio CodeでGoのデバッカを使えるようにする
  - https://qiita.com/TsuyoshiUshio@github/items/ba15b1a7e6c6e5ffaf43
  - https://suruseas.github.io/golang/delve/vsc/2017/10/19/golang-setup.html
  - 参考にしたのは上記2つのサイト。結局、Command Line Toolsの最新版を入れたりしてようやく完了。


