CREATE TABLE IF NOT EXISTS students_project_assignment(
    id SERIAL PRIMARY KEY,
    students_id INT NOT NULL,
    project_id INT NOT NULL,
    score INT,
    is_passed BOOLEAN,
    created_at TIMESTAMP DEFAULT NOW()
);

