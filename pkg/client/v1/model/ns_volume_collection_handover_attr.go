/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsVolumeCollectionHandoverAttr


// NsVolumeCollectionHandoverAttr :
type NsVolumeCollectionHandoverAttr struct {
   // ID
   ID string `json:"id,omitempty"`
   // ReplicationPartnerID
   ReplicationPartnerID string `json:"replication_partner_id,omitempty"`
   // NoReverse
   NoReverse bool `json:"no_reverse,omitempty"`
}
