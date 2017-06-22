package collectors

//import (
//	//"fmt"
//)

type RedisConnection struct {
}

func (self *RedisConnection) ReadInfo() (b []byte, e error) {
	//	if len(self.Raw) == 0 {
	//		var con net.Conn
	//		logger.Printf("connecting %s", self.Address)
	//		con, e = net.Dial("tcp", self.Address)
	//		if e != nil {
	//			logger.Printf("ERROR connecting %s: %q", self.Address, e.Error())
	//			return
	//		}
	//		defer con.Close()

	//		con.Write([]byte("INFO\r\n"))
	//		b = make([]byte, 4096)
	//		var i int
	//		i, e = con.Read(b)
	//		dbg.Printf("read %d bytes", i)
	//		if e != nil {
	//			logger.Printf("ERROR reading: %q", e.Error())
	//			return
	//		}
	//		self.Raw = b
	//	}
	//	b = self.Raw
	return
}
