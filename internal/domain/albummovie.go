package domain

// AlbumMovie represents the many-to-many relationship between Album and Movie
type AlbumMovie struct {
    AlbumID uint `gorm:"primaryKey"`
    MovieID uint `gorm:"primaryKey"`
}
