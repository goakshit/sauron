package constants

const (

	// Create Merchant

	CreateMerchantInvalidParamsErr = "Invalid number of params passed for create merchant call"

	CreateMerchantPercErr = "Invalid percentage passed for create merchant"

	CreateMerchantDuplicateIDErr = "Duplicate merchant id passed. Kindly pass a unique id"

	CreateMerchantDuplicateEmailErr = "Duplicate merchant email passed. Kindly pass a unique email"

	CreateMerchantInvalidDiscountErr = "Invalid discount passed for create merchant call."

	// Update Merchant

	UpdateMerchantInvalidDiscountErr = "Invalid discount passed for update."

	UpdateMerchantInvalidParamsErr = "Invalid number of params passed for update merchant discount call"

	UpdateMerchantPercErr = "Invalid(Unparsable) percentage passed for update merchant discount call"

	UpdateMerchantNameMissingErr = "Name cannot be empty for update merchant discount call"

	UpdateMerchantNotFoundErr = "Merchant with given name not found in database."

	// Create User

	CreateUserInvalidParamsErr = "Invalid number of params passed for create user call"

	CreateUserInvalidCreditLimitErr = "Invalid(or unparsable) credit limit passed for create user call"

	CreateUserDuplicateIDErr = "Duplicate user id passed. Kindly pass a unique id"

	CreateUserDuplicateEmailErr = "Duplicate user email passed. Kindly pass a unique email"

	// Update User

	UpdateUserNameMissingErr = "Name cannot be empty for update user credit call"

	UpdateUserInvalidCreditLimitErr = "Credit limit is invalid(or unparsable)"

	UpdateUserNotFoundErr = "User with given name not found in database."

	UpdateUserInvalidParamsErr = "Invalid number of params passed for update user call"

	// Create Txn

	CreateTxnUserDoesNotExistErr = "User with id passed doesn't exist in database"

	CreateTxnGetUserCreditLimitErr = "Failed to get user credit limit. Please try again later"

	CreateTxnInvalidAmountErr = "The transaction amount is invalid(or unparsable)"

	CreateTxnUserCreditLimitExceededErr = "Credit limit exceeed for this user"

	CreateTxnGetMerchantErr = "Something went wrong while fetching merchant"

	CreateTxnMerchantNotFoundErr = "Merchant with id passed not found"

	CreateTxnInvalidParamsErr = "Invalid number of params passed for create txn call"

	// Create Payback

	CreatePaybackInvalidAmountErr = "The payback amount is invalid(or unparsable)"

	CreatePaybackInvalidParamsErr = "Invalid number of params passed for create payback call"

	CreatePaybackUserNotFoundErr = "User with id passed not found"

	CreatePaybackGetDueAmountErr = "Failed to get user's due amount. Try again after sometime."

	CreatepaybackNoDueAmountErr = "There is no due amount for this user."

	// Report

	ReportUACLGetUsersErr = "Failed to fetch users at credit limit. Try again after sometime."

	ReportUACLNoRecordsErr = "There are no users at credit limit."

	ReportUserDuesGetUsersErr = "Failed to fetch users with pending dues. Try again after sometime."

	ReportUserDuesNoRecordsErr = "There are no users with pending dues."

	ReportUserDuesUserNotFoundErr = "User with id passed, not found"

	ReportUserDuesGetUserDueErr = "Failed to fetch user with pending dues. Try again after sometime."

	ReportInvalidParamErr = "Invalid number of params passed"

	ReportDiscountGetTxnErr = "Failed to fetch transactions for this merchant. Try again later."
)
