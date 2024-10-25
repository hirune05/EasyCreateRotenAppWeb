import { Item, Order, OrderItem } from '@/types/type'

export type MenuItemsProps = {
  orderItems: OrderItem[]
}

export type deliveryOrderProp = {
  deliveryOrder: Order
}

export type deliveryOrderProps = {
  storeStaffId: number
  id: number
  orderItems: OrderItem[]
}

export type MenuItemProps = {
  item: Item
  quantity: number
}
