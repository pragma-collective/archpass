import { Loader2 } from 'lucide-react';
import Link from 'next/link';

export default function TransactionHash({ hash }: { hash: string | null }) {
  return (
    <div className="bg-gray-50 rounded-xl p-6">
      <h3 className="text-xl font-semibold mb-4 text-gray-800">
        Transaction Hash
      </h3>
      {hash ? (
        <Link
          href={`https://basescan.org/tx/${hash}`}
          target="_blank"
          rel="noopener noreferrer"
        >
          View block explorer
        </Link>
      ) : (
        <div className="flex items-center space-x-2">
          <Loader2 className="h-5 w-5 animate-spin text-primary" />
          <p className="text-sm text-gray-600">
            Waiting for transaction hash...
          </p>
        </div>
      )}
    </div>
  );
}
