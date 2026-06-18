create table storefronts (
    id bigint unsigned not null auto_increment,
    storefront_name varchar(255) not null,
    marketplace_storefront_id varchar(100) not null,
    is_active tinyint(1) not null default 1,
    balance decimal(19, 6) null,
    country_code varchar(100) not null,
    cookie varchar(8000) not null,
    data_json json null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime null,
    primary key (id),
    unique key marketplace_storefront_id_unique(marketplace_storefront_id)
) auto_increment = 100001;