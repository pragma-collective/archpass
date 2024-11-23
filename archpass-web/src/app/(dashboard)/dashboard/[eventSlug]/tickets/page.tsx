'use client';

import { Button } from '@/components/ui/button';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { convertIpfsToHttps } from '@/lib/utils';
import { useEventItemQuery } from '@/queries/event-item';
import { useTicketListQuery } from '@/queries/ticket-list';
import type { TTicket } from '@/types';
import { Eye } from 'lucide-react';
import { useParams } from 'next/navigation';
import { CreateTicketModal } from './create-ticket-modal';

export default function TicketsPage() {
  const params = useParams();
  const eventSlug = params.eventSlug as string;
  const { event } = useEventItemQuery(eventSlug);
  const { tickets, refetchTicketList } = useTicketListQuery(event?.id);

  const handleViewTicket = (imageUrl: string) => {
    if (!imageUrl) {
      console.error('No image URL provided');
      return;
    }

    const httpsUrl = convertIpfsToHttps(imageUrl);
    window.open(httpsUrl, '_blank');
  };

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Tickets for {eventSlug}</h1>
        {event ? (
          <CreateTicketModal
            eventContractAddress={event.contractAddress}
            event={event}
            refetchTicketList={refetchTicketList}
          />
        ) : null}
      </div>

      <div className="bg-background rounded-lg border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Price</TableHead>
              <TableHead>Quantity</TableHead>
              {/*<TableHead>Sold</TableHead>*/}
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {tickets?.map((ticket: TTicket) => (
              <TableRow key={ticket.id}>
                <TableCell>{ticket.name}</TableCell>
                <TableCell>{ticket.mintPrice} ETH</TableCell>
                <TableCell>{ticket.quantity}</TableCell>
                {/*<TableCell>{89}</TableCell>*/}
                <TableCell>
                  <div className="flex space-x-2">
                    <Button
                      variant="outline"
                      size="icon"
                      onClick={() => handleViewTicket(ticket.imageUrl)}
                    >
                      <Eye className="h-4 w-4" />
                      <span className="sr-only">View Ticket</span>
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      <p className="text-sm text-muted-foreground">
        Total tickets:{' '}
        {tickets?.reduce(
          (sum: number, ticket: TTicket) => sum + ticket.quantity,
          0,
        )}
      </p>
    </div>
  );
}
