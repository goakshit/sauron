package constants

const (
	CreateMerchantInvalidParamsErr = "Invalid number of params passed for create merchant call"

	CreateMerchantPercErr = "Invalid percentage passed for create merchant"

	CreateMerchantDuplicateIDErr = "Duplicate merchant id passed. Kindly pass a unique id"

	CreateMerchantDuplicateEmailErr = "Duplicate merchant email passed. Kindly pass a unique email"

	UpdateMerchantInvalidDiscountErr = "Invalid discount passed for update."

	UpdateMerchantInvalidParamsErr = "Invalid number of params passed for update merchant discount call"

	UpdateMerchantPercErr = "Invalid(Unparsable) percentage passed for update merchant discount call"

	UpdateMerchantNameMissingErr = "Name cannot be empty for update merchant discount call"

	UpdateMerchantNotFoundErr = "Merchant with given name not found in database."

	CreateUserInvalidParamsErr = "Invalid number of params passed for create user call"

	CreateUserInvalidCreditLimitErr = "Invalid(or unparsable) credit limit passed for create user call"

	CreateUserDuplicateIDErr = "Duplicate user id passed. Kindly pass a unique id"

	CreateUserDuplicateEmailErr = "Duplicate user email passed. Kindly pass a unique email"

	CreateTxnUserDoesNotExistErr = "User with id passed doesn't exist in database"

	CreateTxnGetUserCreditLimitErr = "Failed to get user credit limit. Please try again later"

	CreateTxnInvalidAmountErr = "The transaction amount is invalid(or unparsable)"

	CreateTxnUserCreditLimitExceededErr = "Credit limit exceeed for this user"

	CreateTxnGetMerchantErr = "Something went wrong while fetching merchant"

	CreateTxnMerchantNotFoundErr = "Merchant with id passed not found"

	CreateTxnInvalidParamsErr = "Invalid number of params passed for create txn call"
)
