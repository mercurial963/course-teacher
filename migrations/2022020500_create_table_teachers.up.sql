CREATE TABLE teachers
( teacher_id SERIAL PRIMARY KEY,
  first_name varchar(25) NOT NULL,
  last_name varchar(30) NOT NULL,
  date_of_birth integer NOT NULL
  );