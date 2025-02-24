package mcir

import (
	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
)

type TbImageInfo struct {
	AssociatedObjectList []string `json:"associatedObjectList"`

	ConnectionName  string `json:"connectionName"`
	CreationDate    string `json:"creationDate"`
	CspImageId      string `json:"cspImageId"`
	CspImageName    string `json:"cspImageName"`
	Description     string `json:"description"`
	GuestOS         string `json:"guestOS"`
	ID              string `json:"id"`
	IsAutoGenerated string `json:"isAutoGenerated"`

	KeyValueList []tbcommon.TbKeyValue `json:"keyValueList"`
	Name         string                `json:"name"`
	Namespace    string                `json:"namespace"` //required
	Status       string                `json:"status"`
}

type TbImageInfos []TbImageInfo
