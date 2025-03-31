package enums

type userAuthorityEnum struct {
	Admin              string
	AssetManagement    string
	BorrowManagement   string
	Statistical        string
	CategoryManagement string
	Employee           string
}

var UserAuthority = userAuthorityEnum{
	Admin:              "Admin",
	AssetManagement:    "AssetManagement",
	BorrowManagement:   "BorrowManagement",
	Statistical:        "Statistical",
	CategoryManagement: "CategoryManagement",
	Employee:           "Employee",
}
