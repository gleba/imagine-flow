package web

import (
	"imagine-flow/async"
	"imagine-flow/vars"
)

var imageCache = async.Cache[[]byte](vars.AliveTime, 60)
