CREATE TABLE IF NOT EXISTS player (
  id BIGSERIAL PRIMARY KEY,
  avatar CHARACTER VARYING(255),
  avatar_full CHARACTER VARYING(255),
  avatar_medium CHARACTER VARYING(255),
  country_code CHARACTER VARYING(255),
  creation_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  game_count INTEGER DEFAULT 0,
  identifier CHARACTER VARYING(255) NOT NULL,
  last_updated TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  name CHARACTER VARYING(255),
  CONSTRAINT uk_identifier UNIQUE (identifier)
);