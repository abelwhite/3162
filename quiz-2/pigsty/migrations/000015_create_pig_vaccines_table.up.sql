CREATE TABLE IF NOT EXISTS pig_vaccines (
  pig_vaccine_id bigserial PRIMARY KEY,
  pig_id bigserial,
  vaccine_id bigserial,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

