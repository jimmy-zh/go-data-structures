package graph

import (
	"os"
	"bufio"
	"strings"
	"errors"
	"io"
	"fmt"
)

type UdSymbolGraph struct {
	udg  *UdGraph
	st   map[string]int
	keys []string
}

func NewUdSymbolGraph(stream string, split string) (udsg *UdSymbolGraph, err error) {
	udsg = &UdSymbolGraph{
		st:   make(map[string]int),
		keys: make([]string, 0),
	}

	fd, err := os.OpenFile(stream, os.O_RDONLY, 0)
	if err != nil {
		err = errors.New(fmt.Sprintf("open file error:%s\n", err.Error()))
		return
	}
	defer fd.Close()

	rd := bufio.NewReader(fd)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.New("read error")
		}
		edges := strings.Split(line, split)
		for _, edge := range edges {
			if _, ok := udsg.st[edge]; !ok {
				udsg.keys = append(udsg.keys, edge)
				udsg.st[edge] = len(udsg.st)
			}
		}
	}

	udsg.udg = NewUdGraph(len(udsg.keys))
	//reset fd
	fd.Seek(0, 0)
	//reset rd
	rd.Reset(fd)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.New("read error")
		}
		edges := strings.Split(line, split)
		for i := 1; i < len(edges); i++ {
			udsg.udg.AddEdge(udsg.st[edges[0]], udsg.st[edges[i]])
		}
	}
	return
}

func (udsg *UdSymbolGraph) Contains(edge string) bool {
	_, ok := udsg.st[edge]
	return ok
}

func (udsg *UdSymbolGraph) Index(edge string) int {
	return udsg.st[edge]
}

func (udsg *UdSymbolGraph) Name(v int) string {
	if v < len(udsg.keys) {
		return udsg.keys[v]
	}
	return ""
}

func (udsg *UdSymbolGraph) UdGraph() *UdGraph {
	return udsg.udg
}
