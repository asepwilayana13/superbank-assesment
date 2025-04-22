"use client";
import { MoreHorizontal } from "lucide-react";
import Image from "next/image";
import {
  BarChart,
  Bar,
  Rectangle,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

const data = [
  {
    name: "Mon",
    deposits: 60,
    pockets: 40,
  },
  {
    name: "Tue",
    deposits: 70,
    pockets: 60,
  },
  {
    name: "Wed",
    deposits: 90,
    pockets: 75,
  },
  {
    name: "Thu",
    deposits: 90,
    pockets: 75,
  },
  {
    name: "Fri",
    deposits: 65,
    pockets: 55,
  },
];

const Chart = () => {
  return (
    <div className="bg-blue-300 shadow rounded-lg p-4 h-full">
      <div className="flex justify-between items-center">
        <h1 className="text-white text-lg font-semibold">Deposite & Pocket</h1>
        {/* <Image src="/moreDark.png" alt="" width={20} height={20} /> */}
        <MoreHorizontal size={20} />
      </div>
      <ResponsiveContainer width="100%" height="90%">
        <BarChart width={500} height={300} data={data} barSize={20}>
          <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#ddd" />
          <XAxis
            dataKey="name"
            axisLine={false}
            tick={{ fill: "#d1d5db" }}
            tickLine={false}
          />
          <YAxis axisLine={false} tick={{ fill: "#d1d5db" }} tickLine={false} />
          <Tooltip
            contentStyle={{ borderRadius: "10px", borderColor: "lightgray" }}
          />
          <Legend
            align="left"
            verticalAlign="top"
            wrapperStyle={{ paddingTop: "20px", paddingBottom: "40px" }}
          />
          <Bar
            dataKey="deposits"
            fill="#FEF9C3"
            legendType="circle"
            radius={[10, 10, 0, 0]}
          />
          <Bar
            dataKey="pockets"
            fill="#DBEAFE"
            legendType="circle"
            radius={[10, 10, 0, 0]}
          />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default Chart;