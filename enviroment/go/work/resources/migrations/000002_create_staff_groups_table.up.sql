create table staff_groups (
	id text primary key
	, del boolean not null default false
	, created_at timestamp
	, cre_staff_id text
	, updated_at timestamp
	, update_staff_id text
	,name text not null
);