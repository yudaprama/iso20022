package model

// Specifies the security option of a corporate event.
type SecuritiesOption15 struct {

	// Maximum quantity (or lot) of financial instrument that must be exercised or tendered.
	MaximumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MaxExrcblQty,omitempty"`

	// Minimum quantity (or lot) of financial instrument that must be exercised or tendered.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity (or lot) of financial  instrument that must be exercised or tendered.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotQuantity *FinancialInstrumentQuantity1Choice `xml:"NewBrdLotQty,omitempty"`

	// New denomination of the equity following, for example, an increase or decrease in nominal value.
	NewDenominationQuantity *FinancialInstrumentQuantity1Choice `xml:"NewDnmtnQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotQuantity *FinancialInstrumentQuantity16Choice `xml:"FrntEndOddLotQty,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotQuantity *FinancialInstrumentQuantity16Choice `xml:"BckEndOddLotQty,omitempty"`
}

func (s *SecuritiesOption15) AddMaximumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	s.MaximumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MaximumExercisableQuantity
}

func (s *SecuritiesOption15) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	s.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MinimumExercisableQuantity
}

func (s *SecuritiesOption15) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	s.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.MinimumExercisableMultipleQuantity
}

func (s *SecuritiesOption15) AddNewBoardLotQuantity() *FinancialInstrumentQuantity1Choice {
	s.NewBoardLotQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.NewBoardLotQuantity
}

func (s *SecuritiesOption15) AddNewDenominationQuantity() *FinancialInstrumentQuantity1Choice {
	s.NewDenominationQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.NewDenominationQuantity
}

func (s *SecuritiesOption15) AddFrontEndOddLotQuantity() *FinancialInstrumentQuantity16Choice {
	s.FrontEndOddLotQuantity = new(FinancialInstrumentQuantity16Choice)
	return s.FrontEndOddLotQuantity
}

func (s *SecuritiesOption15) AddBackEndOddLotQuantity() *FinancialInstrumentQuantity16Choice {
	s.BackEndOddLotQuantity = new(FinancialInstrumentQuantity16Choice)
	return s.BackEndOddLotQuantity
}
