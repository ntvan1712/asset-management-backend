package entity

type AllCategoriesEntity struct {
	AssetQualities []AssetQualityEntity `json:"asset_qualities"`
	AssetTypes     []AssetTypeEntity    `json:"asset_types"`
	Locations      []LocationEntity     `json:"locations"`
	Currencies     []CurrencyEntity     `json:"currencies"`
}
