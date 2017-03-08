package check

import stdCheck "sensu-common/check"

var Store = make(map[string]Check)

type Check interface {
	Execute() stdCheck.CheckOutput
}
