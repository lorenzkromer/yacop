CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE vehicles
(
    id              uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name       VARCHAR(255) NOT NULL,
    manufacturer_id uuid         NOT NULL,
    created_at      TIMESTAMPTZ      DEFAULT NULL,
    updated_at      TIMESTAMPTZ      DEFAULT NULL,
    deleted_at      TIMESTAMPTZ      DEFAULT NULL
);

CREATE UNIQUE INDEX ux_vehicles_full_name ON vehicles (full_name) WHERE deleted_at IS NULL;

CREATE TABLE manufacturers
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ      DEFAULT NULL,
    updated_at TIMESTAMPTZ      DEFAULT NULL,
    deleted_at TIMESTAMPTZ      DEFAULT NULL
);

CREATE UNIQUE INDEX ux_manufacturers_name ON manufacturers (name) WHERE deleted_at IS NULL;

ALTER TABLE vehicles
    ADD CONSTRAINT fk_vehicles_manufacturer_id
        FOREIGN KEY (manufacturer_id)
            REFERENCES manufacturers (id);