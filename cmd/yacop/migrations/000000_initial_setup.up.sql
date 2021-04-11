CREATE TYPE fuel_type_enum AS ENUM (
    'PETROL'
    ,'DIESEL'
    ,'ELECTRICITY'
    ,'NATURAL_GAS'
    ,'HYBRID_PETROL'
    ,'HYBRID_DIESEL'
    ,'HYDROGEN'
    );

CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE vehicles
(
    id                          uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name                   VARCHAR(255)   NOT NULL,
    manufacturer_id             uuid           NOT NULL,
    fuel_type                   fuel_type_enum NOT NULL,
    maximum_kilometers_per_hour INTEGER        NOT NULL,
    maximum_kilowatts           INTEGER        NOT NULL,
    weight_in_kilograms         INTEGER        NOT NULL,
    created_at                  TIMESTAMPTZ      DEFAULT NULL,
    updated_at                  TIMESTAMPTZ      DEFAULT NULL,
    deleted_at                  TIMESTAMPTZ      DEFAULT NULL
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