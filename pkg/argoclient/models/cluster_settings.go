// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterSettings cluster settings
//
// swagger:model clusterSettings
type ClusterSettings struct {

	// app label key
	AppLabelKey string `json:"appLabelKey,omitempty"`

	// config management plugins
	ConfigManagementPlugins []*V1alpha1ConfigManagementPlugin `json:"configManagementPlugins"`

	// dex config
	DexConfig *ClusterDexConfig `json:"dexConfig,omitempty"`

	// exec enabled
	ExecEnabled bool `json:"execEnabled,omitempty"`

	// google analytics
	GoogleAnalytics *ClusterGoogleAnalyticsConfig `json:"googleAnalytics,omitempty"`

	// help
	Help *ClusterHelp `json:"help,omitempty"`

	// kustomize options
	KustomizeOptions *V1alpha1KustomizeOptions `json:"kustomizeOptions,omitempty"`

	// kustomize versions
	KustomizeVersions []string `json:"kustomizeVersions"`

	// oidc config
	OidcConfig *ClusterOIDCConfig `json:"oidcConfig,omitempty"`

	// password pattern
	PasswordPattern string `json:"passwordPattern,omitempty"`

	// plugins
	Plugins []*ClusterPlugin `json:"plugins"`

	// resource overrides
	ResourceOverrides map[string]V1alpha1ResourceOverride `json:"resourceOverrides,omitempty"`

	// status badge enabled
	StatusBadgeEnabled bool `json:"statusBadgeEnabled,omitempty"`

	// status badge root Url
	StatusBadgeRootURL string `json:"statusBadgeRootUrl,omitempty"`

	// tracking method
	TrackingMethod string `json:"trackingMethod,omitempty"`

	// ui banner content
	UIBannerContent string `json:"uiBannerContent,omitempty"`

	// ui banner permanent
	UIBannerPermanent bool `json:"uiBannerPermanent,omitempty"`

	// ui banner position
	UIBannerPosition string `json:"uiBannerPosition,omitempty"`

	// ui banner URL
	UIBannerURL string `json:"uiBannerURL,omitempty"`

	// ui Css URL
	UICSSURL string `json:"uiCssURL,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user logins disabled
	UserLoginsDisabled bool `json:"userLoginsDisabled,omitempty"`
}

// Validate validates this cluster settings
func (m *ClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigManagementPlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDexConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogleAnalytics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKustomizeOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidcConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettings) validateConfigManagementPlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigManagementPlugins) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigManagementPlugins); i++ {
		if swag.IsZero(m.ConfigManagementPlugins[i]) { // not required
			continue
		}

		if m.ConfigManagementPlugins[i] != nil {
			if err := m.ConfigManagementPlugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configManagementPlugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configManagementPlugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSettings) validateDexConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DexConfig) { // not required
		return nil
	}

	if m.DexConfig != nil {
		if err := m.DexConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dexConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dexConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) validateGoogleAnalytics(formats strfmt.Registry) error {
	if swag.IsZero(m.GoogleAnalytics) { // not required
		return nil
	}

	if m.GoogleAnalytics != nil {
		if err := m.GoogleAnalytics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("googleAnalytics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("googleAnalytics")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) validateHelp(formats strfmt.Registry) error {
	if swag.IsZero(m.Help) { // not required
		return nil
	}

	if m.Help != nil {
		if err := m.Help.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("help")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("help")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) validateKustomizeOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.KustomizeOptions) { // not required
		return nil
	}

	if m.KustomizeOptions != nil {
		if err := m.KustomizeOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kustomizeOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kustomizeOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) validateOidcConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.OidcConfig) { // not required
		return nil
	}

	if m.OidcConfig != nil {
		if err := m.OidcConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidcConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidcConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) validatePlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	for i := 0; i < len(m.Plugins); i++ {
		if swag.IsZero(m.Plugins[i]) { // not required
			continue
		}

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSettings) validateResourceOverrides(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceOverrides) { // not required
		return nil
	}

	for k := range m.ResourceOverrides {

		if err := validate.Required("resourceOverrides"+"."+k, "body", m.ResourceOverrides[k]); err != nil {
			return err
		}
		if val, ok := m.ResourceOverrides[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceOverrides" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceOverrides" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster settings based on the context it is used
func (m *ClusterSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigManagementPlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDexConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogleAnalytics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHelp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKustomizeOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOidcConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettings) contextValidateConfigManagementPlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConfigManagementPlugins); i++ {

		if m.ConfigManagementPlugins[i] != nil {
			if err := m.ConfigManagementPlugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configManagementPlugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configManagementPlugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSettings) contextValidateDexConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DexConfig != nil {
		if err := m.DexConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dexConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dexConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) contextValidateGoogleAnalytics(ctx context.Context, formats strfmt.Registry) error {

	if m.GoogleAnalytics != nil {
		if err := m.GoogleAnalytics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("googleAnalytics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("googleAnalytics")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) contextValidateHelp(ctx context.Context, formats strfmt.Registry) error {

	if m.Help != nil {
		if err := m.Help.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("help")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("help")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) contextValidateKustomizeOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.KustomizeOptions != nil {
		if err := m.KustomizeOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kustomizeOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kustomizeOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) contextValidateOidcConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.OidcConfig != nil {
		if err := m.OidcConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidcConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oidcConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettings) contextValidatePlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Plugins); i++ {

		if m.Plugins[i] != nil {
			if err := m.Plugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plugins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSettings) contextValidateResourceOverrides(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ResourceOverrides {

		if val, ok := m.ResourceOverrides[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettings) UnmarshalBinary(b []byte) error {
	var res ClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
