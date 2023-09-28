CREATE TABLE users_permissions (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    permission_id UUID REFERENCES permissions(id)
);