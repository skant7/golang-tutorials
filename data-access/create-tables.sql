DROP TABLE IF EXISTS album;
CREATE TABLE album(
    id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(128) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    PRIMARY KEY(`id`)

);

INSERT INTO album (title,artist,price)
VALUES
    ('Midnight Rider','Allman Brothers',45.00),
    ('Willow','Taylor Swift',60.00),
    ('Better','Caracara',89.00),
    ('Takeaway','Chainsmokers',50.00);
    