SELECT c.* FROM customers c 
JOIN orders o ON c.customer_id = o.customer_id 
JOIN employees e ON o.employee_id = e.employee_id 
WHERE e.first_name = 'Adam' AND e.last_name = 'Barr';