-- Create schema
CREATE SCHEMA api;

-- Create sequences
CREATE SEQUENCE pilgrim_sequence START 1;
CREATE SEQUENCE village_sequence START 1;
CREATE SEQUENCE reeve_sequence START 1;
CREATE SEQUENCE innkeeper_sequence START 1;

-- Create standard tables
CREATE TABLE IF NOT EXISTS api.villages (
    id TEXT PRIMARY KEY DEFAULT 'village-' || NEXTVAL('village_sequence'),
    organization VARCHAR(30) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    vic TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS api.pilgrims (
    id TEXT PRIMARY KEY DEFAULT 'pilgrim-' || NEXTVAL('pilgrim_sequence'),
    villageid TEXT NOT NULL,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS api.innkeepers (
    id TEXT PRIMARY KEY DEFAULT 'innkeeper-' || NEXTVAL('innkeeper_sequence'),
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS api.reeves (
    id TEXT PRIMARY KEY DEFAULT 'reeve-' || NEXTVAL('reeve_sequence'),
    villageid TEXT NOT NULL,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS api.projects (
    id SERIAL PRIMARY KEY,
    pilgrimid TEXT NOT NULL DEFAULT current_user,
    villageid TEXT,
    title VARCHAR(30) NOT NULL UNIQUE,
    details TEXT NOT NULL,
    cpu INTEGER NOT NULL,
    memory INTEGER NOT NULL,
    storage INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS api.tickets (
    id SERIAL PRIMARY KEY,
    villageid TEXT NOT NULL DEFAULT current_user,
    email VARCHAR(30) NOT NULL,
    topic VARCHAR(30) NOT NULL,
    details TEXT NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'Open'
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