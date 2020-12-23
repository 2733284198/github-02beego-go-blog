CREATE TABLE `article` (
  `id` int(11) NOT NULL  AUTO_INCREMENT COMMENT 'ID',
  `user_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `category_id` int(11) NOT NULL COMMENT '分类ID',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT 'Tag',
  `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '摘要',
  `desc` longtext NOT NULL COMMENT '详情',
  `pv` int(255) DEFAULT '0' COMMENT 'px',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `review` int(11) DEFAULT '0' COMMENT '评论',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  AUTO_INCREMENT=1  DEFAULT CHARSET=utf8;

ALTER TABLE `article` ADD `recommend` INT(4) NOT NULL DEFAULT '0' COMMENT '是否顶置，0否；1是，默认否';
ALTER TABLE `article` ADD COLUMN `like`  int(11) NOT NULL DEFAULT 0 COMMENT '点赞数量' AFTER `recommend`;
ALTER TABLE `article` ADD COLUMN `html`  longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'html' AFTER `desc`;
ALTER TABLE `article` ADD COLUMN `cover` varchar(255) DEFAULT NULL;
ALTER TABLE `article` ADD COLUMN `url` varchar(255) DEFAULT NULL;
ALTER TABLE `article` ADD COLUMN `other` longtext DEFAULT NULL;
-- --------------------------------------------------------

--
-- 表的结构 `category`
--

CREATE TABLE `category` (
  `id` int(11) NOT NULL COMMENT '主键',
  `name` varchar(255) DEFAULT NULL,
  `pid` int(11) DEFAULT '0' COMMENT '父ID',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` int(11) DEFAULT '1' COMMENT '状态1正常，2删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `log`
--

CREATE TABLE `log` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `ip` varchar(50) NOT NULL,
  `city` varchar(50) NOT NULL,
  `create` datetime NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `page` varchar(255) DEFAULT NULL,
  `uri` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `message`
--

CREATE TABLE `message` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `review` varchar(500) NOT NULL DEFAULT '' COMMENT '评论',
  `reply` varchar(500) NOT NULL COMMENT '回复',
  `site` varchar(500) NOT NULL COMMENT '网址',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '回复时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `review`
--

CREATE TABLE `review` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `review` varchar(500) NOT NULL DEFAULT '' COMMENT '评论',
  `reply` varchar(500) NOT NULL COMMENT '回复',
  `site` varchar(500) NOT NULL COMMENT '网址',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '回复时间',
  `status` int(11) DEFAULT '1' COMMENT '1可用，2禁用',
  `article_id` int(11) DEFAULT '1' COMMENT '文章ID',
  `customer_id` int(11) DEFAULT '1' COMMENT '用户ID',
  `like`  int(11) NOT NULL DEFAULT 0 COMMENT '点赞数量',
  `star` int(11) DEFAULT '0' COMMENT '评分',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL COMMENT 'ID',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `status` int(11) DEFAULT '1' COMMENT '状态'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `setting` (
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `value` text NOT NULL COMMENT '详情',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `wechat_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `openid` varchar(32) NOT NULL COMMENT 'open ID',
  `nickname` varchar(255) DEFAULT NULL COMMENT '用户名',
  `sex` int(11) DEFAULT '0' COMMENT '性别',
  `city` varchar(32) DEFAULT NULL COMMENT '城市',
  `country` varchar(32) DEFAULT NULL COMMENT '国家',
  `province` varchar(32) DEFAULT NULL COMMENT '省份',
  `language` varchar(32) DEFAULT NULL COMMENT '语种',
  `headimgurl` varchar(255) DEFAULT NULL COMMENT '头像',
  `subscribe_time` int(11) DEFAULT NULL COMMENT '订阅时间',
  `created` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '栏目名称',
  `target` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '链接打开方式',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '链接URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  `pid` int(11) DEFAULT '0' COMMENT '顶级栏目',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `link` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '名称',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT 'URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `file` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '标题',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT 'URL',
  `sort` int(11) DEFAULT '100' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

--
-- 转储表的索引
--

--
-- 表的索引 `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `log`
--
ALTER TABLE `log`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `message`
--
ALTER TABLE `message`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `article`
--
ALTER TABLE `article`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `category`
--
ALTER TABLE `category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `log`
--
ALTER TABLE `log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `message`
--
ALTER TABLE `message`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `review`
--
ALTER TABLE `review`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID';

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID';
COMMIT;

INSERT INTO `user` VALUES (1, 'user', '5c0b8081c10ad236fa004adfe685867f', '491126240@qq.com', '2020-2-11 12:22:55', 1);

INSERT INTO `category` (`id`, `name`, `pid`, `sort`, `status`) VALUES
(1, '随笔', 0, 100, 1);

INSERT INTO `article` (`id`, `user_id`, `title`, `category_id`, `tag`, `remark`, `desc`, `pv`, `created`, `updated`, `status`, `review`, `recommend`, `like`,`html`) VALUES
(1, 1, '这是我的第一篇博客', 1, '博客，Go Blog', '这是我的第一篇博客', '## 这是我的第一篇博客', 1, '2020-02-12 23:07:52', '2020-02-13 05:03:22', 1, 0, 0, 0,'<h2>这是我的第一篇博客</h2>');

INSERT INTO `menu` (`id`, `title`, `target`, `url`, `sort`, `pid`) VALUES
(1, '首页', 'none', '/', 100, 0);
INSERT INTO `menu` (`id`, `title`, `target`, `url`, `sort`, `pid`) VALUES
(2, '留言', 'none', '/message.html', 100, 0);
INSERT INTO `menu` (`id`, `title`, `target`, `url`, `sort`, `pid`) VALUES
(3, '关于我', 'none', '/about.html', 100, 0);


INSERT INTO `link` (`id`, `title`, `url`, `sort`) VALUES
(1, 'Go Blog', 'http://go-blog.cn', 100);
INSERT INTO `link` (`id`, `title`, `url`, `sort`) VALUES
(2, 'LeeChan\'Blog', 'http://leechan.online', 100);

INSERT INTO `setting` VALUES 
('about_html_code','<h2>欢迎来到Go Blog!<h2>'),
('about_markdown_doc','## 欢迎来到Go Blog!'),
('description','Go Blog 一个使用 Beego 开发的博客系统'),
('image','/static/uploads/2020021121190681.png'),('keyword','Go Blog 一个使用 Beego 开发的博客系统'),
('limit','20'),('name','123'),
('notice','欢迎来到使用 Go Blog 。'),
('remark_html_code','欢迎来到使用 Go Blog '),
('remark_markdown_doc','欢迎来到使用 Go Blog '),
('tag','欢迎来到使用 Go Blog '),
('template','leechen'),
('title','Go Blog 一个使用 Beego 开发的博客系统');

CREATE TABLE `customer` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` varchar(50) NOT NULL COMMENT '用户ID',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '昵称',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '头像',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '博客地址',
  `signature` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '个性签名',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '邮箱',
  `phone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '电话',
  `wishlist` int DEFAULT '0' COMMENT '收藏',
  `review` int DEFAULT '0' COMMENT '评论',
  `like` int DEFAULT '0' COMMENT '点赞',
  `status` int DEFAULT '1' COMMENT '1可用，2禁用，0删除',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1146 DEFAULT CHARSET=utf8;


CREATE TABLE `ad` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '名称',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '图片链接',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT 'URL',
  `gid` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '组别ID',
  `group` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '组别名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

