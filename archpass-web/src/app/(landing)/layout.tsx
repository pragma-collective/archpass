import type { Metadata } from 'next';
import { NEXT_PUBLIC_URL } from '../../config';

import '../global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';
import LoginButton from '@/components/LoginButton';
import dynamic from 'next/dynamic';
import Link from 'next/link';
import { Inter } from 'next/font/google';
import Image from 'next/image';

const inter = Inter({ subsets: ['latin'] });

const OnchainProviders = dynamic(
  () => import('@/components/OnchainProviders'),
  {
    ssr: false,
  },
);

export const viewport = {
  width: 'device-width',
  initialScale: 1.0,
};

export const metadata: Metadata = {
  title: 'ArchPass',
  description: 'On-chain events',
  openGraph: {
    title: 'ArchPass',
    description: 'Built with OnchainKit',
    images: [`${NEXT_PUBLIC_URL}/vibes/vibes-19.png`],
  },
};

export default function RootLayout({
  children,
}: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body
        className={`bg-background items-center justify-center ${inter.className}`}
      >
        <OnchainProviders hideSmartWallet>
          <div className="min-h-screen bg-background text-foreground">
            <header className="container mx-auto py-6 px-4">
              <div className="container mx-auto px-4 py-4 flex justify-between items-center">
                <Link href="/" className="flex items-center space-x-2">
                  <Image
                    src="/archpass-logo-white.svg"
                    alt="ArchPass Logo"
                    width={32}
                    height={32}
                  />
                  <span className="text-xl">
                    <span className="font-semibold">Arch</span>Pass
                  </span>
                </Link>
                <nav className="">
                  <div className="flex items-center space-x-4">
                    <div className="hidden sm:flex space-x-4">
                      <Link
                        href="/about"
                        className="hover:text-primary transition-colors"
                      >
                        About
                      </Link>
                      <Link
                        href="/contact"
                        className="hover:text-primary transition-colors"
                      >
                        Contact
                      </Link>
                    </div>
                    <LoginButton />
                    {/*<Button variant="outline" asChild>
                  <Link href="/login">Login</Link>
                </Button>
                  <Button asChild>
                    <Link href="/signup">Sign up</Link>
                  </Button>*/}
                  </div>
                </nav>
              </div>
            </header>
            {children}
          </div>
        </OnchainProviders>
      </body>
    </html>
  );
}
