CREATE TABLE if NOT EXISTS pigsty(
    pigsty_id bigserial PRIMARY KEY,
    room_id bigserial,
    name text NOT NULL,
    no_of_pigs bigserial,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW ()
);
