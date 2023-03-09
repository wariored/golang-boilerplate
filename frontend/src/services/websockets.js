// socketService.js
export default class SocketService {
  constructor(url) {
    this.socket = new WebSocket(url)
    this.socket.onopen = this.onOpen.bind(this)
    this.socket.onmessage = this.onMessage.bind(this)
    this.socket.onclose = this.onClose.bind(this)
  }

  onOpen() {
    console.log('WebSocket connected')
  }

  onMessage(event) {
    console.log('WebSocket message received: ', event.data)
  }

  onClose() {
    console.log('WebSocket closed')
  }

  send(message) {
    this.socket.send(message)
  }

  close() {
    this.socket.close()
  }
}
