-- LOGIN (admin & user)
SELECT UserId, Email, Type
FROM Users
WHERE Email = ? AND Password = ?;


-- Admin
-- Create Membership (transaction: account + profile)
INSERT INTO Users (Email, Password, Type) VALUES (?, ?, 'member');
INSERT INTO UserProfiles (UserId, MemberTierId, FirstName, LastName, Address)
VALUES (LAST_INSERT_ID(), ?, ?, ?, ?);

-- View member List
SELECT up.UserProfileId, u.Email, up.FirstName, up.LastName,
       t.TierName, up.Address, up.CreatedAt, up.Status
FROM UserProfiles up
JOIN Users u ON u.UserId = up.UserId
JOIN Tiers t ON t.TierId = up.MemberTierId
ORDER BY up.UserProfileId; 
-- for ACTIVE/INACTIVE member
-- WHERE up.status = TRUE 


--Update Member (select)
SELECT up.UserProfileId, u.UserId, u.Email, up.FirstName, up.LastName,
       up.Address, up.MemberTierId, t.TierName, up.Status, up.CreatedAt
FROM UserProfiles up
JOIN Users u ON u.UserId = up.UserId
JOIN Tiers t ON t.TierId = up.MemberTierId
WHERE up.UserProfileId = ?; -- insert member userprofileID

-- Update Member (apply changes)
UPDATE UserProfiles
SET FirstName = ?, LastName = ?, Address = ?, MemberTierId = ?, Status = ?
WHERE UserProfileId = ?;

-- Delete member (soft delete) -> make them inactive
UPDATE UserProfiles SET Status = FALSE WHERE UserProfileId = ?;

--Create Payment 
INSERT INTO Payments (InvoiceId, PaymentDate, PaymentMethodId, PaymentStatus)
VALUES (?, CURDATE(), ?, 'completed');
UPDATE Invoices SET InvoiceStatus = 'paid' WHERE InvoiceId = ?;


--Admin Reports
--Membership Report(Monthly member joins):
SELECT DATE_FORMAT(CreatedAt, '%Y-%m') AS JoinMonth, COUNT(*) AS NewMembers
FROM UserProfiles
GROUP BY DATE_FORMAT(CreatedAt, '%Y-%m')
ORDER BY JoinMonth;


--Membership Report(member churn):
SELECT up.UserProfileId, up.FirstName, up.LastName, t.TierName
FROM UserProfiles up
JOIN Tiers t ON t.TierId = up.MemberTierId
WHERE up.Status = FALSE;


--Top Highest membership report(mvp member)
SELECT up.UserProfileId, up.FirstName, up.LastName, SUM(i.Amount) AS TotalPaid
FROM Invoices i
JOIN UserProfiles up ON up.UserProfileId = i.UserProfileId
WHERE i.InvoiceStatus = 'paid'
GROUP BY up.UserProfileId, up.FirstName, up.LastName
ORDER BY TotalPaid DESC
LIMIT 10;


--Monthly Income Report
SELECT DATE_FORMAT(p.PaymentDate, '%Y-%m') AS Month, SUM(i.Amount) AS TotalIncome
FROM Payments p
JOIN Invoices i ON i.InvoiceId = p.InvoiceId
WHERE p.PaymentStatus = 'completed'
GROUP BY DATE_FORMAT(p.PaymentDate, '%Y-%m')
ORDER BY Month;



-- USER 
--View schedule
SELECT e.EventName, e.Schedule
FROM UserProfiles up
JOIN Tiers t ON t.TierId = up.MemberTierId
JOIN Events e ON e.MinTierRank <= t.TierRank
WHERE up.UserId = ?
  AND e.Schedule >= NOW()
ORDER BY e.Schedule;


--view pending payment
SELECT i.InvoiceId, i.Amount, i.DueDate, i.InvoiceStatus
FROM Invoices i
JOIN UserProfiles up ON up.UserProfileId = i.UserProfileId
WHERE up.UserId = ?
  AND i.InvoiceStatus IN ('pending', 'overdue')
ORDER BY i.DueDate;