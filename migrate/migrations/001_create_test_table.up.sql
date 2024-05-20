-- 1. Creates a simple test table

CREATE TABLE IF NOT EXISTS test_table (
    id SERIAL NOT NULL PRIMARY KEY,
    random_string TEXT,
    random_integer INTEGER,
    random_datetime TIMESTAMP WITH TIME ZONE
);