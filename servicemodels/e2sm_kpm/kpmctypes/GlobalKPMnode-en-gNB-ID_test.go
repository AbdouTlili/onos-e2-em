// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerDecodeGlobalKPMnodeEnGNbID(t *testing.T) {

	plmnID := []byte{0x21, 0x18}

	globalKPMnodeEnGNbID := &e2sm_kpm_ies.GlobalKpmnodeEnGnbId{
		GlobalGNbId: &e2sm_kpm_ies.GlobalenGnbId{
			PLmnIdentity: &e2sm_kpm_ies.PlmnIdentity{
				Value: plmnID,
			},
			GNbId: &e2sm_kpm_ies.EngnbId{
				EngnbId: &e2sm_kpm_ies.EngnbId_GNbId{
					GNbId: &e2sm_kpm_ies.BitString{
						Value: 0x9bcd4, //uint64
						Len:   22,      //uint32
					},
				},
			},
		},
	}

	xer, err := xerEncodeGlobalKPMnodeEnGNbID(globalKPMnodeEnGNbID)
	assert.NilError(t, err)
	t.Logf("GlobalKPMnodeEnGNb XER\n%s", string(xer))

	result, err := xerDecodeGlobalKPMnodeEnGNbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
}
