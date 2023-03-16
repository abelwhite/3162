CREATE TABLE IF NOT EXISTS rooms (
  room_id bigserial PRIMARY KEY,
  room_name text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
