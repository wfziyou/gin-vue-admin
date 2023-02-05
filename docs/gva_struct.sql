/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : gva

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 05/02/2023 22:41:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 534 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理角色ID',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `exa_file_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint(20) NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `chunk_total` bigint(20) NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_apply`;
CREATE TABLE `hk_apply`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `ower_type` smallint(6) NULL DEFAULT NULL COMMENT '拥有者：0平台、1圈子、2个人',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户_编号',
  `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `type` smallint(6) NULL DEFAULT NULL COMMENT '类型(0小程序、1第三方链接)',
  `mini_program_id` mediumint(9) NULL DEFAULT NULL COMMENT '小程序id',
  `apply_address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问地址',
  `apply_parameters` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问参数',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_apply_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_bug_report
-- ----------------------------
DROP TABLE IF EXISTS `hk_bug_report`;
CREATE TABLE `hk_bug_report`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户编号',
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作步骤',
  `content_attachment` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作步骤附件',
  `expected_result` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '预期结果',
  `actual_result` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实际结果',
  `actual_result_attachment` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '实际结果附件',
  `other_info` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '其他信息',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态：0 未处理 1 处理中 2 拒绝 3 完成',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_bug_report_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle`;
CREATE TABLE `hk_circle`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `type` smallint(6) NULL DEFAULT NULL COMMENT '类型：0官方圈子、1用户圈子、2小区',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
  `logo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子Logo',
  `circle_classify_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子分类_编号',
  `slogan` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子标语',
  `des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子简介',
  `protocol` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子规约',
  `back_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子背景图',
  `support_category` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议',
  `pay` smallint(6) NULL DEFAULT NULL COMMENT '付费：0 否，1是',
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
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_add_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_add_request`;
CREATE TABLE `hk_circle_add_request`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '申请理由',
  `check_user` mediumint(9) NULL DEFAULT NULL COMMENT '审核人',
  `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
  `check_status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态：0 未处理 1 通过，2驳回',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_add_request_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_apply`;
CREATE TABLE `hk_circle_apply`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `apply_group_id` mediumint(9) NULL DEFAULT NULL COMMENT '应用分组_编号',
  `show_name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称别名',
  `apply_id` mediumint(9) NULL DEFAULT NULL COMMENT '应用_编号',
  `apply_parameters` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问参数',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_apply_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_apply_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_apply_group`;
CREATE TABLE `hk_circle_apply_group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `parent_id` mediumint(9) NULL DEFAULT NULL COMMENT '父节点编号',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_apply_group_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_classify
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_classify`;
CREATE TABLE `hk_circle_classify`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `parent_id` mediumint(9) NULL DEFAULT NULL COMMENT '父节点编号',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `des` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `property` smallint(6) NULL DEFAULT NULL COMMENT '属性：0公开 ，1受限',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_classify_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_relation
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_relation`;
CREATE TABLE `hk_circle_relation`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `relation_type` smallint(6) NULL DEFAULT NULL COMMENT '关系类型：0父子节点 1关注',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `other_circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_relation_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_request
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_request`;
CREATE TABLE `hk_circle_request`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `type` smallint(6) NULL DEFAULT NULL COMMENT '类型：0官方圈子 ，1用户圈子',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子名称',
  `logo` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子Logo',
  `circle_classify_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子分类_编号',
  `slogan` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子标语',
  `des` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子简介',
  `protocol` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子规约',
  `back_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '圈子背景图',
  `check_status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态：0 未处理 1 通过，2驳回',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_request_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_circle_user
