package model

// Information related to a vehicle used during a transaction.
type Vehicle1 struct {

	// Number assigned to the vehicle for identification.
	VehicleNumber *Max35NumericText `xml:"VhclNb,omitempty"`

	// Number assigned to the vehicle trailer for identification.
	TrailerNumber *Max35NumericText `xml:"TrlrNb,omitempty"`

	// Registration tag of the vehicle.
	VehicleTag *Max35Text `xml:"VhclTag,omitempty"`

	// Entry mode of the registration tag.
	VehicleTagEntryMode *CardDataReading5Code `xml:"VhclTagNtryMd,omitempty"`

	// Identification of the vehicle in the fleet.
	UnitNumber *Max35NumericText `xml:"UnitNb,omitempty"`

	// True if the car is a replacement car.
	ReplacementCar *TrueFalseIndicator `xml:"RplcmntCar,omitempty"`

	// Odometer reading value indicating the distance travelled by the vehicle.
	Odometer *DecimalNumber `xml:"Odmtr,omitempty"`

	// Hubometer reading value indicating the distance travelled by the trailer.
	Hubometer *DecimalNumber `xml:"Hbmtr,omitempty"`

	// Number of hours the trailer has been in operation.
	TrailerHours *Max35Text `xml:"TrlrHrs,omitempty"`

	// Number of hours the refer unit has been in operation.
	ReferHours *Max35Text `xml:"RefrHrs,omitempty"`

	// Identification assigned to the vehicle related to maintenance.
	MaintenanceIdentification *Max35Text `xml:"MntncId,omitempty"`

	// Second card presented for the payment transaction.
	DriverOrVehicleCard *PlainCardData17 `xml:"DrvrOrVhclCard,omitempty"`

	// Additional information related to the vehicle.
	AdditionalVehicleData []*Vehicle2 `xml:"AddtlVhclData,omitempty"`
}

func (v *Vehicle1) SetVehicleNumber(value string) {
	v.VehicleNumber = (*Max35NumericText)(&value)
}

func (v *Vehicle1) SetTrailerNumber(value string) {
	v.TrailerNumber = (*Max35NumericText)(&value)
}

func (v *Vehicle1) SetVehicleTag(value string) {
	v.VehicleTag = (*Max35Text)(&value)
}

func (v *Vehicle1) SetVehicleTagEntryMode(value string) {
	v.VehicleTagEntryMode = (*CardDataReading5Code)(&value)
}

func (v *Vehicle1) SetUnitNumber(value string) {
	v.UnitNumber = (*Max35NumericText)(&value)
}

func (v *Vehicle1) SetReplacementCar(value string) {
	v.ReplacementCar = (*TrueFalseIndicator)(&value)
}

func (v *Vehicle1) SetOdometer(value string) {
	v.Odometer = (*DecimalNumber)(&value)
}

func (v *Vehicle1) SetHubometer(value string) {
	v.Hubometer = (*DecimalNumber)(&value)
}

func (v *Vehicle1) SetTrailerHours(value string) {
	v.TrailerHours = (*Max35Text)(&value)
}

func (v *Vehicle1) SetReferHours(value string) {
	v.ReferHours = (*Max35Text)(&value)
}

func (v *Vehicle1) SetMaintenanceIdentification(value string) {
	v.MaintenanceIdentification = (*Max35Text)(&value)
}

func (v *Vehicle1) AddDriverOrVehicleCard() *PlainCardData17 {
	v.DriverOrVehicleCard = new(PlainCardData17)
	return v.DriverOrVehicleCard
}

func (v *Vehicle1) AddAdditionalVehicleData() *Vehicle2 {
	newValue := new(Vehicle2)
	v.AdditionalVehicleData = append(v.AdditionalVehicleData, newValue)
	return newValue
}
