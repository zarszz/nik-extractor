create TABLE provinces (
    id VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

create TABLE cities (
    id VARCHAR(255) NOT NULL,
    province_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (province_id) REFERENCES provinces(id)
);

create TABLE districts (
    id VARCHAR(255) NOT NULL,
    city_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (city_id) REFERENCES cities(id)
);

create TABLE villages (
    id VARCHAR(255) NOT NULL,
    district_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (district_id) REFERENCES districts(id)
);