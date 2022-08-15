package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/a-saber-abdelmosen/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRamdomEntry(t *testing.T) Entry {
	account := createRamdomAccount(t)
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(1, 100),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.Amount, args.Amount)
	require.Equal(t, entry.AccountID, account.ID)
	return entry
}

func createEntryForAccount(t *testing.T, accountID int64) Entry {
	args := CreateEntryParams{
		AccountID: accountID,
		Amount:    util.RandomInt(1, 100),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.Amount, args.Amount)
	require.Equal(t, entry.AccountID, accountID)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRamdomEntry(t)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRamdomEntry(t)

	args := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomInt(1, 100),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.Amount, args.Amount)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRamdomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRamdomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRamdomEntry(t)
	}

	args := ListEntriesParams{
		Offset: 5,
		Limit:  5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestListAcountEntries(t *testing.T) {
	account := createRamdomAccount(t)
	for i := 0; i < 10; i++ {
		createEntryForAccount(t, account.ID)
	}

	args := ListAccountEntriesParams{
		Offset:    5,
		Limit:     5,
		AccountID: account.ID,
	}

	entries, err := testQueries.ListAccountEntries(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, ent := range entries {
		require.NotEmpty(t, ent)
		require.Equal(t, ent.AccountID, account.ID)
	}
}
