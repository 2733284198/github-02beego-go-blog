## Go Blog
### 一个快速创建个人博客，cms 的系统

>  包含站点配置，文章管理，文章评论管理，分类管理，最新评论，点击排行，档案统计，留言，评论，回复留言，回复评论，日志，主题颜色修改等等功能

> Go Blog 官网 http://go-blog.cn

> 演示站点 http://leechan.online

### 更新日志
|时间|功能|
|:---|:---|
|2020年1月23日|新增文章顶置功能|
|2020年2月2日|新增自定义导航功能|
|2020年2月4日|新增站点公告功能|
|2020年2月6日|新增友情链接模块|
|2020年2月6日|新增点赞功能|

### 更新 v1.1.0
> 接下来的v1.1.0版本将支持绑定公众号

新功能：
1. 素材管理
1. 自定义菜单
2. 消息群发
3. 关键词回复

### Install 
1. 先下载安装Beego

```
go get github.com/astaxie/beego
```

2. 把go-blog项目拉到本地beego项目的src目录下

```
https://github.com/1920853199/go-blog.git
```

3. 导入数据库，数据库文件/database/blog.sql

4. 修改项目配置信息

```
#conf/app.conf

appname = go-blog
httpport = 8088
runmode = dev
EnableAdmin = false
sessionon = true
url = 127.0.0.1:8088

limit = 10
title = Go Blog
autograph = 如今的我，谈不上幸福，也谈不上不幸。

[db]
dbUser = root
dbPass = root
dbHost = 127.0.0.1
dbPort = 3306
dbName = blog

[redis]
rHost = 127.0.0.1
rPort = 6379

[wechat]
AppID = xxxxxxx
AppSecret = xxxxxxx
Token = xxxxxxx
EncodingAESKey = xxxxxxx


```

5. 在bo-blog 根目录下执行bee run ，访问 http://127.0.0.1:8888 即可

6. 守护进程模式运行 可以了解PM2的相关信息，配置可查看start.sh 文件
