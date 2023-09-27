CREATE TABLE users_roles (
    id serial PRIMARY KEY,
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);