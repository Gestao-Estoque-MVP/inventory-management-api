CREATE TABLE users_permissions (
    id serial PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES users(id),
    permission_id VARCHAR(255) REFERENCES permissions(id)
);