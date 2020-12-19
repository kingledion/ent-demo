package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Merchant holds the schema definition for the Merchant entity.
type Merchant struct {
	ent.Schema
}

// Fields of the Merchant.
func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.String("dba"),
	}
}

// Edges of the Merchant.
func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("m_order", User.Type).
			Ref("order"),
	}
}
