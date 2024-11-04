create table IF NOT EXISTS wellbe_common.c_account_withdrawal_reason (
  account_withdrawal_reason_cd integer not null
  , language_cd integer not null
  , account_withdrawal_reason_name text not null
  , account_withdrawal_reason_abbreviation text not null
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_account_withdrawal_reason_PKC primary key (account_withdrawal_reason_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_account_withdrawal_reason TO wellbe;