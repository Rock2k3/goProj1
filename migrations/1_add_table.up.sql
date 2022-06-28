create table if not exists  tmp (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    title varchar(100) NULL,
    CONSTRAINT specialization_pkey PRIMARY KEY (id)
);