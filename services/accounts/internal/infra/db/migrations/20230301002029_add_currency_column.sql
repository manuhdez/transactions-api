-- +goose Up
-- +goose StatementBegin
ALTER TABLE accounts
    ADD "currency" VARCHAR(3) DEFAULT 'EUR';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE accounts
    DROP COLUMN currency;
-- +goose StatementEnd
