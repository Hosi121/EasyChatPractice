// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/easychat/internal/repository/ent/chat"
	"github.com/Hosi121/easychat/internal/repository/ent/group"
	"github.com/Hosi121/easychat/internal/repository/ent/message"
	"github.com/Hosi121/easychat/internal/repository/ent/predicate"
	"github.com/Hosi121/easychat/internal/repository/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPasswordHash sets the "password_hash" field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.mutation.SetPasswordHash(s)
	return uu
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePasswordHash(s *string) *UserUpdate {
	if s != nil {
		uu.SetPasswordHash(*s)
	}
	return uu
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (uu *UserUpdate) AddMessageIDs(ids ...int) *UserUpdate {
	uu.mutation.AddMessageIDs(ids...)
	return uu
}

// AddMessages adds the "messages" edges to the Message entity.
func (uu *UserUpdate) AddMessages(m ...*Message) *UserUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddMessageIDs(ids...)
}

// AddChatIDs adds the "chats" edge to the Chat entity by IDs.
func (uu *UserUpdate) AddChatIDs(ids ...int) *UserUpdate {
	uu.mutation.AddChatIDs(ids...)
	return uu
}

// AddChats adds the "chats" edges to the Chat entity.
func (uu *UserUpdate) AddChats(c ...*Chat) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddChatIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearMessages clears all "messages" edges to the Message entity.
func (uu *UserUpdate) ClearMessages() *UserUpdate {
	uu.mutation.ClearMessages()
	return uu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (uu *UserUpdate) RemoveMessageIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveMessageIDs(ids...)
	return uu
}

// RemoveMessages removes "messages" edges to Message entities.
func (uu *UserUpdate) RemoveMessages(m ...*Message) *UserUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveMessageIDs(ids...)
}

// ClearChats clears all "chats" edges to the Chat entity.
func (uu *UserUpdate) ClearChats() *UserUpdate {
	uu.mutation.ClearChats()
	return uu
}

// RemoveChatIDs removes the "chats" edge to Chat entities by IDs.
func (uu *UserUpdate) RemoveChatIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveChatIDs(ids...)
	return uu
}

// RemoveChats removes "chats" edges to Chat entities.
func (uu *UserUpdate) RemoveChats(c ...*Chat) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveChatIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if uu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !uu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedChatsIDs(); len(nodes) > 0 && !uu.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPasswordHash sets the "password_hash" field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(s)
	return uuo
}

// SetNillablePasswordHash sets the "password_hash" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePasswordHash(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPasswordHash(*s)
	}
	return uuo
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (uuo *UserUpdateOne) AddMessageIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddMessageIDs(ids...)
	return uuo
}

// AddMessages adds the "messages" edges to the Message entity.
func (uuo *UserUpdateOne) AddMessages(m ...*Message) *UserUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddMessageIDs(ids...)
}

// AddChatIDs adds the "chats" edge to the Chat entity by IDs.
func (uuo *UserUpdateOne) AddChatIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddChatIDs(ids...)
	return uuo
}

// AddChats adds the "chats" edges to the Chat entity.
func (uuo *UserUpdateOne) AddChats(c ...*Chat) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddChatIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearMessages clears all "messages" edges to the Message entity.
func (uuo *UserUpdateOne) ClearMessages() *UserUpdateOne {
	uuo.mutation.ClearMessages()
	return uuo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (uuo *UserUpdateOne) RemoveMessageIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveMessageIDs(ids...)
	return uuo
}

// RemoveMessages removes "messages" edges to Message entities.
func (uuo *UserUpdateOne) RemoveMessages(m ...*Message) *UserUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveMessageIDs(ids...)
}

// ClearChats clears all "chats" edges to the Chat entity.
func (uuo *UserUpdateOne) ClearChats() *UserUpdateOne {
	uuo.mutation.ClearChats()
	return uuo
}

// RemoveChatIDs removes the "chats" edge to Chat entities by IDs.
func (uuo *UserUpdateOne) RemoveChatIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveChatIDs(ids...)
	return uuo
}

// RemoveChats removes "chats" edges to Chat entities.
func (uuo *UserUpdateOne) RemoveChats(c ...*Chat) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveChatIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if uuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !uuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedChatsIDs(); len(nodes) > 0 && !uuo.mutation.ChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChatsTable,
			Columns: []string{user.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
