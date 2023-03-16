CREATE TABLE IF NOT EXISTS humidity (
  humid_id bigserial PRIMARY KEY,
  room_id bigserial REFERENCES rooms(room_id),
  humid_value decimal, 
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

