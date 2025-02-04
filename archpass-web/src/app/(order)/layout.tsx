import type { Metadata } from 'next';
import { NEXT_PUBLIC_URL } from '../../config';
import dynamic from 'next/dynamic';

import '../global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';

const PublicOnchainProviders = dynamic(
  () => import('@/components/PublicOnchainProviders'),
  {
    ssr: false,
  },
);

export const viewport = {
  width: 'device-width',
  initialScale: 1.0,
};

export const metadata: Metadata = {
  title: 'Archpass',
  description: 'Order successful',
  openGraph: {
    title: 'Onchain App Template',
    description: 'Built with OnchainKit',
    images: [`${NEXT_PUBLIC_URL}/vibes/vibes-19.png`],
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <PublicOnchainProviders>{children}</PublicOnchainProviders>
      </body>
    </html>
  );
}
