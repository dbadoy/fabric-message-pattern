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

func Response(obj interface{}) sc.Response {
	b, _ := json.Marshal(obj)
	return shim.Success(b)
}

func Initial() map[string]FuncWrap {
	fwm := make(map[string]FuncWrap)

	var fwl []FuncWrap
	fwl = append(fwl, FuncWrapper("queryCar", &QueryCarRequest{}, QueryCar)
	fwl = append(fwl, FuncWrapper("createCar", &CreateCarRequest{}, CreateCar)
	fwl = append(fwl, FuncWrapper("changeCarOwner", &ChangeCarOwnerRequest{}, ChangeCarOwner)	

	fwl = append(fwl, FuncWrapper("initLedger", &NonParamRequest{}, InitLedger))
	fwl = append(fwl, FuncWrapper("queryAllCars", &NonParamRequest{}, QueryAllCars))

	for _, fw := range fl {
		fwm[fw.fname] = fw
	}

	return fwm
}

func FuncWrapper(fname string, param interface{}, f Function) FuncWrap {
	return FuncWrap{
		fname   : fname,
		param   : param,
		Proceed : f,
	}
}