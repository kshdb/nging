package listener

import (
	"github.com/admpub/nging/application/library/fileupdater"
)

var (
	// GenUpdater 生成Updater
	GenUpdater      = fileupdater.GenUpdater
	NewProperty     = fileupdater.NewProperty
	NewPropertyWith = fileupdater.NewPropertyWith
	ThumbValue      = fileupdater.ThumbValue
)

type (
	// Property 附加属性
	Property   = fileupdater.Property
	ValueFunc  = fileupdater.ValueFunc
	FieldValue = fileupdater.FieldValue
)
