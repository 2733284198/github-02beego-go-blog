## Go Blog
[![GitHub stars](https://img.shields.io/github/stars/1920853199/go-blog)](https://github.com/1920853199/go-blog/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/1920853199/go-blog)](https://github.com/1920853199/go-blog/network)
[![GitHub license](https://img.shields.io/github/license/1920853199/go-blog)](https://github.com/1920853199/go-blog/master/LICENSE)

一个快速创建个人博客，cms 的系统

### 包含功能

[查看](http://go-blog.cn "查看")

> Go Blog 官网 http://go-blog.cn

> 演示站点 http://leechan.online

![Go Blog 官网](/view.png "Go Blog 官网")

### 更新日志
|时间|功能|
|:---|:---|
|2020年1月23日|新增文章顶置功能|
|2020年2月2日|新增自定义导航功能|
|2020年2月4日|新增站点公告功能|
|2020年2月6日|新增友情链接模块|
|2020年2月6日|新增点赞功能|
|2020年2月20日|新增站点用户管理模块，可新增修改后台用户以及密码|
|2020年3月5日|JS渲染页面改为后端渲染，优化页面SEO|
|2020年3月6日|添加标签云|
|2020年3月7日|XSS攻击过滤|
|2020年3月12日|添加点赞限制|
|2020年3月23日|添加硬盘使用监控|
|2020年3月25日|新增图片放大预览|

### 更新 v1.1.0
> 接下来的v1.1.0版本将支持绑定公众号

新功能：
1. 素材管理
2. 自定义菜单
3. 消息群发
4. 关键词回复

### Install 

1. 把Go Blog项目拉到本地

```
https://github.com/1920853199/go-blog.git
```

2. 新建数据库，导入数据库文件，数据库文件/database/blog.sql

3. 修改项目配置信息

```
#conf/app.conf

appname = go-blog
httpport = 8088
runmode = dev
EnableAdmin = false
sessionon = true
url = 127.0.0.1:8088
view = default

limit = 10
title = Go Blog
autograph = 如今的我，谈不上幸福，也谈不上不幸。

[db]
dbType = mysql
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

4. 在bo-blog 根目录下执行bee run ，访问 http://127.0.0.1:8888 即可

5. 守护进程模式运行 可以了解PM2的相关信息，配置可查看start.sh 文件

6. nginx代理示例

```
server {
        listen 80;
        server_name go-blog.cn;
        root    /home/data/go-blog;

        location ~ \.(txt|xml)$ {
                root /home/data/go-blog;
        }

        location / {
            proxy_pass http://127.0.0.1:8889;
            #proxy_redirect off;
            proxy_http_version    1.1;
            proxy_cache_bypass    $http_upgrade;
            proxy_set_header Upgrade            $http_upgrade;
            proxy_set_header Connection         "upgrade";
            proxy_set_header Host               $host;
            proxy_set_header X-Real-IP          $remote_addr;
            proxy_set_header X-Forwarded-For    $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto  $scheme;
            proxy_set_header X-Forwarded-Host   $host;
            proxy_set_header X-Forwarded-Port   $server_port;
        }

        access_log    /home/wwwlogs/go-blog.access.log;
}

```
