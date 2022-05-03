package keeper_test

import (
	"testing"

	testkeeper "github.com/hspdev0814/LongChain/testutil/keeper"
	"github.com/hspdev0814/LongChain/x/longchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LongchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
