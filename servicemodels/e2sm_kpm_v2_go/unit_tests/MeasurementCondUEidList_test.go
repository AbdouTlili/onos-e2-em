// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/pdubuilder"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasCondUeIDL = "00000000  00 00 40 10 6f 6e 66 00  01 1f ff f0 21 22 23 40  |..@.onf.....!\"#@|\n" +
	"00000010  40 01 02 03 00 17 68 18  00 1e 00 01 70 00 00 18  |@.....h.....p...|\n" +
	"00000020  00 00 00 00 00 7a 00 01  c7 00 03 14 28 42 00 01  |.....z......(B..|\n" +
	"00000030  15 00 00 00 06 53 6f 6d  65 55 45                 |.....SomeUE|"

func createMeasurementCondUEIDList() (*e2smkpmv2.MeasurementCondUeidList, error) {

	muei := &e2smkpmv2.MatchingUeidItem{
		UeId: &e2smkpmv2.UeIdentity{
			Value: []byte("SomeUE"),
		},
	}
	muel := &e2smkpmv2.MatchingUeidList{
		Value: make([]*e2smkpmv2.MatchingUeidItem, 0),
	}
	muel.Value = append(muel.Value, muei)

	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 23
	var qfi int32 = 52
	var qci int32 = 24
	var qciMin int32 = 1
	var qciMax int32 = 30
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 25
	var layerMuMimo int32 = 1
	var sum = e2smkpmv2.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2smkpmv2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2smkpmv2.StartEndInd_START_END_IND_END

	labelInfoItem, err := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	if err != nil {
		return nil, err
	}

	mci1, err := pdubuilder.CreateMatchingCondItemMeasLabel(labelInfoItem.GetMeasLabel())
	if err != nil {
		return nil, err
	}

	mci2 := &e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2smkpmv2.TestCondInfo{
				TestType: &e2smkpmv2.TestCondType{
					TestCondType: &e2smkpmv2.TestCondType_AMbr{
						AMbr: e2smkpmv2.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2smkpmv2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2smkpmv2.TestCondValue{
					TestCondValue: &e2smkpmv2.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}

	//if err := mci2.Validate(); err != nil {
	//	return nil, err
	//}

	mcl := &e2smkpmv2.MatchingCondList{
		Value: make([]*e2smkpmv2.MatchingCondItem, 0),
	}
	mcl.Value = append(mcl.Value, mci1)
	mcl.Value = append(mcl.Value, mci2)

	measType := &e2smkpmv2.MeasurementType{
		MeasurementType: &e2smkpmv2.MeasurementType_MeasName{
			MeasName: &e2smkpmv2.MeasurementTypeName{
				Value: "onf",
			},
		},
	}

	mcUEIDi := &e2smkpmv2.MeasurementCondUeidItem{
		MatchingCond:     mcl,
		MeasType:         measType,
		MatchingUeidList: muel,
	}

	mcUEIDl := &e2smkpmv2.MeasurementCondUeidList{
		Value: make([]*e2smkpmv2.MeasurementCondUeidItem, 0),
	}
	mcUEIDl.Value = append(mcUEIDl.Value, mcUEIDi)

	//if err := mcUEIDl.Validate(); err != nil {
	//	return nil, err
	//}
	return mcUEIDl, nil
}

func Test_perEncodeMeasurementCondUEIDList(t *testing.T) {

	mcueIDl, err := createMeasurementCondUEIDList()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mcueIDl, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementCondUEIDList PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementCondUeidList{}
	err = aper.UnmarshalWithParams(per, &result, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementCondUEIDList PER - decoded\n%v", &result)
	assert.Equal(t, mcueIDl.GetValue()[0].GetMeasType().GetMeasName().GetValue(), result.GetValue()[0].GetMeasType().GetMeasName().GetValue())
	assert.DeepEqual(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSD(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQFi().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCi().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimax().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimin().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmax().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmin().GetValue(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetBitrateRange(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetLayerMuMimo(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSUm().Number(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinX(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinX())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinY(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinY())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinZ(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetStartEndInd().Number(), result.GetValue()[0].GetMatchingCond().GetValue()[0].GetMeasLabel().GetStartEndInd().Number())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestValue().GetValueInt(), result.GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestValue().GetValueInt())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestType().GetAMbr().Number(), result.GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestType().GetAMbr().Number())
	assert.Equal(t, mcueIDl.GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestExpr().Number(), result.GetValue()[0].GetMatchingCond().GetValue()[1].GetTestCondInfo().GetTestExpr().Number())
	assert.DeepEqual(t, mcueIDl.GetValue()[0].GetMatchingUeidList().GetValue()[0].GetUeId().GetValue(), result.GetValue()[0].GetMatchingUeidList().GetValue()[0].GetUeId().GetValue())
}

func Test_perMeasurementCondUEIDListCompareBytes(t *testing.T) {

	mcueIDl, err := createMeasurementCondUEIDList()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mcueIDl, "", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementCondUEIDList PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasCondUeIDL)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
