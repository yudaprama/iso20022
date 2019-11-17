package model

// Additional information with update description and date.
type UpdatedAdditionalInformation3 struct {

	// Specifies the amendments made to the narrative since the last message.
	UpdateDescription *Max140Text `xml:"UpdDesc,omitempty"`

	// Specifies the date at which the narrative has been updated.
	UpdateDate *ISODate `xml:"UpdDt,omitempty"`

	// Provides additional textual information.
	AdditionalInformation *Max350Text `xml:"AddtlInf"`
}

func (u *UpdatedAdditionalInformation3) SetUpdateDescription(value string) {
	u.UpdateDescription = (*Max140Text)(&value)
}

func (u *UpdatedAdditionalInformation3) SetUpdateDate(value string) {
	u.UpdateDate = (*ISODate)(&value)
}

func (u *UpdatedAdditionalInformation3) SetAdditionalInformation(value string) {
	u.AdditionalInformation = (*Max350Text)(&value)
}
