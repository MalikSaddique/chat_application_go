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
      ws.current.send(JSON.stringify({
        action: 'ping',
        message: 'hello from client'
      }));
      ws.current.send(JSON.stringify({
        action: 'ping',
        message: 'hi',
      }));
      
    };
  
    ws.current.onmessage = (event) => {
      console.log("WebSocket raw message:", event.data);
      const receivedMessage = JSON.parse(event.data);
      console.log("WebSocket message received:", event.data);
      setMessages((prev) => [...prev, receivedMessage]);
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
      const res = await axiosInstance.post('/protected/message', msgPayload, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        }
      });
      
      setMessages(res.data);
    } catch (error) {
      console.error('Error fetching messages:', error);
    }
  };

  const handleSend = async() => {
    if (newMsg.trim() === '' || receiverId.trim() === '') return;
 
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

      setMessages(prev => [...prev, { ...msgPayload, sender_id: senderId }]);
      setNewMsg('');
    } catch (err) {
      console.error('Error sending message:', err);
    }
  };

  return (
    <div className='container'>
      <h2>📨 ChatBox</h2>
      <UserSearch onUserSelect={(id) => setReceiverId(id)} />
      <h3>Send Message</h3>
      <input
        placeholder="Receiver ID"
        value={receiverId}
        onChange={(e) => setReceiverId(e.target.value)}
        style={{ width: '200px', marginRight: '10px' }}
      />
      <input
        placeholder="Type a message"
        value={newMsg}
        onChange={(e) => setNewMsg(e.target.value)}
        style={{ width: '300px', marginRight: '10px' }}
      />
      <button onClick={handleSend}>Send</button>

      <hr />
      <h3>🗨️ Messages</h3>
      <div className='message-list'>
        {messages.map((msg, i) => (
          <div key={i} className='message-item'>
            <strong>From {msg.sender_id} to {msg.receiver_id}:</strong> {msg.message}
          </div>
        ))}

      </div>
    </div>
  );
}
