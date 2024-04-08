ALTER TABLE users ALTER COLUMN status TYPE text USING status::text;
DROP TYPE user_status;
DROP TABLE users;