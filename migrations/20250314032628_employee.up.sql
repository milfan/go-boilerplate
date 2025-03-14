create table employees (
	id bigserial not null,
	created_by varchar(75) not null,
	created_at timestamptz not null default now(),
	updated_by varchar(75) not null,
	updated_at timestamptz not null default now(),
	deleted_at timestamptz null,
	emp_code varchar(50) not null,
	emp_name varchar(75) not null,
	constraint employees_pkey primary key (id)
);