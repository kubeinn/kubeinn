-- Create schema
CREATE SCHEMA api;
-- Create sequences
CREATE SEQUENCE api.village_sequence START 1;
CREATE SEQUENCE api.pilgrim_sequence START 1;
CREATE SEQUENCE api.innkeeper_sequence START 1;
-- Create group roles
CREATE ROLE pilgrims;
CREATE ROLE innkeepers;
CREATE ROLE postgrest LOGIN PASSWORD 'pgpassword' NOINHERIT;
-- Create tables
CREATE TABLE IF NOT EXISTS api.innkeepers (
    id TEXT PRIMARY KEY DEFAULT 'innkeeper-' || NEXTVAL('api.innkeeper_sequence'),
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    passwd TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS api.pilgrims (
    id TEXT PRIMARY KEY DEFAULT 'pilgrim-' || NEXTVAL('api.pilgrim_sequence'),
    organization TEXT NOT NULL,
    description TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    passwd TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending'
);
CREATE TABLE IF NOT EXISTS api.projects (
    id SERIAL PRIMARY KEY,
    pilgrimid TEXT NOT NULL DEFAULT current_user,
    title TEXT NOT NULL UNIQUE,
    details TEXT NOT NULL,
    cpu INTEGER NOT NULL,
    memory INTEGER NOT NULL,
    storage INTEGER NOT NULL,
    kube_configuration TEXT
);
CREATE TABLE IF NOT EXISTS api.tickets (
    id SERIAL PRIMARY KEY,
    pilgrimid TEXT NOT NULL DEFAULT current_user,
    email TEXT NOT NULL,
    topic TEXT NOT NULL,
    details TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'Open'
);