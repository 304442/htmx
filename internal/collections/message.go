package collections

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

// "Chat" collection: sender, receiver, message, timestamp.

var MessageCollection = &models.Collection{
	Name:       "messages",
	Type:       models.CollectionTypeBase,
	ListRule:   nil,
	DeleteRule: nil,
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:     "sender",
			Type:     schema.FieldNameEmail,
			Required: true,
		},
		&schema.SchemaField{
			Name:     "reciever",
			Type:     schema.FieldNameEmail,
			Required: true,
		},
		&schema.SchemaField{
			Name:     "message",
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
