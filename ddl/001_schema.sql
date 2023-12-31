CREATE SCHEMA IF NOT EXISTS cccar;

CREATE TABLE IF NOT EXISTS cccar.passengers (
	"id" uuid PRIMARY KEY NOT null,
	"name" varchar(255) NOT NULL,
	"email" varchar(255) UNIQUE NOT NULL,
	"document" varchar(12) UNIQUE NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS cccar.drivers (
	"id" uuid PRIMARY KEY NOT null,
	"name" varchar(255) NOT NULL,
	"document" varchar(12) UNIQUE NOT NULL,
	"plate_number" varchar(20) NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz DEFAULT now()
)
