CREATE TABLE IF NOT EXISTS crypto(
     ID SERIAL PRIMARY KEY,
     NAME varchar(250),
     AMOUNT_OWNED double precision,
     image_src text
);

INSERT INTO crypto (NAME, AMOUNT_OWNED, image_src)
    VALUES ('testcoin', 23.6, 'www.fake.com');