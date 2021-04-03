-- test prep
truncate table vehicles cascade;
truncate table manufacturers cascade;

-- demo data
INSERT INTO vehicles (id, full_name)
VALUES ('00000000-0000-0000-0000-000000000001', 'Fast Test Car');

INSERT INTO manufacturers (id, name)
VALUES ('00000000-0000-0000-0000-000000000010', 'Audi');