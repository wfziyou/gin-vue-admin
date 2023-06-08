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
   `birthday` varchar(20) NULL DEFAULT NULL COMMENT '生日',
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
   INDEX `idx_hk_user_uuid`(`uuid`) USING BTREE,
   KEY `phone` (`phone`) USING BTREE
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
  `github` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'github',
  `weibo` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微博',
  `weixin` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信',
  `wx_profile` json NULL COMMENT '微信用户json信息',
  `qq` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'qq',
  `blog` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '博客',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `num_circle` bigint(20) NULL DEFAULT NULL COMMENT '圈子数',
  `num_focus` bigint(20) NULL DEFAULT NULL COMMENT '关注数',
  `num_fan` bigint(20) NULL DEFAULT NULL COMMENT '粉丝数',
  `currency_money` bigint(20) NULL DEFAULT NULL COMMENT '货币_零钱',
  `currency_gold` bigint(20) NULL DEFAULT NULL COMMENT '货币_金币',
  `sign_num` int(11) NOT NULL DEFAULT 0 COMMENT '连续签到天数',
  `spread_uid` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '推广元id',
  `spread_time` datetime NULL DEFAULT NULL COMMENT '推广员关联时间',
  `is_promoter` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否为推广员',
  `spread_count` int(11) NULL DEFAULT 0 COMMENT '下级人数',
  `add_ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '添加ip',
  `last_ip` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '最后一次登录ip',
  `channel_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '频道_编号',
  `update_forum_posts_time` datetime(3) NULL DEFAULT NULL COMMENT '发布帖子时间',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',  
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_extend_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '终端用户拓展' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_user_extend
-- ----------------------------
INSERT INTO `hk_user_extend` VALUES (2, '000000', '', '', '', NULL, '', '', '', 0, 1, 1, 0, 999998, 0, 0, NULL, 0, 0, '', '', '', '2023-04-26 16:55:00.384', '2023-04-25 18:42:39.708', '2023-04-26 16:55:00.384', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_extend` VALUES (3, '000000', '', '', '', NULL, '', '',  '',0, 1, 1, 0, 1000000, 0, 0, NULL, 0, 0, '', '', '', '2023-04-25 18:42:39.719', '2023-04-25 18:42:39.719', '2023-04-25 18:44:05.785', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_user_cover_image
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_cover_image`;
CREATE TABLE `hk_user_cover_image`  (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_extend_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户主页封面' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_user_cover_image
-- ----------------------------
INSERT INTO `hk_user_cover_image` VALUES (1, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230509/d83ddab0eaeab0ecccda2f4f8cee1421.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_cover_image` VALUES (2, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/52f7bf7fb12ea2fa12a7ea7015eb3fe3.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_cover_image` VALUES (3, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/583813a9b9c54bd7b8de1f5e308b869f.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_cover_image` VALUES (4, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/37518339a95d26159d8c8d17c4540030.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_cover_image` VALUES (5, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/2e1d769d8f483c08d57ea5dad21ab847.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_user_header_image
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_header_image`;
CREATE TABLE `hk_user_header_image`  (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `header_img` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_extend_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户主页封面' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_user_header_image
-- ----------------------------
INSERT INTO `hk_user_header_image` VALUES (1, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230509/d83ddab0eaeab0ecccda2f4f8cee1421.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_header_image` VALUES (2, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/52f7bf7fb12ea2fa12a7ea7015eb3fe3.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_header_image` VALUES (3, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/583813a9b9c54bd7b8de1f5e308b869f.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_header_image` VALUES (4, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/37518339a95d26159d8c8d17c4540030.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_user_header_image` VALUES (5, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/2e1d769d8f483c08d57ea5dad21ab847.jpg', 0, '2023-05-10 15:47:57.000', '2023-05-10 15:47:57.000', NULL, NULL, NULL, NULL, 0, 0);

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
  `type` smallint(6) NULL DEFAULT NULL COMMENT '类型：0官方圈子、1用户圈子、2小区',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
  `logo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子Logo',
  `circle_classify_id` bigint(20) NULL DEFAULT NULL COMMENT '圈子分类_编号',
  `slogan` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子标语',
  `des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子简介',
  `protocol` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子规约',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子背景图',
  `support_category` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支持内容类别(json数组)：1视频、2动态、3资讯、4公告、5文章、6问答、7活动',
  `pay` smallint(6) NULL DEFAULT NULL COMMENT '付费：0否、1是',
  `process` smallint(6) NULL DEFAULT NULL COMMENT '是否开启版块内容人工审核：0 否，1是',
  `property` smallint(6) NULL DEFAULT NULL COMMENT '圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）',
  `view` smallint(6) NULL DEFAULT NULL COMMENT '板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到',
  `power_add` smallint(6) NULL DEFAULT NULL COMMENT '圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户',
  `power_view` smallint(6) NULL DEFAULT NULL COMMENT '圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组',
  `power_publish` smallint(6) NULL DEFAULT NULL COMMENT '圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组',
  `power_comment` smallint(6) NULL DEFAULT NULL COMMENT '圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组',
  `power_add_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子加入权限_指定部门和成员(json数组)',
  `power_view_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子内浏览权限_指定部门和用户(json数组)',
  `power_publish_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子内发布权限_指定部门和用户(json数组)',
  `power_comment_user` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子内评论权限_指定部门和用户(json数组)',
  `no_limit_user_group` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '不受限用户组(json数组)',
  `new_user_focus` smallint(6) NULL DEFAULT NULL COMMENT '新注册用户默认关注：0 否，1是',
  `channel_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '频道_编号',
  `user_num` smallint(6) NULL DEFAULT NULL COMMENT '用户数',
  `update_forum_posts_time` datetime(3) NULL DEFAULT NULL COMMENT '发布帖子时间',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_circle
-- ----------------------------
INSERT INTO `hk_circle` VALUES (1, '000000', 0, '伐木累', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/b02c8aad08dd7521970237e5cae29006.jpg', 1, 'we are family!', '圈主很懒，什么都没有留下、、、', '请遵守24字社会主义核心价值观', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230509/d83ddab0eaeab0ecccda2f4f8cee1421.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 1, '2023-02-23 10:16:26.000', '2023-05-10 11:00:42.606', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (2, '000000', 0, '音乐发烧友', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/6f0ae71d842f9a3cafd55d45c6a86f2a.png', 1, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/52f7bf7fb12ea2fa12a7ea7015eb3fe3.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 2, '2023-02-23 10:16:26.000', '2023-05-10 11:00:12.639', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (3, '000000', 0, '开黑', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/f05472796ad05bfce93c7c3d60dfaae7.jpg', 7, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/583813a9b9c54bd7b8de1f5e308b869f.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:58:43.683', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (4, '000000', 0, '有道是', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/135ac35db49bfbd90778eb8e8825ce89.jpg', 1, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/37518339a95d26159d8c8d17c4540030.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 1, 3, '2023-02-23 10:16:26.000', '2023-05-10 11:28:20.143', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (5, '000000', 0, '常言道', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/718d153ee540d036cdc12aab1ed8f527.jpg', 1, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/2e1d769d8f483c08d57ea5dad21ab847.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:59:32.186', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (6, '000000', 0, '日常两三事', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/7b715288b69697a6b14b1a1d76c4bf0f.jpg', 2, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/812fbdf567901e4825ff6c60571e54f9.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:57:33.655', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (7, '000000', 0, '靓仔集结地', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/a084db112e39236d0f6ecadc024664b4.jpg', 3, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/2af6a45790953d0aa60e489bb6d2cf0d.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:49:56.234', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (8, '000000', 0, '冲壳子', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/e4e92d12df8c454fffc1e61c90d21cb1.jpg', 5, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/1cec746aff131db639cb1cb4d8f66042.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:56:53.943', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (9, '000000', 0, '慢跑新生活', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/8dab91bf653219b64a607509ff090d0b.jpg', 6, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/135b0a605988a41c79620820997a7813.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:48:43.518', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (10, '000000', 0, '生命在于运动', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/c46e0a49cf261c6665b76074ad83e244.jpg', 8, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/f4f2a800825062b18fba0fb7114e0b5d.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 11:02:53.527', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (11, '000000', 0, '拒绝emo', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/fa586fc6948931b90c3f3b619bfdb803.jpg', 9, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/370616099cde20906753edd16b809533.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:47:18.889', NULL, -1, -1, 1123598821738675201, 0, 0);
INSERT INTO `hk_circle` VALUES (12, '000000', 0, '舌尖上的美味', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/2d623d0de89ad33114c890ffa9620eb2.jpg', 4, '圈子标语内容', '圈子简介内容', NULL, 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230510/0f7ac726eb31809c746068e720160e70.jpg', NULL, 0, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL, 0, NULL, 0, 3, '2023-02-23 10:16:26.000', '2023-05-10 10:46:37.760', NULL, -1, -1, 1123598821738675201, 0, 0);

-- ----------------------------
-- Table structure for hk_circle_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_request`;
CREATE TABLE `hk_circle_request`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户ID',
  `user_nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `type` int(2) NOT NULL DEFAULT 1 COMMENT '类型：0官方圈子 ，1用户圈子',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
  `logo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子Logo',
  `circle_classify_id` bigint(20) NOT NULL COMMENT '圈子分类_编号',
  `slogan` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子标语',
  `des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子简介',
  `protocol` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子规约',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子背景图',
  `property` smallint(6) NULL DEFAULT NULL COMMENT '圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）',
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
-- Table structure for hk_circle_add_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_add_request`;
CREATE TABLE `hk_circle_add_request`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户_编号',
  `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
  `reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '申请理由',
  `check_user` bigint(20) NULL DEFAULT NULL COMMENT '审核人',
  `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
  `check_status` smallint(6) NOT NULL DEFAULT 0 COMMENT '审核状态：0 未处理 1 通过，2驳回',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_add_request_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '加入圈子申请' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_circle_add_request
-- ----------------------------
INSERT INTO `hk_circle_add_request` VALUES (1, '000000', 1, 1, '原因1', NULL, NULL, 0, '2023-05-16 13:59:10.200', '2023-05-16 13:59:10.200', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_add_request` VALUES (2, '000000', 6, 5, '原因2', NULL, NULL, 0, '2023-05-23 11:47:10.709', '2023-05-23 11:47:10.709', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_circle_relation
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_relation`;
CREATE TABLE `hk_circle_relation`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `relation_type` int(2) NULL DEFAULT NULL COMMENT '关系类型：0父子 1友好',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `circle_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
   `other_circle_id` bigint(20) NOT NULL COMMENT '关系圈子_编号(孩子圈子Id，或者友好圈子Id)',
   `other_circle_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关系圈子名称',
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
-- Table structure for hk_circle_relation_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_relation_request`;
CREATE TABLE `hk_circle_relation_request`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `relation_type` int(2) NULL DEFAULT NULL COMMENT '关系类型：0父子 1友好',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `circle_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
   `request_circle_id` bigint(20) NOT NULL COMMENT '被请求圈子编号',
   `request_circle_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '被请求圈子名称',
   `request_des` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求描述',
   `check_status` int(2) NULL DEFAULT 0 COMMENT '审批状态：0未审批 1同意 2拒绝',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子关系申请';

-- ----------------------------
-- Table structure for hk_circle_tag
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_tag`;
CREATE TABLE `hk_circle_tag`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
  `name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
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
) ENGINE = InnoDB AUTO_INCREMENT = 49 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子标签' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_circle_tag
-- ----------------------------
INSERT INTO `hk_circle_tag` VALUES (1, '000000', 1, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (2, '000000', 1, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (3, '000000', 1, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (4, '000000', 1, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (5, '000000', 2, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (6, '000000', 2, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (7, '000000', 2, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (8, '000000', 2, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (9, '000000', 3, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (10, '000000', 3, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (11, '000000', 3, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (12, '000000', 3, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (13, '000000', 4, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (14, '000000', 4, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (15, '000000', 4, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (16, '000000', 4, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (17, '000000', 5, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (18, '000000', 5, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (19, '000000', 5, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (20, '000000', 5, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (21, '000000', 6, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (22, '000000', 6, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (23, '000000', 6, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (24, '000000', 6, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (25, '000000', 7, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (26, '000000', 7, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (27, '000000', 7, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (28, '000000', 7, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (29, '000000', 8, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (30, '000000', 8, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (31, '000000', 8, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (32, '000000', 8, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (33, '000000', 9, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (34, '000000', 9, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (35, '000000', 9, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (36, '000000', 9, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (37, '000000', 10, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (38, '000000', 10, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (39, '000000', 10, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (40, '000000', 10, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (41, '000000', 11, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (42, '000000', 11, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (43, '000000', 11, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (44, '000000', 11, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (45, '000000', 12, '交友', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (46, '000000', 12, '娱乐', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (47, '000000', 12, '闲聊', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_tag` VALUES (48, '000000', 12, '游戏', NULL, '2023-04-28 15:21:52.000', '2023-04-28 15:21:52.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_circle_user
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_user`;
CREATE TABLE `hk_circle_user`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户ID',
   `remark` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
   `tag` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
   `circle_id` bigint(20) NOT NULL COMMENT '圈子_编号',
   `circle_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
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
INSERT INTO `hk_circle_user` VALUES (1, '000000', 1, '', '', 1, "圈子名称1", 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (2, '000000', 2, '', '', 1, "圈子名称1", 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (3, '000000', 3, '', '', 1, "圈子名称1", 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (4, '000000', 2, '', '', 2, "圈子名称2", 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (5, '000000', 3, '', '', 2, "圈子名称2", 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_user` VALUES (6, '000000', 3, '', '', 3, "圈子名称3", 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);

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
   `type` tinyint(1) NULL DEFAULT NULL COMMENT '类型：0系统 1默认频道',
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
   UNIQUE INDEX `idx_hk_channel_name`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '频道' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_channel
-- ----------------------------
INSERT INTO `hk_channel` VALUES (1, '000000', '资讯', NULL, 1, 1, 1, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:40.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (2, '000000', '关注', NULL, 2, 1, 1, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (3, '000000', '附近', NULL, 3, 1, 1, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (4, '000000', '活动', NULL, 4, 1, 1, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (5, '000000', '问答', NULL, 5, 1, 1, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (6, '000000', '频道1', NULL, 0, 0, 1, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (7, '000000', '频道2', NULL, 0, 0, 0, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (8, '000000', '频道3', NULL, 0, 0, 0, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `hk_channel` VALUES (9, '000000', '频道4', NULL, 0, 0, 0, NULL, '2023-04-20 14:16:38.000', '2023-04-20 14:16:38.000', NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for hk_circle_channel
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_channel`;
CREATE TABLE `hk_circle_channel`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `circle_id` bigint(20) NULL DEFAULT NULL COMMENT '圈子_编号',
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
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子频道' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_circle_channel
-- ----------------------------
INSERT INTO `hk_circle_channel` VALUES (1, '000000', 0, '全部', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_channel` VALUES (2, '000000', 0, '资讯', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_channel` VALUES (3, '000000', 0, '动态', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_channel` VALUES (4, '000000', 0, '问答', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_channel` VALUES (5, '000000', 0, '活动', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_circle_channel` VALUES (6, '000000', 0, '文章', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_forum_posts
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_posts`;
CREATE TABLE `hk_forum_posts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `circle_id` bigint(20) NULL DEFAULT NULL COMMENT '圈子_编号',
  `activity_id` bigint(20) NULL DEFAULT NULL COMMENT '活动_编号',
  `category` smallint(6) NULL DEFAULT NULL COMMENT '类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动',
  `channel_id` bigint(20) NULL DEFAULT NULL COMMENT '频道_编号',
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `seo_key` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO关键词',
  `seo_introduce` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介(SEO简介)',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `source` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源',
  `content_type` smallint(6) NULL DEFAULT NULL COMMENT '内容类型：1markdown、2html',
  `content_markdown` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'markdown内容',
  `content_html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'html内容',
  `video` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '视频地址',
  `attachment` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附件',
  `tag` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `top` smallint(6) NULL DEFAULT NULL COMMENT '置顶：0否、1是',
  `marrow` smallint(6) NULL DEFAULT NULL COMMENT '精华：0否、1是',
  `comment_id` bigint(20) NULL DEFAULT NULL COMMENT '问答最佳答案ID 0未解答，否则已解答',
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
  `pay_currency` smallint(6) NULL DEFAULT NULL COMMENT '付费货币：1人民、2积分、3金币',
  `pay_num` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '付费金额',
  `activity_start_at` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '活动开始时间',
  `activity_end_at` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '活动结束时间',
  `activity_address` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '活动地址',
  `activity_user_num` smallint(6) NULL DEFAULT NULL COMMENT '活动用户人数（0不限制活动人数，否则为活动人数）',
  `activity_cur_user_num` smallint(6) NULL DEFAULT NULL COMMENT '活动参加人数',
  `activity_chat_group_id` bigint(20) NOT NULL COMMENT '活动聊天群编号',
  `activity_add_approve` smallint(6) NULL DEFAULT NULL COMMENT '参加活动是否需要审批：0不审批、1审批',
  `is_public` smallint(6) NULL DEFAULT NULL COMMENT '是否公开：0 否 1 是',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_posts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帖子' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_forum_posts
-- ----------------------------
INSERT INTO `hk_forum_posts` VALUES (1, '000000', 1, 0, 2, 0, '圈子1_标题（用户1）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 2, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-15 20:29:57.000', '2023-04-27 17:08:06.582', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (2, '000000', 1, 0, 2, 0, '圈子1_标题（用户2）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 2, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-15 20:29:57.000', '2023-04-27 17:08:06.582', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (3, '000000', 1, 0, 2, 0, '圈子1_标题（用户3）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 2, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-15 20:29:57.000', '2023-04-28 17:59:37.846', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (4, '000000', 2, 0, 2, 0, '圈子2_标题（用户2）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 2, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-15 20:29:57.000', '2023-04-27 17:08:06.582', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (5, '000000', 2, 0, 2, 0, '圈子2_标题（用户3）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 2, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-15 20:29:57.000', '2023-04-27 17:08:06.582', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (6, '000000', 3, 0, 2, 0, '圈子3_标题（用户3）', NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 2, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-15 20:29:57.000', '2023-04-27 17:08:06.582', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (7, '000000', 1, 0, 0, 0, 'testtest', '', '', '', '', 0, '', '', '', '', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-18 11:41:17.115', '2023-04-27 17:08:06.582', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (8, '000000', 1, 0, 1, 0, '有附件', '', '', '', '', 1, '', '有附件', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker4029652423983233700.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-20/image_picker4029652423983233700.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-20/image_picker4029652423983233700.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-20 09:52:01.938', '2023-05-04 17:45:48.062', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (9, '000000', 1, 0, 1, 0, '两个附件', '', '', '', '', 1, '', '两个附件', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 1, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-20 10:06:26.981', '2023-05-04 14:42:46.248', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (10, '000000', 1, 0, 1, 0, '两个附件2', '', '', '', '', 1, '', '两个附件2', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker9222995987695013449.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-20/image_picker9222995987695013449.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-20/image_picker9222995987695013449.jpg\"},{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker6600262919828096699.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-20/image_picker6600262919828096699.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-20/image_picker6600262919828096699.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-20 10:07:13.517', '2023-05-04 18:13:58.536', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (11, '000000', 1, 0, 1, 0, '今天没赚钱顶级精明和李敏李丽丽你家里', '', '', '', '', 1, '', '今天没赚钱顶级精明和李敏李丽丽你家里', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker4577751758068097584.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-21/image_picker4577751758068097584.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-21/image_picker4577751758068097584.jpg\"},{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker7462903158421350833.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-21/image_picker7462903158421350833.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-02-21/image_picker7462903158421350833.jpg\"},{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2631187621018139806.png\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-02-21/image_picker2631187621018139806.png\",\"tag\":\"png\",\"key\":\"test/uploads/2023-02-21/image_picker2631187621018139806.png\"}]', '', 0, 0, 0, 0, 2, 0, 3, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, '', 0, 0, 0, 0, 1, '2023-02-21 17:42:57.034', '2023-05-05 16:36:40.968', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (12, '000000', 1, 0, 7, 0, '走近科学', '', '', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-26/image_picker1125034918347843884.jpg', '', 2, '', '醒一醒他了吗？发票抬头：***有限公司', '', '', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 1, 30, '2023-04-26 13:45:00', '2023-04-26 15:45:00', '四川省成都市高新区桂溪东路三瓦窑街347号', 0, 0, 0, 0, 1, '2023-04-26 13:46:29.823', '2023-05-06 15:27:16.750', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (13, '000000', 3, 0, 7, 0, '微信小游戏', '', '', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-26/image_picker6175649061815111165.jpg', '', 2, '', '麻烦放收发室，谢谢。', '', '', '', 0, 0, 0, 0, 2, 0, 1, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 1, 100, '2023-04-26 13:51:00', '2023-04-28 13:51:00', '四川省成都市武侯区三瓦窑街202号-附10', 99, 1, 0, 0, 1, '2023-04-26 13:52:07.558', '2023-05-06 15:27:16.741', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (14, '000000', 0, 0, 0, 0, '哦需要啥想日在皮笑肉在啥', '', '', '', '', 0, '', '', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker7097097007589862651.png\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-26/image_picker7097097007589862651.png\",\"tag\":\"png\",\"key\":\"test/uploads/2023-04-26/image_picker7097097007589862651.png\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-04-26 16:49:53.769', '2023-05-06 15:27:16.726', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (15, '000000', 0, 0, 6, 0, '描述问题', '', '', '', '', 2, '', '描述问题', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker8659673069068019080.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-26/image_picker8659673069068019080.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-26/image_picker8659673069068019080.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 1, 2, '', '', '', 0, 0, 0, 0, 1, '2023-04-26 16:55:00.375', '2023-05-06 13:11:37.453', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (16, '000000', 0, 0, 7, 0, '显卡天梯', '', '', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-26/image_picker5725219757837810081.jpg', '', 2, '', '嘻嘻嘻嘻修改环境磨破皮', '', '', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 1, 0, '2023-04-26 18:23:00', '2023-04-27 18:23:00', '金桂路83', 0, 1, 0, 0, 1, '2023-04-26 18:23:17.695', '2023-05-04 17:22:58.242', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (17, '000000', 0, 16, 2, 0, '这个显卡不错', '', '', '', '', 0, '', '', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker5625612915889866833.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker5625612915889866833.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker5625612915889866833.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-04-27 11:06:50.186', '2023-04-27 11:06:50.186', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (18, '000000', 0, 16, 2, 0, '谁说我除了吃啥都不知道，我还知道饿的人了啊你(／_＼)大怨种', '', '', '', '', 0, '', '', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2567766332373301928.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker2567766332373301928.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker2567766332373301928.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-04-27 14:42:07.527', '2023-04-27 14:42:07.527', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (19, '000000', 2, 0, 3, 0, '资讯消息', NULL, NULL, NULL, NULL, 2, NULL, '啊啊啊啊啊啊', NULL, '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2567766332373301928.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker2567766332373301928.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker2567766332373301928.jpg\"}]', NULL, 1, 0, 0, 0, 2, 0, 3, 0, 1, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, 0, 1, 0, 0, 1, '2023-04-27 14:53:59.000', '2023-05-06 16:54:50.493', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (20, '000000', 3, 13, 2, 0, '搜索人', '', '', '', '', 0, '', '', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 0, '2023-04-28 17:58:09.803', '2023-04-28 17:58:09.803', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (21, '000000', 1, 12, 2, 0, '还好哈哈哈', '', '', '', '', 0, '', '', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 0, '2023-04-28 18:00:06.714', '2023-04-28 18:00:06.714', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (23, '000000', 0, 0, 1, 0, '么得', '', '', '', '', 1, '', '么得', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-05-05 09:54:09.120', '2023-05-05 09:54:09.120', '2023-05-05 09:54:31.762', NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (30, '000000', 0, 0, 2, 0, '12345', '', '', '', '', 1, '', '12345', '', '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker7288195821003054315.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-05-05/image_picker7288195821003054315.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-05-05/image_picker7288195821003054315.jpg\"}]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-05-05 14:29:08.005', '2023-05-05 15:30:23.915', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (31, '000000', 0, 0, 5, 0, '套路了', '', '', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-05-05/image_picker7410752799846556025.jpg', '', 1, '', '这里有性命危险的时候回来的时候回来了，', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-05-05 14:29:16.323', '2023-05-05 15:39:40.510', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (32, '000000', 0, 0, 2, 0, '红红火火恍恍惚惚', '', '', '', '', 1, '', '红红火火恍恍惚惚', '', '[]', '', 0, 0, 0, 0, 2, 0, 1, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-05-05 19:01:19.873', '2023-05-05 19:01:46.957', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (33, '000000', 0, 0, 2, 0, '123', '', '', '', '', 1, '', '123', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 1, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-05-06 15:29:43.031', '2023-05-06 15:29:43.031', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (34, '000000', 0, 0, 2, 0, '', '', '', '', '', 1, '', '', '', '[]', '', 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, '', '', '', 0, 0, 0, 0, 1, '2023-05-06 16:53:52.186', '2023-05-06 16:53:52.186', NULL, NULL, NULL, NULL, 0, 0, 0, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (35, '000000', 2, 0, 3, 0, '资讯消息1', NULL, NULL, NULL, NULL, 2, NULL, '资讯消息内容', NULL, '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2567766332373301928.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker2567766332373301928.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker2567766332373301928.jpg\"}]', NULL, 1, 0, 0, 0, 2, 0, 3, 0, 1, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, 0, 1, 0, 0, 1, '2023-04-27 14:53:59.000', '2023-05-06 16:54:50.493', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (36, '000000', 2, 0, 3, 0, '资讯消息2', NULL, NULL, NULL, NULL, 2, NULL, '资讯消息内容', NULL, '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2567766332373301928.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker2567766332373301928.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker2567766332373301928.jpg\"}]', NULL, 1, 0, 0, 0, 2, 0, 3, 0, 1, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, 0, 1, 0, 0, 1, '2023-04-27 14:53:59.000', '2023-05-06 16:54:50.493', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (37, '000000', 2, 0, 3, 0, '资讯消息3', NULL, NULL, NULL, NULL, 2, NULL, '资讯消息内容', NULL, '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2567766332373301928.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker2567766332373301928.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker2567766332373301928.jpg\"}]', NULL, 1, 0, 0, 0, 2, 0, 3, 0, 1, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, 0, 1, 0, 0, 1, '2023-04-27 14:53:59.000', '2023-05-06 16:54:50.493', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_posts` VALUES (38, '000000', 2, 0, 3, 0, '资讯消息4', NULL, NULL, NULL, NULL, 2, NULL, '资讯消息内容', NULL, '[{\"id\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"status\":0,\"createUser\":null,\"createDept\":null,\"updateUser\":null,\"name\":\"image_picker2567766332373301928.jpg\",\"url\":\"https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/test/uploads/2023-04-27/image_picker2567766332373301928.jpg\",\"tag\":\"jpg\",\"key\":\"test/uploads/2023-04-27/image_picker2567766332373301928.jpg\"}]', NULL, 1, 0, 0, 0, 2, 0, 3, 0, 1, 1, 0, NULL, 3, 1, 0, 0, 0, 0, 0, 0, 0, NULL, NULL, NULL, 0, 1, 0, 0, 1, '2023-04-27 14:53:59.000', '2023-05-06 16:54:50.493', NULL, NULL, NULL, NULL, 0, 0);

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
INSERT INTO `hk_apply` VALUES (1, '000000', 0, 0, 0, '应用1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230227/42ebfc53bfcbcf1be1b05b5203c63d0f.png', 1, 1, NULL, 'https://www.baidu.com/', NULL, '2023-02-20 10:58:15.000', '2023-02-27 16:54:42.703', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (2, '000000', 0, 0, 0, '应用2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (3, '000000', 0, 0, 0, '应用3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (4, '000000', 0, 0, 0, '应用4', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (5, '000000', 0, 0, 0, '应用5', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (6, '000000', 0, 0, 0, '应用6', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (11, '000000', 0, 0, 0, '应用1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (12, '000000', 0, 0, 0, '应用2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (13, '000000', 0, 0, 0, '应用3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (14, '000000', 0, 0, 0, '应用4', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (15, '000000', 0, 0, 0, '应用5', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (16, '000000', 0, 0, 0, '应用6', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (21, '000000', 0, 0, 0, '应用1', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (22, '000000', 0, 0, 0, '应用2', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (23, '000000', 0, 0, 0, '应用3', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (24, '000000', 0, 0, 0, '应用4', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (25, '000000', 0, 0, 0, '应用5', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_apply` VALUES (26, '000000', 0, 0, 0, '应用6', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', 1, 0, 1627919042591342593, NULL, NULL, '2023-02-20 10:58:15.421', '2023-02-20 10:58:15.421', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

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
   `browse_num` bigint(20) NULL DEFAULT 0 COMMENT '浏览次数',
   `power` smallint(6) NULL DEFAULT 0 COMMENT '权限：0 所有可见 1 圈内成员可见',
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
INSERT INTO `hk_circle_apply` VALUES (1, '000000', 1, 11, 1, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (2, '000000', 1, 11, 2, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (3, '000000', 1, 12, 3, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (4, '000000', 1, 12, 4, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (5, '000000', 1, 12, 5, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (6, '000000', 1, 13, 6, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

INSERT INTO `hk_circle_apply` VALUES (11, '000000', 1, 21, 11, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (12, '000000', 1, 21, 12, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (13, '000000', 1, 22, 13, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (14, '000000', 1, 22, 14, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (15, '000000', 1, 22, 15, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (16, '000000', 1, 23, 16, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

INSERT INTO `hk_circle_apply` VALUES (21, '000000', 1, 31, 21, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (22, '000000', 1, 31, 22, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (23, '000000', 1, 32, 23, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (24, '000000', 1, 32, 24, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (25, '000000', 1, 32, 25, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);
INSERT INTO `hk_circle_apply` VALUES (26, '000000', 1, 33, 26, 0, 0, 0, '2023-02-21 15:27:07.114', '2023-02-21 15:27:07.114', NULL, 1123598821738675201, 1123598813738675201, 1123598821738675201, 1, 0);

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
INSERT INTO `hk_forum_topic_group` VALUES (1, '000000', NULL, '美食', 1, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (2, '000000', NULL, '时尚美妆', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (3, '000000', 1, '动漫', 1, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (4, '000000', 1, '搞笑', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (5, '000000', 1, '旅游', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (6, '000000', 1, '明星', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (7, '000000', 1, '游戏', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (8, '000000', 1, '摄影', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (9, '000000', 1, '萌宠', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (10, '000000', 1, '运动健身', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (11, '000000', 1, '电视剧', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (12, '000000', 1, '综艺', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (13, '000000', 1, '电影', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (14, '000000', 1, '情感', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (15, '000000', 1, '艺术', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (16, '000000', 1, '读书', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (17, '000000', 1, '互联网', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (18, '000000', 1, '数码', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (19, '000000', 1, '育儿', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (20, '000000', 1, '音乐', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic_group` VALUES (21, '000000', 1, '教育', 2, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);

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
   `hot` int(2) NULL DEFAULT 0 COMMENT '是否热门：0 否 1 是',
   `discuss_num` int(11) NULL DEFAULT NULL COMMENT '讨论数',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
     PRIMARY KEY (`id`) USING BTREE,
     UNIQUE INDEX `idx_hk_topic_name`(`tenant_id`,`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '话题';

-- ----------------------------
-- Records of hk_forum_topic
-- ----------------------------
INSERT INTO `hk_forum_topic` VALUES (1, '000000', 'topic_1_1', NULL, 1, NULL, NULL, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_forum_topic` VALUES (2, '000000', 'topic_1_2', NULL, 1, NULL, NULL, '2023-02-16 12:36:21.000', '2023-02-16 12:36:21.000', NULL, NULL, NULL, NULL, 0, 0);

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
-- Table structure for hk_help_type
-- ----------------------------
DROP TABLE IF EXISTS `hk_help_type`;
CREATE TABLE `hk_help_type`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_help_type_created_at`(`created_at`) USING BTREE,
  INDEX `idx_hk_help_type_updated_at`(`updated_at`) USING BTREE,
  INDEX `idx_hk_help_type_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帮助类型' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_help_type
-- ----------------------------
INSERT INTO `hk_help_type` VALUES (1, '000000', '类型1', NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help_type` VALUES (2, '000000', '类型2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help_type` VALUES (3, '000000', '类型3', NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help_type` VALUES (4, '000000', '类型4', NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help_type` VALUES (5, '000000', '类型5', NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_help
-- ----------------------------
DROP TABLE IF EXISTS `hk_help`;
CREATE TABLE `hk_help`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
  `type_id` bigint(20) NOT NULL COMMENT '类型编号',
  `good_num` smallint(6) NULL DEFAULT NULL COMMENT '好评数',
  `bad_num` smallint(6) NULL DEFAULT NULL COMMENT '差评数',
  `top` smallint(6) NULL DEFAULT NULL COMMENT '是否置顶：0不 1是',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_help_updated_at`(`updated_at`) USING BTREE,
  INDEX `idx_hk_help_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_hk_help_created_at`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帮助' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_help
-- ----------------------------
INSERT INTO `hk_help` VALUES (1, '000000', '帮助标题', '帮助内容', 1, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (2, '000000', '帮助标题', '帮助内容', 1, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (3, '000000', '帮助标题', '帮助内容', 1, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (4, '000000', '帮助标题', '帮助内容', 1, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (5, '000000', '帮助标题', '帮助内容', 1, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (6, '000000', '帮助标题', '帮助内容', 2, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (7, '000000', '帮助标题', '帮助内容', 2, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (8, '000000', '帮助标题', '帮助内容', 2, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (9, '000000', '帮助标题', '帮助内容', 2, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (10, '000000', '帮助标题', '帮助内容', 2, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (11, '000000', '帮助标题', '帮助内容', 3, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (12, '000000', '帮助标题', '帮助内容', 3, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (13, '000000', '帮助标题', '帮助内容', 3, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (14, '000000', '帮助标题', '帮助内容', 3, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (15, '000000', '帮助标题', '帮助内容', 3, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (16, '000000', '帮助标题', '帮助内容', 4, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (17, '000000', '帮助标题', '帮助内容', 4, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (18, '000000', '帮助标题', '帮助内容', 4, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (19, '000000', '帮助标题', '帮助内容', 4, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (20, '000000', '帮助标题', '帮助内容', 4, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (21, '000000', '帮助标题', '帮助内容', 5, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (22, '000000', '帮助标题', '帮助内容', 5, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (23, '000000', '帮助标题', '帮助内容', 5, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (24, '000000', '帮助标题', '帮助内容', 5, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_help` VALUES (25, '000000', '帮助标题', '帮助内容', 5, NULL, NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_help_thumbs_up
-- ----------------------------
DROP TABLE IF EXISTS `hk_help_thumbs_up`;
CREATE TABLE `hk_help_thumbs_up`  (
   `user_id` bigint(20) NOT NULL COMMENT'用户id',
   `help_id` bigint(20) NOT NULL COMMENT '帮助编号',
   `thumbs_up` int(11) NULL DEFAULT 0 COMMENT '是否好评：0未点赞 1有用 2无用',
   PRIMARY KEY (`user_id`, `help_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '帮助点赞' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for hk_feedback
-- ----------------------------
DROP TABLE IF EXISTS `hk_feedback`;
CREATE TABLE `hk_feedback`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `des` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
   `attachment` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附件',
   `check_status` int(2) NULL DEFAULT NULL COMMENT '处理状态：0 未处理、1 处理中、2 拒绝、3 完成',
   `process` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '处理描述',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '反馈';

-- ----------------------------
-- Table structure for hk_report_reason
-- ----------------------------
DROP TABLE IF EXISTS `hk_report_reason`;
CREATE TABLE `hk_report_reason`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `report_type` int(2) NOT NULL COMMENT '举报类型:1用户、2圈子、3群、4帖子、5帖子评论',
  `reason` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报理由',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_report_reason_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '举报原因' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_report_reason
-- ----------------------------
INSERT INTO `hk_report_reason` VALUES (1, '000000', 1, '用户举报原因1', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (2, '000000', 1, '用户举报原因2', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (3, '000000', 1, '用户举报原因3', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (4, '000000', 1, '用户举报原因4', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (5, '000000', 1, '用户举报原因5', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (6, '000000', 1, '用户举报原因6', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (7, '000000', 1, '用户举报原因7', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (8, '000000', 1, '用户举报原因用户举报原因用户举报原因用户举报原因用户举报原因用户举报原因8', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (9, '000000', 1, '用户举报原因9', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (10, '000000', 1, '用户举报原因10', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (11, '000000', 2, '圈子举报原因1', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (12, '000000', 2, '圈子举报原因2', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (13, '000000', 2, '圈子举报原因3', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (14, '000000', 2, '圈子举报原因4', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (15, '000000', 2, '圈子举报原因5', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (16, '000000', 2, '圈子举报原因6', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (17, '000000', 2, '圈子举报原因7', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (18, '000000', 2, '圈子举报原因圈子举报原因圈子举报原因圈子举报原因圈子举报原因圈子举报原因8', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (19, '000000', 2, '圈子举报原因9', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (20, '000000', 2, '圈子举报原因10', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (21, '000000', 3, '群举报原因1', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (22, '000000', 3, '群举报原因2', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (23, '000000', 3, '群举报原因3', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (24, '000000', 3, '群举报原因4', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (25, '000000', 3, '群举报原因5', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (26, '000000', 3, '群举报原因6', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (27, '000000', 3, '群举报原因7', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (28, '000000', 3, '群举报原因群举报原因群举报原因群举报原因群举报原因群举报原因8', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (29, '000000', 3, '群举报原因9', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (30, '000000', 3, '群举报原因10', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (31, '000000', 4, '帖子举报原因1', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (32, '000000', 4, '帖子举报原因2', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (33, '000000', 4, '帖子举报原因3', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (34, '000000', 4, '帖子举报原因4', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (35, '000000', 4, '帖子举报原因5', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (36, '000000', 4, '帖子举报原因6', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (37, '000000', 4, '帖子举报原因7', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (38, '000000', 4, '帖子举报原因帖子举报原因帖子举报原因帖子举报原因帖子举报原因帖子举报原因8', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (39, '000000', 4, '帖子举报原因9', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (40, '000000', 4, '帖子举报原因10', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (41, '000000', 5, '帖子评论举报原因1', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (42, '000000', 5, '帖子评论举报原因2', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (43, '000000', 5, '帖子评论举报原因3', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (44, '000000', 5, '帖子评论举报原因4', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (45, '000000', 5, '帖子评论举报原因5', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (46, '000000', 5, '帖子评论举报原因6', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (47, '000000', 5, '帖子评论举报原因7', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (48, '000000', 5, '帖子评论举报原因8', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (49, '000000', 5, '帖子评论举报原因9', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_report_reason` VALUES (50, '000000', 5, '帖子评论举报原因10', NULL, '2023-05-19 17:48:13.000', '2023-05-19 17:48:13.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_report
-- ----------------------------
DROP TABLE IF EXISTS `hk_report`;
CREATE TABLE `hk_report`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '举报用户编号',
   `reason` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报原因',
   `report_type` int(2) NOT NULL COMMENT '举报类型:1用户、2圈子、3群、4帖子、5帖子评论',
   `report_id` bigint(20) NOT NULL COMMENT '被举报编号',
   `content` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报内容',
   `content_attachment` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '内容附件',
   `cur_status` int(2) NULL DEFAULT NULL COMMENT '处理状态：0 未处理、1 处理中、2 拒绝、3 完成',
   `handle_user_id` bigint(20) NULL DEFAULT NULL COMMENT '操作用户_编号',
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
-- Table structure for hk_user_browsing_history
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_browsing_history`;
CREATE TABLE `hk_user_browsing_history`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `posts_id` bigint(20) NOT NULL COMMENT '帖子编号',
   `category` smallint(6) NULL DEFAULT NULL COMMENT '类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动',
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
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `remark` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `tag` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `focus_user_id` bigint(20) NOT NULL COMMENT '关注用户ID',
  `focus_nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关注用户昵称',
  `mutual` int(2) NULL DEFAULT 0 COMMENT '是否相互关注:0否 1是',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `focus_user` (`user_id`, `focus_user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户关注' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hk_focus_user
-- ----------------------------
INSERT INTO `hk_focus_user` VALUES (2, '000000', 2, '范杰2', '', '', 3, '用户3', 0, '2023-04-21 00:24:22.097', '2023-04-21 00:51:09.831', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_focus_user` VALUES (3, '000000', 3, '用户3', NULL, NULL, 2, '范杰2', 0, '2023-04-21 00:24:22.097', '2023-04-21 00:24:22.097', NULL, NULL, NULL, NULL, 0, 0);


-- ----------------------------
-- Table structure for hk_activity_user 
-- ----------------------------
DROP TABLE IF EXISTS `hk_activity_user`;
CREATE TABLE `hk_activity_user`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `activity_id` bigint(20) NOT NULL COMMENT '活动编号',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `remark` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
   `tag` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
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
   UNIQUE INDEX `idx_hk_activity_user`(`user_id`, `activity_id`) USING BTREE
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
-- Table structure for hk_pay_type
-- ----------------------------
DROP TABLE IF EXISTS `hk_pay_type`;
CREATE TABLE `hk_pay_type` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `icon` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
   `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
   `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果',
   `plat_appid` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付平台appid',
   `plat_key` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付平台key',
   `plat_security` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付平台security',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `name` (`name`) USING BTREE,
   KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='支付类型';

-- ----------------------------
-- Table structure for hk_pay_type
-- ----------------------------
DROP TABLE IF EXISTS `hk_pay_type`;
CREATE TABLE `hk_pay_type`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `icon` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果',
  `plat_appid` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付平台appid',
  `plat_key` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '支付平台key',
  `plat_security` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '支付平台security',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '支付类型' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_pay_type
-- ----------------------------
INSERT INTO `hk_pay_type` VALUES (1, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', '支付宝', 'alipay', '1', NULL, NULL, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_pay_type` VALUES (2, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', '微信', 'wx', '1', NULL, NULL, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);


-- ----------------------------
-- Table structure for hk_extract_type
-- ----------------------------
DROP TABLE IF EXISTS `hk_extract_type`;
CREATE TABLE `hk_extract_type`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `icon` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '提现类型' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_pay_type
-- ----------------------------
INSERT INTO `hk_extract_type` VALUES (1, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', '支付宝', 'alipay', '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_extract_type` VALUES (2, '000000', 'https://hk-uploadfiles.oss-cn-hangzhou.aliyuncs.com/upload/20230221/824d662f5440cf8180ac77fb1e2d68e0.png', '微信', 'wx', '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);

-- ----------------------------
-- Table structure for hk_extract_type
-- ----------------------------
DROP TABLE IF EXISTS `hk_extract_type`;
CREATE TABLE `hk_extract_type` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `icon` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
   `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
   `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标识: bank = 银行卡 alipay = 支付宝wx=微信',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `name` (`name`) USING BTREE,
   KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='提现类型';


-- ----------------------------
-- Table structure for hk_product_gold
-- ----------------------------
DROP TABLE IF EXISTS `hk_product_gold`;
CREATE TABLE `hk_product_gold`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tenant_id` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `count`  bigint(20) NULL DEFAULT 0 COMMENT '数量',
  `price`  bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '商品价格',
  `cost` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '成本价',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `status` int(2) NULL DEFAULT 0 COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '商品' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_product_gold
-- ----------------------------
INSERT INTO `hk_product_gold` VALUES (1, '000000', 60, 600, 600, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_product_gold` VALUES (2, '000000', 300, 3000, 3000, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_product_gold` VALUES (3, '000000', 980, 9800, 9800, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_product_gold` VALUES (4, '000000', 1980, 19800, 19800, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_product_gold` VALUES (5, '000000', 2980, 29800, 29800, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);
INSERT INTO `hk_product_gold` VALUES (6, '000000', 4180, 41800, 41800, '2023-05-18 14:12:18.000', '2023-05-18 14:12:18.000', NULL, NULL, NULL, NULL, 0);

-- ----------------------------
-- Table structure for hk_user_bill
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_bill`;
CREATE TABLE `hk_user_bill` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `link_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '关联id',
   `pm` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0 = 支出 1 = 获得',
   `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账单标题',
   `category` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '明细种类',
   `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账单类型 1收入 2创作收益 3提现 4消费 5其他',
   `number` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '明细数字',
   `balance` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '剩余',
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
-- Table structure for hk_gold_bill
-- ----------------------------
DROP TABLE IF EXISTS `hk_gold_bill`;
CREATE TABLE `hk_gold_bill` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `link_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '关联id',
   `pm` int(2) NOT NULL DEFAULT '0' COMMENT '0 = 支出 1 = 获得',
   `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账单标题',
   `description` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '描述',
   `type` int(2) NOT NULL DEFAULT '5' COMMENT '账单类型 1充值 2付费 3红包 4打赏 5其他',
   `child_type` int(2) NOT NULL DEFAULT '0' COMMENT '子类型',
   `before_number` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '之前金额',
   `change_number` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '改变金额',
   `balance` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '余额',
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
   KEY `pm` (`pm`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='金币账单';

-- ----------------------------
-- Table structure for hk_user_recharge
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_recharge`;
CREATE TABLE `hk_user_recharge` (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
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
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
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
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `order_id` varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单号',
   `extend_order_id` varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '额外订单号',
   `user_id` bigint(20) NOT NULL COMMENT '用户编号',
   `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
   `cart_num` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '商品数量',
   `total_price` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '订单总价',
   `pay_price` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '实际支付金额',
   `cost` bigint(20) unsigned NOT NULL COMMENT '成本价',
   `paid` int(2) unsigned NOT NULL DEFAULT 0 COMMENT '支付状态: 0未支付 1已支付',
   `pay_time` datetime(3) DEFAULT NULL COMMENT '支付时间',
   `pay_type` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付方式',
   `refund_status` int(2) unsigned NOT NULL DEFAULT 0 COMMENT '0 未退款 1 申请中 2 已退款',
   `refund_reason_wap_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '退款图片',
   `refund_reason_wap_explain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '退款用户说明',
   `refund_reason_time` datetime(3) DEFAULT NULL COMMENT '退款时间',
   `refund_reason_wap` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '前台退款原因',
   `refund_reason` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '不退款的理由',
   `refund_price` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '退款金额',
   `gain_integral` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '消费赚取积分',
   `back_integral` bigint(20) unsigned DEFAULT NULL COMMENT '给用户退了多少积分',
   `mark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
   `remark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '管理员备注',  
   `verify_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '核销码',
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
   `browse_data` varchar(2048) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '备注',
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
-- Records of hk_record_browsing_user_homepage
-- ----------------------------
INSERT INTO `hk_record_browsing_user_homepage` VALUES (1, NULL, 2, 2, NULL, NULL, '2023-04-18 16:53:40.000', '2023-04-18 16:53:44.000', NULL, NULL, NULL, NULL, NULL, NULL);

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

-- ----------------------------
-- Table structure for hk_activity_add_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_activity_add_request`;
CREATE TABLE `hk_activity_add_request`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `user_id` bigint(20) NOT NULL COMMENT '用户ID',
   `activity_id` bigint(20) NOT NULL COMMENT '活动_编号',
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '加入活动申请';

-- ----------------------------
-- Table structure for hk_third_platform
-- ----------------------------
DROP TABLE IF EXISTS `hk_third_platform`;
CREATE TABLE `hk_third_platform`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `platform` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台',
   `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台名称',
   `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台图标',
   `appid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '平台的应用编号',
   `key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '秘钥',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '第三方平台';


-- ----------------------------
-- Table structure for hk_app_version
-- ----------------------------
DROP TABLE IF EXISTS `hk_app_version`;
CREATE TABLE `hk_app_version`  (
   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
   `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
   `platform` int(2) NULL DEFAULT 0 COMMENT '平台：1IOS,2Android,3Windows,4OSX,5Web,6MiniWeb,7Linux,8APad,9IPad,10Admin',
   `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本',
   `force_update` int(2) NULL DEFAULT 0 COMMENT '强制更新：0否 1是',
   `packet_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
   `packet_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '下载地址',
   `update_log` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '更新日志',
   `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
   `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
   `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
   `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
   `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
   `status` int(2) NULL DEFAULT 0 COMMENT '状态',
   `is_del` int(2) NULL DEFAULT 0 COMMENT '是否已删除',
   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'app版本';