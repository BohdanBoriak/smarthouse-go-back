CREATE TABLE IF NOT EXISTS public.devices
(
    id              serial PRIMARY KEY,
    user_id         integer references public.users(id),
    house_id        integer references public.houses(id),
    "name"          text NOT NULL,
    "model"         text NOT NULL,
    "type"          varchar(50) NOT NULL,
    "description"   text,
    "units"         text NOT NULL,
    "uuid"          text NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);
