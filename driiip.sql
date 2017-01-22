/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50634
Source Host           : 10.129.233.79:3306
Source Database       : driiip

Target Server Type    : MYSQL
Target Server Version : 50634
File Encoding         : 65001

Date: 2017-01-22 17:24:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for dp_monitor_log
-- ----------------------------
DROP TABLE IF EXISTS `dp_monitor_log`;
CREATE TABLE `dp_monitor_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `device_id` int(10) unsigned NOT NULL,
  `date` int(10) unsigned NOT NULL DEFAULT '0',
  `temperature` int(11) NOT NULL DEFAULT '0',
  `illumination` int(11) NOT NULL DEFAULT '0',
  `water` int(11) NOT NULL DEFAULT '0',
  `humidity` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dp_plant
-- ----------------------------
DROP TABLE IF EXISTS `dp_plant`;
CREATE TABLE `dp_plant` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态：0下架，1上架',
  `common_name` varchar(255) NOT NULL DEFAULT '',
  `latin_name` varchar(255) NOT NULL DEFAULT '',
  `fullname` varchar(255) NOT NULL DEFAULT '',
  `genus_name` varchar(255) NOT NULL DEFAULT '',
  `image` varchar(255) NOT NULL DEFAULT '',
  `subspecies_name` varchar(255) NOT NULL DEFAULT '',
  `bloom_color` char(32) NOT NULL DEFAULT '',
  `leaf_color` char(32) NOT NULL DEFAULT '',
  `plant_type` varchar(255) NOT NULL DEFAULT '',
  `spec_features` varchar(255) NOT NULL DEFAULT '',
  `plant_shape` varchar(255) NOT NULL DEFAULT '',
  `bloom_time` char(255) NOT NULL DEFAULT '',
  `info` blob,
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `upadte_time` int(10) unsigned NOT NULL DEFAULT '0',
  `plant_uuid` char(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dp_user
-- ----------------------------
DROP TABLE IF EXISTS `dp_user`;
CREATE TABLE `dp_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `avatar` varchar(255) NOT NULL DEFAULT '',
  `login_time` int(10) unsigned NOT NULL DEFAULT '0',
  `register_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dp_user_device
-- ----------------------------
DROP TABLE IF EXISTS `dp_user_device`;
CREATE TABLE `dp_user_device` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `device_name` varchar(255) NOT NULL DEFAULT '',
  `per_uuid` varchar(255) NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
SET FOREIGN_KEY_CHECKS=1;
