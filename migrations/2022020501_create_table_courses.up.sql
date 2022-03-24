CREATE TABLE courses
( course_id SERIAL PRIMARY KEY,
  course_name varchar(25) NOT NULL,
  course_description varchar(30) NOT NULL,
  teacher_id int REFERENCES teachers (teacher_id)
  );