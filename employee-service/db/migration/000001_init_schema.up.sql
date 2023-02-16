CREATE TABLE "employee_data" (
    "id" bigserial  PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "primaryTechStack" varchar NOT NULL,
    "secondaryTechStack" varchar NOT NULL
);