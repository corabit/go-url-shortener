package entities

type Url struct {
	Id        string `gorm:"unique;primaryKey;not null"`
	OriginUrl string `gorm:"not null"`
	Clicks    int    `gorm:"not null"`
	LastClick int64  `gorm:"autoUpdateTime"`
	Created   int64  `gorm:"autoCreateTime"`
}
