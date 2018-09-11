package udp

import (
	"encoding/json"
	"fmt"
	"kagami/model"
	"log"
	"net"
	"time"
)

func Loop(localAddress string) {
	server := model.UDPServer{}
	err := createServer(localAddress, &server)
	if err != nil {
		log.Fatal("ERROR: can not create server", localAddress, err)
		return
	}

	log.Println("INFO : Kagami UDP server start at", localAddress)
	for {
		req, addr, err := readReq(&server)
		if err != nil {
			log.Println("ERROR: can not read request", err)
		} else {
			go handleClient(&server, req, addr)
		}
	}
}

func createServer(localAddress string, server *model.UDPServer) error {
	udpAddr, err := net.ResolveUDPAddr("udp4", localAddress)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	addrMap := make(map[string]*model.UDPUser)

	server.LocalAddress = udpAddr
	server.Connection = conn
	server.AddressBook = addrMap

	return nil
}

func readReq(server *model.UDPServer) (*model.ReqUDP, *net.UDPAddr, error) {
	var buf [2048]byte
	req := model.ReqUDP{}

	n, addr, err := server.Connection.ReadFromUDP(buf[0:])
	if err != nil {
		return &req, addr, err
	}

	err = json.Unmarshal(buf[:n], &req)
	if err != nil {
		return &req, addr, err
	}

	return &req, addr, nil
}

func handleClient(server *model.UDPServer, req *model.ReqUDP, addr *net.UDPAddr) {
	rsp := model.RspUDP{}

	switch req.Type {
	case "REG":
		argReg, ok := req.Data.(model.ArgRegister)
		if !ok {
			rsp.Ok = false
			rsp.Desc = "REG request format error"
			rsp.Data = nil
			log.Println("ERROR:", rsp.Desc, req.Data)
			break
		}

		user, ok := server.AddressBook[argReg.Name]
		if ok && user.UDPAddr == addr {
			rsp.Ok = false
			rsp.Desc = fmt.Sprintf("user[%s] is already registed by %s", user.Name, user.UDPAddr)
			rsp.Data = nil
			log.Println("ERROR:", rsp.Desc)
		} else {
			now := time.Now()
			newUser := model.UDPUser{Name: argReg.Name, UDPAddr: addr, RegisterTime: &now}
			server.AddressBook[argReg.Name] = &newUser
			rsp.Ok = true
			rsp.Desc = "ok"
			rsp.Data = model.RetRegister{UDPAddr: addr}
			log.Println("INFO : new user registered", newUser)
		}
	case "LIST":
		argList, ok := req.Data.(model.ArgListNeighbor)
		if !ok {
			rsp.Ok = false
			rsp.Desc = "LIST request format error"
			rsp.Data = nil
			log.Println("ERROR:", rsp.Desc, req.Data)
			break
		}

		user, ok := server.AddressBook[argList.Name]
		if !ok {
			rsp.Ok = false
			rsp.Desc = fmt.Sprintf("can not found user %s", argList.Name)
			rsp.Data = nil
			log.Println("ERROR:", rsp.Desc)
		} else {
			data := model.RetListNeighbor{Neighbors: findNeighbors(server, user)}
			rsp.Ok = true
			rsp.Desc = "ok"
			rsp.Data = data
			log.Println("INFO :", fmt.Sprintf("find %d neighbors for %s", len(data.Neighbors), argList.Name))
		}
	}

	jsonRsp, err := json.Marshal(&rsp)
	if err != nil {
		log.Println("ERROR: fail to parse response", err)
		return
	}
	_, err = server.Connection.WriteToUDP(jsonRsp, addr)
	if err != nil {
		log.Println("ERROR: IO error when sent response to", addr, err)
	}
}

func findNeighbors(server *model.UDPServer, user *model.UDPUser) []*model.UDPUser {
	// TODO: impl actual neighbor search algorithm
	neighbors := make([]*model.UDPUser, 0, len(server.AddressBook))
	for _, value := range server.AddressBook {
		neighbors = append(neighbors, value)
	}
	return neighbors
}
