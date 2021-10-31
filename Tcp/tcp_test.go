package tcp

import (
	"fmt"
	"time"
	"testing"
	"byte"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

func BenchmarkConnection(b *testing.B) {
		log.SetLevel("trace")
		go Run("debug", "localhost", "8283", "pass123", "8284")
		time.Sleep(time.Second)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
				conn, err := Connect("localhost", "8284", "pass123")
				if err != nil {
						b.Error(err)
				}
				conn.Close()
		}
}

func TestTcp(t *testing.T) {
		log.SetLevel("error")
		timeToRoomDeletion = time.Second * 10
		go Run("debug", "localhost", "8283", "pass123", "8284")
		time.Sleep(time.Second)
		err := PingServer(timeToRoomDeletion)
		assert.Nil(t, err)
		
		conn, err := Connect("localhost", "8284", "pass123")
		if err != nil {
				t.Error(err)
		}
		conn.Close()
}