package prop

import "github.com/leanovate/gopter"

func ErrorProp(err error) gopter.Prop {
	return func(genParams *gopter.GenParameters) *gopter.PropResult {
		return &gopter.PropResult{
			Status: gopter.PropError,
			Error:  err,
		}
	}
}
