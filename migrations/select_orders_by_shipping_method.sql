SELECT o.* FROM orders o JOIN shipping_methods sm ON o.shipping_method_id = sm.shipping_method_id
WHERE sm.shipping_method = 'UPS Ground';