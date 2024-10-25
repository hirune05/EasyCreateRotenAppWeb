
import { Item } from '@/types/type';
import { atom } from 'jotai';

// 備考
// お試しで作ったものなので本番では削除する
const sample:Item[] = [
    {
        id:123,
        name:"メニュー1"
    },
    {
        id:124,
        name:"メニュー2"
    }
]

const cartItemsAtom = atom<Item[]>([])
const storeItemsAtom = atom<Item[]>(sample)

export {cartItemsAtom,storeItemsAtom}