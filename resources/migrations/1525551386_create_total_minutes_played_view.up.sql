CREATE OR REPLACE VIEW total_minutes_played AS
  SELECT
    SUM(minutes) as minutes
  FROM all_games_minutes_tracked(to_timestamp(0) at time zone 'utc', current_timestamp at time zone 'utc')
    AS f(game_id BIGINT, minutes BIGINT);