package model

// Describes a transaction and its status.
type StatusReportItems2 struct {

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Entity for which the matching application has generated a report.
	ReportedEntity []*BICIdentification1 `xml:"RptdNtty"`

	// Identifies the status of the transaction by means of a code.
	Status *BaselineStatus3Code `xml:"Sts"`

	// Further description of the transaction status.
	SubStatus *Max140Text `xml:"SubSts,omitempty"`
}

func (s *StatusReportItems2) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *StatusReportItems2) AddReportedEntity() *BICIdentification1 {
	newValue := new(BICIdentification1)
	s.ReportedEntity = append(s.ReportedEntity, newValue)
	return newValue
}

func (s *StatusReportItems2) SetStatus(value string) {
	s.Status = (*BaselineStatus3Code)(&value)
}

func (s *StatusReportItems2) SetSubStatus(value string) {
	s.SubStatus = (*Max140Text)(&value)
}
