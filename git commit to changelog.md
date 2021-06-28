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