const keepAlive = () => {
    let ps = new WebSocket("ws://localhost:8080/ping")
    ps.onopen = () => console.log("Connected")
	ps.onmessage = m => console.log("Got:", m.data)
	ps.onclose = () => console.log("Connection closed")
}

keepAlive()

