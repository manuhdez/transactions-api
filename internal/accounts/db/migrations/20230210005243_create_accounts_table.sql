-- +goose Up
-- +goose StatementBegin
CREATE TABLE `accounts`
(
    `id`      varchar(100) NOT NULL,
    `balance` float DEFAULT 0,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `accounts`;
-- +goose StatementEnd
