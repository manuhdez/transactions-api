CREATE TABLE `accounts` (
    `id` varchar (100) NOT NULL,
    `balance` float DEFAULT 0,
    PRIMARY KEY (`id`)
);

CREATE TABLE `transactions` (
    `id` INT auto_increment NOT NULL,
    `account_id` varchar (100) NOT NULL,
    `amount` float NOT NULL,
    `currency` varchar(3) NOT NULL,
    `type` enum ('deposit', 'withdrawal') NOT NULL,
    `date` timestamp NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (`id`)
);
