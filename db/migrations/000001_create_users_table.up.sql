CREATE TABLE users (
   "id" bigserial PRIMARY KEY,
   "username" VARCHAR(55) NOT NULL UNIQUE,
   "email" VARCHAR(120) NOT NULL UNIQUE,
   "password" VARCHAR(255) NOT NULL,
   "name" VARCHAR(255) NOT NULL,
   "created_at" TIMESTAMP NOT NULL DEFAULT (now())
);