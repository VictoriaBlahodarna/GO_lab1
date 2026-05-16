CREATE TABLE IF NOT EXISTS positions (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    min_salary INT NOT NULL,
    max_salary INT NOT NULL
);

-- Додаємо кілька базових посад для тестів та зручної роботи
INSERT INTO positions (name, min_salary, max_salary) VALUES
('Developer', 1000, 8000),
('Manager', 2000, 10000),
('QA', 800, 4000)
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    position_id INT NOT NULL REFERENCES positions(id),
    salary INT NOT NULL
);
