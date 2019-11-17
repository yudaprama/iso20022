package model

// Geographic location of the ATM specified by geographic coordinates or UTM coordinates.
type GeographicLocation1Choice struct {

	// Location on the earth specified by two numbers representing vertical and horizontal position.
	GeographicCoordinates *GeographicCoordinates1 `xml:"GeogcCordints"`

	// Location on the earth specified by the Universal Transverse Mercator coordinate system, using the WGS84 geodesic system.
	UTMCoordinates *UTMCoordinates1 `xml:"UTMCordints"`
}

func (g *GeographicLocation1Choice) AddGeographicCoordinates() *GeographicCoordinates1 {
	g.GeographicCoordinates = new(GeographicCoordinates1)
	return g.GeographicCoordinates
}

func (g *GeographicLocation1Choice) AddUTMCoordinates() *UTMCoordinates1 {
	g.UTMCoordinates = new(UTMCoordinates1)
	return g.UTMCoordinates
}
