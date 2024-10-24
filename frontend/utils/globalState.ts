
import { Item } from '@/types/type';
import { atom } from 'jotai';

const cartItemsAtom = atom<Item[]>([])

export {cartItemsAtom}