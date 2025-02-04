'use client';

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Label } from '@/components/ui/label';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import { useUpload } from '@/hooks/useUpload';
import { useCreateTicketImageMutation } from '@/queries/create-ticket-image';
import { usePublicEventItemQuery } from '@/queries/public-event-item';
import type { TTicket } from '@/types';

import { CalendarIcon, MapPinIcon, StarIcon, } from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';
import { useState } from 'react';
import { useAccount } from 'wagmi';

// This would typically come from a database or API
const eventData = {
  id: 'eth2024',
  name: 'ETH Global 2024',
  date: 'November 15-17, 2024',
  location: 'San Francisco, CA',
  description:
    "Join the world's largest Ethereum hackathon and conference. Connect with leading developers, entrepreneurs, and enthusiasts in the Ethereum ecosystem. Experience three days of intense coding, workshops, and networking opportunities.",
  imageUrl: 'https://place-hold.it/800x400',
  tickets: [
    { id: 'general', name: 'General Admission', price: '0.1 ETH' },
    { id: 'vip', name: 'VIP Access', price: '0.3 ETH' },
  ],
  organizer: {
    name: 'ETH Global',
    description:
      "ETH Global is the world's largest Ethereum hackathon and conference organizer, bringing together developers, entrepreneurs, and enthusiasts from around the globe.",
    avatarUrl: 'https://place-hold.it/100x100',
    reviewCount: 128,
  },
};

export default function EventPage({ params }: { params: { slug: string } }) {
  const { address } = useAccount();
  const { event } = usePublicEventItemQuery(params.slug);
  const { mutateAsync, isPending: isApiLoading } =
    useCreateTicketImageMutation();
  const { upload, isLoading } = useUpload();
  const [selectedTicketId, setSelectedTicketId] = useState<string>('');

  const selectedTicket = event?.tickets.find(
    (t: TTicket) => t.contractAddress === selectedTicketId,
  );
  // let mintPrice = event?.tickets.find(
  //   (t: TTicket) => t.contractAddress === selectedTicketContract,
  // )?.mintPrice;
  // mintPrice = mintPrice ? parseEther(mintPrice) : 0n;
  // let baseTokenUri = event?.tickets.find(
  //   (t: TTicket) => t.contractAddress === selectedTicketContract,
  // )?.baseTokenUri;
  //
  // const contracts = [
  //   {
  //     address: event?.contractAddress as Address,
  //     abi: eventABI,
  //     functionName: 'mintNFT',
  //     args: [selectedTicketContract, baseTokenUri],
  //     value: mintPrice,
  //   },
  // ];

  // const handleError = useCallback((err: TransactionError) => {
  //   console.error('Transaction error:', err);
  // }, []);
  //
  // const handleSuccess = useCallback(async (response: TransactionResponse) => {
  //   console.log(response);
  // }, []);

  return (
    <div className="min-h-screen bg-background text-foreground">
      <header className="relative h-[400px] w-full">
        <Image
          src={eventData.imageUrl}
          alt={event?.name}
          fill={true}
          style={{ objectFit: 'cover' }}
          priority={true}
        />
        <div className="absolute inset-0 bg-black/60 flex items-end">
          <div className="container mx-auto px-4 py-6">
            <h1 className="text-4xl font-bold text-white mb-2">
              {event?.name}
            </h1>
            <div className="flex items-center text-white/80 space-x-4">
              <span className="flex items-center">
                <CalendarIcon className="mr-2 h-5 w-5" />
                {event?.date}
              </span>
              <span className="flex items-center">
                <MapPinIcon className="mr-2 h-5 w-5" />
                {event?.location}
              </span>
            </div>
          </div>
        </div>
      </header>

      <main className="container mx-auto px-4 py-12">
        <div className="grid md:grid-cols-3 gap-8">
          <div className="md:col-span-2 space-y-8">
            <Card>
              <CardHeader>
                <CardTitle>Organizer</CardTitle>
              </CardHeader>
              <CardContent className="flex items-center space-x-4">
                <Avatar className="h-16 w-16">
                  <AvatarImage
                    src={event?.organizer?.avatarUrl}
                    alt={event?.organizer?.name}
                  />
                  <AvatarFallback>
                    {eventData.organizer.name.slice(0, 2).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div className="space-y-1">
                  <h3 className="text-lg font-semibold">
                    {event?.organizer?.name || event?.organizer?.walletAddress}
                  </h3>
                  <p className="text-sm text-muted-foreground">
                    {event?.organizer?.description}
                  </p>
                  <Link
                    href="#reviews"
                    className="text-sm text-primary flex items-center hover:underline"
                  >
                    <StarIcon className="mr-1 h-4 w-4" />
                    {eventData.organizer.reviewCount} attestations
                  </Link>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>About This Event</CardTitle>
              </CardHeader>
              <CardContent>
                <p>{event?.description}</p>
              </CardContent>
            </Card>
          </div>

          <Card>
            <CardHeader>
              <CardTitle>Ticket Information</CardTitle>
              <CardDescription>
                Secure your spot at {eventData.name}
              </CardDescription>
            </CardHeader>
            <CardContent>
              <RadioGroup
                onValueChange={(value: string) => {
                  value && setSelectedTicketId(value);
                }}
                defaultValue="general"
                className="space-y-4"
              >
                {event?.tickets.map((ticket: TTicket) => (
                  <div key={ticket.id} className="flex items-center space-x-2">
                    <RadioGroupItem
                      value={ticket?.contractAddress || ''}
                      id={ticket.id?.toString()}
                    />
                    <Label
                      htmlFor={ticket.id?.toString()}
                      className="flex justify-between w-full"
                    >
                      <span>{ticket?.name}</span>
                      <span>{ticket?.mintPrice} ETH</span>
                    </Label>
                  </div>
                ))}
              </RadioGroup>
            </CardContent>
            <CardFooter>
              {address ? (
                <Link
                  href={`/checkout?ticket=${btoa(`${selectedTicketId}|${selectedTicket?.imageUrl || ''}`)}`}
                  passHref={true}
                  legacyBehavior={true}
                >
                  <Button
                    as="a"
                    className="w-full"
                    disabled={!selectedTicketId}
                  >
                    {selectedTicketId
                      ? 'Proceed to Checkout'
                      : 'Select a Ticket'}
                  </Button>
                </Link>
              ) : (
                <p className="text-center text-sm text-muted-foreground">
                  Please connect your wallet to purchase tickets.
                </p>
              )}
            </CardFooter>
          </Card>
        </div>
      </main>
    </div>
  );
}
