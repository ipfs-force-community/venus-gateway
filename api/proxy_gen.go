// Code generated by github.com/filecoin-project/venus-market/tools/api_gen. DO NOT EDIT.

package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	proof5 "github.com/filecoin-project/specs-actors/v5/actors/runtime/proof"
	"github.com/filecoin-project/specs-storage/storage"
	sharedTypes "github.com/filecoin-project/venus/venus-shared/types"
	"github.com/google/uuid"
	types2 "github.com/ipfs-force-community/venus-common-utils/types"
	"github.com/ipfs-force-community/venus-gateway/marketevent"
	"github.com/ipfs-force-community/venus-gateway/proofevent"
	"github.com/ipfs-force-community/venus-gateway/types"
	"github.com/ipfs-force-community/venus-gateway/walletevent"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"
)

type GatewayFullNodeStruct struct {
	IProofEventStruct

	IWalletEventStruct

	IMarketEventStruct

	Internal struct {
	}
}

type GatewayFullNodeStub struct {
	IProofEventStub

	IWalletEventStub

	IMarketEventStub
}

type IMarketEventStruct struct {
	Internal struct {
		IsUnsealed func(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 storage.SectorRef, p4 types2.PaddedByteIndex, p5 abi.PaddedPieceSize) (bool, error) `perm:"admin"`

		ListMarketConnectionsState func(p0 context.Context) ([]marketevent.MarketConnectionState, error) `perm:"admin"`

		ListenMarketEvent func(p0 context.Context, p1 *marketevent.MarketRegisterPolicy) (<-chan *types.RequestEvent, error) `perm:"read"`

		ResponseMarketEvent func(p0 context.Context, p1 *types.ResponseEvent) error `perm:"read"`

		SectorsUnsealPiece func(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 storage.SectorRef, p4 types2.PaddedByteIndex, p5 abi.PaddedPieceSize, p6 string) error `perm:"admin"`
	}
}

type IMarketEventStub struct {
}

type IProofEventStruct struct {
	Internal struct {
		ComputeProof func(p0 context.Context, p1 address.Address, p2 []proof5.SectorInfo, p3 abi.PoStRandomness) ([]proof5.PoStProof, error) `perm:"admin"`

		ListConnectedMiners func(p0 context.Context) ([]address.Address, error) `perm:"admin"`

		ListMinerConnection func(p0 context.Context, p1 address.Address) (*proofevent.MinerState, error) `perm:"admin"`

		ListenProofEvent func(p0 context.Context, p1 *proofevent.ProofRegisterPolicy) (<-chan *types.RequestEvent, error) `perm:"read"`

		ResponseProofEvent func(p0 context.Context, p1 *types.ResponseEvent) error `perm:"read"`
	}
}

type IProofEventStub struct {
}

type IWalletEventStruct struct {
	Internal struct {
		AddNewAddress func(p0 context.Context, p1 uuid.UUID, p2 []address.Address) error `perm:"read"`

		ListWalletInfo func(p0 context.Context) ([]*walletevent.WalletDetail, error) `perm:"admin"`

		ListWalletInfoByWallet func(p0 context.Context, p1 string) (*walletevent.WalletDetail, error) `perm:"admin"`

		ListenWalletEvent func(p0 context.Context, p1 *walletevent.WalletRegisterPolicy) (<-chan *types.RequestEvent, error) `perm:"read"`

		RemoveAddress func(p0 context.Context, p1 uuid.UUID, p2 []address.Address) error `perm:"read"`

		ResponseWalletEvent func(p0 context.Context, p1 *types.ResponseEvent) error `perm:"read"`

		SupportNewAccount func(p0 context.Context, p1 uuid.UUID, p2 string) error `perm:"read"`

		WalletHas func(p0 context.Context, p1 string, p2 address.Address) (bool, error) `perm:"admin"`

		WalletSign func(p0 context.Context, p1 string, p2 address.Address, p3 []byte, p4 sharedTypes.MsgMeta) (*crypto.Signature, error) `perm:"admin"`
	}
}

