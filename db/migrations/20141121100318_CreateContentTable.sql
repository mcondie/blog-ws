
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table blocks
(
	content_id int,
	content text,
	name text
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table blocks;