-- ----------------------------
DROP TABLE IF EXISTS `hk_circle_user`;
CREATE TABLE `hk_circle_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户ID',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '用户的圈子排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_circle_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_comment_thumbs_up
-- ----------------------------
DROP TABLE IF EXISTS `hk_comment_thumbs_up`;
CREATE TABLE `hk_comment_thumbs_up`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户id',
  `time` datetime(3) NULL DEFAULT NULL COMMENT '点赞时间',
  `comment_id` mediumint(9) NULL DEFAULT NULL COMMENT '评论编号',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_comment_thumbs_up_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forbidden_speak
-- ----------------------------
DROP TABLE IF EXISTS `hk_forbidden_speak`;
CREATE TABLE `hk_forbidden_speak`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户编号',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `duration_id` mediumint(9) NULL DEFAULT NULL COMMENT '禁言时长_编号',
  `reason_id` mediumint(9) NULL DEFAULT NULL COMMENT '禁言原因_编号',
  `cur_status` smallint(6) NULL DEFAULT NULL COMMENT '当前状态：0 未解锁、1已解锁',
  `unlock_user_id` mediumint(9) NULL DEFAULT NULL COMMENT '解锁用户_编号',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forbidden_speak_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forbidden_speak_duration
-- ----------------------------
DROP TABLE IF EXISTS `hk_forbidden_speak_duration`;
CREATE TABLE `hk_forbidden_speak_duration`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `count` smallint(6) NULL DEFAULT NULL COMMENT '时长',
  `type` smallint(6) NULL DEFAULT NULL COMMENT '时间类型：0 小时、1天',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forbidden_speak_duration_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forbidden_speak_reason
-- ----------------------------
DROP TABLE IF EXISTS `hk_forbidden_speak_reason`;
CREATE TABLE `hk_forbidden_speak_reason`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `reason` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '禁言理由',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forbidden_speak_reason_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_comment
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_comment`;
CREATE TABLE `hk_forum_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `posts_id` mediumint(9) NULL DEFAULT NULL COMMENT '帖子编号',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '评论者',
  `comment_time` datetime(3) NULL DEFAULT NULL COMMENT '评论时间',
  `comment_content` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '评论内容',
  `like_times` smallint(6) NULL DEFAULT NULL COMMENT '点赞数',
  `parent_id` mediumint(9) NULL DEFAULT NULL COMMENT '父评论编号',
  `check_user` mediumint(9) NULL DEFAULT NULL COMMENT '审核人',
  `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
  `check_status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态(0未审批 1通过 2拒绝)',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_comment_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_posts
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_posts`;
CREATE TABLE `hk_forum_posts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `category` smallint(6) NULL DEFAULT NULL COMMENT '类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）',
  `group_id` mediumint(9) NULL DEFAULT NULL COMMENT '帖子分类编号',
  `title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `seo_key` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO关键词',
  `seo_introduce` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介(SEO简介)',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `source` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '来源',
  `time` datetime(3) NULL DEFAULT NULL COMMENT '发布时间',
  `content` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '内容',
  `content_type` smallint(6) NULL DEFAULT NULL COMMENT '内容类型:0 markdown,1 html',
  `content_markdown` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'markdown内容',
  `content_html` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'html内容',
  `video` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '视频地址',
  `attachment` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附件',
  `tag` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `top` smallint(6) NULL DEFAULT NULL COMMENT '置顶(0否 1是)',
  `marrow` smallint(6) NULL DEFAULT NULL COMMENT '精华(0否 1是)',
  `comment_id` mediumint(9) NULL DEFAULT NULL COMMENT '问答最佳答案ID',
  `anonymity` smallint(6) NULL DEFAULT NULL COMMENT '匿名发布(0否 1是)',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '发布者编号',
  `read_num` smallint(6) NULL DEFAULT NULL COMMENT '阅读次数',
  `comment_num` smallint(6) NULL DEFAULT NULL COMMENT '评论次数',
  `share_num` smallint(6) NULL DEFAULT NULL COMMENT '分享次数',
  `collect_num` smallint(6) NULL DEFAULT NULL COMMENT '收藏次数',
  `like_num` smallint(6) NULL DEFAULT NULL COMMENT '点赞次数',
  `check_user` mediumint(9) NULL DEFAULT NULL COMMENT '审核人',
  `check_time` datetime(3) NULL DEFAULT NULL COMMENT '审核时间',
  `check_status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态(0草稿 1未审批 2通过 3拒绝)',
  `power_comment` smallint(6) NULL DEFAULT NULL COMMENT '评论权限(0关闭 1开启)',
  `power_comment_anonymity` smallint(6) NULL DEFAULT NULL COMMENT '匿名评论(0关闭 1开启)',
  `pay` smallint(6) NULL DEFAULT NULL COMMENT '付费：0 否，1是',
  `pay_content` smallint(6) NULL DEFAULT NULL COMMENT '内容付费：0 否，1是',
  `pay_content_look` smallint(6) NULL DEFAULT NULL COMMENT '内容付费可查看百分比例',
  `pay_attachment` smallint(6) NULL DEFAULT NULL COMMENT '附件付费：0 否，1是',
  `pay_currency` smallint(6) NULL DEFAULT NULL COMMENT '付费货币：0 人民、1积分、2代币',
  `pay_num` smallint(6) NULL DEFAULT NULL COMMENT '付费金额',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_posts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_posts_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_posts_group`;
