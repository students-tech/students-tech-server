CREATE TABLE IF NOT EXISTS question(
   id SERIAL PRIMARY KEY,
   students_project_assignment_id INT NOT NULL,
   question TEXT,
   question_number SMALLINT,
   is_correct BOOLEAN,
   created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS answer(
    id SERIAL PRIMARY KEY,
    question_id INT NOT NULL,
    answer TEXT,
    is_correct BOOLEAN,
    is_picked BOOLEAN
);