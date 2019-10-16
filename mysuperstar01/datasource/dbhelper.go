package datasource


import (
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex  //(单利模式的时候用到锁)
)

func InstanceMaster() *xorm.Engine {
	//Todo.....
	return masterEngine
}

func Instanceslave() *xorm.Engine {
	//Todo.....
	return slaveEngine
}

