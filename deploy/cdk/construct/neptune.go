package construct

import (
	"fmt"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	neptune "github.com/aws/aws-cdk-go/awscdk/v2/awsneptune"
	"github.com/aws/jsii-runtime-go"
)

type NeptuneConstruct struct {
	DatabaseCluster  neptune.CfnDBCluster
	DatabaseInstance neptune.CfnDBInstance
	Scope            awscdk.Stack
}

func newClusterProps(name string, props *neptune.CfnDBClusterProps) *neptune.CfnDBClusterProps {
	if props == nil {
		return &neptune.CfnDBClusterProps{
			AssociatedRoles:             nil,
			BackupRetentionPeriod:       jsii.Number(1),
			DbClusterIdentifier:         jsii.String(strings.ToLower(fmt.Sprintf("%s-ClusterId", name))),
			DbClusterParameterGroupName: nil,
			DbSubnetGroupName:           nil,
			DeletionProtection:          jsii.Bool(false),
			EnableCloudwatchLogsExports: nil,
			EngineVersion:               nil,
			IamAuthEnabled:              jsii.Bool(true),
			KmsKeyId:                    nil,
			Port:                        nil,
			PreferredBackupWindow:       nil,
			PreferredMaintenanceWindow:  nil,
			RestoreToTime:               nil,
			RestoreType:                 nil,
			SnapshotIdentifier:          nil,
			SourceDbClusterIdentifier:   nil,
			StorageEncrypted:            nil,
			Tags:                        nil,
			UseLatestRestorableTime:     nil,
			VpcSecurityGroupIds:         nil,
		}
	}

	return props
}

func newInstanceProps(name string, cluster neptune.CfnDBCluster, props *neptune.CfnDBInstanceProps) *neptune.CfnDBInstanceProps {
	if props == nil {
		return &neptune.CfnDBInstanceProps{
			DbInstanceClass:            jsii.String("db.t3.medium"),
			AllowMajorVersionUpgrade:   nil,
			AutoMinorVersionUpgrade:    nil,
			DbClusterIdentifier:        cluster.DbClusterIdentifier(),
			DbInstanceIdentifier:       jsii.String(strings.ToLower(fmt.Sprintf("%s-InstanceId", name))),
			DbParameterGroupName:       nil,
			DbSnapshotIdentifier:       nil,
			DbSubnetGroupName:          nil,
			PreferredMaintenanceWindow: nil,
			Tags:                       nil,
		}
	}

	return props
}

func NewNeptuneConstruct(
	stack awscdk.Stack,
	name string,
	clusterProps *neptune.CfnDBClusterProps,
	instanceProps *neptune.CfnDBInstanceProps,
) *NeptuneConstruct {
	cluster := neptune.NewCfnDBCluster(
		stack,
		jsii.String(fmt.Sprintf("%sClusterConstruct", name)),
		newClusterProps(name, clusterProps),
	)

	instance := neptune.NewCfnDBInstance(
		stack,
		jsii.String(fmt.Sprintf("%sInstanceConstruct", name)),
		newInstanceProps(name, cluster, instanceProps),
	)

	instance.AddDependsOn(cluster)

	return &NeptuneConstruct{
		DatabaseCluster:  cluster,
		DatabaseInstance: instance,
		Scope:            stack,
	}
}
