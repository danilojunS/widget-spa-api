-- USER

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT,
    gravatar TEXT
);

INSERT INTO users (name, gravatar) values ('Colin', 'http://www.gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterid');
INSERT INTO users (name, gravatar) values ('Kyle', 'http://www.gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterid');
INSERT INTO users (name, gravatar) values ('Thomas', 'http://www.gravatar.com/avatar/48009c6a27d25ef0ea03f985d1f186b0?s=64&d=monsterid');
INSERT INTO users (name, gravatar) values ('James', 'http://www.gravatar.com/avatar/9372f138140c8578c82bbc77b2eca602?s=64&d=monsterid');

-- WIDGETS

DROP TABLE IF EXISTS widgets;

CREATE TABLE IF NOT EXISTS widgets(
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT,
    color TEXT,
    price TEXT,
    inventory INT,
    melts BOOL
);

INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Losenoid', 'blue', '9.99', 52, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Rowlow', 'red', '4.00', 7, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Printure', 'green', '5.55', 18, false);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Claster', 'off-white', '12.56', 9, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Pepelexa', 'purple', '0.99', 0, false);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Dropellet', 'speckled', '16.00', 99, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Jeebus', 'depends on the viewing angle', '18.11', 36, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Nodile', 'black', '52.00', 19, false);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Kaloobon', 'white', '8.00', 40, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Bioyino', 'turtle-shell', '90.12', 3, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Dizoolexa', 'magenta', '67.23', 976, false);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('test', 'red', '5.00', 1, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('Foo Foo Widget', 'blue', '5.10', 34, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('123', 'red', '19.99', 42, true);
INSERT INTO widgets (name, color, price, inventory, melts) VALUES ('My other widget', 'red', '19.99', 42, true);
