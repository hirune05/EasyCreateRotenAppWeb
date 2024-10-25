
import { Item } from '@/types/type';
import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils'

// 備考
// お試しで作ったものなので本番では削除する
const sample:Item[] = [
    {
        id:123,
        name:"メニュー1",
        price:200
    },
    {
        id:124,
        name:"メニュー2",
        price:100
    }
]

const cartItemsAtom = atom<Item[]>([])
const storeItemsAtom = atom<Item[]>(sample)

export const tokenAtom = atomWithStorage<string | null>('token', null)
export const studentIdAtom = atomWithStorage<number | null>('studentId', null)
export const nameAtom = atomWithStorage<string | null>('name', null)
export const storeIdAtom = atomWithStorage<number | null>('storeId', null)
export const storeNameAtom = atomWithStorage<string | null>('storeName', null)

export {cartItemsAtom,storeItemsAtom}