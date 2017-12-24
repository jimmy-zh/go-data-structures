package use_case_api

import (
	"os"
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/data-structures/graph"
)

//testing fails now,to be completed
func TestOutputPath(t *testing.T) {
	udsg, err := graph.NewUdSymbolGraph(os.Getenv("GOPATH")+"/graph_test", " ")
	if err != nil {
		t.Fatalf("NewUdSymbolGraph error:%s", err.Error())
	}

	dos := NewDegreesOfSeparation(udsg, "JFK")
	if dos == nil {
		t.Fatal("NewDegreesOfSeparation error")
	}

	dos.OutputPath("LAS")

}
