/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : gva

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 06/02/2023 18:20:14
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
) ENGINE = InnoDB AUTO_INCREMENT = 169 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (3, 'p', '888', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', '888', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', '888', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', '888', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', '888', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', '888', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', '888', '/authority/copyAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (12, 'p', '888', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (89, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (88, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (87, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (62, 'p', '888', '/autoCode/createPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (65, 'p', '888', '/autoCode/createPlug', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (59, 'p', '888', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (64, 'p', '888', '/autoCode/delPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (60, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (57, 'p', '888', '/autoCode/getColumn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (53, 'p', '888', '/autoCode/getDB', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (54, 'p', '888', '/autoCode/getMeta', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (63, 'p', '888', '/autoCode/getPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (61, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (56, 'p', '888', '/autoCode/getTables', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (66, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', '888', '/autoCode/preview', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (58, 'p', '888', '/autoCode/rollback', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1, 'p', '888', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (43, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (51, 'p', '888', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (48, 'p', '888', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (50, 'p', '888', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (49, 'p', '888', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (52, 'p', '888', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (83, 'p', '888', '/email/emailTest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (36, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (35, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (34, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (37, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', '888', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'p', '888', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (85, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (86, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (84, 'p', '888', '/simpleUploader/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (75, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (76, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (72, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (74, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (73, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (69, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (71, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (67, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (70, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (68, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (79, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (81, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (82, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (77, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (80, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', '888', '/system/getServerInfo', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (45, 'p', '888', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', '888', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'p', '888', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', '888', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (29, 'p', '888', '/user/deleteUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'p', '888', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'p', '888', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'p', '888', '/user/resetPassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', '888', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', '888', '/user/setUserInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (92, 'p', '8881', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', '8881', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (97, 'p', '8881', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (94, 'p', '8881', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (93, 'p', '8881', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (96, 'p', '8881', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (98, 'p', '8881', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (99, 'p', '8881', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (100, 'p', '8881', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (101, 'p', '8881', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (90, 'p', '8881', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (119, 'p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', '8881', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', '8881', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', '8881', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', '8881', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', '8881', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (127, 'p', '8881', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', '8881', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (102, 'p', '8881', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (103, 'p', '8881', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', '8881', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', '8881', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (91, 'p', '8881', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', '8881', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (128, 'p', '8881', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', '8881', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', '8881', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (131, 'p', '9528', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (134, 'p', '9528', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (136, 'p', '9528', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (133, 'p', '9528', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (132, 'p', '9528', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (135, 'p', '9528', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (137, 'p', '9528', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (138, 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (139, 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (140, 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (167, 'p', '9528', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (129, 'p', '9528', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (158, 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (157, 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (165, 'p', '9528', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (163, 'p', '9528', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (164, 'p', '9528', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (162, 'p', '9528', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (166, 'p', '9528', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (155, 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (156, 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (154, 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (153, 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (159, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (143, 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (145, 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (147, 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (144, 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'p', '9528', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (146, 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'p', '9528', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (160, 'p', '9528', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (161, 'p', '9528', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (130, 'p', '9528', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (150, 'p', '9528', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (168, 'p', '9528', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (151, 'p', '9528', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (152, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1101, 'p', '888', '/app/circle/getCircleForumPostsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1102, 'p', '888', '/app/circle/deleteCircleForumPosts', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1103, 'p', '888', '/app/circle/getUserCircleForumPostsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1104, 'p', '888', '/app/circle/getSelfCircleList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1105, 'p', '888', '/app/circle/getSelfCircleRequestList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1106, 'p', '888', '/app/circle/selfCircleTop', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1107, 'p', '888', '/app/circle/selfCircleCancelTop', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1108, 'p', '888', '/app/circle/findCircle', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1109, 'p', '888', '/app/circle/getCircleList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1110, 'p', '888', '/app/circle/getChildCircleList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1111, 'p', '888', '/app/circle/updateCircle', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1112, 'p', '888', '/app/circle/deleteCircleUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1113, 'p', '888', '/app/circle/deleteCircleUsers', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1114, 'p', '888', '/app/circle/updateCircleUser', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1115, 'p', '888', '/app/circle/setCircleUserPower', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1116, 'p', '888', '/app/circle/findCircleUser', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1117, 'p', '888', '/app/circle/getCircleUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1118, 'p', '888', '/app/circle/createCircleRequest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1119, 'p', '888', '/app/circle/destroyCircle', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1120, 'p', '888', '/app/circle/findCircleRequest', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1121, 'p', '888', '/app/circle/getCircleRequestList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1122, 'p', '888', '/app/circle/getCircleClassifyList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1123, 'p', '888', '/app/circle/getCircleClassifyListAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1124, 'p', '888', '/app/circle/enterCircle', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1125, 'p', '888', '/app/circle/applyEnterCircle', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1126, 'p', '888', '/app/circle/enterCircleApplyList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1127, 'p', '888', '/app/circle/approveEnterCircleRequest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1128, 'p', '888', '/app/circle/exitCircle', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1129, 'p', '888', '/app/circle/setCircleChannel', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1130, 'p', '888', '/app/circle/addCircleChannel', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1131, 'p', '888', '/app/circle/deleteCircleChannel', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1132, 'p', '888', '/app/circle/getCircleChannelList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1133, 'p', '888', '/app/circle/getCircleChannelListAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1134, 'p', '888', '/app/circle/setCircleTagSort', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1135, 'p', '888', '/app/circle/addCircleTag', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1136, 'p', '888', '/app/circle/deleteCircleTags', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1137, 'p', '888', '/app/circle/getCircleTagList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1138, 'p', '888', '/app/circle/requestBecomeChildCircle', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1139, 'p', '888', '/app/circle/approveChildCircleRequest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1140, 'p', '888', '/app/circle/createChildCircle', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1141, 'p', '888', '/app/circle/deleteChildCircle', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1142, 'p', '888', '/app/circle/getChildCircleRequestList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1143, 'p', '888', '/app/circle/findCircleBaseInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1144, 'p', '888', '/app/circle/findCirclePower', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1145, 'p', '888', '/app/circle/updateCircleBaseInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1146, 'p', '888', '/app/circle/updateCirclePower', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1147, 'p', '888', '/app/circle/createNews', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1148, 'p', '888', '/app/circle/getNewsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1149, 'p', '888', '/app/circle/getNewsDraftList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1150, 'p', '888', '/app/circle/updateNewsDraft', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1151, 'p', '888', '/app/circle/deleteNewsDraft', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1152, 'p', '888', '/app/circle/deleteNewsDraftByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1153, 'p', '888', '/app/circle/deleteNews', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1154, 'p', '888', '/app/circle/deleteNewsByIds', 'DELETE', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1201, 'p', '888', '/app/circleApply/findApply', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1202, 'p', '888', '/app/circleApply/getApplyList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1203, 'p', '888', '/app/circleApply/getApplyListAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1204, 'p', '888', '/app/circleApply/findCircleApply', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1205, 'p', '888', '/app/circleApply/getCircleApplyList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1206, 'p', '888', '/app/circleApply/getCircleApplyListAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1207, 'p', '888', '/app/circleApply/getCircleApplyGroupList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1208, 'p', '888', '/app/circleApply/getCircleApplyGroupListAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1209, 'p', '888', '/app/circleApply/getCircleHotApplyList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1210, 'p', '888', '/app/circleApply/setCircleHotApply', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1211, 'p', '888', '/app/circleApply/addCircleApply', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1212, 'p', '888', '/app/circleApply/updateCircleApply', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1213, 'p', '888', '/app/circleApply/deleteCircleApply', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1214, 'p', '888', '/app/circleApply/addCircleApplyGroup', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1215, 'p', '888', '/app/circleApply/deleteCircleApplyGroup', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1216, 'p', '888', '/app/circleApply/setCircleApplyGroupSort', 'POST', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1301, 'p', '888', '/app/forumPosts/createForumPosts', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1302, 'p', '888', '/app/forumPosts/createNews', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1303, 'p', '888', '/app/forumPosts/deleteForumPosts', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1304, 'p', '888', '/app/forumPosts/deleteForumPostsByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1305, 'p', '888', '/app/forumPosts/deleteSelfForumPosts', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1306, 'p', '888', '/app/forumPosts/updateForumPosts', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1307, 'p', '888', '/app/forumPosts/createForumComment', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1308, 'p', '888', '/app/forumPosts/deleteForumComment', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1309, 'p', '888', '/app/forumPosts/deleteForumCommentByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1310, 'p', '888', '/app/forumPosts/createForumThumbsUp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1311, 'p', '888', '/app/forumPosts/deleteForumThumbsUp', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1312, 'p', '888', '/app/forumPosts/createCommentThumbsUp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1313, 'p', '888', '/app/forumPosts/deleteCommentThumbsUp', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1314, 'p', '888', '/app/forumPosts/getRecommendPostsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1315, 'p', '888', '/app/forumPosts/getGlobalTopInfoList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1316, 'p', '888', '/app/forumPosts/getGlobalRecommendInfoList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1317, 'p', '888', '/app/forumPosts/getNearbyRecommendPostsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1318, 'p', '888', '/app/forumPosts/findForumPosts', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1319, 'p', '888', '/app/forumPosts/getForumPostsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1320, 'p', '888', '/app/forumPosts/findForumComment', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1321, 'p', '888', '/app/forumPosts/getForumCommentList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1322, 'p', '888', '/app/forumPosts/getUserForumPostsList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1401, 'p', '888', '/app/general/findProtocol', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1402, 'p', '888', '/app/general/findProtocolByName', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1403, 'p', '888', '/app/general/findMiniProgram', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14011, 'p', '888', '/app/general/findDraft', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14012, 'p', '888', '/app/general/getDraftList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14013, 'p', '888', '/app/general/deleteAllDraft', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14014, 'p', '888', '/app/general/deleteDraft', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14015, 'p', '888', '/app/general/deleteDraftByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14016, 'p', '888', '/app/general/updateDraft', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14017, 'p', '888', '/app/general/getUserCoverImageList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (14018, 'p', '888', '/app/general/getUserHeaderImageList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1500, 'p', '888', '/app/report/createReport', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1501, 'p', '888', '/app/report/findReport', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1502, 'p', '888', '/app/report/getReportList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1503, 'p', '888', '/app/report/getReportReasonList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1600, 'p', '888', '/app/topic/findForumTopicGroup', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1601, 'p', '888', '/app/topic/getForumTopicGroupList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1602, 'p', '888', '/app/topic/getForumTopicGroupListAll', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1603, 'p', '888', '/app/topic/createForumTopic', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1604, 'p', '888', '/app/topic/deleteForumTopic', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1605, 'p', '888', '/app/topic/deleteForumTopicByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1606, 'p', '888', '/app/topic/updateForumTopic', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1607, 'p', '888', '/app/topic/findForumTopic', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1608, 'p', '888', '/app/topic/getForumTopicList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1609, 'p', '888', '/app/topic/getNearbyHotTopicList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1610, 'p', '888', '/app/topic/getTopicForumPostsList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1701, 'p', '888', '/app/user/bindTelephone', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1702, 'p', '888', '/app/user/bindEmail', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1703, 'p', '888', '/app/user/getUserBaseInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1704, 'p', '888', '/app/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1705, 'p', '888', '/app/user/setSelfBaseInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1706, 'p', '888', '/app/user/getUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1707, 'p', '888', '/app/user/setUserCoverImage', 'POST', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1800, 'p', '888', '/app/userBrowsingHistory/deleteUserBrowsingHistory', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1801, 'p', '888', '/app/userBrowsingHistory/deleteUserBrowsingHistoryByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1802, 'p', '888', '/app/userBrowsingHistory/deleteAllUserBrowsingHistory', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1803, 'p', '888', '/app/userBrowsingHistory/getUserBrowsingHistoryList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (1900, 'p', '888', '/app/userCollect/createUserCollect', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1901, 'p', '888', '/app/userCollect/deleteUserCollect', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1902, 'p', '888', '/app/userCollect/deleteUserCollectByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1903, 'p', '888', '/app/userCollect/deleteAllUserCollect', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (1904, 'p', '888', '/app/userCollect/getUserCollectList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11001, 'p', '888', '/app/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11002, 'p', '888', '/app/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11003, 'p', '888', '/app/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11004, 'p', '888', '/app/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11005, 'p', '888', '/app/fileUploadAndDownload/findFile', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11006, 'p', '888', '/app/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11007, 'p', '888', '/app/fileUploadAndDownload/removeChunk', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11008, 'p', '888', '/app/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11101, 'p', '888', '/app/auth/getThirdPlat', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11102, 'p', '888', '/app/auth/register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11103, 'p', '888', '/app/auth/loginPwd', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11104, 'p', '888', '/app/auth/loginTelephone', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11105, 'p', '888', '/app/auth/loginThird', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11106, 'p', '888', '/app/auth/loginOneClick', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11107, 'p', '888', '/app/auth/getSmsVerification', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11108, 'p', '888', '/app/auth/resetPassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11109, 'p', '888', '/app/auth/resetPasswordCheckCode', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11110, 'p', '888', '/app/auth/getSmsVerificationPrivate', 'POST', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11201, 'p', '888', '/app/activity/createActivity', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11202, 'p', '888', '/app/activity/deleteActivity', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11203, 'p', '888', '/app/activity/deleteActivityByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11204, 'p', '888', '/app/activity/updateActivity', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11205, 'p', '888', '/app/activity/joinActivity', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11206, 'p', '888', '/app/activity/exitActivity', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11207, 'p', '888', '/app/activity/kickOutActivityUsers', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11208, 'p', '888', '/app/activity/getCircleRecommendActivityList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11209, 'p', '888', '/app/activity/getGlobalRecommendActivityList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11210, 'p', '888', '/app/activity/findActivity', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11211, 'p', '888', '/app/activity/findActivityUser', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11212, 'p', '888', '/app/activity/getActivityUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11213, 'p', '888', '/app/activity/createActivityDynamic', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11214, 'p', '888', '/app/activity/deleteActivityDynamic', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11215, 'p', '888', '/app/activity/deleteActivityDynamicByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11216, 'p', '888', '/app/activity/getActivityDynamicList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11217, 'p', '888', '/app/activity/deleteActivityAddRequest', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11218, 'p', '888', '/app/activity/deleteActivityAddRequestByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11219, 'p', '888', '/app/activity/updateActivityAddRequest', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11220, 'p', '888', '/app/activity/findActivityAddRequest', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11221, 'p', '888', '/app/activity/getActivityAddRequestList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11301, 'p', '888', '/app/question/createQuestion', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11302, 'p', '888', '/app/question/closeQuestion', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11303, 'p', '888', '/app/question/answerQuestion', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11304, 'p', '888', '/app/question/delSelfAnswer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11305, 'p', '888', '/app/question/setBestAnswer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11306, 'p', '888', '/app/question/setQuestionScore', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11307, 'p', '888', '/app/question/getGlobalRecommendQuestionList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11308, 'p', '888', '/app/question/getCircleQuestionList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11309, 'p', '888', '/app/question/getAnswerList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11401, 'p', '888', '/app/focus/focusUser', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11402, 'p', '888', '/app/focus/focusUserCancel', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11403, 'p', '888', '/app/focus/updateFocusUser', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11404, 'p', '888', '/app/focus/getFrequentBrowsingUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11405, 'p', '888', '/app/focus/getFocusUserPostsList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11406, 'p', '888', '/app/focus/getFocusUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11407, 'p', '888', '/app/focus/getFansList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11408, 'p', '888', '/app/focus/findFocusUser', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11501, 'p', '888', '/app/channel/getChannelList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11502, 'p', '888', '/app/channel/getUserChannelList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11503, 'p', '888', '/app/channel/setUserChannel', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11504, 'p', '888', '/app/channel/getGeneralChannelList', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11601, 'p', '888', '/app/wallet/getProductGoldList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11602, 'p', '888', '/app/wallet/getPayTypeList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11603, 'p', '888', '/app/wallet/createOrder', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11604, 'p', '888', '/app/wallet/getGoldBillList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11605, 'p', '888', '/app/wallet/findGoldBill', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11606, 'p', '888', '/app/wallet/getExtractTypeList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11607, 'p', '888', '/app/wallet/createUserExtract', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11608, 'p', '888', '/app/wallet/getUserBillList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11609, 'p', '888', '/app/wallet/findUserBill', 'GET', '', '', '', '', '');

INSERT INTO `casbin_rule` VALUES (11701, 'p', '888', '/app/helpCenter/getHelpTypeList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11702, 'p', '888', '/app/helpCenter/getMainHelpList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11703, 'p', '888', '/app/helpCenter/findHelp', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11704, 'p', '888', '/app/helpCenter/helpThumbsUp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11706, 'p', '888', '/app/helpCenter/getHelpList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11707, 'p', '888', '/app/helpCenter/createFeedback', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11708, 'p', '888', '/app/helpCenter/deleteFeedback', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11709, 'p', '888', '/app/helpCenter/deleteFeedbackByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11710, 'p', '888', '/app/helpCenter/updateFeedback', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11711, 'p', '888', '/app/helpCenter/findFeedback', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (11712, 'p', '888', '/app/helpCenter/getFeedbackList', 'GET', '', '', '', '', '');

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `exa_file_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint(20) NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '000000', '2023-02-06 18:19:13.643', '2023-02-06 18:19:13.643', NULL, 0, 0, NULL, NULL, NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '000000', '2023-02-06 18:19:13.643', '2023-02-06 18:19:13.643', NULL, 0, 0, NULL, NULL, NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `chunk_total` bigint(20) NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_files
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis_group
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis_group`;
CREATE TABLE `sys_apis_group`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分组名称',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 90 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统接口分组';

INSERT INTO `sys_apis_group` VALUES (1, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_圈子', '', 1);
INSERT INTO `sys_apis_group` VALUES (2, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_用户圈子应用', '', 2);
INSERT INTO `sys_apis_group` VALUES (3, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_帖子', '', 3);
INSERT INTO `sys_apis_group` VALUES (4, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_常规', '', 4);
INSERT INTO `sys_apis_group` VALUES (5, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_举报', '', 5);
INSERT INTO `sys_apis_group` VALUES (6, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_话题', '', 6);
INSERT INTO `sys_apis_group` VALUES (7, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_用户', '', 7);
INSERT INTO `sys_apis_group` VALUES (8, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_浏览历史', '', 8);
INSERT INTO `sys_apis_group` VALUES (9, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_收藏', '', 9);
INSERT INTO `sys_apis_group` VALUES (10, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_文件上传与下载', '', 10);
INSERT INTO `sys_apis_group` VALUES (11, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_鉴权', '', 11);
INSERT INTO `sys_apis_group` VALUES (12, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_活动', '', 12);
INSERT INTO `sys_apis_group` VALUES (13, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_问答', '', 13);
INSERT INTO `sys_apis_group` VALUES (14, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_关注', '', 14);
INSERT INTO `sys_apis_group` VALUES (15, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_频道', '', 15);
INSERT INTO `sys_apis_group` VALUES (16, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_钱包', '', 16);
INSERT INTO `sys_apis_group` VALUES (17, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'APP权限_反馈中心', '', 17);

INSERT INTO `sys_apis_group` VALUES (51, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'base', '', 51);
INSERT INTO `sys_apis_group` VALUES (52, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'jwt', '', 52);
INSERT INTO `sys_apis_group` VALUES (53, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '系统用户', '', 53);
INSERT INTO `sys_apis_group` VALUES (54, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'api', '', 54);
INSERT INTO `sys_apis_group` VALUES (55, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '角色', '', 55);
INSERT INTO `sys_apis_group` VALUES (56, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'casbin', '', 56);
INSERT INTO `sys_apis_group` VALUES (57, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '菜单', '', 57);
INSERT INTO `sys_apis_group` VALUES (58, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '分片上传', '', 58);
INSERT INTO `sys_apis_group` VALUES (59, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '文件上传与下载', '', 59);
INSERT INTO `sys_apis_group` VALUES (60, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '系统服务', '', 60);
INSERT INTO `sys_apis_group` VALUES (61, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '客户', '', 61);
INSERT INTO `sys_apis_group` VALUES (62, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '代码生成器', '', 62);
INSERT INTO `sys_apis_group` VALUES (63, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '包（pkg）生成器', '', 63);
INSERT INTO `sys_apis_group` VALUES (64, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '代码生成器历史', '', 64);
INSERT INTO `sys_apis_group` VALUES (65, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '系统字典详情', '', 65);
INSERT INTO `sys_apis_group` VALUES (66, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '系统字典', '', 66);
INSERT INTO `sys_apis_group` VALUES (67, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '操作记录', '', 67);
INSERT INTO `sys_apis_group` VALUES (68, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '断点续传(插件版)', '', 68);
INSERT INTO `sys_apis_group` VALUES (69, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, 'email', '', 69);
INSERT INTO `sys_apis_group` VALUES (70, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '按钮权限', '', 70);

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 90 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统接口';

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/base/login', '用户登录(必选)', 51, 'POST');
INSERT INTO `sys_apis` VALUES (2, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 52, 'POST');
INSERT INTO `sys_apis` VALUES (3, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/deleteUser', '删除用户', 53, 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/admin_register', '用户注册', 53, 'POST');
INSERT INTO `sys_apis` VALUES (5, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/getUserList', '获取用户列表', 53, 'POST');
INSERT INTO `sys_apis` VALUES (6, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/setUserInfo', '设置用户信息', 53, 'PUT');
INSERT INTO `sys_apis` VALUES (7, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/setSelfInfo', '设置自身信息(必选)', 53, 'PUT');
INSERT INTO `sys_apis` VALUES (8, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/getUserInfo', '获取自身信息(必选)', 53, 'GET');
INSERT INTO `sys_apis` VALUES (9, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/setUserAuthorities', '设置权限组', 53, 'POST');
INSERT INTO `sys_apis` VALUES (10, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/changePassword', '修改密码（建议选择)', 53, 'POST');
INSERT INTO `sys_apis` VALUES (11, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/setUserAuthority', '修改用户角色(必选)', 53, 'POST');
INSERT INTO `sys_apis` VALUES (12, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/user/resetPassword', '重置用户密码', 53, 'POST');
INSERT INTO `sys_apis` VALUES (13, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/createApi', '创建api', 54, 'POST');
INSERT INTO `sys_apis` VALUES (14, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/deleteApi', '删除Api', 54, 'POST');
INSERT INTO `sys_apis` VALUES (15, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/updateApi', '更新Api', 54, 'POST');
INSERT INTO `sys_apis` VALUES (16, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/getApiList', '获取api列表', 54, 'POST');
INSERT INTO `sys_apis` VALUES (17, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/getAllApis', '获取所有api', 54, 'POST');
INSERT INTO `sys_apis` VALUES (18, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/getApiById', '获取api详细信息', 54, 'POST');
INSERT INTO `sys_apis` VALUES (19, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/api/deleteApisByIds', '批量删除api', 54, 'DELETE');
INSERT INTO `sys_apis` VALUES (20, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authority/copyAuthority', '拷贝角色', 55, 'POST');
INSERT INTO `sys_apis` VALUES (21, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authority/createAuthority', '创建角色', 55, 'POST');
INSERT INTO `sys_apis` VALUES (22, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authority/deleteAuthority', '删除角色', 55, 'POST');
INSERT INTO `sys_apis` VALUES (23, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authority/updateAuthority', '更新角色信息', 55, 'PUT');
INSERT INTO `sys_apis` VALUES (24, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authority/getAuthorityList', '获取角色列表', 55, 'POST');
INSERT INTO `sys_apis` VALUES (25, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authority/setDataAuthority', '设置角色资源权限', 55, 'POST');
INSERT INTO `sys_apis` VALUES (26, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/casbin/updateCasbin', '更改角色api权限', 56, 'POST');
INSERT INTO `sys_apis` VALUES (27, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 56, 'POST');
INSERT INTO `sys_apis` VALUES (28, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/addBaseMenu', '新增菜单', 57, 'POST');
INSERT INTO `sys_apis` VALUES (29, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/getMenu', '获取菜单树(必选)', 57, 'POST');
INSERT INTO `sys_apis` VALUES (30, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/deleteBaseMenu', '删除菜单', 57, 'POST');
INSERT INTO `sys_apis` VALUES (31, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/updateBaseMenu', '更新菜单', 57, 'POST');
INSERT INTO `sys_apis` VALUES (32, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/getBaseMenuById', '根据id获取菜单', 57, 'POST');
INSERT INTO `sys_apis` VALUES (33, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/getMenuList', '分页获取基础menu列表', 57, 'POST');
INSERT INTO `sys_apis` VALUES (34, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 57, 'POST');
INSERT INTO `sys_apis` VALUES (35, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/getMenuAuthority', '获取指定角色menu', 57, 'POST');
INSERT INTO `sys_apis` VALUES (36, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 57, 'POST');
INSERT INTO `sys_apis` VALUES (37, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', 58, 'GET');
INSERT INTO `sys_apis` VALUES (38, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', 58, 'POST');
INSERT INTO `sys_apis` VALUES (39, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', 58, 'POST');
INSERT INTO `sys_apis` VALUES (40, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', 58, 'POST');
INSERT INTO `sys_apis` VALUES (41, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/upload', '文件上传示例', 59, 'POST');
INSERT INTO `sys_apis` VALUES (42, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/deleteFile', '删除文件', 59, 'POST');
INSERT INTO `sys_apis` VALUES (43, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', 59, 'POST');
INSERT INTO `sys_apis` VALUES (44, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 59, 'POST');
INSERT INTO `sys_apis` VALUES (45, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/system/getServerInfo', '获取服务器信息', 60, 'POST');
INSERT INTO `sys_apis` VALUES (46, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/system/getSystemConfig', '获取配置文件内容', 60, 'POST');
INSERT INTO `sys_apis` VALUES (47, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/system/setSystemConfig', '设置配置文件内容', 60, 'POST');
INSERT INTO `sys_apis` VALUES (48, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/customer/customer', '更新客户', 61, 'PUT');
INSERT INTO `sys_apis` VALUES (49, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/customer/customer', '创建客户', 61, 'POST');
INSERT INTO `sys_apis` VALUES (50, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/customer/customer', '删除客户', 61, 'DELETE');
INSERT INTO `sys_apis` VALUES (51, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/customer/customer', '获取单一客户', 61, 'GET');
INSERT INTO `sys_apis` VALUES (52, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/customer/customerList', '获取客户列表', 61, 'GET');
INSERT INTO `sys_apis` VALUES (53, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/getDB', '获取所有数据库', 62, 'GET');
INSERT INTO `sys_apis` VALUES (54, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/getTables', '获取数据库表', 62, 'GET');
INSERT INTO `sys_apis` VALUES (55, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/createTemp', '自动化代码', 62, 'POST');
INSERT INTO `sys_apis` VALUES (56, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/preview', '预览自动化代码', 62, 'POST');
INSERT INTO `sys_apis` VALUES (57, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/getColumn', '获取所选table的所有字段', 62, 'GET');
INSERT INTO `sys_apis` VALUES (58, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/createPlug', '自动创建插件包', 62, 'POST');
INSERT INTO `sys_apis` VALUES (59, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/installPlugin', '安装插件', 62, 'POST');
INSERT INTO `sys_apis` VALUES (60, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/createPackage', '生成包(package)', 63, 'POST');
INSERT INTO `sys_apis` VALUES (61, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/getPackage', '获取所有包(package)', 63, 'POST');
INSERT INTO `sys_apis` VALUES (62, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/delPackage', '删除包(package)', 63, 'POST');
INSERT INTO `sys_apis` VALUES (63, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/getMeta', '获取meta信息', 64, 'POST');
INSERT INTO `sys_apis` VALUES (64, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/rollback', '回滚自动生成代码', 64, 'POST');
INSERT INTO `sys_apis` VALUES (65, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/getSysHistory', '查询回滚记录', 64, 'POST');
INSERT INTO `sys_apis` VALUES (66, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/autoCode/delSysHistory', '删除回滚记录', 64, 'POST');
INSERT INTO `sys_apis` VALUES (67, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', 65, 'PUT');
INSERT INTO `sys_apis` VALUES (68, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', 65, 'POST');
INSERT INTO `sys_apis` VALUES (69, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', 65, 'DELETE');
INSERT INTO `sys_apis` VALUES (70, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', 65, 'GET');
INSERT INTO `sys_apis` VALUES (71, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', 65, 'GET');
INSERT INTO `sys_apis` VALUES (72, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionary/createSysDictionary', '新增字典', 66, 'POST');
INSERT INTO `sys_apis` VALUES (73, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionary/deleteSysDictionary', '删除字典', 66, 'DELETE');
INSERT INTO `sys_apis` VALUES (74, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionary/updateSysDictionary', '更新字典', 66, 'PUT');
INSERT INTO `sys_apis` VALUES (75, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', 66, 'GET');
INSERT INTO `sys_apis` VALUES (76, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', 66, 'GET');
INSERT INTO `sys_apis` VALUES (77, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', 67, 'POST');
INSERT INTO `sys_apis` VALUES (78, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', 67, 'GET');
INSERT INTO `sys_apis` VALUES (79, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', 67, 'GET');
INSERT INTO `sys_apis` VALUES (80, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', 67, 'DELETE');
INSERT INTO `sys_apis` VALUES (81, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', 67, 'DELETE');
INSERT INTO `sys_apis` VALUES (82, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/simpleUploader/upload', '插件版分片上传', 68, 'POST');
INSERT INTO `sys_apis` VALUES (83, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', 68, 'GET');
INSERT INTO `sys_apis` VALUES (84, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', 68, 'GET');
INSERT INTO `sys_apis` VALUES (85, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/email/emailTest', '发送测试邮件', 69, 'POST');
INSERT INTO `sys_apis` VALUES (86, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/email/emailSend', '发送邮件示例', 69, 'POST');
INSERT INTO `sys_apis` VALUES (87, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', 70, 'POST');
INSERT INTO `sys_apis` VALUES (88, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', 70, 'POST');
INSERT INTO `sys_apis` VALUES (89, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', 70, 'POST');

INSERT INTO `sys_apis` VALUES (101, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleForumPostsList', '分页获取圈子ForumPosts列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (102, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteCircleForumPosts', '(圈子管理员)删除圈子的帖子', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (103, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getUserCircleForumPostsList', '用户加入圈子的所有动态列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (104, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getSelfCircleList', '分页获取用户加入的Circle列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (105, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getSelfCircleRequestList', '分页获取用户创建圈子申请记录列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (106, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/selfCircleTop', '用户圈子置顶', 1, 'POST');
INSERT INTO `sys_apis` VALUES (107, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/selfCircleCancelTop', '用户圈子取消置顶', 1, 'POST');
INSERT INTO `sys_apis` VALUES (108, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/findCircle', '用id查询Circle', 1, 'GET');
INSERT INTO `sys_apis` VALUES (109, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleList', '分页获取Circle列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (110, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getChildCircleList', '分页获取子圈子列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (111, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/updateCircle', '(圈子管理者)更新Circle', 1, 'PUT');
INSERT INTO `sys_apis` VALUES (112, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteCircleUser', '(圈子管理者)删除圈子用户', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (113, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteCircleUsers', '(圈子管理者)批量删除圈子用户', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (114, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/updateCircleUser', '更新CircleUser', 1, 'PUT');
INSERT INTO `sys_apis` VALUES (115, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/setCircleUserPower', '(圈子管理者)设置圈子用户权限', 1, 'PUT');
INSERT INTO `sys_apis` VALUES (116, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/findCircleUser', '查询圈子用户信息', 1, 'GET');
INSERT INTO `sys_apis` VALUES (117, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleUserList', '分页获取CircleUser列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (118, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/createCircleRequest', '创建圈子请求', 1, 'POST');
INSERT INTO `sys_apis` VALUES (119, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/destroyCircle', '(圈子管理者)注销圈子', 1, 'POST');
INSERT INTO `sys_apis` VALUES (120, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/findCircleRequest', '用id查询CircleRequest', 1, 'GET');
INSERT INTO `sys_apis` VALUES (121, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleRequestList', '分页获取CircleRequest列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (122, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleClassifyList', '分页获取CircleClassify列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (123, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleClassifyListAll', '获取CircleClassify列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (124, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/enterCircle', '加入圈子', 1, 'POST');
INSERT INTO `sys_apis` VALUES (125, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/applyEnterCircle', '申请加入圈子', 1, 'POST');
INSERT INTO `sys_apis` VALUES (126, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/enterCircleApplyList', '分页获取加入圈子申请', 1, 'GET');
INSERT INTO `sys_apis` VALUES (127, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/approveEnterCircleRequest', '审批加入圈子申请', 1, 'POST');
INSERT INTO `sys_apis` VALUES (128, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/exitCircle', '退出圈子', 1, 'POST');
INSERT INTO `sys_apis` VALUES (129, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/setCircleChannel', '(圈子管理者)设置圈子频道', 1, 'POST');
INSERT INTO `sys_apis` VALUES (130, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/addCircleChannel', '(圈子管理者)添加圈子频道', 1, 'POST');
INSERT INTO `sys_apis` VALUES (131, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteCircleChannel', '(圈子管理者)删除圈子频道', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (132, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleChannelList', '获取圈子频道', 1, 'GET');
INSERT INTO `sys_apis` VALUES (133, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleChannelListAll', '(圈子管理员)获取圈子频道全部', 1, 'GET');
INSERT INTO `sys_apis` VALUES (134, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/setCircleTagSort', '(圈子管理员)设置圈子标签顺序', 1, 'POST');
INSERT INTO `sys_apis` VALUES (135, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/addCircleTag', '(圈子管理员)添加圈子标签', 1, 'POST');
INSERT INTO `sys_apis` VALUES (136, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteCircleTags', '(圈子管理员)删除圈子标签', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (137, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getCircleTagList', '(圈子管理员)获取圈子的标签', 1, 'GET');
INSERT INTO `sys_apis` VALUES (138, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/requestBecomeChildCircle', '(圈子管理者)申请成为子圈子', 1, 'POST');
INSERT INTO `sys_apis` VALUES (139, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/approveChildCircleRequest', '(圈子管理者)审批子圈子申请', 1, 'POST');
INSERT INTO `sys_apis` VALUES (140, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/createChildCircle', '(圈子管理者)创建子圈子', 1, 'POST');
INSERT INTO `sys_apis` VALUES (141, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteChildCircle', '(圈子管理者)删除子圈子', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (142, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getChildCircleRequestList', '(圈子管理者)获取子圈子申请列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (143, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/findCircleBaseInfo', '用id查询圈子基本信息', 1, 'GET');
INSERT INTO `sys_apis` VALUES (144, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/findCirclePower', '用id查询圈子权限', 1, 'GET');
INSERT INTO `sys_apis` VALUES (145, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/updateCircleBaseInfo', '(圈子管理者)更新圈子基本信息', 1, 'PUT');
INSERT INTO `sys_apis` VALUES (146, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/updateCirclePower', '(圈子管理者)更新圈子权限', 1, 'PUT');
INSERT INTO `sys_apis` VALUES (147, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/createNews', '(圈子管理者)创建资讯', 1, 'POST');
INSERT INTO `sys_apis` VALUES (148, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getNewsList', '(圈子管理者)获取资讯列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (149, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/getNewsDraftList', '(圈子管理者)获取资讯草稿列表', 1, 'GET');
INSERT INTO `sys_apis` VALUES (150, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/updateNewsDraft', '(圈子管理者)更新资讯草稿', 1, 'POST');
INSERT INTO `sys_apis` VALUES (151, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteNewsDraft', '(圈子管理者)删除资讯草稿', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (152, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteNewsDraftByIds', '(圈子管理者)批量删除资讯草稿', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (153, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteNews', '(圈子管理者)删除资讯', 1, 'DELETE');
INSERT INTO `sys_apis` VALUES (154, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circle/deleteNewsByIds', '(圈子管理者)批量删除资讯', 1, 'DELETE');

INSERT INTO `sys_apis` VALUES (201, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/findApply', '用id查询Apply', 2, 'GET');
INSERT INTO `sys_apis` VALUES (202, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getApplyList', '分页获取Apply列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (203, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getApplyListAll', '获取Apply列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (204, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/findCircleApply', '用id查询CircleApply', 2, 'GET');
INSERT INTO `sys_apis` VALUES (205, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getCircleApplyList', '分页获取CircleApply列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (206, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getCircleApplyListAll', '获取CircleApply列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (207, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getCircleApplyGroupList', '分页获取CircleApplyGroup列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (208, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getCircleApplyGroupListAll', '获取CircleApplyGroup列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (209, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/getCircleHotApplyList', '获取圈子热门应用列表', 2, 'GET');
INSERT INTO `sys_apis` VALUES (210, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/setCircleHotApply', '(圈子管理者)设置圈子热门应用', 2, 'POST');
INSERT INTO `sys_apis` VALUES (211, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/addCircleApply', '(圈子管理者)添加圈子应用', 2, 'POST');
INSERT INTO `sys_apis` VALUES (212, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/updateCircleApply', '(圈子管理者)更新圈子应用', 2, 'POST');
INSERT INTO `sys_apis` VALUES (213, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/deleteCircleApply', '(圈子管理者)删除圈子应用', 2, 'DELETE');
INSERT INTO `sys_apis` VALUES (214, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/addCircleApplyGroup', '(圈子管理者)添加圈子应用分组', 2, 'POST');
INSERT INTO `sys_apis` VALUES (215, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/deleteCircleApplyGroup', '(圈子管理者)删除圈子应用分组', 2, 'DELETE');
INSERT INTO `sys_apis` VALUES (216, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/circleApply/setCircleApplyGroupSort', '(圈子管理者)设置圈子应用分组顺序', 2, 'POST');

INSERT INTO `sys_apis` VALUES (301, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/createForumPosts', '创建帖子', 3, 'POST');
INSERT INTO `sys_apis` VALUES (302, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/createNews', '创建资讯', 3, 'POST');
INSERT INTO `sys_apis` VALUES (303, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteForumPosts', '删除帖子', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (304, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteForumPostsByIds', '批量删除帖子', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (305, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteSelfForumPosts', '删除自己的帖子', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (306, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/updateForumPosts', '更新帖子', 3, 'PUT');
INSERT INTO `sys_apis` VALUES (307, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/createForumComment', '创建评论', 3, 'POST');
INSERT INTO `sys_apis` VALUES (308, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteForumComment', '删除评论', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (309, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteForumCommentByIds', '批量删除评论', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (310, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/createForumThumbsUp', '帖子点赞', 3, 'POST');
INSERT INTO `sys_apis` VALUES (311, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteForumThumbsUp', '删除帖子点赞', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (312, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/createCommentThumbsUp', '评论点赞', 3, 'POST');
INSERT INTO `sys_apis` VALUES (313, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/deleteCommentThumbsUp', '删除评论点赞', 3, 'DELETE');
INSERT INTO `sys_apis` VALUES (314, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getRecommendPostsList', '分页获取推荐帖子列表', 3, 'GET');
INSERT INTO `sys_apis` VALUES (315, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getGlobalTopInfoList', '获全局置顶资讯列表', 3, 'GET');
INSERT INTO `sys_apis` VALUES (316, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getGlobalRecommendInfoList', '分页获全局推荐资讯列表', 3, 'GET');
INSERT INTO `sys_apis` VALUES (317, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getNearbyRecommendPostsList', '分页获附近推荐帖子列表', 3, 'GET');
INSERT INTO `sys_apis` VALUES (318, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/findForumPosts', '用id查询帖子', 3, 'GET');
INSERT INTO `sys_apis` VALUES (319, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getForumPostsList', '分页获取帖子列表', 3, 'GET');
INSERT INTO `sys_apis` VALUES (320, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/findForumComment', '用id查询评论', 3, 'GET');
INSERT INTO `sys_apis` VALUES (321, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getForumCommentList', '分页获取评论列表', 3, 'GET');
INSERT INTO `sys_apis` VALUES (322, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/forumPosts/getUserForumPostsList', '分页获取用户帖子列表', 7, 'GET');

INSERT INTO `sys_apis` VALUES (401, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/findProtocol', '用id查询协议', 4, 'GET');
INSERT INTO `sys_apis` VALUES (402, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/findProtocolByName', '用名字查询协议', 4, 'GET');
INSERT INTO `sys_apis` VALUES (403, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/findMiniProgram', '用id查询MiniProgram', 4, 'GET');
INSERT INTO `sys_apis` VALUES (411, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/findDraft', '用id查询草稿', 4, 'GET');
INSERT INTO `sys_apis` VALUES (412, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/getDraftList', '分页获取草稿列表', 4, 'GET');
INSERT INTO `sys_apis` VALUES (413, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/deleteAllDraft', '删除所有草稿', 4, 'DELETE');
INSERT INTO `sys_apis` VALUES (414, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/deleteDraft', '删除草稿', 4, 'DELETE');
INSERT INTO `sys_apis` VALUES (415, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/deleteDraftByIds', '批量删除草稿', 4, 'DELETE');
INSERT INTO `sys_apis` VALUES (416, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/updateDraft', '更新草稿', 4, 'POST');
INSERT INTO `sys_apis` VALUES (417, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/getUserCoverImageList', '获取用户主页封面列表', 4, 'GET');
INSERT INTO `sys_apis` VALUES (418, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/general/getUserHeaderImageList', '获取用户头像列表', 4, 'GET');

INSERT INTO `sys_apis` VALUES (500, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/report/createReport', '举报', 5, 'POST');
INSERT INTO `sys_apis` VALUES (501, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/report/findReport', '用id查询Report', 5, 'GET');
INSERT INTO `sys_apis` VALUES (502, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/report/getReportList', '分页获取Report列表', 5, 'GET');
INSERT INTO `sys_apis` VALUES (503, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/report/getReportReasonList', '分页获取ReportReason列表', 5, 'GET');

INSERT INTO `sys_apis` VALUES (600, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/findForumTopicGroup', '用id查询ForumTopicGroup', 6, 'GET');
INSERT INTO `sys_apis` VALUES (601, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/getForumTopicGroupList', '分页获取ForumTopicGroup列表', 6, 'GET');
INSERT INTO `sys_apis` VALUES (602, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/getForumTopicGroupListAll', '获取ForumTopicGroup列表', 6, 'GET');
INSERT INTO `sys_apis` VALUES (603, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/createForumTopic', '创建ForumTopic', 6, 'POST');
INSERT INTO `sys_apis` VALUES (604, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/deleteForumTopic', '删除ForumTopic', 6, 'DELETE');
INSERT INTO `sys_apis` VALUES (605, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/deleteForumTopicByIds', '批量删除ForumTopic', 6, 'DELETE');
INSERT INTO `sys_apis` VALUES (606, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/updateForumTopic', '更新ForumTopic', 6, 'PUT');
INSERT INTO `sys_apis` VALUES (607, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/findForumTopic', '用id查询ForumTopic', 6, 'GET');
INSERT INTO `sys_apis` VALUES (608, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/getForumTopicList', '分页获取ForumTopic列表', 6, 'GET');
INSERT INTO `sys_apis` VALUES (609, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/getNearbyHotTopicList', '获取附近热门话题列表', 6, 'GET');
INSERT INTO `sys_apis` VALUES (610, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/topic/getTopicForumPostsList', '分页获取话题帖子列表', 6, 'GET');

INSERT INTO `sys_apis` VALUES (701, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/bindTelephone', '绑定手机', 7, 'POST');
INSERT INTO `sys_apis` VALUES (702, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/bindEmail', '绑定邮箱', 7, 'POST');
INSERT INTO `sys_apis` VALUES (703, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/getUserBaseInfo', '用id查询用户基本信息', 7, 'GET');
INSERT INTO `sys_apis` VALUES (704, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/getUserInfo', '用id查询用户信息', 7, 'GET');
INSERT INTO `sys_apis` VALUES (705, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/setSelfBaseInfo', '设置用户基础信息', 7, 'PUT');
INSERT INTO `sys_apis` VALUES (706, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/getUserList', '分页获取User列表', 7, 'GET');
INSERT INTO `sys_apis` VALUES (707, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/user/setUserCoverImage', '设置用户主页封面', 7, 'POST');

INSERT INTO `sys_apis` VALUES (800, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userBrowsingHistory/deleteUserBrowsingHistory', '删除浏览历史', 8, 'DELETE');
INSERT INTO `sys_apis` VALUES (801, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userBrowsingHistory/deleteUserBrowsingHistoryByIds', '批量删除浏览历史', 8, 'DELETE');
INSERT INTO `sys_apis` VALUES (802, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userBrowsingHistory/deleteAllUserBrowsingHistory', '删除所有浏览历史', 8, 'DELETE');
INSERT INTO `sys_apis` VALUES (803, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userBrowsingHistory/getUserBrowsingHistoryList', '分页获取浏览历史列表', 8, 'GET');

INSERT INTO `sys_apis` VALUES (900, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userCollect/createUserCollect', '收藏帖子', 9, 'POST');
INSERT INTO `sys_apis` VALUES (901, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userCollect/deleteUserCollect', '删除收藏', 9, 'DELETE');
INSERT INTO `sys_apis` VALUES (902, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userCollect/deleteUserCollectByIds', '批量删除收藏', 9, 'DELETE');
INSERT INTO `sys_apis` VALUES (903, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userCollect/deleteAllUserCollect', '删除所有收藏', 9, 'DELETE');
INSERT INTO `sys_apis` VALUES (904, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/userCollect/getUserCollectList', '分页获取收藏列表', 9, 'GET');

INSERT INTO `sys_apis` VALUES (1001, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', 10, 'GET');
INSERT INTO `sys_apis` VALUES (1002, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/breakpointContinue', '断点续传', 10, 'POST');
INSERT INTO `sys_apis` VALUES (1003, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', 10, 'POST');
INSERT INTO `sys_apis` VALUES (1004, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/removeChunk', '上传完成移除文件', 10, 'POST');
INSERT INTO `sys_apis` VALUES (1005, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/upload', '文件上传示例', 10, 'POST');
INSERT INTO `sys_apis` VALUES (1006, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/deleteFile', '删除文件', 10, 'POST');
INSERT INTO `sys_apis` VALUES (1007, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/editFileName', '文件名或者备注编辑',10, 'POST');
INSERT INTO `sys_apis` VALUES (1008, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/fileUploadAndDownload/getFileList', '获取上传文件列表', 10, 'POST');

INSERT INTO `sys_apis` VALUES (1101, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/getThirdPlat', '获取第三方平台信息', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1102, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/register', '注册', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1103, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/loginPwd', '用户登录(账号密码)', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1104, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/loginTelephone', '用户登录(手机)', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1105, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/loginThird', '用户登录(第三方授权)', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1106, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/loginOneClick', '一键登录', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1107, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/getSmsVerification', '获取短信验证码', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1108, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/resetPassword', '重置密码', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1109, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/resetPasswordCheckCode', '重置密码检验验证码', 11, 'POST');
INSERT INTO `sys_apis` VALUES (1110, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/auth/getSmsVerificationPrivate', '获取短信验证码', 11, 'POST');

INSERT INTO `sys_apis` VALUES (1201, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/createActivity', '创建活动', 12, 'POST');
INSERT INTO `sys_apis` VALUES (1202, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/deleteActivity', '删除活动', 12, 'DELETE');
INSERT INTO `sys_apis` VALUES (1203, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/deleteActivityByIds', '批量删除活动', 12, 'DELETE');
INSERT INTO `sys_apis` VALUES (1204, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/updateActivity', '更新活动', 12, 'PUT');
INSERT INTO `sys_apis` VALUES (1205, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/joinActivity', '加入活动', 12, 'POST');
INSERT INTO `sys_apis` VALUES (1206, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/exitActivity', '退出活动', 12, 'POST');
INSERT INTO `sys_apis` VALUES (1207, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/kickOutActivityUsers', '踢出活动的用户', 12, 'DELETE');
INSERT INTO `sys_apis` VALUES (1208, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/getCircleRecommendActivityList', '分页获取圈子推荐活动列表', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1209, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/getGlobalRecommendActivityList', '分页获取全局推荐活动列表', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1210, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/findActivity', '用id查询活动详情', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1211, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/findActivityUser', '查询活动的用户', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1212, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/getActivityUserList', '分页获取活动的用户列表', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1213, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/createActivityDynamic', '创建活动动态', 12, 'POST');
INSERT INTO `sys_apis` VALUES (1214, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/deleteActivityDynamic', '删除活动动态', 12, 'DELETE');
INSERT INTO `sys_apis` VALUES (1215, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/deleteActivityDynamicByIds', '批量删除活动动态', 12, 'DELETE');

INSERT INTO `sys_apis` VALUES (1216, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/getActivityDynamicList', '分页获取活动的动态列表', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1217, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/deleteActivityAddRequest', '删除活动报名申请', 12, 'DELETE');
INSERT INTO `sys_apis` VALUES (1218, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/deleteActivityAddRequestByIds', '批量删除活动报名申请', 12, 'DELETE');
INSERT INTO `sys_apis` VALUES (1219, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/updateActivityAddRequest', '活动报名申请审批', 12, 'PUT');
INSERT INTO `sys_apis` VALUES (1220, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/findActivityAddRequest', '用id查询活动报名申请', 12, 'GET');
INSERT INTO `sys_apis` VALUES (1221, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/activity/getActivityAddRequestList', '分页获取活动报名申请列表', 12, 'GET');

INSERT INTO `sys_apis` VALUES (1301, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/createQuestion', '创建问题', 13, 'POST');
INSERT INTO `sys_apis` VALUES (1302, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/closeQuestion', '关闭问题', 13, 'POST');
INSERT INTO `sys_apis` VALUES (1303, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/answerQuestion', '回答问题', 13, 'POST');
INSERT INTO `sys_apis` VALUES (1304, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/delSelfAnswer', '删除自己的回答', 13, 'DELETE');
INSERT INTO `sys_apis` VALUES (1305, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/setBestAnswer', '设置最佳答案', 13, 'POST');
INSERT INTO `sys_apis` VALUES (1306, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/setQuestionScore', '给问题打分（1-5星）', 13, 'POST');
INSERT INTO `sys_apis` VALUES (1307, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/getGlobalRecommendQuestionList', '分页获取全局推荐问题列表', 13, 'GET');
INSERT INTO `sys_apis` VALUES (1308, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/getCircleQuestionList', '分页获取圈子问题列表', 13, 'GET');
INSERT INTO `sys_apis` VALUES (1309, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/question/getAnswerList', '分页获取问题的回答列表', 13, 'GET');

INSERT INTO `sys_apis` VALUES (1401, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/focusUser', '关注用户', 14, 'POST');
INSERT INTO `sys_apis` VALUES (1402, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/focusUserCancel', '取消用户关注', 14, 'POST');
INSERT INTO `sys_apis` VALUES (1403, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/updateFocusUser', '更新关注用户', 14, 'POST');
INSERT INTO `sys_apis` VALUES (1404, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/getFrequentBrowsingUserList', '分页获取经常浏览用户列表', 14, 'GET');
INSERT INTO `sys_apis` VALUES (1405, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/getFocusUserPostsList', '分页获取关注用户动态列表', 14, 'GET');
INSERT INTO `sys_apis` VALUES (1406, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/getFocusUserList', '分页获取关注用户列表', 14, 'GET');
INSERT INTO `sys_apis` VALUES (1407, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/getFansList', '分页获取粉丝列表', 14, 'GET');
INSERT INTO `sys_apis` VALUES (1408, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/focus/findFocusUser', '用id查询关注用户', 14, 'GET');

INSERT INTO `sys_apis` VALUES (1501, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/channel/getChannelList', '获取频道列表', 15, 'GET');
INSERT INTO `sys_apis` VALUES (1502, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/channel/getUserChannelList', '获取用户频道', 15, 'GET');
INSERT INTO `sys_apis` VALUES (1503, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/channel/setUserChannel', '设置用户频道', 15, 'POST');
INSERT INTO `sys_apis` VALUES (1504, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/channel/getGeneralChannelList', '获取常规频道列表', 15, 'GET');

INSERT INTO `sys_apis` VALUES (1601, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/getProductGoldList', '获取金币列表', 16, 'GET');
INSERT INTO `sys_apis` VALUES (1602, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/getPayTypeList', '获取支付类型列表', 16, 'GET');
INSERT INTO `sys_apis` VALUES (1603, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/createOrder', '创建订单', 16, 'POST');
INSERT INTO `sys_apis` VALUES (1604, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/getGoldBillList', '获取金币账单列表', 16, 'GET');
INSERT INTO `sys_apis` VALUES (1605, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/findGoldBill', '根据ID获取金币账单', 16, 'GET');
INSERT INTO `sys_apis` VALUES (1606, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/getExtractTypeList', '获取提现类型列表', 16, 'GET');
INSERT INTO `sys_apis` VALUES (1607, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/createUserExtract', '创建用户提现', 16, 'POST');
INSERT INTO `sys_apis` VALUES (1608, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/getUserBillList', '获取用户账单列表', 16, 'GET');
INSERT INTO `sys_apis` VALUES (1609, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/wallet/findUserBill', '根据ID获取用户账单', 16, 'GET');

INSERT INTO `sys_apis` VALUES (1701, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/getHelpTypeList', '(管理员)获取帮助类型列表', 17, 'GET');
INSERT INTO `sys_apis` VALUES (1702, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/getMainHelpList', '获取主页帮助列表', 17, 'GET');
INSERT INTO `sys_apis` VALUES (1703, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/findHelp', '用id查询帮助', 17, 'GET');
INSERT INTO `sys_apis` VALUES (1704, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/helpThumbsUp', '帮助点赞', 17, 'POST');
INSERT INTO `sys_apis` VALUES (1706, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/getHelpList', '获取帮助列表', 17, 'GET');
INSERT INTO `sys_apis` VALUES (1707, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/createFeedback', '创建反馈', 17, 'POST');
INSERT INTO `sys_apis` VALUES (1708, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/deleteFeedback', '(管理员)删除反馈', 17, 'DELETE');
INSERT INTO `sys_apis` VALUES (1709, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/deleteFeedbackByIds', '(管理员)批量删除反馈', 17, 'DELETE');
INSERT INTO `sys_apis` VALUES (1710, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/updateFeedback', '(管理员)更新反馈', 17, 'PUT');
INSERT INTO `sys_apis` VALUES (1711, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/findFeedback', '用id查询反馈', 17, 'GET');
INSERT INTO `sys_apis` VALUES (1712, '000000', '2023-02-06 18:19:13.303', '2023-02-06 18:19:13.303', NULL, 0, 0, NULL, NULL, NULL, '/app/helpCenter/getFeedbackList', '分页获取反馈列表', 17, 'GET');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `create_user` bigint NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint NULL DEFAULT NULL COMMENT '修改人',
  `is_del` int NULL DEFAULT 0 COMMENT '是否已删除',
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `status` int NULL DEFAULT 0 COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9529 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2023-02-06 18:19:13.322', '2023-02-06 18:19:13.602', NULL, 888, '普通用户', 0, 'dashboard', 0, NULL, NULL, NULL, 0, '000000', 0);
INSERT INTO `sys_authorities` VALUES ('2023-02-06 18:19:13.322', '2023-02-06 18:19:13.633', NULL, 8881, '普通用户子角色', 888, 'dashboard', 0, NULL, NULL, NULL, 0, '000000', 0);
INSERT INTO `sys_authorities` VALUES ('2023-02-06 18:19:13.322', '2023-02-06 18:19:13.614', NULL, 9528, '测试角色', 0, 'dashboard', 0, NULL, NULL, NULL, 0, '000000', 0);

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '角色按钮';

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '角色菜单';

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, 888);
INSERT INTO `sys_authority_menus` VALUES (1, 8881);
INSERT INTO `sys_authority_menus` VALUES (1, 9528);
INSERT INTO `sys_authority_menus` VALUES (2, 888);
INSERT INTO `sys_authority_menus` VALUES (2, 8881);
INSERT INTO `sys_authority_menus` VALUES (2, 9528);
INSERT INTO `sys_authority_menus` VALUES (3, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 8881);
INSERT INTO `sys_authority_menus` VALUES (5, 888);
INSERT INTO `sys_authority_menus` VALUES (5, 8881);
INSERT INTO `sys_authority_menus` VALUES (6, 888);
INSERT INTO `sys_authority_menus` VALUES (6, 8881);
INSERT INTO `sys_authority_menus` VALUES (7, 888);
INSERT INTO `sys_authority_menus` VALUES (7, 8881);
INSERT INTO `sys_authority_menus` VALUES (8, 888);
INSERT INTO `sys_authority_menus` VALUES (8, 8881);
INSERT INTO `sys_authority_menus` VALUES (8, 9528);
INSERT INTO `sys_authority_menus` VALUES (9, 888);
INSERT INTO `sys_authority_menus` VALUES (9, 8881);
INSERT INTO `sys_authority_menus` VALUES (10, 888);
INSERT INTO `sys_authority_menus` VALUES (10, 8881);
INSERT INTO `sys_authority_menus` VALUES (11, 888);
INSERT INTO `sys_authority_menus` VALUES (11, 8881);
INSERT INTO `sys_authority_menus` VALUES (12, 888);
INSERT INTO `sys_authority_menus` VALUES (13, 888);
INSERT INTO `sys_authority_menus` VALUES (13, 8881);
INSERT INTO `sys_authority_menus` VALUES (14, 888);
INSERT INTO `sys_authority_menus` VALUES (14, 8881);
INSERT INTO `sys_authority_menus` VALUES (15, 888);
INSERT INTO `sys_authority_menus` VALUES (15, 8881);
INSERT INTO `sys_authority_menus` VALUES (16, 888);
INSERT INTO `sys_authority_menus` VALUES (16, 8881);
INSERT INTO `sys_authority_menus` VALUES (17, 888);
INSERT INTO `sys_authority_menus` VALUES (17, 8881);
INSERT INTO `sys_authority_menus` VALUES (18, 888);
INSERT INTO `sys_authority_menus` VALUES (19, 888);
INSERT INTO `sys_authority_menus` VALUES (20, 888);
INSERT INTO `sys_authority_menus` VALUES (21, 888);
INSERT INTO `sys_authority_menus` VALUES (22, 888);
INSERT INTO `sys_authority_menus` VALUES (23, 888);
INSERT INTO `sys_authority_menus` VALUES (24, 888);
INSERT INTO `sys_authority_menus` VALUES (25, 888);
INSERT INTO `sys_authority_menus` VALUES (26, 888);
INSERT INTO `sys_authority_menus` VALUES (27, 888);
INSERT INTO `sys_authority_menus` VALUES (28, 888);
INSERT INTO `sys_authority_menus` VALUES (29, 888);

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------

-- ----------------------------
-- Table structure for sys_auto_codes
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_codes`;
CREATE TABLE `sys_auto_codes`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_codes_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_codes
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统菜单按钮';

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统菜单参数';

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统菜单';

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '仪表盘', 'odometer', 0);
INSERT INTO `sys_base_menus` VALUES (2, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0);
INSERT INTO `sys_base_menus` VALUES (3, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '超级管理员', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (4, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '', 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` VALUES (6, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'api管理', 'platform', 0);
INSERT INTO `sys_base_menus` VALUES (7, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '', 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` VALUES (8, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` VALUES (9, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 'dictionary', 0, 0, '字典详情-${id}', 'list', 0);
INSERT INTO `sys_base_menus` VALUES (10, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` VALUES (11, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (12, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 7, '', 0, 0, '示例文件', 'management', 0);
INSERT INTO `sys_base_menus` VALUES (13, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '12', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '', 0, 0, '媒体库（上传下载）', 'upload', 0);
INSERT INTO `sys_base_menus` VALUES (14, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '12', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` VALUES (15, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '12', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '', 0, 0, '客户列表（资源示例）', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (16, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (17, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '16', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` VALUES (18, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '16', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '', 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (19, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '16', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '', 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (20, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '16', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '', 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (21, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '16', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (22, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '16', 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '自动化package', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (23, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, '', 0, 0, '官方网站', 'home-filled', 0);
INSERT INTO `sys_base_menus` VALUES (24, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` VALUES (25, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '0', 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0);
INSERT INTO `sys_base_menus` VALUES (26, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '25', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0);
INSERT INTO `sys_base_menus` VALUES (27, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '25', 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0);
INSERT INTO `sys_base_menus` VALUES (28, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '25', 'autoPlug', 'autoPlug', 0, 'view/systemTools/autoPlug/autoPlug.vue', 2, '', 0, 0, '插件模板', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (29, '000000', '2023-02-06 18:19:13.412', '2023-02-06 18:19:13.412', NULL, 0, 0, NULL, NULL, NULL, 0, '25', 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 3, '', 0, 0, '邮件插件', 'message', 0);

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic ;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES (888, 888);
INSERT INTO `sys_data_authority_id` VALUES (888, 8881);
INSERT INTO `sys_data_authority_id` VALUES (888, 9528);
INSERT INTO `sys_data_authority_id` VALUES (9528, 8881);
INSERT INTO `sys_data_authority_id` VALUES (9528, 9528);

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统字典';

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '000000', '2023-02-06 18:19:13.350', '2023-02-06 18:19:13.358', NULL, 0, 1, NULL, NULL, NULL, '性别', 'gender', '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '000000', '2023-02-06 18:19:13.350', '2023-02-06 18:19:13.368', NULL, 0, 1, NULL, NULL, NULL, '数据库int类型', 'int', 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '000000', '2023-02-06 18:19:13.350', '2023-02-06 18:19:13.376', NULL, 0, 1, NULL, NULL, NULL, '数据库时间日期类型', 'time.Time', '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '000000', '2023-02-06 18:19:13.350', '2023-02-06 18:19:13.384', NULL, 0, 1, NULL, NULL, NULL, '数据库浮点型', 'float64', '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '000000', '2023-02-06 18:19:13.350', '2023-02-06 18:19:13.393', NULL, 0, 1, NULL, NULL, NULL, '数据库字符串', 'string', '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '000000', '2023-02-06 18:19:13.350', '2023-02-06 18:19:13.402', NULL, 0, 1, NULL, NULL, NULL, '数据库bool类型', 'bool', '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) NULL DEFAULT NULL COMMENT '字典值',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统字典详情';

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '000000', '2023-02-06 18:19:13.360', '2023-02-06 18:19:13.360', NULL, 0, 1, NULL, NULL, NULL, '男', 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '000000', '2023-02-06 18:19:13.360', '2023-02-06 18:19:13.360', NULL, 0, 1, NULL, NULL, NULL, '女', 2, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '000000', '2023-02-06 18:19:13.369', '2023-02-06 18:19:13.369', NULL, 0, 1, NULL, NULL, NULL, 'smallint', 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '000000', '2023-02-06 18:19:13.369', '2023-02-06 18:19:13.369', NULL, 0, 1, NULL, NULL, NULL, 'mediumint', 2, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '000000', '2023-02-06 18:19:13.369', '2023-02-06 18:19:13.369', NULL, 0, 1, NULL, NULL, NULL, 'int', 3, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '000000', '2023-02-06 18:19:13.369', '2023-02-06 18:19:13.369', NULL, 0, 1, NULL, NULL, NULL, 'bigint', 4, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '000000', '2023-02-06 18:19:13.377', '2023-02-06 18:19:13.377', NULL, 0, 1, NULL, NULL, NULL, 'date', 0, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '000000', '2023-02-06 18:19:13.377', '2023-02-06 18:19:13.377', NULL, 0, 1, NULL, NULL, NULL, 'time', 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '000000', '2023-02-06 18:19:13.377', '2023-02-06 18:19:13.377', NULL, 0, 1, NULL, NULL, NULL, 'year', 2, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '000000', '2023-02-06 18:19:13.377', '2023-02-06 18:19:13.377', NULL, 0, 1, NULL, NULL, NULL, 'datetime', 3, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (11, '000000', '2023-02-06 18:19:13.377', '2023-02-06 18:19:13.377', NULL, 0, 1, NULL, NULL, NULL, 'timestamp', 5, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '000000', '2023-02-06 18:19:13.386', '2023-02-06 18:19:13.386', NULL, 0, 1, NULL, NULL, NULL, 'float', 0, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '000000', '2023-02-06 18:19:13.386', '2023-02-06 18:19:13.386', NULL, 0, 1, NULL, NULL, NULL, 'double', 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (14, '000000', '2023-02-06 18:19:13.386', '2023-02-06 18:19:13.386', NULL, 0, 1, NULL, NULL, NULL, 'decimal', 2, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (15, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'char', 0, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'varchar', 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'tinyblob', 2, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'tinytext', 3, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'text', 4, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'blob', 5, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'mediumblob', 6, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'mediumtext', 7, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'longblob', 8, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '000000', '2023-02-06 18:19:13.394', '2023-02-06 18:19:13.394', NULL, 0, 1, NULL, NULL, NULL, 'longtext', 9, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '000000', '2023-02-06 18:19:13.403', '2023-02-06 18:19:13.403', NULL, 0, 1, NULL, NULL, NULL, 'tinyint', 0, 0, 6);

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` bigint(20) NULL DEFAULT NULL COMMENT '请求状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '延迟',
  `agent` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint(20) UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic COMMENT '系统字典详';

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 888);
INSERT INTO `sys_user_authority` VALUES (1, 8881);
INSERT INTO `sys_user_authority` VALUES (1, 9528);
INSERT INTO `sys_user_authority` VALUES (2, 888);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '000000' COMMENT '租户ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `is_del` smallint(5) UNSIGNED NULL DEFAULT 0 COMMENT '刪除标志',
  `status` smallint(6) NULL DEFAULT 0 COMMENT '状态',
  `create_user` bigint(20) NULL DEFAULT NULL COMMENT '创建人',
  `create_dept` bigint(20) NULL DEFAULT NULL COMMENT '创建部门',
  `update_user` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_users_uuid`(`uuid`) USING BTREE,
  INDEX `idx_sys_users_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '000000', '2023-02-06 18:19:13.573', '2023-02-06 18:19:13.579', NULL, 0, 0, NULL, NULL, NULL, '499e4d4c-6d35-4592-969b-97a922c8e6be', 'admin', '$2a$10$0hE2/AS2CYN.aHD.WERkeuu1eF3.oiyCWOsEzyReVbBFj6wMDW6E.', 'Mr.奇淼', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', 888, '17611111111', '333333333@qq.com', 1);
INSERT INTO `sys_users` VALUES (2, '000000', '2023-02-06 18:19:13.573', '2023-02-06 18:19:13.590', NULL, 0, 0, NULL, NULL, NULL, '51435709-11ff-4fa2-9e32-24b46ba4b1b5', 'a303176530', '$2a$10$dJ/gfz8EijiJhfFxm9JVvOTiTGDDCR1JuEe/JyZKAolWg5SvuYtwC', '用户1', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', 9528, '17611111111', '333333333@qq.com', 1);

SET FOREIGN_KEY_CHECKS = 1;
