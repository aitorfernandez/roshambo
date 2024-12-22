create table if not exists profile (
  id varchar primary key,
  account_id varchar not null unique,
  avatar varchar,
  created_at bigint not null default (extract(epoch from now() at time zone 'utc')),
  updated_at bigint not null default (extract(epoch from now() at time zone 'utc')),
  username varchar not null unique
);

create trigger update_at_timestamp before update on profile
  for each row execute procedure trigger_set_timestamp();
