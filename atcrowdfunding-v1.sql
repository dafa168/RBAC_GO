/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.20.235
 Source Server Type    : MySQL
 Source Server Version : 50173
 Source Host           : 192.168.20.235:3306
 Source Schema         : atcrowdfunding-v1

 Target Server Type    : MySQL
 Target Server Version : 50173
 File Encoding         : 65001

 Date: 02/06/2020 14:02:48
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  INDEX `IDX_casbin_rule_v2`(`v2`) USING BTREE,
  INDEX `IDX_casbin_rule_v3`(`v3`) USING BTREE,
  INDEX `IDX_casbin_rule_v4`(`v4`) USING BTREE,
  INDEX `IDX_casbin_rule_v5`(`v5`) USING BTREE,
  INDEX `IDX_casbin_rule_p_type`(`p_type`) USING BTREE,
  INDEX `IDX_casbin_rule_v0`(`v0`) USING BTREE,
  INDEX `IDX_casbin_rule_v1`(`v1`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/back/*', '*', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'root', 'root', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'useradmin', '/back/user/*', '*', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'ccadmin', 'useradmin', '', '', '', '');

-- ----------------------------
-- Table structure for oauth_token
-- ----------------------------
DROP TABLE IF EXISTS `oauth_token`;
CREATE TABLE `oauth_token`  (
  `token` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Token',
  `user_id` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'UserId',
  `secret` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'Secret',
  `express_in` bigint(20) NOT NULL DEFAULT 0 COMMENT '是否是标准库',
  `revoked` tinyint(1) NULL DEFAULT NULL
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of oauth_token
-- ----------------------------
INSERT INTO `oauth_token` VALUES ('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTA3MTU0OTIsImlhdCI6MTU5MDcxMTg5Mn0.MhxW8uCbqGemENUHx6EhyiizBv1V25Gz5-qR9Nh8QBc', '3', 'secret', 1590715492, 0);
INSERT INTO `oauth_token` VALUES ('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTA2NjAxMzAsImlhdCI6MTU5MDY1NjUzMH0.9nTQUNRhlZRn1S5Z8IQuvetUYnz_drMJDQZZk2tT3wk', '4', 'secret', 1590660130, 0);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `Ename` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_role_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 'PM-项目经理', '');
INSERT INTO `role` VALUES (2, 'SE-软件工程师', '');
INSERT INTO `role` VALUES (3, 'PG-程序员', '');
INSERT INTO `role` VALUES (4, 'TL-组长', '');
INSERT INTO `role` VALUES (5, 'GL-组长', '');
INSERT INTO `role` VALUES (6, 'QC-品质控制', '');
INSERT INTO `role` VALUES (7, 'SA-软件架构师', '');
INSERT INTO `role` VALUES (8, 'SYSTEM-系统管理', '');
INSERT INTO `role` VALUES (9, '管理员', 'admin');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `login_acct` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_ps_wd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_user_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (3, 'username', 'admin', 'admin', 'email', NULL);
INSERT INTO `user` VALUES (4, 'root', 'root', 'root', NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
