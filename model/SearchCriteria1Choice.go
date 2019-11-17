package model

// Choice of search criteria for the financial investigation.
type SearchCriteria1Choice struct {

	// Identifies the account as the search criteria for the financial institution to do the investigation.
	Account *AccountAndParties1 `xml:"Acct"`

	// Identifies a customer identification as the search criteria for the financial institution to do the investigation.
	CustomerIdentification *CustomerIdentification1 `xml:"CstmrId"`

	// Identifies a payment instrument as the search criteria for the financial institution to do the investigation.
	PaymentInstrument *PaymentInstrumentType1 `xml:"PmtInstrm"`

	// Specifies the original transaction number.
	OriginalTransactionNumber []*RequestType1 `xml:"OrgnlTxNb"`
}

func (s *SearchCriteria1Choice) AddAccount() *AccountAndParties1 {
	s.Account = new(AccountAndParties1)
	return s.Account
}

func (s *SearchCriteria1Choice) AddCustomerIdentification() *CustomerIdentification1 {
	s.CustomerIdentification = new(CustomerIdentification1)
	return s.CustomerIdentification
}

func (s *SearchCriteria1Choice) AddPaymentInstrument() *PaymentInstrumentType1 {
	s.PaymentInstrument = new(PaymentInstrumentType1)
	return s.PaymentInstrument
}

func (s *SearchCriteria1Choice) AddOriginalTransactionNumber() *RequestType1 {
	newValue := new(RequestType1)
	s.OriginalTransactionNumber = append(s.OriginalTransactionNumber, newValue)
	return newValue
}
