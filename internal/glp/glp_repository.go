package glp

import (
	"context"
	"math/big"
	"os"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/internal/contracts/glp"
	glpmanager "github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/internal/contracts/glp_manager"
	rewardrouter "github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/internal/contracts/reward_router"
	util "github.com/RafGDev/gmx-delta-neutral/gmx-neutral.command/internal/util"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

type Repository interface {
	BuyGlp(amount *big.Int) error
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
var glpManagerAddress = common.HexToAddress("0xe1ae4d4b06A5Fe1fc288f6B4CD72f9F8323B107F")
var glpAddress = common.HexToAddress("0x01234181085565ed162a948b6a5e88758CD7c7b8")
var rewardRouterAddress = common.HexToAddress("0x82147c5a7e850ea4e28155df107f2590fd4ba327")

var glpDecimals = 18
var usdcDecimals = 18

func (r *GlpRepository) BuyGlp(dollarAmount *big.Int) error {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")

	if err != nil {
		return err
	}

	ctr, err := glpmanager.NewGlpmanager(glpManagerAddress, client)

	if err != nil {
		return err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	aum, err := ctr.GetAumInUsdg(callOpts, true)
	expandedDollarAmount := util.ExpandDecimals(dollarAmount, int64(usdcDecimals))

	if err != nil {
		return err
	}

	glpContract, err := glp.NewGlp(glpAddress, client)

	if err != nil {
		return err
	}

	supply, err := glpContract.TotalSupply(callOpts)

	if err != nil {
		return err
	}

	expandedGlpPrice := new(big.Int).Div(util.ExpandDecimals(aum, int64(glpDecimals)), supply)

	rewardCtr, err := rewardrouter.NewRewardrouter(rewardRouterAddress, client)

	if err != nil {
		return err
	}

	amountToPurchase := new(big.Int).Div(util.ExpandDecimals(expandedDollarAmount, int64(usdcDecimals)), expandedGlpPrice)
	minUsdg := new(big.Int)

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		return err
	}

	chainId := big.NewInt(43114)
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	if err != nil {
		return err
	}

	_, err = rewardCtr.MintAndStakeGlpETH(opts, minUsdg, amountToPurchase)

	if err != nil {
		return err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)

	if err != nil {
		return err
	}

	svc := dynamodb.NewFromConfig(cfg)

	guid := uuid.New()

	_, err = svc.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("TokenPositions"),
		Item: map[string]types.AttributeValue{
			"TransactionId": &types.AttributeValueMemberS{Value: guid.String()},
			"TokenAddress":  &types.AttributeValueMemberS{Value: tokenAddress.String()},
			"WalletAddress": &types.AttributeValueMemberS{Value: walletAddress.String()},
			"Symbol":        &types.AttributeValueMemberS{Value: "GLP"},
			"Amount":        &types.AttributeValueMemberS{Value: amountToPurchase.String()},
			"Decimals":      &types.AttributeValueMemberN{Value: "18"},
			"PurchasePrice": &types.AttributeValueMemberS{Value: expandedGlpPrice.String()},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
