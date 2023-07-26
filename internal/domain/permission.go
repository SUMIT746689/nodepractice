package domain

type USER_PERMISSION string

const (
	CREATE_USER       USER_PERMISSION = "create_user"
	CREATE_ADMIN_USER USER_PERMISSION = "create_admin_user"
)
