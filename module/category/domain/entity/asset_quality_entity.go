package entity

type AssetQualityEntity struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

