-- +goose Up
-- +goose StatementBegin
CREATE TABLE "accounts"
(
    "id"      varchar NOT NULL,
    "balance" NUMERIC DEFAULT 0,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "accounts";
-- +goose StatementEnd
