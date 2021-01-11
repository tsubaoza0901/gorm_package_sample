
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- メモ：company_idを使用しない場合は削除してgoose up
CREATE TABLE IF NOT EXISTS users (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(128) NOT NULL COMMENT "氏名",
    age int(11) NOT NULL COMMENT "年齢",
    company_id int(11) COMMENT "カンパニーID",
    PRIMARY KEY(id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;