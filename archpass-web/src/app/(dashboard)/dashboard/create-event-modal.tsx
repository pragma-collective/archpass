import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { AP_EVENT_FACTORY_CONTRACT_ADDRESS } from '@/config';
import { BASE_SEPOLIA_CHAIN_ID, eventFactoryABI } from '@/constants';
import { useCreateEventMutation } from '@/queries/create-event';
import {
  Transaction,
  TransactionButton,
  type TransactionError,
  type TransactionResponse,
  TransactionStatus,
  TransactionStatusAction,
  TransactionStatusLabel,
} from '@coinbase/onchainkit/transaction';
import { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';
import type { ContractFunctionParameters } from 'viem';

type FormData = {
  eventName: string;
  description: string;
  location: string;
  eventDate: string;
  headerImage: string;
};

type CreateEventModalProps = {
  refetchEventList: () => void;
};

export function CreateEventModal({ refetchEventList }: CreateEventModalProps) {
  const [open, setOpen] = useState(false);
  const {
    register,
    formState: { errors },
    getValues,
    reset,
  } = useForm<FormData>();
  const { mutateAsync } = useCreateEventMutation();
  const contracts = [
    {
      address: AP_EVENT_FACTORY_CONTRACT_ADDRESS,
      abi: eventFactoryABI,
      functionName: 'createEvent',
      args: ['0xtesthash'],
    },
  ] as unknown as ContractFunctionParameters[];

  const handleError = useCallback((err: TransactionError) => {
    console.error('Transaction error:', err);
  }, []);

  const handleSuccess = useCallback(
    (response: TransactionResponse) => {
      const formValues = getValues();
      const eventAddress = response.transactionReceipts?.[0].logs?.[0].address;
      const payload = {
        name: formValues.eventName,
        description: formValues.description,
        date: formValues.eventDate,
        location: formValues.location,
        contractAddress: eventAddress,
      };

      mutateAsync(payload).then(() => {
        setOpen(false);
        reset();
        refetchEventList();
      });
    },
    [getValues, mutateAsync],
  );

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild={true}>
        <Button>Create Event</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create New Event</DialogTitle>
        </DialogHeader>
        <form className="space-y-4">
          <div>
            <Label htmlFor="eventName">Event Name</Label>
            <Input
              id="eventName"
              {...register('eventName', { required: 'Event name is required' })}
            />
            {errors.eventName && (
              <p className="text-sm text-red-500">{errors.eventName.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="description">Description</Label>
            <Textarea
              id="description"
              {...register('description', {
                required: 'Description is required',
              })}
            />
            {errors.description && (
              <p className="text-sm text-red-500">
                {errors.description.message}
              </p>
            )}
          </div>
          <div>
            <Label htmlFor="location">Location</Label>
            <Input
              id="location"
              {...register('location', { required: 'Location is required' })}
            />
            {errors.location && (
              <p className="text-sm text-red-500">{errors.location.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="eventDate">Event Date</Label>
            <Input
              id="eventDate"
              {...register('eventDate', { required: 'Event date is required' })}
              placeholder="Aug 15-19, 2024"
            />
            {errors.eventDate && (
              <p className="text-sm text-red-500">{errors.eventDate.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="headerImage">Header Image URL</Label>
            <Input
              id="headerImage"
              type="url"
              {...register('headerImage', {
                required: 'Header image URL is required',
                pattern: {
                  value:
                    /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/,
                  message: 'Invalid URL format',
                },
              })}
            />
            {errors.headerImage && (
              <p className="text-sm text-red-500">
                {errors.headerImage.message}
              </p>
            )}
          </div>
          <Transaction
            contracts={contracts}
            chainId={BASE_SEPOLIA_CHAIN_ID}
            onError={handleError}
            onSuccess={handleSuccess}
          >
            <TransactionButton
              text="Create event"
              className="mt-0 mr-auto ml-auto max-w-full text-[white]"
            />
            <TransactionStatus>
              <TransactionStatusLabel />
              <TransactionStatusAction />
            </TransactionStatus>
          </Transaction>
        </form>
      </DialogContent>
    </Dialog>
  );
}
