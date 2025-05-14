package model

type CategoryStatisticModel struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type CategoryStatisticResponse struct {
	TotalAssetCount int                      `json:"total_asset_count"`
	Items           []CategoryStatisticModel `json:"items"`
}
