package pool

import (
	"github.com/alitto/pond/v2"
)

var IOPool = pond.NewPool(1000)
