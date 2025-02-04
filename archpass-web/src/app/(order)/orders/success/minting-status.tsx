import { Loader2 } from 'lucide-react';
import Link from 'next/link';

export default function MintingStatus({
  mintStatus,
  transferTxHash,
  mintTxHash,
}: { mintStatus: string; mintTxHash: string; transferTxHash: string }) {
  return (
    <div className="bg-gray-50 rounded-xl p-6">
      <h3 className="text-xl font-semibold mb-4 text-gray-800">
        Ticket Minting Status
      </h3>
      <div>
        <p className="text-sm text-gray-600 mb-4">
          {mintStatus === 'pending' &&
            'Your ticket is being minted. This process may take a few minutes.'}
          {mintStatus === 'completed' &&
            'Your ticket has been successfully minted and sent to your wallet!'}
          {mintStatus === 'failed' &&
            'There was an issue minting your ticket. Please contact support.'}
        </p>
        {mintStatus === 'pending' && (
          <div className="flex items-center space-x-2">
            <Loader2 className="h-5 w-5 animate-spin text-primary" />
            <p className="text-sm text-gray-600">Minting in progress...</p>
          </div>
        )}
        {mintStatus === 'completed' && (
          <div className="space-y-2">
            {transferTxHash && (
              <div className="text-sm text-gray-600">
                <Link
                  href={`https://basescan.org/tx/${transferTxHash}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  View transfer transaction
                </Link>
              </div>
            )}
            {mintTxHash && (
              <div className="text-sm text-gray-600">
                <Link
                  href={`https://basescan.org/tx/${mintTxHash}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline"
                >
                  View mint transaction
                </Link>
              </div>
            )}
          </div>
        )}
        {mintStatus === 'failed' && (
          <div className="text-center text-red-600 font-semibold">
            ✗ Minting Failed
          </div>
        )}
      </div>
    </div>
  );
}
