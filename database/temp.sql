INSERT INTO Employees 
(first_name, last_name, email, title, is_active) 
VALUES 
('Silvester', 'Quartly', 'squartly0@ebay.com', 'Civil Engineer', true),
('Mala', 'Paddingdon', 'mpaddingdon1@scientificamerican.com', 'Recruiter', true),
('Romy', 'Mountjoy', 'rmountjoy2@studiopress.com', 'Chemical Engineer', true),
('Gardie', 'Varian', 'gvarian3@weibo.com', 'Senior Financial Analyst', true),
('Lynsey', 'Rylstone', 'lrylstone4@opensource.org', 'Senior Sales Associate', true),
('Tandy', 'Oury', 'toury5@google.de', 'Database Administrator II', true),
('Eugenio', 'Dunbobbin', 'edunbobbin6@tuttocitta.it', 'Programmer I', true),
('Marcus', 'Eckhard', 'meckhard7@sphinn.com', 'Programmer Analyst IV', true),
('Wilhelm', 'Acton', 'wacton8@redcross.org', 'Structural Engineer', true),
('Mireille', 'Einchcombe', 'meinchcombe9@washingtonpost.com', 'Editor', true);

--testing manipulating email
SELECT 
	email, 
	lower(first_name) AS lower_first, 
	lower(last_name) AS lower_last,
	(lower(first_name) || '.' || lower(last_name) || '@business.com') AS new_email
FROM Employees;

--testing update of one employee email
UPDATE Employees
SET email=(
	SELECT lower(first_name) || '.' || lower(last_name) || '@business.com' 
	FROM Employees
	)
WHERE employee_id=1;

--testing update of multiple employee email (is updating both employees to the first records email)
UPDATE Employees
SET email=(
	SELECT lower(first_name) || '.' || lower(last_name) || '@business.com' 
	FROM Employees
	)
WHERE employee_id=2 OR employee_id=3;

--testing update of multiple employee email (works!)
UPDATE Employees
SET email=(
	SELECT lower(first_name) || '.' || lower(last_name) || '@business.com' 
	FROM Employees e2
	WHERE Employees.employee_id=e2.employee_id
	)
WHERE employee_id=2 OR employee_id=3;

--select emails that haven't been modified to the correct format
SELECT first_name, last_name, email
FROM Employees
WHERE email NOT like '%@business.com%';

--update all emails that haven't been updated to the correct format
UPDATE Employees
SET email=(
	SELECT lower(first_name) || '.' || lower(last_name) || '@business.com' 
	FROM Employees e2
	WHERE Employees.employee_id=e2.employee_id
	)
WHERE email NOT like '%@business.com%';


INSERT INTO Departments 
(name, is_active) 
VALUES 
('Human Resources', true),
('Marketing', true),
('Engineering', true),
('Finance', true),
('Sales', true),
('Executive', true),
('Customer Relations', true);

/* Realized this model is wrong, because in the video interview they said that each employee could be in 
multiple departments. 

--see if I can get department_id and name based on employee title inside the select clause
SELECT 
	first_name, 
	last_name, 
	title, 
	department_id,
	CASE
		WHEN title like '%Engineer%' or title like '%Programmer%' or title like '%Database%'
			THEN (SELECT name FROM Department WHERE name='Engineering')
		WHEN title like '%Recruiter%' THEN (SELECT name FROM Department WHERE name='HR')
		ELSE (SELECT name FROM Department WHERE name='Marketing')
	END as department_name
FROM Employees;

--see if I can get department_id and name based on employee title with a left join 
SELECT e.first_name, e.last_name, e.title, d.department_id, d.name
FROM Employees e
LEFT JOIN Department d on 
	CASE
		WHEN title like '%Engineer%' or title like '%Programmer%' or title like '%Database%'
			THEN d.name='Engineering'
		WHEN title like '%Recruiter%' THEN d.name='HR'
		ELSE d.name='Marketing'
	END
WHERE e.employee_id = 1 or e.employee_id=2;

--check how to select department
SELECT department_id, name
FROM Department
WHERE name='Engineering';

--update employees to have a department_id based on their title
UPDATE Employees
SET department_id=(
	SELECT d.department_id
	FROM Department d
	WHERE
		CASE
			WHEN Employees.title like '%Engineer%' or Employees.title like '%Programmer%' or Employees.title like '%Database%'
				THEN d.name='Engineering'
			WHEN Employees.title like '%Recruiter%'
				THEN d.name='HR'
			ELSE d.name='Marketing'
		END
	);
	
--validate setting the employee department worked properly
SELECT e.first_name, e.last_name, e.title, d.name
FROM Employees e
LEFT JOIN Department d ON e.department_id=d.department_id;
*/
