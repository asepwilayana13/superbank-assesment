
export const getAccountBankDetail = async (id: string | number) => {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/bank_account/detail/${id}`);

  if (!res.ok) {
    throw new Error("Failed to fetch students");
  }

  return res.json();
};