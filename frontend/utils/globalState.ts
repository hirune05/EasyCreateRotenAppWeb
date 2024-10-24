
import { Item } from '@/types/type';
import { atom } from 'jotai';

const cartItemsAtom = atom<Item[]>([])
const storeItemsAtom = atom<Item[]>([])

export {cartItemsAtom,storeItemsAtom}