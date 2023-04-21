// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var IapTunnelInstanceIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"zone": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"instance": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type IapTunnelInstanceIamUpdater struct {
	project  string
	zone     string
	instance string
	d        TerraformResourceData
	Config   *transport_tpg.Config
}

func IapTunnelInstanceIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	zone, _ := getZone(d, config)
	if zone != "" {
		if err := d.Set("zone", zone); err != nil {
			return nil, fmt.Errorf("Error setting zone: %s", err)
		}
	}
	values["zone"] = zone
	if v, ok := d.GetOk("instance"); ok {
		values["instance"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_tunnel/zones/(?P<zone>[^/]+)/instances/(?P<instance>[^/]+)", "projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/instances/(?P<instance>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<instance>[^/]+)", "(?P<zone>[^/]+)/(?P<instance>[^/]+)", "(?P<instance>[^/]+)"}, d, config, d.Get("instance").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapTunnelInstanceIamUpdater{
		project:  values["project"],
		zone:     values["zone"],
		instance: values["instance"],
		d:        d,
		Config:   config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("zone", u.zone); err != nil {
		return nil, fmt.Errorf("Error setting zone: %s", err)
	}
	if err := d.Set("instance", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting instance: %s", err)
	}

	return u, nil
}

func IapTunnelInstanceIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	zone, _ := getZone(d, config)
	if zone != "" {
		values["zone"] = zone
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/iap_tunnel/zones/(?P<zone>[^/]+)/instances/(?P<instance>[^/]+)", "projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/instances/(?P<instance>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<instance>[^/]+)", "(?P<zone>[^/]+)/(?P<instance>[^/]+)", "(?P<instance>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &IapTunnelInstanceIamUpdater{
		project:  values["project"],
		zone:     values["zone"],
		instance: values["instance"],
		d:        d,
		Config:   config,
	}
	if err := d.Set("instance", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting instance: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *IapTunnelInstanceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyTunnelInstanceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}
	obj = map[string]interface{}{
		"options": map[string]interface{}{
			"requestedPolicyVersion": IamPolicyVersion,
		},
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *IapTunnelInstanceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyTunnelInstanceUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *IapTunnelInstanceIamUpdater) qualifyTunnelInstanceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{IapBasePath}}%s:%s", fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s", u.project, u.zone, u.instance), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *IapTunnelInstanceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s", u.project, u.zone, u.instance)
}

func (u *IapTunnelInstanceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-iap-tunnelinstance-%s", u.GetResourceId())
}

func (u *IapTunnelInstanceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("iap tunnelinstance %q", u.GetResourceId())
}
