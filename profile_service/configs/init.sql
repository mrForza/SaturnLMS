CREATE TABLE IF NOT EXISTS user_profile (
    id UUID PRIMARY KEY,
    first_name VARCHAR(64),
    last_name VARCHAR(64),
    father_name VARCHAR(64),
    age INT,
    gender BOOLEAN,
    about_me VARCHAR(65536),
    interests VARCHAR(65536)
);

CREATE TABLE IF NOT EXISTS student_profile (
    id UUID PRIMARY KEY,
    facultaty_name VARCHAR(64),
    program_name VARCHAR(64),
    group_number INT,
    course_number INT,
    university_name VARCHAR(128),
    profile_id UUID UNIQUE REFERENCES user_profile(id)
);

CREATE TABLE IF NOT EXISTS teacher_profile (
    id UUID PRIMARY KEY,
    education VARCHAR(1024),
    scientific_experience VARCHAR(2048),
    teaching_experience VARCHAR(2048),
    professional_interests VARCHAR(2048),
    achievements VARCHAR(1024),
    languages VARCHAR(64),
    profile_id UUID UNIQUE REFERENCES user_profile(id)
);

CREATE TABLE IF NOT EXISTS administrative_profile (
    id UUID PRIMARY KEY,
    education VARCHAR(1024),
    work_experience VARCHAR(2048),
    achievements VARCHAR(1024),
    languages VARCHAR(64),
    profile_id UUID UNIQUE REFERENCES user_profile(id)
);