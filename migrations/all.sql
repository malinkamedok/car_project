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

/*
 Триггеры
*/

/* Триггер 1 - подсчёт prod_cost при добавлении */
create or replace function update_prod_cost()
    returns trigger as $$
begin
    update model set prod_cost = new.prod_cost * 1.3 where model.id = new.id;
    return new;
end;
$$ language 'plpgsql';

create trigger auto_update_prod_cost after insert on model
    for each row execute procedure update_prod_cost();



/* Триггер 2 - Тригер на увелечение sales при изменении количества в order */
create or replace function update_sales()
returns trigger as $$
begin
    update model set sales = model.sales + new.quantity where model.id = new.model_id;
    return new;
end;
$$ language 'plpgsql';

create trigger auto_update_sales after insert or update on "order"
    for each row execute procedure update_sales();



/* Триггер 3 - Тригер на проверку того, что количество машин в заказе меньше или равно количеству машин на складе */
create or replace function check_count_model()
returns trigger as $$
begin
    if ((select s.amount from storage as s inner join model m on m.id = s.model_id
        inner join "order" o on m.id = new.model_id where o.model_id = new.model_id) < new.quantity) then
        return null;
    end if;
    return new;
end;
$$ language 'plpgsql';

create trigger auto_check_count_model before insert on "order"
    for each row execute procedure check_count_model();



/*
 Функции
*/


/* функция 1 - Вывод списка всех моделей */
create or replace function get_all_models()
returns table(id integer, vendor_id integer, name varchar(20), wheeldrive varchar(3),
significance integer, price decimal, prod_cost decimal, engineer_id bigint, factory_id bigint, sales bigint, vendor_name varchar(50), engineer_name varchar(50), country_name varchar(50)) as $$
begin
    return query select m.id, m.vendor_id, m.name, m.wheeldrive, m.significance, m.price, m.prod_cost, m.engineer_id, m.factory_id, m.sales, v.name, e.name, c.name from model as m inner join engineer e on e.id = m.engineer_id inner join vendor as v on m.vendor_id = v.id inner join country c on v.country_id = c.id;
end;
$$ language 'plpgsql';

--select * from get_all_models();


/* функций 2 - Вывод всех супсидий */
create or replace function get_all_subsidies()
returns table(id integer, country_id bigint, require_price decimal, required_wd varchar(3), country_name varchar(30)) as $$
begin
    return query select s.id, s.country_id, s.require_price, s.required_wd, c.name from subsidy as s inner join country c on c.id = s.country_id;
end;
$$ language 'plpgsql';



/* функция 3 - Вывод инженеров по id вендора */
create or replace function get_engineer_by_vendor(vendor_id_by bigint)
returns table(id integer, vendor_id bigint, name varchar(100), gender varchar(6),
experience decimal, salary decimal, factory_id bigint) as $$
begin
    return query select * from engineer where engineer.vendor_id = vendor_id_by;
end;
$$ language 'plpgsql';



/* функция 4 - Вывод заводов по id вендора */
create or replace function get_factory_by_vendor(vendor_id_by bigint)
returns table(id integer, vendor_id bigint, max_workers bigint, productivity decimal) as $$
begin
    return query select * from factory where factory.vendor_id = vendor_id_by;
end;
$$ language 'plpgsql';



/* функция 5 - Вывод компонентов по id вендора и id типа*/
create or replace function get_components_by_vendor_and_type(vendor_id_by bigint, type_id_by bigint)
returns table(id integer, vendor_id bigint, type_id bigint, name varchar(30), additional_info varchar(30)) as $$
begin
    return query select * from component where component.vendor_id = vendor_id_by and component.type_id = type_id_by;
end;
$$ language 'plpgsql';



/* функция 6 - Вывод типов */
create or replace function get_all_types()
returns table(id integer, type varchar(20), additional_info varchar(30)) as $$
begin
    return query select * from type;
end;
$$ language 'plpgsql';



/* функция 7 - Размещение супсидии */
create or replace function create_subsidy(country_id_by bigint, require_price_by decimal, required_wd_by varchar(3))
returns void as $$
begin
    insert into subsidy (country_id, require_price, required_wd) values (country_id_by, require_price_by, required_wd_by);
end;
$$ language 'plpgsql';



