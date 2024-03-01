package collections

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

// "Service" collection: title, description, price.

var ServicesCollection = &models.Collection{
	Name:       "services",
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
			Name:     "description",
			Type:     schema.FieldTypeText,
			Required: true,
		},
		&schema.SchemaField{
			Name:     "price",
			Type:     schema.FieldTypeNumber,
			Required: true,
		},
	),
}
