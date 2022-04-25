package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/maxeinhorn/hellohello/testutil/keeper"
	"github.com/maxeinhorn/hellohello/x/hellohello/keeper"
	"github.com/maxeinhorn/hellohello/x/hellohello/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HellohelloKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