/* функция 8 - создание заказа */
create or replace function create_order(model_id_by bigint, quantity_by bigint, order_type_by varchar(50),
  country_to_id_new bigint)
returns void as $$
declare
    new_order_id integer;
    quantity integer;
    price_d integer;
begin
    insert into "order" (model_id, quantity, order_type)  values (model_id_by, quantity_by, order_type_by);

    new_order_id = (select max(id) from "order");
    quantity = (select "order".quantity from "order" where "order".id = new_order_id);
    price_d = (select model.price from model where model.id = model_id_by);
    insert into shipment (order_id, country_to_id, date, cost) values (new_order_id, country_to_id_new, current_timestamp, quantity * price_d);
end
$$ language 'plpgsql';

--select create_order(1, 10, 'ddd', 1);
--select * from get_orders_by_vendor_id(1);

/* функция 9 - оформление доставки */
create or replace function create_shipment(order_id_by integer, country_to_id_by integer, date_by timestamp, cost_by decimal)
returns void as $$
begin
    insert into shipment (order_id, country_to_id, date, cost) values (order_id_by, country_to_id_by, date_by, cost_by);
end;
$$ language 'plpgsql';


/* функция 10 - прием субсидии - большая функция*/
create or replace function accept_subsidies(subsidy_id_new integer,
                                            vendor_id_new integer,
                                            name_new varchar(20),
                                            significance_new integer,
                                            engineer_id_new bigint,
                                            factory_id_new bigint,
                                            component_id_new_engine bigint,
                                            component_id_new_door bigint,
                                            component_id_new_bumper bigint,
                                            component_id_new_transmission bigint)
    returns void as $$
declare
    required_price_by_id decimal = (select s.require_price from subsidy s where s.id = subsidy_id_new);
    required_wd_by_id varchar(3) = (select s.required_wd from subsidy s where s.id = subsidy_id_new);
    new_model_id integer;
begin

    insert into model (vendor_id, name, wheeldrive, significance, price, prod_cost, engineer_id, factory_id, sales)
    values (vendor_id_new, name_new, required_wd_by_id, significance_new, required_price_by_id, required_price_by_id, engineer_id_new, factory_id_new, 0);

    new_model_id = (select max(id) from model);

    insert into model_component (model_id, component_id) values (new_model_id, component_id_new_engine);
    insert into model_component (model_id, component_id) values (new_model_id, component_id_new_door);
    insert into model_component (model_id, component_id) values (new_model_id, component_id_new_bumper);
    insert into model_component (model_id, component_id) values (new_model_id, component_id_new_transmission);

    delete from subsidy s where s.id = subsidy_id_new;

end;
$$ language 'plpgsql';

/* функция 11 - создаине новой модели */
create or replace function create_model(
    vendor_id_new integer,
    name_new varchar(20),
    wheeldrive_new varchar(3),
    significance_new integer,
    price_new bigint,
    engineer_id_new bigint,
    factory_id_new bigint
)
    returns void as $$
begin
    insert into model (vendor_id, name, wheeldrive, significance, price, prod_cost, engineer_id, factory_id, sales)
    values (vendor_id_new, name_new, wheeldrive_new, significance_new, price_new, price_new, engineer_id_new, factory_id_new, 0);
end;
$$ language 'plpgsql';

/* функция 12 - получение заказов по vendor_id */
create or replace function get_orders_by_vendor_id(vendor_id_new integer)
returns table(model_name varchar(50),model_id integer,country_name varchar(50), order_id integer, quantity bigint, order_type varchar(50), shipment_cost numeric, shipment_date timestamp) as $$
begin
    return query select model.name as model_name, model.id as model_id, c.name, "order".id as order_id, "order".quantity, "order".order_type, s.cost, s.date from "order"
                                                                                                                             inner join model on "order".model_id = model.id
                                                                                                                             inner join vendor on model.vendor_id = vendor.id
                                                                                                                             inner join shipment s on "order".id = s.order_id
                                                                                                                             inner join country c on s.country_to_id = c.id
    where vendor.id = vendor_id_new;
end;
$$ language 'plpgsql';

--insert into shipment (order_id, country_to_id,date, cost) values (7, 1, '2022-11-10 00:00:00.000000', 10);



