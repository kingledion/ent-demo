// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"

	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"

	// Table holds the table name of the user in the database.
	Table = "users"
	// OrderTable is the table the holds the order relation/edge. The primary key declared below.
	OrderTable = "user_order"
	// OrderInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	OrderInverseTable = "merchants"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldFirstname,
	FieldLastname,
}

var (
	// OrderPrimaryKey and OrderColumn2 are the table columns denoting the
	// primary key for the order relation (M2M).
	OrderPrimaryKey = []string{"user_id", "merchant_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
