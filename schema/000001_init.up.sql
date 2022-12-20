CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    created_at    date
);

CREATE TABLE wallets
(
    id          serial                                      not null unique,
    user_id     int references users (id) on delete cascade not null,
    created_at  date,
    updated_at  date,
    private_key varchar,
    public_key  varchar,
    balance     numeric(18, 2)
);


CREATE TABLE transactions
(
    id          serial                                        not null unique,
    created_at  date,
    sender_id   int references users (id) on delete cascade   not null,
    receiver_id int references users (id) on delete cascade   not null,
    wallet_id   int references wallets (id) on delete cascade not null,
    amount      numeric(18, 2)
);