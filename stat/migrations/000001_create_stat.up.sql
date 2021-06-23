create table if not exists stat (
  id serial primary key,
  account_id varchar not null,
  computer_move smallint not null,
  player_move smallint not null,
  round int not null
);
