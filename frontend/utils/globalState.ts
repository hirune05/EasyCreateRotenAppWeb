
import { Item, CartItem } from '@/types/type';
import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils'

const cartItemsAtom = atom<CartItem[]>([])
const storeItemsAtom = atom<Item[] | null>(null)

const tokenAtom = atomWithStorage<string | null>('token', null)
const studentIdAtom = atomWithStorage<number | null>('studentId', null)
const storeStaffIdAtom = atomWithStorage<number | null>('storeStaffId', null)
const nameAtom = atomWithStorage<string | null>('name', null)
const storeIdAtom = atomWithStorage<number>('storeId', 1)
const storeNameAtom = atomWithStorage<string | null>('storeName', null)

export {cartItemsAtom,storeItemsAtom,tokenAtom,studentIdAtom,nameAtom,storeIdAtom,storeNameAtom,storeStaffIdAtom}