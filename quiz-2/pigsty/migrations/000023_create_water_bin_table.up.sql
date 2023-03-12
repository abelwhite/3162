CREATE TABLE IF NOT EXISTS water_bin (
  water_bin_id bigserial PRIMARY KEY,
  status text NO NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

