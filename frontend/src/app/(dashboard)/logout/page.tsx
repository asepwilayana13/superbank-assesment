"use client"
import { useEffect } from "react";
import { destroyCookie } from "nookies";
import { useRouter } from "next/navigation";

export default function LogoutPage() {
  const router = useRouter();

  useEffect(() => {
    destroyCookie(null, "token");
   
    router.push("/login");
  }, []);

  return (
    <div className="flex justify-center items-center h-screen text-gray-600 text-sm">
      Logging out...
    </div>
  );
}