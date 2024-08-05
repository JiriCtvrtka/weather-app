/*
USERS

	First name
	Last name
    Username
    Password
    Email
	City
    ZIP codeStreet 
    Number
    Addition info
    Age

*/

CREATE TABLE users (
  firstname VARCHAR(255),
  lastname VARCHAR(255),
  username VARCHAR(255) PRIMARY KEY,
  password VARCHAR(255),
  email VARCHAR(255),
  city VARCHAR(255),
  zipcode INT,
  street VARCHAR(255),
  number VARCHAR(255),
  additional_info VARCHAR(255),
  age INT
);


INSERT INTO users (firstname,
  lastname,
  username,
  password,
  email,
  city,
  zipcode,
  street,
  number,
  additional_info,
  age)
VALUES 
	('Jane', 'Doe', 'janedoe1','password1', 'jane@yahoo.com', 'New York', 223423, 'Street1', '24A','',27),
  ('John', 'Doe', 'johndoe', 'password2', 'john@yahoo.com', 'New York', 223423, 'Street1', '24A','',27);


/*
PRODUCTS

	ID 
	Name
	Count 
	Description
	Price
	Currency

*/


CREATE TABLE products (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255),
  description VARCHAR(255),
  currency VARCHAR(255),
  count INT,
  price FLOAT
);


INSERT INTO products (id,
  name,
  description,
  currency,
  count,
  price)
VALUES 
	('P432','TV1', 'desc1', 'euro', 4, 78.9),
  ('P433','TV2', 'desc2', 'euro', 7, 100.56);


DROP TABLE products_json;
CREATE TABLE products_json (
  id VARCHAR(255) PRIMARY KEY,
  values JSONB
);
INSERT INTO products_json (id, values) VALUES
('P432', '{"name": "TV1", "description": "desc3", "currency": "euro", "count": 5, "price": 5.4}');
SELECT values->>'name' as XXX FROM products_json;

/*
ORDERS

	ID
	Username
	Items -> [product_id:count, …] (json) { “product_id”: 5 }
	Status
	Delivery
	Delivery price
	Total price
	Currency


*/

CREATE TABLE orders (
  id VARCHAR(255) PRIMARY KEY,
  username VARCHAR(255),
  items JSONB,
  status VARCHAR(255),
  delivery VARCHAR(255),
  delivery_price FLOAT,
  total_price FLOAT,
  currency VARCHAR(255)
);


INSERT INTO orders (id,
  username,
  items,
  status,
  delivery,
  delivery_price,
  total_price,
  currency)              
VALUES 
	('O1','janedoe1', '{"P432": 1}', 'In Progress', 'Easybox', 10, 700, 'euro'),
  ('O2','johndoe', '{"P432": 1,"P433": 1}', 'Finished', 'DPD', 10, 999.9, 'usd');
