CREATE TABLE IF NOT EXISTS game (
  app_id BIGINT NOT NULL,
  icon_url CHARACTER VARYING(255) NOT NULL,
  logo_url CHARACTER VARYING(255) NOT NULL,
  name CHARACTER VARYING(255) NOT NULL,
  CONSTRAINT pk_game_id PRIMARY KEY (app_id)
);