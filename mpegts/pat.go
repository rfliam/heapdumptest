package mpegts

type PAT interface {
	NumPrograms() int
	PMT(packet Packet) PMT
	ProgramMapPid() uint16
}

type PMT interface {
	Pids() []uint16
	RemoveElementaryStreams(pids []uint16)
}

type Packet interface {
	PacketBytes() []byte
	PayloadUnitStartIndicator() bool
	Pid() uint16
	Payload() []byte
	Header() []byte
	PAT() PAT
}
