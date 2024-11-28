-- +goose Up
-- +goose StatementBegin
alter table accounts add "user_id" varchar;
alter table transactions add "user_id" varchar;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table accounts drop column "user_id";
alter table transactions drop column "user_id";
-- +goose StatementEnd
