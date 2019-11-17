package model

// Provides detailed information on the number of reported items with their respective acceptance status.
type NumberOfItemsPerStatus1 struct {

	// Common status of the report items for which the number of report items is specified in NumberOfItems.
	Status *ReportItemStatus1Code `xml:"Sts"`

	// Number of items for the status.
	NumberOfItems *Max15NumericText `xml:"NbOfItms"`
}

func (n *NumberOfItemsPerStatus1) SetStatus(value string) {
	n.Status = (*ReportItemStatus1Code)(&value)
}

func (n *NumberOfItemsPerStatus1) SetNumberOfItems(value string) {
	n.NumberOfItems = (*Max15NumericText)(&value)
}
