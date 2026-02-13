-- File di test: test_trains.sql
-- Scopo: inserire dati di esempio coerenti con le tabelle definite in init_trains.sql
-- Ordine degli insert: stations -> trains -> routes -> schedules -> schedules_stops
-- Gli INSERT sono scritti in modo idempotente (ON DUPLICATE KEY UPDATE) per poter essere rieseguiti senza errori.
USE trains_db;
START TRANSACTION;

-- 1) Stations (inserisco id espliciti per poterli referenziare facilmente)
INSERT INTO stations (id, name, city, region, status)
VALUES
  (1, 'Roma Termini', 'Roma', 'Lazio', 'active'),
  (2, 'Milano Centrale', 'Milano', 'Lombardia', 'active'),
  (3, 'Napoli Centrale', 'Napoli', 'Campania', 'active'),
  (4, 'Firenze Santa Maria Novella', 'Firenze', 'Toscana', 'active'),
  (5, 'Bologna Centrale', 'Bologna', 'Emilia-Romagna', 'active')
ON DUPLICATE KEY UPDATE
  city = VALUES(city), region = VALUES(region), status = VALUES(status), name = VALUES(name);

-- 2) Trains (uuid espliciti)
INSERT INTO trains (uuid, train_number, type, capacity, status)
VALUES
  ('11111111-1111-1111-1111-111111111111', 'TR-100', 'regional', 300, 'active'),
  ('22222222-2222-2222-2222-222222222222', 'TR-200', 'intercity', 600, 'active'),
  ('33333333-3333-3333-3333-333333333333', 'TR-300', 'highspeed', 800, 'inactive')
ON DUPLICATE KEY UPDATE
  train_number = VALUES(train_number), type = VALUES(type), capacity = VALUES(capacity), status = VALUES(status);

-- 3) Routes (imposto id espliciti per chiarezza)
INSERT INTO routes (id, train_id, departure_station, arrival_station, departure_time, arrival_time, distance)
VALUES
  (1, '11111111-1111-1111-1111-111111111111', 'Roma Termini', 'Milano Centrale', '2026-02-11 07:00:00', '2026-02-11 11:30:00', 570),
  (2, '22222222-2222-2222-2222-222222222222', 'Milano Centrale', 'Napoli Centrale', '2026-02-11 08:00:00', '2026-02-11 14:00:00', 760),
  (3, '33333333-3333-3333-3333-333333333333', 'Bologna Centrale', 'Firenze Santa Maria Novella', '2026-02-11 09:00:00', '2026-02-11 10:15:00', 120)
ON DUPLICATE KEY UPDATE
  train_id = VALUES(train_id), departure_station = VALUES(departure_station), arrival_station = VALUES(arrival_station), departure_time = VALUES(departure_time), arrival_time = VALUES(arrival_time), distance = VALUES(distance);

-- 4) Schedules (id espliciti per poterli usare in schedules_stops)
INSERT INTO schedules (id, train_id, station_id, departure_station, arrival_station, status, price)
VALUES
  (1, '11111111-1111-1111-1111-111111111111', 1, 'Roma Termini', 'Milano Centrale', 'active', 50),
  (2, '22222222-2222-2222-2222-222222222222', 2, 'Milano Centrale', 'Napoli Centrale', 'active', 70),
  (3, '33333333-3333-3333-3333-333333333333', 5, 'Bologna Centrale', 'Firenze Santa Maria Novella', 'active', 30)
ON DUPLICATE KEY UPDATE
  train_id = VALUES(train_id), station_id = VALUES(station_id), departure_station = VALUES(departure_station), arrival_station = VALUES(arrival_station), status = VALUES(status), price = VALUES(price);

-- 5) Schedule stops (coerenti con station_id e name)
-- Schedule 1 (Roma -> Bologna -> Milano)
INSERT INTO schedules_stops (id, schedule_id, station_id, station_name, stop_order, arrival_time, departure_time)
VALUES
  (1, 1, 1, 'Roma Termini', 1, '2026-02-11 07:00:00', '2026-02-11 07:05:00'),
  (2, 1, 5, 'Bologna Centrale', 2, '2026-02-11 09:00:00', '2026-02-11 09:05:00'),
  (3, 1, 2, 'Milano Centrale', 3, '2026-02-11 11:30:00', '2026-02-11 11:35:00')
ON DUPLICATE KEY UPDATE
  schedule_id = VALUES(schedule_id), station_id = VALUES(station_id), station_name = VALUES(station_name), stop_order = VALUES(stop_order), arrival_time = VALUES(arrival_time), departure_time = VALUES(departure_time);

-- Schedule 2 (Milano -> Bologna -> Napoli)
INSERT INTO schedules_stops (id, schedule_id, station_id, station_name, stop_order, arrival_time, departure_time)
VALUES
  (4, 2, 2, 'Milano Centrale', 1, '2026-02-11 08:00:00', '2026-02-11 08:05:00'),
  (5, 2, 5, 'Bologna Centrale', 2, '2026-02-11 10:30:00', '2026-02-11 10:35:00'),
  (6, 2, 3, 'Napoli Centrale', 3, '2026-02-11 14:00:00', '2026-02-11 14:05:00')
ON DUPLICATE KEY UPDATE
  schedule_id = VALUES(schedule_id), station_id = VALUES(station_id), station_name = VALUES(station_name), stop_order = VALUES(stop_order), arrival_time = VALUES(arrival_time), departure_time = VALUES(departure_time);

-- Schedule 3 (Bologna -> Firenze)
INSERT INTO schedules_stops (id, schedule_id, station_id, station_name, stop_order, arrival_time, departure_time)
VALUES
  (7, 3, 5, 'Bologna Centrale', 1, '2026-02-11 09:00:00', '2026-02-11 09:05:00'),
  (8, 3, 4, 'Firenze Santa Maria Novella', 2, '2026-02-11 10:15:00', '2026-02-11 10:20:00')
ON DUPLICATE KEY UPDATE
  schedule_id = VALUES(schedule_id), station_id = VALUES(station_id), station_name = VALUES(station_name), stop_order = VALUES(stop_order), arrival_time = VALUES(arrival_time), departure_time = VALUES(departure_time);

COMMIT;

