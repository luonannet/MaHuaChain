package main

import (
	"Bamboo/component"
	"Bamboo/data"
)

func main() {
	machine := component.NewMachine()

	machine.PingNeighbor()
	go machine.JoinNet(initData(machine.IP))
	machine.Server.Peer.ListenAndServe()
}

//initData 携带本机数据加入网络
func initData(ip string) *data.NodeData {
	thisData := new(data.NodeData)
	thisData.IP = ip
	thisData.StartIndex = 30000
	thisData.EndIndex = 40000
	return thisData
}
