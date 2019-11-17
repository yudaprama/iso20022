package model

// Provides the details of each individual currency control record.
type CurrencyControlRecordStatus1 struct {

	// Unique and unambiguous identification of each entry/record within the package of transactions.
	RecordIdentification *Max35Text `xml:"RcrdId"`

	// Defines the status of the reported record.
	Status *StatisticalReportingStatus1Code `xml:"Sts"`

	// Provides detailed information on the status reason.
	StatusReason []*ValidationStatusReason1 `xml:"StsRsn,omitempty"`

	// Provides the date and time when the status was issued.
	StatusDateTime *ISODateTime `xml:"StsDtTm,omitempty"`

	// Unique and unambiguous identification of the document.
	DocumentIdentification *DocumentIdentification28 `xml:"DocId,omitempty"`
}

func (c *CurrencyControlRecordStatus1) SetRecordIdentification(value string) {
	c.RecordIdentification = (*Max35Text)(&value)
}

func (c *CurrencyControlRecordStatus1) SetStatus(value string) {
	c.Status = (*StatisticalReportingStatus1Code)(&value)
}

func (c *CurrencyControlRecordStatus1) AddStatusReason() *ValidationStatusReason1 {
	newValue := new(ValidationStatusReason1)
	c.StatusReason = append(c.StatusReason, newValue)
	return newValue
}

func (c *CurrencyControlRecordStatus1) SetStatusDateTime(value string) {
	c.StatusDateTime = (*ISODateTime)(&value)
}

func (c *CurrencyControlRecordStatus1) AddDocumentIdentification() *DocumentIdentification28 {
	c.DocumentIdentification = new(DocumentIdentification28)
	return c.DocumentIdentification
}
