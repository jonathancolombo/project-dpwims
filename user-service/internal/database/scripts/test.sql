INSERT INTO users (username, password, password_salt, email, fiscal_code, telephone, role)
VALUES
    ('admin', 'admin123', 'salt1', 'admin@example.com', 'RSSMRA80A01H501U', '3331112222', 'admin'),
    ('jonathan', 'password1', 'salt2', 'jonathan@example.com', 'BNCLNZ90B15F205X', '3332223333', 'user'),
    ('maria', 'password2', 'salt3', 'maria@example.com', 'VRDLGI85C60H501Z', '3334445555', 'user'),
    ('luca', 'password3', 'salt4', 'luca@example.com', 'FRNGPP92D10F205Y', '3336667777', 'user'),
    ('testuser', 'test123', 'salt5', 'testuser@example.com', 'PLLMRA70A01H501T', '3338889999', 'user');
