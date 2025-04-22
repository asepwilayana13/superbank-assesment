// /features/auth/hooks/useAuth.ts
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { loginService } from '@/services/authService';

export function useAuth() {
  const [error, setError] = useState('');
  const router = useRouter();

  const login = async (email: string, password: string) => {
    try {
      const data = await loginService(email, password);
      
      localStorage.setItem('token', data.data.session.token);
      router.push('/user');
    } catch (err: any) {
      setError(err.message);
    }
  };

  return { login, error };
}
