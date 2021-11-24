package resources

import (
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildDmsReplicationInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDMSClient(ctrl)
	instances := databasemigrationservice.DescribeReplicationInstancesOutput{}
	err := faker.FakeData(&instances)
	if err != nil {
		t.Fatal(err)
	}
	instances.Marker = nil
	ip := faker.IPv4()
	instances.ReplicationInstances[0].ReplicationInstancePrivateIpAddress = &ip
	instances.ReplicationInstances[0].ReplicationInstancePublicIpAddress = &ip
	instances.ReplicationInstances[0].ReplicationInstancePublicIpAddresses[0] = ip
	instances.ReplicationInstances[0].ReplicationInstancePrivateIpAddresses[0] = ip
	m.EXPECT().DescribeReplicationInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&instances, nil)
	return client.Services{
		DMS: m,
	}
}

func TestDmsReplicationInstances(t *testing.T) {
	awsTestHelper(t, DmsReplicationInstances(), buildDmsReplicationInstances, TestOptions{})
}
