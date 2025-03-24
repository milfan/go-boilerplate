create table order_details (
	id bigserial not null,
	created_by varchar(75) not null,
	created_at timestamptz not null default now(),
	updated_by varchar(75) not null,
	updated_at timestamptz not null default now(),
	deleted_at timestamptz null,
	order_id bigint not null,
	product_id bigint not null,
	product_qty int not null,
	product_price bigint not null,
	constraint order_details_pkey primary key (id),
    constraint order_order_details_order_id_fkey foreign key (order_id) references orders(id)  on update cascade on delete no action,
    constraint order_details_products_product_id_fkey foreign key (product_id) references products(id)  on update cascade on delete no action
);