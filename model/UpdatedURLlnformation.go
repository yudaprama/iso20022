package model

// Additional information with update description and date.
type UpdatedURLlnformation struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides the web address, that is, the address for the Universal Resource Locator (URL), to use over the www (HTTP) service where additional information may be found.
	URLAddress *Max256Text `xml:"URLAdr"`
}

func (u *UpdatedURLlnformation) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedURLlnformation) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedURLlnformation) SetURLAddress(value string) {
	u.URLAddress = (*Max256Text)(&value)
}
