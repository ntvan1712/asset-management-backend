package enums

type assetHistoryTypeEnum struct {
	Create         string
	Update         string
	ByManager      string
	Transfer       string
	CancelTransfer string
	RejectTransfer string
	OnBorrow       string
	Return         string
}

var AssetHistoryType = assetHistoryTypeEnum{
	Create:         "create",
	Update:         "update",
	Transfer:       "transfer",
	ByManager:      "by_manager",
	CancelTransfer: "cancel_transfer",
	RejectTransfer: "reject_transfer",
	OnBorrow:       "on_borrow",
	Return:         "return",
}
