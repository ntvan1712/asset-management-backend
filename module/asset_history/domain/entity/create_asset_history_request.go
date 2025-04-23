package entity

type CreateAssetHistoryRequest struct {
	Title   string  `json:"title"  validate:"required"`
	Content *string `json:"content,omitempty"`

	// From params 
	AssetID   *int `json:"asset_id"`
	// From token
	CreatorID *int `json:"creator_id"`
}
