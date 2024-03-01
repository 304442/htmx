package collections

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

// title, content, author, date.

var BlogCollection = &models.Collection{
	Name:       "blog",
	Type:       models.CollectionTypeBase,
	ListRule:   nil,
	DeleteRule: nil,
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:     "title",
			Type:     schema.FieldTypeText,
			Required: true,
			Options: &schema.TextOptions{
				Max: types.Pointer(10),
			},
		},
		&schema.SchemaField{
			Name:     "author",
			Type:     schema.FieldTypeText,
			Required: true,
		},
		&schema.SchemaField{
			Name:     "content",
			Type:     schema.FieldTypeText,
			Required: true,
		},
	),
}
