package enums

type assetStatusEnum struct {
	Available        string
	OnBorrow         string
	AwaitingTransfer string
	Unusable         string
}

var AssetStatusEnum = assetStatusEnum{
	Available:        "available",
	OnBorrow:         "on_borrow",
	AwaitingTransfer: "awaiting_transfer",
	Unusable:         "unusable",
}
