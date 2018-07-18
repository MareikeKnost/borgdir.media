CREATE TABLE user(
ID integer primary key AUTOINCREMENT,
name text,
email text,
password text,
status text DEFAULT "aktiv",
activeUntil text DEFAULT "immer",
image text default "static/images/profilbild.jpg",
role text default "client");


CREATE TABLE borrow(
equipmentID integer,
clientID integer,
borrowedOn text,
returnUntil text,
quantity integer);

CREATE TABLE reserve(
equipmentID integer,
clientID integer,
returnUntil text,
reservedOn text);

CREATE TABLE equipment(
ID integer primary key AUTOINCREMENT,
name text,
description text,
image text,
category text,
content text,
location text,
quantity integer DEFAULT 1,
status text DEFAULT "verfügbar",
maxQuantity integer);



INSERT INTO equipment
(ID, name, description, image, category, content, location, quantity, status, maxQuantity)
VALUES(1, 'Canon Eos 6d', 'Vollformatkamera
  Funktion: WLAN-Fähig
  mit Videofunktion', 'static/images/canon-eos-6d.jpg', 'Kameras', 'nur die Kamera ohne Objektiv', 'Schrank 1', 3, 'verfügbar', 3);
INSERT INTO equipment
(ID, name, description, image, category, content, location, quantity, status, maxQuantity)
VALUES(2, 'Canon Eos 5d Mark 3', 'Vollformatkamera', 'static/images/canon-eos-5d-iii.jpg', 'Kameras', 'Kamera mit kleiner Kameratasche', 'Schrank 1', 2, 'verfügbar', 2);
INSERT INTO equipment
(ID, name, description, image, category, content, location, quantity, status, maxQuantity)
VALUES(3, '1D X Mark II', 'Auflösung: 20,2 MP
Funktion: WLAN-Fähig, Mit Videofunktion
Sensortyp: CMOS', 'static/images/canon-eos-1d-x-mark-ii.jpg', 'Kameras', 'Kamera mit Tragegurt', 'Schrank 2', 1, 'verfügbar', 1);
INSERT INTO equipment
(ID, name, description, image, category, content, location, quantity, status, maxQuantity)
VALUES(4, 'Speedlite 470EX', 'Leitzahl 47 bei ISO 100', 'static/images/Speedlite-470EX-AI.jpg', 'Blitzgeräte', 'Blitz mit kleiner Tasche', 'Schrank 3', 3, 'verfügbar', 3);
INSERT INTO equipment
(ID, name, description, image, category, content, location, quantity, status, maxQuantity)
VALUES(5, 'Rollei-Stativ', 'leichtes Staiv für unterwegs', 'static/images/rollei-statiiv-c5i.jpg', 'Stative', 'Stativ mit Schnellwechelplatte', 'Schrank 4', 3, 'verfügbar', 3);
