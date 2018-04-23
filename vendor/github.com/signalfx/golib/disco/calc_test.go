// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package disco

import (
	"fmt"
	"math"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type Calculator interface {
	// Parameters:
	//  - Num1
	//  - Num2
	Add(num1 int32, num2 int32) (r int32, err error)
}

type CalculatorClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewCalculatorClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *CalculatorClient {
	return &CalculatorClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewCalculatorClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *CalculatorClient {
	return &CalculatorClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Num1
//  - Num2
func (p *CalculatorClient) Add(num1 int32, num2 int32) (r int32, err error) {
	if err = p.sendAdd(num1, num2); err != nil {
		return
	}
	return p.recvAdd()
}

func (p *CalculatorClient) sendAdd(num1 int32, num2 int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("add", thrift.CALL, p.SeqId)
	args0 := NewAddArgs()
	args0.Num1 = num1
	args0.Num2 = num2
	err = args0.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *CalculatorClient) recvAdd() (value int32, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result1 := NewAddResult()
	err = result1.Read(iprot)
	iprot.ReadMessageEnd()
	value = result1.Success
	return
}

type CalculatorProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Calculator
}

func (p *CalculatorProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *CalculatorProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *CalculatorProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewCalculatorProcessor(handler Calculator) *CalculatorProcessor {

	self4 := &CalculatorProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["add"] = &calculatorProcessorAdd{handler: handler}
	return self4
}

func (p *CalculatorProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type calculatorProcessorAdd struct {
	handler Calculator
}

func (p *calculatorProcessorAdd) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewAddArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewAddResult()
	if result.Success, err = p.handler.Add(args.Num1, args.Num2); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: "+err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("add", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type AddArgs struct {
	Num1 int32 `thrift:"num1,1"`
	Num2 int32 `thrift:"num2,2"`
}

func NewAddArgs() *AddArgs {
	return &AddArgs{}
}

func (p *AddArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", p)
	} else {
		p.Num1 = v
	}
	return nil
}

func (p *AddArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", p)
	} else {
		p.Num2 = v
	}
	return nil
}

func (p *AddArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", p, err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", p, err)
	}
	return nil
}

func (p *AddArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num1", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:num1: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num1)); err != nil {
		return fmt.Errorf("%T.num1 (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:num1: %s", p, err)
	}
	return err
}

func (p *AddArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num2", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:num2: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num2)); err != nil {
		return fmt.Errorf("%T.num2 (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:num2: %s", p, err)
	}
	return err
}

func (p *AddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddArgs(%+v)", *p)
}

type AddResult struct {
	Success int32 `thrift:"success,0"`
}

func NewAddResult() *AddResult {
	return &AddResult{}
}

func (p *AddResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s", p)
	} else {
		p.Success = v
	}
	return nil
}

func (p *AddResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", p, err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", p, err)
	}
	return nil
}

func (p *AddResult) writeField0(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
		return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Success)); err != nil {
		return fmt.Errorf("%T.success (0) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 0:success: %s", p, err)
	}
	return err
}

func (p *AddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddResult(%+v)", *p)
}