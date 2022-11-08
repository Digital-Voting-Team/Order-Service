-- +migrate Up

create table if not exists public.statuses
(
    status_id   integer generated always as identity
        constraint statuses_pk
            primary key,
    status_name varchar not null
);

alter table public.statuses
    owner to postgres;

create table if not exists public.orders
(
    order_id       integer generated always as identity
        constraint orders_pk
            primary key,
    customer_id    integer          not null,
    staff_id       integer          not null,
    total_price    double precision not null,
    payment_method integer          not null,
    is_take_away   boolean          not null,
    status_id      integer          not null
        constraint orders_statuses_status_id_fk
            references public.statuses,
    cafe_id        integer          not null,
    order_date     timestamp with time zone not null
);

alter table public.orders
    owner to postgres;

create table if not exists public.order_items
(
    order_item_id integer generated always as identity
        constraint order_items_pk
            primary key,
    meal_id       integer not null,
    quantity      integer not null,
    order_id      integer not null
        constraint order_items_orders_order_id_fk
            references public.orders
);

alter table public.order_items
    owner to postgres;

create table if not exists public.addresses
(
    address_id   integer generated always as identity
        constraint addresses_pk
            primary key,
    building_number integer not null,
    street       varchar not null,
    city         varchar not null,
    district     varchar not null,
    region       varchar not null,
    postal_code  varchar not null
);

alter table public.addresses
    owner to postgres;

create table if not exists public.deliveries
(
    delivery_id    integer generated always as identity
        constraint deliveries_pk
            primary key,
    order_id       integer          not null
        constraint deliveries_orders_order_id_fk
            references public.orders,
    address_id     integer          not null
        constraint deliveries_addresses_address_id_fk
            references public.addresses,
    staff_id       integer          not null,
    delivery_price double precision not null,
    delivery_date  timestamp with time zone not null
);

alter table public.deliveries
    owner to postgres;

-- +migrate Down

drop table if exists public.deliveries;
drop table if exists public.addresses;
drop table if exists public.order_items;
drop table if exists public.orders;
drop table if exists public.statuses;
