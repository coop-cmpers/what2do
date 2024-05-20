-- 3. Creates a table for events

CREATE TABLE IF NOT EXISTS events (
    id UUID NOT NULL PRIMARY KEY,
    name TEXT,
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    location TEXT
);