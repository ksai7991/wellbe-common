create table IF NOT EXISTS wellbe_common.c_mail_template (
  mail_template_cd integer not null
  , language_cd integer not null
  , subject text
  , body text
  , create_datetime character varying not null
  , create_function character varying not null
  , update_datetime character varying
  , update_function character varying
  , constraint c_mail_template_PKC primary key (mail_template_cd,language_cd)
) ;
GRANT ALL ON wellbe_common.c_mail_template TO wellbe;