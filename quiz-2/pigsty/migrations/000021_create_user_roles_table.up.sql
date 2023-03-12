CREATE TABLE IF NOT EXISTS pig_roles (
  user_role_id bigserial PRIMARY KEY,
  user_id bigserial,
  role text NO NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

