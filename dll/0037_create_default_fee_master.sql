create table IF NOT EXISTS wellbe_common.default_fee_master (
  id text not null
  , fee_rate numeric(5,3) not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint default_fee_master_PKC primary key (id)
) ;
GRANT ALL ON wellbe_common.default_fee_master TO wellbe;