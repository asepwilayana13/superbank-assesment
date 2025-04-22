'use client';

import { useState } from 'react';
import { useAuth } from '@/hooks/useAuth';
import { useRouter } from 'next/navigation';

export default function LoginForm() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { login, error } = useAuth();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await login(email, password);
  };
  
  return (
    <div className="flex max-w-4xl mx-auto mt-20 shadow-lg rounded-lg overflow-hidden">
      {/* LEFT - SIGN IN */}
      <div className="w-1/2 bg-white p-10">
        <h2 className="text-3xl font-bold mb-6 text-center">Sign in</h2>

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

        <p className="text-center text-gray-500 text-sm mb-4">or use your account</p>

        <form className="space-y-4"  onSubmit={handleSubmit}>
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full bg-gray-300 rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500"
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full bg-gray-300 rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-orange-500"
          />
          <p className="text-sm text-right text-gray-400 hover:underline cursor-pointer">Forgot your password?</p>
          <button
            type="submit"
            className="w-full bg-pink-500 text-white font-bold py-2 rounded hover:bg-pink-600 transition"
          >
            SIGN IN
          </button>
          {error && <p>{error}</p>}
        </form>
      </div>

      {/* RIGHT - SIGN UP CTA */}
      <div className="w-1/2 bg-gradient-to-tr from-pink-500 to-red-500 text-white flex flex-col justify-center items-center p-10">
        <h2 className="text-3xl font-bold mb-4">Hello, Friend!</h2>
        <p className="text-center mb-6">
          Enter your personal details and start <br /> journey with us
        </p>
        <button 
          onClick={() => router.push('/register')}
          className="border border-white px-6 py-2 rounded-full font-semibold hover:bg-white hover:text-pink-600 transition"
        >
          SIGN UP
        </button>
      </div>
    </div>
  );
}
