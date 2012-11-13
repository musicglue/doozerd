package gc

import (
	"github.com/musicglue/doozerd/consensus"
	"github.com/musicglue/doozerd/store"
	"log"
	"strconv"
	"time"
)

func Pulse(node string, seqns <-chan int64, p consensus.Proposer, sleep int64) {
	path := "/ctl/node/" + node + "/applied"
	for {
		seqn, ok := <-seqns
		if !ok {
			break
		}

		e := consensus.Set(p, path, []byte(strconv.FormatInt(seqn, 10)), store.Clobber)
		if e.Err != nil {
			log.Println(e.Err)
		}

		time.Sleep(time.Duration(sleep))
	}
}
