package walletevent

import (
	"context"
	"testing"

	"github.com/ipfs-force-community/venus-gateway/testhelper"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/venus-auth/auth"
	"github.com/filecoin-project/venus-auth/cmd/jwtclient"
	_ "github.com/filecoin-project/venus/pkg/crypto/secp"
	sharedTypes "github.com/filecoin-project/venus/venus-shared/types"
	"github.com/ipfs-force-community/venus-gateway/types"
	"github.com/ipfs-force-community/venus-gateway/validator/mocks"
	"github.com/stretchr/testify/require"
)

func TestListenWalletEvent(t *testing.T) {
	walletAccount := "walletAccount"
	t.Run("correct", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		walletEvent := setupWalletEvent(t)
		client := setupClient(t, ctx, walletAccount, []string{"admin"}, walletEvent)
		go client.listenWalletEvent(ctx)
		client.walletEventClient.WaitReady(ctx)

		walletInfo, err := walletEvent.ListWalletInfoByWallet(ctx, walletAccount)
		require.NoError(t, err)
		require.Equal(t, walletInfo.Account, walletAccount)
		require.Equal(t, walletInfo.SupportAccounts, []string{"admin"})
	})

	t.Run("multiple listen", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		walletEvent := setupWalletEvent(t)
		client := setupClient(t, ctx, walletAccount, []string{}, walletEvent)
		go client.listenWalletEvent(ctx)
		client.walletEventClient.WaitReady(ctx)

		client2 := setupClient(t, ctx, walletAccount, []string{}, walletEvent)
		go client2.listenWalletEvent(ctx)
		client2.walletEventClient.WaitReady(ctx)

		walletInfo, err := walletEvent.ListWalletInfoByWallet(ctx, walletAccount)
		require.NoError(t, err)
		require.Len(t, walletInfo.ConnectStates, 2)
	})

	t.Run("wallet account not found", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		walletEvent := setupWalletEvent(t)
		//register
		client := setupClient(t, ctx, walletAccount, []string{"admin"}, walletEvent)
		err := client.walletEventClient.listenWalletRequestOnce(ctx)
		require.Contains(t, err.Error(), "unable to get account name in method ListenWalletEvent request")
	})
}

func TestSupportNewAccount(t *testing.T) {
	walletAccount := "walletAccount"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	walletEvent := setupWalletEvent(t)
	client := setupClient(t, ctx, walletAccount, []string{"admin"}, walletEvent)
	go client.listenWalletEvent(ctx)
	client.walletEventClient.WaitReady(ctx)

	err := client.supportNewAccount(ctx, "ac1")
	require.NoError(t, err)

	wallets, err := walletEvent.ListWalletInfo(ctx)
	require.NoError(t, err)
	require.Len(t, wallets, 1)
	require.Len(t, wallets[0].SupportAccounts, 2)
	require.Contains(t, wallets[0].SupportAccounts, "ac1")
	require.Contains(t, wallets[0].SupportAccounts, "admin")

	err = client.walletEventClient.SupportAccount(ctx, "fake_acc")
	require.EqualError(t, err, "unable to get account name in method SupportNewAccount request")

	ctx = jwtclient.CtxWithName(ctx, "fac_acc")
	err = client.walletEventClient.SupportAccount(ctx, "__")
	require.NoError(t, err)
}

func TestAddNewAddress(t *testing.T) {
	walletAccount := "walletAccount"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	walletEvent := setupWalletEvent(t)
	client := setupClient(t, ctx, walletAccount, []string{"admin"}, walletEvent)
	go client.listenWalletEvent(ctx)
	client.walletEventClient.WaitReady(ctx)

	addr1 := client.newkey()
	addr2 := client.newkey()

	err := client.walletEventClient.AddNewAddress(ctx, []address.Address{addr1})
	require.EqualError(t, err, "unable to get account name in method AddNewAddress request")

	ctx = jwtclient.CtxWithName(ctx, walletAccount)
	err = client.walletEventClient.AddNewAddress(ctx, []address.Address{addr1})
	require.NoError(t, err)

	ctx = jwtclient.CtxWithName(ctx, walletAccount)
	err = client.walletEventClient.AddNewAddress(ctx, []address.Address{addr1}) //allow dup add
	require.NoError(t, err)

	ctx = jwtclient.CtxWithName(ctx, walletAccount)
	err = client.walletEventClient.AddNewAddress(ctx, []address.Address{addr1, addr1, addr2})
	require.NoError(t, err)

	wallet, err := walletEvent.ListWalletInfoByWallet(ctx, walletAccount)
	require.NoError(t, err)
	require.Contains(t, wallet.ConnectStates[0].Addrs, addr2)
	require.Contains(t, wallet.ConnectStates[0].Addrs, addr1)
}

