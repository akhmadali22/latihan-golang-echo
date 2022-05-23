DROP TABLE IF EXISTS siswa;
CREATE SEQUENCE siswa_id START 1;
CREATE TABLE siswa (
    id serial PRIMARY KEY,
    nama VARCHAR(255),
    kelas INTEGER,
    jenjang VARCHAR(255), 
    smt INTEGER
);