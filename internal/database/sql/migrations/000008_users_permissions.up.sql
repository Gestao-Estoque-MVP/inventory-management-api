CREATE TABLE users_permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    permission_id UUID REFERENCES permissions(id)
);