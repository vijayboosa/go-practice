-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT u_email UNIQUE (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT IF EXISTS u_email;
-- +goose StatementEnd
