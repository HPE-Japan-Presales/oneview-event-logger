package oneview

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/rest"
)

type EventList struct {
	Type        string  `json:"type,omitempty"`
	URI         string  `json:"uri,omitempty"`
	Category    string  `json:"category,omitempty"`
	ETAG        string  `json:"eTag,omitempty"`
	Created     string  `json:"created,omitempty"`
	Modified    string  `json:"modified,omitempty"`
	Start       int     `json:"start,omitempty"`
	Count       int     `json:"count,omitempty"`
	Total       int     `json:"total,omitempty"`
	PrevPageURI string  `json:"prevPageUri,omitempty"`
	NextPageURI string  `json:"nextPageUri,omitempty"`
	Members     []Event `json:"members,omitempty"`
}

type Event struct {
	Type                string              `json:"type,omitempty"`
	URI                 string              `json:"uri,omitempty"`
	Vategory            string              `json:"category,omitempty"`
	ETAG                string              `json:"eventItemSnmpOid,omitempty"`
	Created             string              `json:"created,omitempty"`
	Modified            string              `json:"modified,omitempty"`
	HealthCategory      string              `json:"healthCategory,omitempty"`
	Description         string              `json:"description,omitempty"`
	EventTypeID         string              `json:"eventTypeID,omitempty"`
	EventDetails        []EventDetail       `json:"eventDetails,omitempty"`
	RxTime              string              `json:"rxTime,omitempty"`
	Processed           bool                `json:"processed,omitempty"`
	Severity            string              `json:"severity,omitempty"`
	Urgency             string              `json:"urgency,omitempty"`
	ServiceEventSource  bool                `json:"serviceEventSource,omitempty"`
	ServiceEventDetails ServiceEventDetails `json:"serviceEventDetails,omitempty"`
}

type EventDetail struct {
	EventItemName        string `json:"eventItemName"`
	EventItemValue       string `json:"eventItemValue,omitempty"`
	VarBindOrderIndex    int    `json:"varBindOrderIndex,omitempty"`
	IsThisVarbindData    bool   `json:"isThisVarbindData,omitempty"`
	EventItemDescription string `json:"eventItemDescription,omitempty"`
	EventItemSnmpOid     string `json:"eventItemSnmpOid,omitempty"`
}

type ServiceEventDetails struct {
	CaseId             string `json:"caseId,omitempty"`
	PrimaryContact     string `json:"primaryContact,omitempty"`
	RemoteSupportState string `json:"remoteSupportState,omitempty"`
}

func GetEventList(c *ov.OVClient) (EventList, error) {
	var (
		uri       = "/rest/events"
		eventlist EventList
	)

	// refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	data, err := c.RestAPICall(rest.GET, uri, nil)
	if err != nil {
		return eventlist, err
	}
	if err := json.Unmarshal([]byte(data), &eventlist); err != nil {
		return eventlist, err
	}

	return eventlist, nil
}