/* функций 13 - получение заказов по country_id  сделать id модели!!!!!!! */
create or replace function get_orders_by_country_id(country_id_new integer)
returns table(vendor_name varchar(50), model_name varchar(50), model_id integer,order_id integer, quantity bigint, order_type varchar(50), shipment_cost numeric, shipment_date timestamp) as $$
begin
    return query select vendor.name as vendor, model.name as model, model.id as model_is, "order".id as order_id, "order".quantity, "order".order_type, shipment.cost as shipment_cost, shipment.date from "order"
                                                                                                                                                                                                  inner join model on "order".model_id = model.id
                                                                                                                                                                                                  inner join vendor on model.vendor_id = vendor.id
                                                                                                                                                                                                  inner join shipment on "order".id = shipment.order_id
                                                                                                                                                                                                  inner join country on shipment.country_to_id = country.id
    where country.id = country_id_new;
end;
$$ language 'plpgsql';

/* функция 14 - получение всех субсидий по country id */
create or replace function get_subsidies_by_country_id(country_id_new integer)
returns table(id_subsidy integer, country_id bigint, required_price numeric, required_wd varchar(3)) as $$
begin
    return query select * from subsidy where subsidy.country_id = country_id_new;
end;
$$ language 'plpgsql';

create or replace function get_all_components(type_new varchar(20))
returns table(id integer, vendor_id bigint, vendor_name varchar(50), type_id bigint, name varchar(30), additional_info varchar(30)) as $$
begin
    return query select com.id, com.vendor_id, v.name, com.type_id, com.name, com.additional_info from component as com inner join vendor v on v.id = com.vendor_id inner join type t on t.id = com.type_id
    where t.type = type_new;
end;
$$ language 'plpgsql';

/* функция 15 - выполнение заказа */
create or replace function do_order(order_id_new integer)
returns void as $$
begin
    delete from shipment where shipment.order_id = order_id_new;
    delete from "order" where "order".id = order_id_new;

end;
$$ language 'plpgsql';


insert into country(gdp_usd, name) values
                                       (124425.64, 'Russia'),
                                       (11122134.65, 'USA'),
                                       (1212125.15, 'Japan'),
                                       (24214.51, 'France'),
                                       (14144.51, 'Mongolia'),
                                       (1414.51, 'Kazakhstan'),
                                       (1667.13, 'Uzbekistan'),
                                       (200000, 'United Kingdom'),
                                       (443242331, 'Germany');

insert into subsidy (country_id, require_price, required_wd) values
                                                                 (1, 15000, '4wd'),
                                                                 (2, 250000, '4wd'),
                                                                 (3, 300000, 'rwd'),
                                                                 (4, 1100000, 'fwd'),
                                                                 (5, 15000, '4wd'),
                                                                 (6, 15000, 'fwd'),
                                                                 (7, 15000, '4wd');

insert into country_tax (country_id, tax) values
                                              (1, 1000),
                                              (2, 2000),
                                              (3, 3000),
                                              (4, 4000),
                                              (5, 5000),
                                              (6, 6000),
                                              (7, 7000),
                                              (8, 8000),
                                              (9, 9000);

insert into vendor(country_id, capitalization, name) values
                                                         (1, 13135.24, 'UAZ'),
                                                         (1, 1556.67, 'LADA'),
                                                         (1, 35635.77, 'ZIL'),
                                                         (1, 23725.88, 'KAMAZ'),
                                                         (2, 356365.78, 'Ford'),
                                                         (2, 36363572.35, 'Chevrolet'),
                                                         (2, 24747247.66, 'Cadillac'),
                                                         (2, 24574474.12, 'Lincoln'),
                                                         (2, 247247.22, 'Buick'),
                                                         (2, 496559.23, 'Jeep'),
                                                         (2, 3559468.33, 'Pontiac'),
                                                         (3, 45385974.66, 'Toyota'),
                                                         (3, 358453.67, 'Mazda'),
                                                         (3, 34687456945383466.79, 'Nissan'),
                                                         (4, 3684374.76, 'Renault'),
                                                         (4, 34678453.56, 'Citroen'),
                                                         (4, 586947958.54, 'Peugeot'),
                                                         (8, 42234212, 'Land Rover'),
                                                         (8, 32131312, 'Jaguar'),
                                                         (8, 321312, 'Aston Martin'),
                                                         (9, 5453453, 'BMW'),
                                                         (9, 324235245, 'Mercedes-Benz'),
                                                         (9, 35463635346356, 'Volkswagen');

