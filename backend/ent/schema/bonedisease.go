package schema

import "github.com/facebookincubator/ent"

// Bonedisease holds the schema definition for the Bonedisease entity.
type Bonedisease struct {
	ent.Schema
}

// Fields of the Bonedisease.
func (Bonedisease) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserEmail").NotEmpty(),
		field.String("UserName").NotEmpty(),
	}
}

// Edges of the Bonedisease.
func (Bonedisease) Edges() []ent.Edge {
	return nil
}
