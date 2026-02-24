USE tickets_db;
START TRANSACTION;

-- ============================
--  TICKETS (10 records)
-- ============================

INSERT INTO tickets (uuid, user_id, train_uuid, schedule_id, seat_number, price, status)
VALUES
  -- Schedule 1 (Roma → Milano, price 50)
  ('aaa11111-1111-1111-1111-111111111111', 1, '11111111-1111-1111-1111-111111111111', 1, 'A01', 50, 'issued'),
  ('aaa22222-2222-2222-2222-222222222222', 2, '11111111-1111-1111-1111-111111111111', 1, 'A02', 50, 'booked'),
  ('aaa33333-3333-3333-3333-333333333333', 3, '11111111-1111-1111-1111-111111111111', 1, 'A03', 50, 'cancelled'),

  -- Schedule 2 (Milano → Napoli, price 70)
  ('bbb11111-1111-1111-1111-111111111111', 4, '22222222-2222-2222-2222-222222222222', 2, 'B01', 70, 'issued'),
  ('bbb22222-2222-2222-2222-222222222222', 5, '22222222-2222-2222-2222-222222222222', 2, 'B02', 70, 'issued'),
  ('bbb33333-3333-3333-3333-333333333333', 1, '22222222-2222-2222-2222-222222222222', 2, 'B03', 70, 'booked'),
  ('bbb44444-4444-4444-4444-444444444444', 2, '22222222-2222-2222-2222-222222222222', 2, 'B04', 70, 'cancelled'),

  -- Schedule 3 (Bologna → Firenze, price 30)
  ('ccc11111-1111-1111-1111-111111111111', 3, '33333333-3333-3333-3333-333333333333', 3, 'C01', 30, 'issued'),
  ('ccc22222-2222-2222-2222-222222222222', 4, '33333333-3333-3333-3333-333333333333', 3, 'C02', 30, 'booked'),
  ('ccc33333-3333-3333-3333-333333333333', 5, '33333333-3333-3333-3333-333333333333', 3, 'C03', 30, 'issued')
ON DUPLICATE KEY UPDATE
  user_id = VALUES(user_id),
  train_uuid = VALUES(train_uuid),
  schedule_id = VALUES(schedule_id),
  seat_number = VALUES(seat_number),
  price = VALUES(price),
  status = VALUES(status);

-- ============================
--  PAYMENTS (10 records)
--  Only for tickets with status = 'issued'
-- ============================

INSERT INTO payments (ticket_id, amount, payment_method, provider_reference)
VALUES
  ('aaa11111-1111-1111-1111-111111111111', 50, 'credit_card', 'PAY-AAA-001'),
  ('bbb11111-1111-1111-1111-111111111111', 70, 'credit_card', 'PAY-BBB-001'),
  ('bbb22222-2222-2222-2222-222222222222', 70, 'paypal', 'PAY-BBB-002'),
  ('ccc11111-1111-1111-1111-111111111111', 30, 'credit_card', 'PAY-CCC-001'),
  ('ccc33333-3333-3333-3333-333333333333', 30, 'paypal', 'PAY-CCC-002'),

  -- Extra payments for variety
  ('bbb11111-1111-1111-1111-111111111111', 70, 'credit_card', 'PAY-BBB-003'),
  ('aaa11111-1111-1111-1111-111111111111', 50, 'paypal', 'PAY-AAA-002'),
  ('ccc33333-3333-3333-3333-333333333333', 30, 'credit_card', 'PAY-CCC-003'),
  ('bbb22222-2222-2222-2222-222222222222', 70, 'credit_card', 'PAY-BBB-004'),
  ('ccc11111-1111-1111-1111-111111111111', 30, 'paypal', 'PAY-CCC-004');

COMMIT;
