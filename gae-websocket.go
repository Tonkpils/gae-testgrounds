package gaewebsocket

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

var externalIPURL = "http://metadata/computeMetadata/v1beta1/instance/network-interfaces/0/access-configs/0/external-ip"

func init() {
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/ping", pingHandler)
	http.Handle("/ws", websocket.Handler(echoServer))
}

func getHostname() (string, error) {
	resp, err := http.Get(externalIPURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	name, err := getHostname()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	wsInfo := struct {
		WebSocketURL string
	}{
		WebSocketURL: "ws://" + name + ":8080" + "/ws",
	}
	t, err := template.ParseFiles("gae.html")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if err := t.Execute(w, wsInfo); err != nil {
		fmt.Fprint(w, err)
		return
	}
}

func echoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PONG")
}
