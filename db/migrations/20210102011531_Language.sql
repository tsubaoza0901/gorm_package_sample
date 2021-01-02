
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS languages (
    id int(11) NOT NULL AUTO_INCREMENT,
    language_name varchar(128) NOT NULL COMMENT "言語名",
    PRIMARY KEY(id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE languages;