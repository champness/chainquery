// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Addresses", testAddresses)
	t.Run("Blocks", testBlocks)
	t.Run("Claims", testClaims)
	t.Run("Inputs", testInputs)
	t.Run("Outputs", testOutputs)
	t.Run("Transactions", testTransactions)
	t.Run("TransactionAddresses", testTransactionAddresses)
	t.Run("UnknownClaims", testUnknownClaims)
}

func TestDelete(t *testing.T) {
	t.Run("Addresses", testAddressesDelete)
	t.Run("Blocks", testBlocksDelete)
	t.Run("Claims", testClaimsDelete)
	t.Run("Inputs", testInputsDelete)
	t.Run("Outputs", testOutputsDelete)
	t.Run("Transactions", testTransactionsDelete)
	t.Run("TransactionAddresses", testTransactionAddressesDelete)
	t.Run("UnknownClaims", testUnknownClaimsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Addresses", testAddressesQueryDeleteAll)
	t.Run("Blocks", testBlocksQueryDeleteAll)
	t.Run("Claims", testClaimsQueryDeleteAll)
	t.Run("Inputs", testInputsQueryDeleteAll)
	t.Run("Outputs", testOutputsQueryDeleteAll)
	t.Run("Transactions", testTransactionsQueryDeleteAll)
	t.Run("TransactionAddresses", testTransactionAddressesQueryDeleteAll)
	t.Run("UnknownClaims", testUnknownClaimsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Addresses", testAddressesSliceDeleteAll)
	t.Run("Blocks", testBlocksSliceDeleteAll)
	t.Run("Claims", testClaimsSliceDeleteAll)
	t.Run("Inputs", testInputsSliceDeleteAll)
	t.Run("Outputs", testOutputsSliceDeleteAll)
	t.Run("Transactions", testTransactionsSliceDeleteAll)
	t.Run("TransactionAddresses", testTransactionAddressesSliceDeleteAll)
	t.Run("UnknownClaims", testUnknownClaimsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Addresses", testAddressesExists)
	t.Run("Blocks", testBlocksExists)
	t.Run("Claims", testClaimsExists)
	t.Run("Inputs", testInputsExists)
	t.Run("Outputs", testOutputsExists)
	t.Run("Transactions", testTransactionsExists)
	t.Run("TransactionAddresses", testTransactionAddressesExists)
	t.Run("UnknownClaims", testUnknownClaimsExists)
}

func TestFind(t *testing.T) {
	t.Run("Addresses", testAddressesFind)
	t.Run("Blocks", testBlocksFind)
	t.Run("Claims", testClaimsFind)
	t.Run("Inputs", testInputsFind)
	t.Run("Outputs", testOutputsFind)
	t.Run("Transactions", testTransactionsFind)
	t.Run("TransactionAddresses", testTransactionAddressesFind)
	t.Run("UnknownClaims", testUnknownClaimsFind)
}

func TestBind(t *testing.T) {
	t.Run("Addresses", testAddressesBind)
	t.Run("Blocks", testBlocksBind)
	t.Run("Claims", testClaimsBind)
	t.Run("Inputs", testInputsBind)
	t.Run("Outputs", testOutputsBind)
	t.Run("Transactions", testTransactionsBind)
	t.Run("TransactionAddresses", testTransactionAddressesBind)
	t.Run("UnknownClaims", testUnknownClaimsBind)
}

func TestOne(t *testing.T) {
	t.Run("Addresses", testAddressesOne)
	t.Run("Blocks", testBlocksOne)
	t.Run("Claims", testClaimsOne)
	t.Run("Inputs", testInputsOne)
	t.Run("Outputs", testOutputsOne)
	t.Run("Transactions", testTransactionsOne)
	t.Run("TransactionAddresses", testTransactionAddressesOne)
	t.Run("UnknownClaims", testUnknownClaimsOne)
}

func TestAll(t *testing.T) {
	t.Run("Addresses", testAddressesAll)
	t.Run("Blocks", testBlocksAll)
	t.Run("Claims", testClaimsAll)
	t.Run("Inputs", testInputsAll)
	t.Run("Outputs", testOutputsAll)
	t.Run("Transactions", testTransactionsAll)
	t.Run("TransactionAddresses", testTransactionAddressesAll)
	t.Run("UnknownClaims", testUnknownClaimsAll)
}

func TestCount(t *testing.T) {
	t.Run("Addresses", testAddressesCount)
	t.Run("Blocks", testBlocksCount)
	t.Run("Claims", testClaimsCount)
	t.Run("Inputs", testInputsCount)
	t.Run("Outputs", testOutputsCount)
	t.Run("Transactions", testTransactionsCount)
	t.Run("TransactionAddresses", testTransactionAddressesCount)
	t.Run("UnknownClaims", testUnknownClaimsCount)
}

func TestInsert(t *testing.T) {
	t.Run("Addresses", testAddressesInsert)
	t.Run("Addresses", testAddressesInsertWhitelist)
	t.Run("Blocks", testBlocksInsert)
	t.Run("Blocks", testBlocksInsertWhitelist)
	t.Run("Claims", testClaimsInsert)
	t.Run("Claims", testClaimsInsertWhitelist)
	t.Run("Inputs", testInputsInsert)
	t.Run("Inputs", testInputsInsertWhitelist)
	t.Run("Outputs", testOutputsInsert)
	t.Run("Outputs", testOutputsInsertWhitelist)
	t.Run("Transactions", testTransactionsInsert)
	t.Run("Transactions", testTransactionsInsertWhitelist)
	t.Run("TransactionAddresses", testTransactionAddressesInsert)
	t.Run("TransactionAddresses", testTransactionAddressesInsertWhitelist)
	t.Run("UnknownClaims", testUnknownClaimsInsert)
	t.Run("UnknownClaims", testUnknownClaimsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ClaimToTransactionUsingTransactionByHash", testClaimToOneTransactionUsingTransactionByHash)
	t.Run("ClaimToClaimUsingPublisher", testClaimToOneClaimUsingPublisher)
	t.Run("InputToAddressUsingInputAddress", testInputToOneAddressUsingInputAddress)
	t.Run("InputToTransactionUsingTransaction", testInputToOneTransactionUsingTransaction)
	t.Run("OutputToTransactionUsingTransaction", testOutputToOneTransactionUsingTransaction)
	t.Run("OutputToInputUsingSpentByInput", testOutputToOneInputUsingSpentByInput)
	t.Run("TransactionToBlockUsingBlockByHash", testTransactionToOneBlockUsingBlockByHash)
	t.Run("TransactionAddressToTransactionUsingTransaction", testTransactionAddressToOneTransactionUsingTransaction)
	t.Run("TransactionAddressToAddressUsingAddress", testTransactionAddressToOneAddressUsingAddress)
	t.Run("UnknownClaimToOutputUsingOutput", testUnknownClaimToOneOutputUsingOutput)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AddressToInputAddressInputs", testAddressToManyInputAddressInputs)
	t.Run("AddressToInputs", testAddressToManyInputs)
	t.Run("AddressToOutputs", testAddressToManyOutputs)
	t.Run("AddressToTransactionAddresses", testAddressToManyTransactionAddresses)
	t.Run("BlockToBlockByHashTransactions", testBlockToManyBlockByHashTransactions)
	t.Run("ClaimToPublisherClaims", testClaimToManyPublisherClaims)
	t.Run("InputToAddresses", testInputToManyAddresses)
	t.Run("InputToSpentByInputOutputs", testInputToManySpentByInputOutputs)
	t.Run("OutputToAddresses", testOutputToManyAddresses)
	t.Run("OutputToUnknownClaims", testOutputToManyUnknownClaims)
	t.Run("TransactionToTransactionByHashClaims", testTransactionToManyTransactionByHashClaims)
	t.Run("TransactionToInputs", testTransactionToManyInputs)
	t.Run("TransactionToOutputs", testTransactionToManyOutputs)
	t.Run("TransactionToTransactionAddresses", testTransactionToManyTransactionAddresses)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ClaimToTransactionUsingTransactionByHash", testClaimToOneSetOpTransactionUsingTransactionByHash)
	t.Run("ClaimToClaimUsingPublisher", testClaimToOneSetOpClaimUsingPublisher)
	t.Run("InputToAddressUsingInputAddress", testInputToOneSetOpAddressUsingInputAddress)
	t.Run("InputToTransactionUsingTransaction", testInputToOneSetOpTransactionUsingTransaction)
	t.Run("OutputToTransactionUsingTransaction", testOutputToOneSetOpTransactionUsingTransaction)
	t.Run("OutputToInputUsingSpentByInput", testOutputToOneSetOpInputUsingSpentByInput)
	t.Run("TransactionToBlockUsingBlockByHash", testTransactionToOneSetOpBlockUsingBlockByHash)
	t.Run("TransactionAddressToTransactionUsingTransaction", testTransactionAddressToOneSetOpTransactionUsingTransaction)
	t.Run("TransactionAddressToAddressUsingAddress", testTransactionAddressToOneSetOpAddressUsingAddress)
	t.Run("UnknownClaimToOutputUsingOutput", testUnknownClaimToOneSetOpOutputUsingOutput)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ClaimToTransactionUsingTransactionByHash", testClaimToOneRemoveOpTransactionUsingTransactionByHash)
	t.Run("ClaimToClaimUsingPublisher", testClaimToOneRemoveOpClaimUsingPublisher)
	t.Run("InputToAddressUsingInputAddress", testInputToOneRemoveOpAddressUsingInputAddress)
	t.Run("OutputToInputUsingSpentByInput", testOutputToOneRemoveOpInputUsingSpentByInput)
	t.Run("TransactionToBlockUsingBlockByHash", testTransactionToOneRemoveOpBlockUsingBlockByHash)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AddressToInputAddressInputs", testAddressToManyAddOpInputAddressInputs)
	t.Run("AddressToInputs", testAddressToManyAddOpInputs)
	t.Run("AddressToOutputs", testAddressToManyAddOpOutputs)
	t.Run("AddressToTransactionAddresses", testAddressToManyAddOpTransactionAddresses)
	t.Run("BlockToBlockByHashTransactions", testBlockToManyAddOpBlockByHashTransactions)
	t.Run("ClaimToPublisherClaims", testClaimToManyAddOpPublisherClaims)
	t.Run("InputToAddresses", testInputToManyAddOpAddresses)
	t.Run("InputToSpentByInputOutputs", testInputToManyAddOpSpentByInputOutputs)
	t.Run("OutputToAddresses", testOutputToManyAddOpAddresses)
	t.Run("OutputToUnknownClaims", testOutputToManyAddOpUnknownClaims)
	t.Run("TransactionToTransactionByHashClaims", testTransactionToManyAddOpTransactionByHashClaims)
	t.Run("TransactionToInputs", testTransactionToManyAddOpInputs)
	t.Run("TransactionToOutputs", testTransactionToManyAddOpOutputs)
	t.Run("TransactionToTransactionAddresses", testTransactionToManyAddOpTransactionAddresses)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AddressToInputAddressInputs", testAddressToManySetOpInputAddressInputs)
	t.Run("AddressToInputs", testAddressToManySetOpInputs)
	t.Run("AddressToOutputs", testAddressToManySetOpOutputs)
	t.Run("BlockToBlockByHashTransactions", testBlockToManySetOpBlockByHashTransactions)
	t.Run("ClaimToPublisherClaims", testClaimToManySetOpPublisherClaims)
	t.Run("InputToAddresses", testInputToManySetOpAddresses)
	t.Run("InputToSpentByInputOutputs", testInputToManySetOpSpentByInputOutputs)
	t.Run("OutputToAddresses", testOutputToManySetOpAddresses)
	t.Run("TransactionToTransactionByHashClaims", testTransactionToManySetOpTransactionByHashClaims)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AddressToInputAddressInputs", testAddressToManyRemoveOpInputAddressInputs)
	t.Run("AddressToInputs", testAddressToManyRemoveOpInputs)
	t.Run("AddressToOutputs", testAddressToManyRemoveOpOutputs)
	t.Run("BlockToBlockByHashTransactions", testBlockToManyRemoveOpBlockByHashTransactions)
	t.Run("ClaimToPublisherClaims", testClaimToManyRemoveOpPublisherClaims)
	t.Run("InputToAddresses", testInputToManyRemoveOpAddresses)
	t.Run("InputToSpentByInputOutputs", testInputToManyRemoveOpSpentByInputOutputs)
	t.Run("OutputToAddresses", testOutputToManyRemoveOpAddresses)
	t.Run("TransactionToTransactionByHashClaims", testTransactionToManyRemoveOpTransactionByHashClaims)
}

func TestReload(t *testing.T) {
	t.Run("Addresses", testAddressesReload)
	t.Run("Blocks", testBlocksReload)
	t.Run("Claims", testClaimsReload)
	t.Run("Inputs", testInputsReload)
	t.Run("Outputs", testOutputsReload)
	t.Run("Transactions", testTransactionsReload)
	t.Run("TransactionAddresses", testTransactionAddressesReload)
	t.Run("UnknownClaims", testUnknownClaimsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Addresses", testAddressesReloadAll)
	t.Run("Blocks", testBlocksReloadAll)
	t.Run("Claims", testClaimsReloadAll)
	t.Run("Inputs", testInputsReloadAll)
	t.Run("Outputs", testOutputsReloadAll)
	t.Run("Transactions", testTransactionsReloadAll)
	t.Run("TransactionAddresses", testTransactionAddressesReloadAll)
	t.Run("UnknownClaims", testUnknownClaimsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Addresses", testAddressesSelect)
	t.Run("Blocks", testBlocksSelect)
	t.Run("Claims", testClaimsSelect)
	t.Run("Inputs", testInputsSelect)
	t.Run("Outputs", testOutputsSelect)
	t.Run("Transactions", testTransactionsSelect)
	t.Run("TransactionAddresses", testTransactionAddressesSelect)
	t.Run("UnknownClaims", testUnknownClaimsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Addresses", testAddressesUpdate)
	t.Run("Blocks", testBlocksUpdate)
	t.Run("Claims", testClaimsUpdate)
	t.Run("Inputs", testInputsUpdate)
	t.Run("Outputs", testOutputsUpdate)
	t.Run("Transactions", testTransactionsUpdate)
	t.Run("TransactionAddresses", testTransactionAddressesUpdate)
	t.Run("UnknownClaims", testUnknownClaimsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Addresses", testAddressesSliceUpdateAll)
	t.Run("Blocks", testBlocksSliceUpdateAll)
	t.Run("Claims", testClaimsSliceUpdateAll)
	t.Run("Inputs", testInputsSliceUpdateAll)
	t.Run("Outputs", testOutputsSliceUpdateAll)
	t.Run("Transactions", testTransactionsSliceUpdateAll)
	t.Run("TransactionAddresses", testTransactionAddressesSliceUpdateAll)
	t.Run("UnknownClaims", testUnknownClaimsSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Addresses", testAddressesUpsert)
	t.Run("Blocks", testBlocksUpsert)
	t.Run("Claims", testClaimsUpsert)
	t.Run("Inputs", testInputsUpsert)
	t.Run("Outputs", testOutputsUpsert)
	t.Run("Transactions", testTransactionsUpsert)
	t.Run("TransactionAddresses", testTransactionAddressesUpsert)
	t.Run("UnknownClaims", testUnknownClaimsUpsert)
}
