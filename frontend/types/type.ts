
type LoginStudentResponse = {
  token: string;
  studentId: number;
  name: string;
  storeId: number;
  storeName: string;
  storeStaffId: number;
};

type LoginStudentRequest = {
  studentId: number;
  password: string;
  eventId: number;
}

type AddOrderComplexItem = {
  itemId: number;
  quantity: number;
  arranges?: string;
}

type AddOrderComplexRequest = {
  storeId: number;
  storeStaffId: number;
  items: AddOrderComplexItem[]
}

type getItemByStoreIdResponse = {
  storeId: number;
  storeName: string;
  Items: Item[];
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
  id: number;
  name: string;
  year: number;
  startTime: Date;
  endTime: Date;
  createdAt: Date;
  updatedAt: Date;
};

// 本番用の型
type Item = {
  id: number;
  storeId?: number;
  Store?: Store;
  name: string;
  description: string | null;
  price: number;
  status: number;
  imageUrl: string | null;
  createdAt: Date; 
  updatedAt: Date; 
  OrderItems: OrderItem[]; 
}

type Order = {
  id: number;
  storeId: number;
  Store: Store | null;
  pickedUpAt: Date | null;
  status: number;
  storeStaffId: number;
  StoreStaff: StoreStaff | null;
  createdAt: Date;
  updatedAt: Date;
  OrderItems: OrderItem[];
};

type OrderItem = {
  id: number;
  orderId: number;
  Order: Order | null;
  itemId: number;
  /* item: Item | null; */
  Item: Item;
  arranges?: string;
  quantity: number;
  createdAt: Date;
  updatedAt: Date;
};

type Store = {
  id: number;
  name: string;
  imageUrl: string | null;
  eventId: number;
  Event: EventInfo | null;
  createdAt: Date;
  updatedAt: Date;
  Staffs: StoreStaff[];
  Items: Item[];
};

type StoreStaff = {
  id: number;
  role: number;
  studentId: number;
  Student: Student | null;
  storeId: number;
  Store: Store | null;
  createdAt: Date;
  updatedAt: Date;
  Orders: Order[];
};

type Student = {
  id: number;
  name: string;
  password: string;
  createdAt: Date;
};

type OrderedItem = {
  name:string
  quantity: number;
  arranges?:string; 
}

type CartItem = {
  id: number;
  name: string;
  price?: number;
  arranges?: string;
}

type ReportItems = {
  storeId?:number
  storeStaffId?:number
  description:string
  title:string

}

export type {LoginStudentResponse,LoginStudentRequest,AddOrderComplexItem,getItemByStoreIdResponse,AdminUser,EventInfo,Order,Item,AddOrderComplexRequest,OrderItem,OrderedItem,CartItem,ReportItems,StoreStaff}
