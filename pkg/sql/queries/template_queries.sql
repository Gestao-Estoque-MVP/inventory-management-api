-- name: CreateTemplate :one
INSERT INTO template_email (id, name, url, description, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name;

-- name: GetTemplateS3 :one
SELECT url FROM template_email WHERE id = $1;

-- name: GetTemplate :one
SELECT * FROM template_email WHERE id = $1;