package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sideprotocol/sidechain/testutil/keeper"
	"github.com/sideprotocol/sidechain/x/devearn/keeper"
	"github.com/sideprotocol/sidechain/x/devearn/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DevearnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
