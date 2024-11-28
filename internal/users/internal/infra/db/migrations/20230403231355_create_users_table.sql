-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    "id"         varchar(100) NOT NULL UNIQUE,
    "first_name" varchar(50)  NOT NULL,
    "last_name"  varchar(80)  NOT NULL,
    "email"      varchar(80)  NOT NULL UNIQUE,
    "password"   varchar(100) NOT NULL,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
