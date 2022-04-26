```bash
$ npm install -g conventional-changelog
$ cd my-project
$ conventional-changelog -p angular -i CHANGELOG.md -w
```

发布changelog.md

``` bash
$ conventional-changelog -p angular -i CHANGELOG.md -w -r 0
```

配置成 package.json

```javascript
{
  "scripts": {
    "changelog": "conventional-changelog -p angular -i CHANGELOG.md -w -r 0"
  }
}
```

然后

```bash
$ npm run changelog
```

git merge flow
  
  ```bash
  $ git checkout master
$ git checkout -b newFeature
$ ...
$ git commit -m 'work from day 1'
$ ...
$ git commit -m 'work from day 2'
$ ...
$ git commit -m 'finish the feature'
$ git checkout master
$ git merge --no-ff newFeature -m 'add new feature'
$ git log
// something like below
commit 'add new feature'         // => commit created at merge with proper message
commit 'finish the feature'
commit 'work from day 2'
commit 'work from day 1'
$ gitk                           // => see details with graph

$ git checkout -b anotherFeature        // => create a new branch (*)
$ ...
$ git commit -m 'work from day 3'
$ ...
$ git commit -m 'work from day 4'
$ ...
$ git commit -m 'finish another feature'
$ git checkout master
$ git merge anotherFeature       // --ff is by default, message will be ignored
$ git log
// something like below
commit 'work from day 4'
commit 'work from day 3'
commit 'add new feature'
commit 'finish the feature'
commit ...
$ gitk                    
```