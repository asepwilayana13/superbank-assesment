type QueryParams = Record<string, string | number | boolean>;

export const getCustomers = async (params: string, page:number, limit: number) => {
const query_params: QueryParams =  {
  limit: limit ? limit : 10 ,
  page: page ? page : 1,
  search: params
}
const query = Object.keys(query_params).map(k => encodeURIComponent(k) + '=' + encodeURIComponent(query_params[k])).join('&')
const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/customer?${query}`, {})

  if (!res.ok) {
    throw new Error("Failed to fetch students");
  }

  return res.json();
};