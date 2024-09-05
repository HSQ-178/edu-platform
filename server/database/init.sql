CREATE DATABASE IF NOT EXISTS `edu-platform` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

USE `edu-platform`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `role_id` INT(0) UNSIGNED NOT NULL COMMENT '角色ID',
    `username` VARCHAR(24) NOT NULL COMMENT '用户名',
    `password` VARCHAR(24) NOT NULL COMMENT '密码',
    `name` VARCHAR(15) NOT NULL COMMENT '姓名',
    `nickname` VARCHAR(15) NOT NULL COMMENT '昵称',
    `email` VARCHAR(50) COMMENT '邮箱',
    `phone` VARCHAR(18) COMMENT '手机号',
    `avatar` VARCHAR(255) COMMENT '头像',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
    `created_at` DATETIME NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_user_username` (`username`),
    UNIQUE INDEX `idx_user_nickname` (`nickname`),
    UNIQUE INDEX `idx_user_phone` (`phone`),
    UNIQUE INDEX `idx_user_email` (`email`),
    INDEX `idx_user_role` (`role_id`),
    INDEX `idx_user_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户表';

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE IF NOT EXISTS `user_role`(
    `id` INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `role_name` VARCHAR(10) NOT NULL COMMENT '角色名',
    `created_at` DATETIME NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_user_role_name` (`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户角色表';

DROP TABLE IF EXISTS `user_role_permission`;
CREATE TABLE IF NOT EXISTS `permission` (
    `id` INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `role_id` INT(0) UNSIGNED NOT NULL COMMENT '绑定的角色ID',
    `resource` VARCHAR(100) NOT NULL COMMENT '允许的资源',
    `action` VARCHAR(100) NOT NULL COMMENT '允许的行为',
    PRIMARY KEY (`id`),
    INDEX `idx_user_role_permission` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户角色权限表';

DROP TABLE IF EXISTS `course`;
CREATE TABLE IF NOT EXISTS `course` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `course_classification_id` INT(0) NOT NULL COMMENT '课程分类ID',
    `user_id` BIGINT COMMENT '用户ID',
    `student_name` VARCHAR(255) NOT NULL COMMENT '学生名',
    `multiple` TINYINT(1) NOT NULL DEFAULT 2 COMMENT '是否多名学生 1:是 2:否',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
    `created_at` DATETIME NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_course_user_id` (`user_id`),
    INDEX `idx_course_multiple` (`multiple`),
    INDEX `idx_course_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '课程表';

DROP TABLE IF EXISTS `course_classification`;
CREATE TABLE IF NOT EXISTS `course_classification` (
    `id` INT(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title` VARCHAR(255) NOT NULL COMMENT '分类标题',
    `profit` INT(0) NOT NULL COMMENT '每课时课时费',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:启用 2:冻结 3:删除',
    `created_at` DATETIME NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_course_classification_title` (`title`),
    INDEX `idx_course_classification_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '课程分类表';

DROP TABLE IF EXISTS `class_record`;
CREATE TABLE IF NOT EXISTS `class_record` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `course_id` BIGINT UNSIGNED NOT NULL COMMENT '课程ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '授课用户ID',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 1:签到 2:未签到 3:删除',
    `created_at` DATETIME NOT NULL COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_class_record_course_id` (`course_id`),
    INDEX `idx_class_record_user_id` (`user_id`),
    INDEX `idx_class_record_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '课时记录表';