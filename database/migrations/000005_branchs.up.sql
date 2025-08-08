-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create branchs table
CREATE TABLE branchs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

-- Index for soft deletes
CREATE INDEX idx_branchs_deleted_at ON branchs(deleted_at);