package main

import "net"

type YuntiConnPool struct {
	TCPConns       *map[uint64](*net.Conn)
	socketidNext   uint64
	insocketidNext uint64
	SocketBuffers  map[uint64](*YuntiBufferStream)
}

func (pool *YuntiConnPool) Init() {
	pool.TCPConns = make(map[unint64](*net.Conn))
	pool.socketidNext = 1
}

func (pool *YuntiConnPool) OpenTCPConn(target string) uint64 {
	dialing, err := net.Dial("tcp", target)
	if err != nil {
		return 0
	}
	pool.TCPConns[pool.socketidNext] = dialing
	pool.InitStreamBuffer(pool.socketidNext)
	socketid := pool.socketidNext
	pool.socketidNext += 1
	return socketid
}
func (pool *YuntiConnPool) InitStreamBuffer(socketid uint64) {

}
func (pool *YuntiConnPool) DropTCPConn(socketid uint64) {
	pool.DestoryBuffer(socketid)
	pool.TCPConns[socketid].Close()
}

type YuntiBufferStream struct {
	StreamRxNext           uint64
	StreamTxNext           uint64
	StreamBufferRx         map[uint64](YuntiPacket)
	StreamBufferTx         map[uint64](YuntiPacket)
	streamid               uint64
	OutputStream           io.Writer
	OutputPutWorkerTxInput chan YuntiPacket
}

func (stream *YuntiBufferStream) Init() {

}

func (stream *YuntiBufferStream) InsertPacketTx(inserting YuntiPacket) {
	stream.OutputPutWorkerTxInput <- inserting
}

func (stream *YuntiBufferStream) OutputWorkerTx(input chan YuntiPacket) {
	for {
		select {
		case buffer := <-input:
			stream.StreamBufferTx[buffer.Seqid] = buffer
			for {
				if nextSend, nextSendOK := stream.StreamBufferTx[stream.StreamTxNext]; nextSendOK {
					_, err := stream.OutputStream.Write(nextSend.Payload)
					if err != nil {
						return err
					}
				} else {
					break
				}
				delete(stream.StreamBufferTx, stream.StreamTxNext)
				stream.StreamTxNext += 1
			}

		}
	}

}

type YuntiPacket struct {
	Payload []byte
	Seqid   uint64
}

type InterPoolSyncer struct {
}
