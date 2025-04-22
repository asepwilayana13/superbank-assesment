// /features/auth/hooks/useAuth.ts
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { loginService, registerService } from '../services/authService';
import { setCookie } from 'nookies';

export function useAuth() {
  const [error, setError] = useState('');
  const router = useRouter();

  const login = async (email: string, password: string) => {
    try {
      const data = await loginService(email, password);
      
      setCookie({},"token",data.data.session.token)

      router.push('/admin');
    } catch (err: any) {
      setError(err.message);
    }
  };

  const register = async (username: string, email: string, password: string) => {
    try {
      await registerService(username, email, password);
      router.push('/login');
    } catch (err: any) {
      setError(err.message);
    }
  };

  return { register, login, error };
}
