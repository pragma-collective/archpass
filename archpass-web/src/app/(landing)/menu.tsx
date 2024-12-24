'use client';

import LoginButton from '@/components/LoginButton';
import Link from 'next/link';
import { useAccount } from 'wagmi';

export default function Menu() {
  const { address } = useAccount();
  return (
    <nav className="">
      <div className="flex items-center space-x-4">
        <div className="hidden sm:flex space-x-4">
          <Link href="/about" className="hover:text-primary transition-colors">
            About
          </Link>
          <Link
            href="/contact"
            className="hover:text-primary transition-colors"
          >
            Contact
          </Link>
          {address ? (
            <Link
              href="/dashboard"
              className="hover:text-primary transition-colors"
            >
              Dashboard
            </Link>
          ) : null}
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
  );
}
