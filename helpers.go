package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ConnectionItem struct {
	ConnectionID string `json:"connectionID"`
}

const (
	AccessKeyID = "AKIAICBTL6PBD4F3EWYA"
	SecretAccessKey = "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"
	APIGatewayEndpoint = "https://7b4qjmnofj.execute-api.us-east-2.amazonaws.com/development"
	Region = "us-east-1"
)

func NewDynamoDBSession() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	})
	return dynamodb.New(sess)
}

func NewAPIGatewaySession() *apigatewaymanagementapi.ApiGatewayManagementApi {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
		Endpoint: aws.String(APIGatewayEndpoint),
	})
	return apigatewaymanagementapi.New(sess)
}

