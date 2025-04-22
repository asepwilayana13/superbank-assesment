export const getCustomers = async () => {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/castomers`);

  if (!res.ok) {
    throw new Error("Failed to fetch students");
  }

  return res.json();
};