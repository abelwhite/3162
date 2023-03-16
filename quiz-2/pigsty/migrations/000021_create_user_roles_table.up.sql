CREATE TABLE IF NOT EXISTS user_roles (
  user_role_id bigserial PRIMARY KEY,
  user_id bigserial REFERENCES user(user_id),
  role text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

