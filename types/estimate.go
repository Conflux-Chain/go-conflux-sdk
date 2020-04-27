// Copyright 2019 Conflux Foundation. All rights reserved.
// Conflux is free software and distributed under GNU General Public License.
// See http://www.gnu.org/licenses/

package types

import "github.com/ethereum/go-ethereum/common/hexutil"

//Estimate represents estimated gas will be used and storage will be collateralized when transaction excutes
type Estimate struct {
	GasUsed               *hexutil.Big `json:"gasUsed"`
	StorageCollateralized *hexutil.Big `json:"storageCollateralized"`
}
