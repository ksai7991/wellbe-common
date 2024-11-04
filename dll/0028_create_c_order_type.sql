create table IF NOT EXISTS wellbe_common.c_order_type (
  order_type_cd integer not null
  , language_cd integer not null
  , order_type_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_order_type_PKC primary key (order_type_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_order_type TO wellbe;