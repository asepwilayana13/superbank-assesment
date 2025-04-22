import Announcements from "@/components/Announcements";
import CustomerChart from "@/components/CustomerChart";
import UserCard from "@/components/CustomerCard";

const AdminPage = () => {
  return (
    <div className="bg-white p-4 rounded-md flex-1 m-4 mt-0">
      <div className="p-4 flex gap-4 flex-col md:flex-row">
        {/* LEFT */}
        <div className="w-full lg:w-2/3 flex flex-col gap-8">
          {/* USER CARDS */}
          <div className="flex gap-4 justify-between flex-wrap">
            <UserCard type="Customer" color="bg-green-100" />
            <UserCard type="Deposit" color="bg-yellow-100"/>
            <UserCard type="Pocket" color="bg-blue-100"/>
          </div>
          {/* MIDDLE CHARTS */}
          <div className="flex gap-4 flex-col lg:flex-row">
            {/* COUNT CHART */}
            <div className="w-full h-[450px]">
              <CustomerChart />
            </div>
          </div>
          {/* BOTTOM CHART */}
          <div className="w-full h-[500px]">
            {/* <FinanceChart /> */}
          </div>
        </div>
        {/* RIGHT */}
        <div className="w-full lg:w-1/3 flex flex-col gap-8">
          {/* <EventCalendar /> */}
          <Announcements/>
        </div>
      </div>
    </div>
  );
};

export default AdminPage;