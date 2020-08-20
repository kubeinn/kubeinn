-- Create standard tables
CREATE TABLE IF NOT EXISTS villages (
    villageID SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    details TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS pilgrims (
    pilgrimID SERIAL PRIMARY KEY,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL,
    villageID INTEGER,
    CONSTRAINT fk_village FOREIGN KEY(villageID) REFERENCES villages(villageID)
);
CREATE TABLE IF NOT EXISTS innkeepers (
    innkeeperID SERIAL PRIMARY KEY,
    email VARCHAR(30) NOT NULL,
    passwd CHAR(60) NOT NULL
);
CREATE TABLE IF NOT EXISTS projects (
    projectID SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    details TEXT NOT NULL,
    cpu INTEGER NOT NULL,
    memory INTEGER NOT NULL,
    storage INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS usage (
    projectID INTEGER NOT NULL,
    pilgrimID INTEGER NOT NULL,
    startTime TIMESTAMPTZ NOT NULL,
    endTime TIMESTAMPTZ NOT NULL,
    cpuMinutesUsed INTEGER NOT NULL,
    memoryMinutesUsed INTEGER NOT NULL,
    PRIMARY KEY(projectID, pilgrimID)
);