'use client';
import { AuthProvider } from '@/providers/auth-provider';
import { OnchainKitProvider } from '@coinbase/onchainkit';
import { RainbowKitProvider } from '@rainbow-me/rainbowkit';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import type { ReactNode } from 'react';
import { base } from 'viem/chains';
import { WagmiProvider } from 'wagmi';
import { NEXT_PUBLIC_CDP_API_KEY, RPC_URL } from '../config';
import { useWagmiConfig } from '../wagmi';

type Props = { children: ReactNode; hideSmartWallet?: boolean };

const queryClient = new QueryClient();

function OnchainProviders({ children, hideSmartWallet = false }: Props) {
  const wagmiConfig = useWagmiConfig(hideSmartWallet);

  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <OnchainKitProvider
          apiKey={NEXT_PUBLIC_CDP_API_KEY}
          rpcUrl={RPC_URL}
          chain={base}
        >
          <RainbowKitProvider modalSize="compact">
            <AuthProvider>{children}</AuthProvider>
          </RainbowKitProvider>
        </OnchainKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}

export default OnchainProviders;
