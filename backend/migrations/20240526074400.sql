-- Create "addresses" table
CREATE TABLE `addresses` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` bigint unsigned NULL,
  `street` varchar(100) NULL,
  `city` varchar(50) NULL,
  `state` varchar(50) NULL,
  `zip` varchar(10) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_addresses_user` (`user_id`),
  INDEX `idx_addresses_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "cart_updates" table
CREATE TABLE `cart_updates` (
  `quantity` bigint NULL,
  `product_id` bigint unsigned NULL,
  INDEX `fk_cart_updates_product` (`product_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "carts" table
CREATE TABLE `carts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `quantity` bigint NULL,
  `product_id` bigint unsigned NULL,
  `user_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_carts_product` (`product_id`),
  INDEX `fk_users_carts` (`user_id`),
  INDEX `idx_carts_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "order_items" table
CREATE TABLE `order_items` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint unsigned NULL,
  `product_id` bigint unsigned NULL,
  `quantity` bigint NULL,
  `price` decimal(10,2) NULL,
  `currency` char(3) NULL DEFAULT "myr",
  PRIMARY KEY (`id`),
  INDEX `fk_order_items_product` (`product_id`),
  INDEX `fk_orders_order_items` (`order_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "orders" table
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` bigint unsigned NULL,
  `address_id` bigint unsigned NULL,
  `payment_at` datetime(3) NULL,
  `order_status` enum('to_pay','to_ship','to_receive','to_review','complete') NULL DEFAULT "to_pay",
  PRIMARY KEY (`id`),
  INDEX `fk_orders_address` (`address_id`),
  INDEX `fk_orders_user` (`user_id`),
  INDEX `idx_orders_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "payments" table
CREATE TABLE `payments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `order_id` bigint unsigned NULL,
  `stripe_session_id` char(66) NULL,
  `is_complete` bool NULL DEFAULT 0,
  `payment_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_payments_order` (`order_id`),
  INDEX `idx_payments_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "products" table
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(255) NULL,
  `description` varchar(255) NULL,
  `unit_price` decimal(10,2) NULL,
  `currency` char(3) NULL DEFAULT "myr",
  `stock_quantity` bigint NULL DEFAULT 0,
  `is_active` bool NULL DEFAULT 1,
  `image_path` varchar(500) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_products_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` varchar(100) NULL,
  `email` varchar(100) NULL,
  `password` varchar(255) NULL,
  `role` enum('admin','member') NULL DEFAULT "member",
  `profile_pic` varchar(255) NULL,
  `sub` varchar(100) NULL,
  `default_address_id` bigint unsigned NULL,
  `contact_number` varchar(20) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_users_default_address` (`default_address_id`),
  INDEX `idx_users_deleted_at` (`deleted_at`),
  UNIQUE INDEX `uni_users_sub` (`sub`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Modify "addresses" table
ALTER TABLE `addresses` ADD CONSTRAINT `fk_addresses_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "cart_updates" table
ALTER TABLE `cart_updates` ADD CONSTRAINT `fk_cart_updates_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "carts" table
ALTER TABLE `carts` ADD CONSTRAINT `fk_carts_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT `fk_users_carts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "order_items" table
ALTER TABLE `order_items` ADD CONSTRAINT `fk_order_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT `fk_orders_order_items` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "orders" table
ALTER TABLE `orders` ADD CONSTRAINT `fk_orders_address` FOREIGN KEY (`address_id`) REFERENCES `addresses` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT `fk_orders_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "payments" table
ALTER TABLE `payments` ADD CONSTRAINT `fk_payments_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "users" table
ALTER TABLE `users` ADD CONSTRAINT `fk_users_default_address` FOREIGN KEY (`default_address_id`) REFERENCES `addresses` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
