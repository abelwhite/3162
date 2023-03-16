CREATE TABLE IF NOT EXISTS sties (
  sty_id bigserial PRIMARY KEY,
  sty_type text NOT NULL,
  pig_id bigserial REFERENCES pigs(pig_id),
  water_bin_id bigserial REFERENCES water_bin(water_bin_id),
  feed_bin_id bigserial REFERENCES feed_bin(feed_bin_id),
  room_id bigserial REFERENCES rooms(room_id),
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
