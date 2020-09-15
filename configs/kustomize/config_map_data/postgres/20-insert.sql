INSERT INTO api.villages (title, details)
VALUES ('test-village-1', 'This is an example village.');
INSERT INTO api.pilgrims (email, passwd)
VALUES ('test-email', 'test-password');
INSERT INTO api.innkeepers (email, passwd)
VALUES ('test-email', 'test-password');
INSERT INTO api.projects (title, details, cpu, memory, storage)
VALUES (
        'test-project-1',
        'This is an example project.',
        1,
        2,
        3
    );
INSERT INTO api.tickets (email, topic, details, isOpen)
VALUES (
        'test-email',
        'test-topic',
        'This is an example ticket.',
        TRUE
    );
INSERT INTO api.usage (
        projectID,
        pilgrimID,
        startTime,
        endTime,
        cpuMinutesUsed,
        memoryMinutesUsed
    )
VALUES (
        1,
        1,
        '2016-06-22 19:10:25-07',
        '2016-06-22 19:10:25-07',
        1,
        1
    );