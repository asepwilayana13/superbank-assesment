import { TAccountBankDetail } from "@/interface/bankAccount.interface"
import { formatCardNumber } from "@/utils/formatCardNumber"
import { Check } from "lucide-react"
import moment from "moment"


const CardInformation = (props:TAccountBankDetail) => {
  return(
    <div className="bg-white p-6 mt-4 rounded-lg w-[60%] relative shadow flex flex-col lg:flex-col items-start justify-between space-y-4 lg:space-y-0">
      <div className="flex gap-4 w-full">
        {/* Header */}
        <div className="flex justify-between relative items-center mb-4">
          <div>
            <h2 className="text-lg font-semibold text-blue-400 mb-2">Card Information</h2>
            <p className="text-sm text-gray-500">We use maximum encryption so that your credit card and personal information are safe and secure.</p>
          </div>
        </div>
      </div>
      {/* Content */}
      <div className="flex flex-row lg:flex-row gap-6">
        {/* Credit Card Visual */}
        <div className="bg-gradient-to-br from-blue-200 to-blue-200 rounded-xl p-5 h-full w-[70%] flex flex-col justify-between shadow-inner">
          <div className="h-8 w-16 bg-blue-400 rounded mb-8"></div>
          <div className="text-lg font-semibold text-blue-400 tracking-widest">{formatCardNumber(String(props.accountNumber))}</div>
          <div className="flex justify-between mt-4 text-sm gap-4">
            <div>
              <p className="text-black">Cardholder Name</p>
              <p className="font-semibold text-blue-400">{props.fullName}</p>
            </div>
            <div>
              <p className="text-black">Ex. Date</p>
              <p className="font-semibold text-blue-400">24/30</p>
            </div>
            <div>
              <p className="text-black">CVC Code</p>
              <p className="font-semibold text-blue-400">345</p>
            </div>
          </div>
        </div>

        {/* Card Form */}
        <div className="space-y-4">
          <div>
            <label className="block text-sm text-gray-600 mb-1">Full Name</label>
            <input disabled type="text" className="w-full px-3 py-2 text-gray-700 shadow rounded-md focus:outline-none focus:ring focus:border-blue-300" value={props.fullName} />
          </div>
          <div>
            <label className="block text-sm text-gray-600 mb-1">Address</label>
            <input disabled type="text" className="w-full px-3 py-2 text-gray-700 shadow rounded-md focus:outline-none focus:ring focus:border-blue-300" value={props.address} />
          </div>
          <div className="flex flex-row gap-4">
            <div>
              <label className="block text-sm text-gray-600 mb-1">Phone Number</label>
              <input disabled type="text" className="w-full px-3 py-2 text-gray-700 shadow rounded-md focus:outline-none focus:ring focus:border-blue-300" value={props.phoneNumber} />
            </div>
            <div>
              <label className="block text-sm text-gray-600 mb-1">Date of Birth</label>
              <input disabled type="text" className="w-full px-3 py-2 text-gray-700 shadow rounded-md focus:outline-none focus:ring focus:border-blue-300" value={moment(props.dateOfBirth).format('DD MMMM YYYY')} />
            </div>
          </div>
        </div>
      </div>
      <button className="text-sm border rounded-full px-4 py-1 text-blue-400 absolute right-0 mr-4 hover:text-black flex items-center gap-1">
        <Check />
        {/* <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg> */}
        Saved Cards
      </button>
    </div>
  )
}

export default CardInformation