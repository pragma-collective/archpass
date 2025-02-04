import { NEXT_PUBLIC_API_BASE_URL } from '@/config';
import { createAuthenticatedClient } from '@/lib/client';
import { useMutation } from '@tanstack/react-query';

export function useCreateCheckoutMutation() {
  return useMutation({
    mutationFn: async (ticket: Record<string, any>) => {
      try {
        const api = createAuthenticatedClient();
        const { data, status } = await api.post(
          `${NEXT_PUBLIC_API_BASE_URL}/checkout.create`,
          ticket,
        );

        if (status === 404) {
          throw new Error('Ticket not found');
        }

        return data;
      } catch (err) {
        if (err?.response?.status === 404) {
          throw new Error('Ticket not found');
        }
        throw err;
      }
    },
  });
}
