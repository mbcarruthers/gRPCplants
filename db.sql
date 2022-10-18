
CREATE DATABASE native;
CREATE TABLE native.plants(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    common_name STRING NOT NULL,
    genus STRING NOT NULL,
    species STRING NOT NULL
);

INSERT INTO native.plants(common_name, genus, species)
VALUES('Salt-marsh Morning Glory','Ipomoea','sagittata');

INSERT INTO native.plants(common_name, genus, species)
VALUES('Beach Morning Glory','Ipomoea','pes-caprae');

INSERT INTO native.plants(common_name, genus, species)
VALUES('Semaphore Thoroughwart','Eupatorium','mikaniodes');

INSERT INTO native.plants(common_name, genus, species)
VALUES('Seabeach Evening-primrose','Oenothera','humifusa');

INSERT INTO native.plants(common_name, genus, species)
VALUES('Dune Marsh-Elder','Iva','imbricata');