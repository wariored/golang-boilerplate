import {createEffect, createSignal, For} from 'solid-js'
import SocketService from '../services/websockets'

function ChatComponent () {
  const [user, setUser] = createSignal(null)
  const [messages, setMessages] = createSignal([])
  const [users, setUsers] = createSignal([])
  const [socketService, setSocketService] = createSignal(null)

  createEffect(() => {
	  const newSocketService = new SocketService('ws://localhost:8080/ws/messages')
    setSocketService(newSocketService)
    return () => {
      newSocketService.close()
    }
  }, [])

  const handleUserClick = (e, selectedUser) => {
    setUser(selectedUser)
    socket.send('getMessages', selectedUser)
  }

  return () => (
    <div class="flex h-screen">
      <div class="w-1/4 bg-gray-200">
        <div class="p-4">
          <h2 class="text-lg font-medium">Users</h2>
		  <For each={users} fallback={<div>Loading...</div>}>
			  {(user) => {
              <li class="py-2 cursor-pointer" onClick={e => handleUserClick(e, user)}>
                {user.name}
				  </li>
			}}
          </For>
        </div>
      </div>
      <div class="w-1/2 bg-white">
        <div class="p-4">
          <h2 class="text-lg font-medium">Chat</h2>
          {user && <h3 class="text-sm font-medium">{user.name}</h3>}
          <div class="h-64 overflow-y-scroll">
			  <For each={messages} fallback={<div>Loading...</div>}>
				  {(message) => {
				  <div class="py-2">
					<p class="text-sm font-medium">{message.text}</p>
					<p class="text-xs text-gray-600">{message.timestamp}</p>
				  </div>
				  }}
		  </For>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ChatComponent
