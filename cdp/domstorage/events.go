package domstorage

// AUTOGENERATED. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventDomStorageItemsCleared [no description].
type EventDomStorageItemsCleared struct {
	StorageID *StorageID `json:"storageId,omitempty"`
}

// EventDomStorageItemRemoved [no description].
type EventDomStorageItemRemoved struct {
	StorageID *StorageID `json:"storageId,omitempty"`
	Key       string     `json:"key,omitempty"`
}

// EventDomStorageItemAdded [no description].
type EventDomStorageItemAdded struct {
	StorageID *StorageID `json:"storageId,omitempty"`
	Key       string     `json:"key,omitempty"`
	NewValue  string     `json:"newValue,omitempty"`
}

// EventDomStorageItemUpdated [no description].
type EventDomStorageItemUpdated struct {
	StorageID *StorageID `json:"storageId,omitempty"`
	Key       string     `json:"key,omitempty"`
	OldValue  string     `json:"oldValue,omitempty"`
	NewValue  string     `json:"newValue,omitempty"`
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventDOMStorageDomStorageItemsCleared,
	cdp.EventDOMStorageDomStorageItemRemoved,
	cdp.EventDOMStorageDomStorageItemAdded,
	cdp.EventDOMStorageDomStorageItemUpdated,
}
