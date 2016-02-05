package main

import "bytes"
import "encoding/binary"

const (
	YuntiNCPProto_undefined = iota
	// Write to a remote stream
	//# uint64 poolid, uint64 streamid,uint64 streamseqid, []byte payload
	YuntiNCPProto_RemoteWrite
	//Open a TCP connection as a remote stream
	//# uint64 poolid,uint64 pcseqid,(string) []byte target
	YuntiNCPProto_RemoteTCPOpen
	// Add a stream pool and also engage push
	//# uint64 cseqid
	YuntiNCPProto_PoolAdd
	// Associate current socket with existing pool for incoming push
	//# uint64 cseqid, uint64 poolid
	YuntiNCPProto_PoolPullAssociate
	// Set pulling Option
	//# uint64 poolid ,uint64 pcseqid,uint64 optid, []byte value
	YuntiNCPProto_PoolSetOpt
	// Close a Pool and all connection associated with it
	//# uint64 poolid ,uint64 pcseqid
	YuntiNCPProto_PoolClose
)

func InterpretPacket(input chan []byte) error {

	for {

		select {
		case buffer := <-input:
			var operationTypeOpcode uint64
			packetBuffer := bytes.NewBuffer(buffer)
			erro := binary.Read(packetBuffer, binary.LittleEndian, &operationTypeOpcode)
			if erro != nil {
				return erro
			}

			go PacketProgress(operationTypeOpcode, packetBuffer)

			//default:
		}

	}

}

func PacketProgress(opCode, packetBuffer) {
	switch opCode {
	default:
		//unknow opcode
		//Verion not match or tranmission error

	}
}
