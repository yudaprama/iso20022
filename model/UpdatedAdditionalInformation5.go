package model

// Additional information with update description and date.
type UpdatedAdditionalInformation5 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *RestrictedFINXMax140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation []*RestrictedFINXMax350Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation5) SetUpdateDescription(value string) {
	u.UpdateDescription = (*RestrictedFINXMax140Text)(&value)
}

func (u *UpdatedAdditionalInformation5) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation5) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*RestrictedFINXMax350Text)(&value))
}