CREATE TABLE `hk_forum_posts_group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `parent_id` mediumint(9) NULL DEFAULT NULL COMMENT '父节点编号',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态(0未审批 1通过 2拒绝)',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_posts_group_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_tag
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_tag`;
CREATE TABLE `hk_forum_tag`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `group_id` mediumint(9) NULL DEFAULT NULL COMMENT '分组id',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_tag_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_tag_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_tag_group`;
CREATE TABLE `hk_forum_tag_group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_tag_group_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_thumbs_up
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_thumbs_up`;
CREATE TABLE `hk_forum_thumbs_up`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户id',
  `time` datetime(3) NULL DEFAULT NULL COMMENT '点赞时间',
  `posts_id` mediumint(9) NULL DEFAULT NULL COMMENT '帖子编号',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_thumbs_up_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_topic
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_topic`;
CREATE TABLE `hk_forum_topic`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `cover_image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '封面',
  `topic_group_id` mediumint(9) NULL DEFAULT NULL COMMENT '分组id',
  `type` smallint(6) NULL DEFAULT NULL COMMENT '话题类型：0 全局，1圈子',
  `hot` smallint(6) NULL DEFAULT NULL COMMENT '是否热门：0 否 1 是',
  `discuss_num` smallint(6) NULL DEFAULT NULL COMMENT '讨论数',
  `attention_num` smallint(6) NULL DEFAULT NULL COMMENT '关注数',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `review_status` smallint(6) NULL DEFAULT NULL COMMENT '审核状态：0 未处理 1 通过，2驳回',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_topic_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_forum_topic_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_forum_topic_group`;
CREATE TABLE `hk_forum_topic_group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_forum_topic_group_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_mini_program
-- ----------------------------
DROP TABLE IF EXISTS `hk_mini_program`;
CREATE TABLE `hk_mini_program`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `company_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称',
  `cur_version` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前版本',
  `cur_packet_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '当前包id',
  `hidden` smallint(6) NULL DEFAULT NULL COMMENT '隐藏(0否 1是)',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_mini_program_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_mini_program_packet
-- ----------------------------
DROP TABLE IF EXISTS `hk_mini_program_packet`;
CREATE TABLE `hk_mini_program_packet`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `mini_program_id` mediumint(9) NULL DEFAULT NULL COMMENT '小程序id',
  `name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
  `version` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本',
  `packet_address` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问地址',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_mini_program_packet_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_plat_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_plat_apply`;
CREATE TABLE `hk_plat_apply`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `apply_group_id` mediumint(9) NULL DEFAULT NULL COMMENT '应用分组_编号',
  `show_name` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称别名',
  `apply_id` mediumint(9) NULL DEFAULT NULL COMMENT '应用_编号',
  `apply_parameters` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '访问参数',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_plat_apply_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_plat_apply_group
-- ----------------------------
DROP TABLE IF EXISTS `hk_plat_apply_group`;
CREATE TABLE `hk_plat_apply_group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `parent_id` mediumint(9) NULL DEFAULT NULL COMMENT '父节点编号',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_plat_apply_group_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_protocol
-- ----------------------------
DROP TABLE IF EXISTS `hk_protocol`;
CREATE TABLE `hk_protocol`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议名称',
  `des` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议说明',
  `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '协议内容',
  `module` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所属模块/插件',
  `pos` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '显示位置',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_protocol_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_report
