// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenstackMetadata openstack metadata
// swagger:model OpenstackMetadata
type OpenstackMetadata struct {

	// availability zones
	AvailabilityZones []AvailabilityZone `json:"availabilityZones"`

	// flavors
	Flavors []Flavor `json:"flavors"`

	// key pairs
	KeyPairs []*KeyPair `json:"keyPairs"`

	// routers
	Routers []*Router `json:"routers"`

	// security groups
	SecurityGroups []*SecurityGroup `json:"securityGroups"`
}

// Validate validates this openstack metadata
func (m *OpenstackMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZones(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlavors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKeyPairs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRouters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecurityGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenstackMetadata) validateAvailabilityZones(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailabilityZones) { // not required
		return nil
	}

	return nil
}

func (m *OpenstackMetadata) validateFlavors(formats strfmt.Registry) error {

	if swag.IsZero(m.Flavors) { // not required
		return nil
	}

	return nil
}

func (m *OpenstackMetadata) validateKeyPairs(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyPairs) { // not required
		return nil
	}

	for i := 0; i < len(m.KeyPairs); i++ {

		if swag.IsZero(m.KeyPairs[i]) { // not required
			continue
		}

		if m.KeyPairs[i] != nil {

			if err := m.KeyPairs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyPairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenstackMetadata) validateRouters(formats strfmt.Registry) error {

	if swag.IsZero(m.Routers) { // not required
		return nil
	}

	for i := 0; i < len(m.Routers); i++ {

		if swag.IsZero(m.Routers[i]) { // not required
			continue
		}

		if m.Routers[i] != nil {

			if err := m.Routers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OpenstackMetadata) validateSecurityGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroups); i++ {

		if swag.IsZero(m.SecurityGroups[i]) { // not required
			continue
		}

		if m.SecurityGroups[i] != nil {

			if err := m.SecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackMetadata) UnmarshalBinary(b []byte) error {
	var res OpenstackMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AvailabilityZone availability zone
// swagger:model AvailabilityZone
type AvailabilityZone struct {

	// available
	Available bool `json:"available,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this availability zone
func (m *AvailabilityZone) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AvailabilityZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailabilityZone) UnmarshalBinary(b []byte) error {
	var res AvailabilityZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Flavor flavor
// swagger:model Flavor
type Flavor struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ram
	RAM int64 `json:"ram"`

	// vcpus
	Vcpus int64 `json:"vcpus"`
}

// Validate validates this flavor
func (m *Flavor) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Flavor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Flavor) UnmarshalBinary(b []byte) error {
	var res Flavor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KeyPair key pair
// swagger:model KeyPair
type KeyPair struct {

	// name
	Name string `json:"name,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`
}

// Validate validates this key pair
func (m *KeyPair) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *KeyPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyPair) UnmarshalBinary(b []byte) error {
	var res KeyPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Router router
// swagger:model Router
type Router struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// networks
	Networks []*Network `json:"networks"`
}

// Validate validates this router
func (m *Router) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Router) validateNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {

		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {

			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Router) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Router) UnmarshalBinary(b []byte) error {
	var res Router
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Network network
// swagger:model Network
type Network struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// subnets
	Subnets []*Subnet `json:"subnets"`
}

// Validate validates this network
func (m *Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubnets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) validateSubnets(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(m.Subnets); i++ {

		if swag.IsZero(m.Subnets[i]) { // not required
			continue
		}

		if m.Subnets[i] != nil {

			if err := m.Subnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Network) UnmarshalBinary(b []byte) error {
	var res Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Subnet subnet
// swagger:model Subnet
type Subnet struct {

	// c ID r
	CIDR string `json:"CIDR,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this subnet
func (m *Subnet) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Subnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subnet) UnmarshalBinary(b []byte) error {
	var res Subnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityGroup security group
// swagger:model SecurityGroup
type SecurityGroup struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this security group
func (m *SecurityGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroup) UnmarshalBinary(b []byte) error {
	var res SecurityGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
