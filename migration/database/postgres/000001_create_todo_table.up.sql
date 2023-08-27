-- Create transaction --
BEGIN;

-- set time zone --
SET TIME ZONE 'Asia/Bangkok';

-- install Extension UUID --
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE todo_status AS ENUM (
    'draft',
    'in-progress',
    'done'
);

CREATE TABLE "todo" (
  "id" uuid NOT NULL UNIQUE PRIMARY KEY DEFAULT uuid_generate_v4(),
  "task_name" VARCHAR(255) NOT NULL,
  "status" todo_status NOT NULL DEFAULT ('draft'),
  "creator_name" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP
);

ALTER TABLE todo
ADD CONSTRAINT TODO_NAME_UNIQUE UNIQUE (task_name);

COMMIT;