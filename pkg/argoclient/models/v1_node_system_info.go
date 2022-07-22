// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NodeSystemInfo NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
//
// swagger:model v1NodeSystemInfo
type V1NodeSystemInfo struct {

	// The Architecture reported by the node
	Architecture string `json:"architecture,omitempty"`

	// Boot ID reported by the node.
	BootID string `json:"bootID,omitempty"`

	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `json:"containerRuntimeVersion,omitempty"`

	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `json:"kernelVersion,omitempty"`

	// KubeProxy Version reported by the node.
	KubeProxyVersion string `json:"kubeProxyVersion,omitempty"`

	// Kubelet Version reported by the node.
	KubeletVersion string `json:"kubeletVersion,omitempty"`

	// MachineID reported by the node. For unique machine identification
	// in the cluster this field is preferred. Learn more from man(5)
	// machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	MachineID string `json:"machineID,omitempty"`

	// The Operating System reported by the node
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage string `json:"osImage,omitempty"`

	// SystemUUID reported by the node. For unique machine identification
	// MachineID is preferred. This field is specific to Red Hat hosts
	// https://access.redhat.com/documentation/en-us/red_hat_subscription_management/1/html/rhsm/uuid
	SystemUUID string `json:"systemUUID,omitempty"`
}

// Validate validates this v1 node system info
func (m *V1NodeSystemInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 node system info based on context it is used
func (m *V1NodeSystemInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1NodeSystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NodeSystemInfo) UnmarshalBinary(b []byte) error {
	var res V1NodeSystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
