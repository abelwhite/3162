CREATE TABLE IF NOT EXISTS vaccines (
  vaccine_id bigserial PRIMARY KEY,
  vaccine text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

