package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Girl01star/learning/lesson_12/internal/documentstore"
	"github.com/Girl01star/learning/lesson_12/internal/proto"
)

func main() {
	store := documentstore.NewStore()

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("server started on 127.0.0.1:8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handleConn(conn, store)
	}
}

func handleConn(conn net.Conn, store *documentstore.Store) {
	defer conn.Close()
	fmt.Println("client connected:", conn.RemoteAddr())

	in := bufio.NewScanner(conn)
	in.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	out := bufio.NewWriter(conn)
	defer out.Flush()

	for in.Scan() {
		line := in.Bytes()

		var req proto.Request
		if err := json.Unmarshal(line, &req); err != nil {
			writeResp(out, proto.Response{Ok: false, Error: "bad json: " + err.Error()})
			continue
		}

		resp := process(store, req)
		writeResp(out, resp)
	}

	if err := in.Err(); err != nil {
		fmt.Println("conn read error:", err)
	}
	fmt.Println("client disconnected:", conn.RemoteAddr())
}

func writeResp(w *bufio.Writer, resp proto.Response) {
	b, _ := json.Marshal(resp)
	b = append(b, '\n')
	_, _ = w.Write(b)
	_ = w.Flush()
}

func process(store *documentstore.Store, req proto.Request) proto.Response {
	switch req.Cmd {

	case "ping":
		return proto.Response{Ok: true, Data: map[string]any{"time": time.Now().Format(time.RFC3339)}}

	case "create_collection":
		cfg := &documentstore.CollectionConfig{}
		if req.PrimaryKey != "" {
			cfg.PrimaryKey = req.PrimaryKey
		}
		_, err := store.CreateCollection(req.Collection, cfg)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true}

	case "delete_collection":
		if err := store.DeleteCollection(req.Collection); err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true}

	case "get_collection":
		_, err := store.Collection(req.Collection)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true}

	case "put":
		col, err := store.Collection(req.Collection)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		if err := col.Put(req.Doc); err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true}

	case "get":
		col, err := store.Collection(req.Collection)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		doc, err := col.Get(req.Key)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true, Data: doc}

	case "delete":
		col, err := store.Collection(req.Collection)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		if err := col.Delete(req.Key); err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true}

	case "list":
		col, err := store.Collection(req.Collection)
		if err != nil {
			return proto.Response{Ok: false, Error: err.Error()}
		}
		return proto.Response{Ok: true, Data: col.List()}

	case "logs":
		return proto.Response{Ok: true, Data: store.Logs()}

	default:
		return proto.Response{Ok: false, Error: "unknown cmd: " + req.Cmd}
	}
}
