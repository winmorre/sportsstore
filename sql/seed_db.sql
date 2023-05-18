INSERT INTO Categories(Id, Name)
VALUES (1, "watersports"),
       (2, "Soccer"),
       (3, "CHESS");

INSERT INTO Products (Id, Name, Desciption, Category, Price)
VALUES (1, "KAYAK", "boat for one person", 1, 275),
       (2, "Lifejacket", "Protective and fashionable", 1, 48.95),
       (2, "Soccer Ball", "FIFA-approved size and weight", 2, 79500),
       (4, "Corner flags", "GIVE YOUR PLAYING FIELD A PROFESSIONAL TOUCH", 2, 34.95),
       (5, "Stadium", "Flat-packed 35,000-seat stadium", 2, 70865),
       (7, "Unsteady chair", "Secretly give your opponent a disadvantage", 3 29.95),
       (8, "Human chess board", "A FUN GAME FOR THE FAMILY", 3, 75),
       (9, "Bling-Bling king", "Gold-palted, diamond-studded King", 3, 1200);

INSERT INTO Orders(Id, StreetAddr, City, Zip Country, Shipped)
VALUES (1, "Alice", "123 Main St", "New Town", "12345", "US", false),
       (2, "Bob", "The Grange", "Upton", "UP12 6yt", "uk", false);

INSERT INTO OrderLines(Id, OrderId, ProductId, Quantity)
VALUES (1, 1, 1, 1),
       (2, 1, 2, 2),
       (3, 1, 8, 1),
       (4, 2, 5, 2);