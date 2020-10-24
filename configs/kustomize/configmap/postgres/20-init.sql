-- Create schema
CREATE SCHEMA api;

-- Create sequences
CREATE SEQUENCE pilgrim_sequence START 1;
CREATE SEQUENCE village_sequence START 1;
CREATE SEQUENCE reeve_sequence START 1;

-- Create standard tables
CREATE TABLE IF NOT EXISTS api.villages (
    id TEXT PRIMARY KEY DEFAULT 'village-' || NEXTVAL('village_sequence'),
    organization VARCHAR(30) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    vic TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS api.pilgrims (
    id TEXT PRIMARY KEY DEFAULT 'pilgrim-' || NEXTVAL('pilgrim_sequence'),
    villageID TEXT NOT NULL,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS api.innkeepers (
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS api.reeves (
    id TEXT PRIMARY KEY DEFAULT 'reeve-' || NEXTVAL('reeve_sequence'),
    villageID TEXT NOT NULL,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS api.projects (
    id SERIAL PRIMARY KEY,
    pilgrimID TEXT NOT NULL DEFAULT current_user,
    villageID TEXT NOT NULL,
    title VARCHAR(30) NOT NULL UNIQUE,
    details TEXT NOT NULL,
    cpu INTEGER NOT NULL,
    memory INTEGER NOT NULL,
    storage INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS api.tickets (
    id SERIAL PRIMARY KEY,
    villageID TEXT NOT NULL DEFAULT current_user,
    email VARCHAR(30) NOT NULL,
    topic VARCHAR(30) NOT NULL,
    details TEXT NOT NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'Open'
);
CREATE TABLE IF NOT EXISTS api.usage (
    id SERIAL PRIMARY KEY,
    projectID INTEGER NOT NULL,
    pilgrimID TEXT NOT NULL,
    startTime TIMESTAMPTZ NOT NULL,
    endTime TIMESTAMPTZ NOT NULL,
    cpuMinutesUsed INTEGER NOT NULL,
    memoryMinutesUsed INTEGER NOT NULL
);