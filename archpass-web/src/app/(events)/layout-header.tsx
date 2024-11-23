'use client';

import LoginButton from '@/components/LoginButton';
import SignupButton from '@/components/SignupButton';
import Image from 'next/image';
import Link from 'next/link';
import { useAccount } from 'wagmi';

export function EventLayoutHeader() {
  const { address } = useAccount();
  return (
    <header className="container flex justify-between mx-auto py-6 px-4">
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
      <nav className="flex items-center">
        <div className="flex items-center space-x-4">
          {!address && <LoginButton requireSIWE={false} />}

          <SignupButton requireSIWE={false} />
        </div>
      </nav>
    </header>
  );
}
