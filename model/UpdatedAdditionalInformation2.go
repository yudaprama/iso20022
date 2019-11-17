package model

// Additional information with update description and date.
type UpdatedAdditionalInformation2 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*Max8000Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation2) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedAdditionalInformation2) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max8000Text)(&value))
}
