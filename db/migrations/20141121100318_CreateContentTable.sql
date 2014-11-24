
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table blocks
(
	content_id serial,
	content text not null,
	key text not null UNIQUE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table blocks;