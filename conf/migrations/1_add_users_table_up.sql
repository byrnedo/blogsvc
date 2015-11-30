create table users (
  "id" SERIAL,
  "first_name" varchar(255),
  "last_name" varchar(255),
  "email" varchar(255),
  "alias" varchar(255),
  "creation_time" timestamp with time zone,
  "update_time" timestamp with time zone
);
