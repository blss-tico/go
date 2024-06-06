/* create table for users */
CREATE TABLE users (
    id TEXT NOT NULL PRIMARY KEY, 
    name TEXT NOT NULL, 
    age NUMBER NOT NULL, 
    UNIQUE(name)
);

/* select all users */
SELECT id, name, age FROM users;

/* create a new user */
INSERT INTO uses (id, name, age) VALUES ("<id>", "<name>", "<age>");

/* read user by id */
SELECT id, name, age FROM users WHERE id="<id>";

/* update user by id */
UPDATE users SET name="<name>", age="<age>" WHERE id="<id>";

/* delete user by id */
DELETE users WHERE id="<id>";

/* OBS: fields inside <> must to be substituted by values from variables */