/* simple psql database set up
 * Edouard Paris
 */

CREATE TABLE task(
  id serial PRIMARY KEY,
  name VARCHAR(50),
  description TEXT,
  progression INTEGER,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

INSERT INTO task (name, description, progression, created_at, updated_at) VALUES(
  'My first task',
  'get up and put on my pants',
  '100',
  '2016-10-03 08:42:42',
  '2016-10-03 08:42:42'
);
