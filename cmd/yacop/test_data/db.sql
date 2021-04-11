-- test prep
truncate table vehicles cascade;
truncate table manufacturers cascade;

-- demo data
INSERT INTO manufacturers (id, name)
VALUES ('00000000-0000-0000-0000-000000000010', 'Audi');

INSERT INTO vehicles (id, full_name, manufacturer_id, fuel_type, maximum_kilometers_per_hour, maximum_kilowatts,
                      weight_in_kilograms)
VALUES ('00000000-0000-0000-0000-000000000001', 'A6 2.0 TDI', '00000000-0000-0000-0000-000000000010', 'DIESEL', 228,
        130, 1650);