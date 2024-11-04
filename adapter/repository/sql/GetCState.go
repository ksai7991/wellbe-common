package sql

const (
	GetCState = `
  WITH PARAMS AS (
    SELECT
      $1 AS STATE_NAME
  )
  SELECT
  A.state_cd
  ,A.language_cd
  ,A.country_cd
  ,A.state_name
  ,A.state_cd_iso
  ,A.timezone_iana
 FROM wellbe_common.c_state AS A
 CROSS JOIN PARAMS AS P
 WHERE 1=1
   AND (P.STATE_NAME = '' OR A.STATE_NAME ILIKE '%'||P.STATE_NAME||'%')
 ORDER BY A.country_cd, A.state_cd_iso
`
)
