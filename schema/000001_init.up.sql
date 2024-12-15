CREATE TABLE class (
    id uuid default gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE student (
    id uuid default gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    birth_day DATE NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE student_class (
    id uuid default gen_random_uuid() PRIMARY KEY,
    student_id uuid NOT NULL,
    class_id uuid NOT NULL,
    FOREIGN KEY (student_id) REFERENCES student(id),
    FOREIGN KEY (class_id) REFERENCES class(id)
);

CREATE TABLE teacher (
    id uuid default gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255)
);

CREATE TABLE subject (
    id uuid default gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE timepair (
    id uuid default gen_random_uuid() PRIMARY KEY,
    start_pair VARCHAR(255) NOT NULL,
    end_pair VARCHAR(255) NOT NULL
);

CREATE TABLE schedule (
    id uuid default gen_random_uuid() PRIMARY KEY,
    date DATE,
    class_id uuid NOT NULL,
    number_pair INT NOT NULL,
    teacher_id uuid NOT NULL,
    subject_id uuid NOT NULL,
    class_room VARCHAR(255) NOT NULL,
    FOREIGN KEY (class_id) REFERENCES class(id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(id),
    FOREIGN KEY (subject_id) REFERENCES subject(id)
);