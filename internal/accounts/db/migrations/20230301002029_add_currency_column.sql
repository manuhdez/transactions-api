-- +goose Up
-- +goose StatementBegin
ALTER TABLE `transactions-api`.`accounts`
    ADD `currency` varchar(3) DEFAULT 'EUR';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `transactions-api`.`accounts`
    DROP COLUMN `currency`;
-- +goose StatementEnd
