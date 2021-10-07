// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdudecoder

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/encoder"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func createE2SmKpmRanFunctionDescription() (*e2sm_kpm_v2_go.E2SmKpmRanfunctionDescription, error) {

	var rfSn = "onf"
	var rfE2SMoid = "1.3.6.1.4.1.53148.1.2.2.2"
	var rfd = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellIDbits := []byte{0x12, 0xF0, 0xDE, 0xBC, 0x50}
	cellGlobalID, err := pdubuilder.CreateCellGlobalIDNRCGI(plmnID, cellIDbits) // 36 bits
	if err != nil {
		return nil, err
	}

	var cellObjID = "ONF"
	cellMeasObjItem := pdubuilder.CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalKpmnodeID, err := pdubuilder.CreateGlobalKpmnodeIDgNBID(&bs, plmnID)
	if err != nil {
		return nil, err
	}
	globalKpmnodeID.GetGNb().GNbCuUpId = &e2sm_kpm_v2_go.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalKpmnodeID.GetGNb().GNbDuId = &e2sm_kpm_v2_go.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2sm_kpm_v2_go.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	kpmNodeItem := pdubuilder.CreateRicKpmnodeItem(globalKpmnodeID).SetCellMeasurementObjectList(cmol)

	rknl := make([]*e2sm_kpm_v2_go.RicKpmnodeItem, 0)
	rknl = append(rknl, kpmNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15
	retsi := pdubuilder.CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2sm_kpm_v2_go.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2sm_kpm_v2_go.MeasurementInfoActionList{
		Value: make([]*e2sm_kpm_v2_go.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := pdubuilder.CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2sm_kpm_v2_go.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 1
	var indHdrFormat int32 = 2
	rrsi := pdubuilder.CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2sm_kpm_v2_go.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmKpmPdu := pdubuilder.CreateE2SmKpmRanfunctionDescription(rfSn, rfE2SMoid, rfd).SetRanFunctionInstance(rfi).SetRicEventTriggerStyleList(retsl).SetRicReportStyleList(rrsl).SetRicKpmNodeList(rknl)
	fmt.Printf("Created E2SM-KPM-RanFunctionDescription is \n %v \n", newE2SmKpmPdu)

	return newE2SmKpmPdu, nil
}

func Test_DecodeE2SmKpmRanFunctionDescription(t *testing.T) {

	e2SmKpmPdu, err := createE2SmKpmRanFunctionDescription()
	assert.NilError(t, err)

	ranFunctionName, ricEventList, ricReportList, err := DecodeE2SmKpmRanFunctionDescription(e2SmKpmPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", string(ranFunctionName.RanFunctionE2SmOid))
	//assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	//assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 1, len(*ricReportList))
}

func Test_DecodeE2SmKpmRanFunctionDescriptionBytes(t *testing.T) {
	rfd, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(ranFuncDescBytes)
	assert.NilError(t, err)
	//err = rfd.Validate()
	//assert.NilError(t, err)
	assert.Equal(t, "ORAN-E2SM-KPM", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "KPM monitor", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
}

var ranFuncDescBytes = []byte{0x70, 0x18, 0x4F, 0x52, 0x41, 0x4E, 0x2D, 0x45, 0x32, 0x53, 0x4D, 0x2D, 0x4B, 0x50, 0x4D, 0x00,
	0x00, 0x18, 0x31, 0x2E, 0x33, 0x2E, 0x36, 0x2E, 0x31, 0x2E, 0x34, 0x2E, 0x31, 0x2E, 0x35, 0x33,
	0x31, 0x34, 0x38, 0x2E, 0x31, 0x2E, 0x32, 0x2E, 0x32, 0x2E, 0x32, 0x05, 0x00, 0x4B, 0x50, 0x4D,
	0x20, 0x6D, 0x6F, 0x6E, 0x69, 0x74, 0x6F, 0x72, 0x00, 0x00, 0x40, 0x00, 0x13, 0xF1, 0x84, 0x50,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x30, 0x00, 0x13, 0xF1, 0x84, 0x00, 0x00,
	0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6F, 0x64, 0x69, 0x63, 0x20,
	0x72, 0x65, 0x70, 0x6F, 0x72, 0x74, 0x00, 0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4E,
	0x6F, 0x64, 0x65, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x00,
	0x01, 0x00, 0x07, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74,
	0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x53, 0x75, 0x6D, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52,
	0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2E,
	0x53, 0x75, 0x6D, 0x00, 0x00, 0x01, 0x42, 0xA0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E,
	0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x53, 0x75, 0x6D, 0x00, 0x00,
	0x02, 0x43, 0xC0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45, 0x73, 0x74,
	0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x72, 0x65, 0x63, 0x6F, 0x6E, 0x66, 0x69, 0x67, 0x46, 0x61,
	0x69, 0x6C, 0x00, 0x00, 0x03, 0x43, 0x00, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52,
	0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x48, 0x4F, 0x46, 0x61, 0x69, 0x6C,
	0x00, 0x00, 0x04, 0x42, 0xE0, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x52, 0x65, 0x45,
	0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2E, 0x4F, 0x74, 0x68, 0x65, 0x72, 0x00, 0x00, 0x05,
	0x41, 0x60, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x4D, 0x65, 0x61, 0x6E, 0x00, 0x00,
	0x06, 0x41, 0x40, 0x52, 0x52, 0x43, 0x2E, 0x43, 0x6F, 0x6E, 0x6E, 0x4D, 0x61, 0x78, 0x00, 0x00,
	0x07, 0x00, 0x01, 0x00, 0x01}

func Test_DecodeE2SmKpmRanFunctionDescriptionBytesRadisys(t *testing.T) {
	t.Logf("Bytes in HEX format are \n%v\n", hex.Dump(rSysBytes))
	rfd, err := encoder.PerDecodeE2SmKpmRanFunctionDescription(rSysBytes)
	assert.NilError(t, err)
	//err = rfd.Validate()
	//assert.NilError(t, err)
	t.Logf("Protobuf message is \n%v", rfd)
	assert.Equal(t, "ORAN-E2SM-KPM", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "KPM monitor", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "1.3.6.1.4.1.53148.1.2.2.2", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
}

var rSysBytes = []byte{
	0x70, 0x18, 0x4f, 0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x4b, 0x50, 0x4d, 0x00, 0x00, 0x18, 0x31, 0x2e,
	0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x32, 0x2e,
	0x32, 0x2e, 0x32, 0x05, 0x00, 0x4b, 0x50, 0x4d, 0x20, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x00, 0x00, 0x40, 0x00,
	0x13, 0xf1, 0x84, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x30, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00,
	0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x20, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x00, 0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x20, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x01, 0x00, 0x07, 0x42, 0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52,
	0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x00,
	0x00, 0x01, 0x42, 0xa0, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41,
	0x74, 0x74, 0x2e, 0x53, 0x75, 0x6d, 0x00, 0x00, 0x02, 0x43, 0xc0, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52,
	0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x61,
	0x69, 0x6c, 0x00, 0x00, 0x03, 0x43, 0x00, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74,
	0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x48, 0x4f, 0x46, 0x61, 0x69, 0x6c, 0x00, 0x00, 0x04, 0x42, 0xe0, 0x52, 0x52, 0x43,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x00, 0x00, 0x05, 0x41, 0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x65, 0x61, 0x6e, 0x00, 0x00,
	0x06, 0x41, 0x40, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x00, 0x00, 0x07, 0x00, 0x01, 0x00,
	0x01}