type IWalletEventStub struct {
}

func (s *IMarketEventStruct) IsUnsealed(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 storage.SectorRef, p4 types2.PaddedByteIndex, p5 abi.PaddedPieceSize) (bool, error) {
	return s.Internal.IsUnsealed(p0, p1, p2, p3, p4, p5)
}

func (s *IMarketEventStub) IsUnsealed(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 storage.SectorRef, p4 types2.PaddedByteIndex, p5 abi.PaddedPieceSize) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *IMarketEventStruct) ListMarketConnectionsState(p0 context.Context) ([]marketevent.MarketConnectionState, error) {
	return s.Internal.ListMarketConnectionsState(p0)
}

func (s *IMarketEventStub) ListMarketConnectionsState(p0 context.Context) ([]marketevent.MarketConnectionState, error) {
	return *new([]marketevent.MarketConnectionState), xerrors.New("method not supported")
}

func (s *IMarketEventStruct) ListenMarketEvent(p0 context.Context, p1 *marketevent.MarketRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return s.Internal.ListenMarketEvent(p0, p1)
}

func (s *IMarketEventStub) ListenMarketEvent(p0 context.Context, p1 *marketevent.MarketRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return nil, xerrors.New("method not supported")
}

func (s *IMarketEventStruct) ResponseMarketEvent(p0 context.Context, p1 *types.ResponseEvent) error {
	return s.Internal.ResponseMarketEvent(p0, p1)
}

func (s *IMarketEventStub) ResponseMarketEvent(p0 context.Context, p1 *types.ResponseEvent) error {
	return xerrors.New("method not supported")
}

func (s *IMarketEventStruct) SectorsUnsealPiece(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 storage.SectorRef, p4 types2.PaddedByteIndex, p5 abi.PaddedPieceSize, p6 string) error {
	return s.Internal.SectorsUnsealPiece(p0, p1, p2, p3, p4, p5, p6)
}

func (s *IMarketEventStub) SectorsUnsealPiece(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 storage.SectorRef, p4 types2.PaddedByteIndex, p5 abi.PaddedPieceSize, p6 string) error {
	return xerrors.New("method not supported")
}

func (s *IProofEventStruct) ComputeProof(p0 context.Context, p1 address.Address, p2 []proof5.SectorInfo, p3 abi.PoStRandomness) ([]proof5.PoStProof, error) {
	return s.Internal.ComputeProof(p0, p1, p2, p3)
}

func (s *IProofEventStub) ComputeProof(p0 context.Context, p1 address.Address, p2 []proof5.SectorInfo, p3 abi.PoStRandomness) ([]proof5.PoStProof, error) {
	return *new([]proof5.PoStProof), xerrors.New("method not supported")
}

func (s *IProofEventStruct) ListConnectedMiners(p0 context.Context) ([]address.Address, error) {
	return s.Internal.ListConnectedMiners(p0)
}

func (s *IProofEventStub) ListConnectedMiners(p0 context.Context) ([]address.Address, error) {
	return *new([]address.Address), xerrors.New("method not supported")
}

func (s *IProofEventStruct) ListMinerConnection(p0 context.Context, p1 address.Address) (*proofevent.MinerState, error) {
	return s.Internal.ListMinerConnection(p0, p1)
}

func (s *IProofEventStub) ListMinerConnection(p0 context.Context, p1 address.Address) (*proofevent.MinerState, error) {
	return nil, xerrors.New("method not supported")
}

func (s *IProofEventStruct) ListenProofEvent(p0 context.Context, p1 *proofevent.ProofRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return s.Internal.ListenProofEvent(p0, p1)
}

func (s *IProofEventStub) ListenProofEvent(p0 context.Context, p1 *proofevent.ProofRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return nil, xerrors.New("method not supported")
}

func (s *IProofEventStruct) ResponseProofEvent(p0 context.Context, p1 *types.ResponseEvent) error {
	return s.Internal.ResponseProofEvent(p0, p1)
}

