import { AlertCircle } from 'lucide-react';

export default function ErrorOrder() {
  return (
    <div className="min-h-screen bg-white flex items-center justify-center p-4">
      <div className="w-full max-w-xl">
        {' '}
        <div className="bg-white shadow-2xl rounded-3xl overflow-hidden">
          <div className="px-8 py-10">
            {' '}
            <div className="flex flex-col items-center">
              <AlertCircle className="h-20 w-20 text-red-500 mb-6" />
              <h2 className="text-3xl font-bold text-gray-900 mb-4">
                Oops! Something went wrong
              </h2>{' '}
              <p className="text-lg text-gray-600 text-center mb-8">
                Failed to fetch order data. Please try again later.
              </p>{' '}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
