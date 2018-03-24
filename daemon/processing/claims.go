package processing

import (
	"encoding/hex"
	"strings"

	"github.com/lbryio/chainquery/lbrycrd"
	"github.com/lbryio/chainquery/model"
	"github.com/lbryio/errors.go"

	"encoding/binary"
	"github.com/sirupsen/logrus"
)

func processAsClaim(script []byte, vout model.Output) (address *string, err error) {
	var pubkeyscript []byte
	var name string
	var claimid string
	if lbrycrd.IsClaimNameScript(script) {
		name, claimid, pubkeyscript, err = processClaimNameScript(&script, vout)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if lbrycrd.IsClaimSupportScript(script) {
		name, claimid, pubkeyscript, err = processClaimSupportScript(&script, vout)
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if lbrycrd.IsClaimUpdateScript(script) {
		name, claimid, pubkeyscript, err = processClaimUpdateScript(&script, vout)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	pksAddress := lbrycrd.GetAddressFromPublicKeyScript(pubkeyscript)
	address = &pksAddress
	logrus.Debug("Handled Claim: ", " Name ", name, ", ClaimID ", claimid)

	return nil, errors.Base("Not a claim -- " + hex.EncodeToString(script))
}

func processClaimNameScript(script *[]byte, vout model.Output) (name string, claimid string, pkscript []byte, err error) {
	name, value, pubkeyscript, err := lbrycrd.ParseClaimNameScript(*script)
	if err != nil {
		errors.Prefix("Claim name processing error: ", err)
		return name, claimid, pubkeyscript, err
	}
	claim, err := lbrycrd.DecodeClaimValue(name, value)
	if claim != nil {
		claimId := getClaimIdFromOutput(&vout)
		if claimId != "" {
			logrus.Debug("ClaimName ", name, " ClaimId ", claimId)
		}
	}

	return name, claimid, pubkeyscript, err
}

func processClaimSupportScript(script *[]byte, vout model.Output) (name string, claimid string, pubkeyscript []byte, err error) {
	name, claimid, pubkeyscript, err = lbrycrd.ParseClaimSupportScript(*script)
	if err != nil {
		errors.Prefix("Claim support processing error: ", err)
		return name, claimid, pubkeyscript, err
	}
	//log.Debug("ClaimSupport ", name, " ClaimId ", claimid)

	return name, claimid, pubkeyscript, err
}

func processClaimUpdateScript(script *[]byte, vout model.Output) (name string, claimId string, pubkeyscript []byte, err error) {
	name, claimId, value, pubkeyscript, err := lbrycrd.ParseClaimUpdateScript(*script)
	if err != nil {
		errors.Prefix("Claim update processing error: ", err)
		return name, claimId, pubkeyscript, err
	}
	claim, err := lbrycrd.DecodeClaimValue(name, value)
	if claim != nil {
		//log.Debug("ClaimUpdate ", name, " ClaimId ", claimId)
	}
	return name, claimId, pubkeyscript, err
}

func GetAddressFromClaimASM(asm string) string {
	sections := strings.Split(asm, " ")
	address, err := lbrycrd.GetAddressFromP2PKH(sections[7])
	if err != nil {
		panic(err)
	}

	return address
}

func getClaimIdFromOutput(vout *model.Output) string {

	txHashBytes, err := hex.DecodeString(vout.TransactionHash)
	println("Hash bytes: ", len(txHashBytes))
	if err != nil {
		logrus.Error("Could not decode hex string -> ", vout.TransactionHash, " ", err)
	}
	bs := make([]byte, 4) //uint32 byte array
	println("txHash: ", hex.EncodeToString(txHashBytes), "vout(uint32): ", uint32(vout.Vout))
	binary.BigEndian.PutUint32(bs, uint32(vout.Vout))

	claimIdBytes := append(txHashBytes, bs...)
	println("Concatenated ", hex.EncodeToString(claimIdBytes))
	claimIdBytes = lbrycrd.Hash160(claimIdBytes)

	return hex.EncodeToString(claimIdBytes)
}
