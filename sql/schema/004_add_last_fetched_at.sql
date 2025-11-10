-- +goose Up
-- +goose StatementBegin

ALTER TABLE feeds 
ADD COLUMN IF NOT EXISTS last_fetched_at TIMESTAMP;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
