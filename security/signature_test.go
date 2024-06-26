package security

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/imgproxy/imgproxy/v3/config"
)

type SignatureTestSuite struct {
	suite.Suite
}

func (s *SignatureTestSuite) SetupTest() {
	config.Reset()

	config.Keys = [][]byte{[]byte("test-key")}
	config.Salts = [][]byte{[]byte("test-salt")}
}

func (s *SignatureTestSuite) TestVerifySignature() {
	err := VerifySignature("dtLwhdnPPiu_epMl1LrzheLpvHas-4mwvY6L3Z8WwlY", "asd")
	s.Require().NoError(err)
}

func (s *SignatureTestSuite) TestVerifySignatureTruncated() {
	config.SignatureSize = 8

	err := VerifySignature("dtLwhdnPPis", "asd")
	s.Require().NoError(err)
}

func (s *SignatureTestSuite) TestVerifySignatureInvalid() {
	err := VerifySignature("dtLwhdnPPis", "asd")
	s.Require().Error(err)
}

func (s *SignatureTestSuite) TestVerifySignatureMultiplePairs() {
	config.Keys = append(config.Keys, []byte("test-key2"))
	config.Salts = append(config.Salts, []byte("test-salt2"))

	err := VerifySignature("dtLwhdnPPiu_epMl1LrzheLpvHas-4mwvY6L3Z8WwlY", "asd")
	s.Require().NoError(err)

	err = VerifySignature("jbDffNPt1-XBgDccsaE-XJB9lx8JIJqdeYIZKgOqZpg", "asd")
	s.Require().NoError(err)

	err = VerifySignature("dtLwhdnPPis", "asd")
	s.Require().Error(err)
}

func (s *SignatureTestSuite) TestVerifySignatureTrusted() {
	config.TrustedSignatures = []string{"truested"}
	defer func() {
		config.TrustedSignatures = []string{}
	}()

	err := VerifySignature("truested", "asd")
	s.Require().NoError(err)

	err = VerifySignature("untrusted", "asd")
	s.Require().Error(err)
}

func TestSignature(t *testing.T) {
	suite.Run(t, new(SignatureTestSuite))
}
