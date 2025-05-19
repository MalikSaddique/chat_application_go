import ChatBox from '../components/chat_box';

export default function Chat({onLogout}) {
  return (
    <div>
      <h2>Chat Room</h2>
      <ChatBox />
      <button onClick={() => {
        localStorage.removeItem('accessToken'); 
        onLogout(); 
      }}>
        Logout
      </button>
    </div>
  );
}
