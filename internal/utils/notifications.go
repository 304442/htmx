package utils

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func SendNotifications(app *pocketbase.PocketBase, reason string , message string) error {
	notificationsCollection, err := app.Dao().FindCollectionByNameOrId("notifications")
	if err != nil{
		return err
	}
	newNotificationRecord := models.NewRecord(notificationsCollection)
	newNotificationRecord.Set("reason", reason)
	newNotificationRecord.Set("message", message)
	return app.Dao().Save(newNotificationRecord)
}
