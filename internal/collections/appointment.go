package collections

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)
// "Appointment" collection: user, date, time.
var AppointmentCollection = &models.Collection{
    Name:       "appointment",
    Type:       models.CollectionTypeBase,
    ListRule:   nil,
    DeleteRule: nil,
    Schema:     schema.NewSchema(
        &schema.SchemaField{
            Name:     "user",
            Type:     schema.FieldTypeEmail,
            Required: true,
        },
        &schema.SchemaField{
            Name:     "date",
            Type:     schema.FieldTypeDate,
            Required: true,
        },
    ),
}
