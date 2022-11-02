1package main

import (
    "bytes"
    "fmt"
    "strings"

    xj "github.com/basgys/goxml2json"
)

const prefix = "veryuniqueattrprefix-"

func main() {
    xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?>
    <osm version="0.6" generator="CGImap 0.0.2">
     <bounds minlat="54.0889580" minlon="12.2487570" maxlat="54.0913900" maxlon="12.2524800"/>
     <foo>bar</foo>
    </osm>`)

    /*
    // Decode XML document
    root := &xj.Node{}
    err := xj.NewDecoder(
        xml,
        xj.WithAttrPrefix(prefix),
    ).Decode(root)
    if err != nil {
	fmt.Println("error decoding XML:",err)
        panic(err)
    }

    RemoveAttr(root)

    // Then encode it in JSON
    buf := new(bytes.Buffer)
    e := xj.NewEncoder(buf)
    err = e.Encode(root)
    if err != nil {
	fmt.Println("error encoding JSON:",err)
        panic(err)
    }

    fmt.Println(buf.String())
	*/
    json, err := xj.Convert(xml)
    if err != nil {
        panic("ERROR converting xml to json")
    }

    fmt.Println(json.String())
}

func RemoveAttr(n *xj.Node) {
    for k, v := range n.Children {
        if strings.HasPrefix(k, prefix) {
            delete(n.Children, k)
        } else {
            for _, n := range v {
                RemoveAttr(n)
            }
        }
    }
}
