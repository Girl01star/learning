package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Girl01star/learning/lesson_12/internal/documentstore"
	"github.com/Girl01star/learning/lesson_12/internal/proto"
)

func main() {
	addr := "127.0.0.1:8080"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("connected to", addr)
	fmt.Println("commands:")
	fmt.Println("  ping")
	fmt.Println("  create_collection <name> [primaryKey]")
	fmt.Println("  delete_collection <name>")
	fmt.Println("  get_collection <name>")
	fmt.Println("  put <collection> <jsonDoc>")
	fmt.Println("  get <collection> <key>")
	fmt.Println("  delete <collection> <key>")
	fmt.Println("  list <collection>")
	fmt.Println("  logs")
	fmt.Println("  exit")

	in := bufio.NewScanner(os.Stdin)

	srv := bufio.NewReader(conn)

	out := bufio.NewWriter(conn)

	for {
		fmt.Print("> ")
		if !in.Scan() {
			fmt.Println()
			return
		}
		line := strings.TrimSpace(in.Text())
		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" {
			return
		}

		req, err := parseCommand(line)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		b, err := json.Marshal(req)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		b = append(b, '\n')

		if _, err := out.Write(b); err != nil {
			fmt.Println("write error:", err)
			return
		}
		if err := out.Flush(); err != nil {
			fmt.Println("flush error:", err)
			return
		}

		respLine, err := srv.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		var resp proto.Response
		if err := json.Unmarshal([]byte(respLine), &resp); err != nil {
			fmt.Println("bad response json:", err)
			fmt.Println("raw:", respLine)
			continue
		}

		printResponse(resp)
	}
}

func printResponse(resp proto.Response) {
	if !resp.Ok {
		fmt.Println("❌", resp.Error)
		return
	}
	if resp.Data == nil {
		fmt.Println("✅ ok")
		return
	}

	pretty, err := json.MarshalIndent(resp.Data, "", "  ")
	if err != nil {
		fmt.Println("✅ ok:", resp.Data)
		return
	}
	fmt.Println("✅ ok")
	fmt.Println(string(pretty))
}

func parseCommand(line string) (proto.Request, error) {
	parts := strings.Fields(line)
	cmd := parts[0]

	switch cmd {
	case "ping":
		return proto.Request{Cmd: "ping"}, nil

	case "create_collection":
		if len(parts) < 2 {
			return proto.Request{}, fmt.Errorf("usage: create_collection <name> [primaryKey]")
		}
		req := proto.Request{
			Cmd:        "create_collection",
			Collection: parts[1],
		}
		if len(parts) >= 3 {
			req.PrimaryKey = parts[2]
		}
		return req, nil

	case "delete_collection":
		if len(parts) != 2 {
			return proto.Request{}, fmt.Errorf("usage: delete_collection <name>")
		}
		return proto.Request{
			Cmd:        "delete_collection",
			Collection: parts[1],
		}, nil

	case "get_collection":
		if len(parts) != 2 {
			return proto.Request{}, fmt.Errorf("usage: get_collection <name>")
		}
		return proto.Request{
			Cmd:        "get_collection",
			Collection: parts[1],
		}, nil

	case "put":
		rest := strings.TrimSpace(line[len("put"):])
		firstSpace := strings.Index(rest, " ")
		if firstSpace == -1 {
			return proto.Request{}, fmt.Errorf("usage: put <collection> <jsonDoc>")
		}
		collection := strings.TrimSpace(rest[:firstSpace])
		jsonDoc := strings.TrimSpace(rest[firstSpace+1:])
		if collection == "" || jsonDoc == "" {
			return proto.Request{}, fmt.Errorf("usage: put <collection> <jsonDoc>")
		}

		var doc documentstore.Document
		if err := json.Unmarshal([]byte(jsonDoc), &doc); err != nil {
			return proto.Request{}, fmt.Errorf("bad doc json: %w", err)
		}

		return proto.Request{
			Cmd:        "put",
			Collection: collection,
			Doc:        doc,
		}, nil

	case "get":
		if len(parts) != 3 {
			return proto.Request{}, fmt.Errorf("usage: get <collection> <key>")
		}
		return proto.Request{
			Cmd:        "get",
			Collection: parts[1],
			Key:        parts[2],
		}, nil

	case "delete":
		if len(parts) != 3 {
			return proto.Request{}, fmt.Errorf("usage: delete <collection> <key>")
		}
		return proto.Request{
			Cmd:        "delete",
			Collection: parts[1],
			Key:        parts[2],
		}, nil

	case "list":
		if len(parts) != 2 {
			return proto.Request{}, fmt.Errorf("usage: list <collection>")
		}
		return proto.Request{
			Cmd:        "list",
			Collection: parts[1],
		}, nil

	case "logs":
		return proto.Request{Cmd: "logs"}, nil

	default:
		return proto.Request{}, fmt.Errorf("unknown command: %s", cmd)
	}
}
