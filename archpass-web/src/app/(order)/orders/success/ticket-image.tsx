import Image from 'next/image';

export default function TicketImage({ imageUrl }: { imageUrl: string }) {
  return (
    <div className="w-full max-w-[370px] mx-auto relative rounded-lg overflow-hidden shadow-lg">
      <Image
        src={imageUrl || '/placeholder-ticket.svg'}
        alt="Event Ticket"
        width={741}
        height={1202}
        layout="responsive"
        priority={true}
      />
    </div>
  );
}
