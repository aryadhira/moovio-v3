"use client"
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";

const Navbar = () => {
  const pathname = usePathname();

  const menuItems = [
    { name: "Home", path: "/" },
    { name: "Discovery", path: "/discovery" },
    { name: "Top Rated", path: "/toprated" },
  ]

  return (
    <nav className="bg-neutral-900 shadow-lg">
      <div className="container mx-auto p-4 flex items-center justify-between">
        <div className="flex items-center gap-5">
          <Image src="/clapperboard.png" alt="brand-logo" height={30} width={30}/>
          <h1 className="text-xl font-bold">Moovio</h1>
        </div>
        <ul className="flex items-center space-x-5">
          {menuItems.map((item) => (
            <li
              key={item.path}
              className={`${pathname === item.path ? "text-white" : "text-gray-500"} hover:text-white`}
            >
              <Link href={item.path} className="font-bold">{item.name}</Link>
            </li>
          ))}
        </ul>
      </div>
    </nav>
  )
}

export default Navbar;