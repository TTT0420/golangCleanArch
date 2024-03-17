
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNIQUE COMMENT 'ユーザー固有ID',
    user_name VARCHAR(255) COMMENT 'ユーザー名',
    user_type TINYINT(1)  DEFAULT 1 COMMENT 'ユーザータイプ（1:通常,2:VIP）',
    is_deleted TINYINT(1) DEFAULT 0  COMMENT '削除ステータス（0:未削除,1:削除）',
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_date DATETIME ON UPDATE CURRENT_TIMESTAMP
)COMMENT 'ユーザー管理テーブル';


CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT COMMENT 'ユーザー固有ID',
    title VARCHAR(255) COMMENT '投稿タイトル',
    content VARCHAR(255) COMMENT '投稿内容',
    status TINYINT(1) COMMENT '投稿ステータス（1:下書き,2:非公開,3:限定公開,4:公開）',
    is_deleted TINYINT(1) DEFAULT 0  COMMENT '削除ステータス（0:未削除,1:削除）',
    created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_date DATETIME ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE

)COMMENT '投稿管理テーブル';

