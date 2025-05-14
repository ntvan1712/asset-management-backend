package enums

type assetStatisticEnum struct {
	AssetType    string
	AssetQuality string
	Status       string
	Location     string
}

var AssetStatisticEnum = assetStatisticEnum{
	AssetType:    "asset_type",
	AssetQuality: "asset_quality",
	Status:       "status",
	Location:     "location",
}
