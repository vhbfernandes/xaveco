package models

import (
	"github.com/kamva/mgm/v3"
)

// Xaveco is a generic database model with Content and Tags as the fields;
// Content must be string and Tags is a slice of strings
type Xaveco struct {
	mgm.DefaultModel `bson:",inline"`
	Content              string     `json:"content" bson:"content"`
	Tags                 []string   `json:"tags" bson:"tags"`
}
