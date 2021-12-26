// Code generated by Kitex v0.1.2. DO NOT EDIT.

package externalstorage

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/pingcap/kvproto/proto/kitex_gen/backup"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return externalStorageServiceInfo
}

var externalStorageServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ExternalStorage"
	handlerType := (*backup.ExternalStorage)(nil)
	methods := map[string]kitex.MethodInfo{
		"restore": kitex.NewMethodInfo(restoreHandler, newRestoreArgs, newRestoreResult, false),
		"save":    kitex.NewMethodInfo(saveHandler, newSaveArgs, newSaveResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "backup",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.1.2",
		Extra:           extra,
	}
	return svcInfo
}

func restoreHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(backup.ExternalStorageRestoreRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(backup.ExternalStorage).Restore(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RestoreArgs:
		success, err := handler.(backup.ExternalStorage).Restore(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RestoreResult)
		realResult.Success = success
	}
	return nil
}
func newRestoreArgs() interface{} {
	return &RestoreArgs{}
}

func newRestoreResult() interface{} {
	return &RestoreResult{}
}

type RestoreArgs struct {
	Req *backup.ExternalStorageRestoreRequest
}

func (p *RestoreArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RestoreArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RestoreArgs) Unmarshal(in []byte) error {
	msg := new(backup.ExternalStorageRestoreRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RestoreArgs_Req_DEFAULT *backup.ExternalStorageRestoreRequest

func (p *RestoreArgs) GetReq() *backup.ExternalStorageRestoreRequest {
	if !p.IsSetReq() {
		return RestoreArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RestoreArgs) IsSetReq() bool {
	return p.Req != nil
}

type RestoreResult struct {
	Success *backup.ExternalStorageRestoreResponse
}

var RestoreResult_Success_DEFAULT *backup.ExternalStorageRestoreResponse

func (p *RestoreResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RestoreResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RestoreResult) Unmarshal(in []byte) error {
	msg := new(backup.ExternalStorageRestoreResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RestoreResult) GetSuccess() *backup.ExternalStorageRestoreResponse {
	if !p.IsSetSuccess() {
		return RestoreResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RestoreResult) SetSuccess(x interface{}) {
	p.Success = x.(*backup.ExternalStorageRestoreResponse)
}

func (p *RestoreResult) IsSetSuccess() bool {
	return p.Success != nil
}

func saveHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(backup.ExternalStorageSaveRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(backup.ExternalStorage).Save(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SaveArgs:
		success, err := handler.(backup.ExternalStorage).Save(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SaveResult)
		realResult.Success = success
	}
	return nil
}
func newSaveArgs() interface{} {
	return &SaveArgs{}
}

func newSaveResult() interface{} {
	return &SaveResult{}
}

type SaveArgs struct {
	Req *backup.ExternalStorageSaveRequest
}

func (p *SaveArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SaveArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SaveArgs) Unmarshal(in []byte) error {
	msg := new(backup.ExternalStorageSaveRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SaveArgs_Req_DEFAULT *backup.ExternalStorageSaveRequest

func (p *SaveArgs) GetReq() *backup.ExternalStorageSaveRequest {
	if !p.IsSetReq() {
		return SaveArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SaveArgs) IsSetReq() bool {
	return p.Req != nil
}

type SaveResult struct {
	Success *backup.ExternalStorageSaveResponse
}

var SaveResult_Success_DEFAULT *backup.ExternalStorageSaveResponse

func (p *SaveResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SaveResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SaveResult) Unmarshal(in []byte) error {
	msg := new(backup.ExternalStorageSaveResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SaveResult) GetSuccess() *backup.ExternalStorageSaveResponse {
	if !p.IsSetSuccess() {
		return SaveResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SaveResult) SetSuccess(x interface{}) {
	p.Success = x.(*backup.ExternalStorageSaveResponse)
}

func (p *SaveResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Restore(ctx context.Context, Req *backup.ExternalStorageRestoreRequest) (r *backup.ExternalStorageRestoreResponse, err error) {
	var _args RestoreArgs
	_args.Req = Req
	var _result RestoreResult
	if err = p.c.Call(ctx, "restore", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Save(ctx context.Context, Req *backup.ExternalStorageSaveRequest) (r *backup.ExternalStorageSaveResponse, err error) {
	var _args SaveArgs
	_args.Req = Req
	var _result SaveResult
	if err = p.c.Call(ctx, "save", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
