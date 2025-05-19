import axios from 'axios';
// import { API_BASE_URL } from '../config';


const axiosInstance = axios.create({
  method : "POST",
  baseURL: import.meta.env.VITE_API_URL,
  
  // headers: {
  //   'Content-Type': 'application/json',
  // },
  body: JSON.stringify({
    email: "email",
    password: "password"
  }),
  // withCredentials: true
});

export default axiosInstance;
