
type LoginStudentResponse = {
  token: string;
  student_id: number;
  name: string;
  store_id: number;
  store_name: string;
};

type LoginStudentRequest = {
  student_id: number;
  password: string;
  event_id: number;
}

type AddOrderComplexItem = {
  item_id: number;
  quantity: number;
  arranges: string | undefined;
}

type AddOrderComplexRequest = {
  store_id: number;
  store_staff_id: number;
  items: AddOrderComplexItem[]
}

type StoreAndItems = {
  store_id: number;
  store_name: string;
  items: Item;
}

type getItemByStoreIdResponse = {
  StoreAndItems: StoreAndItems
}

type AdminUser = {
  id: number;
  username: string;
  password: string;
  email: string;
  createdAt: Date;
  updatedAt: Date;
};

type EventInfo = {
  Event: {
  id: number;
  name: string;
  year: number;
  startTime: Date;
  endTime: Date;
  createdAt: Date;
  updatedAt: Date;
  }
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
  pickedUpAt?: Date;
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
  imageUrl?: string;
  eventId: number;
  event: EventInfo;
  createdAt: Date;
  updatedAt: Date;
  staffs: StoreStaff[];
  items: Item[];
};

// 備考
// サンプルデータを作るときに型を変えたので本番では戻す
type Item = {
  id: number;
  storeId?: number;
  store?: Store;
  name: string;
  description?: string;
  price?: number;
  imageUrl?: string;
  createdAt?: Date; 
  updatedAt?: Date; 
  orderItems?: OrderItem[]; 
};

type OrderItem = {
  id: number;
  orderId: number;
  order: Order;
  itemId: number; 
  item: Item;
  arranges?: string;
  quantity: number;
  createdAt: Date;
  updatedAt: Date;
};

type OrderedItem = {
  name:string
  quantity: number;
  arranges?:string; 
}

export type {LoginStudentResponse,LoginStudentRequest,AddOrderComplexItem,getItemByStoreIdResponse,AdminUser,EventInfo,Order,Item,AddOrderComplexRequest,OrderItem,OrderedItem}
