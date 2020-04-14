package care

import (
	"fmt"
	"time"

	"github.com/fanux/lvscare/service"
)

//VsAndRsCare is
func VsAndRsCare(vs string, rs []string, beat int64, path string, schem string) error {
	var lvs service.Lvser
	var err error
	fmt.Println("start lvscare.")
	lvs, err = service.BuildLvscare(vs, rs)
	if err != nil {
		fmt.Printf("new lvs failed %s\n",err)
		return err
	}

	t := time.NewTicker(time.Duration(beat) * time.Second)
	for {
		select {
		case <-t.C:
			//check virturl server
			virs, _:= lvs.GetVirtualServer()
			if virs == nil {
				lvs.Close()
				lvs, err = service.BuildLvscare(vs, rs)
				if err != nil {
					fmt.Printf("new lvs failed %s\n",err)
					return err
				}
				err = lvs.CreateVirtualServer()
				if err != nil {
					fmt.Println("create vs failed", err)
				}
			}

			//check real server
			lvs.CheckRealServers(path, schem)
		}
	}

	return nil
}
