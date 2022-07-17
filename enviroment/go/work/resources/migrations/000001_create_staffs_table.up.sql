create table staffs (
	id text primary key
	, del boolean not null default false
	, created_at timestamp
	, cre_staff_id text
	, updated_at timestamp
	, update_staff_id text
	, account_id text not null
	, password text not null
	, name text not null
);
