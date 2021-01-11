
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS credit_cards (
    id int(11) NOT NULL AUTO_INCREMENT,
    card_number varchar(128) NOT NULL COMMENT "カード番号",
    user_id int(11) NOT NULL COMMENT "ユーザーID",
    PRIMARY KEY(id)
) ENGINE=InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE credit_cards;
