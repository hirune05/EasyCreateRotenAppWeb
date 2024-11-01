import type { Item, Order, OrderItem, StoreStaff } from '@/types/type'

export type MenuItemsProps = {
  orderItems: OrderItem[]
}

export type deliveryOrderProp = {
  deliveryOrder: Order
}

export type deliveryOrderProps = {
  storeStaff: StoreStaff | null
  id: number
  orderItems: OrderItem[]
}

export type MenuItemProps = {
  item: Item
  quantity: number
}
