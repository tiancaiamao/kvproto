// Code generated by Kitex v0.1.2. DO NOT EDIT.

package tikv

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
	"github.com/pingcap/kvproto/proto/kitex_gen/coprocessor"
	"github.com/pingcap/kvproto/proto/kitex_gen/kvrpcpb"
	"github.com/pingcap/kvproto/proto/kitex_gen/mpp"
	"github.com/pingcap/kvproto/proto/kitex_gen/raft_serverpb"
	"github.com/pingcap/kvproto/proto/kitex_gen/tikvpb"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	KvGet(ctx context.Context, Req *kvrpcpb.GetRequest, callOptions ...callopt.Option) (r *kvrpcpb.GetResponse, err error)
	KvScan(ctx context.Context, Req *kvrpcpb.ScanRequest, callOptions ...callopt.Option) (r *kvrpcpb.ScanResponse, err error)
	KvPrewrite(ctx context.Context, Req *kvrpcpb.PrewriteRequest, callOptions ...callopt.Option) (r *kvrpcpb.PrewriteResponse, err error)
	KvPessimisticLock(ctx context.Context, Req *kvrpcpb.PessimisticLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.PessimisticLockResponse, err error)
	KVPessimisticRollback(ctx context.Context, Req *kvrpcpb.PessimisticRollbackRequest, callOptions ...callopt.Option) (r *kvrpcpb.PessimisticRollbackResponse, err error)
	KvTxnHeartBeat(ctx context.Context, Req *kvrpcpb.TxnHeartBeatRequest, callOptions ...callopt.Option) (r *kvrpcpb.TxnHeartBeatResponse, err error)
	KvCheckTxnStatus(ctx context.Context, Req *kvrpcpb.CheckTxnStatusRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckTxnStatusResponse, err error)
	KvCheckSecondaryLocks(ctx context.Context, Req *kvrpcpb.CheckSecondaryLocksRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckSecondaryLocksResponse, err error)
	KvCommit(ctx context.Context, Req *kvrpcpb.CommitRequest, callOptions ...callopt.Option) (r *kvrpcpb.CommitResponse, err error)
	KvImport(ctx context.Context, Req *kvrpcpb.ImportRequest, callOptions ...callopt.Option) (r *kvrpcpb.ImportResponse, err error)
	KvCleanup(ctx context.Context, Req *kvrpcpb.CleanupRequest, callOptions ...callopt.Option) (r *kvrpcpb.CleanupResponse, err error)
	KvBatchGet(ctx context.Context, Req *kvrpcpb.BatchGetRequest, callOptions ...callopt.Option) (r *kvrpcpb.BatchGetResponse, err error)
	KvBatchRollback(ctx context.Context, Req *kvrpcpb.BatchRollbackRequest, callOptions ...callopt.Option) (r *kvrpcpb.BatchRollbackResponse, err error)
	KvScanLock(ctx context.Context, Req *kvrpcpb.ScanLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.ScanLockResponse, err error)
	KvResolveLock(ctx context.Context, Req *kvrpcpb.ResolveLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.ResolveLockResponse, err error)
	KvGC(ctx context.Context, Req *kvrpcpb.GCRequest, callOptions ...callopt.Option) (r *kvrpcpb.GCResponse, err error)
	KvDeleteRange(ctx context.Context, Req *kvrpcpb.DeleteRangeRequest, callOptions ...callopt.Option) (r *kvrpcpb.DeleteRangeResponse, err error)
	RawGet(ctx context.Context, Req *kvrpcpb.RawGetRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawGetResponse, err error)
	RawBatchGet(ctx context.Context, Req *kvrpcpb.RawBatchGetRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchGetResponse, err error)
	RawPut(ctx context.Context, Req *kvrpcpb.RawPutRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawPutResponse, err error)
	RawBatchPut(ctx context.Context, Req *kvrpcpb.RawBatchPutRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchPutResponse, err error)
	RawDelete(ctx context.Context, Req *kvrpcpb.RawDeleteRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawDeleteResponse, err error)
	RawBatchDelete(ctx context.Context, Req *kvrpcpb.RawBatchDeleteRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchDeleteResponse, err error)
	RawScan(ctx context.Context, Req *kvrpcpb.RawScanRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawScanResponse, err error)
	RawDeleteRange(ctx context.Context, Req *kvrpcpb.RawDeleteRangeRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawDeleteRangeResponse, err error)
	RawBatchScan(ctx context.Context, Req *kvrpcpb.RawBatchScanRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchScanResponse, err error)
	RawGetKeyTTL(ctx context.Context, Req *kvrpcpb.RawGetKeyTTLRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawGetKeyTTLResponse, err error)
	RawCompareAndSwap(ctx context.Context, Req *kvrpcpb.RawCASRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawCASResponse, err error)
	RawChecksum(ctx context.Context, Req *kvrpcpb.RawChecksumRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawChecksumResponse, err error)
	UnsafeDestroyRange(ctx context.Context, Req *kvrpcpb.UnsafeDestroyRangeRequest, callOptions ...callopt.Option) (r *kvrpcpb.UnsafeDestroyRangeResponse, err error)
	RegisterLockObserver(ctx context.Context, Req *kvrpcpb.RegisterLockObserverRequest, callOptions ...callopt.Option) (r *kvrpcpb.RegisterLockObserverResponse, err error)
	CheckLockObserver(ctx context.Context, Req *kvrpcpb.CheckLockObserverRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckLockObserverResponse, err error)
	RemoveLockObserver(ctx context.Context, Req *kvrpcpb.RemoveLockObserverRequest, callOptions ...callopt.Option) (r *kvrpcpb.RemoveLockObserverResponse, err error)
	PhysicalScanLock(ctx context.Context, Req *kvrpcpb.PhysicalScanLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.PhysicalScanLockResponse, err error)
	Coprocessor(ctx context.Context, Req *coprocessor.Request, callOptions ...callopt.Option) (r *coprocessor.Response, err error)
	CoprocessorStream(ctx context.Context, Req *coprocessor.Request, callOptions ...callopt.Option) (stream Tikv_CoprocessorStreamClient, err error)
	BatchCoprocessor(ctx context.Context, Req *coprocessor.BatchRequest, callOptions ...callopt.Option) (stream Tikv_BatchCoprocessorClient, err error)
	RawCoprocessor(ctx context.Context, Req *kvrpcpb.RawCoprocessorRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawCoprocessorResponse, err error)
	Raft(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_RaftClient, err error)
	BatchRaft(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_BatchRaftClient, err error)
	Snapshot(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_SnapshotClient, err error)
	SplitRegion(ctx context.Context, Req *kvrpcpb.SplitRegionRequest, callOptions ...callopt.Option) (r *kvrpcpb.SplitRegionResponse, err error)
	ReadIndex(ctx context.Context, Req *kvrpcpb.ReadIndexRequest, callOptions ...callopt.Option) (r *kvrpcpb.ReadIndexResponse, err error)
	MvccGetByKey(ctx context.Context, Req *kvrpcpb.MvccGetByKeyRequest, callOptions ...callopt.Option) (r *kvrpcpb.MvccGetByKeyResponse, err error)
	MvccGetByStartTs(ctx context.Context, Req *kvrpcpb.MvccGetByStartTsRequest, callOptions ...callopt.Option) (r *kvrpcpb.MvccGetByStartTsResponse, err error)
	BatchCommands(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_BatchCommandsClient, err error)
	DispatchMPPTask(ctx context.Context, Req *mpp.DispatchTaskRequest, callOptions ...callopt.Option) (r *mpp.DispatchTaskResponse, err error)
	CancelMPPTask(ctx context.Context, Req *mpp.CancelTaskRequest, callOptions ...callopt.Option) (r *mpp.CancelTaskResponse, err error)
	EstablishMPPConnection(ctx context.Context, Req *mpp.EstablishMPPConnectionRequest, callOptions ...callopt.Option) (stream Tikv_EstablishMPPConnectionClient, err error)
	IsAlive(ctx context.Context, Req *mpp.IsAliveRequest, callOptions ...callopt.Option) (r *mpp.IsAliveResponse, err error)
	CheckLeader(ctx context.Context, Req *kvrpcpb.CheckLeaderRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckLeaderResponse, err error)
	GetStoreSafeTS(ctx context.Context, Req *kvrpcpb.StoreSafeTSRequest, callOptions ...callopt.Option) (r *kvrpcpb.StoreSafeTSResponse, err error)
	GetLockWaitInfo(ctx context.Context, Req *kvrpcpb.GetLockWaitInfoRequest, callOptions ...callopt.Option) (r *kvrpcpb.GetLockWaitInfoResponse, err error)
}

