package keyper

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"

	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/app"
)

// MakePrivkeyGeneratedEvent creates a PrivkeyGeneratedEvent from the given tendermint event of
// type "shutter.privkey-generated"
func MakePrivkeyGeneratedEvent(ev abcitypes.Event) (PrivkeyGeneratedEvent, error) {
	if len(ev.Attributes) < 2 {
		return PrivkeyGeneratedEvent{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("BatchIndex")) || !bytes.Equal(ev.Attributes[1].Key, []byte("Privkey")) {
		return PrivkeyGeneratedEvent{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	b, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return PrivkeyGeneratedEvent{}, err
	}
	privkey, err := app.DecodePrivkeyFromEvent(string(ev.Attributes[1].Value))
	if err != nil {
		return PrivkeyGeneratedEvent{}, err
	}

	return PrivkeyGeneratedEvent{uint64(b), privkey}, nil
}

// MakePubkeyGeneratedEvent creates a PubkeyGeneratedEvent from the given tendermint event of type
// type "shutter.pubkey-generated"
func MakePubkeyGeneratedEvent(ev abcitypes.Event) (PubkeyGeneratedEvent, error) {
	if len(ev.Attributes) < 2 {
		return PubkeyGeneratedEvent{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("BatchIndex")) || !bytes.Equal(ev.Attributes[1].Key, []byte("Pubkey")) {
		return PubkeyGeneratedEvent{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	b, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return PubkeyGeneratedEvent{}, err
	}
	pubkey, err := app.DecodePubkeyFromEvent(string(ev.Attributes[1].Value))
	if err != nil {
		return PubkeyGeneratedEvent{}, err
	}

	return PubkeyGeneratedEvent{uint64(b), pubkey}, nil
}

// MakeBatchConfigEvent creates a BatchConfigEvent from the given tendermint event of type
// "shutter.batch-config"
func MakeBatchConfigEvent(ev abcitypes.Event) (BatchConfigEvent, error) {
	if len(ev.Attributes) < 3 {
		return BatchConfigEvent{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("StartBatchIndex")) ||
		!bytes.Equal(ev.Attributes[1].Key, []byte("Threshold")) ||
		!bytes.Equal(ev.Attributes[2].Key, []byte("Keypers")) {
		return BatchConfigEvent{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	b, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return BatchConfigEvent{}, err
	}

	threshold, err := strconv.Atoi(string(ev.Attributes[1].Value))
	if err != nil {
		return BatchConfigEvent{}, err
	}
	keypers := app.DecodeAddressesFromEvent(string(ev.Attributes[2].Value))
	return BatchConfigEvent{uint64(b), uint32(threshold), keypers}, nil
}

// MakeEncryptionKeySignatureAddedEvent creates a EncryptionKeySignatureAddedEvent from the given
// tendermint event of type "shutter.encryption-key-signature-added"
func MakeEncryptionKeySignatureAddedEvent(ev abcitypes.Event) (EncryptionKeySignatureAddedEvent, error) {
	if len(ev.Attributes) < 4 {
		return EncryptionKeySignatureAddedEvent{}, fmt.Errorf("event contains not enough attributes: %+v", ev)
	}
	if !bytes.Equal(ev.Attributes[0].Key, []byte("KeyperIndex")) ||
		!bytes.Equal(ev.Attributes[1].Key, []byte("BatchIndex")) ||
		!bytes.Equal(ev.Attributes[2].Key, []byte("EncryptionKey")) ||
		!bytes.Equal(ev.Attributes[3].Key, []byte("Signature")) {
		return EncryptionKeySignatureAddedEvent{}, fmt.Errorf("bad event attributes: %+v", ev)
	}

	keyperIndex, err := strconv.Atoi(string(ev.Attributes[0].Value))
	if err != nil {
		return EncryptionKeySignatureAddedEvent{}, err
	}

	batchIndex, err := strconv.Atoi(string(ev.Attributes[1].Value))
	if err != nil {
		return EncryptionKeySignatureAddedEvent{}, err
	}

	key, err := base64.RawURLEncoding.DecodeString(string(ev.Attributes[2].Value))
	if err != nil {
		return EncryptionKeySignatureAddedEvent{}, err
	}

	signature, err := base64.RawURLEncoding.DecodeString(string(ev.Attributes[3].Value))
	if err != nil {
		return EncryptionKeySignatureAddedEvent{}, err
	}

	return EncryptionKeySignatureAddedEvent{
		KeyperIndex:   uint64(keyperIndex),
		BatchIndex:    uint64(batchIndex),
		EncryptionKey: key,
		Signature:     signature,
	}, nil
}

// MakeEvent creates an Event from the given tendermint event. It will return a
// PubkeyGeneratedEvent, PrivkeyGeneratedEvent or BatchConfigEvent based on the event's type.
func MakeEvent(ev abcitypes.Event) (IEvent, error) {
	if ev.Type == "shutter.privkey-generated" {
		res, err := MakePrivkeyGeneratedEvent(ev)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if ev.Type == "shutter.pubkey-generated" {
		res, err := MakePubkeyGeneratedEvent(ev)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if ev.Type == "shutter.batch-config" {
		res, err := MakeBatchConfigEvent(ev)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if ev.Type == "shutter.encryption-key-signature-added" {
		res, err := MakeEncryptionKeySignatureAddedEvent(ev)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, fmt.Errorf("cannot make event from %+v", ev)
}
