CREATE TABLE IF NOT EXISTS pig_health (
  pig_health_id bigserial PRIMARY KEY,
  pig_id bigserial,
  healt text NO NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

