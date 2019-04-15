package broker

var brokers = map[string]Broker{}

// Register Register
func Register(b Broker) {
	if _, ok := brokers[b.String()]; ok {

	}
}

// NewBroker NewBroker
func NewBroker(name string) Broker {

}