func (s *IProofEventStub) ResponseProofEvent(p0 context.Context, p1 *types.ResponseEvent) error {
	return xerrors.New("method not supported")
}

func (s *IWalletEventStruct) AddNewAddress(p0 context.Context, p1 uuid.UUID, p2 []address.Address) error {
	return s.Internal.AddNewAddress(p0, p1, p2)
}

func (s *IWalletEventStub) AddNewAddress(p0 context.Context, p1 uuid.UUID, p2 []address.Address) error {
	return xerrors.New("method not supported")
}

func (s *IWalletEventStruct) ListWalletInfo(p0 context.Context) ([]*walletevent.WalletDetail, error) {
	return s.Internal.ListWalletInfo(p0)
}

func (s *IWalletEventStub) ListWalletInfo(p0 context.Context) ([]*walletevent.WalletDetail, error) {
	return *new([]*walletevent.WalletDetail), xerrors.New("method not supported")
}

func (s *IWalletEventStruct) ListWalletInfoByWallet(p0 context.Context, p1 string) (*walletevent.WalletDetail, error) {
	return s.Internal.ListWalletInfoByWallet(p0, p1)
}

func (s *IWalletEventStub) ListWalletInfoByWallet(p0 context.Context, p1 string) (*walletevent.WalletDetail, error) {
	return nil, xerrors.New("method not supported")
}

func (s *IWalletEventStruct) ListenWalletEvent(p0 context.Context, p1 *walletevent.WalletRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return s.Internal.ListenWalletEvent(p0, p1)
}

func (s *IWalletEventStub) ListenWalletEvent(p0 context.Context, p1 *walletevent.WalletRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return nil, xerrors.New("method not supported")
}

func (s *IWalletEventStruct) RemoveAddress(p0 context.Context, p1 uuid.UUID, p2 []address.Address) error {
	return s.Internal.RemoveAddress(p0, p1, p2)
}

func (s *IWalletEventStub) RemoveAddress(p0 context.Context, p1 uuid.UUID, p2 []address.Address) error {
	return xerrors.New("method not supported")
}

func (s *IWalletEventStruct) ResponseWalletEvent(p0 context.Context, p1 *types.ResponseEvent) error {
	return s.Internal.ResponseWalletEvent(p0, p1)
}

func (s *IWalletEventStub) ResponseWalletEvent(p0 context.Context, p1 *types.ResponseEvent) error {
	return xerrors.New("method not supported")
}

func (s *IWalletEventStruct) SupportNewAccount(p0 context.Context, p1 uuid.UUID, p2 string) error {
	return s.Internal.SupportNewAccount(p0, p1, p2)
}

func (s *IWalletEventStub) SupportNewAccount(p0 context.Context, p1 uuid.UUID, p2 string) error {
	return xerrors.New("method not supported")
}

func (s *IWalletEventStruct) WalletHas(p0 context.Context, p1 string, p2 address.Address) (bool, error) {
	return s.Internal.WalletHas(p0, p1, p2)
}

func (s *IWalletEventStub) WalletHas(p0 context.Context, p1 string, p2 address.Address) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *IWalletEventStruct) WalletSign(p0 context.Context, p1 string, p2 address.Address, p3 []byte, p4 sharedTypes.MsgMeta) (*crypto.Signature, error) {
	return s.Internal.WalletSign(p0, p1, p2, p3, p4)
}

func (s *IWalletEventStub) WalletSign(p0 context.Context, p1 string, p2 address.Address, p3 []byte, p4 sharedTypes.MsgMeta) (*crypto.Signature, error) {
	return nil, xerrors.New("method not supported")
}

var _ GatewayFullNode = new(GatewayFullNodeStruct)
var _ IMarketEvent = new(IMarketEventStruct)
var _ IProofEvent = new(IProofEventStruct)
var _ IWalletEvent = new(IWalletEventStruct)
