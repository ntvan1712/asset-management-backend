package enums

type assetHistoryTypeEnum struct {
	Create    string
	Update    string
	ByManager string
}

var AssetHistoryType = assetHistoryTypeEnum{
	Create:    "create",
	Update:    "update",
	ByManager: "by_manager",
}
