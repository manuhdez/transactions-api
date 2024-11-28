-- +goose Up
-- +goose StatementBegin
CREATE TABLE "accounts"
(
    id      VARCHAR(100) NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    balance NUMERIC DEFAULT 0,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "accounts";
-- +goose StatementEnd
