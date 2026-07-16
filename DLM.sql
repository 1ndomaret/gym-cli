Create database if not exists Gym_Manager;
use Gym_Manager;

create table Users(
    UserId Int Primary key AUTO_INCREMENT,
    Email varchar(50) not null unique,
    Password varchar(50) not null,
    Type ENUM('member','admin') NOT NULL DEFAULT 'member'
);

create table PaymentMethod(
    PaymentMethodId int primary key Auto_Increment,
    MethodName varchar(50) not null
);

create table Tiers(
    TierId int Primary key AUTO_INCREMENT,
    TierName varchar(50) not null,
    TierRank INT NOT NULL UNIQUE,
    MonthlyCost Decimal(10,2)
);

create table Events(
    EventId int Primary key AUTO_INCREMENT,
    EventName varchar(50) not null,
    MinTierRank int not null,
    Schedule DATETIME
);

create table UserProfiles(
    UserProfileId int Primary key Auto_Increment,
    UserId int not null unique,
    MemberTierId int not null,
    FirstName varchar(50) not null,
    LastName varchar(50),
    Address varchar(100),
    CreatedAt DATETIME not null DEFAULT current_timestamp,
    Status boolean not null DEFAULT true,
    FOREIGN KEY (UserId) references Users(UserId) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (MemberTierId) references Tiers(TierId) ON DELETE RESTRICT ON UPDATE CASCADE
);

create table Invoices(
    InvoiceId int primary key Auto_Increment,
    UserProfileId Int not null,
    MemberTierId Int not null,
    Amount decimal(10,2),
    DueDate date,
    InvoiceStatus ENUM('pending','paid','overdue','cancelled') NOT NULL DEFAULT 'pending',
    FOREIGN KEY (UserProfileId) references UserProfiles(UserProfileId) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (MemberTierId) References Tiers(TierId) ON DELETE RESTRICT ON UPDATE CASCADE
);

create table Payments(
    PaymentId int primary key Auto_Increment,
    InvoiceId int not null,
    PaymentDate date,
    PaymentMethodId int not null,
    PaymentStatus ENUM('pending','completed','failed','refunded') NOT NULL DEFAULT 'pending',
    FOREIGN KEY (InvoiceId) References Invoices(InvoiceId) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (PaymentMethodId) References PaymentMethod(PaymentMethodId) ON DELETE RESTRICT ON UPDATE CASCADE
);
