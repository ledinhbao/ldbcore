package ldbcore

import "github.com/jinzhu/gorm"

// Setting contains server configuration, database-based.
//
// Read from database, table [settings], column key & value
type Setting struct {
	gorm.Model
	Key   string
	Value string
}

// Set (key, value) used to quickly set key-value for Setting object.
func (s *Setting) Set(key string, value string) {
	s.Key = key
	s.Value = value
}
