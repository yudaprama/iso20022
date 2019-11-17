package model

// Additional information with update description and date.
type UpdatedAdditionalInformation10 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *RestrictedFINXMax140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*RestrictedFINZMax8000Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation10) SetUpdateDescription(value string) {
	u.UpdateDescription = (*RestrictedFINXMax140Text)(&value)
}

func (u *UpdatedAdditionalInformation10) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation10) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*RestrictedFINZMax8000Text)(&value))
}
