'use client';
import { useAuth } from '@/hooks/useAuth';
import { useRouter } from 'next/navigation';
import { useState } from 'react';

export default function RegisterCard() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const { register, error } = useAuth();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await register(username, email, password);
  };
  
  return (
    <div className="flex max-w-4xl mx-auto mt-20 shadow-lg rounded-lg overflow-hidden">
      {/* LEFT - CTA TO LOGIN */}
      <div className="w-1/2 bg-gradient-to-tr from-pink-500 to-red-500 text-white flex flex-col justify-center items-center p-10">
        <h2 className="text-3xl font-bold mb-4">Welcome Back!</h2>
        <p className="text-center mb-6">
          To keep connected with us <br />
          please login with your personal info
        </p>
        <button 
          onClick={() => router.push('/login')}
          className="border border-white px-6 py-2 rounded-full font-semibold hover:bg-white hover:text-pink-600 transition"
        >
          SIGN IN
        </button>
      </div>

      {/* RIGHT - SIGN UP FORM */}
      <div className="w-1/2 bg-white p-10">
        <h2 className="text-3xl font-bold mb-6 text-center">Create Account</h2>

        <div className="flex justify-center gap-4 mb-6">
          <button className="border rounded-full w-10 h-10 flex items-center justify-center">
            <i className="fab fa-facebook-f text-sm"></i>
          </button>
          <button className="border rounded-full w-10 h-10 flex items-center justify-center">
            <i className="fab fa-google text-sm"></i>
          </button>
          <button className="border rounded-full w-10 h-10 flex items-center justify-center">
            <i className="fab fa-linkedin-in text-sm"></i>
          </button>
        </div>

        <p className="text-center text-gray-500 text-sm mb-4">or use your email for registration</p>

        <form className="space-y-4"  onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Username"
            onChange={(e) => setUsername(e.target.value)}
            className="w-full border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-pink-500"
          />
          <input
            type="email"
            placeholder="Email"
            onChange={(e) => setEmail(e.target.value)}
            className="w-full border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-pink-500"
          />
          <input
            type="password"
            placeholder="Password"
            onChange={(e) => setPassword(e.target.value)}
            className="w-full border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-pink-500"
          />
          <button
            type="submit"
            className="w-full bg-pink-500 text-white font-bold py-2 rounded hover:bg-pink-600 transition"
          >
            SIGN UP
          </button>
        </form>
      </div>
    </div>
  );
}
