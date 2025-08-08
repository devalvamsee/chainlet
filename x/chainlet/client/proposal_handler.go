package client

import (
	"github.com/devalvamsee/chainlet/x/chainlet/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the token mapping change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitTokenMappingChangeProposalTxCmd)
