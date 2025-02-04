'use client';

import { AlertCircle } from 'lucide-react';
import { Card, CardContent, CardHeader } from '@/components/ui/card';
import Image from 'next/image';

export default function ErrorComponent() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-4">
      <Card className="w-full max-w-xl">
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
          <AlertCircle className="h-16 w-16 text-red-500" />
          <p className="text-lg font-medium text-center">
            An error occurred while processing your order.
          </p>
          <p className="text-lg text-gray-600 text-center mb-8">
            Please try again later or refresh the page.
          </p>{' '}
        </CardContent>
      </Card>
    </div>
  );
}
