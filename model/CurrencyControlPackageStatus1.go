package model

// Provides the details of each package of currency control records.
type CurrencyControlPackageStatus1 struct {

	// Unique and unambiguous identification of each package of transactions and optionally the entry/record within the package of transactions.
	PackageIdentification *Max35Text `xml:"PackgId"`

	// Defines the status of the reported transaction.
	Status *StatisticalReportingStatus1Code `xml:"Sts"`

	// Provides detailed information on the status reason.
	StatusReason []*ValidationStatusReason1 `xml:"StsRsn,omitempty"`

	// Provides the date and time when the status was issued.
	StatusDateTime *ISODateTime `xml:"StsDtTm,omitempty"`

	// Provides the status of the individual records in the package.
	RecordStatus []*CurrencyControlRecordStatus1 `xml:"RcrdSts,omitempty"`
}

func (c *CurrencyControlPackageStatus1) SetPackageIdentification(value string) {
	c.PackageIdentification = (*Max35Text)(&value)
}

func (c *CurrencyControlPackageStatus1) SetStatus(value string) {
	c.Status = (*StatisticalReportingStatus1Code)(&value)
}

func (c *CurrencyControlPackageStatus1) AddStatusReason() *ValidationStatusReason1 {
	newValue := new(ValidationStatusReason1)
	c.StatusReason = append(c.StatusReason, newValue)
	return newValue
}

func (c *CurrencyControlPackageStatus1) SetStatusDateTime(value string) {
	c.StatusDateTime = (*ISODateTime)(&value)
}

func (c *CurrencyControlPackageStatus1) AddRecordStatus() *CurrencyControlRecordStatus1 {
	newValue := new(CurrencyControlRecordStatus1)
	c.RecordStatus = append(c.RecordStatus, newValue)
	return newValue
}
