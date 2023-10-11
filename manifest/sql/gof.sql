CREATE TABLE `channel_group` (
                                 `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                 `channel` varchar(256) NOT NULL,
                                 `group` int(10) unsigned NOT NULL,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `channel_group_uk` (`channel`,`group`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `server` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                          `name` varchar(100) NOT NULL,
                          `ip` varchar(100) NOT NULL,
                          `group` int(10) unsigned NOT NULL,
                          `http` varchar(255) NOT NULL,
                          `max_reg` int(10) unsigned NOT NULL,
                          `socket` varchar(255) NOT NULL,
                          `max_online` int(10) unsigned NOT NULL,
                          `status` tinyint(1) NOT NULL,
                          `hot` tinyint(1) NOT NULL,
                          `sort` int(10) unsigned NOT NULL,
                          `open` datetime NOT NULL,
                          `created_at` datetime DEFAULT NULL,
                          `updated_at` datetime DEFAULT NULL,
                          `deleted_at` datetime DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `user` (
                        `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                        `passport` varchar(32) NOT NULL,
                        `password` varchar(32) NOT NULL,
                        `name` varchar(128) NOT NULL,
                        `avatar` varchar(256) NOT NULL,
                        `role_ids` varchar(1024) NOT NULL,
                        `email` varchar(128) NOT NULL,
                        `phone` varchar(32) NOT NULL,
                        `channel` varchar(256) NOT NULL,
                        `created_at` datetime DEFAULT NULL,
                        `updated_at` datetime DEFAULT NULL,
                        `deleted_at` datetime DEFAULT NULL,
                        `token` varchar(100) NOT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `user_uk` (`passport`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;