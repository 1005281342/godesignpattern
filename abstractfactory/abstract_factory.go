package abstractfactory

import "log"

// IAbstractFactory ...
type IAbstractFactory interface {
	CreateProxy() Proxy
	CreateNotify() Notify
}

func AbstractFactory() IAbstractFactory {
	return &demo{}
}

type demo struct{}

func (d *demo) CreateProxy() Proxy {
	return NewRPCProxy()
}

func (d *demo) CreateNotify() Notify {
	return NewNotify()
}

// Proxy ...
type Proxy interface {
	Call() error
}

// Notify ...
type Notify interface {
	Exec() error
}

type rpcProxy struct{}

func NewRPCProxy() Proxy {
	return &rpcProxy{}
}

func (r *rpcProxy) Call() error {
	log.Println("this is rpc proxy")
	return nil
}

type notify struct{}

func NewNotify() Notify {
	return &notify{}
}

func (n *notify) Exec() error {
	log.Println("notify hhh!!")
	return nil
}
