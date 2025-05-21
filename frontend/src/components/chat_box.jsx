import { useEffect, useState, useRef } from 'react';
import axiosInstance from '../utils/axiosInstance';
import { WS_URL } from '../config';
import UserSearch from './search_user';

export default function ChatBox() {
  const [messages, setMessages] = useState([]);
  const [newMsg, setNewMsg] = useState('');
  const [receiverId, setReceiverId] = useState('');
  const ws = useRef(null);

  useEffect(() => {
    const token = localStorage.getItem('accessToken');
    if (!token) return;

    const wsUrl = `${WS_URL}?token=${token}`;
    ws.current = new WebSocket(wsUrl);

    ws.current.onopen = () => {
      console.log('WebSocket connected');
      ws.current.send(JSON.stringify({ action: 'ping', message: 'hello from client' }));
      ws.current.send(JSON.stringify({ action: 'ping', message: 'hi' }));
    };

    ws.current.onmessage = (event) => {
      console.log("WebSocket raw message:", event.data);
      const receivedMessage = JSON.parse(event.data);
      console.log("WebSocket message received:", event.data);
      setMessages((prev) => [...prev, { ...receivedMessage, timestamp: new Date().toLocaleTimeString() }]);
    };

    ws.current.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.current.onclose = () => {
      console.log('WebSocket closed');
    };

    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, []);

  const fetchMessages = async () => {
    try {
      const res = await axiosInstance.post('/protected/message', { receiver_id: parseInt(receiverId) }, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        }
      });
      setMessages(res.data.map(msg => ({ ...msg, timestamp: new Date().toLocaleTimeString() })));
    } catch (error) {
      console.error('Error fetching messages:', error);
    }
  };

  const handleSend = async () => {
    if (newMsg.trim() === '' || String(receiverId).trim() === '') return;

    const msgPayload = {
      receiver_id: parseInt(receiverId),
      message: newMsg
    };
    console.log("Sending message:", msgPayload);

    try {
      await axiosInstance.post('/protected/send', msgPayload, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        }
      });
      function parseJwt(token) {
        try {
          return JSON.parse(atob(token.split('.')[1]));
        } catch (e) {
          return null;
        }
      }
      const senderId = parseJwt(localStorage.getItem('accessToken'))?.user_id;

      setMessages(prev => [...prev, { ...msgPayload, sender_id: senderId, timestamp: new Date().toLocaleTimeString() }]);
      setNewMsg('');
    } catch (err) {
      console.error('Error sending message:', err);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('accessToken');
    window.location.reload();
  };

  return (
    <div className='container'>
      <div className='chat-header'>
        <h2>ChatBox</h2>
      </div>

      <UserSearch onUserSelect={(id) => setReceiverId(id)} />

      <div className='message-list'>
        {messages.map((msg, i) => {
          function parseJwt(token) {
            try {
              return JSON.parse(atob(token.split('.')[1]));
            } catch (e) {
              return null;
            }
          }
          const currentUserId = parseJwt(localStorage.getItem('accessToken'))?.user_id;
          const isSent = msg.sender_id === currentUserId;

          return (
            <div key={i} className={`message-item ${isSent ? 'sent' : 'received'}`}>
              <strong>From {msg.sender_id} to {msg.receiver_id}:</strong> {msg.message}
              <span className="message-timestamp">{msg.timestamp || '2:44PM'}</span>
            </div>
          );
        })}
      </div>

      <div className='message-input-container'>
        <input
          className='message-input'
          placeholder="Send a message"
          value={newMsg}
          onChange={(e) => setNewMsg(e.target.value)}
        />
        <button className='send-button' onClick={handleSend}>Send</button>
      </div>

      <button className='logout-button' onClick={handleLogout}>Logout</button>
    </div>
  );
}