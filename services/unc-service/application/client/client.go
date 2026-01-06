package client

type Clients struct {
	EmailClient EmailClient
}

func SetupClients() *Clients {
	return &Clients{
		EmailClient: NewEmailClient(),
	}
}
