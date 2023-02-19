SELECT p.* FROM products p 
JOIN order_details od ON p.product_id = od.product_id
JOIN orders o ON od.order_id = o.order_id
JOIN customers c ON o.customer_id = c.customer_id
WHERE c.company_name = 'Contonso, Ltd';
