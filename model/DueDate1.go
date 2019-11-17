package model

// Specifies the date by when the financial institutiion needs to provide a response.
type DueDate1 struct {

	// Specifies the date when the authority needs the response in situations where the response or part of it will not be given electronically but on paper in manual process.
	DueDate *ISODate `xml:"DueDt,omitempty"`

	// Specifies the reason why the authority needs the information on due date.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`
}

func (d *DueDate1) SetDueDate(value string) {
	d.DueDate = (*ISODate)(&value)
}

func (d *DueDate1) SetAdditionalInformation(value string) {
	d.AdditionalInformation = (*Max140Text)(&value)
}
