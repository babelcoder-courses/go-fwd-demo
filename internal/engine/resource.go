package engine

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Title string
}
