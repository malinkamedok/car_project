-- # Поиск автомобилей с заданным компонентом

select model.name from model
    inner join model_component on model.id = model_component.model_id
    inner join component on model_component.component_id = component.id
    inner join type on component.type_id = type.id
        where type.type = 'engine';

-- # Просмотр всех компонентов конкретного автомобиля

select component.name, component.additional_info, type.type, type.additional_info from component
    inner join type on component.type_id = type.id
    inner join model_component on component.id = model_component.component_id
    inner join model on model_component.model_id = model.id
        where model.name = 'Patriot';

-- # Получение заказов по вендору

select vendor.name as vendor, model.name as model, "order".id as order_id, "order".quantity, "order".order_type from "order"
    inner join model on "order".model_id = model.id
    inner join vendor on model.vendor_id = vendor.id
        where vendor.id = 1;

-- # Получение заказов по стране

select country.id as country_id, country.name as country, vendor.name as vendor, model.name as model, "order".id as order_id, "order".quantity, "order".order_type, shipment.cost as shipment_cost, shipment.date from "order"
   inner join model on "order".model_id = model.id
   inner join vendor on model.vendor_id = vendor.id
   inner join shipment on "order".id = shipment.order_id
   inner join country on shipment.country_to_id = country.id
        where country.id = 1;

-- # Получение субсидий по стране

select country.name as country, subsidy.require_price, subsidy.required_wd from subsidy
    inner join country on subsidy.country_id = country.id
        where country.id = 1;

-- # Получение доставок по стране

select country.name as country, shipment.id, shipment.date, shipment.cost from shipment
    inner join country on shipment.country_to_id = country.id
        where country_to_id = 1;