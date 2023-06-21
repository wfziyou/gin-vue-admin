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
-- Table structure for hk_shop_channel
-- ----------------------------
DROP TABLE IF EXISTS `hk_shop_channel`;
CREATE TABLE `hk_shop_channel`  (
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
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '圈子频道' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of hk_shop_channel
-- ----------------------------
INSERT INTO `hk_shop_channel` VALUES (1, '000000', '推荐', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_shop_channel` VALUES (2, '000000', '家用', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_shop_channel` VALUES (3, '000000', '户外', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);
INSERT INTO `hk_shop_channel` VALUES (4, '000000', '学生', '2023-05-09 17:57:30.000', '2023-05-09 17:57:30.000', NULL, NULL, NULL, NULL, 0, 0);

-- ----------------------------
-- Table structure for hk_shop_product
-- ----------------------------
DROP TABLE IF EXISTS `hk_shop_product`;
CREATE TABLE `hk_shop_product` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `image` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品图片',
  `slider_image` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '轮播图',
  `store_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品名称',
  `store_info` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品简介',
  `keyword` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '关键字',
  `bar_code` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '产品条码（一维码）',
  `cate_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分类id',
  `price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品价格',
  `vip_price` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '会员价格',
  `ot_price` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '市场价',
  `postage` decimal(8,2) unsigned DEFAULT '0.00' COMMENT '邮费',
  `unit_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '单位名',
  `sort` smallint DEFAULT '0' COMMENT '排序',
  `sales` mediumint unsigned DEFAULT '0' COMMENT '销量',
  `stock` mediumint unsigned DEFAULT '0' COMMENT '库存',
  `is_show` tinyint(1) DEFAULT '1' COMMENT '状态（0：未上架，1：上架）',
  `is_hot` tinyint(1) DEFAULT '0' COMMENT '是否热卖',
  `is_benefit` tinyint(1) DEFAULT '0' COMMENT '是否优惠',
  `is_best` tinyint(1) DEFAULT '0' COMMENT '是否精品',
  `is_new` tinyint(1) DEFAULT '0' COMMENT '是否新品',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '产品描述',
  `create_time` datetime DEFAULT NULL COMMENT '添加时间',
  `update_time` datetime DEFAULT NULL,
  `is_postage` tinyint unsigned DEFAULT '0' COMMENT '是否包邮',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否删除',
  `mer_use` tinyint unsigned DEFAULT '0' COMMENT '商户是否代理 0不可代理1可代理',
  `give_integral` decimal(8,2) unsigned DEFAULT NULL COMMENT '获得积分',
  `cost` decimal(8,2) unsigned DEFAULT NULL COMMENT '成本价',
  `is_good` tinyint(1) DEFAULT '0' COMMENT '是否优品推荐',
  `ficti` mediumint DEFAULT '100' COMMENT '虚拟销量',
  `browse` int DEFAULT '0' COMMENT '浏览量',
  `is_sub` tinyint(1) DEFAULT '0' COMMENT '是否单独分佣',
  `temp_id` int DEFAULT NULL COMMENT '运费模板ID',
  `spec_type` tinyint(1) DEFAULT '0' COMMENT '规格 0单 1多',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_hot` (`is_hot`) USING BTREE,
  KEY `is_benefit` (`is_benefit`) USING BTREE,
  KEY `is_best` (`is_best`) USING BTREE,
  KEY `is_new` (`is_new`) USING BTREE,
  KEY `toggle_on_sale, is_del` (`is_del`) USING BTREE,
  KEY `price` (`price`) USING BTREE,
  KEY `is_show` (`is_show`) USING BTREE,
  KEY `sort` (`sort`) USING BTREE,
  KEY `sales` (`sales`) USING BTREE,
  KEY `add_time` (`create_time`) USING BTREE,
  KEY `is_postage` (`is_postage`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='商品表';

-- ----------------------------
-- Records of hk_shop_product
-- ----------------------------
BEGIN;
INSERT INTO `hk_shop_product` VALUES (1, 'https://image.dayouqiantu.cn/5ca04fa9c08ef.jpg', 'https://image.dayouqiantu.cn/5ca081af6183f.jpg,https://image.dayouqiantu.cn/5ca04fa9c08ef.jpg', '照片打印', '照片打印', '照片打印', '', '5', 2.00, 0.00, 1.00, 0.00, '张', 0, 8, 992, 1, 0, 0, 0, 0, '<p><br/></p><p><img src=\"https://image.dayouqiantu.cn/5ca04fa9c08ef.jpg\"/></p>', '2020-09-04 17:41:30', '2020-09-05 20:57:49', 0, 0, 0, 0.00, 3.00, 0, 0, 3, 0, 34, 0);
INSERT INTO `hk_shop_product` VALUES (2, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg,https://consoleapi.xinxintuan.co/file/pic/20200903171807647067.jpg,https://consoleapi.xinxintuan.co/file/pic/20200903171807641535.jpg', 'X型展架', 'X型展架', 'X型展架', '', '3', 10.00, 0.00, 12.00, 0.00, '张', 0, 20, 3310, 1, 1, 1, 0, 0, '<p><br/></p><p><img src=\"https://consoleapi.xinxintuan.co/file/pic/20200903171807641535.jpg\"/></p>', '2020-09-04 17:55:16', '2021-04-04 15:13:02', 0, 1, 0, 0.00, 11.00, 0, 69, 25, 0, 34, 1);
INSERT INTO `hk_shop_product` VALUES (3, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807647067.jpg', 'https://consoleapi.xinxintuan.co/file/pic/20200903171807647067.jpg', '彩色复印-量大有优惠', '彩色复印', '彩色复印', '', '3', 2.00, 0.00, 1.00, 0.00, '张', 0, 8, 9996, 1, 1, 1, 0, 0, '<p>1111</p>', '2020-09-07 13:14:50', '2020-09-13 22:18:46', 0, 1, 0, 0.00, 1.00, 0, 0, 43, 0, 34, 0);
INSERT INTO `hk_shop_product` VALUES (4, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807625149.jpg', 'https://consoleapi.xinxintuan.co/file/pic/20200903171807625149.jpg', '户外写真背胶', '户外写真背胶', '户外写真背胶', '', '3', 2.00, 0.00, 1.00, 0.00, '米', 0, 24, 87, 1, 1, 1, 0, 0, '<p>户外写真背胶</p>', '2020-09-07 13:16:02', NULL, 0, 1, 0, 0.00, 1.00, 0, 0, 32, 0, 34, 0);
INSERT INTO `hk_shop_product` VALUES (5, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807568739.jpg', 'https://consoleapi.xinxintuan.co/file/pic/20200903171807568739.jpg', '广告版定制', '广告版定制', '广告版定制', '', '3', 2.00, 0.00, 11.00, 0.00, '米', 0, 16, 2, 1, 1, 1, 1, 1, '<p>广告版定制</p>', '2020-09-07 13:17:49', '2020-09-12 00:14:32', 0, 1, 0, 0.00, 1.00, 0, 0, 26, 0, 34, 0);
INSERT INTO `hk_shop_product` VALUES (6, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807641535.jpg', 'https://consoleapi.xinxintuan.co/file/pic/20200903171807641535.jpg', '印刷各种印刷', '印刷各种印刷', '印刷各种印刷', '', '3', 11.00, 0.00, 33.00, 0.00, '张', 1, 35, 9965, 1, 1, 1, 0, 0, '<p>印刷各种印刷</p>', '2020-09-07 13:18:45', '2020-09-11 17:18:18', 0, 1, 0, 1.00, 22.00, 0, 1, 139, 0, 34, 0);
INSERT INTO `hk_shop_product` VALUES (7, 'https://consoleapi.xinxintuan.co/file/pic/20200903171208889668.png', 'https://image.dayouqiantu.cn/5ca04fa9c08ef.jpg,https://image.dayouqiantu.cn/5ca04fa9c08ef.jpg,https://image.dayouqiantu.cn/5ca081af6183f.jpg', '专属测试商品', '', '专属测试商品', '', '3', 100.00, 0.00, 120.00, 0.00, '张', 0, 3, 7, 1, 0, 0, 0, 0, '<p><br/></p><p><img src=\"https://consoleapi.xinxintuan.co/file/pic/20200903171807647067.jpg\"/></p>', '2020-09-12 16:11:05', '2020-09-12 17:03:58', 0, 1, 0, 10.00, 110.00, 0, 9, 9, 0, 34, 0);
INSERT INTO `hk_shop_product` VALUES (8, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', 'https://image.dayouqiantu.cn/5ca081af6183f.jpg,https://image.dayouqiantu.cn/5ca081af6183f.jpg,https://image.dayouqiantu.cn/5ca04fa9c08ef.jpg', '多规格商品测试', '多规格商品测试', '多规格商品测试', '', '5', 4.00, 0.00, 9.00, 0.00, '张', 0, 5, 501, 1, 1, 1, 1, 1, '<p><br/></p><p><img src=\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\"/></p>', '2020-09-12 17:24:51', NULL, 0, 0, 0, 0.00, 1.00, 0, 0, 19, 1, 34, 1);
COMMIT;


-- ----------------------------
-- Table structure for hk_shop_product_attr
-- ----------------------------
DROP TABLE IF EXISTS `hk_shop_product_attr`;
CREATE TABLE `hk_shop_product_attr` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `attr_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '属性名',
  `attr_values` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '属性值',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `store_id` (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='商品属性表';

-- ----------------------------
-- Records of hk_shop_product_attr
-- ----------------------------
BEGIN;
INSERT INTO `hk_shop_product_attr` VALUES (32, 1, '规格', '默认');
INSERT INTO `hk_shop_product_attr` VALUES (55, 4, '规格', '默认');
INSERT INTO `hk_shop_product_attr` VALUES (57, 2, '纸张', 'A4,A3');
INSERT INTO `hk_shop_product_attr` VALUES (58, 2, '颜色', '白色,红色');
INSERT INTO `hk_shop_product_attr` VALUES (59, 0, '规格', '默认');
INSERT INTO `hk_shop_product_attr` VALUES (60, 6, '规格', '默认');
INSERT INTO `hk_shop_product_attr` VALUES (62, 7, '规格', '默认');
INSERT INTO `hk_shop_product_attr` VALUES (67, 8, '纸张', 'A4,A3,A5');
INSERT INTO `hk_shop_product_attr` VALUES (68, 8, '颜色', '白色,红色');
INSERT INTO `hk_shop_product_attr` VALUES (69, 5, '规格', '默认');
INSERT INTO `hk_shop_product_attr` VALUES (70, 3, '规格', '默认');
COMMIT;

-- ----------------------------
-- Table structure for hk_shop_product_attr_result
-- ----------------------------
DROP TABLE IF EXISTS `hk_shop_product_attr_result`;
CREATE TABLE `hk_shop_product_attr_result` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `result` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品属性参数',
  `change_time` datetime NOT NULL COMMENT '上次修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `product_id` (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='商品属性详情表';

-- ----------------------------
-- Records of hk_shop_product_attr_result
-- ----------------------------
BEGIN;
INSERT INTO `hk_shop_product_attr_result` VALUES (26, 1, '{\"attr\":[{\"attrHidden\":\"\",\"detail\":[\"默认\"],\"detailValue\":\"\",\"value\":\"规格\"}],\"value\":[{\"barCode\":\"00005\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":3.0,\"detail\":{\"规格\":\"默认\"},\"otPrice\":1.0,\"pic\":\"\",\"pinkPrice\":0.1,\"pinkStock\":100,\"price\":2.0,\"seckillPrice\":0.2,\"seckillStock\":100,\"stock\":999,\"value1\":\"规格\",\"value2\":\"\",\"volume\":0.0,\"weight\":0.0}]}', '2020-09-08 21:18:03');
INSERT INTO `hk_shop_product_attr_result` VALUES (44, 4, '{\"attr\":[{\"attrHidden\":\"\",\"detail\":[\"默认\"],\"detailValue\":\"\",\"value\":\"规格\"}],\"value\":[{\"barCode\":\"\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":0.0,\"detail\":{\"规格\":\"默认\"},\"otPrice\":0.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg\",\"pinkPrice\":22.0,\"pinkStock\":18,\"price\":0.0,\"seckillPrice\":0.0,\"seckillStock\":0,\"stock\":93,\"value1\":\"规格\",\"volume\":0.0,\"weight\":0.0}]}', '2020-09-12 15:55:17');
INSERT INTO `hk_shop_product_attr_result` VALUES (46, 2, '{\"attr\":[{\"detail\":[\"A4\",\"A3\"],\"value\":\"纸张\"},{\"detail\":[\"白色\",\"红色\"],\"value\":\"颜色\"}],\"value\":[{\"barCode\":\"00001\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":11.0,\"detail\":{\"纸张\":\"A4\",\"颜色\":\"白色\"},\"otPrice\":12.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg\",\"pinkPrice\":11.09,\"pinkStock\":100,\"price\":10.0,\"seckillPrice\":1.0,\"seckillStock\":100,\"stock\":991,\"value1\":\"A4\",\"value2\":\"白色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"00002\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":11.0,\"detail\":{\"纸张\":\"A4\",\"颜色\":\"红色\"},\"otPrice\":12.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg\",\"pinkPrice\":22.88,\"pinkStock\":99,\"price\":11.0,\"seckillPrice\":2.0,\"seckillStock\":100,\"stock\":661,\"value1\":\"A4\",\"value2\":\"红色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"00003\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":11.0,\"detail\":{\"纸张\":\"A3\",\"颜色\":\"白色\"},\"otPrice\":12.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg\",\"pinkPrice\":33.89,\"pinkStock\":99,\"price\":12.0,\"seckillPrice\":3.0,\"seckillStock\":100,\"stock\":995,\"value1\":\"A3\",\"value2\":\"白色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"00004\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":11.0,\"detail\":{\"纸张\":\"A3\",\"颜色\":\"红色\"},\"otPrice\":12.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg\",\"pinkPrice\":45.66,\"pinkStock\":100,\"price\":13.0,\"seckillPrice\":4.0,\"seckillStock\":100,\"stock\":665,\"value1\":\"A3\",\"value2\":\"红色\",\"volume\":0.0,\"weight\":0.0}]}', '2020-09-12 16:10:09');
INSERT INTO `hk_shop_product_attr_result` VALUES (48, 6, '{\"attr\":[{\"attrHidden\":\"\",\"detail\":[\"默认\"],\"detailValue\":\"\",\"value\":\"规格\"}],\"value\":[{\"barCode\":\"00002\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":22.0,\"detail\":{\"规格\":\"默认\"},\"otPrice\":33.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg\",\"pinkPrice\":19.9,\"pinkStock\":99,\"price\":11.0,\"seckillPrice\":22.98,\"seckillStock\":2000,\"stock\":9989,\"value1\":\"规格\",\"volume\":1.0,\"weight\":1.0}]}', '2020-09-12 16:11:31');
INSERT INTO `hk_shop_product_attr_result` VALUES (50, 7, '{\"attr\":[{\"attrHidden\":\"\",\"detail\":[\"默认\"],\"detailValue\":\"\",\"value\":\"规格\"}],\"value\":[{\"barCode\":\"\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":110.0,\"detail\":{\"规格\":\"默认\"},\"otPrice\":120.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200911093912577832.png\",\"pinkPrice\":0.0,\"pinkStock\":0,\"price\":100.0,\"seckillPrice\":0.0,\"seckillStock\":0,\"stock\":10,\"value1\":\"规格\",\"volume\":0.0,\"weight\":0.0}]}', '2020-09-12 17:03:58');
INSERT INTO `hk_shop_product_attr_result` VALUES (53, 8, '{\"attr\":[{\"detail\":[\"A4\",\"A3\",\"A5\"],\"value\":\"纸张\"},{\"detail\":[\"白色\",\"红色\"],\"value\":\"颜色\"}],\"value\":[{\"barCode\":\"\",\"brokerage\":1.0,\"brokerageTwo\":1.0,\"cost\":1.0,\"detail\":{\"纸张\":\"A4\",\"颜色\":\"白色\"},\"otPrice\":12.0,\"pic\":\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\",\"pinkPrice\":1.0,\"pinkStock\":9,\"price\":9.0,\"seckillPrice\":1.0,\"seckillStock\":10,\"stock\":96,\"value1\":\"A4\",\"value2\":\"白色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"\",\"brokerage\":1.0,\"brokerageTwo\":1.0,\"cost\":1.0,\"detail\":{\"纸张\":\"A4\",\"颜色\":\"红色\"},\"otPrice\":9.0,\"pic\":\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\",\"pinkPrice\":2.0,\"pinkStock\":10,\"price\":8.0,\"seckillPrice\":1.0,\"seckillStock\":10,\"stock\":99,\"value1\":\"A4\",\"value2\":\"红色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"\",\"brokerage\":1.0,\"brokerageTwo\":1.0,\"cost\":1.0,\"detail\":{\"纸张\":\"A3\",\"颜色\":\"白色\"},\"otPrice\":9.0,\"pic\":\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\",\"pinkPrice\":0.5,\"pinkStock\":10,\"price\":7.0,\"seckillPrice\":0.1,\"seckillStock\":10,\"stock\":88,\"value1\":\"A3\",\"value2\":\"白色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"\",\"brokerage\":1.0,\"brokerageTwo\":1.0,\"cost\":1.0,\"detail\":{\"纸张\":\"A3\",\"颜色\":\"红色\"},\"otPrice\":9.0,\"pic\":\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\",\"pinkPrice\":1.0,\"pinkStock\":9,\"price\":6.0,\"seckillPrice\":0.1,\"seckillStock\":9,\"stock\":88,\"value1\":\"A3\",\"value2\":\"红色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"\",\"brokerage\":1.0,\"brokerageTwo\":1.0,\"cost\":1.0,\"detail\":{\"纸张\":\"A5\",\"颜色\":\"白色\"},\"otPrice\":9.0,\"pic\":\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\",\"pinkPrice\":2.0,\"pinkStock\":9,\"price\":5.0,\"seckillPrice\":0.1,\"seckillStock\":9,\"stock\":65,\"value1\":\"A5\",\"value2\":\"白色\",\"volume\":0.0,\"weight\":0.0},{\"barCode\":\"\",\"brokerage\":1.0,\"brokerageTwo\":1.0,\"cost\":1.0,\"detail\":{\"纸张\":\"A5\",\"颜色\":\"红色\"},\"otPrice\":9.0,\"pic\":\"https://image.dayouqiantu.cn/5ca081af6183f.jpg\",\"pinkPrice\":3.0,\"pinkStock\":9,\"price\":4.0,\"seckillPrice\":0.1,\"seckillStock\":9,\"stock\":66,\"value1\":\"A5\",\"value2\":\"红色\",\"volume\":0.0,\"weight\":0.0}]}', '2020-09-12 18:23:33');
INSERT INTO `hk_shop_product_attr_result` VALUES (54, 5, '{\"attr\":[{\"attrHidden\":\"\",\"detail\":[\"默认\"],\"detailValue\":\"\",\"value\":\"规格\"}],\"value\":[{\"barCode\":\"1231321\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":1.0,\"detail\":{\"规格\":\"默认\"},\"otPrice\":11.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807641535.jpg\",\"pinkPrice\":29.0,\"pinkStock\":10,\"price\":2.0,\"seckillPrice\":11.99,\"seckillStock\":10,\"stock\":11,\"value1\":\"规格\",\"volume\":1.0,\"weight\":1.0}]}', '2020-09-13 10:15:24');
INSERT INTO `hk_shop_product_attr_result` VALUES (55, 3, '{\"attr\":[{\"attrHidden\":\"\",\"detail\":[\"默认\"],\"detailValue\":\"\",\"value\":\"规格\"}],\"value\":[{\"barCode\":\"\",\"brokerage\":0.0,\"brokerageTwo\":0.0,\"cost\":1.0,\"detail\":{\"规格\":\"默认\"},\"otPrice\":1.0,\"pic\":\"https://consoleapi.xinxintuan.co/file/pic/20200903171807647067.jpg\",\"pinkPrice\":119.99,\"pinkStock\":10,\"price\":2.0,\"seckillPrice\":0.0,\"seckillStock\":0,\"stock\":9999,\"value1\":\"规格\",\"volume\":0.0,\"weight\":0.0}]}', '2020-09-13 22:18:46');
COMMIT;

-- ----------------------------
-- Table structure for hk_shop_product_attr_value
-- ----------------------------
DROP TABLE IF EXISTS `hk_shop_product_attr_value`;
CREATE TABLE `hk_shop_product_attr_value` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `sku` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品属性索引值 (attr_value|attr_value[|....])',
  `stock` int unsigned NOT NULL COMMENT '属性对应的库存',
  `sales` int unsigned DEFAULT '0' COMMENT '销量',
  `price` decimal(8,2) unsigned NOT NULL COMMENT '属性金额',
  `image` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '图片',
  `unique` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '唯一值',
  `cost` decimal(8,2) unsigned NOT NULL COMMENT '成本价',
  `bar_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品条码',
  `ot_price` decimal(10,2) DEFAULT '0.00' COMMENT '原价',
  `weight` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '重量',
  `volume` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '体积',
  `brokerage` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '一级返佣',
  `brokerage_two` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '二级返佣',
  `pink_price` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '拼团价',
  `pink_stock` int NOT NULL DEFAULT '0' COMMENT '拼团库存',
  `seckill_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '秒杀价',
  `seckill_stock` int NOT NULL DEFAULT '0' COMMENT '秒杀库存',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique` (`unique`,`sku`) USING BTREE,
  KEY `store_id` (`product_id`,`sku`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=107 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='商品属性值表';

-- ----------------------------
-- Records of hk_shop_product_attr_value
-- ----------------------------
BEGIN;
INSERT INTO `hk_shop_product_attr_value` VALUES (44, 1, '默认', 993, 6, 2.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', 'd4ce8cd4bda84d479c11348e060afb8e', 3.00, '00005', 1.00, 0.00, 0.00, 0.00, 0.00, 0.10, 100, 0.20, 97);
INSERT INTO `hk_shop_product_attr_value` VALUES (77, 4, '默认', 91, 2, 0.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', '6a1919d9178b48f3a75fe4640e51693e', 0.00, '', 0.00, 0.00, 0.00, 0.00, 0.00, 22.00, 17, 0.00, 0);
INSERT INTO `hk_shop_product_attr_value` VALUES (79, 2, 'A4,白色', 987, 4, 10.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', '04e9d586cef8464aaee0a45439f12520', 11.00, '00001', 12.00, 0.00, 0.00, 0.00, 0.00, 11.09, 99, 1.00, 100);
INSERT INTO `hk_shop_product_attr_value` VALUES (80, 2, 'A4,红色', 661, 0, 11.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', '3d1ffa92e6064c36b614d3046c268968', 11.00, '00002', 12.00, 0.00, 0.00, 0.00, 0.00, 22.88, 99, 2.00, 100);
INSERT INTO `hk_shop_product_attr_value` VALUES (81, 2, 'A3,白色', 995, 0, 12.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', '344f3a0e9d644845ae483c8b9a84a112', 11.00, '00003', 12.00, 0.00, 0.00, 0.00, 0.00, 33.89, 99, 3.00, 100);
INSERT INTO `hk_shop_product_attr_value` VALUES (82, 2, 'A3,红色', 665, 0, 13.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', '4a9ca8587855412ea7be9ba53f829052', 11.00, '00004', 12.00, 0.00, 0.00, 0.00, 0.00, 45.66, 100, 4.00, 100);
INSERT INTO `hk_shop_product_attr_value` VALUES (83, 0, '默认', 10, 0, 100.00, 'https://consoleapi.xinxintuan.co/file/pic/20200911093912577832.png', 'bbdc071ff67c4892839eac9f62e7eb18', 110.00, '', 120.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 0.00, 0);
INSERT INTO `hk_shop_product_attr_value` VALUES (84, 6, '默认', 9966, 23, 11.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807584910.jpg', '90babee9cea24645870c6027fc420d42', 22.00, '00002', 33.00, 1.00, 1.00, 0.00, 0.00, 19.90, 95, 22.98, 2000);
INSERT INTO `hk_shop_product_attr_value` VALUES (86, 7, '默认', 7, 3, 100.00, 'https://consoleapi.xinxintuan.co/file/pic/20200911093912577832.png', '789984b858bb489bb6313107e11fdc15', 110.00, '', 120.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0, 0.00, 0);
INSERT INTO `hk_shop_product_attr_value` VALUES (99, 8, 'A4,白色', 95, 1, 9.00, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', '013f89ec740f478d8144f052909c0807', 1.00, '', 12.00, 0.00, 0.00, 1.00, 1.00, 1.00, 9, 1.00, 10);
INSERT INTO `hk_shop_product_attr_value` VALUES (100, 8, 'A4,红色', 99, 0, 8.00, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', '4e1d7c450a8e48f6ae4ec7788f0cfd90', 1.00, '', 9.00, 0.00, 0.00, 1.00, 1.00, 2.00, 10, 1.00, 10);
INSERT INTO `hk_shop_product_attr_value` VALUES (101, 8, 'A3,白色', 88, 0, 7.00, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', '3ce61e5e6f4b46a7a210862af85212bc', 1.00, '', 9.00, 0.00, 0.00, 1.00, 1.00, 0.50, 10, 0.10, 10);
INSERT INTO `hk_shop_product_attr_value` VALUES (102, 8, 'A3,红色', 88, 0, 6.00, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', 'b1e48ae88299479899d95b86dd1be700', 1.00, '', 9.00, 0.00, 0.00, 1.00, 1.00, 1.00, 9, 0.10, 9);
INSERT INTO `hk_shop_product_attr_value` VALUES (103, 8, 'A5,白色', 65, 0, 5.00, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', '81b7a068f84a498ca8ad4e1916011e00', 1.00, '', 9.00, 0.00, 0.00, 1.00, 1.00, 2.00, 9, 0.10, 9);
INSERT INTO `hk_shop_product_attr_value` VALUES (104, 8, 'A5,红色', 66, 0, 4.00, 'https://image.dayouqiantu.cn/5ca081af6183f.jpg', '8112da382da7420f961203372f7c1d31', 1.00, '', 9.00, 0.00, 0.00, 1.00, 1.00, 3.00, 9, 0.10, 9);
INSERT INTO `hk_shop_product_attr_value` VALUES (105, 5, '默认', 3, 8, 2.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807641535.jpg', '040fb7ab2a8b4ff3b6c9d1c78ea3e889', 1.00, '1231321', 11.00, 1.00, 1.00, 0.00, 0.00, 29.00, 5, 11.99, 8);
INSERT INTO `hk_shop_product_attr_value` VALUES (106, 3, '默认', 9996, 3, 2.00, 'https://consoleapi.xinxintuan.co/file/pic/20200903171807647067.jpg', '43cdbd1043474b4e97a1cffffa18071c', 1.00, '', 1.00, 0.00, 0.00, 0.00, 0.00, 119.99, 10, 0.00, 0);
COMMIT;

-- ----------------------------
-- Table structure for hk_user_address
-- ----------------------------
DROP TABLE IF EXISTS `hk_user_address`;
CREATE TABLE `hk_user_address` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户地址id',
  `uid` bigint unsigned NOT NULL COMMENT '用户id',
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
  KEY `uid` (`uid`) USING BTREE,
  KEY `is_default` (`is_default`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户地址表';

-- ----------------------------
-- Records of hk_user_address
-- ----------------------------
BEGIN;
INSERT INTO `hk_user_address` VALUES (24, 40, '陶子', '13666666666', '北京市', '北京市', 2, '朝阳区', '哈哈哈家', '', '0', '0', 0, 0, '2020-09-12 14:54:31', NULL);
INSERT INTO `hk_user_address` VALUES (25, 43, '张先生', '18888888888', '天津市', '天津市', 7363, '和平区', '6666', '', '0', '0', 0, 0, '2020-09-12 17:04:50', '2020-09-13 02:29:33');
INSERT INTO `hk_user_address` VALUES (26, 53, 'hahsdhsd', '15601064107', '北京市', '北京市', 2, '东城区', 'wqwewe', '', '0', '0', 0, 0, '2020-09-13 02:29:33', '2020-09-13 11:01:30');
INSERT INTO `hk_user_address` VALUES (27, 44, '胡先生', '18888888888', '北京市', '北京市', 2, '东城区', '999', '', '0', '0', 0, 0, '2020-09-13 10:17:00', '2020-09-13 11:01:30');
INSERT INTO `hk_user_address` VALUES (28, 53, '12', '15601064107', '山西省', '长治市', 72975, '长治县', '21231', '', '0', '0', 0, 0, '2020-09-13 11:01:30', '2020-09-13 21:33:36');
INSERT INTO `hk_user_address` VALUES (29, 42, '阿萨德', '18888888888', '天津市', '天津市', 7363, '河东区', '阿萨打算打', '', '0', '0', 0, 0, '2020-09-13 21:33:36', '2020-09-13 22:04:49');
INSERT INTO `hk_user_address` VALUES (30, 50, 'Hhhh', '18888888888', '河北省', '邯郸市', 26779, '复兴区', 'Hhhh', '', '0', '0', 0, 0, '2020-09-13 22:04:49', '2020-09-13 22:05:05');
INSERT INTO `hk_user_address` VALUES (31, 54, '胡新生', '15136175234', '北京市', '北京市', 2, '东城区', '999', '', '0', '0', 0, 0, '2020-09-13 22:05:05', '2020-09-14 00:00:20');
INSERT INTO `hk_user_address` VALUES (32, 49, 'Hhhh', '18855555555', '天津市', '天津市', 7363, '河东区', 'Yyyy', '', '0', '0', 0, 0, '2020-09-14 00:00:20', '2020-09-14 00:03:47');
INSERT INTO `hk_user_address` VALUES (33, 59, '胡先生', '15136175245', '天津市', '天津市', 7363, '和平区', '666', '', '0', '0', 0, 0, '2020-09-14 00:03:47', '2020-09-14 10:02:22');
INSERT INTO `hk_user_address` VALUES (34, 61, '胡先生', '15136175246', '北京市', '北京市', 2, '东城区', '666', '', '0', '0', 0, 0, '2020-09-14 10:02:22', '2020-09-14 12:53:03');
INSERT INTO `hk_user_address` VALUES (35, 71, '666', '18888888888', '天津市', '天津市', 7363, '河东区', '摸摸', '', '0', '0', 1, 0, '2020-09-14 12:53:03', NULL);
COMMIT;


-- ----------------------------
-- Table structure for hk_shop_order
-- ----------------------------
DROP TABLE IF EXISTS `hk_shop_order`;
CREATE TABLE `hk_shop_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `order_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单号',
  `extend_order_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '额外订单号',
  `uid` bigint unsigned NOT NULL COMMENT '用户id',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户姓名',
  `user_phone` varchar(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户电话',
  `user_address` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '详细地址',
  `cart_id` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '[]' COMMENT '购物车id',
  `freight_price` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '运费金额',
  `total_num` int unsigned NOT NULL DEFAULT '0' COMMENT '订单商品总数',
  `total_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '订单总价',
  `total_postage` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '邮费',
  `pay_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '实际支付金额',
  `pay_postage` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '支付邮费',
  `deduction_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '抵扣金额',
  `coupon_id` int unsigned NOT NULL DEFAULT '0' COMMENT '优惠券id',
  `coupon_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '优惠券金额',
  `paid` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '支付状态',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `pay_type` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付方式',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：已完成；-1：已退款）',
  `refund_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0 未退款 1 申请中 2 已退款',
  `refund_reason_wap_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '退款图片',
  `refund_reason_wap_explain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '退款用户说明',
  `refund_reason_time` datetime DEFAULT NULL COMMENT '退款时间',
  `refund_reason_wap` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '前台退款原因',
  `refund_reason` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '不退款的理由',
  `refund_price` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '退款金额',
  `delivery_sn` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '快递公司编号',
  `delivery_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '快递名称/送货人姓名',
  `delivery_type` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '发货类型',
  `delivery_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '快递单号/手机号',
  `gain_integral` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '消费赚取积分',
  `use_integral` decimal(8,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '使用积分',
  `back_integral` decimal(8,2) unsigned DEFAULT NULL COMMENT '给用户退了多少积分',
  `mark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '备注',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `unique` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一id(md5加密)类似id',
  `remark` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '管理员备注',
  `combination_id` bigint unsigned DEFAULT '0' COMMENT '拼团产品id0一般产品',
  `pink_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '拼团id 0没有拼团',
  `cost` decimal(8,2) unsigned NOT NULL COMMENT '成本价',
  `seckill_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '秒杀产品ID',
  `bargain_id` int unsigned DEFAULT '0' COMMENT '砍价id',
  `verify_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '核销码',
  `store_id` int NOT NULL DEFAULT '0' COMMENT '门店id',
  `shipping_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '配送方式 1=快递 ，2=门店自提',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_id_2` (`order_id`,`uid`) USING BTREE,
  UNIQUE KEY `unique` (`unique`) USING BTREE,
  KEY `uid` (`uid`) USING BTREE,
  KEY `add_time` (`create_time`) USING BTREE,
  KEY `pay_price` (`pay_price`) USING BTREE,
  KEY `paid` (`paid`) USING BTREE,
  KEY `pay_time` (`pay_time`) USING BTREE,
  KEY `pay_type` (`pay_type`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `is_del` (`is_del`) USING BTREE,
  KEY `coupon_id` (`coupon_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='订单表';

-- ----------------------------
-- Records of hk_shop_order
-- ----------------------------
BEGIN;
INSERT INTO `hk_shop_order` VALUES (1, '1302228251506114560', NULL, 29, '哈哈', '18888888888', '北京市 北京市 朝阳区 老两口儿', '2', 0.00, 1, 10.00, 0.00, 10.00, 0.00, 0.00, 0, 0.00, 1, '2020-09-05 20:55:40', 'yue', '2020-09-05 20:53:04', '2020-09-05 20:55:40', 0, 0, NULL, NULL, NULL, NULL, NULL, 0.00, '', NULL, NULL, NULL, 0.00, 0.00, NULL, '', 0, '8b53cc02e2ed45dda4356979fc7f2f87', NULL, 0, 0, 11.00, 0, 0, '', 0, 1);
COMMIT;