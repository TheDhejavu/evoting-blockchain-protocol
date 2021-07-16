package blockchain

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

// INITIALIZE ELECTION
// Init  election TxOutput
type TxElectionOutput struct {
	ID              string
	Signers         [][]byte
	SigWitnesses    [][]byte
	ElectionKeyHash []byte
	Title           string
	Description     string
	TotalPeople     int64
	Candidates      [][]byte
}

// End Election TxInput
type TxElectionInput struct {
	TxID            []byte
	Signers         [][]byte
	SigWitnesses    [][]byte
	TxOut           string
	ElectionKeyHash []byte
}

// NewTxAccreditationInput Stops Accreditation  Phase
func NewElectionTxInput(keyHash, txId []byte, txOut string, signers, SigWitnesses [][]byte) *TxInput {
	tx := &TxInput{
		ElectionTx: TxElectionInput{
			TxID:            txId,
			Signers:         signers,
			SigWitnesses:    SigWitnesses,
			ElectionKeyHash: keyHash,
			TxOut:           txOut,
		},
	}
	return tx
}

// NewTxAccreditationTxOutput Starts Accreditation Phase
func NewElectionTxOutput(title, desp string, keyHash []byte, signers, SigWitnesses, candidates [][]byte, totalPeople int64) *TxOutput {
	tx := &TxOutput{
		ElectionTx: TxElectionOutput{
			Signers:         signers,
			SigWitnesses:    SigWitnesses,
			ElectionKeyHash: keyHash,
			Title:           title,
			Description:     desp,
			TotalPeople:     totalPeople,
			Candidates:      candidates,
		},
	}
	uuid, _ := uuid.NewUUID()
	tx.ElectionTx.ID = uuid.String()
	return tx
}

// Convert Election output to Byte for verification and signing purposes
func (tx TxElectionOutput) TrimmedCopy() *TxElectionOutput {
	txCopy := &TxElectionOutput{
		"",
		nil,
		nil,
		tx.ElectionKeyHash,
		tx.Title,
		tx.Description,
		tx.TotalPeople,
		tx.Candidates,
	}
	return txCopy
}

// Convert Election output to Byte for verification and signing purposes
func (tx *TxElectionOutput) ToByte() []byte {
	var hash [32]byte

	txCopy := tx.TrimmedCopy()

	hash = sha256.Sum256([]byte(fmt.Sprintf("%x", txCopy)))
	return hash[:]
}

func (tx *TxElectionOutput) IsSet() bool {
	return reflect.DeepEqual(tx, &TxElectionOutput{}) == false
}

// Trim election input data
func (tx *TxElectionInput) TrimmedCopy() *TxElectionInput {
	txCopy := &TxElectionInput{
		tx.TxID,
		nil,
		nil,
		tx.TxOut,
		tx.ElectionKeyHash,
	}
	return txCopy
}

// Convert Election output to Byte for verification and signing purposes
func (tx *TxElectionInput) ToByte() []byte {
	var hash [32]byte

	txCopy := tx.TrimmedCopy()

	hash = sha256.Sum256([]byte(fmt.Sprintf("%x", txCopy)))
	return hash[:]
}

func (tx *TxElectionInput) IsSet() bool {
	return reflect.DeepEqual(tx, &TxElectionInput{}) == false
}

// Helper function for displaying transaction data in the console
func (tx *TxElectionInput) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf("--TX_INPUT: %x", tx.TxID))
	if tx.IsSet() {
		for i := 0; i < len(tx.Signers); i++ {
			lines = append(lines, fmt.Sprintf("(Signers) \n --(%d): %x", i, tx.Signers[i]))
		}
		for i := 0; i < len(tx.SigWitnesses); i++ {
			lines = append(lines, fmt.Sprintf("(Signature Witness): \n --(%d): %x", i, tx.SigWitnesses[i]))
		}
		lines = append(lines, fmt.Sprintf("TxOut: %s", tx.TxOut))
		lines = append(lines, fmt.Sprintf("Election Keyhash: %x", tx.ElectionKeyHash))
	}

	return strings.Join(lines, "\n")
}

// Helper function for displaying transaction data in the console
func (tx *TxElectionOutput) String() string {
	var lines []string
	if tx.IsSet() {
		lines = append(lines, fmt.Sprintf("--TX_OUTPUT \n(ID): %x", tx.ID))
		lines = append(lines, fmt.Sprintf("(TxID): %x", tx.ID))
		lines = append(lines, fmt.Sprintf("(Title): %s", tx.Title))
		for i := 0; i < len(tx.Signers); i++ {
			lines = append(lines, fmt.Sprintf("(Signers) \n --(%d): %x", i, tx.Signers[i]))
		}
		for i := 0; i < len(tx.SigWitnesses); i++ {
			lines = append(lines, fmt.Sprintf("(Signature Witness): \n --(%d): %x", i, tx.SigWitnesses[i]))
		}
		lines = append(lines, fmt.Sprintf("(Description): %s", tx.Description))
		lines = append(lines, fmt.Sprintf("(People): %d", tx.TotalPeople))
		lines = append(lines, fmt.Sprintf("(Election Keyhash): %x", tx.ElectionKeyHash))
	}
	return strings.Join(lines, "\n")
}