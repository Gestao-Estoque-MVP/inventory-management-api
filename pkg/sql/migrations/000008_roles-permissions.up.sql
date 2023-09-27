CREATE TABLE  roles_permissions (
    id UUID PRIMARY KEY,
    role_id VARCHAR(255) NOT NULL,
    permission_id VARCHAR(255) NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE
);
