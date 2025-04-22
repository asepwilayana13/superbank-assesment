export type TAccountBankDetailRespone = {
  id: string,
  customer_id: string,
  account_number: string,
  account_type: string,
  balance: number,
  customer?: Customer,
  totalBalance?: totalBalance,
}

export type TAccountBankDetail = {
  id?: string,
  customerId?: string,
  accountNumber?: string,
  accountType?: string,
  balance?: number,
  fullName?: string,
  phoneNumber?: string,
  address?: string,
  dateOfBirth?: number,
  totalBalance?: number,
  totalPocketBalance?: number,
  totalDepositeBalance?: number,
}

type Customer = {
  id: number;
  full_name: string;
  phone_number: string;
  date_of_birth?: number;
  address: string;
};

type totalBalance = {
  total_pocket_balance: number;
  total_deposite_balance: number;
  total_balance: number;
};
