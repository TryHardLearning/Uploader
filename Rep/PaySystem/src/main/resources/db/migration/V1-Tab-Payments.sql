create table payments(
    id serial primary key;

    price decimal(19,1);
    date timezone;
    client varchar(22);
    status varchar(15);

    demand_fk integer;
)