create table if not exists country (
    id serial primary key,
    gdp_usd decimal not null,
    name varchar(50) not null
);

create table if not exists subsidy (
    id serial primary key,
    country_id bigint not null references country(id),
    require_price decimal not null,
    required_wd varchar(3) not null
);

create table if not exists country_tax (
    id serial primary key,
    country_id bigint not null references country(id),
    tax decimal not null
);

create table if not exists vendor (
    id serial primary key,
    country_id serial references country(id),
    capitalization decimal,
    name varchar(50) not null
);


create table if not exists "type" (
    id serial primary key,
    "type" varchar(20) not null,
    additional_info varchar(30) not null
);

create table if not exists component (
    id serial primary key,
    vendor_id bigint not null references vendor(id),
    type_id bigint not null references "type"(id),
    name varchar(30) not null,
    additional_info varchar(30) not null
);

create table if not exists factory (
    id serial primary key,
    vendor_id bigint not null references vendor(id),
    max_workers bigint not null check ( max_workers >= 0 ),
    productivity decimal not null check ( productivity >= 0 )
);

create table if not exists engineer (
    id serial primary key,
    vendor_id bigint not null references vendor(id),
    name varchar(100) not null,
    gender varchar(6) not null,
    experience decimal not null check ( experience >= 0 ),
    salary decimal check ( salary >= 0 ),
    factory_id bigint not null references factory(id)
);

create table if not exists model (
    id serial primary key,
    vendor_id serial references vendor(id),
    name varchar(20) not null,
    wheeldrive varchar(3) not null,
    significance int not null,
    price decimal not null,
    prod_cost decimal not null,
    engineer_id bigint not null references engineer(id),
    factory_id bigint not null references factory(id),
    sales bigint not null check ( sales >= 0 )
);

create table if not exists model_component (
    id serial primary key,
    model_id bigint not null references model(id),
    component_id bigint not null references component(id)
);

create table if not exists storage (
    id serial primary key,
    model_id bigint references model(id),
    amount bigint  check ( amount >= 0 ),
    vendor_id bigint not null references vendor(id)
);

create table if not exists "order" (
    id serial primary key,
    model_id bigint not null references model(id),
    quantity bigint not null,
    order_type varchar(50) not null
);

create table if not exists shipment (
    id serial primary key,
    order_id serial references "order"(id),
    country_to_id serial references country(id),
    "date" timestamp not null,
    cost decimal not null check ( cost >= 0 )
);