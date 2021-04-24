/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-MHO-IEs"
 * 	found in "e2sm-mho.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_UE_Identity_H_
#define	_UE_Identity_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"

#ifdef __cplusplus
extern "C" {
#endif

/* UE-Identity */
typedef OCTET_STRING_t	 UE_Identity_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_UE_Identity;
asn_struct_free_f UE_Identity_free;
asn_struct_print_f UE_Identity_print;
asn_constr_check_f UE_Identity_constraint;
ber_type_decoder_f UE_Identity_decode_ber;
der_type_encoder_f UE_Identity_encode_der;
xer_type_decoder_f UE_Identity_decode_xer;
xer_type_encoder_f UE_Identity_encode_xer;
per_type_decoder_f UE_Identity_decode_uper;
per_type_encoder_f UE_Identity_encode_uper;
per_type_decoder_f UE_Identity_decode_aper;
per_type_encoder_f UE_Identity_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _UE_Identity_H_ */
#include "asn_internal.h"
