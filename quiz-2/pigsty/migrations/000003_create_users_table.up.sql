CREATE TABLE IF NOT EXISTS users (
  user_id bigserial PRIMARY KEY,
  firstname text NOT NULL,
  lastname text NOT NULL,
  email text NOT NULL, 
  password text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
