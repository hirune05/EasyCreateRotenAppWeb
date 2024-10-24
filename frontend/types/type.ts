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
  storeId: number; // store_idをstoreIdに変更
  store: Store;
  pickedUpAt?: Date | null; // picked_up_atをpickedUpAtに変更
  status: number;
  storeStaffId: number; // store_staff_idをstoreStaffIdに変更
  storeStaff: StoreStaff;
  createdAt: Date; // created_atをcreatedAtに変更
  updatedAt: Date; // updated_atをupdatedAtに変更
  orderItems: OrderItem[]; // order_itemsをorderItemsに変更
};

type StoreStaff = {
  id: number;
  role: number;
  studentId: number; // student_idをstudentIdに変更
  student: Student;
  storeId: number; // store_idをstoreIdに変更
  store: Store;
  createdAt: Date; // created_atをcreatedAtに変更
  updatedAt: Date; // updated_atをupdatedAtに変更
  orders: Order[];
};

type Store = {
  id: number;
  name: string;
  imageUrl?: string | null; // image_urlをimageUrlに変更
  eventId: number; // event_idをeventIdに変更
  event: Event;
  createdAt: Date; // created_atをcreatedAtに変更
  updatedAt: Date; // updated_atをupdatedAtに変更
  staffs: StoreStaff[];
  items: Item[];
};

type Item = {
  id: number;
  storeId: number; // store_idをstoreIdに変更
  store: Store;
  name: string;
  description?: string | null;
  price: number;
  imageUrl?: string | null; // image_urlをimageUrlに変更
  createdAt: Date; // created_atをcreatedAtに変更
  updatedAt: Date; // updated_atをupdatedAtに変更
  orderItems: OrderItem[]; // order_itemsをorderItemsに変更
};

type OrderItem = {
  id: number;
  orderId: number; // order_idをorderIdに変更
  order: Order;
  itemId: number; // item_idをitemIdに変更
  item: Item;
  arranges?: string | null;
  quantity: number;
  createdAt: Date; // created_atをcreatedAtに変更
  updatedAt: Date; // updated_atをupdatedAtに変更
};

export type {LoginStudentResponse,AdminUser,EventInfo,Order}