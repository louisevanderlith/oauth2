package core

import "github.com/louisevanderlith/husk"

type authcodeFilter func(obj AuthCode) bool

func (f authcodeFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(AuthCode))
}

//byID filter will filter by client_id
func byAuthCode(code string) authcodeFilter {
	return func(obj AuthCode) bool {
		return obj.Code == code
	}
}