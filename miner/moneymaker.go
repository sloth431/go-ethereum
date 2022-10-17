package miner

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type Moneymaker struct {
}

func NewMoneyMaker() *Moneymaker {
	return &Moneymaker{}
}

func (mm *Moneymaker) preCheck() {

}

func (mm *Moneymaker) parseLogs(env *environment, logs []*types.Log, tx *types.Transaction) {
	for _, item := range logs {
		//log.Println(item.Address)
		log.Trace("asdf", item.Address)
	}
}

func (mm *Moneymaker) calculateProfit() {

}

func (mm *Moneymaker) sandwitch() {

}

func (mm *Moneymaker) afterTx() {

}

func (mm *Moneymaker) CheckTx(env *environment, logs []*types.Log, tx *types.Transaction) {
	// 0. checks and lock
	mm.preCheck()
	// 1. parse uniswap V2, V3 logs
	mm.parseLogs(env, logs, tx)
	// 2. calc profit
	mm.calculateProfit()
	// 3. send front back tx
	mm.sandwitch()
	// 4. watch out and alert profit
	mm.afterTx()
}
