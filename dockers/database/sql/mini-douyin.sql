SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- table structure for comment;
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '评论记录 id',
    `user_id` int(11) NOT NULL COMMENT '评论者 id',
    `video_id` int(11) NOT NULL COMMENT '视频 id',
    `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '评论内容',
    `action_type` int(11) NOT NULL DEFAULT '1' COMMENT '评论类型，1 为评论，2 为删除',
    `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发布评论时间',
    `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新评论时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `fk_comment_user`(`user_id`) USING BTREE,
    INDEX `fk_comment_video`(`video_id`) USING BTREE,
    CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_comment_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- table structure for like;
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '点赞记录 id',
    `user_id` int(11) NOT NULL COMMENT '点赞者 id',
    `video_id` int(11) NOT NULL COMMENT '视频 id',
    `liked` int(11) NOT NULL COMMENT '1 为已点赞，2 未点赞',
    `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞时间',
    `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新点赞时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `fk_like_user`(`user_id`) USING BTREE,
    INDEX `fk_like_video`(`video_id`) USING BTREE,
    CONSTRAINT `fk_like_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_like_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- table structure for `message`;
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` ( 
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '消息记录 id',
    `user_id` int(11) NOT NULL COMMENT '发送消息的 user_id',
    `receiver_id` int(11) NOT NULL COMMENT '接收消息的 user_id',
    `msg_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '消息内容',
    `action_type` int(11) NOT NULL COMMENT '消息类型，1 为发送，2 为撤回',
    `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发送消息时间',
    `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新消息时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `fk_message_user_1`(`user_id`) USING BTREE,
    INDEX `fk_message_user_2`(`receiver_id`) USING BTREE,
    CONSTRAINT `fk_message_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_message_user_2` FOREIGN KEY (`receiver_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- table structure for `relation`;
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '关注记录 id',
    `user_id` int(11) NOT NULL COMMENT '关注者 id',
    `following_id` int(11) NOT NULL COMMENT '被关注者 id',
    `followed` int(11) NOT NULL COMMENT '1 为已关注，2 为未关注',
    `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关注时间',
    `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新关注时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `fk_relation_user_1`(`user_id`) USING BTREE,
    INDEX `fk_relation_user_2`(`following_id`) USING BTREE,
    CONSTRAINT `fk_relation_user_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_relation_user_2` FOREIGN KEY (`following_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- table structure for `user`;
drop table if exists `user`;
CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户 id',
    `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
    `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
    `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像',
    `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建用户时间',
    `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新用户时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `name`(`name`) USING BTREE COMMENT '用户名唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- table structure for `video`;
drop table if exists `video`;
CREATE TABLE `video` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '视频 id',
    `author_id` int(11) NOT NULL COMMENT '视频上传者 id',
    `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频标题',
    `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频地址',
    `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频封面地址',
    `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '上传视频时间',
    `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新视频时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `fk_video_user`(`author_id`) USING BTREE,
    CONSTRAINT `fk_video_user` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
