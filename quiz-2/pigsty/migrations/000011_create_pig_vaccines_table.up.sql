CREATE TABLE IF NOT EXISTS pig_vaccines (
  pig_vaccine_id bigserial PRIMARY KEY,
  pig_id bigserial REFERENCES pigs(pig_id),
  vaccine_id bigserial REFERENCES vaccines(vaccine_id),
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

