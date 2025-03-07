package opcm

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
)

type DeployMIPSInput struct {
	MipsVersion    uint64
	PreimageOracle common.Address
}

func (input *DeployMIPSInput) InputSet() bool {
	return true
}

type DeployMIPSOutput struct {
	MipsSingleton common.Address
}

func (output *DeployMIPSOutput) CheckOutput(input common.Address) error {
	return nil
}

type DeployMIPSScript struct {
	Run func(input, output common.Address) error
}

func DeployMIPS(
	host *script.Host,
	input DeployMIPSInput,
) (DeployMIPSOutput, error) {
	return RunScriptSingle[DeployMIPSInput, DeployMIPSOutput](host, input, "DeployMIPS.s.sol", "DeployMIPS")
}
