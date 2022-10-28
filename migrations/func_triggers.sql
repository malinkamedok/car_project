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
significance integer, price decimal, prod_cost decimal, engineer_id bigint, factory_id bigint, sales bigint) as $$
begin
    return query select * from model;
end;
$$ language 'plpgsql';



/* функций 2 - Вывод всех супсидий */
create or replace function get_all_subsidies()
returns table(id integer, country_id bigint, require_price decimal, required_wd varchar(3)) as $$
begin
    return query select * from subsidy;
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
create or replace function create_order(model_id_by bigint, quantity_by bigint, order_type_by varchar(50))
returns void as $$
begin
    insert into "order" (model_id, quantity, order_type)  values (model_id_by, quantity_by, order_type_by);
end
$$ language 'plpgsql';



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


/* пример запроса с входнымми данными к 10 функции */
select accept_subsidies(
    1,
    1,
    'loooool',
    69,
    1,
    1,
    1,
    2,
    3,
    4
    );

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


/* пример запроса с входными данными к функции 11 */
select create_model(1, 'Санечка, снимаешь?', 1, 1, 1, 1, 1);


select vendor.name as vendor, model.name as model, "order".id as order_id, "order".quantity, "order".order_type from "order"
                                                                                                                         inner join model on "order".model_id = model.id
                                                                                                                         inner join vendor on model.vendor_id = vendor.id
where vendor.id = 1;