/*
 Navicat Premium Data Transfer

 Source Server         : mysql_localhost
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : hk

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 07/08/2022 17:06:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hk_user
-- ----------------------------
DROP TABLE IF EXISTS `hk_user`;
CREATE TABLE `hk_user`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户编号',
   `user_type` int(2) NOT NULL DEFAULT 2 COMMENT '用户平台: 1web、2app、3other',
   `account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账号',
   `password` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
   `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
   `real_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真名',
   `header_img` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '头像',
   `email` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
   `phone` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
   `birthday` datetime(3) NULL DEFAULT NULL COMMENT '生日',
   `sex` int(2) NOT NULL DEFAULT 0 COMMENT '性别： 0未知、1男、2女',
   `description` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',   
   `role_id` bigint(20) UNSIGNED NULL DEFAULT 888 COMMENT '用户角色ID',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE INDEX `idx_hk_user_account`(`tenant_id`,`account`) USING BTREE,
   INDEX `idx_hk_user_uuid`(`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '终端用户';

-- ----------------------------
-- Records of hk_user
-- ----------------------------
INSERT INTO `hk_user` VALUES (1, '000000', '0acdb369-4e8d-49dc-8f34-352a5079a43e', 2, 'test', '$2a$10$xlA24ZuTjCjtG.zZYIVsJeT1/IQ2fItgLH4fLe4DNLMS6Hn9U.ahS', '用户1', '', '', '', '', NULL, 0, '描述信息',888, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user` VALUES (2, '000000', '0acdb369-4e8d-49dc-8f34-352a5079a43f', 2, 'admin', '$2a$10$oC3wU0ZeutgWib.oh5rmM.hVtdeqy8/mhJh3bIGP0lP2DTg/dMmiK', '范杰2', '', '', '', '', NULL, 1, '描述信息',888, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user` VALUES (3, '000001', '5da4fe91-b5e4-47aa-ab1f-d49df3415447', 2, 'test', '$2a$10$xlA24ZuTjCjtG.zZYIVsJeT1/IQ2fItgLH4fLe4DNLMS6Hn9U.ahS', '用户3', '', '', '', '', NULL, 1, '描述信息',888, '2023-02-08 14:32:12.598', '2023-02-08 14:32:12.598', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user` VALUES (4, '000000', '7d43462c-a441-46e8-8981-2d7c332268b9', 2, 'login', '$2a$10$0PJORNdD/HynCQEtePuTAuvPQSsHneAfXzpNIJHX2cFw2JHZBwM8G', '用户4', '', '', '', '', NULL, 2, '描述信息',888, '2023-02-08 14:32:19.098', '2023-02-08 14:32:19.098', NULL, NULL, NULL, NULL, 0, 0);


-- ----------------------------
-- Table structure for hk_user_extend
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_extend`;
CREATE TABLE `hk_user_extend`  (
   `id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `circle_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '当前圈子编号',
   `github` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL comment 'github',
   `weibo` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL comment '微博',
   `weixin` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL comment '微信',
   `wx_profile` json DEFAULT NULL COMMENT '微信用户json信息',
   `qq`varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL comment 'qq',
   `blog` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL comment '博客',
   `num_circle` bigint(20) NOT NULL DEFAULT 0 comment '圈子数',
   `num_focus` bigint(20) NOT NULL DEFAULT 0 comment '关注数',
   `num_fan` bigint(20) NOT NULL DEFAULT 0 comment '粉丝数',
   `now_money` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '用户余额',
   `integral` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '用户剩余积分',
   `sign_num` int NOT NULL DEFAULT '0' COMMENT '连续签到天数',
   `spread_uid` bigint(20) unsigned DEFAULT '0' COMMENT '推广元id',
   `spread_time` datetime DEFAULT NULL COMMENT '推广员关联时间',
   `is_promoter` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否为推广员',
   `spread_count` int DEFAULT '0' COMMENT '下级人数',
   `add_ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '添加ip',
   `last_ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '最后一次登录ip',
   `channel_id` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '频道_编号',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '终端用户拓展';

-- ----------------------------
-- Table structure for hk_user_address
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_address`;
CREATE TABLE `hk_user_address` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户ID',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收货人姓名',
  `phone` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收货人电话',
  `province` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收货人所在省',
  `city` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收货人所在市',
  `city_id` int DEFAULT NULL,
  `district` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收货人所在区',
  `detail` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '收货人详细地址',
  `post_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮编',
  `longitude` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '经度',
  `latitude` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '纬度',
  `is_default` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否默认',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime NOT NULL COMMENT '添加时间',
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `is_default` (`is_default`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户地址表';

-- ----------------------------
-- Table structure for hk_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_authority`;
CREATE TABLE `hk_user_authority`  (
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) UNSIGNED NOT NULL,
   `role_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for hk_user_oauth
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_oauth`;
CREATE TABLE `hk_user_oauth`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '第三方系统用户ID',
   `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户ID',
   `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账号',
   `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
   `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
   `blog` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '应用主页',
   `company` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名',
   `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
   `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮件',
   `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
   `gender` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别 1男 2女 3未知',
   `source` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '终端用户第三方认证';

-- ----------------------------
-- Table structure for hk_circle_classify
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_classify`;
CREATE TABLE `hk_circle_classify`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父节点编号',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `des` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
   `property` int(2) NOT NULL DEFAULT 0 COMMENT '属性：0公开 ，1受限',
   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子分类';

-- ----------------------------
-- Records of hk_circle_classify
-- ----------------------------
INSERT INTO `hk_circle_classify` VALUES (1, '000000', NULL, '分类1', NULL, 0, 1, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (2, '000000', NULL, '分类2', NULL, 0, 2, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (3, '000000', NULL, '分类3', NULL, 0, 3, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (4, '000000', NULL, '分类4', NULL, 0, 4, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (5, '000000', NULL, '分类5', NULL, 0, 5, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (6, '000000', NULL, '分类6', NULL, 0, 6, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (7, '000000', NULL, '分类7', NULL, 0, 7, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (8, '000000', NULL, '分类8', NULL, 0, 8, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_classify` VALUES (9, '000000', NULL, '分类9', NULL, 0, 9, '2023-02-08 13:42:55.885', '2023-02-08 13:42:55.885', NULL, NULL, NULL, NULL, 0, 0);


-- ----------------------------
-- Table structure for hk_circle
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle`;
CREATE TABLE `hk_circle`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `type` int(2) NOT NULL DEFAULT 0 COMMENT '类型：0官方圈子、1用户圈子、2小区',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
   `logo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子Logo',
   `circle_classify_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '圈子分类_编号',
   `slogan` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子标语',
   `des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子简介',
   `protocol` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子规约',
   `back_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子背景图',
   `support_category` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '[1,2,3,4,5,6,7]' COMMENT '支持内容类别(json数组)：1视频、2动态、3资讯、4公告、5文章、6问答、7活动',
   `pay` int(2) NOT NULL DEFAULT 0 COMMENT '付费：0 否，1是',
   `process` int(2) NOT NULL DEFAULT 0 COMMENT '是否开启版块内容人工审核：0 否，1是',
   `property` int(2) NOT NULL DEFAULT 0 COMMENT '圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）',
   `view` int(2) NOT NULL DEFAULT 0 COMMENT '板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到',
   `power_add` int(2) NOT NULL DEFAULT 0 COMMENT '圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户',
   `power_view` int(2) NOT NULL DEFAULT 0 COMMENT '圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组',
   `power_publish` int(2) NOT NULL DEFAULT 0 COMMENT '圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组',
   `power_comment` int(2) NOT NULL DEFAULT 0 COMMENT '圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组',
   `power_add_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子加入权限_指定部门和成员(json数组)',
   `power_view_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子内浏览权限_指定部门和用户(json数组)',
   `power_publish_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子内发布权限_指定部门和用户(json数组)',
   `power_comment_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子内评论权限_指定部门和用户(json数组)',
   `no_limit_user_group` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '不受限用户组(json数组)',  
   `new_user_focus` int(2) NOT NULL DEFAULT 0 COMMENT '新注册用户默认关注：0 否，1是',
   `channel_id` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '频道_编号',
   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子';
-- ----------------------------
-- Records of hk_circle
-- ----------------------------
INSERT INTO `hk_circle` VALUES (1, '000000', 0, '圈子名称1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, '圈子标语内容', '圈子简介内容', NULL, NULL, NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 1, '2023-02-23 10:16:26.493', '2023-02-23 10:16:26.493', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle` VALUES (2, '000000', 0, '圈子名称2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, '圈子标语内容', '圈子简介内容', NULL, NULL, NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 2, '2023-02-23 10:16:26.493', '2023-02-23 10:16:26.493', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle` VALUES (3, '000000', 0, '圈子名称3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, '圈子标语内容', '圈子简介内容', NULL, NULL, NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 3, '2023-02-23 10:16:26.493', '2023-02-23 10:16:26.493', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_circle_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_request`;
CREATE TABLE `hk_circle_request`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `type` int(2) NOT NULL DEFAULT 1 COMMENT '类型：0官方圈子 ，1用户圈子',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
  `logo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子Logo',
  `circle_classify_id` bigint(20) NOT NULL COMMENT '圈子分类_编号',
  `slogan` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子标语',
  `des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子简介',
  `protocol` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子规约',
  `back_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子背景图',
  `check_status` smallint(6) NOT NULL DEFAULT 0 COMMENT '审核状态：0 未处理 1 通过，2驳回',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` int(2) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_request_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子申请' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_circle_request
-- ----------------------------
INSERT INTO `hk_circle_request` VALUES (1, '000000', 1, 'aa', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, NULL, NULL, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_circle_add_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_add_request`;
CREATE TABLE `hk_circle_add_request`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户ID',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '申请理由',
   `check_user` bigint(20) NULL DEFAULT NULL COMMENT '审核人',
   `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
   `check_status` int(2) NULL DEFAULT 0 COMMENT '审核状态：0 未处理、1通过、2驳回',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '加入圈子申请';

-- ----------------------------
-- Table structure for hk_circle_relation
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_relation`;
CREATE TABLE `hk_circle_relation`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `relation_type` int(2) NULL DEFAULT NULL COMMENT '关系类型：0父子节点 1关注',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `other_circle_id` bigint(20) NOT NULL COMMENT '关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子关系';

-- ----------------------------
-- Table structure for hk_circle_user
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_user`;
CREATE TABLE `hk_circle_user`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户ID',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `power` int(2) NULL DEFAULT 0 COMMENT '权限：0普通 1圈主',
   `sort` int(11) NULL DEFAULT NULL COMMENT '用户的圈子排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子成员';

-- ----------------------------
-- Records of hk_circle_user
-- ----------------------------
INSERT INTO `hk_circle_user` VALUES (1, '000000', 1, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (2, '000000', 2, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (3, '000000', 3, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (4, '000000', 2, 2, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (5, '000000', 3, 2, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (6, '000000', 3, 3, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_channel
-- ----------------------------
DROP TABLE IF EXISTS `hk_channel`;
CREATE TABLE `hk_channel`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
   `code` int(11) NULL DEFAULT 0 COMMENT '标识：0常规、1资讯、2关注、3附近、4活动、5问答',
   `fixed_type` tinyint(1) NULL DEFAULT NULL COMMENT '固定标识(不固定:0，固定:1)',
   `type` tinyint(1) NULL DEFAULT NULL COMMENT '类型：0系统 1自定义',
   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE INDEX `idx_hk_channel_name`(`name`) USING BTREE,
   UNIQUE INDEX `idx_hk_channel_code`(`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '频道' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for hk_circle_channel
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_channel`;
CREATE TABLE `hk_circle_channel`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `circle_id` bigint(20) NULL DEFAULT NULL COMMENT '圈子_编号',
   `channel_id` bigint(20) NULL DEFAULT NULL COMMENT '频道_编号',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子频道' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for hk_user_channel
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_channel`;
CREATE TABLE `hk_user_channel`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户ID',
   `channel_id` bigint(20) NULL DEFAULT NULL COMMENT '频道_编号',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户频道' ROW_FORMAT = Dynamic;



-- ----------------------------
-- Table structure for hk_forum_posts
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_posts`;
CREATE TABLE `hk_forum_posts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `circle_id` bigint(20) NULL DEFAULT NULL COMMENT '圈子_编号',
  `category` smallint(6) NULL DEFAULT NULL COMMENT '类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动',
  `channel_id` bigint(20) NULL DEFAULT NULL COMMENT '频道_编号',
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `seo_key` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO关键词',
  `seo_introduce` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介(SEO简介)',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `source` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源',
  `content_type` smallint(6) NULL DEFAULT NULL COMMENT '内容类型：1markdown、2html',
  `content_markdown` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'markdown内容',
  `content_html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'html内容',
  `video` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '视频地址',
  `attachment` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附件',
  `tag` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `top` smallint(6) NULL DEFAULT NULL COMMENT '置顶：0否、1是',
  `marrow` smallint(6) NULL DEFAULT NULL COMMENT '精华：0否、1是',
  `comment_id` bigint(20) NULL DEFAULT NULL COMMENT '问答最佳答案ID',
  `anonymity` smallint(6) NULL DEFAULT NULL COMMENT '匿名发布：0否、1是',
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '发布者编号',
  `read_num` smallint(6) NULL DEFAULT NULL COMMENT '阅读次数',
  `comment_num` smallint(6) NULL DEFAULT NULL COMMENT '评论次数',
  `share_num` smallint(6) NULL DEFAULT NULL COMMENT '分享次数',
  `collect_num` smallint(6) NULL DEFAULT NULL COMMENT '收藏次数',
  `like_num` smallint(6) NULL DEFAULT NULL COMMENT '点赞次数',
  `check_user` bigint(20) NULL DEFAULT NULL COMMENT '审核人',
  `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
  `check_status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态：1草稿、2未审批、3通过、4拒绝',
  `power_comment` smallint(6) NULL DEFAULT NULL COMMENT '评论权限：0关闭、1开启',
  `power_comment_anonymity` smallint(6) NULL DEFAULT NULL COMMENT '匿名评论：0关闭、1开启',
  `pay` smallint(6) NULL DEFAULT NULL COMMENT '付费：0否、1是',
  `pay_content` smallint(6) NULL DEFAULT NULL COMMENT '内容付费：0否、1是',
  `pay_content_look` smallint(6) NULL DEFAULT NULL COMMENT '内容付费可查看百分比例',
  `pay_attachment` smallint(6) NULL DEFAULT NULL COMMENT '附件付费：0否、1是',
  `pay_currency` smallint(6) NULL DEFAULT NULL COMMENT '付费货币：1人民、2积分、3代币',
  `pay_num` smallint(6) NULL DEFAULT NULL COMMENT '付费金额',
  `activity_start_at` datetime(3) NULL DEFAULT NULL COMMENT '活动开始时间',
  `activity_end_at` datetime(3) NULL DEFAULT NULL COMMENT '活动结束时间',
  `activity_address` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '活动地址',
  `activity_user_num` int(11) NULL DEFAULT 0 COMMENT '活动用户人数（0不限制活动人数，否则为活动人数）',
  `activity_cur_user_num` int(11) NULL DEFAULT 0 COMMENT '活动参加人数',
  `activity_chat_group_id` bigint(20) NOT NULL COMMENT '活动聊天群编号',
  `is_public` int(2) NULL DEFAULT 0 COMMENT '是否公开：0 否 1 是',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `thumbs_up` smallint(6) NULL DEFAULT 0 COMMENT '是否点赞：0否、1是',
  `collect` smallint(6) NULL DEFAULT 0 COMMENT '是否收藏：0否、1是',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_posts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帖子' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_forum_posts
-- ----------------------------
INSERT INTO `hk_forum_posts` VALUES (1, '000000', 1, 2, 0, '圈子1_标题（用户1）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, NULL, 2, 0, 0, 0, 0, 0, 0, 2, 0, '2023-02-15 20:29:57.000', '2023-02-15 20:29:57.000', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (2, '000000', 1, 2, 0, '圈子1_标题（用户2）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 2, 0, 0, 0, 0, 0, 0, 2, 0, '2023-02-15 20:29:57.000', '2023-02-15 20:29:57.000', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (3, '000000', 1, 2, 0, '圈子1_标题（用户3）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, NULL, 2, 0, 0, 0, 0, 0, 0, 2, 0, '2023-02-15 20:29:57.000', '2023-02-15 20:29:57.000', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (4, '000000', 2, 2, 0, '圈子2_标题（用户2）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 2, 0, 0, 0, 0, 0, 0, 2, 0, '2023-02-15 20:29:57.000', '2023-02-15 20:29:57.000', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (5, '000000', 2, 2, 0, '圈子2_标题（用户3）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, NULL, 2, 0, 0, 0, 0, 0, 0, 2, 0, '2023-02-15 20:29:57.000', '2023-02-15 20:29:57.000', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (6, '000000', 3, 2, 0, '圈子3_标题（用户3）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, NULL, 2, 0, 0, 0, 0, 0, 0, 2, 0, '2023-02-15 20:29:57.000', '2023-02-15 20:29:57.000', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (7, '000000', 1, 0, 0, 'testtest', '', '', '', '', 0, '', '', '', '', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-02-18 11:41:17.115', '2023-02-18 11:41:17.115', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (8, '000000', 1, 1, 0, '有附件', '', '', '', '', 1, '', '有附件', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker4029652423983233700.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-20/image_picker4029652423983233700.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-20/image_picker4029652423983233700.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-02-20 09:52:01.938', '2023-02-20 09:52:01.938', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (9, '000000', 1, 1, 0, '两个附件', '', '', '', '', 1, '', '两个附件', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-02-20 10:06:26.981', '2023-02-24 14:40:05.945', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (10, '000000', 1, 1, 0, '两个附件2', '', '', '', '', 1, '', '两个附件2', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker9222995987695013449.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-20/image_picker9222995987695013449.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-20/image_picker9222995987695013449.jpg\"},{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker6600262919828096699.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-20/image_picker6600262919828096699.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-20/image_picker6600262919828096699.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-02-20 10:07:13.517', '2023-02-24 14:39:58.682', NULL, NULL, NULL, NULL, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (11, '000000', 1, 1, 0, '今天没赚钱顶级精明和李敏李丽丽你家里', '', '', '', '', 1, '', '今天没赚钱顶级精明和李敏李丽丽你家里', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker4577751758068097584.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-21/image_picker4577751758068097584.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-21/image_picker4577751758068097584.jpg\"},{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker7462903158421350833.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-21/image_picker7462903158421350833.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-21/image_picker7462903158421350833.jpg\"},{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2631187621018139806.png\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-21/image_picker2631187621018139806.png\",\"tag\":\"png\",\"key\":\"test/uploads/2023-02-21/image_picker2631187621018139806.png\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 1, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '2023-02-21 17:42:57.034', '2023-02-27 11:15:12.257', NULL, NULL, NULL, NULL, 0, 0, 0, 0);


-- ----------------------------
-- Table structure for hk_forum_comment 修改：删除comment_time
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_comment`;
CREATE TABLE `hk_forum_comment`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子编号',
   `user_id` bigint(20) NOT NULL COMMENT '评论者',
   `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '评论内容',
   `like_num` int(11) NULL DEFAULT 0 COMMENT '点赞数',
   `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父评论编号',
   `check_user` bigint(20) NULL DEFAULT NULL COMMENT '审核人',
   `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
   `check_status` int(2) NULL DEFAULT NULL COMMENT '审核状态(0未审批 1通过 2拒绝)',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帖子评论';

-- ----------------------------
-- Records of hk_forum_comment
-- ----------------------------
INSERT INTO `hk_forum_comment` VALUES (1, '000000', 4, 2, '评论内容', NULL, NULL, NULL, NULL, NULL, '2023-02-18 13:33:43.000', '2023-02-18 13:33:43.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_forum_thumbs_up 修改：删除 time
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_thumbs_up`;
CREATE TABLE `hk_forum_thumbs_up`  (
   `user_id` bigint(20) NOT NULL COMMENT'用户id',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子编号',
   PRIMARY KEY (`user_id`, `posts_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帖子点赞' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_forum_thumbs_up
-- ----------------------------
INSERT INTO `hk_forum_thumbs_up` VALUES (2, 4);

-- ----------------------------
-- Table structure for hk_comment_thumbs_up 修改：删除 time
-- ----------------------------
DROP TABLE IF EXISTS `hk_comment_thumbs_up`;
CREATE TABLE `hk_comment_thumbs_up`  (
   `user_id` bigint(20) NOT NULL COMMENT'用户id',
   `comment_id` bigint(20) NOT NULL COMMENT '评论编号',
   PRIMARY KEY (`user_id`, `comment_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '评论点赞' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_comment_thumbs_up
-- ----------------------------
INSERT INTO `hk_comment_thumbs_up` VALUES (2, 1);

-- ----------------------------
-- Table structure for hk_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_apply`;
CREATE TABLE `hk_apply`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `ower_type` int(2) NULL DEFAULT NULL COMMENT '拥有者：0平台、1圈子、2个人',
   `circle_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '圈子_编号',
   `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户_编号',
   `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
   `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `type` int(2) NOT NULL DEFAULT 0 COMMENT '类型(0小程序、1第三方链接)',
   `mini_program_id` bigint(20) NULL DEFAULT NULL COMMENT '小程序id',
   `program_id` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '小程序id',
   `apply_address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问地址',
   `apply_parameters` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问参数',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '应用';

-- ----------------------------
-- Records of hk_apply
-- ----------------------------
INSERT INTO `hk_apply` VALUES (1, '000000', 0, 0, 0, '应用1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230227/42ebfc53bfcbcf1be1b05b5203c63d0f.png', 1, 1, NULL, NULL, 'https://www.baidu.com/', NULL, '2023-02-20 10:58:15.000', '2023-02-27 16:54:42.703', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (2, '000000', 0, 0, 0, '应用2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (3, '000000', 0, 0, 0, '应用3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (4, '000000', 0, 0, 0, '应用4', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (5, '000000', 0, 0, 0, '应用5', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (6, '000000', 0, 0, 0, '应用6', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (11, '000000', 0, 0, 0, '应用1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (12, '000000', 0, 0, 0, '应用2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (13, '000000', 0, 0, 0, '应用3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (14, '000000', 0, 0, 0, '应用4', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (15, '000000', 0, 0, 0, '应用5', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (16, '000000', 0, 0, 0, '应用6', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (21, '000000', 0, 0, 0, '应用1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (22, '000000', 0, 0, 0, '应用2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (23, '000000', 0, 0, 0, '应用3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (24, '000000', 0, 0, 0, '应用4', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (25, '000000', 0, 0, 0, '应用5', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (26, '000000', 0, 0, 0, '应用6', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, '__UNI__A0B6451', NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

-- ----------------------------
-- Table structure for hk_plat_apply_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_plat_apply_group`;
CREATE TABLE `hk_plat_apply_group`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',   
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
   `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父节点编号',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台应用分组';

-- ----------------------------
-- Table structure for hk_plat_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_plat_apply`;
CREATE TABLE `hk_plat_apply`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `apply_group_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '应用分组_编号',
   `show_name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称别名',
   `apply_id` bigint(20) NULL DEFAULT NULL COMMENT '应用_编号',
   `apply_parameters` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问参数',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '平台应用';

-- ----------------------------
-- Table structure for hk_circle_apply_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_apply_group`;
CREATE TABLE `hk_circle_apply_group`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',   
   `circle_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '圈子_编号',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
   `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父节点编号',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子应用分组';

-- ----------------------------
-- Records of hk_circle_apply_group
-- ----------------------------
INSERT INTO `hk_circle_apply_group` VALUES (11, '000000', 1, '分组1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 0, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (12, '000000', 1, '分组2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 11, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (13, '000000', 1, '分组3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 11, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (21, '000000', 2, '分组1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 0, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (22, '000000', 2, '分组2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 21, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (23, '000000', 2, '分组3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 21, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (31, '000000', 3, '分组1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 0, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (32, '000000', 3, '分组2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 31, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply_group` VALUES (33, '000000', 3, '分组3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 31, 0, '2023-02-21 15:26:57.441', '2023-02-21 15:26:57.441', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

-- ----------------------------
-- Table structure for hk_circle_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_apply`;
CREATE TABLE `hk_circle_apply`  (
   `id` bigint(20)UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `circle_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '圈子_编号',
   `apply_group_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '应用分组_编号',
   `apply_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '应用_编号',
   `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_hk_circle_apply_apply_id`(`circle_id`,`apply_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子应用';

-- ----------------------------
-- Records of hk_circle_apply
-- ----------------------------
INSERT INTO `hk_circle_apply` VALUES (1, '000000', 1, 11, 1, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (2, '000000', 1, 11, 2, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (3, '000000', 1, 12, 3, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (4, '000000', 1, 12, 4, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (5, '000000', 1, 12, 5, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (6, '000000', 1, 13, 6, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

INSERT INTO `hk_circle_apply` VALUES (11, '000000', 1, 21, 11, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (12, '000000', 1, 21, 12, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (13, '000000', 1, 22, 13, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (14, '000000', 1, 22, 14, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (15, '000000', 1, 22, 15, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (16, '000000', 1, 23, 16, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

INSERT INTO `hk_circle_apply` VALUES (21, '000000', 1, 31, 21, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (22, '000000', 1, 31, 22, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (23, '000000', 1, 32, 23, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (24, '000000', 1, 32, 24, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (25, '000000', 1, 32, 25, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (26, '000000', 1, 33, 26, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

-- ----------------------------
-- Table structure for hk_user_circle_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_circle_apply`;
CREATE TABLE `hk_user_circle_apply`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `apply_id` bigint(20) NULL DEFAULT NULL COMMENT '应用_编号',
   `sort` int(2) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户圈子应用';

-- ----------------------------
-- Table structure for hk_forum_topic_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_topic_group`;
CREATE TABLE `hk_forum_topic_group`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `circle_id` bigint(20) NULL DEFAULT NULL COMMENT '圈子_编号',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '话题分组';

-- ----------------------------
-- Records of hk_forum_topic_group
-- ----------------------------
INSERT INTO `hk_forum_topic_group` VALUES (1, '000000', NULL, 'topic_1组', 1, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (2, '000000', NULL, 'topic_2组', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (3, '000000', 1, 'topic_1组', 1, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (4, '000000', 1, 'topic_2组', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_forum_topic
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_topic`;
CREATE TABLE `hk_forum_topic`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
   `topic_group_id` bigint(20) NOT NULL COMMENT'分组id',
   `type` int(2) NOT NULL DEFAULT 0 COMMENT '话题类型：0 全局，1圈子',
   `hot` int(2) NULL DEFAULT 0 COMMENT '是否热门：0 否 1 是',
   `discuss_num` int(11) NULL DEFAULT NULL COMMENT '讨论数',
   `attention_num` int(11) NULL DEFAULT NULL COMMENT '关注数',
   `circle_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '圈子_编号',
   `review_status` int(2) NULL DEFAULT 0 COMMENT '审核状态：0 未处理 1 通过，2驳回',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '话题';

-- ----------------------------
-- Records of hk_forum_topic
-- ----------------------------
INSERT INTO `hk_forum_topic` VALUES (1, '000000', 'topic_1_1', NULL, 1, 0, NULL, NULL, NULL, 0, 0, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic` VALUES (2, '000000', 'topic_1_2', NULL, 1, 0, NULL, NULL, NULL, 0, 0, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_forum_topic_posts_mapping
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_topic_posts_mapping`;
CREATE TABLE `hk_forum_topic_posts_mapping`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `topic_id` bigint(20) NOT NULL COMMENT '话题id',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子id',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '话题-帖子关系表';

-- ----------------------------
-- Records of hk_forum_topic_posts_mapping
-- ----------------------------
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (1, '000000', 1, 1, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (2, '000000', 1, 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (3, '000000', 1, 3, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (4, '000000', 1, 4, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (5, '000000', 1, 5, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (6, '000000', 1, 6, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (7, '000000', 2, 1, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (8, '000000', 2, 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (9, '000000', 2, 3, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (10, '000000', 2, 4, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (11, '000000', 2, 5, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_posts_mapping` VALUES (12, '000000', 2, 6, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_forum_tag_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_tag_group`;
CREATE TABLE `hk_forum_tag_group`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签分组';

-- ----------------------------
-- Table structure for hk_forum_tag
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_tag`;
CREATE TABLE `hk_forum_tag`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
   `group_id` bigint(20) NOT NULL COMMENT'分组id',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签';

-- ----------------------------
-- Table structure for hk_forum_tag_posts_mapping
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_tag_posts_mapping`;
CREATE TABLE `hk_forum_tag_posts_mapping`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `tag_id` bigint(20) NOT NULL COMMENT '标签ID',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子ID',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签-帖子关系表';

-- ----------------------------
-- Table structure for hk_forbidden_speak_reason
-- ----------------------------
DROP TABLE IF EXISTS `hk_forbidden_speak_reason`;
CREATE TABLE `hk_forbidden_speak_reason`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `reason` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '禁言理由',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '禁言原因';

-- ----------------------------
-- Table structure for hk_forbidden_speak_duration
-- ----------------------------
DROP TABLE IF EXISTS `hk_forbidden_speak_duration`;
CREATE TABLE `hk_forbidden_speak_duration`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `count` int(11) NULL DEFAULT NULL COMMENT '时长',
   `type` int(2) NULL DEFAULT NULL COMMENT '时间类型：0 小时、1天',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '禁言时长';

-- ----------------------------
-- Table structure for hk_forbidden_speak
-- ----------------------------
DROP TABLE IF EXISTS `hk_forbidden_speak`;
CREATE TABLE `hk_forbidden_speak`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `duration_id` bigint(20) NOT NULL COMMENT '禁言时长_编号',
   `reason_id` bigint(20) NOT NULL COMMENT '禁言原因_编号',
   `cur_status` int(2) NULL DEFAULT 0 COMMENT '当前状态：0 未解锁、1已解锁',
   `unlock_user_id` bigint(20) NULL DEFAULT NULL COMMENT '解锁用户_编号',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '禁言';

-- ----------------------------
-- Table structure for hk_report_reason
-- ----------------------------
DROP TABLE IF EXISTS `hk_report_reason`;
CREATE TABLE `hk_report_reason`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `reason` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报理由',
   `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '举报原因';


-- ----------------------------
-- Table structure for hk_report
-- ----------------------------
DROP TABLE IF EXISTS `hk_report`;
CREATE TABLE `hk_report`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `report_user_id` bigint(20) NOT NULL COMMENT '被举报用户编号',
   `user_id` bigint(20) NOT NULL COMMENT '举报用户编号',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `reason_id` bigint(20) NOT NULL COMMENT '举报原因_编号',
   `report_type` int(2) NOT NULL COMMENT '举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题',
   `content` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报内容',
   `content_attachment` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '内容附件',
   `cur_status` int(2) NULL DEFAULT NULL COMMENT '处理状态：0 未处理、1 处理中、2 拒绝、3 完成',
   `handle_user_id` bigint(20) NOT NULL COMMENT '操作用户_编号',
   `handle_type` bigint(20) NULL DEFAULT NULL COMMENT '处理方式:0无效处理（不予处理）、账号禁言',
   `duration_id` bigint(20) NULL DEFAULT NULL COMMENT '禁言时长_编号',
   `handle_score` int(2) NULL DEFAULT NULL COMMENT '积分处理:0不扣分、1扣分',
   `score_experience` int(11) NULL DEFAULT NULL COMMENT '经验',
   `score_community` int(11) NULL DEFAULT NULL COMMENT '社区积分',
   `score_buy` int(11) NULL DEFAULT NULL COMMENT '购物积分',
   `score_download` int(11) NULL DEFAULT NULL COMMENT '下载值',   
   `unlock` int(2) NULL DEFAULT NULL COMMENT '是否解除：0 否、是',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '举报';

-- ----------------------------
-- 平台业务
-- ----------------------------

-- ----------------------------
-- Table structure for hk_mini_program
-- ----------------------------
DROP TABLE IF EXISTS `hk_mini_program`;
CREATE TABLE `hk_mini_program`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
   `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
   `company_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称',
   `program_id` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '小程序id',
   `cur_packet_id` bigint(20) NULL DEFAULT NULL COMMENT '当前包id',
   `hidden` int(2) NOT NULL DEFAULT 0 COMMENT '隐藏(0否 1是)',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '小程序';

-- ----------------------------
-- Records of hk_mini_program
-- ----------------------------
INSERT INTO `hk_mini_program` VALUES (1627919042591342593, '000000', '小程序名称', NULL, '四川东联汇科','__UNI__A0B6451', 1627927490628435969, 0, '2023-02-21 14:31:59.677', '2023-02-21 14:31:59.677', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_mini_program` VALUES (1629014834567827458, '000000', '宠物猫', NULL, '四川东联汇科', '__UNI__A0B6451', 1630083744511602690, 1, '2023-02-24 15:06:16.815', '2023-02-24 15:06:16.815', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

-- ----------------------------
-- Table structure for hk_mini_program_packet
-- ----------------------------
DROP TABLE IF EXISTS `hk_mini_program_packet`;
CREATE TABLE `hk_mini_program_packet`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `mini_program_id` bigint(20) NOT NULL COMMENT '小程序id',
   `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '包名',
   `version` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '版本',
   `code` bigint(20) NOT NULL COMMENT 'code',
   `packet_address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问地址',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '小程序包';

-- ----------------------------
-- Records of hk_mini_program_packet
-- ----------------------------
INSERT INTO `hk_mini_program_packet` VALUES (1627927490628435969, '000000', 1627919042591342593, '123', 'v2.0',101, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230228/259d4a4197af616373672d2bbc774ff5.wgt', '2023-02-21 15:05:33.000', '2023-02-24 16:41:01.669', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_mini_program_packet` VALUES (1630083744511602690, '000000', 1629014834567827458, 'aaaa', 'v2.1',101, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230228/259d4a4197af616373672d2bbc774ff5.wgt', '2023-02-24 14:43:28.000', '2023-02-24 15:08:26.394', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

-- ----------------------------
-- Table structure for hk_bug_report
-- ----------------------------
DROP TABLE IF EXISTS `hk_bug_report`;
CREATE TABLE `hk_bug_report`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
   `content` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作步骤',
   `content_attachment` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作步骤附件',
   `expected_result` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '预期结果',
   `actual_result` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实际结果',
   `actual_result_attachment` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实际结果附件',
   `other_info` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '其他信息',
   `check_status` int(2) NULL DEFAULT NULL COMMENT '处理状态：0 未处理、1 处理中、2 拒绝、3 完成',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Bug反馈';

-- ----------------------------
-- Table structure for hk_user_cur_circle
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_cur_circle`;
CREATE TABLE `hk_user_cur_circle`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户当前圈子';

-- ----------------------------
-- Table structure for hk_user_browsing_history
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_browsing_history`;
CREATE TABLE `hk_user_browsing_history`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子编号',
   `time` datetime(3) NULL DEFAULT NULL COMMENT '浏览时间',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '浏览历史';

-- ----------------------------
-- Table structure for hk_user_collect 修改：删除time
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_collect`;
CREATE TABLE `hk_user_collect`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子编号',
   `category` smallint(6) NULL DEFAULT NULL COMMENT '类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`user_id`, `posts_id`) USING BTREE,
   UNIQUE INDEX `idx_user_collect_id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '收藏';

-- ----------------------------
-- Table structure for hk_protocol
-- ----------------------------
DROP TABLE IF EXISTS `hk_protocol`;
CREATE TABLE `hk_protocol`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议名称',
   `des` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议说明',
   `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议内容',
   `module` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所属模块/插件',
   `pos` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '显示位置',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE INDEX `idx_hk_protocol_name`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '协议';

-- ----------------------------
-- Records of hk_protocol
-- ----------------------------
INSERT INTO `hk_protocol` VALUES (1, '000000', '用户协议', '用户协议描述', '用户协议内容', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_protocol` VALUES (2, '000000', '隐私政策', '隐私政策秒杀', '隐私政策内容', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_config_region
-- ----------------------------
DROP TABLE IF EXISTS `hk_config_region`;
CREATE TABLE `hk_config_region`  (
  `code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '区划编号',
  `parent_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父区划编号',
  `ancestors` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '祖区划编号',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区划名称',
  `province_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '省级区划编号',
  `province_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '省级名称',
  `city_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '市级区划编号',
  `city_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '市级名称',
  `district_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区级区划编号',
  `district_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区级名称',
  `town_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '镇级区划编号',
  `town_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '镇级名称',
  `village_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '村级区划编号',
  `village_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '村级名称',
  `region_level` int(2) NULL DEFAULT NULL COMMENT '层级',
  `sort` int(2) NULL DEFAULT 0 COMMENT '排序',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '(配置)行政区划表';


-- ----------------------------
-- Table structure for hk_files
-- ----------------------------
DROP TABLE IF EXISTS `hk_files`;
CREATE TABLE `hk_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `file_address` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `file_domain_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '域名',
  `file_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT'md5码',
  `file_original_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '原名',
  `file_extension` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '拓展名',
  `file_chunk_total` bigint(20) NULL DEFAULT NULL COMMENT '块数',
  `file_size` bigint(20) NULL DEFAULT NULL COMMENT '大小',
  `is_finish` tinyint(1) NULL DEFAULT NULL COMMENT '是否上传完',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT NULL COMMENT '状态',
  `is_del` int(2) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for hk_focus_user
-- ----------------------------
DROP TABLE IF EXISTS `hk_focus_user`;
CREATE TABLE `hk_focus_user`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户ID',
   `focus_user_id` bigint(20) NOT NULL COMMENT '关注用户ID',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户关注';

-- ----------------------------
-- Table structure for hk_activity_user 
-- ----------------------------
DROP TABLE IF EXISTS `hk_activity_user`;
CREATE TABLE `hk_activity_user`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `activity_id` bigint(20) NOT NULL COMMENT '活动编号',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `power` int(2) NULL DEFAULT 0 COMMENT '权限：0普通 1管理员 2发起者',
   `pay` int(2) NULL DEFAULT 0 COMMENT '是否付费：0 否 1 是',
   `order_id` bigint(20) NULL DEFAULT NULL COMMENT '订单编号',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `user_activity` (`user_id`, `activity_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '活动用户';

-- ----------------------------
-- Table structure for hk_user_sign
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_sign`;
CREATE TABLE `hk_user_sign` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '签到说明',
   `number` int NOT NULL DEFAULT '0' COMMENT '获得积分',
   `balance` int NOT NULL DEFAULT '0' COMMENT '剩余积分',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='签到记录表';

-- ----------------------------
-- Table structure for hk_user_bill
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_bill`;
CREATE TABLE `hk_user_bill` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `link_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '关联id',
   `pm` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0 = 支出 1 = 获得',
   `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账单标题',
   `category` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '明细种类',
   `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '明细类型',
   `number` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '明细数字',
   `balance` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '剩余',
   `mark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态:0带确定 1有效 -1无效',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `user_id` (`user_id`) USING BTREE,
   KEY `status` (`status`) USING BTREE,
   KEY `add_time` (`created_at`) USING BTREE,
   KEY `pm` (`pm`) USING BTREE,
   KEY `type` (`category`,`type`,`link_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户账单表';

-- ----------------------------
-- Table structure for hk_user_recharge
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_recharge`;
CREATE TABLE `hk_user_recharge` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
   `order_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '订单号',
   `price` decimal(8,2) DEFAULT NULL COMMENT '充值金额',
   `give_price` decimal(8,2) DEFAULT '0.00' COMMENT '购买赠送金额',
   `recharge_type` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '充值类型',
   `paid` tinyint(1) DEFAULT NULL COMMENT '是否充值',
   `pay_time` datetime DEFAULT NULL COMMENT '充值支付时间',
   `refund_price` decimal(10,2) DEFAULT '0.00' COMMENT '退款金额',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态:0带确定 1有效 -1无效',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_id` (`order_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `recharge_type` (`recharge_type`) USING BTREE,
  KEY `paid` (`paid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户充值表';

-- ----------------------------
-- Table structure for hk_user_extract
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_extract`;
CREATE TABLE `hk_user_extract` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `real_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名称',
   `extract_type` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'bank' COMMENT 'bank = 银行卡 alipay = 支付宝wx=微信',
   `bank_code` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '0' COMMENT '银行卡',
   `bank_address` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '开户地址',
   `alipay_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '支付宝账号',
   `wechat` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '微信号',
   `extract_price` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '提现金额',
   `mark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
   `balance` decimal(8,2) unsigned DEFAULT '0.00',
   `fail_msg` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '无效原因',
   `fail_time` datetime DEFAULT NULL,
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态:0审核中 1已提现 2未通过',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `extract_type` (`extract_type`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `add_time` (`created_at`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `fail_time` (`fail_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户提现表';

-- ----------------------------
-- Table structure for hk_order
-- ----------------------------
DROP TABLE IF EXISTS `hk_order`;
CREATE TABLE `hk_order` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
   `order_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单号',
   `extend_order_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '额外订单号',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
   `cart_num` smallint unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
   `total_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '订单总价',
   `pay_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '实际支付金额',
   `paid` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '支付状态',
   `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
   `pay_type` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付方式',
   `refund_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0 未退款 1 申请中 2 已退款',
   `refund_reason_wap_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '退款图片',
   `refund_reason_wap_explain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '退款用户说明',
   `refund_reason_time` datetime DEFAULT NULL COMMENT '退款时间',
   `refund_reason_wap` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '前台退款原因',
   `refund_reason` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '不退款的理由',
   `refund_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '退款金额',
   `gain_integral` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '消费赚取积分',
   `back_integral` decimal(8,2) unsigned DEFAULT NULL COMMENT '给用户退了多少积分',
   `mark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '备注',
   `unique` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一id(md5加密)类似id',
   `remark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '管理员备注',
   `cost` decimal(8,2) unsigned NOT NULL COMMENT '成本价',
   `verify_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '核销码',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态:0待发货 1已完成 2申请退款 3已退货 4已退款 5拒绝退款',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_id_2` (`order_id`,`user_id`) USING BTREE,
  UNIQUE KEY `unique` (`unique`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `pay_price` (`pay_price`) USING BTREE,
  KEY `paid` (`paid`) USING BTREE,
  KEY `pay_time` (`pay_time`) USING BTREE,
  KEY `pay_type` (`pay_type`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='订单表';

-- ----------------------------
-- Table structure for hk_record_browsing_user_homepage
-- ----------------------------
DROP TABLE IF EXISTS `hk_record_browsing_user_homepage`;
CREATE TABLE `hk_record_browsing_user_homepage`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `browser` bigint(20) NOT NULL COMMENT '浏览者编号',
   `browse_time` datetime(3) NULL DEFAULT NULL COMMENT '浏览时间',
   `browse_num` int(11) NULL DEFAULT NULL COMMENT '浏览次数',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '记录浏览用户主页';

-- ----------------------------
-- Table structure for hk_record_browsing_circle_homepage
-- ----------------------------
DROP TABLE IF EXISTS `hk_record_browsing_circle_homepage`;
CREATE TABLE `hk_record_browsing_circle_homepage`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子编号',
   `browser` bigint(20) NOT NULL COMMENT '浏览者编号',
   `browse_time` datetime(3) NULL DEFAULT NULL COMMENT '浏览时间',
   `browse_num` int(11) NULL DEFAULT NULL COMMENT '浏览次数',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '记录浏览圈子主页';