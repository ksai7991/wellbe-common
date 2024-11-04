create table IF NOT EXISTS wellbe_common.c_booking_status (
  booking_status_cd integer not null
  , language_cd integer not null
  , booking_status_name text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_booking_status_PKC primary key (booking_status_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_booking_status TO wellbe;