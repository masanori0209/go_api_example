CREATE DATABASE IF NOT EXISTS api CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE USER IF NOT EXISTS 'user'@'%' IDENTIFIED BY 'user';
GRANT ALL PRIVILEGES ON api.* TO 'user'@'%';
SET session wait_timeout=259200;
SET global max_allowed_packet=10485760;

FLUSH PRIVILEGES;
