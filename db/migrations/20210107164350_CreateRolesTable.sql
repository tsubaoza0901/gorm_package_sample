
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS roles (
    id int(11) NOT NULL AUTO_INCREMENT,
    role_name varchar(128) NOT NULL COMMENT "担当名",
    user_id int(11) NOT NULL COMMENT "ユーザーID",
    PRIMARY KEY(id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE roles;