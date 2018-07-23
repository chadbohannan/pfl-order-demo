package gopfl

import (
	"context"

	"google.golang.org/appengine/datastore"
)

// TODO store test service hostname
// TODO store Basic Authentication credentials

const settingsKind = "Settings"

// Setting encapsulates a string value
type Setting struct {
	Value string
}

// PutSetting persists a value for a key
func PutSetting(c context.Context, key, value string) error {
	k := datastore.NewKey(c, settingsKind, key, 0, nil)
	if _, err := datastore.Put(c, k, &Setting{value}); err != nil {
		return err
	}
	return nil
}

// GetSetting retrieves a value stored for a key
func GetSetting(c context.Context, key string) (string, error) {
	var setting *Setting
	k := datastore.NewKey(c, settingsKind, key, 0, nil)
	if err := datastore.Get(c, k, setting); err != nil {
		return "", err
	}
	return setting.Value, nil
}
