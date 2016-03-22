package handshake_test

import (
	"bytes"
	"testing"

	"github.com/WatchBeam/rtmp/handshake"
	"github.com/stretchr/testify/assert"
)

func TestItConstructsVersionSequences(t *testing.T) {
	v := handshake.NewVersionSequence()

	assert.IsType(t, new(handshake.VerisonSequence), v)
}

func TestItReadsSupportedVersionNumbers(t *testing.T) {
	v := handshake.NewVersionSequence()
	err := v.Read(bytes.NewBuffer([]byte{0x3}))

	assert.Nil(t, err)
}

func TestItRejectsUnsupportedVersionNumbers(t *testing.T) {
	v := handshake.NewVersionSequence()
	err := v.Read(bytes.NewBuffer([]byte{0x4}))

	assert.Equal(t, "rtmp/handshake: unsupported version 4", err.Error())
}

func TestItWritesVersionAndS1(t *testing.T) {
	buf := new(bytes.Buffer)
	v := handshake.NewVersionSequence()

	assert.Nil(t, v.Write(buf))
	assert.Equal(t, []byte{0x3}, buf.Bytes()[:1])
	assert.Len(t, buf.Bytes(), 1+4+4+1528)
}

func TestItReturnsAnInitializedAckSequence(t *testing.T) {
	v := handshake.NewVersionSequence()
	next := v.Next()

	switch typ := next.(type) {
	case *handshake.AckSequence:
		assert.Equal(t, v.S1.Payload, typ.S1Payload)
	default:
		t.Fatalf("handshake: got unknown type for next sequence (%T)", v)
	}
}
