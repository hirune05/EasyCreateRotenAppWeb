
import { Item } from '@/types/type';
import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils'

const cartItemsAtom = atom<Item[]>([])
const storeItemsAtom = atom<Item[]>([])

export const tokenAtom = atomWithStorage<string | null>('token', null)
export const studentIdAtom = atomWithStorage<number | null>('studentId', null)
export const nameAtom = atomWithStorage<string | null>('name', null)
export const storeIdAtom = atomWithStorage<number | null>('storeId', null)
export const storeNameAtom = atomWithStorage<string | null>('storeName', null)

export {cartItemsAtom,storeItemsAtom}