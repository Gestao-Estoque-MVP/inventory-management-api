CREATE TABLE users_permissions (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES users(id),
    permission_id VARCHAR(255) REFERENCES permissions(id)
);