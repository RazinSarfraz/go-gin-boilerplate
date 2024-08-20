package context

import (
	"context"
	"time"
	"user-backend/models"

	"gorm.io/gorm"
)

type Context struct {
	baseCtx context.Context
	tx      *gorm.DB
	user    *models.User
}

// Read-only accessors
func (c Context) Context() context.Context { return c.baseCtx }
func (c Context) Tx() *gorm.DB             { return c.tx }

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.baseCtx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.baseCtx.Done()
}

func (c Context) Err() error {
	return c.baseCtx.Err()
}

// create a new context
func NewContext(ctx context.Context) Context {

	return Context{
		baseCtx: ctx,
	}
}

// WithContext returns a Context with an updated context.Context.
func (c Context) WithContext(ctx context.Context) Context {
	c.baseCtx = ctx
	return c
}

// WithTx returns a Context with an updated tx.
func (c Context) WithTx(tx *gorm.DB) Context {
	c.tx = tx
	return c
}

// WithContext returns a Context with an updated user.
func (c Context) WithUser(user *models.User) Context {
	c.user = user
	return c
}

// Return user from ctx
func (c Context) User() *models.User {
	return c.user
}
