/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.60.100
 Source Server Type    : MySQL
 Source Server Version : 80020
 Source Host           : 192.168.60.100:3306
 Source Schema         : atcrowdfunding-v1

 Target Server Type    : MySQL
 Target Server Version : 80020
 File Encoding         : 65001

 Date: 22/05/2020 17:34:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pid` int(0) NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_permission_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES (1, '系统菜单', 0, NULL, NULL);
INSERT INTO `permission` VALUES (2, '控制面板', 1, NULL, 'glyphicon glyphicon-dashboard');
INSERT INTO `permission` VALUES (3, '权限管理', 1, NULL, 'glyphicon glyphicon glyphicon-tasks');
INSERT INTO `permission` VALUES (4, '用户维护', 3, '/user/index', 'glyphicon glyphicon-user');
INSERT INTO `permission` VALUES (5, '角色维护', 3, '/role/index', 'glyphicon glyphicon-king');
INSERT INTO `permission` VALUES (6, '许可维护', 3, '/permission/index', 'glyphicon glyphicon-lock');
INSERT INTO `permission` VALUES (7, '业务审核', 1, '/error', NULL);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_role_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 'PM-项目经理');
INSERT INTO `role` VALUES (2, 'SE-软件工程师');
INSERT INTO `role` VALUES (3, 'PG-程序员');
INSERT INTO `role` VALUES (4, 'TL-组长');
INSERT INTO `role` VALUES (5, 'GL-组长');
INSERT INTO `role` VALUES (6, 'QC-品质控制');
INSERT INTO `role` VALUES (7, 'SA-软件架构师');
INSERT INTO `role` VALUES (8, 'SYSTEM-系统管理');

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `role_id` int(0) NULL DEFAULT NULL,
  `permission_id` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_role_permission_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES (4, 1, 1);
INSERT INTO `role_permission` VALUES (6, 1, 2);
INSERT INTO `role_permission` VALUES (7, 2, 2);
INSERT INTO `role_permission` VALUES (16, 1, 3);
INSERT INTO `role_permission` VALUES (17, 1, 4);
INSERT INTO `role_permission` VALUES (18, 1, 5);
INSERT INTO `role_permission` VALUES (19, 1, 6);
INSERT INTO `role_permission` VALUES (20, 6, 1);
INSERT INTO `role_permission` VALUES (21, 6, 2);
INSERT INTO `role_permission` VALUES (22, 6, 3);
INSERT INTO `role_permission` VALUES (23, 6, 5);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `login_acct` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_ps_wd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_user_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (3, 'zhangsan', 'admin', 'admin', NULL, NULL);
INSERT INTO `user` VALUES (4, 'root', 'root', 'root', NULL, NULL);

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `user_id` int(0) NULL DEFAULT NULL,
  `role_id` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `UQE_user_role_id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1, 1);
INSERT INTO `user_role` VALUES (2, 4, 3);

SET FOREIGN_KEY_CHECKS = 1;
