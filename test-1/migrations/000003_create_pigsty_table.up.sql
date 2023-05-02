CREATE TABLE if NOT EXISTS pigsty(
    pigsty_id bigserial PRIMARY KEY,
    room text NOT NULL,
    name text NOT NULL,
    num_of_pigs bigserial,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW ()
);
