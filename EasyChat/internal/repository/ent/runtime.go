// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Hosi121/easychat/internal/repository/ent/chat"
	"github.com/Hosi121/easychat/internal/repository/ent/group"
	"github.com/Hosi121/easychat/internal/repository/ent/message"
	"github.com/Hosi121/easychat/internal/repository/ent/schema"
	"github.com/Hosi121/easychat/internal/repository/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := schema.Chat{}.Fields()
	_ = chatFields
	// chatDescCreatedAt is the schema descriptor for created_at field.
	chatDescCreatedAt := chatFields[0].Descriptor()
	// chat.DefaultCreatedAt holds the default value on creation for the created_at field.
	chat.DefaultCreatedAt = chatDescCreatedAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupFields[1].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescContent is the schema descriptor for content field.
	messageDescContent := messageFields[0].Descriptor()
	// message.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	message.ContentValidator = messageDescContent.Validators[0].(func(string) error)
	// messageDescSentAt is the schema descriptor for sent_at field.
	messageDescSentAt := messageFields[1].Descriptor()
	// message.DefaultSentAt holds the default value on creation for the sent_at field.
	message.DefaultSentAt = messageDescSentAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[2].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = userDescPasswordHash.Validators[0].(func(string) error)
}
