"use client";

import { Plus, Trash2 } from "lucide-react";
import dynamic from "next/dynamic";
import Image from "next/image";
import { useState,JSX, useRef } from "react";
import useClickOutside from "@/hooks/useClickOutSide";

// USE LAZY LOADING

const CustomerForm = dynamic(() => import("./forms/CustomerForm"), {
  loading: () => <h1>Loading...</h1>,
});

const forms: {
  [key: string]: (type: "create" | "update", data?: any) => JSX.Element;
} = {
  customer: (type, data) => <CustomerForm type={type} data={data} />,
};

const FormModal = ({
  table,
  type,
  data,
  id,
}: {
  table:
    | "customer"
    | "bankAccount";
  type: "create" | "update" | "delete";
  data?: any;
  id?: number;
}) => {
  const modalRef = useRef<HTMLDivElement>(null)
  const size = type === "create" ? "w-8 h-8" : "w-10 h-10";
  const bgColor =
    type === "create"
      ? "bg-green-50"
      : type === "update"
      ? "bg-Sky"
      : "bg-red-50";
  const icon =
    type === "create"
      ? <Plus size={20} color="#22C55E"/>
      : type === "update"
      ? "bg-Sky"
      : <Trash2 size={20} color="#EF4444"/>;

  const [open, setOpen] = useState(false);

  const Form = () => {
    return type === "delete" && id ? (
      <form action="" className="p-4 flex flex-col gap-4">
        <span className="text-center font-medium">
          All data will be lost. Are you sure you want to delete this {table}?
        </span>
        <button className="bg-red-700 text-white py-2 px-4 rounded-md border-none w-max self-center">
          Delete
        </button>
      </form>
    ) : type === "create" || type === "update" ? (
      forms[table](type, data)
    ) : (
      "Form not found!"
    );
  };

  useClickOutside({ref: modalRef, callback: () => {
    console.log('here')
    setOpen(false)
  }})

  return (
    <>
      <button
        className={`${size} flex items-center justify-center rounded-full ${bgColor}`}
        onClick={() => setOpen(true)}
      >
        <span>{icon}</span>
        {/* <Image src={`/${type}.png`} alt="" width={16} height={16} /> */}
      </button>
      {open && (
        <div className="w-screen h-screen fixed left-0 top-0 bg-gray-400/25 backdrop-blur-xs z-50 flex items-center justify-center">
          <div ref={modalRef} className="bg-white p-4 rounded-md relative w-[90%] md:w-[70%] lg:w-[60%] xl:w-[50%] 2xl:w-[40%]">
            <Form />
            <div
              className="absolute top-4 right-4 cursor-pointer"
              onClick={() => setOpen(false)}
            >
              <Image src="/close.png" alt="" width={14} height={14} />
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default FormModal;