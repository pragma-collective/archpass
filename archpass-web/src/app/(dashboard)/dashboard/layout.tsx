'use client';
import { Card, CardContent } from '@/components/ui/card';
import { usePathname } from 'next/navigation';

import '@/app/global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';
import LoginButton from '@/components/LoginButton';
import dynamic from 'next/dynamic';
import { DashboardSidebar } from './[eventSlug]/sidebar';
import { Inter } from 'next/font/google';
import Image from 'next/image';
import Link from 'next/link';

const inter = Inter({ subsets: ['latin'] });

const OnchainProviders = dynamic(
  () => import('@/components/OnchainProviders'),
  {
    ssr: false,
  },
);

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();
  const showSidebar = pathname !== '/dashboard';

  return (
    <html lang="en" className="h-full">
      <body
        className={`h-full bg-background text-foreground ${inter.className}`}
      >
        <OnchainProviders>
          <div className="flex min-h-screen">
            {/* Sidebar */}
            {showSidebar && <DashboardSidebar />}

            {/* Main content */}
            <div className="flex-1">
              <header className="flex justify-between items-center bg-card text-card-foreground p-4 border-b">
                {showSidebar ? (
                  <span />
                ) : (
                  <Link href="/" className="flex items-center space-x-2">
                    <Image
                      src="/archpass-logo-solid.png"
                      alt="ArchPass Logo"
                      width={32}
                      height={32}
                    />
                    <span className="text-xl">
                      <span className="font-semibold">Arch</span>Pass
                    </span>
                  </Link>
                )}
                <LoginButton />
              </header>

              <main className="p-4 md:p-8">
                <Card>
                  <CardContent className="p-6">{children}</CardContent>
                </Card>
              </main>
            </div>
          </div>
        </OnchainProviders>
      </body>
    </html>
  );
}
