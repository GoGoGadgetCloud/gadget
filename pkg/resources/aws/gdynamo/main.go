package gdynamo

import (
	"reflect"

	"github.com/stefan79/gadgeto/pkg/modes"
	"github.com/stefan79/gadgeto/pkg/resources/aws"
)

type (
	DynamoDBResource[ID any, Target any] interface {
		GetAWSSDKClient(interface{}, error)
		GetTableName() string
		CreateCRUDHelper(CRUDSHelper[ID, Target], error)
	}

	CRUDSHelper[ID any, Target any] interface {
		Create(Target) error
		Read(ID) (Target, error)
		Update(Target) error
		Delete(ID) error
	}
)

func NewDynamoDB[ID any, Type any](name string, mode modes.Mode[interface{}]) DynamoDBBBuilder[ID, Type] {

	idType := reflect.TypeOf((*ID)(nil)).Elem()
	valueType := reflect.TypeOf((*ID)(nil)).Elem()

	primaryKeys, err := getPrimaryKeyAttributeDefinitions(idType)
	if err != nil {
		panic(err)
	}
	tableAttributes, err := geTableAttributeDefinitions(valueType, idType)
	if err != nil {
		panic(err)
	}

	return &builderConf[ID, Type]{
		BaseAWSResource: aws.BaseAWSResource[DynamoDBResource[ID, Type]]{
			Name: &name,
			Mode: modes.DowncastMode[DynamoDBResource[ID, Type]](mode),
		},
		primaryKeys:     primaryKeys,
		tableAttributes: tableAttributes,
		tags:            map[string]string{},
	}
}
