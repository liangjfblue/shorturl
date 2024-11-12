CREATE TABLE tb_shortUrl_0 (
    id bigint PRIMARY KEY,
    short_url VARCHAR(7) NOT NULL,
    long_url VARCHAR(64) NOT NULL,
    short_type VARCHAR(5) NOT NULL DEFAULT '0',
    timestamp DATETIME NOT NULL,
    UNIQUE KEY `uqi_short` (`short_url`)
) COMMENT '长短链映射' ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci;

create table tb_shortUrl_1 (
    id bigint PRIMARY KEY,
    short_url VARCHAR(7) NOT NULL,
    long_url VARCHAR(64) NOT NULL,
    short_type VARCHAR(5) NOT NULL DEFAULT '0',
    timestamp DATETIME NOT NULL,
    UNIQUE KEY `uqi_short` (`short_url`)
) COMMENT '长短链映射' ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci;

create table tb_shortUrl_2 (
   id bigint PRIMARY KEY,
   short_url VARCHAR(7) NOT NULL,
   long_url VARCHAR(64) NOT NULL,
   short_type VARCHAR(5) NOT NULL DEFAULT '0',
   timestamp DATETIME NOT NULL,
   UNIQUE KEY `uqi_short` (`short_url`)
) COMMENT '长短链映射' ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE=utf8mb4_general_ci;



