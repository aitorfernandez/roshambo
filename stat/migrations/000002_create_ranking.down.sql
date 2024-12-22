drop trigger if exists trigger_ranking_stat on stat;

drop function if exists trigger_ranking_func;
drop function if exists moves(numeric, numeric);

drop table if exists ranking;
