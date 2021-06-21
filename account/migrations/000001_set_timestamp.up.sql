create or replace function trigger_set_timestamp()
returns trigger as $$
begin
  new.updated_at = (extract(epoch from now() at time zone 'utc'));
  return new;
end;
$$ language plpgsql;
