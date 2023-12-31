CREATE TABLE IF NOT EXISTS students(
  id SERIAL PRIMARY KEY,
  user_id VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  residency TEXT NOT NULL,
  phone_number VARCHAR(255) NOT NULL,
  linkedin_url TEXT,
  github_username VARCHAR(255),
  gitlab_username VARCHAR(255),
  project_count VARCHAR(31),
  university_name VARCHAR(255) NOT NULL,
  major VARCHAR(255) NOT NULL,
  current_semester SMALLINT NOT NULL,
  relevant_coursework TEXT,
  hours_availability_per_week VARCHAR(255) NOT NULL,
  role TEXT NOT NULL,
  university_email VARCHAR(255) NOT NULL,
  ktm_url TEXT,
  is_student_verified BOOLEAN DEFAULT FALSE,
  resume_url TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
