#!/bin/sh

# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

proto_imports=".:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_kpm/v1beta1/e2sm_kpm_ies.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_kpm_v2/v2/e2sm_kpm_v2.proto

#For future modules based on onos-lib-go/api/asn1 BitString
#protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
#  --validate_out=lang=go:. --proto_path=servicemodels \
#  --go_out=../../.. \
#  e2sm_kpm_v2/v2/e2sm_kpm_v2.proto
#protoc-go-inject-tag -input=servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2/e2sm_kpm_v2.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_kpm_v2_go/v2/e2sm_kpm_v2_go.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go/e2sm_kpm_v2_go.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_kpm_v2_go/v2/e2sm_kpm_v2_go.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_rc_pre_go/v2/e2sm_rc_pre_v2_go.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go/e2sm_rc_pre_v2_go.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rc_pre_go/v2/e2sm_rc_pre_v2_go.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_rsm/v1/e2sm_rsm_v1.proto e2sm_rsm/v1/e2sm_v2.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-rsm-ies/e2sm_rsm_v1.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rsm/v1/e2sm_rsm_v1.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_rsm/v1/e2sm_v2.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rsm/v1/e2sm-v2-ies/e2sm_v2.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rsm/v1/e2sm_v2.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_mho_go/v2/e2sm_mho_go.proto e2sm_mho_go/v2/e2sm_v2.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go/e2sm_mho_go.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_mho_go/v2/e2sm_mho_go.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_mho_go/v2/e2sm_v2.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies/e2sm_v2.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_mho_go/v2/e2sm_v2.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_rc/v1/e2sm_rc.proto e2sm_rc/v1/e2sm_common_ies.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies/e2sm_rc.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rc/v1/e2sm_rc.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_rc/v1/e2sm_common_ies.proto
protoc-go-inject-tag -input=servicemodels/github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies/e2sm_common_ies.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rc/v1/e2sm_common_ies.proto

#protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
#  --proto_path=servicemodels \
#  --go_out=. \
#  e2sm_ni/v1beta1/e2sm_ni_ies.proto
#protoc-go-inject-tag -input=e2sm_ni/v1beta1/e2sm-ni-ies/e2sm_ni_ies.pb.go
#
#protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
#  --validate_out=lang=go:. --proto_path=servicemodels \
#  --go_out=. \
#  e2sm_ni/v1beta1/e2sm_ni_ies.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rc_pre/v2/e2sm_rc_pre_v2.proto

protoc -I=$proto_imports --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_mho/v1/e2sm_mho.proto

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  test_sm_aper_go_lib/v1/test_sm.proto
protoc-go-inject-tag -input=servicemodels/test_sm_aper_go_lib/v1/test-sm-ies/test_sm.pb.go

### RMET SM 
protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_rmet/v1/e2sm_rmet.proto

protoc-go-inject-tag -input=servicemodels/e2sm_rmet/v1/e2sm-rmet-go/e2sm_rmet.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_rmet/v1/e2sm_rmet.proto



### MET SM 
protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_met/v1/e2sm_met.proto

protoc-go-inject-tag -input=servicemodels/e2sm_met/v1/e2sm-met-go/e2sm_met.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_met/v1/e2sm_met.proto




### MET SM 
protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --go_out=./servicemodels/ \
  e2sm_xtdd/v1/e2sm_xtdd.proto

protoc-go-inject-tag -input=servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go/e2sm_xtdd.pb.go

protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api \
  --validate_out=lang=go:. --proto_path=servicemodels \
  --go_out=. \
  e2sm_xtdd/v1/e2sm_xtdd.proto




cp -r github.com/onosproject/onos-e2-sm/* .
cp -r github.com/AbdouTlili/onos-e2-sm/* .
# rm -rf github.com

cp -r servicemodels/github.com/onosproject/onos-e2-sm/* .
cp -r servicemodels/github.com/AbdouTlili/onos-e2-sm/* .
# rm -rf servicemodels/github.com

