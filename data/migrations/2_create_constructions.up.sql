CREATE TABLE IF NOT EXISTS construction (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    company_id uuid NOT NULL,
    created_at timestamptz default now() NOT NULL,
    updated_at timestamptz NULL,
    name text NOT NULL,
    initial_date timestamptz NOT NULL,
    due_date timestamptz NOT NULL,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id)
            REFERENCES company(id)
)