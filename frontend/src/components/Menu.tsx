'use client';

import { usePathname } from "next/navigation";
import { Cog, CreditCard, House, LayoutDashboard, LogOut, User, Users } from "lucide-react";
import Link from "next/link";

const menuItems = [
  {
    title: "MENU",
    items: [
      {
        icon: <LayoutDashboard size={20}/>,
        label: "Dashboard",
        href: "/admin",
        visible: ["admin", "customer"],
      },
      {
        icon: <Users size={20}/>,
        label: "Customer",
        href: "/list/customer",
        visible: ["admin", "customer"],
      },
    ],
  },
  {
    title: "OTHER",
    items: [
      {
        icon: <LogOut size={20}/>,
        label: "Logout",
        href: "/logout",
        visible: ["admin", "customer"],
      },
    ],
  },
];

const Menu = () => {
  const pathname = usePathname();
  return (
    <div className="mt-4 text-md">
      {menuItems.map((section) => (
        <div className="flex flex-col gap-2" key={section.title}>
          <span className="hidden lg:block text-black font-semibold my-4">
            {section.title}
          </span>
          {section.items.map((item) => {
            const isActive = pathname === item.href;

            return (
              <Link
                href={item.href}
                key={item.label}
                className={`flex items-center justify-center lg:justify-start gap-4 text-black py-2 md:px-2 rounded-md 
                ${isActive ? 'bg-blue-50 text-blue-500' : 'hover:bg-gray-100'}`}
              >
                <span className="text-md">{item.icon}</span>
                <span className="hidden lg:block">{item.label}</span>
              </Link>
            );
          })}
        </div>
      ))}
    </div>
  );
};

export default Menu;