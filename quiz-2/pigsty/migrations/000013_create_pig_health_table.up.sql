CREATE TABLE IF NOT EXISTS pig_health (
  pig_health_id bigserial PRIMARY KEY,
  pig_id bigserial REFERENCES pigs(pig_id),
  health text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

