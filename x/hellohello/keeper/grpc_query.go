package keeper

import (
	"github.com/maxeinhorn/hellohello/x/hellohello/types"
)

var _ types.QueryServer = Keeper{}
