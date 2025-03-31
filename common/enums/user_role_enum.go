package enums

type userRoleIDEnum struct {
	Admin    int
	Manager  int
	Employee int
}

var UserRoleID = userRoleIDEnum{
	Admin:    1,
	Manager:  2,
	Employee: 3,
}

