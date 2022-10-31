package token

import (
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiGPBDP4VCc+cp7f02lvo
xEs0vlTYuf2rmNuKyVX2Gf3Mf8sDBK9KTSBbTSfRJ0OSYt4f/TSrAgbb5Z/BWFT/
uN1kJ/Yz9HnKIAvahTC6HVLxnyoqZGnDht9OB2LOhC23eR+pUAxcLTqSvcs0NbZ1
nnZqMAlLcjf5J5hsYqhEWWV4E2jgbI4lPLz5mig5+e34GA78i7OvpycBV2JtuGUL
mAIwYEN5PyxBNByv/8FYzKZX9ReKSdaFwUX0hGDCSe01qZazDvqEsR5Q2XQGeLyv
uWBZOJaSp9Qnl+DoTKEBrBi9+UDQvCZoHWWmMvhdfJsAzp3uX9SlYEij4qIBn6L5
RwIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("can not parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	tkn := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9nZW4iLCJzdWIiOiI2MTBjZTBiYjgwMzY1MGJmYmVmYTBlMjQifQ.RoGL7OpNa56BPIRSNBdz4Du_nisM2rZVqHTjd9YZIJfWDbY_EihRAbTjV6gx4dWK7e20M3CQaLMIpNhT2B1p1m5yleXfhShIVm9cCVMPm719AoOYzxP5Ol9PFSkmvNtePVBvtuS-aR0boeStmXsQSbCF30avNWAeta3niSbA_BHEJDxLo2NYsHa_qZrvNYOf_Qz3zaW07iDNyn3tb-68m8Y1PLjbw_J2KPU392sf2j0uedIwY6CjX7N8-m4cJoRXxX40_5vOINMZb-9eS-sgyi1W1YUcVUCYMtAGa8HtnKHm7HhD-fjTL7M5c0HMOL5GIoO311u9cw1C_sVsg2_M2Q`

	jwt.TimeFunc = func() time.Time {
		return time.Unix(1516239122, 0)
	}

	accountID, err := v.Verify(tkn)
	if err != nil {
		t.Errorf("verification failed: %v", err)
	}
	want := "610ce0bb803650bfbefa0e24"
	if accountID != want {
		t.Errorf("wrong account id, want: %q, got: %q", want, accountID)
	}
}
