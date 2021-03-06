// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/technician"
)

// TechnicianDelete is the builder for deleting a Technician entity.
type TechnicianDelete struct {
	config
	predicates []predicate.Technician
}

// Where adds a new predicate to the delete builder.
func (td *TechnicianDelete) Where(ps ...predicate.Technician) *TechnicianDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TechnicianDelete) Exec(ctx context.Context) (int, error) {
	return td.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TechnicianDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TechnicianDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: technician.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: technician.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, spec)
}

// TechnicianDeleteOne is the builder for deleting a single Technician entity.
type TechnicianDeleteOne struct {
	td *TechnicianDelete
}

// Exec executes the deletion query.
func (tdo *TechnicianDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{technician.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TechnicianDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
