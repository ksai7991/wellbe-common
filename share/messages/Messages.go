package messages

const (
	MESSAGE_EN_SUCCESS                      = "Process completed successfully."
	MESSAGE_EN_NUMBERING_OVERFLOW           = "The numbering process has overflowed."
	MESSAGE_EN_NUMBERING_NOTEXISTS          = "The numbering key is not exists. Item:"
	MESSAGE_EN_REQUEST_VALUE_INVALID        = "Request value is invalid. Item:"
	MESSAGE_EN_REQUEST_ITEM_MANDATORY       = "Request item is mandatory. Item: %v"
	MESSAGE_EN_REQUEST_ITEM_OUT_OF_LENGTH   = "Request value need to less then %v length. Item: %v"
	MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT = "Request value is invalid format. Item: %v Format: %v"
	MESSAGE_EN_NOTEXISTS_UPDATE             = "Update data does not exist."
	MESSAGE_EN_NOTEXISTS_DELETE             = "Delete data does not exist."
	MESSAGE_EN_NOTEXISTS_FOREINGKEY         = "Specified key does not exist. key: "
	MESSAGE_EN_SERVER_ERROR                 = "There is a problem on the server. Please wait for a while."
	MESSAGE_EN_PLEASE_LOGIN                 = "Please login."
	MESSAGE_EN_MASTER_DATA_IS_UNSETUP       = "Mater data is unsetuped. Table: %v Reason: %v"
	MESSAGE_EN_EXCHANGE_PAIRE_UNSETUP       = "Exchange paire is unsetuped. currency_cd1: %v currency_cd2: %v"
	MESSAGE_EN_COMPARE_STRING_IS_INVALID    = "Compare string is invalid."
	MESSAGE_EN_RECAPTCHA_FAIL               = "Recaptch has failed. detail: %v"
	MESSAGE_EN_DATETIME_FORMAT_IS_INVALID   = "Date time format is invalid. s: %v"
	MESSAGE_EN_S3_GETOBJECT_ERROR           = "S3 Get Object Error. detail: %v"
	MESSAGE_EN_IMAGE_CONVERT_ERROR          = "Failed Image Convert Error"
)