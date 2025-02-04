'use client';

import { Loader2 } from 'lucide-react';
import { Card, CardContent, CardHeader } from '@/components/ui/card';
import Image from 'next/image';

export default function LoadingComponent() {
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
          <div className="w-full max-w-[370px] mx-auto relative rounded-lg overflow-hidden shadow-lg bg-gray-200 animate-pulse h-[600px]" />
          <div className="w-full text-center">
            <Loader2 className="h-8 w-8 animate-spin text-primary mx-auto mb-4" />
            <p className="text-lg font-medium">Loading...</p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
