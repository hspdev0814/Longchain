package longchain_test

import (
	"testing"

	keepertest "github.com/hspdev0814/LongChain/testutil/keeper"
	"github.com/hspdev0814/LongChain/testutil/nullify"
	"github.com/hspdev0814/LongChain/x/longchain"
	"github.com/hspdev0814/LongChain/x/longchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LongchainKeeper(t)
	longchain.InitGenesis(ctx, *k, genesisState)
	got := longchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
