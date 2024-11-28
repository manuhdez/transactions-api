-- +goose Up
-- +goose StatementBegin
CREATE TYPE transaction_type AS ENUM ('deposit', 'withdrawal', 'transfer');

CREATE TABLE transactions
(
    id         SERIAL           NOT NULL,
    account_id VARCHAR(100)     NOT NULL,
    user_id    VARCHAR(100)     NOT NULL,
    amount     NUMERIC          NOT NULL,
    balance    NUMERIC          NOT NULL,
    type       transaction_type NOT NULL,
    date       TIMESTAMP        NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
DROP TYPE transaction_type;
-- +goose StatementEnd
