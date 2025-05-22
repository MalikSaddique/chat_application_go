import { useState} from 'react';
import Login from './pages/login';
import SignUp from './pages/sign_up';
import Chat from './pages/chat';

export default function App() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [showSignup, setShowSignup] = useState(false);

  if (!loggedIn) {
    return showSignup ? (
      <>
        <SignUp />
        <button id='signup' onClick={() => setShowSignup(false)}>Already have an account?</button>
      </>
    ) : (
      <>
        <Login onLogin={() => setLoggedIn(true)} />
        <button id='login' onClick={() => setShowSignup(true)}>New user? Sign Up</button>
      </>
    );
  }

  return <Chat onLogout={()=> setLoggedIn(false)}/>;
}
