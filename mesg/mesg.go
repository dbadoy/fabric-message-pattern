package mesg

import (
	"encoding/json"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type FuncWrap struct {
	fname   string
	param   Request
	Proceed Function
}

type Function func(APIstub shim.ChaincodeStubInterface, request Request) sc.Response

type Request interface {
	FieldCount() int
	Set([]string) error
	Get() interface{}
}

func Initial() map[string]FuncWrap {
	fwm := make(map[string]FuncWrap)

	var fwl []FuncWrap
	fwl = append(fwl, funcWrapper("queryCar", &QueryCarRequest{}, QueryCar)
	fwl = append(fwl, funcWrapper("createCar", &CreateCarRequest{}, CreateCar)
	fwl = append(fwl, funcWrapper("changeCarOwner", &ChangeCarOwnerRequest{}, ChangeCarOwner)

	fwl = append(fwl, funcWrapper("initLedger", &NonParamRequest{}, InitLedger))
	fwl = append(fwl, funcWrapper("queryAllCars", &NonParamRequest{}, QueryAllCars))

	for _, fw := range fl {
		fwm[fw.fname] = fw
	}

	return fwm
}

func funcWrapper(fname string, param interface{}, f Function) FuncWrap {
	return FuncWrap{
		fname   : fname,
		param   : param,
		Proceed : f,
	}
}