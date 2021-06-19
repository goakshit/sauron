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
)
