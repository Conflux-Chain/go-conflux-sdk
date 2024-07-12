package sdk

import (
	"testing"
)

func TestPrivatekeyAccountManagerInterface(t *testing.T) {
	var _ AccountManagerOperator = NewPrivatekeyAccountManager(nil, 1)
}
