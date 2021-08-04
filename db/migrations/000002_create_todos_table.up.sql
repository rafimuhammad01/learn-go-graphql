CREATE TABLE todos (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR (55) NOT NULL,
    "description" VARCHAR (255),
    "user_id" bigserial NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now())
);

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");