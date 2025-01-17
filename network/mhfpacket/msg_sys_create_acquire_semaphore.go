package mhfpacket

import (
	"errors"
	"erupe-ce/common/bfutil"
	"erupe-ce/common/byteframe"
	_config "erupe-ce/config"
	"erupe-ce/network"
	"erupe-ce/network/clientctx"
)

// MsgSysCreateAcquireSemaphore represents the MSG_SYS_CREATE_ACQUIRE_SEMAPHORE
type MsgSysCreateAcquireSemaphore struct {
	AckHandle   uint32
	Unk0        uint16
	PlayerCount uint8
	SemaphoreID string
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysCreateAcquireSemaphore) Opcode() network.PacketID {
	return network.MSG_SYS_CREATE_ACQUIRE_SEMAPHORE
}

// Parse parses the packet from binary
func (m *MsgSysCreateAcquireSemaphore) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	if _config.ErupeConfig.RealClientMode >= _config.S7 { // Assuming this was added with Ravi?
		m.PlayerCount = bf.ReadUint8()
	}
	SemaphoreIDLength := bf.ReadUint8()
	m.SemaphoreID = string(bfutil.UpToNull(bf.ReadBytes(uint(SemaphoreIDLength))))
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgSysCreateAcquireSemaphore) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
