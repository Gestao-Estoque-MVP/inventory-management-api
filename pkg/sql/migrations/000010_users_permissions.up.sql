CREATE TABLE users_permissions (
    id CHAR(36) PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES users(id),
    permission_id VARCHAR(255) REFERENCES permissions(id)
);