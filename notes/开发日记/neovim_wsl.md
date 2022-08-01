安装nvim，默认配置无效，安装的版本不 是最新的版本
之前搜索了一个版本是通过获取安装包，解压来安装，不太喜欢这种方式
以及今天又折腾了一下，在公司发现好像是没有打勾使用wsl，回家检测发现是有打勾，真是奇怪呢。

于是看插件出错信息，vscode是需要大于0.5版本，又google了一下，vpn续费之后没那么垃圾了，终于可以搜到东西了耶。
非常的顺利，一下命令亲测可用
``` bash
sudo add-apt-repository ppa:neovim-ppa/unstable
sudo apt-get update
sudo apt-get install neovim
```

安装成功之后，nvim --version显示0.8了，大成功！！！

后记：之后vscode就能成功使用neovim插件了，但是输入法有点跳，先不管吧...下次解决