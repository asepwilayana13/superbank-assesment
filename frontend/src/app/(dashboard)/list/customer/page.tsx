'use client'
import FormModal from "@/components/FormModal";
import Pagination from "@/components/Pagination";
// import FormModal from "@/components/FormModal";
import Table from "@/components/Table";
import { useCustomers } from "@/hooks/useCustomer";
import { Filter, Search, SortAsc, View } from "lucide-react";
import moment from "moment";
// import TableSearch from "@/components/TableSearch";
import Image from "next/image";
import Link from "next/link";
import { useState } from "react";

type Customer = {
  id: number;
  full_name: string;
  phone_number: string;
  date_of_birth?: number;
  address: string;
  bankAccount: BankAccount;
};

type BankAccount = {
  id: number
}

const columns = [
  {
    header: "Info",
    accessor: "info",
  },
  {
    header: "Fullname",
    accessor: "fullName",
    className: "hidden md:table-cell",
  },
  {
    header: "Phone number",
    accessor: "phoneNumber",
    className: "hidden md:table-cell",
  },
  {
    header: "Date Of Birth",
    accessor: "dateOfBirth",
    className: "hidden lg:table-cell",
  },
  {
    header: "Address",
    accessor: "address",
    className: "hidden lg:table-cell",
  },
  {
    header: "Actions",
    accessor: "action",
  },
];


const CustomerPage = () => {
  const [search, setSearch] = useState('')
  const [page, setPage] = useState<number>(1)
  const [limit, setlimit] = useState(10)
  const { customers, totalPages, totalCount, loading, error } = useCustomers(search, page, limit);

  const renderRow = (item: Customer) => (
    
    <tr key={item.id} className="border-b border-gray-200 even:bg-slate-50 text-sm hover:bg-gray-100">
      <td className="flex items-center gap-4 p-4">
        <Image
          src="https://images.pexels.com/photos/2888150/pexels-photo-2888150.jpeg?auto=compress&cs=tinysrgb&w=1200"
          alt=""
          width={40}
          height={40}
          className="md:hidden xl:block w-10 h-10 rounded-full object-cover"
        />
      </td>
      <td className="hidden md:table-cell">{item.full_name}</td>
      <td className="hidden md:table-cell">{item.phone_number}</td>
      <td className="hidden md:table-cell">{moment(item.date_of_birth).format('DD MMMM YYYY')}</td>
      <td className="hidden md:table-cell">{item.address}</td>
      <td>
        <div className="flex items-center gap-2">
        <Link href={`/list/bank-account/${item?.bankAccount?.id}`}>
          <button className="w-10 h-10 flex items-center justify-center rounded-full bg-blue-50 cursor-pointer hover:bg-gray-700">
            <span><View size={20} color="#3B82F6"/></span>
          </button>
        </Link>
        {/* {role === "admin" && ( */}
          {/* <button className="w-10 h-10 flex items-center justify-center rounded-full bg-red-50">
            <span><Trash2 size={20} color="#EF4444"/></span>
          </button> */}
          <FormModal table="customer" type="delete" id={item.id}/>
        {/* )} */}
        </div>
      </td>
    </tr>

  );
  return (
    <div className="bg-white p-4 rounded-md flex-1 m-4 mt-0">
      {/* TOP */}
      <div className="flex items-center justify-between">
          <h1 className="hidden md:block text-lg font-semibold">All Customer</h1>
        <div className="flex flex-col md:flex-row items-center gap-4 w-full md:w-auto">
          {/* <TableSearch /> */}
          <div className='hidden md:flex items-center gap-2 text-xs rounded-full ring-[1.5px] ring-gray-300 px-2'>
            <span><Search size={20}/></span>
            <input type="text" placeholder="Search..." onChange={(e) => setSearch(e.target.value)} className="w-[200px] p-2 bg-transparent outline-none"/>
          </div>
          <div className="flex items-center gap-4 self-end">
            <button className="w-8 h-8 flex items-center justify-center rounded-full bg-green-50">
              <span className="text-md"><Filter size={20} color="#22C55E"/></span>
            </button>
            <button className="w-8 h-8 flex items-center justify-center rounded-full bg-green-50">
              <span className="text"><SortAsc size={20} color="#22C55E"/></span>
            </button>
            {/* {role === "admin" && ( */}
            {/* <button className="w-8 h-8 flex items-center justify-center rounded-full bg-green-50">
              <span className="text"><Plus size={20} color="#22C55E"/></span>
            </button> */}

              <FormModal table="customer" type="create" />
            {/* )} */}
          </div>
        </div>
      </div>
      {/* LIST */}
      <Table columns={columns} renderRow={renderRow} data={customers} />
      {/* PAGINATION */}
      <div className="flex justify-end">
        <Pagination currentPage={page}
          totalPages={totalPages}
          onPageChange={(newPage) => {
            setPage(newPage)
          }}
        />
      </div>
    </div>
  );
};

export default CustomerPage;