func TestRemoveNewAddressAndWalletHas(t *testing.T) {
	walletAccount := "walletAccount"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	walletEvent := setupWalletEvent(t)
	client := setupClient(t, ctx, walletAccount, []string{"admin"}, walletEvent)
	go client.listenWalletEvent(ctx)
	client.walletEventClient.WaitReady(ctx)

	addr1 := client.newkey()
	accCtx := jwtclient.CtxWithName(ctx, walletAccount)
	err := client.walletEventClient.AddNewAddress(accCtx, []address.Address{addr1})
	require.NoError(t, err)
	has, err := walletEvent.WalletHas(ctx, "admin", addr1)
	require.NoError(t, err)
	require.True(t, has)

	addr2 := client.newkey()
	err = client.walletEventClient.AddNewAddress(accCtx, []address.Address{addr2})
	require.NoError(t, err)
	has, err = walletEvent.WalletHas(ctx, "admin", addr2)
	require.NoError(t, err)
	require.True(t, has)

	has, err = walletEvent.WalletHas(ctx, "fak_acc", addr1)
	require.NoError(t, err)
	require.False(t, has)

	err = client.walletEventClient.RemoveAddress(ctx, []address.Address{addr1})
	require.EqualError(t, err, "unable to get account name in method RemoveAddress request")

	err = client.walletEventClient.RemoveAddress(accCtx, []address.Address{addr1})
	require.NoError(t, err)

	has, err = walletEvent.WalletHas(ctx, "admin", addr1)
	require.NoError(t, err)
	require.False(t, has)

	wallet, err := walletEvent.ListWalletInfoByWallet(ctx, walletAccount)
	require.NoError(t, err)
	require.Len(t, wallet.ConnectStates[0].Addrs, 2)
	require.Contains(t, wallet.ConnectStates[0].Addrs, addr2)
	require.NotContains(t, wallet.ConnectStates[0].Addrs, addr1)

	err = client.walletEventClient.RemoveAddress(accCtx, []address.Address{addr2})
	require.NoError(t, err)
	wallet, err = walletEvent.ListWalletInfoByWallet(ctx, walletAccount)
	require.NoError(t, err)
	require.Len(t, wallet.ConnectStates[0].Addrs, 1)

	has, err = walletEvent.WalletHas(ctx, "admin", addr1)
	require.NoError(t, err)
	require.False(t, has)
}

func TestWalletSign(t *testing.T) {
	walletAccount := "walletAccount"
	//register
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	walletEvent := setupWalletEvent(t)
	client := setupClient(t, ctx, walletAccount, []string{}, walletEvent)
	go client.listenWalletEvent(ctx)
	client.walletEventClient.WaitReady(ctx)

	addrs, err := client.wallet.WalletList(ctx)
	for _, addr := range addrs {
		require.NoError(t, err)
		_, err = walletEvent.WalletSign(ctx, "admin", addr, []byte{1, 2, 3}, sharedTypes.MsgMeta{
			Type:  sharedTypes.MTUnknown,
			Extra: nil,
		})

		require.Error(t, err)

		err = client.supportNewAccount(ctx, "admin")
		require.NoError(t, err)

		_, err = walletEvent.WalletSign(ctx, "admin", addr, []byte{1, 2, 3}, sharedTypes.MsgMeta{
			Type:  sharedTypes.MTUnknown,
			Extra: nil,
		})
		require.NoError(t, err)

		client.wallet.SetFail(ctx, true)
		_, err = walletEvent.WalletSign(ctx, "admin", addr, []byte{1, 2, 3}, sharedTypes.MsgMeta{
			Type:  sharedTypes.MTUnknown,
			Extra: nil,
		})
		require.EqualError(t, err, "mock error")
	}

}

func setupWalletEvent(t *testing.T) *WalletEventStream {
	authClient := mocks.NewMockAuthClient()
	authClient.AddMockUser(&auth.OutputUser{
		Id:         "id",
		Name:       "admin",
		SourceType: 0,
		Comment:    "",
		State:      0,
		CreateTime: 0,
		UpdateTime: 0,
	})

	ctx := context.Background()
	return NewWalletEventStream(ctx, authClient, types.DefaultConfig())
}

func setupClient(t *testing.T, ctx context.Context, walletAccount string, supportAccounts []string, event *WalletEventStream) *mockClient {
	wallet := testhelper.NewMemWallet()
	_, err := wallet.AddKey(context.Background())
	require.NoError(t, err)
	walletEventClient := NewWalletEventClient(ctx, wallet, event, logging.Logger("test").With(), supportAccounts)
	return &mockClient{
		t:                 t,
		walletEventClient: walletEventClient,
		wallet:            wallet,
		walletAccount:     walletAccount,
	}
}

type mockClient struct {
	t                 *testing.T
	wallet            *testhelper.MemWallet
	walletEventClient *WalletEventClient
	walletAccount     string
}

func (m *mockClient) newkey() address.Address {
	addr, err := m.wallet.AddKey(context.Background())
	require.NoError(m.t, err)
	return addr
}

func (m *mockClient) listenWalletEvent(ctx context.Context) {
	ctx = jwtclient.CtxWithTokenLocation(ctx, "127.1.1.1")
	ctx = jwtclient.CtxWithName(ctx, m.walletAccount)
	m.walletEventClient.ListenWalletRequest(ctx)
}

func (m *mockClient) supportNewAccount(ctx context.Context, account string) error {
	ctx = jwtclient.CtxWithTokenLocation(ctx, "127.1.1.1")
	ctx = jwtclient.CtxWithName(ctx, m.walletAccount)
	return m.walletEventClient.SupportAccount(ctx, account)
}
