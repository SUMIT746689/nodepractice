package domain

type USER_PERMISSION string

const (
	CREATE_USER          USER_PERMISSION = "create_user"
	CREATE_ADMIN_USER    USER_PERMISSION = "create_admin_user"
	CREATE_CASHIER_USER  USER_PERMISSION = "create_cashier_user"
	CREATE_CUSTOMER_USER USER_PERMISSION = "create_customer_user"

	UPDATE_USER          USER_PERMISSION = "update_user"
	UPDATE_ADMIN_USER    USER_PERMISSION = "update_admin_user"
	UPDATE_CASHIER_USER  USER_PERMISSION = "update_cashier_user"
	UPDATE_CUSTOMER_USER USER_PERMISSION = "update_customer_user"

	DELETE_USER          USER_PERMISSION = "delete_user"
	DELETE_ADMIN_USER    USER_PERMISSION = "delete_admin_user"
	DELETE_CASHIER_USER  USER_PERMISSION = "delete_cashier_user"
	DELETE_CUSTOMER_USER USER_PERMISSION = "delete_customer_user"

	INDEX_COMPANY  USER_PERMISSION = "index_company"
	CREATE_COMPANY USER_PERMISSION = "create_company"
	UPDATE_COMPANY USER_PERMISSION = "update_company"
	DELETE_COMPANY USER_PERMISSION = "delete_company"
)
