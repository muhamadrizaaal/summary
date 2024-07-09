-- Membuat Skema Database dan Tabel
Misalkan kita memiliki tiga tabel: sales, customers, dan products.

CREATE TABLE sales (
    id SERIAL PRIMARY KEY,
    product_id INT,
    customer_id INT,
    amount DECIMAL,
    created_at TIMESTAMP
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP
);

CREATE TABLE summary (
    id SERIAL PRIMARY KEY,
    total_sales DECIMAL,
    total_customers INT,
    total_products INT,
    summary_date TIMESTAMP
);

-- Menulis Query SQL untuk Summary
Query ini akan menghitung total penjualan, jumlah pelanggan, dan jumlah produk.

INSERT INTO summary (total_sales, total_customers, total_products, summary_date)
SELECT 
    (SELECT SUM(amount) FROM sales) AS total_sales,
    (SELECT COUNT(DISTINCT id) FROM customers) AS total_customers,
    (SELECT COUNT(DISTINCT id) FROM products) AS total_products,
    NOW() AS summary_date;