insert into type (type, additional_info) values
                                             ('engine', 'v6'),
                                             ('engine', 'i4'),
                                             ('engine', 'v8'),
                                             ('engine', 'i6'),
                                             ('engine', 'electric'),
                                             ('door', 'front left'),
                                             ('door', 'front right'),
                                             ('door', 'back left'),
                                             ('door', 'back right'),
                                             ('bumper', 'front'),
                                             ('bumper', 'back'),
                                             ('transmission', 'manual'),
                                             ('transmission', 'automatic'),
                                             ('transmission', 'variator'),
                                             ('transmission', 'robot');

insert into component (vendor_id, type_id, name, additional_info) values
                                                                      (1, 1, 'ZMZ406', 'petrol'),
                                                                      (15, 2, 'K4M', 'petrol'),
                                                                      (15, 3, 'D4F', 'petrol'),
                                                                      (15, 2, 'F7R', 'petrol'),
                                                                      (15, 2, 'K9K', 'diesel'),
                                                                      (15, 1, 'P9X', 'diesel'),
                                                                      (14, 2, 'MR20DE', 'petrol'),
                                                                      (14, 2, 'QR20DE', 'petrol'),
                                                                      (14, 3, 'VH45DE', 'petrol'),
                                                                      (14, 4, 'QD32', 'diesel'),
                                                                      (15, 6, 'RRD', 'diesel'),
                                                                      (21, 5, 'EV', 'EvEngine'),
                                                                      (15, 6, 'RenaultDoorFL', 'Red'),
                                                                      (15, 7, 'RenaultDoorFR', 'Red'),
                                                                      (15, 8, 'RenaultDoorBL', 'Red'),
                                                                      (15, 9, 'RenaultDoorBR', 'Red'),
                                                                      (14, 6, 'NissanDoorFL', 'Yellow'),
                                                                      (14, 7, 'NissanDoorFR', 'Yellow'),
                                                                      (14, 8, 'NissanDoorBL', 'Yellow'),
                                                                      (14, 9, 'NissanDoorBR', 'Yellow'),
                                                                      (21, 6, 'BMWEvDoorFL', 'Black'),
                                                                      (21, 7, 'BMWEvDoorFR', 'Black'),
                                                                      (21, 8, 'BMWEvDoorBL', 'Black'),
                                                                      (21, 9, 'BMWEvDoorBR', 'Black'),
                                                                      (14, 12, 'RS6F94R', 'MT 6-speed'),
                                                                      (14, 14, 'JF010E', 'CVT'),
                                                                      (15, 12, 'JH1', 'MT 5-speed'),
                                                                      (15, 13, 'DP2', 'AT 3-speed'),
                                                                      (21, 15, 'DCT', 'Robot'),
                                                                      (21, 13, '6HP19', 'AT 6-speed'),
                                                                      (15, 10, 'JFE3', 'REDFR'),
                                                                      (14, 10, 'F44E3', 'RE4DFR'),
                                                                      (15, 10, 'Jflr3', 'RFR');


insert into factory (vendor_id, max_workers, productivity) values
                                                               (1, 1000, 10),
                                                               (2, 1000, 10),
                                                               (3, 1000, 10),
                                                               (4, 1000, 10),
                                                               (5, 1000, 10),
                                                               (6, 1000, 10),
                                                               (7, 1000, 10),
                                                               (8, 1000, 10),
                                                               (9, 1000, 10),
                                                               (10, 1000, 10),
                                                               (11, 1000, 10),
                                                               (12, 1000, 10),
                                                               (13, 1000, 10),
                                                               (14, 1000, 10),
                                                               (15, 1000, 10),
                                                               (16, 1000, 10),
                                                               (17, 1000, 10),
                                                               (18, 1000, 10),
                                                               (19, 1000, 10),
                                                               (20, 1000, 10),
                                                               (21, 1000, 10),
                                                               (22, 1000, 10),
                                                               (23, 1000, 10);


