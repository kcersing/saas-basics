// Code generated by thriftgo (0.3.15). DO NOT EDIT.

package admin

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"saas/idl_gen/model/base"
)

type AdminService interface {
	//检查系统状态
	HealthCheck(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error)
	//获取验证码
	Captcha(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error)
}

type AdminServiceClient struct {
	c thrift.TClient
}

func NewAdminServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AdminServiceClient {
	return &AdminServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAdminServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AdminServiceClient {
	return &AdminServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAdminServiceClient(c thrift.TClient) *AdminServiceClient {
	return &AdminServiceClient{
		c: c,
	}
}

func (p *AdminServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *AdminServiceClient) HealthCheck(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error) {
	var _args AdminServiceHealthCheckArgs
	_args.Req = req
	var _result AdminServiceHealthCheckResult
	if err = p.Client_().Call(ctx, "HealthCheck", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *AdminServiceClient) Captcha(ctx context.Context, req *base.Empty) (r *base.NilResponse, err error) {
	var _args AdminServiceCaptchaArgs
	_args.Req = req
	var _result AdminServiceCaptchaResult
	if err = p.Client_().Call(ctx, "Captcha", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type AdminServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AdminService
}

func (p *AdminServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AdminServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AdminServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAdminServiceProcessor(handler AdminService) *AdminServiceProcessor {
	self := &AdminServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("HealthCheck", &adminServiceProcessorHealthCheck{handler: handler})
	self.AddToProcessorMap("Captcha", &adminServiceProcessorCaptcha{handler: handler})
	return self
}
func (p *AdminServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type adminServiceProcessorHealthCheck struct {
	handler AdminService
}

func (p *adminServiceProcessorHealthCheck) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AdminServiceHealthCheckArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("HealthCheck", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := AdminServiceHealthCheckResult{}
	var retval *base.NilResponse
	if retval, err2 = p.handler.HealthCheck(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing HealthCheck: "+err2.Error())
		oprot.WriteMessageBegin("HealthCheck", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("HealthCheck", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type adminServiceProcessorCaptcha struct {
	handler AdminService
}

func (p *adminServiceProcessorCaptcha) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AdminServiceCaptchaArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Captcha", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := AdminServiceCaptchaResult{}
	var retval *base.NilResponse
	if retval, err2 = p.handler.Captcha(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Captcha: "+err2.Error())
		oprot.WriteMessageBegin("Captcha", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Captcha", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type AdminServiceHealthCheckArgs struct {
	Req *base.Empty `thrift:"req,1"`
}

func NewAdminServiceHealthCheckArgs() *AdminServiceHealthCheckArgs {
	return &AdminServiceHealthCheckArgs{}
}

func (p *AdminServiceHealthCheckArgs) InitDefault() {
}

var AdminServiceHealthCheckArgs_Req_DEFAULT *base.Empty

func (p *AdminServiceHealthCheckArgs) GetReq() (v *base.Empty) {
	if !p.IsSetReq() {
		return AdminServiceHealthCheckArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_AdminServiceHealthCheckArgs = map[int16]string{
	1: "req",
}

func (p *AdminServiceHealthCheckArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AdminServiceHealthCheckArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AdminServiceHealthCheckArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AdminServiceHealthCheckArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := base.NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *AdminServiceHealthCheckArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("HealthCheck_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AdminServiceHealthCheckArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AdminServiceHealthCheckArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdminServiceHealthCheckArgs(%+v)", *p)

}

type AdminServiceHealthCheckResult struct {
	Success *base.NilResponse `thrift:"success,0,optional"`
}

func NewAdminServiceHealthCheckResult() *AdminServiceHealthCheckResult {
	return &AdminServiceHealthCheckResult{}
}

func (p *AdminServiceHealthCheckResult) InitDefault() {
}

var AdminServiceHealthCheckResult_Success_DEFAULT *base.NilResponse

func (p *AdminServiceHealthCheckResult) GetSuccess() (v *base.NilResponse) {
	if !p.IsSetSuccess() {
		return AdminServiceHealthCheckResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_AdminServiceHealthCheckResult = map[int16]string{
	0: "success",
}

func (p *AdminServiceHealthCheckResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AdminServiceHealthCheckResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AdminServiceHealthCheckResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AdminServiceHealthCheckResult) ReadField0(iprot thrift.TProtocol) error {
	_field := base.NewNilResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *AdminServiceHealthCheckResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("HealthCheck_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AdminServiceHealthCheckResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *AdminServiceHealthCheckResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdminServiceHealthCheckResult(%+v)", *p)

}

type AdminServiceCaptchaArgs struct {
	Req *base.Empty `thrift:"req,1"`
}

func NewAdminServiceCaptchaArgs() *AdminServiceCaptchaArgs {
	return &AdminServiceCaptchaArgs{}
}

func (p *AdminServiceCaptchaArgs) InitDefault() {
}

var AdminServiceCaptchaArgs_Req_DEFAULT *base.Empty

func (p *AdminServiceCaptchaArgs) GetReq() (v *base.Empty) {
	if !p.IsSetReq() {
		return AdminServiceCaptchaArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_AdminServiceCaptchaArgs = map[int16]string{
	1: "req",
}

func (p *AdminServiceCaptchaArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AdminServiceCaptchaArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AdminServiceCaptchaArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AdminServiceCaptchaArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := base.NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *AdminServiceCaptchaArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Captcha_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AdminServiceCaptchaArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *AdminServiceCaptchaArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdminServiceCaptchaArgs(%+v)", *p)

}

type AdminServiceCaptchaResult struct {
	Success *base.NilResponse `thrift:"success,0,optional"`
}

func NewAdminServiceCaptchaResult() *AdminServiceCaptchaResult {
	return &AdminServiceCaptchaResult{}
}

func (p *AdminServiceCaptchaResult) InitDefault() {
}

var AdminServiceCaptchaResult_Success_DEFAULT *base.NilResponse

func (p *AdminServiceCaptchaResult) GetSuccess() (v *base.NilResponse) {
	if !p.IsSetSuccess() {
		return AdminServiceCaptchaResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_AdminServiceCaptchaResult = map[int16]string{
	0: "success",
}

func (p *AdminServiceCaptchaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AdminServiceCaptchaResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_AdminServiceCaptchaResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *AdminServiceCaptchaResult) ReadField0(iprot thrift.TProtocol) error {
	_field := base.NewNilResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *AdminServiceCaptchaResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Captcha_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *AdminServiceCaptchaResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *AdminServiceCaptchaResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AdminServiceCaptchaResult(%+v)", *p)

}
