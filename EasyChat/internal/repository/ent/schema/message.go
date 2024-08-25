package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
    "time"
)

// Message holds the schema definition for the Message entity.
type Message struct {
    ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
    return []ent.Field{
        field.Text("content").NotEmpty(),
        field.Time("sent_at").Default(time.Now),
    }
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("chat", Chat.Type).Ref("messages").Unique(),
        edge.From("user", User.Type).Ref("messages").Unique(),
    }
}

