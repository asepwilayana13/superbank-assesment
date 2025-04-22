'use client'

import Card from "@/components/Card";
import CardInformation from "@/components/CardInformation";
import CardMethod from "@/components/CardMethod";
import { useAccountBankDetail } from "@/hooks/useAccountBank";
import { useParams } from "next/navigation";

const BankAccoutPage = () => {
  const { id } = useParams();
  
  const { accountBankDetail, loading, error } = useAccountBankDetail(id as string);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  
  return (
    <div className="bg-white p-4 rounded-md flex-1 m-4 mt-0">
      <Card 
        totalBalance={accountBankDetail?.totalBalance?.total_balance}
        balance={accountBankDetail?.balance}
        totalPocketBalance={accountBankDetail?.totalBalance?.total_pocket_balance}
        totalDepositeBalance={accountBankDetail?.totalBalance?.total_deposite_balance}
        fullName={accountBankDetail?.customer?.full_name}
      />
      <div className="flex flex-row gap-4">
        <CardInformation 
          accountNumber={accountBankDetail?.account_number} 
          fullName={accountBankDetail?.customer?.full_name}
          phoneNumber={accountBankDetail?.customer?.phone_number}
          address={accountBankDetail?.customer?.address}
          dateOfBirth={accountBankDetail?.customer?.date_of_birth}
        />
        <CardMethod />
      </div>
    </div>
  );
};

export default BankAccoutPage;