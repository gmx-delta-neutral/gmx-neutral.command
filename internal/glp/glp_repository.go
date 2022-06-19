package glp

import (
	"context"
	"errors"
	"math/big"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ethereum/go-ethereum/common"
)

type Repository interface {
	BuyGlp(amount big.Int) error
}

type GlpRepository struct {
}

type PositionType int

const (
	Buy PositionType = iota
	Sell
)

func NewGlpRepository() Repository {
	return &GlpRepository{}
}

func (p PositionType) String() string {
	switch p {
	case Buy:
		return "Buy"
	case Sell:
		return "Sell"
	}
	return ""
}

var walletAddress = common.HexToAddress("0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664")
var tokenAddress = common.HexToAddress("0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664")

func (r *GlpRepository) BuyGlp(amount big.Int) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)

	if err != nil {
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("TokenPositions"),
		Key: map[string]types.AttributeValue{
			"TokenAddress":  &types.AttributeValueMemberS{Value: tokenAddress.String()},
			"WalletAddress": &types.AttributeValueMemberS{Value: walletAddress.String()},
		},
	})

	if err != nil {
		return err
	}

	if out.Item != nil {
		oldBalance := ""
		attributevalue.Unmarshal(out.Item["Balance"], &oldBalance)
		balance, success := new(big.Int).SetString(oldBalance, 10)

		if !success {
			return errors.New("Error unmarshaling balance")
		}

		newBalance := new(big.Int).Add(balance, &amount)

		_, err := svc.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String("TokenPositions"),
			Key: map[string]types.AttributeValue{
				"TokenAddress":  &types.AttributeValueMemberS{Value: tokenAddress.String()},
				"WalletAddress": &types.AttributeValueMemberS{Value: walletAddress.String()},
			},
			UpdateExpression: aws.String("set Balance = :balance"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":balance": &types.AttributeValueMemberS{Value: newBalance.String()},
			},
		})

		if err != nil {
			return errors.New("Error updating balance")
		}

		return nil
	}

	_, err = svc.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("TokenPositions"),
		Item: map[string]types.AttributeValue{
			"TokenAddress":  &types.AttributeValueMemberS{Value: tokenAddress.String()},
			"WalletAddress": &types.AttributeValueMemberS{Value: walletAddress.String()},
			"Symbol":        &types.AttributeValueMemberS{Value: "GLP"},
			"Balance":       &types.AttributeValueMemberS{Value: amount.String()},
			"Decimals":      &types.AttributeValueMemberN{Value: "18"},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
