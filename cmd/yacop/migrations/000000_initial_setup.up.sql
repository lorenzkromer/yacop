CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE vehicles
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMPTZ      DEFAULT NULL,
    updated_at TIMESTAMPTZ      DEFAULT NULL,
    deleted_at TIMESTAMPTZ      DEFAULT NULL,
    full_name  VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX ux_vehicles_full_name ON vehicles (full_name) WHERE deleted_at IS NULL;