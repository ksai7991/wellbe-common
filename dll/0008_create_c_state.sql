CREATE TABLE IF NOT EXISTS wellbe_common.c_state (
  state_cd integer not null
  , language_cd integer not null
  , country_cd integer not null
  , state_name text not null
  , state_cd_iso text
  , timezone_iana text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_state_PKC primary key (state_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_state TO wellbe;