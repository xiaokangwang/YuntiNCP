package main

import "io"
import "encoding/binary"
import "bytes"

func StreamSpilit(output chan []byte, input io.Reader, cancel chan int) error {
	for {
		select {
		case <-cancel:
			return nil
		default:
			var nextReadN uint64
			erro := binary.Read(input, binary.LittleEndian, &nextReadN)
			if erro != nil {
				return erro
			}
			buffer := make([]byte, nextReadN)
			_, erro = io.ReadFull(input, buffer)
			if erro != nil {
				return erro
			}
			output <- buffer
		}
	}
}

func StreamConcrete(input chan []byte, output io.Writer, cancel chan int) error {
	for {
		select {
		case <-cancel:
			return nil
		default:
			var nextWriteN uint64
			buffer <- input
			nextWriteN = len(buffer)
			erro := binary.Write(output, binary.LittleEndian, nextWriteN)
			if erro != nil {
				return erro
			}
			writeBuffer := bytes.NewBuffer(buffer)
			_, erro = io.Copy(output, writeBuffer)
			if erro != nil {
				return erro
			}
		}
	}
}
