-- +goose Up
-- +goose StatementBegin
CREATE TABLE `transactions`
(
    `id`         INT auto_increment NOT NULL,
    `account_id` varchar(100) NOT NULL,
    `amount`     float        NOT NULL,
    `balance`    float        NOT NULL,
    `type`       enum ('deposit', 'withdrawal') NOT NULL,
    `date`       timestamp    NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `transactions`
-- +goose StatementEnd
