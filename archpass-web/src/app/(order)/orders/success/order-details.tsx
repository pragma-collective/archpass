import type { OrderData } from './order-data-provider';

export default function OrderDetails({ orderData }: { orderData: OrderData }) {
  return (
    <div className="bg-gray-50 rounded-xl p-6">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">
        Order Details
      </h2>
      <dl className="grid grid-cols-1 gap-x-4 gap-y-4 sm:grid-cols-2">
        <div className="sm:col-span-1">
          <dt className="text-sm font-medium text-gray-500">Order ID</dt>
          <dd className="mt-1 text-sm text-gray-900">{orderData.id}</dd>
        </div>
        <div className="sm:col-span-1">
          <dt className="text-sm font-medium text-gray-500">Event Name</dt>
          <dd className="mt-1 text-sm text-gray-900">{orderData.eventName}</dd>
        </div>
        <div className="sm:col-span-1">
          <dt className="text-sm font-medium text-gray-500">Reference</dt>
          <dd className="mt-1 text-sm text-gray-900">
            {orderData.paymentReference}
          </dd>
        </div>
        <div className="sm:col-span-1">
          <dt className="text-sm font-medium text-gray-500">Ticket Quantity</dt>
          <dd className="mt-1 text-sm text-gray-900">{orderData.quantity}</dd>
        </div>
        <div className="sm:col-span-1">
          <dt className="text-sm font-medium text-gray-500">Total Amount</dt>
          <dd className="mt-1 text-sm text-gray-900">{orderData.price} USDC</dd>
        </div>
      </dl>
    </div>
  );
}
