SELECT COUNT (Products.Id)
FROM Products,Categoreis
WHERE Products.Category = Categories.Id AND Products.Category = ?