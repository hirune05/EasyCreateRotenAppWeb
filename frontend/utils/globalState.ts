
import { Item } from '@/types/type';
import { atom } from 'jotai';

const itemAtom = atom<Item[]>([])

export default itemAtom