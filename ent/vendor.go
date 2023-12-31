// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"pos/ent/vendor"
	"pos/internal/domain"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Vendor is the model entity for the Vendor schema.
type Vendor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Representative holds the value of the "representative" field.
	Representative domain.Representative `json:"representative,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vendor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vendor.FieldRepresentative:
			values[i] = new([]byte)
		case vendor.FieldID:
			values[i] = new(sql.NullInt64)
		case vendor.FieldName, vendor.FieldAddress, vendor.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vendor fields.
func (v *Vendor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vendor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vendor.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case vendor.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				v.Address = value.String
			}
		case vendor.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				v.Email = value.String
			}
		case vendor.FieldRepresentative:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field representative", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &v.Representative); err != nil {
					return fmt.Errorf("unmarshal field representative: %w", err)
				}
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vendor.
// This includes values selected through modifiers, order, etc.
func (v *Vendor) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// Update returns a builder for updating this Vendor.
// Note that you need to call Vendor.Unwrap() before calling this method if this Vendor
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vendor) Update() *VendorUpdateOne {
	return NewVendorClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vendor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vendor) Unwrap() *Vendor {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vendor is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vendor) String() string {
	var builder strings.Builder
	builder.WriteString("Vendor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("name=")
	builder.WriteString(v.Name)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(v.Address)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(v.Email)
	builder.WriteString(", ")
	builder.WriteString("representative=")
	builder.WriteString(fmt.Sprintf("%v", v.Representative))
	builder.WriteByte(')')
	return builder.String()
}

// Vendors is a parsable slice of Vendor.
type Vendors []*Vendor
