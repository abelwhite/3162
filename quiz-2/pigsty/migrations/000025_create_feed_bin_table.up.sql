CREATE TABLE IF NOT EXISTS feed_bin (
  feed_bin_id bigserial PRIMARY KEY,
  status text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

