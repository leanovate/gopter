package prop

import "github.com/untoldwind/gopter"

// SampleCheck performs a property check on a specific SampleCheck
// This is what testers actually have to implmeent
type SampleCheck func(...interface{}) (interface{}, error)

func convertResult(args []gopter.PropArg, result interface{}, err error) gopter.PropResult {
	if err != nil {
		return gopter.PropResult{
			Status: gopter.PropError,
			Error:  err,
			Args:   args,
		}
	}
	return gopter.PropResult{}
}