-- ----------------------------
DROP TABLE IF EXISTS `hk_report`;
CREATE TABLE `hk_report`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `report_user_id` mediumint(9) NULL DEFAULT NULL COMMENT '被举报用户编号',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '举报用户编号',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `reason_id` mediumint(9) NULL DEFAULT NULL COMMENT '举报原因_编号',
  `report_type` smallint(6) NULL DEFAULT NULL COMMENT '举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题',
  `content` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报内容',
  `content_attachment` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '内容附件',
  `cur_status` smallint(6) NULL DEFAULT NULL COMMENT '处理状态：0 未处理、1已处理',
  `handle_user_id` mediumint(9) NULL DEFAULT NULL COMMENT '操作用户_编号',
  `handle_type` mediumint(9) NULL DEFAULT NULL COMMENT '处理方式:0无效处理（不予处理）、账号禁言',
  `duration_id` mediumint(9) NULL DEFAULT NULL COMMENT '禁言时长_编号',
  `handle_score` smallint(6) NULL DEFAULT NULL COMMENT '积分处理:0不扣分、1扣分',
  `score_experience` smallint(6) NULL DEFAULT NULL COMMENT '经验',
  `score_community` smallint(6) NULL DEFAULT NULL COMMENT '社区积分',
  `score_buy` smallint(6) NULL DEFAULT NULL COMMENT '购物积分',
  `score_download` smallint(6) NULL DEFAULT NULL COMMENT '下载值',
  `unlock` smallint(6) NULL DEFAULT NULL COMMENT '是否解除：0 否、是',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_report_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_report_reason
-- ----------------------------
DROP TABLE IF EXISTS `hk_report_reason`;
CREATE TABLE `hk_report_reason`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `reason` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '举报理由',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_report_reason_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_user
-- ----------------------------
DROP TABLE IF EXISTS `hk_user`;
CREATE TABLE `hk_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户编号',
  `user_type` smallint(6) NULL DEFAULT NULL COMMENT '用户平台',
  `account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '账号',
  `password` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真名',
  `header_img` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `email` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `birthday` datetime(3) NULL DEFAULT NULL COMMENT '生日',
  `sex` smallint(6) NULL DEFAULT NULL COMMENT '性别',
  `role_id` mediumint(8) UNSIGNED NULL DEFAULT NULL COMMENT '用户角色ID',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `is_deleted` smallint(5) UNSIGNED NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_user_browsing_history
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_browsing_history`;
CREATE TABLE `hk_user_browsing_history`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户编号',
  `posts_id` mediumint(9) NULL DEFAULT NULL COMMENT '帖子编号',
  `time` datetime(3) NULL DEFAULT NULL COMMENT '浏览时间',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_browsing_history_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_user_circle_apply
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_circle_apply`;
CREATE TABLE `hk_user_circle_apply`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户编号',
  `circle_id` mediumint(9) NULL DEFAULT NULL COMMENT '圈子_编号',
  `apply_id` mediumint(9) NULL DEFAULT NULL COMMENT '应用_编号',
  `sort` smallint(6) NULL DEFAULT NULL COMMENT '排序',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_circle_apply_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_user_collect
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_collect`;
CREATE TABLE `hk_user_collect`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户编号',
  `posts_id` mediumint(9) NULL DEFAULT NULL COMMENT '帖子编号',
  `time` datetime(3) NULL DEFAULT NULL COMMENT '收藏时间',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_collect_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for hk_user_extend
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_extend`;
CREATE TABLE `hk_user_extend`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  `user_id` mediumint(9) NULL DEFAULT NULL COMMENT '用户ID',
  `github` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'github',
  `weibo` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微博',
  `weixin` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信',
  `qq` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'qq',
  `blog` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '博客',
  `num_circle` mediumint(9) NULL DEFAULT NULL COMMENT '圈子数',
  `num_focus` mediumint(9) NULL DEFAULT NULL COMMENT '关注数',
  `num_fan` mediumint(9) NULL DEFAULT NULL COMMENT '粉丝数',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_hk_user_extend_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `is_del` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 306 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `authority_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9529 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `flag` bigint(20) NULL DEFAULT NULL,
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_auto_codes
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_codes`;
CREATE TABLE `sys_auto_codes`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_codes_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `menu_level` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) NULL DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) NULL DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 79 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  `hk_user_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint(20) UNSIGNED NULL DEFAULT 888 COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint(20) NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  `create_time` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) NULL DEFAULT NULL COMMENT '修改时间',
  `update_user` mediumint(9) NULL DEFAULT NULL COMMENT '修改人',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '状态',
  `is_deleted` smallint(6) NULL DEFAULT NULL COMMENT '是否已删除',
  `create_user` mediumint(9) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` mediumint(9) NULL DEFAULT NULL COMMENT '创建部门',
  `is_del` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_username`(`username`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_users_uuid`(`uuid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
