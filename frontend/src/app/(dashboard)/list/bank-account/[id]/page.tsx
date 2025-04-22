'use client'

import Card from "@/components/Card";
import CardInformation from "@/components/CardInformation";
import CardMethod from "@/components/CardMethod";
import { useAccountBankDetail } from "@/hooks/useAccountBank";
import Table from "@/components/Table";
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
      <div className="bg-white shadow mt-4 p-4 rounded-md flex-1 m-4">
        {/* TOP */}
        <div className="flex items-center justify-between">
          <h1 className="hidden md:block text-lg font-semibold">list deposite</h1>
        </div>
        {/* LIST */}
        {/* <Table columns={columns} renderRow={renderRow}/> */}
        <table className="w-full mt-4">
          <thead className="text-left text-black text-md">
            <td className="px-6 py-3 text-left text-md font-medium text-black uppercase tracking-wider">deposit amount</td>
            <td className="px-6 py-3 text-left text-md font-medium text-black uppercase tracking-wider">Intereset Rate</td>
            <td className="px-6 py-3 text-left text-md font-medium text-black uppercase tracking-wider">tenor</td>
            <td className="px-6 py-3 text-left text-md font-medium text-black uppercase tracking-wider">Balance plus interest</td>
          </thead>

        <tr className="border-b border-gray-200 even:bg-slate-50 hover:bg-gray-100 text-left text-gray-500 text-sm">
          <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-700">1000000</td>
          <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-700">5</td>
          <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-700">12</td>
          <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-700">105000</td>
         </tr>
        </table>
        {/* PAGINATION */}
        <div className="flex justify-end">
      </div>
    </div>
    </div>
  );
};

export default BankAccoutPage;