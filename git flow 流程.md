公司的流程非常混乱
但是向公司推荐`git flow`流程的时候，有些地方我却不是很清晰。

**release的版本号在哪里修改？**

我是直接在`develop`修改了，然后`Merge`到`Master`分支

`gitflow`流程是创建`release`分支，修改好版本号后合到`Master`和`Develop`分支，之后再删除掉

**hotfix应该如何处理**

我认为`master`上`release`的版本会存在多个，所以有时候`Hotfix`的版本存在不能`Merge`到`Master`，也不能`Merge`到`release`的情况；

如果不存在多个版本，都是最新版本那就很好处理了，`Hotfix`肯定都`Merge`到`Develop`和`Master`

**版本号如何命名？**

### 常用git flow命令

``` shell
$ git flow feature start rss-feed
Switched to a new branch 'feature/rss-feed'

Summary of actions:
- A new branch 'feature/rss-feed' was created, based on 'develop'
- You are now on branch 'feature/rss-feed'
```

``` shell
$ git flow feature finish rss-feed
Switched to branch 'develop'
Updating 6bcf266..41748ad
Fast-forward
    feed.xml | 0
    1 file changed, 0 insertions(+), 0 deletions(-)
    create mode 100644 feed.xml
Deleted branch feature/rss-feed (was 41748ad).
```

``` shell
$ git flow release start 1.1.5
Switched to a new branch 'release/1.1.5'
```

``` shell
$ git flow release finish 1.1.5
```

``` shell

$ git flow hotfix start missing-link
```

``` shell

$ git flow hotfix finish missing-link
```



### **常用`git`命令**

**feature**

``` shell
$ git checkout -b myfeature develop
Switched to a new branch "myfeature"
```

``` shell
$ git checkout develop
Switched to branch 'develop'
$ git merge --no-ff myfeature
Updating ea1b82a..05e9557
(Summary of changes)
$ git branch -d myfeature
Deleted branch myfeature (was 05e9557).
$ git push origin develop
```

**release**

``` shell
$ git checkout -b release-1.2 develop
Switched to a new branch "release-1.2"
$ ./bump-version.sh 1.2
Files modified successfully, version bumped to 1.2.
$ git commit -a -m "Bumped version number to 1.2"
[release-1.2 74d9424] Bumped version number to 1.2
1 files changed, 1 insertions(+), 1 deletions(-)

```

``` shell

$ git checkout master
Switched to branch 'master'
$ git merge --no-ff release-1.2
Merge made by recursive.
(Summary of changes)
$ git tag -a 1.2
```

``` shell

$ git checkout develop
Switched to branch 'develop'
$ git merge --no-ff release-1.2
Merge made by recursive.
```

``` shell

$ git branch -d release-1.2
Deleted branch release-1.2 (was ff452fe).
```

**hotfix**

``` shell 

$ git checkout -b hotfix-1.2.1 master
Switched to a new branch "hotfix-1.2.1"
$ ./bump-version.sh 1.2.1
Files modified successfully, version bumped to 1.2.1.
$ git commit -a -m "Bumped version number to 1.2.1"
[hotfix-1.2.1 41e61bb] Bumped version number to 1.2.1
1 files changed, 1 insertions(+), 1 deletions(-)
```

``` shell

$ git commit -m "Fixed severe production problem"
[hotfix-1.2.1 abbe5d6] Fixed severe production problem
5 files changed, 32 insertions(+), 17 deletions(-)

```

``` shell
$ git checkout master
Switched to branch 'master'
$ git merge --no-ff hotfix-1.2.1
Merge made by recursive.
(Summary of changes)
$ git tag -a 1.2.1

```

