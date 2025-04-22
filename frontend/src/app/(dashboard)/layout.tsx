import Menu from "@/components/Menu";
import Navbar from "@/components/Navbar";
import { CreditCard } from "lucide-react";
import Image from "next/image";
import Link from "next/link";

export default function DashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="min-h-screen flex">
      {/* LEFT */}
      <div className="w-[14%] md:w-[8%] lg:w-[16%] xl:w-[14%] p-4  border-b-gray-300  bg-white">
        <Link
          href="/"
          className="flex items-center justify-center lg:justify-start gap-2"
        >
          <div className="w-10 h-10 bg-blue-50 rounded-full flex items-center justify-center text-blue-600"><CreditCard /></div>
          <span className="hidden lg:block font-bold">SuperBank</span>
        </Link>
        <Menu/>
      </div>
      {/* RIGHT */}
      <div className="w-[86%] md:w-[92%] lg:w-[84%] xl:w-[86%] bg-[#F7F8FA] overflow-hidden flex flex-col">
        <Navbar />
        {children}
      </div>
    </div>
  );
}