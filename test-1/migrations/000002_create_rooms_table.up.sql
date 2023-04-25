CREATE TABLE if NOT EXISTS rooms(
    room_id bigserial PRIMARY KEY,
    name text NOT NULL,
    no_of_pigsty bigserial,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW ()
);