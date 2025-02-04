import { FileQuestion } from 'lucide-react';

export default function NoOrder() {
  return (
    <div className="min-h-screen bg-white flex items-center justify-center">
      <div className="w-full max-w-md px-4 py-8 sm:px-6 lg:px-8">
        <div className="bg-white shadow-2xl rounded-3xl overflow-hidden">
          <div className="px-6 py-8 sm:p-10">
            <div className="flex flex-col items-center mb-8">
              <FileQuestion className="h-16 w-16 text-gray-400 mb-4" />
              <h2 className="text-2xl font-bold text-gray-900 mb-2">
                No Order Found
              </h2>
              <p className="text-gray-600 text-center mb-6">
                We couldn't find any order data for the provided ID.
              </p>
              <a
                href="/"
                className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition-colors"
              >
                Return to Home
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
