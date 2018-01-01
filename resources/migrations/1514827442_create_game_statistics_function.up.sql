CREATE OR REPLACE FUNCTION all_games_minutes_tracked(dateFrom TIMESTAMP, dateTo TIMESTAMP)
  RETURNS SETOF RECORD AS $$
DECLARE
  rec RECORD;
BEGIN
  FOR rec IN (
    WITH aggregate AS (
        SELECT
          player_id,
          game_id,
          minutes_played - lag(minutes_played) OVER (PARTITION BY game_id, player_id ORDER BY date ASC) AS diff,
          date
        FROM game_snapshot AS snapshot_diff
    )
    SELECT game_id, SUM(agg.diff) AS sum
    FROM aggregate agg
    WHERE agg.diff > 0 AND agg.date >= dateFrom AND agg.date <= dateTo
    GROUP BY game_id
    ORDER BY sum DESC) LOOP
    RETURN NEXT rec;
  END LOOP;
  RETURN;
END;
$$ LANGUAGE plpgsql;