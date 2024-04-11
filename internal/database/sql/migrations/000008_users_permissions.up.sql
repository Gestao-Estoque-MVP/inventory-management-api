CREATE TABLE users_permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX idx_users_permissions ON users_permissions (user_id, permission_id);