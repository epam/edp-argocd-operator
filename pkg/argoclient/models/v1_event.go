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

// V1Event Event is a report of an event somewhere in the cluster.  Events
// have a limited retention time and triggers and messages may evolve
// with time.  Event consumers should not rely on the timing of an event
// with a given Reason reflecting a consistent underlying trigger, or the
// continued existence of events with that Reason.  Events should be
// treated as informative, best-effort, supplemental data.
//
// swagger:model v1Event
type V1Event struct {

	// What action was taken/failed regarding to the Regarding object.
	// +optional
	Action string `json:"action,omitempty"`

	// The number of times this event has occurred.
	// +optional
	Count int32 `json:"count,omitempty"`

	// event time
	EventTime *V1MicroTime `json:"eventTime,omitempty"`

	// first timestamp
	// Format: date-time
	FirstTimestamp V1Time `json:"firstTimestamp,omitempty"`

	// involved object
	InvolvedObject *V1ObjectReference `json:"involvedObject,omitempty"`

	// last timestamp
	// Format: date-time
	LastTimestamp V1Time `json:"lastTimestamp,omitempty"`

	// A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	// +optional
	Message string `json:"message,omitempty"`

	// metadata
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`

	// This should be a short, machine understandable string that gives the reason
	// for the transition into the object's current status.
	// TODO: provide exact specification for format.
	// +optional
	Reason string `json:"reason,omitempty"`

	// related
	Related *V1ObjectReference `json:"related,omitempty"`

	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	// +optional
	ReportingComponent string `json:"reportingComponent,omitempty"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	// +optional
	ReportingInstance string `json:"reportingInstance,omitempty"`

	// series
	Series *V1EventSeries `json:"series,omitempty"`

	// source
	Source *V1EventSource `json:"source,omitempty"`

	// Type of this event (Normal, Warning), new types could be added in the future
	// +optional
	Type string `json:"type,omitempty"`
}

// Validate validates this v1 event
func (m *V1Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvolvedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Event) validateEventTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTime) { // not required
		return nil
	}

	if m.EventTime != nil {
		if err := m.EventTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventTime")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateFirstTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.FirstTimestamp) { // not required
		return nil
	}

	if err := m.FirstTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("firstTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("firstTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1Event) validateInvolvedObject(formats strfmt.Registry) error {
	if swag.IsZero(m.InvolvedObject) { // not required
		return nil
	}

	if m.InvolvedObject != nil {
		if err := m.InvolvedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("involvedObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("involvedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateLastTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTimestamp) { // not required
		return nil
	}

	if err := m.LastTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1Event) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateRelated(formats strfmt.Registry) error {
	if swag.IsZero(m.Related) { // not required
		return nil
	}

	if m.Related != nil {
		if err := m.Related.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("related")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("related")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateSeries(formats strfmt.Registry) error {
	if swag.IsZero(m.Series) { // not required
		return nil
	}

	if m.Series != nil {
		if err := m.Series.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("series")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("series")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 event based on the context it is used
func (m *V1Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvolvedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Event) contextValidateEventTime(ctx context.Context, formats strfmt.Registry) error {

	if m.EventTime != nil {
		if err := m.EventTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventTime")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) contextValidateFirstTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FirstTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("firstTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("firstTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1Event) contextValidateInvolvedObject(ctx context.Context, formats strfmt.Registry) error {

	if m.InvolvedObject != nil {
		if err := m.InvolvedObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("involvedObject")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("involvedObject")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) contextValidateLastTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastTimestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTimestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTimestamp")
		}
		return err
	}

	return nil
}

func (m *V1Event) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) contextValidateRelated(ctx context.Context, formats strfmt.Registry) error {

	if m.Related != nil {
		if err := m.Related.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("related")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("related")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) contextValidateSeries(ctx context.Context, formats strfmt.Registry) error {

	if m.Series != nil {
		if err := m.Series.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("series")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("series")
			}
			return err
		}
	}

	return nil
}

func (m *V1Event) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Event) UnmarshalBinary(b []byte) error {
	var res V1Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
