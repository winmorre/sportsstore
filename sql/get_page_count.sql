SELECT COUNT (Products.Id)
FROM Proucts,Categories
WHERE Products.Category = Categories.Id;