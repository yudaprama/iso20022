package model

// Specifies the security option of a corporate event.
type SecuritiesOption23 struct {

	// Maximum quantity (or lot) of financial instrument that may be exercised or tendered.
	MaximumExercisableQuantity *FinancialInstrumentQuantity19Choice `xml:"MaxExrcblQty,omitempty"`

	// Minimum quantity (or lot) of financial instrument that may be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity19Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial  instrument that may be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity20Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity20Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity20Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity20Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption23) AddMaximumExercisableQuantity() *FinancialInstrumentQuantity19Choice {
	s.MaximumExercisableQuantity = new(FinancialInstrumentQuantity19Choice)
	return s.MaximumExercisableQuantity
}

func (s *SecuritiesOption23) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity19Choice {
	s.MinimumExercisableQuantity = new(FinancialInstrumentQuantity19Choice)
	return s.MinimumExercisableQuantity
}

func (s *SecuritiesOption23) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity20Choice {
	s.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.MinimumExercisableMultipleQuantity
}

func (s *SecuritiesOption23) AddNewBoardLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption23) AddNewDenominationQuantity() *FinancialInstrumentQuantity20Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption23) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption23) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity20Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity20Choice)
	return s.BackEndOddLotQuantity
}
