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

INSERT INTO "Pastebin" ("content", "user_id", "password")
VALUES ('Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.', 1, '$2a$08$j9.Fd5z8HGRJpwY/5JLg0uhv/.8HcW4L4x0xB9Dl0v8.ZW6GzU6Tm'),
('Curabitur diam dolor, malesuada quis varius id, suscipit nec dui. Nunc auctor, quam et posuere interdum, nibh ligula convallis turpis, eget fringilla metus orci vel metus. Aenean velit mi, fringilla sit amet ultricies id, rutrum ac magna.', 2, '$2a$08$Mz.DZ5z8HGRJpwY/5JLg0uhv/.8HcW4L4x0xB9Dl0v8.ZW6GzU6Tm'),
('Suspendisse potenti. In hac habitasse platea dictumst. Integer fringilla ultricies ipsum, et ornare tellus venenatis at. Mauris vestibulum, risus in pulvinar finibus, nisi elit gravida erat, eget tincidunt diam nisi et massa.', 3, '$2a$08$Y5.Ee5z8HGRJpwY/5JLg0uhv/.8HcW4L4x0xB9Dl0v8.ZW6GzU6Tm');



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