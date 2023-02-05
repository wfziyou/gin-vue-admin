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

 Date: 05/02/2023 22:46:07
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
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (440, 'p', '888', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (441, 'p', '888', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (446, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (444, 'p', '888', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (445, 'p', '888', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (443, 'p', '888', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (442, 'p', '888', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (533, 'p', '888', '/app/user/getUserBaseInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (447, 'p', '888', '/authority/copyAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (448, 'p', '888', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (449, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (451, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (452, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (450, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (515, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (514, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (513, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (487, 'p', '888', '/autoCode/createPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (485, 'p', '888', '/autoCode/createPlug', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (482, 'p', '888', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (489, 'p', '888', '/autoCode/delPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (493, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (484, 'p', '888', '/autoCode/getColumn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (480, 'p', '888', '/autoCode/getDB', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (490, 'p', '888', '/autoCode/getMeta', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (488, 'p', '888', '/autoCode/getPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (492, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (481, 'p', '888', '/autoCode/getTables', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (486, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (483, 'p', '888', '/autoCode/preview', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (491, 'p', '888', '/autoCode/rollback', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (428, 'p', '888', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (454, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (453, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (477, 'p', '888', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (478, 'p', '888', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (476, 'p', '888', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (475, 'p', '888', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (479, 'p', '888', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (512, 'p', '888', '/email/emailTest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (465, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (466, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (469, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (470, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (464, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (471, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (467, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (468, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (516, 'p', '888', '/hkUser/createHkUser', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (522, 'p', '888', '/hkUser/createUser', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (517, 'p', '888', '/hkUser/deleteHkUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (518, 'p', '888', '/hkUser/deleteHkUserByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (523, 'p', '888', '/hkUser/deleteUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (524, 'p', '888', '/hkUser/deleteUserByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (520, 'p', '888', '/hkUser/findHkUser', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (526, 'p', '888', '/hkUser/findUser', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (521, 'p', '888', '/hkUser/getHkUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (527, 'p', '888', '/hkUser/getUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (519, 'p', '888', '/hkUser/updateHkUser', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (525, 'p', '888', '/hkUser/updateUser', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (429, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (455, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (463, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (457, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (459, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (461, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (456, 'p', '888', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (462, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (460, 'p', '888', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (458, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (510, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (511, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (509, 'p', '888', '/simpleUploader/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (499, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (500, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (502, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (503, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (501, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (495, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (496, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (497, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (498, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (494, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (504, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (507, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (508, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (505, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (506, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (472, 'p', '888', '/system/getServerInfo', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (473, 'p', '888', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (474, 'p', '888', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (431, 'p', '888', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (437, 'p', '888', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (528, 'p', '888', '/user/createUser', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (430, 'p', '888', '/user/deleteUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (529, 'p', '888', '/user/deleteUserByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (531, 'p', '888', '/user/findUser', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (435, 'p', '888', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (532, 'p', '888', '/user/getUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (432, 'p', '888', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (439, 'p', '888', '/user/resetPassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (434, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (436, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (438, 'p', '888', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (433, 'p', '888', '/user/setUserInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (530, 'p', '888', '/user/updateUser', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (395, 'p', '8881', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (396, 'p', '8881', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (399, 'p', '8881', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (400, 'p', '8881', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (398, 'p', '8881', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (397, 'p', '8881', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (427, 'p', '8881', '/app/user/getUserBaseInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (401, 'p', '8881', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (402, 'p', '8881', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (403, 'p', '8881', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (404, 'p', '8881', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (388, 'p', '8881', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (406, 'p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (405, 'p', '8881', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (424, 'p', '8881', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (425, 'p', '8881', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (423, 'p', '8881', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (422, 'p', '8881', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (426, 'p', '8881', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (417, 'p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (418, 'p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (419, 'p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (416, 'p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (389, 'p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (407, 'p', '8881', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (415, 'p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (409, 'p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (411, 'p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (413, 'p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (408, 'p', '8881', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (414, 'p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (412, 'p', '8881', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (410, 'p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (420, 'p', '8881', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (421, 'p', '8881', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (390, 'p', '8881', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (393, 'p', '8881', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (392, 'p', '8881', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (391, 'p', '8881', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (394, 'p', '8881', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (265, 'p', '9528', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (266, 'p', '9528', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (269, 'p', '9528', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (270, 'p', '9528', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (268, 'p', '9528', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (267, 'p', '9528', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (298, 'p', '9528', '/app/user/getUserBaseInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (271, 'p', '9528', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (272, 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (273, 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (274, 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (297, 'p', '9528', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (258, 'p', '9528', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (276, 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (275, 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (294, 'p', '9528', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (295, 'p', '9528', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (293, 'p', '9528', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (292, 'p', '9528', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (296, 'p', '9528', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (287, 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (288, 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (289, 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (286, 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (259, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (277, 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (285, 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (279, 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (281, 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (283, 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (278, 'p', '9528', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (284, 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (282, 'p', '9528', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (280, 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (290, 'p', '9528', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (291, 'p', '9528', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (260, 'p', '9528', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (263, 'p', '9528', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (262, 'p', '9528', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (261, 'p', '9528', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (264, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '', '', '');

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
-- Records of exa_customers
-- ----------------------------

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
-- Records of exa_file_chunks
-- ----------------------------

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
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2023-01-31 14:44:25.043', '2023-01-31 14:44:25.043', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2023-01-31 14:44:25.043', '2023-01-31 14:44:25.043', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of exa_files
-- ----------------------------

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
-- Records of hk_apply
-- ----------------------------

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
-- Records of hk_bug_report
-- ----------------------------

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
-- Records of hk_circle
-- ----------------------------

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
-- Records of hk_circle_add_request
-- ----------------------------

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
-- Records of hk_circle_apply
-- ----------------------------

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
-- Records of hk_circle_apply_group
-- ----------------------------

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
-- Records of hk_circle_classify
-- ----------------------------

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
-- Records of hk_circle_relation
-- ----------------------------

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
-- Records of hk_circle_request
-- ----------------------------

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
-- Records of hk_circle_user
-- ----------------------------

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
-- Records of hk_comment_thumbs_up
-- ----------------------------

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
-- Records of hk_forbidden_speak
-- ----------------------------

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
-- Records of hk_forbidden_speak_duration
-- ----------------------------

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
-- Records of hk_forbidden_speak_reason
-- ----------------------------

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
-- Records of hk_forum_comment
-- ----------------------------

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
-- Records of hk_forum_posts
-- ----------------------------

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
-- Records of hk_forum_posts_group
-- ----------------------------

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
-- Records of hk_forum_tag
-- ----------------------------

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
-- Records of hk_forum_tag_group
-- ----------------------------

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
-- Records of hk_forum_thumbs_up
-- ----------------------------

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
-- Records of hk_forum_topic
-- ----------------------------

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
-- Records of hk_forum_topic_group
-- ----------------------------

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
-- Records of hk_mini_program
-- ----------------------------

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
-- Records of hk_mini_program_packet
-- ----------------------------

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
-- Records of hk_plat_apply
-- ----------------------------

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
-- Records of hk_plat_apply_group
-- ----------------------------

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
-- Records of hk_protocol
-- ----------------------------

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
-- Records of hk_report
-- ----------------------------

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
-- Records of hk_report_reason
-- ----------------------------

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
-- Records of hk_user
-- ----------------------------

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
-- Records of hk_user_browsing_history
-- ----------------------------

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
-- Records of hk_user_circle_apply
-- ----------------------------

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
-- Records of hk_user_collect
-- ----------------------------

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
-- Records of hk_user_extend
-- ----------------------------

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
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (1, '2023-01-31 21:08:08.434', '2023-01-31 21:08:08.434', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMmM4YTZlNWYtZmIxMC00NGFmLTllMWItZmQ0N2FmODRmYWNhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzU3NTIyNzQsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3NTE0NjQ3NH0.cMIoCU66mCaJNbf18UJAUIvYxxfOXZYIEv5xe0pDT6s', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (2, '2023-02-01 08:13:22.955', '2023-02-01 08:13:22.955', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMmM4YTZlNWYtZmIxMC00NGFmLTllMWItZmQ0N2FmODRmYWNhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzU4MTUwOTQsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3NTIwOTI5NH0.oupJI9zcyo_Au_M0EVYiGlQiyLIArW0rbkq8O87Zv4A', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (3, '2023-02-01 15:18:25.856', '2023-02-01 15:18:25.856', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMmM4YTZlNWYtZmIxMC00NGFmLTllMWItZmQ0N2FmODRmYWNhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzU4NDA2OTgsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3NTIzNDg5OH0.kYk9wllPMUsJE3-ZTMablWHV8RfdMZOdt61rI_8gEtM', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `jwt_blacklists` VALUES (4, '2023-02-05 13:46:35.213', '2023-02-05 13:46:35.213', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMmM4YTZlNWYtZmIxMC00NGFmLTllMWItZmQ0N2FmODRmYWNhIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzYwMDE4ODYsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3NTM5NjA4Nn0.JlbZ8Hn7lOaZ3_aPXdKVqXlEt3rRoLvum1RVqomplgg', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/base/login', '用户登录(必选)', 'base', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (2, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (3, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (4, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (5, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (6, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (7, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (8, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (9, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (10, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (11, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (12, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (13, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/createApi', '创建api', 'api', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (14, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/deleteApi', '删除Api', 'api', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (15, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/updateApi', '更新Api', 'api', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (16, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/getApiList', '获取api列表', 'api', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (17, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (18, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (19, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (20, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (21, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (22, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (23, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (24, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (25, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (26, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (27, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (28, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (29, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (30, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (31, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (32, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (33, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (34, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (35, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (36, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (37, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (38, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (39, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (40, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (41, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (42, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (43, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (44, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (45, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (46, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (47, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (48, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/customer/customer', '更新客户', '客户', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (49, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/customer/customer', '创建客户', '客户', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (50, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/customer/customer', '删除客户', '客户', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (51, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/customer/customer', '获取单一客户', '客户', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (52, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (53, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (54, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (55, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (56, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (57, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (58, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (59, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (60, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (61, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (62, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (63, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (64, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (65, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (66, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (67, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (68, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (69, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (70, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (71, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (72, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (73, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (74, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (75, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (76, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (77, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (78, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (79, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (80, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (81, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (82, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (83, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (84, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (85, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (86, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (87, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (88, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (89, '2023-01-31 14:44:24.725', '2023-01-31 14:44:24.725', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (90, '2023-01-31 14:51:53.325', '2023-01-31 14:51:53.325', NULL, '/hkForumPosts/createHkForumPosts', '新增hkForumPosts表', 'hkForumPosts', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (91, '2023-01-31 14:51:53.329', '2023-01-31 14:51:53.329', NULL, '/hkForumPosts/deleteHkForumPosts', '删除hkForumPosts表', 'hkForumPosts', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (92, '2023-01-31 14:51:53.332', '2023-01-31 14:51:53.332', NULL, '/hkForumPosts/deleteHkForumPostsByIds', '批量删除hkForumPosts表', 'hkForumPosts', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (93, '2023-01-31 14:51:53.336', '2023-01-31 14:51:53.336', NULL, '/hkForumPosts/updateHkForumPosts', '更新hkForumPosts表', 'hkForumPosts', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (94, '2023-01-31 14:51:53.339', '2023-01-31 14:51:53.339', NULL, '/hkForumPosts/findHkForumPosts', '根据ID获取hkForumPosts表', 'hkForumPosts', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (95, '2023-01-31 14:51:53.342', '2023-01-31 14:51:53.342', NULL, '/hkForumPosts/getHkForumPostsList', '获取hkForumPosts表列表', 'hkForumPosts', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (96, '2023-01-31 15:22:01.150', '2023-01-31 15:22:01.150', NULL, '/hkUser/createHkUser', '新增hkUser表', 'hkUser', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (97, '2023-01-31 15:22:01.154', '2023-01-31 15:22:01.154', NULL, '/hkUser/deleteHkUser', '删除hkUser表', 'hkUser', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (98, '2023-01-31 15:22:01.157', '2023-01-31 15:22:01.157', NULL, '/hkUser/deleteHkUserByIds', '批量删除hkUser表', 'hkUser', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (99, '2023-01-31 15:22:01.161', '2023-01-31 15:22:01.161', NULL, '/hkUser/updateHkUser', '更新hkUser表', 'hkUser', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (100, '2023-01-31 15:22:01.165', '2023-01-31 15:22:01.165', NULL, '/hkUser/findHkUser', '根据ID获取hkUser表', 'hkUser', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (101, '2023-01-31 15:22:01.168', '2023-01-31 15:22:01.168', NULL, '/hkUser/getHkUserList', '获取hkUser表列表', 'hkUser', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (102, '2023-01-31 16:59:40.142', '2023-01-31 16:59:40.142', NULL, '/hkUser/createUser', '新增hkUser表', 'hkUser', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (103, '2023-01-31 16:59:40.147', '2023-01-31 16:59:40.147', NULL, '/hkUser/deleteUser', '删除hkUser表', 'hkUser', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (104, '2023-01-31 16:59:40.150', '2023-01-31 16:59:40.150', NULL, '/hkUser/deleteUserByIds', '批量删除hkUser表', 'hkUser', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (105, '2023-01-31 16:59:40.153', '2023-01-31 16:59:40.153', NULL, '/hkUser/updateUser', '更新hkUser表', 'hkUser', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (106, '2023-01-31 16:59:40.157', '2023-01-31 16:59:40.157', NULL, '/hkUser/findUser', '根据ID获取hkUser表', 'hkUser', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (107, '2023-01-31 16:59:40.160', '2023-01-31 16:59:40.160', NULL, '/hkUser/getUserList', '获取hkUser表列表', 'hkUser', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (108, '2023-01-31 17:03:35.179', '2023-01-31 17:03:35.179', NULL, '/user/createUser', '新增hkUser表', 'user', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (109, '2023-01-31 17:03:35.185', '2023-01-31 17:03:35.185', NULL, '/user/deleteUserByIds', '批量删除hkUser表', 'user', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (110, '2023-01-31 17:03:35.188', '2023-01-31 17:03:35.188', NULL, '/user/updateUser', '更新hkUser表', 'user', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (111, '2023-01-31 17:03:35.191', '2023-01-31 17:03:35.191', NULL, '/user/findUser', '根据ID获取hkUser表', 'user', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (112, '2023-01-31 17:03:35.195', '2023-01-31 17:03:35.195', NULL, '/user/getUserList', '获取hkUser表列表', 'user', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (113, '2023-02-01 08:19:17.081', '2023-02-01 08:19:17.081', NULL, '/hkApply/createHkApply', '新增hkApply表', 'hkApply', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (114, '2023-02-01 08:19:17.084', '2023-02-01 08:19:17.084', NULL, '/hkApply/deleteHkApply', '删除hkApply表', 'hkApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (115, '2023-02-01 08:19:17.087', '2023-02-01 08:19:17.087', NULL, '/hkApply/deleteHkApplyByIds', '批量删除hkApply表', 'hkApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (116, '2023-02-01 08:19:17.091', '2023-02-01 08:19:17.091', NULL, '/hkApply/updateHkApply', '更新hkApply表', 'hkApply', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (117, '2023-02-01 08:19:17.094', '2023-02-01 08:19:17.094', NULL, '/hkApply/findHkApply', '根据ID获取hkApply表', 'hkApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (118, '2023-02-01 08:19:17.097', '2023-02-01 08:19:17.097', NULL, '/hkApply/getHkApplyList', '获取hkApply表列表', 'hkApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (119, '2023-02-01 08:19:59.449', '2023-02-01 08:19:59.449', NULL, '/hkCircleApply/createHkCircleApply', '新增hkCircleApply表', 'hkCircleApply', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (120, '2023-02-01 08:19:59.453', '2023-02-01 08:19:59.453', NULL, '/hkCircleApply/deleteHkCircleApply', '删除hkCircleApply表', 'hkCircleApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (121, '2023-02-01 08:19:59.456', '2023-02-01 08:19:59.456', NULL, '/hkCircleApply/deleteHkCircleApplyByIds', '批量删除hkCircleApply表', 'hkCircleApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (122, '2023-02-01 08:19:59.459', '2023-02-01 08:19:59.459', NULL, '/hkCircleApply/updateHkCircleApply', '更新hkCircleApply表', 'hkCircleApply', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (123, '2023-02-01 08:19:59.463', '2023-02-01 08:19:59.463', NULL, '/hkCircleApply/findHkCircleApply', '根据ID获取hkCircleApply表', 'hkCircleApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (124, '2023-02-01 08:19:59.466', '2023-02-01 08:19:59.466', NULL, '/hkCircleApply/getHkCircleApplyList', '获取hkCircleApply表列表', 'hkCircleApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (125, '2023-02-01 08:20:25.479', '2023-02-01 08:20:25.479', NULL, '/hkCircleApplyGroup/createHkCircleApplyGroup', '新增hkCircleApplyGroup表', 'hkCircleApplyGroup', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (126, '2023-02-01 08:20:25.484', '2023-02-01 08:20:25.484', NULL, '/hkCircleApplyGroup/deleteHkCircleApplyGroup', '删除hkCircleApplyGroup表', 'hkCircleApplyGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (127, '2023-02-01 08:20:25.487', '2023-02-01 08:20:25.487', NULL, '/hkCircleApplyGroup/deleteHkCircleApplyGroupByIds', '批量删除hkCircleApplyGroup表', 'hkCircleApplyGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (128, '2023-02-01 08:20:25.490', '2023-02-01 08:20:25.490', NULL, '/hkCircleApplyGroup/updateHkCircleApplyGroup', '更新hkCircleApplyGroup表', 'hkCircleApplyGroup', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (129, '2023-02-01 08:20:25.495', '2023-02-01 08:20:25.495', NULL, '/hkCircleApplyGroup/findHkCircleApplyGroup', '根据ID获取hkCircleApplyGroup表', 'hkCircleApplyGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (130, '2023-02-01 08:20:25.499', '2023-02-01 08:20:25.499', NULL, '/hkCircleApplyGroup/getHkCircleApplyGroupList', '获取hkCircleApplyGroup表列表', 'hkCircleApplyGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (131, '2023-02-01 08:21:20.945', '2023-02-01 08:21:20.945', NULL, '/hkPlatApply/createHkPlatApply', '新增hkPlatApply表', 'hkPlatApply', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (132, '2023-02-01 08:21:20.948', '2023-02-01 08:21:20.948', NULL, '/hkPlatApply/deleteHkPlatApply', '删除hkPlatApply表', 'hkPlatApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (133, '2023-02-01 08:21:20.951', '2023-02-01 08:21:20.951', NULL, '/hkPlatApply/deleteHkPlatApplyByIds', '批量删除hkPlatApply表', 'hkPlatApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (134, '2023-02-01 08:21:20.955', '2023-02-01 08:21:20.955', NULL, '/hkPlatApply/updateHkPlatApply', '更新hkPlatApply表', 'hkPlatApply', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (135, '2023-02-01 08:21:20.959', '2023-02-01 08:21:20.959', NULL, '/hkPlatApply/findHkPlatApply', '根据ID获取hkPlatApply表', 'hkPlatApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (136, '2023-02-01 08:21:20.962', '2023-02-01 08:21:20.962', NULL, '/hkPlatApply/getHkPlatApplyList', '获取hkPlatApply表列表', 'hkPlatApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (137, '2023-02-01 08:21:42.953', '2023-02-01 08:21:42.953', NULL, '/hkPlatApplyGroup/createHkPlatApplyGroup', '新增hkPlatApplyGroup表', 'hkPlatApplyGroup', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (138, '2023-02-01 08:21:42.957', '2023-02-01 08:21:42.957', NULL, '/hkPlatApplyGroup/deleteHkPlatApplyGroup', '删除hkPlatApplyGroup表', 'hkPlatApplyGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (139, '2023-02-01 08:21:42.960', '2023-02-01 08:21:42.960', NULL, '/hkPlatApplyGroup/deleteHkPlatApplyGroupByIds', '批量删除hkPlatApplyGroup表', 'hkPlatApplyGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (140, '2023-02-01 08:21:42.964', '2023-02-01 08:21:42.964', NULL, '/hkPlatApplyGroup/updateHkPlatApplyGroup', '更新hkPlatApplyGroup表', 'hkPlatApplyGroup', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (141, '2023-02-01 08:21:42.967', '2023-02-01 08:21:42.967', NULL, '/hkPlatApplyGroup/findHkPlatApplyGroup', '根据ID获取hkPlatApplyGroup表', 'hkPlatApplyGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (142, '2023-02-01 08:21:42.971', '2023-02-01 08:21:42.971', NULL, '/hkPlatApplyGroup/getHkPlatApplyGroupList', '获取hkPlatApplyGroup表列表', 'hkPlatApplyGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (143, '2023-02-01 08:35:37.452', '2023-02-01 08:35:37.452', NULL, '/hkBugReport/createHkBugReport', '新增hkBugReport表', 'hkBugReport', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (144, '2023-02-01 08:35:37.455', '2023-02-01 08:35:37.455', NULL, '/hkBugReport/deleteHkBugReport', '删除hkBugReport表', 'hkBugReport', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (145, '2023-02-01 08:35:37.459', '2023-02-01 08:35:37.459', NULL, '/hkBugReport/deleteHkBugReportByIds', '批量删除hkBugReport表', 'hkBugReport', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (146, '2023-02-01 08:35:37.462', '2023-02-01 08:35:37.462', NULL, '/hkBugReport/updateHkBugReport', '更新hkBugReport表', 'hkBugReport', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (147, '2023-02-01 08:35:37.466', '2023-02-01 08:35:37.466', NULL, '/hkBugReport/findHkBugReport', '根据ID获取hkBugReport表', 'hkBugReport', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (148, '2023-02-01 08:35:37.469', '2023-02-01 08:35:37.469', NULL, '/hkBugReport/getHkBugReportList', '获取hkBugReport表列表', 'hkBugReport', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (149, '2023-02-01 08:36:46.258', '2023-02-01 08:36:46.258', NULL, '/hkCircle/createHkCircle', '新增hkCircle表', 'hkCircle', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (150, '2023-02-01 08:36:46.261', '2023-02-01 08:36:46.261', NULL, '/hkCircle/deleteHkCircle', '删除hkCircle表', 'hkCircle', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (151, '2023-02-01 08:36:46.265', '2023-02-01 08:36:46.265', NULL, '/hkCircle/deleteHkCircleByIds', '批量删除hkCircle表', 'hkCircle', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (152, '2023-02-01 08:36:46.268', '2023-02-01 08:36:46.268', NULL, '/hkCircle/updateHkCircle', '更新hkCircle表', 'hkCircle', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (153, '2023-02-01 08:36:46.271', '2023-02-01 08:36:46.271', NULL, '/hkCircle/findHkCircle', '根据ID获取hkCircle表', 'hkCircle', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (154, '2023-02-01 08:36:46.275', '2023-02-01 08:36:46.275', NULL, '/hkCircle/getHkCircleList', '获取hkCircle表列表', 'hkCircle', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (155, '2023-02-01 08:38:39.156', '2023-02-01 08:38:39.156', NULL, '/hkCircleAddRequest/createHkCircleAddRequest', '新增hkCircleAddRequest表', 'hkCircleAddRequest', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (156, '2023-02-01 08:38:39.160', '2023-02-01 08:38:39.160', NULL, '/hkCircleAddRequest/deleteHkCircleAddRequest', '删除hkCircleAddRequest表', 'hkCircleAddRequest', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (157, '2023-02-01 08:38:39.164', '2023-02-01 08:38:39.164', NULL, '/hkCircleAddRequest/deleteHkCircleAddRequestByIds', '批量删除hkCircleAddRequest表', 'hkCircleAddRequest', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (158, '2023-02-01 08:38:39.167', '2023-02-01 08:38:39.167', NULL, '/hkCircleAddRequest/updateHkCircleAddRequest', '更新hkCircleAddRequest表', 'hkCircleAddRequest', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (159, '2023-02-01 08:38:39.171', '2023-02-01 08:38:39.171', NULL, '/hkCircleAddRequest/findHkCircleAddRequest', '根据ID获取hkCircleAddRequest表', 'hkCircleAddRequest', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (160, '2023-02-01 08:38:39.174', '2023-02-01 08:38:39.174', NULL, '/hkCircleAddRequest/getHkCircleAddRequestList', '获取hkCircleAddRequest表列表', 'hkCircleAddRequest', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (161, '2023-02-01 08:39:55.356', '2023-02-01 08:39:55.356', NULL, '/hkCircleClassify/createHkCircleClassify', '新增hkCircleClassify表', 'hkCircleClassify', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (162, '2023-02-01 08:39:55.359', '2023-02-01 08:39:55.359', NULL, '/hkCircleClassify/deleteHkCircleClassify', '删除hkCircleClassify表', 'hkCircleClassify', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (163, '2023-02-01 08:39:55.362', '2023-02-01 08:39:55.362', NULL, '/hkCircleClassify/deleteHkCircleClassifyByIds', '批量删除hkCircleClassify表', 'hkCircleClassify', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (164, '2023-02-01 08:39:55.366', '2023-02-01 08:39:55.366', NULL, '/hkCircleClassify/updateHkCircleClassify', '更新hkCircleClassify表', 'hkCircleClassify', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (165, '2023-02-01 08:39:55.370', '2023-02-01 08:39:55.370', NULL, '/hkCircleClassify/findHkCircleClassify', '根据ID获取hkCircleClassify表', 'hkCircleClassify', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (166, '2023-02-01 08:39:55.373', '2023-02-01 08:39:55.373', NULL, '/hkCircleClassify/getHkCircleClassifyList', '获取hkCircleClassify表列表', 'hkCircleClassify', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (167, '2023-02-01 08:40:25.556', '2023-02-01 08:40:25.556', NULL, '/hkCircleRelation/createHkCircleRelation', '新增hkCircleRelation表', 'hkCircleRelation', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (168, '2023-02-01 08:40:25.560', '2023-02-01 08:40:25.560', NULL, '/hkCircleRelation/deleteHkCircleRelation', '删除hkCircleRelation表', 'hkCircleRelation', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (169, '2023-02-01 08:40:25.563', '2023-02-01 08:40:25.563', NULL, '/hkCircleRelation/deleteHkCircleRelationByIds', '批量删除hkCircleRelation表', 'hkCircleRelation', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (170, '2023-02-01 08:40:25.566', '2023-02-01 08:40:25.566', NULL, '/hkCircleRelation/updateHkCircleRelation', '更新hkCircleRelation表', 'hkCircleRelation', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (171, '2023-02-01 08:40:25.570', '2023-02-01 08:40:25.570', NULL, '/hkCircleRelation/findHkCircleRelation', '根据ID获取hkCircleRelation表', 'hkCircleRelation', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (172, '2023-02-01 08:40:25.574', '2023-02-01 08:40:25.574', NULL, '/hkCircleRelation/getHkCircleRelationList', '获取hkCircleRelation表列表', 'hkCircleRelation', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (173, '2023-02-01 08:40:58.589', '2023-02-01 08:40:58.589', NULL, '/hkCircleRequest/createHkCircleRequest', '新增hkCircleRequest表', 'hkCircleRequest', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (174, '2023-02-01 08:40:58.593', '2023-02-01 08:40:58.593', NULL, '/hkCircleRequest/deleteHkCircleRequest', '删除hkCircleRequest表', 'hkCircleRequest', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (175, '2023-02-01 08:40:58.597', '2023-02-01 08:40:58.597', NULL, '/hkCircleRequest/deleteHkCircleRequestByIds', '批量删除hkCircleRequest表', 'hkCircleRequest', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (176, '2023-02-01 08:40:58.602', '2023-02-01 08:40:58.602', NULL, '/hkCircleRequest/updateHkCircleRequest', '更新hkCircleRequest表', 'hkCircleRequest', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (177, '2023-02-01 08:40:58.606', '2023-02-01 08:40:58.606', NULL, '/hkCircleRequest/findHkCircleRequest', '根据ID获取hkCircleRequest表', 'hkCircleRequest', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (178, '2023-02-01 08:40:58.610', '2023-02-01 08:40:58.610', NULL, '/hkCircleRequest/getHkCircleRequestList', '获取hkCircleRequest表列表', 'hkCircleRequest', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (179, '2023-02-01 08:41:43.683', '2023-02-01 08:41:43.683', NULL, '/hkCircleUser/createHkCircleUser', '新增hkCircleUser表', 'hkCircleUser', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (180, '2023-02-01 08:41:43.687', '2023-02-01 08:41:43.687', NULL, '/hkCircleUser/deleteHkCircleUser', '删除hkCircleUser表', 'hkCircleUser', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (181, '2023-02-01 08:41:43.690', '2023-02-01 08:41:43.690', NULL, '/hkCircleUser/deleteHkCircleUserByIds', '批量删除hkCircleUser表', 'hkCircleUser', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (182, '2023-02-01 08:41:43.694', '2023-02-01 08:41:43.694', NULL, '/hkCircleUser/updateHkCircleUser', '更新hkCircleUser表', 'hkCircleUser', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (183, '2023-02-01 08:41:43.697', '2023-02-01 08:41:43.697', NULL, '/hkCircleUser/findHkCircleUser', '根据ID获取hkCircleUser表', 'hkCircleUser', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (184, '2023-02-01 08:41:43.701', '2023-02-01 08:41:43.701', NULL, '/hkCircleUser/getHkCircleUserList', '获取hkCircleUser表列表', 'hkCircleUser', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (185, '2023-02-01 08:42:12.758', '2023-02-01 08:42:12.758', NULL, '/hkCommentThumbsUp/createHkCommentThumbsUp', '新增hkCommentThumbsUp表', 'hkCommentThumbsUp', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (186, '2023-02-01 08:42:12.762', '2023-02-01 08:42:12.762', NULL, '/hkCommentThumbsUp/deleteHkCommentThumbsUp', '删除hkCommentThumbsUp表', 'hkCommentThumbsUp', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (187, '2023-02-01 08:42:12.765', '2023-02-01 08:42:12.765', NULL, '/hkCommentThumbsUp/deleteHkCommentThumbsUpByIds', '批量删除hkCommentThumbsUp表', 'hkCommentThumbsUp', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (188, '2023-02-01 08:42:12.768', '2023-02-01 08:42:12.768', NULL, '/hkCommentThumbsUp/updateHkCommentThumbsUp', '更新hkCommentThumbsUp表', 'hkCommentThumbsUp', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (189, '2023-02-01 08:42:12.772', '2023-02-01 08:42:12.772', NULL, '/hkCommentThumbsUp/findHkCommentThumbsUp', '根据ID获取hkCommentThumbsUp表', 'hkCommentThumbsUp', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (190, '2023-02-01 08:42:12.776', '2023-02-01 08:42:12.776', NULL, '/hkCommentThumbsUp/getHkCommentThumbsUpList', '获取hkCommentThumbsUp表列表', 'hkCommentThumbsUp', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (191, '2023-02-01 08:43:05.489', '2023-02-01 08:43:05.489', NULL, '/hkForbiddenSpeak/createHkForbiddenSpeak', '新增hkForbiddenSpeak表', 'hkForbiddenSpeak', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (192, '2023-02-01 08:43:05.492', '2023-02-01 08:43:05.492', NULL, '/hkForbiddenSpeak/deleteHkForbiddenSpeak', '删除hkForbiddenSpeak表', 'hkForbiddenSpeak', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (193, '2023-02-01 08:43:05.495', '2023-02-01 08:43:05.495', NULL, '/hkForbiddenSpeak/deleteHkForbiddenSpeakByIds', '批量删除hkForbiddenSpeak表', 'hkForbiddenSpeak', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (194, '2023-02-01 08:43:05.498', '2023-02-01 08:43:05.498', NULL, '/hkForbiddenSpeak/updateHkForbiddenSpeak', '更新hkForbiddenSpeak表', 'hkForbiddenSpeak', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (195, '2023-02-01 08:43:05.501', '2023-02-01 08:43:05.501', NULL, '/hkForbiddenSpeak/findHkForbiddenSpeak', '根据ID获取hkForbiddenSpeak表', 'hkForbiddenSpeak', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (196, '2023-02-01 08:43:05.505', '2023-02-01 08:43:05.505', NULL, '/hkForbiddenSpeak/getHkForbiddenSpeakList', '获取hkForbiddenSpeak表列表', 'hkForbiddenSpeak', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (197, '2023-02-01 08:43:29.984', '2023-02-01 08:43:29.984', NULL, '/hkForbiddenSpeakDuration/createHkForbiddenSpeakDuration', '新增hkForbiddenSpeakDuration表', 'hkForbiddenSpeakDuration', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (198, '2023-02-01 08:43:29.990', '2023-02-01 08:43:29.990', NULL, '/hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDuration', '删除hkForbiddenSpeakDuration表', 'hkForbiddenSpeakDuration', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (199, '2023-02-01 08:43:29.994', '2023-02-01 08:43:29.994', NULL, '/hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDurationByIds', '批量删除hkForbiddenSpeakDuration表', 'hkForbiddenSpeakDuration', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (200, '2023-02-01 08:43:29.998', '2023-02-01 08:43:29.998', NULL, '/hkForbiddenSpeakDuration/updateHkForbiddenSpeakDuration', '更新hkForbiddenSpeakDuration表', 'hkForbiddenSpeakDuration', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (201, '2023-02-01 08:43:30.001', '2023-02-01 08:43:30.001', NULL, '/hkForbiddenSpeakDuration/findHkForbiddenSpeakDuration', '根据ID获取hkForbiddenSpeakDuration表', 'hkForbiddenSpeakDuration', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (202, '2023-02-01 08:43:30.005', '2023-02-01 08:43:30.005', NULL, '/hkForbiddenSpeakDuration/getHkForbiddenSpeakDurationList', '获取hkForbiddenSpeakDuration表列表', 'hkForbiddenSpeakDuration', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (203, '2023-02-01 08:43:51.481', '2023-02-01 08:43:51.481', NULL, '/hkForbiddenSpeakReason/createHkForbiddenSpeakReason', '新增hkForbiddenSpeakReason表', 'hkForbiddenSpeakReason', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (204, '2023-02-01 08:43:51.486', '2023-02-01 08:43:51.486', NULL, '/hkForbiddenSpeakReason/deleteHkForbiddenSpeakReason', '删除hkForbiddenSpeakReason表', 'hkForbiddenSpeakReason', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (205, '2023-02-01 08:43:51.489', '2023-02-01 08:43:51.489', NULL, '/hkForbiddenSpeakReason/deleteHkForbiddenSpeakReasonByIds', '批量删除hkForbiddenSpeakReason表', 'hkForbiddenSpeakReason', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (206, '2023-02-01 08:43:51.492', '2023-02-01 08:43:51.492', NULL, '/hkForbiddenSpeakReason/updateHkForbiddenSpeakReason', '更新hkForbiddenSpeakReason表', 'hkForbiddenSpeakReason', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (207, '2023-02-01 08:43:51.496', '2023-02-01 08:43:51.496', NULL, '/hkForbiddenSpeakReason/findHkForbiddenSpeakReason', '根据ID获取hkForbiddenSpeakReason表', 'hkForbiddenSpeakReason', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (208, '2023-02-01 08:43:51.499', '2023-02-01 08:43:51.499', NULL, '/hkForbiddenSpeakReason/getHkForbiddenSpeakReasonList', '获取hkForbiddenSpeakReason表列表', 'hkForbiddenSpeakReason', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (209, '2023-02-01 08:45:09.499', '2023-02-01 08:45:09.499', NULL, '/hkForumComment/createHkForumComment', '新增hkForumComment表', 'hkForumComment', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (210, '2023-02-01 08:45:09.502', '2023-02-01 08:45:09.502', NULL, '/hkForumComment/deleteHkForumComment', '删除hkForumComment表', 'hkForumComment', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (211, '2023-02-01 08:45:09.505', '2023-02-01 08:45:09.505', NULL, '/hkForumComment/deleteHkForumCommentByIds', '批量删除hkForumComment表', 'hkForumComment', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (212, '2023-02-01 08:45:09.509', '2023-02-01 08:45:09.509', NULL, '/hkForumComment/updateHkForumComment', '更新hkForumComment表', 'hkForumComment', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (213, '2023-02-01 08:45:09.513', '2023-02-01 08:45:09.513', NULL, '/hkForumComment/findHkForumComment', '根据ID获取hkForumComment表', 'hkForumComment', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (214, '2023-02-01 08:45:09.516', '2023-02-01 08:45:09.516', NULL, '/hkForumComment/getHkForumCommentList', '获取hkForumComment表列表', 'hkForumComment', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (215, '2023-02-01 08:45:39.580', '2023-02-01 08:45:39.580', NULL, '/hkForumPostsGroup/createHkForumPostsGroup', '新增hkForumPostsGroup表', 'hkForumPostsGroup', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (216, '2023-02-01 08:45:39.584', '2023-02-01 08:45:39.584', NULL, '/hkForumPostsGroup/deleteHkForumPostsGroup', '删除hkForumPostsGroup表', 'hkForumPostsGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (217, '2023-02-01 08:45:39.588', '2023-02-01 08:45:39.588', NULL, '/hkForumPostsGroup/deleteHkForumPostsGroupByIds', '批量删除hkForumPostsGroup表', 'hkForumPostsGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (218, '2023-02-01 08:45:39.592', '2023-02-01 08:45:39.592', NULL, '/hkForumPostsGroup/updateHkForumPostsGroup', '更新hkForumPostsGroup表', 'hkForumPostsGroup', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (219, '2023-02-01 08:45:39.595', '2023-02-01 08:45:39.595', NULL, '/hkForumPostsGroup/findHkForumPostsGroup', '根据ID获取hkForumPostsGroup表', 'hkForumPostsGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (220, '2023-02-01 08:45:39.598', '2023-02-01 08:45:39.598', NULL, '/hkForumPostsGroup/getHkForumPostsGroupList', '获取hkForumPostsGroup表列表', 'hkForumPostsGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (221, '2023-02-01 08:46:04.291', '2023-02-01 08:46:04.291', NULL, '/hkForumTag/createHkForumTag', '新增hkForumTag表', 'hkForumTag', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (222, '2023-02-01 08:46:04.295', '2023-02-01 08:46:04.295', NULL, '/hkForumTag/deleteHkForumTag', '删除hkForumTag表', 'hkForumTag', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (223, '2023-02-01 08:46:04.298', '2023-02-01 08:46:04.298', NULL, '/hkForumTag/deleteHkForumTagByIds', '批量删除hkForumTag表', 'hkForumTag', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (224, '2023-02-01 08:46:04.302', '2023-02-01 08:46:04.302', NULL, '/hkForumTag/updateHkForumTag', '更新hkForumTag表', 'hkForumTag', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (225, '2023-02-01 08:46:04.306', '2023-02-01 08:46:04.306', NULL, '/hkForumTag/findHkForumTag', '根据ID获取hkForumTag表', 'hkForumTag', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (226, '2023-02-01 08:46:04.310', '2023-02-01 08:46:04.310', NULL, '/hkForumTag/getHkForumTagList', '获取hkForumTag表列表', 'hkForumTag', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (227, '2023-02-01 08:46:25.657', '2023-02-01 08:46:25.657', NULL, '/hkForumTagGroup/createHkForumTagGroup', '新增hkForumTagGroup表', 'hkForumTagGroup', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (228, '2023-02-01 08:46:25.660', '2023-02-01 08:46:25.660', NULL, '/hkForumTagGroup/deleteHkForumTagGroup', '删除hkForumTagGroup表', 'hkForumTagGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (229, '2023-02-01 08:46:25.664', '2023-02-01 08:46:25.664', NULL, '/hkForumTagGroup/deleteHkForumTagGroupByIds', '批量删除hkForumTagGroup表', 'hkForumTagGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (230, '2023-02-01 08:46:25.667', '2023-02-01 08:46:25.667', NULL, '/hkForumTagGroup/updateHkForumTagGroup', '更新hkForumTagGroup表', 'hkForumTagGroup', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (231, '2023-02-01 08:46:25.671', '2023-02-01 08:46:25.671', NULL, '/hkForumTagGroup/findHkForumTagGroup', '根据ID获取hkForumTagGroup表', 'hkForumTagGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (232, '2023-02-01 08:46:25.675', '2023-02-01 08:46:25.675', NULL, '/hkForumTagGroup/getHkForumTagGroupList', '获取hkForumTagGroup表列表', 'hkForumTagGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (233, '2023-02-01 08:48:25.180', '2023-02-01 08:48:25.180', NULL, '/hkForumThumbsUp/createHkForumThumbsUp', '新增hkForumThumbsUp表', 'hkForumThumbsUp', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (234, '2023-02-01 08:48:25.184', '2023-02-01 08:48:25.184', NULL, '/hkForumThumbsUp/deleteHkForumThumbsUp', '删除hkForumThumbsUp表', 'hkForumThumbsUp', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (235, '2023-02-01 08:48:25.188', '2023-02-01 08:48:25.188', NULL, '/hkForumThumbsUp/deleteHkForumThumbsUpByIds', '批量删除hkForumThumbsUp表', 'hkForumThumbsUp', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (236, '2023-02-01 08:48:25.191', '2023-02-01 08:48:25.191', NULL, '/hkForumThumbsUp/updateHkForumThumbsUp', '更新hkForumThumbsUp表', 'hkForumThumbsUp', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (237, '2023-02-01 08:48:25.194', '2023-02-01 08:48:25.194', NULL, '/hkForumThumbsUp/findHkForumThumbsUp', '根据ID获取hkForumThumbsUp表', 'hkForumThumbsUp', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (238, '2023-02-01 08:48:25.198', '2023-02-01 08:48:25.198', NULL, '/hkForumThumbsUp/getHkForumThumbsUpList', '获取hkForumThumbsUp表列表', 'hkForumThumbsUp', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (239, '2023-02-01 08:48:42.752', '2023-02-01 08:48:42.752', NULL, '/hkForumTopic/createHkForumTopic', '新增hkForumTopic表', 'hkForumTopic', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (240, '2023-02-01 08:48:42.756', '2023-02-01 08:48:42.756', NULL, '/hkForumTopic/deleteHkForumTopic', '删除hkForumTopic表', 'hkForumTopic', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (241, '2023-02-01 08:48:42.759', '2023-02-01 08:48:42.759', NULL, '/hkForumTopic/deleteHkForumTopicByIds', '批量删除hkForumTopic表', 'hkForumTopic', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (242, '2023-02-01 08:48:42.763', '2023-02-01 08:48:42.763', NULL, '/hkForumTopic/updateHkForumTopic', '更新hkForumTopic表', 'hkForumTopic', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (243, '2023-02-01 08:48:42.766', '2023-02-01 08:48:42.766', NULL, '/hkForumTopic/findHkForumTopic', '根据ID获取hkForumTopic表', 'hkForumTopic', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (244, '2023-02-01 08:48:42.770', '2023-02-01 08:48:42.770', NULL, '/hkForumTopic/getHkForumTopicList', '获取hkForumTopic表列表', 'hkForumTopic', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (245, '2023-02-01 08:49:10.089', '2023-02-01 08:49:10.089', NULL, '/hkForumTopicGroup/createHkForumTopicGroup', '新增hkForumTopicGroup表', 'hkForumTopicGroup', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (246, '2023-02-01 08:49:10.093', '2023-02-01 08:49:10.093', NULL, '/hkForumTopicGroup/deleteHkForumTopicGroup', '删除hkForumTopicGroup表', 'hkForumTopicGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (247, '2023-02-01 08:49:10.096', '2023-02-01 08:49:10.096', NULL, '/hkForumTopicGroup/deleteHkForumTopicGroupByIds', '批量删除hkForumTopicGroup表', 'hkForumTopicGroup', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (248, '2023-02-01 08:49:10.099', '2023-02-01 08:49:10.099', NULL, '/hkForumTopicGroup/updateHkForumTopicGroup', '更新hkForumTopicGroup表', 'hkForumTopicGroup', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (249, '2023-02-01 08:49:10.103', '2023-02-01 08:49:10.103', NULL, '/hkForumTopicGroup/findHkForumTopicGroup', '根据ID获取hkForumTopicGroup表', 'hkForumTopicGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (250, '2023-02-01 08:49:10.107', '2023-02-01 08:49:10.107', NULL, '/hkForumTopicGroup/getHkForumTopicGroupList', '获取hkForumTopicGroup表列表', 'hkForumTopicGroup', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (251, '2023-02-01 08:50:01.501', '2023-02-01 08:50:01.501', NULL, '/hkMiniProgram/createHkMiniProgram', '新增hkMiniProgram表', 'hkMiniProgram', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (252, '2023-02-01 08:50:01.505', '2023-02-01 08:50:01.505', NULL, '/hkMiniProgram/deleteHkMiniProgram', '删除hkMiniProgram表', 'hkMiniProgram', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (253, '2023-02-01 08:50:01.509', '2023-02-01 08:50:01.509', NULL, '/hkMiniProgram/deleteHkMiniProgramByIds', '批量删除hkMiniProgram表', 'hkMiniProgram', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (254, '2023-02-01 08:50:01.512', '2023-02-01 08:50:01.512', NULL, '/hkMiniProgram/updateHkMiniProgram', '更新hkMiniProgram表', 'hkMiniProgram', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (255, '2023-02-01 08:50:01.516', '2023-02-01 08:50:01.516', NULL, '/hkMiniProgram/findHkMiniProgram', '根据ID获取hkMiniProgram表', 'hkMiniProgram', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (256, '2023-02-01 08:50:01.519', '2023-02-01 08:50:01.519', NULL, '/hkMiniProgram/getHkMiniProgramList', '获取hkMiniProgram表列表', 'hkMiniProgram', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (257, '2023-02-01 08:50:25.781', '2023-02-01 08:50:25.781', NULL, '/hkMiniProgramPacket/createHkMiniProgramPacket', '新增hkMiniProgramPacket表', 'hkMiniProgramPacket', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (258, '2023-02-01 08:50:25.785', '2023-02-01 08:50:25.785', NULL, '/hkMiniProgramPacket/deleteHkMiniProgramPacket', '删除hkMiniProgramPacket表', 'hkMiniProgramPacket', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (259, '2023-02-01 08:50:25.788', '2023-02-01 08:50:25.788', NULL, '/hkMiniProgramPacket/deleteHkMiniProgramPacketByIds', '批量删除hkMiniProgramPacket表', 'hkMiniProgramPacket', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (260, '2023-02-01 08:50:25.791', '2023-02-01 08:50:25.791', NULL, '/hkMiniProgramPacket/updateHkMiniProgramPacket', '更新hkMiniProgramPacket表', 'hkMiniProgramPacket', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (261, '2023-02-01 08:50:25.795', '2023-02-01 08:50:25.795', NULL, '/hkMiniProgramPacket/findHkMiniProgramPacket', '根据ID获取hkMiniProgramPacket表', 'hkMiniProgramPacket', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (262, '2023-02-01 08:50:25.798', '2023-02-01 08:50:25.798', NULL, '/hkMiniProgramPacket/getHkMiniProgramPacketList', '获取hkMiniProgramPacket表列表', 'hkMiniProgramPacket', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (263, '2023-02-01 08:50:57.659', '2023-02-01 08:50:57.659', NULL, '/hkProtocol/createHkProtocol', '新增hkProtocol表', 'hkProtocol', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (264, '2023-02-01 08:50:57.663', '2023-02-01 08:50:57.663', NULL, '/hkProtocol/deleteHkProtocol', '删除hkProtocol表', 'hkProtocol', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (265, '2023-02-01 08:50:57.666', '2023-02-01 08:50:57.666', NULL, '/hkProtocol/deleteHkProtocolByIds', '批量删除hkProtocol表', 'hkProtocol', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (266, '2023-02-01 08:50:57.671', '2023-02-01 08:50:57.671', NULL, '/hkProtocol/updateHkProtocol', '更新hkProtocol表', 'hkProtocol', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (267, '2023-02-01 08:50:57.675', '2023-02-01 08:50:57.675', NULL, '/hkProtocol/findHkProtocol', '根据ID获取hkProtocol表', 'hkProtocol', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (268, '2023-02-01 08:50:57.679', '2023-02-01 08:50:57.679', NULL, '/hkProtocol/getHkProtocolList', '获取hkProtocol表列表', 'hkProtocol', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (269, '2023-02-01 08:57:58.760', '2023-02-01 08:57:58.760', NULL, '/hkReport/createHkReport', '新增hkReport表', 'hkReport', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (270, '2023-02-01 08:57:58.763', '2023-02-01 08:57:58.763', NULL, '/hkReport/deleteHkReport', '删除hkReport表', 'hkReport', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (271, '2023-02-01 08:57:58.767', '2023-02-01 08:57:58.767', NULL, '/hkReport/deleteHkReportByIds', '批量删除hkReport表', 'hkReport', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (272, '2023-02-01 08:57:58.771', '2023-02-01 08:57:58.771', NULL, '/hkReport/updateHkReport', '更新hkReport表', 'hkReport', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (273, '2023-02-01 08:57:58.774', '2023-02-01 08:57:58.774', NULL, '/hkReport/findHkReport', '根据ID获取hkReport表', 'hkReport', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (274, '2023-02-01 08:57:58.778', '2023-02-01 08:57:58.778', NULL, '/hkReport/getHkReportList', '获取hkReport表列表', 'hkReport', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (275, '2023-02-01 08:58:39.559', '2023-02-01 08:58:39.559', NULL, '/hkReportReason/createHkReportReason', '新增hkReportReason表', 'hkReportReason', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (276, '2023-02-01 08:58:39.563', '2023-02-01 08:58:39.563', NULL, '/hkReportReason/deleteHkReportReason', '删除hkReportReason表', 'hkReportReason', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (277, '2023-02-01 08:58:39.567', '2023-02-01 08:58:39.567', NULL, '/hkReportReason/deleteHkReportReasonByIds', '批量删除hkReportReason表', 'hkReportReason', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (278, '2023-02-01 08:58:39.570', '2023-02-01 08:58:39.570', NULL, '/hkReportReason/updateHkReportReason', '更新hkReportReason表', 'hkReportReason', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (279, '2023-02-01 08:58:39.574', '2023-02-01 08:58:39.574', NULL, '/hkReportReason/findHkReportReason', '根据ID获取hkReportReason表', 'hkReportReason', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (280, '2023-02-01 08:58:39.577', '2023-02-01 08:58:39.577', NULL, '/hkReportReason/getHkReportReasonList', '获取hkReportReason表列表', 'hkReportReason', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (281, '2023-02-01 09:00:28.082', '2023-02-01 09:00:28.082', NULL, '/hkUserBrowsingHistory/createHkUserBrowsingHistory', '新增hkUserBrowsingHistory表', 'hkUserBrowsingHistory', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (282, '2023-02-01 09:00:28.086', '2023-02-01 09:00:28.086', NULL, '/hkUserBrowsingHistory/deleteHkUserBrowsingHistory', '删除hkUserBrowsingHistory表', 'hkUserBrowsingHistory', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (283, '2023-02-01 09:00:28.090', '2023-02-01 09:00:28.090', NULL, '/hkUserBrowsingHistory/deleteHkUserBrowsingHistoryByIds', '批量删除hkUserBrowsingHistory表', 'hkUserBrowsingHistory', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (284, '2023-02-01 09:00:28.094', '2023-02-01 09:00:28.094', NULL, '/hkUserBrowsingHistory/updateHkUserBrowsingHistory', '更新hkUserBrowsingHistory表', 'hkUserBrowsingHistory', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (285, '2023-02-01 09:00:28.097', '2023-02-01 09:00:28.097', NULL, '/hkUserBrowsingHistory/findHkUserBrowsingHistory', '根据ID获取hkUserBrowsingHistory表', 'hkUserBrowsingHistory', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (286, '2023-02-01 09:00:28.100', '2023-02-01 09:00:28.100', NULL, '/hkUserBrowsingHistory/getHkUserBrowsingHistoryList', '获取hkUserBrowsingHistory表列表', 'hkUserBrowsingHistory', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (287, '2023-02-01 09:05:01.250', '2023-02-01 09:05:01.250', NULL, '/hkUserCircleApply/createHkUserCircleApply', '新增hkUserCircleApply表', 'hkUserCircleApply', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (288, '2023-02-01 09:05:01.253', '2023-02-01 09:05:01.253', NULL, '/hkUserCircleApply/deleteHkUserCircleApply', '删除hkUserCircleApply表', 'hkUserCircleApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (289, '2023-02-01 09:05:01.257', '2023-02-01 09:05:01.257', NULL, '/hkUserCircleApply/deleteHkUserCircleApplyByIds', '批量删除hkUserCircleApply表', 'hkUserCircleApply', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (290, '2023-02-01 09:05:01.260', '2023-02-01 09:05:01.260', NULL, '/hkUserCircleApply/updateHkUserCircleApply', '更新hkUserCircleApply表', 'hkUserCircleApply', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (291, '2023-02-01 09:05:01.264', '2023-02-01 09:05:01.264', NULL, '/hkUserCircleApply/findHkUserCircleApply', '根据ID获取hkUserCircleApply表', 'hkUserCircleApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (292, '2023-02-01 09:05:01.268', '2023-02-01 09:05:01.268', NULL, '/hkUserCircleApply/getHkUserCircleApplyList', '获取hkUserCircleApply表列表', 'hkUserCircleApply', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (293, '2023-02-01 09:05:52.857', '2023-02-01 09:05:52.857', NULL, '/hkUserCollect/createHkUserCollect', '新增hkUserCollect表', 'hkUserCollect', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (294, '2023-02-01 09:05:52.861', '2023-02-01 09:05:52.861', NULL, '/hkUserCollect/deleteHkUserCollect', '删除hkUserCollect表', 'hkUserCollect', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (295, '2023-02-01 09:05:52.865', '2023-02-01 09:05:52.865', NULL, '/hkUserCollect/deleteHkUserCollectByIds', '批量删除hkUserCollect表', 'hkUserCollect', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (296, '2023-02-01 09:05:52.868', '2023-02-01 09:05:52.868', NULL, '/hkUserCollect/updateHkUserCollect', '更新hkUserCollect表', 'hkUserCollect', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (297, '2023-02-01 09:05:52.871', '2023-02-01 09:05:52.871', NULL, '/hkUserCollect/findHkUserCollect', '根据ID获取hkUserCollect表', 'hkUserCollect', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (298, '2023-02-01 09:05:52.875', '2023-02-01 09:05:52.875', NULL, '/hkUserCollect/getHkUserCollectList', '获取hkUserCollect表列表', 'hkUserCollect', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (299, '2023-02-01 09:09:52.781', '2023-02-01 09:09:52.781', NULL, '/hkUserExtend/createHkUserExtend', '新增hkUserExtend表', 'hkUserExtend', 'POST', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (300, '2023-02-01 09:09:52.785', '2023-02-01 09:09:52.785', NULL, '/hkUserExtend/deleteHkUserExtend', '删除hkUserExtend表', 'hkUserExtend', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (301, '2023-02-01 09:09:52.788', '2023-02-01 09:09:52.788', NULL, '/hkUserExtend/deleteHkUserExtendByIds', '批量删除hkUserExtend表', 'hkUserExtend', 'DELETE', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (302, '2023-02-01 09:09:52.792', '2023-02-01 09:09:52.792', NULL, '/hkUserExtend/updateHkUserExtend', '更新hkUserExtend表', 'hkUserExtend', 'PUT', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (303, '2023-02-01 09:09:52.795', '2023-02-01 09:09:52.795', NULL, '/hkUserExtend/findHkUserExtend', '根据ID获取hkUserExtend表', 'hkUserExtend', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (304, '2023-02-01 09:09:52.799', '2023-02-01 09:09:52.799', NULL, '/hkUserExtend/getHkUserExtendList', '获取hkUserExtend表列表', 'hkUserExtend', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_apis` VALUES (305, '2023-02-05 14:18:46.323', '2023-02-05 14:18:46.323', NULL, '/app/user/getUserBaseInfo', '获取用户基本信息', 'app', 'GET', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2023-01-31 14:44:24.741', '2023-02-05 13:48:11.707', NULL, 888, '普通用户', 0, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-01-31 14:44:24.741', '2023-02-05 13:48:22.099', NULL, 8881, '普通用户子角色', 888, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-01-31 14:44:24.741', '2023-02-05 13:48:29.560', NULL, 9528, '测试角色', 0, 'dashboard');

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `sys_authority_menus` VALUES (3, 8881);
INSERT INTO `sys_authority_menus` VALUES (3, 9528);
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
INSERT INTO `sys_authority_menus` VALUES (12, 8881);
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
INSERT INTO `sys_authority_menus` VALUES (30, 888);
INSERT INTO `sys_authority_menus` VALUES (30, 8881);
INSERT INTO `sys_authority_menus` VALUES (30, 9528);

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
-- Records of sys_auto_code_histories
-- ----------------------------
INSERT INTO `sys_auto_code_histories` VALUES (1, '2023-01-31 14:51:53.824', '2023-01-31 14:51:53.824', NULL, 'community', '', 'hk_forum_posts', '{\"structName\":\"HkForumPosts\",\"tableName\":\"hk_forum_posts\",\"packageName\":\"hkForumPosts\",\"humpPackageName\":\"hk_forum_posts\",\"abbreviation\":\"hkForumPosts\",\"description\":\"hkForumPosts表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Category\",\"fieldDesc\":\"类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）\",\"fieldType\":\"int\",\"fieldJson\":\"category\",\"dataTypeLong\":\"10\",\"comment\":\"类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）\",\"columnName\":\"category\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"GroupId\",\"fieldDesc\":\"帖子分类编号\",\"fieldType\":\"int\",\"fieldJson\":\"groupId\",\"dataTypeLong\":\"19\",\"comment\":\"帖子分类编号\",\"columnName\":\"group_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Title\",\"fieldDesc\":\"标题\",\"fieldType\":\"string\",\"fieldJson\":\"title\",\"dataTypeLong\":\"80\",\"comment\":\"标题\",\"columnName\":\"title\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"SeoKey\",\"fieldDesc\":\"SEO关键词\",\"fieldType\":\"string\",\"fieldJson\":\"seoKey\",\"dataTypeLong\":\"500\",\"comment\":\"SEO关键词\",\"columnName\":\"seo_key\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"SeoIntroduce\",\"fieldDesc\":\"简介(SEO简介)\",\"fieldType\":\"string\",\"fieldJson\":\"seoIntroduce\",\"dataTypeLong\":\"150\",\"comment\":\"简介(SEO简介)\",\"columnName\":\"seo_introduce\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CoverImage\",\"fieldDesc\":\"封面\",\"fieldType\":\"string\",\"fieldJson\":\"coverImage\",\"dataTypeLong\":\"500\",\"comment\":\"封面\",\"columnName\":\"cover_image\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Source\",\"fieldDesc\":\"来源\",\"fieldType\":\"string\",\"fieldJson\":\"source\",\"dataTypeLong\":\"40\",\"comment\":\"来源\",\"columnName\":\"source\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Time\",\"fieldDesc\":\"发布时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"time\",\"dataTypeLong\":\"\",\"comment\":\"发布时间\",\"columnName\":\"time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"内容\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"\",\"comment\":\"内容\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContentType\",\"fieldDesc\":\"内容类型:0 markdown,1 html\",\"fieldType\":\"int\",\"fieldJson\":\"contentType\",\"dataTypeLong\":\"10\",\"comment\":\"内容类型:0 markdown,1 html\",\"columnName\":\"content_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContentMarkdown\",\"fieldDesc\":\"markdown内容\",\"fieldType\":\"string\",\"fieldJson\":\"contentMarkdown\",\"dataTypeLong\":\"\",\"comment\":\"markdown内容\",\"columnName\":\"content_markdown\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContentHtml\",\"fieldDesc\":\"html内容\",\"fieldType\":\"string\",\"fieldJson\":\"contentHtml\",\"dataTypeLong\":\"\",\"comment\":\"html内容\",\"columnName\":\"content_html\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Video\",\"fieldDesc\":\"视频地址\",\"fieldType\":\"string\",\"fieldJson\":\"video\",\"dataTypeLong\":\"500\",\"comment\":\"视频地址\",\"columnName\":\"video\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Attachment\",\"fieldDesc\":\"附件\",\"fieldType\":\"string\",\"fieldJson\":\"attachment\",\"dataTypeLong\":\"2000\",\"comment\":\"附件\",\"columnName\":\"attachment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Tag\",\"fieldDesc\":\"标签\",\"fieldType\":\"string\",\"fieldJson\":\"tag\",\"dataTypeLong\":\"400\",\"comment\":\"标签\",\"columnName\":\"tag\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Top\",\"fieldDesc\":\"置顶(0否 1是)\",\"fieldType\":\"int\",\"fieldJson\":\"top\",\"dataTypeLong\":\"10\",\"comment\":\"置顶(0否 1是)\",\"columnName\":\"top\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Marrow\",\"fieldDesc\":\"精华(0否 1是)\",\"fieldType\":\"int\",\"fieldJson\":\"marrow\",\"dataTypeLong\":\"10\",\"comment\":\"精华(0否 1是)\",\"columnName\":\"marrow\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CommentId\",\"fieldDesc\":\"问答最佳答案ID\",\"fieldType\":\"int\",\"fieldJson\":\"commentId\",\"dataTypeLong\":\"19\",\"comment\":\"问答最佳答案ID\",\"columnName\":\"comment_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Anonymity\",\"fieldDesc\":\"匿名发布(0否 1是)\",\"fieldType\":\"int\",\"fieldJson\":\"anonymity\",\"dataTypeLong\":\"10\",\"comment\":\"匿名发布(0否 1是)\",\"columnName\":\"anonymity\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"发布者编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"发布者编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ReadNum\",\"fieldDesc\":\"阅读次数\",\"fieldType\":\"int\",\"fieldJson\":\"readNum\",\"dataTypeLong\":\"10\",\"comment\":\"阅读次数\",\"columnName\":\"read_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CommentNum\",\"fieldDesc\":\"评论次数\",\"fieldType\":\"int\",\"fieldJson\":\"commentNum\",\"dataTypeLong\":\"10\",\"comment\":\"评论次数\",\"columnName\":\"comment_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ShareNum\",\"fieldDesc\":\"分享次数\",\"fieldType\":\"int\",\"fieldJson\":\"shareNum\",\"dataTypeLong\":\"10\",\"comment\":\"分享次数\",\"columnName\":\"share_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CollectNum\",\"fieldDesc\":\"收藏次数\",\"fieldType\":\"int\",\"fieldJson\":\"collectNum\",\"dataTypeLong\":\"10\",\"comment\":\"收藏次数\",\"columnName\":\"collect_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"LikeNum\",\"fieldDesc\":\"点赞次数\",\"fieldType\":\"int\",\"fieldJson\":\"likeNum\",\"dataTypeLong\":\"10\",\"comment\":\"点赞次数\",\"columnName\":\"like_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckUser\",\"fieldDesc\":\"审核人\",\"fieldType\":\"int\",\"fieldJson\":\"checkUser\",\"dataTypeLong\":\"19\",\"comment\":\"审核人\",\"columnName\":\"check_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckTime\",\"fieldDesc\":\"审核时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"checkTime\",\"dataTypeLong\":\"\",\"comment\":\"审核时间\",\"columnName\":\"check_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckStatus\",\"fieldDesc\":\"审核状态(0草稿 1未审批 2通过 3拒绝)\",\"fieldType\":\"int\",\"fieldJson\":\"checkStatus\",\"dataTypeLong\":\"10\",\"comment\":\"审核状态(0草稿 1未审批 2通过 3拒绝)\",\"columnName\":\"check_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerComment\",\"fieldDesc\":\"评论权限(0关闭 1开启)\",\"fieldType\":\"int\",\"fieldJson\":\"powerComment\",\"dataTypeLong\":\"10\",\"comment\":\"评论权限(0关闭 1开启)\",\"columnName\":\"power_comment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerCommentAnonymity\",\"fieldDesc\":\"匿名评论(0关闭 1开启)\",\"fieldType\":\"int\",\"fieldJson\":\"powerCommentAnonymity\",\"dataTypeLong\":\"10\",\"comment\":\"匿名评论(0关闭 1开启)\",\"columnName\":\"power_comment_anonymity\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Pay\",\"fieldDesc\":\"付费：0 否，1是\",\"fieldType\":\"int\",\"fieldJson\":\"pay\",\"dataTypeLong\":\"10\",\"comment\":\"付费：0 否，1是\",\"columnName\":\"pay\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PayContent\",\"fieldDesc\":\"内容付费：0 否，1是\",\"fieldType\":\"int\",\"fieldJson\":\"payContent\",\"dataTypeLong\":\"10\",\"comment\":\"内容付费：0 否，1是\",\"columnName\":\"pay_content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PayContentLook\",\"fieldDesc\":\"内容付费可查看百分比例\",\"fieldType\":\"int\",\"fieldJson\":\"payContentLook\",\"dataTypeLong\":\"10\",\"comment\":\"内容付费可查看百分比例\",\"columnName\":\"pay_content_look\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PayAttachment\",\"fieldDesc\":\"附件付费：0 否，1是\",\"fieldType\":\"int\",\"fieldJson\":\"payAttachment\",\"dataTypeLong\":\"10\",\"comment\":\"附件付费：0 否，1是\",\"columnName\":\"pay_attachment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PayCurrency\",\"fieldDesc\":\"付费货币：0 人民、1积分、2代币\",\"fieldType\":\"int\",\"fieldJson\":\"payCurrency\",\"dataTypeLong\":\"10\",\"comment\":\"付费货币：0 人民、1积分、2代币\",\"columnName\":\"pay_currency\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PayNum\",\"fieldDesc\":\"付费金额\",\"fieldType\":\"int\",\"fieldJson\":\"payNum\",\"dataTypeLong\":\"10\",\"comment\":\"付费金额\",\"columnName\":\"pay_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_posts.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_posts.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_posts.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_posts.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_posts.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumPosts.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumPosts\\hkForumPostsForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumPosts\\hkForumPosts.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumPosts{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumPostsRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumPostsApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumPostsRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumPostsService;', 'HkForumPosts', 'hkForumPosts表', '90;91;92;93;94;95;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (2, '2023-01-31 15:22:01.681', '2023-01-31 15:22:01.681', '2023-01-31 16:57:41.618', 'community', '', 'hk_user', '{\"structName\":\"HkUser\",\"tableName\":\"hk_user\",\"packageName\":\"hkUser\",\"humpPackageName\":\"hk_user\",\"abbreviation\":\"hkUser\",\"description\":\"hkUser表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Code\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"string\",\"fieldJson\":\"code\",\"dataTypeLong\":\"12\",\"comment\":\"用户编号\",\"columnName\":\"code\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserType\",\"fieldDesc\":\"用户平台\",\"fieldType\":\"int\",\"fieldJson\":\"userType\",\"dataTypeLong\":\"10\",\"comment\":\"用户平台\",\"columnName\":\"user_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Account\",\"fieldDesc\":\"账号\",\"fieldType\":\"string\",\"fieldJson\":\"account\",\"dataTypeLong\":\"45\",\"comment\":\"账号\",\"columnName\":\"account\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Password\",\"fieldDesc\":\"密码\",\"fieldType\":\"string\",\"fieldJson\":\"password\",\"dataTypeLong\":\"45\",\"comment\":\"密码\",\"columnName\":\"password\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"昵称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"昵称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RealName\",\"fieldDesc\":\"真名\",\"fieldType\":\"string\",\"fieldJson\":\"realName\",\"dataTypeLong\":\"10\",\"comment\":\"真名\",\"columnName\":\"real_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Avatar\",\"fieldDesc\":\"头像\",\"fieldType\":\"string\",\"fieldJson\":\"avatar\",\"dataTypeLong\":\"500\",\"comment\":\"头像\",\"columnName\":\"avatar\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Email\",\"fieldDesc\":\"邮箱\",\"fieldType\":\"string\",\"fieldJson\":\"email\",\"dataTypeLong\":\"45\",\"comment\":\"邮箱\",\"columnName\":\"email\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Phone\",\"fieldDesc\":\"手机\",\"fieldType\":\"string\",\"fieldJson\":\"phone\",\"dataTypeLong\":\"45\",\"comment\":\"手机\",\"columnName\":\"phone\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Birthday\",\"fieldDesc\":\"生日\",\"fieldType\":\"time.Time\",\"fieldJson\":\"birthday\",\"dataTypeLong\":\"\",\"comment\":\"生日\",\"columnName\":\"birthday\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sex\",\"fieldDesc\":\"性别\",\"fieldType\":\"int\",\"fieldJson\":\"sex\",\"dataTypeLong\":\"10\",\"comment\":\"性别\",\"columnName\":\"sex\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleId\",\"fieldDesc\":\"角色id\",\"fieldType\":\"string\",\"fieldJson\":\"roleId\",\"dataTypeLong\":\"1000\",\"comment\":\"角色id\",\"columnName\":\"role_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkUser.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUser\\hkUserForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUser\\hkUser.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkUser{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkUserRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkUserApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkUserRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkUserService;', 'HkUser', 'hkUser表', '96;97;98;99;100;101;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (3, '2023-01-31 16:59:40.679', '2023-01-31 16:59:40.679', '2023-01-31 17:01:19.503', 'community', '', 'hk_user', '{\"structName\":\"User\",\"tableName\":\"hk_user\",\"packageName\":\"app_User\",\"humpPackageName\":\"app__user\",\"abbreviation\":\"hkUser\",\"description\":\"hkUser表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Uuid\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"string\",\"fieldJson\":\"uuid\",\"dataTypeLong\":\"50\",\"comment\":\"用户编号\",\"columnName\":\"uuid\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserType\",\"fieldDesc\":\"用户平台\",\"fieldType\":\"int\",\"fieldJson\":\"userType\",\"dataTypeLong\":\"10\",\"comment\":\"用户平台\",\"columnName\":\"user_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Account\",\"fieldDesc\":\"账号\",\"fieldType\":\"string\",\"fieldJson\":\"account\",\"dataTypeLong\":\"45\",\"comment\":\"账号\",\"columnName\":\"account\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Password\",\"fieldDesc\":\"密码\",\"fieldType\":\"string\",\"fieldJson\":\"password\",\"dataTypeLong\":\"45\",\"comment\":\"密码\",\"columnName\":\"password\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NickName\",\"fieldDesc\":\"昵称\",\"fieldType\":\"string\",\"fieldJson\":\"nickName\",\"dataTypeLong\":\"20\",\"comment\":\"昵称\",\"columnName\":\"nick_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RealName\",\"fieldDesc\":\"真名\",\"fieldType\":\"string\",\"fieldJson\":\"realName\",\"dataTypeLong\":\"10\",\"comment\":\"真名\",\"columnName\":\"real_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"HeaderImg\",\"fieldDesc\":\"头像\",\"fieldType\":\"string\",\"fieldJson\":\"headerImg\",\"dataTypeLong\":\"500\",\"comment\":\"头像\",\"columnName\":\"header_img\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Email\",\"fieldDesc\":\"邮箱\",\"fieldType\":\"string\",\"fieldJson\":\"email\",\"dataTypeLong\":\"45\",\"comment\":\"邮箱\",\"columnName\":\"email\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Phone\",\"fieldDesc\":\"手机\",\"fieldType\":\"string\",\"fieldJson\":\"phone\",\"dataTypeLong\":\"45\",\"comment\":\"手机\",\"columnName\":\"phone\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Birthday\",\"fieldDesc\":\"生日\",\"fieldType\":\"time.Time\",\"fieldJson\":\"birthday\",\"dataTypeLong\":\"\",\"comment\":\"生日\",\"columnName\":\"birthday\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sex\",\"fieldDesc\":\"性别\",\"fieldType\":\"int\",\"fieldJson\":\"sex\",\"dataTypeLong\":\"10\",\"comment\":\"性别\",\"columnName\":\"sex\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleId\",\"fieldDesc\":\"用户角色ID\",\"fieldType\":\"int\",\"fieldJson\":\"roleId\",\"dataTypeLong\":\"20\",\"comment\":\"用户角色ID\",\"columnName\":\"role_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态(用户是否被冻结) 1正常 2冻结\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态(用户是否被冻结) 1正常 2冻结\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\app__user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\app__user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\app__user.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\app__user.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\app__user.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\app_User.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\app_User\\app_UserForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\app_User\\app_User.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.User{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitUserRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@UserApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@UserRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@UserService;', 'User', 'hkUser表', '102;103;104;105;106;107;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (4, '2023-01-31 17:10:03.369', '2023-01-31 17:10:03.369', NULL, 'community', '', 'hk_user', '{\"structName\":\"User\",\"tableName\":\"hk_user\",\"packageName\":\"app_user\",\"humpPackageName\":\"app_user\",\"abbreviation\":\"user\",\"description\":\"hkUser表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Uuid\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"string\",\"fieldJson\":\"uuid\",\"dataTypeLong\":\"50\",\"comment\":\"用户编号\",\"columnName\":\"uuid\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserType\",\"fieldDesc\":\"用户平台\",\"fieldType\":\"int\",\"fieldJson\":\"userType\",\"dataTypeLong\":\"10\",\"comment\":\"用户平台\",\"columnName\":\"user_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Account\",\"fieldDesc\":\"账号\",\"fieldType\":\"string\",\"fieldJson\":\"account\",\"dataTypeLong\":\"45\",\"comment\":\"账号\",\"columnName\":\"account\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Password\",\"fieldDesc\":\"密码\",\"fieldType\":\"string\",\"fieldJson\":\"password\",\"dataTypeLong\":\"45\",\"comment\":\"密码\",\"columnName\":\"password\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NickName\",\"fieldDesc\":\"昵称\",\"fieldType\":\"string\",\"fieldJson\":\"nickName\",\"dataTypeLong\":\"20\",\"comment\":\"昵称\",\"columnName\":\"nick_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RealName\",\"fieldDesc\":\"真名\",\"fieldType\":\"string\",\"fieldJson\":\"realName\",\"dataTypeLong\":\"10\",\"comment\":\"真名\",\"columnName\":\"real_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"HeaderImg\",\"fieldDesc\":\"头像\",\"fieldType\":\"string\",\"fieldJson\":\"headerImg\",\"dataTypeLong\":\"500\",\"comment\":\"头像\",\"columnName\":\"header_img\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Email\",\"fieldDesc\":\"邮箱\",\"fieldType\":\"string\",\"fieldJson\":\"email\",\"dataTypeLong\":\"45\",\"comment\":\"邮箱\",\"columnName\":\"email\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Phone\",\"fieldDesc\":\"手机\",\"fieldType\":\"string\",\"fieldJson\":\"phone\",\"dataTypeLong\":\"45\",\"comment\":\"手机\",\"columnName\":\"phone\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Birthday\",\"fieldDesc\":\"生日\",\"fieldType\":\"time.Time\",\"fieldJson\":\"birthday\",\"dataTypeLong\":\"\",\"comment\":\"生日\",\"columnName\":\"birthday\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sex\",\"fieldDesc\":\"性别\",\"fieldType\":\"int\",\"fieldJson\":\"sex\",\"dataTypeLong\":\"10\",\"comment\":\"性别\",\"columnName\":\"sex\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleId\",\"fieldDesc\":\"用户角色ID\",\"fieldType\":\"int\",\"fieldJson\":\"roleId\",\"dataTypeLong\":\"20\",\"comment\":\"用户角色ID\",\"columnName\":\"role_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态(用户是否被冻结) 1正常 2冻结\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态(用户是否被冻结) 1正常 2冻结\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\app_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\app_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\app_user.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\app_user.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\app_user.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\app_user.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\app_user\\app_userForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\app_user\\app_user.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.User{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitUserRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@UserApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@UserRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@UserService;', 'User', 'hkUser表', '', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (5, '2023-02-01 08:19:17.697', '2023-02-01 08:19:17.697', NULL, 'apply', '', 'hk_apply', '{\"structName\":\"HkApply\",\"tableName\":\"hk_apply\",\"packageName\":\"hkApply\",\"humpPackageName\":\"hk_apply\",\"abbreviation\":\"hkApply\",\"description\":\"hkApply表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"OwerType\",\"fieldDesc\":\"拥有者：0平台、1圈子、2个人\",\"fieldType\":\"int\",\"fieldJson\":\"owerType\",\"dataTypeLong\":\"10\",\"comment\":\"拥有者：0平台、1圈子、2个人\",\"columnName\":\"ower_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户_编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户_编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"80\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Icon\",\"fieldDesc\":\"图标\",\"fieldType\":\"string\",\"fieldJson\":\"icon\",\"dataTypeLong\":\"256\",\"comment\":\"图标\",\"columnName\":\"icon\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Type\",\"fieldDesc\":\"类型(0小程序、1第三方链接)\",\"fieldType\":\"int\",\"fieldJson\":\"type\",\"dataTypeLong\":\"10\",\"comment\":\"类型(0小程序、1第三方链接)\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MiniProgramId\",\"fieldDesc\":\"小程序id\",\"fieldType\":\"int\",\"fieldJson\":\"miniProgramId\",\"dataTypeLong\":\"19\",\"comment\":\"小程序id\",\"columnName\":\"mini_program_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyAddress\",\"fieldDesc\":\"访问地址\",\"fieldType\":\"string\",\"fieldJson\":\"applyAddress\",\"dataTypeLong\":\"256\",\"comment\":\"访问地址\",\"columnName\":\"apply_address\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyParameters\",\"fieldDesc\":\"访问参数\",\"fieldType\":\"string\",\"fieldJson\":\"applyParameters\",\"dataTypeLong\":\"256\",\"comment\":\"访问参数\",\"columnName\":\"apply_parameters\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"apply\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\hk_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\hk_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\request\\hk_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\hk_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\hk_apply.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkApply.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkApply\\hkApplyForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkApply\\hkApply.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@apply.HkApply{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@applyRouter.InitHkApplyRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\enter.go@ApiGroup@HkApplyApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\enter.go@RouterGroup@HkApplyRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\enter.go@ServiceGroup@HkApplyService;', 'HkApply', 'hkApply表', '113;114;115;116;117;118;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (6, '2023-02-01 08:20:00.069', '2023-02-01 08:20:00.069', NULL, 'apply', '', 'hk_circle_apply', '{\"structName\":\"HkCircleApply\",\"tableName\":\"hk_circle_apply\",\"packageName\":\"hkCircleApply\",\"humpPackageName\":\"hk_circle_apply\",\"abbreviation\":\"hkCircleApply\",\"description\":\"hkCircleApply表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyGroupId\",\"fieldDesc\":\"应用分组_编号\",\"fieldType\":\"int\",\"fieldJson\":\"applyGroupId\",\"dataTypeLong\":\"19\",\"comment\":\"应用分组_编号\",\"columnName\":\"apply_group_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ShowName\",\"fieldDesc\":\"名称别名\",\"fieldType\":\"string\",\"fieldJson\":\"showName\",\"dataTypeLong\":\"80\",\"comment\":\"名称别名\",\"columnName\":\"show_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyId\",\"fieldDesc\":\"应用_编号\",\"fieldType\":\"int\",\"fieldJson\":\"applyId\",\"dataTypeLong\":\"19\",\"comment\":\"应用_编号\",\"columnName\":\"apply_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyParameters\",\"fieldDesc\":\"访问参数\",\"fieldType\":\"string\",\"fieldJson\":\"applyParameters\",\"dataTypeLong\":\"256\",\"comment\":\"访问参数\",\"columnName\":\"apply_parameters\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"apply\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\hk_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\hk_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\request\\hk_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\hk_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\hk_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleApply.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleApply\\hkCircleApplyForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleApply\\hkCircleApply.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@apply.HkCircleApply{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@applyRouter.InitHkCircleApplyRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\enter.go@ApiGroup@HkCircleApplyApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\enter.go@RouterGroup@HkCircleApplyRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\enter.go@ServiceGroup@HkCircleApplyService;', 'HkCircleApply', 'hkCircleApply表', '119;120;121;122;123;124;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (7, '2023-02-01 08:20:26.122', '2023-02-01 08:20:26.122', NULL, 'apply', '', 'hk_circle_apply_group', '{\"structName\":\"HkCircleApplyGroup\",\"tableName\":\"hk_circle_apply_group\",\"packageName\":\"hkCircleApplyGroup\",\"humpPackageName\":\"hk_circle_apply_group\",\"abbreviation\":\"hkCircleApplyGroup\",\"description\":\"hkCircleApplyGroup表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Icon\",\"fieldDesc\":\"图标\",\"fieldType\":\"string\",\"fieldJson\":\"icon\",\"dataTypeLong\":\"256\",\"comment\":\"图标\",\"columnName\":\"icon\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ParentId\",\"fieldDesc\":\"父节点编号\",\"fieldType\":\"int\",\"fieldJson\":\"parentId\",\"dataTypeLong\":\"19\",\"comment\":\"父节点编号\",\"columnName\":\"parent_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"apply\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\hk_circle_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\hk_circle_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\request\\hk_circle_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\hk_circle_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\hk_circle_apply_group.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleApplyGroup.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleApplyGroup\\hkCircleApplyGroupForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleApplyGroup\\hkCircleApplyGroup.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@apply.HkCircleApplyGroup{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@applyRouter.InitHkCircleApplyGroupRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\enter.go@ApiGroup@HkCircleApplyGroupApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\enter.go@RouterGroup@HkCircleApplyGroupRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\enter.go@ServiceGroup@HkCircleApplyGroupService;', 'HkCircleApplyGroup', 'hkCircleApplyGroup表', '125;126;127;128;129;130;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (8, '2023-02-01 08:21:21.578', '2023-02-01 08:21:21.578', NULL, 'apply', '', 'hk_plat_apply', '{\"structName\":\"HkPlatApply\",\"tableName\":\"hk_plat_apply\",\"packageName\":\"hkPlatApply\",\"humpPackageName\":\"hk_plat_apply\",\"abbreviation\":\"hkPlatApply\",\"description\":\"hkPlatApply表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyGroupId\",\"fieldDesc\":\"应用分组_编号\",\"fieldType\":\"int\",\"fieldJson\":\"applyGroupId\",\"dataTypeLong\":\"19\",\"comment\":\"应用分组_编号\",\"columnName\":\"apply_group_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ShowName\",\"fieldDesc\":\"名称别名\",\"fieldType\":\"string\",\"fieldJson\":\"showName\",\"dataTypeLong\":\"80\",\"comment\":\"名称别名\",\"columnName\":\"show_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyId\",\"fieldDesc\":\"应用_编号\",\"fieldType\":\"int\",\"fieldJson\":\"applyId\",\"dataTypeLong\":\"19\",\"comment\":\"应用_编号\",\"columnName\":\"apply_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyParameters\",\"fieldDesc\":\"访问参数\",\"fieldType\":\"string\",\"fieldJson\":\"applyParameters\",\"dataTypeLong\":\"256\",\"comment\":\"访问参数\",\"columnName\":\"apply_parameters\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"apply\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\hk_plat_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\hk_plat_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\request\\hk_plat_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\hk_plat_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\hk_plat_apply.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkPlatApply.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkPlatApply\\hkPlatApplyForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkPlatApply\\hkPlatApply.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@apply.HkPlatApply{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@applyRouter.InitHkPlatApplyRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\enter.go@ApiGroup@HkPlatApplyApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\enter.go@RouterGroup@HkPlatApplyRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\enter.go@ServiceGroup@HkPlatApplyService;', 'HkPlatApply', 'hkPlatApply表', '131;132;133;134;135;136;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (9, '2023-02-01 08:21:43.613', '2023-02-01 08:21:43.613', NULL, 'apply', '', 'hk_plat_apply_group', '{\"structName\":\"HkPlatApplyGroup\",\"tableName\":\"hk_plat_apply_group\",\"packageName\":\"hkPlatApplyGroup\",\"humpPackageName\":\"hk_plat_apply_group\",\"abbreviation\":\"hkPlatApplyGroup\",\"description\":\"hkPlatApplyGroup表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Icon\",\"fieldDesc\":\"图标\",\"fieldType\":\"string\",\"fieldJson\":\"icon\",\"dataTypeLong\":\"256\",\"comment\":\"图标\",\"columnName\":\"icon\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ParentId\",\"fieldDesc\":\"父节点编号\",\"fieldType\":\"int\",\"fieldJson\":\"parentId\",\"dataTypeLong\":\"19\",\"comment\":\"父节点编号\",\"columnName\":\"parent_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"apply\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\hk_plat_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\hk_plat_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\apply\\request\\hk_plat_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\hk_plat_apply_group.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\hk_plat_apply_group.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkPlatApplyGroup.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkPlatApplyGroup\\hkPlatApplyGroupForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkPlatApplyGroup\\hkPlatApplyGroup.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@apply.HkPlatApplyGroup{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@applyRouter.InitHkPlatApplyGroupRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\apply\\enter.go@ApiGroup@HkPlatApplyGroupApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\apply\\enter.go@RouterGroup@HkPlatApplyGroupRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\apply\\enter.go@ServiceGroup@HkPlatApplyGroupService;', 'HkPlatApplyGroup', 'hkPlatApplyGroup表', '137;138;139;140;141;142;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (10, '2023-02-01 08:35:38.161', '2023-02-01 08:35:38.161', NULL, 'general', '', 'hk_bug_report', '{\"structName\":\"HkBugReport\",\"tableName\":\"hk_bug_report\",\"packageName\":\"hkBugReport\",\"humpPackageName\":\"hk_bug_report\",\"abbreviation\":\"hkBugReport\",\"description\":\"hkBugReport表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Title\",\"fieldDesc\":\"标题\",\"fieldType\":\"string\",\"fieldJson\":\"title\",\"dataTypeLong\":\"80\",\"comment\":\"标题\",\"columnName\":\"title\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"操作步骤\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"400\",\"comment\":\"操作步骤\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContentAttachment\",\"fieldDesc\":\"操作步骤附件\",\"fieldType\":\"string\",\"fieldJson\":\"contentAttachment\",\"dataTypeLong\":\"400\",\"comment\":\"操作步骤附件\",\"columnName\":\"content_attachment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ExpectedResult\",\"fieldDesc\":\"预期结果\",\"fieldType\":\"string\",\"fieldJson\":\"expectedResult\",\"dataTypeLong\":\"400\",\"comment\":\"预期结果\",\"columnName\":\"expected_result\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ActualResult\",\"fieldDesc\":\"实际结果\",\"fieldType\":\"string\",\"fieldJson\":\"actualResult\",\"dataTypeLong\":\"400\",\"comment\":\"实际结果\",\"columnName\":\"actual_result\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ActualResultAttachment\",\"fieldDesc\":\"实际结果附件\",\"fieldType\":\"string\",\"fieldJson\":\"actualResultAttachment\",\"dataTypeLong\":\"2000\",\"comment\":\"实际结果附件\",\"columnName\":\"actual_result_attachment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"OtherInfo\",\"fieldDesc\":\"其他信息\",\"fieldType\":\"string\",\"fieldJson\":\"otherInfo\",\"dataTypeLong\":\"400\",\"comment\":\"其他信息\",\"columnName\":\"other_info\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"general\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\hk_bug_report.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\hk_bug_report.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\request\\hk_bug_report.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\hk_bug_report.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\hk_bug_report.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkBugReport.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkBugReport\\hkBugReportForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkBugReport\\hkBugReport.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@general.HkBugReport{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@generalRouter.InitHkBugReportRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\enter.go@ApiGroup@HkBugReportApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\enter.go@RouterGroup@HkBugReportRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\enter.go@ServiceGroup@HkBugReportService;', 'HkBugReport', 'hkBugReport表', '143;144;145;146;147;148;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (11, '2023-02-01 08:36:46.932', '2023-02-01 08:36:46.932', NULL, 'community', '', 'hk_circle', '{\"structName\":\"HkCircle\",\"tableName\":\"hk_circle\",\"packageName\":\"hkCircle\",\"humpPackageName\":\"hk_circle\",\"abbreviation\":\"hkCircle\",\"description\":\"hkCircle表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Type\",\"fieldDesc\":\"类型：0官方圈子、1用户圈子、2小区\",\"fieldType\":\"int\",\"fieldJson\":\"type\",\"dataTypeLong\":\"10\",\"comment\":\"类型：0官方圈子、1用户圈子、2小区\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"圈子名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"圈子名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Logo\",\"fieldDesc\":\"圈子Logo\",\"fieldType\":\"string\",\"fieldJson\":\"logo\",\"dataTypeLong\":\"500\",\"comment\":\"圈子Logo\",\"columnName\":\"logo\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleClassifyId\",\"fieldDesc\":\"圈子分类_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleClassifyId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子分类_编号\",\"columnName\":\"circle_classify_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Slogan\",\"fieldDesc\":\"圈子标语\",\"fieldType\":\"string\",\"fieldJson\":\"slogan\",\"dataTypeLong\":\"20\",\"comment\":\"圈子标语\",\"columnName\":\"slogan\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Des\",\"fieldDesc\":\"圈子简介\",\"fieldType\":\"string\",\"fieldJson\":\"des\",\"dataTypeLong\":\"1000\",\"comment\":\"圈子简介\",\"columnName\":\"des\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Protocol\",\"fieldDesc\":\"圈子规约\",\"fieldType\":\"string\",\"fieldJson\":\"protocol\",\"dataTypeLong\":\"1000\",\"comment\":\"圈子规约\",\"columnName\":\"protocol\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"BackImage\",\"fieldDesc\":\"圈子背景图\",\"fieldType\":\"string\",\"fieldJson\":\"backImage\",\"dataTypeLong\":\"500\",\"comment\":\"圈子背景图\",\"columnName\":\"back_image\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"SupportCategory\",\"fieldDesc\":\"支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议\",\"fieldType\":\"string\",\"fieldJson\":\"supportCategory\",\"dataTypeLong\":\"500\",\"comment\":\"支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议\",\"columnName\":\"support_category\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Pay\",\"fieldDesc\":\"付费：0 否，1是\",\"fieldType\":\"int\",\"fieldJson\":\"pay\",\"dataTypeLong\":\"10\",\"comment\":\"付费：0 否，1是\",\"columnName\":\"pay\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Process\",\"fieldDesc\":\"是否开启版块内容人工审核：0 否，1是\",\"fieldType\":\"int\",\"fieldJson\":\"process\",\"dataTypeLong\":\"10\",\"comment\":\"是否开启版块内容人工审核：0 否，1是\",\"columnName\":\"process\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Property\",\"fieldDesc\":\"圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）\",\"fieldType\":\"int\",\"fieldJson\":\"property\",\"dataTypeLong\":\"10\",\"comment\":\"圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）\",\"columnName\":\"property\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"View\",\"fieldDesc\":\"板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到\",\"fieldType\":\"int\",\"fieldJson\":\"view\",\"dataTypeLong\":\"10\",\"comment\":\"板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到\",\"columnName\":\"view\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerAdd\",\"fieldDesc\":\"圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户\",\"fieldType\":\"int\",\"fieldJson\":\"powerAdd\",\"dataTypeLong\":\"10\",\"comment\":\"圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户\",\"columnName\":\"power_add\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerView\",\"fieldDesc\":\"圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组\",\"fieldType\":\"int\",\"fieldJson\":\"powerView\",\"dataTypeLong\":\"10\",\"comment\":\"圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组\",\"columnName\":\"power_view\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerPublish\",\"fieldDesc\":\"圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组\",\"fieldType\":\"int\",\"fieldJson\":\"powerPublish\",\"dataTypeLong\":\"10\",\"comment\":\"圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组\",\"columnName\":\"power_publish\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerComment\",\"fieldDesc\":\"圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组\",\"fieldType\":\"int\",\"fieldJson\":\"powerComment\",\"dataTypeLong\":\"10\",\"comment\":\"圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组\",\"columnName\":\"power_comment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerAddUser\",\"fieldDesc\":\"圈子加入权限_指定部门和成员(json数组)\",\"fieldType\":\"string\",\"fieldJson\":\"powerAddUser\",\"dataTypeLong\":\"500\",\"comment\":\"圈子加入权限_指定部门和成员(json数组)\",\"columnName\":\"power_add_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerViewUser\",\"fieldDesc\":\"圈子内浏览权限_指定部门和用户(json数组)\",\"fieldType\":\"string\",\"fieldJson\":\"powerViewUser\",\"dataTypeLong\":\"500\",\"comment\":\"圈子内浏览权限_指定部门和用户(json数组)\",\"columnName\":\"power_view_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerPublishUser\",\"fieldDesc\":\"圈子内发布权限_指定部门和用户(json数组)\",\"fieldType\":\"string\",\"fieldJson\":\"powerPublishUser\",\"dataTypeLong\":\"500\",\"comment\":\"圈子内发布权限_指定部门和用户(json数组)\",\"columnName\":\"power_publish_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PowerCommentUser\",\"fieldDesc\":\"圈子内评论权限_指定部门和用户(json数组)\",\"fieldType\":\"string\",\"fieldJson\":\"powerCommentUser\",\"dataTypeLong\":\"500\",\"comment\":\"圈子内评论权限_指定部门和用户(json数组)\",\"columnName\":\"power_comment_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NoLimitUserGroup\",\"fieldDesc\":\"不受限用户组(json数组)\",\"fieldType\":\"string\",\"fieldJson\":\"noLimitUserGroup\",\"dataTypeLong\":\"500\",\"comment\":\"不受限用户组(json数组)\",\"columnName\":\"no_limit_user_group\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NewUserFocus\",\"fieldDesc\":\"新注册用户默认关注：0 否，1是\",\"fieldType\":\"int\",\"fieldJson\":\"newUserFocus\",\"dataTypeLong\":\"10\",\"comment\":\"新注册用户默认关注：0 否，1是\",\"columnName\":\"new_user_focus\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_circle.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_circle.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_circle.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_circle.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_circle.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircle.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircle\\hkCircleForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircle\\hkCircle.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCircle{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCircleRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCircleApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCircleRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCircleService;', 'HkCircle', 'hkCircle表', '149;150;151;152;153;154;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (12, '2023-02-01 08:38:39.858', '2023-02-01 08:38:39.858', NULL, 'community', '', 'hk_circle_add_request', '{\"structName\":\"HkCircleAddRequest\",\"tableName\":\"hk_circle_add_request\",\"packageName\":\"hkCircleAddRequest\",\"humpPackageName\":\"hk_circle_add_request\",\"abbreviation\":\"hkCircleAddRequest\",\"description\":\"hkCircleAddRequest表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Reason\",\"fieldDesc\":\"申请理由\",\"fieldType\":\"string\",\"fieldJson\":\"reason\",\"dataTypeLong\":\"500\",\"comment\":\"申请理由\",\"columnName\":\"reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckUser\",\"fieldDesc\":\"审核人\",\"fieldType\":\"int\",\"fieldJson\":\"checkUser\",\"dataTypeLong\":\"19\",\"comment\":\"审核人\",\"columnName\":\"check_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckTime\",\"fieldDesc\":\"审核时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"checkTime\",\"dataTypeLong\":\"\",\"comment\":\"审核时间\",\"columnName\":\"check_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckStatus\",\"fieldDesc\":\"审核状态：0 未处理 1 通过，2驳回\",\"fieldType\":\"int\",\"fieldJson\":\"checkStatus\",\"dataTypeLong\":\"10\",\"comment\":\"审核状态：0 未处理 1 通过，2驳回\",\"columnName\":\"check_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_circle_add_request.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_circle_add_request.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_circle_add_request.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_circle_add_request.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_circle_add_request.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleAddRequest.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleAddRequest\\hkCircleAddRequestForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleAddRequest\\hkCircleAddRequest.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCircleAddRequest{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCircleAddRequestRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCircleAddRequestApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCircleAddRequestRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCircleAddRequestService;', 'HkCircleAddRequest', 'hkCircleAddRequest表', '155;156;157;158;159;160;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (13, '2023-02-01 08:39:56.064', '2023-02-01 08:39:56.064', NULL, 'community', '', 'hk_circle_classify', '{\"structName\":\"HkCircleClassify\",\"tableName\":\"hk_circle_classify\",\"packageName\":\"hkCircleClassify\",\"humpPackageName\":\"hk_circle_classify\",\"abbreviation\":\"hkCircleClassify\",\"description\":\"hkCircleClassify表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ParentId\",\"fieldDesc\":\"父节点编号\",\"fieldType\":\"int\",\"fieldJson\":\"parentId\",\"dataTypeLong\":\"19\",\"comment\":\"父节点编号\",\"columnName\":\"parent_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Des\",\"fieldDesc\":\"描述\",\"fieldType\":\"string\",\"fieldJson\":\"des\",\"dataTypeLong\":\"20\",\"comment\":\"描述\",\"columnName\":\"des\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Property\",\"fieldDesc\":\"属性：0公开 ，1受限\",\"fieldType\":\"int\",\"fieldJson\":\"property\",\"dataTypeLong\":\"10\",\"comment\":\"属性：0公开 ，1受限\",\"columnName\":\"property\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_circle_classify.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_circle_classify.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_circle_classify.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_circle_classify.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_circle_classify.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleClassify.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleClassify\\hkCircleClassifyForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleClassify\\hkCircleClassify.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCircleClassify{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCircleClassifyRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCircleClassifyApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCircleClassifyRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCircleClassifyService;', 'HkCircleClassify', 'hkCircleClassify表', '161;162;163;164;165;166;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (14, '2023-02-01 08:40:26.295', '2023-02-01 08:40:26.295', NULL, 'community', '', 'hk_circle_relation', '{\"structName\":\"HkCircleRelation\",\"tableName\":\"hk_circle_relation\",\"packageName\":\"hkCircleRelation\",\"humpPackageName\":\"hk_circle_relation\",\"abbreviation\":\"hkCircleRelation\",\"description\":\"hkCircleRelation表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RelationType\",\"fieldDesc\":\"关系类型：0父子节点 1关注\",\"fieldType\":\"int\",\"fieldJson\":\"relationType\",\"dataTypeLong\":\"10\",\"comment\":\"关系类型：0父子节点 1关注\",\"columnName\":\"relation_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"OtherCircleId\",\"fieldDesc\":\"关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）\",\"fieldType\":\"int\",\"fieldJson\":\"otherCircleId\",\"dataTypeLong\":\"19\",\"comment\":\"关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）\",\"columnName\":\"other_circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_circle_relation.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_circle_relation.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_circle_relation.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_circle_relation.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_circle_relation.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleRelation.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleRelation\\hkCircleRelationForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleRelation\\hkCircleRelation.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCircleRelation{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCircleRelationRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCircleRelationApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCircleRelationRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCircleRelationService;', 'HkCircleRelation', 'hkCircleRelation表', '167;168;169;170;171;172;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (15, '2023-02-01 08:40:59.317', '2023-02-01 08:40:59.317', NULL, 'community', '', 'hk_circle_request', '{\"structName\":\"HkCircleRequest\",\"tableName\":\"hk_circle_request\",\"packageName\":\"hkCircleRequest\",\"humpPackageName\":\"hk_circle_request\",\"abbreviation\":\"hkCircleRequest\",\"description\":\"hkCircleRequest表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Type\",\"fieldDesc\":\"类型：0官方圈子 ，1用户圈子\",\"fieldType\":\"int\",\"fieldJson\":\"type\",\"dataTypeLong\":\"10\",\"comment\":\"类型：0官方圈子 ，1用户圈子\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"圈子名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"圈子名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Logo\",\"fieldDesc\":\"圈子Logo\",\"fieldType\":\"string\",\"fieldJson\":\"logo\",\"dataTypeLong\":\"500\",\"comment\":\"圈子Logo\",\"columnName\":\"logo\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleClassifyId\",\"fieldDesc\":\"圈子分类_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleClassifyId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子分类_编号\",\"columnName\":\"circle_classify_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Slogan\",\"fieldDesc\":\"圈子标语\",\"fieldType\":\"string\",\"fieldJson\":\"slogan\",\"dataTypeLong\":\"20\",\"comment\":\"圈子标语\",\"columnName\":\"slogan\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Des\",\"fieldDesc\":\"圈子简介\",\"fieldType\":\"string\",\"fieldJson\":\"des\",\"dataTypeLong\":\"1000\",\"comment\":\"圈子简介\",\"columnName\":\"des\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Protocol\",\"fieldDesc\":\"圈子规约\",\"fieldType\":\"string\",\"fieldJson\":\"protocol\",\"dataTypeLong\":\"1000\",\"comment\":\"圈子规约\",\"columnName\":\"protocol\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"BackImage\",\"fieldDesc\":\"圈子背景图\",\"fieldType\":\"string\",\"fieldJson\":\"backImage\",\"dataTypeLong\":\"500\",\"comment\":\"圈子背景图\",\"columnName\":\"back_image\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckStatus\",\"fieldDesc\":\"审核状态：0 未处理 1 通过，2驳回\",\"fieldType\":\"int\",\"fieldJson\":\"checkStatus\",\"dataTypeLong\":\"10\",\"comment\":\"审核状态：0 未处理 1 通过，2驳回\",\"columnName\":\"check_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_circle_request.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_circle_request.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_circle_request.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_circle_request.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_circle_request.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleRequest.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleRequest\\hkCircleRequestForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleRequest\\hkCircleRequest.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCircleRequest{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCircleRequestRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCircleRequestApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCircleRequestRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCircleRequestService;', 'HkCircleRequest', 'hkCircleRequest表', '173;174;175;176;177;178;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (16, '2023-02-01 08:41:44.425', '2023-02-01 08:41:44.425', NULL, 'community', '', 'hk_circle_user', '{\"structName\":\"HkCircleUser\",\"tableName\":\"hk_circle_user\",\"packageName\":\"hkCircleUser\",\"humpPackageName\":\"hk_circle_user\",\"abbreviation\":\"hkCircleUser\",\"description\":\"hkCircleUser表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户ID\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户ID\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"用户的圈子排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"用户的圈子排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_circle_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_circle_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_circle_user.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_circle_user.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_circle_user.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCircleUser.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleUser\\hkCircleUserForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCircleUser\\hkCircleUser.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCircleUser{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCircleUserRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCircleUserApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCircleUserRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCircleUserService;', 'HkCircleUser', 'hkCircleUser表', '179;180;181;182;183;184;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (17, '2023-02-01 08:42:13.484', '2023-02-01 08:42:13.484', NULL, 'community', '', 'hk_comment_thumbs_up', '{\"structName\":\"HkCommentThumbsUp\",\"tableName\":\"hk_comment_thumbs_up\",\"packageName\":\"hkCommentThumbsUp\",\"humpPackageName\":\"hk_comment_thumbs_up\",\"abbreviation\":\"hkCommentThumbsUp\",\"description\":\"hkCommentThumbsUp表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户id\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Time\",\"fieldDesc\":\"点赞时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"time\",\"dataTypeLong\":\"\",\"comment\":\"点赞时间\",\"columnName\":\"time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CommentId\",\"fieldDesc\":\"评论编号\",\"fieldType\":\"int\",\"fieldJson\":\"commentId\",\"dataTypeLong\":\"19\",\"comment\":\"评论编号\",\"columnName\":\"comment_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_comment_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_comment_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_comment_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_comment_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_comment_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkCommentThumbsUp.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCommentThumbsUp\\hkCommentThumbsUpForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkCommentThumbsUp\\hkCommentThumbsUp.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkCommentThumbsUp{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkCommentThumbsUpRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkCommentThumbsUpApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkCommentThumbsUpRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkCommentThumbsUpService;', 'HkCommentThumbsUp', 'hkCommentThumbsUp表', '185;186;187;188;189;190;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (18, '2023-02-01 08:43:06.282', '2023-02-01 08:43:06.282', NULL, 'community', '', 'hk_forbidden_speak', '{\"structName\":\"HkForbiddenSpeak\",\"tableName\":\"hk_forbidden_speak\",\"packageName\":\"hkForbiddenSpeak\",\"humpPackageName\":\"hk_forbidden_speak\",\"abbreviation\":\"hkForbiddenSpeak\",\"description\":\"hkForbiddenSpeak表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"DurationId\",\"fieldDesc\":\"禁言时长_编号\",\"fieldType\":\"int\",\"fieldJson\":\"durationId\",\"dataTypeLong\":\"19\",\"comment\":\"禁言时长_编号\",\"columnName\":\"duration_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ReasonId\",\"fieldDesc\":\"禁言原因_编号\",\"fieldType\":\"int\",\"fieldJson\":\"reasonId\",\"dataTypeLong\":\"19\",\"comment\":\"禁言原因_编号\",\"columnName\":\"reason_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CurStatus\",\"fieldDesc\":\"当前状态：0 未解锁、1已解锁\",\"fieldType\":\"int\",\"fieldJson\":\"curStatus\",\"dataTypeLong\":\"10\",\"comment\":\"当前状态：0 未解锁、1已解锁\",\"columnName\":\"cur_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UnlockUserId\",\"fieldDesc\":\"解锁用户_编号\",\"fieldType\":\"int\",\"fieldJson\":\"unlockUserId\",\"dataTypeLong\":\"19\",\"comment\":\"解锁用户_编号\",\"columnName\":\"unlock_user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forbidden_speak.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forbidden_speak.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forbidden_speak.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forbidden_speak.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forbidden_speak.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForbiddenSpeak.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForbiddenSpeak\\hkForbiddenSpeakForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForbiddenSpeak\\hkForbiddenSpeak.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForbiddenSpeak{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForbiddenSpeakRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForbiddenSpeakApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForbiddenSpeakRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForbiddenSpeakService;', 'HkForbiddenSpeak', 'hkForbiddenSpeak表', '191;192;193;194;195;196;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (19, '2023-02-01 08:43:30.723', '2023-02-01 08:43:30.723', NULL, 'community', '', 'hk_forbidden_speak_duration', '{\"structName\":\"HkForbiddenSpeakDuration\",\"tableName\":\"hk_forbidden_speak_duration\",\"packageName\":\"hkForbiddenSpeakDuration\",\"humpPackageName\":\"hk_forbidden_speak_duration\",\"abbreviation\":\"hkForbiddenSpeakDuration\",\"description\":\"hkForbiddenSpeakDuration表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Count\",\"fieldDesc\":\"时长\",\"fieldType\":\"int\",\"fieldJson\":\"count\",\"dataTypeLong\":\"10\",\"comment\":\"时长\",\"columnName\":\"count\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Type\",\"fieldDesc\":\"时间类型：0 小时、1天\",\"fieldType\":\"int\",\"fieldJson\":\"type\",\"dataTypeLong\":\"10\",\"comment\":\"时间类型：0 小时、1天\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forbidden_speak_duration.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forbidden_speak_duration.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forbidden_speak_duration.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forbidden_speak_duration.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forbidden_speak_duration.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForbiddenSpeakDuration.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForbiddenSpeakDuration\\hkForbiddenSpeakDurationForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForbiddenSpeakDuration\\hkForbiddenSpeakDuration.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForbiddenSpeakDuration{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForbiddenSpeakDurationRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForbiddenSpeakDurationApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForbiddenSpeakDurationRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForbiddenSpeakDurationService;', 'HkForbiddenSpeakDuration', 'hkForbiddenSpeakDuration表', '197;198;199;200;201;202;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (20, '2023-02-01 08:43:52.234', '2023-02-01 08:43:52.234', NULL, 'community', '', 'hk_forbidden_speak_reason', '{\"structName\":\"HkForbiddenSpeakReason\",\"tableName\":\"hk_forbidden_speak_reason\",\"packageName\":\"hkForbiddenSpeakReason\",\"humpPackageName\":\"hk_forbidden_speak_reason\",\"abbreviation\":\"hkForbiddenSpeakReason\",\"description\":\"hkForbiddenSpeakReason表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Reason\",\"fieldDesc\":\"禁言理由\",\"fieldType\":\"string\",\"fieldJson\":\"reason\",\"dataTypeLong\":\"20\",\"comment\":\"禁言理由\",\"columnName\":\"reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forbidden_speak_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forbidden_speak_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forbidden_speak_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forbidden_speak_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forbidden_speak_reason.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForbiddenSpeakReason.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForbiddenSpeakReason\\hkForbiddenSpeakReasonForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForbiddenSpeakReason\\hkForbiddenSpeakReason.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForbiddenSpeakReason{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForbiddenSpeakReasonRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForbiddenSpeakReasonApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForbiddenSpeakReasonRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForbiddenSpeakReasonService;', 'HkForbiddenSpeakReason', 'hkForbiddenSpeakReason表', '203;204;205;206;207;208;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (21, '2023-02-01 08:45:10.277', '2023-02-01 08:45:10.277', NULL, 'community', '', 'hk_forum_comment', '{\"structName\":\"HkForumComment\",\"tableName\":\"hk_forum_comment\",\"packageName\":\"hkForumComment\",\"humpPackageName\":\"hk_forum_comment\",\"abbreviation\":\"hkForumComment\",\"description\":\"hkForumComment表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PostsId\",\"fieldDesc\":\"帖子编号\",\"fieldType\":\"int\",\"fieldJson\":\"postsId\",\"dataTypeLong\":\"19\",\"comment\":\"帖子编号\",\"columnName\":\"posts_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"评论者\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"评论者\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CommentTime\",\"fieldDesc\":\"评论时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"commentTime\",\"dataTypeLong\":\"\",\"comment\":\"评论时间\",\"columnName\":\"comment_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CommentContent\",\"fieldDesc\":\"评论内容\",\"fieldType\":\"string\",\"fieldJson\":\"commentContent\",\"dataTypeLong\":\"\",\"comment\":\"评论内容\",\"columnName\":\"comment_content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"LikeTimes\",\"fieldDesc\":\"点赞数\",\"fieldType\":\"int\",\"fieldJson\":\"likeTimes\",\"dataTypeLong\":\"10\",\"comment\":\"点赞数\",\"columnName\":\"like_times\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ParentId\",\"fieldDesc\":\"父评论编号\",\"fieldType\":\"int\",\"fieldJson\":\"parentId\",\"dataTypeLong\":\"19\",\"comment\":\"父评论编号\",\"columnName\":\"parent_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckUser\",\"fieldDesc\":\"审核人\",\"fieldType\":\"int\",\"fieldJson\":\"checkUser\",\"dataTypeLong\":\"19\",\"comment\":\"审核人\",\"columnName\":\"check_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckTime\",\"fieldDesc\":\"审核时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"checkTime\",\"dataTypeLong\":\"\",\"comment\":\"审核时间\",\"columnName\":\"check_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CheckStatus\",\"fieldDesc\":\"审核状态(0未审批 1通过 2拒绝)\",\"fieldType\":\"int\",\"fieldJson\":\"checkStatus\",\"dataTypeLong\":\"10\",\"comment\":\"审核状态(0未审批 1通过 2拒绝)\",\"columnName\":\"check_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_comment.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_comment.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_comment.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_comment.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_comment.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumComment.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumComment\\hkForumCommentForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumComment\\hkForumComment.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumComment{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumCommentRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumCommentApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumCommentRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumCommentService;', 'HkForumComment', 'hkForumComment表', '209;210;211;212;213;214;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (22, '2023-02-01 08:45:40.376', '2023-02-01 08:45:40.376', NULL, 'community', '', 'hk_forum_posts_group', '{\"structName\":\"HkForumPostsGroup\",\"tableName\":\"hk_forum_posts_group\",\"packageName\":\"hkForumPostsGroup\",\"humpPackageName\":\"hk_forum_posts_group\",\"abbreviation\":\"hkForumPostsGroup\",\"description\":\"hkForumPostsGroup表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ParentId\",\"fieldDesc\":\"父节点编号\",\"fieldType\":\"int\",\"fieldJson\":\"parentId\",\"dataTypeLong\":\"19\",\"comment\":\"父节点编号\",\"columnName\":\"parent_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"审核状态(0未审批 1通过 2拒绝)\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"审核状态(0未审批 1通过 2拒绝)\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_posts_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_posts_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_posts_group.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_posts_group.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_posts_group.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumPostsGroup.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumPostsGroup\\hkForumPostsGroupForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumPostsGroup\\hkForumPostsGroup.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumPostsGroup{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumPostsGroupRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumPostsGroupApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumPostsGroupRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumPostsGroupService;', 'HkForumPostsGroup', 'hkForumPostsGroup表', '215;216;217;218;219;220;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (23, '2023-02-01 08:46:05.086', '2023-02-01 08:46:05.086', NULL, 'community', '', 'hk_forum_tag', '{\"structName\":\"HkForumTag\",\"tableName\":\"hk_forum_tag\",\"packageName\":\"hkForumTag\",\"humpPackageName\":\"hk_forum_tag\",\"abbreviation\":\"hkForumTag\",\"description\":\"hkForumTag表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"GroupId\",\"fieldDesc\":\"分组id\",\"fieldType\":\"int\",\"fieldJson\":\"groupId\",\"dataTypeLong\":\"19\",\"comment\":\"分组id\",\"columnName\":\"group_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_tag.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_tag.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_tag.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_tag.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_tag.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumTag.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTag\\hkForumTagForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTag\\hkForumTag.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumTag{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumTagRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumTagApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumTagRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumTagService;', 'HkForumTag', 'hkForumTag表', '221;222;223;224;225;226;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (24, '2023-02-01 08:46:26.426', '2023-02-01 08:46:26.426', NULL, 'community', '', 'hk_forum_tag_group', '{\"structName\":\"HkForumTagGroup\",\"tableName\":\"hk_forum_tag_group\",\"packageName\":\"hkForumTagGroup\",\"humpPackageName\":\"hk_forum_tag_group\",\"abbreviation\":\"hkForumTagGroup\",\"description\":\"hkForumTagGroup表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_tag_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_tag_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_tag_group.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_tag_group.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_tag_group.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumTagGroup.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTagGroup\\hkForumTagGroupForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTagGroup\\hkForumTagGroup.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumTagGroup{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumTagGroupRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumTagGroupApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumTagGroupRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumTagGroupService;', 'HkForumTagGroup', 'hkForumTagGroup表', '227;228;229;230;231;232;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (25, '2023-02-01 08:48:26.017', '2023-02-01 08:48:26.017', NULL, 'community', '', 'hk_forum_thumbs_up', '{\"structName\":\"HkForumThumbsUp\",\"tableName\":\"hk_forum_thumbs_up\",\"packageName\":\"hkForumThumbsUp\",\"humpPackageName\":\"hk_forum_thumbs_up\",\"abbreviation\":\"hkForumThumbsUp\",\"description\":\"hkForumThumbsUp表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户id\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Time\",\"fieldDesc\":\"点赞时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"time\",\"dataTypeLong\":\"\",\"comment\":\"点赞时间\",\"columnName\":\"time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PostsId\",\"fieldDesc\":\"帖子编号\",\"fieldType\":\"int\",\"fieldJson\":\"postsId\",\"dataTypeLong\":\"19\",\"comment\":\"帖子编号\",\"columnName\":\"posts_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_thumbs_up.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumThumbsUp.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumThumbsUp\\hkForumThumbsUpForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumThumbsUp\\hkForumThumbsUp.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumThumbsUp{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumThumbsUpRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumThumbsUpApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumThumbsUpRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumThumbsUpService;', 'HkForumThumbsUp', 'hkForumThumbsUp表', '233;234;235;236;237;238;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (26, '2023-02-01 08:48:43.563', '2023-02-01 08:48:43.563', NULL, 'community', '', 'hk_forum_topic', '{\"structName\":\"HkForumTopic\",\"tableName\":\"hk_forum_topic\",\"packageName\":\"hkForumTopic\",\"humpPackageName\":\"hk_forum_topic\",\"abbreviation\":\"hkForumTopic\",\"description\":\"hkForumTopic表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CoverImage\",\"fieldDesc\":\"封面\",\"fieldType\":\"string\",\"fieldJson\":\"coverImage\",\"dataTypeLong\":\"500\",\"comment\":\"封面\",\"columnName\":\"cover_image\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TopicGroupId\",\"fieldDesc\":\"分组id\",\"fieldType\":\"int\",\"fieldJson\":\"topicGroupId\",\"dataTypeLong\":\"19\",\"comment\":\"分组id\",\"columnName\":\"topic_group_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Type\",\"fieldDesc\":\"话题类型：0 全局，1圈子\",\"fieldType\":\"int\",\"fieldJson\":\"type\",\"dataTypeLong\":\"10\",\"comment\":\"话题类型：0 全局，1圈子\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Hot\",\"fieldDesc\":\"是否热门：0 否 1 是\",\"fieldType\":\"int\",\"fieldJson\":\"hot\",\"dataTypeLong\":\"10\",\"comment\":\"是否热门：0 否 1 是\",\"columnName\":\"hot\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"DiscussNum\",\"fieldDesc\":\"讨论数\",\"fieldType\":\"int\",\"fieldJson\":\"discussNum\",\"dataTypeLong\":\"10\",\"comment\":\"讨论数\",\"columnName\":\"discuss_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"AttentionNum\",\"fieldDesc\":\"关注数\",\"fieldType\":\"int\",\"fieldJson\":\"attentionNum\",\"dataTypeLong\":\"10\",\"comment\":\"关注数\",\"columnName\":\"attention_num\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ReviewStatus\",\"fieldDesc\":\"审核状态：0 未处理 1 通过，2驳回\",\"fieldType\":\"int\",\"fieldJson\":\"reviewStatus\",\"dataTypeLong\":\"10\",\"comment\":\"审核状态：0 未处理 1 通过，2驳回\",\"columnName\":\"review_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_topic.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_topic.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_topic.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_topic.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_topic.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumTopic.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTopic\\hkForumTopicForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTopic\\hkForumTopic.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumTopic{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumTopicRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumTopicApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumTopicRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumTopicService;', 'HkForumTopic', 'hkForumTopic表', '239;240;241;242;243;244;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (27, '2023-02-01 08:49:10.933', '2023-02-01 08:49:10.933', NULL, 'community', '', 'hk_forum_topic_group', '{\"structName\":\"HkForumTopicGroup\",\"tableName\":\"hk_forum_topic_group\",\"packageName\":\"hkForumTopicGroup\",\"humpPackageName\":\"hk_forum_topic_group\",\"abbreviation\":\"hkForumTopicGroup\",\"description\":\"hkForumTopicGroup表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_forum_topic_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_forum_topic_group.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_forum_topic_group.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_forum_topic_group.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_forum_topic_group.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkForumTopicGroup.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTopicGroup\\hkForumTopicGroupForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkForumTopicGroup\\hkForumTopicGroup.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkForumTopicGroup{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkForumTopicGroupRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkForumTopicGroupApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkForumTopicGroupRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkForumTopicGroupService;', 'HkForumTopicGroup', 'hkForumTopicGroup表', '245;246;247;248;249;250;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (28, '2023-02-01 08:50:02.335', '2023-02-01 08:50:02.335', NULL, 'general', '', 'hk_mini_program', '{\"structName\":\"HkMiniProgram\",\"tableName\":\"hk_mini_program\",\"packageName\":\"hkMiniProgram\",\"humpPackageName\":\"hk_mini_program\",\"abbreviation\":\"hkMiniProgram\",\"description\":\"hkMiniProgram表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"80\",\"comment\":\"名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Icon\",\"fieldDesc\":\"图标\",\"fieldType\":\"string\",\"fieldJson\":\"icon\",\"dataTypeLong\":\"256\",\"comment\":\"图标\",\"columnName\":\"icon\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompanyName\",\"fieldDesc\":\"公司名称\",\"fieldType\":\"string\",\"fieldJson\":\"companyName\",\"dataTypeLong\":\"256\",\"comment\":\"公司名称\",\"columnName\":\"company_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CurVersion\",\"fieldDesc\":\"当前版本\",\"fieldType\":\"string\",\"fieldJson\":\"curVersion\",\"dataTypeLong\":\"12\",\"comment\":\"当前版本\",\"columnName\":\"cur_version\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CurPacketId\",\"fieldDesc\":\"当前包id\",\"fieldType\":\"string\",\"fieldJson\":\"curPacketId\",\"dataTypeLong\":\"12\",\"comment\":\"当前包id\",\"columnName\":\"cur_packet_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Hidden\",\"fieldDesc\":\"隐藏(0否 1是)\",\"fieldType\":\"int\",\"fieldJson\":\"hidden\",\"dataTypeLong\":\"10\",\"comment\":\"隐藏(0否 1是)\",\"columnName\":\"hidden\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"general\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\hk_mini_program.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\hk_mini_program.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\request\\hk_mini_program.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\hk_mini_program.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\hk_mini_program.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkMiniProgram.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkMiniProgram\\hkMiniProgramForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkMiniProgram\\hkMiniProgram.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@general.HkMiniProgram{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@generalRouter.InitHkMiniProgramRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\enter.go@ApiGroup@HkMiniProgramApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\enter.go@RouterGroup@HkMiniProgramRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\enter.go@ServiceGroup@HkMiniProgramService;', 'HkMiniProgram', 'hkMiniProgram表', '251;252;253;254;255;256;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (29, '2023-02-01 08:50:26.620', '2023-02-01 08:50:26.620', NULL, 'general', '', 'hk_mini_program_packet', '{\"structName\":\"HkMiniProgramPacket\",\"tableName\":\"hk_mini_program_packet\",\"packageName\":\"hkMiniProgramPacket\",\"humpPackageName\":\"hk_mini_program_packet\",\"abbreviation\":\"hkMiniProgramPacket\",\"description\":\"hkMiniProgramPacket表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MiniProgramId\",\"fieldDesc\":\"小程序id\",\"fieldType\":\"int\",\"fieldJson\":\"miniProgramId\",\"dataTypeLong\":\"19\",\"comment\":\"小程序id\",\"columnName\":\"mini_program_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"包名\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"80\",\"comment\":\"包名\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Version\",\"fieldDesc\":\"版本\",\"fieldType\":\"string\",\"fieldJson\":\"version\",\"dataTypeLong\":\"80\",\"comment\":\"版本\",\"columnName\":\"version\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PacketAddress\",\"fieldDesc\":\"访问地址\",\"fieldType\":\"string\",\"fieldJson\":\"packetAddress\",\"dataTypeLong\":\"256\",\"comment\":\"访问地址\",\"columnName\":\"packet_address\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"general\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\hk_mini_program_packet.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\hk_mini_program_packet.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\request\\hk_mini_program_packet.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\hk_mini_program_packet.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\hk_mini_program_packet.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkMiniProgramPacket.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkMiniProgramPacket\\hkMiniProgramPacketForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkMiniProgramPacket\\hkMiniProgramPacket.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@general.HkMiniProgramPacket{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@generalRouter.InitHkMiniProgramPacketRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\enter.go@ApiGroup@HkMiniProgramPacketApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\enter.go@RouterGroup@HkMiniProgramPacketRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\enter.go@ServiceGroup@HkMiniProgramPacketService;', 'HkMiniProgramPacket', 'hkMiniProgramPacket表', '257;258;259;260;261;262;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (30, '2023-02-01 08:50:58.509', '2023-02-01 08:50:58.509', NULL, 'general', '', 'hk_protocol', '{\"structName\":\"HkProtocol\",\"tableName\":\"hk_protocol\",\"packageName\":\"hkProtocol\",\"humpPackageName\":\"hk_protocol\",\"abbreviation\":\"hkProtocol\",\"description\":\"hkProtocol表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"协议名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"20\",\"comment\":\"协议名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Des\",\"fieldDesc\":\"协议说明\",\"fieldType\":\"string\",\"fieldJson\":\"des\",\"dataTypeLong\":\"20\",\"comment\":\"协议说明\",\"columnName\":\"des\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"协议内容\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"2000\",\"comment\":\"协议内容\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Module\",\"fieldDesc\":\"所属模块/插件\",\"fieldType\":\"string\",\"fieldJson\":\"module\",\"dataTypeLong\":\"20\",\"comment\":\"所属模块/插件\",\"columnName\":\"module\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Pos\",\"fieldDesc\":\"显示位置\",\"fieldType\":\"string\",\"fieldJson\":\"pos\",\"dataTypeLong\":\"20\",\"comment\":\"显示位置\",\"columnName\":\"pos\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"general\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\hk_protocol.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\hk_protocol.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\request\\hk_protocol.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\hk_protocol.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\hk_protocol.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkProtocol.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkProtocol\\hkProtocolForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkProtocol\\hkProtocol.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@general.HkProtocol{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@generalRouter.InitHkProtocolRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\enter.go@ApiGroup@HkProtocolApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\enter.go@RouterGroup@HkProtocolRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\enter.go@ServiceGroup@HkProtocolService;', 'HkProtocol', 'hkProtocol表', '263;264;265;266;267;268;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (31, '2023-02-01 08:57:59.648', '2023-02-01 08:57:59.648', NULL, 'community', '', 'hk_report', '{\"structName\":\"HkReport\",\"tableName\":\"hk_report\",\"packageName\":\"hkReport\",\"humpPackageName\":\"hk_report\",\"abbreviation\":\"hkReport\",\"description\":\"hkReport表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ReportUserId\",\"fieldDesc\":\"被举报用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"reportUserId\",\"dataTypeLong\":\"19\",\"comment\":\"被举报用户编号\",\"columnName\":\"report_user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"举报用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"举报用户编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ReasonId\",\"fieldDesc\":\"举报原因_编号\",\"fieldType\":\"int\",\"fieldJson\":\"reasonId\",\"dataTypeLong\":\"19\",\"comment\":\"举报原因_编号\",\"columnName\":\"reason_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ReportType\",\"fieldDesc\":\"举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题\",\"fieldType\":\"int\",\"fieldJson\":\"reportType\",\"dataTypeLong\":\"10\",\"comment\":\"举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题\",\"columnName\":\"report_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"举报内容\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"200\",\"comment\":\"举报内容\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContentAttachment\",\"fieldDesc\":\"内容附件\",\"fieldType\":\"string\",\"fieldJson\":\"contentAttachment\",\"dataTypeLong\":\"400\",\"comment\":\"内容附件\",\"columnName\":\"content_attachment\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CurStatus\",\"fieldDesc\":\"处理状态：0 未处理、1已处理\",\"fieldType\":\"int\",\"fieldJson\":\"curStatus\",\"dataTypeLong\":\"10\",\"comment\":\"处理状态：0 未处理、1已处理\",\"columnName\":\"cur_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"HandleUserId\",\"fieldDesc\":\"操作用户_编号\",\"fieldType\":\"int\",\"fieldJson\":\"handleUserId\",\"dataTypeLong\":\"19\",\"comment\":\"操作用户_编号\",\"columnName\":\"handle_user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"HandleType\",\"fieldDesc\":\"处理方式:0无效处理（不予处理）、账号禁言\",\"fieldType\":\"int\",\"fieldJson\":\"handleType\",\"dataTypeLong\":\"19\",\"comment\":\"处理方式:0无效处理（不予处理）、账号禁言\",\"columnName\":\"handle_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"DurationId\",\"fieldDesc\":\"禁言时长_编号\",\"fieldType\":\"int\",\"fieldJson\":\"durationId\",\"dataTypeLong\":\"19\",\"comment\":\"禁言时长_编号\",\"columnName\":\"duration_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"HandleScore\",\"fieldDesc\":\"积分处理:0不扣分、1扣分\",\"fieldType\":\"int\",\"fieldJson\":\"handleScore\",\"dataTypeLong\":\"10\",\"comment\":\"积分处理:0不扣分、1扣分\",\"columnName\":\"handle_score\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ScoreExperience\",\"fieldDesc\":\"经验\",\"fieldType\":\"int\",\"fieldJson\":\"scoreExperience\",\"dataTypeLong\":\"10\",\"comment\":\"经验\",\"columnName\":\"score_experience\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ScoreCommunity\",\"fieldDesc\":\"社区积分\",\"fieldType\":\"int\",\"fieldJson\":\"scoreCommunity\",\"dataTypeLong\":\"10\",\"comment\":\"社区积分\",\"columnName\":\"score_community\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ScoreBuy\",\"fieldDesc\":\"购物积分\",\"fieldType\":\"int\",\"fieldJson\":\"scoreBuy\",\"dataTypeLong\":\"10\",\"comment\":\"购物积分\",\"columnName\":\"score_buy\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ScoreDownload\",\"fieldDesc\":\"下载值\",\"fieldType\":\"int\",\"fieldJson\":\"scoreDownload\",\"dataTypeLong\":\"10\",\"comment\":\"下载值\",\"columnName\":\"score_download\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Unlock\",\"fieldDesc\":\"是否解除：0 否、是\",\"fieldType\":\"int\",\"fieldJson\":\"unlock\",\"dataTypeLong\":\"10\",\"comment\":\"是否解除：0 否、是\",\"columnName\":\"unlock\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_report.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_report.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_report.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_report.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_report.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkReport.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkReport\\hkReportForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkReport\\hkReport.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkReport{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkReportRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkReportApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkReportRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkReportService;', 'HkReport', 'hkReport表', '269;270;271;272;273;274;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (32, '2023-02-01 08:58:40.443', '2023-02-01 08:58:40.443', NULL, 'community', '', 'hk_report_reason', '{\"structName\":\"HkReportReason\",\"tableName\":\"hk_report_reason\",\"packageName\":\"hkReportReason\",\"humpPackageName\":\"hk_report_reason\",\"abbreviation\":\"hkReportReason\",\"description\":\"hkReportReason表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Reason\",\"fieldDesc\":\"举报理由\",\"fieldType\":\"string\",\"fieldJson\":\"reason\",\"dataTypeLong\":\"20\",\"comment\":\"举报理由\",\"columnName\":\"reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_report_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_report_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_report_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_report_reason.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_report_reason.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkReportReason.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkReportReason\\hkReportReasonForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkReportReason\\hkReportReason.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkReportReason{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkReportReasonRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkReportReasonApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkReportReasonRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkReportReasonService;', 'HkReportReason', 'hkReportReason表', '275;276;277;278;279;280;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (33, '2023-02-01 09:00:28.955', '2023-02-01 09:00:28.955', NULL, 'general', '', 'hk_user_browsing_history', '{\"structName\":\"HkUserBrowsingHistory\",\"tableName\":\"hk_user_browsing_history\",\"packageName\":\"hkUserBrowsingHistory\",\"humpPackageName\":\"hk_user_browsing_history\",\"abbreviation\":\"hkUserBrowsingHistory\",\"description\":\"hkUserBrowsingHistory表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PostsId\",\"fieldDesc\":\"帖子编号\",\"fieldType\":\"int\",\"fieldJson\":\"postsId\",\"dataTypeLong\":\"19\",\"comment\":\"帖子编号\",\"columnName\":\"posts_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Time\",\"fieldDesc\":\"浏览时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"time\",\"dataTypeLong\":\"\",\"comment\":\"浏览时间\",\"columnName\":\"time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"general\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\hk_user_browsing_history.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\hk_user_browsing_history.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\request\\hk_user_browsing_history.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\hk_user_browsing_history.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\hk_user_browsing_history.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkUserBrowsingHistory.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserBrowsingHistory\\hkUserBrowsingHistoryForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserBrowsingHistory\\hkUserBrowsingHistory.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@general.HkUserBrowsingHistory{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@generalRouter.InitHkUserBrowsingHistoryRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\enter.go@ApiGroup@HkUserBrowsingHistoryApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\enter.go@RouterGroup@HkUserBrowsingHistoryRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\enter.go@ServiceGroup@HkUserBrowsingHistoryService;', 'HkUserBrowsingHistory', 'hkUserBrowsingHistory表', '281;282;283;284;285;286;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (34, '2023-02-01 09:05:02.117', '2023-02-01 09:05:02.117', NULL, 'community', '', 'hk_user_circle_apply', '{\"structName\":\"HkUserCircleApply\",\"tableName\":\"hk_user_circle_apply\",\"packageName\":\"hkUserCircleApply\",\"humpPackageName\":\"hk_user_circle_apply\",\"abbreviation\":\"hkUserCircleApply\",\"description\":\"hkUserCircleApply表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CircleId\",\"fieldDesc\":\"圈子_编号\",\"fieldType\":\"int\",\"fieldJson\":\"circleId\",\"dataTypeLong\":\"19\",\"comment\":\"圈子_编号\",\"columnName\":\"circle_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ApplyId\",\"fieldDesc\":\"应用_编号\",\"fieldType\":\"int\",\"fieldJson\":\"applyId\",\"dataTypeLong\":\"19\",\"comment\":\"应用_编号\",\"columnName\":\"apply_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"10\",\"comment\":\"排序\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_user_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_user_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_user_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_user_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_user_circle_apply.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkUserCircleApply.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserCircleApply\\hkUserCircleApplyForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserCircleApply\\hkUserCircleApply.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkUserCircleApply{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkUserCircleApplyRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkUserCircleApplyApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkUserCircleApplyRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkUserCircleApplyService;', 'HkUserCircleApply', 'hkUserCircleApply表', '287;288;289;290;291;292;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (35, '2023-02-01 09:05:53.755', '2023-02-01 09:05:53.755', NULL, 'general', '', 'hk_user_collect', '{\"structName\":\"HkUserCollect\",\"tableName\":\"hk_user_collect\",\"packageName\":\"hkUserCollect\",\"humpPackageName\":\"hk_user_collect\",\"abbreviation\":\"hkUserCollect\",\"description\":\"hkUserCollect表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户编号\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PostsId\",\"fieldDesc\":\"帖子编号\",\"fieldType\":\"int\",\"fieldJson\":\"postsId\",\"dataTypeLong\":\"19\",\"comment\":\"帖子编号\",\"columnName\":\"posts_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Time\",\"fieldDesc\":\"收藏时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"time\",\"dataTypeLong\":\"\",\"comment\":\"收藏时间\",\"columnName\":\"time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"general\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\hk_user_collect.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\hk_user_collect.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\general\\request\\hk_user_collect.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\hk_user_collect.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\hk_user_collect.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkUserCollect.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserCollect\\hkUserCollectForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserCollect\\hkUserCollect.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@general.HkUserCollect{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@generalRouter.InitHkUserCollectRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\general\\enter.go@ApiGroup@HkUserCollectApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\general\\enter.go@RouterGroup@HkUserCollectRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\general\\enter.go@ServiceGroup@HkUserCollectService;', 'HkUserCollect', 'hkUserCollect表', '293;294;295;296;297;298;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (36, '2023-02-01 09:09:53.713', '2023-02-01 09:09:53.713', NULL, 'community', '', 'hk_user_extend', '{\"structName\":\"HkUserExtend\",\"tableName\":\"hk_user_extend\",\"packageName\":\"hkUserExtend\",\"humpPackageName\":\"hk_user_extend\",\"abbreviation\":\"hkUserExtend\",\"description\":\"hkUserExtend表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserId\",\"fieldDesc\":\"用户ID\",\"fieldType\":\"int\",\"fieldJson\":\"userId\",\"dataTypeLong\":\"19\",\"comment\":\"用户ID\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Github\",\"fieldDesc\":\"github\",\"fieldType\":\"string\",\"fieldJson\":\"github\",\"dataTypeLong\":\"64\",\"comment\":\"github\",\"columnName\":\"github\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Weibo\",\"fieldDesc\":\"微博\",\"fieldType\":\"string\",\"fieldJson\":\"weibo\",\"dataTypeLong\":\"32\",\"comment\":\"微博\",\"columnName\":\"weibo\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Weixin\",\"fieldDesc\":\"微信\",\"fieldType\":\"string\",\"fieldJson\":\"weixin\",\"dataTypeLong\":\"32\",\"comment\":\"微信\",\"columnName\":\"weixin\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Qq\",\"fieldDesc\":\"qq\",\"fieldType\":\"string\",\"fieldJson\":\"qq\",\"dataTypeLong\":\"32\",\"comment\":\"qq\",\"columnName\":\"qq\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Blog\",\"fieldDesc\":\"博客\",\"fieldType\":\"string\",\"fieldJson\":\"blog\",\"dataTypeLong\":\"500\",\"comment\":\"博客\",\"columnName\":\"blog\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NumCircle\",\"fieldDesc\":\"圈子数\",\"fieldType\":\"int\",\"fieldJson\":\"numCircle\",\"dataTypeLong\":\"19\",\"comment\":\"圈子数\",\"columnName\":\"num_circle\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NumFocus\",\"fieldDesc\":\"关注数\",\"fieldType\":\"int\",\"fieldJson\":\"numFocus\",\"dataTypeLong\":\"19\",\"comment\":\"关注数\",\"columnName\":\"num_focus\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NumFan\",\"fieldDesc\":\"粉丝数\",\"fieldType\":\"int\",\"fieldJson\":\"numFan\",\"dataTypeLong\":\"19\",\"comment\":\"粉丝数\",\"columnName\":\"num_fan\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_user_extend.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_user_extend.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_user_extend.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_user_extend.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_user_extend.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkUserExtend.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserExtend\\hkUserExtendForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUserExtend\\hkUserExtend.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkUserExtend{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkUserExtendRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkUserExtendApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkUserExtendRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkUserExtendService;', 'HkUserExtend', 'hkUserExtend表', '299;300;301;302;303;304;', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_code_histories` VALUES (37, '2023-02-01 09:19:47.067', '2023-02-01 09:19:47.067', NULL, 'community', '', 'hk_user', '{\"structName\":\"HkUser\",\"tableName\":\"hk_user\",\"packageName\":\"hkUser\",\"humpPackageName\":\"hk_user\",\"abbreviation\":\"hkUser\",\"description\":\"hkUser表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TenantId\",\"fieldDesc\":\"租户ID\",\"fieldType\":\"string\",\"fieldJson\":\"tenantId\",\"dataTypeLong\":\"12\",\"comment\":\"租户ID\",\"columnName\":\"tenant_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Uuid\",\"fieldDesc\":\"用户编号\",\"fieldType\":\"string\",\"fieldJson\":\"uuid\",\"dataTypeLong\":\"50\",\"comment\":\"用户编号\",\"columnName\":\"uuid\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserType\",\"fieldDesc\":\"用户平台\",\"fieldType\":\"int\",\"fieldJson\":\"userType\",\"dataTypeLong\":\"10\",\"comment\":\"用户平台\",\"columnName\":\"user_type\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Account\",\"fieldDesc\":\"账号\",\"fieldType\":\"string\",\"fieldJson\":\"account\",\"dataTypeLong\":\"45\",\"comment\":\"账号\",\"columnName\":\"account\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Password\",\"fieldDesc\":\"密码\",\"fieldType\":\"string\",\"fieldJson\":\"password\",\"dataTypeLong\":\"45\",\"comment\":\"密码\",\"columnName\":\"password\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"NickName\",\"fieldDesc\":\"昵称\",\"fieldType\":\"string\",\"fieldJson\":\"nickName\",\"dataTypeLong\":\"20\",\"comment\":\"昵称\",\"columnName\":\"nick_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RealName\",\"fieldDesc\":\"真名\",\"fieldType\":\"string\",\"fieldJson\":\"realName\",\"dataTypeLong\":\"10\",\"comment\":\"真名\",\"columnName\":\"real_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"HeaderImg\",\"fieldDesc\":\"头像\",\"fieldType\":\"string\",\"fieldJson\":\"headerImg\",\"dataTypeLong\":\"500\",\"comment\":\"头像\",\"columnName\":\"header_img\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Email\",\"fieldDesc\":\"邮箱\",\"fieldType\":\"string\",\"fieldJson\":\"email\",\"dataTypeLong\":\"45\",\"comment\":\"邮箱\",\"columnName\":\"email\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Phone\",\"fieldDesc\":\"手机\",\"fieldType\":\"string\",\"fieldJson\":\"phone\",\"dataTypeLong\":\"45\",\"comment\":\"手机\",\"columnName\":\"phone\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Birthday\",\"fieldDesc\":\"生日\",\"fieldType\":\"time.Time\",\"fieldJson\":\"birthday\",\"dataTypeLong\":\"\",\"comment\":\"生日\",\"columnName\":\"birthday\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sex\",\"fieldDesc\":\"性别\",\"fieldType\":\"int\",\"fieldJson\":\"sex\",\"dataTypeLong\":\"10\",\"comment\":\"性别\",\"columnName\":\"sex\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleId\",\"fieldDesc\":\"用户角色ID\",\"fieldType\":\"int\",\"fieldJson\":\"roleId\",\"dataTypeLong\":\"20\",\"comment\":\"用户角色ID\",\"columnName\":\"role_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateUser\",\"fieldDesc\":\"创建人\",\"fieldType\":\"int\",\"fieldJson\":\"createUser\",\"dataTypeLong\":\"19\",\"comment\":\"创建人\",\"columnName\":\"create_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateDept\",\"fieldDesc\":\"创建部门\",\"fieldType\":\"int\",\"fieldJson\":\"createDept\",\"dataTypeLong\":\"19\",\"comment\":\"创建部门\",\"columnName\":\"create_dept\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CreateTime\",\"fieldDesc\":\"创建时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"createTime\",\"dataTypeLong\":\"\",\"comment\":\"创建时间\",\"columnName\":\"create_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateUser\",\"fieldDesc\":\"修改人\",\"fieldType\":\"int\",\"fieldJson\":\"updateUser\",\"dataTypeLong\":\"19\",\"comment\":\"修改人\",\"columnName\":\"update_user\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UpdateTime\",\"fieldDesc\":\"修改时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"updateTime\",\"dataTypeLong\":\"\",\"comment\":\"修改时间\",\"columnName\":\"update_time\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态(用户是否被冻结) 1正常 2冻结\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"10\",\"comment\":\"状态(用户是否被冻结) 1正常 2冻结\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsDeleted\",\"fieldDesc\":\"是否已删除\",\"fieldType\":\"int\",\"fieldJson\":\"isDeleted\",\"dataTypeLong\":\"10\",\"comment\":\"是否已删除\",\"columnName\":\"is_deleted\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":true,\"package\":\"community\"}', 'D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\model\\community\\request\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\hk_user.go;D:\\_works\\community\\gin-vue-admin\\web\\src\\api\\hkUser.js;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUser\\hkUserForm.vue;D:\\_works\\community\\gin-vue-admin\\web\\src\\view\\hkUser\\hkUser.vue;', 'D:\\_works\\community\\gin-vue-admin\\server\\initialize\\gorm.go@MysqlTables@community.HkUser{},;D:\\_works\\community\\gin-vue-admin\\server\\initialize\\router.go@Routers@communityRouter.InitHkUserRouter(PrivateGroup);D:\\_works\\community\\gin-vue-admin\\server\\api\\v1\\community\\enter.go@ApiGroup@HkUserApi;D:\\_works\\community\\gin-vue-admin\\server\\router\\community\\enter.go@RouterGroup@HkUserRouter;D:\\_works\\community\\gin-vue-admin\\server\\service\\community\\enter.go@ServiceGroup@HkUserService;', 'HkUser', 'hkUser表', '', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_auto_codes
-- ----------------------------
INSERT INTO `sys_auto_codes` VALUES (1, '2023-01-31 14:45:05.767', '2023-01-31 14:45:05.767', NULL, 'community', '社区', '社区描述', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_codes` VALUES (2, '2023-02-01 08:19:03.660', '2023-02-01 08:19:03.660', NULL, 'apply', '应用', '应用描述', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_codes` VALUES (3, '2023-02-01 08:33:41.552', '2023-02-01 08:33:41.552', NULL, 'general', '常规业务', '常规业务描述', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_auto_codes` VALUES (4, '2023-02-01 10:15:19.157', '2023-02-01 10:15:19.157', NULL, 'app', 'app包', 'app包描述', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_base_menu_btns
-- ----------------------------

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
-- Records of sys_base_menu_parameters
-- ----------------------------

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
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '仪表盘', 'odometer', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (2, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (3, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '超级管理员', 'user', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (4, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, '角色管理', 'avatar', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (5, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '', 1, 0, '菜单管理', 'tickets', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (6, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'api管理', 'platform', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (7, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '', 0, 0, '用户管理', 'coordinate', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (8, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (9, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 'dictionary', 0, 0, '字典详情-${id}', 'list', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (10, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (11, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (12, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 7, '', 0, 0, '示例文件', 'management', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (13, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '12', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '', 0, 0, '媒体库（上传下载）', 'upload', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (14, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '12', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (15, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '12', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '', 0, 0, '客户列表（资源示例）', 'avatar', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (16, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '系统工具', 'tools', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (17, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '16', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (18, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '16', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '', 1, 0, '表单生成器', 'magic-stick', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (19, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '16', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '', 0, 0, '系统配置', 'operation', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (20, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '16', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '', 0, 0, '自动化代码管理', 'magic-stick', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (21, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '16', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (22, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '16', 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '自动化package', 'folder', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (23, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, '', 0, 0, '官方网站', 'home-filled', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (24, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (25, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '0', 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (26, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '25', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (27, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '25', 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (28, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '25', 'autoPlug', 'autoPlug', 0, 'view/systemTools/autoPlug/autoPlug.vue', 2, '', 0, 0, '插件模板', 'folder', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (29, '2023-01-31 14:44:24.822', '2023-01-31 14:44:24.822', NULL, 0, '25', 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 3, '', 0, 0, '邮件插件', 'message', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (30, '2023-02-05 13:45:48.820', '2023-02-05 15:29:57.213', NULL, 0, '0', 'hkUser', 'hkUser', 0, 'view/hkUser/hkUser.vue', 0, '', 0, 0, '终端用户', 'user', 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2023-01-31 14:44:24.766', '2023-01-31 14:44:24.773', NULL, '性别', 'gender', 1, '性别字典', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionaries` VALUES (2, '2023-01-31 14:44:24.766', '2023-01-31 14:44:24.781', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionaries` VALUES (3, '2023-01-31 14:44:24.766', '2023-01-31 14:44:24.789', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionaries` VALUES (4, '2023-01-31 14:44:24.766', '2023-01-31 14:44:24.797', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionaries` VALUES (5, '2023-01-31 14:44:24.766', '2023-01-31 14:44:24.805', NULL, '数据库字符串', 'string', 1, '数据库字符串', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionaries` VALUES (6, '2023-01-31 14:44:24.766', '2023-01-31 14:44:24.813', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型', NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2023-01-31 14:44:24.774', '2023-01-31 14:44:24.774', NULL, '男', 1, 1, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (2, '2023-01-31 14:44:24.774', '2023-01-31 14:44:24.774', NULL, '女', 2, 1, 2, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (3, '2023-01-31 14:44:24.782', '2023-01-31 14:44:24.782', NULL, 'smallint', 1, 1, 1, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (4, '2023-01-31 14:44:24.782', '2023-01-31 14:44:24.782', NULL, 'mediumint', 2, 1, 2, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (5, '2023-01-31 14:44:24.782', '2023-01-31 14:44:24.782', NULL, 'int', 3, 1, 3, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (6, '2023-01-31 14:44:24.782', '2023-01-31 14:44:24.782', NULL, 'bigint', 4, 1, 4, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (7, '2023-01-31 14:44:24.791', '2023-01-31 14:44:24.791', NULL, 'date', 0, 1, 0, 3, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (8, '2023-01-31 14:44:24.791', '2023-01-31 14:44:24.791', NULL, 'time', 1, 1, 1, 3, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (9, '2023-01-31 14:44:24.791', '2023-01-31 14:44:24.791', NULL, 'year', 2, 1, 2, 3, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (10, '2023-01-31 14:44:24.791', '2023-01-31 14:44:24.791', NULL, 'datetime', 3, 1, 3, 3, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (11, '2023-01-31 14:44:24.791', '2023-01-31 14:44:24.791', NULL, 'timestamp', 5, 1, 5, 3, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (12, '2023-01-31 14:44:24.799', '2023-01-31 14:44:24.799', NULL, 'float', 0, 1, 0, 4, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (13, '2023-01-31 14:44:24.799', '2023-01-31 14:44:24.799', NULL, 'double', 1, 1, 1, 4, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (14, '2023-01-31 14:44:24.799', '2023-01-31 14:44:24.799', NULL, 'decimal', 2, 1, 2, 4, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (15, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'char', 0, 1, 0, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (16, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'varchar', 1, 1, 1, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (17, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'tinyblob', 2, 1, 2, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (18, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'tinytext', 3, 1, 3, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (19, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'text', 4, 1, 4, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (20, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'blob', 5, 1, 5, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (21, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'mediumblob', 6, 1, 6, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (22, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'mediumtext', 7, 1, 7, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (23, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'longblob', 8, 1, 8, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (24, '2023-01-31 14:44:24.806', '2023-01-31 14:44:24.806', NULL, 'longtext', 9, 1, 9, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dictionary_details` VALUES (25, '2023-01-31 14:44:24.815', '2023-01-31 14:44:24.815', NULL, 'tinyint', 0, 1, 0, 6, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
-- Records of sys_operation_records
-- ----------------------------

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
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 888, NULL);
INSERT INTO `sys_user_authority` VALUES (1, 8881, NULL);
INSERT INTO `sys_user_authority` VALUES (1, 9528, NULL);
INSERT INTO `sys_user_authority` VALUES (2, 888, NULL);

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

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2023-01-31 14:44:24.978', '2023-02-05 15:48:50.281', NULL, '2c8a6e5f-fb10-44af-9e1b-fd47af84faca', 'admin', '$2a$10$nhay0n5NsDvIdlKLstqFW.a4iGOhLgtg4sTTp0Hfhz6KKL4pE27Ma', 'Mr.奇淼', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', 888, '17611111111', '333333333@qq.com', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_users` VALUES (2, '2023-01-31 14:44:24.978', '2023-01-31 14:44:24.992', NULL, 'cb771ae5-7280-424f-8675-328ccf44bd03', 'a303176530', '$2a$10$YUPZkAWInC5qjW1DbyV0IeIgSZeCRU8jQQgbGPmewvU.rnTeLboma', '用户1', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', 9528, '17611111111', '333333333@qq.com', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
