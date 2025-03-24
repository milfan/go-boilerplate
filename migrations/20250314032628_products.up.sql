create table products (
	id bigserial not null,
	created_by varchar(75) not null,
	created_at timestamptz not null default now(),
	updated_by varchar(75) not null,
	updated_at timestamptz not null default now(),
	deleted_at timestamptz null,
	product_code varchar(50) not null,
	product_name varchar(100) not null,
	product_price bigint not null,
	constraint products_pkey primary key (id)
);