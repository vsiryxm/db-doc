## 💡 简介

db-doc 是一款生成在线数据库文档的小工具

## 📸 截图

### 在线文档

![在线文档](https://gitee.com/czxtc/note-pic-repo/raw/master//note/20210323153924.png)

### 离线文档

#### Markdown

![Markdown](https://gitee.com/czxtc/note-pic-repo/raw/master//note/20210323154242.png)

#### HTML

![HTML](https://gitee.com/czxtc/note-pic-repo/raw/master//note/20210323154126.png)

#### PDF

![PDF](https://gitee.com/czxtc/note-pic-repo/raw/master//note/20210324091920.png)



## ⚒️ 支持

* [x] Mysql  
* [x] SQL Server
* [x] Postgre SQL    

## ⚗ 用法

### 下载

[点击下载](https://github.com/viodo/db-doc/releases) 

解压后双击打开DbDoc, 按照提示输入操作即可

```shell
? Database type:
1:MySQL
2:SQL Server
1
? Database host (127.0.0.1) :
127.0.0.1
? Database port (3306) :
3306
? Database username (root) :
root
? Database password (123456) :
123456
? Database name:
test
? Document type:
1:Online(在线文档)
2:Offline(离线文档)
1
Doc generate successfully!
Doc server is running : http://127.0.0.1:3000
```
浏览器访问: http://127.0.0.1:3000

## 🙏 鸣谢

* [Docsify - A magical documentation site generator](https://docsify.js.org)

* [blackfriday - a markdown processor for Go](https://github.com/russross/blackfriday)

* [chromedp - A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol.](https://github.com/chromedp/chromedp)