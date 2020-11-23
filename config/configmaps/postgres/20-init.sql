-- Create schema
CREATE SCHEMA api;

-- Create sequences
CREATE SEQUENCE api.pilgrim_sequence START 1;
CREATE SEQUENCE api.village_sequence START 1;
CREATE SEQUENCE api.reeve_sequence START 1;
CREATE SEQUENCE api.innkeeper_sequence START 1;

-- Create group roles
CREATE ROLE pilgrims;
CREATE ROLE villages;
CREATE ROLE innkeepers;

-- Create standard tables
CREATE TABLE IF NOT EXISTS api.villages (
    id TEXT PRIMARY KEY DEFAULT 'village-' || NEXTVAL('api.village_sequence'),
    organization TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS api.pilgrims (
    id TEXT PRIMARY KEY DEFAULT 'pilgrim-' || NEXTVAL('api.pilgrim_sequence'),
    villageid TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    passwd TEXT NOT NULL,
    regcode TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS api.innkeepers (
    id TEXT PRIMARY KEY DEFAULT 'innkeeper-' || NEXTVAL('api.innkeeper_sequence'),
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    passwd TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS api.reeves (
    id TEXT PRIMARY KEY DEFAULT 'reeve-' || NEXTVAL('api.reeve_sequence'),
    villageid TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    passwd TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending'
);
CREATE TABLE IF NOT EXISTS api.projects (
    id SERIAL PRIMARY KEY,
    pilgrimid TEXT NOT NULL DEFAULT current_user,
    villageid TEXT,
    title TEXT NOT NULL UNIQUE,
    details TEXT NOT NULL,
    cpu INTEGER NOT NULL,
    memory INTEGER NOT NULL,
    storage INTEGER NOT NULL,
    kube_configuration TEXT
);
CREATE TABLE IF NOT EXISTS api.tickets (
    id SERIAL PRIMARY KEY,
    villageid TEXT NOT NULL DEFAULT current_user,
    email TEXT NOT NULL,
    topic TEXT NOT NULL,
    details TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'Open'
);
CREATE TABLE IF NOT EXISTS api.usage (
    id SERIAL PRIMARY KEY,
    projectid INTEGER NOT NULL,
    pilgrimid TEXT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    cpu_minutes_used INTEGER NOT NULL,
    memory_minutes_used INTEGER NOT NULL
);