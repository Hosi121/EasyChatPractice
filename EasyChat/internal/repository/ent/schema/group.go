// internal/repository/ent/schema/group.go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Group holds the schema definition for the Group entity.
type Group struct {
    ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("users", User.Type).Ref("groups"),
        edge.To("messages", Message.Type),
    }
}

