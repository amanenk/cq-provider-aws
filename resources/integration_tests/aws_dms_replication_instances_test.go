package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationDmsReplicationInstances(t *testing.T) {
	resource := resources.DmsReplicationInstances()
	awsTestIntegrationHelper(t, resource, []string{"aws_dms_replication_instances.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resource.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"type":       "ipsec.1",
						"ip_address": "172.83.124.10",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ec2-cgw-%s-%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
