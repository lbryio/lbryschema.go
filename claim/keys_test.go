package claim

import (
	"testing"

	"gotest.tools/assert"
)

// The purpose of this test, is to make sure the function converts btcec.PublicKey to DER format the same way
// lbry SDK does as this is the bytes that are put into protobuf and the same bytes are used for verify signatures.
// Making sure these
func TestPublicKeyToDER(t *testing.T) {
	cert_claim_hex := "08011002225e0801100322583056301006072a8648ce3d020106052b8104000a03420004d015365a40f3e5c03c87227168e5851f44659837bcf6a3398ae633bc37d04ee19baeb26dc888003bd728146dbea39f5344bf8c52cedaf1a3a1623a0166f4a367"
	cert_claim, err := DecodeClaimHex(cert_claim_hex, "lbrycrd_main")
	if err != nil {
		t.Error(err)
	}
	p1, err := getPublicKeyFromBytes(cert_claim.Claim.GetChannel().PublicKey)
	if err != nil {
		t.Error(err)
	}

	pubkeyBytes2, err := PublicKeyToDER(p1)
	if err != nil {
		t.Error(err)
	}
	for i, b := range cert_claim.Claim.GetChannel().PublicKey {
		assert.Assert(t, b == pubkeyBytes2[i], "DER format in bytes must match!")
	}

	p2, err := getPublicKeyFromBytes(pubkeyBytes2)
	if err != nil {
		t.Error(err)
	}
	assert.Assert(t, p1.IsEqual(p2), "The keys produced must be the same key!")
}
