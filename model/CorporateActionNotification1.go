package model

// Provides information about the notification advice.
type CorporateActionNotification1 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat4Choice `xml:"AnncmntDt,omitempty"`

	// Date/time at which additional information on the event will be announced, eg, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat4Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, eg, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat4Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Specifies the status of the details of the event.
	ProcessingStatus *ProcessingStatus1FormatChoice `xml:"PrcgSts"`
}

func (c *CorporateActionNotification1) AddAnnouncementDate() *DateFormat4Choice {
	c.AnnouncementDate = new(DateFormat4Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionNotification1) AddFurtherDetailedAnnouncementDate() *DateFormat4Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat4Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionNotification1) AddOfficialAnnouncementPublicationDate() *DateFormat4Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat4Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionNotification1) AddProcessingStatus() *ProcessingStatus1FormatChoice {
	c.ProcessingStatus = new(ProcessingStatus1FormatChoice)
	return c.ProcessingStatus
}
