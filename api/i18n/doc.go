package i18n

const (
	ValidationErrorList                     = "Failed to pass list info. %s"
	ValidationWarningLabelToLong            = "The label cannot be longer than 20 characters."
	ValidationWarningLabelNotEmpty          = "The label cannot be empty."
	ValidationAlistTypeV3                   = "Please refer to the documentation on list type v3"
	SuccessAlistNotFound                    = "List not found."
	SuccessUserNotFound                     = "User not found."
	InternalServerErrorMissingAlistUuid     = "Uuid is missing, possibly an internal error"
	InternalServerErrorMissingUserUuid      = "User.Uuid is missing, possibly an internal error"
	InternalServerErrorTalkingToDatabase    = "Issue with talking to the database in %s."
	InputDeleteAlistOperationOwnerOnly      = "Only the owner of the list can remove it."
	InputSaveAlistOperationOwnerOnly        = "Only the owner of the list can modify it."
	PostUserLabelJSONFailure                = "Your input is invalid json."
	PostShareListJSONFailure                = "Your input is invalid json."
	InputMissingListUuid                    = "The uuid is missing."
	InternalServerErrorDeleteAlist          = "We have failed to remove your list."
	ApiMethodNotSupported                   = "This method is not supported."
	ApiAlistNotFound                        = "Failed to find alist with uuid: %s"
	ApiDeleteAlistSuccess                   = "List %s was removed."
	ApiDeleteUserLabelSuccess               = "Label %s was removed."
	UserInsertAlreadyExistsPasswordNotMatch = "Failed to save."
	UserInsertUsernameExists                = "Username already exists."
	DatabaseLookupNotFound                  = "sql: no rows in result set"
	AclHttpAccessDeny                       = "Access Denied"
)
