create table if not exists ranking (
  id serial primary key,
  account_id varchar not null unique,
  draw smallint not null default 0,
  lose smallint not null default 0,
  total_rounds int not null,
  win smallint not null default 0
);

create or replace function moves(
   player_move numeric,
   computer_move numeric,
   out draw numeric,
   out lose numeric,
   out win numeric
) as $$
begin
  if player_move = computer_move then
    draw := 1;
  elsif (player_move + 1) % 3 = computer_move then
    lose := 1;
  else
    win := 1;
  end if;
end;
$$ language plpgsql;

create or replace function trigger_ranking_func()
returns trigger as $$
begin
  if (TG_OP = 'INSERT') then
    with w as (
      select coalesce((select win from moves(new.player_move, new.computer_move)), 0)
    ), l as (
      select coalesce((select lose from moves(new.player_move, new.computer_move)), 0)
    ), d as (
      select coalesce((select draw from moves(new.player_move, new.computer_move)), 0)
    )
    insert into ranking (account_id, total_rounds, win, lose, draw)
    values (
      new.account_id,
      1,
      (select * from w),
      (select * from l),
      (select * from d)
    )
    on conflict
      (account_id)
    do update set
      total_rounds = (select max(total_rounds) + 1 from ranking where account_id = new.account_id),
      win = (select (max(win) + (select * from w)) from ranking where account_id = new.account_id),
      lose = (select (max(lose) + (select * from l)) from ranking where account_id = new.account_id),
      draw = (select (max(draw) + (select * from d)) from ranking where account_id = new.account_id)
    ;
    return new;
  end if;
end;
$$ language plpgsql;

create trigger trigger_ranking_stat after insert on stat
  for each row execute function trigger_ranking_func();
