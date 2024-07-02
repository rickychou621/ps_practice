package webserver

type Hub struct {
	client     map[*Client]bool
	broadcast  chan []byte  //廣播用channel
	register   chan *Client //註冊連入的client
	unregister chan *Client //刪除斷線的client
}

func newHub() *Hub {
	return &Hub{
		client:     make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}
