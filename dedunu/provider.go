package dedunu

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	awsbase "github.com/hashicorp/aws-sdk-go-base"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
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
				Type:     schema.TypeString,
				Required: true,
				Description: "The profile for API operations. If not set, the default profile\n" +
					"created with `aws configure` will be used.",
				InputDefault: "",
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The region where AWS operations will take place. Examples\n" +
					"are us-east-1, us-west-2, etc.",
				InputDefault: "us-west-1",
			},
			"role_arn": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Amazon Resource Name of an IAM Role to assume prior to making\n" +
					"the AWS Lambda call.",
				InputDefault: "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"awslambda_invocation": resourceAwsLambdaInvocation(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion

		awsbaseConfig := &awsbase.Config{
			AccessKey:                   "",
			AssumeRoleARN:               d.Get("role_arn").(string),
			AssumeRoleDurationSeconds:   0,
			AssumeRoleExternalID:        "",
			AssumeRolePolicy:            "",
			AssumeRolePolicyARNs:        []string{},
			AssumeRoleSessionName:       "",
			AssumeRoleTags:              map[string]string{},
			AssumeRoleTransitiveTagKeys: []string{},
			CallerDocumentationURL:      "https://registry.terraform.io/providers/hashicorp/aws",
			CallerName:                  "Terraform AWS Provider",
			CredsFilename:               "",
			DebugLogging:                logging.IsDebugOrHigher(),
			IamEndpoint:                 "",
			Insecure:                    false,
			MaxRetries:                  25,
			Profile:                     d.Get("profile").(string),
			Region:                      d.Get("region").(string),
			SecretKey:                   "",
			SkipCredsValidation:         false,
			SkipMetadataApiCheck:        false,
			SkipRequestingAccountId:     false,
			StsEndpoint:                 "",
			Token:                       "",
			UserAgentProducts: []*awsbase.UserAgentProduct{
				{Name: "APN", Version: "1.0"},
				{Name: "HashiCorp", Version: "1.0"},
				{Name: "Terraform", Version: terraformVersion,
					Extra: []string{"+https://www.terraform.io"}},
			},
		}

		sess, err := awsbase.GetSession(awsbaseConfig)

		if err != nil {
			return nil, fmt.Errorf("error configuring AWS Lamda provider: %w", err)
		}

		config := &providerConfig{
			session: sess,
		}

		return config, nil
	}

	return provider
}
