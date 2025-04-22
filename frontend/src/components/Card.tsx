import { TAccountBankDetail } from "@/interface/bankAccount.interface"
import { formatCurrency } from "@/utils/formatCurrency"
import { HandCoins, User, Wallet } from "lucide-react"

const Card = (props:TAccountBankDetail) => {
  return (
    <div className="bg-white p-6 rounded-lg relative shadow flex flex-col lg:flex-col items-start justify-between space-y-4 lg:space-y-0">
      <div className="flex gap-4">
        <div className="flex items-center justify-center w-12 h-12 bg-blue-400 rounded-full">
          <User />
        </div>
        <div>
          <h2 className="text-xl font-semibold text-gray-800">{props.fullName}</h2>
          <p className="text-sm text-gray-500">22-03-2024 12:00 PM</p>
        </div>
      </div>
      <div className="flex w-full gap-4 mt-4 text-sm text-gray-700">
        <div className="bg-white px-4 py-2 w-full rounded border border-white shadow">
          <p className="text-xs text-gray-500 mb-1">SALDO UTAMA</p>
          <div className="flex items-center gap-2">
            <span className="text-blue-400 text-xl font-semibold">{formatCurrency(props.balance)}</span>
          </div>
        </div>
        <div className="bg-white px-4 py-2 w-full rounded border border-white shadow">
          <p className="text-xs text-gray-500 mb-1">TOTAL SALDO</p>
          <div className="flex items-center gap-2">
            <span className="text-blue-400 text-xl font-semibold">{formatCurrency(props.totalBalance)}</span>
          </div>
        </div>
        <div className="bg-white px-4 py-2 w-full relative rounded border border-white shadow">
          <div className="flex flex-col">
            <p className="text-xs text-gray-500 mb-1">DEPOSITE</p>
            <div className="flex items-center gap-2">
              <span className="text-blue-400 text-xl font-semibold">{formatCurrency(props.totalDepositeBalance)}</span>
            </div>
            <div className="h-10 w-10 flex items-center justify-center absolute right-2 bg-green-50 rounded-full">
              <HandCoins size={20} color="#22C55E"/>
            </div>
          </div>
        </div>
        <div className="bg-white px-4 py-2 w-full relative rounded border border-white shadow">
          <div className="flex flex-col">
            <p className="text-xs text-gray-500 mb-1">POCKETS</p>
            <div className="flex items-center gap-2">
              <span className="text-blue-400 text-xl font-semibold">{formatCurrency(props.totalPocketBalance)}</span>
            </div>
            <div className="h-10 w-10 flex items-center justify-center absolute right-4 bg-blue-50 rounded-full">
              <Wallet size={20} color="#3B82F6" />
            </div>
          </div>
        </div>
        {/* <div className="bg-white px-4 py-2 w-full rounded border border-white shadow">
          <p className="text-xs text-gray-500 mb-1">LAST LOGIN</p>
          <span className="font-semibold">22-03-2024 12:00 PM</span>
        </div> */}
      </div>
        <button className="text-sm border rounded-full px-4 py-1 absolute right-0 mr-4 text-gray-600 border-gray-300 hover:bg-gray-200">
          Hide info
        </button>
    </div>

  )
}

export default Card