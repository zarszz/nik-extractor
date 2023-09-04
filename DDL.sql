CREATE TABLE provinces
(
    id   VARCHAR(2)   NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE regencies
(
    id          VARCHAR(4)   NOT NULL,
    province_id VARCHAR(255) NOT NULL,
    name        VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (province_id) REFERENCES provinces (id)
);

CREATE TABLE districts
(
    id         VARCHAR(7)   NOT NULL,
    regency_id VARCHAR(4)   NOT NULL,
    name       VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (regency_id) REFERENCES regencies (id)
);

CREATE TABLE villages
(
    id          VARCHAR(10)  NOT NULL,
    district_id VARCHAR(7)   NOT NULL,
    name        VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (district_id) REFERENCES districts (id)
);

CREATE TABLE `users`
(
    `id`   varchar(16)  NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
)