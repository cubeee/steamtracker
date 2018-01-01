CREATE TABLE IF NOT EXISTS game_snapshot (
  id BIGSERIAL PRIMARY KEY ,
  date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  minutes_played INTEGER NOT NULL,
  game_id BIGINT NOT NULL,
  player_id BIGINT NOT NULL,
  CONSTRAINT fk_player_id FOREIGN KEY (player_id) REFERENCES player(id),
  CONSTRAINT fk_game_id FOREIGN KEY (game_id) REFERENCES game(app_id)
);