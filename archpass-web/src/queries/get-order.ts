import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function useGetOrderQuery(orderId: string) {
  const { data, isLoading, isFetching, refetch, error } = useQuery({
    queryKey: ['orderItem', orderId],
    enabled: !!orderId,
    retry: false,
    refetchInterval: (data) => {
      if (data && !data.state.error) {
        return 10000;
      }

      return false;
    },
    queryFn: async () => {
      const api = createAuthenticatedClient();

      try {
        const { data: orderItem, status } = await api.get(`/order.get`, {
          params: { orderId },
        });

        if (status === 404) {
          throw new Error('Order not found');
        }

        return orderItem;
      } catch (err) {
        if (err?.response?.status === 404) {
          throw new Error('Order not found');
        }
        throw err;
      }
    },
  });

  return {
    order: data,
    pending: isLoading || isFetching,
    refetch,
    error,
    notFound: error?.message === 'Order not found',
  };
}
