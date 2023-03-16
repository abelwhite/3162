CREATE TABLE IF NOT EXISTS temperatures (
  temo_id bigserial PRIMARY KEY,
  room_id bigserial REFERENCES rooms(room_id),
  temp_value decimal, 
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

