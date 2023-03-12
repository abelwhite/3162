CREATE TABLE IF NOT EXISTS sties (
  sty_id bigserial PRIMARY KEY,
  sty_type text NOT NULL,
  pig_id bigserial,
  water_bin_id bigserial,
  feed_bin_id bigserial,
  room_id bigserial,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
