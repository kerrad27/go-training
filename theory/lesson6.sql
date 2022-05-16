-- create table customers
-- (
--     customer_id   integer default nextval('"Customers_CustomerID_seq"'::regclass) not null
--         constraint customers_pk
--             primary key,
--     customer_name varchar,
--     contact_name  varchar,
--     address       varchar,
--     city          varchar,
--     postal_code   varchar,
--     country       varchar
-- );
--
-- alter table customers
--     owner to user_db;
--
-- create unique index customers_customerid_uindex
--     on customers (customer_id);




-- select * from CUSTOMERS

-- insert into customers values (1, 'Alfreds Futterkiste', 'Maria Anders',
--                               'Obere Str. 57', 'Berlin', 12209, 'Germany')

-- insert into customers values (2, 'Ana Trujillo Emparedados y helados', 'Ana Trujillo',
--                               'Avda. de la Constitución 2222', 'México D.F.', 05021, 'Mexico')

-- insert into customers values (3, 'Antonio Moreno Taquería', 'Antonio Moreno',
--                               'Mataderos 2312', 'México D.F.', 05023, 'Mexico')

-- insert into customers values (4, 'Around the Horn', 'Thomas Hardy',
--                               '120 Hanover Sq.', 'London', 'WA1 1DP', 'UK')

-- insert into customers values (6, 'Blauer See Delikatessen', 'Hanna Moos',
--                               'Forsterstr. 57', 'Mannheim', '68306', 'Germany')



-- SELECT - extracts data from a database
-- UPDATE - updates data in a database
-- DELETE - deletes data from a database
-- INSERT INTO - inserts new data into a database
-- CREATE DATABASE - creates a new database
-- ALTER DATABASE - modifies a database
-- CREATE TABLE - creates a new table
-- ALTER TABLE - modifies a table
-- DROP TABLE - deletes a table
-- CREATE INDEX - creates an index (search key)
-- DROP INDEX - deletes an index

-- SELECT column1, column2, ...
--     FROM table_name;

-- SELECT DISTINCT column1, column2, ...
--     FROM table_name;




-- SELECT Country FROM Customers;
--
-- SELECT COUNT(DISTINCT Country) FROM Customers;

-- SELECT Count(*) AS DistinctCountries
-- FROM (SELECT DISTINCT Country FROM Customers);


-- SELECT column1, column2, ...
--     FROM table_name
-- WHERE condition;

-- SELECT * FROM Customers WHERE Country='Mexico'

-- SELECT * FROM Customers WHERE Customer_ID=1

-- =	Equal
-- >	Greater than
-- <	Less than
-- >=	Greater than or equal
-- <=	Less than or equal
-- <>	Not equal. Note: In some versions of SQL this operator may be written as !=
-- BETWEEN	Between a certain range
-- LIKE	Search for a pattern
-- IN	To specify multiple possible values for a column


-- AND Syntax
--
-- SELECT column1, column2, ...
--     FROM table_name
-- WHERE condition1 AND condition2 AND condition3 ...;
--
-- OR Syntax
--
-- SELECT column1, column2, ...
--     FROM table_name
-- WHERE condition1 OR condition2 OR condition3 ...;
--
-- NOT Syntax
--
-- SELECT column1, column2, ...
--     FROM table_name
-- WHERE NOT condition;

-- SELECT * FROM Customers
-- WHERE Country='Germany' AND City='Berlin'

-- SELECT * FROM Customers
-- WHERE Country='Germany' OR Country='UK'

-- SELECT * FROM Customers
-- WHERE NOT Country='Germany';

-- SELECT * FROM Customers
-- WHERE Country='Germany' AND (City='Berlin' OR City='München');

-- SELECT * FROM Customers
-- WHERE NOT Country='Germany' AND NOT Country='UK';

-- SELECT * FROM Customers
-- ORDER BY Country;

-- SELECT * FROM Customers
-- ORDER BY Country DESC;

