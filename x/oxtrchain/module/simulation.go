package oxtrchain

import (
	"cosmossdk.io/core/address"
	sdk "cosmossdk.io/types"
	"cosmossdk.io/types/module"
	simtypes "cosmossdk.io/types/simulation"

	"github.com/oxtr/oxtrchain/x/oxtrchain/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oxtrchainGenesis := types.GenesisState{
		Params:	types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oxtrchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
	}
}
