'use client';

import { useEffect, useState } from 'react';
import TicketImage from './ticket-image';
import OrderDetails from './order-details';
import TransactionHash from './transaction-hash';
import MintingStatus from './minting-status';
import NoOrderData from './no-order-data';
import ErrorOrder from './order-error';
import { useGetOrderQuery } from '@/queries/get-order';
import { Loader2 } from 'lucide-react';

export default function OrderSuccessPage({
  searchParams,
}: {
  searchParams: { orderId: string };
}) {
  const { orderId } = searchParams;
  const [isInitialLoad, setIsInitialLoad] = useState(true);
  const {
    order: orderData,
    error,
    refetch,
    notFound,
  } = useGetOrderQuery(orderId);

  useEffect(() => {
    if (orderData || error) {
      setIsInitialLoad(false);
    }
  }, [orderData, error]);

  if (isInitialLoad) {
    return (
      <div className="min-h-screen bg-white flex items-center justify-center">
        <div className="w-full max-w-3xl px-4 py-8 sm:px-6 lg:px-8">
          <div className="bg-white shadow-2xl rounded-3xl overflow-hidden">
            <div className="px-6 py-8 sm:p-10">
              <div className="flex flex-col items-center mb-8">
                <Loader2 className="h-16 w-16 text-black-500 animate-spin mb-4" />
                <div className="h-8 bg-gray-200 rounded w-3/4 mb-4"></div>
                <div className="h-4 bg-gray-200 rounded w-1/2"></div>
              </div>
              <div className="flex flex-col md:flex-row gap-8">
                <div className="md:w-1/2 bg-gray-200 rounded-lg h-64"></div>
                <div className="md:w-1/2 space-y-6">
                  <div className="h-4 bg-gray-200 rounded w-3/4"></div>
                  <div className="h-4 bg-gray-200 rounded w-1/2"></div>
                  <div className="h-4 bg-gray-200 rounded w-2/3"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  if (error && notFound) {
    return <NoOrderData />;
  }
  if (error) {
    return <ErrorOrder refetch={refetch} />;
  }

  return (
    <div className="min-h-screen bg-white flex items-center justify-center">
      <div className="w-full max-w-3xl px-4 py-8 sm:px-6 lg:px-8">
        <div className="bg-white shadow-2xl rounded-3xl overflow-hidden">
          <div className="px-6 py-8 sm:p-10">
            <div className="flex flex-col items-center mb-8">
              {orderData.paymentStatus === 'completed' ? (
                <div className="mb-6">
                  <svg
                    className="h-16 w-16 text-green-500"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                </div>
              ) : (
                <div className="mb-6">
                  <svg
                    className="h-16 w-16 text-yellow-500 mx-auto animate-spin"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                </div>
              )}
              <h1 className="text-3xl font-extrabold text-gray-900 text-center mb-2">
                {orderData.paymentStatus === 'completed'
                  ? 'Order Successful!'
                  : 'Order Pending'}
              </h1>
              <p className="text-lg text-gray-600 text-center">
                {orderData.paymentStatus === 'completed'
                  ? 'Thank you for your purchase. Your ticket has been processed.'
                  : 'Thank you for your purchase. Your ticket is being processed.'}
              </p>
            </div>
            <div className="flex flex-col md:flex-row gap-8">
              <div className="md:w-1/2">
                <TicketImage imageUrl={orderData.ticketImageUrl} />
              </div>
              <div className="md:w-1/2 space-y-6">
                <OrderDetails orderData={orderData} />
                <TransactionHash hash={orderData.paymentTransactionHash} />
                <MintingStatus
                  mintStatus={orderData.mintingStatus}
                  transferTxHash={orderData.transferTransactionHash}
                  mintTxHash={orderData.mintTransactionHash}
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