-- SELECT * FROM Customers
-- ORDER BY Country, Customer_Name;

-- SELECT * FROM Customers
-- ORDER BY Country ASC, Customer_Name DESC;

-- INSERT INTO Customers (Customer_Name, Contact_Name, Address, City, Postal_Code, Country)
-- VALUES ('Cardinal', 'Tom B. Erichsen', 'Skagen 21', 'Stavanger', '4006', 'Norway');

-- INSERT INTO Customers (Customer_Name, City, Country)
-- VALUES ('Cardinal', 'Stavanger', 'Norway');

-- SELECT Customer_Name, Contact_Name, Address
-- FROM Customers
-- WHERE Address IS NULL;

-- SELECT Customer_Name, Contact_Name, Address
-- FROM Customers
-- WHERE Address IS NOT NULL;

-- UPDATE Customers
-- SET Contact_Name = 'Alfred Schmidt', City= 'Frankfurt'
-- WHERE Customer_ID = 1;

-- multiple

-- UPDATE Customers
-- SET Contact_Name='Juan'
-- WHERE Country='Mexico';

-- DELETE FROM Customers WHERE Customer_Name='Alfreds Futterkiste';

-- SELECT * FROM Customers
-- LIMIT 3;

-- SELECT * FROM Customers
-- WHERE Country='Germany'
-- LIMIT 3;

-- SELECT MIN(customer_id) AS SmallestCustomerId
-- FROM customers;

-- LIKE

-- WHERE CustomerName LIKE 'a%'	Finds any values that start with "a"
-- WHERE CustomerName LIKE '%a'	Finds any values that end with "a"
-- WHERE CustomerName LIKE '%or%'	Finds any values that have "or" in any position
-- WHERE CustomerName LIKE '_r%'	Finds any values that have "r" in the second position
-- WHERE CustomerName LIKE 'a_%'	Finds any values that start with "a" and are at least 2 characters in length
-- WHERE CustomerName LIKE 'a__%'	Finds any values that start with "a" and are at least 3 characters in length
-- WHERE ContactName LIKE 'a%o'	Finds any values that start with "a" and ends with "o"

-- SELECT * FROM Customers
-- WHERE Customer_Name LIKE 'Al%';

-- SELECT * FROM Customers
-- WHERE Country IN ('Germany', 'France', 'UK');

-- SELECT * FROM Customers
-- WHERE Country NOT IN ('Germany', 'France', 'UK');

-- SELECT * FROM Customers
-- WHERE customer_id BETWEEN 3 AND 6;

-- INSERT INTO orders VALUES (10248,	90,	5,	'1996-07-04')
-- INSERT INTO orders VALUES (10249,	1,	6,	'1996-07-05');
-- INSERT INTO orders VALUES (10250,	2,	6,	'1996-07-06');
-- INSERT INTO orders VALUES (10251,	1,	7,	'1996-07-08');

-- SELECT Orders.Order_ID, Customers.Customer_Name, Orders.Order_Date
-- FROM Orders FULL JOIN Customers ON Orders.Customer_ID=Customers.Customer_ID;

-- SELECT A.Customer_Name AS CustomerName1, B.Customer_Name AS CustomerName2, A.City
-- FROM Customers A, Customers B
-- WHERE A.Customer_ID <> B.Customer_ID
--   AND A.City = B.City
-- ORDER BY A.City;

-- SELECT customer_id FROM Customers
-- UNION
-- SELECT Order_ID FROM Orders;

-- SELECT COUNT(Customer_ID), Country
-- FROM Customers
-- GROUP BY Country;

-- SELECT COUNT(Customer_ID), Country
-- FROM Customers
-- GROUP BY Country
-- ORDER BY COUNT(Customer_ID) DESC;

-- SELECT COUNT(Customer_ID), Country
-- FROM Customers
-- GROUP BY Country
-- HAVING COUNT(Customer_ID) > 1;