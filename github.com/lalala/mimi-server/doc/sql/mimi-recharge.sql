SET client_min_messages = warning;


CREATE TABLE recharge_order (
    id SERIAL primary key not null,
    order_no text NOT NULL,
    account_id integer NOT NULL,
    money integer NOT NULL,
    status integer NOT NULL,
    os character varying(7) NOT NULL,
    client_req_time integer NOT NULL,
    create_time integer NOT NULL,
    uc_resp_time integer NOT NULL,
    payment_method integer NOT NULL
);