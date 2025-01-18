import {Source_Sans_3} from 'next/font/google';
import "./globals.css";
import Sidebar from '@/components/sidebar';

const sourceSans = Source_Sans_3({
  family: 'Source+Sans+3',
  subsets: ["latin"],
  weight:["200","300","400","700","900"]
})
export const metadata = {
  title: "Moovio | Stream Free Movies Unlimited",
  description: "Stream Free Movies Unlimited",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={sourceSans.className}>
        <div className="flex flex-row gap-5">
          <div className='w-1/6'>
            <Sidebar/>
          </div>
          <div className='w-5/6'>
            {children}
          </div>
        </div>
      </body>
    </html>
  );
}
