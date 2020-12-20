package internal

import "github.com/kingledion/ent-demo/internal/ent"

// The types in this package should really belong to an imported client from sezzle-pay or whoever owns these types
// The conversion functions will belong somewhere as well; maybe in a service/typeConversion.go

// User is exportable representation of user information
type User struct {
	UUID      string `json:"uuid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Merchant is exportable representation of merchant information
type Merchant struct {
	UUID string `json:"uuid"`
	DBA  string `json:"dba"`
}

// MerchantFromEnt converts an `ent` merchant to a publically sharable Merchant
func MerchantFromEnt(merch ent.Merchant) Merchant {
	return Merchant{
		UUID: merch.UUID,
		DBA:  merch.Dba,
	}
}

// Order is exportable representation of order information
type Order struct {
	Amount   uint     `json:"amount,omitempty"`
	User     User     `json:"user"`
	Merchant Merchant `json:"merchant"`
}
