package gdynamo

import (
	"github.com/stefan79/gadgeto/pkg/resources/aws"
)

type (
	BillingMode string
	TableClass  string
	Action      string

	//TODO Rename the Specification Methos
	DynamoDBBBuilder[ID any, Type any] interface {
		WithTag(key string, value string) DynamoDBBBuilder[ID, Type]
		WithPermissions(action ...Action) DynamoDBBBuilder[ID, Type]
		Build() DynamoDBResource[ID, Type]

		//TODO add this
		//WithGlobalSecondaryIndex(name string, keyschema interface{}, projection interface{}, proivisionedThroughput interface{}) DynamoDBBBuilder
		//WithLocalSecondaryIndex(name string, keyschema interface{}, projection interface{}) DynamoDBBBuilder
		//WithKinesisStreamSpecification(spec interface{}) DynamoDBBBuilder
		//WithStreamSpecification() DynamoDBBBuilder
		//WithTTL() DynamoDBBBuilder

		//TODO get this from the stage
		//WithBillingMode(billingMode BillingMode) DynamoDBBBuilder
		//WithDeletionProtectionEnabled() DynamoDBBBuilder
		//Create a Bucket definition
		//WithPointInTimeRecoverySpecification(spec interface{}) DynamoDBBBuilder
		//WithImportFromS3(gs3.Client)
		//WithProvisionedThroughput(spec interface{}) DynamoDBBBuilder
		//WithServerSideEncryptionSpecification(spec interface{}) DynamoDBBBuilder

		//TODO create a meta class which coordinates ProvisionedThroughput and BillingMode
		//TODO make this an own type so it can be dropped in as trigger
	}

	builderConf[ID any, Type any] struct {
		aws.BaseAWSResource[DynamoDBResource[ID, Type]]
		primaryKeys     []attribute
		tableAttributes []attribute
		tags            map[string]string
		permissions     []Action
	}
)

var (
	Provisioned   BillingMode = "PROVISIONED"
	PayPerRequest BillingMode = "PAY_PER_REQUEST"

	Standard         TableClass = "STANDARD"
	InfrequentAccess TableClass = "STANDARD_INFREQUENT_ACCESS"

	BatchGetItem                        Action = "BatchGetItem"
	BatchWriteItem                      Action = "BatchWriteItem"
	ConditionCheckItem                  Action = "ConditionCheckItem"
	CreateBackup                        Action = "CreateBackup"
	CreateGlobalTable                   Action = "CreateGlobalTable"
	CreateTable                         Action = "CreateTable"
	DeleteBackup                        Action = "DeleteBackup"
	DeleteItem                          Action = "DeleteItem"
	DeleteTable                         Action = "DeleteTable"
	DescribeBackup                      Action = "DescribeBackup"
	DescribeContinuousBackups           Action = "DescribeContinuousBackups"
	DescribeContributorInsights         Action = "DescribeContributorInsights"
	DescribeEndpoints                   Action = "DescribeEndpoints"
	DescribeExport                      Action = "DescribeExport"
	DescribeGlobalTable                 Action = "DescribeGlobalTable"
	DescribeGlobalTableSettings         Action = "DescribeGlobalTableSettings"
	DescribeKinesisStreamingDestination Action = "DescribeKinesisStreamingDestination"
	DescribeLimits                      Action = "DescribeLimits"
	DescribeTable                       Action = "DescribeTable"
	DescribeTableReplicaAutoScaling     Action = "DescribeTableReplicaAutoScaling"
	DescribeTimeToLive                  Action = "DescribeTimeToLive"
	DisableKinesisStreamingDestination  Action = "DisableKinesisStreamingDestination"
	EnableKinesisStreamingDestination   Action = "EnableKinesisStreamingDestination"
	ExecuteStatement                    Action = "ExecuteStatement"
	ExecuteTransaction                  Action = "ExecuteTransaction"
	ExportTableToPointInTime            Action = "ExportTableToPointInTime"
	GetItem                             Action = "GetItem"
	ListBackups                         Action = "ListBackups"
	ListContributorInsights             Action = "ListContributorInsights"
	ListExports                         Action = "ListExports"
	ListGlobalTables                    Action = "ListGlobalTables"
	ListTables                          Action = "ListTables"
	ListTagsOfResource                  Action = "ListTagsOfResource"
	PutItem                             Action = "PutItem"
	Query                               Action = "Query"
	RestoreTableFromBackup              Action = "RestoreTableFromBackup"
	RestoreTableToPointInTime           Action = "RestoreTableToPointInTime"
	Scan                                Action = "Scan"
	TagResource                         Action = "TagResource"
	UntagResource                       Action = "UntagResource"
	UpdateContinuousBackups             Action = "UpdateContinuousBackups"
	UpdateContributorInsights           Action = "UpdateContributorInsights"
	UpdateGlobalTable                   Action = "UpdateGlobalTable"
	UpdateGlobalTableSettings           Action = "UpdateGlobalTableSettings"
	UpdateItem                          Action = "UpdateItem"
	UpdateTable                         Action = "UpdateTable"
	UpdateTableReplicaAutoScaling       Action = "UpdateTableReplicaAutoScaling"
	UpdateTimeToLive                    Action = "UpdateTimeToLive"

	ControlPlanActions = []Action{
		CreateBackup,
		CreateGlobalTable,
		CreateTable,
		DeleteBackup,
		DeleteTable,
		DescribeBackup,
		DescribeContinuousBackups,
		DescribeContributorInsights,
		DescribeEndpoints,
		DescribeExport,
		DescribeGlobalTable,
		DescribeGlobalTableSettings,
		DescribeKinesisStreamingDestination,
		DescribeLimits,
		DescribeTable,
		DescribeTableReplicaAutoScaling,
		DescribeTimeToLive,
		DisableKinesisStreamingDestination,
		EnableKinesisStreamingDestination,
		ExportTableToPointInTime,
		ListBackups,
		ListContributorInsights,
		ListExports,
		ListGlobalTables,
		ListTables,
		ListTagsOfResource,
		RestoreTableFromBackup,
		RestoreTableToPointInTime,
		TagResource,
		UntagResource,
		UpdateContinuousBackups,
		UpdateContributorInsights,
		UpdateGlobalTable,
		UpdateGlobalTableSettings,
		UpdateTable,
		UpdateTableReplicaAutoScaling,
		UpdateTimeToLive,
	}

	TablePlaneReadActions = []Action{
		BatchGetItem,
		GetItem,
		Query,
		Scan,
	}

	TablePlaneWriteActions = []Action{
		BatchWriteItem,
		ConditionCheckItem,
		UpdateItem,
		PutItem,
	}

	TablePlaneDeleteActions = []Action{
		DeleteItem,
	}

	TablePlaneActions = append(
		TablePlaneReadActions,
		append(
			TablePlaneWriteActions,
			TablePlaneDeleteActions...,
		)...,
	)
)

func (b *builderConf[ID, Type]) WithTableDefinition(spec interface{}) DynamoDBBBuilder[ID, Type] {
	return b
}

func (b *builderConf[ID, Type]) WithTag(key string, value string) DynamoDBBBuilder[ID, Type] {
	b.tags[key] = value
	return b
}

func (b *builderConf[ID, Type]) WithPermissions(action ...Action) DynamoDBBBuilder[ID, Type] {
	b.permissions = append(b.permissions, action...)
	return b
}

func (b *builderConf[ID, Type]) Build() DynamoDBResource[ID, Type] {
	trigger, err := b.Mode.Dispatch(b)
	if err != nil {
		panic(err)
	}
	return trigger
}
