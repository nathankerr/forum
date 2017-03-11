CREATE TABLE users (
  "uuid" UUID PRIMARY KEY,
  "username" VARCHAR(20) NOT NULL,
  "password" CHAR(60) NOT NULL,
  "created" BIGINT NOT NULL,
  "removed" BIGINT NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX on "users" ("username", "password", "removed");

CREATE TABLE tokens (
  "uuid" UUID PRIMARY KEY,
  "user" UUID NOT NULL,
  "token" CHAR(32) NOT NULL,
  "created" BIGINT NOT NULL,
  "removed" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX on "tokens" ("user", "token", "removed");

CREATE TABLE categories (
  "uuid" UUID PRIMARY KEY,
  "name" VARCHAR(128) NOT NULL,
  "position" SMALLINT NOT NULL,
  "removed" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX on "categories" ("removed", "position");

CREATE TABLE boards (
  "uuid" UUID PRIMARY KEY,
  "category" UUID NOT NULL,
  "name" VARCHAR(128) NOT NULL,
  "position" SMALLINT NOT NULL,
  "removed" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX on "boards" ("category", "removed", "position");

CREATE TABLE threads (
  "uuid" UUID PRIMARY KEY,
  "board" UUID NOT NULL,
  "sticky" BOOLEAN NOT NULL DEFAULT FALSE,
  "title" VARCHAR(128) NOT NULL,
  "removed" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX on "threads" ("board", "sticky", "removed");

CREATE TABLE posts (
  "uuid" UUID PRIMARY KEY,
  "thread" UUID NOT NULL,
  "user" UUID NOT NULL,
  "title" VARCHAR(128) NOT NULL,
  "content" TEXT NOT NULL,
  "created" BIGINT NOT NULL,
  "modified" BIGINT NOT NULL,
  "removed" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX on posts ("thread", "removed", "modified");