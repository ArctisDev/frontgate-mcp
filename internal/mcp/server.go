package mcp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/ArctisDev/frontgate/internal/output"
	"github.com/ArctisDev/frontgate/internal/types"
)

type Server struct {
	logger *log.Logger
	tools  []Tool
}

type TransportMode int

const (
	TransportUnknown TransportMode = iota
	TransportRawJSON
	TransportFramed
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RespError  `json:"error,omitempty"`
}

type RespError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
		tools:  NewTools(),
	}
}

func (s *Server) Run(ctx context.Context, in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	decoder := json.NewDecoder(reader)
	writer := bufio.NewWriter(out)
	mode := TransportUnknown
	defer writer.Flush()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		payload, detectedMode, err := readMessage(reader, decoder, mode)
		if err != nil {
			if err == io.EOF {
				s.logger.Printf("stdin closed; shutting down")
				return nil
			}
			s.logger.Printf("read frame failed: %v", err)
			return err
		}
		if mode == TransportUnknown {
			mode = detectedMode
			s.logger.Printf("detected transport mode=%v", mode)
		}

		var req Request
		if err := json.Unmarshal(payload, &req); err != nil {
			s.logger.Printf("invalid request: %v", err)
			if err := writeFrame(writer, Response{
				JSONRPC: "2.0",
				Error:   &RespError{Code: -32700, Message: "parse error"},
			}); err != nil {
				return err
			}
			continue
		}

		s.logger.Printf("received method=%s has_id=%t", req.Method, len(req.ID) > 0)

		if len(req.ID) == 0 {
			if req.Method == "notifications/initialized" || req.Method == "initialized" {
				continue
			}
			s.logger.Printf("ignoring notification method=%s", req.Method)
			continue
		}

		id := decodeID(req.ID)
		resp := s.handleRequest(ctx, req, id)
		if err := writeResponse(writer, resp, mode); err != nil {
			return err
		}
	}
}

func (s *Server) handleRequest(ctx context.Context, req Request, id interface{}) Response {
	switch req.Method {
	case "initialize":
		return Response{
			JSONRPC: "2.0",
			ID:      id,
			Result: map[string]interface{}{
				"protocolVersion": "2025-06-18",
				"serverInfo": map[string]string{
					"name":    "frontgate",
					"version": "0.1.0",
				},
				"capabilities": map[string]interface{}{
					"tools": map[string]interface{}{
						"listChanged": false,
					},
				},
			},
		}
	case "ping":
		return Response{
			JSONRPC: "2.0",
			ID:      id,
			Result:  map[string]interface{}{},
		}
	case "tools/list":
		return Response{
			JSONRPC: "2.0",
			ID:      id,
			Result: map[string]interface{}{
				"tools": ListToolDefinitions(s.tools),
			},
		}
	case "tools/call":
		var params struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			return errorResponse(id, -32602, fmt.Sprintf("invalid tools/call params: %v", err))
		}

		tool, ok := findTool(s.tools, params.Name)
		if !ok {
			return errorResponse(id, -32601, "tool not found")
		}

		result, err := tool.Handler(ctx, params.Arguments)
		if err != nil {
			result = output.Error(err)
		}

		return Response{
			JSONRPC: "2.0",
			ID:      id,
			Result: map[string]interface{}{
				"content": []map[string]string{
					{
						"type": "text",
						"text": result.Text,
					},
				},
				"structuredContent": result.StructuredData,
				"isError":           result.IsError,
			},
		}
	default:
		return errorResponse(id, -32601, "method not found")
	}
}

func errorResponse(id interface{}, code int, message string) Response {
	return Response{
		JSONRPC: "2.0",
		ID:      id,
		Error:   &RespError{Code: code, Message: message},
	}
}

func readMessage(r *bufio.Reader, decoder *json.Decoder, mode TransportMode) ([]byte, TransportMode, error) {
	if mode == TransportRawJSON {
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil {
			return nil, TransportRawJSON, err
		}
		return raw, TransportRawJSON, nil
	}

	if err := skipWhitespace(r); err != nil {
		return nil, TransportUnknown, err
	}

	peek, err := r.Peek(1)
	if err != nil {
		return nil, TransportUnknown, err
	}

	if len(peek) == 1 && (peek[0] == '{' || peek[0] == '[') {
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil {
			return nil, TransportUnknown, err
		}
		return raw, TransportRawJSON, nil
	}

	body, err := readFrame(r)
	return body, TransportFramed, err
}

func readFrame(r *bufio.Reader) ([]byte, error) {
	headers := map[string]string{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header line: %q", line)
		}
		headers[strings.ToLower(strings.TrimSpace(parts[0]))] = strings.TrimSpace(parts[1])
	}

	cl, ok := headers["content-length"]
	if !ok {
		return nil, fmt.Errorf("missing Content-Length header")
	}
	size, err := strconv.Atoi(cl)
	if err != nil || size <= 0 {
		return nil, fmt.Errorf("invalid Content-Length: %q", cl)
	}

	body := make([]byte, size)
	if _, err := io.ReadFull(r, body); err != nil {
		return nil, err
	}
	return body, nil
}

func skipWhitespace(r *bufio.Reader) error {
	for {
		peek, err := r.Peek(1)
		if err != nil {
			return err
		}
		if len(peek) == 0 {
			return io.EOF
		}
		if !unicode.IsSpace(rune(peek[0])) {
			return nil
		}
		if _, err := r.ReadByte(); err != nil {
			return err
		}
	}
}

func writeFrame(w *bufio.Writer, resp Response) error {
	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body)))
	buf.Write(body)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}
	return w.Flush()
}

func writeRawJSON(w *bufio.Writer, resp Response) error {
	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := w.Write(body); err != nil {
		return err
	}
	if err := w.WriteByte('\n'); err != nil {
		return err
	}
	return w.Flush()
}

func writeResponse(w *bufio.Writer, resp Response, mode TransportMode) error {
	switch mode {
	case TransportRawJSON:
		return writeRawJSON(w, resp)
	default:
		return writeFrame(w, resp)
	}
}

func decodeID(raw json.RawMessage) interface{} {
	var id interface{}
	if err := json.Unmarshal(raw, &id); err != nil {
		return string(raw)
	}
	return id
}

func findTool(tools []Tool, name string) (Tool, bool) {
	for _, tool := range tools {
		if tool.Name == name {
			return tool, true
		}
	}
	return Tool{}, false
}

var _ = types.ToolResult{}
