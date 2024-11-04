CREATE TABLE IF NOT EXISTS wellbe_common.currency_exchange_rate (
  base_currency_cd integer not null
  , target_currency_cd integer not null
  , paire_name varchar(6) not null
  , rate numeric(30,10) not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint currency_exchange_rate_PKC primary key (base_currency_cd,target_currency_cd)
) ;
GRANT ALL ON wellbe_common.currency_exchange_rate TO wellbe;