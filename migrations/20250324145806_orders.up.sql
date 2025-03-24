create table orders (
	id bigserial not null,
	created_by varchar(75) not null,
	created_at timestamptz not null default now(),
	updated_by varchar(75) not null,
	updated_at timestamptz not null default now(),
	deleted_at timestamptz null,
	order_code varchar(50) not null,
	order_date timestamptz not null default now(),
	constraint orders_pkey primary key (id)
);