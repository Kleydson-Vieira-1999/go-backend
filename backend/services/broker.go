package services

import (
	"log"
)

type Client struct {
	EstablishmentID string
	MessageChan     chan string
}

type Broker struct {
	clients    map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan struct {
		EstablishmentID string
		Payload         string
	}
}

func NewBroker() *Broker {
	return &Broker{
		clients:    make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast: make(chan struct {
			EstablishmentID string
			Payload         string
		}),
	}
}

func (b *Broker) Start() {
	for {
		select {
		case client := <-b.register:
			if b.clients[client.EstablishmentID] == nil {
				b.clients[client.EstablishmentID] = make(map[*Client]bool)
			}
			b.clients[client.EstablishmentID][client] = true
			log.Printf("Novo cliente conectado na cozinha do estabelecimento: %s", client.EstablishmentID)

		case client := <-b.unregister:
			if _, ok := b.clients[client.EstablishmentID][client]; ok {
				delete(b.clients[client.EstablishmentID], client)
				close(client.MessageChan)
				if len(b.clients[client.EstablishmentID]) == 0 {
					delete(b.clients, client.EstablishmentID)
				}
				log.Printf("Cliente desconectado do estabelecimento: %s", client.EstablishmentID)
			}

		case msg := <-b.broadcast:
			if subs, ok := b.clients[msg.EstablishmentID]; ok {
				for client := range subs {
					select {
					case client.MessageChan <- msg.Payload:
					default:
						delete(subs, client)
						close(client.MessageChan)
					}
				}
			}
		}
	}
}

func (b *Broker) Broadcast(establishmentID string, jsonPayload string) {
	b.broadcast <- struct {
		EstablishmentID string
		Payload         string
	}{EstablishmentID: establishmentID, Payload: jsonPayload}
}

func (b *Broker) Register(c *Client)   { b.register <- c }
func (b *Broker) Unregister(c *Client) { b.unregister <- c }
