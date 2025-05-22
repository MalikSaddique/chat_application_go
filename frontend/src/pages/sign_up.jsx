import { useState } from 'react';
import { signup } from '../auth';
import "../App.css"

export default function SignUp() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSignup = async (e) => {
    e.preventDefault();
    try {
      await signup(email, password);
      alert("Signup successful. Please login.");
    } catch (error) {
      alert("Signup failed");
    }
  };

  return (
    <div className="wrapper signIn">
    <div className="form">
    <form onSubmit={handleSignup}>
    <div className="heading">SIGN-UP</div>
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
