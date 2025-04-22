import { ChevronRight, HandCoins, Wallet } from "lucide-react"

const CardMethod = () => {
  return (
    <div className="bg-blue-400 p-6 mt-4 rounded-lg w-[40%] relative shadow flex flex-col lg:flex-col items-start space-y-4 lg:space-y-0">
      <h2 className="text-md font-semibold text-gray-800 mb-4">Deposit or pocket through these methods.</h2>
      <div className="space-y-3 w-full">
        {/* Deposit Item  */}
        <div className="flex items-center w-full bg-white p-3 rounded-lg relative shadow-sm hover:bg-gray-50 cursor-pointer">
          <div className="flex items-center gap-3">
            <div className="h-10 w-10 flex items-center justify-center bg-green-50 rounded-full">
              <HandCoins size={20} color="#22C55E"/>
            </div>
            <div>
              <p className="font-medium text-sm text-gray-800">Deposit</p>
              <p className="text-xs text-gray-500">Deposit</p>
            </div>
          </div>
          <div className="text-gray-400 absolute right-4">
            <ChevronRight />
          </div>
        </div>
        {/* Pockets Item */}
        <div className="flex items-center w-full bg-white p-3 rounded-lg relative shadow-sm hover:bg-gray-50 cursor-pointer">
          <div className="flex items-center gap-3">
            <div className="h-10 w-10 flex items-center justify-center bg-blue-50 rounded-full">
              <Wallet size={20} color="#3B82F6" />
            </div>
            <div>
              <p className="font-medium text-sm text-gray-800">Pockets</p>
              <p className="text-xs text-gray-500">create pocket</p>
            </div>
          </div>
          <div className="text-gray-400 absolute right-4">
            <ChevronRight />
          </div>
        </div>
      </div>
    </div>
  )
}

export default CardMethod