USE identity_users;

INSERT INTO users (username, password, password_salt, email, fiscal_code, telephone, role)
VALUES ('mrossi', '$2a$10$A1B2C3D4E5F6G7H8I9J0K', 'salt123', 'marco.rossi@example.com', 'RSSMRC90A01H501U',
        '3331112222', 'customer'),
       ('lbianchi', '$2a$10$Z9Y8X7W6V5U4T3S2R1Q0P', 'salt456', 'laura.bianchi@example.com', 'BNCLRA85C41F205Z',
        '3332223333', 'admin'),
       ('gverdi', '$2a$10$Q1W2E3R4T5Y6U7I8O9P0A', 'salt789', 'giuseppe.verdi@example.com', 'VRDGPP70D15H501T',
        '3333334444', 'customer'),
       ('sconti', '$2a$10$M1N2B3V4C5X6Z7L8K9J0H', 'saltABC', 'sara.conti@example.com', 'CNTSSR95E62D612X', '3334245555',
        'customer'),
       ('fneri', '$2a$10$H1G2F3D4S5A6Q7W8E9R0T', 'saltXYZ', 'franco.neri@example.com', 'NRIFNC88H20A794Q', '3334445555',
        'customer');