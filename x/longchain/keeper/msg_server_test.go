package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hspdev0814/LongChain/testutil/keeper"
	"github.com/hspdev0814/LongChain/x/longchain/keeper"
	"github.com/hspdev0814/LongChain/x/longchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LongchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
