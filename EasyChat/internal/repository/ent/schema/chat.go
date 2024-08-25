package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
    ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
    return []ent.Field{
        field.Time("created_at").Default(time.Now),
    }
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("users", User.Type).Ref("chats").Unique(),
        edge.To("messages", Message.Type),
    }
}

