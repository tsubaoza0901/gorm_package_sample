
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS user_language_relations (
    user_id int(11) NOT NULL COMMENT "ユーザーID",
    language_id int(11) NOT NULL COMMENT "言語ID"
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_language_relations;