package cmd

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AA(t *testing.T) {

	const hx = "0a89010a86010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412660a2b6f736d6f316c70366535383635797879307666757a66726e776e3339327871616e67646e717a6a63743963122b6f736d6f317638756a657279647a6a367a306761377a716635336568393834396c367071387575373276721a0a0a05756f736d6f12013112490a0612040a020801123d0a090a046f736d6f1201301080c2d72f1a2b6f736d6f316c70366535383635797879307666757a66726e776e3339327871616e67646e717a6a637439631a001a20aa903f6e2719008feaa2417501a1d46754a0d4bf4ab8fdeb4312c1b1a7e4497d"
	_, enc := NewRootCmd()

	decoded, err := hex.DecodeString(hx)
	require.NoError(t, err)

	tx, err := enc.TxConfig.TxDecoder()(decoded)

	require.NoError(t, err)

	t.Logf("%#v", tx)
}