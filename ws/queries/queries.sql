CREATE TABLE users (
  username VARCHAR(255) PRIMARY KEY,
  password VARCHAR(255),
  email VARCHAR(255)
);

CREATE TABLE history (
  message_id VARCHAR(255),
  message TEXT,
  username VARCHAR(255),
  target VARCHAR(255)[]
);
