-- Filename: migrations/000001_create_pigs_table.up.sql
CREATE TABLE IF NOT EXISTS pigs (
  pig_id bigserial PRIMARY KEY,
  breed text NOT NULL,
  dob date, 
  gender text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
