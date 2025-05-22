import { useState } from 'react';
import { login } from '../auth';
import './css/login_signup.css'


export default function Login({ onLogin }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await login(email, password);
      if (res.data.access_token) {
        localStorage.setItem('accessToken', res.data.access_token);
        localStorage.setItem('refreshToken', res.data.refresh_token);
        onLogin(); 
      } else {
        console.log('No token received');
      }

  } catch (error) {
    console.error('Login error:', error.res?.data || error.message);
  }
  };

  return (
    <div className="wrapper signIn">
    <div className="form">
    <form onSubmit={handleLogin}>
    <div className="heading">LOGIN</div>
           <div>
						<label htmlFor="name">E-MAIL</label>
						<input type="email" id="e-mail" placeholder="Enter your email" value={email} onChange={e => setEmail(e.target.value)} />
					</div>
          <div>
						<label htmlFor="password">Password</label>
						<input type="password" id="password" placeholder="Enter your password" value={password} onChange={e => setPassword(e.target.value)} />
					</div>
          <button type="submit" >
						Submit
					</button>
    </form>
    </div>
    </div>

  );
}
