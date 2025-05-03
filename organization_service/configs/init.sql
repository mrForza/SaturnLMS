CREATE TABLE IF NOT EXISTS university (
    name VARCHAR(64) PRIMARY KEY,
    description VARCHAR(65536),
    legal_address VARCHAR(512),
    actual_address VARCHAR(512),
    inn VARCHAR(10),
    bank_name VARCHAR(64),
    owner_id UUID
);

CREATE TABLE IF NOT EXISTS facultaty (
    name VARCHAR(64) PRIMARY KEY,
    description VARCHAR(65536),
    university_name VARCHAR(64) UNIQUE REFERENCES university(name)
);

CREATE TABLE IF NOT EXISTS program (
    name VARCHAR(64) PRIMARY KEY,
    description VARCHAR(65536),
    type VARCHAR(16),
    languages VARCHAR(64),
    facultaty_name VARCHAR(64) UNIQUE REFERENCES facultaty(name)
);

CREATE TABLE IF NOT EXISTS program_group (
    number INT PRIMARY KEY,
    name VARCHAR(64),
    course_number INT,
    program_name VARCHAR(64) UNIQUE REFERENCES program(name)
);