insert into engineer (vendor_id, name, gender, experience, salary, factory_id) values
                                                                                   (1, 'Andrew', 'male', 13, 10000, 1),
                                                                                   (2, 'Peter', 'male', 13, 10000, 2),
                                                                                   (3, 'Amanda', 'female', 13, 10000, 3),
                                                                                   (3, 'Sergey', 'male', 13, 153764, 3),
                                                                                   (3, 'Anton', 'male', 13, 10000, 3),
                                                                                   (4, 'Zakhar', 'male', 13, 13566, 4),
                                                                                   (5, 'Vladimir', 'male', 13, 135135, 5),
                                                                                   (6, 'Josef', 'male', 10, 21143, 6),
                                                                                   (7, 'Takida', 'male', 13, 10000, 7),
                                                                                   (8, 'Gregory', 'male', 13, 111113, 8),
                                                                                   (9, 'Hans', 'male', 13, 26242, 9),
                                                                                   (9, 'Frans', 'male', 13, 357357, 9),
                                                                                   (10, 'Mikhail', 'male', 13, 12454, 10),
                                                                                   (11, 'Alexey', 'male', 13, 124524, 11),
                                                                                   (11, 'Kirill', 'male', 13, 12455, 11),
                                                                                   (12, 'Pavel', 'male', 13, 999999, 12),
                                                                                   (13, 'Dmitriy', 'male', 4, 14342342, 13),
                                                                                   (14, 'Kate', 'female', 20, 143112, 14),
                                                                                   (15, 'Kostyan', 'male', 40, 23253453453, 15),
                                                                                   (16, 'Alex', 'male', 10, 12423423, 16),
                                                                                   (17, 'Thomas', 'male', 15, 214134, 17),
                                                                                   (18, 'Abgdish', 'male', 12, 12143242, 18),
                                                                                   (19, 'Gadya', 'female', 4, 14234, 19),
                                                                                   (20, 'Pepepopo', 'male', 10, 42523423, 20),
                                                                                   (21, 'Vladimir', 'male', 9, 314342, 21),
                                                                                   (22, 'Adolf', 'male', 41, 534532, 22),
                                                                                   (23, 'Mikhail', 'male', 12, 445345, 23);


insert into model (vendor_id, name, wheeldrive, significance, price, prod_cost, engineer_id, factory_id, sales) values
                                                                                                                    (1, 'Patriot', '4wd', 1000000, 2990, 1000, 1, 1, 10000),
                                                                                                                    (1, 'PatrNoEngine', '4wd', 1000000, 2990, 1000, 1, 1, 10000),
                                                                                                                    (1, 'Buhanka', '4wd', 1000000, 2990, 1000, 1, 1, 10000),
                                                                                                                    (14, 'Qashqai', 'fwd', 100, 5000, 2000, 18, 14, 10000000),
                                                                                                                    (15, 'Megane', 'fwd', 100, 5000, 2500, 19, 15, 100000),
                                                                                                                    (21, '318', 'rwd', 200, 6000, 3000, 23, 21, 1000);

insert into model_component (model_id, component_id) values
                                                         (1, 1),
                                                         (3, 1),
                                                         (4, 7),
                                                         (4, 17),
                                                         (4, 18),
                                                         (4, 19),
                                                         (4, 20),
                                                         (4, 25),
                                                         (5, 2),
                                                         (5, 13),
                                                         (5, 14),
                                                         (5, 15),
                                                         (5, 16),
                                                         (5, 28),
                                                         (6, 30),
                                                         (6, 21),
                                                         (6, 21),
                                                         (6, 22),
                                                         (6, 23),
                                                         (6, 24);

insert into "order" (model_id, quantity, order_type) values
                                                         (1, 1000, 'gos'),
                                                         (2, 1, 'gos'),
                                                         (3, 5000, 'retail'),
                                                         (4, 1000000, 'retail'),
                                                         (5, 294340, 'retail'),
                                                         (6, 421341, 'retail');

insert into storage (model_id, amount, vendor_id) values
                                                      (1, 1500, 1),
                                                      (2, 1000, 1),
                                                      (3, 100000, 1),
                                                      (4, 10000000, 14),
                                                      (5, 294341, 15),
                                                      (6, 4213411, 21);

insert into shipment (order_id, country_to_id, date, cost) values
                                                               (1, 5, '2022-10-20', 100),
                                                               (2, 3, '2022-12-31', 100),
                                                               (3, 6, '2022-09-30', 100),
                                                               (4, 8, '2022-11-10', 2000),
                                                               (5, 9, '2022-12-21', 1000),
                                                               (6, 1, '2024-01-01', 50000);