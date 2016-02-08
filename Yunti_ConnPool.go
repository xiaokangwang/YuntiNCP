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
	pool.TCPConns[socketidNext] = dialing
	pool.InitStreamBuffer(socketidNext)
	return socketidNext
}
func (pool *YuntiConnPool) InitStreamBuffer(socketid uint64) {

}
func (pool *YuntiConnPool) DropTCPConn(socketid uint64) {
	pool.DestoryBuffer(socketid)
	pool.TCPConns[socketid].Close()
}

type YuntiBufferStream struct {
	StreamRxNext   uint64
	StreamTxNext   uint64
	StreamBufferRx map[uint64]([]byte)
	StreamBufferTx map[uint64]([]byte)
  streamid       uint64
  OutputStream   io.Writer
}

func(*stream YuntiBufferStream)Init(){

}

func(*stream YuntiBufferStream)InsertPacket(){

}

func(*stream YuntiBufferStream)OutputWorkerTx(input chan YuntiPacket){
  for{
    select{
    case buffer:<-input:
      stream.StreamBufferTx[buffer.Seqid]=buffer.Payload
    }
  }

}

type YuntiPacket struct{
  Payload []byte
  Seqid uint64
}
