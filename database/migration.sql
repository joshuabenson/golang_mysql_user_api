CREATE TABLE `users`
(
    id bigint AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `users` (`name`)
VALUES ('Faj Mccain'),
       ('Freddy KRUGER'),
       ('Frank Anthony');