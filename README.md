# Goならわかるシステムプログラミング 読書メモ
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


