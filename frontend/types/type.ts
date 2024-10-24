
type LoginStudentResponse = {
  token: string;
  studentId: number;
  name: string;
  storeId: number;
  storeName: string;
};

type AdminUser = {
  id: number;
  username: string;
  password: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
};

type EventInfo = {
  id: number;
  name: string;
  year: number;
  startTime: Date;
  endTime: Date;
  createdAt: Date;
  updatedAt: Date;
};

type Student = {
  id: number;
  name: string;
  password: string;
  createdAt: Date;
};

type Order = {
  id: number;
  storeId: number;
  store: Store;
  pickedUpAt?: Date | null;
  status: number;
  storeStaffId: number;
  storeStaff: StoreStaff;
  createdAt: Date;
  updatedAt: Date;
  orderItems: OrderItem[];
};

type StoreStaff = {
  id: number;
  role: number;
  studentId: number;
  student: Student;
  storeId: number;
  store: Store;
  createdAt: Date;
  updatedAt: Date;
  orders: Order[];
};

type Store = {
  id: number;
  name: string;
  imageUrl?: string | null;
  eventId: number;
  event: Event;
  createdAt: Date;
  updatedAt: Date;
  staffs: StoreStaff[];
  items: Item[];
};

type Item = {
  id: number;
  storeId: number;
  store: Store;
  name: string;
  description?: string | null;
  price: number;
  imageUrl?: string | null;
  createdAt: Date; 
  updatedAt: Date; 
  orderItems: OrderItem[]; 
};

type OrderItem = {
  id: number;
  orderId: number;
  order: Order;
  itemId: number; 
  item: Item;
  arranges?: string | null;
  quantity: number;
  createdAt: Date;
  updatedAt: Date;
};

export type {LoginStudentResponse,AdminUser,EventInfo,Order,Item}
