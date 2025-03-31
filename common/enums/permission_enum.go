package enums

type permissionIDEnum struct {
	AssetManagement    int
	BorrowManagement   int
	Statistical        int
	CategoryManagement int
}

var PermissionID = permissionIDEnum{
	AssetManagement:    1,
	BorrowManagement:   2,
	Statistical:        3,
	CategoryManagement: 4,
}
