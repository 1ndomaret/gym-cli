use Gym_Manager;

INSERT INTO Users (Email, Password, Type) VALUES
('budi.santoso@example.com', 'password123', 'member'),
('siti.aminah@example.com', 'password123', 'member'),
('agus.wijaya@example.com', 'password123', 'member'),
('dewi.anggraini@example.com', 'password123', 'member'),
('eko.prasetyo@example.com', 'password123', 'member'),
('rina.wati@example.com', 'password123', 'member'),
('hendra.gunawan@example.com', 'password123', 'member'),
('ayu.wulandari@example.com', 'password123', 'member'),
('bambang.kurniawan@example.com', 'password123', 'member'),
('admin@example.com', 'admin123', 'admin');

INSERT INTO PaymentMethod (MethodName) VALUES
('Cash'),
('Bank Transfer'),
('Credit Card'),
('GoPay'),
('OVO');

INSERT INTO Tiers (TierName, TierRank, MonthlyCost) VALUES
('Bronze', 1, 150000.00),
('Silver', 2, 300000.00),
('Gold', 3, 500000.00);

INSERT INTO Events (EventName, MinTierRank, Schedule) VALUES
('Open Gym Day', 1, '2026-08-01 08:00:00'),
('Group Workout Session', 1, '2026-08-03 18:00:00'),
('Aqua Aerobics Class', 2, '2026-08-05 07:00:00'),
('Spin Class', 2, '2026-08-07 19:00:00'),
('Personal Training Session', 3, '2026-08-10 10:00:00'),
('Wellness and Spa Day', 3, '2026-08-12 14:00:00');

INSERT INTO UserProfiles (UserId, MemberTierId, FirstName, LastName, Address, CreatedAt, Status) VALUES
(1, 1, 'Budi', 'Santoso', 'Jl. Merdeka No. 10, Jakarta', '2026-01-10 09:30:00', TRUE),
(2, 1, 'Siti', 'Aminah', 'Jl. Diponegoro No. 22, Bandung', '2026-01-15 14:20:00', TRUE),
(3, 1, 'Agus', 'Wijaya', 'Jl. Gatot Subroto No. 5, Surabaya', '2026-02-01 11:00:00', TRUE),
(4, 2, 'Dewi', 'Anggraini', 'Jl. Ahmad Yani No. 8, Semarang', '2026-02-12 16:45:00', TRUE),
(5, 2, 'Eko', 'Prasetyo', 'Jl. Pahlawan No. 17, Yogyakarta', '2026-03-03 08:15:00', TRUE),
(6, 2, 'Rina', 'Wati', 'Jl. Sudirman No. 30, Jakarta', '2026-03-20 13:30:00', TRUE),
(7, 3, 'Hendra', 'Gunawan', 'Jl. Malioboro No. 12, Yogyakarta', '2026-04-05 10:00:00', TRUE),
(8, 3, 'Ayu', 'Wulandari', 'Jl. Thamrin No. 25, Jakarta', '2026-04-18 15:10:00', TRUE),
(9, 3, 'Bambang', 'Kurniawan', 'Jl. Asia Afrika No. 40, Bandung', '2026-05-02 09:00:00', FALSE);

INSERT INTO Invoices (UserProfileId, MemberTierId, Amount, DueDate, InvoiceStatus) VALUES
(1, 1, 150000.00, '2026-07-01', 'paid'),
(2, 1, 150000.00, '2026-07-01', 'pending'),
(3, 1, 150000.00, '2026-06-01', 'overdue'),
(4, 2, 300000.00, '2026-07-01', 'paid'),
(5, 2, 300000.00, '2026-07-01', 'paid'),
(6, 2, 300000.00, '2026-07-05', 'pending'),
(7, 3, 500000.00, '2026-07-01', 'paid'),
(8, 3, 500000.00, '2026-06-01', 'overdue'),
(9, 3, 500000.00, '2026-07-01', 'cancelled');

INSERT INTO Payments (InvoiceId, PaymentDate, PaymentMethodId, PaymentStatus) VALUES
(1, '2026-06-28', 4, 'completed'),
(4, '2026-06-29', 2, 'completed'),
(5, '2026-06-30', 3, 'completed'),
(7, '2026-06-27', 1, 'completed'),
(3, '2026-06-15', 5, 'failed'),
(8, '2026-06-10', 4, 'pending');
