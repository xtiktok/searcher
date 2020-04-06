package notify

import (
	"errors"
	"fmt"
	"searcher/common/consts"
	"searcher/common/dto"
	"sync"
)

type NotifyStore struct {
	sync.Map
}

var store *NotifyStore

func init() {
	store = &NotifyStore{}
}

func DoWatch(args []string, conn *dto.TsConn) error {
	init := []*dto.TsConn{conn}
	actual, ok := store.LoadOrStore(fmt.Sprintf("%s.%s", args[0], args[1]), init)
	if ok {
		return nil
	}
	list, ok := actual.([]*dto.TsConn)
	if !ok {
		return errors.New(consts.ErrorInternalError)
	}
	list = append(list, conn)
	store.Store(fmt.Sprintf("%s.%s", args[0], args[1]), list)
	return nil
}

func DoNotify(key string, operate string) {
	actual, ok := store.Load(fmt.Sprintf("%s.%s", key, operate))
	if !ok {
		return
	}

	list, ok := actual.([]*dto.TsConn)
	if !ok {
		return
	}

	for i := range list {
		fmt.Println(fmt.Sprintf("start notify %d", i))
		//_, err := list[i].Conn.Write([]byte(fmt.Sprintf("%s %s", key, operate)))
		//if err != nil {
		//	list[i].Conn.Close()
		//	list[i] = nil
		//	continue
		//}
	}

}
