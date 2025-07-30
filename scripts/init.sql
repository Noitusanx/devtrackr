-- Create database if it doesn't exist
SELECT 'CREATE DATABASE tracker_db_test' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'tracker_db_test')\gexec

-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Set timezone
SET timezone = 'Asia/Jakarta';