package streams

import (
	"encoding/json"
	"errors"

	"github.com/rfliam/heapdumptest/mpegts"
)

type ElementaryStreamType uint8

const (
	AudioElementaryStream ElementaryStreamType = iota
	VideoElementaryStream
	DataElementaryStream
	UknownElementaryStream
)

const (
	AACLCCodec        = "mp4a.40.2"
	HEAACCodec        = "mp4a.40.5"
	HEAACv2Codec      = "mp4a.40.29"
	DolbyDigitalCodec = "ec-3"
)

// A StreamInfo contains all the metadata about a mp2t elementary stream.
type ElementaryStreamInfo struct {
	StreamIndex int
	Err         error                `json:",omitempty"`
	Pid         uint16               `json:",omitempty"` // TS Pid
	Type        ElementaryStreamType `json:"-"`
	Bandwidth   uint64
	Codec       string `json:",omitempty"`
	Lang        string `json:"Language,omitempty"`
}

// TrasportStreamInfo contains all the metatadata/information for a
// single ts.
type TransportStreamInfo struct {
	StreamIndex int
	Err         error

	// The mux info
	PAT       mpegts.PAT
	PATPacket mpegts.Packet
	PMT       mpegts.PMT
	PMTPacket mpegts.Packet

	PMTPid uint16

	EsInfos map[uint16]ElementaryStreamInfo // From pid to elementarystreaminfo, will only contain up streams
}

func (tsi TransportStreamInfo) VideoElementaryStream() (esi ElementaryStreamInfo, found bool) {
	for _, esi := range tsi.EsInfos {
		if esi.Type == VideoElementaryStream {
			return esi, true
		}
	}
	return ElementaryStreamInfo{}, false
}

func (tsi TransportStreamInfo) AudioElementaryStreams() map[uint16]ElementaryStreamInfo {
	audioEs := make(map[uint16]ElementaryStreamInfo, len(tsi.EsInfos))
	for _, esi := range tsi.EsInfos {
		if esi.Type == AudioElementaryStream {
			audioEs[esi.Pid] = esi
		}
	}
	return audioEs
}

func (est ElementaryStreamType) MarshalJSON() ([]byte, error) {
	switch est {
	case AudioElementaryStream:
		return []byte("\"AudioElementaryStream\""), nil
	case VideoElementaryStream:
		return []byte("\"VideoElementaryStream\""), nil
	case DataElementaryStream:
		return []byte("\"DataElementaryStream\""), nil
	case UknownElementaryStream:
		return []byte("\"UknownElementaryStream\""), nil
	default:
		return []byte{}, errors.New("Invalid ElementaryStreamType")
	}
}

func (esi ElementaryStreamInfo) MarshalJSON() ([]byte, error) {
	res := map[string]interface{}{
		"Bandwidth": esi.Bandwidth,
		"Codec":     esi.Codec,
		"Pid":       esi.Pid,
		"Type":      esi.Type,
	}

	return json.Marshal(res)
}
