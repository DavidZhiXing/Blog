type server struct {
	host string
	port string
	debuglevel string
	banner string
	password string
	rooms roomMap
}

type roomMap struct {
	rooms map[string]*roomInfo
	sync.RWMutex
}

type roomInfo struct {
	firstUser *comm.comm
	secondUser *comm.comm
	opened time.Time
	full bool
}

func Run(debuglevel, host, port, password string, banner ...string) (err error) {
	var s server
	s.host = host
	s.port = port
	s.debuglevel = debuglevel
	s.password = password
	s.rooms = roomMap{rooms: make(map[string]*roomInfo)}
	if len(banner) > 0 {
		s.banner = banner[0]
	} else {
		s.banner = "Welcome to the server"
	}

	return s.start()
}

func (tart(s *server) (err error) {
	log.SetLevel(log.DebugLevel)
	log.Debugf("Starting server on %s:%s", s.host, s.port)
	s.rooms.Lock()
	s.rooms.rooms=make(map[string]*roomInfo)
	s.rooms.Unlock()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			var roomsToDelete []string
			s.rooms.Lock()
			for room := range s.rooms.rooms {
				if time.Since(s.rooms.rooms[room].opened) > time.Second * 60 {
					roomsToDelete = append(roomsToDelete, room)
				}
			}
			s.rooms.Unlock()

			for _, room := range roomsToDelete {
				s.deleteRoom(room)
			}
		}
	}()

	err = s.run()]
	if err != nil {
		log.Errorf("Error running server: %s", err)
	}
	return
}

func (s *server) run() (err error) {
	network := "tcp"
	addr := net.JoinHostPort(s.host, s.port)
	if s.host != "" {
		ip := net.ParseIP(s.host)
		if ip == nil {
			tcpIp, err := net.ResolveIPAddr("ip", s.host)
			if err != nil {	
				return err
			}
			ip = tcpIp.IP
		}

		addr = net.JoinHostPort(ip.String(), s.port)
		if s.host != "" {
			if ip.To4() != nil {
				network = "tcp4"
			} else {
				network = "tcp6"
			}
		}

		
	}
	addr = strings.Replace(addr, "127.0.0.1", "0.0.0.0", 1)
	log.Infof("Listening on %s", addr)
	server, err := net.Listen(network, addr)
	if err != nil {
		return fmt.Errorf("Error listening: %s", err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Errorf("Error accepting: %s", err)
			continue
		}
		go s.handleRequest(conn)
	}