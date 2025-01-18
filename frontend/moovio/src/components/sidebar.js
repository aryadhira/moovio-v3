"use client"
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";

const Sidebar = () => {
    const pathname = usePathname();

    const menuItems = [
        { name: "Home", path: "/", icon: "/home.png" },
        { name: "Discovery", path: "/discovery", icon: "/compass.png" },
        { name: "Top Rated", path: "/toprated", icon: "/favorite.png" },
    ]

    return (
        <div className="w-full bg-neutral-900 h-screen p-4 top-0 left-0">
            <div className="flex flex-row items-center gap-3 py-5 pl-5">
                <Image src="/clapperboard.png" alt="brand-logo" height={30} width={30}/>
                <h1 className="text-2xl font-bold">Moovio</h1>
            </div>
            <h3 className="text-lg font-bold py-5">MENU</h3>
            <ul className="list-none">
                {menuItems.map((item) => (
                    <li
                        key={item.path}
                        className={`flex flex-row gap-3 items-center py-2 pl-2 ${
                        pathname === item.path ? "bg-neutral-700" : "hover:bg-neutral-700"
                        }`}
                    >
                        <Image src={item.icon} alt={item.name} width={25} height={25} />
                        <Link href={item.path}>{item.name}</Link>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default Sidebar;