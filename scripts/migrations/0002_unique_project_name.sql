ALTER TABLE projects
ADD CONSTRAINT uniq_user_project_name
UNIQUE (user_id, name);