'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Image from 'next/image';
import { Loader2 } from 'lucide-react';
import { useAccount } from 'wagmi';
import { Card, CardContent, CardHeader } from '@/components/ui/card';
import { useCreateCheckoutMutation } from '@/queries/create-checkout';
import Loader from './loading';
import Error from './error';

export default function CheckoutPage({
  searchParams,
}: {
  searchParams: { ticket: string };
}) {
  const { address } = useAccount();
  const [status, setStatus] = useState<'processing' | 'redirecting'>(
    'processing',
  );
  const router = useRouter();
  const { mutateAsync, isPending, error, i } = useCreateCheckoutMutation();

  const ticketDetails = searchParams.ticket;
  const parsedTicketDetails = atob(ticketDetails);
  const [ticketContract, ticketImage] = parsedTicketDetails.split('|');

  useEffect(() => {
    const processOrder = async () => {
      if (!address) {
        return;
      }

      try {
        const response = await mutateAsync({
          contractAddress: ticketContract,
          walletAddress: address,
        });

        setStatus('redirecting');

        setTimeout(() => {
          router.push(response.paymentUrl);
        }, 2500);
      } catch (error) {
        console.error('Error processing order:', error);
      }
    };

    processOrder();
  }, [router, address]);

  if (isPending) {
    return <Loader />;
  }

  if (error) {
    return <Error />;
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-4">
      <Card className="w-full max-w-md">
        <CardHeader className="flex flex-col items-center space-y-2">
          <div className="flex items-center space-x-2">
            <Image
              src="/archpass-logo-solid.png"
              alt="ArchPass Logo"
              width={32}
              height={32}
            />
            <span className="text-xl">
              <span className="font-semibold">Arch</span>Pass
            </span>
          </div>
        </CardHeader>
        <CardContent className="flex flex-col items-center space-y-6">
          <div className="w-full max-w-[370px] mx-auto relative rounded-lg overflow-hidden shadow-lg">
            <Image
              src={ticketImage || '/placeholder.svg'}
              alt="Event Ticket"
              width={741}
              height={1202}
              layout="responsive"
              priority={true}
            />
          </div>
          <div className="w-full text-center">
            <Loader2 className="h-8 w-8 animate-spin text-primary mx-auto mb-4" />
            <p className="text-lg font-medium">
              {status === 'processing'
                ? 'Processing your order...'
                : 'Redirecting you to payment page...'}
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
