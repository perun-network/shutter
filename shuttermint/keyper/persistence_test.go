package keyper

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/google/go-cmp/cmp"
	"github.com/shutter-network/shutter/shuttermint/keyper/observe"
	"github.com/shutter-network/shutter/shuttermint/keyper/tee"
	"github.com/spf13/viper"

	"gotest.tools/v3/assert"
)

func randomAddress() (ret common.Address) {
	_, _ = rand.Read(ret[:])
	return
}

// Test that the loading and saving of the TOML config works.
func TestConfigPersistence(t *testing.T) {
	var cfg = Config{
		ShuttermintURL:              "ShuttermintURL",
		EthereumURL:                 "EthereumURL",
		DBDir:                       "DBDir",
		ConfigContractAddress:       randomAddress(),
		BatcherContractAddress:      randomAddress(),
		KeyBroadcastContractAddress: randomAddress(),
		ExecutorContractAddress:     randomAddress(),
		DepositContractAddress:      randomAddress(),
		KeyperSlasherAddress:        randomAddress(),
		MainChainFollowDistance:     11111,
		ExecutionStaggering:         22222,
		DKGPhaseLength:              33333,
		GasPriceMultiplier:          44444.4444,
	}

	assert.NilError(t, cfg.GenerateNewKeys(), "Generating keys")

	var encoding bytes.Buffer
	var v = viper.New()
	initial_setup_ran := false
	t.Run("TOML config", func(t *testing.T) {
		assert.NilError(t, cfg.WriteTOML(&encoding), "Writing TOML")
		encodingCopy := encoding

		v.SetConfigType("toml")
		assert.NilError(t, v.ReadConfig(&encodingCopy), "Reading TOML")

		assert.Equal(t, v.Get("ShuttermintURL").(string), cfg.ShuttermintURL)
		assert.Equal(t, v.Get("EthereumURL").(string), cfg.EthereumURL)
		assert.Equal(t, v.Get("DBDir").(string), cfg.DBDir)
		assert.Equal(t, v.Get("SigningKey").(string), hex.EncodeToString(
			crypto.FromECDSA(cfg.SigningKey)))
		assert.Equal(t, v.Get("EncryptionKey").(string), hex.EncodeToString(
			crypto.FromECDSA(cfg.EncryptionKey.ExportECDSA())))
		initial_setup_ran = true
	})

	t.Run("Viper unmarshaling", func(t *testing.T) {
		if !initial_setup_ran {
			t.Skip()
		}
		var cfg2 Config
		cfg2.Unmarshal(v)
		assert.DeepEqual(t, cfg, cfg2,
			cmp.Comparer(func(a, b *big.Int) bool { return a.Cmp(b) == 0 }),
			cmp.Comparer(func(a, b ecies.ECIESParams) bool { return true }))

		var encoding2 bytes.Buffer
		assert.NilError(t, cfg2.WriteTOML(&encoding2), "Writing TOML")
		assert.Equal(t, string(encoding.Bytes()), string(encoding2.Bytes()))
	})
}

// Test that the saving and loading of state from Gob works.
func TestGobStatePersistence(t *testing.T) {
	keypers := []common.Address{randomAddress(), randomAddress(), randomAddress()}
	var st = storedState{
		State: &State{
			CheckInMessageSent:       true,
			LastSentBatchConfigIndex: 12356,
			LastEonStarted:           948523,
			DKGs: []DKG{
				DKG{Eon: 2, StartBatchIndex: 98, Keypers: keypers},
				DKG{Eon: 4, PhaseLength: PhaseLength{1, 2, 3, 4}, Keypers: keypers}},
			EKGs: []*EKG{&EKG{Eon: 5, Keypers: keypers}},
		},
		Shutter: &observe.Shutter{
			CurrentBlock:        435,
			LastCommittedHeight: 745,
		},
		MainChain: &observe.MainChain{
			FollowDistance:        43,
			CurrentBlock:          3426,
			NumExecutionHalfSteps: 348,
		},
	}

	var enc bytes.Buffer
	tee.DoWithoutTEE(func() {
		assert.NilError(t, writeStoredState(&enc, st))
	})
	encBytes := enc.Bytes()

	st2, err := loadStoredState(&enc)
	assert.NilError(t, err)

	var enc2 bytes.Buffer
	tee.DoWithoutTEE(func() {
		assert.NilError(t, writeStoredState(&enc2, st2))
	})

	assert.DeepEqual(t, encBytes, enc2.Bytes())
}
