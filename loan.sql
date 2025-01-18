/*
 Navicat Premium Dump SQL

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 90100 (9.1.0)
 Source Host           : 192.168.188.233:3306
 Source Schema         : loan

 Target Server Type    : MySQL
 Target Server Version : 90100 (9.1.0)
 File Encoding         : 65001

 Date: 18/01/2025 17:37:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cache
-- ----------------------------
DROP TABLE IF EXISTS `cache`;
CREATE TABLE `cache` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `cache_key` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cache_value` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for deposit
-- ----------------------------
DROP TABLE IF EXISTS `deposit`;
CREATE TABLE `deposit` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `loan_id` int DEFAULT NULL,
  `aleo_address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `aleo_amount` decimal(18,0) DEFAULT NULL,
  `aleo_price` decimal(18,6) DEFAULT NULL,
  `usdt_value` decimal(30,0) DEFAULT NULL,
  `hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `at` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for gorm_templete
-- ----------------------------
DROP TABLE IF EXISTS `gorm_templete`;
CREATE TABLE `gorm_templete` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for image_assets
-- ----------------------------
DROP TABLE IF EXISTS `image_assets`;
CREATE TABLE `image_assets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for income_record
-- ----------------------------
DROP TABLE IF EXISTS `income_record`;
CREATE TABLE `income_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `amount` decimal(30,0) DEFAULT NULL,
  `at` int DEFAULT NULL,
  `is_negative` int DEFAULT NULL,
  `type` int DEFAULT NULL,
  `split_days` int DEFAULT NULL,
  `end_at` int DEFAULT NULL,
  `hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for leo_price_record
-- ----------------------------
DROP TABLE IF EXISTS `leo_price_record`;
CREATE TABLE `leo_price_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `price` decimal(18,6) DEFAULT NULL,
  `at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=120 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for leo_rate_record
-- ----------------------------
DROP TABLE IF EXISTS `leo_rate_record`;
CREATE TABLE `leo_rate_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `rate` decimal(10,6) DEFAULT NULL,
  `at` int DEFAULT NULL,
  `days` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for loan
-- ----------------------------
DROP TABLE IF EXISTS `loan`;
CREATE TABLE `loan` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `aleo_address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bsc_address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` int DEFAULT NULL,
  `stages` int DEFAULT NULL,
  `pay_stages` int DEFAULT NULL,
  `day_per_stage` int DEFAULT NULL,
  `start_at` int DEFAULT NULL,
  `health` decimal(10,6) DEFAULT NULL,
  `rate` decimal(10,6) DEFAULT NULL,
  `release_rate` decimal(10,6) DEFAULT NULL,
  `hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` int DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bsc_loan_id` int DEFAULT NULL,
  `release_at` int DEFAULT NULL,
  `release_hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `release_amount` decimal(30,0) DEFAULT NULL,
  `interest_amount` decimal(30,0) DEFAULT NULL,
  `pay_back_at` int DEFAULT NULL,
  `pay_back_hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pay_back_amount` decimal(30,0) DEFAULT NULL,
  `release_aleo_hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `release_aleo_at` int DEFAULT NULL,
  `release_aleo_amount` decimal(18,0) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for loan_config
-- ----------------------------
DROP TABLE IF EXISTS `loan_config`;
CREATE TABLE `loan_config` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `rate` decimal(10,6) DEFAULT NULL,
  `release_rate` decimal(10,6) DEFAULT NULL,
  `available_stages` int DEFAULT NULL,
  `day_per_stage` int DEFAULT NULL,
  `allow_types` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `banner_ids` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `aleo_price` decimal(10,6) DEFAULT NULL,
  `min_loan_amount` decimal(18,0) DEFAULT NULL,
  `max_loan_amount` decimal(18,0) DEFAULT NULL,
  `platform_income_rate` decimal(10,6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for provide_liquid
-- ----------------------------
DROP TABLE IF EXISTS `provide_liquid`;
CREATE TABLE `provide_liquid` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `amount` decimal(30,0) DEFAULT NULL,
  `duration` int DEFAULT NULL,
  `start` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  `provider` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `create_at` int DEFAULT NULL,
  `create_hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `retrieve_at` int DEFAULT NULL,
  `retrieve_hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for provide_record
-- ----------------------------
DROP TABLE IF EXISTS `provide_record`;
CREATE TABLE `provide_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `type` int DEFAULT NULL,
  `provider` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `amount` decimal(30,0) DEFAULT NULL,
  `hash` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for provide_reward_record
-- ----------------------------
DROP TABLE IF EXISTS `provide_reward_record`;
CREATE TABLE `provide_reward_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `type` int DEFAULT NULL,
  `provider` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `amount` decimal(30,0) DEFAULT NULL,
  `hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `at` int DEFAULT NULL,
  `source_type` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