type Tikv_CoprocessorStreamClient interface {
	streaming.Stream
	Recv() (*coprocessor.Response, error)
}

type Tikv_BatchCoprocessorClient interface {
	streaming.Stream
	Recv() (*coprocessor.BatchResponse, error)
}

type Tikv_RaftClient interface {
	streaming.Stream
	Send(*raft_serverpb.RaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
}

type Tikv_BatchRaftClient interface {
	streaming.Stream
	Send(*tikvpb.BatchRaftMessage) error
	CloseAndRecv() (*raft_serverpb.Done, error)
}

type Tikv_SnapshotClient interface {
	streaming.Stream
	Send(*raft_serverpb.SnapshotChunk) error
	CloseAndRecv() (*raft_serverpb.Done, error)
}

type Tikv_BatchCommandsClient interface {
	streaming.Stream
	Send(*tikvpb.BatchCommandsRequest) error
	Recv() (*tikvpb.BatchCommandsResponse, error)
}

type Tikv_EstablishMPPConnectionClient interface {
	streaming.Stream
	Recv() (*mpp.MPPDataPacket, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTikvClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTikvClient struct {
	*kClient
}

func (p *kTikvClient) KvGet(ctx context.Context, Req *kvrpcpb.GetRequest, callOptions ...callopt.Option) (r *kvrpcpb.GetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvGet(ctx, Req)
}

func (p *kTikvClient) KvScan(ctx context.Context, Req *kvrpcpb.ScanRequest, callOptions ...callopt.Option) (r *kvrpcpb.ScanResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvScan(ctx, Req)
}

func (p *kTikvClient) KvPrewrite(ctx context.Context, Req *kvrpcpb.PrewriteRequest, callOptions ...callopt.Option) (r *kvrpcpb.PrewriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvPrewrite(ctx, Req)
}

func (p *kTikvClient) KvPessimisticLock(ctx context.Context, Req *kvrpcpb.PessimisticLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.PessimisticLockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvPessimisticLock(ctx, Req)
}

func (p *kTikvClient) KVPessimisticRollback(ctx context.Context, Req *kvrpcpb.PessimisticRollbackRequest, callOptions ...callopt.Option) (r *kvrpcpb.PessimisticRollbackResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KVPessimisticRollback(ctx, Req)
}

func (p *kTikvClient) KvTxnHeartBeat(ctx context.Context, Req *kvrpcpb.TxnHeartBeatRequest, callOptions ...callopt.Option) (r *kvrpcpb.TxnHeartBeatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvTxnHeartBeat(ctx, Req)
}

func (p *kTikvClient) KvCheckTxnStatus(ctx context.Context, Req *kvrpcpb.CheckTxnStatusRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckTxnStatusResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvCheckTxnStatus(ctx, Req)
}

func (p *kTikvClient) KvCheckSecondaryLocks(ctx context.Context, Req *kvrpcpb.CheckSecondaryLocksRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckSecondaryLocksResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvCheckSecondaryLocks(ctx, Req)
}

func (p *kTikvClient) KvCommit(ctx context.Context, Req *kvrpcpb.CommitRequest, callOptions ...callopt.Option) (r *kvrpcpb.CommitResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvCommit(ctx, Req)
}

func (p *kTikvClient) KvImport(ctx context.Context, Req *kvrpcpb.ImportRequest, callOptions ...callopt.Option) (r *kvrpcpb.ImportResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvImport(ctx, Req)
}

func (p *kTikvClient) KvCleanup(ctx context.Context, Req *kvrpcpb.CleanupRequest, callOptions ...callopt.Option) (r *kvrpcpb.CleanupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvCleanup(ctx, Req)
}

func (p *kTikvClient) KvBatchGet(ctx context.Context, Req *kvrpcpb.BatchGetRequest, callOptions ...callopt.Option) (r *kvrpcpb.BatchGetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvBatchGet(ctx, Req)
}

func (p *kTikvClient) KvBatchRollback(ctx context.Context, Req *kvrpcpb.BatchRollbackRequest, callOptions ...callopt.Option) (r *kvrpcpb.BatchRollbackResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvBatchRollback(ctx, Req)
}

func (p *kTikvClient) KvScanLock(ctx context.Context, Req *kvrpcpb.ScanLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.ScanLockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvScanLock(ctx, Req)
}

func (p *kTikvClient) KvResolveLock(ctx context.Context, Req *kvrpcpb.ResolveLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.ResolveLockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvResolveLock(ctx, Req)
}

func (p *kTikvClient) KvGC(ctx context.Context, Req *kvrpcpb.GCRequest, callOptions ...callopt.Option) (r *kvrpcpb.GCResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvGC(ctx, Req)
}

func (p *kTikvClient) KvDeleteRange(ctx context.Context, Req *kvrpcpb.DeleteRangeRequest, callOptions ...callopt.Option) (r *kvrpcpb.DeleteRangeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.KvDeleteRange(ctx, Req)
}

func (p *kTikvClient) RawGet(ctx context.Context, Req *kvrpcpb.RawGetRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawGetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawGet(ctx, Req)
}

func (p *kTikvClient) RawBatchGet(ctx context.Context, Req *kvrpcpb.RawBatchGetRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchGetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawBatchGet(ctx, Req)
}

func (p *kTikvClient) RawPut(ctx context.Context, Req *kvrpcpb.RawPutRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawPutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawPut(ctx, Req)
}

func (p *kTikvClient) RawBatchPut(ctx context.Context, Req *kvrpcpb.RawBatchPutRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchPutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawBatchPut(ctx, Req)
}

func (p *kTikvClient) RawDelete(ctx context.Context, Req *kvrpcpb.RawDeleteRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawDeleteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawDelete(ctx, Req)
}

func (p *kTikvClient) RawBatchDelete(ctx context.Context, Req *kvrpcpb.RawBatchDeleteRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchDeleteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawBatchDelete(ctx, Req)
}

func (p *kTikvClient) RawScan(ctx context.Context, Req *kvrpcpb.RawScanRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawScanResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawScan(ctx, Req)
}

func (p *kTikvClient) RawDeleteRange(ctx context.Context, Req *kvrpcpb.RawDeleteRangeRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawDeleteRangeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawDeleteRange(ctx, Req)
}

func (p *kTikvClient) RawBatchScan(ctx context.Context, Req *kvrpcpb.RawBatchScanRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawBatchScanResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawBatchScan(ctx, Req)
}

func (p *kTikvClient) RawGetKeyTTL(ctx context.Context, Req *kvrpcpb.RawGetKeyTTLRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawGetKeyTTLResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawGetKeyTTL(ctx, Req)
}

func (p *kTikvClient) RawCompareAndSwap(ctx context.Context, Req *kvrpcpb.RawCASRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawCASResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawCompareAndSwap(ctx, Req)
}

func (p *kTikvClient) RawChecksum(ctx context.Context, Req *kvrpcpb.RawChecksumRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawChecksumResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawChecksum(ctx, Req)
}

func (p *kTikvClient) UnsafeDestroyRange(ctx context.Context, Req *kvrpcpb.UnsafeDestroyRangeRequest, callOptions ...callopt.Option) (r *kvrpcpb.UnsafeDestroyRangeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnsafeDestroyRange(ctx, Req)
}

func (p *kTikvClient) RegisterLockObserver(ctx context.Context, Req *kvrpcpb.RegisterLockObserverRequest, callOptions ...callopt.Option) (r *kvrpcpb.RegisterLockObserverResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterLockObserver(ctx, Req)
}

func (p *kTikvClient) CheckLockObserver(ctx context.Context, Req *kvrpcpb.CheckLockObserverRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckLockObserverResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckLockObserver(ctx, Req)
}

func (p *kTikvClient) RemoveLockObserver(ctx context.Context, Req *kvrpcpb.RemoveLockObserverRequest, callOptions ...callopt.Option) (r *kvrpcpb.RemoveLockObserverResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveLockObserver(ctx, Req)
}

func (p *kTikvClient) PhysicalScanLock(ctx context.Context, Req *kvrpcpb.PhysicalScanLockRequest, callOptions ...callopt.Option) (r *kvrpcpb.PhysicalScanLockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhysicalScanLock(ctx, Req)
}

func (p *kTikvClient) Coprocessor(ctx context.Context, Req *coprocessor.Request, callOptions ...callopt.Option) (r *coprocessor.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Coprocessor(ctx, Req)
}

func (p *kTikvClient) CoprocessorStream(ctx context.Context, Req *coprocessor.Request, callOptions ...callopt.Option) (stream Tikv_CoprocessorStreamClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CoprocessorStream(ctx, Req)
}

func (p *kTikvClient) BatchCoprocessor(ctx context.Context, Req *coprocessor.BatchRequest, callOptions ...callopt.Option) (stream Tikv_BatchCoprocessorClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchCoprocessor(ctx, Req)
}

func (p *kTikvClient) RawCoprocessor(ctx context.Context, Req *kvrpcpb.RawCoprocessorRequest, callOptions ...callopt.Option) (r *kvrpcpb.RawCoprocessorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RawCoprocessor(ctx, Req)
}

func (p *kTikvClient) Raft(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_RaftClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Raft(ctx)
}

func (p *kTikvClient) BatchRaft(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_BatchRaftClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchRaft(ctx)
}

func (p *kTikvClient) Snapshot(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_SnapshotClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Snapshot(ctx)
}

func (p *kTikvClient) SplitRegion(ctx context.Context, Req *kvrpcpb.SplitRegionRequest, callOptions ...callopt.Option) (r *kvrpcpb.SplitRegionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SplitRegion(ctx, Req)
}

func (p *kTikvClient) ReadIndex(ctx context.Context, Req *kvrpcpb.ReadIndexRequest, callOptions ...callopt.Option) (r *kvrpcpb.ReadIndexResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadIndex(ctx, Req)
}

func (p *kTikvClient) MvccGetByKey(ctx context.Context, Req *kvrpcpb.MvccGetByKeyRequest, callOptions ...callopt.Option) (r *kvrpcpb.MvccGetByKeyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MvccGetByKey(ctx, Req)
}

func (p *kTikvClient) MvccGetByStartTs(ctx context.Context, Req *kvrpcpb.MvccGetByStartTsRequest, callOptions ...callopt.Option) (r *kvrpcpb.MvccGetByStartTsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MvccGetByStartTs(ctx, Req)
}

func (p *kTikvClient) BatchCommands(ctx context.Context, callOptions ...callopt.Option) (stream Tikv_BatchCommandsClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchCommands(ctx)
}

func (p *kTikvClient) DispatchMPPTask(ctx context.Context, Req *mpp.DispatchTaskRequest, callOptions ...callopt.Option) (r *mpp.DispatchTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DispatchMPPTask(ctx, Req)
}

func (p *kTikvClient) CancelMPPTask(ctx context.Context, Req *mpp.CancelTaskRequest, callOptions ...callopt.Option) (r *mpp.CancelTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelMPPTask(ctx, Req)
}

func (p *kTikvClient) EstablishMPPConnection(ctx context.Context, Req *mpp.EstablishMPPConnectionRequest, callOptions ...callopt.Option) (stream Tikv_EstablishMPPConnectionClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EstablishMPPConnection(ctx, Req)
}

func (p *kTikvClient) IsAlive(ctx context.Context, Req *mpp.IsAliveRequest, callOptions ...callopt.Option) (r *mpp.IsAliveResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsAlive(ctx, Req)
}

func (p *kTikvClient) CheckLeader(ctx context.Context, Req *kvrpcpb.CheckLeaderRequest, callOptions ...callopt.Option) (r *kvrpcpb.CheckLeaderResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckLeader(ctx, Req)
}

func (p *kTikvClient) GetStoreSafeTS(ctx context.Context, Req *kvrpcpb.StoreSafeTSRequest, callOptions ...callopt.Option) (r *kvrpcpb.StoreSafeTSResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetStoreSafeTS(ctx, Req)
}

func (p *kTikvClient) GetLockWaitInfo(ctx context.Context, Req *kvrpcpb.GetLockWaitInfoRequest, callOptions ...callopt.Option) (r *kvrpcpb.GetLockWaitInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLockWaitInfo(ctx, Req)
}
