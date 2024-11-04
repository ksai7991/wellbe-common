CREATE TABLE IF NOT EXISTS wellbe_common.c_currency (
  currency_cd integer not null
  , language_cd integer not null
  , currency_name text not null
  , currency_cd_iso varchar(3) not null
  , significant_digit integer not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_currency_PKC primary key (currency_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_currency TO wellbe;