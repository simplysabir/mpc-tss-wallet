package tss

import (
	"fmt"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
)

// TssParty represents a participant in the TSS protocol
type TssParty struct {
    Party tss.Party
    OutCh chan tss.Message
    EndCh chan keygen.LocalPartySaveData
}

// NewTssParty creates a new TSS party
func NewTssParty(params *tss.Parameters, partyID *tss.PartyID, partyCount int, threshold int) (*TssParty, error) {
    outCh := make(chan tss.Message, partyCount)
    endCh := make(chan keygen.LocalPartySaveData, 1)
    
    party := keygen.NewLocalParty(params, outCh, endCh, partyID)
    if party == nil {
        return nil, fmt.Errorf("failed to create local party")
    }

    return &TssParty{
        Party: party,
        OutCh: outCh,
        EndCh: endCh,
    }, nil
}

// StartKeyGeneration starts the key generation process for this party
func (p *TssParty) StartKeyGeneration() error {
    go func() {
        if err := p.Party.Start(); err != nil {
            fmt.Printf("Error starting party: %v\n", err)
        }
    }()

    return nil
}

// HandleMessage processes an incoming message from another party
func (p *TssParty) HandleMessage(msg tss.Message) error {
    dest := msg.GetTo()
    if dest == nil {
        for _, id := range p.Party.WrapInfo().PartyIDMap {
            if id.Index == msg.GetFrom().Index {
                dest = []*tss.PartyID{id}
                break
            }
        }
    }
    err := p.Party.UpdateFromBytes(msg.GetContent(), msg.GetFrom(), dest)
    if err != nil {
        return fmt.Errorf("failed to update party: %v", err)
    }
    return nil
}

// GetResult waits for and returns the key generation result
func (p *TssParty) GetResult() (*keygen.LocalPartySaveData, error) {
    select {
    case save := <-p.EndCh:
        return &save, nil
    case <-p.Party.WrapInfo().AbortingChan():
        return nil, fmt.Errorf("key generation aborted")
    }
}