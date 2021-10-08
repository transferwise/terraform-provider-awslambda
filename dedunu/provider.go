package dedunu

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	awsbase "github.com/hashicorp/aws-sdk-go-base"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type providerConfig struct {
	session *session.Session
}

// Provider - AWS Lambda provider
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: descriptions["profile"],
			},

			"region": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"AWS_REGION",
					"AWS_DEFAULT_REGION",
				}, nil),
				Description:  descriptions["region"],
				InputDefault: "us-east-1", // lintignore:AWSAT003
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"awslambda_invocation": resourceAwsLambdaInvocation(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		return providerConfigure(d, provider.TerraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	sess, err := awsbase.GetSession(&awsbase.Config{
		Profile: d.Get("profile").(string),
		Region:  d.Get("region").(string),
		UserAgentProducts: []*awsbase.UserAgentProduct{
			{Name: "APN", Version: "1.0"},
			{Name: "HashiCorp", Version: "1.0"},
			{Name: "Terraform", Version: terraformVersion, Extra: []string{"+https://www.terraform.io"}},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error configuring Terraform AWS Provider: %w", err)
	}

	config := &providerConfig{
		session: sess,
	}
	return config, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"region": "The region where AWS operations will take place. Examples\n" +
			"are us-east-1, us-west-2, etc.", // lintignore:AWSAT003

		"profile": "The profile for API operations. If not set, the default profile\n" +
			"created with `aws configure` will be used.",
	}
}
