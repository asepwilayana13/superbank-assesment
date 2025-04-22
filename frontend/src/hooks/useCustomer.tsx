import { useEffect, useState } from 'react';
import { getCustomers } from '../services/customerService';

export const useCustomers = (search:string, page:number, limit:number) => {
  const [customers, setCustomers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [totalPages, setTotalPages] = useState(0);
  const [totalCount, setTotalCount] = useState(0);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCustomers = async () => {
      try {
        const data = await getCustomers(search, page, limit);
        setCustomers(data.data.customer);
        setTotalPages(data.data.total_pages);
        setTotalCount(data.data.total_count);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchCustomers();
  }, [search, page, limit]);

  return { customers, totalPages, totalCount, loading, error };
};
