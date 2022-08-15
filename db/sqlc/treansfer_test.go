package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/a-saber-abdelmosen/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRamdomAccount(t)
	account2 := createRamdomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(1, 100),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)
	require.Equal(t, transfer.Amount, args.Amount)

	return transfer
}

func createRandomTransferForAccounts(
	t *testing.T,
	fromAccountId,
	toAccountId int64) Transfer {

	args := CreateTransferParams{
		FromAccountID: fromAccountId,
		ToAccountID:   toAccountId,
		Amount:        util.RandomInt(1, 100),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)
	require.Equal(t, transfer.Amount, args.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTranfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.CreatedAt, transfer2.CreatedAt)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	args := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomInt(1, 100),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer2.Amount, args.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	args := ListTransfersParams{
		Offset: 5,
		Limit:  5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, entry := range transfers {
		require.NotEmpty(t, entry)
	}
}

func TestListAccountsTransfer(t *testing.T) {
	account1 := createRamdomAccount(t)
	account2 := createRamdomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransferForAccounts(t, account1.ID, account2.ID)
	}

	args := ListTransfersBetweenAccountsParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Offset:        5,
		Limit:         5,
	}

	transfers, err := testQueries.ListTransfersBetweenAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, entry := range transfers {
		require.NotEmpty(t, entry)
	}
}
