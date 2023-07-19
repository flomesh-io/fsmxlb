// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FirewallEntry firewall entry
//
// swagger:model FirewallEntry
type FirewallEntry struct {

	// opts
	Opts *FirewallOptionEntry `json:"opts,omitempty"`

	// rule arguments
	RuleArguments *FirewallRuleEntry `json:"ruleArguments,omitempty"`
}

// Validate validates this firewall entry
func (m *FirewallEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirewallEntry) validateOpts(formats strfmt.Registry) error {
	if swag.IsZero(m.Opts) { // not required
		return nil
	}

	if m.Opts != nil {
		if err := m.Opts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opts")
			}
			return err
		}
	}

	return nil
}

func (m *FirewallEntry) validateRuleArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleArguments) { // not required
		return nil
	}

	if m.RuleArguments != nil {
		if err := m.RuleArguments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ruleArguments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ruleArguments")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this firewall entry based on the context it is used
func (m *FirewallEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOpts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuleArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FirewallEntry) contextValidateOpts(ctx context.Context, formats strfmt.Registry) error {

	if m.Opts != nil {

		if swag.IsZero(m.Opts) { // not required
			return nil
		}

		if err := m.Opts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opts")
			}
			return err
		}
	}

	return nil
}

func (m *FirewallEntry) contextValidateRuleArguments(ctx context.Context, formats strfmt.Registry) error {

	if m.RuleArguments != nil {

		if swag.IsZero(m.RuleArguments) { // not required
			return nil
		}

		if err := m.RuleArguments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ruleArguments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ruleArguments")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FirewallEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FirewallEntry) UnmarshalBinary(b []byte) error {
	var res FirewallEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}