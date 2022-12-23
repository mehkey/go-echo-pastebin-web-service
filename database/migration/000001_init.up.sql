CREATE TABLE "Users" (
                         "id" SERIAL PRIMARY KEY,
                         "name" varchar,
                         "email" varchar,
                         "password" hashed,
                         "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                         "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO "Users" ("name","email","password") VALUES ('Veronica Smith','Veronica.Smith@gmail.com','XXXX');



CREATE TABLE "Pastebin" (
                         "id" SERIAL PRIMARY KEY,
                         "content" varchar,
                         "user_id" int,
                         "password" hashed,
                         "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                         "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);



ALTER TABLE "Pastebin"
    ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");

/*
Users:
user_id (pk)
user_name
email_address
password (hashed)
creation_date

Pastebin
pastebin_id (pk)
content (1000 words)
password (optional)
creation_date
user_id (creator)

*/