package collections

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

// "QuoteRequest" collection: user, details, timestamp.

var QuoteCollection = &models.Collection{
	Name:       "quotes",
	Type:       models.CollectionTypeBase,
	ListRule:   nil,
	DeleteRule: nil,
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:     "user",
			Type:     schema.FieldNameEmail,
			Required: true,
		},
		&schema.SchemaField{
			Name:     "details",
			Type:     schema.FieldTypeText,
			Required: true,
		},
		&schema.SchemaField{
			Name:     "date",
			Type:     schema.FieldTypeDate,
			Required: true,
		},
	),
}
