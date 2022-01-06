CREATE TABLE IF NOT EXISTS construction (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamptz default now() NOT NULL,
    updated_at timestamptz NULL,
    name text NOT NULL
)