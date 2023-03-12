CREATE TABLE IF NOT EXISTS pig_weight (
  pig_weight_id bigserial PRIMARY KEY,
  pig_id bigserial,
  weight float,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

