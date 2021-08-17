package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSsoAdminInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSSOAdminClient(ctrl)

	i := ssoadmin.ListInstancesOutput{}
	if err := faker.FakeData(&i); err != nil {
		t.Fatal(err)
	}
	i.NextToken = nil
	m.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(&i, nil)

	p := ssoadmin.ListPermissionSetsOutput{}
	if err := faker.FakeData(&p); err != nil {
		t.Fatal(err)
	}
	p.NextToken = nil
	m.EXPECT().ListPermissionSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&p, nil)

	ps := ssoadmin.DescribePermissionSetOutput{}
	if err := faker.FakeData(&ps); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribePermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ps, nil)

	aa := ssoadmin.ListAccountAssignmentsOutput{}
	if err := faker.FakeData(&aa); err != nil {
		t.Fatal(err)
	}
	aa.NextToken = nil
	m.EXPECT().ListAccountAssignments(gomock.Any(), gomock.Any(), gomock.Any()).Return(&aa, nil)

	policy := "{\"hello\":1}"
	m.EXPECT().GetInlinePolicyForPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ssoadmin.GetInlinePolicyForPermissionSetOutput{
		InlinePolicy: &policy,
	}, nil)

	return client.Services{SSOAdmin: m}
}

func TestSsoAdminInstances(t *testing.T) {
	awsTestHelper(t, SsoAdminInstances(), buildSsoAdminInstances, TestOptions{})
}
