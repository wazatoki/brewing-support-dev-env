create table join_staffs_staff_groups (
    staffs_id text not null REFERENCES staffs (id)
    , staff_groups_id text not null REFERENCES staff_groups (id)
    , primary key(staffs_id, staff_groups_id)
);