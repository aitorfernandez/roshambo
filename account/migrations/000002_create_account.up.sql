create table if not exists account (
  id varchar primary key,
  created_at bigint not null default (extract(epoch from now() at time zone 'utc')),
  email varchar not null unique,
  last_request_at bigint not null,
  updated_at bigint not null default (extract(epoch from now() at time zone 'utc'))
);

create trigger update_at_timestamp before update on account
  for each row execute procedure trigger_set_timestamp();
