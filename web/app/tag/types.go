package tag

import "gorm.io/gorm"

type tag struct {
	gorm.Model
	Name string `json:"name"`
}