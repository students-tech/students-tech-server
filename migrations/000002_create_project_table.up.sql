CREATE TABLE IF NOT EXISTS project(
   id SERIAL PRIMARY KEY,
   email_owner VARCHAR(255) NOT NULL,
    phone_owner VARCHAR(255)  NOT NULL,
    project_name VARCHAR(255) NOT NULL,
    raw_project_description TEXT NOT NULL,
    platform TEXT NOT NULL,
    tech_stack TEXT NOT NULL,
    project_objective TEXT NOT NULL,
    is_have_design BOOLEAN NOT NULL,
    budget_range INT NOT NULL,
    timeline VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
