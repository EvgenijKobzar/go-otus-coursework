-- +goose Up
-- +goose StatementBegin
alter table movies_online.seasons add moderated    boolean default false;
alter table movies_online.episodes add moderated    boolean default false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table movies_online.seasons drop column moderated;
alter table movies_online.episodes drop column moderated;
-- +goose StatementEnd
