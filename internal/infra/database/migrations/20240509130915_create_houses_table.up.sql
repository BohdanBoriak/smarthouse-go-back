CREATE TABLE IF NOT EXISTS public.houses
(
    id              serial PRIMARY KEY,
    user_id         integer references public.users(id),
    "name"          text NOT NULL,
    "address"       text NOT NULL,
    lat             double precision NOT NULL,
    lon             double precision NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);
