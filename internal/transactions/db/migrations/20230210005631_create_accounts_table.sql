-- +goose Up
-- +goose StatementBegin
CREATE TABLE `accounts`
(
    `id`      varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `accounts`;
-- +goose StatementEnd
