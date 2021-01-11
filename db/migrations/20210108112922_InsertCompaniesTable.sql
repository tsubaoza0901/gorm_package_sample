
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO companies (id, name) 
VALUES 
    (1, 'A社'),
    (2, 'B社'),
    (3, 'C社'),
    (4, 'D社'),
    (5, 'E社');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM companies;
