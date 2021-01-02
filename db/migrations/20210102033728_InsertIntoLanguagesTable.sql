
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO languages (id, language_name) 
VALUES 
    (1, '日本語'),
    (2, '英語'),
    (3, 'フランス語'),
    (4, 'ドイツ語'),
    (5, '中国語');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM languages;