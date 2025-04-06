package model

type AllCategories struct {
	AssetQualities []AssetQuality `json:"asset_qualities"`
	AssetTypes     []AssetType    `json:"asset_types"`
	Locations      []Location     `json:"locations"`
	Currencies     []Currency     `json:"currencies"`
}
