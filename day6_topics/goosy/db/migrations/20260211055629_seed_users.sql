-- +goose Up 
-- +goose StatementBegin 
INSERT INTO users (name, email) VALUES 
  ('goosy1', 'goosy1@mail.com'),
  ('goosy2', 'goosy2@mail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM users WHERE email IN ('goosy1@mail.com', 'goosy2@mail.com');

-- +goose StatementEnd
