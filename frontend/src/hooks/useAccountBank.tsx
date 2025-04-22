import { useEffect, useState } from 'react';
import { getAccountBankDetail } from '@/services/accountBankService';
import {TAccountBankDetailRespone } from '@/interface/bankAccount.interface';

// type TAccountBankDetail = {
//   id: string,
//   customer_id: string,
//   account_number: string,
//   account_type: string,
//   balance: number,
// }

export const useAccountBankDetail = (id: string | number) => {
  const [accountBankDetail, setaccountBankDetail] = useState<TAccountBankDetailRespone>();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    const fetchAccountBankDetail = async () => {
      try {
        const data = await getAccountBankDetail(id);
        
        setaccountBankDetail(data.data);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchAccountBankDetail();
  }, [id]);

  return { accountBankDetail, loading, error };
};