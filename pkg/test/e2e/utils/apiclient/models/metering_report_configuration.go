// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MeteringReportConfiguration MeteringReportConfiguration holds report configuration
//
// swagger:model MeteringReportConfiguration
type MeteringReportConfiguration struct {

	// Interval defines the number of days consulted in the metering report.
	Interval uint32 `json:"interval,omitempty"`

	// Retention defines a number of days after which reports are queued for removal. If not set, reports are kept forever.
	// Please note that this functionality works only for object storage that supports an object lifecycle management mechanism.
	Retention uint32 `json:"retention,omitempty"`

	// Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron. Please take a note that Schedule is responsible
	// only for setting the time when a report generation mechanism kicks off. The Interval MUST be set independently.
	Schedule string `json:"schedule,omitempty"`
}

// Validate validates this metering report configuration
func (m *MeteringReportConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metering report configuration based on context it is used
func (m *MeteringReportConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeteringReportConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteringReportConfiguration) UnmarshalBinary(b []byte) error {
	var res MeteringReportConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
