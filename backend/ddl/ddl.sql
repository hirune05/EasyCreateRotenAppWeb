CREATE TABLE `log` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` int NOT NULL,
    `action` varchar(255),
    `details` text,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- store staff table
CREATE TABLE `store_staff` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `role` int NOT NULL,  -- role is either store staff or store owner
    `store_id` int,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
    -- FOREIGN KEY (`store_id`) REFERENCES `store`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- store table
CREATE TABLE `store` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(255) NOT NULL,
    `grade` int,
    `major` varchar(255),
    `year` int,
    `image_url` text,
    `store_owner_id` int NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`store_owner_id`) REFERENCES `store_staff`(`id`) ON DELETE CASCADE
);

-- admin_user table
CREATE TABLE `admin_user` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `email` varchar(255),
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- product table
CREATE TABLE `product` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(255) NOT NULL,
    `description` text,
    `price` int NOT NULL,
    `image_url` text,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- order table
CREATE TABLE `order` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `store_id` int NOT NULL,
    `prepared_at` datetime,
    `completed_at` datetime,
    `pickedup_at` datetime,
    `status` int,  -- 0: Preparing, 1: Ready for Pickup, 2: Completed
    `user_id` int,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`store_id`) REFERENCES `store`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- order_item table
CREATE TABLE `order_item` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `order_id` int NOT NULL,
    `product_id` int NOT NULL,
    `arranges` text,
    `quantity` int NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`order_id`) REFERENCES `order`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`product_id`) REFERENCES `product`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- report table
CREATE TABLE `report` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `store_id` int NOT NULL,
    `store_staff_id` int NOT NULL,
    `description` text,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`store_id`) REFERENCES `store`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`store_staff_id`) REFERENCES `store_staff`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
