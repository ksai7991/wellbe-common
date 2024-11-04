create table IF NOT EXISTS wellbe_common.currency_for_payment (
  currency_cd integer not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint currency_for_payment_PKC primary key (currency_cd)
) ;
GRANT ALL ON wellbe_common.currency_for_payment TO wellbe;